// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDBInstanceEndpointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteDBInstanceEndpointRequest
	GetClientToken() *string
	SetDBInstanceEndpointId(v string) *DeleteDBInstanceEndpointRequest
	GetDBInstanceEndpointId() *string
	SetDBInstanceId(v string) *DeleteDBInstanceEndpointRequest
	GetDBInstanceId() *string
	SetResourceOwnerId(v int64) *DeleteDBInstanceEndpointRequest
	GetResourceOwnerId() *int64
}

type DeleteDBInstanceEndpointRequest struct {
	// The client token that is used to ensure the idempotency of requests. You can use the client to generate the value, but you must ensure that it is unique among different requests. The token can only contain ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
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

func (s DeleteDBInstanceEndpointRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDBInstanceEndpointRequest) GoString() string {
	return s.String()
}

func (s *DeleteDBInstanceEndpointRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteDBInstanceEndpointRequest) GetDBInstanceEndpointId() *string {
	return s.DBInstanceEndpointId
}

func (s *DeleteDBInstanceEndpointRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DeleteDBInstanceEndpointRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteDBInstanceEndpointRequest) SetClientToken(v string) *DeleteDBInstanceEndpointRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteDBInstanceEndpointRequest) SetDBInstanceEndpointId(v string) *DeleteDBInstanceEndpointRequest {
	s.DBInstanceEndpointId = &v
	return s
}

func (s *DeleteDBInstanceEndpointRequest) SetDBInstanceId(v string) *DeleteDBInstanceEndpointRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteDBInstanceEndpointRequest) SetResourceOwnerId(v int64) *DeleteDBInstanceEndpointRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteDBInstanceEndpointRequest) Validate() error {
	return dara.Validate(s)
}
