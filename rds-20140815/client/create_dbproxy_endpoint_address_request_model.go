// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBProxyEndpointAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConnectionStringPrefix(v string) *CreateDBProxyEndpointAddressRequest
	GetConnectionStringPrefix() *string
	SetDBInstanceId(v string) *CreateDBProxyEndpointAddressRequest
	GetDBInstanceId() *string
	SetDBProxyConnectStringNetType(v string) *CreateDBProxyEndpointAddressRequest
	GetDBProxyConnectStringNetType() *string
	SetDBProxyEndpointId(v string) *CreateDBProxyEndpointAddressRequest
	GetDBProxyEndpointId() *string
	SetDBProxyEngineType(v string) *CreateDBProxyEndpointAddressRequest
	GetDBProxyEngineType() *string
	SetDBProxyNewConnectStringPort(v string) *CreateDBProxyEndpointAddressRequest
	GetDBProxyNewConnectStringPort() *string
	SetRegionId(v string) *CreateDBProxyEndpointAddressRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateDBProxyEndpointAddressRequest
	GetResourceGroupId() *string
	SetVPCId(v string) *CreateDBProxyEndpointAddressRequest
	GetVPCId() *string
	SetVSwitchId(v string) *CreateDBProxyEndpointAddressRequest
	GetVSwitchId() *string
}

type CreateDBProxyEndpointAddressRequest struct {
	// This parameter is required.
	ConnectionStringPrefix *string `json:"ConnectionStringPrefix,omitempty" xml:"ConnectionStringPrefix,omitempty"`
	// This parameter is required.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// This parameter is required.
	DBProxyConnectStringNetType *string `json:"DBProxyConnectStringNetType,omitempty" xml:"DBProxyConnectStringNetType,omitempty"`
	// This parameter is required.
	DBProxyEndpointId           *string `json:"DBProxyEndpointId,omitempty" xml:"DBProxyEndpointId,omitempty"`
	DBProxyEngineType           *string `json:"DBProxyEngineType,omitempty" xml:"DBProxyEngineType,omitempty"`
	DBProxyNewConnectStringPort *string `json:"DBProxyNewConnectStringPort,omitempty" xml:"DBProxyNewConnectStringPort,omitempty"`
	RegionId                    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId             *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	VPCId                       *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	VSwitchId                   *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s CreateDBProxyEndpointAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDBProxyEndpointAddressRequest) GoString() string {
	return s.String()
}

func (s *CreateDBProxyEndpointAddressRequest) GetConnectionStringPrefix() *string {
	return s.ConnectionStringPrefix
}

func (s *CreateDBProxyEndpointAddressRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateDBProxyEndpointAddressRequest) GetDBProxyConnectStringNetType() *string {
	return s.DBProxyConnectStringNetType
}

func (s *CreateDBProxyEndpointAddressRequest) GetDBProxyEndpointId() *string {
	return s.DBProxyEndpointId
}

func (s *CreateDBProxyEndpointAddressRequest) GetDBProxyEngineType() *string {
	return s.DBProxyEngineType
}

func (s *CreateDBProxyEndpointAddressRequest) GetDBProxyNewConnectStringPort() *string {
	return s.DBProxyNewConnectStringPort
}

func (s *CreateDBProxyEndpointAddressRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDBProxyEndpointAddressRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateDBProxyEndpointAddressRequest) GetVPCId() *string {
	return s.VPCId
}

func (s *CreateDBProxyEndpointAddressRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateDBProxyEndpointAddressRequest) SetConnectionStringPrefix(v string) *CreateDBProxyEndpointAddressRequest {
	s.ConnectionStringPrefix = &v
	return s
}

func (s *CreateDBProxyEndpointAddressRequest) SetDBInstanceId(v string) *CreateDBProxyEndpointAddressRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CreateDBProxyEndpointAddressRequest) SetDBProxyConnectStringNetType(v string) *CreateDBProxyEndpointAddressRequest {
	s.DBProxyConnectStringNetType = &v
	return s
}

func (s *CreateDBProxyEndpointAddressRequest) SetDBProxyEndpointId(v string) *CreateDBProxyEndpointAddressRequest {
	s.DBProxyEndpointId = &v
	return s
}

func (s *CreateDBProxyEndpointAddressRequest) SetDBProxyEngineType(v string) *CreateDBProxyEndpointAddressRequest {
	s.DBProxyEngineType = &v
	return s
}

func (s *CreateDBProxyEndpointAddressRequest) SetDBProxyNewConnectStringPort(v string) *CreateDBProxyEndpointAddressRequest {
	s.DBProxyNewConnectStringPort = &v
	return s
}

func (s *CreateDBProxyEndpointAddressRequest) SetRegionId(v string) *CreateDBProxyEndpointAddressRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDBProxyEndpointAddressRequest) SetResourceGroupId(v string) *CreateDBProxyEndpointAddressRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateDBProxyEndpointAddressRequest) SetVPCId(v string) *CreateDBProxyEndpointAddressRequest {
	s.VPCId = &v
	return s
}

func (s *CreateDBProxyEndpointAddressRequest) SetVSwitchId(v string) *CreateDBProxyEndpointAddressRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateDBProxyEndpointAddressRequest) Validate() error {
	return dara.Validate(s)
}
