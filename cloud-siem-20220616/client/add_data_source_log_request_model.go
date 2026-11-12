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
	// The ID of the cloud account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123xxxxxxx
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The code of the multi-cloud provider. Valid values:
	//
	// - qcloud: Tencent Cloud.
	//
	// - hcloud: Huawei Cloud.
	//
	// - aliyun: Alibaba Cloud. This value is only used for Alibaba Cloud data source scenarios. When you query available multi-cloud data source types, the ListDataSourceTypes operation only supports qcloud and hcloud.
	//
	// This parameter is required.
	//
	// example:
	//
	// hcloud
	CloudCode *string `json:"CloudCode,omitempty" xml:"CloudCode,omitempty"`
	// The ID of the data source instance. You can call the ListBindDataSources operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 220ba97c9d1fdb0b9c7e8c7ca328d7ea
	DataSourceInstanceId *string `json:"DataSourceInstanceId,omitempty" xml:"DataSourceInstanceId,omitempty"`
	// The details of the data source parameters. The value is a JSON array.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"LogCode":"cloud_siem_qcloud_waf_alert_log","LogParas":"[{\\"ParaCode\\":\\"api_name\\",\\"ParaValue\\":\\"GetAttackDownloadRecords\\"}]"}]
	DataSourceInstanceLogs *string `json:"DataSourceInstanceLogs,omitempty" xml:"DataSourceInstanceLogs,omitempty"`
	// The log code.
	//
	// example:
	//
	// cloud_siem_waf_xxxxx
	LogCode *string `json:"LogCode,omitempty" xml:"LogCode,omitempty"`
	// The region where the data management center of Cloud Threat Detection and Response (CTDR) is located. Select the region of the management center based on the region where your assets are located. Valid values:
	//
	// - cn-hangzhou: Your assets are located in the Chinese mainland and Hong Kong (China).
	//
	// - ap-southeast-1: Your assets are located outside China.
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
