// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApplicationFederatedCredentialRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationFederatedCredentialName(v string) *CreateApplicationFederatedCredentialRequest
	GetApplicationFederatedCredentialName() *string
	SetApplicationFederatedCredentialType(v string) *CreateApplicationFederatedCredentialRequest
	GetApplicationFederatedCredentialType() *string
	SetApplicationId(v string) *CreateApplicationFederatedCredentialRequest
	GetApplicationId() *string
	SetAttributeMappings(v []*CreateApplicationFederatedCredentialRequestAttributeMappings) *CreateApplicationFederatedCredentialRequest
	GetAttributeMappings() []*CreateApplicationFederatedCredentialRequestAttributeMappings
	SetDescription(v string) *CreateApplicationFederatedCredentialRequest
	GetDescription() *string
	SetFederatedCredentialProviderId(v string) *CreateApplicationFederatedCredentialRequest
	GetFederatedCredentialProviderId() *string
	SetInstanceId(v string) *CreateApplicationFederatedCredentialRequest
	GetInstanceId() *string
	SetOidcVerificationConfig(v *CreateApplicationFederatedCredentialRequestOidcVerificationConfig) *CreateApplicationFederatedCredentialRequest
	GetOidcVerificationConfig() *CreateApplicationFederatedCredentialRequestOidcVerificationConfig
	SetPkcs7VerificationConfig(v *CreateApplicationFederatedCredentialRequestPkcs7VerificationConfig) *CreateApplicationFederatedCredentialRequest
	GetPkcs7VerificationConfig() *CreateApplicationFederatedCredentialRequestPkcs7VerificationConfig
	SetVerificationCondition(v string) *CreateApplicationFederatedCredentialRequest
	GetVerificationCondition() *string
	SetVerificationMode(v string) *CreateApplicationFederatedCredentialRequest
	GetVerificationMode() *string
}

