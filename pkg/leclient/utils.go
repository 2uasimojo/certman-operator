/*
Copyright 2019 Red Hat, Inc.

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

package leclient

import (
	"context"
	"crypto/x509/pkix"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func GetSecret(kubeClient client.Client, secretName, namespace string) (*corev1.Secret, error) {

	s := &corev1.Secret{}

	err := kubeClient.Get(context.TODO(), types.NamespacedName{Name: secretName, Namespace: namespace}, s)

	if err != nil {
		return nil, err
	}

	return s, nil
}

// IsCertificateIssuerLE takes an issuer name on a certificate and determines if it's a Let's Encrypt CA
func IsCertificateIssuerLE(issuer pkix.Name) bool {
	if len(issuer.Organization) > 0 {
		for _, o := range issuer.Organization {
			if o == "Let's Encrypt" {
				return true
			}
		}
	}

	if issuer.CommonName == "Fake LE Intermediate X1" {
		return true
	}

	return false
}
