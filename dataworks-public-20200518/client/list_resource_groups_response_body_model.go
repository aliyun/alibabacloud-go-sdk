// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListResourceGroupsResponseBodyData) *ListResourceGroupsResponseBody
	GetData() []*ListResourceGroupsResponseBodyData
	SetHttpStatusCode(v int32) *ListResourceGroupsResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ListResourceGroupsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListResourceGroupsResponseBody
	GetSuccess() *bool
}

type ListResourceGroupsResponseBody struct {
	// The list of resource groups.
	Data []*ListResourceGroupsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0bc1411515937635973****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListResourceGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListResourceGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourceGroupsResponseBody) GetData() []*ListResourceGroupsResponseBodyData {
	return s.Data
}

func (s *ListResourceGroupsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListResourceGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListResourceGroupsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListResourceGroupsResponseBody) SetData(v []*ListResourceGroupsResponseBodyData) *ListResourceGroupsResponseBody {
	s.Data = v
	return s
}

func (s *ListResourceGroupsResponseBody) SetHttpStatusCode(v int32) *ListResourceGroupsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListResourceGroupsResponseBody) SetRequestId(v string) *ListResourceGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourceGroupsResponseBody) SetSuccess(v bool) *ListResourceGroupsResponseBody {
	s.Success = &v
	return s
}

func (s *ListResourceGroupsResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListResourceGroupsResponseBodyData struct {
	// The category of the resource group. Valid values:
	//
	// - default: public resource group.
	//
	// - single: dedicated resource group.
	//
	// example:
	//
	// default
	BizExtKey *string `json:"BizExtKey,omitempty" xml:"BizExtKey,omitempty"`
	// The name of the cluster. This parameter is valid only for MaxCompute and PAI resource group types.
	//
	// example:
	//
	// AY18G
	Cluster *string `json:"Cluster,omitempty" xml:"Cluster,omitempty"`
	// The time when the cluster was created. The format is Jul 9, 2018 2:43:37 PM.
	//
	// example:
	//
	// Jul 9, 2018 2:43:37 PM
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Indicates whether Kp (key person) access is used. Valid values:
	//
	// - true: The MaxCompute engine uses the Alibaba Cloud account UID as the display name of the access account.
	//
	// - false: The MaxCompute engine uses the Alibaba Cloud account name as the display name of the access account.
	//
	// This parameter is meaningless for other types and is valid only for the MaxCompute engine.
	//
	// example:
	//
	// false
	EnableKp *bool `json:"EnableKp,omitempty" xml:"EnableKp,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// 1234567
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The identifier of the resource group.
	//
	// example:
	//
	// e1815577-2f4e-4c5e-b29****
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	// Indicates whether the resource group is the default resource group. Valid values:
	//
	// - true: The resource group is the default resource group.
	//
	// - false: The resource group is not the default resource group.
	//
	// example:
	//
	// true
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// The type of the resource group. Valid values:
	//
	// - ISOLATE: an upfront dedicated resource group.
	//
	// - SHARE: a pay-as-you-go public resource group.
	//
	// - DEVELOP: a developer edition.
	//
	// example:
	//
	// SHARE
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The name of the resource group.
	//
	// example:
	//
	// abc
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the resource group. Valid values:
	//
	// - 0: DataWorks
	//
	// - 2: MaxCompute
	//
	// - 3: PAI
	//
	// - 4: data integration
	//
	// - 7: scheduling
	//
	// - 9: dataService
	//
	// example:
	//
	// 3
	ResourceGroupType *string `json:"ResourceGroupType,omitempty" xml:"ResourceGroupType,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmzbn7pti3zfa
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The sequence field. Used to sort created resource groups in ascending order by creation sequence number.
	//
	// example:
	//
	// 300
	Sequence *int32 `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// The detailed information of the resource group. The content displayed in {} is the detailed information of the resource group.
	//
	// example:
	//
	// {}
	Specs map[string]interface{} `json:"Specs,omitempty" xml:"Specs,omitempty"`
	// The status of the resource group. Valid values:
	//
	// - NORMAL(0): The resource group is running or in service.
	//
	// - STOP(1): The resource group has expired and is frozen.
	//
	// - DELETED(2): The resource group has been released or destroyed.
	//
	// - CREATING(3): The resource group is being created or started.
	//
	// - CREATE_FAILED(4): The resource group failed to be created or started.
	//
	// - UPDATING(5): The resource group is being scaled out or upgraded.
	//
	// - UPDATE_FAILED(6): The resource group failed to be scaled out or upgraded.
	//
	// - DELETING(7): The resource group is being released or destroyed.
	//
	// - DELETE_FAILED(8): The resource group failed to be released or destroyed.
	//
	// - TIMEOUT(9): The operation performed on the resource group timed out. All change operations may time out. This value is temporarily available only for DataService.
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The list of tags.
	Tags []*ListResourceGroupsResponseBodyDataTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The tenant ID.
	//
	// example:
	//
	// 1234567
	TenantId *int64 `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// The time when the resource group was last updated.
	//
	// The format is `MMM d, yyyy h:mm:ss a`, for example, `Jul 9, 2018 2:43:37 PM`.
	//
	// example:
	//
	// Jul 9, 2018 2:43:37 PM
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListResourceGroupsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListResourceGroupsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListResourceGroupsResponseBodyData) GetBizExtKey() *string {
	return s.BizExtKey
}

func (s *ListResourceGroupsResponseBodyData) GetCluster() *string {
	return s.Cluster
}

func (s *ListResourceGroupsResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListResourceGroupsResponseBodyData) GetEnableKp() *bool {
	return s.EnableKp
}

func (s *ListResourceGroupsResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *ListResourceGroupsResponseBodyData) GetIdentifier() *string {
	return s.Identifier
}

func (s *ListResourceGroupsResponseBodyData) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *ListResourceGroupsResponseBodyData) GetMode() *string {
	return s.Mode
}

