// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTransitRouterCidrAllocationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttachmentId(v string) *ListTransitRouterCidrAllocationRequest
	GetAttachmentId() *string
	SetAttachmentName(v string) *ListTransitRouterCidrAllocationRequest
	GetAttachmentName() *string
	SetCidr(v string) *ListTransitRouterCidrAllocationRequest
	GetCidr() *string
	SetCidrBlock(v string) *ListTransitRouterCidrAllocationRequest
	GetCidrBlock() *string
	SetClientToken(v string) *ListTransitRouterCidrAllocationRequest
	GetClientToken() *string
	SetDedicatedOwnerId(v string) *ListTransitRouterCidrAllocationRequest
	GetDedicatedOwnerId() *string
	SetDryRun(v bool) *ListTransitRouterCidrAllocationRequest
	GetDryRun() *bool
	SetMaxResults(v int32) *ListTransitRouterCidrAllocationRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListTransitRouterCidrAllocationRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *ListTransitRouterCidrAllocationRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListTransitRouterCidrAllocationRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ListTransitRouterCidrAllocationRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ListTransitRouterCidrAllocationRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListTransitRouterCidrAllocationRequest
	GetResourceOwnerId() *int64
	SetTransitRouterCidrId(v string) *ListTransitRouterCidrAllocationRequest
	GetTransitRouterCidrId() *string
	SetTransitRouterId(v string) *ListTransitRouterCidrAllocationRequest
	GetTransitRouterId() *string
}

