// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSslVpnServersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeSslVpnServersResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSslVpnServersResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeSslVpnServersResponseBody
	GetRequestId() *string
	SetSslVpnServers(v *DescribeSslVpnServersResponseBodySslVpnServers) *DescribeSslVpnServersResponseBody
	GetSslVpnServers() *DescribeSslVpnServersResponseBodySslVpnServers
	SetTotalCount(v int32) *DescribeSslVpnServersResponseBody
	GetTotalCount() *int32
}

type DescribeSslVpnServersResponseBody struct {
	// The number of the page to return.
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
	// The request ID.
	//
	// example:
	//
	// D350187B-EA41-4577-950B-95434C8302E1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The detailed information about the SSL-VPN servers.
	SslVpnServers *DescribeSslVpnServersResponseBodySslVpnServers `json:"SslVpnServers,omitempty" xml:"SslVpnServers,omitempty" type:"Struct"`
	// The number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSslVpnServersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSslVpnServersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSslVpnServersResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSslVpnServersResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSslVpnServersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSslVpnServersResponseBody) GetSslVpnServers() *DescribeSslVpnServersResponseBodySslVpnServers {
	return s.SslVpnServers
}

func (s *DescribeSslVpnServersResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeSslVpnServersResponseBody) SetPageNumber(v int32) *DescribeSslVpnServersResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSslVpnServersResponseBody) SetPageSize(v int32) *DescribeSslVpnServersResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSslVpnServersResponseBody) SetRequestId(v string) *DescribeSslVpnServersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSslVpnServersResponseBody) SetSslVpnServers(v *DescribeSslVpnServersResponseBodySslVpnServers) *DescribeSslVpnServersResponseBody {
	s.SslVpnServers = v
	return s
}