type CreateApplicationFederatedCredentialRequest struct {
	// The name of the application federated identity credential.
	//
	// This parameter is required.
	//
	// example:
	//
	// example_name
	ApplicationFederatedCredentialName *string `json:"ApplicationFederatedCredentialName,omitempty" xml:"ApplicationFederatedCredentialName,omitempty"`
	// The type of the application federated identity credential.
	//
	// This parameter is required.
	//
	// example:
	//
	// oidc
	ApplicationFederatedCredentialType *string `json:"ApplicationFederatedCredentialType,omitempty" xml:"ApplicationFederatedCredentialType,omitempty"`
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The attribute mappings.
	AttributeMappings []*CreateApplicationFederatedCredentialRequestAttributeMappings `json:"AttributeMappings,omitempty" xml:"AttributeMappings,omitempty" type:"Repeated"`
	// The description.
	//
	// example:
	//
	// description_text
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The federated credential provider ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// fcp_adasd
	FederatedCredentialProviderId *string `json:"FederatedCredentialProviderId,omitempty" xml:"FederatedCredentialProviderId,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The OIDC structured configuration. This parameter applies when the verification mode is structured and the credential type is oidc.
	OidcVerificationConfig *CreateApplicationFederatedCredentialRequestOidcVerificationConfig `json:"OidcVerificationConfig,omitempty" xml:"OidcVerificationConfig,omitempty" type:"Struct"`
	// The PKCS#7 structured configuration. This parameter applies when the verification mode is structured and the credential type is pkcs7.
	Pkcs7VerificationConfig *CreateApplicationFederatedCredentialRequestPkcs7VerificationConfig `json:"Pkcs7VerificationConfig,omitempty" xml:"Pkcs7VerificationConfig,omitempty" type:"Struct"`
	// The verification condition.
	//
	// example:
	//
	// IsNullOrEmpty("jwt.issuer")
	VerificationCondition *string `json:"VerificationCondition,omitempty" xml:"VerificationCondition,omitempty"`
	// The verification mode. Valid values:
	//
	// - freedom (default)
	//
	// - structured
	//
	// example:
	//
	// freedom
	VerificationMode *string `json:"VerificationMode,omitempty" xml:"VerificationMode,omitempty"`
}

func (s CreateApplicationFederatedCredentialRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationFederatedCredentialRequest) GoString() string {
	return s.String()
}

func (s *CreateApplicationFederatedCredentialRequest) GetApplicationFederatedCredentialName() *string {
	return s.ApplicationFederatedCredentialName
}

func (s *CreateApplicationFederatedCredentialRequest) GetApplicationFederatedCredentialType() *string {
	return s.ApplicationFederatedCredentialType
}

func (s *CreateApplicationFederatedCredentialRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *CreateApplicationFederatedCredentialRequest) GetAttributeMappings() []*CreateApplicationFederatedCredentialRequestAttributeMappings {
	return s.AttributeMappings
}

func (s *CreateApplicationFederatedCredentialRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateApplicationFederatedCredentialRequest) GetFederatedCredentialProviderId() *string {
	return s.FederatedCredentialProviderId
}

func (s *CreateApplicationFederatedCredentialRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateApplicationFederatedCredentialRequest) GetOidcVerificationConfig() *CreateApplicationFederatedCredentialRequestOidcVerificationConfig {
	return s.OidcVerificationConfig
}

func (s *CreateApplicationFederatedCredentialRequest) GetPkcs7VerificationConfig() *CreateApplicationFederatedCredentialRequestPkcs7VerificationConfig {
	return s.Pkcs7VerificationConfig
}

func (s *CreateApplicationFederatedCredentialRequest) GetVerificationCondition() *string {
	return s.VerificationCondition
}

func (s *CreateApplicationFederatedCredentialRequest) GetVerificationMode() *string {
	return s.VerificationMode
}

func (s *CreateApplicationFederatedCredentialRequest) SetApplicationFederatedCredentialName(v string) *CreateApplicationFederatedCredentialRequest {
	s.ApplicationFederatedCredentialName = &v
	return s
}

func (s *CreateApplicationFederatedCredentialRequest) SetApplicationFederatedCredentialType(v string) *CreateApplicationFederatedCredentialRequest {
	s.ApplicationFederatedCredentialType = &v
	return s
}

func (s *CreateApplicationFederatedCredentialRequest) SetApplicationId(v string) *CreateApplicationFederatedCredentialRequest {
	s.ApplicationId = &v
	return s
}

func (s *CreateApplicationFederatedCredentialRequest) SetAttributeMappings(v []*CreateApplicationFederatedCredentialRequestAttributeMappings) *CreateApplicationFederatedCredentialRequest {
	s.AttributeMappings = v
	return s
}

func (s *CreateApplicationFederatedCredentialRequest) SetDescription(v string) *CreateApplicationFederatedCredentialRequest {
	s.Description = &v
	return s
}

func (s *CreateApplicationFederatedCredentialRequest) SetFederatedCredentialProviderId(v string) *CreateApplicationFederatedCredentialRequest {
	s.FederatedCredentialProviderId = &v
	return s
}

func (s *CreateApplicationFederatedCredentialRequest) SetInstanceId(v string) *CreateApplicationFederatedCredentialRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateApplicationFederatedCredentialRequest) SetOidcVerificationConfig(v *CreateApplicationFederatedCredentialRequestOidcVerificationConfig) *CreateApplicationFederatedCredentialRequest {
	s.OidcVerificationConfig = v
	return s
}

func (s *CreateApplicationFederatedCredentialRequest) SetPkcs7VerificationConfig(v *CreateApplicationFederatedCredentialRequestPkcs7VerificationConfig) *CreateApplicationFederatedCredentialRequest {
	s.Pkcs7VerificationConfig = v
	return s
}

func (s *CreateApplicationFederatedCredentialRequest) SetVerificationCondition(v string) *CreateApplicationFederatedCredentialRequest {
	s.VerificationCondition = &v
	return s
}

func (s *CreateApplicationFederatedCredentialRequest) SetVerificationMode(v string) *CreateApplicationFederatedCredentialRequest {
	s.VerificationMode = &v
	return s
}

func (s *CreateApplicationFederatedCredentialRequest) Validate() error {
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

type CreateApplicationFederatedCredentialRequestAttributeMappings struct {
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

func (s CreateApplicationFederatedCredentialRequestAttributeMappings) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationFederatedCredentialRequestAttributeMappings) GoString() string {
	return s.String()
}

func (s *CreateApplicationFederatedCredentialRequestAttributeMappings) GetSourceValueExpression() *string {
	return s.SourceValueExpression
}

func (s *CreateApplicationFederatedCredentialRequestAttributeMappings) GetTargetField() *string {
	return s.TargetField
}

func (s *CreateApplicationFederatedCredentialRequestAttributeMappings) SetSourceValueExpression(v string) *CreateApplicationFederatedCredentialRequestAttributeMappings {
	s.SourceValueExpression = &v
	return s
}

func (s *CreateApplicationFederatedCredentialRequestAttributeMappings) SetTargetField(v string) *CreateApplicationFederatedCredentialRequestAttributeMappings {
	s.TargetField = &v
	return s
}

func (s *CreateApplicationFederatedCredentialRequestAttributeMappings) Validate() error {
	return dara.Validate(s)
}

type CreateApplicationFederatedCredentialRequestOidcVerificationConfig struct {
	// The Azure VM scenario configuration.
	AzureVmConfig *CreateApplicationFederatedCredentialRequestOidcVerificationConfigAzureVmConfig `json:"AzureVmConfig,omitempty" xml:"AzureVmConfig,omitempty" type:"Struct"`
	// The GCP VM scenario configuration.
	GcpVmConfig   *CreateApplicationFederatedCredentialRequestOidcVerificationConfigGcpVmConfig   `json:"GcpVmConfig,omitempty" xml:"GcpVmConfig,omitempty" type:"Struct"`
	GenericConfig *CreateApplicationFederatedCredentialRequestOidcVerificationConfigGenericConfig `json:"GenericConfig,omitempty" xml:"GenericConfig,omitempty" type:"Struct"`
	// The Kubernetes scenario configuration.
	KubernetesConfig *CreateApplicationFederatedCredentialRequestOidcVerificationConfigKubernetesConfig `json:"KubernetesConfig,omitempty" xml:"KubernetesConfig,omitempty" type:"Struct"`
	// The OIDC scenario profile. Valid values:
	//
	// - generic
	//
	// - kubernetes
	//
	// - gcp_vm
	//
	// - azure_vm
	//
	// example:
	//
	// kubernetes
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
}

func (s CreateApplicationFederatedCredentialRequestOidcVerificationConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationFederatedCredentialRequestOidcVerificationConfig) GoString() string {
	return s.String()
}

func (s *CreateApplicationFederatedCredentialRequestOidcVerificationConfig) GetAzureVmConfig() *CreateApplicationFederatedCredentialRequestOidcVerificationConfigAzureVmConfig {
	return s.AzureVmConfig
}

func (s *CreateApplicationFederatedCredentialRequestOidcVerificationConfig) GetGcpVmConfig() *CreateApplicationFederatedCredentialRequestOidcVerificationConfigGcpVmConfig {
	return s.GcpVmConfig
}

func (s *CreateApplicationFederatedCredentialRequestOidcVerificationConfig) GetGenericConfig() *CreateApplicationFederatedCredentialRequestOidcVerificationConfigGenericConfig {
	return s.GenericConfig
}

func (s *CreateApplicationFederatedCredentialRequestOidcVerificationConfig) GetKubernetesConfig() *CreateApplicationFederatedCredentialRequestOidcVerificationConfigKubernetesConfig {
	return s.KubernetesConfig
}

func (s *CreateApplicationFederatedCredentialRequestOidcVerificationConfig) GetProfile() *string {
	return s.Profile
}

func (s *CreateApplicationFederatedCredentialRequestOidcVerificationConfig) SetAzureVmConfig(v *CreateApplicationFederatedCredentialRequestOidcVerificationConfigAzureVmConfig) *CreateApplicationFederatedCredentialRequestOidcVerificationConfig {
	s.AzureVmConfig = v
	return s
}

func (s *CreateApplicationFederatedCredentialRequestOidcVerificationConfig) SetGcpVmConfig(v *CreateApplicationFederatedCredentialRequestOidcVerificationConfigGcpVmConfig) *CreateApplicationFederatedCredentialRequestOidcVerificationConfig {
	s.GcpVmConfig = v
	return s
}

func (s *CreateApplicationFederatedCredentialRequestOidcVerificationConfig) SetGenericConfig(v *CreateApplicationFederatedCredentialRequestOidcVerificationConfigGenericConfig) *CreateApplicationFederatedCredentialRequestOidcVerificationConfig {
	s.GenericConfig = v
	return s
}

func (s *CreateApplicationFederatedCredentialRequestOidcVerificationConfig) SetKubernetesConfig(v *CreateApplicationFederatedCredentialRequestOidcVerificationConfigKubernetesConfig) *CreateApplicationFederatedCredentialRequestOidcVerificationConfig {
	s.KubernetesConfig = v
	return s
}

func (s *CreateApplicationFederatedCredentialRequestOidcVerificationConfig) SetProfile(v string) *CreateApplicationFederatedCredentialRequestOidcVerificationConfig {
	s.Profile = &v
	return s
}

func (s *CreateApplicationFederatedCredentialRequestOidcVerificationConfig) Validate() error {
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

type CreateApplicationFederatedCredentialRequestOidcVerificationConfigAzureVmConfig struct {
	PrincipalId       *string   `json:"PrincipalId,omitempty" xml:"PrincipalId,omitempty"`
	ResourceGroupName *string   `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
	SubscriptionId    *string   `json:"SubscriptionId,omitempty" xml:"SubscriptionId,omitempty"`
	VmNames           []*string `json:"VmNames,omitempty" xml:"VmNames,omitempty" type:"Repeated"`
}

