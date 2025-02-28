const core = require('@actions/core');
const github = require('@actions/github');
const exec = require('@actions/exec');

function run() {
    core.notice('Hello from my custom JS action');
    const bucket = core.getInput('bucket', {required: true});
    const bucketRegion = core.getInput('bucket-region', {required: true});
    const disFolder = core.getInput('dist-folder', {required: true});

    // log inputs
    core.notice(`bucket: ${bucket}`);
    core.notice(`bucketRegion: ${bucketRegion}`);
    core.notice(`disFolder: ${disFolder}`);

    exec.exec('ls', ['-la']);

    // test print dummy from env
    core.info(`AWS_ACCESS_KEY_ID: ${process.env.AWS_ACCESS_KEY_ID}`);

}

run();