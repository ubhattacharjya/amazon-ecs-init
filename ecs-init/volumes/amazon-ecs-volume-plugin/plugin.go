// Copyright 2019 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//	http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package main

import (
	"os/user"
	"strconv"

	"github.com/aws/amazon-ecs-init/ecs-init/volumes"
	"github.com/docker/go-plugins-helpers/volume"
)

func main() {
	plugin := volumes.NewAmazonECSVolumePlugin()
	handler := volume.NewHandler(plugin)
	rootUser, _ := user.Lookup("root")
	gid, _ := strconv.Atoi(rootUser.Gid)
	handler.ServeUnix("amazon-ecs-volume-plugin", gid)
}
