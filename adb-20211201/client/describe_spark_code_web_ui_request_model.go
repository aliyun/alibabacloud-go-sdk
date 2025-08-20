// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSparkCodeWebUiRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeSparkCodeWebUiRequest
	GetDBClusterId() *string
	SetJobId(v int64) *DescribeSparkCodeWebUiRequest
	GetJobId() *int64
	SetRegionId(v string) *DescribeSparkCodeWebUiRequest
	GetRegionId() *string
}

type DescribeSparkCodeWebUiRequest struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition cluster.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the IDs of all AnalyticDB for MySQL clusters within a region.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bp1v6usq6m65****
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

func (s DescribeSparkCodeWebUiRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSparkCodeWebUiRequest) GoString() string {
	return s.String()
}

func (s *DescribeSparkCodeWebUiRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeSparkCodeWebUiRequest) GetJobId() *int64 {
	return s.JobId
}

func (s *DescribeSparkCodeWebUiRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSparkCodeWebUiRequest) SetDBClusterId(v string) *DescribeSparkCodeWebUiRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeSparkCodeWebUiRequest) SetJobId(v int64) *DescribeSparkCodeWebUiRequest {
	s.JobId = &v
	return s
}

func (s *DescribeSparkCodeWebUiRequest) SetRegionId(v string) *DescribeSparkCodeWebUiRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSparkCodeWebUiRequest) Validate() error {
	return dara.Validate(s)
}
