// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBInstanceEndpointAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateDBInstanceEndpointAddressRequest
	GetClientToken() *string
	SetConnectionStringPrefix(v string) *CreateDBInstanceEndpointAddressRequest
	GetConnectionStringPrefix() *string
	SetDBInstanceEndpointId(v string) *CreateDBInstanceEndpointAddressRequest
	GetDBInstanceEndpointId() *string
	SetDBInstanceId(v string) *CreateDBInstanceEndpointAddressRequest
	GetDBInstanceId() *string
	SetIpType(v string) *CreateDBInstanceEndpointAddressRequest
	GetIpType() *string
	SetPort(v string) *CreateDBInstanceEndpointAddressRequest
	GetPort() *string
	SetResourceGroupId(v string) *CreateDBInstanceEndpointAddressRequest
	GetResourceGroupId() *string
	SetResourceOwnerId(v int64) *CreateDBInstanceEndpointAddressRequest
	GetResourceOwnerId() *int64
}

type CreateDBInstanceEndpointAddressRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	ConnectionStringPrefix *string `json:"ConnectionStringPrefix,omitempty" xml:"ConnectionStringPrefix,omitempty"`
	// This parameter is required.
	DBInstanceEndpointId *string `json:"DBInstanceEndpointId,omitempty" xml:"DBInstanceEndpointId,omitempty"`
	// This parameter is required.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// This parameter is required.
	IpType *string `json:"IpType,omitempty" xml:"IpType,omitempty"`
	// This parameter is required.
	Port            *string `json:"Port,omitempty" xml:"Port,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateDBInstanceEndpointAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDBInstanceEndpointAddressRequest) GoString() string {
	return s.String()
}

func (s *CreateDBInstanceEndpointAddressRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateDBInstanceEndpointAddressRequest) GetConnectionStringPrefix() *string {
	return s.ConnectionStringPrefix
}

func (s *CreateDBInstanceEndpointAddressRequest) GetDBInstanceEndpointId() *string {
	return s.DBInstanceEndpointId
}

func (s *CreateDBInstanceEndpointAddressRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateDBInstanceEndpointAddressRequest) GetIpType() *string {
	return s.IpType
}

func (s *CreateDBInstanceEndpointAddressRequest) GetPort() *string {
	return s.Port
}

func (s *CreateDBInstanceEndpointAddressRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateDBInstanceEndpointAddressRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateDBInstanceEndpointAddressRequest) SetClientToken(v string) *CreateDBInstanceEndpointAddressRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDBInstanceEndpointAddressRequest) SetConnectionStringPrefix(v string) *CreateDBInstanceEndpointAddressRequest {
	s.ConnectionStringPrefix = &v
	return s
}

func (s *CreateDBInstanceEndpointAddressRequest) SetDBInstanceEndpointId(v string) *CreateDBInstanceEndpointAddressRequest {
	s.DBInstanceEndpointId = &v
	return s
}

func (s *CreateDBInstanceEndpointAddressRequest) SetDBInstanceId(v string) *CreateDBInstanceEndpointAddressRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CreateDBInstanceEndpointAddressRequest) SetIpType(v string) *CreateDBInstanceEndpointAddressRequest {
	s.IpType = &v
	return s
}

func (s *CreateDBInstanceEndpointAddressRequest) SetPort(v string) *CreateDBInstanceEndpointAddressRequest {
	s.Port = &v
	return s
}

func (s *CreateDBInstanceEndpointAddressRequest) SetResourceGroupId(v string) *CreateDBInstanceEndpointAddressRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateDBInstanceEndpointAddressRequest) SetResourceOwnerId(v int64) *CreateDBInstanceEndpointAddressRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateDBInstanceEndpointAddressRequest) Validate() error {
	return dara.Validate(s)
}
