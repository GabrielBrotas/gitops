
## Process
- Execute tests
- Linter
- Verify code quality
- Verify security
- Generate artifacts to the deployment process
- Identify the next software version (1.2, 2.0, 1.2.2,...)
- Generate tags and releases

## Some tools
- Jenkins
- Github Actions
- Circle CI
- AWS Code Build
- Azure DevOps
- Google Cloud Build
- GitLab Pipelines / CI

## Status Check
Guarantee that a pull request will not be merged to the main branch without having gone through the CI process or Code Review.

Activate status check on github:
Settings -> Branches -> Branch protection rules -> Add rule
- [x] Require status check to pass before merging
    - select the job that needs to pass to accept the merge
- [x] Include administrators
- [x] Restrict who can push to matching branches
    - No one can commit direct to this branch, you have to create a pull request in order to commit some change

Add Rule
- add the same rules to the develop branch

# update readme