func (s *ListResourceGroupsResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ListResourceGroupsResponseBodyData) GetResourceGroupType() *string {
	return s.ResourceGroupType
}

func (s *ListResourceGroupsResponseBodyData) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *ListResourceGroupsResponseBodyData) GetSequence() *int32 {
	return s.Sequence
}

func (s *ListResourceGroupsResponseBodyData) GetSpecs() map[string]interface{} {
	return s.Specs
}

func (s *ListResourceGroupsResponseBodyData) GetStatus() *int32 {
	return s.Status
}

func (s *ListResourceGroupsResponseBodyData) GetTags() []*ListResourceGroupsResponseBodyDataTags {
	return s.Tags
}

func (s *ListResourceGroupsResponseBodyData) GetTenantId() *int64 {
	return s.TenantId
}

func (s *ListResourceGroupsResponseBodyData) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListResourceGroupsResponseBodyData) SetBizExtKey(v string) *ListResourceGroupsResponseBodyData {
	s.BizExtKey = &v
	return s
}

func (s *ListResourceGroupsResponseBodyData) SetCluster(v string) *ListResourceGroupsResponseBodyData {
	s.Cluster = &v
	return s
}

func (s *ListResourceGroupsResponseBodyData) SetCreateTime(v string) *ListResourceGroupsResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListResourceGroupsResponseBodyData) SetEnableKp(v bool) *ListResourceGroupsResponseBodyData {
	s.EnableKp = &v
	return s
}

func (s *ListResourceGroupsResponseBodyData) SetId(v int64) *ListResourceGroupsResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListResourceGroupsResponseBodyData) SetIdentifier(v string) *ListResourceGroupsResponseBodyData {
	s.Identifier = &v
	return s
}

func (s *ListResourceGroupsResponseBodyData) SetIsDefault(v bool) *ListResourceGroupsResponseBodyData {
	s.IsDefault = &v
	return s
}

func (s *ListResourceGroupsResponseBodyData) SetMode(v string) *ListResourceGroupsResponseBodyData {
	s.Mode = &v
	return s
}

func (s *ListResourceGroupsResponseBodyData) SetName(v string) *ListResourceGroupsResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListResourceGroupsResponseBodyData) SetResourceGroupType(v string) *ListResourceGroupsResponseBodyData {
	s.ResourceGroupType = &v
	return s
}

func (s *ListResourceGroupsResponseBodyData) SetResourceManagerResourceGroupId(v string) *ListResourceGroupsResponseBodyData {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *ListResourceGroupsResponseBodyData) SetSequence(v int32) *ListResourceGroupsResponseBodyData {
	s.Sequence = &v
	return s
}

func (s *ListResourceGroupsResponseBodyData) SetSpecs(v map[string]interface{}) *ListResourceGroupsResponseBodyData {
	s.Specs = v
	return s
}

func (s *ListResourceGroupsResponseBodyData) SetStatus(v int32) *ListResourceGroupsResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListResourceGroupsResponseBodyData) SetTags(v []*ListResourceGroupsResponseBodyDataTags) *ListResourceGroupsResponseBodyData {
	s.Tags = v
	return s
}

func (s *ListResourceGroupsResponseBodyData) SetTenantId(v int64) *ListResourceGroupsResponseBodyData {
	s.TenantId = &v
	return s
}

func (s *ListResourceGroupsResponseBodyData) SetUpdateTime(v string) *ListResourceGroupsResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *ListResourceGroupsResponseBodyData) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListResourceGroupsResponseBodyDataTags struct {
	// The tag key.
	//
	// example:
	//
	// Env
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// Test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListResourceGroupsResponseBodyDataTags) String() string {
	return dara.Prettify(s)
}

func (s ListResourceGroupsResponseBodyDataTags) GoString() string {
	return s.String()
}

func (s *ListResourceGroupsResponseBodyDataTags) GetKey() *string {
	return s.Key
}

func (s *ListResourceGroupsResponseBodyDataTags) GetValue() *string {
	return s.Value
}

func (s *ListResourceGroupsResponseBodyDataTags) SetKey(v string) *ListResourceGroupsResponseBodyDataTags {
	s.Key = &v
	return s
}

func (s *ListResourceGroupsResponseBodyDataTags) SetValue(v string) *ListResourceGroupsResponseBodyDataTags {
	s.Value = &v
	return s
}

func (s *ListResourceGroupsResponseBodyDataTags) Validate() error {
	return dara.Validate(s)
}
