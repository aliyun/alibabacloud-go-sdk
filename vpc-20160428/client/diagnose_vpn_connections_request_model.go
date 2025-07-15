// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDiagnoseVpnConnectionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DiagnoseVpnConnectionsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DiagnoseVpnConnectionsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DiagnoseVpnConnectionsRequest
	GetRegionId() *string
	SetResourceOwnerId(v int64) *DiagnoseVpnConnectionsRequest
	GetResourceOwnerId() *int64
	SetTunnelIds(v []*string) *DiagnoseVpnConnectionsRequest
	GetTunnelIds() []*string
	SetVpnConnectionIds(v []*string) *DiagnoseVpnConnectionsRequest
	GetVpnConnectionIds() []*string
	SetVpnGatewayId(v string) *DiagnoseVpnConnectionsRequest
	GetVpnGatewayId() *string
}

type DiagnoseVpnConnectionsRequest struct {
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the IPsec-VPN connection.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-qingdao
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The list of tunnel IDs.
	TunnelIds []*string `json:"TunnelIds,omitempty" xml:"TunnelIds,omitempty" type:"Repeated"`
	// The IDs of IPsec-VPN connections.
	VpnConnectionIds []*string `json:"VpnConnectionIds,omitempty" xml:"VpnConnectionIds,omitempty" type:"Repeated"`
	// The ID of the VPN gateway.
	//
	// example:
	//
	// vpn-bp10hz6b0mbp39flt****
	VpnGatewayId *string `json:"VpnGatewayId,omitempty" xml:"VpnGatewayId,omitempty"`
}

func (s DiagnoseVpnConnectionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DiagnoseVpnConnectionsRequest) GoString() string {
	return s.String()
}

func (s *DiagnoseVpnConnectionsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DiagnoseVpnConnectionsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DiagnoseVpnConnectionsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DiagnoseVpnConnectionsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DiagnoseVpnConnectionsRequest) GetTunnelIds() []*string {
	return s.TunnelIds
}

func (s *DiagnoseVpnConnectionsRequest) GetVpnConnectionIds() []*string {
	return s.VpnConnectionIds
}

func (s *DiagnoseVpnConnectionsRequest) GetVpnGatewayId() *string {
	return s.VpnGatewayId
}

func (s *DiagnoseVpnConnectionsRequest) SetPageNumber(v int32) *DiagnoseVpnConnectionsRequest {
	s.PageNumber = &v
	return s
}

func (s *DiagnoseVpnConnectionsRequest) SetPageSize(v int32) *DiagnoseVpnConnectionsRequest {
	s.PageSize = &v
	return s
}

func (s *DiagnoseVpnConnectionsRequest) SetRegionId(v string) *DiagnoseVpnConnectionsRequest {
	s.RegionId = &v
	return s
}

func (s *DiagnoseVpnConnectionsRequest) SetResourceOwnerId(v int64) *DiagnoseVpnConnectionsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DiagnoseVpnConnectionsRequest) SetTunnelIds(v []*string) *DiagnoseVpnConnectionsRequest {
	s.TunnelIds = v
	return s
}

func (s *DiagnoseVpnConnectionsRequest) SetVpnConnectionIds(v []*string) *DiagnoseVpnConnectionsRequest {
	s.VpnConnectionIds = v
	return s
}

func (s *DiagnoseVpnConnectionsRequest) SetVpnGatewayId(v string) *DiagnoseVpnConnectionsRequest {
	s.VpnGatewayId = &v
	return s
}

func (s *DiagnoseVpnConnectionsRequest) Validate() error {
	return dara.Validate(s)
}
