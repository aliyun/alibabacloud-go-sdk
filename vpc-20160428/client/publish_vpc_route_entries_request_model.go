// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishVpcRouteEntriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDryRun(v bool) *PublishVpcRouteEntriesRequest
	GetDryRun() *bool
	SetOwnerAccount(v string) *PublishVpcRouteEntriesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *PublishVpcRouteEntriesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *PublishVpcRouteEntriesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *PublishVpcRouteEntriesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *PublishVpcRouteEntriesRequest
	GetResourceOwnerId() *int64
	SetRouteEntries(v []*PublishVpcRouteEntriesRequestRouteEntries) *PublishVpcRouteEntriesRequest
	GetRouteEntries() []*PublishVpcRouteEntriesRequestRouteEntries
	SetTargetInstanceId(v string) *PublishVpcRouteEntriesRequest
	GetTargetInstanceId() *string
	SetTargetType(v string) *PublishVpcRouteEntriesRequest
	GetTargetType() *string
}

type PublishVpcRouteEntriesRequest struct {
	// Indicates whether to perform a dry run of this request. Values:
	//
	// - **true**: Sends a check request without publishing the route. The checks include whether the AccessKey is valid, the authorization status of the RAM user, and if all required parameters are filled out. If the check fails, the corresponding error is returned. If the check passes, the `DryRunOperation` error code is returned.
	//
	// - **false*	- (default): Sends a normal request. After passing the check, it returns a 2xx HTTP status code and directly queries the resource status.
	//
	// example:
	//
	// false
	DryRun       *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the instance is located. You can obtain the region ID by calling the DescribeRegions interface.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// List of route entries to be published, supporting up to 50 routes at most.
	RouteEntries []*PublishVpcRouteEntriesRequestRouteEntries `json:"RouteEntries,omitempty" xml:"RouteEntries,omitempty" type:"Repeated"`
	// The ID of the target instance for route publication.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecr-dhw2xsds5****
	TargetInstanceId *string `json:"TargetInstanceId,omitempty" xml:"TargetInstanceId,omitempty"`
	// The type of the target for route publication.
	//
	// This parameter is required.
	//
	// example:
	//
	// ECR
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s PublishVpcRouteEntriesRequest) String() string {
	return dara.Prettify(s)
}

func (s PublishVpcRouteEntriesRequest) GoString() string {
	return s.String()
}

func (s *PublishVpcRouteEntriesRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *PublishVpcRouteEntriesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *PublishVpcRouteEntriesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *PublishVpcRouteEntriesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *PublishVpcRouteEntriesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *PublishVpcRouteEntriesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *PublishVpcRouteEntriesRequest) GetRouteEntries() []*PublishVpcRouteEntriesRequestRouteEntries {
	return s.RouteEntries
}

func (s *PublishVpcRouteEntriesRequest) GetTargetInstanceId() *string {
	return s.TargetInstanceId
}

func (s *PublishVpcRouteEntriesRequest) GetTargetType() *string {
	return s.TargetType
}

func (s *PublishVpcRouteEntriesRequest) SetDryRun(v bool) *PublishVpcRouteEntriesRequest {
	s.DryRun = &v
	return s
}

func (s *PublishVpcRouteEntriesRequest) SetOwnerAccount(v string) *PublishVpcRouteEntriesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *PublishVpcRouteEntriesRequest) SetOwnerId(v int64) *PublishVpcRouteEntriesRequest {
	s.OwnerId = &v
	return s
}

func (s *PublishVpcRouteEntriesRequest) SetRegionId(v string) *PublishVpcRouteEntriesRequest {
	s.RegionId = &v
	return s
}

func (s *PublishVpcRouteEntriesRequest) SetResourceOwnerAccount(v string) *PublishVpcRouteEntriesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *PublishVpcRouteEntriesRequest) SetResourceOwnerId(v int64) *PublishVpcRouteEntriesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *PublishVpcRouteEntriesRequest) SetRouteEntries(v []*PublishVpcRouteEntriesRequestRouteEntries) *PublishVpcRouteEntriesRequest {
	s.RouteEntries = v
	return s
}

func (s *PublishVpcRouteEntriesRequest) SetTargetInstanceId(v string) *PublishVpcRouteEntriesRequest {
	s.TargetInstanceId = &v
	return s
}

func (s *PublishVpcRouteEntriesRequest) SetTargetType(v string) *PublishVpcRouteEntriesRequest {
	s.TargetType = &v
	return s
}

func (s *PublishVpcRouteEntriesRequest) Validate() error {
	return dara.Validate(s)
}

type PublishVpcRouteEntriesRequestRouteEntries struct {
	// The destination CIDR block for the route entry.
	//
	// This parameter is required.
	//
	// example:
	//
	// 121.41.165.123/32
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	// The ID of the route table for the route entry.
	//
	// This parameter is required.
	//
	// example:
	//
	// vtb-2ze3jgygk9bmsj23s****
	RouteTableId *string `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
}

func (s PublishVpcRouteEntriesRequestRouteEntries) String() string {
	return dara.Prettify(s)
}

func (s PublishVpcRouteEntriesRequestRouteEntries) GoString() string {
	return s.String()
}

func (s *PublishVpcRouteEntriesRequestRouteEntries) GetDestinationCidrBlock() *string {
	return s.DestinationCidrBlock
}

func (s *PublishVpcRouteEntriesRequestRouteEntries) GetRouteTableId() *string {
	return s.RouteTableId
}

func (s *PublishVpcRouteEntriesRequestRouteEntries) SetDestinationCidrBlock(v string) *PublishVpcRouteEntriesRequestRouteEntries {
	s.DestinationCidrBlock = &v
	return s
}

func (s *PublishVpcRouteEntriesRequestRouteEntries) SetRouteTableId(v string) *PublishVpcRouteEntriesRequestRouteEntries {
	s.RouteTableId = &v
	return s
}

func (s *PublishVpcRouteEntriesRequestRouteEntries) Validate() error {
	return dara.Validate(s)
}
