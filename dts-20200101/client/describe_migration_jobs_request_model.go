// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMigrationJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *DescribeMigrationJobsRequest
	GetAccountId() *string
	SetMigrationJobName(v string) *DescribeMigrationJobsRequest
	GetMigrationJobName() *string
	SetOwnerId(v string) *DescribeMigrationJobsRequest
	GetOwnerId() *string
	SetPageNum(v int32) *DescribeMigrationJobsRequest
	GetPageNum() *int32
	SetPageSize(v int32) *DescribeMigrationJobsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeMigrationJobsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeMigrationJobsRequest
	GetResourceGroupId() *string
	SetTag(v []*DescribeMigrationJobsRequestTag) *DescribeMigrationJobsRequest
	GetTag() []*DescribeMigrationJobsRequestTag
}

type DescribeMigrationJobsRequest struct {
	// The ID of the Alibaba Cloud account. You do not need to specify this parameter because this parameter will be removed in the future.
	//
	// example:
	//
	// 12323344****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The name of the data migration task.
	//
	// >  Fuzzy match is supported.
	//
	// example:
	//
	// MySQL迁移
	MigrationJobName *string `json:"MigrationJobName,omitempty" xml:"MigrationJobName,omitempty"`
	OwnerId          *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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
	// The ID of the region where the data migration instances reside. For more information, see [List of supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource group ID.
	//
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tag of the data migration instance, used as a filter. When this is not empty, only the instance tasks with this tag will be returned.
	Tag []*DescribeMigrationJobsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeMigrationJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMigrationJobsRequest) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobsRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *DescribeMigrationJobsRequest) GetMigrationJobName() *string {
	return s.MigrationJobName
}

func (s *DescribeMigrationJobsRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *DescribeMigrationJobsRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *DescribeMigrationJobsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeMigrationJobsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeMigrationJobsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeMigrationJobsRequest) GetTag() []*DescribeMigrationJobsRequestTag {
	return s.Tag
}

func (s *DescribeMigrationJobsRequest) SetAccountId(v string) *DescribeMigrationJobsRequest {
	s.AccountId = &v
	return s
}

func (s *DescribeMigrationJobsRequest) SetMigrationJobName(v string) *DescribeMigrationJobsRequest {
	s.MigrationJobName = &v
	return s
}

func (s *DescribeMigrationJobsRequest) SetOwnerId(v string) *DescribeMigrationJobsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeMigrationJobsRequest) SetPageNum(v int32) *DescribeMigrationJobsRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeMigrationJobsRequest) SetPageSize(v int32) *DescribeMigrationJobsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeMigrationJobsRequest) SetRegionId(v string) *DescribeMigrationJobsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeMigrationJobsRequest) SetResourceGroupId(v string) *DescribeMigrationJobsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeMigrationJobsRequest) SetTag(v []*DescribeMigrationJobsRequestTag) *DescribeMigrationJobsRequest {
	s.Tag = v
	return s
}

func (s *DescribeMigrationJobsRequest) Validate() error {
	return dara.Validate(s)
}

type DescribeMigrationJobsRequestTag struct {
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

func (s DescribeMigrationJobsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeMigrationJobsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobsRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeMigrationJobsRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeMigrationJobsRequestTag) SetKey(v string) *DescribeMigrationJobsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeMigrationJobsRequestTag) SetValue(v string) *DescribeMigrationJobsRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeMigrationJobsRequestTag) Validate() error {
	return dara.Validate(s)
}
