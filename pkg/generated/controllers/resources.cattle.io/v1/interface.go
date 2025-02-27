/*
Copyright 2023 Rancher Labs, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by main. DO NOT EDIT.

package v1

import (
	v1 "github.com/rancher/backup-restore-operator/pkg/apis/resources.cattle.io/v1"
	"github.com/rancher/lasso/pkg/controller"
	"github.com/rancher/wrangler/pkg/schemes"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func init() {
	schemes.Register(v1.AddToScheme)
}

type Interface interface {
	Backup() BackupController
	ResourceSet() ResourceSetController
	Restore() RestoreController
}

func New(controllerFactory controller.SharedControllerFactory) Interface {
	return &version{
		controllerFactory: controllerFactory,
	}
}

type version struct {
	controllerFactory controller.SharedControllerFactory
}

func (c *version) Backup() BackupController {
	return NewBackupController(schema.GroupVersionKind{Group: "resources.cattle.io", Version: "v1", Kind: "Backup"}, "backups", false, c.controllerFactory)
}
func (c *version) ResourceSet() ResourceSetController {
	return NewResourceSetController(schema.GroupVersionKind{Group: "resources.cattle.io", Version: "v1", Kind: "ResourceSet"}, "resourcesets", false, c.controllerFactory)
}
func (c *version) Restore() RestoreController {
	return NewRestoreController(schema.GroupVersionKind{Group: "resources.cattle.io", Version: "v1", Kind: "Restore"}, "restores", false, c.controllerFactory)
}
