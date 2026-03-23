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
	DBProxyEngineType *string `json:"DBProxyEngineType,omitempty" xml:"DBProxyEngineType,omitempty"`
	// This parameter is required.
	DbInstanceId *string `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
	// This parameter is required.
	DbProxyConnectString *string `json:"DbProxyConnectString,omitempty" xml:"DbProxyConnectString,omitempty"`
	// This parameter is required.
	DbProxyEndpointId *string `json:"DbProxyEndpointId,omitempty" xml:"DbProxyEndpointId,omitempty"`
	// This parameter is required.
	DbProxySslEnabled *string `json:"DbProxySslEnabled,omitempty" xml:"DbProxySslEnabled,omitempty"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
