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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests.
	//
	// The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 6000170000591aed949d0f****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The public endpoint.
	//
	// This parameter is required.
	//
	// example:
	//
	// new****.mysql.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The endpoint ID of the instance. You can call the DescribeDBInstanceEndpoints operation to query the endpoint ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ep-****
	DBInstanceEndpointId *string `json:"DBInstanceEndpointId,omitempty" xml:"DBInstanceEndpointId,omitempty"`
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-****
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
