// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterConfigInXMLRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeDBClusterConfigInXMLRequest
	GetDBClusterId() *string
	SetRegionId(v string) *DescribeDBClusterConfigInXMLRequest
	GetRegionId() *string
}

type DescribeDBClusterConfigInXMLRequest struct {
	// The cluster ID. You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/170879.html) operation to query information about all the clusters that are deployed in a specific region. The information includes the cluster IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp108z124a8o7****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/170875.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDBClusterConfigInXMLRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterConfigInXMLRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterConfigInXMLRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeDBClusterConfigInXMLRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBClusterConfigInXMLRequest) SetDBClusterId(v string) *DescribeDBClusterConfigInXMLRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterConfigInXMLRequest) SetRegionId(v string) *DescribeDBClusterConfigInXMLRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBClusterConfigInXMLRequest) Validate() error {
	return dara.Validate(s)
}
