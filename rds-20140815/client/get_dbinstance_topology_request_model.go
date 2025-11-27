// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDBInstanceTopologyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *GetDBInstanceTopologyRequest
	GetDBInstanceId() *string
	SetOwnerId(v int64) *GetDBInstanceTopologyRequest
	GetOwnerId() *int64
	SetResourceOwnerId(v int64) *GetDBInstanceTopologyRequest
	GetResourceOwnerId() *int64
}

type GetDBInstanceTopologyRequest struct {
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5*******
	DBInstanceId    *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetDBInstanceTopologyRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDBInstanceTopologyRequest) GoString() string {
	return s.String()
}

func (s *GetDBInstanceTopologyRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *GetDBInstanceTopologyRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetDBInstanceTopologyRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetDBInstanceTopologyRequest) SetDBInstanceId(v string) *GetDBInstanceTopologyRequest {
	s.DBInstanceId = &v
	return s
}

func (s *GetDBInstanceTopologyRequest) SetOwnerId(v int64) *GetDBInstanceTopologyRequest {
	s.OwnerId = &v
	return s
}

func (s *GetDBInstanceTopologyRequest) SetResourceOwnerId(v int64) *GetDBInstanceTopologyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetDBInstanceTopologyRequest) Validate() error {
	return dara.Validate(s)
}
