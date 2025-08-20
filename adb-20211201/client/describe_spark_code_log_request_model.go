// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSparkCodeLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeSparkCodeLogRequest
	GetDBClusterId() *string
	SetJobId(v int64) *DescribeSparkCodeLogRequest
	GetJobId() *int64
	SetRegionId(v string) *DescribeSparkCodeLogRequest
	GetRegionId() *string
}

type DescribeSparkCodeLogRequest struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition cluster.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/454250.html) operation to query the IDs of all AnalyticDB for MySQL clusters within a region.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-uf6o6m8p6x7v****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The ID of the Spark job.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1248
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The region ID of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeSparkCodeLogRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSparkCodeLogRequest) GoString() string {
	return s.String()
}

func (s *DescribeSparkCodeLogRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeSparkCodeLogRequest) GetJobId() *int64 {
	return s.JobId
}

func (s *DescribeSparkCodeLogRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSparkCodeLogRequest) SetDBClusterId(v string) *DescribeSparkCodeLogRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeSparkCodeLogRequest) SetJobId(v int64) *DescribeSparkCodeLogRequest {
	s.JobId = &v
	return s
}

func (s *DescribeSparkCodeLogRequest) SetRegionId(v string) *DescribeSparkCodeLogRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSparkCodeLogRequest) Validate() error {
	return dara.Validate(s)
}
