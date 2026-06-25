// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListResourcesResponseBody
	GetRequestId() *string
	SetResources(v []*ListResourcesResponseBodyResources) *ListResourcesResponseBody
	GetResources() []*ListResourcesResponseBodyResources
	SetTotalCount(v int64) *ListResourcesResponseBody
	GetTotalCount() *int64
}

type ListResourcesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1e195c5116124202371861018d5bde
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of resources.
	Resources []*ListResourcesResponseBodyResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
	// The total number of entries that match the filter criteria.
	//
	// example:
	//
	// 2
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListResourcesResponseBody) GetResources() []*ListResourcesResponseBodyResources {
	return s.Resources
}

func (s *ListResourcesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListResourcesResponseBody) SetRequestId(v string) *ListResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourcesResponseBody) SetResources(v []*ListResourcesResponseBodyResources) *ListResourcesResponseBody {
	s.Resources = v
	return s
}

func (s *ListResourcesResponseBody) SetTotalCount(v int64) *ListResourcesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListResourcesResponseBody) Validate() error {
	if s.Resources != nil {
		for _, item := range s.Resources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListResourcesResponseBodyResources struct {
	// The encryption details. This parameter is valid only for MaxCompute resources.
	Encryption *ListResourcesResponseBodyResourcesEncryption `json:"Encryption,omitempty" xml:"Encryption,omitempty" type:"Struct"`
	// The environment type. Valid values:
	//
	// - `dev`: development environment
	//
	// - `prod`: production environment
	//
	// example:
	//
	// prod
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// **Deprecated.*	- This parameter is no longer used.
	Executor *ListResourcesResponseBodyResourcesExecutor `json:"Executor,omitempty" xml:"Executor,omitempty" type:"Struct"`
	// The time when the resource was created. The time is displayed in UTC and is formatted in ISO 8601.
	//
	// example:
	//
	// 2021-01-21T17:12:35.232Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// The name of the resource group. The name must be unique within an Alibaba Cloud account.
	//
	// example:
	//
	// groupName
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// 123
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Indicates whether the resource is the default resource of its type. Each resource type has only one default resource. Valid values:
	//
	// - `true`: The resource is the default resource.
	//
	// - `false`: The resource is not the default resource.
	//
	// example:
	//
	// true
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// The list of labels.
	Labels []*ListResourcesResponseBodyResourcesLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The resource name.
	//
	// example:
	//
	// ResourceName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// **Deprecated.*	- This parameter is deprecated and will be removed in a future release. Use the `ResourceType` parameter instead.
	//
	// example:
	//
	// MaxCompute
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The list of quotas.
	Quotas []*ListResourcesResponseBodyResourcesQuotas `json:"Quotas,omitempty" xml:"Quotas,omitempty" type:"Repeated"`
	// The type of the resource. Valid values:
	//
	// - `MaxCompute`: MaxCompute resources
	//
	// - `ECS`: ECS resources
	//
	// - `Lingjun`: Lingjun intelligent computing resources
	//
	// - `ACS`: ACS computing resources
	//
	// - `Flink`: Flink resources
	//
	// - `SelfManagedAckPro`: self-managed cluster resources for AckPro
	//
	// - `SelfManagedAckLingjun`: self-managed cluster resources for AckLingjun
	//
	// - `SelfManagedASI`: self-managed cluster resources for ASI (third-party cloud)
	//
	// example:
	//
	// MaxCompute
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The resource specifications.
	//
	// example:
	//
	// 对于MaxCompute {"Endpoint": "odps.alibaba-inc.com", "Project": "mignshi"}
	Spec map[string]interface{} `json:"Spec,omitempty" xml:"Spec,omitempty"`
	// The ID of the workspace to which the resource belongs.
	//
	// example:
	//
	// 123
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListResourcesResponseBodyResources) String() string {
	return dara.Prettify(s)
}

func (s ListResourcesResponseBodyResources) GoString() string {
	return s.String()
}

func (s *ListResourcesResponseBodyResources) GetEncryption() *ListResourcesResponseBodyResourcesEncryption {
	return s.Encryption
}

func (s *ListResourcesResponseBodyResources) GetEnvType() *string {
	return s.EnvType
}

func (s *ListResourcesResponseBodyResources) GetExecutor() *ListResourcesResponseBodyResourcesExecutor {
	return s.Executor
}

func (s *ListResourcesResponseBodyResources) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *ListResourcesResponseBodyResources) GetGroupName() *string {
	return s.GroupName
}

func (s *ListResourcesResponseBodyResources) GetId() *string {
	return s.Id
}

func (s *ListResourcesResponseBodyResources) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *ListResourcesResponseBodyResources) GetLabels() []*ListResourcesResponseBodyResourcesLabels {
	return s.Labels
}

