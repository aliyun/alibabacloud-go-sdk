// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDataSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *ModifyDataSourceRequest
	GetAccountId() *string
	SetCloudCode(v string) *ModifyDataSourceRequest
	GetCloudCode() *string
	SetDataSourceInstanceId(v string) *ModifyDataSourceRequest
	GetDataSourceInstanceId() *string
	SetDataSourceInstanceName(v string) *ModifyDataSourceRequest
	GetDataSourceInstanceName() *string
	SetDataSourceInstanceParams(v string) *ModifyDataSourceRequest
	GetDataSourceInstanceParams() *string
	SetDataSourceInstanceRemark(v string) *ModifyDataSourceRequest
	GetDataSourceInstanceRemark() *string
	SetDataSourceType(v string) *ModifyDataSourceRequest
	GetDataSourceType() *string
	SetRegionId(v string) *ModifyDataSourceRequest
	GetRegionId() *string
}

type ModifyDataSourceRequest struct {
	// The ID of the cloud account.
	//
	// example:
	//
	// 123xxxxxx
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
	// 220ba97c9d1fdb0b9c7e8c7ca328d7ea
	DataSourceInstanceId *string `json:"DataSourceInstanceId,omitempty" xml:"DataSourceInstanceId,omitempty"`
	// The name of the data source.
	//
	// example:
	//
	// beijing_waf_kafka
	DataSourceInstanceName *string `json:"DataSourceInstanceName,omitempty" xml:"DataSourceInstanceName,omitempty"`
	// The parameters of the data source in the JSON string format.
	//
	// example:
	//
	// [{"paraCode":"region_code","paraValue":"ap-guangzhou"}]
	DataSourceInstanceParams *string `json:"DataSourceInstanceParams,omitempty" xml:"DataSourceInstanceParams,omitempty"`
	// The remarks on the data source.
	//
	// example:
	//
	// waf_alert_log
	DataSourceInstanceRemark *string `json:"DataSourceInstanceRemark,omitempty" xml:"DataSourceInstanceRemark,omitempty"`
	// The type of the data source. Valid values:
	//
	// 	- ckafka: Tencent Cloud Kafka (CKafka)
	//
	// 	- obs: Huawei Cloud Object Storage Service (OBS)
	//
	// 	- wafApi: download API of Tencent Cloud Web Application Firewall (WAF)
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

func (s ModifyDataSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDataSourceRequest) GoString() string {
	return s.String()
}

func (s *ModifyDataSourceRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *ModifyDataSourceRequest) GetCloudCode() *string {
	return s.CloudCode
}

func (s *ModifyDataSourceRequest) GetDataSourceInstanceId() *string {
	return s.DataSourceInstanceId
}

func (s *ModifyDataSourceRequest) GetDataSourceInstanceName() *string {
	return s.DataSourceInstanceName
}

func (s *ModifyDataSourceRequest) GetDataSourceInstanceParams() *string {
	return s.DataSourceInstanceParams
}

func (s *ModifyDataSourceRequest) GetDataSourceInstanceRemark() *string {
	return s.DataSourceInstanceRemark
}

func (s *ModifyDataSourceRequest) GetDataSourceType() *string {
	return s.DataSourceType
}

func (s *ModifyDataSourceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDataSourceRequest) SetAccountId(v string) *ModifyDataSourceRequest {
	s.AccountId = &v
	return s
}

func (s *ModifyDataSourceRequest) SetCloudCode(v string) *ModifyDataSourceRequest {
	s.CloudCode = &v
	return s
}

func (s *ModifyDataSourceRequest) SetDataSourceInstanceId(v string) *ModifyDataSourceRequest {
	s.DataSourceInstanceId = &v
	return s
}

func (s *ModifyDataSourceRequest) SetDataSourceInstanceName(v string) *ModifyDataSourceRequest {
	s.DataSourceInstanceName = &v
	return s
}

func (s *ModifyDataSourceRequest) SetDataSourceInstanceParams(v string) *ModifyDataSourceRequest {
	s.DataSourceInstanceParams = &v
	return s
}

func (s *ModifyDataSourceRequest) SetDataSourceInstanceRemark(v string) *ModifyDataSourceRequest {
	s.DataSourceInstanceRemark = &v
	return s
}

func (s *ModifyDataSourceRequest) SetDataSourceType(v string) *ModifyDataSourceRequest {
	s.DataSourceType = &v
	return s
}

func (s *ModifyDataSourceRequest) SetRegionId(v string) *ModifyDataSourceRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDataSourceRequest) Validate() error {
	return dara.Validate(s)
}
