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
	// The code of the cloud service provider.
	//
	// Valid values:
	//
	// 	- qcloud
	//
	// 	- hcloud
	//
	// 	- aliyun
	//
	// This parameter is required.
	//
	// example:
	//
	// hcloud
	CloudCode *string `json:"CloudCode,omitempty" xml:"CloudCode,omitempty"`
	// The type of the data source. Valid values:
	//
	// 	- **ckafka**: Tencent Cloud TDMQ for CKafka
	//
	// 	- **obs**: Huawei Cloud Object Storage Service (OBS)
	//
	// 	- **wafApi**: download API of Tencent Cloud Web Application Firewall (WAF)
	//
	// This parameter is required.
	//
	// example:
	//
	// obs
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	// The region in which the data management center of the threat analysis feature resides. Specify this parameter based on the regions in which your assets reside. Valid values:
	//
	// 	- cn-hangzhou: Your assets reside in regions in China.
	//
	// 	- ap-southeast-1: Your assets reside in regions outside China.
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
