// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWithdrawVpcPublishedRouteEntriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDryRun(v bool) *WithdrawVpcPublishedRouteEntriesRequest
	GetDryRun() *bool
	SetOwnerAccount(v string) *WithdrawVpcPublishedRouteEntriesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *WithdrawVpcPublishedRouteEntriesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *WithdrawVpcPublishedRouteEntriesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *WithdrawVpcPublishedRouteEntriesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *WithdrawVpcPublishedRouteEntriesRequest
	GetResourceOwnerId() *int64
	SetRouteEntries(v []*WithdrawVpcPublishedRouteEntriesRequestRouteEntries) *WithdrawVpcPublishedRouteEntriesRequest
	GetRouteEntries() []*WithdrawVpcPublishedRouteEntriesRequestRouteEntries
	SetTargetInstanceId(v string) *WithdrawVpcPublishedRouteEntriesRequest
	GetTargetInstanceId() *string
	SetTargetType(v string) *WithdrawVpcPublishedRouteEntriesRequest
	GetTargetType() *string
}

type WithdrawVpcPublishedRouteEntriesRequest struct {
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun       *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region. Call the DescribeRegions operation to access it.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The route entries to be withdrawn. Maximum value: 50.
	RouteEntries []*WithdrawVpcPublishedRouteEntriesRequestRouteEntries `json:"RouteEntries,omitempty" xml:"RouteEntries,omitempty" type:"Repeated"`
	// Target instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecr-dhw2xsds5****
	TargetInstanceId *string `json:"TargetInstanceId,omitempty" xml:"TargetInstanceId,omitempty"`
	// The type of target instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ECR
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s WithdrawVpcPublishedRouteEntriesRequest) String() string {
	return dara.Prettify(s)
}

func (s WithdrawVpcPublishedRouteEntriesRequest) GoString() string {
	return s.String()
}

func (s *WithdrawVpcPublishedRouteEntriesRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *WithdrawVpcPublishedRouteEntriesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *WithdrawVpcPublishedRouteEntriesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *WithdrawVpcPublishedRouteEntriesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *WithdrawVpcPublishedRouteEntriesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *WithdrawVpcPublishedRouteEntriesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *WithdrawVpcPublishedRouteEntriesRequest) GetRouteEntries() []*WithdrawVpcPublishedRouteEntriesRequestRouteEntries {
	return s.RouteEntries
}

func (s *WithdrawVpcPublishedRouteEntriesRequest) GetTargetInstanceId() *string {
	return s.TargetInstanceId
}

func (s *WithdrawVpcPublishedRouteEntriesRequest) GetTargetType() *string {
	return s.TargetType
}

func (s *WithdrawVpcPublishedRouteEntriesRequest) SetDryRun(v bool) *WithdrawVpcPublishedRouteEntriesRequest {
	s.DryRun = &v
	return s
}

func (s *WithdrawVpcPublishedRouteEntriesRequest) SetOwnerAccount(v string) *WithdrawVpcPublishedRouteEntriesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *WithdrawVpcPublishedRouteEntriesRequest) SetOwnerId(v int64) *WithdrawVpcPublishedRouteEntriesRequest {
	s.OwnerId = &v
	return s
}

func (s *WithdrawVpcPublishedRouteEntriesRequest) SetRegionId(v string) *WithdrawVpcPublishedRouteEntriesRequest {
	s.RegionId = &v
	return s
}

func (s *WithdrawVpcPublishedRouteEntriesRequest) SetResourceOwnerAccount(v string) *WithdrawVpcPublishedRouteEntriesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *WithdrawVpcPublishedRouteEntriesRequest) SetResourceOwnerId(v int64) *WithdrawVpcPublishedRouteEntriesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *WithdrawVpcPublishedRouteEntriesRequest) SetRouteEntries(v []*WithdrawVpcPublishedRouteEntriesRequestRouteEntries) *WithdrawVpcPublishedRouteEntriesRequest {
	s.RouteEntries = v
	return s
}

func (s *WithdrawVpcPublishedRouteEntriesRequest) SetTargetInstanceId(v string) *WithdrawVpcPublishedRouteEntriesRequest {
	s.TargetInstanceId = &v
	return s
}

func (s *WithdrawVpcPublishedRouteEntriesRequest) SetTargetType(v string) *WithdrawVpcPublishedRouteEntriesRequest {
	s.TargetType = &v
	return s
}

func (s *WithdrawVpcPublishedRouteEntriesRequest) Validate() error {
	return dara.Validate(s)
}

type WithdrawVpcPublishedRouteEntriesRequestRouteEntries struct {
	// The destination CIDR block
	//
	// This parameter is required.
	//
	// example:
	//
	// 10.0.0.0/24
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	// The ID of the route table.
	//
	// This parameter is required.
	//
	// example:
	//
	// vtb-bp145q7glnuzd****
	RouteTableId *string `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
}

func (s WithdrawVpcPublishedRouteEntriesRequestRouteEntries) String() string {
	return dara.Prettify(s)
}

func (s WithdrawVpcPublishedRouteEntriesRequestRouteEntries) GoString() string {
	return s.String()
}

func (s *WithdrawVpcPublishedRouteEntriesRequestRouteEntries) GetDestinationCidrBlock() *string {
	return s.DestinationCidrBlock
}

func (s *WithdrawVpcPublishedRouteEntriesRequestRouteEntries) GetRouteTableId() *string {
	return s.RouteTableId
}

func (s *WithdrawVpcPublishedRouteEntriesRequestRouteEntries) SetDestinationCidrBlock(v string) *WithdrawVpcPublishedRouteEntriesRequestRouteEntries {
	s.DestinationCidrBlock = &v
	return s
}

func (s *WithdrawVpcPublishedRouteEntriesRequestRouteEntries) SetRouteTableId(v string) *WithdrawVpcPublishedRouteEntriesRequestRouteEntries {
	s.RouteTableId = &v
	return s
}

func (s *WithdrawVpcPublishedRouteEntriesRequestRouteEntries) Validate() error {
	return dara.Validate(s)
}
