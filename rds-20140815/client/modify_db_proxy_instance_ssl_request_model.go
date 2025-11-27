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
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-t4n3axxxxx
	DbInstanceId *string `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
	// The dedicated proxy endpoint of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// test123456.rwlb.rds.aliyuncs.com
	DbProxyConnectString *string `json:"DbProxyConnectString,omitempty" xml:"DbProxyConnectString,omitempty"`
	// The ID of the proxy endpoint. You can call the DescribeDBProxyEndpoint operation to query the ID of the proxy endpoint.
	//
	// This parameter is required.
	//
	// example:
	//
	// ta9um4xxxxx
	DbProxyEndpointId *string `json:"DbProxyEndpointId,omitempty" xml:"DbProxyEndpointId,omitempty"`
	// The SSL configuration setting that you want to apply on the instance. Valid values:
	//
	// 	- 0: disables SSL encryption.
	//
	// 	- 1: enables SSL encryption or modifies the endpoint that requires SSL encryption.
	//
	// 	- 2: updates the validity period of the SSL certificate.
	//
	// > This setting causes your instance to restart. Proceed with caution.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DbProxySslEnabled *string `json:"DbProxySslEnabled,omitempty" xml:"DbProxySslEnabled,omitempty"`
	// The region ID. You can call the DescribeRegions operation to query the most recent region list.
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
