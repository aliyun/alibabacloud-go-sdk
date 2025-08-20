// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSparkAppTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeSparkAppTypeRequest
	GetAppId() *string
	SetDBClusterId(v string) *DescribeSparkAppTypeRequest
	GetDBClusterId() *string
	SetRegionId(v string) *DescribeSparkAppTypeRequest
	GetRegionId() *string
}

type DescribeSparkAppTypeRequest struct {
	// The application ID.
	//
	// >  You can call the [ListSparkApps](https://help.aliyun.com/document_detail/612475.html) operation to query the Spark application IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// s202407161205sza4c07c1000****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The cluster ID.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/98094.html) operation to query the information about all AnalyticDB for MySQL clusters within a region, including cluster IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-wz9w49b12933****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The region ID.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/143074.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shenzhen
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeSparkAppTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSparkAppTypeRequest) GoString() string {
	return s.String()
}

func (s *DescribeSparkAppTypeRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeSparkAppTypeRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeSparkAppTypeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSparkAppTypeRequest) SetAppId(v string) *DescribeSparkAppTypeRequest {
	s.AppId = &v
	return s
}

func (s *DescribeSparkAppTypeRequest) SetDBClusterId(v string) *DescribeSparkAppTypeRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeSparkAppTypeRequest) SetRegionId(v string) *DescribeSparkAppTypeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSparkAppTypeRequest) Validate() error {
	return dara.Validate(s)
}
