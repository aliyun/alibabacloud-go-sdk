// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDpiGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDpiGroupIds(v []*string) *ListDpiGroupsRequest
	GetDpiGroupIds() []*string
	SetDpiGroupNames(v []*string) *ListDpiGroupsRequest
	GetDpiGroupNames() []*string
	SetMaxResults(v int32) *ListDpiGroupsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListDpiGroupsRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *ListDpiGroupsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListDpiGroupsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ListDpiGroupsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ListDpiGroupsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListDpiGroupsRequest
	GetResourceOwnerId() *int64
}

type ListDpiGroupsRequest struct {
	// example:
	//
	// 1
	DpiGroupIds []*string `json:"DpiGroupIds,omitempty" xml:"DpiGroupIds,omitempty" type:"Repeated"`
	// example:
	//
	// P2P
	DpiGroupNames []*string `json:"DpiGroupNames,omitempty" xml:"DpiGroupNames,omitempty" type:"Repeated"`
	// The number of entries to return on each page.
	//
	// Valid values: **1*	- to **100**.
	//
	// Default value: **20**.
	//
	// example:
	//
	// 3
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token used to query the next page.
	//
	// example:
	//
	// FFPSpX59Eb****
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region to which the application groups belong.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/69813.html) to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ListDpiGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDpiGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListDpiGroupsRequest) GetDpiGroupIds() []*string {
	return s.DpiGroupIds
}

func (s *ListDpiGroupsRequest) GetDpiGroupNames() []*string {
	return s.DpiGroupNames
}

func (s *ListDpiGroupsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDpiGroupsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDpiGroupsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListDpiGroupsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListDpiGroupsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListDpiGroupsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListDpiGroupsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListDpiGroupsRequest) SetDpiGroupIds(v []*string) *ListDpiGroupsRequest {
	s.DpiGroupIds = v
	return s
}

func (s *ListDpiGroupsRequest) SetDpiGroupNames(v []*string) *ListDpiGroupsRequest {
	s.DpiGroupNames = v
	return s
}

func (s *ListDpiGroupsRequest) SetMaxResults(v int32) *ListDpiGroupsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDpiGroupsRequest) SetNextToken(v string) *ListDpiGroupsRequest {
	s.NextToken = &v
	return s
}

func (s *ListDpiGroupsRequest) SetOwnerAccount(v string) *ListDpiGroupsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListDpiGroupsRequest) SetOwnerId(v int64) *ListDpiGroupsRequest {
	s.OwnerId = &v
	return s
}

func (s *ListDpiGroupsRequest) SetRegionId(v string) *ListDpiGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *ListDpiGroupsRequest) SetResourceOwnerAccount(v string) *ListDpiGroupsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListDpiGroupsRequest) SetResourceOwnerId(v int64) *ListDpiGroupsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListDpiGroupsRequest) Validate() error {
	return dara.Validate(s)
}
