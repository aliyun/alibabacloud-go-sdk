// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBClusterBindingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *CreateDBClusterBindingRequest
	GetDBClusterId() *string
	SetDBClusterIdBak(v string) *CreateDBClusterBindingRequest
	GetDBClusterIdBak() *string
	SetDBInstanceId(v string) *CreateDBClusterBindingRequest
	GetDBInstanceId() *string
	SetRegionId(v string) *CreateDBClusterBindingRequest
	GetRegionId() *string
}

type CreateDBClusterBindingRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// selectdb-cn-7213cjv2ez-be
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// selectdb-xxxb9f2w-be
	DBClusterIdBak *string `json:"DBClusterIdBak,omitempty" xml:"DBClusterIdBak,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// selectdb-cn-7213cjv2ez
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The region ID
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateDBClusterBindingRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDBClusterBindingRequest) GoString() string {
	return s.String()
}

func (s *CreateDBClusterBindingRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateDBClusterBindingRequest) GetDBClusterIdBak() *string {
	return s.DBClusterIdBak
}

func (s *CreateDBClusterBindingRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateDBClusterBindingRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDBClusterBindingRequest) SetDBClusterId(v string) *CreateDBClusterBindingRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateDBClusterBindingRequest) SetDBClusterIdBak(v string) *CreateDBClusterBindingRequest {
	s.DBClusterIdBak = &v
	return s
}

func (s *CreateDBClusterBindingRequest) SetDBInstanceId(v string) *CreateDBClusterBindingRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CreateDBClusterBindingRequest) SetRegionId(v string) *CreateDBClusterBindingRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDBClusterBindingRequest) Validate() error {
	return dara.Validate(s)
}
