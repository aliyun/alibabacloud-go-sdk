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
	RequestId     *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Cipher                *string `json:"Cipher,omitempty" xml:"Cipher,omitempty"`
	ClientIpPool          *string `json:"ClientIpPool,omitempty" xml:"ClientIpPool,omitempty"`
	Compress              *bool   `json:"Compress,omitempty" xml:"Compress,omitempty"`
	Connections           *int32  `json:"Connections,omitempty" xml:"Connections,omitempty"`
	CreateTime            *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	EnableMultiFactorAuth *bool   `json:"EnableMultiFactorAuth,omitempty" xml:"EnableMultiFactorAuth,omitempty"`
	IDaaSApplicationId    *string `json:"IDaaSApplicationId,omitempty" xml:"IDaaSApplicationId,omitempty"`
	IDaaSInstanceId       *string `json:"IDaaSInstanceId,omitempty" xml:"IDaaSInstanceId,omitempty"`
	IDaaSInstanceVersion  *string `json:"IDaaSInstanceVersion,omitempty" xml:"IDaaSInstanceVersion,omitempty"`
	IDaaSRegionId         *string `json:"IDaaSRegionId,omitempty" xml:"IDaaSRegionId,omitempty"`
	InternetIp            *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	LocalSubnet           *string `json:"LocalSubnet,omitempty" xml:"LocalSubnet,omitempty"`
	MaxConnections        *int32  `json:"MaxConnections,omitempty" xml:"MaxConnections,omitempty"`
	Name                  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Port                  *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Proto                 *string `json:"Proto,omitempty" xml:"Proto,omitempty"`
	RegionId              *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId       *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SslVpnServerId        *string `json:"SslVpnServerId,omitempty" xml:"SslVpnServerId,omitempty"`
	VpnGatewayId          *string `json:"VpnGatewayId,omitempty" xml:"VpnGatewayId,omitempty"`
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
