// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationFederatedCredentialsForProviderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationFederatedCredentials(v []*ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials) *ListApplicationFederatedCredentialsForProviderResponseBody
	GetApplicationFederatedCredentials() []*ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials
	SetMaxResults(v int32) *ListApplicationFederatedCredentialsForProviderResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListApplicationFederatedCredentialsForProviderResponseBody
	GetNextToken() *string
	SetPreviousToken(v string) *ListApplicationFederatedCredentialsForProviderResponseBody
	GetPreviousToken() *string
	SetRequestId(v string) *ListApplicationFederatedCredentialsForProviderResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListApplicationFederatedCredentialsForProviderResponseBody
	GetTotalCount() *int32
}

type ListApplicationFederatedCredentialsForProviderResponseBody struct {
	// The list of application federated credentials.
	ApplicationFederatedCredentials []*ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials `json:"ApplicationFederatedCredentials,omitempty" xml:"ApplicationFederatedCredentials,omitempty" type:"Repeated"`
	// The maximum number of entries returned per page in a paged query. This parameter is used for paging.
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
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListApplicationFederatedCredentialsForProviderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationFederatedCredentialsForProviderResponseBody) GoString() string {
	return s.String()
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBody) GetApplicationFederatedCredentials() []*ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials {
	return s.ApplicationFederatedCredentials
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBody) GetPreviousToken() *string {
	return s.PreviousToken
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBody) SetApplicationFederatedCredentials(v []*ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials) *ListApplicationFederatedCredentialsForProviderResponseBody {
	s.ApplicationFederatedCredentials = v
	return s
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBody) SetMaxResults(v int32) *ListApplicationFederatedCredentialsForProviderResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBody) SetNextToken(v string) *ListApplicationFederatedCredentialsForProviderResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBody) SetPreviousToken(v string) *ListApplicationFederatedCredentialsForProviderResponseBody {
	s.PreviousToken = &v
	return s
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBody) SetRequestId(v string) *ListApplicationFederatedCredentialsForProviderResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBody) SetTotalCount(v int32) *ListApplicationFederatedCredentialsForProviderResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBody) Validate() error {
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

type ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials struct {
	// The application federated credential ID.
	//
	// example:
	//
	// afc_dads12sadxxxxx
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
	// app_asda1dsadxxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The time when the credential was created.
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
	// fcp_adasd12dxxxxx
	FederatedCredentialProviderId *string `json:"FederatedCredentialProviderId,omitempty" xml:"FederatedCredentialProviderId,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The time when the credential was last used.
	//
	// example:
	//
	// 1758785994982
	LastUsedTime *int64 `json:"LastUsedTime,omitempty" xml:"LastUsedTime,omitempty"`
	// The OIDC structured configuration. This applies to structured mode with the OIDC type.
	OidcVerificationConfig *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsOidcVerificationConfig `json:"OidcVerificationConfig,omitempty" xml:"OidcVerificationConfig,omitempty" type:"Struct"`
	// The PKCS#7 structured configuration. This applies to structured mode with the PKCS#7 type.
	Pkcs7VerificationConfig *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsPkcs7VerificationConfig `json:"Pkcs7VerificationConfig,omitempty" xml:"Pkcs7VerificationConfig,omitempty" type:"Struct"`
	// The application federated credential status.
	//
	// example:
	//
	// enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the credential was last updated.
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

func (s ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials) GoString() string {
	return s.String()
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials) GetApplicationFederatedCredentialId() *string {
	return s.ApplicationFederatedCredentialId
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials) GetApplicationFederatedCredentialName() *string {
	return s.ApplicationFederatedCredentialName
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials) GetApplicationFederatedCredentialType() *string {
	return s.ApplicationFederatedCredentialType
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials) GetDescription() *string {
	return s.Description
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials) GetFederatedCredentialProviderId() *string {
	return s.FederatedCredentialProviderId
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials) GetLastUsedTime() *int64 {
	return s.LastUsedTime
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials) GetOidcVerificationConfig() *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsOidcVerificationConfig {
	return s.OidcVerificationConfig
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials) GetPkcs7VerificationConfig() *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsPkcs7VerificationConfig {
	return s.Pkcs7VerificationConfig
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials) GetStatus() *string {
	return s.Status
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials) GetVerificationCondition() *string {
	return s.VerificationCondition
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials) GetVerificationMode() *string {
	return s.VerificationMode
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials) SetApplicationFederatedCredentialId(v string) *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials {
	s.ApplicationFederatedCredentialId = &v
	return s
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials) SetApplicationFederatedCredentialName(v string) *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials {
	s.ApplicationFederatedCredentialName = &v
	return s
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials) SetApplicationFederatedCredentialType(v string) *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials {
	s.ApplicationFederatedCredentialType = &v
	return s
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials) SetApplicationId(v string) *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials {
	s.ApplicationId = &v
	return s
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials) SetCreateTime(v int64) *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials {
	s.CreateTime = &v
	return s
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials) SetDescription(v string) *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials {
	s.Description = &v
	return s
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials) SetFederatedCredentialProviderId(v string) *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials {
	s.FederatedCredentialProviderId = &v
	return s
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials) SetInstanceId(v string) *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials {
	s.InstanceId = &v
	return s
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials) SetLastUsedTime(v int64) *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials {
	s.LastUsedTime = &v
	return s
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials) SetOidcVerificationConfig(v *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsOidcVerificationConfig) *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials {
	s.OidcVerificationConfig = v
	return s
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials) SetPkcs7VerificationConfig(v *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsPkcs7VerificationConfig) *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials {
	s.Pkcs7VerificationConfig = v
	return s
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials) SetStatus(v string) *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials {
	s.Status = &v
	return s
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials) SetUpdateTime(v int64) *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials {
	s.UpdateTime = &v
	return s
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials) SetVerificationCondition(v string) *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials {
	s.VerificationCondition = &v
	return s
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials) SetVerificationMode(v string) *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials {
	s.VerificationMode = &v
	return s
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentials) Validate() error {
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

type ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsOidcVerificationConfig struct {
	// The Azure VM scenario configuration.
	AzureVmConfig *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsOidcVerificationConfigAzureVmConfig `json:"AzureVmConfig,omitempty" xml:"AzureVmConfig,omitempty" type:"Struct"`
	// The GCP VM scenario configuration.
	GcpVmConfig   *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsOidcVerificationConfigGcpVmConfig   `json:"GcpVmConfig,omitempty" xml:"GcpVmConfig,omitempty" type:"Struct"`
	GenericConfig *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsOidcVerificationConfigGenericConfig `json:"GenericConfig,omitempty" xml:"GenericConfig,omitempty" type:"Struct"`
	// The Kubernetes scenario configuration.
	KubernetesConfig *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsOidcVerificationConfigKubernetesConfig `json:"KubernetesConfig,omitempty" xml:"KubernetesConfig,omitempty" type:"Struct"`
	// The OIDC scenario profile. Valid values: generic, kubernetes, gcp_vm, and azure_vm.
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
}

func (s ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsOidcVerificationConfig) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsOidcVerificationConfig) GoString() string {
	return s.String()
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsOidcVerificationConfig) GetAzureVmConfig() *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsOidcVerificationConfigAzureVmConfig {
	return s.AzureVmConfig
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsOidcVerificationConfig) GetGcpVmConfig() *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsOidcVerificationConfigGcpVmConfig {
	return s.GcpVmConfig
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsOidcVerificationConfig) GetGenericConfig() *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsOidcVerificationConfigGenericConfig {
	return s.GenericConfig
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsOidcVerificationConfig) GetKubernetesConfig() *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsOidcVerificationConfigKubernetesConfig {
	return s.KubernetesConfig
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsOidcVerificationConfig) GetProfile() *string {
	return s.Profile
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsOidcVerificationConfig) SetAzureVmConfig(v *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsOidcVerificationConfigAzureVmConfig) *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsOidcVerificationConfig {
	s.AzureVmConfig = v
	return s
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsOidcVerificationConfig) SetGcpVmConfig(v *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsOidcVerificationConfigGcpVmConfig) *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsOidcVerificationConfig {
	s.GcpVmConfig = v
	return s
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsOidcVerificationConfig) SetGenericConfig(v *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsOidcVerificationConfigGenericConfig) *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsOidcVerificationConfig {
	s.GenericConfig = v
	return s
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsOidcVerificationConfig) SetKubernetesConfig(v *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsOidcVerificationConfigKubernetesConfig) *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsOidcVerificationConfig {
	s.KubernetesConfig = v
	return s
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsOidcVerificationConfig) SetProfile(v string) *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsOidcVerificationConfig {
	s.Profile = &v
	return s
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsOidcVerificationConfig) Validate() error {
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

type ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsOidcVerificationConfigAzureVmConfig struct {
	PrincipalId       *string   `json:"PrincipalId,omitempty" xml:"PrincipalId,omitempty"`
	ResourceGroupName *string   `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
	SubscriptionId    *string   `json:"SubscriptionId,omitempty" xml:"SubscriptionId,omitempty"`
	VmNames           []*string `json:"VmNames,omitempty" xml:"VmNames,omitempty" type:"Repeated"`
}

func (s ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsOidcVerificationConfigAzureVmConfig) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsOidcVerificationConfigAzureVmConfig) GoString() string {
	return s.String()
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsOidcVerificationConfigAzureVmConfig) GetPrincipalId() *string {
	return s.PrincipalId
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsOidcVerificationConfigAzureVmConfig) GetResourceGroupName() *string {
	return s.ResourceGroupName
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsOidcVerificationConfigAzureVmConfig) GetSubscriptionId() *string {
	return s.SubscriptionId
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsOidcVerificationConfigAzureVmConfig) GetVmNames() []*string {
	return s.VmNames
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsOidcVerificationConfigAzureVmConfig) SetPrincipalId(v string) *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsOidcVerificationConfigAzureVmConfig {
	s.PrincipalId = &v
	return s
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsOidcVerificationConfigAzureVmConfig) SetResourceGroupName(v string) *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsOidcVerificationConfigAzureVmConfig {
	s.ResourceGroupName = &v
	return s
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsOidcVerificationConfigAzureVmConfig) SetSubscriptionId(v string) *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsOidcVerificationConfigAzureVmConfig {
	s.SubscriptionId = &v
	return s
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsOidcVerificationConfigAzureVmConfig) SetVmNames(v []*string) *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsOidcVerificationConfigAzureVmConfig {
	s.VmNames = v
	return s
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsOidcVerificationConfigAzureVmConfig) Validate() error {
	return dara.Validate(s)
}

type ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsOidcVerificationConfigGcpVmConfig struct {
	// The list of VM instance IDs. A maximum of 10 IDs are supported.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	ProjectId   *string   `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The sub claim that corresponds to the service account.
	ServiceAccountId *string `json:"ServiceAccountId,omitempty" xml:"ServiceAccountId,omitempty"`
}

func (s ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsOidcVerificationConfigGcpVmConfig) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsOidcVerificationConfigGcpVmConfig) GoString() string {
	return s.String()
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsOidcVerificationConfigGcpVmConfig) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsOidcVerificationConfigGcpVmConfig) GetProjectId() *string {
	return s.ProjectId
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsOidcVerificationConfigGcpVmConfig) GetServiceAccountId() *string {
	return s.ServiceAccountId
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsOidcVerificationConfigGcpVmConfig) SetInstanceIds(v []*string) *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsOidcVerificationConfigGcpVmConfig {
	s.InstanceIds = v
	return s
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsOidcVerificationConfigGcpVmConfig) SetProjectId(v string) *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsOidcVerificationConfigGcpVmConfig {
	s.ProjectId = &v
	return s
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsOidcVerificationConfigGcpVmConfig) SetServiceAccountId(v string) *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsOidcVerificationConfigGcpVmConfig {
	s.ServiceAccountId = &v
	return s
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsOidcVerificationConfigGcpVmConfig) Validate() error {
	return dara.Validate(s)
}

type ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsOidcVerificationConfigGenericConfig struct {
	Subject *string `json:"Subject,omitempty" xml:"Subject,omitempty"`
}

func (s ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsOidcVerificationConfigGenericConfig) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsOidcVerificationConfigGenericConfig) GoString() string {
	return s.String()
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsOidcVerificationConfigGenericConfig) GetSubject() *string {
	return s.Subject
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsOidcVerificationConfigGenericConfig) SetSubject(v string) *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsOidcVerificationConfigGenericConfig {
	s.Subject = &v
	return s
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsOidcVerificationConfigGenericConfig) Validate() error {
	return dara.Validate(s)
}

type ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsOidcVerificationConfigKubernetesConfig struct {
	// The Kubernetes namespace.
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The pod name prefix.
	PodNamePrefix *string `json:"PodNamePrefix,omitempty" xml:"PodNamePrefix,omitempty"`
	// The Kubernetes service account name.
	ServiceAccountName *string `json:"ServiceAccountName,omitempty" xml:"ServiceAccountName,omitempty"`
}

func (s ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsOidcVerificationConfigKubernetesConfig) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsOidcVerificationConfigKubernetesConfig) GoString() string {
	return s.String()
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsOidcVerificationConfigKubernetesConfig) GetNamespace() *string {
	return s.Namespace
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsOidcVerificationConfigKubernetesConfig) GetPodNamePrefix() *string {
	return s.PodNamePrefix
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsOidcVerificationConfigKubernetesConfig) GetServiceAccountName() *string {
	return s.ServiceAccountName
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsOidcVerificationConfigKubernetesConfig) SetNamespace(v string) *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsOidcVerificationConfigKubernetesConfig {
	s.Namespace = &v
	return s
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsOidcVerificationConfigKubernetesConfig) SetPodNamePrefix(v string) *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsOidcVerificationConfigKubernetesConfig {
	s.PodNamePrefix = &v
	return s
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsOidcVerificationConfigKubernetesConfig) SetServiceAccountName(v string) *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsOidcVerificationConfigKubernetesConfig {
	s.ServiceAccountName = &v
	return s
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsOidcVerificationConfigKubernetesConfig) Validate() error {
	return dara.Validate(s)
}

type ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsPkcs7VerificationConfig struct {
	// The list of allowed instance IDs. A maximum of 10 IDs are supported.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
}

func (s ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsPkcs7VerificationConfig) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsPkcs7VerificationConfig) GoString() string {
	return s.String()
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsPkcs7VerificationConfig) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsPkcs7VerificationConfig) SetInstanceIds(v []*string) *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsPkcs7VerificationConfig {
	s.InstanceIds = v
	return s
}

func (s *ListApplicationFederatedCredentialsForProviderResponseBodyApplicationFederatedCredentialsPkcs7VerificationConfig) Validate() error {
	return dara.Validate(s)
}
