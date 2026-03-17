// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudConnectNetworksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCcnId(v string) *DescribeCloudConnectNetworksRequest
	GetCcnId() *string
	SetName(v string) *DescribeCloudConnectNetworksRequest
	GetName() *string
	SetOwnerAccount(v string) *DescribeCloudConnectNetworksRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeCloudConnectNetworksRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeCloudConnectNetworksRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeCloudConnectNetworksRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeCloudConnectNetworksRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeCloudConnectNetworksRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeCloudConnectNetworksRequest
	GetResourceOwnerId() *int64
	SetTag(v []*DescribeCloudConnectNetworksRequestTag) *DescribeCloudConnectNetworksRequest
	GetTag() []*DescribeCloudConnectNetworksRequestTag
}

type DescribeCloudConnectNetworksRequest struct {
	// The ID of the CCN instance.
	//
	// example:
	//
	// ccn-l9340rlu5enst*****
	CcnId *string `json:"CcnId,omitempty" xml:"CcnId,omitempty"`
	// The name of the CCN instance.
	//
	// The name must be 2 to 100 characters in length and can contain letters, digits, periods (.), underscores (_),and hyphens (-). The name must start with a letter.
	//
	// example:
	//
	// ccnname
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: **10**. Maximum value: **50**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region where the CCN instances are deployed.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string                                   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string                                   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                                    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Tag                  []*DescribeCloudConnectNetworksRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeCloudConnectNetworksRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudConnectNetworksRequest) GoString() string {
	return s.String()
}

func (s *DescribeCloudConnectNetworksRequest) GetCcnId() *string {
	return s.CcnId
}

func (s *DescribeCloudConnectNetworksRequest) GetName() *string {
	return s.Name
}

func (s *DescribeCloudConnectNetworksRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeCloudConnectNetworksRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeCloudConnectNetworksRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeCloudConnectNetworksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCloudConnectNetworksRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCloudConnectNetworksRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeCloudConnectNetworksRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeCloudConnectNetworksRequest) GetTag() []*DescribeCloudConnectNetworksRequestTag {
	return s.Tag
}

func (s *DescribeCloudConnectNetworksRequest) SetCcnId(v string) *DescribeCloudConnectNetworksRequest {
	s.CcnId = &v
	return s
}

func (s *DescribeCloudConnectNetworksRequest) SetName(v string) *DescribeCloudConnectNetworksRequest {
	s.Name = &v
	return s
}

func (s *DescribeCloudConnectNetworksRequest) SetOwnerAccount(v string) *DescribeCloudConnectNetworksRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeCloudConnectNetworksRequest) SetOwnerId(v int64) *DescribeCloudConnectNetworksRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCloudConnectNetworksRequest) SetPageNumber(v int32) *DescribeCloudConnectNetworksRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCloudConnectNetworksRequest) SetPageSize(v int32) *DescribeCloudConnectNetworksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCloudConnectNetworksRequest) SetRegionId(v string) *DescribeCloudConnectNetworksRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCloudConnectNetworksRequest) SetResourceOwnerAccount(v string) *DescribeCloudConnectNetworksRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeCloudConnectNetworksRequest) SetResourceOwnerId(v int64) *DescribeCloudConnectNetworksRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeCloudConnectNetworksRequest) SetTag(v []*DescribeCloudConnectNetworksRequestTag) *DescribeCloudConnectNetworksRequest {
	s.Tag = v
	return s
}

func (s *DescribeCloudConnectNetworksRequest) Validate() error {
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

type DescribeCloudConnectNetworksRequestTag struct {
	// The key of the tag that is bound to the CCN instance.
	//
	// example:
	//
	// test
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag that is bound to the CCN instance.
	//
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeCloudConnectNetworksRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudConnectNetworksRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeCloudConnectNetworksRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeCloudConnectNetworksRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeCloudConnectNetworksRequestTag) SetKey(v string) *DescribeCloudConnectNetworksRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeCloudConnectNetworksRequestTag) SetValue(v string) *DescribeCloudConnectNetworksRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeCloudConnectNetworksRequestTag) Validate() error {
	return dara.Validate(s)
}
