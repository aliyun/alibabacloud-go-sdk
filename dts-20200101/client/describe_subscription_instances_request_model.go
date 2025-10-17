// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSubscriptionInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *DescribeSubscriptionInstancesRequest
	GetAccountId() *string
	SetClientToken(v string) *DescribeSubscriptionInstancesRequest
	GetClientToken() *string
	SetOwnerId(v string) *DescribeSubscriptionInstancesRequest
	GetOwnerId() *string
	SetPageNum(v int32) *DescribeSubscriptionInstancesRequest
	GetPageNum() *int32
	SetPageSize(v int32) *DescribeSubscriptionInstancesRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeSubscriptionInstancesRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeSubscriptionInstancesRequest
	GetResourceGroupId() *string
	SetSubscriptionInstanceName(v string) *DescribeSubscriptionInstancesRequest
	GetSubscriptionInstanceName() *string
	SetTag(v []*DescribeSubscriptionInstancesRequestTag) *DescribeSubscriptionInstancesRequest
	GetTag() []*DescribeSubscriptionInstancesRequestTag
}

type DescribeSubscriptionInstancesRequest struct {
	// The ID of the Alibaba Cloud account. You do not need to specify this parameter because this parameter will be removed in the future.
	//
	// example:
	//
	// 12323344****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that it is unique among different requests. The **ClientToken*	- parameter can contain only ASCII characters and cannot exceed 64 characters in length.
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
	// The ID of the region where the change tracking instance resides. For more information, see [List of supported regions](https://help.aliyun.com/document_detail/49442.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource group ID.
	//
	// example:
	//
	// rg-aekz4us4iruleja
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The name of the change tracking instance.
	//
	// >  If you specify this parameter, DTS returns all the change tracking instances that match the specified name.
	//
	// example:
	//
	// MySQL订阅
	SubscriptionInstanceName *string `json:"SubscriptionInstanceName,omitempty" xml:"SubscriptionInstanceName,omitempty"`
	// Tags of the data migration instance, used as a filter. When this is not empty, only instances with this tag will be returned.
	Tag []*DescribeSubscriptionInstancesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeSubscriptionInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSubscriptionInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionInstancesRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *DescribeSubscriptionInstancesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeSubscriptionInstancesRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *DescribeSubscriptionInstancesRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *DescribeSubscriptionInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSubscriptionInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSubscriptionInstancesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeSubscriptionInstancesRequest) GetSubscriptionInstanceName() *string {
	return s.SubscriptionInstanceName
}

func (s *DescribeSubscriptionInstancesRequest) GetTag() []*DescribeSubscriptionInstancesRequestTag {
	return s.Tag
}

func (s *DescribeSubscriptionInstancesRequest) SetAccountId(v string) *DescribeSubscriptionInstancesRequest {
	s.AccountId = &v
	return s
}

func (s *DescribeSubscriptionInstancesRequest) SetClientToken(v string) *DescribeSubscriptionInstancesRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeSubscriptionInstancesRequest) SetOwnerId(v string) *DescribeSubscriptionInstancesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSubscriptionInstancesRequest) SetPageNum(v int32) *DescribeSubscriptionInstancesRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeSubscriptionInstancesRequest) SetPageSize(v int32) *DescribeSubscriptionInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSubscriptionInstancesRequest) SetRegionId(v string) *DescribeSubscriptionInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSubscriptionInstancesRequest) SetResourceGroupId(v string) *DescribeSubscriptionInstancesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeSubscriptionInstancesRequest) SetSubscriptionInstanceName(v string) *DescribeSubscriptionInstancesRequest {
	s.SubscriptionInstanceName = &v
	return s
}

func (s *DescribeSubscriptionInstancesRequest) SetTag(v []*DescribeSubscriptionInstancesRequestTag) *DescribeSubscriptionInstancesRequest {
	s.Tag = v
	return s
}

func (s *DescribeSubscriptionInstancesRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSubscriptionInstancesRequestTag struct {
	// The tag key.
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
	// The tag value.
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

func (s DescribeSubscriptionInstancesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeSubscriptionInstancesRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionInstancesRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeSubscriptionInstancesRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeSubscriptionInstancesRequestTag) SetKey(v string) *DescribeSubscriptionInstancesRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeSubscriptionInstancesRequestTag) SetValue(v string) *DescribeSubscriptionInstancesRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeSubscriptionInstancesRequestTag) Validate() error {
	return dara.Validate(s)
}
