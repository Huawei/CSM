/*
Copyright (c) Huawei Technologies Co., Ltd. 2024-2024. All rights reserved.

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

// Package resource is used to obtain core resources in Kubernetes.
package resource

import (
	"context"

	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// SecretOps can get information of secret
type SecretOps interface {
	// GetSecret gets the secrets object given its name and namespace
	GetSecret(name string, namespace string) (*coreV1.Secret, error)
}

// GetSecret gets the secrets object given its name and namespace
func (c *Client) GetSecret(name string, namespace string) (*coreV1.Secret, error) {
	if err := c.initClient(); err != nil {
		return nil, err
	}

	return c.kubernetes.CoreV1().Secrets(namespace).Get(context.TODO(), name, metaV1.GetOptions{})
}
