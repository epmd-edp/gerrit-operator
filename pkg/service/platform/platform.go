package platform

import (
	"gerrit-operator/pkg/apis/edp/v1alpha1"
	"gerrit-operator/pkg/service/helpers"
	"gerrit-operator/pkg/service/platform/openshift"
	appsV1Api "github.com/openshift/api/apps/v1"
	coreV1Api "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/clientcmd"
)

// PlatformService defines common behaviour of the services for the supported platforms
type PlatformService interface {
	CreateExternalEndpoint(gerrit *v1alpha1.Gerrit) error
	CreateSecurityContext(gerrit *v1alpha1.Gerrit, sa *coreV1Api.ServiceAccount) error
	CreateService(gerrit *v1alpha1.Gerrit) error
	CreateSecret(gerrit *v1alpha1.Gerrit, name string, data map[string][]byte) error
	CreateVolume(gerrit *v1alpha1.Gerrit) error
	CreateServiceAccount(gerrit *v1alpha1.Gerrit) (*coreV1Api.ServiceAccount, error)
	GetSecret(namespace string, name string) (map[string][]byte, error)
	CreateDeployConf(gerrit *v1alpha1.Gerrit) error
	GetDeploymentConfig(instance v1alpha1.Gerrit) (*appsV1Api.DeploymentConfig, error)
}

// NewService creates a new instance of the platform.Service type using scheme parameter provided
func NewService(scheme *runtime.Scheme) PlatformService {
	config := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		clientcmd.NewDefaultClientConfigLoadingRules(),
		&clientcmd.ConfigOverrides{},
	)

	restConfig, err := config.ClientConfig()
	if err != nil {
		helpers.LogErrorAndReturn(err)
		return nil
	}

	platform := &openshift.OpenshiftService{}
	if err = platform.Init(restConfig, scheme); err != nil {
		helpers.LogErrorAndReturn(err)
		return nil
	}
	return platform
}
