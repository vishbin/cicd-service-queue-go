#!groovy



def project = new Project();

properties([
  /* Only keep the 10 most recent builds. */
  [$class: 'BuildDiscarderProperty', strategy: [$class: 'LogRotator', numToKeepStr: '10']],
])

try {

  stage 'Checkout Code'
  node('bogie-slave') {
    deleteDir()
    checkoutCode {}
    project.loadProperties()
    stash 'sources'
  }

  stage 'Maven Build'
  node('bogie-slave') {
    deleteDir()
    unstash 'sources'
    mavenBuild {}
    dockerTasks {
      name    = project['name']
      version = project['version']
    }
    stash includes: '**/*.jar,**/*.war', name: 'binaries'
  }

  stage name: 'Sonar Analysis', concurrency: 1
  node('bogie-slave') {
    deleteDir()
    unstash 'sources'
    sonarAnalysis {}
  }



  stage name: 'Build Dev Environment', concurrency: 1
  bogieInfra {
    service     = project['name']
    environment = 'dev'
  }

  stage name: 'Dev Deploy', concurrency: 1
  node('bogie-slave') {
    bogieDeploy {
      serviceName    = project['name']
      serviceVersion = project['version']
      deployEnv      = 'dev'
    }
  }

  

} catch (BotTriggerException ex) {
  echo "Skipping this build - bots are excluded from triggering the pipeline"
  return // exit the pipeline without failing the build
} catch (Exception ex) {
  sendFailureNotifications {
    emailCulprits     = true // emails the suspected commit culprit(s)
    // notifyHipChatRoom = '<room_name>' // send notification to the team's hipchat room
  }

  throw ex // re-throw so that the exception bubbles up to the console
}