// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataSourceParametersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCloudCode(v string) *DescribeDataSourceParametersRequest
	GetCloudCode() *string
	SetDataSourceType(v string) *DescribeDataSourceParametersRequest
	GetDataSourceType() *string
	SetRegionId(v string) *DescribeDataSourceParametersRequest
	GetRegionId() *string
}

type DescribeDataSourceParametersRequest struct {
	// The code for the multicloud environment.
	//
	// This parameter is required.
	//
	// example:
	//
	// hcloud
	CloudCode *string `json:"CloudCode,omitempty" xml:"CloudCode,omitempty"`
	// The type of the data source. Valid values:
	//
	// - **ckafka**: Tencent Cloud CKafka.
	//
	// - **obs**: Huawei Cloud OBS.
	//
	// - **wafApi**: Tencent Cloud WAF attack log download API.
	//
	// This parameter is required.
	//
	// example:
	//
	// obs
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	// The region where the Data Management center for threat analysis is deployed. Select a region based on the region where your assets are located. Valid values:
	//
	// - cn-hangzhou: Your assets are in the Chinese mainland or China (Hong Kong).
	//
	// - ap-southeast-1: Your assets are outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDataSourceParametersRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataSourceParametersRequest) GoString() string {
	return s.String()
}

func (s *DescribeDataSourceParametersRequest) GetCloudCode() *string {
	return s.CloudCode
}

func (s *DescribeDataSourceParametersRequest) GetDataSourceType() *string {
	return s.DataSourceType
}

func (s *DescribeDataSourceParametersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDataSourceParametersRequest) SetCloudCode(v string) *DescribeDataSourceParametersRequest {
	s.CloudCode = &v
	return s
}

func (s *DescribeDataSourceParametersRequest) SetDataSourceType(v string) *DescribeDataSourceParametersRequest {
	s.DataSourceType = &v
	return s
}

func (s *DescribeDataSourceParametersRequest) SetRegionId(v string) *DescribeDataSourceParametersRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDataSourceParametersRequest) Validate() error {
	return dara.Validate(s)
}
