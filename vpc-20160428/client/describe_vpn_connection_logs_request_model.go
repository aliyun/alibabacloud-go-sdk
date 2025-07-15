// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpnConnectionLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFrom(v int32) *DescribeVpnConnectionLogsRequest
	GetFrom() *int32
	SetMinutePeriod(v int32) *DescribeVpnConnectionLogsRequest
	GetMinutePeriod() *int32
	SetOwnerAccount(v string) *DescribeVpnConnectionLogsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeVpnConnectionLogsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeVpnConnectionLogsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeVpnConnectionLogsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeVpnConnectionLogsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeVpnConnectionLogsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeVpnConnectionLogsRequest
	GetResourceOwnerId() *int64
	SetTo(v int32) *DescribeVpnConnectionLogsRequest
	GetTo() *int32
	SetTunnelId(v string) *DescribeVpnConnectionLogsRequest
	GetTunnelId() *string
	SetVpnConnectionId(v string) *DescribeVpnConnectionLogsRequest
	GetVpnConnectionId() *string
}

type DescribeVpnConnectionLogsRequest struct {
	// The start time of the flow log. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// >  If you specify **From**, you must specify **To*	- or **MinutePeriod**.
	//
	// example:
	//
	// 1671003744
	From *int32 `json:"From,omitempty" xml:"From,omitempty"`
	// The interval at which log data is collected. Valid values: **1*	- to **10**. Unit: minutes.
	//
	// >  If you do not specify **From*	- and **To**, you must specify **MinutePeriod**.
	//
	// example:
	//
	// 10
	MinutePeriod *int32  `json:"MinutePeriod,omitempty" xml:"MinutePeriod,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: **1*	- to **50**. Default value: **10**.
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
	// eu-central-1
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The end time of the flow log. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// >  If you specify **To**, you must specify **From*	- or **MinutePeriod**.
	//
	// example:
	//
	// 1671004344
	To *int32 `json:"To,omitempty" xml:"To,omitempty"`
	// The ID of the IPsec-VPN connection.
	//
	// This parameter is available only for a dual-tunnel IPsec-VPN connection.
	//
	// example:
	//
	// tun-opsqc4d97wni27****
	TunnelId *string `json:"TunnelId,omitempty" xml:"TunnelId,omitempty"`
	// The ID of the IPsec-VPN connection.
	//
	// This parameter is required.
	//
	// example:
	//
	// vco-m5evqnds4y459flt3****
	VpnConnectionId *string `json:"VpnConnectionId,omitempty" xml:"VpnConnectionId,omitempty"`
}

func (s DescribeVpnConnectionLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnConnectionLogsRequest) GoString() string {
	return s.String()
}

func (s *DescribeVpnConnectionLogsRequest) GetFrom() *int32 {
	return s.From
}

func (s *DescribeVpnConnectionLogsRequest) GetMinutePeriod() *int32 {
	return s.MinutePeriod
}

func (s *DescribeVpnConnectionLogsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeVpnConnectionLogsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVpnConnectionLogsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeVpnConnectionLogsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeVpnConnectionLogsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeVpnConnectionLogsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeVpnConnectionLogsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeVpnConnectionLogsRequest) GetTo() *int32 {
	return s.To
}

func (s *DescribeVpnConnectionLogsRequest) GetTunnelId() *string {
	return s.TunnelId
}

func (s *DescribeVpnConnectionLogsRequest) GetVpnConnectionId() *string {
	return s.VpnConnectionId
}

func (s *DescribeVpnConnectionLogsRequest) SetFrom(v int32) *DescribeVpnConnectionLogsRequest {
	s.From = &v
	return s
}

func (s *DescribeVpnConnectionLogsRequest) SetMinutePeriod(v int32) *DescribeVpnConnectionLogsRequest {
	s.MinutePeriod = &v
	return s
}

func (s *DescribeVpnConnectionLogsRequest) SetOwnerAccount(v string) *DescribeVpnConnectionLogsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeVpnConnectionLogsRequest) SetOwnerId(v int64) *DescribeVpnConnectionLogsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVpnConnectionLogsRequest) SetPageNumber(v int32) *DescribeVpnConnectionLogsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeVpnConnectionLogsRequest) SetPageSize(v int32) *DescribeVpnConnectionLogsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeVpnConnectionLogsRequest) SetRegionId(v string) *DescribeVpnConnectionLogsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeVpnConnectionLogsRequest) SetResourceOwnerAccount(v string) *DescribeVpnConnectionLogsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeVpnConnectionLogsRequest) SetResourceOwnerId(v int64) *DescribeVpnConnectionLogsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeVpnConnectionLogsRequest) SetTo(v int32) *DescribeVpnConnectionLogsRequest {
	s.To = &v
	return s
}

func (s *DescribeVpnConnectionLogsRequest) SetTunnelId(v string) *DescribeVpnConnectionLogsRequest {
	s.TunnelId = &v
	return s
}

func (s *DescribeVpnConnectionLogsRequest) SetVpnConnectionId(v string) *DescribeVpnConnectionLogsRequest {
	s.VpnConnectionId = &v
	return s
}

func (s *DescribeVpnConnectionLogsRequest) Validate() error {
	return dara.Validate(s)
}
