// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataSourceLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *DeleteDataSourceLogRequest
	GetAccountId() *string
	SetCloudCode(v string) *DeleteDataSourceLogRequest
	GetCloudCode() *string
	SetDataSourceInstanceId(v string) *DeleteDataSourceLogRequest
	GetDataSourceInstanceId() *string
	SetLogInstanceId(v string) *DeleteDataSourceLogRequest
	GetLogInstanceId() *string
	SetRegionId(v string) *DeleteDataSourceLogRequest
	GetRegionId() *string
}

type DeleteDataSourceLogRequest struct {
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
	// The ID of the log. The ID is an MD5 hash value that is calculated by the threat analysis feature based on specific parameters. You can call the [ListDataSourceLogs](https://api.aliyun-inc.com/#/publishment/document/cloud-siem/863fdf54478f4cc5877e27c2a5fe9e44?tenantUuid=f382fccd88b94c5c8c864def6815b854\\&activeTabKey=api%7CListDataSourceLogs) operation to query the IDs of logs.
	//
	// This parameter is required.
	//
	// example:
	//
	// ef33097c9d1fdb0b9c7e8c7ca320pkl1
	LogInstanceId *string `json:"LogInstanceId,omitempty" xml:"LogInstanceId,omitempty"`
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

func (s DeleteDataSourceLogRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataSourceLogRequest) GoString() string {
	return s.String()
}

func (s *DeleteDataSourceLogRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *DeleteDataSourceLogRequest) GetCloudCode() *string {
	return s.CloudCode
}

func (s *DeleteDataSourceLogRequest) GetDataSourceInstanceId() *string {
	return s.DataSourceInstanceId
}

func (s *DeleteDataSourceLogRequest) GetLogInstanceId() *string {
	return s.LogInstanceId
}

func (s *DeleteDataSourceLogRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteDataSourceLogRequest) SetAccountId(v string) *DeleteDataSourceLogRequest {
	s.AccountId = &v
	return s
}

func (s *DeleteDataSourceLogRequest) SetCloudCode(v string) *DeleteDataSourceLogRequest {
	s.CloudCode = &v
	return s
}

func (s *DeleteDataSourceLogRequest) SetDataSourceInstanceId(v string) *DeleteDataSourceLogRequest {
	s.DataSourceInstanceId = &v
	return s
}

func (s *DeleteDataSourceLogRequest) SetLogInstanceId(v string) *DeleteDataSourceLogRequest {
	s.LogInstanceId = &v
	return s
}

func (s *DeleteDataSourceLogRequest) SetRegionId(v string) *DeleteDataSourceLogRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteDataSourceLogRequest) Validate() error {
	return dara.Validate(s)
}
