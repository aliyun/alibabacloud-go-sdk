// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVpcPrefixListEntriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int64) *GetVpcPrefixListEntriesRequest
	GetMaxResults() *int64
	SetNextToken(v string) *GetVpcPrefixListEntriesRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *GetVpcPrefixListEntriesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *GetVpcPrefixListEntriesRequest
	GetOwnerId() *int64
	SetPrefixListId(v string) *GetVpcPrefixListEntriesRequest
	GetPrefixListId() *string
	SetRegionId(v string) *GetVpcPrefixListEntriesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *GetVpcPrefixListEntriesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetVpcPrefixListEntriesRequest
	GetResourceOwnerId() *int64
}

type GetVpcPrefixListEntriesRequest struct {
	// The number of entries per page. Valid values: **1*	- to **100**. Default value: **20**.
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- You do not need to specify this parameter for the first request.
	//
	// 	- You must specify the token that is obtained from the previous query as the value of NextToken.
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

func (s GetVpcPrefixListEntriesRequest) String() string {
	return dara.Prettify(s)
}

func (s GetVpcPrefixListEntriesRequest) GoString() string {
	return s.String()
}

func (s *GetVpcPrefixListEntriesRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *GetVpcPrefixListEntriesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *GetVpcPrefixListEntriesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GetVpcPrefixListEntriesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetVpcPrefixListEntriesRequest) GetPrefixListId() *string {
	return s.PrefixListId
}

func (s *GetVpcPrefixListEntriesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetVpcPrefixListEntriesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetVpcPrefixListEntriesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetVpcPrefixListEntriesRequest) SetMaxResults(v int64) *GetVpcPrefixListEntriesRequest {
	s.MaxResults = &v
	return s
}

func (s *GetVpcPrefixListEntriesRequest) SetNextToken(v string) *GetVpcPrefixListEntriesRequest {
	s.NextToken = &v
	return s
}

func (s *GetVpcPrefixListEntriesRequest) SetOwnerAccount(v string) *GetVpcPrefixListEntriesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetVpcPrefixListEntriesRequest) SetOwnerId(v int64) *GetVpcPrefixListEntriesRequest {
	s.OwnerId = &v
	return s
}

func (s *GetVpcPrefixListEntriesRequest) SetPrefixListId(v string) *GetVpcPrefixListEntriesRequest {
	s.PrefixListId = &v
	return s
}

func (s *GetVpcPrefixListEntriesRequest) SetRegionId(v string) *GetVpcPrefixListEntriesRequest {
	s.RegionId = &v
	return s
}

func (s *GetVpcPrefixListEntriesRequest) SetResourceOwnerAccount(v string) *GetVpcPrefixListEntriesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetVpcPrefixListEntriesRequest) SetResourceOwnerId(v int64) *GetVpcPrefixListEntriesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetVpcPrefixListEntriesRequest) Validate() error {
	return dara.Validate(s)
}
