// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDBInstanceEndpointAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteDBInstanceEndpointAddressRequest
	GetClientToken() *string
	SetConnectionString(v string) *DeleteDBInstanceEndpointAddressRequest
	GetConnectionString() *string
	SetDBInstanceEndpointId(v string) *DeleteDBInstanceEndpointAddressRequest
	GetDBInstanceEndpointId() *string
	SetDBInstanceId(v string) *DeleteDBInstanceEndpointAddressRequest
	GetDBInstanceId() *string
	SetResourceOwnerId(v int64) *DeleteDBInstanceEndpointAddressRequest
	GetResourceOwnerId() *int64
}

type DeleteDBInstanceEndpointAddressRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// This parameter is required.
	DBInstanceEndpointId *string `json:"DBInstanceEndpointId,omitempty" xml:"DBInstanceEndpointId,omitempty"`
	// This parameter is required.
	DBInstanceId    *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteDBInstanceEndpointAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDBInstanceEndpointAddressRequest) GoString() string {
	return s.String()
}

func (s *DeleteDBInstanceEndpointAddressRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteDBInstanceEndpointAddressRequest) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *DeleteDBInstanceEndpointAddressRequest) GetDBInstanceEndpointId() *string {
	return s.DBInstanceEndpointId
}

func (s *DeleteDBInstanceEndpointAddressRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DeleteDBInstanceEndpointAddressRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteDBInstanceEndpointAddressRequest) SetClientToken(v string) *DeleteDBInstanceEndpointAddressRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteDBInstanceEndpointAddressRequest) SetConnectionString(v string) *DeleteDBInstanceEndpointAddressRequest {
	s.ConnectionString = &v
	return s
}

func (s *DeleteDBInstanceEndpointAddressRequest) SetDBInstanceEndpointId(v string) *DeleteDBInstanceEndpointAddressRequest {
	s.DBInstanceEndpointId = &v
	return s
}

func (s *DeleteDBInstanceEndpointAddressRequest) SetDBInstanceId(v string) *DeleteDBInstanceEndpointAddressRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteDBInstanceEndpointAddressRequest) SetResourceOwnerId(v int64) *DeleteDBInstanceEndpointAddressRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteDBInstanceEndpointAddressRequest) Validate() error {
	return dara.Validate(s)
}
