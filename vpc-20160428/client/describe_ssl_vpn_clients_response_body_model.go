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
	// The list of client information.
	ClientInfoList []*DescribeSslVpnClientsResponseBodyClientInfoList `json:"ClientInfoList,omitempty" xml:"ClientInfoList,omitempty" type:"Repeated"`
	// The page number of the list.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The maximum number of entries per page in a paged query.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the VPN gateway instance.
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
	// The total number of entries in the list.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The VPN gateway instance ID.
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
	if s.ClientInfoList != nil {
		for _, item := range s.ClientInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSslVpnClientsResponseBodyClientInfoList struct {
	// The SSL client certificate used when the client establishes an SSL-VPN connection to Alibaba Cloud.
	//
	// > If the client uses two-factor identity authentication to establish an SSL-VPN connection to Alibaba Cloud, the value of this parameter is the username of the client.
	//
	// example:
	//
	// CN=vsc-gw8gkh6gtilf1whgc****
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// The timestamp when the client establishes an SSL-VPN connection to Alibaba Cloud. Unit: milliseconds.
	//
	// The timestamp is in the UNIX format and represents the total duration from 00:00:00 on January 1, 1970 (UTC) to the time when the SSL-VPN connection is established.
	//
	// example:
	//
	// 1670985008000
	ConnectedTime *int64 `json:"ConnectedTime,omitempty" xml:"ConnectedTime,omitempty"`
	// The public IP address that the client uses when establishing an SSL-VPN connection to Alibaba Cloud.
	//
	// example:
	//
	// 8.XX.XX.15
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The port number that the client uses when establishing an SSL-VPN connection to Alibaba Cloud.
	//
	// example:
	//
	// 4****
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The private IP address that the VPN gateway assigns to the client when the client establishes an SSL-VPN connection to Alibaba Cloud.
	//
	// example:
	//
	// 10.10.10.10
	PrivateIp *string `json:"PrivateIp,omitempty" xml:"PrivateIp,omitempty"`
	// The traffic that the VPN gateway receives from the client over the SSL-VPN connection. Unit: bytes.
	//
	// example:
	//
	// 60782
	ReceiveBytes *int64 `json:"ReceiveBytes,omitempty" xml:"ReceiveBytes,omitempty"`
	// The traffic that the VPN gateway sends to the client over the SSL-VPN connection. Unit: bytes.
	//
	// example:
	//
	// 57144
	SendBytes *int64 `json:"SendBytes,omitempty" xml:"SendBytes,omitempty"`
	// The status of the SSL-VPN connection.
	//
	// Valid values: **online**, which indicates that the client has successfully established an SSL-VPN connection to Alibaba Cloud.
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