func (s *ListResourcesResponseBodyResources) GetName() *string {
	return s.Name
}

func (s *ListResourcesResponseBodyResources) GetProductType() *string {
	return s.ProductType
}

func (s *ListResourcesResponseBodyResources) GetQuotas() []*ListResourcesResponseBodyResourcesQuotas {
	return s.Quotas
}

func (s *ListResourcesResponseBodyResources) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListResourcesResponseBodyResources) GetSpec() map[string]interface{} {
	return s.Spec
}

func (s *ListResourcesResponseBodyResources) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListResourcesResponseBodyResources) SetEncryption(v *ListResourcesResponseBodyResourcesEncryption) *ListResourcesResponseBodyResources {
	s.Encryption = v
	return s
}

func (s *ListResourcesResponseBodyResources) SetEnvType(v string) *ListResourcesResponseBodyResources {
	s.EnvType = &v
	return s
}

func (s *ListResourcesResponseBodyResources) SetExecutor(v *ListResourcesResponseBodyResourcesExecutor) *ListResourcesResponseBodyResources {
	s.Executor = v
	return s
}

func (s *ListResourcesResponseBodyResources) SetGmtCreateTime(v string) *ListResourcesResponseBodyResources {
	s.GmtCreateTime = &v
	return s
}

func (s *ListResourcesResponseBodyResources) SetGroupName(v string) *ListResourcesResponseBodyResources {
	s.GroupName = &v
	return s
}

func (s *ListResourcesResponseBodyResources) SetId(v string) *ListResourcesResponseBodyResources {
	s.Id = &v
	return s
}

func (s *ListResourcesResponseBodyResources) SetIsDefault(v bool) *ListResourcesResponseBodyResources {
	s.IsDefault = &v
	return s
}

func (s *ListResourcesResponseBodyResources) SetLabels(v []*ListResourcesResponseBodyResourcesLabels) *ListResourcesResponseBodyResources {
	s.Labels = v
	return s
}

func (s *ListResourcesResponseBodyResources) SetName(v string) *ListResourcesResponseBodyResources {
	s.Name = &v
	return s
}

func (s *ListResourcesResponseBodyResources) SetProductType(v string) *ListResourcesResponseBodyResources {
	s.ProductType = &v
	return s
}

func (s *ListResourcesResponseBodyResources) SetQuotas(v []*ListResourcesResponseBodyResourcesQuotas) *ListResourcesResponseBodyResources {
	s.Quotas = v
	return s
}

func (s *ListResourcesResponseBodyResources) SetResourceType(v string) *ListResourcesResponseBodyResources {
	s.ResourceType = &v
	return s
}

func (s *ListResourcesResponseBodyResources) SetSpec(v map[string]interface{}) *ListResourcesResponseBodyResources {
	s.Spec = v
	return s
}

func (s *ListResourcesResponseBodyResources) SetWorkspaceId(v string) *ListResourcesResponseBodyResources {
	s.WorkspaceId = &v
	return s
}

