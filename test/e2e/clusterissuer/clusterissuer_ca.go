/*
Copyright 2018 The Jetstack cert-manager contributors.

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

package clusterissuer

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/jetstack/cert-manager/pkg/apis/certmanager/v1alpha1"
	"github.com/jetstack/cert-manager/test/e2e/framework"
	"github.com/jetstack/cert-manager/test/util"
)

const clusterResourceNamespace = "cert-manager"

var _ = framework.CertManagerDescribe("CA ClusterIssuer", func() {
	f := framework.NewDefaultFramework("create-ca-clusterissuer")

	issuerName := "test-ca-clusterissuer"
	secretName := "ca-clusterissuer-signing-keypair"

	BeforeEach(func() {
		By("Creating a signing keypair fixture")
		_, err := f.KubeClientSet.CoreV1().Secrets(clusterResourceNamespace).Create(util.NewSigningKeypairSecret(secretName))
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		By("Cleaning up")
		f.KubeClientSet.CoreV1().Secrets(clusterResourceNamespace).Delete(secretName, nil)
	})

	It("should validate a signing keypair", func() {
		By("Creating an Issuer")
		clusterIssuer := util.NewCertManagerCAClusterIssuer(issuerName, secretName)
		_, err := f.CertManagerClientSet.CertmanagerV1alpha1().ClusterIssuers().Create(clusterIssuer)
		Expect(err).NotTo(HaveOccurred())
		defer f.CertManagerClientSet.CertmanagerV1alpha1().ClusterIssuers().Delete(clusterIssuer.Name, nil)
		By("Waiting for Issuer to become Ready")
		err = util.WaitForClusterIssuerCondition(f.CertManagerClientSet.CertmanagerV1alpha1().ClusterIssuers(),
			issuerName,
			v1alpha1.IssuerCondition{
				Type:   v1alpha1.IssuerConditionReady,
				Status: v1alpha1.ConditionTrue,
			})
		Expect(err).NotTo(HaveOccurred())
	})
})