type ListTransitRouterCidrAllocationRequest struct {
	// The ID of the network instance connection.
	//
	// example:
	//
	// tr-attach-2nalp6yksc805w****
	AttachmentId *string `json:"AttachmentId,omitempty" xml:"AttachmentId,omitempty"`
	// The name of the network instance connection.
	//
	// example:
	//
	// nametest
	AttachmentName *string `json:"AttachmentName,omitempty" xml:"AttachmentName,omitempty"`
	// The transit router CIDR block.
	//
	// example:
	//
	// 192.168.10.0/24
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// The allocated CIDR block under the transit router CIDR block.
	//
	// example:
	//
	// 192.168.10.0/28
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- may be different for each API request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The dedicated CIDR block.
	//
	// Set the value to **VPN**, which specifies that you want to query the CIDR block reserved by the system for creating VPN connections in the backend.
	//
	// example:
	//
	// VPN
	DedicatedOwnerId *string `json:"DedicatedOwnerId,omitempty" xml:"DedicatedOwnerId,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: performs a dry run. The system checks the required parameters, request syntax, and business restrictions. If the request fails the dry run, the corresponding error is returned. If the request passes the dry run, the error code `DryRunOperation` is returned.
	//
	// - **false*	- (default): performs a dry run and sends the request. If the request passes the dry run, the transit router CIDR block allocation details are queried.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The number of entries per page.
	//
	// - If you do not specify a value for **MaxResults**, it indicates that you do not need to query results by page. The value of **MaxResults*	- in the response indicates the total number of entries.
	//
	// - If you specify a value for **MaxResults**, it indicates that you need to query results by page. Valid values: **1*	- to **100**. We recommend that you set **MaxResults*	- to **20**.
	//
	//   The value of **MaxResults*	- in the response indicates the number of entries on the current page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// - You do not need to specify this parameter for the first request or if no subsequent request exists.
	//
	// - If a subsequent request exists, set the value to the **NextToken*	- value returned in the previous API call.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the transit router instance.
	//
	// You can call the [DescribeChildInstanceRegions](https://help.aliyun.com/document_detail/132080.html) operation to query region IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the transit router CIDR block.
	//
	// You can call the [ListTransitRouterCidr](https://help.aliyun.com/document_detail/462772.html) operation to query the transit router CIDR block ID.
	//
	// example:
	//
	// cidr-0zv0q9crqpntzz****
	TransitRouterCidrId *string `json:"TransitRouterCidrId,omitempty" xml:"TransitRouterCidrId,omitempty"`
	// The forward router instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// tr-p0w3x8c9em72a40nw****
	TransitRouterId *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
}

func (s ListTransitRouterCidrAllocationRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRouterCidrAllocationRequest) GoString() string {
	return s.String()
}

func (s *ListTransitRouterCidrAllocationRequest) GetAttachmentId() *string {
	return s.AttachmentId
}

func (s *ListTransitRouterCidrAllocationRequest) GetAttachmentName() *string {
	return s.AttachmentName
}

func (s *ListTransitRouterCidrAllocationRequest) GetCidr() *string {
	return s.Cidr
}

func (s *ListTransitRouterCidrAllocationRequest) GetCidrBlock() *string {
	return s.CidrBlock
}

func (s *ListTransitRouterCidrAllocationRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ListTransitRouterCidrAllocationRequest) GetDedicatedOwnerId() *string {
	return s.DedicatedOwnerId
}

func (s *ListTransitRouterCidrAllocationRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ListTransitRouterCidrAllocationRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTransitRouterCidrAllocationRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTransitRouterCidrAllocationRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListTransitRouterCidrAllocationRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListTransitRouterCidrAllocationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListTransitRouterCidrAllocationRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListTransitRouterCidrAllocationRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListTransitRouterCidrAllocationRequest) GetTransitRouterCidrId() *string {
	return s.TransitRouterCidrId
}

func (s *ListTransitRouterCidrAllocationRequest) GetTransitRouterId() *string {
	return s.TransitRouterId
}

func (s *ListTransitRouterCidrAllocationRequest) SetAttachmentId(v string) *ListTransitRouterCidrAllocationRequest {
	s.AttachmentId = &v
	return s
}

func (s *ListTransitRouterCidrAllocationRequest) SetAttachmentName(v string) *ListTransitRouterCidrAllocationRequest {
	s.AttachmentName = &v
	return s
}

func (s *ListTransitRouterCidrAllocationRequest) SetCidr(v string) *ListTransitRouterCidrAllocationRequest {
	s.Cidr = &v
	return s
}

func (s *ListTransitRouterCidrAllocationRequest) SetCidrBlock(v string) *ListTransitRouterCidrAllocationRequest {
	s.CidrBlock = &v
	return s
}

func (s *ListTransitRouterCidrAllocationRequest) SetClientToken(v string) *ListTransitRouterCidrAllocationRequest {
	s.ClientToken = &v
	return s
}

func (s *ListTransitRouterCidrAllocationRequest) SetDedicatedOwnerId(v string) *ListTransitRouterCidrAllocationRequest {
	s.DedicatedOwnerId = &v
	return s
}

func (s *ListTransitRouterCidrAllocationRequest) SetDryRun(v bool) *ListTransitRouterCidrAllocationRequest {
	s.DryRun = &v
	return s
}

func (s *ListTransitRouterCidrAllocationRequest) SetMaxResults(v int32) *ListTransitRouterCidrAllocationRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTransitRouterCidrAllocationRequest) SetNextToken(v string) *ListTransitRouterCidrAllocationRequest {
	s.NextToken = &v
	return s
}

func (s *ListTransitRouterCidrAllocationRequest) SetOwnerAccount(v string) *ListTransitRouterCidrAllocationRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListTransitRouterCidrAllocationRequest) SetOwnerId(v int64) *ListTransitRouterCidrAllocationRequest {
	s.OwnerId = &v
	return s
}

func (s *ListTransitRouterCidrAllocationRequest) SetRegionId(v string) *ListTransitRouterCidrAllocationRequest {
	s.RegionId = &v
	return s
}

func (s *ListTransitRouterCidrAllocationRequest) SetResourceOwnerAccount(v string) *ListTransitRouterCidrAllocationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListTransitRouterCidrAllocationRequest) SetResourceOwnerId(v int64) *ListTransitRouterCidrAllocationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListTransitRouterCidrAllocationRequest) SetTransitRouterCidrId(v string) *ListTransitRouterCidrAllocationRequest {
	s.TransitRouterCidrId = &v
	return s
}

func (s *ListTransitRouterCidrAllocationRequest) SetTransitRouterId(v string) *ListTransitRouterCidrAllocationRequest {
	s.TransitRouterId = &v
	return s
}

func (s *ListTransitRouterCidrAllocationRequest) Validate() error {
	return dara.Validate(s)
}
