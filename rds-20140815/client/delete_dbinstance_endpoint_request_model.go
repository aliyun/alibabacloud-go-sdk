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
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	DBInstanceEndpointId *string `json:"DBInstanceEndpointId,omitempty" xml:"DBInstanceEndpointId,omitempty"`
	// This parameter is required.
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
