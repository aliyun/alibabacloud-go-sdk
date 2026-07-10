// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApplicationFederatedCredentialRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationFederatedCredentialId(v string) *UpdateApplicationFederatedCredentialRequest
	GetApplicationFederatedCredentialId() *string
	SetApplicationId(v string) *UpdateApplicationFederatedCredentialRequest
	GetApplicationId() *string
	SetAttributeMappings(v []*UpdateApplicationFederatedCredentialRequestAttributeMappings) *UpdateApplicationFederatedCredentialRequest
	GetAttributeMappings() []*UpdateApplicationFederatedCredentialRequestAttributeMappings
	SetInstanceId(v string) *UpdateApplicationFederatedCredentialRequest
	GetInstanceId() *string
	SetOidcVerificationConfig(v *UpdateApplicationFederatedCredentialRequestOidcVerificationConfig) *UpdateApplicationFederatedCredentialRequest
	GetOidcVerificationConfig() *UpdateApplicationFederatedCredentialRequestOidcVerificationConfig
	SetPkcs7VerificationConfig(v *UpdateApplicationFederatedCredentialRequestPkcs7VerificationConfig) *UpdateApplicationFederatedCredentialRequest
	GetPkcs7VerificationConfig() *UpdateApplicationFederatedCredentialRequestPkcs7VerificationConfig
	SetVerificationCondition(v string) *UpdateApplicationFederatedCredentialRequest
	GetVerificationCondition() *string
}