func (s CreateApplicationFederatedCredentialRequestOidcVerificationConfigAzureVmConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationFederatedCredentialRequestOidcVerificationConfigAzureVmConfig) GoString() string {
	return s.String()
}

func (s *CreateApplicationFederatedCredentialRequestOidcVerificationConfigAzureVmConfig) GetPrincipalId() *string {
	return s.PrincipalId
}

func (s *CreateApplicationFederatedCredentialRequestOidcVerificationConfigAzureVmConfig) GetResourceGroupName() *string {
	return s.ResourceGroupName
}

func (s *CreateApplicationFederatedCredentialRequestOidcVerificationConfigAzureVmConfig) GetSubscriptionId() *string {
	return s.SubscriptionId
}

func (s *CreateApplicationFederatedCredentialRequestOidcVerificationConfigAzureVmConfig) GetVmNames() []*string {
	return s.VmNames
}

func (s *CreateApplicationFederatedCredentialRequestOidcVerificationConfigAzureVmConfig) SetPrincipalId(v string) *CreateApplicationFederatedCredentialRequestOidcVerificationConfigAzureVmConfig {
	s.PrincipalId = &v
	return s
}

func (s *CreateApplicationFederatedCredentialRequestOidcVerificationConfigAzureVmConfig) SetResourceGroupName(v string) *CreateApplicationFederatedCredentialRequestOidcVerificationConfigAzureVmConfig {
	s.ResourceGroupName = &v
	return s
}

