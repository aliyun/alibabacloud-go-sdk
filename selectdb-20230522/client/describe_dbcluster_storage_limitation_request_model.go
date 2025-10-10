// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterStorageLimitationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeDBClusterStorageLimitationRequest
	GetDBClusterId() *string
	SetDBInstanceId(v string) *DescribeDBClusterStorageLimitationRequest
	GetDBInstanceId() *string
	SetRegionId(v string) *DescribeDBClusterStorageLimitationRequest
	GetRegionId() *string
}

type DescribeDBClusterStorageLimitationRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// selectdb-cn-7213c8****-be
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// selectdb-cn-28t3qs****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDBClusterStorageLimitationRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterStorageLimitationRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterStorageLimitationRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeDBClusterStorageLimitationRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBClusterStorageLimitationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBClusterStorageLimitationRequest) SetDBClusterId(v string) *DescribeDBClusterStorageLimitationRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterStorageLimitationRequest) SetDBInstanceId(v string) *DescribeDBClusterStorageLimitationRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBClusterStorageLimitationRequest) SetRegionId(v string) *DescribeDBClusterStorageLimitationRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBClusterStorageLimitationRequest) Validate() error {
	return dara.Validate(s)
}
