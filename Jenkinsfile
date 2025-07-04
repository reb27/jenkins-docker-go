pipeline {
    agent any

    stages {
        stage('Build') {
            steps {
                script {
                    docker.build("go-build", "-f Dockerfile.build .").inside {
                        sh './converter'
                    }
                }
            }
        }
        stage('Test') {
            steps {
                script {
                    docker.build("go-test", "-f Dockerfile.test .").inside {
                        sh 'go test -v'
                    }
                }
            }
        }
    }
}
