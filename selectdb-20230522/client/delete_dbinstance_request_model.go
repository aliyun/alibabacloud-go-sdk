// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDBInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DeleteDBInstanceRequest
	GetDBInstanceId() *string
	SetRegionId(v string) *DeleteDBInstanceRequest
	GetRegionId() *string
	SetResourceOwnerId(v int64) *DeleteDBInstanceRequest
	GetResourceOwnerId() *int64
}

type DeleteDBInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// selectdb-cn-7213cjv****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteDBInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteDBInstanceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DeleteDBInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteDBInstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteDBInstanceRequest) SetDBInstanceId(v string) *DeleteDBInstanceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteDBInstanceRequest) SetRegionId(v string) *DeleteDBInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteDBInstanceRequest) SetResourceOwnerId(v int64) *DeleteDBInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteDBInstanceRequest) Validate() error {
	return dara.Validate(s)
}