func (s *CreateApplicationFederatedCredentialRequestOidcVerificationConfigAzureVmConfig) SetSubscriptionId(v string) *CreateApplicationFederatedCredentialRequestOidcVerificationConfigAzureVmConfig {
	s.SubscriptionId = &v
	return s
}

func (s *CreateApplicationFederatedCredentialRequestOidcVerificationConfigAzureVmConfig) SetVmNames(v []*string) *CreateApplicationFederatedCredentialRequestOidcVerificationConfigAzureVmConfig {
	s.VmNames = v
	return s
}

func (s *CreateApplicationFederatedCredentialRequestOidcVerificationConfigAzureVmConfig) Validate() error {
	return dara.Validate(s)
}

type CreateApplicationFederatedCredentialRequestOidcVerificationConfigGcpVmConfig struct {
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	ProjectId   *string   `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The sub claim that corresponds to the service account.
	//
	// example:
	//
	// 123456789
	ServiceAccountId *string `json:"ServiceAccountId,omitempty" xml:"ServiceAccountId,omitempty"`
}

func (s CreateApplicationFederatedCredentialRequestOidcVerificationConfigGcpVmConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationFederatedCredentialRequestOidcVerificationConfigGcpVmConfig) GoString() string {
	return s.String()
}

func (s *CreateApplicationFederatedCredentialRequestOidcVerificationConfigGcpVmConfig) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *CreateApplicationFederatedCredentialRequestOidcVerificationConfigGcpVmConfig) GetProjectId() *string {
	return s.ProjectId
}

func (s *CreateApplicationFederatedCredentialRequestOidcVerificationConfigGcpVmConfig) GetServiceAccountId() *string {
	return s.ServiceAccountId
}

func (s *CreateApplicationFederatedCredentialRequestOidcVerificationConfigGcpVmConfig) SetInstanceIds(v []*string) *CreateApplicationFederatedCredentialRequestOidcVerificationConfigGcpVmConfig {
	s.InstanceIds = v
	return s
}

func (s *CreateApplicationFederatedCredentialRequestOidcVerificationConfigGcpVmConfig) SetProjectId(v string) *CreateApplicationFederatedCredentialRequestOidcVerificationConfigGcpVmConfig {
	s.ProjectId = &v
	return s
}

func (s *CreateApplicationFederatedCredentialRequestOidcVerificationConfigGcpVmConfig) SetServiceAccountId(v string) *CreateApplicationFederatedCredentialRequestOidcVerificationConfigGcpVmConfig {
	s.ServiceAccountId = &v
	return s
}

func (s *CreateApplicationFederatedCredentialRequestOidcVerificationConfigGcpVmConfig) Validate() error {
	return dara.Validate(s)
}

type CreateApplicationFederatedCredentialRequestOidcVerificationConfigGenericConfig struct {
	Subject *string `json:"Subject,omitempty" xml:"Subject,omitempty"`
}

func (s CreateApplicationFederatedCredentialRequestOidcVerificationConfigGenericConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationFederatedCredentialRequestOidcVerificationConfigGenericConfig) GoString() string {
	return s.String()
}

func (s *CreateApplicationFederatedCredentialRequestOidcVerificationConfigGenericConfig) GetSubject() *string {
	return s.Subject
}

func (s *CreateApplicationFederatedCredentialRequestOidcVerificationConfigGenericConfig) SetSubject(v string) *CreateApplicationFederatedCredentialRequestOidcVerificationConfigGenericConfig {
	s.Subject = &v
	return s
}

func (s *CreateApplicationFederatedCredentialRequestOidcVerificationConfigGenericConfig) Validate() error {
	return dara.Validate(s)
}

type CreateApplicationFederatedCredentialRequestOidcVerificationConfigKubernetesConfig struct {
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

func (s CreateApplicationFederatedCredentialRequestOidcVerificationConfigKubernetesConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationFederatedCredentialRequestOidcVerificationConfigKubernetesConfig) GoString() string {
	return s.String()
}

func (s *CreateApplicationFederatedCredentialRequestOidcVerificationConfigKubernetesConfig) GetNamespace() *string {
	return s.Namespace
}

func (s *CreateApplicationFederatedCredentialRequestOidcVerificationConfigKubernetesConfig) GetPodNamePrefix() *string {
	return s.PodNamePrefix
}

func (s *CreateApplicationFederatedCredentialRequestOidcVerificationConfigKubernetesConfig) GetServiceAccountName() *string {
	return s.ServiceAccountName
}

func (s *CreateApplicationFederatedCredentialRequestOidcVerificationConfigKubernetesConfig) SetNamespace(v string) *CreateApplicationFederatedCredentialRequestOidcVerificationConfigKubernetesConfig {
	s.Namespace = &v
	return s
}

func (s *CreateApplicationFederatedCredentialRequestOidcVerificationConfigKubernetesConfig) SetPodNamePrefix(v string) *CreateApplicationFederatedCredentialRequestOidcVerificationConfigKubernetesConfig {
	s.PodNamePrefix = &v
	return s
}

func (s *CreateApplicationFederatedCredentialRequestOidcVerificationConfigKubernetesConfig) SetServiceAccountName(v string) *CreateApplicationFederatedCredentialRequestOidcVerificationConfigKubernetesConfig {
	s.ServiceAccountName = &v
	return s
}

func (s *CreateApplicationFederatedCredentialRequestOidcVerificationConfigKubernetesConfig) Validate() error {
	return dara.Validate(s)
}

type CreateApplicationFederatedCredentialRequestPkcs7VerificationConfig struct {
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
}

func (s CreateApplicationFederatedCredentialRequestPkcs7VerificationConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationFederatedCredentialRequestPkcs7VerificationConfig) GoString() string {
	return s.String()
}

func (s *CreateApplicationFederatedCredentialRequestPkcs7VerificationConfig) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *CreateApplicationFederatedCredentialRequestPkcs7VerificationConfig) SetInstanceIds(v []*string) *CreateApplicationFederatedCredentialRequestPkcs7VerificationConfig {
	s.InstanceIds = v
	return s
}

func (s *CreateApplicationFederatedCredentialRequestPkcs7VerificationConfig) Validate() error {
	return dara.Validate(s)
}