type UpdateApplicationFederatedCredentialRequest struct {
	// The application federated credential ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// afc_aaaaa1111
	ApplicationFederatedCredentialId *string `json:"ApplicationFederatedCredentialId,omitempty" xml:"ApplicationFederatedCredentialId,omitempty"`
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The attribute mappings.
	AttributeMappings []*UpdateApplicationFederatedCredentialRequestAttributeMappings `json:"AttributeMappings,omitempty" xml:"AttributeMappings,omitempty" type:"Repeated"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The OIDC structured configuration (structured mode + oidc type).
	OidcVerificationConfig *UpdateApplicationFederatedCredentialRequestOidcVerificationConfig `json:"OidcVerificationConfig,omitempty" xml:"OidcVerificationConfig,omitempty" type:"Struct"`
	// The PKCS#7 structured configuration (structured mode + pkcs7 type).
	Pkcs7VerificationConfig *UpdateApplicationFederatedCredentialRequestPkcs7VerificationConfig `json:"Pkcs7VerificationConfig,omitempty" xml:"Pkcs7VerificationConfig,omitempty" type:"Struct"`
	// The verification condition.
	//
	// example:
	//
	// IsNullOrEmpty("")
	VerificationCondition *string `json:"VerificationCondition,omitempty" xml:"VerificationCondition,omitempty"`
}

func (s UpdateApplicationFederatedCredentialRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationFederatedCredentialRequest) GoString() string {
	return s.String()
}

func (s *UpdateApplicationFederatedCredentialRequest) GetApplicationFederatedCredentialId() *string {
	return s.ApplicationFederatedCredentialId
}

func (s *UpdateApplicationFederatedCredentialRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *UpdateApplicationFederatedCredentialRequest) GetAttributeMappings() []*UpdateApplicationFederatedCredentialRequestAttributeMappings {
	return s.AttributeMappings
}

func (s *UpdateApplicationFederatedCredentialRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateApplicationFederatedCredentialRequest) GetOidcVerificationConfig() *UpdateApplicationFederatedCredentialRequestOidcVerificationConfig {
	return s.OidcVerificationConfig
}

func (s *UpdateApplicationFederatedCredentialRequest) GetPkcs7VerificationConfig() *UpdateApplicationFederatedCredentialRequestPkcs7VerificationConfig {
	return s.Pkcs7VerificationConfig
}

func (s *UpdateApplicationFederatedCredentialRequest) GetVerificationCondition() *string {
	return s.VerificationCondition
}

func (s *UpdateApplicationFederatedCredentialRequest) SetApplicationFederatedCredentialId(v string) *UpdateApplicationFederatedCredentialRequest {
	s.ApplicationFederatedCredentialId = &v
	return s
}

func (s *UpdateApplicationFederatedCredentialRequest) SetApplicationId(v string) *UpdateApplicationFederatedCredentialRequest {
	s.ApplicationId = &v
	return s
}

func (s *UpdateApplicationFederatedCredentialRequest) SetAttributeMappings(v []*UpdateApplicationFederatedCredentialRequestAttributeMappings) *UpdateApplicationFederatedCredentialRequest {
	s.AttributeMappings = v
	return s
}

func (s *UpdateApplicationFederatedCredentialRequest) SetInstanceId(v string) *UpdateApplicationFederatedCredentialRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateApplicationFederatedCredentialRequest) SetOidcVerificationConfig(v *UpdateApplicationFederatedCredentialRequestOidcVerificationConfig) *UpdateApplicationFederatedCredentialRequest {
	s.OidcVerificationConfig = v
	return s
}

func (s *UpdateApplicationFederatedCredentialRequest) SetPkcs7VerificationConfig(v *UpdateApplicationFederatedCredentialRequestPkcs7VerificationConfig) *UpdateApplicationFederatedCredentialRequest {
	s.Pkcs7VerificationConfig = v
	return s
}

func (s *UpdateApplicationFederatedCredentialRequest) SetVerificationCondition(v string) *UpdateApplicationFederatedCredentialRequest {
	s.VerificationCondition = &v
	return s
}

func (s *UpdateApplicationFederatedCredentialRequest) Validate() error {
	if s.AttributeMappings != nil {
		for _, item := range s.AttributeMappings {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.OidcVerificationConfig != nil {
		if err := s.OidcVerificationConfig.Validate(); err != nil {
			return err
		}
	}
	if s.Pkcs7VerificationConfig != nil {
		if err := s.Pkcs7VerificationConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateApplicationFederatedCredentialRequestAttributeMappings struct {
	// The source value expression.
	//
	// example:
	//
	// Append(client.applicationFederatedCredentialId, ":", cert.subject.CN, ":", cert.serialNumber)
	SourceValueExpression *string `json:"SourceValueExpression,omitempty" xml:"SourceValueExpression,omitempty"`
	// The target field.
	//
	// example:
	//
	// client.activeSubjectUrn
	TargetField *string `json:"TargetField,omitempty" xml:"TargetField,omitempty"`
}

func (s UpdateApplicationFederatedCredentialRequestAttributeMappings) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationFederatedCredentialRequestAttributeMappings) GoString() string {
	return s.String()
}

func (s *UpdateApplicationFederatedCredentialRequestAttributeMappings) GetSourceValueExpression() *string {
	return s.SourceValueExpression
}

func (s *UpdateApplicationFederatedCredentialRequestAttributeMappings) GetTargetField() *string {
	return s.TargetField
}

func (s *UpdateApplicationFederatedCredentialRequestAttributeMappings) SetSourceValueExpression(v string) *UpdateApplicationFederatedCredentialRequestAttributeMappings {
	s.SourceValueExpression = &v
	return s
}

func (s *UpdateApplicationFederatedCredentialRequestAttributeMappings) SetTargetField(v string) *UpdateApplicationFederatedCredentialRequestAttributeMappings {
	s.TargetField = &v
	return s
}

func (s *UpdateApplicationFederatedCredentialRequestAttributeMappings) Validate() error {
	return dara.Validate(s)
}

type UpdateApplicationFederatedCredentialRequestOidcVerificationConfig struct {
	// The Azure VM scenario configuration.
	AzureVmConfig *UpdateApplicationFederatedCredentialRequestOidcVerificationConfigAzureVmConfig `json:"AzureVmConfig,omitempty" xml:"AzureVmConfig,omitempty" type:"Struct"`
	// The GCP VM scenario configuration.
	GcpVmConfig   *UpdateApplicationFederatedCredentialRequestOidcVerificationConfigGcpVmConfig   `json:"GcpVmConfig,omitempty" xml:"GcpVmConfig,omitempty" type:"Struct"`
	GenericConfig *UpdateApplicationFederatedCredentialRequestOidcVerificationConfigGenericConfig `json:"GenericConfig,omitempty" xml:"GenericConfig,omitempty" type:"Struct"`
	// The Kubernetes scenario configuration.
	KubernetesConfig *UpdateApplicationFederatedCredentialRequestOidcVerificationConfigKubernetesConfig `json:"KubernetesConfig,omitempty" xml:"KubernetesConfig,omitempty" type:"Struct"`
	// The OIDC scenario profile. Valid values: generic, kubernetes, gcp_vm, and azure_vm.
	//
	// example:
	//
	// kubernetes
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
}

func (s UpdateApplicationFederatedCredentialRequestOidcVerificationConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationFederatedCredentialRequestOidcVerificationConfig) GoString() string {
	return s.String()
}

func (s *UpdateApplicationFederatedCredentialRequestOidcVerificationConfig) GetAzureVmConfig() *UpdateApplicationFederatedCredentialRequestOidcVerificationConfigAzureVmConfig {
	return s.AzureVmConfig
}

func (s *UpdateApplicationFederatedCredentialRequestOidcVerificationConfig) GetGcpVmConfig() *UpdateApplicationFederatedCredentialRequestOidcVerificationConfigGcpVmConfig {
	return s.GcpVmConfig
}

func (s *UpdateApplicationFederatedCredentialRequestOidcVerificationConfig) GetGenericConfig() *UpdateApplicationFederatedCredentialRequestOidcVerificationConfigGenericConfig {
	return s.GenericConfig
}

func (s *UpdateApplicationFederatedCredentialRequestOidcVerificationConfig) GetKubernetesConfig() *UpdateApplicationFederatedCredentialRequestOidcVerificationConfigKubernetesConfig {
	return s.KubernetesConfig
}

func (s *UpdateApplicationFederatedCredentialRequestOidcVerificationConfig) GetProfile() *string {
	return s.Profile
}

func (s *UpdateApplicationFederatedCredentialRequestOidcVerificationConfig) SetAzureVmConfig(v *UpdateApplicationFederatedCredentialRequestOidcVerificationConfigAzureVmConfig) *UpdateApplicationFederatedCredentialRequestOidcVerificationConfig {
	s.AzureVmConfig = v
	return s
}

func (s *UpdateApplicationFederatedCredentialRequestOidcVerificationConfig) SetGcpVmConfig(v *UpdateApplicationFederatedCredentialRequestOidcVerificationConfigGcpVmConfig) *UpdateApplicationFederatedCredentialRequestOidcVerificationConfig {
	s.GcpVmConfig = v
	return s
}

func (s *UpdateApplicationFederatedCredentialRequestOidcVerificationConfig) SetGenericConfig(v *UpdateApplicationFederatedCredentialRequestOidcVerificationConfigGenericConfig) *UpdateApplicationFederatedCredentialRequestOidcVerificationConfig {
	s.GenericConfig = v
	return s
}

func (s *UpdateApplicationFederatedCredentialRequestOidcVerificationConfig) SetKubernetesConfig(v *UpdateApplicationFederatedCredentialRequestOidcVerificationConfigKubernetesConfig) *UpdateApplicationFederatedCredentialRequestOidcVerificationConfig {
	s.KubernetesConfig = v
	return s
}

func (s *UpdateApplicationFederatedCredentialRequestOidcVerificationConfig) SetProfile(v string) *UpdateApplicationFederatedCredentialRequestOidcVerificationConfig {
	s.Profile = &v
	return s
}

func (s *UpdateApplicationFederatedCredentialRequestOidcVerificationConfig) Validate() error {
	if s.AzureVmConfig != nil {
		if err := s.AzureVmConfig.Validate(); err != nil {
			return err
		}
	}
	if s.GcpVmConfig != nil {
		if err := s.GcpVmConfig.Validate(); err != nil {
			return err
		}
	}
	if s.GenericConfig != nil {
		if err := s.GenericConfig.Validate(); err != nil {
			return err
		}
	}
	if s.KubernetesConfig != nil {
		if err := s.KubernetesConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateApplicationFederatedCredentialRequestOidcVerificationConfigAzureVmConfig struct {
	PrincipalId       *string   `json:"PrincipalId,omitempty" xml:"PrincipalId,omitempty"`
	ResourceGroupName *string   `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
	SubscriptionId    *string   `json:"SubscriptionId,omitempty" xml:"SubscriptionId,omitempty"`
	VmNames           []*string `json:"VmNames,omitempty" xml:"VmNames,omitempty" type:"Repeated"`
}

