// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBProxyEndpointAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ModifyDBProxyEndpointAddressRequest
	GetDBInstanceId() *string
	SetDBProxyConnectStringNetType(v string) *ModifyDBProxyEndpointAddressRequest
	GetDBProxyConnectStringNetType() *string
	SetDBProxyEndpointId(v string) *ModifyDBProxyEndpointAddressRequest
	GetDBProxyEndpointId() *string
	SetDBProxyEngineType(v string) *ModifyDBProxyEndpointAddressRequest
	GetDBProxyEngineType() *string
	SetDBProxyNewConnectString(v string) *ModifyDBProxyEndpointAddressRequest
	GetDBProxyNewConnectString() *string
	SetDBProxyNewConnectStringPort(v string) *ModifyDBProxyEndpointAddressRequest
	GetDBProxyNewConnectStringPort() *string
	SetOwnerId(v int64) *ModifyDBProxyEndpointAddressRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyDBProxyEndpointAddressRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBProxyEndpointAddressRequest
	GetResourceOwnerId() *int64
}

type ModifyDBProxyEndpointAddressRequest struct {
	// This parameter is required.
	DBInstanceId                *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DBProxyConnectStringNetType *string `json:"DBProxyConnectStringNetType,omitempty" xml:"DBProxyConnectStringNetType,omitempty"`
	// This parameter is required.
	DBProxyEndpointId           *string `json:"DBProxyEndpointId,omitempty" xml:"DBProxyEndpointId,omitempty"`
	DBProxyEngineType           *string `json:"DBProxyEngineType,omitempty" xml:"DBProxyEngineType,omitempty"`
	DBProxyNewConnectString     *string `json:"DBProxyNewConnectString,omitempty" xml:"DBProxyNewConnectString,omitempty"`
	DBProxyNewConnectStringPort *string `json:"DBProxyNewConnectStringPort,omitempty" xml:"DBProxyNewConnectStringPort,omitempty"`
	OwnerId                     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount        *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId             *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyDBProxyEndpointAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBProxyEndpointAddressRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBProxyEndpointAddressRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDBProxyEndpointAddressRequest) GetDBProxyConnectStringNetType() *string {
	return s.DBProxyConnectStringNetType
}

func (s *ModifyDBProxyEndpointAddressRequest) GetDBProxyEndpointId() *string {
	return s.DBProxyEndpointId
}

func (s *ModifyDBProxyEndpointAddressRequest) GetDBProxyEngineType() *string {
	return s.DBProxyEngineType
}

func (s *ModifyDBProxyEndpointAddressRequest) GetDBProxyNewConnectString() *string {
	return s.DBProxyNewConnectString
}

func (s *ModifyDBProxyEndpointAddressRequest) GetDBProxyNewConnectStringPort() *string {
	return s.DBProxyNewConnectStringPort
}

func (s *ModifyDBProxyEndpointAddressRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBProxyEndpointAddressRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBProxyEndpointAddressRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBProxyEndpointAddressRequest) SetDBInstanceId(v string) *ModifyDBProxyEndpointAddressRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBProxyEndpointAddressRequest) SetDBProxyConnectStringNetType(v string) *ModifyDBProxyEndpointAddressRequest {
	s.DBProxyConnectStringNetType = &v
	return s
}

func (s *ModifyDBProxyEndpointAddressRequest) SetDBProxyEndpointId(v string) *ModifyDBProxyEndpointAddressRequest {
	s.DBProxyEndpointId = &v
	return s
}

func (s *ModifyDBProxyEndpointAddressRequest) SetDBProxyEngineType(v string) *ModifyDBProxyEndpointAddressRequest {
	s.DBProxyEngineType = &v
	return s
}

func (s *ModifyDBProxyEndpointAddressRequest) SetDBProxyNewConnectString(v string) *ModifyDBProxyEndpointAddressRequest {
	s.DBProxyNewConnectString = &v
	return s
}

func (s *ModifyDBProxyEndpointAddressRequest) SetDBProxyNewConnectStringPort(v string) *ModifyDBProxyEndpointAddressRequest {
	s.DBProxyNewConnectStringPort = &v
	return s
}

func (s *ModifyDBProxyEndpointAddressRequest) SetOwnerId(v int64) *ModifyDBProxyEndpointAddressRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBProxyEndpointAddressRequest) SetResourceOwnerAccount(v string) *ModifyDBProxyEndpointAddressRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBProxyEndpointAddressRequest) SetResourceOwnerId(v int64) *ModifyDBProxyEndpointAddressRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBProxyEndpointAddressRequest) Validate() error {
	return dara.Validate(s)
}
