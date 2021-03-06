const { awscdk } = require('projen');
const { NpmAccess } = require('projen/lib/javascript');
const project = new awscdk.AwsCdkConstructLibrary({
  author: 'maafk',
  authorAddress: 'maafk@users.noreply.github.com',
  cdkVersion: '2.25.0',
  defaultReleaseBranch: 'main',
  name: '@professionalaf/enterprise-utils',
  autoApproveOptions: {
    allowedUsernames: ['maafk'],
  },
  autoApproveUpgrades: true,
  npmAccess: NpmAccess.PUBLIC,
  repositoryUrl: 'https://github.com/professionalaf/enterprise-utils.git',
  publishToPypi: {
    distName: 'enterprise-utils',
    module: 'enterprise_utils',
    homepage: 'https://github.com/professionalaf/enterprise-utils',
  },
  publishToNuget: {
    packageId: 'professionalaf.EnterpriseUtils',
    dotNetNamespace: 'professionalaf.EnterpriseUtils',
  },
  publishToMaven: {
    mavenGroupId: 'io.github.professionalaf',
    javaPackage: 'io.github.professionalaf.EnterpriseUtils',
    mavenArtifactId: 'EnterpriseUtils',
    mavenEndpoint: 'https://s01.oss.sonatype.org',
  },
  publishToGo: {
    moduleName: 'github.com/professionalaf/enterprise-utils',
  },
});
project.synth();