func (s *ListResourcesResponseBodyResources) Validate() error {
	if s.Encryption != nil {
		if err := s.Encryption.Validate(); err != nil {
			return err
		}
	}
	if s.Executor != nil {
		if err := s.Executor.Validate(); err != nil {
			return err
		}
	}
	if s.Labels != nil {
		for _, item := range s.Labels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Quotas != nil {
		for _, item := range s.Quotas {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListResourcesResponseBodyResourcesEncryption struct {
	// The encryption algorithm.
	//
	// example:
	//
	// AESCTR
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// Indicates whether encryption is enabled.
	//
	// example:
	//
	// false
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The encryption key.
	//
	// example:
	//
	// DEFAULT
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
}

func (s ListResourcesResponseBodyResourcesEncryption) String() string {
	return dara.Prettify(s)
}

func (s ListResourcesResponseBodyResourcesEncryption) GoString() string {
	return s.String()
}

func (s *ListResourcesResponseBodyResourcesEncryption) GetAlgorithm() *string {
	return s.Algorithm
}

func (s *ListResourcesResponseBodyResourcesEncryption) GetEnabled() *bool {
	return s.Enabled
}

func (s *ListResourcesResponseBodyResourcesEncryption) GetKey() *string {
	return s.Key
}

func (s *ListResourcesResponseBodyResourcesEncryption) SetAlgorithm(v string) *ListResourcesResponseBodyResourcesEncryption {
	s.Algorithm = &v
	return s
}

func (s *ListResourcesResponseBodyResourcesEncryption) SetEnabled(v bool) *ListResourcesResponseBodyResourcesEncryption {
	s.Enabled = &v
	return s
}

func (s *ListResourcesResponseBodyResourcesEncryption) SetKey(v string) *ListResourcesResponseBodyResourcesEncryption {
	s.Key = &v
	return s
}

func (s *ListResourcesResponseBodyResourcesEncryption) Validate() error {
	return dara.Validate(s)
}

type ListResourcesResponseBodyResourcesExecutor struct {
	// **Deprecated.*	- This parameter is no longer used.
	//
	// example:
	//
	// 110973******7793
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s ListResourcesResponseBodyResourcesExecutor) String() string {
	return dara.Prettify(s)
}

func (s ListResourcesResponseBodyResourcesExecutor) GoString() string {
	return s.String()
}

func (s *ListResourcesResponseBodyResourcesExecutor) GetOwnerId() *string {
	return s.OwnerId
}

func (s *ListResourcesResponseBodyResourcesExecutor) SetOwnerId(v string) *ListResourcesResponseBodyResourcesExecutor {
	s.OwnerId = &v
	return s
}

func (s *ListResourcesResponseBodyResourcesExecutor) Validate() error {
	return dara.Validate(s)
}

type ListResourcesResponseBodyResourcesLabels struct {
	// The key of the label.
	//
	// example:
	//
	// system.supported.dsw
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the label.
	//
	// example:
	//
	// true
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListResourcesResponseBodyResourcesLabels) String() string {
	return dara.Prettify(s)
}

func (s ListResourcesResponseBodyResourcesLabels) GoString() string {
	return s.String()
}

func (s *ListResourcesResponseBodyResourcesLabels) GetKey() *string {
	return s.Key
}

func (s *ListResourcesResponseBodyResourcesLabels) GetValue() *string {
	return s.Value
}

func (s *ListResourcesResponseBodyResourcesLabels) SetKey(v string) *ListResourcesResponseBodyResourcesLabels {
	s.Key = &v
	return s
}

func (s *ListResourcesResponseBodyResourcesLabels) SetValue(v string) *ListResourcesResponseBodyResourcesLabels {
	s.Value = &v
	return s
}

func (s *ListResourcesResponseBodyResourcesLabels) Validate() error {
	return dara.Validate(s)
}

type ListResourcesResponseBodyResourcesQuotas struct {
	// The card type. Valid values:
	//
	// - `CPU`
	//
	// - `GPU`
	//
	// example:
	//
	// CPU
	CardType *string `json:"CardType,omitempty" xml:"CardType,omitempty"`
	// The display name of the quota.
	//
	// example:
	//
	// 默认后付费Quota
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The quota ID.
	//
	// example:
	//
	// 123
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The billing method. Valid values:
	//
	// - `isolate`: subscription
	//
	// - `share`: pay-as-you-go
	//
	// example:
	//
	// isolate
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The quota name.
	//
	// example:
	//
	// QuotaName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The product code. Valid values:
	//
	// - `PAI_isolate`: PAI subscription resource group (PAI CPU)
	//
	// - `PAI_share`: PAI pay-as-you-go resource group (PAI GPU)
	//
	// - `MaxCompute_share`: MaxCompute pay-as-you-go resource group
	//
	// - `MaxCompute_isolate`: MaxCompute subscription resource group
	//
	// - `DataWorks_isolate`: DataWorks subscription resource group
	//
	// - `DataWorks_share`: DataWorks pay-as-you-go resource group
	//
	// - `DLC_share`: DLC pay-as-you-go resource group
	//
	// example:
	//
	// MaxCompute_isolate
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The quota type. Valid values:
	//
	// - `PAI`
	//
	// - `MaxCompute`
	//
	// - `DLC`
	//
	// example:
	//
	// MaxCompute
	QuotaType *string `json:"QuotaType,omitempty" xml:"QuotaType,omitempty"`
	// The list of specifications.
	//
	// example:
	//
	// {\\"cu\\":\\"11500\\",\\"minCu\\":\\"2300\\",\\"parentId\\":\\"0\\"}
	Specs []*ListResourcesResponseBodyResourcesQuotasSpecs `json:"Specs,omitempty" xml:"Specs,omitempty" type:"Repeated"`
}

func (s ListResourcesResponseBodyResourcesQuotas) String() string {
	return dara.Prettify(s)
}

func (s ListResourcesResponseBodyResourcesQuotas) GoString() string {
	return s.String()
}

func (s *ListResourcesResponseBodyResourcesQuotas) GetCardType() *string {
	return s.CardType
}

func (s *ListResourcesResponseBodyResourcesQuotas) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListResourcesResponseBodyResourcesQuotas) GetId() *string {
	return s.Id
}

func (s *ListResourcesResponseBodyResourcesQuotas) GetMode() *string {
	return s.Mode
}

func (s *ListResourcesResponseBodyResourcesQuotas) GetName() *string {
	return s.Name
}

func (s *ListResourcesResponseBodyResourcesQuotas) GetProductCode() *string {
	return s.ProductCode
}

func (s *ListResourcesResponseBodyResourcesQuotas) GetQuotaType() *string {
	return s.QuotaType
}

func (s *ListResourcesResponseBodyResourcesQuotas) GetSpecs() []*ListResourcesResponseBodyResourcesQuotasSpecs {
	return s.Specs
}

func (s *ListResourcesResponseBodyResourcesQuotas) SetCardType(v string) *ListResourcesResponseBodyResourcesQuotas {
	s.CardType = &v
	return s
}

func (s *ListResourcesResponseBodyResourcesQuotas) SetDisplayName(v string) *ListResourcesResponseBodyResourcesQuotas {
	s.DisplayName = &v
	return s
}

func (s *ListResourcesResponseBodyResourcesQuotas) SetId(v string) *ListResourcesResponseBodyResourcesQuotas {
	s.Id = &v
	return s
}

func (s *ListResourcesResponseBodyResourcesQuotas) SetMode(v string) *ListResourcesResponseBodyResourcesQuotas {
	s.Mode = &v
	return s
}

func (s *ListResourcesResponseBodyResourcesQuotas) SetName(v string) *ListResourcesResponseBodyResourcesQuotas {
	s.Name = &v
	return s
}

func (s *ListResourcesResponseBodyResourcesQuotas) SetProductCode(v string) *ListResourcesResponseBodyResourcesQuotas {
	s.ProductCode = &v
	return s
}

func (s *ListResourcesResponseBodyResourcesQuotas) SetQuotaType(v string) *ListResourcesResponseBodyResourcesQuotas {
	s.QuotaType = &v
	return s
}

func (s *ListResourcesResponseBodyResourcesQuotas) SetSpecs(v []*ListResourcesResponseBodyResourcesQuotasSpecs) *ListResourcesResponseBodyResourcesQuotas {
	s.Specs = v
	return s
}

func (s *ListResourcesResponseBodyResourcesQuotas) Validate() error {
	if s.Specs != nil {
		for _, item := range s.Specs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListResourcesResponseBodyResourcesQuotasSpecs struct {
	// The specification name.
	//
	// example:
	//
	// cu
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The specification value.
	//
	// example:
	//
	// 11500
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListResourcesResponseBodyResourcesQuotasSpecs) String() string {
	return dara.Prettify(s)
}

func (s ListResourcesResponseBodyResourcesQuotasSpecs) GoString() string {
	return s.String()
}

func (s *ListResourcesResponseBodyResourcesQuotasSpecs) GetName() *string {
	return s.Name
}

func (s *ListResourcesResponseBodyResourcesQuotasSpecs) GetValue() *string {
	return s.Value
}

func (s *ListResourcesResponseBodyResourcesQuotasSpecs) SetName(v string) *ListResourcesResponseBodyResourcesQuotasSpecs {
	s.Name = &v
	return s
}

func (s *ListResourcesResponseBodyResourcesQuotasSpecs) SetValue(v string) *ListResourcesResponseBodyResourcesQuotasSpecs {
	s.Value = &v
	return s
}

func (s *ListResourcesResponseBodyResourcesQuotasSpecs) Validate() error {
	return dara.Validate(s)
}
