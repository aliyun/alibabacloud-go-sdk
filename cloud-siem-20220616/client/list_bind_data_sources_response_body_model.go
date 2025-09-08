// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBindDataSourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListBindDataSourcesResponseBodyData) *ListBindDataSourcesResponseBody
	GetData() []*ListBindDataSourcesResponseBodyData
	SetRequestId(v string) *ListBindDataSourcesResponseBody
	GetRequestId() *string
}

type ListBindDataSourcesResponseBody struct {
	// The data returned.
	Data []*ListBindDataSourcesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListBindDataSourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListBindDataSourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListBindDataSourcesResponseBody) GetData() []*ListBindDataSourcesResponseBodyData {
	return s.Data
}

func (s *ListBindDataSourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListBindDataSourcesResponseBody) SetData(v []*ListBindDataSourcesResponseBodyData) *ListBindDataSourcesResponseBody {
	s.Data = v
	return s
}

func (s *ListBindDataSourcesResponseBody) SetRequestId(v string) *ListBindDataSourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBindDataSourcesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListBindDataSourcesResponseBodyData struct {
	// The ID of the cloud account.
	//
	// example:
	//
	// 123xxxxxxx
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The username of the cloud account.
	//
	// example:
	//
	// sas_tq_account_xxxx
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The code of the cloud service provider. Valid values:
	//
	// 	- qcloud: Tencent Cloud
	//
	// 	- aliyun: Alibaba Cloud
	//
	// 	- hcloud: Huawei Cloud
	//
	// example:
	//
	// hcloud
	CloudCode *string `json:"CloudCode,omitempty" xml:"CloudCode,omitempty"`
	// The ID of the data source. The ID is an MD5 hash value that is calculated by the threat analysis feature based on specific parameters.
	//
	// example:
	//
	// 220ba97c9d1fdb0b9c7e8c7ca328d7ea
	DataSourceInstanceId *string `json:"DataSourceInstanceId,omitempty" xml:"DataSourceInstanceId,omitempty"`
	// The name of the data source.
	//
	// example:
	//
	// waf_kafka
	DataSourceName *string `json:"DataSourceName,omitempty" xml:"DataSourceName,omitempty"`
	// The remarks on the data source.
	//
	// example:
	//
	// waf_kafka
	DataSourceRemark *string `json:"DataSourceRemark,omitempty" xml:"DataSourceRemark,omitempty"`
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
	// The number of logs that are added within the data source.
	//
	// example:
	//
	// 1
	LogCount *int32 `json:"LogCount,omitempty" xml:"LogCount,omitempty"`
	// The number of existing tasks that are created to add logs within the data source.
	//
	// example:
	//
	// 0
	TaskCount *int32 `json:"TaskCount,omitempty" xml:"TaskCount,omitempty"`
}

func (s ListBindDataSourcesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListBindDataSourcesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListBindDataSourcesResponseBodyData) GetAccountId() *string {
	return s.AccountId
}

func (s *ListBindDataSourcesResponseBodyData) GetAccountName() *string {
	return s.AccountName
}

func (s *ListBindDataSourcesResponseBodyData) GetCloudCode() *string {
	return s.CloudCode
}

func (s *ListBindDataSourcesResponseBodyData) GetDataSourceInstanceId() *string {
	return s.DataSourceInstanceId
}

func (s *ListBindDataSourcesResponseBodyData) GetDataSourceName() *string {
	return s.DataSourceName
}

func (s *ListBindDataSourcesResponseBodyData) GetDataSourceRemark() *string {
	return s.DataSourceRemark
}

func (s *ListBindDataSourcesResponseBodyData) GetDataSourceType() *string {
	return s.DataSourceType
}

func (s *ListBindDataSourcesResponseBodyData) GetLogCount() *int32 {
	return s.LogCount
}

func (s *ListBindDataSourcesResponseBodyData) GetTaskCount() *int32 {
	return s.TaskCount
}

func (s *ListBindDataSourcesResponseBodyData) SetAccountId(v string) *ListBindDataSourcesResponseBodyData {
	s.AccountId = &v
	return s
}

func (s *ListBindDataSourcesResponseBodyData) SetAccountName(v string) *ListBindDataSourcesResponseBodyData {
	s.AccountName = &v
	return s
}

func (s *ListBindDataSourcesResponseBodyData) SetCloudCode(v string) *ListBindDataSourcesResponseBodyData {
	s.CloudCode = &v
	return s
}

func (s *ListBindDataSourcesResponseBodyData) SetDataSourceInstanceId(v string) *ListBindDataSourcesResponseBodyData {
	s.DataSourceInstanceId = &v
	return s
}

func (s *ListBindDataSourcesResponseBodyData) SetDataSourceName(v string) *ListBindDataSourcesResponseBodyData {
	s.DataSourceName = &v
	return s
}

func (s *ListBindDataSourcesResponseBodyData) SetDataSourceRemark(v string) *ListBindDataSourcesResponseBodyData {
	s.DataSourceRemark = &v
	return s
}

func (s *ListBindDataSourcesResponseBodyData) SetDataSourceType(v string) *ListBindDataSourcesResponseBodyData {
	s.DataSourceType = &v
	return s
}

func (s *ListBindDataSourcesResponseBodyData) SetLogCount(v int32) *ListBindDataSourcesResponseBodyData {
	s.LogCount = &v
	return s
}

func (s *ListBindDataSourcesResponseBodyData) SetTaskCount(v int32) *ListBindDataSourcesResponseBodyData {
	s.TaskCount = &v
	return s
}

func (s *ListBindDataSourcesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
