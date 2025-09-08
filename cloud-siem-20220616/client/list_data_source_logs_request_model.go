// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataSourceLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *ListDataSourceLogsRequest
	GetAccountId() *string
	SetCloudCode(v string) *ListDataSourceLogsRequest
	GetCloudCode() *string
	SetDataSourceInstanceId(v string) *ListDataSourceLogsRequest
	GetDataSourceInstanceId() *string
	SetRegionId(v string) *ListDataSourceLogsRequest
	GetRegionId() *string
}

type ListDataSourceLogsRequest struct {
	// The ID of the cloud account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123xxxxxx
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The code that is used for multi-cloud environments. Valid values:
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
	// The ID of the data source. The value is obtained after the threat analysis feature calculates the MD5 hash value of a parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 220ba97c9d1fdb0b9c7e8c7ca328d7ea
	DataSourceInstanceId *string `json:"DataSourceInstanceId,omitempty" xml:"DataSourceInstanceId,omitempty"`
	// The data management center of the threat analysis feature. Specify this parameter based on the region where your assets reside. Valid values:
	//
	// 	- cn-hangzhou: Your assets reside in regions inside China.
	//
	// 	- ap-southeast-1: Your assets reside in regions outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListDataSourceLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataSourceLogsRequest) GoString() string {
	return s.String()
}

func (s *ListDataSourceLogsRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *ListDataSourceLogsRequest) GetCloudCode() *string {
	return s.CloudCode
}

func (s *ListDataSourceLogsRequest) GetDataSourceInstanceId() *string {
	return s.DataSourceInstanceId
}

func (s *ListDataSourceLogsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListDataSourceLogsRequest) SetAccountId(v string) *ListDataSourceLogsRequest {
	s.AccountId = &v
	return s
}

func (s *ListDataSourceLogsRequest) SetCloudCode(v string) *ListDataSourceLogsRequest {
	s.CloudCode = &v
	return s
}

func (s *ListDataSourceLogsRequest) SetDataSourceInstanceId(v string) *ListDataSourceLogsRequest {
	s.DataSourceInstanceId = &v
	return s
}

func (s *ListDataSourceLogsRequest) SetRegionId(v string) *ListDataSourceLogsRequest {
	s.RegionId = &v
	return s
}

func (s *ListDataSourceLogsRequest) Validate() error {
	return dara.Validate(s)
}
