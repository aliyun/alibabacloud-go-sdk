// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationFederatedCredentialsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationFederatedCredentials(v []*ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials) *ListApplicationFederatedCredentialsResponseBody
	GetApplicationFederatedCredentials() []*ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials
	SetMaxResults(v int32) *ListApplicationFederatedCredentialsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListApplicationFederatedCredentialsResponseBody
	GetNextToken() *string
	SetPreviousToken(v string) *ListApplicationFederatedCredentialsResponseBody
	GetPreviousToken() *string
	SetRequestId(v string) *ListApplicationFederatedCredentialsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListApplicationFederatedCredentialsResponseBody
	GetTotalCount() *int64
}

type ListApplicationFederatedCredentialsResponseBody struct {
	// The list of application federated credentials.
	ApplicationFederatedCredentials []*ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials `json:"ApplicationFederatedCredentials,omitempty" xml:"ApplicationFederatedCredentials,omitempty" type:"Repeated"`
	// The number of entries per page in a paged query. This parameter is used for paging.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token returned by this call.
	//
	// example:
	//
	// NTxxxexample
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The pagination token returned by this call.
	//
	// example:
	//
	// PTxxxexample
	PreviousToken *string `json:"PreviousToken,omitempty" xml:"PreviousToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListApplicationFederatedCredentialsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationFederatedCredentialsResponseBody) GoString() string {
	return s.String()
}

func (s *ListApplicationFederatedCredentialsResponseBody) GetApplicationFederatedCredentials() []*ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials {
	return s.ApplicationFederatedCredentials
}

func (s *ListApplicationFederatedCredentialsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListApplicationFederatedCredentialsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListApplicationFederatedCredentialsResponseBody) GetPreviousToken() *string {
	return s.PreviousToken
}

func (s *ListApplicationFederatedCredentialsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListApplicationFederatedCredentialsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListApplicationFederatedCredentialsResponseBody) SetApplicationFederatedCredentials(v []*ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials) *ListApplicationFederatedCredentialsResponseBody {
	s.ApplicationFederatedCredentials = v
	return s
}

func (s *ListApplicationFederatedCredentialsResponseBody) SetMaxResults(v int32) *ListApplicationFederatedCredentialsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListApplicationFederatedCredentialsResponseBody) SetNextToken(v string) *ListApplicationFederatedCredentialsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListApplicationFederatedCredentialsResponseBody) SetPreviousToken(v string) *ListApplicationFederatedCredentialsResponseBody {
	s.PreviousToken = &v
	return s
}

func (s *ListApplicationFederatedCredentialsResponseBody) SetRequestId(v string) *ListApplicationFederatedCredentialsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApplicationFederatedCredentialsResponseBody) SetTotalCount(v int64) *ListApplicationFederatedCredentialsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListApplicationFederatedCredentialsResponseBody) Validate() error {
	if s.ApplicationFederatedCredentials != nil {
		for _, item := range s.ApplicationFederatedCredentials {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials struct {
	// The application federated credential ID.
	//
	// example:
	//
	// afc_adsa1sdaxxxxx
	ApplicationFederatedCredentialId *string `json:"ApplicationFederatedCredentialId,omitempty" xml:"ApplicationFederatedCredentialId,omitempty"`
	// The name of the application federated credential.
	//
	// example:
	//
	// test
	ApplicationFederatedCredentialName *string `json:"ApplicationFederatedCredentialName,omitempty" xml:"ApplicationFederatedCredentialName,omitempty"`
	// The type of the application federated credential.
	//
	// example:
	//
	// oidc
	ApplicationFederatedCredentialType *string `json:"ApplicationFederatedCredentialType,omitempty" xml:"ApplicationFederatedCredentialType,omitempty"`
	// The application ID.
	//
	// example:
	//
	// app_xxxasda1
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The time when the application federated credential was created.
	//
	// example:
	//
	// 1758785994982
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the application federated credential.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The federated trust source ID.
	//
	// example:
	//
	// fcp_das1asda1xxxx
	FederatedCredentialProviderId *string `json:"FederatedCredentialProviderId,omitempty" xml:"FederatedCredentialProviderId,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The time when the application federated credential was last used.
	//
	// example:
	//
	// 1758785994982
	LastUsedTime *int64 `json:"LastUsedTime,omitempty" xml:"LastUsedTime,omitempty"`
	// The OIDC structured configuration. This parameter applies to the structured mode with the OIDC type.
	OidcVerificationConfig *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsOidcVerificationConfig `json:"OidcVerificationConfig,omitempty" xml:"OidcVerificationConfig,omitempty" type:"Struct"`
	// The PKCS#7 structured configuration. This parameter applies to the structured mode with the PKCS#7 type.
	Pkcs7VerificationConfig *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsPkcs7VerificationConfig `json:"Pkcs7VerificationConfig,omitempty" xml:"Pkcs7VerificationConfig,omitempty" type:"Struct"`
	// The status of the application federated credential.
	//
	// example:
	//
	// enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the application federated credential was last updated.
	//
	// example:
	//
	// 1758785994982
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The verification condition. In freedom mode, this is a manually entered value. In structured mode, this is the final compiled value.
	VerificationCondition *string `json:"VerificationCondition,omitempty" xml:"VerificationCondition,omitempty"`
	// The verification mode. Valid values: freedom and structured.
	VerificationMode *string `json:"VerificationMode,omitempty" xml:"VerificationMode,omitempty"`
}

