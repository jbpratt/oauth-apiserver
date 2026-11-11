package v1alpha1

import (
	authenticationv1alpha1 "github.com/openshift/api/authentication/v1alpha1"
	"github.com/openshift/oauth-apiserver/pkg/externaloidc/apis/authentication"
	"k8s.io/apimachinery/pkg/conversion"
)

// Convert_v1alpha1_AuthenticationConfiguration_To_authentication_AuthenticationConfiguration drops
// TrustedCAFilePath: the internal type has no peer for it.
func Convert_v1alpha1_AuthenticationConfiguration_To_authentication_AuthenticationConfiguration(in *authenticationv1alpha1.AuthenticationConfiguration, out *authentication.AuthenticationConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_AuthenticationConfiguration_To_authentication_AuthenticationConfiguration(in, out, s)
}
