// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSslVpnClientsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClientInfoList(v []*DescribeSslVpnClientsResponseBodyClientInfoList) *DescribeSslVpnClientsResponseBody
	GetClientInfoList() []*DescribeSslVpnClientsResponseBodyClientInfoList
	SetPageNumber(v int32) *DescribeSslVpnClientsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSslVpnClientsResponseBody
	GetPageSize() *int32
	SetRegionId(v string) *DescribeSslVpnClientsResponseBody
	GetRegionId() *string
	SetRequestId(v string) *DescribeSslVpnClientsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeSslVpnClientsResponseBody
	GetTotalCount() *int32
	SetVpnGatewayId(v string) *DescribeSslVpnClientsResponseBody
	GetVpnGatewayId() *string
}

type DescribeSslVpnClientsResponseBody struct {
	// The list of clients.
	ClientInfoList []*DescribeSslVpnClientsResponseBodyClientInfoList `json:"ClientInfoList,omitempty" xml:"ClientInfoList,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// example:
	//
	// eu-central-1
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 885E117D-06A9-38A3-8DD2-40BDAC429FFC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of returned entries.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The ID of the VPN gateway.
	//
	// example:
	//
	// vpn-gw8gfb947ctddabja****
	VpnGatewayId *string `json:"VpnGatewayId,omitempty" xml:"VpnGatewayId,omitempty"`
}

func (s DescribeSslVpnClientsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSslVpnClientsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSslVpnClientsResponseBody) GetClientInfoList() []*DescribeSslVpnClientsResponseBodyClientInfoList {
	return s.ClientInfoList
}

func (s *DescribeSslVpnClientsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSslVpnClientsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSslVpnClientsResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSslVpnClientsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSslVpnClientsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeSslVpnClientsResponseBody) GetVpnGatewayId() *string {
	return s.VpnGatewayId
}

func (s *DescribeSslVpnClientsResponseBody) SetClientInfoList(v []*DescribeSslVpnClientsResponseBodyClientInfoList) *DescribeSslVpnClientsResponseBody {
	s.ClientInfoList = v
	return s
}

func (s *DescribeSslVpnClientsResponseBody) SetPageNumber(v int32) *DescribeSslVpnClientsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSslVpnClientsResponseBody) SetPageSize(v int32) *DescribeSslVpnClientsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSslVpnClientsResponseBody) SetRegionId(v string) *DescribeSslVpnClientsResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeSslVpnClientsResponseBody) SetRequestId(v string) *DescribeSslVpnClientsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSslVpnClientsResponseBody) SetTotalCount(v int32) *DescribeSslVpnClientsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeSslVpnClientsResponseBody) SetVpnGatewayId(v string) *DescribeSslVpnClientsResponseBody {
	s.VpnGatewayId = &v
	return s
}

func (s *DescribeSslVpnClientsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeSslVpnClientsResponseBodyClientInfoList struct {
	// The SSL client certificate used by the client.
	//
	// >  If the client uses two-factor authentication to establish an SSL-VPN connection to Alibaba Cloud, the return value is the username of the client.
	//
	// example:
	//
	// CN=vsc-gw8gkh6gtilf1whgc****
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// The timestamp that indicates when the client connected to Alibaba Cloud through an SSL-VPN connection. Unit: milliseconds.
	//
	// This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1670985008000
	ConnectedTime *int64 `json:"ConnectedTime,omitempty" xml:"ConnectedTime,omitempty"`
	// The actual public IP address used by the client when the client established an SSL-VPN connection to Alibaba Cloud.
	//
	// example:
	//
	// 8.XX.XX.15
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The port used by the client when the client established an SSL-VPN connection to Alibaba Cloud.
	//
	// example:
	//
	// 4****
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The private IP address allocated to the client by the VPN gateway when the client established an SSL-VPN connection to Alibaba Cloud.
	//
	// example:
	//
	// 10.10.10.10
	PrivateIp *string `json:"PrivateIp,omitempty" xml:"PrivateIp,omitempty"`
	// The amount of data transferred from the client to the VPN gateway through the SSL-VPN connection. Unit: bytes.
	//
	// example:
	//
	// 60782
	ReceiveBytes *int64 `json:"ReceiveBytes,omitempty" xml:"ReceiveBytes,omitempty"`
	// The amount of data transferred from the VPN gateway to the client through the SSL-VPN connection. Unit: bytes.
	//
	// example:
	//
	// 57144
	SendBytes *int64 `json:"SendBytes,omitempty" xml:"SendBytes,omitempty"`
	// The status of the SSL-VPN connection.
	//
	// The value is set to **online**, which indicates that the client has connected to Alibaba Cloud through an SSL-VPN connection.
	//
	// example:
	//
	// online
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeSslVpnClientsResponseBodyClientInfoList) String() string {
	return dara.Prettify(s)
}

func (s DescribeSslVpnClientsResponseBodyClientInfoList) GoString() string {
	return s.String()
}

func (s *DescribeSslVpnClientsResponseBodyClientInfoList) GetCommonName() *string {
	return s.CommonName
}

func (s *DescribeSslVpnClientsResponseBodyClientInfoList) GetConnectedTime() *int64 {
	return s.ConnectedTime
}

func (s *DescribeSslVpnClientsResponseBodyClientInfoList) GetIp() *string {
	return s.Ip
}

func (s *DescribeSslVpnClientsResponseBodyClientInfoList) GetPort() *string {
	return s.Port
}

func (s *DescribeSslVpnClientsResponseBodyClientInfoList) GetPrivateIp() *string {
	return s.PrivateIp
}

func (s *DescribeSslVpnClientsResponseBodyClientInfoList) GetReceiveBytes() *int64 {
	return s.ReceiveBytes
}

func (s *DescribeSslVpnClientsResponseBodyClientInfoList) GetSendBytes() *int64 {
	return s.SendBytes
}

func (s *DescribeSslVpnClientsResponseBodyClientInfoList) GetStatus() *string {
	return s.Status
}

func (s *DescribeSslVpnClientsResponseBodyClientInfoList) SetCommonName(v string) *DescribeSslVpnClientsResponseBodyClientInfoList {
	s.CommonName = &v
	return s
}

func (s *DescribeSslVpnClientsResponseBodyClientInfoList) SetConnectedTime(v int64) *DescribeSslVpnClientsResponseBodyClientInfoList {
	s.ConnectedTime = &v
	return s
}

func (s *DescribeSslVpnClientsResponseBodyClientInfoList) SetIp(v string) *DescribeSslVpnClientsResponseBodyClientInfoList {
	s.Ip = &v
	return s
}

func (s *DescribeSslVpnClientsResponseBodyClientInfoList) SetPort(v string) *DescribeSslVpnClientsResponseBodyClientInfoList {
	s.Port = &v
	return s
}

func (s *DescribeSslVpnClientsResponseBodyClientInfoList) SetPrivateIp(v string) *DescribeSslVpnClientsResponseBodyClientInfoList {
	s.PrivateIp = &v
	return s
}

func (s *DescribeSslVpnClientsResponseBodyClientInfoList) SetReceiveBytes(v int64) *DescribeSslVpnClientsResponseBodyClientInfoList {
	s.ReceiveBytes = &v
	return s
}

func (s *DescribeSslVpnClientsResponseBodyClientInfoList) SetSendBytes(v int64) *DescribeSslVpnClientsResponseBodyClientInfoList {
	s.SendBytes = &v
	return s
}

func (s *DescribeSslVpnClientsResponseBodyClientInfoList) SetStatus(v string) *DescribeSslVpnClientsResponseBodyClientInfoList {
	s.Status = &v
	return s
}

func (s *DescribeSslVpnClientsResponseBodyClientInfoList) Validate() error {
	return dara.Validate(s)
}