func (s UpdateApplicationFederatedCredentialRequestOidcVerificationConfigAzureVmConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationFederatedCredentialRequestOidcVerificationConfigAzureVmConfig) GoString() string {
	return s.String()
}

func (s *UpdateApplicationFederatedCredentialRequestOidcVerificationConfigAzureVmConfig) GetPrincipalId() *string {
	return s.PrincipalId
}

func (s *UpdateApplicationFederatedCredentialRequestOidcVerificationConfigAzureVmConfig) GetResourceGroupName() *string {
	return s.ResourceGroupName
}

func (s *UpdateApplicationFederatedCredentialRequestOidcVerificationConfigAzureVmConfig) GetSubscriptionId() *string {
	return s.SubscriptionId
}

func (s *UpdateApplicationFederatedCredentialRequestOidcVerificationConfigAzureVmConfig) GetVmNames() []*string {
	return s.VmNames
}

func (s *UpdateApplicationFederatedCredentialRequestOidcVerificationConfigAzureVmConfig) SetPrincipalId(v string) *UpdateApplicationFederatedCredentialRequestOidcVerificationConfigAzureVmConfig {
	s.PrincipalId = &v
	return s
}

func (s *UpdateApplicationFederatedCredentialRequestOidcVerificationConfigAzureVmConfig) SetResourceGroupName(v string) *UpdateApplicationFederatedCredentialRequestOidcVerificationConfigAzureVmConfig {
	s.ResourceGroupName = &v
	return s
}

