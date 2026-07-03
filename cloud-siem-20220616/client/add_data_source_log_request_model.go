// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDataSourceLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *AddDataSourceLogRequest
	GetAccountId() *string
	SetCloudCode(v string) *AddDataSourceLogRequest
	GetCloudCode() *string
	SetDataSourceInstanceId(v string) *AddDataSourceLogRequest
	GetDataSourceInstanceId() *string
	SetDataSourceInstanceLogs(v string) *AddDataSourceLogRequest
	GetDataSourceInstanceLogs() *string
	SetLogCode(v string) *AddDataSourceLogRequest
	GetLogCode() *string
	SetRegionId(v string) *AddDataSourceLogRequest
	GetRegionId() *string
}

type AddDataSourceLogRequest struct {
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
	// The ID of the data source. Threat Analysis calculates this ID as an MD5 hash value based on specific parameters. To obtain the data source ID, call the [ListDataSourceLogs](https://api.aliyun-inc.com/#/publishment/document/cloud-siem/863fdf54478f4cc5877e27c2a5fe9e44?tenantUuid=f382fccd88b94c5c8c864def6815b854\\&activeTabKey=api%7CListDataSourceLogs) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 220ba97c9d1fdb0b9c7e8c7ca328d7ea
	DataSourceInstanceId *string `json:"DataSourceInstanceId,omitempty" xml:"DataSourceInstanceId,omitempty"`
	// The details of the data source parameters. The value must be a JSON array.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"LogCode":"cloud_siem_qcloud_waf_alert_log","LogParas":"[{\\"ParaCode\\":\\"api_name\\",\\"ParaValue\\":\\"GetAttackDownloadRecords\\"}]"}]
	DataSourceInstanceLogs *string `json:"DataSourceInstanceLogs,omitempty" xml:"DataSourceInstanceLogs,omitempty"`
	// The code of the log.
	//
	// example:
	//
	// cloud_siem_waf_xxxxx
	LogCode *string `json:"LogCode,omitempty" xml:"LogCode,omitempty"`
	// The region where the Data Management Center of Threat Analysis is deployed. Select the region where your assets are located. Valid values:
	//
	// - cn-hangzhou: Assets are in the Chinese mainland and Hong Kong (China).
	//
	// - ap-southeast-1: Assets are outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AddDataSourceLogRequest) String() string {
	return dara.Prettify(s)
}

func (s AddDataSourceLogRequest) GoString() string {
	return s.String()
}

func (s *AddDataSourceLogRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *AddDataSourceLogRequest) GetCloudCode() *string {
	return s.CloudCode
}

func (s *AddDataSourceLogRequest) GetDataSourceInstanceId() *string {
	return s.DataSourceInstanceId
}

func (s *AddDataSourceLogRequest) GetDataSourceInstanceLogs() *string {
	return s.DataSourceInstanceLogs
}

func (s *AddDataSourceLogRequest) GetLogCode() *string {
	return s.LogCode
}

func (s *AddDataSourceLogRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddDataSourceLogRequest) SetAccountId(v string) *AddDataSourceLogRequest {
	s.AccountId = &v
	return s
}

func (s *AddDataSourceLogRequest) SetCloudCode(v string) *AddDataSourceLogRequest {
	s.CloudCode = &v
	return s
}

func (s *AddDataSourceLogRequest) SetDataSourceInstanceId(v string) *AddDataSourceLogRequest {
	s.DataSourceInstanceId = &v
	return s
}

func (s *AddDataSourceLogRequest) SetDataSourceInstanceLogs(v string) *AddDataSourceLogRequest {
	s.DataSourceInstanceLogs = &v
	return s
}

func (s *AddDataSourceLogRequest) SetLogCode(v string) *AddDataSourceLogRequest {
	s.LogCode = &v
	return s
}

func (s *AddDataSourceLogRequest) SetRegionId(v string) *AddDataSourceLogRequest {
	s.RegionId = &v
	return s
}

func (s *AddDataSourceLogRequest) Validate() error {
	return dara.Validate(s)
}
