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

package certificaterequest

import (
	"time"

	"github.com/go-logr/logr"

	certmanv1alpha1 "github.com/openshift/certman-operator/pkg/apis/certman/v1alpha1"
)

func (r *ReconcileCertificateRequest) ShouldRenew(reqLogger logr.Logger, cr *certmanv1alpha1.CertificateRequest) (bool, error) {

	renewBeforeDays := cr.Spec.RenewBeforeDays

	if renewBeforeDays <= 0 {
		renewBeforeDays = RenewCertificateBeforeDays
	}

	certificate, err := GetCertificate(r.client, cr)
	if err != nil || certificate == nil {
		log.Error(err, "There was problem loading existing certificate")
		return false, err
	}

	if certificate != nil {

		notAfter := certificate.NotAfter
		currentTime := time.Now().In(time.UTC)
		timeDiff := notAfter.Sub(currentTime)
		daysCertificateValidFor := int(timeDiff.Hours() / 24)
		shouldRenew := daysCertificateValidFor <= renewBeforeDays

		reqLogger.Info("Checking if certificate should be renewed", "RenewCertificateBeforeDays", renewBeforeDays, "notAfter", notAfter.String(), "daysCertificateValidFor", daysCertificateValidFor, "shouldRenew", shouldRenew)

		return shouldRenew, nil
	}

	return false, nil
}
