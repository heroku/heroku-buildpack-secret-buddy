How to Test and Debug
=====================

**First of all, if you didn't already do it, please read [this](https://devcenter.heroku.com/articles/buildpack-api)**

### Description
Operationally, the secretbuddy buildpack consists of the buildpack per-se (the shell scripts located in `/bin`) and the worker (`secret-buddy-buildpack`) go program.

### Shell scripts
We only need to add functionality to the `compile` script. More info [here](https://devcenter.heroku.com/articles/buildpack-api#bin-compile). In this script, we validate if the `SECRETBUDDY_ENV` is not empty, and we create the structure needed for the `.profile.d` directory and the worker, which will be used during the app startup. More info [here](https://devcenter.heroku.com/articles/buildpack-api#profile-d-scripts).

### Worker
This go program has all the functionality to populate the secrets envars after reading both `SECRETBUDDY_ENV` and `HEROKU_SECRETS_CONFIG`. The output is the list of `export` commands for each one of the env vars. More info to understand `HEROKU_SECRETS_CONFIG` can be found [here](https://salesforce.quip.com/lZ0BAq4Factj).

## How to Test
Any change of the buildpack must be published. Our buildpacks are published [here](https://addons-next.heroku.com/buildpacks).

### Changes in the scripts

#### How to Run locally
You can run the scripts with the local terminal, in your buildpack source directory, by executing:
```
$./bin/compile . <any param> <path to the ENV_DIR>
```

The `<path to the ENV_DIR>` is a directory named `ENV_DIR` where you must re-create the envvars by creating one file per each env var. For example, if the app receives the env var `SHOGUN_USER`, you have to make a file named `SHOGUN_USER` whose content is the content of that env var.
Scripts run when the application is built and deployed. So any log printout will appear in the build log of the Dashboard. [Example](https://dashboard.heroku.com/apps/herokai-addons-opex-staging/activity/builds/256571ac-70fe-4648-938e-f326160f28b0)
So, add the buildpack to an app, and deploy this app. Check the `Activity > Build log`.

## Changes in the worker
Any change in the worker must be built, pushed, and published. So, make any necessary changes, and locally run `make build` and the `secret-buddy-buildpack` binary will be built for Linux. Push this change, and publish the new buildpack version as explained above. Then re-deploy the application that has this buildpack.

## check-secretbuddy script
This Python script will be copied into the `bin` directory of the application, so that it can be run with the command `heroku run -a <APP NAME> -- bin/check-secretbuddy`. The copy process is done in the `compile` script, along with the change of attributes to `755` (executable). For testing, the Python script can be run it locally, and it performs a series of checks, like the listing of all secretbuddy vars (from `SECRETBUDDY_ENV`, the verification that all the secretbuddy vars are OS env vars, etc).