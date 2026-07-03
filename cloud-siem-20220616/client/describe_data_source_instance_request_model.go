// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataSourceInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *DescribeDataSourceInstanceRequest
	GetAccountId() *string
	SetCloudCode(v string) *DescribeDataSourceInstanceRequest
	GetCloudCode() *string
	SetDataSourceInstanceId(v string) *DescribeDataSourceInstanceRequest
	GetDataSourceInstanceId() *string
	SetRegionId(v string) *DescribeDataSourceInstanceRequest
	GetRegionId() *string
}

type DescribeDataSourceInstanceRequest struct {
	// The ID of the Alibaba Cloud account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123xxxxxxx
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The code for the multicloud environment. Valid values:
	//
	// - qcloud: Tencent Cloud.
	//
	// - aliyun: Alibaba Cloud.
	//
	// - hcloud: Huawei Cloud.
	//
	// This parameter is required.
	//
	// example:
	//
	// hcloud
	CloudCode *string `json:"CloudCode,omitempty" xml:"CloudCode,omitempty"`
	// The ID of the data source. The threat analysis service generates the ID by calculating the MD5 hash of specific parameters. Call the [ListDataSourceLogs](https://api.aliyun-inc.com/#/publishment/document/cloud-siem/863fdf54478f4cc5877e27c2a5fe9e44?tenantUuid=f382fccd88b94c5c8c864def6815b854\\&activeTabKey=api%7CListDataSourceLogs) operation to obtain the data source ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 220ba97c9d1fdb0b9c7e8c7ca328d7ea
	DataSourceInstanceId *string `json:"DataSourceInstanceId,omitempty" xml:"DataSourceInstanceId,omitempty"`
	// The region where the Data Management center of the threat analysis service is located. Select a region based on the region where your assets are located. Valid values:
	//
	// - cn-hangzhou: Your assets are in the Chinese mainland or China (Hong Kong).
	//
	// - ap-southeast-1: Your assets are in a region outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDataSourceInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataSourceInstanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeDataSourceInstanceRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *DescribeDataSourceInstanceRequest) GetCloudCode() *string {
	return s.CloudCode
}

func (s *DescribeDataSourceInstanceRequest) GetDataSourceInstanceId() *string {
	return s.DataSourceInstanceId
}

func (s *DescribeDataSourceInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDataSourceInstanceRequest) SetAccountId(v string) *DescribeDataSourceInstanceRequest {
	s.AccountId = &v
	return s
}

func (s *DescribeDataSourceInstanceRequest) SetCloudCode(v string) *DescribeDataSourceInstanceRequest {
	s.CloudCode = &v
	return s
}

func (s *DescribeDataSourceInstanceRequest) SetDataSourceInstanceId(v string) *DescribeDataSourceInstanceRequest {
	s.DataSourceInstanceId = &v
	return s
}

func (s *DescribeDataSourceInstanceRequest) SetRegionId(v string) *DescribeDataSourceInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDataSourceInstanceRequest) Validate() error {
	return dara.Validate(s)
}
