#!/bin/bash

set -e

version_from_conf=`awk '$1=="versioncode" {print $3}' ./conf/app.conf`

if [ $# == 0 ] && [ -z $version_from_conf ]; then
    echo "baby,we need a version code"
    exit 1
fi

version=$version_from_conf
if [ $# == 1 ]; then
    version=$1
fi

echo $version

default_runmode="test"
runmode=`awk '$1=="runmode" {print $3}' ./conf/app.conf`

if [ $default_runmode != $runmode ]
then
    echo "$runmode is err,you should in $default_runmode"
	exit 1
fi

echo "redeploy publicservice"

cd GoWorkTest/src/zhangmai_micro/publicservice/
if docker build -t publicservicetest:$version .
then
    echo "stop and rm old container,start new one..."
    docker stop publicservicetest
    docker rm publicservicetest
    docker run --restart=always --name  publicservicetest -v /opt/lvdoutest/zhangmailog/publicservicelog:/opt/zhangmaiLogs -v /data:/data -d -p 9012:9012 -p 9013:9013 publicservicetest:$version
    docker ps
    docker rmi -f  `docker images | grep '<none>' | awk '{print $3}'`
fi


