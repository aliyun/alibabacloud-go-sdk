// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceEndpointsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DescribeDBInstanceEndpointsRequest
	GetClientToken() *string
	SetDBInstanceEndpointId(v string) *DescribeDBInstanceEndpointsRequest
	GetDBInstanceEndpointId() *string
	SetDBInstanceId(v string) *DescribeDBInstanceEndpointsRequest
	GetDBInstanceId() *string
	SetResourceOwnerId(v int64) *DescribeDBInstanceEndpointsRequest
	GetResourceOwnerId() *int64
}

type DescribeDBInstanceEndpointsRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests.
	//
	// The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOC****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The endpoint ID of the instance.
	//
	// >  If this parameter is not specified, the information about all endpoints is returned.
	//
	// example:
	//
	// ep-****-ro
	DBInstanceEndpointId *string `json:"DBInstanceEndpointId,omitempty" xml:"DBInstanceEndpointId,omitempty"`
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-u****
	DBInstanceId    *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeDBInstanceEndpointsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceEndpointsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceEndpointsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeDBInstanceEndpointsRequest) GetDBInstanceEndpointId() *string {
	return s.DBInstanceEndpointId
}

func (s *DescribeDBInstanceEndpointsRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstanceEndpointsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDBInstanceEndpointsRequest) SetClientToken(v string) *DescribeDBInstanceEndpointsRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeDBInstanceEndpointsRequest) SetDBInstanceEndpointId(v string) *DescribeDBInstanceEndpointsRequest {
	s.DBInstanceEndpointId = &v
	return s
}

func (s *DescribeDBInstanceEndpointsRequest) SetDBInstanceId(v string) *DescribeDBInstanceEndpointsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceEndpointsRequest) SetResourceOwnerId(v int64) *DescribeDBInstanceEndpointsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBInstanceEndpointsRequest) Validate() error {
	return dara.Validate(s)
}
