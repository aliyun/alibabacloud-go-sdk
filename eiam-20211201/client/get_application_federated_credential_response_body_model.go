// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApplicationFederatedCredentialResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationFederatedCredential(v *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential) *GetApplicationFederatedCredentialResponseBody
	GetApplicationFederatedCredential() *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential
	SetRequestId(v string) *GetApplicationFederatedCredentialResponseBody
	GetRequestId() *string
}

type GetApplicationFederatedCredentialResponseBody struct {
	// The application federated credential object.
	ApplicationFederatedCredential *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential `json:"ApplicationFederatedCredential,omitempty" xml:"ApplicationFederatedCredential,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetApplicationFederatedCredentialResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationFederatedCredentialResponseBody) GoString() string {
	return s.String()
}

func (s *GetApplicationFederatedCredentialResponseBody) GetApplicationFederatedCredential() *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential {
	return s.ApplicationFederatedCredential
}

func (s *GetApplicationFederatedCredentialResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetApplicationFederatedCredentialResponseBody) SetApplicationFederatedCredential(v *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential) *GetApplicationFederatedCredentialResponseBody {
	s.ApplicationFederatedCredential = v
	return s
}

func (s *GetApplicationFederatedCredentialResponseBody) SetRequestId(v string) *GetApplicationFederatedCredentialResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetApplicationFederatedCredentialResponseBody) Validate() error {
	if s.ApplicationFederatedCredential != nil {
		if err := s.ApplicationFederatedCredential.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential struct {
	// The application federated credential ID.
	//
	// example:
	//
	// afc_aaaaa1111
	ApplicationFederatedCredentialId *string `json:"ApplicationFederatedCredentialId,omitempty" xml:"ApplicationFederatedCredentialId,omitempty"`
	// The application federated credential name.
	//
	// example:
	//
	// test
	ApplicationFederatedCredentialName *string `json:"ApplicationFederatedCredentialName,omitempty" xml:"ApplicationFederatedCredentialName,omitempty"`
	// The application federated credential type.
	//
	// example:
	//
	// oidc
	ApplicationFederatedCredentialType *string `json:"ApplicationFederatedCredentialType,omitempty" xml:"ApplicationFederatedCredentialType,omitempty"`
	// The application ID.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The attribute mappings.
	AttributeMappings []*GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialAttributeMappings `json:"AttributeMappings,omitempty" xml:"AttributeMappings,omitempty" type:"Repeated"`
	// The creation time.
	//
	// example:
	//
	// 1758785994982
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The application federated credential description.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The federated trust source ID.
	//
	// example:
	//
	// fcp_asda1dasdxxxx
	FederatedCredentialProviderId *string `json:"FederatedCredentialProviderId,omitempty" xml:"FederatedCredentialProviderId,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The last used time.
	//
	// example:
	//
	// 1758785994982
	LastUsedTime *int64 `json:"LastUsedTime,omitempty" xml:"LastUsedTime,omitempty"`
	// The OIDC structured configuration.
	OidcVerificationConfig *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialOidcVerificationConfig `json:"OidcVerificationConfig,omitempty" xml:"OidcVerificationConfig,omitempty" type:"Struct"`
	// The PKCS#7 structured configuration.
	Pkcs7VerificationConfig *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialPkcs7VerificationConfig `json:"Pkcs7VerificationConfig,omitempty" xml:"Pkcs7VerificationConfig,omitempty" type:"Struct"`
	// The application federated credential status.
	//
	// example:
	//
	// enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The update time.
	//
	// example:
	//
	// 1758785994982
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The verification condition.
	//
	// example:
	//
	// Equals(cert.subject.CN, "test")
	VerificationCondition *string `json:"VerificationCondition,omitempty" xml:"VerificationCondition,omitempty"`
	// The verification mode. Valid values:
	//
	// - freedom: Free mode.
	//
	// - structured: Structured mode.
	//
	// example:
	//
	// structured
	VerificationMode *string `json:"VerificationMode,omitempty" xml:"VerificationMode,omitempty"`
}

