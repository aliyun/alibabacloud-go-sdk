// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSynchronizationJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *DescribeSynchronizationJobsRequest
	GetAccountId() *string
	SetClientToken(v string) *DescribeSynchronizationJobsRequest
	GetClientToken() *string
	SetOwnerId(v string) *DescribeSynchronizationJobsRequest
	GetOwnerId() *string
	SetPageNum(v int32) *DescribeSynchronizationJobsRequest
	GetPageNum() *int32
	SetPageSize(v int32) *DescribeSynchronizationJobsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeSynchronizationJobsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeSynchronizationJobsRequest
	GetResourceGroupId() *string
	SetSynchronizationJobName(v string) *DescribeSynchronizationJobsRequest
	GetSynchronizationJobName() *string
	SetTag(v []*DescribeSynchronizationJobsRequestTag) *DescribeSynchronizationJobsRequest
	GetTag() []*DescribeSynchronizationJobsRequestTag
}

type DescribeSynchronizationJobsRequest struct {
	// The ID of the Alibaba Cloud account. You do not need to specify this parameter because this parameter will be removed in the future.
	//
	// example:
	//
	// 12323344****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that the value is unique among different requests. The **ClientToken*	- value can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe63****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerId     *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return. The value must be an integer that is greater than **0*	- and does not exceed the maximum value of the Integer data type. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries to return on each page. Valid values: **30**, **50**, and **100**. Default value: **30**.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region where the data synchronization instance resides. For more information, see [List of supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The name of the data synchronization task.
	//
	// >  Fuzzy matching is supported.
	//
	// example:
	//
	// dtstest
	SynchronizationJobName *string `json:"SynchronizationJobName,omitempty" xml:"SynchronizationJobName,omitempty"`
	// The tag key. You can call the [ListTagResources](https://help.aliyun.com/document_detail/191187.html) operation to query the tag key.
	//
	// >
	//
	// 	- N specifies the serial number of the tag. For example, Tag.1.Key specifies the key of the first tag and Tag.2.Key specifies the key of the second tag. You can specify 1 to 20 tag keys at a time.
	//
	// 	- This parameter cannot be an empty string.
	Tag []*DescribeSynchronizationJobsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeSynchronizationJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSynchronizationJobsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobsRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *DescribeSynchronizationJobsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeSynchronizationJobsRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *DescribeSynchronizationJobsRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *DescribeSynchronizationJobsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSynchronizationJobsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSynchronizationJobsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeSynchronizationJobsRequest) GetSynchronizationJobName() *string {
	return s.SynchronizationJobName
}

func (s *DescribeSynchronizationJobsRequest) GetTag() []*DescribeSynchronizationJobsRequestTag {
	return s.Tag
}

func (s *DescribeSynchronizationJobsRequest) SetAccountId(v string) *DescribeSynchronizationJobsRequest {
	s.AccountId = &v
	return s
}

func (s *DescribeSynchronizationJobsRequest) SetClientToken(v string) *DescribeSynchronizationJobsRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeSynchronizationJobsRequest) SetOwnerId(v string) *DescribeSynchronizationJobsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSynchronizationJobsRequest) SetPageNum(v int32) *DescribeSynchronizationJobsRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeSynchronizationJobsRequest) SetPageSize(v int32) *DescribeSynchronizationJobsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSynchronizationJobsRequest) SetRegionId(v string) *DescribeSynchronizationJobsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSynchronizationJobsRequest) SetResourceGroupId(v string) *DescribeSynchronizationJobsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeSynchronizationJobsRequest) SetSynchronizationJobName(v string) *DescribeSynchronizationJobsRequest {
	s.SynchronizationJobName = &v
	return s
}

func (s *DescribeSynchronizationJobsRequest) SetTag(v []*DescribeSynchronizationJobsRequestTag) *DescribeSynchronizationJobsRequest {
	s.Tag = v
	return s
}

func (s *DescribeSynchronizationJobsRequest) Validate() error {
	return dara.Validate(s)
}

type DescribeSynchronizationJobsRequestTag struct {
	// The tag key. You can call the [ListTagResources](https://help.aliyun.com/document_detail/191187.html) operation to query the tag key.
	//
	// >
	//
	// 	- N specifies the serial number of the tag. For example, Tag.1.Key specifies the key of the first tag and Tag.2.Key specifies the key of the second tag. You can specify 1 to 20 tag keys at a time.
	//
	// 	- This parameter cannot be an empty string.
	//
	// This parameter is required.
	//
	// example:
	//
	// testkey1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. You can call the [ListTagResources](https://help.aliyun.com/document_detail/191187.html) operation to query the tag value.
	//
	// >
	//
	// 	- N specifies the serial number of the tag. For example, Tag.1.Value specifies the value of the first tag and Tag.2.Value specifies the value of the second tag. You can specify 1 to 20 tag values at a time.
	//
	// 	- This parameter can be an empty string.
	//
	// example:
	//
	// testvalue1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeSynchronizationJobsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeSynchronizationJobsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobsRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeSynchronizationJobsRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeSynchronizationJobsRequestTag) SetKey(v string) *DescribeSynchronizationJobsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeSynchronizationJobsRequestTag) SetValue(v string) *DescribeSynchronizationJobsRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeSynchronizationJobsRequestTag) Validate() error {
	return dara.Validate(s)
}
