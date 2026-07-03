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
	// The ID of the Alibaba Cloud account.
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
	// The ID of the data source. The threat analysis feature generates this ID by calculating an MD5 hash of the parameters.
	//
	// Call the [DescribeDataSourceInstance](https://help.aliyun.com/document_detail/2639736.html) operation to obtain the data source ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ef33097c9d1fdb0b9c7e8c7ca320pkl1
	DataSourceInstanceId *string `json:"DataSourceInstanceId,omitempty" xml:"DataSourceInstanceId,omitempty"`
	// The details of the data source parameters, in a JSON array format.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"LogCode":"cloud_siem_qcloud_waf_alert_log","LogParas":"[{\\"ParaCode\\":\\"api_name\\",\\"ParaValue\\":\\"GetAttackDownloadRecords\\"}]"}]
	DataSourceInstanceLogs *string `json:"DataSourceInstanceLogs,omitempty" xml:"DataSourceInstanceLogs,omitempty"`
	// The type of the data source. Valid values:
	//
	// - obs: Huawei Cloud Object Storage Service (OBS).
	//
	// - wafApi: Tencent Cloud Web Application Firewall (WAF) download API.
	//
	// - ckafka: Tencent Cloud CKafka.
	//
	// example:
	//
	// obs
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	// The code of the log.
	//
	// example:
	//
	// cloud_siem_waf_xxxxx
	LogCode *string `json:"LogCode,omitempty" xml:"LogCode,omitempty"`
	// The ID of the log. The threat analysis feature generates this ID by calculating an MD5 hash of the parameters. Call the [ListDataSourceLogs](https://help.aliyun.com/document_detail/2639707.html) operation to obtain the log ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ef33097c9d1fdb0b9c7e8c7ca320pkl1
	LogInstanceId *string `json:"LogInstanceId,omitempty" xml:"LogInstanceId,omitempty"`
	// The region where the Data Management hub is located. Select a region based on the location of your assets. Valid values:
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