func (s GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential) GoString() string {
	return s.String()
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential) GetApplicationFederatedCredentialId() *string {
	return s.ApplicationFederatedCredentialId
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential) GetApplicationFederatedCredentialName() *string {
	return s.ApplicationFederatedCredentialName
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential) GetApplicationFederatedCredentialType() *string {
	return s.ApplicationFederatedCredentialType
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential) GetAttributeMappings() []*GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialAttributeMappings {
	return s.AttributeMappings
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential) GetDescription() *string {
	return s.Description
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential) GetFederatedCredentialProviderId() *string {
	return s.FederatedCredentialProviderId
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential) GetLastUsedTime() *int64 {
	return s.LastUsedTime
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential) GetOidcVerificationConfig() *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialOidcVerificationConfig {
	return s.OidcVerificationConfig
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential) GetPkcs7VerificationConfig() *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialPkcs7VerificationConfig {
	return s.Pkcs7VerificationConfig
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential) GetStatus() *string {
	return s.Status
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential) GetVerificationCondition() *string {
	return s.VerificationCondition
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential) GetVerificationMode() *string {
	return s.VerificationMode
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential) SetApplicationFederatedCredentialId(v string) *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential {
	s.ApplicationFederatedCredentialId = &v
	return s
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential) SetApplicationFederatedCredentialName(v string) *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential {
	s.ApplicationFederatedCredentialName = &v
	return s
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential) SetApplicationFederatedCredentialType(v string) *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential {
	s.ApplicationFederatedCredentialType = &v
	return s
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential) SetApplicationId(v string) *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential {
	s.ApplicationId = &v
	return s
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential) SetAttributeMappings(v []*GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialAttributeMappings) *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential {
	s.AttributeMappings = v
	return s
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential) SetCreateTime(v int64) *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential {
	s.CreateTime = &v
	return s
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential) SetDescription(v string) *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential {
	s.Description = &v
	return s
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential) SetFederatedCredentialProviderId(v string) *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential {
	s.FederatedCredentialProviderId = &v
	return s
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential) SetInstanceId(v string) *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential {
	s.InstanceId = &v
	return s
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential) SetLastUsedTime(v int64) *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential {
	s.LastUsedTime = &v
	return s
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential) SetOidcVerificationConfig(v *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialOidcVerificationConfig) *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential {
	s.OidcVerificationConfig = v
	return s
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential) SetPkcs7VerificationConfig(v *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialPkcs7VerificationConfig) *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential {
	s.Pkcs7VerificationConfig = v
	return s
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential) SetStatus(v string) *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential {
	s.Status = &v
	return s
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential) SetUpdateTime(v int64) *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential {
	s.UpdateTime = &v
	return s
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential) SetVerificationCondition(v string) *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential {
	s.VerificationCondition = &v
	return s
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential) SetVerificationMode(v string) *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential {
	s.VerificationMode = &v
	return s
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential) Validate() error {
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

type GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialAttributeMappings struct {
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

func (s GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialAttributeMappings) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialAttributeMappings) GoString() string {
	return s.String()
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialAttributeMappings) GetSourceValueExpression() *string {
	return s.SourceValueExpression
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialAttributeMappings) GetTargetField() *string {
	return s.TargetField
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialAttributeMappings) SetSourceValueExpression(v string) *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialAttributeMappings {
	s.SourceValueExpression = &v
	return s
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialAttributeMappings) SetTargetField(v string) *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialAttributeMappings {
	s.TargetField = &v
	return s
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialAttributeMappings) Validate() error {
	return dara.Validate(s)
}

type GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialOidcVerificationConfig struct {
	// The Azure VM scenario configuration.
	AzureVmConfig *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialOidcVerificationConfigAzureVmConfig `json:"AzureVmConfig,omitempty" xml:"AzureVmConfig,omitempty" type:"Struct"`
	// The GCP VM scenario configuration.
	GcpVmConfig *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialOidcVerificationConfigGcpVmConfig `json:"GcpVmConfig,omitempty" xml:"GcpVmConfig,omitempty" type:"Struct"`
	// The generic scenario configuration.
	GenericConfig *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialOidcVerificationConfigGenericConfig `json:"GenericConfig,omitempty" xml:"GenericConfig,omitempty" type:"Struct"`
	// The Kubernetes scenario configuration.
	KubernetesConfig *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialOidcVerificationConfigKubernetesConfig `json:"KubernetesConfig,omitempty" xml:"KubernetesConfig,omitempty" type:"Struct"`
	// The OIDC scenario profile. Different profiles correspond to different configurations. Valid values:
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

func (s GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialOidcVerificationConfig) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialOidcVerificationConfig) GoString() string {
	return s.String()
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialOidcVerificationConfig) GetAzureVmConfig() *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialOidcVerificationConfigAzureVmConfig {
	return s.AzureVmConfig
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialOidcVerificationConfig) GetGcpVmConfig() *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialOidcVerificationConfigGcpVmConfig {
	return s.GcpVmConfig
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialOidcVerificationConfig) GetGenericConfig() *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialOidcVerificationConfigGenericConfig {
	return s.GenericConfig
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialOidcVerificationConfig) GetKubernetesConfig() *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialOidcVerificationConfigKubernetesConfig {
	return s.KubernetesConfig
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialOidcVerificationConfig) GetProfile() *string {
	return s.Profile
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialOidcVerificationConfig) SetAzureVmConfig(v *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialOidcVerificationConfigAzureVmConfig) *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialOidcVerificationConfig {
	s.AzureVmConfig = v
	return s
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialOidcVerificationConfig) SetGcpVmConfig(v *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialOidcVerificationConfigGcpVmConfig) *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialOidcVerificationConfig {
	s.GcpVmConfig = v
	return s
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialOidcVerificationConfig) SetGenericConfig(v *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialOidcVerificationConfigGenericConfig) *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialOidcVerificationConfig {
	s.GenericConfig = v
	return s
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialOidcVerificationConfig) SetKubernetesConfig(v *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialOidcVerificationConfigKubernetesConfig) *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialOidcVerificationConfig {
	s.KubernetesConfig = v
	return s
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialOidcVerificationConfig) SetProfile(v string) *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialOidcVerificationConfig {
	s.Profile = &v
	return s
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialOidcVerificationConfig) Validate() error {
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

type GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialOidcVerificationConfigAzureVmConfig struct {
	// The principal ID.
	//
	// example:
	//
	// 5dee234a-1b4c-4ad7-a19f-fxxxxx
	PrincipalId *string `json:"PrincipalId,omitempty" xml:"PrincipalId,omitempty"`
	// The Azure resource group name.
	//
	// example:
	//
	// user_default
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
	// The subscription ID.
	//
	// example:
	//
	// 4342a1f4-7e5d-4371-97dc-d4f33f4xxxx
	SubscriptionId *string `json:"SubscriptionId,omitempty" xml:"SubscriptionId,omitempty"`
	// The list of virtual machine names.
	VmNames []*string `json:"VmNames,omitempty" xml:"VmNames,omitempty" type:"Repeated"`
}

func (s GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialOidcVerificationConfigAzureVmConfig) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialOidcVerificationConfigAzureVmConfig) GoString() string {
	return s.String()
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialOidcVerificationConfigAzureVmConfig) GetPrincipalId() *string {
	return s.PrincipalId
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialOidcVerificationConfigAzureVmConfig) GetResourceGroupName() *string {
	return s.ResourceGroupName
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialOidcVerificationConfigAzureVmConfig) GetSubscriptionId() *string {
	return s.SubscriptionId
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialOidcVerificationConfigAzureVmConfig) GetVmNames() []*string {
	return s.VmNames
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialOidcVerificationConfigAzureVmConfig) SetPrincipalId(v string) *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialOidcVerificationConfigAzureVmConfig {
	s.PrincipalId = &v
	return s
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialOidcVerificationConfigAzureVmConfig) SetResourceGroupName(v string) *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialOidcVerificationConfigAzureVmConfig {
	s.ResourceGroupName = &v
	return s
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialOidcVerificationConfigAzureVmConfig) SetSubscriptionId(v string) *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialOidcVerificationConfigAzureVmConfig {
	s.SubscriptionId = &v
	return s
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialOidcVerificationConfigAzureVmConfig) SetVmNames(v []*string) *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialOidcVerificationConfigAzureVmConfig {
	s.VmNames = v
	return s
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialOidcVerificationConfigAzureVmConfig) Validate() error {
	return dara.Validate(s)
}

type GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialOidcVerificationConfigGcpVmConfig struct {
	// The list of VM instance IDs.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The GCP project ID.
	//
	// example:
	//
	// turnkey-axiom-475109-xx
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The subject corresponding to the service account.
	//
	// example:
	//
	// 123456789
	ServiceAccountId *string `json:"ServiceAccountId,omitempty" xml:"ServiceAccountId,omitempty"`
}

func (s GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialOidcVerificationConfigGcpVmConfig) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialOidcVerificationConfigGcpVmConfig) GoString() string {
	return s.String()
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialOidcVerificationConfigGcpVmConfig) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialOidcVerificationConfigGcpVmConfig) GetProjectId() *string {
	return s.ProjectId
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialOidcVerificationConfigGcpVmConfig) GetServiceAccountId() *string {
	return s.ServiceAccountId
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialOidcVerificationConfigGcpVmConfig) SetInstanceIds(v []*string) *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialOidcVerificationConfigGcpVmConfig {
	s.InstanceIds = v
	return s
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialOidcVerificationConfigGcpVmConfig) SetProjectId(v string) *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialOidcVerificationConfigGcpVmConfig {
	s.ProjectId = &v
	return s
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialOidcVerificationConfigGcpVmConfig) SetServiceAccountId(v string) *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialOidcVerificationConfigGcpVmConfig {
	s.ServiceAccountId = &v
	return s
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialOidcVerificationConfigGcpVmConfig) Validate() error {
	return dara.Validate(s)
}

type GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialOidcVerificationConfigGenericConfig struct {
	// The subject identifier.
	//
	// example:
	//
	// test_subject
	Subject *string `json:"Subject,omitempty" xml:"Subject,omitempty"`
}

func (s GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialOidcVerificationConfigGenericConfig) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialOidcVerificationConfigGenericConfig) GoString() string {
	return s.String()
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialOidcVerificationConfigGenericConfig) GetSubject() *string {
	return s.Subject
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialOidcVerificationConfigGenericConfig) SetSubject(v string) *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialOidcVerificationConfigGenericConfig {
	s.Subject = &v
	return s
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialOidcVerificationConfigGenericConfig) Validate() error {
	return dara.Validate(s)
}

type GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialOidcVerificationConfigKubernetesConfig struct {
	// The K8s namespace.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The pod name prefix.
	//
	// example:
	//
	// test-pod
	PodNamePrefix *string `json:"PodNamePrefix,omitempty" xml:"PodNamePrefix,omitempty"`
	// The K8s service account name.
	//
	// example:
	//
	// default
	ServiceAccountName *string `json:"ServiceAccountName,omitempty" xml:"ServiceAccountName,omitempty"`
}

func (s GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialOidcVerificationConfigKubernetesConfig) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialOidcVerificationConfigKubernetesConfig) GoString() string {
	return s.String()
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialOidcVerificationConfigKubernetesConfig) GetNamespace() *string {
	return s.Namespace
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialOidcVerificationConfigKubernetesConfig) GetPodNamePrefix() *string {
	return s.PodNamePrefix
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialOidcVerificationConfigKubernetesConfig) GetServiceAccountName() *string {
	return s.ServiceAccountName
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialOidcVerificationConfigKubernetesConfig) SetNamespace(v string) *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialOidcVerificationConfigKubernetesConfig {
	s.Namespace = &v
	return s
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialOidcVerificationConfigKubernetesConfig) SetPodNamePrefix(v string) *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialOidcVerificationConfigKubernetesConfig {
	s.PodNamePrefix = &v
	return s
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialOidcVerificationConfigKubernetesConfig) SetServiceAccountName(v string) *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialOidcVerificationConfigKubernetesConfig {
	s.ServiceAccountName = &v
	return s
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialOidcVerificationConfigKubernetesConfig) Validate() error {
	return dara.Validate(s)
}

type GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialPkcs7VerificationConfig struct {
	// The list of allowed instance IDs.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
}

func (s GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialPkcs7VerificationConfig) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialPkcs7VerificationConfig) GoString() string {
	return s.String()
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialPkcs7VerificationConfig) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialPkcs7VerificationConfig) SetInstanceIds(v []*string) *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialPkcs7VerificationConfig {
	s.InstanceIds = v
	return s
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialPkcs7VerificationConfig) Validate() error {
	return dara.Validate(s)
}
