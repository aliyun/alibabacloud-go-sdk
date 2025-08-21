// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNatGatewaysRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnsRegionId(v string) *DescribeNatGatewaysRequest
	GetEnsRegionId() *string
	SetEnsRegionIds(v []*string) *DescribeNatGatewaysRequest
	GetEnsRegionIds() []*string
	SetName(v string) *DescribeNatGatewaysRequest
	GetName() *string
	SetNatGatewayId(v string) *DescribeNatGatewaysRequest
	GetNatGatewayId() *string
	SetNatGatewayIds(v []*string) *DescribeNatGatewaysRequest
	GetNatGatewayIds() []*string
	SetNetworkId(v string) *DescribeNatGatewaysRequest
	GetNetworkId() *string
	SetPageNumber(v int32) *DescribeNatGatewaysRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeNatGatewaysRequest
	GetPageSize() *int32
	SetVSwitchId(v string) *DescribeNatGatewaysRequest
	GetVSwitchId() *string
}

type DescribeNatGatewaysRequest struct {
	// The ID of the Edge Node Service (ENS) node.
	//
	// example:
	//
	// cn-wuxi-9
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// The IDs of edge nodes. You can specify 1 to 100 IDs.
	EnsRegionIds []*string `json:"EnsRegionIds,omitempty" xml:"EnsRegionIds,omitempty" type:"Repeated"`
	// The name of the NAT gateway.
	//
	// example:
	//
	// test0
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the NAT gateway.
	//
	// example:
	//
	// nat-5t7nh1cfm6kxiszlttr38****
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// The IDs of the NAT gateways. You can specify 1 to 100 IDs.
	NatGatewayIds []*string `json:"NatGatewayIds,omitempty" xml:"NatGatewayIds,omitempty" type:"Repeated"`
	// The ID of the network.
	//
	// example:
	//
	// n-2zeuphj08tt7q3brd****
	NetworkId *string `json:"NetworkId,omitempty" xml:"NetworkId,omitempty"`
	// The page number. Pages start from page **1**.
	//
	// Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. The maximum value is **100**.
	//
	// Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the vSwitch.
	//
	// example:
	//
	// vsw-5rllcjb3ol6duzjdnbm1o****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s DescribeNatGatewaysRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNatGatewaysRequest) GoString() string {
	return s.String()
}

func (s *DescribeNatGatewaysRequest) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *DescribeNatGatewaysRequest) GetEnsRegionIds() []*string {
	return s.EnsRegionIds
}

func (s *DescribeNatGatewaysRequest) GetName() *string {
	return s.Name
}

func (s *DescribeNatGatewaysRequest) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *DescribeNatGatewaysRequest) GetNatGatewayIds() []*string {
	return s.NatGatewayIds
}

func (s *DescribeNatGatewaysRequest) GetNetworkId() *string {
	return s.NetworkId
}

func (s *DescribeNatGatewaysRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeNatGatewaysRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeNatGatewaysRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeNatGatewaysRequest) SetEnsRegionId(v string) *DescribeNatGatewaysRequest {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeNatGatewaysRequest) SetEnsRegionIds(v []*string) *DescribeNatGatewaysRequest {
	s.EnsRegionIds = v
	return s
}

func (s *DescribeNatGatewaysRequest) SetName(v string) *DescribeNatGatewaysRequest {
	s.Name = &v
	return s
}

func (s *DescribeNatGatewaysRequest) SetNatGatewayId(v string) *DescribeNatGatewaysRequest {
	s.NatGatewayId = &v
	return s
}

func (s *DescribeNatGatewaysRequest) SetNatGatewayIds(v []*string) *DescribeNatGatewaysRequest {
	s.NatGatewayIds = v
	return s
}

func (s *DescribeNatGatewaysRequest) SetNetworkId(v string) *DescribeNatGatewaysRequest {
	s.NetworkId = &v
	return s
}

func (s *DescribeNatGatewaysRequest) SetPageNumber(v int32) *DescribeNatGatewaysRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeNatGatewaysRequest) SetPageSize(v int32) *DescribeNatGatewaysRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeNatGatewaysRequest) SetVSwitchId(v string) *DescribeNatGatewaysRequest {
	s.VSwitchId = &v
	return s
}

func (s *DescribeNatGatewaysRequest) Validate() error {
	return dara.Validate(s)
}
