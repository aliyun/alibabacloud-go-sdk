// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListInstancesResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListInstancesResponseBody
	GetHttpStatusCode() *int32
	SetInstances(v []*ListInstancesResponseBodyInstances) *ListInstancesResponseBody
	GetInstances() []*ListInstancesResponseBodyInstances
	SetMessage(v string) *ListInstancesResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *ListInstancesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListInstancesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListInstancesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListInstancesResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListInstancesResponseBody
	GetTotalCount() *int32
}

type ListInstancesResponseBody struct {
	// The response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// List of business instances.
	Instances []*ListInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// The API message.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The page size for paging.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation succeeded.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of instances.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListInstancesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListInstancesResponseBody) GetInstances() []*ListInstancesResponseBodyInstances {
	return s.Instances
}

func (s *ListInstancesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListInstancesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListInstancesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstancesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListInstancesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListInstancesResponseBody) SetCode(v string) *ListInstancesResponseBody {
	s.Code = &v
	return s
}

func (s *ListInstancesResponseBody) SetHttpStatusCode(v int32) *ListInstancesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListInstancesResponseBody) SetInstances(v []*ListInstancesResponseBodyInstances) *ListInstancesResponseBody {
	s.Instances = v
	return s
}

func (s *ListInstancesResponseBody) SetMessage(v string) *ListInstancesResponseBody {
	s.Message = &v
	return s
}

func (s *ListInstancesResponseBody) SetPageNumber(v int32) *ListInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListInstancesResponseBody) SetPageSize(v int32) *ListInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListInstancesResponseBody) SetRequestId(v string) *ListInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstancesResponseBody) SetSuccess(v bool) *ListInstancesResponseBody {
	s.Success = &v
	return s
}

func (s *ListInstancesResponseBody) SetTotalCount(v int32) *ListInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListInstancesResponseBody) Validate() error {
	if s.Instances != nil {
		for _, item := range s.Instances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListInstancesResponseBodyInstances struct {
	// Creation Time.
	//
	// example:
	//
	// 1578469042851
	CreationTime *int64 `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The creator ID.
	//
	// example:
	//
	// 34234
	CreatorId *int64 `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	// The creator name.
	//
	// example:
	//
	// xxx
	CreatorName *string `json:"CreatorName,omitempty" xml:"CreatorName,omitempty"`
	// The description of the Intelligent Outbound Call business instance.
	//
	// example:
	//
	// 这是一个实例
	InstanceDescription *string `json:"InstanceDescription,omitempty" xml:"InstanceDescription,omitempty"`
	// Business instance ID.
	//
	// example:
	//
	// 90515b5-6115-4ccf-83e2-52d5bfaf2ddf
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Name of the Intelligent Outbound Calling business instance.
	//
	// example:
	//
	// 回访
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// Indicates whether the instance is a system preset instance.
	//
	// example:
	//
	// false
	IsPreset *bool `json:"IsPreset,omitempty" xml:"IsPreset,omitempty"`
	// Maximum allowed outbound call concurrency.
	//
	// example:
	//
	// 10
	MaxConcurrentConversation *int32 `json:"MaxConcurrentConversation,omitempty" xml:"MaxConcurrentConversation,omitempty"`
	// Account name.
	//
	// example:
	//
	// xxx
	OwnerName *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	// Resource group ID.
	//
	// example:
	//
	// 90515b5-6115-4ccf-83e2-52d5bfaf2ddf
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Tag information.
	//
	// example:
	//
	// []
	ResourceTags []*ListInstancesResponseBodyInstancesResourceTags `json:"ResourceTags,omitempty" xml:"ResourceTags,omitempty" type:"Repeated"`
}

func (s ListInstancesResponseBodyInstances) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstances) GetCreationTime() *int64 {
	return s.CreationTime
}

func (s *ListInstancesResponseBodyInstances) GetCreatorId() *int64 {
	return s.CreatorId
}

func (s *ListInstancesResponseBodyInstances) GetCreatorName() *string {
	return s.CreatorName
}

func (s *ListInstancesResponseBodyInstances) GetInstanceDescription() *string {
	return s.InstanceDescription
}

func (s *ListInstancesResponseBodyInstances) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListInstancesResponseBodyInstances) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListInstancesResponseBodyInstances) GetIsPreset() *bool {
	return s.IsPreset
}

func (s *ListInstancesResponseBodyInstances) GetMaxConcurrentConversation() *int32 {
	return s.MaxConcurrentConversation
}

func (s *ListInstancesResponseBodyInstances) GetOwnerName() *string {
	return s.OwnerName
}

func (s *ListInstancesResponseBodyInstances) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListInstancesResponseBodyInstances) GetResourceTags() []*ListInstancesResponseBodyInstancesResourceTags {
	return s.ResourceTags
}

func (s *ListInstancesResponseBodyInstances) SetCreationTime(v int64) *ListInstancesResponseBodyInstances {
	s.CreationTime = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetCreatorId(v int64) *ListInstancesResponseBodyInstances {
	s.CreatorId = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetCreatorName(v string) *ListInstancesResponseBodyInstances {
	s.CreatorName = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetInstanceDescription(v string) *ListInstancesResponseBodyInstances {
	s.InstanceDescription = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetInstanceId(v string) *ListInstancesResponseBodyInstances {
	s.InstanceId = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetInstanceName(v string) *ListInstancesResponseBodyInstances {
	s.InstanceName = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetIsPreset(v bool) *ListInstancesResponseBodyInstances {
	s.IsPreset = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetMaxConcurrentConversation(v int32) *ListInstancesResponseBodyInstances {
	s.MaxConcurrentConversation = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetOwnerName(v string) *ListInstancesResponseBodyInstances {
	s.OwnerName = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetResourceGroupId(v string) *ListInstancesResponseBodyInstances {
	s.ResourceGroupId = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetResourceTags(v []*ListInstancesResponseBodyInstancesResourceTags) *ListInstancesResponseBodyInstances {
	s.ResourceTags = v
	return s
}

func (s *ListInstancesResponseBodyInstances) Validate() error {
	if s.ResourceTags != nil {
		for _, item := range s.ResourceTags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListInstancesResponseBodyInstancesResourceTags struct {
	// Tag key.
	//
	// example:
	//
	// age
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// Tag value.
	//
	// example:
	//
	// 20
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListInstancesResponseBodyInstancesResourceTags) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyInstancesResourceTags) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstancesResourceTags) GetKey() *string {
	return s.Key
}

func (s *ListInstancesResponseBodyInstancesResourceTags) GetValue() *string {
	return s.Value
}

func (s *ListInstancesResponseBodyInstancesResourceTags) SetKey(v string) *ListInstancesResponseBodyInstancesResourceTags {
	s.Key = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesResourceTags) SetValue(v string) *ListInstancesResponseBodyInstancesResourceTags {
	s.Value = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesResourceTags) Validate() error {
	return dara.Validate(s)
}
