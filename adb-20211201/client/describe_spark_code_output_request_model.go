// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSparkCodeOutputRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeSparkCodeOutputRequest
	GetDBClusterId() *string
	SetJobId(v int64) *DescribeSparkCodeOutputRequest
	GetJobId() *int64
	SetRegionId(v string) *DescribeSparkCodeOutputRequest
	GetRegionId() *string
}

type DescribeSparkCodeOutputRequest struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition cluster.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/612397.html) operation to query the IDs of all AnalyticDB for MySQL clusters within a region.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-uf6210mmev07****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The ID of the Spark job.
	//
	// This parameter is required.
	//
	// example:
	//
	// 620
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

func (s DescribeSparkCodeOutputRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSparkCodeOutputRequest) GoString() string {
	return s.String()
}

func (s *DescribeSparkCodeOutputRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeSparkCodeOutputRequest) GetJobId() *int64 {
	return s.JobId
}

func (s *DescribeSparkCodeOutputRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSparkCodeOutputRequest) SetDBClusterId(v string) *DescribeSparkCodeOutputRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeSparkCodeOutputRequest) SetJobId(v int64) *DescribeSparkCodeOutputRequest {
	s.JobId = &v
	return s
}

func (s *DescribeSparkCodeOutputRequest) SetRegionId(v string) *DescribeSparkCodeOutputRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSparkCodeOutputRequest) Validate() error {
	return dara.Validate(s)
}
