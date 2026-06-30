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
	// You can specify up to five filter conditions.
	Filter       []*DescribeCensRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	OwnerAccount *string                      `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64                       `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Default value: **1**.
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
	// You can specify up to 20 tags.
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
	// The filter key. Valid values:
	//
	// - **CenId**: The ID of the CEN instance.
	//
	// - **Name**: The name of the CEN instance.
	//
	// The logical relationship among multiple filter conditions is **AND**. All filter conditions must be met.
	//
	// You can specify up to five filter conditions.
	//
	// example:
	//
	// CenId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The filter value.
	//
	// Specify the filter value based on the **Key**. You can specify multiple values for a key. The logical relationship among the values is **OR**. A resource is a match if it meets any of the specified values.
	//
	// You can specify up to five filter values for a filter condition.
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
	// The tag key of the resource.
	//
	// The tag key cannot be an empty string. The tag key can be up to 64 characters in length and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// You can specify up to 20 tag keys.
	//
	// example:
	//
	// tagtest
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the resource.
	//
	// The tag value can be an empty string or a string of up to 128 characters. It cannot start with `aliyun` or `acs:` and cannot contain `http://` or `https://`.
	//
	// Each tag key must have a corresponding tag value. You can specify up to 20 tag values.
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
