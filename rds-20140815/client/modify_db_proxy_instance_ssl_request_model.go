// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDbProxyInstanceSslRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBProxyEngineType(v string) *ModifyDbProxyInstanceSslRequest
	GetDBProxyEngineType() *string
	SetDbInstanceId(v string) *ModifyDbProxyInstanceSslRequest
	GetDbInstanceId() *string
	SetDbProxyConnectString(v string) *ModifyDbProxyInstanceSslRequest
	GetDbProxyConnectString() *string
	SetDbProxyEndpointId(v string) *ModifyDbProxyInstanceSslRequest
	GetDbProxyEndpointId() *string
	SetDbProxySslEnabled(v string) *ModifyDbProxyInstanceSslRequest
	GetDbProxySslEnabled() *string
	SetRegionId(v string) *ModifyDbProxyInstanceSslRequest
	GetRegionId() *string
}

type ModifyDbProxyInstanceSslRequest struct {
	// A reserved parameter. You do not need to specify this parameter.
	//
	// example:
	//
	// normal
	DBProxyEngineType *string `json:"DBProxyEngineType,omitempty" xml:"DBProxyEngineType,omitempty"`
	// The instance ID. You can call DescribeDBInstances to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-t4n3a****
	DbInstanceId *string `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
	// The endpoint for which you want to enable SSL encryption.
	//
	// This parameter is required.
	//
	// example:
	//
	// test123456.rwlb.rds.aliyuncs.com
	DbProxyConnectString *string `json:"DbProxyConnectString,omitempty" xml:"DbProxyConnectString,omitempty"`
	// The ID of the database proxy endpoint. You can call DescribeDBProxyEndpoint to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ta9um4****
	DbProxyEndpointId *string `json:"DbProxyEndpointId,omitempty" xml:"DbProxyEndpointId,omitempty"`
	// The operation that you want to perform on SSL encryption. Valid values:
	//
	// 	- 0: Disables SSL encryption.
	//
	// 	- 1: Enables SSL encryption or changes the endpoint for which SSL encryption is enabled.
	//
	// 	- 2: Updates the validity period of the SSL certificate.
	//
	// >The preceding operations restart the instance. Proceed with caution.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DbProxySslEnabled *string `json:"DbProxySslEnabled,omitempty" xml:"DbProxySslEnabled,omitempty"`
	// The region ID. You can call DescribeRegions to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyDbProxyInstanceSslRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDbProxyInstanceSslRequest) GoString() string {
	return s.String()
}

func (s *ModifyDbProxyInstanceSslRequest) GetDBProxyEngineType() *string {
	return s.DBProxyEngineType
}

func (s *ModifyDbProxyInstanceSslRequest) GetDbInstanceId() *string {
	return s.DbInstanceId
}

func (s *ModifyDbProxyInstanceSslRequest) GetDbProxyConnectString() *string {
	return s.DbProxyConnectString
}

func (s *ModifyDbProxyInstanceSslRequest) GetDbProxyEndpointId() *string {
	return s.DbProxyEndpointId
}

func (s *ModifyDbProxyInstanceSslRequest) GetDbProxySslEnabled() *string {
	return s.DbProxySslEnabled
}

func (s *ModifyDbProxyInstanceSslRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDbProxyInstanceSslRequest) SetDBProxyEngineType(v string) *ModifyDbProxyInstanceSslRequest {
	s.DBProxyEngineType = &v
	return s
}

func (s *ModifyDbProxyInstanceSslRequest) SetDbInstanceId(v string) *ModifyDbProxyInstanceSslRequest {
	s.DbInstanceId = &v
	return s
}

func (s *ModifyDbProxyInstanceSslRequest) SetDbProxyConnectString(v string) *ModifyDbProxyInstanceSslRequest {
	s.DbProxyConnectString = &v
	return s
}

func (s *ModifyDbProxyInstanceSslRequest) SetDbProxyEndpointId(v string) *ModifyDbProxyInstanceSslRequest {
	s.DbProxyEndpointId = &v
	return s
}

func (s *ModifyDbProxyInstanceSslRequest) SetDbProxySslEnabled(v string) *ModifyDbProxyInstanceSslRequest {
	s.DbProxySslEnabled = &v
	return s
}

func (s *ModifyDbProxyInstanceSslRequest) SetRegionId(v string) *ModifyDbProxyInstanceSslRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDbProxyInstanceSslRequest) Validate() error {
	return dara.Validate(s)
}
