// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseInstancePublicConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConnectionString(v string) *ReleaseInstancePublicConnectionRequest
	GetConnectionString() *string
	SetDBInstanceId(v string) *ReleaseInstancePublicConnectionRequest
	GetDBInstanceId() *string
	SetRegionId(v string) *ReleaseInstancePublicConnectionRequest
	GetRegionId() *string
	SetResourceOwnerId(v int64) *ReleaseInstancePublicConnectionRequest
	GetResourceOwnerId() *int64
}

type ReleaseInstancePublicConnectionRequest struct {
	// The connection string of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// selectdb-cn-7213c8y****-public.selectdbfe.pre.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// selectdb-cn-7213cjv****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ReleaseInstancePublicConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s ReleaseInstancePublicConnectionRequest) GoString() string {
	return s.String()
}

func (s *ReleaseInstancePublicConnectionRequest) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *ReleaseInstancePublicConnectionRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ReleaseInstancePublicConnectionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ReleaseInstancePublicConnectionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ReleaseInstancePublicConnectionRequest) SetConnectionString(v string) *ReleaseInstancePublicConnectionRequest {
	s.ConnectionString = &v
	return s
}

func (s *ReleaseInstancePublicConnectionRequest) SetDBInstanceId(v string) *ReleaseInstancePublicConnectionRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ReleaseInstancePublicConnectionRequest) SetRegionId(v string) *ReleaseInstancePublicConnectionRequest {
	s.RegionId = &v
	return s
}

func (s *ReleaseInstancePublicConnectionRequest) SetResourceOwnerId(v int64) *ReleaseInstancePublicConnectionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ReleaseInstancePublicConnectionRequest) Validate() error {
	return dara.Validate(s)
}