func (s *UpdateApplicationFederatedCredentialRequestOidcVerificationConfigAzureVmConfig) SetSubscriptionId(v string) *UpdateApplicationFederatedCredentialRequestOidcVerificationConfigAzureVmConfig {
	s.SubscriptionId = &v
	return s
}

func (s *UpdateApplicationFederatedCredentialRequestOidcVerificationConfigAzureVmConfig) SetVmNames(v []*string) *UpdateApplicationFederatedCredentialRequestOidcVerificationConfigAzureVmConfig {
	s.VmNames = v
	return s
}

func (s *UpdateApplicationFederatedCredentialRequestOidcVerificationConfigAzureVmConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateApplicationFederatedCredentialRequestOidcVerificationConfigGcpVmConfig struct {
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	ProjectId   *string   `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The sub value corresponding to the service account.
	//
	// example:
	//
	// 123456789
	ServiceAccountId *string `json:"ServiceAccountId,omitempty" xml:"ServiceAccountId,omitempty"`
}

func (s UpdateApplicationFederatedCredentialRequestOidcVerificationConfigGcpVmConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationFederatedCredentialRequestOidcVerificationConfigGcpVmConfig) GoString() string {
	return s.String()
}

func (s *UpdateApplicationFederatedCredentialRequestOidcVerificationConfigGcpVmConfig) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *UpdateApplicationFederatedCredentialRequestOidcVerificationConfigGcpVmConfig) GetProjectId() *string {
	return s.ProjectId
}

func (s *UpdateApplicationFederatedCredentialRequestOidcVerificationConfigGcpVmConfig) GetServiceAccountId() *string {
	return s.ServiceAccountId
}

func (s *UpdateApplicationFederatedCredentialRequestOidcVerificationConfigGcpVmConfig) SetInstanceIds(v []*string) *UpdateApplicationFederatedCredentialRequestOidcVerificationConfigGcpVmConfig {
	s.InstanceIds = v
	return s
}

func (s *UpdateApplicationFederatedCredentialRequestOidcVerificationConfigGcpVmConfig) SetProjectId(v string) *UpdateApplicationFederatedCredentialRequestOidcVerificationConfigGcpVmConfig {
	s.ProjectId = &v
	return s
}

func (s *UpdateApplicationFederatedCredentialRequestOidcVerificationConfigGcpVmConfig) SetServiceAccountId(v string) *UpdateApplicationFederatedCredentialRequestOidcVerificationConfigGcpVmConfig {
	s.ServiceAccountId = &v
	return s
}

func (s *UpdateApplicationFederatedCredentialRequestOidcVerificationConfigGcpVmConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateApplicationFederatedCredentialRequestOidcVerificationConfigGenericConfig struct {
	Subject *string `json:"Subject,omitempty" xml:"Subject,omitempty"`
}

func (s UpdateApplicationFederatedCredentialRequestOidcVerificationConfigGenericConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationFederatedCredentialRequestOidcVerificationConfigGenericConfig) GoString() string {
	return s.String()
}

func (s *UpdateApplicationFederatedCredentialRequestOidcVerificationConfigGenericConfig) GetSubject() *string {
	return s.Subject
}

func (s *UpdateApplicationFederatedCredentialRequestOidcVerificationConfigGenericConfig) SetSubject(v string) *UpdateApplicationFederatedCredentialRequestOidcVerificationConfigGenericConfig {
	s.Subject = &v
	return s
}

func (s *UpdateApplicationFederatedCredentialRequestOidcVerificationConfigGenericConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateApplicationFederatedCredentialRequestOidcVerificationConfigKubernetesConfig struct {
	// The Kubernetes namespace.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The pod name prefix.
	//
	// example:
	//
	// my-pod-
	PodNamePrefix *string `json:"PodNamePrefix,omitempty" xml:"PodNamePrefix,omitempty"`
	// The Kubernetes service account name.
	//
	// example:
	//
	// my-sa
	ServiceAccountName *string `json:"ServiceAccountName,omitempty" xml:"ServiceAccountName,omitempty"`
}

func (s UpdateApplicationFederatedCredentialRequestOidcVerificationConfigKubernetesConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationFederatedCredentialRequestOidcVerificationConfigKubernetesConfig) GoString() string {
	return s.String()
}

func (s *UpdateApplicationFederatedCredentialRequestOidcVerificationConfigKubernetesConfig) GetNamespace() *string {
	return s.Namespace
}

func (s *UpdateApplicationFederatedCredentialRequestOidcVerificationConfigKubernetesConfig) GetPodNamePrefix() *string {
	return s.PodNamePrefix
}

func (s *UpdateApplicationFederatedCredentialRequestOidcVerificationConfigKubernetesConfig) GetServiceAccountName() *string {
	return s.ServiceAccountName
}

func (s *UpdateApplicationFederatedCredentialRequestOidcVerificationConfigKubernetesConfig) SetNamespace(v string) *UpdateApplicationFederatedCredentialRequestOidcVerificationConfigKubernetesConfig {
	s.Namespace = &v
	return s
}

func (s *UpdateApplicationFederatedCredentialRequestOidcVerificationConfigKubernetesConfig) SetPodNamePrefix(v string) *UpdateApplicationFederatedCredentialRequestOidcVerificationConfigKubernetesConfig {
	s.PodNamePrefix = &v
	return s
}

func (s *UpdateApplicationFederatedCredentialRequestOidcVerificationConfigKubernetesConfig) SetServiceAccountName(v string) *UpdateApplicationFederatedCredentialRequestOidcVerificationConfigKubernetesConfig {
	s.ServiceAccountName = &v
	return s
}

func (s *UpdateApplicationFederatedCredentialRequestOidcVerificationConfigKubernetesConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateApplicationFederatedCredentialRequestPkcs7VerificationConfig struct {
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
}

func (s UpdateApplicationFederatedCredentialRequestPkcs7VerificationConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationFederatedCredentialRequestPkcs7VerificationConfig) GoString() string {
	return s.String()
}

func (s *UpdateApplicationFederatedCredentialRequestPkcs7VerificationConfig) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *UpdateApplicationFederatedCredentialRequestPkcs7VerificationConfig) SetInstanceIds(v []*string) *UpdateApplicationFederatedCredentialRequestPkcs7VerificationConfig {
	s.InstanceIds = v
	return s
}

func (s *UpdateApplicationFederatedCredentialRequestPkcs7VerificationConfig) Validate() error {
	return dara.Validate(s)
}
