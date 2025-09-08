// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *DeleteDataSourceRequest
	GetAccountId() *string
	SetCloudCode(v string) *DeleteDataSourceRequest
	GetCloudCode() *string
	SetDataSourceInstanceId(v string) *DeleteDataSourceRequest
	GetDataSourceInstanceId() *string
	SetRegionId(v string) *DeleteDataSourceRequest
	GetRegionId() *string
}

type DeleteDataSourceRequest struct {
	// The ID of the cloud account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123xxxxxxx
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The code of the cloud service provider. Valid values:
	//
	// 	- qcloud: Tencent Cloud
	//
	// 	- aliyun: Alibaba Cloud
	//
	// 	- hcloud: Huawei Cloud
	//
	// This parameter is required.
	//
	// example:
	//
	// hcloud
	CloudCode *string `json:"CloudCode,omitempty" xml:"CloudCode,omitempty"`
	// The ID of the data source. The ID is an MD5 hash value that is calculated by the threat analysis feature based on specific parameters. You can call the [ListDataSourceLogs](https://api.aliyun-inc.com/#/publishment/document/cloud-siem/863fdf54478f4cc5877e27c2a5fe9e44?tenantUuid=f382fccd88b94c5c8c864def6815b854\\&activeTabKey=api%7CListDataSourceLogs) operation to query the IDs of data sources.
	//
	// This parameter is required.
	//
	// example:
	//
	// 220ba97c9d1fdb0b9c7e8c7ca328d7ea
	DataSourceInstanceId *string `json:"DataSourceInstanceId,omitempty" xml:"DataSourceInstanceId,omitempty"`
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

func (s DeleteDataSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataSourceRequest) GoString() string {
	return s.String()
}

func (s *DeleteDataSourceRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *DeleteDataSourceRequest) GetCloudCode() *string {
	return s.CloudCode
}

func (s *DeleteDataSourceRequest) GetDataSourceInstanceId() *string {
	return s.DataSourceInstanceId
}

func (s *DeleteDataSourceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteDataSourceRequest) SetAccountId(v string) *DeleteDataSourceRequest {
	s.AccountId = &v
	return s
}

func (s *DeleteDataSourceRequest) SetCloudCode(v string) *DeleteDataSourceRequest {
	s.CloudCode = &v
	return s
}

func (s *DeleteDataSourceRequest) SetDataSourceInstanceId(v string) *DeleteDataSourceRequest {
	s.DataSourceInstanceId = &v
	return s
}

func (s *DeleteDataSourceRequest) SetRegionId(v string) *DeleteDataSourceRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteDataSourceRequest) Validate() error {
	return dara.Validate(s)
}
