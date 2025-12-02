// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterSSLRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeDBClusterSSLRequest
	GetDBClusterId() *string
	SetRegionId(v string) *DescribeDBClusterSSLRequest
	GetRegionId() *string
}

type DescribeDBClusterSSLRequest struct {
	// The cluster ID.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/454250.html) operation to query the information about all AnalyticDB for MySQL clusters within a region, including cluster IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The region ID.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/612393.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDBClusterSSLRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterSSLRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterSSLRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeDBClusterSSLRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBClusterSSLRequest) SetDBClusterId(v string) *DescribeDBClusterSSLRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterSSLRequest) SetRegionId(v string) *DescribeDBClusterSSLRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBClusterSSLRequest) Validate() error {
	return dara.Validate(s)
}
