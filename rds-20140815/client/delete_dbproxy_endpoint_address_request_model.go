// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDBProxyEndpointAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DeleteDBProxyEndpointAddressRequest
	GetDBInstanceId() *string
	SetDBProxyConnectStringNetType(v string) *DeleteDBProxyEndpointAddressRequest
	GetDBProxyConnectStringNetType() *string
	SetDBProxyEndpointId(v string) *DeleteDBProxyEndpointAddressRequest
	GetDBProxyEndpointId() *string
	SetDBProxyEngineType(v string) *DeleteDBProxyEndpointAddressRequest
	GetDBProxyEngineType() *string
	SetRegionId(v string) *DeleteDBProxyEndpointAddressRequest
	GetRegionId() *string
}

type DeleteDBProxyEndpointAddressRequest struct {
	// This parameter is required.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// This parameter is required.
	DBProxyConnectStringNetType *string `json:"DBProxyConnectStringNetType,omitempty" xml:"DBProxyConnectStringNetType,omitempty"`
	// This parameter is required.
	DBProxyEndpointId *string `json:"DBProxyEndpointId,omitempty" xml:"DBProxyEndpointId,omitempty"`
	DBProxyEngineType *string `json:"DBProxyEngineType,omitempty" xml:"DBProxyEngineType,omitempty"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteDBProxyEndpointAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDBProxyEndpointAddressRequest) GoString() string {
	return s.String()
}

func (s *DeleteDBProxyEndpointAddressRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DeleteDBProxyEndpointAddressRequest) GetDBProxyConnectStringNetType() *string {
	return s.DBProxyConnectStringNetType
}

func (s *DeleteDBProxyEndpointAddressRequest) GetDBProxyEndpointId() *string {
	return s.DBProxyEndpointId
}

func (s *DeleteDBProxyEndpointAddressRequest) GetDBProxyEngineType() *string {
	return s.DBProxyEngineType
}

func (s *DeleteDBProxyEndpointAddressRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteDBProxyEndpointAddressRequest) SetDBInstanceId(v string) *DeleteDBProxyEndpointAddressRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteDBProxyEndpointAddressRequest) SetDBProxyConnectStringNetType(v string) *DeleteDBProxyEndpointAddressRequest {
	s.DBProxyConnectStringNetType = &v
	return s
}

func (s *DeleteDBProxyEndpointAddressRequest) SetDBProxyEndpointId(v string) *DeleteDBProxyEndpointAddressRequest {
	s.DBProxyEndpointId = &v
	return s
}

func (s *DeleteDBProxyEndpointAddressRequest) SetDBProxyEngineType(v string) *DeleteDBProxyEndpointAddressRequest {
	s.DBProxyEngineType = &v
	return s
}

func (s *DeleteDBProxyEndpointAddressRequest) SetRegionId(v string) *DeleteDBProxyEndpointAddressRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteDBProxyEndpointAddressRequest) Validate() error {
	return dara.Validate(s)
}
