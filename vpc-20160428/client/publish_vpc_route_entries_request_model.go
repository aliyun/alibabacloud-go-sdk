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
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: performs a dry run without publishing route entries. The system checks the AccessKey pair, the authorization of the Resource Access Management (RAM) user, and the required parameters. If the check fails, the corresponding error is returned. If the check succeeds, the error code `DryRunOperation` is returned.
	//
	// - **false*	- (default): sends a normal request. If the check succeeds, a 2xx HTTP status code is returned and the resource status is queried.
	//
	// example:
	//
	// false
	DryRun       *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the instance. You can call the DescribeRegions operation to query the region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The list of route entries to publish. You can specify up to 50 routes.
	RouteEntries []*PublishVpcRouteEntriesRequestRouteEntries `json:"RouteEntries,omitempty" xml:"RouteEntries,omitempty" type:"Repeated"`
	// The publish route entry target instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecr-dhw2xsds5****
	TargetInstanceId *string `json:"TargetInstanceId,omitempty" xml:"TargetInstanceId,omitempty"`
	// The type of the route publish target.
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
	if s.RouteEntries != nil {
		for _, item := range s.RouteEntries {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type PublishVpcRouteEntriesRequestRouteEntries struct {
	// The destination CIDR block of the route entry.
	//
	// This parameter is required.
	//
	// example:
	//
	// 121.41.165.123/32
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	// The route table ID of the route entry.
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
