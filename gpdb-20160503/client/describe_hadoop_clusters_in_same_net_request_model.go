// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHadoopClustersInSameNetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeHadoopClustersInSameNetRequest
	GetDBInstanceId() *string
	SetRegionId(v string) *DescribeHadoopClustersInSameNetRequest
	GetRegionId() *string
}

type DescribeHadoopClustersInSameNetRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeHadoopClustersInSameNetRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeHadoopClustersInSameNetRequest) GoString() string {
	return s.String()
}

func (s *DescribeHadoopClustersInSameNetRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeHadoopClustersInSameNetRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeHadoopClustersInSameNetRequest) SetDBInstanceId(v string) *DescribeHadoopClustersInSameNetRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeHadoopClustersInSameNetRequest) SetRegionId(v string) *DescribeHadoopClustersInSameNetRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeHadoopClustersInSameNetRequest) Validate() error {
	return dara.Validate(s)
}