func (s ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials) GoString() string {
	return s.String()
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials) GetApplicationFederatedCredentialId() *string {
	return s.ApplicationFederatedCredentialId
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials) GetApplicationFederatedCredentialName() *string {
	return s.ApplicationFederatedCredentialName
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials) GetApplicationFederatedCredentialType() *string {
	return s.ApplicationFederatedCredentialType
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials) GetDescription() *string {
	return s.Description
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials) GetFederatedCredentialProviderId() *string {
	return s.FederatedCredentialProviderId
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials) GetLastUsedTime() *int64 {
	return s.LastUsedTime
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials) GetOidcVerificationConfig() *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsOidcVerificationConfig {
	return s.OidcVerificationConfig
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials) GetPkcs7VerificationConfig() *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsPkcs7VerificationConfig {
	return s.Pkcs7VerificationConfig
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials) GetStatus() *string {
	return s.Status
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials) GetVerificationCondition() *string {
	return s.VerificationCondition
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials) GetVerificationMode() *string {
	return s.VerificationMode
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials) SetApplicationFederatedCredentialId(v string) *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials {
	s.ApplicationFederatedCredentialId = &v
	return s
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials) SetApplicationFederatedCredentialName(v string) *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials {
	s.ApplicationFederatedCredentialName = &v
	return s
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials) SetApplicationFederatedCredentialType(v string) *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials {
	s.ApplicationFederatedCredentialType = &v
	return s
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials) SetApplicationId(v string) *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials {
	s.ApplicationId = &v
	return s
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials) SetCreateTime(v int64) *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials {
	s.CreateTime = &v
	return s
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials) SetDescription(v string) *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials {
	s.Description = &v
	return s
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials) SetFederatedCredentialProviderId(v string) *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials {
	s.FederatedCredentialProviderId = &v
	return s
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials) SetInstanceId(v string) *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials {
	s.InstanceId = &v
	return s
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials) SetLastUsedTime(v int64) *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials {
	s.LastUsedTime = &v
	return s
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials) SetOidcVerificationConfig(v *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsOidcVerificationConfig) *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials {
	s.OidcVerificationConfig = v
	return s
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials) SetPkcs7VerificationConfig(v *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsPkcs7VerificationConfig) *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials {
	s.Pkcs7VerificationConfig = v
	return s
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials) SetStatus(v string) *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials {
	s.Status = &v
	return s
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials) SetUpdateTime(v int64) *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials {
	s.UpdateTime = &v
	return s
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials) SetVerificationCondition(v string) *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials {
	s.VerificationCondition = &v
	return s
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials) SetVerificationMode(v string) *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials {
	s.VerificationMode = &v
	return s
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentials) Validate() error {
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

type ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsOidcVerificationConfig struct {
	// The Azure VM scenario configuration.
	AzureVmConfig *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsOidcVerificationConfigAzureVmConfig `json:"AzureVmConfig,omitempty" xml:"AzureVmConfig,omitempty" type:"Struct"`
	// The GCP VM scenario configuration.
	GcpVmConfig   *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsOidcVerificationConfigGcpVmConfig   `json:"GcpVmConfig,omitempty" xml:"GcpVmConfig,omitempty" type:"Struct"`
	GenericConfig *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsOidcVerificationConfigGenericConfig `json:"GenericConfig,omitempty" xml:"GenericConfig,omitempty" type:"Struct"`
	// The Kubernetes scenario configuration.
	KubernetesConfig *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsOidcVerificationConfigKubernetesConfig `json:"KubernetesConfig,omitempty" xml:"KubernetesConfig,omitempty" type:"Struct"`
	// The OIDC scenario profile. Valid values: generic, kubernetes, gcp_vm, and azure_vm.
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
}

func (s ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsOidcVerificationConfig) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsOidcVerificationConfig) GoString() string {
	return s.String()
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsOidcVerificationConfig) GetAzureVmConfig() *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsOidcVerificationConfigAzureVmConfig {
	return s.AzureVmConfig
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsOidcVerificationConfig) GetGcpVmConfig() *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsOidcVerificationConfigGcpVmConfig {
	return s.GcpVmConfig
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsOidcVerificationConfig) GetGenericConfig() *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsOidcVerificationConfigGenericConfig {
	return s.GenericConfig
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsOidcVerificationConfig) GetKubernetesConfig() *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsOidcVerificationConfigKubernetesConfig {
	return s.KubernetesConfig
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsOidcVerificationConfig) GetProfile() *string {
	return s.Profile
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsOidcVerificationConfig) SetAzureVmConfig(v *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsOidcVerificationConfigAzureVmConfig) *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsOidcVerificationConfig {
	s.AzureVmConfig = v
	return s
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsOidcVerificationConfig) SetGcpVmConfig(v *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsOidcVerificationConfigGcpVmConfig) *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsOidcVerificationConfig {
	s.GcpVmConfig = v
	return s
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsOidcVerificationConfig) SetGenericConfig(v *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsOidcVerificationConfigGenericConfig) *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsOidcVerificationConfig {
	s.GenericConfig = v
	return s
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsOidcVerificationConfig) SetKubernetesConfig(v *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsOidcVerificationConfigKubernetesConfig) *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsOidcVerificationConfig {
	s.KubernetesConfig = v
	return s
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsOidcVerificationConfig) SetProfile(v string) *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsOidcVerificationConfig {
	s.Profile = &v
	return s
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsOidcVerificationConfig) Validate() error {
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

type ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsOidcVerificationConfigAzureVmConfig struct {
	PrincipalId       *string   `json:"PrincipalId,omitempty" xml:"PrincipalId,omitempty"`
	ResourceGroupName *string   `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
	SubscriptionId    *string   `json:"SubscriptionId,omitempty" xml:"SubscriptionId,omitempty"`
	VmNames           []*string `json:"VmNames,omitempty" xml:"VmNames,omitempty" type:"Repeated"`
}

func (s ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsOidcVerificationConfigAzureVmConfig) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsOidcVerificationConfigAzureVmConfig) GoString() string {
	return s.String()
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsOidcVerificationConfigAzureVmConfig) GetPrincipalId() *string {
	return s.PrincipalId
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsOidcVerificationConfigAzureVmConfig) GetResourceGroupName() *string {
	return s.ResourceGroupName
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsOidcVerificationConfigAzureVmConfig) GetSubscriptionId() *string {
	return s.SubscriptionId
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsOidcVerificationConfigAzureVmConfig) GetVmNames() []*string {
	return s.VmNames
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsOidcVerificationConfigAzureVmConfig) SetPrincipalId(v string) *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsOidcVerificationConfigAzureVmConfig {
	s.PrincipalId = &v
	return s
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsOidcVerificationConfigAzureVmConfig) SetResourceGroupName(v string) *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsOidcVerificationConfigAzureVmConfig {
	s.ResourceGroupName = &v
	return s
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsOidcVerificationConfigAzureVmConfig) SetSubscriptionId(v string) *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsOidcVerificationConfigAzureVmConfig {
	s.SubscriptionId = &v
	return s
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsOidcVerificationConfigAzureVmConfig) SetVmNames(v []*string) *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsOidcVerificationConfigAzureVmConfig {
	s.VmNames = v
	return s
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsOidcVerificationConfigAzureVmConfig) Validate() error {
	return dara.Validate(s)
}

type ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsOidcVerificationConfigGcpVmConfig struct {
	// The list of VM instance IDs. A maximum of 10 IDs are supported.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	ProjectId   *string   `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The sub claim that corresponds to the service account.
	ServiceAccountId *string `json:"ServiceAccountId,omitempty" xml:"ServiceAccountId,omitempty"`
}

func (s ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsOidcVerificationConfigGcpVmConfig) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsOidcVerificationConfigGcpVmConfig) GoString() string {
	return s.String()
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsOidcVerificationConfigGcpVmConfig) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsOidcVerificationConfigGcpVmConfig) GetProjectId() *string {
	return s.ProjectId
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsOidcVerificationConfigGcpVmConfig) GetServiceAccountId() *string {
	return s.ServiceAccountId
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsOidcVerificationConfigGcpVmConfig) SetInstanceIds(v []*string) *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsOidcVerificationConfigGcpVmConfig {
	s.InstanceIds = v
	return s
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsOidcVerificationConfigGcpVmConfig) SetProjectId(v string) *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsOidcVerificationConfigGcpVmConfig {
	s.ProjectId = &v
	return s
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsOidcVerificationConfigGcpVmConfig) SetServiceAccountId(v string) *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsOidcVerificationConfigGcpVmConfig {
	s.ServiceAccountId = &v
	return s
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsOidcVerificationConfigGcpVmConfig) Validate() error {
	return dara.Validate(s)
}

type ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsOidcVerificationConfigGenericConfig struct {
	Subject *string `json:"Subject,omitempty" xml:"Subject,omitempty"`
}

func (s ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsOidcVerificationConfigGenericConfig) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsOidcVerificationConfigGenericConfig) GoString() string {
	return s.String()
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsOidcVerificationConfigGenericConfig) GetSubject() *string {
	return s.Subject
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsOidcVerificationConfigGenericConfig) SetSubject(v string) *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsOidcVerificationConfigGenericConfig {
	s.Subject = &v
	return s
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsOidcVerificationConfigGenericConfig) Validate() error {
	return dara.Validate(s)
}

type ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsOidcVerificationConfigKubernetesConfig struct {
	// The Kubernetes namespace.
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The pod name prefix.
	PodNamePrefix *string `json:"PodNamePrefix,omitempty" xml:"PodNamePrefix,omitempty"`
	// The Kubernetes service account name.
	ServiceAccountName *string `json:"ServiceAccountName,omitempty" xml:"ServiceAccountName,omitempty"`
}

func (s ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsOidcVerificationConfigKubernetesConfig) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsOidcVerificationConfigKubernetesConfig) GoString() string {
	return s.String()
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsOidcVerificationConfigKubernetesConfig) GetNamespace() *string {
	return s.Namespace
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsOidcVerificationConfigKubernetesConfig) GetPodNamePrefix() *string {
	return s.PodNamePrefix
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsOidcVerificationConfigKubernetesConfig) GetServiceAccountName() *string {
	return s.ServiceAccountName
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsOidcVerificationConfigKubernetesConfig) SetNamespace(v string) *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsOidcVerificationConfigKubernetesConfig {
	s.Namespace = &v
	return s
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsOidcVerificationConfigKubernetesConfig) SetPodNamePrefix(v string) *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsOidcVerificationConfigKubernetesConfig {
	s.PodNamePrefix = &v
	return s
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsOidcVerificationConfigKubernetesConfig) SetServiceAccountName(v string) *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsOidcVerificationConfigKubernetesConfig {
	s.ServiceAccountName = &v
	return s
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsOidcVerificationConfigKubernetesConfig) Validate() error {
	return dara.Validate(s)
}

type ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsPkcs7VerificationConfig struct {
	// The list of allowed instance IDs. A maximum of 10 IDs are supported.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
}

func (s ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsPkcs7VerificationConfig) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsPkcs7VerificationConfig) GoString() string {
	return s.String()
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsPkcs7VerificationConfig) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsPkcs7VerificationConfig) SetInstanceIds(v []*string) *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsPkcs7VerificationConfig {
	s.InstanceIds = v
	return s
}

func (s *ListApplicationFederatedCredentialsResponseBodyApplicationFederatedCredentialsPkcs7VerificationConfig) Validate() error {
	return dara.Validate(s)
}
