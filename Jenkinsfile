pipeline {
    agent any
    tools {
        go 'go1.22.2' // Ensure this Go version is installed and configured in Jenkins
    }
    environment {
        GO111MODULE = 'on'
    }

    stages {
        stage('Clone Repository') {
            steps {
                git 'https://github.com/mayankgongal5/go-pipeline.git'
            }
        }
        stage('Build') {
            steps {
                echo 'Building the app'
                sh 'go build -v ./...'
            }
        }
        stage('Test') {
            steps {
                echo 'Running tests'
                sh 'go test -v ./...'
            }
        }
    }
}
