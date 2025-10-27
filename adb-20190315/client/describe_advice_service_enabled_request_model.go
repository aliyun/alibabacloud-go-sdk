// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAdviceServiceEnabledRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeAdviceServiceEnabledRequest
	GetDBClusterId() *string
	SetRegionId(v string) *DescribeAdviceServiceEnabledRequest
	GetRegionId() *string
}

type DescribeAdviceServiceEnabledRequest struct {
	// The cluster ID.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the IDs of all AnalyticDB for MySQL Data Warehouse Edition clusters within a region.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-2ze0vp0j6t3to****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The region ID of the cluster.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/143074.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeAdviceServiceEnabledRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAdviceServiceEnabledRequest) GoString() string {
	return s.String()
}

func (s *DescribeAdviceServiceEnabledRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeAdviceServiceEnabledRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAdviceServiceEnabledRequest) SetDBClusterId(v string) *DescribeAdviceServiceEnabledRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAdviceServiceEnabledRequest) SetRegionId(v string) *DescribeAdviceServiceEnabledRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAdviceServiceEnabledRequest) Validate() error {
	return dara.Validate(s)
}
