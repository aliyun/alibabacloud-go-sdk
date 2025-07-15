// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPublicIpAddressPoolCidrBlocksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCidrBlock(v string) *ListPublicIpAddressPoolCidrBlocksRequest
	GetCidrBlock() *string
	SetDryRun(v bool) *ListPublicIpAddressPoolCidrBlocksRequest
	GetDryRun() *bool
	SetMaxResults(v int32) *ListPublicIpAddressPoolCidrBlocksRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListPublicIpAddressPoolCidrBlocksRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *ListPublicIpAddressPoolCidrBlocksRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListPublicIpAddressPoolCidrBlocksRequest
	GetOwnerId() *int64
	SetPublicIpAddressPoolId(v string) *ListPublicIpAddressPoolCidrBlocksRequest
	GetPublicIpAddressPoolId() *string
	SetRegionId(v string) *ListPublicIpAddressPoolCidrBlocksRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ListPublicIpAddressPoolCidrBlocksRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListPublicIpAddressPoolCidrBlocksRequest
	GetResourceOwnerId() *int64
}

type ListPublicIpAddressPoolCidrBlocksRequest struct {
	// The CIDR blocks.
	//
	// example:
	//
	// 47.0.XX.XX/24
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// 	- **true**: performs a dry run. The system checks the required parameters, request syntax, and limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and sends the request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The maximum number of entries to return. Valid values: **10*	- to **100**. Default value: **10**.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used for the next query. Valid values:
	//
	// 	- If this is your first query and no next queries are to be sent, ignore this parameter.
	//
	// 	- If a subsequent query is to be sent, set the parameter to the value of NextToken that is returned from the last call.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the IP address pool.
	//
	// This parameter is required.
	//
	// example:
	//
	// pippool-6wetvn6fumkgycssx****
	PublicIpAddressPoolId *string `json:"PublicIpAddressPoolId,omitempty" xml:"PublicIpAddressPoolId,omitempty"`
	// The region ID of the CIDR blocks.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-chengdu
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ListPublicIpAddressPoolCidrBlocksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPublicIpAddressPoolCidrBlocksRequest) GoString() string {
	return s.String()
}

func (s *ListPublicIpAddressPoolCidrBlocksRequest) GetCidrBlock() *string {
	return s.CidrBlock
}

func (s *ListPublicIpAddressPoolCidrBlocksRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ListPublicIpAddressPoolCidrBlocksRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListPublicIpAddressPoolCidrBlocksRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPublicIpAddressPoolCidrBlocksRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListPublicIpAddressPoolCidrBlocksRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListPublicIpAddressPoolCidrBlocksRequest) GetPublicIpAddressPoolId() *string {
	return s.PublicIpAddressPoolId
}

func (s *ListPublicIpAddressPoolCidrBlocksRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListPublicIpAddressPoolCidrBlocksRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListPublicIpAddressPoolCidrBlocksRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListPublicIpAddressPoolCidrBlocksRequest) SetCidrBlock(v string) *ListPublicIpAddressPoolCidrBlocksRequest {
	s.CidrBlock = &v
	return s
}

func (s *ListPublicIpAddressPoolCidrBlocksRequest) SetDryRun(v bool) *ListPublicIpAddressPoolCidrBlocksRequest {
	s.DryRun = &v
	return s
}

func (s *ListPublicIpAddressPoolCidrBlocksRequest) SetMaxResults(v int32) *ListPublicIpAddressPoolCidrBlocksRequest {
	s.MaxResults = &v
	return s
}

func (s *ListPublicIpAddressPoolCidrBlocksRequest) SetNextToken(v string) *ListPublicIpAddressPoolCidrBlocksRequest {
	s.NextToken = &v
	return s
}

func (s *ListPublicIpAddressPoolCidrBlocksRequest) SetOwnerAccount(v string) *ListPublicIpAddressPoolCidrBlocksRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListPublicIpAddressPoolCidrBlocksRequest) SetOwnerId(v int64) *ListPublicIpAddressPoolCidrBlocksRequest {
	s.OwnerId = &v
	return s
}

func (s *ListPublicIpAddressPoolCidrBlocksRequest) SetPublicIpAddressPoolId(v string) *ListPublicIpAddressPoolCidrBlocksRequest {
	s.PublicIpAddressPoolId = &v
	return s
}

func (s *ListPublicIpAddressPoolCidrBlocksRequest) SetRegionId(v string) *ListPublicIpAddressPoolCidrBlocksRequest {
	s.RegionId = &v
	return s
}

func (s *ListPublicIpAddressPoolCidrBlocksRequest) SetResourceOwnerAccount(v string) *ListPublicIpAddressPoolCidrBlocksRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListPublicIpAddressPoolCidrBlocksRequest) SetResourceOwnerId(v int64) *ListPublicIpAddressPoolCidrBlocksRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListPublicIpAddressPoolCidrBlocksRequest) Validate() error {
	return dara.Validate(s)
}