func (s *DescribeSslVpnServersResponseBody) SetTotalCount(v int32) *DescribeSslVpnServersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeSslVpnServersResponseBody) Validate() error {
	if s.SslVpnServers != nil {
		if err := s.SslVpnServers.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSslVpnServersResponseBodySslVpnServers struct {
	SslVpnServer []*DescribeSslVpnServersResponseBodySslVpnServersSslVpnServer `json:"SslVpnServer,omitempty" xml:"SslVpnServer,omitempty" type:"Repeated"`
}

func (s DescribeSslVpnServersResponseBodySslVpnServers) String() string {
	return dara.Prettify(s)
}

func (s DescribeSslVpnServersResponseBodySslVpnServers) GoString() string {
	return s.String()
}

func (s *DescribeSslVpnServersResponseBodySslVpnServers) GetSslVpnServer() []*DescribeSslVpnServersResponseBodySslVpnServersSslVpnServer {
	return s.SslVpnServer
}

func (s *DescribeSslVpnServersResponseBodySslVpnServers) SetSslVpnServer(v []*DescribeSslVpnServersResponseBodySslVpnServersSslVpnServer) *DescribeSslVpnServersResponseBodySslVpnServers {
	s.SslVpnServer = v
	return s
}

func (s *DescribeSslVpnServersResponseBodySslVpnServers) Validate() error {
	if s.SslVpnServer != nil {
		for _, item := range s.SslVpnServer {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSslVpnServersResponseBodySslVpnServersSslVpnServer struct {
	// The encryption algorithm.
	//
	// example:
	//
	// AES-128-CBC
	Cipher *string `json:"Cipher,omitempty" xml:"Cipher,omitempty"`
	// The client CIDR block.
	//
	// example:
	//
	// 10.10.1.0/24
	ClientIpPool *string `json:"ClientIpPool,omitempty" xml:"ClientIpPool,omitempty"`
	// Indicates whether data compression is enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// example:
	//
	// false
	Compress *bool `json:"Compress,omitempty" xml:"Compress,omitempty"`
	// The total number of current connections.
	//
	// example:
	//
	// 0
	Connections *int32 `json:"Connections,omitempty" xml:"Connections,omitempty"`
	// The timestamp generated when the SSL-VPN server was created.
	//
	// example:
	//
	// 1613800884000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Indicates whether two-factor authentication is enabled.
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// example:
	//
	// true
	EnableMultiFactorAuth *bool `json:"EnableMultiFactorAuth,omitempty" xml:"EnableMultiFactorAuth,omitempty"`
	// The ID of the IDaaS application.
	//
	// example:
	//
	// app_my6g4qmvnwxzj2f****
	IDaaSApplicationId *string `json:"IDaaSApplicationId,omitempty" xml:"IDaaSApplicationId,omitempty"`
	// The ID of the IDaaS EIAM instance.
	//
	// example:
	//
	// idaas-cn-hangzhou-****
	IDaaSInstanceId *string `json:"IDaaSInstanceId,omitempty" xml:"IDaaSInstanceId,omitempty"`
	// The version of the IDaaS EIAM instance.
	//
	// 	- This parameter is returned only if the SSL server is associated with an IDaaS EIAM 2.0 instance. Only **EIAM 2.0*	- is returned.
	//
	// 	- If the SSL server is associated with an IDaaS EIAM 1.0 instance, no value is returned.
	//
	// example:
	//
	// EIAM 2.0
	IDaaSInstanceVersion *string `json:"IDaaSInstanceVersion,omitempty" xml:"IDaaSInstanceVersion,omitempty"`
	// The region ID of the IDaaS EIAM instance.
	//
	// example:
	//
	// cn-hangzhou
	IDaaSRegionId *string `json:"IDaaSRegionId,omitempty" xml:"IDaaSRegionId,omitempty"`
	// The public IP address of the VPN gateway.
	//
	// example:
	//
	// 47.5.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The local CIDR block.
	//
	// example:
	//
	// 192.168.0.0/24
	LocalSubnet *string `json:"LocalSubnet,omitempty" xml:"LocalSubnet,omitempty"`
	// The maximum number of connections.
	//
	// example:
	//
	// 5
	MaxConnections *int32 `json:"MaxConnections,omitempty" xml:"MaxConnections,omitempty"`
	// The name of the SSL server.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The port that is used by the SSL-VPN server.
	//
	// example:
	//
	// 1194
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The protocol that is used by the SSL server.
	//
	// example:
	//
	// UDP
	Proto *string `json:"Proto,omitempty" xml:"Proto,omitempty"`
	// The region ID of the SSL server.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID of the SSL server.
	//
	// You can call the [ListResourceGroups](https://help.aliyun.com/document_detail/158855.html) operation to query resource groups.
	//
	// example:
	//
	// rg-acfmzs372yg****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the SSL server.
	//
	// example:
	//
	// vss-bp15j3du13gq1dgey****
	SslVpnServerId *string `json:"SslVpnServerId,omitempty" xml:"SslVpnServerId,omitempty"`
	// The ID of the VPN gateway.
	//
	// example:
	//
	// vpn-bp1on0xae9d771ggi****
	VpnGatewayId *string `json:"VpnGatewayId,omitempty" xml:"VpnGatewayId,omitempty"`
}

func (s DescribeSslVpnServersResponseBodySslVpnServersSslVpnServer) String() string {
	return dara.Prettify(s)
}

func (s DescribeSslVpnServersResponseBodySslVpnServersSslVpnServer) GoString() string {
	return s.String()
}

func (s *DescribeSslVpnServersResponseBodySslVpnServersSslVpnServer) GetCipher() *string {
	return s.Cipher
}

func (s *DescribeSslVpnServersResponseBodySslVpnServersSslVpnServer) GetClientIpPool() *string {
	return s.ClientIpPool
}

func (s *DescribeSslVpnServersResponseBodySslVpnServersSslVpnServer) GetCompress() *bool {
	return s.Compress
}

func (s *DescribeSslVpnServersResponseBodySslVpnServersSslVpnServer) GetConnections() *int32 {
	return s.Connections
}

func (s *DescribeSslVpnServersResponseBodySslVpnServersSslVpnServer) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DescribeSslVpnServersResponseBodySslVpnServersSslVpnServer) GetEnableMultiFactorAuth() *bool {
	return s.EnableMultiFactorAuth
}

func (s *DescribeSslVpnServersResponseBodySslVpnServersSslVpnServer) GetIDaaSApplicationId() *string {
	return s.IDaaSApplicationId
}

func (s *DescribeSslVpnServersResponseBodySslVpnServersSslVpnServer) GetIDaaSInstanceId() *string {
	return s.IDaaSInstanceId
}

func (s *DescribeSslVpnServersResponseBodySslVpnServersSslVpnServer) GetIDaaSInstanceVersion() *string {
	return s.IDaaSInstanceVersion
}

func (s *DescribeSslVpnServersResponseBodySslVpnServersSslVpnServer) GetIDaaSRegionId() *string {
	return s.IDaaSRegionId
}

func (s *DescribeSslVpnServersResponseBodySslVpnServersSslVpnServer) GetInternetIp() *string {
	return s.InternetIp
}

func (s *DescribeSslVpnServersResponseBodySslVpnServersSslVpnServer) GetLocalSubnet() *string {
	return s.LocalSubnet
}

func (s *DescribeSslVpnServersResponseBodySslVpnServersSslVpnServer) GetMaxConnections() *int32 {
	return s.MaxConnections
}

func (s *DescribeSslVpnServersResponseBodySslVpnServersSslVpnServer) GetName() *string {
	return s.Name
}

func (s *DescribeSslVpnServersResponseBodySslVpnServersSslVpnServer) GetPort() *int32 {
	return s.Port
}

func (s *DescribeSslVpnServersResponseBodySslVpnServersSslVpnServer) GetProto() *string {
	return s.Proto
}

func (s *DescribeSslVpnServersResponseBodySslVpnServersSslVpnServer) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSslVpnServersResponseBodySslVpnServersSslVpnServer) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeSslVpnServersResponseBodySslVpnServersSslVpnServer) GetSslVpnServerId() *string {
	return s.SslVpnServerId
}

func (s *DescribeSslVpnServersResponseBodySslVpnServersSslVpnServer) GetVpnGatewayId() *string {
	return s.VpnGatewayId
}

func (s *DescribeSslVpnServersResponseBodySslVpnServersSslVpnServer) SetCipher(v string) *DescribeSslVpnServersResponseBodySslVpnServersSslVpnServer {
	s.Cipher = &v
	return s
}

func (s *DescribeSslVpnServersResponseBodySslVpnServersSslVpnServer) SetClientIpPool(v string) *DescribeSslVpnServersResponseBodySslVpnServersSslVpnServer {
	s.ClientIpPool = &v
	return s
}

func (s *DescribeSslVpnServersResponseBodySslVpnServersSslVpnServer) SetCompress(v bool) *DescribeSslVpnServersResponseBodySslVpnServersSslVpnServer {
	s.Compress = &v
	return s
}

func (s *DescribeSslVpnServersResponseBodySslVpnServersSslVpnServer) SetConnections(v int32) *DescribeSslVpnServersResponseBodySslVpnServersSslVpnServer {
	s.Connections = &v
	return s
}

func (s *DescribeSslVpnServersResponseBodySslVpnServersSslVpnServer) SetCreateTime(v int64) *DescribeSslVpnServersResponseBodySslVpnServersSslVpnServer {
	s.CreateTime = &v
	return s
}

func (s *DescribeSslVpnServersResponseBodySslVpnServersSslVpnServer) SetEnableMultiFactorAuth(v bool) *DescribeSslVpnServersResponseBodySslVpnServersSslVpnServer {
	s.EnableMultiFactorAuth = &v
	return s
}

func (s *DescribeSslVpnServersResponseBodySslVpnServersSslVpnServer) SetIDaaSApplicationId(v string) *DescribeSslVpnServersResponseBodySslVpnServersSslVpnServer {
	s.IDaaSApplicationId = &v
	return s
}

func (s *DescribeSslVpnServersResponseBodySslVpnServersSslVpnServer) SetIDaaSInstanceId(v string) *DescribeSslVpnServersResponseBodySslVpnServersSslVpnServer {
	s.IDaaSInstanceId = &v
	return s
}

func (s *DescribeSslVpnServersResponseBodySslVpnServersSslVpnServer) SetIDaaSInstanceVersion(v string) *DescribeSslVpnServersResponseBodySslVpnServersSslVpnServer {
	s.IDaaSInstanceVersion = &v
	return s
}

func (s *DescribeSslVpnServersResponseBodySslVpnServersSslVpnServer) SetIDaaSRegionId(v string) *DescribeSslVpnServersResponseBodySslVpnServersSslVpnServer {
	s.IDaaSRegionId = &v
	return s
}

func (s *DescribeSslVpnServersResponseBodySslVpnServersSslVpnServer) SetInternetIp(v string) *DescribeSslVpnServersResponseBodySslVpnServersSslVpnServer {
	s.InternetIp = &v
	return s
}

func (s *DescribeSslVpnServersResponseBodySslVpnServersSslVpnServer) SetLocalSubnet(v string) *DescribeSslVpnServersResponseBodySslVpnServersSslVpnServer {
	s.LocalSubnet = &v
	return s
}

func (s *DescribeSslVpnServersResponseBodySslVpnServersSslVpnServer) SetMaxConnections(v int32) *DescribeSslVpnServersResponseBodySslVpnServersSslVpnServer {
	s.MaxConnections = &v
	return s
}

func (s *DescribeSslVpnServersResponseBodySslVpnServersSslVpnServer) SetName(v string) *DescribeSslVpnServersResponseBodySslVpnServersSslVpnServer {
	s.Name = &v
	return s
}

func (s *DescribeSslVpnServersResponseBodySslVpnServersSslVpnServer) SetPort(v int32) *DescribeSslVpnServersResponseBodySslVpnServersSslVpnServer {
	s.Port = &v
	return s
}

func (s *DescribeSslVpnServersResponseBodySslVpnServersSslVpnServer) SetProto(v string) *DescribeSslVpnServersResponseBodySslVpnServersSslVpnServer {
	s.Proto = &v
	return s
}

func (s *DescribeSslVpnServersResponseBodySslVpnServersSslVpnServer) SetRegionId(v string) *DescribeSslVpnServersResponseBodySslVpnServersSslVpnServer {
	s.RegionId = &v
	return s
}

func (s *DescribeSslVpnServersResponseBodySslVpnServersSslVpnServer) SetResourceGroupId(v string) *DescribeSslVpnServersResponseBodySslVpnServersSslVpnServer {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeSslVpnServersResponseBodySslVpnServersSslVpnServer) SetSslVpnServerId(v string) *DescribeSslVpnServersResponseBodySslVpnServersSslVpnServer {
	s.SslVpnServerId = &v
	return s
}

func (s *DescribeSslVpnServersResponseBodySslVpnServersSslVpnServer) SetVpnGatewayId(v string) *DescribeSslVpnServersResponseBodySslVpnServersSslVpnServer {
	s.VpnGatewayId = &v
	return s
}

func (s *DescribeSslVpnServersResponseBodySslVpnServersSslVpnServer) Validate() error {
	return dara.Validate(s)
}
