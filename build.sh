#!/bin/sh
# For CI pipline, temp build
VERSION_TAG=`date +%m``date +%d`
docker build -t dockerleslie/fortunecow:$VERSION_TAG
docker tag dockerleslie/fortunecow:$VERSION_TAG dockerleslie/fortunecow:latest
docker login
docker push dockerleslie/fortunecow:$VERSION_TAG
docker push dockerleslie/fortunecow:latest
