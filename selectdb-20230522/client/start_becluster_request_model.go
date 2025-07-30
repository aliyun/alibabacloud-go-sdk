// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartBEClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *StartBEClusterRequest
	GetDBClusterId() *string
	SetDBInstanceId(v string) *StartBEClusterRequest
	GetDBInstanceId() *string
	SetRegionId(v string) *StartBEClusterRequest
	GetRegionId() *string
	SetResourceOwnerId(v int64) *StartBEClusterRequest
	GetResourceOwnerId() *int64
}

type StartBEClusterRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// selectdb-cn-7213c8yvv09-be
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
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

func (s StartBEClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s StartBEClusterRequest) GoString() string {
	return s.String()
}

func (s *StartBEClusterRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *StartBEClusterRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *StartBEClusterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StartBEClusterRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *StartBEClusterRequest) SetDBClusterId(v string) *StartBEClusterRequest {
	s.DBClusterId = &v
	return s
}

func (s *StartBEClusterRequest) SetDBInstanceId(v string) *StartBEClusterRequest {
	s.DBInstanceId = &v
	return s
}

func (s *StartBEClusterRequest) SetRegionId(v string) *StartBEClusterRequest {
	s.RegionId = &v
	return s
}

func (s *StartBEClusterRequest) SetResourceOwnerId(v int64) *StartBEClusterRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *StartBEClusterRequest) Validate() error {
	return dara.Validate(s)
}
