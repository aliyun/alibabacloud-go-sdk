// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVpcPrefixListAssociationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int64) *GetVpcPrefixListAssociationsRequest
	GetMaxResults() *int64
	SetNextToken(v string) *GetVpcPrefixListAssociationsRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *GetVpcPrefixListAssociationsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *GetVpcPrefixListAssociationsRequest
	GetOwnerId() *int64
	SetPrefixListId(v string) *GetVpcPrefixListAssociationsRequest
	GetPrefixListId() *string
	SetRegionId(v string) *GetVpcPrefixListAssociationsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *GetVpcPrefixListAssociationsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetVpcPrefixListAssociationsRequest
	GetResourceOwnerId() *int64
}

type GetVpcPrefixListAssociationsRequest struct {
	// The number of entries to return in each call. Valid values: **1*	- to **100**. Default value: **20**.
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- You do not need to specify this parameter for the first request.
	//
	// 	- If a value is returned for NextToken, specify the value in the next request to retrieve a new page of results.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the prefix list that you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// pl-0b7hwu67****
	PrefixListId *string `json:"PrefixListId,omitempty" xml:"PrefixListId,omitempty"`
	// The region ID of the prefix list.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetVpcPrefixListAssociationsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetVpcPrefixListAssociationsRequest) GoString() string {
	return s.String()
}

func (s *GetVpcPrefixListAssociationsRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *GetVpcPrefixListAssociationsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *GetVpcPrefixListAssociationsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GetVpcPrefixListAssociationsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetVpcPrefixListAssociationsRequest) GetPrefixListId() *string {
	return s.PrefixListId
}

func (s *GetVpcPrefixListAssociationsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetVpcPrefixListAssociationsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetVpcPrefixListAssociationsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetVpcPrefixListAssociationsRequest) SetMaxResults(v int64) *GetVpcPrefixListAssociationsRequest {
	s.MaxResults = &v
	return s
}

func (s *GetVpcPrefixListAssociationsRequest) SetNextToken(v string) *GetVpcPrefixListAssociationsRequest {
	s.NextToken = &v
	return s
}

func (s *GetVpcPrefixListAssociationsRequest) SetOwnerAccount(v string) *GetVpcPrefixListAssociationsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetVpcPrefixListAssociationsRequest) SetOwnerId(v int64) *GetVpcPrefixListAssociationsRequest {
	s.OwnerId = &v
	return s
}

func (s *GetVpcPrefixListAssociationsRequest) SetPrefixListId(v string) *GetVpcPrefixListAssociationsRequest {
	s.PrefixListId = &v
	return s
}

func (s *GetVpcPrefixListAssociationsRequest) SetRegionId(v string) *GetVpcPrefixListAssociationsRequest {
	s.RegionId = &v
	return s
}

func (s *GetVpcPrefixListAssociationsRequest) SetResourceOwnerAccount(v string) *GetVpcPrefixListAssociationsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetVpcPrefixListAssociationsRequest) SetResourceOwnerId(v int64) *GetVpcPrefixListAssociationsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetVpcPrefixListAssociationsRequest) Validate() error {
	return dara.Validate(s)
}
