// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVcoRouteEntriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DescribeVcoRouteEntriesRequest
	GetClientToken() *string
	SetOwnerAccount(v string) *DescribeVcoRouteEntriesRequest
	GetOwnerAccount() *string
	SetPageNumber(v int32) *DescribeVcoRouteEntriesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeVcoRouteEntriesRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeVcoRouteEntriesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeVcoRouteEntriesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeVcoRouteEntriesRequest
	GetResourceOwnerId() *int64
	SetRouteEntryType(v string) *DescribeVcoRouteEntriesRequest
	GetRouteEntryType() *string
	SetVpnConnectionId(v string) *DescribeVcoRouteEntriesRequest
	GetVpnConnectionId() *string
}

type DescribeVcoRouteEntriesRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that it is unique among all requests. ClientToken can contain only ASCII characters.
	//
	// >  If you do not set this parameter, the system uses **RequestId*	- as **ClientToken**. The value of **RequestId*	- for each API request may be different.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-4266****
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: **10**. Valid values: **1*	- to **50**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the IPsec-VPN connection.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent list of regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The route type. Valid values:
	//
	// 	- **custom*	- (default): a destination-based route
	//
	// 	- **bgp**: a BGP route
	//
	// example:
	//
	// custom
	RouteEntryType *string `json:"RouteEntryType,omitempty" xml:"RouteEntryType,omitempty"`
	// The ID of the IPsec-VPN connection.
	//
	// This parameter is required.
	//
	// example:
	//
	// vco-p0w2jpkhi2eeop6q6****
	VpnConnectionId *string `json:"VpnConnectionId,omitempty" xml:"VpnConnectionId,omitempty"`
}

func (s DescribeVcoRouteEntriesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVcoRouteEntriesRequest) GoString() string {
	return s.String()
}

func (s *DescribeVcoRouteEntriesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeVcoRouteEntriesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeVcoRouteEntriesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeVcoRouteEntriesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeVcoRouteEntriesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeVcoRouteEntriesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeVcoRouteEntriesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeVcoRouteEntriesRequest) GetRouteEntryType() *string {
	return s.RouteEntryType
}

func (s *DescribeVcoRouteEntriesRequest) GetVpnConnectionId() *string {
	return s.VpnConnectionId
}

func (s *DescribeVcoRouteEntriesRequest) SetClientToken(v string) *DescribeVcoRouteEntriesRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeVcoRouteEntriesRequest) SetOwnerAccount(v string) *DescribeVcoRouteEntriesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeVcoRouteEntriesRequest) SetPageNumber(v int32) *DescribeVcoRouteEntriesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeVcoRouteEntriesRequest) SetPageSize(v int32) *DescribeVcoRouteEntriesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeVcoRouteEntriesRequest) SetRegionId(v string) *DescribeVcoRouteEntriesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeVcoRouteEntriesRequest) SetResourceOwnerAccount(v string) *DescribeVcoRouteEntriesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeVcoRouteEntriesRequest) SetResourceOwnerId(v int64) *DescribeVcoRouteEntriesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeVcoRouteEntriesRequest) SetRouteEntryType(v string) *DescribeVcoRouteEntriesRequest {
	s.RouteEntryType = &v
	return s
}

func (s *DescribeVcoRouteEntriesRequest) SetVpnConnectionId(v string) *DescribeVcoRouteEntriesRequest {
	s.VpnConnectionId = &v
	return s
}

func (s *DescribeVcoRouteEntriesRequest) Validate() error {
	return dara.Validate(s)
}
