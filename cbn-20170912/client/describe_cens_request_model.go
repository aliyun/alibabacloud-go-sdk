// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCensRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v []*DescribeCensRequestFilter) *DescribeCensRequest
	GetFilter() []*DescribeCensRequestFilter
	SetOwnerAccount(v string) *DescribeCensRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeCensRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeCensRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeCensRequest
	GetPageSize() *int32
	SetResourceGroupId(v string) *DescribeCensRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeCensRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeCensRequest
	GetResourceOwnerId() *int64
	SetTag(v []*DescribeCensRequestTag) *DescribeCensRequest
	GetTag() []*DescribeCensRequestTag
}

type DescribeCensRequest struct {
	// The filter conditions.
	//
	// You can specify at most five filter conditions in each call.
	Filter       []*DescribeCensRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	OwnerAccount *string                      `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64                       `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values: **1*	- to **50**. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the resource group to which the CEN instance belongs.
	//
	// example:
	//
	// rg-acfm3unpnuw****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tags.
	//
	// You can specify at most 20 tags in each call.
	Tag []*DescribeCensRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeCensRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCensRequest) GoString() string {
	return s.String()
}

func (s *DescribeCensRequest) GetFilter() []*DescribeCensRequestFilter {
	return s.Filter
}

func (s *DescribeCensRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeCensRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeCensRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeCensRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCensRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeCensRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeCensRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeCensRequest) GetTag() []*DescribeCensRequestTag {
	return s.Tag
}

func (s *DescribeCensRequest) SetFilter(v []*DescribeCensRequestFilter) *DescribeCensRequest {
	s.Filter = v
	return s
}

func (s *DescribeCensRequest) SetOwnerAccount(v string) *DescribeCensRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeCensRequest) SetOwnerId(v int64) *DescribeCensRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCensRequest) SetPageNumber(v int32) *DescribeCensRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCensRequest) SetPageSize(v int32) *DescribeCensRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCensRequest) SetResourceGroupId(v string) *DescribeCensRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeCensRequest) SetResourceOwnerAccount(v string) *DescribeCensRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeCensRequest) SetResourceOwnerId(v int64) *DescribeCensRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeCensRequest) SetTag(v []*DescribeCensRequestTag) *DescribeCensRequest {
	s.Tag = v
	return s
}

func (s *DescribeCensRequest) Validate() error {
	if s.Filter != nil {
		for _, item := range s.Filter {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
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

type DescribeCensRequestFilter struct {
	// The key of the filter. Valid values:
	//
	// 	- **CenId**: the ID of a CEN instance.
	//
	// 	- **Name**: the name of a CEN instance.
	//
	// By default, the logical operator among filter conditions is **AND**. Information about a CEN instance is returned only if the CEN instance matches all filter conditions.
	//
	// You can specify at most five filter conditions in each call.
	//
	// example:
	//
	// CenId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the filter condition.
	//
	// Specify a filter value based on the **Key*	- parameter. You can specify multiple values for a filter **key**. The logical operator among multiple filter values is **OR**. If a CEN instance matches one or more of the values that you specify, the CEN instance matches the filter condition.
	//
	// You can specify at most five values in each filter condition.
	//
	// example:
	//
	// cen-0xyeagctz5sfg9****
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s DescribeCensRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s DescribeCensRequestFilter) GoString() string {
	return s.String()
}

func (s *DescribeCensRequestFilter) GetKey() *string {
	return s.Key
}

func (s *DescribeCensRequestFilter) GetValue() []*string {
	return s.Value
}

func (s *DescribeCensRequestFilter) SetKey(v string) *DescribeCensRequestFilter {
	s.Key = &v
	return s
}

func (s *DescribeCensRequestFilter) SetValue(v []*string) *DescribeCensRequestFilter {
	s.Value = v
	return s
}

func (s *DescribeCensRequestFilter) Validate() error {
	return dara.Validate(s)
}

type DescribeCensRequestTag struct {
	// The tag keys.
	//
	// The tag keys cannot be an empty string. The tag keys can be up to 64 characters in length and cannot start with `acs:` or `aliyun`. It cannot contain `http://` or `https://`.
	//
	// You can specify at most 20 tag keys.
	//
	// example:
	//
	// tagtest
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag values.
	//
	// The tag values can be 0 to 128 characters in length, and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// Each tag key must have a unique tag value. You can specify at most 20 tag values in each call.
	//
	// example:
	//
	// tagtest
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeCensRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeCensRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeCensRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeCensRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeCensRequestTag) SetKey(v string) *DescribeCensRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeCensRequestTag) SetValue(v string) *DescribeCensRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeCensRequestTag) Validate() error {
	return dara.Validate(s)
}
