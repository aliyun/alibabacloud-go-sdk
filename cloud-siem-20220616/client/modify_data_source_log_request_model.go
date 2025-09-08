// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDataSourceLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *ModifyDataSourceLogRequest
	GetAccountId() *string
	SetCloudCode(v string) *ModifyDataSourceLogRequest
	GetCloudCode() *string
	SetDataSourceInstanceId(v string) *ModifyDataSourceLogRequest
	GetDataSourceInstanceId() *string
	SetDataSourceInstanceLogs(v string) *ModifyDataSourceLogRequest
	GetDataSourceInstanceLogs() *string
	SetDataSourceType(v string) *ModifyDataSourceLogRequest
	GetDataSourceType() *string
	SetLogCode(v string) *ModifyDataSourceLogRequest
	GetLogCode() *string
	SetLogInstanceId(v string) *ModifyDataSourceLogRequest
	GetLogInstanceId() *string
	SetRegionId(v string) *ModifyDataSourceLogRequest
	GetRegionId() *string
}

type ModifyDataSourceLogRequest struct {
	// The ID of the cloud account.
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
	// The ID of the data source. The ID is an MD5 hash value that is calculated by the threat analysis feature based on specific parameters. You can call the [DescribeDataSourceInstance](https://api.aliyun-inc.com/#/publishment/document/cloud-siem/863fdf54478f4cc5877e27c2a5fe9e44?tenantUuid=f382fccd88b94c5c8c864def6815b854\\&activeTabKey=api%7CDescribeDataSourceInstance) operation to query the IDs of data sources.
	//
	// This parameter is required.
	//
	// example:
	//
	// ef33097c9d1fdb0b9c7e8c7ca320pkl1
	DataSourceInstanceId *string `json:"DataSourceInstanceId,omitempty" xml:"DataSourceInstanceId,omitempty"`
	// The parameters of the data source. Set this parameter to a JSON string.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"LogCode":"cloud_siem_qcloud_waf_alert_log","LogParas":"[{\\"ParaCode\\":\\"api_name\\",\\"ParaValue\\":\\"GetAttackDownloadRecords\\"}]"}]
	DataSourceInstanceLogs *string `json:"DataSourceInstanceLogs,omitempty" xml:"DataSourceInstanceLogs,omitempty"`
	// The type of the data source. Valid values:
	//
	// 	- obs: Huawei Cloud Object Storage Service (OBS)
	//
	// 	- wafApi: download API of Tencent Cloud Web Application Firewall (WAF)
	//
	// 	- ckafka: Tencent Cloud Kafka (CKafka)
	//
	// example:
	//
	// obs
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	// The log code.
	//
	// example:
	//
	// cloud_siem_waf_xxxxx
	LogCode *string `json:"LogCode,omitempty" xml:"LogCode,omitempty"`
	// The ID of the log. The ID is an MD5 hash value that is calculated by the threat analysis feature based on specific parameters. You can call the [ListDataSourceLogs](https://api.aliyun-inc.com/#/publishment/document/cloud-siem/863fdf54478f4cc5877e27c2a5fe9e44?tenantUuid=f382fccd88b94c5c8c864def6815b854\\&activeTabKey=api%7CListDataSourceLogs) to query log IDs.
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

func (s ModifyDataSourceLogRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDataSourceLogRequest) GoString() string {
	return s.String()
}

func (s *ModifyDataSourceLogRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *ModifyDataSourceLogRequest) GetCloudCode() *string {
	return s.CloudCode
}

func (s *ModifyDataSourceLogRequest) GetDataSourceInstanceId() *string {
	return s.DataSourceInstanceId
}

func (s *ModifyDataSourceLogRequest) GetDataSourceInstanceLogs() *string {
	return s.DataSourceInstanceLogs
}

func (s *ModifyDataSourceLogRequest) GetDataSourceType() *string {
	return s.DataSourceType
}

func (s *ModifyDataSourceLogRequest) GetLogCode() *string {
	return s.LogCode
}

func (s *ModifyDataSourceLogRequest) GetLogInstanceId() *string {
	return s.LogInstanceId
}

func (s *ModifyDataSourceLogRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDataSourceLogRequest) SetAccountId(v string) *ModifyDataSourceLogRequest {
	s.AccountId = &v
	return s
}

func (s *ModifyDataSourceLogRequest) SetCloudCode(v string) *ModifyDataSourceLogRequest {
	s.CloudCode = &v
	return s
}

func (s *ModifyDataSourceLogRequest) SetDataSourceInstanceId(v string) *ModifyDataSourceLogRequest {
	s.DataSourceInstanceId = &v
	return s
}

func (s *ModifyDataSourceLogRequest) SetDataSourceInstanceLogs(v string) *ModifyDataSourceLogRequest {
	s.DataSourceInstanceLogs = &v
	return s
}

func (s *ModifyDataSourceLogRequest) SetDataSourceType(v string) *ModifyDataSourceLogRequest {
	s.DataSourceType = &v
	return s
}

func (s *ModifyDataSourceLogRequest) SetLogCode(v string) *ModifyDataSourceLogRequest {
	s.LogCode = &v
	return s
}

func (s *ModifyDataSourceLogRequest) SetLogInstanceId(v string) *ModifyDataSourceLogRequest {
	s.LogInstanceId = &v
	return s
}

func (s *ModifyDataSourceLogRequest) SetRegionId(v string) *ModifyDataSourceLogRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDataSourceLogRequest) Validate() error {
	return dara.Validate(s)
}
