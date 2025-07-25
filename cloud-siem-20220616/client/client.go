// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type DataProductListLogMapValue struct {
	// The code of the log.
	//
	// example:
	//
	// cloud_siem_config_log
	LogCode *string `json:"LogCode,omitempty" xml:"LogCode,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// audit log
	LogName *string `json:"LogName,omitempty" xml:"LogName,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// audit log
	LogNameEn *string `json:"LogNameEn,omitempty" xml:"LogNameEn,omitempty"`
	// The language code of the log that is used to indicate the language in which the log is displayed.
	//
	// example:
	//
	// ${sas.cloudsiem.prod.cloud_siem_aegis_crack_from_beaver}
	LogNameKey *string `json:"LogNameKey,omitempty" xml:"LogNameKey,omitempty"`
	// The status of the log delivery. Valid values:
	//
	// 	- true: The logs are being delivered.
	//
	// 	- false: The log delivery feature is disabled.
	//
	// example:
	//
	// true
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the log delivery feature can be enabled or disabled. The feature can be enabled or disabled only by the administrator of the threat analysis feature. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	CanOperateOrNot *bool `json:"CanOperateOrNot,omitempty" xml:"CanOperateOrNot,omitempty"`
	// The topic of the log in the Logstore. The value is an index field in the Logstore that can be used to distinguish different logs.
	//
	// example:
	//
	// sas_login_event
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// The extended parameter.
	ExtraParameters []*DataProductListLogMapValueExtraParameters `json:"ExtraParameters,omitempty" xml:"ExtraParameters,omitempty" type:"Repeated"`
}

func (s DataProductListLogMapValue) String() string {
	return tea.Prettify(s)
}

func (s DataProductListLogMapValue) GoString() string {
	return s.String()
}

func (s *DataProductListLogMapValue) SetLogCode(v string) *DataProductListLogMapValue {
	s.LogCode = &v
	return s
}

func (s *DataProductListLogMapValue) SetLogName(v string) *DataProductListLogMapValue {
	s.LogName = &v
	return s
}

func (s *DataProductListLogMapValue) SetLogNameEn(v string) *DataProductListLogMapValue {
	s.LogNameEn = &v
	return s
}

func (s *DataProductListLogMapValue) SetLogNameKey(v string) *DataProductListLogMapValue {
	s.LogNameKey = &v
	return s
}

func (s *DataProductListLogMapValue) SetStatus(v bool) *DataProductListLogMapValue {
	s.Status = &v
	return s
}

func (s *DataProductListLogMapValue) SetCanOperateOrNot(v bool) *DataProductListLogMapValue {
	s.CanOperateOrNot = &v
	return s
}

func (s *DataProductListLogMapValue) SetTopic(v string) *DataProductListLogMapValue {
	s.Topic = &v
	return s
}

func (s *DataProductListLogMapValue) SetExtraParameters(v []*DataProductListLogMapValueExtraParameters) *DataProductListLogMapValue {
	s.ExtraParameters = v
	return s
}

type DataProductListLogMapValueExtraParameters struct {
	// The ID of the extended parameter.
	//
	// example:
	//
	// flag
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the extended parameter.
	//
	// example:
	//
	// value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DataProductListLogMapValueExtraParameters) String() string {
	return tea.Prettify(s)
}

func (s DataProductListLogMapValueExtraParameters) GoString() string {
	return s.String()
}

func (s *DataProductListLogMapValueExtraParameters) SetKey(v string) *DataProductListLogMapValueExtraParameters {
	s.Key = &v
	return s
}

func (s *DataProductListLogMapValueExtraParameters) SetValue(v string) *DataProductListLogMapValueExtraParameters {
	s.Value = &v
	return s
}

type AddDataSourceRequest struct {
	// The ID of the cloud account.
	//
	// example:
	//
	// 123xxxxxxxx
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The code of the cloud service provider.
	//
	// Valid values:
	//
	// 	- qcloud
	//
	// 	- hcloud
	//
	// 	- aliyun
	//
	// This parameter is required.
	//
	// example:
	//
	// hcloud
	CloudCode *string `json:"CloudCode,omitempty" xml:"CloudCode,omitempty"`
	// The name of the data source.
	//
	// example:
	//
	// beijing_waf_kafka
	DataSourceInstanceName *string `json:"DataSourceInstanceName,omitempty" xml:"DataSourceInstanceName,omitempty"`
	// The parameters of the data source. Set this parameter to a JSON array.
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

func (s AddDataSourceRequest) String() string {
	return tea.Prettify(s)
}

func (s AddDataSourceRequest) GoString() string {
	return s.String()
}

func (s *AddDataSourceRequest) SetAccountId(v string) *AddDataSourceRequest {
	s.AccountId = &v
	return s
}

func (s *AddDataSourceRequest) SetCloudCode(v string) *AddDataSourceRequest {
	s.CloudCode = &v
	return s
}

func (s *AddDataSourceRequest) SetDataSourceInstanceName(v string) *AddDataSourceRequest {
	s.DataSourceInstanceName = &v
	return s
}

func (s *AddDataSourceRequest) SetDataSourceInstanceParams(v string) *AddDataSourceRequest {
	s.DataSourceInstanceParams = &v
	return s
}

func (s *AddDataSourceRequest) SetDataSourceInstanceRemark(v string) *AddDataSourceRequest {
	s.DataSourceInstanceRemark = &v
	return s
}

func (s *AddDataSourceRequest) SetDataSourceType(v string) *AddDataSourceRequest {
	s.DataSourceType = &v
	return s
}

func (s *AddDataSourceRequest) SetRegionId(v string) *AddDataSourceRequest {
	s.RegionId = &v
	return s
}

type AddDataSourceResponseBody struct {
	// The data returned.
	Data *AddDataSourceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddDataSourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *AddDataSourceResponseBody) SetData(v *AddDataSourceResponseBodyData) *AddDataSourceResponseBody {
	s.Data = v
	return s
}

func (s *AddDataSourceResponseBody) SetRequestId(v string) *AddDataSourceResponseBody {
	s.RequestId = &v
	return s
}

type AddDataSourceResponseBodyData struct {
	// The number of data sources that are added. The value 1 indicates that data source is added, and a value less than or equal to 0 indicates that the data source failed to be added.
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The ID of the data source. The ID is an MD5 hash value that is calculated by the threat analysis feature based on specific parameters.
	//
	// example:
	//
	// 220ba97c9d1fdb0b9c7e8c7ca328d7ea
	DataSourceInstanceId *string `json:"DataSourceInstanceId,omitempty" xml:"DataSourceInstanceId,omitempty"`
}

func (s AddDataSourceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AddDataSourceResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddDataSourceResponseBodyData) SetCount(v int32) *AddDataSourceResponseBodyData {
	s.Count = &v
	return s
}

func (s *AddDataSourceResponseBodyData) SetDataSourceInstanceId(v string) *AddDataSourceResponseBodyData {
	s.DataSourceInstanceId = &v
	return s
}

type AddDataSourceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddDataSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddDataSourceResponse) String() string {
	return tea.Prettify(s)
}

func (s AddDataSourceResponse) GoString() string {
	return s.String()
}

func (s *AddDataSourceResponse) SetHeaders(v map[string]*string) *AddDataSourceResponse {
	s.Headers = v
	return s
}

func (s *AddDataSourceResponse) SetStatusCode(v int32) *AddDataSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *AddDataSourceResponse) SetBody(v *AddDataSourceResponseBody) *AddDataSourceResponse {
	s.Body = v
	return s
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
	// The parameters of the data source. Set this parameter to a JSON array.
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

func (s AddDataSourceLogRequest) String() string {
	return tea.Prettify(s)
}

func (s AddDataSourceLogRequest) GoString() string {
	return s.String()
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

type AddDataSourceLogResponseBody struct {
	// The data returned.
	Data *AddDataSourceLogResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddDataSourceLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddDataSourceLogResponseBody) GoString() string {
	return s.String()
}

func (s *AddDataSourceLogResponseBody) SetData(v *AddDataSourceLogResponseBodyData) *AddDataSourceLogResponseBody {
	s.Data = v
	return s
}

func (s *AddDataSourceLogResponseBody) SetRequestId(v string) *AddDataSourceLogResponseBody {
	s.RequestId = &v
	return s
}

type AddDataSourceLogResponseBodyData struct {
	// The number of logs that are added. The value 1 indicates that the log is added, and a value less than or equal to 0 indicates that the log failed to be added.
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The ID of the log. The ID is an MD5 hash value that is calculated by the threat analysis feature based on specific parameters.
	//
	// example:
	//
	// ef33097c9d1fdb0b9c7e8c7ca320pkl1
	LogInstanceId *string `json:"LogInstanceId,omitempty" xml:"LogInstanceId,omitempty"`
}

func (s AddDataSourceLogResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AddDataSourceLogResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddDataSourceLogResponseBodyData) SetCount(v int32) *AddDataSourceLogResponseBodyData {
	s.Count = &v
	return s
}

func (s *AddDataSourceLogResponseBodyData) SetLogInstanceId(v string) *AddDataSourceLogResponseBodyData {
	s.LogInstanceId = &v
	return s
}

type AddDataSourceLogResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddDataSourceLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddDataSourceLogResponse) String() string {
	return tea.Prettify(s)
}

func (s AddDataSourceLogResponse) GoString() string {
	return s.String()
}

func (s *AddDataSourceLogResponse) SetHeaders(v map[string]*string) *AddDataSourceLogResponse {
	s.Headers = v
	return s
}

func (s *AddDataSourceLogResponse) SetStatusCode(v int32) *AddDataSourceLogResponse {
	s.StatusCode = &v
	return s
}

func (s *AddDataSourceLogResponse) SetBody(v *AddDataSourceLogResponseBody) *AddDataSourceLogResponse {
	s.Body = v
	return s
}

type AddUserSourceLogConfigRequest struct {
	// Specifies whether to add logs or delete added logs. Valid values:
	//
	// 	- \\-1: deletes added logs.
	//
	// 	- 0: adds logs.
	//
	// example:
	//
	// 0
	Deleted *int32 `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	// The display details of the Logstore.
	//
	// example:
	//
	// cn-shanghai.siem-project.siem-logstore
	DisPlayLine *string `json:"DisPlayLine,omitempty" xml:"DisPlayLine,omitempty"`
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
	// The log code.
	//
	// example:
	//
	// cloud_siem_aegis_proc
	SourceLogCode *string `json:"SourceLogCode,omitempty" xml:"SourceLogCode,omitempty"`
	// The details of the Logstore that you want to use in the JSON string format.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"project":"wafnew-project-1335759343513432-cn-hangzhou","logStore":"wafnew-logstore","regionCode":"cn-hangzhou","prodCode":"waf"}
	SourceLogInfo *string `json:"SourceLogInfo,omitempty" xml:"SourceLogInfo,omitempty"`
	// The code of the cloud service.
	//
	// example:
	//
	// sas
	SourceProdCode *string `json:"SourceProdCode,omitempty" xml:"SourceProdCode,omitempty"`
	// The ID of the Alibaba Cloud account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123XXXXXX
	SubUserId *int64 `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
}

func (s AddUserSourceLogConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s AddUserSourceLogConfigRequest) GoString() string {
	return s.String()
}

func (s *AddUserSourceLogConfigRequest) SetDeleted(v int32) *AddUserSourceLogConfigRequest {
	s.Deleted = &v
	return s
}

func (s *AddUserSourceLogConfigRequest) SetDisPlayLine(v string) *AddUserSourceLogConfigRequest {
	s.DisPlayLine = &v
	return s
}

func (s *AddUserSourceLogConfigRequest) SetRegionId(v string) *AddUserSourceLogConfigRequest {
	s.RegionId = &v
	return s
}

func (s *AddUserSourceLogConfigRequest) SetSourceLogCode(v string) *AddUserSourceLogConfigRequest {
	s.SourceLogCode = &v
	return s
}

func (s *AddUserSourceLogConfigRequest) SetSourceLogInfo(v string) *AddUserSourceLogConfigRequest {
	s.SourceLogInfo = &v
	return s
}

func (s *AddUserSourceLogConfigRequest) SetSourceProdCode(v string) *AddUserSourceLogConfigRequest {
	s.SourceProdCode = &v
	return s
}

func (s *AddUserSourceLogConfigRequest) SetSubUserId(v int64) *AddUserSourceLogConfigRequest {
	s.SubUserId = &v
	return s
}

type AddUserSourceLogConfigResponseBody struct {
	// The data returned.
	Data *AddUserSourceLogConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddUserSourceLogConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddUserSourceLogConfigResponseBody) GoString() string {
	return s.String()
}

func (s *AddUserSourceLogConfigResponseBody) SetData(v *AddUserSourceLogConfigResponseBodyData) *AddUserSourceLogConfigResponseBody {
	s.Data = v
	return s
}

func (s *AddUserSourceLogConfigResponseBody) SetRequestId(v string) *AddUserSourceLogConfigResponseBody {
	s.RequestId = &v
	return s
}

type AddUserSourceLogConfigResponseBodyData struct {
	// The display details of the Logstore.
	//
	// example:
	//
	// cn-shanghai.siem-project.siem-logstore
	DiplayLine *string `json:"DiplayLine,omitempty" xml:"DiplayLine,omitempty"`
	// Indicates whether the details of added logs are returned. Valid values: true false
	//
	// example:
	//
	// 0
	Displayed *bool `json:"Displayed,omitempty" xml:"Displayed,omitempty"`
	// Indicates whether the logs are added to the threat analysis feature. Valid values: true false
	//
	// example:
	//
	// 0
	Imported *bool `json:"Imported,omitempty" xml:"Imported,omitempty"`
	// The ID of the Alibaba Cloud account that is used to purchase the threat analysis feature.
	//
	// example:
	//
	// 123XXXXXXXXX
	MainUserId *int64 `json:"MainUserId,omitempty" xml:"MainUserId,omitempty"`
	// The log code.
	//
	// example:
	//
	// cloud_siem_aegis_proc
	SourceLogCode *string `json:"SourceLogCode,omitempty" xml:"SourceLogCode,omitempty"`
	// The code of the cloud service.
	//
	// example:
	//
	// sas
	SourceProdCode *string `json:"SourceProdCode,omitempty" xml:"SourceProdCode,omitempty"`
	// The ID of the Alibaba Cloud account that can be used to perform operations supported by the threat analysis feature.
	//
	// example:
	//
	// 123XXXXXXXX
	SubUserId *int64 `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
	// The username of the Alibaba Cloud account that can be used to perform operations supported by the threat analysis feature.
	//
	// example:
	//
	// sas_account_xxx
	SubUserName *string `json:"SubUserName,omitempty" xml:"SubUserName,omitempty"`
}

func (s AddUserSourceLogConfigResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AddUserSourceLogConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddUserSourceLogConfigResponseBodyData) SetDiplayLine(v string) *AddUserSourceLogConfigResponseBodyData {
	s.DiplayLine = &v
	return s
}

func (s *AddUserSourceLogConfigResponseBodyData) SetDisplayed(v bool) *AddUserSourceLogConfigResponseBodyData {
	s.Displayed = &v
	return s
}

func (s *AddUserSourceLogConfigResponseBodyData) SetImported(v bool) *AddUserSourceLogConfigResponseBodyData {
	s.Imported = &v
	return s
}

func (s *AddUserSourceLogConfigResponseBodyData) SetMainUserId(v int64) *AddUserSourceLogConfigResponseBodyData {
	s.MainUserId = &v
	return s
}

func (s *AddUserSourceLogConfigResponseBodyData) SetSourceLogCode(v string) *AddUserSourceLogConfigResponseBodyData {
	s.SourceLogCode = &v
	return s
}

func (s *AddUserSourceLogConfigResponseBodyData) SetSourceProdCode(v string) *AddUserSourceLogConfigResponseBodyData {
	s.SourceProdCode = &v
	return s
}

func (s *AddUserSourceLogConfigResponseBodyData) SetSubUserId(v int64) *AddUserSourceLogConfigResponseBodyData {
	s.SubUserId = &v
	return s
}

func (s *AddUserSourceLogConfigResponseBodyData) SetSubUserName(v string) *AddUserSourceLogConfigResponseBodyData {
	s.SubUserName = &v
	return s
}

type AddUserSourceLogConfigResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddUserSourceLogConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddUserSourceLogConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s AddUserSourceLogConfigResponse) GoString() string {
	return s.String()
}

func (s *AddUserSourceLogConfigResponse) SetHeaders(v map[string]*string) *AddUserSourceLogConfigResponse {
	s.Headers = v
	return s
}

func (s *AddUserSourceLogConfigResponse) SetStatusCode(v int32) *AddUserSourceLogConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *AddUserSourceLogConfigResponse) SetBody(v *AddUserSourceLogConfigResponseBody) *AddUserSourceLogConfigResponse {
	s.Body = v
	return s
}

type BindAccountRequest struct {
	// The AccessKey ID of the cloud account.
	//
	// This parameter is required.
	//
	// example:
	//
	// ABCXXXXXXXX
	AccessId *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	// The ID of the cloud account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123xxxxxxx
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The username of the cloud account.
	//
	// This parameter is required.
	//
	// example:
	//
	// xxxx
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The code of the cloud service provider. Valid values:
	//
	// 	- aliyun: Alibaba Cloud
	//
	// 	- hcloud: Huawei Cloud
	//
	// 	- qcloud: Tencent Cloud
	//
	// This parameter is required.
	//
	// example:
	//
	// hcloud
	CloudCode *string `json:"CloudCode,omitempty" xml:"CloudCode,omitempty"`
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
	RoleFor  *int64  `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	RoleType *int32  `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s BindAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s BindAccountRequest) GoString() string {
	return s.String()
}

func (s *BindAccountRequest) SetAccessId(v string) *BindAccountRequest {
	s.AccessId = &v
	return s
}

func (s *BindAccountRequest) SetAccountId(v string) *BindAccountRequest {
	s.AccountId = &v
	return s
}

func (s *BindAccountRequest) SetAccountName(v string) *BindAccountRequest {
	s.AccountName = &v
	return s
}

func (s *BindAccountRequest) SetCloudCode(v string) *BindAccountRequest {
	s.CloudCode = &v
	return s
}

func (s *BindAccountRequest) SetRegionId(v string) *BindAccountRequest {
	s.RegionId = &v
	return s
}

func (s *BindAccountRequest) SetRoleFor(v int64) *BindAccountRequest {
	s.RoleFor = &v
	return s
}

func (s *BindAccountRequest) SetRoleType(v int32) *BindAccountRequest {
	s.RoleType = &v
	return s
}

type BindAccountResponseBody struct {
	// The data returned.
	Data *BindAccountResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BindAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BindAccountResponseBody) GoString() string {
	return s.String()
}

func (s *BindAccountResponseBody) SetData(v *BindAccountResponseBodyData) *BindAccountResponseBody {
	s.Data = v
	return s
}

func (s *BindAccountResponseBody) SetRequestId(v string) *BindAccountResponseBody {
	s.RequestId = &v
	return s
}

type BindAccountResponseBodyData struct {
	// The number of the cloud accounts that are added to the threat analysis feature.
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s BindAccountResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s BindAccountResponseBodyData) GoString() string {
	return s.String()
}

func (s *BindAccountResponseBodyData) SetCount(v int32) *BindAccountResponseBodyData {
	s.Count = &v
	return s
}

type BindAccountResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s BindAccountResponse) GoString() string {
	return s.String()
}

func (s *BindAccountResponse) SetHeaders(v map[string]*string) *BindAccountResponse {
	s.Headers = v
	return s
}

func (s *BindAccountResponse) SetStatusCode(v int32) *BindAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *BindAccountResponse) SetBody(v *BindAccountResponseBody) *BindAccountResponse {
	s.Body = v
	return s
}

type CloseDeliveryRequest struct {
	// The log code of the cloud service, such as the code of the process log for Security Center. You can obtain the log code from the response of the ListDelivery operation.
	//
	// example:
	//
	// cloud_siem_aegis_proc
	LogCode *string `json:"LogCode,omitempty" xml:"LogCode,omitempty"`
	// The code of the cloud service. Valid values:
	//
	// 	- qcloud_waf
	//
	// 	- qlcoud_cfw
	//
	// 	- hcloud_waf
	//
	// 	- hcloud_cfw
	//
	// 	- ddos
	//
	// 	- sas
	//
	// 	- cfw
	//
	// 	- config
	//
	// 	- csk
	//
	// 	- fc
	//
	// 	- rds
	//
	// 	- nas
	//
	// 	- apigateway
	//
	// 	- cdn
	//
	// 	- mongodb
	//
	// 	- eip
	//
	// 	- slb
	//
	// 	- vpc
	//
	// 	- actiontrail
	//
	// 	- waf
	//
	// 	- bastionhost
	//
	// 	- oss
	//
	// 	- polardb
	//
	// This parameter is required.
	//
	// example:
	//
	// sas
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The region in which the data management center of the threat analysis feature resides. Specify this parameter based on the region where your assets reside. Valid values:
	//
	// 	- cn-hangzhou: Your assets reside in regions in China.
	//
	// 	- ap-southeast-1: Your assets reside in regions outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the account that you switch from the management account.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view. Valid values:
	//
	// - 0: the current Alibaba Cloud account
	//
	// - 1: the global account
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s CloseDeliveryRequest) String() string {
	return tea.Prettify(s)
}

func (s CloseDeliveryRequest) GoString() string {
	return s.String()
}

func (s *CloseDeliveryRequest) SetLogCode(v string) *CloseDeliveryRequest {
	s.LogCode = &v
	return s
}

func (s *CloseDeliveryRequest) SetProductCode(v string) *CloseDeliveryRequest {
	s.ProductCode = &v
	return s
}

func (s *CloseDeliveryRequest) SetRegionId(v string) *CloseDeliveryRequest {
	s.RegionId = &v
	return s
}

func (s *CloseDeliveryRequest) SetRoleFor(v int64) *CloseDeliveryRequest {
	s.RoleFor = &v
	return s
}

func (s *CloseDeliveryRequest) SetRoleType(v int32) *CloseDeliveryRequest {
	s.RoleType = &v
	return s
}

type CloseDeliveryResponseBody struct {
	// Indicates whether the threat analysis feature was disabled. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F375A043-4F5B-55F2-A564-CC47FFC6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloseDeliveryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CloseDeliveryResponseBody) GoString() string {
	return s.String()
}

func (s *CloseDeliveryResponseBody) SetData(v bool) *CloseDeliveryResponseBody {
	s.Data = &v
	return s
}

func (s *CloseDeliveryResponseBody) SetRequestId(v string) *CloseDeliveryResponseBody {
	s.RequestId = &v
	return s
}

type CloseDeliveryResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloseDeliveryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloseDeliveryResponse) String() string {
	return tea.Prettify(s)
}

func (s CloseDeliveryResponse) GoString() string {
	return s.String()
}

func (s *CloseDeliveryResponse) SetHeaders(v map[string]*string) *CloseDeliveryResponse {
	s.Headers = v
	return s
}

func (s *CloseDeliveryResponse) SetStatusCode(v int32) *CloseDeliveryResponse {
	s.StatusCode = &v
	return s
}

func (s *CloseDeliveryResponse) SetBody(v *CloseDeliveryResponseBody) *CloseDeliveryResponse {
	s.Body = v
	return s
}

type DeleteAutomateResponseConfigRequest struct {
	// The ID of the rule.
	//
	// example:
	//
	// 123
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
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
	// The ID of the account that you switch from the management account.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view. Valid values:
	//
	// - 0: the current Alibaba Cloud account
	//
	// - 1: the global account
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s DeleteAutomateResponseConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAutomateResponseConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteAutomateResponseConfigRequest) SetId(v int64) *DeleteAutomateResponseConfigRequest {
	s.Id = &v
	return s
}

func (s *DeleteAutomateResponseConfigRequest) SetRegionId(v string) *DeleteAutomateResponseConfigRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteAutomateResponseConfigRequest) SetRoleFor(v int64) *DeleteAutomateResponseConfigRequest {
	s.RoleFor = &v
	return s
}

func (s *DeleteAutomateResponseConfigRequest) SetRoleType(v int32) *DeleteAutomateResponseConfigRequest {
	s.RoleType = &v
	return s
}

type DeleteAutomateResponseConfigResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	//
	// example:
	//
	// 123456
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteAutomateResponseConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAutomateResponseConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAutomateResponseConfigResponseBody) SetCode(v int32) *DeleteAutomateResponseConfigResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteAutomateResponseConfigResponseBody) SetData(v string) *DeleteAutomateResponseConfigResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteAutomateResponseConfigResponseBody) SetMessage(v string) *DeleteAutomateResponseConfigResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteAutomateResponseConfigResponseBody) SetRequestId(v string) *DeleteAutomateResponseConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAutomateResponseConfigResponseBody) SetSuccess(v bool) *DeleteAutomateResponseConfigResponseBody {
	s.Success = &v
	return s
}

type DeleteAutomateResponseConfigResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAutomateResponseConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAutomateResponseConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAutomateResponseConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteAutomateResponseConfigResponse) SetHeaders(v map[string]*string) *DeleteAutomateResponseConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteAutomateResponseConfigResponse) SetStatusCode(v int32) *DeleteAutomateResponseConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAutomateResponseConfigResponse) SetBody(v *DeleteAutomateResponseConfigResponseBody) *DeleteAutomateResponseConfigResponse {
	s.Body = v
	return s
}

type DeleteBindAccountRequest struct {
	// The AccessKey ID of the cloud account.
	//
	// This parameter is required.
	//
	// example:
	//
	// ABCXXXXXXXX
	AccessId *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	// The ID of the cloud account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123xxxxxxx
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The ID generated when the account is added to the threat analysis feature. You can call the [ListBindAccount](https://api.aliyun-inc.com/#/publishment/document/cloud-siem/863fdf54478f4cc5877e27c2a5fe9e44?tenantUuid=f382fccd88b94c5c8c864def6815b854\\&activeTabKey=api%7CListBindAccount) operation to query the ID.
	//
	// example:
	//
	// 10
	BindId *int64 `json:"BindId,omitempty" xml:"BindId,omitempty"`
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
	RoleFor  *int64  `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	RoleType *int32  `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s DeleteBindAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteBindAccountRequest) GoString() string {
	return s.String()
}

func (s *DeleteBindAccountRequest) SetAccessId(v string) *DeleteBindAccountRequest {
	s.AccessId = &v
	return s
}

func (s *DeleteBindAccountRequest) SetAccountId(v string) *DeleteBindAccountRequest {
	s.AccountId = &v
	return s
}

func (s *DeleteBindAccountRequest) SetBindId(v int64) *DeleteBindAccountRequest {
	s.BindId = &v
	return s
}

func (s *DeleteBindAccountRequest) SetCloudCode(v string) *DeleteBindAccountRequest {
	s.CloudCode = &v
	return s
}

func (s *DeleteBindAccountRequest) SetRegionId(v string) *DeleteBindAccountRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteBindAccountRequest) SetRoleFor(v int64) *DeleteBindAccountRequest {
	s.RoleFor = &v
	return s
}

func (s *DeleteBindAccountRequest) SetRoleType(v int32) *DeleteBindAccountRequest {
	s.RoleType = &v
	return s
}

type DeleteBindAccountResponseBody struct {
	// The data returned.
	Data *DeleteBindAccountResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteBindAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteBindAccountResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBindAccountResponseBody) SetData(v *DeleteBindAccountResponseBodyData) *DeleteBindAccountResponseBody {
	s.Data = v
	return s
}

func (s *DeleteBindAccountResponseBody) SetRequestId(v string) *DeleteBindAccountResponseBody {
	s.RequestId = &v
	return s
}

type DeleteBindAccountResponseBodyData struct {
	// The number of cloud accounts that are removed. The value 1 indicates that cloud account is removed, and a value less than or equal to 0 indicates that the cloud account failed to be removed.
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DeleteBindAccountResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DeleteBindAccountResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteBindAccountResponseBodyData) SetCount(v int32) *DeleteBindAccountResponseBodyData {
	s.Count = &v
	return s
}

type DeleteBindAccountResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteBindAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteBindAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteBindAccountResponse) GoString() string {
	return s.String()
}

func (s *DeleteBindAccountResponse) SetHeaders(v map[string]*string) *DeleteBindAccountResponse {
	s.Headers = v
	return s
}

func (s *DeleteBindAccountResponse) SetStatusCode(v int32) *DeleteBindAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBindAccountResponse) SetBody(v *DeleteBindAccountResponseBody) *DeleteBindAccountResponse {
	s.Body = v
	return s
}

type DeleteCustomizeRuleRequest struct {
	// The region in which the service is deployed.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the account that you switch from the management account.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view. Valid values:
	//
	// - 0: the current Alibaba Cloud account
	//
	// - 1: the global account
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// The ID of the rule.
	//
	// example:
	//
	// 123456789
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DeleteCustomizeRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCustomizeRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteCustomizeRuleRequest) SetRegionId(v string) *DeleteCustomizeRuleRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteCustomizeRuleRequest) SetRoleFor(v int64) *DeleteCustomizeRuleRequest {
	s.RoleFor = &v
	return s
}

func (s *DeleteCustomizeRuleRequest) SetRoleType(v int32) *DeleteCustomizeRuleRequest {
	s.RoleType = &v
	return s
}

func (s *DeleteCustomizeRuleRequest) SetRuleId(v int64) *DeleteCustomizeRuleRequest {
	s.RuleId = &v
	return s
}

type DeleteCustomizeRuleResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	//
	// example:
	//
	// 123456
	Data *int32 `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteCustomizeRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCustomizeRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCustomizeRuleResponseBody) SetCode(v int32) *DeleteCustomizeRuleResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteCustomizeRuleResponseBody) SetData(v int32) *DeleteCustomizeRuleResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteCustomizeRuleResponseBody) SetMessage(v string) *DeleteCustomizeRuleResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteCustomizeRuleResponseBody) SetRequestId(v string) *DeleteCustomizeRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCustomizeRuleResponseBody) SetSuccess(v bool) *DeleteCustomizeRuleResponseBody {
	s.Success = &v
	return s
}

type DeleteCustomizeRuleResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCustomizeRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCustomizeRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCustomizeRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteCustomizeRuleResponse) SetHeaders(v map[string]*string) *DeleteCustomizeRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteCustomizeRuleResponse) SetStatusCode(v int32) *DeleteCustomizeRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCustomizeRuleResponse) SetBody(v *DeleteCustomizeRuleResponseBody) *DeleteCustomizeRuleResponse {
	s.Body = v
	return s
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
	return tea.Prettify(s)
}

func (s DeleteDataSourceRequest) GoString() string {
	return s.String()
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

type DeleteDataSourceResponseBody struct {
	// The data returned.
	Data *DeleteDataSourceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDataSourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDataSourceResponseBody) SetData(v *DeleteDataSourceResponseBodyData) *DeleteDataSourceResponseBody {
	s.Data = v
	return s
}

func (s *DeleteDataSourceResponseBody) SetRequestId(v string) *DeleteDataSourceResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDataSourceResponseBodyData struct {
	// The number of data sources that are removed. The value 1 indicates that data source is removed, and a value less than or equal to 0 indicates that the data source failed to be removed.
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DeleteDataSourceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DeleteDataSourceResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteDataSourceResponseBodyData) SetCount(v int32) *DeleteDataSourceResponseBodyData {
	s.Count = &v
	return s
}

type DeleteDataSourceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDataSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDataSourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDataSourceResponse) GoString() string {
	return s.String()
}

func (s *DeleteDataSourceResponse) SetHeaders(v map[string]*string) *DeleteDataSourceResponse {
	s.Headers = v
	return s
}

func (s *DeleteDataSourceResponse) SetStatusCode(v int32) *DeleteDataSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDataSourceResponse) SetBody(v *DeleteDataSourceResponseBody) *DeleteDataSourceResponse {
	s.Body = v
	return s
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
	return tea.Prettify(s)
}

func (s DeleteDataSourceLogRequest) GoString() string {
	return s.String()
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

type DeleteDataSourceLogResponseBody struct {
	// The data returned.
	Data *DeleteDataSourceLogResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDataSourceLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDataSourceLogResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDataSourceLogResponseBody) SetData(v *DeleteDataSourceLogResponseBodyData) *DeleteDataSourceLogResponseBody {
	s.Data = v
	return s
}

func (s *DeleteDataSourceLogResponseBody) SetRequestId(v string) *DeleteDataSourceLogResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDataSourceLogResponseBodyData struct {
	// The number of logs that are removed. The value 1 indicates that the log is removed, and a value less than or equal to 0 indicates that the log failed to be removed.
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The ID of the log. The ID is an MD5 hash value that is calculated by the threat analysis feature based on specific parameters.
	//
	// example:
	//
	// ef33097c9d1fdb0b9c7e8c7ca320pkl1
	LogInstanceId *string `json:"LogInstanceId,omitempty" xml:"LogInstanceId,omitempty"`
}

func (s DeleteDataSourceLogResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DeleteDataSourceLogResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteDataSourceLogResponseBodyData) SetCount(v int32) *DeleteDataSourceLogResponseBodyData {
	s.Count = &v
	return s
}

func (s *DeleteDataSourceLogResponseBodyData) SetLogInstanceId(v string) *DeleteDataSourceLogResponseBodyData {
	s.LogInstanceId = &v
	return s
}

type DeleteDataSourceLogResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDataSourceLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDataSourceLogResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDataSourceLogResponse) GoString() string {
	return s.String()
}

func (s *DeleteDataSourceLogResponse) SetHeaders(v map[string]*string) *DeleteDataSourceLogResponse {
	s.Headers = v
	return s
}

func (s *DeleteDataSourceLogResponse) SetStatusCode(v int32) *DeleteDataSourceLogResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDataSourceLogResponse) SetBody(v *DeleteDataSourceLogResponseBody) *DeleteDataSourceLogResponse {
	s.Body = v
	return s
}

type DeleteWhiteRuleListRequest struct {
	// The unique ID of the whitelist rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456789
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
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
	// The ID of the account that you switch from the management account.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view. Valid values:
	//
	// - 0: the current Alibaba Cloud account
	//
	// - 1: the global account
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s DeleteWhiteRuleListRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteWhiteRuleListRequest) GoString() string {
	return s.String()
}

func (s *DeleteWhiteRuleListRequest) SetId(v int64) *DeleteWhiteRuleListRequest {
	s.Id = &v
	return s
}

func (s *DeleteWhiteRuleListRequest) SetRegionId(v string) *DeleteWhiteRuleListRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteWhiteRuleListRequest) SetRoleFor(v int64) *DeleteWhiteRuleListRequest {
	s.RoleFor = &v
	return s
}

func (s *DeleteWhiteRuleListRequest) SetRoleType(v int32) *DeleteWhiteRuleListRequest {
	s.RoleType = &v
	return s
}

type DeleteWhiteRuleListResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	//
	// example:
	//
	// 123456
	Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteWhiteRuleListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteWhiteRuleListResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWhiteRuleListResponseBody) SetCode(v int32) *DeleteWhiteRuleListResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteWhiteRuleListResponseBody) SetData(v interface{}) *DeleteWhiteRuleListResponseBody {
	s.Data = v
	return s
}

func (s *DeleteWhiteRuleListResponseBody) SetMessage(v string) *DeleteWhiteRuleListResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteWhiteRuleListResponseBody) SetRequestId(v string) *DeleteWhiteRuleListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteWhiteRuleListResponseBody) SetSuccess(v bool) *DeleteWhiteRuleListResponseBody {
	s.Success = &v
	return s
}

type DeleteWhiteRuleListResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteWhiteRuleListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteWhiteRuleListResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteWhiteRuleListResponse) GoString() string {
	return s.String()
}

func (s *DeleteWhiteRuleListResponse) SetHeaders(v map[string]*string) *DeleteWhiteRuleListResponse {
	s.Headers = v
	return s
}

func (s *DeleteWhiteRuleListResponse) SetStatusCode(v int32) *DeleteWhiteRuleListResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteWhiteRuleListResponse) SetBody(v *DeleteWhiteRuleListResponseBody) *DeleteWhiteRuleListResponse {
	s.Body = v
	return s
}

type DescribeAggregateFunctionRequest struct {
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
	// The ID of the account that you switch from the management account.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view. Valid values:
	//
	// - 0: the current Alibaba Cloud account
	//
	// - 1: the global account
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s DescribeAggregateFunctionRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAggregateFunctionRequest) GoString() string {
	return s.String()
}

func (s *DescribeAggregateFunctionRequest) SetRegionId(v string) *DescribeAggregateFunctionRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAggregateFunctionRequest) SetRoleFor(v int64) *DescribeAggregateFunctionRequest {
	s.RoleFor = &v
	return s
}

func (s *DescribeAggregateFunctionRequest) SetRoleType(v int32) *DescribeAggregateFunctionRequest {
	s.RoleType = &v
	return s
}

type DescribeAggregateFunctionResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	//
	// example:
	//
	// 123456
	Data []*DescribeAggregateFunctionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeAggregateFunctionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAggregateFunctionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAggregateFunctionResponseBody) SetCode(v int32) *DescribeAggregateFunctionResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAggregateFunctionResponseBody) SetData(v []*DescribeAggregateFunctionResponseBodyData) *DescribeAggregateFunctionResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAggregateFunctionResponseBody) SetMessage(v string) *DescribeAggregateFunctionResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAggregateFunctionResponseBody) SetRequestId(v string) *DescribeAggregateFunctionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAggregateFunctionResponseBody) SetSuccess(v bool) *DescribeAggregateFunctionResponseBody {
	s.Success = &v
	return s
}

type DescribeAggregateFunctionResponseBodyData struct {
	// The aggregate function.
	//
	// example:
	//
	// count
	Function *string `json:"Function,omitempty" xml:"Function,omitempty"`
	// The display name of the aggregate function.
	//
	// example:
	//
	// Count
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
}

func (s DescribeAggregateFunctionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeAggregateFunctionResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAggregateFunctionResponseBodyData) SetFunction(v string) *DescribeAggregateFunctionResponseBodyData {
	s.Function = &v
	return s
}

func (s *DescribeAggregateFunctionResponseBodyData) SetFunctionName(v string) *DescribeAggregateFunctionResponseBodyData {
	s.FunctionName = &v
	return s
}

type DescribeAggregateFunctionResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAggregateFunctionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAggregateFunctionResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAggregateFunctionResponse) GoString() string {
	return s.String()
}

func (s *DescribeAggregateFunctionResponse) SetHeaders(v map[string]*string) *DescribeAggregateFunctionResponse {
	s.Headers = v
	return s
}

func (s *DescribeAggregateFunctionResponse) SetStatusCode(v int32) *DescribeAggregateFunctionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAggregateFunctionResponse) SetBody(v *DescribeAggregateFunctionResponseBody) *DescribeAggregateFunctionResponse {
	s.Body = v
	return s
}

type DescribeAlertSceneRequest struct {
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
	// The ID of the account that you switch from the management account.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view. Valid values:
	//
	// - 0: the current Alibaba Cloud account
	//
	// - 1: the global account
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s DescribeAlertSceneRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertSceneRequest) GoString() string {
	return s.String()
}

func (s *DescribeAlertSceneRequest) SetRegionId(v string) *DescribeAlertSceneRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAlertSceneRequest) SetRoleFor(v int64) *DescribeAlertSceneRequest {
	s.RoleFor = &v
	return s
}

func (s *DescribeAlertSceneRequest) SetRoleType(v int32) *DescribeAlertSceneRequest {
	s.RoleType = &v
	return s
}

type DescribeAlertSceneResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	//
	// example:
	//
	// 123456
	Data []*DescribeAlertSceneResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeAlertSceneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertSceneResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAlertSceneResponseBody) SetCode(v int32) *DescribeAlertSceneResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAlertSceneResponseBody) SetData(v []*DescribeAlertSceneResponseBodyData) *DescribeAlertSceneResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAlertSceneResponseBody) SetMessage(v string) *DescribeAlertSceneResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAlertSceneResponseBody) SetRequestId(v string) *DescribeAlertSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAlertSceneResponseBody) SetSuccess(v bool) *DescribeAlertSceneResponseBody {
	s.Success = &v
	return s
}

type DescribeAlertSceneResponseBodyData struct {
	// The name of the alert. The value varies based on the display language (Chinese or English) of the Security Center console.
	//
	// example:
	//
	// login_common_ip
	AlertName *string `json:"AlertName,omitempty" xml:"AlertName,omitempty"`
	// The ID of the alert name.
	//
	// example:
	//
	// login_common_ip
	AlertNameId *string `json:"AlertNameId,omitempty" xml:"AlertNameId,omitempty"`
	// The title of the alert notification. The value varies based on the display language (Chinese or English) of the Security Center console.
	//
	// example:
	//
	// unusual login-login_common_ip
	AlertTile *string `json:"AlertTile,omitempty" xml:"AlertTile,omitempty"`
	// The ID of the alert title.
	//
	// example:
	//
	// unusual login-login_common_ip
	AlertTileId *string `json:"AlertTileId,omitempty" xml:"AlertTileId,omitempty"`
	// The type of the alert. The value varies based on the display language (Chinese or English) of the Security Center console.
	//
	// example:
	//
	// unusual login
	AlertType *string `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
	// The ID of the alert type.
	//
	// example:
	//
	// unusual login
	AlertTypeId *string `json:"AlertTypeId,omitempty" xml:"AlertTypeId,omitempty"`
	// The information about the entities for which you need to add the alert to the whitelist.
	//
	// example:
	//
	// [{"Type": "host_uuid","Value": "441862da-a539-4cc0-a00d-473955826881","Values": ["441862da-a539-4cc0-a00d-473955826881"],"Name": "${aliyun.siem.entity.host_uuid}"}]
	Targets []*DescribeAlertSceneResponseBodyDataTargets `json:"Targets,omitempty" xml:"Targets,omitempty" type:"Repeated"`
}

func (s DescribeAlertSceneResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertSceneResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAlertSceneResponseBodyData) SetAlertName(v string) *DescribeAlertSceneResponseBodyData {
	s.AlertName = &v
	return s
}

func (s *DescribeAlertSceneResponseBodyData) SetAlertNameId(v string) *DescribeAlertSceneResponseBodyData {
	s.AlertNameId = &v
	return s
}

func (s *DescribeAlertSceneResponseBodyData) SetAlertTile(v string) *DescribeAlertSceneResponseBodyData {
	s.AlertTile = &v
	return s
}

func (s *DescribeAlertSceneResponseBodyData) SetAlertTileId(v string) *DescribeAlertSceneResponseBodyData {
	s.AlertTileId = &v
	return s
}

func (s *DescribeAlertSceneResponseBodyData) SetAlertType(v string) *DescribeAlertSceneResponseBodyData {
	s.AlertType = &v
	return s
}

func (s *DescribeAlertSceneResponseBodyData) SetAlertTypeId(v string) *DescribeAlertSceneResponseBodyData {
	s.AlertTypeId = &v
	return s
}

func (s *DescribeAlertSceneResponseBodyData) SetTargets(v []*DescribeAlertSceneResponseBodyDataTargets) *DescribeAlertSceneResponseBodyData {
	s.Targets = v
	return s
}

type DescribeAlertSceneResponseBodyDataTargets struct {
	// The display name of the attribute for the entity.
	//
	// example:
	//
	// HOST UUID
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The attribute of the entity.
	//
	// example:
	//
	// host_uuid
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The right operand that is displayed by default in the whitelist rule.
	//
	// example:
	//
	// 441862da-a539-4cc0-a00d-47395582****
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// The right operands supported by the whitelist rule.
	//
	// example:
	//
	// ["441862da-a539-4cc0-a00d-473955826881"]
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s DescribeAlertSceneResponseBodyDataTargets) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertSceneResponseBodyDataTargets) GoString() string {
	return s.String()
}

func (s *DescribeAlertSceneResponseBodyDataTargets) SetName(v string) *DescribeAlertSceneResponseBodyDataTargets {
	s.Name = &v
	return s
}

func (s *DescribeAlertSceneResponseBodyDataTargets) SetType(v string) *DescribeAlertSceneResponseBodyDataTargets {
	s.Type = &v
	return s
}

func (s *DescribeAlertSceneResponseBodyDataTargets) SetValue(v string) *DescribeAlertSceneResponseBodyDataTargets {
	s.Value = &v
	return s
}

func (s *DescribeAlertSceneResponseBodyDataTargets) SetValues(v []*string) *DescribeAlertSceneResponseBodyDataTargets {
	s.Values = v
	return s
}

type DescribeAlertSceneResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAlertSceneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAlertSceneResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertSceneResponse) GoString() string {
	return s.String()
}

func (s *DescribeAlertSceneResponse) SetHeaders(v map[string]*string) *DescribeAlertSceneResponse {
	s.Headers = v
	return s
}

func (s *DescribeAlertSceneResponse) SetStatusCode(v int32) *DescribeAlertSceneResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAlertSceneResponse) SetBody(v *DescribeAlertSceneResponseBody) *DescribeAlertSceneResponse {
	s.Body = v
	return s
}

type DescribeAlertSceneByEventRequest struct {
	// The ID of the event.
	//
	// This parameter is required.
	//
	// example:
	//
	// 85ea4241-798f-4684-a876-65d4f0c3****
	IncidentUuid *string `json:"IncidentUuid,omitempty" xml:"IncidentUuid,omitempty"`
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
	// The ID of the account that you switch from the management account.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view. Valid values:
	//
	// - 0: the current Alibaba Cloud account
	//
	// - 1: the global account
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s DescribeAlertSceneByEventRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertSceneByEventRequest) GoString() string {
	return s.String()
}

func (s *DescribeAlertSceneByEventRequest) SetIncidentUuid(v string) *DescribeAlertSceneByEventRequest {
	s.IncidentUuid = &v
	return s
}

func (s *DescribeAlertSceneByEventRequest) SetRegionId(v string) *DescribeAlertSceneByEventRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAlertSceneByEventRequest) SetRoleFor(v int64) *DescribeAlertSceneByEventRequest {
	s.RoleFor = &v
	return s
}

func (s *DescribeAlertSceneByEventRequest) SetRoleType(v int32) *DescribeAlertSceneByEventRequest {
	s.RoleType = &v
	return s
}

type DescribeAlertSceneByEventResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	//
	// example:
	//
	// 123456
	Data []*DescribeAlertSceneByEventResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeAlertSceneByEventResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertSceneByEventResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAlertSceneByEventResponseBody) SetCode(v int32) *DescribeAlertSceneByEventResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAlertSceneByEventResponseBody) SetData(v []*DescribeAlertSceneByEventResponseBodyData) *DescribeAlertSceneByEventResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAlertSceneByEventResponseBody) SetMessage(v string) *DescribeAlertSceneByEventResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAlertSceneByEventResponseBody) SetRequestId(v string) *DescribeAlertSceneByEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAlertSceneByEventResponseBody) SetSuccess(v bool) *DescribeAlertSceneByEventResponseBody {
	s.Success = &v
	return s
}

type DescribeAlertSceneByEventResponseBodyData struct {
	// The alert name. The display name of the alert name varies based on the language of the system, such as Chinese and English.
	//
	// example:
	//
	// login_common_ip
	AlertName *string `json:"AlertName,omitempty" xml:"AlertName,omitempty"`
	// The ID of the alert name.
	//
	// example:
	//
	// login_common_ip
	AlertNameId *string `json:"AlertNameId,omitempty" xml:"AlertNameId,omitempty"`
	// The alert title. The display name of the alert title varies based on the language of the system, such as Chinese and English.
	//
	// example:
	//
	// Unusual Logon-login_common_ip
	AlertTile *string `json:"AlertTile,omitempty" xml:"AlertTile,omitempty"`
	// The ID of the alert title.
	//
	// example:
	//
	// Unusual Logon-login_common_ip
	AlertTileId *string `json:"AlertTileId,omitempty" xml:"AlertTileId,omitempty"`
	// The alert type. The display name of the alert type varies based on the language of the system, such as Chinese and English.
	//
	// example:
	//
	// Unusual Logon
	AlertType *string `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
	// The ID of the alert type.
	//
	// example:
	//
	// Unusual Logon
	AlertTypeId *string `json:"AlertTypeId,omitempty" xml:"AlertTypeId,omitempty"`
	// The objects that can be added to the whitelist.
	//
	// example:
	//
	// [{"Type": "host_uuid","Value": "441862da-a539-4cc0-a00d-473955826881","Values": ["441862da-a539-4cc0-a00d-473955826881"],"Name": "${aliyun.siem.entity.host_uuid}"}]
	Targets []*DescribeAlertSceneByEventResponseBodyDataTargets `json:"Targets,omitempty" xml:"Targets,omitempty" type:"Repeated"`
}

func (s DescribeAlertSceneByEventResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertSceneByEventResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAlertSceneByEventResponseBodyData) SetAlertName(v string) *DescribeAlertSceneByEventResponseBodyData {
	s.AlertName = &v
	return s
}

func (s *DescribeAlertSceneByEventResponseBodyData) SetAlertNameId(v string) *DescribeAlertSceneByEventResponseBodyData {
	s.AlertNameId = &v
	return s
}

func (s *DescribeAlertSceneByEventResponseBodyData) SetAlertTile(v string) *DescribeAlertSceneByEventResponseBodyData {
	s.AlertTile = &v
	return s
}

func (s *DescribeAlertSceneByEventResponseBodyData) SetAlertTileId(v string) *DescribeAlertSceneByEventResponseBodyData {
	s.AlertTileId = &v
	return s
}

func (s *DescribeAlertSceneByEventResponseBodyData) SetAlertType(v string) *DescribeAlertSceneByEventResponseBodyData {
	s.AlertType = &v
	return s
}

func (s *DescribeAlertSceneByEventResponseBodyData) SetAlertTypeId(v string) *DescribeAlertSceneByEventResponseBodyData {
	s.AlertTypeId = &v
	return s
}

func (s *DescribeAlertSceneByEventResponseBodyData) SetTargets(v []*DescribeAlertSceneByEventResponseBodyDataTargets) *DescribeAlertSceneByEventResponseBodyData {
	s.Targets = v
	return s
}

type DescribeAlertSceneByEventResponseBodyDataTargets struct {
	// The display name of the entity attribute field that can be added to the whitelist.
	//
	// example:
	//
	// host uuid
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The entity attribute field that can be added to the whitelist.
	//
	// example:
	//
	// host_uuid
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The right operand that is displayed by default in the whitelist rule.
	//
	// example:
	//
	// 441862da-a539-4cc0-a00d-47395582****
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// The supported right operands of the whitelist rule.
	//
	// example:
	//
	// ["441862da-a539-4cc0-a00d-473955826881"]
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s DescribeAlertSceneByEventResponseBodyDataTargets) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertSceneByEventResponseBodyDataTargets) GoString() string {
	return s.String()
}

func (s *DescribeAlertSceneByEventResponseBodyDataTargets) SetName(v string) *DescribeAlertSceneByEventResponseBodyDataTargets {
	s.Name = &v
	return s
}

func (s *DescribeAlertSceneByEventResponseBodyDataTargets) SetType(v string) *DescribeAlertSceneByEventResponseBodyDataTargets {
	s.Type = &v
	return s
}

func (s *DescribeAlertSceneByEventResponseBodyDataTargets) SetValue(v string) *DescribeAlertSceneByEventResponseBodyDataTargets {
	s.Value = &v
	return s
}

func (s *DescribeAlertSceneByEventResponseBodyDataTargets) SetValues(v []*string) *DescribeAlertSceneByEventResponseBodyDataTargets {
	s.Values = v
	return s
}

type DescribeAlertSceneByEventResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAlertSceneByEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAlertSceneByEventResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertSceneByEventResponse) GoString() string {
	return s.String()
}

func (s *DescribeAlertSceneByEventResponse) SetHeaders(v map[string]*string) *DescribeAlertSceneByEventResponse {
	s.Headers = v
	return s
}

func (s *DescribeAlertSceneByEventResponse) SetStatusCode(v int32) *DescribeAlertSceneByEventResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAlertSceneByEventResponse) SetBody(v *DescribeAlertSceneByEventResponseBody) *DescribeAlertSceneByEventResponse {
	s.Body = v
	return s
}

type DescribeAlertSourceRequest struct {
	// The end of the time range to query. Unit: milliseconds.
	//
	// example:
	//
	// 1577808000000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The risk levels. The value is a JSON array. Valid values:
	//
	// 	- serious: high
	//
	// 	- suspicious: medium
	//
	// 	- remind: low
	//
	// example:
	//
	// ["serious","suspicious","remind"]
	Level []*string `json:"Level,omitempty" xml:"Level,omitempty" type:"Repeated"`
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
	// The ID of the account that you switch from the management account.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view. Valid values:
	//
	// - 0: the current Alibaba Cloud account
	//
	// - 1: the global account
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// The beginning of the time range to query. Unit: milliseconds.
	//
	// example:
	//
	// 1577808000000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeAlertSourceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertSourceRequest) GoString() string {
	return s.String()
}

func (s *DescribeAlertSourceRequest) SetEndTime(v int64) *DescribeAlertSourceRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeAlertSourceRequest) SetLevel(v []*string) *DescribeAlertSourceRequest {
	s.Level = v
	return s
}

func (s *DescribeAlertSourceRequest) SetRegionId(v string) *DescribeAlertSourceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAlertSourceRequest) SetRoleFor(v int64) *DescribeAlertSourceRequest {
	s.RoleFor = &v
	return s
}

func (s *DescribeAlertSourceRequest) SetRoleType(v int32) *DescribeAlertSourceRequest {
	s.RoleType = &v
	return s
}

func (s *DescribeAlertSourceRequest) SetStartTime(v int64) *DescribeAlertSourceRequest {
	s.StartTime = &v
	return s
}

type DescribeAlertSourceResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	//
	// example:
	//
	// 123456
	Data []*DescribeAlertSourceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeAlertSourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertSourceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAlertSourceResponseBody) SetCode(v int32) *DescribeAlertSourceResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAlertSourceResponseBody) SetData(v []*DescribeAlertSourceResponseBodyData) *DescribeAlertSourceResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAlertSourceResponseBody) SetMessage(v string) *DescribeAlertSourceResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAlertSourceResponseBody) SetRequestId(v string) *DescribeAlertSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAlertSourceResponseBody) SetSuccess(v bool) *DescribeAlertSourceResponseBody {
	s.Success = &v
	return s
}

type DescribeAlertSourceResponseBodyData struct {
	// The internal code of the alert data source.
	//
	// example:
	//
	// aliyun.siem.alert_datasource.sas
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The name of the alert data source.
	//
	// example:
	//
	// sas
	SourceName *string `json:"SourceName,omitempty" xml:"SourceName,omitempty"`
}

func (s DescribeAlertSourceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertSourceResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAlertSourceResponseBodyData) SetSource(v string) *DescribeAlertSourceResponseBodyData {
	s.Source = &v
	return s
}

func (s *DescribeAlertSourceResponseBodyData) SetSourceName(v string) *DescribeAlertSourceResponseBodyData {
	s.SourceName = &v
	return s
}

type DescribeAlertSourceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAlertSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAlertSourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertSourceResponse) GoString() string {
	return s.String()
}

func (s *DescribeAlertSourceResponse) SetHeaders(v map[string]*string) *DescribeAlertSourceResponse {
	s.Headers = v
	return s
}

func (s *DescribeAlertSourceResponse) SetStatusCode(v int32) *DescribeAlertSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAlertSourceResponse) SetBody(v *DescribeAlertSourceResponseBody) *DescribeAlertSourceResponse {
	s.Body = v
	return s
}

type DescribeAlertSourceWithEventRequest struct {
	// The UUID of the event.
	//
	// example:
	//
	// 85ea4241-798f-4684-a876-65d4f0c3****
	IncidentUuid *string `json:"IncidentUuid,omitempty" xml:"IncidentUuid,omitempty"`
	// The region in which the data management center of the threat analysis feature resides. Specify this parameter based on the region where your assets reside. Valid values:
	//
	// 	- Valid values: Your assets reside in regions in China.
	//
	// 	- ap-southeast-1: Your assets reside in regions outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the account that you switch from the management account.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view. Valid values:
	//
	// - 0: the current Alibaba Cloud account
	//
	// - 1: the global account
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s DescribeAlertSourceWithEventRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertSourceWithEventRequest) GoString() string {
	return s.String()
}

func (s *DescribeAlertSourceWithEventRequest) SetIncidentUuid(v string) *DescribeAlertSourceWithEventRequest {
	s.IncidentUuid = &v
	return s
}

func (s *DescribeAlertSourceWithEventRequest) SetRegionId(v string) *DescribeAlertSourceWithEventRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAlertSourceWithEventRequest) SetRoleFor(v int64) *DescribeAlertSourceWithEventRequest {
	s.RoleFor = &v
	return s
}

func (s *DescribeAlertSourceWithEventRequest) SetRoleType(v int32) *DescribeAlertSourceWithEventRequest {
	s.RoleType = &v
	return s
}

type DescribeAlertSourceWithEventResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	//
	// example:
	//
	// 123456
	Data []*DescribeAlertSourceWithEventResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeAlertSourceWithEventResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertSourceWithEventResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAlertSourceWithEventResponseBody) SetCode(v int32) *DescribeAlertSourceWithEventResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAlertSourceWithEventResponseBody) SetData(v []*DescribeAlertSourceWithEventResponseBodyData) *DescribeAlertSourceWithEventResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAlertSourceWithEventResponseBody) SetMessage(v string) *DescribeAlertSourceWithEventResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAlertSourceWithEventResponseBody) SetRequestId(v string) *DescribeAlertSourceWithEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAlertSourceWithEventResponseBody) SetSuccess(v bool) *DescribeAlertSourceWithEventResponseBody {
	s.Success = &v
	return s
}

type DescribeAlertSourceWithEventResponseBodyData struct {
	// The internal code of the alert data source.
	//
	// example:
	//
	// aliyun.siem.alert_datasource.sas
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The name of the alert data source.
	//
	// example:
	//
	// sas
	SourceName *string `json:"SourceName,omitempty" xml:"SourceName,omitempty"`
}

func (s DescribeAlertSourceWithEventResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertSourceWithEventResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAlertSourceWithEventResponseBodyData) SetSource(v string) *DescribeAlertSourceWithEventResponseBodyData {
	s.Source = &v
	return s
}

func (s *DescribeAlertSourceWithEventResponseBodyData) SetSourceName(v string) *DescribeAlertSourceWithEventResponseBodyData {
	s.SourceName = &v
	return s
}

type DescribeAlertSourceWithEventResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAlertSourceWithEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAlertSourceWithEventResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertSourceWithEventResponse) GoString() string {
	return s.String()
}

func (s *DescribeAlertSourceWithEventResponse) SetHeaders(v map[string]*string) *DescribeAlertSourceWithEventResponse {
	s.Headers = v
	return s
}

func (s *DescribeAlertSourceWithEventResponse) SetStatusCode(v int32) *DescribeAlertSourceWithEventResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAlertSourceWithEventResponse) SetBody(v *DescribeAlertSourceWithEventResponseBody) *DescribeAlertSourceWithEventResponse {
	s.Body = v
	return s
}

type DescribeAlertTypeRequest struct {
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
	// The ID of the account that you switch from the management account.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view. Valid values:
	//
	// - 0: the current Alibaba Cloud account
	//
	// - 1: the global account
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// The type of rule. Valid values:
	//
	// - predefine: the defined rule by system
	//
	// - customize: the customed rule by user
	//
	// example:
	//
	// customize
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
}

func (s DescribeAlertTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertTypeRequest) GoString() string {
	return s.String()
}

func (s *DescribeAlertTypeRequest) SetRegionId(v string) *DescribeAlertTypeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAlertTypeRequest) SetRoleFor(v int64) *DescribeAlertTypeRequest {
	s.RoleFor = &v
	return s
}

func (s *DescribeAlertTypeRequest) SetRoleType(v int32) *DescribeAlertTypeRequest {
	s.RoleType = &v
	return s
}

func (s *DescribeAlertTypeRequest) SetRuleType(v string) *DescribeAlertTypeRequest {
	s.RuleType = &v
	return s
}

type DescribeAlertTypeResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	//
	// example:
	//
	// 123456
	Data []*DescribeAlertTypeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeAlertTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertTypeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAlertTypeResponseBody) SetCode(v int32) *DescribeAlertTypeResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAlertTypeResponseBody) SetData(v []*DescribeAlertTypeResponseBodyData) *DescribeAlertTypeResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAlertTypeResponseBody) SetMessage(v string) *DescribeAlertTypeResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAlertTypeResponseBody) SetRequestId(v string) *DescribeAlertTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAlertTypeResponseBody) SetSuccess(v bool) *DescribeAlertTypeResponseBody {
	s.Success = &v
	return s
}

type DescribeAlertTypeResponseBodyData struct {
	// The type of the risk.
	//
	// example:
	//
	// WEBSHELL
	AlertType *string `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
	// The internal code of the risk type.
	//
	// example:
	//
	// siem_rule_type_process_abnormal_command
	AlertTypeMds *string `json:"AlertTypeMds,omitempty" xml:"AlertTypeMds,omitempty"`
}

func (s DescribeAlertTypeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertTypeResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAlertTypeResponseBodyData) SetAlertType(v string) *DescribeAlertTypeResponseBodyData {
	s.AlertType = &v
	return s
}

func (s *DescribeAlertTypeResponseBodyData) SetAlertTypeMds(v string) *DescribeAlertTypeResponseBodyData {
	s.AlertTypeMds = &v
	return s
}

type DescribeAlertTypeResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAlertTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAlertTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertTypeResponse) GoString() string {
	return s.String()
}

func (s *DescribeAlertTypeResponse) SetHeaders(v map[string]*string) *DescribeAlertTypeResponse {
	s.Headers = v
	return s
}

func (s *DescribeAlertTypeResponse) SetStatusCode(v int32) *DescribeAlertTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAlertTypeResponse) SetBody(v *DescribeAlertTypeResponseBody) *DescribeAlertTypeResponse {
	s.Body = v
	return s
}

type DescribeAlertsRequest struct {
	// example:
	//
	// 异常连接-TFTP恶意扫描
	AlertName *string `json:"AlertName,omitempty" xml:"AlertName,omitempty"`
	// The title of the alert.
	//
	// example:
	//
	// Unusual Logon-login_common_account
	AlertTitle *string `json:"AlertTitle,omitempty" xml:"AlertTitle,omitempty"`
	// example:
	//
	// 异常连接-TFTP恶意扫描
	AlertType *string `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
	// The UUID of the alert.
	//
	// example:
	//
	// sas_71e24437d2797ce8fc59692905a4****
	AlertUuid *string `json:"AlertUuid,omitempty" xml:"AlertUuid,omitempty"`
	// example:
	//
	// 异常连接-TFTP恶意扫描
	AssetId *string `json:"AssetId,omitempty" xml:"AssetId,omitempty"`
	// example:
	//
	// 异常连接-TFTP恶意扫描
	AssetName *string `json:"AssetName,omitempty" xml:"AssetName,omitempty"`
	// The page number. Pages start from page 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The end of the time range to query. Unit: milliseconds.
	//
	// example:
	//
	// 1577808000000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 异常连接-TFTP恶意扫描
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// example:
	//
	// 异常连接-TFTP恶意扫描
	EntityName *string `json:"EntityName,omitempty" xml:"EntityName,omitempty"`
	// Specifies whether an attack is defended. Valid values:
	//
	// 	- 0: detected.
	//
	// 	- 1: blocked.
	//
	// example:
	//
	// 1
	IsDefend *string `json:"IsDefend,omitempty" xml:"IsDefend,omitempty"`
	// example:
	//
	// 176555323***
	LabelType *string `json:"LabelType,omitempty" xml:"LabelType,omitempty"`
	// The risk level. The value is a JSON array. Valid values:
	//
	// 	- serious: high
	//
	// 	- suspicious: medium
	//
	// 	- remind: low
	//
	// example:
	//
	// ["serious","suspicious","remind"]
	Level []*string `json:"Level,omitempty" xml:"Level,omitempty" type:"Repeated"`
	// The number of entries per page. Maximum value: 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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
	// The ID of the account that you switch from the management account.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view. Valid values:
	//
	// - 0: the current Alibaba Cloud account
	//
	// - 1: the global account
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// The source of the alert.
	//
	// example:
	//
	// sas
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The beginning of the time range to query. Unit: milliseconds.
	//
	// example:
	//
	// 1577808000000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The ID of the Alibaba Cloud account within which the alert is generated.
	//
	// example:
	//
	// 176555323***
	SubUserId *string `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
}

func (s DescribeAlertsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAlertsRequest) SetAlertName(v string) *DescribeAlertsRequest {
	s.AlertName = &v
	return s
}

func (s *DescribeAlertsRequest) SetAlertTitle(v string) *DescribeAlertsRequest {
	s.AlertTitle = &v
	return s
}

func (s *DescribeAlertsRequest) SetAlertType(v string) *DescribeAlertsRequest {
	s.AlertType = &v
	return s
}

func (s *DescribeAlertsRequest) SetAlertUuid(v string) *DescribeAlertsRequest {
	s.AlertUuid = &v
	return s
}

func (s *DescribeAlertsRequest) SetAssetId(v string) *DescribeAlertsRequest {
	s.AssetId = &v
	return s
}

func (s *DescribeAlertsRequest) SetAssetName(v string) *DescribeAlertsRequest {
	s.AssetName = &v
	return s
}

func (s *DescribeAlertsRequest) SetCurrentPage(v int32) *DescribeAlertsRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeAlertsRequest) SetEndTime(v int64) *DescribeAlertsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeAlertsRequest) SetEntityId(v string) *DescribeAlertsRequest {
	s.EntityId = &v
	return s
}

func (s *DescribeAlertsRequest) SetEntityName(v string) *DescribeAlertsRequest {
	s.EntityName = &v
	return s
}

func (s *DescribeAlertsRequest) SetIsDefend(v string) *DescribeAlertsRequest {
	s.IsDefend = &v
	return s
}

func (s *DescribeAlertsRequest) SetLabelType(v string) *DescribeAlertsRequest {
	s.LabelType = &v
	return s
}

func (s *DescribeAlertsRequest) SetLevel(v []*string) *DescribeAlertsRequest {
	s.Level = v
	return s
}

func (s *DescribeAlertsRequest) SetPageSize(v int32) *DescribeAlertsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAlertsRequest) SetRegionId(v string) *DescribeAlertsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAlertsRequest) SetRoleFor(v int64) *DescribeAlertsRequest {
	s.RoleFor = &v
	return s
}

func (s *DescribeAlertsRequest) SetRoleType(v int32) *DescribeAlertsRequest {
	s.RoleType = &v
	return s
}

func (s *DescribeAlertsRequest) SetSource(v string) *DescribeAlertsRequest {
	s.Source = &v
	return s
}

func (s *DescribeAlertsRequest) SetStartTime(v int64) *DescribeAlertsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeAlertsRequest) SetSubUserId(v string) *DescribeAlertsRequest {
	s.SubUserId = &v
	return s
}

type DescribeAlertsResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	//
	// example:
	//
	// 123456
	Data *DescribeAlertsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeAlertsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAlertsResponseBody) SetCode(v int32) *DescribeAlertsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAlertsResponseBody) SetData(v *DescribeAlertsResponseBodyData) *DescribeAlertsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAlertsResponseBody) SetMessage(v string) *DescribeAlertsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAlertsResponseBody) SetRequestId(v string) *DescribeAlertsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAlertsResponseBody) SetSuccess(v bool) *DescribeAlertsResponseBody {
	s.Success = &v
	return s
}

type DescribeAlertsResponseBodyData struct {
	// The pagination information.
	PageInfo *DescribeAlertsResponseBodyDataPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The detailed data.
	ResponseData []*DescribeAlertsResponseBodyDataResponseData `json:"ResponseData,omitempty" xml:"ResponseData,omitempty" type:"Repeated"`
}

func (s DescribeAlertsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAlertsResponseBodyData) SetPageInfo(v *DescribeAlertsResponseBodyDataPageInfo) *DescribeAlertsResponseBodyData {
	s.PageInfo = v
	return s
}

func (s *DescribeAlertsResponseBodyData) SetResponseData(v []*DescribeAlertsResponseBodyDataResponseData) *DescribeAlertsResponseBodyData {
	s.ResponseData = v
	return s
}

type DescribeAlertsResponseBodyDataPageInfo struct {
	// The current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAlertsResponseBodyDataPageInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertsResponseBodyDataPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeAlertsResponseBodyDataPageInfo) SetCurrentPage(v int32) *DescribeAlertsResponseBodyDataPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeAlertsResponseBodyDataPageInfo) SetPageSize(v int32) *DescribeAlertsResponseBodyDataPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeAlertsResponseBodyDataPageInfo) SetTotalCount(v int64) *DescribeAlertsResponseBodyDataPageInfo {
	s.TotalCount = &v
	return s
}

type DescribeAlertsResponseBodyDataResponseData struct {
	// The description of the alert.
	//
	// example:
	//
	// The detection model found a suspicious Webshell file on your server, which may be a backdoor file implanted to maintain permissions after the attacker successfully invaded the website.
	AlertDesc *string `json:"AlertDesc,omitempty" xml:"AlertDesc,omitempty"`
	// The internal code of the alert description.
	//
	// example:
	//
	// security_event_config.event_name.webshell
	AlertDescCode *string `json:"AlertDescCode,omitempty" xml:"AlertDescCode,omitempty"`
	// The description of the alert in English.
	//
	// example:
	//
	// The detection model found a suspicious Webshell file on your server, which may be a backdoor file implanted to maintain permissions after the attacker successfully invaded the website.
	AlertDescEn *string `json:"AlertDescEn,omitempty" xml:"AlertDescEn,omitempty"`
	// The details of the alert.
	//
	// example:
	//
	// {"main_user_id": "165295629792****";"log_uuid_count": "99";"attack_ip": "21.92.*.*"}
	AlertDetail *string `json:"AlertDetail,omitempty" xml:"AlertDetail,omitempty"`
	// The displayed details of the alert.
	//
	// example:
	//
	// aliyun
	AlertInfoList []*DescribeAlertsResponseBodyDataResponseDataAlertInfoList `json:"AlertInfoList,omitempty" xml:"AlertInfoList,omitempty" type:"Repeated"`
	// The threat level. Valid values:
	//
	// 	- serious: high
	//
	// 	- suspicious: medium
	//
	// 	- remind: low
	//
	// example:
	//
	// remind
	AlertLevel *string `json:"AlertLevel,omitempty" xml:"AlertLevel,omitempty"`
	// The name of the alert.
	//
	// example:
	//
	// Try SNMP weak password
	AlertName *string `json:"AlertName,omitempty" xml:"AlertName,omitempty"`
	// The internal code of the alert name.
	//
	// example:
	//
	// security_event_config.event_name.webshell
	AlertNameCode *string `json:"AlertNameCode,omitempty" xml:"AlertNameCode,omitempty"`
	// The name of the alert in English.
	//
	// example:
	//
	// Try SNMP weak password
	AlertNameEn *string `json:"AlertNameEn,omitempty" xml:"AlertNameEn,omitempty"`
	// The service for which the alert associated with the event is generated.
	//
	// example:
	//
	// sas
	AlertSrcProd *string `json:"AlertSrcProd,omitempty" xml:"AlertSrcProd,omitempty"`
	// The sub-module of ther alert source.
	//
	// example:
	//
	// waf
	AlertSrcProdModule *string `json:"AlertSrcProdModule,omitempty" xml:"AlertSrcProdModule,omitempty"`
	// The title of the alert.
	//
	// example:
	//
	// Scan-Try SNMP weak password
	AlertTitle *string `json:"AlertTitle,omitempty" xml:"AlertTitle,omitempty"`
	// The title of the alert in English.
	//
	// example:
	//
	// Scan-Try SNMP weak password
	AlertTitleEn *string `json:"AlertTitleEn,omitempty" xml:"AlertTitleEn,omitempty"`
	// The alert type.
	//
	// example:
	//
	// Scan
	AlertType *string `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
	// The internal code of the alert type.
	//
	// example:
	//
	// security_event_config.event_name.webshellName
	AlertTypeCode *string `json:"AlertTypeCode,omitempty" xml:"AlertTypeCode,omitempty"`
	// The type of the alert in English.
	//
	// example:
	//
	// Scan
	AlertTypeEn *string `json:"AlertTypeEn,omitempty" xml:"AlertTypeEn,omitempty"`
	// The UUID of the alert.
	//
	// example:
	//
	// sas_71e24437d2797ce8fc59692905a4****
	AlertUuid *string `json:"AlertUuid,omitempty" xml:"AlertUuid,omitempty"`
	// The details of the asset.
	//
	// example:
	//
	// [
	//
	//       {
	//
	//             "is_main_asset": "1",
	//
	//             "asset_name": "47.245.*",
	//
	//             "port": "22",
	//
	//             "ip": "47.245.*",
	//
	//             "asset_type": "ip",
	//
	//             "location": "ap-southeast-1",
	//
	//             "asset_id": "47.245.*",
	//
	//             "net_connect_dir": "in"
	//
	//       }
	//
	// ]
	AssetList *string `json:"AssetList,omitempty" xml:"AssetList,omitempty"`
	// The tag of the ATT\\&CK attack.
	//
	// example:
	//
	// T1595.002 Vulnerability Scanning
	AttCk *string `json:"AttCk,omitempty" xml:"AttCk,omitempty"`
	// The cloud code. Valid values:
	//
	// 	- aliyun: Alibaba Cloud
	//
	// 	- qcloud: Tencent Cloud
	//
	// 	- hcloud: Huawei Cloud
	//
	// example:
	//
	// aliyun
	CloudCode *string `json:"CloudCode,omitempty" xml:"CloudCode,omitempty"`
	// The time when the alert was closed.
	//
	// example:
	//
	// 2021-01-06 16:37:29
	EndTime       *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	EntityList    *string `json:"EntityList,omitempty" xml:"EntityList,omitempty"`
	ExtendContent *string `json:"ExtendContent,omitempty" xml:"ExtendContent,omitempty"`
	// The time when the alert was received.
	//
	// example:
	//
	// 2021-01-06 16:37:29
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the alert was last updated.
	//
	// example:
	//
	// 2021-01-06 16:37:29
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The unique ID of the alert.
	//
	// example:
	//
	// 123456789
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The UUID of the event.
	//
	// example:
	//
	// 85ea4241-798f-4684-a876-65d4f0c3****
	IncidentUuid *string `json:"IncidentUuid,omitempty" xml:"IncidentUuid,omitempty"`
	// Indicates whether an attack is defended. Valid values:
	//
	// 	- 0: detected.
	//
	// 	- 1: blocked.
	//
	// example:
	//
	// 1
	IsDefend *string `json:"IsDefend,omitempty" xml:"IsDefend,omitempty"`
	// The time when the alert was recorded.
	//
	// example:
	//
	// 2021-01-06 16:37:29
	LogTime *string `json:"LogTime,omitempty" xml:"LogTime,omitempty"`
	// The UUID of the alert log.
	//
	// example:
	//
	// cfw_d12e285a-a042-4d7e-be89-f8a795ef****
	LogUuid *string `json:"LogUuid,omitempty" xml:"LogUuid,omitempty"`
	// The ID of the Alibaba Cloud account that is associated with the alert in SIEM.
	//
	// example:
	//
	// 127608589417****
	MainUserId *int64 `json:"MainUserId,omitempty" xml:"MainUserId,omitempty"`
	// The time when the alert is triggered.
	//
	// example:
	//
	// 2021-01-06 16:37:29
	OccurTime *string `json:"OccurTime,omitempty" xml:"OccurTime,omitempty"`
	// The time at which the alert was first generated.
	//
	// example:
	//
	// 2021-01-06 16:37:29
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The ID of the Alibaba Cloud account within which the alert is generated.
	//
	// example:
	//
	// 176555323***
	SubUserId *int64 `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
	// example:
	//
	// 176555323***
	SubUserName *string `json:"SubUserName,omitempty" xml:"SubUserName,omitempty"`
}

func (s DescribeAlertsResponseBodyDataResponseData) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertsResponseBodyDataResponseData) GoString() string {
	return s.String()
}

func (s *DescribeAlertsResponseBodyDataResponseData) SetAlertDesc(v string) *DescribeAlertsResponseBodyDataResponseData {
	s.AlertDesc = &v
	return s
}

func (s *DescribeAlertsResponseBodyDataResponseData) SetAlertDescCode(v string) *DescribeAlertsResponseBodyDataResponseData {
	s.AlertDescCode = &v
	return s
}

func (s *DescribeAlertsResponseBodyDataResponseData) SetAlertDescEn(v string) *DescribeAlertsResponseBodyDataResponseData {
	s.AlertDescEn = &v
	return s
}

func (s *DescribeAlertsResponseBodyDataResponseData) SetAlertDetail(v string) *DescribeAlertsResponseBodyDataResponseData {
	s.AlertDetail = &v
	return s
}

func (s *DescribeAlertsResponseBodyDataResponseData) SetAlertInfoList(v []*DescribeAlertsResponseBodyDataResponseDataAlertInfoList) *DescribeAlertsResponseBodyDataResponseData {
	s.AlertInfoList = v
	return s
}

func (s *DescribeAlertsResponseBodyDataResponseData) SetAlertLevel(v string) *DescribeAlertsResponseBodyDataResponseData {
	s.AlertLevel = &v
	return s
}

func (s *DescribeAlertsResponseBodyDataResponseData) SetAlertName(v string) *DescribeAlertsResponseBodyDataResponseData {
	s.AlertName = &v
	return s
}

func (s *DescribeAlertsResponseBodyDataResponseData) SetAlertNameCode(v string) *DescribeAlertsResponseBodyDataResponseData {
	s.AlertNameCode = &v
	return s
}

func (s *DescribeAlertsResponseBodyDataResponseData) SetAlertNameEn(v string) *DescribeAlertsResponseBodyDataResponseData {
	s.AlertNameEn = &v
	return s
}

func (s *DescribeAlertsResponseBodyDataResponseData) SetAlertSrcProd(v string) *DescribeAlertsResponseBodyDataResponseData {
	s.AlertSrcProd = &v
	return s
}

func (s *DescribeAlertsResponseBodyDataResponseData) SetAlertSrcProdModule(v string) *DescribeAlertsResponseBodyDataResponseData {
	s.AlertSrcProdModule = &v
	return s
}

func (s *DescribeAlertsResponseBodyDataResponseData) SetAlertTitle(v string) *DescribeAlertsResponseBodyDataResponseData {
	s.AlertTitle = &v
	return s
}

func (s *DescribeAlertsResponseBodyDataResponseData) SetAlertTitleEn(v string) *DescribeAlertsResponseBodyDataResponseData {
	s.AlertTitleEn = &v
	return s
}

func (s *DescribeAlertsResponseBodyDataResponseData) SetAlertType(v string) *DescribeAlertsResponseBodyDataResponseData {
	s.AlertType = &v
	return s
}

func (s *DescribeAlertsResponseBodyDataResponseData) SetAlertTypeCode(v string) *DescribeAlertsResponseBodyDataResponseData {
	s.AlertTypeCode = &v
	return s
}

func (s *DescribeAlertsResponseBodyDataResponseData) SetAlertTypeEn(v string) *DescribeAlertsResponseBodyDataResponseData {
	s.AlertTypeEn = &v
	return s
}

func (s *DescribeAlertsResponseBodyDataResponseData) SetAlertUuid(v string) *DescribeAlertsResponseBodyDataResponseData {
	s.AlertUuid = &v
	return s
}

func (s *DescribeAlertsResponseBodyDataResponseData) SetAssetList(v string) *DescribeAlertsResponseBodyDataResponseData {
	s.AssetList = &v
	return s
}

func (s *DescribeAlertsResponseBodyDataResponseData) SetAttCk(v string) *DescribeAlertsResponseBodyDataResponseData {
	s.AttCk = &v
	return s
}

func (s *DescribeAlertsResponseBodyDataResponseData) SetCloudCode(v string) *DescribeAlertsResponseBodyDataResponseData {
	s.CloudCode = &v
	return s
}

func (s *DescribeAlertsResponseBodyDataResponseData) SetEndTime(v string) *DescribeAlertsResponseBodyDataResponseData {
	s.EndTime = &v
	return s
}

func (s *DescribeAlertsResponseBodyDataResponseData) SetEntityList(v string) *DescribeAlertsResponseBodyDataResponseData {
	s.EntityList = &v
	return s
}

func (s *DescribeAlertsResponseBodyDataResponseData) SetExtendContent(v string) *DescribeAlertsResponseBodyDataResponseData {
	s.ExtendContent = &v
	return s
}

func (s *DescribeAlertsResponseBodyDataResponseData) SetGmtCreate(v string) *DescribeAlertsResponseBodyDataResponseData {
	s.GmtCreate = &v
	return s
}

func (s *DescribeAlertsResponseBodyDataResponseData) SetGmtModified(v string) *DescribeAlertsResponseBodyDataResponseData {
	s.GmtModified = &v
	return s
}

func (s *DescribeAlertsResponseBodyDataResponseData) SetId(v int64) *DescribeAlertsResponseBodyDataResponseData {
	s.Id = &v
	return s
}

func (s *DescribeAlertsResponseBodyDataResponseData) SetIncidentUuid(v string) *DescribeAlertsResponseBodyDataResponseData {
	s.IncidentUuid = &v
	return s
}

func (s *DescribeAlertsResponseBodyDataResponseData) SetIsDefend(v string) *DescribeAlertsResponseBodyDataResponseData {
	s.IsDefend = &v
	return s
}

func (s *DescribeAlertsResponseBodyDataResponseData) SetLogTime(v string) *DescribeAlertsResponseBodyDataResponseData {
	s.LogTime = &v
	return s
}

func (s *DescribeAlertsResponseBodyDataResponseData) SetLogUuid(v string) *DescribeAlertsResponseBodyDataResponseData {
	s.LogUuid = &v
	return s
}

func (s *DescribeAlertsResponseBodyDataResponseData) SetMainUserId(v int64) *DescribeAlertsResponseBodyDataResponseData {
	s.MainUserId = &v
	return s
}

func (s *DescribeAlertsResponseBodyDataResponseData) SetOccurTime(v string) *DescribeAlertsResponseBodyDataResponseData {
	s.OccurTime = &v
	return s
}

func (s *DescribeAlertsResponseBodyDataResponseData) SetStartTime(v string) *DescribeAlertsResponseBodyDataResponseData {
	s.StartTime = &v
	return s
}

func (s *DescribeAlertsResponseBodyDataResponseData) SetSubUserId(v int64) *DescribeAlertsResponseBodyDataResponseData {
	s.SubUserId = &v
	return s
}

func (s *DescribeAlertsResponseBodyDataResponseData) SetSubUserName(v string) *DescribeAlertsResponseBodyDataResponseData {
	s.SubUserName = &v
	return s
}

type DescribeAlertsResponseBodyDataResponseDataAlertInfoList struct {
	// The attribute key.
	//
	// example:
	//
	// suspicious.wbd.wb.trojanpath
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The name of the key.
	//
	// example:
	//
	// Trojan Path
	KeyName *string `json:"KeyName,omitempty" xml:"KeyName,omitempty"`
	// The value of the key.
	//
	// example:
	//
	// /root/test33.php
	Values *string `json:"Values,omitempty" xml:"Values,omitempty"`
}

func (s DescribeAlertsResponseBodyDataResponseDataAlertInfoList) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertsResponseBodyDataResponseDataAlertInfoList) GoString() string {
	return s.String()
}

func (s *DescribeAlertsResponseBodyDataResponseDataAlertInfoList) SetKey(v string) *DescribeAlertsResponseBodyDataResponseDataAlertInfoList {
	s.Key = &v
	return s
}

func (s *DescribeAlertsResponseBodyDataResponseDataAlertInfoList) SetKeyName(v string) *DescribeAlertsResponseBodyDataResponseDataAlertInfoList {
	s.KeyName = &v
	return s
}

func (s *DescribeAlertsResponseBodyDataResponseDataAlertInfoList) SetValues(v string) *DescribeAlertsResponseBodyDataResponseDataAlertInfoList {
	s.Values = &v
	return s
}

type DescribeAlertsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAlertsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAlertsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAlertsResponse) SetHeaders(v map[string]*string) *DescribeAlertsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAlertsResponse) SetStatusCode(v int32) *DescribeAlertsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAlertsResponse) SetBody(v *DescribeAlertsResponseBody) *DescribeAlertsResponse {
	s.Body = v
	return s
}

type DescribeAlertsCountRequest struct {
	// The end of the time range to query. Unit: milliseconds.
	//
	// example:
	//
	// 1577808000000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// bySrcProd
	QueryType *string `json:"QueryType,omitempty" xml:"QueryType,omitempty"`
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
	// The ID of the account that you switch from the management account.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view. Valid values:
	//
	// - 0: the current Alibaba Cloud account
	//
	// - 1: the global account
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// The beginning of the time range to query. Unit: milliseconds.
	//
	// example:
	//
	// 1577808000000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeAlertsCountRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertsCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeAlertsCountRequest) SetEndTime(v int64) *DescribeAlertsCountRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeAlertsCountRequest) SetQueryType(v string) *DescribeAlertsCountRequest {
	s.QueryType = &v
	return s
}

func (s *DescribeAlertsCountRequest) SetRegionId(v string) *DescribeAlertsCountRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAlertsCountRequest) SetRoleFor(v int64) *DescribeAlertsCountRequest {
	s.RoleFor = &v
	return s
}

func (s *DescribeAlertsCountRequest) SetRoleType(v int32) *DescribeAlertsCountRequest {
	s.RoleType = &v
	return s
}

func (s *DescribeAlertsCountRequest) SetStartTime(v int64) *DescribeAlertsCountRequest {
	s.StartTime = &v
	return s
}

type DescribeAlertsCountResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	//
	// example:
	//
	// 123456
	Data *DescribeAlertsCountResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeAlertsCountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertsCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAlertsCountResponseBody) SetCode(v int32) *DescribeAlertsCountResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAlertsCountResponseBody) SetData(v *DescribeAlertsCountResponseBodyData) *DescribeAlertsCountResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAlertsCountResponseBody) SetMessage(v string) *DescribeAlertsCountResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAlertsCountResponseBody) SetRequestId(v string) *DescribeAlertsCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAlertsCountResponseBody) SetSuccess(v bool) *DescribeAlertsCountResponseBody {
	s.Success = &v
	return s
}

type DescribeAlertsCountResponseBodyData struct {
	// The total number of alerts.
	//
	// example:
	//
	// 75
	All      *int64            `json:"All,omitempty" xml:"All,omitempty"`
	CountMap map[string]*int64 `json:"CountMap,omitempty" xml:"CountMap,omitempty"`
	// The number of high-risk alerts.
	//
	// example:
	//
	// 25
	High *int64 `json:"High,omitempty" xml:"High,omitempty"`
	// The number of low-risk alerts.
	//
	// example:
	//
	// 25
	Low *int64 `json:"Low,omitempty" xml:"Low,omitempty"`
	// The number of medium-risk alerts.
	//
	// example:
	//
	// 25
	Medium *int64 `json:"Medium,omitempty" xml:"Medium,omitempty"`
	// The number of connected services.
	//
	// example:
	//
	// 3
	ProductNum *int32 `json:"ProductNum,omitempty" xml:"ProductNum,omitempty"`
}

func (s DescribeAlertsCountResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertsCountResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAlertsCountResponseBodyData) SetAll(v int64) *DescribeAlertsCountResponseBodyData {
	s.All = &v
	return s
}

func (s *DescribeAlertsCountResponseBodyData) SetCountMap(v map[string]*int64) *DescribeAlertsCountResponseBodyData {
	s.CountMap = v
	return s
}

func (s *DescribeAlertsCountResponseBodyData) SetHigh(v int64) *DescribeAlertsCountResponseBodyData {
	s.High = &v
	return s
}

func (s *DescribeAlertsCountResponseBodyData) SetLow(v int64) *DescribeAlertsCountResponseBodyData {
	s.Low = &v
	return s
}

func (s *DescribeAlertsCountResponseBodyData) SetMedium(v int64) *DescribeAlertsCountResponseBodyData {
	s.Medium = &v
	return s
}

func (s *DescribeAlertsCountResponseBodyData) SetProductNum(v int32) *DescribeAlertsCountResponseBodyData {
	s.ProductNum = &v
	return s
}

type DescribeAlertsCountResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAlertsCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAlertsCountResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertsCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeAlertsCountResponse) SetHeaders(v map[string]*string) *DescribeAlertsCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeAlertsCountResponse) SetStatusCode(v int32) *DescribeAlertsCountResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAlertsCountResponse) SetBody(v *DescribeAlertsCountResponseBody) *DescribeAlertsCountResponse {
	s.Body = v
	return s
}

type DescribeAlertsWithEntityRequest struct {
	// The page number. Pages start from page 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 1577808000000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the entity.
	//
	// example:
	//
	// 123456789
	EntityId *int64 `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// example:
	//
	// 123456789
	EntityUuid *string `json:"EntityUuid,omitempty" xml:"EntityUuid,omitempty"`
	// The UUID of the event.
	//
	// example:
	//
	// 85ea4241-798f-4684-a876-65d4f0c3****
	IncidentUuid *string `json:"IncidentUuid,omitempty" xml:"IncidentUuid,omitempty"`
	// The number of entries per page. Maximum value: 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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
	// The ID of the account that you switch from the management account.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view. Valid values:
	//
	// - 0: the current Alibaba Cloud account
	//
	// - 1: the global account
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// The ID of the SOAR handing policy.
	//
	// example:
	//
	// 577bbf90-a770-44a7-8154-586aa2d318fa
	SophonTaskId *string `json:"SophonTaskId,omitempty" xml:"SophonTaskId,omitempty"`
	// example:
	//
	// 1577808000000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeAlertsWithEntityRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertsWithEntityRequest) GoString() string {
	return s.String()
}

func (s *DescribeAlertsWithEntityRequest) SetCurrentPage(v int32) *DescribeAlertsWithEntityRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeAlertsWithEntityRequest) SetEndTime(v int64) *DescribeAlertsWithEntityRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeAlertsWithEntityRequest) SetEntityId(v int64) *DescribeAlertsWithEntityRequest {
	s.EntityId = &v
	return s
}

func (s *DescribeAlertsWithEntityRequest) SetEntityUuid(v string) *DescribeAlertsWithEntityRequest {
	s.EntityUuid = &v
	return s
}

func (s *DescribeAlertsWithEntityRequest) SetIncidentUuid(v string) *DescribeAlertsWithEntityRequest {
	s.IncidentUuid = &v
	return s
}

func (s *DescribeAlertsWithEntityRequest) SetPageSize(v int32) *DescribeAlertsWithEntityRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAlertsWithEntityRequest) SetRegionId(v string) *DescribeAlertsWithEntityRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAlertsWithEntityRequest) SetRoleFor(v int64) *DescribeAlertsWithEntityRequest {
	s.RoleFor = &v
	return s
}

func (s *DescribeAlertsWithEntityRequest) SetRoleType(v int32) *DescribeAlertsWithEntityRequest {
	s.RoleType = &v
	return s
}

func (s *DescribeAlertsWithEntityRequest) SetSophonTaskId(v string) *DescribeAlertsWithEntityRequest {
	s.SophonTaskId = &v
	return s
}

func (s *DescribeAlertsWithEntityRequest) SetStartTime(v int64) *DescribeAlertsWithEntityRequest {
	s.StartTime = &v
	return s
}

type DescribeAlertsWithEntityResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	//
	// example:
	//
	// 123456
	Data *DescribeAlertsWithEntityResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeAlertsWithEntityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertsWithEntityResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAlertsWithEntityResponseBody) SetCode(v int32) *DescribeAlertsWithEntityResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAlertsWithEntityResponseBody) SetData(v *DescribeAlertsWithEntityResponseBodyData) *DescribeAlertsWithEntityResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAlertsWithEntityResponseBody) SetMessage(v string) *DescribeAlertsWithEntityResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAlertsWithEntityResponseBody) SetRequestId(v string) *DescribeAlertsWithEntityResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAlertsWithEntityResponseBody) SetSuccess(v bool) *DescribeAlertsWithEntityResponseBody {
	s.Success = &v
	return s
}

type DescribeAlertsWithEntityResponseBodyData struct {
	// The pagination information.
	PageInfo *DescribeAlertsWithEntityResponseBodyDataPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The detailed data.
	ResponseData []*DescribeAlertsWithEntityResponseBodyDataResponseData `json:"ResponseData,omitempty" xml:"ResponseData,omitempty" type:"Repeated"`
}

func (s DescribeAlertsWithEntityResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertsWithEntityResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAlertsWithEntityResponseBodyData) SetPageInfo(v *DescribeAlertsWithEntityResponseBodyDataPageInfo) *DescribeAlertsWithEntityResponseBodyData {
	s.PageInfo = v
	return s
}

func (s *DescribeAlertsWithEntityResponseBodyData) SetResponseData(v []*DescribeAlertsWithEntityResponseBodyDataResponseData) *DescribeAlertsWithEntityResponseBodyData {
	s.ResponseData = v
	return s
}

type DescribeAlertsWithEntityResponseBodyDataPageInfo struct {
	// The current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAlertsWithEntityResponseBodyDataPageInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertsWithEntityResponseBodyDataPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeAlertsWithEntityResponseBodyDataPageInfo) SetCurrentPage(v int32) *DescribeAlertsWithEntityResponseBodyDataPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeAlertsWithEntityResponseBodyDataPageInfo) SetPageSize(v int32) *DescribeAlertsWithEntityResponseBodyDataPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeAlertsWithEntityResponseBodyDataPageInfo) SetTotalCount(v int64) *DescribeAlertsWithEntityResponseBodyDataPageInfo {
	s.TotalCount = &v
	return s
}

type DescribeAlertsWithEntityResponseBodyDataResponseData struct {
	// The description of the alert.
	//
	// example:
	//
	// The detection model found a suspicious Webshell file on your server, which may be a backdoor file implanted to maintain permissions after the attacker successfully invaded the website.
	AlertDesc *string `json:"AlertDesc,omitempty" xml:"AlertDesc,omitempty"`
	// The internal code of the alert description.
	//
	// example:
	//
	// security_event_config.event_name.webshell
	AlertDescCode *string `json:"AlertDescCode,omitempty" xml:"AlertDescCode,omitempty"`
	// The alert description in English.
	//
	// example:
	//
	// The detection model found a suspicious Webshell file on your server, which may be a backdoor file implanted to maintain permissions after the attacker successfully invaded the website.
	AlertDescEn *string `json:"AlertDescEn,omitempty" xml:"AlertDescEn,omitempty"`
	// The details of the alert.
	//
	// example:
	//
	// {"main_user_id": "165295629792****";"log_uuid_count": "99";"attack_ip": "21.92.*.*"}
	AlertDetail *string `json:"AlertDetail,omitempty" xml:"AlertDetail,omitempty"`
	// The displayed details of the alert.
	//
	// example:
	//
	// aliyun
	AlertInfoList []*DescribeAlertsWithEntityResponseBodyDataResponseDataAlertInfoList `json:"AlertInfoList,omitempty" xml:"AlertInfoList,omitempty" type:"Repeated"`
	// The risk level. Valid values:
	//
	// 	- serious: high
	//
	// 	- suspicious: medium
	//
	// 	- remind: low
	//
	// example:
	//
	// remind
	AlertLevel *string `json:"AlertLevel,omitempty" xml:"AlertLevel,omitempty"`
	// The name of the alert.
	//
	// example:
	//
	// Try SNMP weak password
	AlertName *string `json:"AlertName,omitempty" xml:"AlertName,omitempty"`
	// The internal code of the alert name.
	//
	// example:
	//
	// security_event_config.event_name.webshell
	AlertNameCode *string `json:"AlertNameCode,omitempty" xml:"AlertNameCode,omitempty"`
	// The name of the alert.
	//
	// example:
	//
	// Try SNMP weak password
	AlertNameEn *string `json:"AlertNameEn,omitempty" xml:"AlertNameEn,omitempty"`
	// The source of the alert.
	//
	// example:
	//
	// sas
	AlertSrcProd *string `json:"AlertSrcProd,omitempty" xml:"AlertSrcProd,omitempty"`
	// The sub-module of the alert source.
	//
	// example:
	//
	// waf
	AlertSrcProdModule *string `json:"AlertSrcProdModule,omitempty" xml:"AlertSrcProdModule,omitempty"`
	// The title of the alert.
	//
	// example:
	//
	// Scan-Try SNMP weak password
	AlertTitle *string `json:"AlertTitle,omitempty" xml:"AlertTitle,omitempty"`
	// The alert title in English.
	//
	// example:
	//
	// Scan-Try SNMP weak password
	AlertTitleEn *string `json:"AlertTitleEn,omitempty" xml:"AlertTitleEn,omitempty"`
	// The type of the alert.
	//
	// example:
	//
	// Scan
	AlertType *string `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
	// The internal code of the alert type.
	//
	// example:
	//
	// security_event_config.event_name.webshellName
	AlertTypeCode *string `json:"AlertTypeCode,omitempty" xml:"AlertTypeCode,omitempty"`
	// The alert type in English.
	//
	// example:
	//
	// Scan
	AlertTypeEn *string `json:"AlertTypeEn,omitempty" xml:"AlertTypeEn,omitempty"`
	// The UUID of the alert.
	//
	// example:
	//
	// sas_71e24437d2797ce8fc59692905a4****
	AlertUuid *string `json:"AlertUuid,omitempty" xml:"AlertUuid,omitempty"`
	// The details of the asset.
	//
	// example:
	//
	// [
	//
	//       {
	//
	//             "is_main_asset": "1",
	//
	//             "asset_name": "47.245.*",
	//
	//             "port": "22",
	//
	//             "ip": "47.245.*",
	//
	//             "asset_type": "ip",
	//
	//             "location": "ap-southeast-1",
	//
	//             "asset_id": "47.245.*",
	//
	//             "net_connect_dir": "in"
	//
	//       }
	//
	// ]
	AssetList *string `json:"AssetList,omitempty" xml:"AssetList,omitempty"`
	// The tag of the ATT\\&CK attack.
	//
	// example:
	//
	// T1595.002 Vulnerability Scanning
	AttCk *string `json:"AttCk,omitempty" xml:"AttCk,omitempty"`
	// The cloud code. Valid values:
	//
	// 	- aliyun: Alibaba Cloud
	//
	// 	- qcloud: Tencent Cloud
	//
	// 	- hcloud: Huawei Cloud
	//
	// example:
	//
	// aliyun
	CloudCode *string `json:"CloudCode,omitempty" xml:"CloudCode,omitempty"`
	// The time when the alert was closed.
	//
	// example:
	//
	// 2021-01-06 16:37:29
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	EntityList *string `json:"EntityList,omitempty" xml:"EntityList,omitempty"`
	// The time when the alert was received.
	//
	// example:
	//
	// 2021-01-06 16:37:29
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the alert was last updated.
	//
	// example:
	//
	// 2021-01-06 16:37:29
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The unique ID of the alert.
	//
	// example:
	//
	// 123456789
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The UUID of the event.
	//
	// example:
	//
	// 85ea4241-798f-4684-a876-65d4f0c3****
	IncidentUuid *string `json:"IncidentUuid,omitempty" xml:"IncidentUuid,omitempty"`
	// Specifies whether an attack is defended. Valid values:
	//
	// 	- 0: detected
	//
	// 	- 1: blocked
	//
	// example:
	//
	// 1
	IsDefend *string `json:"IsDefend,omitempty" xml:"IsDefend,omitempty"`
	// The time when the alert was recorded.
	//
	// example:
	//
	// 2021-01-06 16:37:29
	LogTime *string `json:"LogTime,omitempty" xml:"LogTime,omitempty"`
	// The UUID of the alert log.
	//
	// example:
	//
	// cfw_d12e285a-a042-4d7e-be89-f8a795ef****
	LogUuid *string `json:"LogUuid,omitempty" xml:"LogUuid,omitempty"`
	// The ID of the Alibaba Cloud account that is associated with the alert in SIEM.
	//
	// example:
	//
	// 127608589417****
	MainUserId *int64 `json:"MainUserId,omitempty" xml:"MainUserId,omitempty"`
	// The time when the alert was triggered.
	//
	// example:
	//
	// 2021-01-06 16:37:29
	OccurTime *string `json:"OccurTime,omitempty" xml:"OccurTime,omitempty"`
	// The time at which the alert was first generated.
	//
	// example:
	//
	// 2021-01-06 16:37:29
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The ID of the Alibaba Cloud account within which the alert is generated.
	//
	// example:
	//
	// 176555323***
	SubUserId *int64 `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
	// example:
	//
	// 176555323***
	SubUserName *string `json:"SubUserName,omitempty" xml:"SubUserName,omitempty"`
}

func (s DescribeAlertsWithEntityResponseBodyDataResponseData) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertsWithEntityResponseBodyDataResponseData) GoString() string {
	return s.String()
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) SetAlertDesc(v string) *DescribeAlertsWithEntityResponseBodyDataResponseData {
	s.AlertDesc = &v
	return s
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) SetAlertDescCode(v string) *DescribeAlertsWithEntityResponseBodyDataResponseData {
	s.AlertDescCode = &v
	return s
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) SetAlertDescEn(v string) *DescribeAlertsWithEntityResponseBodyDataResponseData {
	s.AlertDescEn = &v
	return s
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) SetAlertDetail(v string) *DescribeAlertsWithEntityResponseBodyDataResponseData {
	s.AlertDetail = &v
	return s
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) SetAlertInfoList(v []*DescribeAlertsWithEntityResponseBodyDataResponseDataAlertInfoList) *DescribeAlertsWithEntityResponseBodyDataResponseData {
	s.AlertInfoList = v
	return s
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) SetAlertLevel(v string) *DescribeAlertsWithEntityResponseBodyDataResponseData {
	s.AlertLevel = &v
	return s
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) SetAlertName(v string) *DescribeAlertsWithEntityResponseBodyDataResponseData {
	s.AlertName = &v
	return s
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) SetAlertNameCode(v string) *DescribeAlertsWithEntityResponseBodyDataResponseData {
	s.AlertNameCode = &v
	return s
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) SetAlertNameEn(v string) *DescribeAlertsWithEntityResponseBodyDataResponseData {
	s.AlertNameEn = &v
	return s
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) SetAlertSrcProd(v string) *DescribeAlertsWithEntityResponseBodyDataResponseData {
	s.AlertSrcProd = &v
	return s
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) SetAlertSrcProdModule(v string) *DescribeAlertsWithEntityResponseBodyDataResponseData {
	s.AlertSrcProdModule = &v
	return s
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) SetAlertTitle(v string) *DescribeAlertsWithEntityResponseBodyDataResponseData {
	s.AlertTitle = &v
	return s
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) SetAlertTitleEn(v string) *DescribeAlertsWithEntityResponseBodyDataResponseData {
	s.AlertTitleEn = &v
	return s
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) SetAlertType(v string) *DescribeAlertsWithEntityResponseBodyDataResponseData {
	s.AlertType = &v
	return s
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) SetAlertTypeCode(v string) *DescribeAlertsWithEntityResponseBodyDataResponseData {
	s.AlertTypeCode = &v
	return s
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) SetAlertTypeEn(v string) *DescribeAlertsWithEntityResponseBodyDataResponseData {
	s.AlertTypeEn = &v
	return s
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) SetAlertUuid(v string) *DescribeAlertsWithEntityResponseBodyDataResponseData {
	s.AlertUuid = &v
	return s
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) SetAssetList(v string) *DescribeAlertsWithEntityResponseBodyDataResponseData {
	s.AssetList = &v
	return s
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) SetAttCk(v string) *DescribeAlertsWithEntityResponseBodyDataResponseData {
	s.AttCk = &v
	return s
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) SetCloudCode(v string) *DescribeAlertsWithEntityResponseBodyDataResponseData {
	s.CloudCode = &v
	return s
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) SetEndTime(v string) *DescribeAlertsWithEntityResponseBodyDataResponseData {
	s.EndTime = &v
	return s
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) SetEntityList(v string) *DescribeAlertsWithEntityResponseBodyDataResponseData {
	s.EntityList = &v
	return s
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) SetGmtCreate(v string) *DescribeAlertsWithEntityResponseBodyDataResponseData {
	s.GmtCreate = &v
	return s
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) SetGmtModified(v string) *DescribeAlertsWithEntityResponseBodyDataResponseData {
	s.GmtModified = &v
	return s
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) SetId(v int64) *DescribeAlertsWithEntityResponseBodyDataResponseData {
	s.Id = &v
	return s
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) SetIncidentUuid(v string) *DescribeAlertsWithEntityResponseBodyDataResponseData {
	s.IncidentUuid = &v
	return s
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) SetIsDefend(v string) *DescribeAlertsWithEntityResponseBodyDataResponseData {
	s.IsDefend = &v
	return s
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) SetLogTime(v string) *DescribeAlertsWithEntityResponseBodyDataResponseData {
	s.LogTime = &v
	return s
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) SetLogUuid(v string) *DescribeAlertsWithEntityResponseBodyDataResponseData {
	s.LogUuid = &v
	return s
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) SetMainUserId(v int64) *DescribeAlertsWithEntityResponseBodyDataResponseData {
	s.MainUserId = &v
	return s
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) SetOccurTime(v string) *DescribeAlertsWithEntityResponseBodyDataResponseData {
	s.OccurTime = &v
	return s
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) SetStartTime(v string) *DescribeAlertsWithEntityResponseBodyDataResponseData {
	s.StartTime = &v
	return s
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) SetSubUserId(v int64) *DescribeAlertsWithEntityResponseBodyDataResponseData {
	s.SubUserId = &v
	return s
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) SetSubUserName(v string) *DescribeAlertsWithEntityResponseBodyDataResponseData {
	s.SubUserName = &v
	return s
}

type DescribeAlertsWithEntityResponseBodyDataResponseDataAlertInfoList struct {
	// The attribute key.
	//
	// example:
	//
	// suspicious.wbd.wb.trojanpath
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The name of the key.
	//
	// example:
	//
	// Trojan Path
	KeyName *string `json:"KeyName,omitempty" xml:"KeyName,omitempty"`
	// The value of the key.
	//
	// example:
	//
	// /root/test33.php
	Values *string `json:"Values,omitempty" xml:"Values,omitempty"`
}

func (s DescribeAlertsWithEntityResponseBodyDataResponseDataAlertInfoList) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertsWithEntityResponseBodyDataResponseDataAlertInfoList) GoString() string {
	return s.String()
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseDataAlertInfoList) SetKey(v string) *DescribeAlertsWithEntityResponseBodyDataResponseDataAlertInfoList {
	s.Key = &v
	return s
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseDataAlertInfoList) SetKeyName(v string) *DescribeAlertsWithEntityResponseBodyDataResponseDataAlertInfoList {
	s.KeyName = &v
	return s
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseDataAlertInfoList) SetValues(v string) *DescribeAlertsWithEntityResponseBodyDataResponseDataAlertInfoList {
	s.Values = &v
	return s
}

type DescribeAlertsWithEntityResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAlertsWithEntityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAlertsWithEntityResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertsWithEntityResponse) GoString() string {
	return s.String()
}

func (s *DescribeAlertsWithEntityResponse) SetHeaders(v map[string]*string) *DescribeAlertsWithEntityResponse {
	s.Headers = v
	return s
}

func (s *DescribeAlertsWithEntityResponse) SetStatusCode(v int32) *DescribeAlertsWithEntityResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAlertsWithEntityResponse) SetBody(v *DescribeAlertsWithEntityResponseBody) *DescribeAlertsWithEntityResponse {
	s.Body = v
	return s
}

type DescribeAlertsWithEventRequest struct {
	// example:
	//
	// 异常连接-TFTP恶意扫描
	AlertName *string `json:"AlertName,omitempty" xml:"AlertName,omitempty"`
	// The title of the alert.
	//
	// example:
	//
	// Scan-Try SNMP weak password
	AlertTitle *string `json:"AlertTitle,omitempty" xml:"AlertTitle,omitempty"`
	// example:
	//
	// 异常连接-TFTP恶意扫描
	AlertType *string `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
	// example:
	//
	// 异常连接-TFTP恶意扫描
	AssetId *string `json:"AssetId,omitempty" xml:"AssetId,omitempty"`
	// example:
	//
	// 异常连接-TFTP恶意扫描
	AssetName *string `json:"AssetName,omitempty" xml:"AssetName,omitempty"`
	// The page number. Pages start from page 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 1577808000000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 异常连接-TFTP恶意扫描
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// example:
	//
	// 异常连接-TFTP恶意扫描
	EntityName *string `json:"EntityName,omitempty" xml:"EntityName,omitempty"`
	// The ID of the event.
	//
	// example:
	//
	// 85ea4241-798f-4684-a876-65d4f0c3****
	IncidentUuid *string `json:"IncidentUuid,omitempty" xml:"IncidentUuid,omitempty"`
	// Specifies whether an attack is defended. Valid values:
	//
	// 	- 0: detected
	//
	// 	- 1: blocked
	//
	// example:
	//
	// 1
	IsDefend *string `json:"IsDefend,omitempty" xml:"IsDefend,omitempty"`
	// The risk levels. The value is a JSON array. Valid values:
	//
	// 	- serious: high
	//
	// 	- suspicious: medium
	//
	// 	- remind: low
	//
	// example:
	//
	// ["serious","suspicious","remind"]
	Level []*string `json:"Level,omitempty" xml:"Level,omitempty" type:"Repeated"`
	// The number of entries per page. Maximum value: 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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
	// The ID of the member in the resource directory.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view.
	//
	// 	- 0: view of the current Alibaba Cloud account.
	//
	// 	- 1: view of all accounts for the enterprise.
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// The data source of the alert.
	//
	// example:
	//
	// sas
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// 1577808000000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The ID of the account within which the alert is generated.
	//
	// example:
	//
	// 176555323***
	SubUserId *int64 `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
}

func (s DescribeAlertsWithEventRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertsWithEventRequest) GoString() string {
	return s.String()
}

func (s *DescribeAlertsWithEventRequest) SetAlertName(v string) *DescribeAlertsWithEventRequest {
	s.AlertName = &v
	return s
}

func (s *DescribeAlertsWithEventRequest) SetAlertTitle(v string) *DescribeAlertsWithEventRequest {
	s.AlertTitle = &v
	return s
}

func (s *DescribeAlertsWithEventRequest) SetAlertType(v string) *DescribeAlertsWithEventRequest {
	s.AlertType = &v
	return s
}

func (s *DescribeAlertsWithEventRequest) SetAssetId(v string) *DescribeAlertsWithEventRequest {
	s.AssetId = &v
	return s
}

func (s *DescribeAlertsWithEventRequest) SetAssetName(v string) *DescribeAlertsWithEventRequest {
	s.AssetName = &v
	return s
}

func (s *DescribeAlertsWithEventRequest) SetCurrentPage(v int32) *DescribeAlertsWithEventRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeAlertsWithEventRequest) SetEndTime(v int64) *DescribeAlertsWithEventRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeAlertsWithEventRequest) SetEntityId(v string) *DescribeAlertsWithEventRequest {
	s.EntityId = &v
	return s
}

func (s *DescribeAlertsWithEventRequest) SetEntityName(v string) *DescribeAlertsWithEventRequest {
	s.EntityName = &v
	return s
}

func (s *DescribeAlertsWithEventRequest) SetIncidentUuid(v string) *DescribeAlertsWithEventRequest {
	s.IncidentUuid = &v
	return s
}

func (s *DescribeAlertsWithEventRequest) SetIsDefend(v string) *DescribeAlertsWithEventRequest {
	s.IsDefend = &v
	return s
}

func (s *DescribeAlertsWithEventRequest) SetLevel(v []*string) *DescribeAlertsWithEventRequest {
	s.Level = v
	return s
}

func (s *DescribeAlertsWithEventRequest) SetPageSize(v int32) *DescribeAlertsWithEventRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAlertsWithEventRequest) SetRegionId(v string) *DescribeAlertsWithEventRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAlertsWithEventRequest) SetRoleFor(v int64) *DescribeAlertsWithEventRequest {
	s.RoleFor = &v
	return s
}

func (s *DescribeAlertsWithEventRequest) SetRoleType(v int32) *DescribeAlertsWithEventRequest {
	s.RoleType = &v
	return s
}

func (s *DescribeAlertsWithEventRequest) SetSource(v string) *DescribeAlertsWithEventRequest {
	s.Source = &v
	return s
}

func (s *DescribeAlertsWithEventRequest) SetStartTime(v int64) *DescribeAlertsWithEventRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeAlertsWithEventRequest) SetSubUserId(v int64) *DescribeAlertsWithEventRequest {
	s.SubUserId = &v
	return s
}

type DescribeAlertsWithEventResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	//
	// example:
	//
	// 123456
	Data *DescribeAlertsWithEventResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeAlertsWithEventResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertsWithEventResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAlertsWithEventResponseBody) SetCode(v int32) *DescribeAlertsWithEventResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAlertsWithEventResponseBody) SetData(v *DescribeAlertsWithEventResponseBodyData) *DescribeAlertsWithEventResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAlertsWithEventResponseBody) SetMessage(v string) *DescribeAlertsWithEventResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAlertsWithEventResponseBody) SetRequestId(v string) *DescribeAlertsWithEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAlertsWithEventResponseBody) SetSuccess(v bool) *DescribeAlertsWithEventResponseBody {
	s.Success = &v
	return s
}

type DescribeAlertsWithEventResponseBodyData struct {
	// The pagination information.
	PageInfo *DescribeAlertsWithEventResponseBodyDataPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The detailed data.
	ResponseData []*DescribeAlertsWithEventResponseBodyDataResponseData `json:"ResponseData,omitempty" xml:"ResponseData,omitempty" type:"Repeated"`
}

func (s DescribeAlertsWithEventResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertsWithEventResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAlertsWithEventResponseBodyData) SetPageInfo(v *DescribeAlertsWithEventResponseBodyDataPageInfo) *DescribeAlertsWithEventResponseBodyData {
	s.PageInfo = v
	return s
}

func (s *DescribeAlertsWithEventResponseBodyData) SetResponseData(v []*DescribeAlertsWithEventResponseBodyDataResponseData) *DescribeAlertsWithEventResponseBodyData {
	s.ResponseData = v
	return s
}

type DescribeAlertsWithEventResponseBodyDataPageInfo struct {
	// The current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAlertsWithEventResponseBodyDataPageInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertsWithEventResponseBodyDataPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeAlertsWithEventResponseBodyDataPageInfo) SetCurrentPage(v int32) *DescribeAlertsWithEventResponseBodyDataPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeAlertsWithEventResponseBodyDataPageInfo) SetPageSize(v int32) *DescribeAlertsWithEventResponseBodyDataPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeAlertsWithEventResponseBodyDataPageInfo) SetTotalCount(v int64) *DescribeAlertsWithEventResponseBodyDataPageInfo {
	s.TotalCount = &v
	return s
}

type DescribeAlertsWithEventResponseBodyDataResponseData struct {
	// The description of the alert.
	//
	// example:
	//
	// The detection model found a suspicious Webshell file on your server, which may be a backdoor file implanted to maintain permissions after the attacker successfully invaded the website.
	AlertDesc *string `json:"AlertDesc,omitempty" xml:"AlertDesc,omitempty"`
	// The internal code of the alert description.
	//
	// example:
	//
	// security_event_config.event_name.webshell
	AlertDescCode *string `json:"AlertDescCode,omitempty" xml:"AlertDescCode,omitempty"`
	// The alert description in English.
	//
	// example:
	//
	// The detection model found a suspicious Webshell file on your server, which may be a backdoor file implanted to maintain permissions after the attacker successfully invaded the website.
	AlertDescEn *string `json:"AlertDescEn,omitempty" xml:"AlertDescEn,omitempty"`
	// The details of the alert.
	//
	// example:
	//
	// {"main_user_id": "165295629792****";"log_uuid_count": "99";"attack_ip": "21.92.*.*"}
	AlertDetail *string `json:"AlertDetail,omitempty" xml:"AlertDetail,omitempty"`
	// The displayed details of the alert.
	//
	// example:
	//
	// aliyun
	AlertInfoList []*DescribeAlertsWithEventResponseBodyDataResponseDataAlertInfoList `json:"AlertInfoList,omitempty" xml:"AlertInfoList,omitempty" type:"Repeated"`
	// The risk level. Valid values:
	//
	// 	- serious: high.
	//
	// 	- suspicious: medium.
	//
	// 	- remind: low.
	//
	// example:
	//
	// remind
	AlertLevel *string `json:"AlertLevel,omitempty" xml:"AlertLevel,omitempty"`
	// The name of the alert.
	//
	// example:
	//
	// Try SNMP weak password
	AlertName *string `json:"AlertName,omitempty" xml:"AlertName,omitempty"`
	// The internal code of the alert name.
	//
	// example:
	//
	// security_event_config.event_name.webshell
	AlertNameCode *string `json:"AlertNameCode,omitempty" xml:"AlertNameCode,omitempty"`
	// The alert name in English.
	//
	// example:
	//
	// Try SNMP weak password
	AlertNameEn *string `json:"AlertNameEn,omitempty" xml:"AlertNameEn,omitempty"`
	// The source of the alert.
	//
	// example:
	//
	// sas
	AlertSrcProd *string `json:"AlertSrcProd,omitempty" xml:"AlertSrcProd,omitempty"`
	// The sub-module of the alert source.
	//
	// example:
	//
	// waf
	AlertSrcProdModule *string `json:"AlertSrcProdModule,omitempty" xml:"AlertSrcProdModule,omitempty"`
	// The title of the alert.
	//
	// example:
	//
	// Scan-Try SNMP weak password
	AlertTitle *string `json:"AlertTitle,omitempty" xml:"AlertTitle,omitempty"`
	// The alert title in English.
	//
	// example:
	//
	// Scan-Try SNMP weak password
	AlertTitleEn *string `json:"AlertTitleEn,omitempty" xml:"AlertTitleEn,omitempty"`
	// The type of the alert.
	//
	// example:
	//
	// Scan
	AlertType *string `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
	// The internal code of the alert type.
	//
	// example:
	//
	// security_event_config.event_name.webshellName
	AlertTypeCode *string `json:"AlertTypeCode,omitempty" xml:"AlertTypeCode,omitempty"`
	// The alert type in English.
	//
	// example:
	//
	// Scan
	AlertTypeEn *string `json:"AlertTypeEn,omitempty" xml:"AlertTypeEn,omitempty"`
	// The UUID of the alert.
	//
	// example:
	//
	// sas_71e24437d2797ce8fc59692905a4****
	AlertUuid *string `json:"AlertUuid,omitempty" xml:"AlertUuid,omitempty"`
	// The details of the asset.
	//
	// example:
	//
	// [
	//
	//       {
	//
	//             "is_main_asset": "1",
	//
	//             "asset_name": "47.245.*",
	//
	//             "port": "22",
	//
	//             "ip": "47.245.*",
	//
	//             "asset_type": "ip",
	//
	//             "location": "ap-southeast-1",
	//
	//             "asset_id": "47.245.*",
	//
	//             "net_connect_dir": "in"
	//
	//       }
	//
	// ]
	AssetList *string `json:"AssetList,omitempty" xml:"AssetList,omitempty"`
	// The tag of the ATT\\&CK technique.
	//
	// example:
	//
	// T1595.002 Vulnerability Scanning
	AttCk *string `json:"AttCk,omitempty" xml:"AttCk,omitempty"`
	// The code of the cloud service provider. Valid values:
	//
	// 	- aliyun: Alibaba Cloud.
	//
	// 	- qcloud: Tencent Cloud.
	//
	// 	- hcloud: Huawei Cloud.
	//
	// example:
	//
	// aliyun
	CloudCode *string `json:"CloudCode,omitempty" xml:"CloudCode,omitempty"`
	// The time when the alert was closed.
	//
	// example:
	//
	// 2021-01-06 16:37:29
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The details of the entity.
	//
	// example:
	//
	// [{&quot;entity_user_id&quot;:&quot;198921674491****&quot;,&quot;entity_account_id&quot;:&quot;N/A&quot;,&quot;entity_uuid&quot;:&quot;6245f979d5dd9ef8dd19bdc72228****&quot;,&quot;entity_type&quot;:&quot;host&quot;,&quot;entity_name&quot;:&quot;zhh-test-20240409&quot;,&quot;is_comprised&quot;:&quot;1&quot;,&quot;os_type&quot;:&quot;linux&quot;,&quot;entity_id&quot;:&quot;a88f44dd-b8d4-4ded-831c-77a4835****&quot;,&quot;host_uuid&quot;:&quot;a88f44dd-b8d4-4ded-831c-77a4835****&quot;,&quot;host_name&quot;:&quot;zhh-test-2024****&quot;}]
	EntityList    *string `json:"EntityList,omitempty" xml:"EntityList,omitempty"`
	ExtendContent *string `json:"ExtendContent,omitempty" xml:"ExtendContent,omitempty"`
	// The time when the alert was received.
	//
	// example:
	//
	// 2021-01-06 16:37:29
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the alert was last updated.
	//
	// example:
	//
	// 2021-01-06 16:37:29
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The unique ID of the alert.
	//
	// example:
	//
	// 123456789
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The UUID of the event.
	//
	// example:
	//
	// 85ea4241-798f-4684-a876-65d4f0c3****
	IncidentUuid *string `json:"IncidentUuid,omitempty" xml:"IncidentUuid,omitempty"`
	// Indicates whether an attack is defended against. Valid values:
	//
	// 	- 0: detected.
	//
	// 	- 1: blocked.
	//
	// example:
	//
	// 1
	IsDefend *string `json:"IsDefend,omitempty" xml:"IsDefend,omitempty"`
	// The time when the alert was recorded.
	//
	// example:
	//
	// 2021-01-06 16:37:29
	LogTime *string `json:"LogTime,omitempty" xml:"LogTime,omitempty"`
	// The UUID of the alert log.
	//
	// example:
	//
	// cfw_d12e285a-a042-4d7e-be89-f8a795ef****
	LogUuid *string `json:"LogUuid,omitempty" xml:"LogUuid,omitempty"`
	// The ID of the Alibaba Cloud account that is associated with the alert in SIEM.
	//
	// example:
	//
	// 127608589417****
	MainUserId *int64 `json:"MainUserId,omitempty" xml:"MainUserId,omitempty"`
	// The time when the alert was triggered.
	//
	// example:
	//
	// 2021-01-06 16:37:29
	OccurTime *string `json:"OccurTime,omitempty" xml:"OccurTime,omitempty"`
	// The time at which the alert was first generated.
	//
	// example:
	//
	// 2021-01-06 16:37:29
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The ID of the Alibaba Cloud account within which the alert is generated.
	//
	// example:
	//
	// 176555323***
	SubUserId *int64 `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
	// example:
	//
	// 176555323***
	SubUserName *string `json:"SubUserName,omitempty" xml:"SubUserName,omitempty"`
}

func (s DescribeAlertsWithEventResponseBodyDataResponseData) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertsWithEventResponseBodyDataResponseData) GoString() string {
	return s.String()
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) SetAlertDesc(v string) *DescribeAlertsWithEventResponseBodyDataResponseData {
	s.AlertDesc = &v
	return s
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) SetAlertDescCode(v string) *DescribeAlertsWithEventResponseBodyDataResponseData {
	s.AlertDescCode = &v
	return s
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) SetAlertDescEn(v string) *DescribeAlertsWithEventResponseBodyDataResponseData {
	s.AlertDescEn = &v
	return s
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) SetAlertDetail(v string) *DescribeAlertsWithEventResponseBodyDataResponseData {
	s.AlertDetail = &v
	return s
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) SetAlertInfoList(v []*DescribeAlertsWithEventResponseBodyDataResponseDataAlertInfoList) *DescribeAlertsWithEventResponseBodyDataResponseData {
	s.AlertInfoList = v
	return s
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) SetAlertLevel(v string) *DescribeAlertsWithEventResponseBodyDataResponseData {
	s.AlertLevel = &v
	return s
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) SetAlertName(v string) *DescribeAlertsWithEventResponseBodyDataResponseData {
	s.AlertName = &v
	return s
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) SetAlertNameCode(v string) *DescribeAlertsWithEventResponseBodyDataResponseData {
	s.AlertNameCode = &v
	return s
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) SetAlertNameEn(v string) *DescribeAlertsWithEventResponseBodyDataResponseData {
	s.AlertNameEn = &v
	return s
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) SetAlertSrcProd(v string) *DescribeAlertsWithEventResponseBodyDataResponseData {
	s.AlertSrcProd = &v
	return s
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) SetAlertSrcProdModule(v string) *DescribeAlertsWithEventResponseBodyDataResponseData {
	s.AlertSrcProdModule = &v
	return s
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) SetAlertTitle(v string) *DescribeAlertsWithEventResponseBodyDataResponseData {
	s.AlertTitle = &v
	return s
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) SetAlertTitleEn(v string) *DescribeAlertsWithEventResponseBodyDataResponseData {
	s.AlertTitleEn = &v
	return s
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) SetAlertType(v string) *DescribeAlertsWithEventResponseBodyDataResponseData {
	s.AlertType = &v
	return s
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) SetAlertTypeCode(v string) *DescribeAlertsWithEventResponseBodyDataResponseData {
	s.AlertTypeCode = &v
	return s
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) SetAlertTypeEn(v string) *DescribeAlertsWithEventResponseBodyDataResponseData {
	s.AlertTypeEn = &v
	return s
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) SetAlertUuid(v string) *DescribeAlertsWithEventResponseBodyDataResponseData {
	s.AlertUuid = &v
	return s
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) SetAssetList(v string) *DescribeAlertsWithEventResponseBodyDataResponseData {
	s.AssetList = &v
	return s
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) SetAttCk(v string) *DescribeAlertsWithEventResponseBodyDataResponseData {
	s.AttCk = &v
	return s
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) SetCloudCode(v string) *DescribeAlertsWithEventResponseBodyDataResponseData {
	s.CloudCode = &v
	return s
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) SetEndTime(v string) *DescribeAlertsWithEventResponseBodyDataResponseData {
	s.EndTime = &v
	return s
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) SetEntityList(v string) *DescribeAlertsWithEventResponseBodyDataResponseData {
	s.EntityList = &v
	return s
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) SetExtendContent(v string) *DescribeAlertsWithEventResponseBodyDataResponseData {
	s.ExtendContent = &v
	return s
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) SetGmtCreate(v string) *DescribeAlertsWithEventResponseBodyDataResponseData {
	s.GmtCreate = &v
	return s
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) SetGmtModified(v string) *DescribeAlertsWithEventResponseBodyDataResponseData {
	s.GmtModified = &v
	return s
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) SetId(v int64) *DescribeAlertsWithEventResponseBodyDataResponseData {
	s.Id = &v
	return s
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) SetIncidentUuid(v string) *DescribeAlertsWithEventResponseBodyDataResponseData {
	s.IncidentUuid = &v
	return s
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) SetIsDefend(v string) *DescribeAlertsWithEventResponseBodyDataResponseData {
	s.IsDefend = &v
	return s
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) SetLogTime(v string) *DescribeAlertsWithEventResponseBodyDataResponseData {
	s.LogTime = &v
	return s
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) SetLogUuid(v string) *DescribeAlertsWithEventResponseBodyDataResponseData {
	s.LogUuid = &v
	return s
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) SetMainUserId(v int64) *DescribeAlertsWithEventResponseBodyDataResponseData {
	s.MainUserId = &v
	return s
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) SetOccurTime(v string) *DescribeAlertsWithEventResponseBodyDataResponseData {
	s.OccurTime = &v
	return s
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) SetStartTime(v string) *DescribeAlertsWithEventResponseBodyDataResponseData {
	s.StartTime = &v
	return s
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) SetSubUserId(v int64) *DescribeAlertsWithEventResponseBodyDataResponseData {
	s.SubUserId = &v
	return s
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) SetSubUserName(v string) *DescribeAlertsWithEventResponseBodyDataResponseData {
	s.SubUserName = &v
	return s
}

type DescribeAlertsWithEventResponseBodyDataResponseDataAlertInfoList struct {
	// The attribute key.
	//
	// example:
	//
	// suspicious.wbd.wb.trojanpath
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The name of the key.
	//
	// example:
	//
	// Trojan Path
	KeyName *string `json:"KeyName,omitempty" xml:"KeyName,omitempty"`
	// The value of the key.
	//
	// example:
	//
	// /root/test33.php
	Values *string `json:"Values,omitempty" xml:"Values,omitempty"`
}

func (s DescribeAlertsWithEventResponseBodyDataResponseDataAlertInfoList) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertsWithEventResponseBodyDataResponseDataAlertInfoList) GoString() string {
	return s.String()
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseDataAlertInfoList) SetKey(v string) *DescribeAlertsWithEventResponseBodyDataResponseDataAlertInfoList {
	s.Key = &v
	return s
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseDataAlertInfoList) SetKeyName(v string) *DescribeAlertsWithEventResponseBodyDataResponseDataAlertInfoList {
	s.KeyName = &v
	return s
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseDataAlertInfoList) SetValues(v string) *DescribeAlertsWithEventResponseBodyDataResponseDataAlertInfoList {
	s.Values = &v
	return s
}

type DescribeAlertsWithEventResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAlertsWithEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAlertsWithEventResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertsWithEventResponse) GoString() string {
	return s.String()
}

func (s *DescribeAlertsWithEventResponse) SetHeaders(v map[string]*string) *DescribeAlertsWithEventResponse {
	s.Headers = v
	return s
}

func (s *DescribeAlertsWithEventResponse) SetStatusCode(v int32) *DescribeAlertsWithEventResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAlertsWithEventResponse) SetBody(v *DescribeAlertsWithEventResponseBody) *DescribeAlertsWithEventResponse {
	s.Body = v
	return s
}

type DescribeAuthRequest struct {
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

func (s DescribeAuthRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAuthRequest) GoString() string {
	return s.String()
}

func (s *DescribeAuthRequest) SetRegionId(v string) *DescribeAuthRequest {
	s.RegionId = &v
	return s
}

type DescribeAuthResponseBody struct {
	// Indicates whether the SIEM system is granted the required permissions. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4F539347-7D9A-51EA-8ABF-5D5507045C5C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAuthResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAuthResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAuthResponseBody) SetData(v bool) *DescribeAuthResponseBody {
	s.Data = &v
	return s
}

func (s *DescribeAuthResponseBody) SetRequestId(v string) *DescribeAuthResponseBody {
	s.RequestId = &v
	return s
}

type DescribeAuthResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAuthResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAuthResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAuthResponse) GoString() string {
	return s.String()
}

func (s *DescribeAuthResponse) SetHeaders(v map[string]*string) *DescribeAuthResponse {
	s.Headers = v
	return s
}

func (s *DescribeAuthResponse) SetStatusCode(v int32) *DescribeAuthResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAuthResponse) SetBody(v *DescribeAuthResponseBody) *DescribeAuthResponse {
	s.Body = v
	return s
}

type DescribeAutomateResponseConfigCounterRequest struct {
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
	// The ID of the account that you switch from the management account.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view. Valid values:
	//
	// - 0: the current Alibaba Cloud account
	//
	// - 1: the global account
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s DescribeAutomateResponseConfigCounterRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutomateResponseConfigCounterRequest) GoString() string {
	return s.String()
}

func (s *DescribeAutomateResponseConfigCounterRequest) SetRegionId(v string) *DescribeAutomateResponseConfigCounterRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAutomateResponseConfigCounterRequest) SetRoleFor(v int64) *DescribeAutomateResponseConfigCounterRequest {
	s.RoleFor = &v
	return s
}

func (s *DescribeAutomateResponseConfigCounterRequest) SetRoleType(v int32) *DescribeAutomateResponseConfigCounterRequest {
	s.RoleType = &v
	return s
}

type DescribeAutomateResponseConfigCounterResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	//
	// example:
	//
	// 123456
	Data *DescribeAutomateResponseConfigCounterResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeAutomateResponseConfigCounterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutomateResponseConfigCounterResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAutomateResponseConfigCounterResponseBody) SetCode(v int32) *DescribeAutomateResponseConfigCounterResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAutomateResponseConfigCounterResponseBody) SetData(v *DescribeAutomateResponseConfigCounterResponseBodyData) *DescribeAutomateResponseConfigCounterResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAutomateResponseConfigCounterResponseBody) SetMessage(v string) *DescribeAutomateResponseConfigCounterResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAutomateResponseConfigCounterResponseBody) SetRequestId(v string) *DescribeAutomateResponseConfigCounterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAutomateResponseConfigCounterResponseBody) SetSuccess(v bool) *DescribeAutomateResponseConfigCounterResponseBody {
	s.Success = &v
	return s
}

type DescribeAutomateResponseConfigCounterResponseBodyData struct {
	// The total number of rules.
	//
	// example:
	//
	// 20
	All *int64 `json:"All,omitempty" xml:"All,omitempty"`
	// The number of enabled rules.
	//
	// example:
	//
	// 10
	Online *int64 `json:"Online,omitempty" xml:"Online,omitempty"`
}

func (s DescribeAutomateResponseConfigCounterResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutomateResponseConfigCounterResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAutomateResponseConfigCounterResponseBodyData) SetAll(v int64) *DescribeAutomateResponseConfigCounterResponseBodyData {
	s.All = &v
	return s
}

func (s *DescribeAutomateResponseConfigCounterResponseBodyData) SetOnline(v int64) *DescribeAutomateResponseConfigCounterResponseBodyData {
	s.Online = &v
	return s
}

type DescribeAutomateResponseConfigCounterResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAutomateResponseConfigCounterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAutomateResponseConfigCounterResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutomateResponseConfigCounterResponse) GoString() string {
	return s.String()
}

func (s *DescribeAutomateResponseConfigCounterResponse) SetHeaders(v map[string]*string) *DescribeAutomateResponseConfigCounterResponse {
	s.Headers = v
	return s
}

func (s *DescribeAutomateResponseConfigCounterResponse) SetStatusCode(v int32) *DescribeAutomateResponseConfigCounterResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAutomateResponseConfigCounterResponse) SetBody(v *DescribeAutomateResponseConfigCounterResponseBody) *DescribeAutomateResponseConfigCounterResponse {
	s.Body = v
	return s
}

type DescribeAutomateResponseConfigFeatureRequest struct {
	// The type of the automated response rule. Valid values:
	//
	// 	- event
	//
	// 	- alert
	//
	// example:
	//
	// event
	AutoResponseType *string `json:"AutoResponseType,omitempty" xml:"AutoResponseType,omitempty"`
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
	// The ID of the account that you switch from the management account.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view. Valid values:
	//
	// - 0: the current Alibaba Cloud account
	//
	// - 1: the global account
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s DescribeAutomateResponseConfigFeatureRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutomateResponseConfigFeatureRequest) GoString() string {
	return s.String()
}

func (s *DescribeAutomateResponseConfigFeatureRequest) SetAutoResponseType(v string) *DescribeAutomateResponseConfigFeatureRequest {
	s.AutoResponseType = &v
	return s
}

func (s *DescribeAutomateResponseConfigFeatureRequest) SetRegionId(v string) *DescribeAutomateResponseConfigFeatureRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAutomateResponseConfigFeatureRequest) SetRoleFor(v int64) *DescribeAutomateResponseConfigFeatureRequest {
	s.RoleFor = &v
	return s
}

func (s *DescribeAutomateResponseConfigFeatureRequest) SetRoleType(v int32) *DescribeAutomateResponseConfigFeatureRequest {
	s.RoleType = &v
	return s
}

type DescribeAutomateResponseConfigFeatureResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	//
	// example:
	//
	// 123456
	Data []*DescribeAutomateResponseConfigFeatureResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeAutomateResponseConfigFeatureResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutomateResponseConfigFeatureResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAutomateResponseConfigFeatureResponseBody) SetCode(v int32) *DescribeAutomateResponseConfigFeatureResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAutomateResponseConfigFeatureResponseBody) SetData(v []*DescribeAutomateResponseConfigFeatureResponseBodyData) *DescribeAutomateResponseConfigFeatureResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAutomateResponseConfigFeatureResponseBody) SetMessage(v string) *DescribeAutomateResponseConfigFeatureResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAutomateResponseConfigFeatureResponseBody) SetRequestId(v string) *DescribeAutomateResponseConfigFeatureResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAutomateResponseConfigFeatureResponseBody) SetSuccess(v bool) *DescribeAutomateResponseConfigFeatureResponseBody {
	s.Success = &v
	return s
}

type DescribeAutomateResponseConfigFeatureResponseBodyData struct {
	// The data type of the condition field in the automated response rule.
	//
	// example:
	//
	// varchar
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// The name of the condition field in the automated response rule.
	//
	// example:
	//
	// alert_desc
	Feature *string `json:"Feature,omitempty" xml:"Feature,omitempty"`
	// The enumerated values of the right operand for the field.
	RightValueEnums []*DescribeAutomateResponseConfigFeatureResponseBodyDataRightValueEnums `json:"RightValueEnums,omitempty" xml:"RightValueEnums,omitempty" type:"Repeated"`
	// The operators that are supported for the condition field.
	SupportOperators []*DescribeAutomateResponseConfigFeatureResponseBodyDataSupportOperators `json:"SupportOperators,omitempty" xml:"SupportOperators,omitempty" type:"Repeated"`
}

func (s DescribeAutomateResponseConfigFeatureResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutomateResponseConfigFeatureResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAutomateResponseConfigFeatureResponseBodyData) SetDataType(v string) *DescribeAutomateResponseConfigFeatureResponseBodyData {
	s.DataType = &v
	return s
}

func (s *DescribeAutomateResponseConfigFeatureResponseBodyData) SetFeature(v string) *DescribeAutomateResponseConfigFeatureResponseBodyData {
	s.Feature = &v
	return s
}

func (s *DescribeAutomateResponseConfigFeatureResponseBodyData) SetRightValueEnums(v []*DescribeAutomateResponseConfigFeatureResponseBodyDataRightValueEnums) *DescribeAutomateResponseConfigFeatureResponseBodyData {
	s.RightValueEnums = v
	return s
}

func (s *DescribeAutomateResponseConfigFeatureResponseBodyData) SetSupportOperators(v []*DescribeAutomateResponseConfigFeatureResponseBodyDataSupportOperators) *DescribeAutomateResponseConfigFeatureResponseBodyData {
	s.SupportOperators = v
	return s
}

type DescribeAutomateResponseConfigFeatureResponseBodyDataRightValueEnums struct {
	// The enumerated value of the right operand.
	//
	// example:
	//
	// serious
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// The internal code of the enumerated value.
	//
	// example:
	//
	// aliyun.siem.automate.feature.alert_level.serious
	ValueMds *string `json:"ValueMds,omitempty" xml:"ValueMds,omitempty"`
}

func (s DescribeAutomateResponseConfigFeatureResponseBodyDataRightValueEnums) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutomateResponseConfigFeatureResponseBodyDataRightValueEnums) GoString() string {
	return s.String()
}

func (s *DescribeAutomateResponseConfigFeatureResponseBodyDataRightValueEnums) SetValue(v string) *DescribeAutomateResponseConfigFeatureResponseBodyDataRightValueEnums {
	s.Value = &v
	return s
}

func (s *DescribeAutomateResponseConfigFeatureResponseBodyDataRightValueEnums) SetValueMds(v string) *DescribeAutomateResponseConfigFeatureResponseBodyDataRightValueEnums {
	s.ValueMds = &v
	return s
}

type DescribeAutomateResponseConfigFeatureResponseBodyDataSupportOperators struct {
	// Indicates whether the right operand is required. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	HasRightValue *bool `json:"HasRightValue,omitempty" xml:"HasRightValue,omitempty"`
	// The position of the operator in the operator list.
	//
	// example:
	//
	// 3
	Index *int32 `json:"Index,omitempty" xml:"Index,omitempty"`
	// The operator.
	//
	// example:
	//
	// <=
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// The description of the operator in Chinese.
	//
	// example:
	//
	// larger than or equal to
	OperatorDescCn *string `json:"OperatorDescCn,omitempty" xml:"OperatorDescCn,omitempty"`
	// The description of the operator in English.
	//
	// example:
	//
	// larger than or equal to
	OperatorDescEn *string `json:"OperatorDescEn,omitempty" xml:"OperatorDescEn,omitempty"`
	// The name of the operator.
	//
	// example:
	//
	// <=
	OperatorName *string `json:"OperatorName,omitempty" xml:"OperatorName,omitempty"`
	// The data types that are supported by the operator. The data types are separated by commas (,).
	//
	// example:
	//
	// varchar
	SupportDataType *string `json:"SupportDataType,omitempty" xml:"SupportDataType,omitempty"`
	// The scenarios that are supported by the operator. Multiple scenarios are separated by commas (,), such as aggregation scenarios. By default, this parameter is empty.
	//
	// example:
	//
	// [AGGREGATE]
	SupportTag []*string `json:"SupportTag,omitempty" xml:"SupportTag,omitempty" type:"Repeated"`
}

func (s DescribeAutomateResponseConfigFeatureResponseBodyDataSupportOperators) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutomateResponseConfigFeatureResponseBodyDataSupportOperators) GoString() string {
	return s.String()
}

func (s *DescribeAutomateResponseConfigFeatureResponseBodyDataSupportOperators) SetHasRightValue(v bool) *DescribeAutomateResponseConfigFeatureResponseBodyDataSupportOperators {
	s.HasRightValue = &v
	return s
}

func (s *DescribeAutomateResponseConfigFeatureResponseBodyDataSupportOperators) SetIndex(v int32) *DescribeAutomateResponseConfigFeatureResponseBodyDataSupportOperators {
	s.Index = &v
	return s
}

func (s *DescribeAutomateResponseConfigFeatureResponseBodyDataSupportOperators) SetOperator(v string) *DescribeAutomateResponseConfigFeatureResponseBodyDataSupportOperators {
	s.Operator = &v
	return s
}

func (s *DescribeAutomateResponseConfigFeatureResponseBodyDataSupportOperators) SetOperatorDescCn(v string) *DescribeAutomateResponseConfigFeatureResponseBodyDataSupportOperators {
	s.OperatorDescCn = &v
	return s
}

func (s *DescribeAutomateResponseConfigFeatureResponseBodyDataSupportOperators) SetOperatorDescEn(v string) *DescribeAutomateResponseConfigFeatureResponseBodyDataSupportOperators {
	s.OperatorDescEn = &v
	return s
}

func (s *DescribeAutomateResponseConfigFeatureResponseBodyDataSupportOperators) SetOperatorName(v string) *DescribeAutomateResponseConfigFeatureResponseBodyDataSupportOperators {
	s.OperatorName = &v
	return s
}

func (s *DescribeAutomateResponseConfigFeatureResponseBodyDataSupportOperators) SetSupportDataType(v string) *DescribeAutomateResponseConfigFeatureResponseBodyDataSupportOperators {
	s.SupportDataType = &v
	return s
}

func (s *DescribeAutomateResponseConfigFeatureResponseBodyDataSupportOperators) SetSupportTag(v []*string) *DescribeAutomateResponseConfigFeatureResponseBodyDataSupportOperators {
	s.SupportTag = v
	return s
}

type DescribeAutomateResponseConfigFeatureResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAutomateResponseConfigFeatureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAutomateResponseConfigFeatureResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutomateResponseConfigFeatureResponse) GoString() string {
	return s.String()
}

func (s *DescribeAutomateResponseConfigFeatureResponse) SetHeaders(v map[string]*string) *DescribeAutomateResponseConfigFeatureResponse {
	s.Headers = v
	return s
}

func (s *DescribeAutomateResponseConfigFeatureResponse) SetStatusCode(v int32) *DescribeAutomateResponseConfigFeatureResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAutomateResponseConfigFeatureResponse) SetBody(v *DescribeAutomateResponseConfigFeatureResponseBody) *DescribeAutomateResponseConfigFeatureResponse {
	s.Body = v
	return s
}

type DescribeAutomateResponseConfigPlayBooksRequest struct {
	// The type of the automated response rule. Valid values:
	//
	// 	- event
	//
	// 	- alert
	//
	// example:
	//
	// event
	AutoResponseType *string `json:"AutoResponseType,omitempty" xml:"AutoResponseType,omitempty"`
	// The entity type of the playbook. Valid values:
	//
	// 	- ip
	//
	// 	- process
	//
	// 	- file
	//
	// example:
	//
	// ip
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
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
	// The ID of the account that you switch from the management account.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view. Valid values:
	//
	// - 0: the current Alibaba Cloud account
	//
	// - 1: the global account
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s DescribeAutomateResponseConfigPlayBooksRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutomateResponseConfigPlayBooksRequest) GoString() string {
	return s.String()
}

func (s *DescribeAutomateResponseConfigPlayBooksRequest) SetAutoResponseType(v string) *DescribeAutomateResponseConfigPlayBooksRequest {
	s.AutoResponseType = &v
	return s
}

func (s *DescribeAutomateResponseConfigPlayBooksRequest) SetEntityType(v string) *DescribeAutomateResponseConfigPlayBooksRequest {
	s.EntityType = &v
	return s
}

func (s *DescribeAutomateResponseConfigPlayBooksRequest) SetRegionId(v string) *DescribeAutomateResponseConfigPlayBooksRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAutomateResponseConfigPlayBooksRequest) SetRoleFor(v int64) *DescribeAutomateResponseConfigPlayBooksRequest {
	s.RoleFor = &v
	return s
}

func (s *DescribeAutomateResponseConfigPlayBooksRequest) SetRoleType(v int32) *DescribeAutomateResponseConfigPlayBooksRequest {
	s.RoleType = &v
	return s
}

type DescribeAutomateResponseConfigPlayBooksResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	//
	// example:
	//
	// 123456
	Data []*DescribeAutomateResponseConfigPlayBooksResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeAutomateResponseConfigPlayBooksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutomateResponseConfigPlayBooksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAutomateResponseConfigPlayBooksResponseBody) SetCode(v int32) *DescribeAutomateResponseConfigPlayBooksResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAutomateResponseConfigPlayBooksResponseBody) SetData(v []*DescribeAutomateResponseConfigPlayBooksResponseBodyData) *DescribeAutomateResponseConfigPlayBooksResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAutomateResponseConfigPlayBooksResponseBody) SetMessage(v string) *DescribeAutomateResponseConfigPlayBooksResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAutomateResponseConfigPlayBooksResponseBody) SetRequestId(v string) *DescribeAutomateResponseConfigPlayBooksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAutomateResponseConfigPlayBooksResponseBody) SetSuccess(v bool) *DescribeAutomateResponseConfigPlayBooksResponseBody {
	s.Success = &v
	return s
}

type DescribeAutomateResponseConfigPlayBooksResponseBodyData struct {
	// The description of the playbook.
	//
	// example:
	//
	// Waf Block IP
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The display name of the playbook.
	//
	// example:
	//
	// WafBlockIP
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The unique identifier name of the playbook.
	//
	// example:
	//
	// WafBlockIP
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The input parameter template of the playbook. Valid values:
	//
	// 	- template-ip: IP address
	//
	// 	- template-process: process
	//
	// 	- template-filee: file
	//
	// example:
	//
	// template-ip
	ParamType *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
	// The UUID of the playbook.
	//
	// example:
	//
	// system_aliyun_clb_process_book
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeAutomateResponseConfigPlayBooksResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutomateResponseConfigPlayBooksResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAutomateResponseConfigPlayBooksResponseBodyData) SetDescription(v string) *DescribeAutomateResponseConfigPlayBooksResponseBodyData {
	s.Description = &v
	return s
}

func (s *DescribeAutomateResponseConfigPlayBooksResponseBodyData) SetDisplayName(v string) *DescribeAutomateResponseConfigPlayBooksResponseBodyData {
	s.DisplayName = &v
	return s
}

func (s *DescribeAutomateResponseConfigPlayBooksResponseBodyData) SetName(v string) *DescribeAutomateResponseConfigPlayBooksResponseBodyData {
	s.Name = &v
	return s
}

func (s *DescribeAutomateResponseConfigPlayBooksResponseBodyData) SetParamType(v string) *DescribeAutomateResponseConfigPlayBooksResponseBodyData {
	s.ParamType = &v
	return s
}

func (s *DescribeAutomateResponseConfigPlayBooksResponseBodyData) SetUuid(v string) *DescribeAutomateResponseConfigPlayBooksResponseBodyData {
	s.Uuid = &v
	return s
}

type DescribeAutomateResponseConfigPlayBooksResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAutomateResponseConfigPlayBooksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAutomateResponseConfigPlayBooksResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutomateResponseConfigPlayBooksResponse) GoString() string {
	return s.String()
}

func (s *DescribeAutomateResponseConfigPlayBooksResponse) SetHeaders(v map[string]*string) *DescribeAutomateResponseConfigPlayBooksResponse {
	s.Headers = v
	return s
}

func (s *DescribeAutomateResponseConfigPlayBooksResponse) SetStatusCode(v int32) *DescribeAutomateResponseConfigPlayBooksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAutomateResponseConfigPlayBooksResponse) SetBody(v *DescribeAutomateResponseConfigPlayBooksResponseBody) *DescribeAutomateResponseConfigPlayBooksResponse {
	s.Body = v
	return s
}

type DescribeCloudSiemAssetsRequest struct {
	// example:
	//
	// test123
	AssetName *string `json:"AssetName,omitempty" xml:"AssetName,omitempty"`
	// The type of the asset. Valid values:
	//
	// 	- ip
	//
	// 	- domain
	//
	// 	- url
	//
	// 	- process
	//
	// 	- file
	//
	// 	- host
	//
	// example:
	//
	// ip
	AssetType *string `json:"AssetType,omitempty" xml:"AssetType,omitempty"`
	// example:
	//
	// 123456-2222-3333-5555-3435345****
	AssetUuid *string `json:"AssetUuid,omitempty" xml:"AssetUuid,omitempty"`
	// The page number. Pages start from page 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The UUID of the event.
	//
	// example:
	//
	// 85ea4241-798f-4684-a876-65d4f0c3****
	IncidentUuid *string `json:"IncidentUuid,omitempty" xml:"IncidentUuid,omitempty"`
	// The number of entries per page. Maximum value: 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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
	// The ID of the account that you switch from the management account.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view. Valid values:
	//
	// - 0: the current Alibaba Cloud account
	//
	// - 1: the global account
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s DescribeCloudSiemAssetsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudSiemAssetsRequest) GoString() string {
	return s.String()
}

func (s *DescribeCloudSiemAssetsRequest) SetAssetName(v string) *DescribeCloudSiemAssetsRequest {
	s.AssetName = &v
	return s
}

func (s *DescribeCloudSiemAssetsRequest) SetAssetType(v string) *DescribeCloudSiemAssetsRequest {
	s.AssetType = &v
	return s
}

func (s *DescribeCloudSiemAssetsRequest) SetAssetUuid(v string) *DescribeCloudSiemAssetsRequest {
	s.AssetUuid = &v
	return s
}

func (s *DescribeCloudSiemAssetsRequest) SetCurrentPage(v int32) *DescribeCloudSiemAssetsRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeCloudSiemAssetsRequest) SetIncidentUuid(v string) *DescribeCloudSiemAssetsRequest {
	s.IncidentUuid = &v
	return s
}

func (s *DescribeCloudSiemAssetsRequest) SetPageSize(v int32) *DescribeCloudSiemAssetsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCloudSiemAssetsRequest) SetRegionId(v string) *DescribeCloudSiemAssetsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCloudSiemAssetsRequest) SetRoleFor(v int64) *DescribeCloudSiemAssetsRequest {
	s.RoleFor = &v
	return s
}

func (s *DescribeCloudSiemAssetsRequest) SetRoleType(v int32) *DescribeCloudSiemAssetsRequest {
	s.RoleType = &v
	return s
}

type DescribeCloudSiemAssetsResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	//
	// example:
	//
	// 123456
	Data *DescribeCloudSiemAssetsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeCloudSiemAssetsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudSiemAssetsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCloudSiemAssetsResponseBody) SetCode(v int32) *DescribeCloudSiemAssetsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeCloudSiemAssetsResponseBody) SetData(v *DescribeCloudSiemAssetsResponseBodyData) *DescribeCloudSiemAssetsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeCloudSiemAssetsResponseBody) SetMessage(v string) *DescribeCloudSiemAssetsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeCloudSiemAssetsResponseBody) SetRequestId(v string) *DescribeCloudSiemAssetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCloudSiemAssetsResponseBody) SetSuccess(v bool) *DescribeCloudSiemAssetsResponseBody {
	s.Success = &v
	return s
}

type DescribeCloudSiemAssetsResponseBodyData struct {
	// The pagination information.
	PageInfo *DescribeCloudSiemAssetsResponseBodyDataPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The detailed data.
	ResponseData []*DescribeCloudSiemAssetsResponseBodyDataResponseData `json:"ResponseData,omitempty" xml:"ResponseData,omitempty" type:"Repeated"`
}

func (s DescribeCloudSiemAssetsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudSiemAssetsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeCloudSiemAssetsResponseBodyData) SetPageInfo(v *DescribeCloudSiemAssetsResponseBodyDataPageInfo) *DescribeCloudSiemAssetsResponseBodyData {
	s.PageInfo = v
	return s
}

func (s *DescribeCloudSiemAssetsResponseBodyData) SetResponseData(v []*DescribeCloudSiemAssetsResponseBodyDataResponseData) *DescribeCloudSiemAssetsResponseBodyData {
	s.ResponseData = v
	return s
}

type DescribeCloudSiemAssetsResponseBodyDataPageInfo struct {
	// The current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCloudSiemAssetsResponseBodyDataPageInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudSiemAssetsResponseBodyDataPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeCloudSiemAssetsResponseBodyDataPageInfo) SetCurrentPage(v int32) *DescribeCloudSiemAssetsResponseBodyDataPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeCloudSiemAssetsResponseBodyDataPageInfo) SetPageSize(v int32) *DescribeCloudSiemAssetsResponseBodyDataPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeCloudSiemAssetsResponseBodyDataPageInfo) SetTotalCount(v int64) *DescribeCloudSiemAssetsResponseBodyDataPageInfo {
	s.TotalCount = &v
	return s
}

type DescribeCloudSiemAssetsResponseBodyDataResponseData struct {
	// The UUID of the alert associated with the event.
	//
	// example:
	//
	// sas_71e24437d2797ce8fc59692905a4****
	AlertUuid *string `json:"AlertUuid,omitempty" xml:"AlertUuid,omitempty"`
	// The ID of the Alibaba Cloud account in SIEM.
	//
	// example:
	//
	// 1276085894174392
	Aliuid *int64 `json:"Aliuid,omitempty" xml:"Aliuid,omitempty"`
	// The logical ID of the asset.
	//
	// example:
	//
	// 0616caeb-acb8-45e0-8520-4ee5fbe251f0
	AssetId *string `json:"AssetId,omitempty" xml:"AssetId,omitempty"`
	// The display information of the asset is in the JSON format.
	//
	// example:
	//
	// [{"KeyName": "${aliyun.siem.asset.asset_name}","Values": "zsw-agentless-ubuntu20","Key": "asset_name"}]
	AssetInfo []*DescribeCloudSiemAssetsResponseBodyDataResponseDataAssetInfo `json:"AssetInfo,omitempty" xml:"AssetInfo,omitempty" type:"Repeated"`
	// The name of the asset.
	//
	// example:
	//
	// zsw-agentless-centos****
	AssetName *string `json:"AssetName,omitempty" xml:"AssetName,omitempty"`
	// The type of the asset. Valid values:
	//
	// 	- ip
	//
	// 	- domain
	//
	// 	- url
	//
	// 	- process
	//
	// 	- file
	//
	// 	- host
	//
	// example:
	//
	// domain
	AssetType *string `json:"AssetType,omitempty" xml:"AssetType,omitempty"`
	// The cloud code of the entity. Valid values:
	//
	// 	- aliyun: Alibaba Cloud
	//
	// 	- qcloud: Tencent Cloud
	//
	// 	- hcloud: Huawei Cloud
	//
	// example:
	//
	// aliyun
	CloudCode *string `json:"CloudCode,omitempty" xml:"CloudCode,omitempty"`
	// The time when the asset was synchronized.
	//
	// example:
	//
	// 2021-01-06 16:37:29
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the asset was last updated.
	//
	// example:
	//
	// 2021-01-06 16:37:29
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the asset.
	//
	// example:
	//
	// 123
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The UUID of the event.
	//
	// example:
	//
	// 85ea4241-798f-4684-a876-65d4f0c3****
	IncidentUuid *string `json:"IncidentUuid,omitempty" xml:"IncidentUuid,omitempty"`
	// The ID of the associated account to which the asset belongs.
	//
	// example:
	//
	// 176555323***
	SubUserId *int64 `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
}

func (s DescribeCloudSiemAssetsResponseBodyDataResponseData) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudSiemAssetsResponseBodyDataResponseData) GoString() string {
	return s.String()
}

func (s *DescribeCloudSiemAssetsResponseBodyDataResponseData) SetAlertUuid(v string) *DescribeCloudSiemAssetsResponseBodyDataResponseData {
	s.AlertUuid = &v
	return s
}

func (s *DescribeCloudSiemAssetsResponseBodyDataResponseData) SetAliuid(v int64) *DescribeCloudSiemAssetsResponseBodyDataResponseData {
	s.Aliuid = &v
	return s
}

func (s *DescribeCloudSiemAssetsResponseBodyDataResponseData) SetAssetId(v string) *DescribeCloudSiemAssetsResponseBodyDataResponseData {
	s.AssetId = &v
	return s
}

func (s *DescribeCloudSiemAssetsResponseBodyDataResponseData) SetAssetInfo(v []*DescribeCloudSiemAssetsResponseBodyDataResponseDataAssetInfo) *DescribeCloudSiemAssetsResponseBodyDataResponseData {
	s.AssetInfo = v
	return s
}

func (s *DescribeCloudSiemAssetsResponseBodyDataResponseData) SetAssetName(v string) *DescribeCloudSiemAssetsResponseBodyDataResponseData {
	s.AssetName = &v
	return s
}

func (s *DescribeCloudSiemAssetsResponseBodyDataResponseData) SetAssetType(v string) *DescribeCloudSiemAssetsResponseBodyDataResponseData {
	s.AssetType = &v
	return s
}

func (s *DescribeCloudSiemAssetsResponseBodyDataResponseData) SetCloudCode(v string) *DescribeCloudSiemAssetsResponseBodyDataResponseData {
	s.CloudCode = &v
	return s
}

func (s *DescribeCloudSiemAssetsResponseBodyDataResponseData) SetGmtCreate(v string) *DescribeCloudSiemAssetsResponseBodyDataResponseData {
	s.GmtCreate = &v
	return s
}

func (s *DescribeCloudSiemAssetsResponseBodyDataResponseData) SetGmtModified(v string) *DescribeCloudSiemAssetsResponseBodyDataResponseData {
	s.GmtModified = &v
	return s
}

func (s *DescribeCloudSiemAssetsResponseBodyDataResponseData) SetId(v int64) *DescribeCloudSiemAssetsResponseBodyDataResponseData {
	s.Id = &v
	return s
}

func (s *DescribeCloudSiemAssetsResponseBodyDataResponseData) SetIncidentUuid(v string) *DescribeCloudSiemAssetsResponseBodyDataResponseData {
	s.IncidentUuid = &v
	return s
}

func (s *DescribeCloudSiemAssetsResponseBodyDataResponseData) SetSubUserId(v int64) *DescribeCloudSiemAssetsResponseBodyDataResponseData {
	s.SubUserId = &v
	return s
}

type DescribeCloudSiemAssetsResponseBodyDataResponseDataAssetInfo struct {
	// The attribute key.
	//
	// example:
	//
	// suspicious.wbd.wb.trojanpath
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The name of the key.
	//
	// example:
	//
	// Trojan Path
	KeyName *string `json:"KeyName,omitempty" xml:"KeyName,omitempty"`
	// The value of the key.
	//
	// example:
	//
	// /root/test33.php
	Values *string `json:"Values,omitempty" xml:"Values,omitempty"`
}

func (s DescribeCloudSiemAssetsResponseBodyDataResponseDataAssetInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudSiemAssetsResponseBodyDataResponseDataAssetInfo) GoString() string {
	return s.String()
}

func (s *DescribeCloudSiemAssetsResponseBodyDataResponseDataAssetInfo) SetKey(v string) *DescribeCloudSiemAssetsResponseBodyDataResponseDataAssetInfo {
	s.Key = &v
	return s
}

func (s *DescribeCloudSiemAssetsResponseBodyDataResponseDataAssetInfo) SetKeyName(v string) *DescribeCloudSiemAssetsResponseBodyDataResponseDataAssetInfo {
	s.KeyName = &v
	return s
}

func (s *DescribeCloudSiemAssetsResponseBodyDataResponseDataAssetInfo) SetValues(v string) *DescribeCloudSiemAssetsResponseBodyDataResponseDataAssetInfo {
	s.Values = &v
	return s
}

type DescribeCloudSiemAssetsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCloudSiemAssetsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCloudSiemAssetsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudSiemAssetsResponse) GoString() string {
	return s.String()
}

func (s *DescribeCloudSiemAssetsResponse) SetHeaders(v map[string]*string) *DescribeCloudSiemAssetsResponse {
	s.Headers = v
	return s
}

func (s *DescribeCloudSiemAssetsResponse) SetStatusCode(v int32) *DescribeCloudSiemAssetsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCloudSiemAssetsResponse) SetBody(v *DescribeCloudSiemAssetsResponseBody) *DescribeCloudSiemAssetsResponse {
	s.Body = v
	return s
}

type DescribeCloudSiemAssetsCounterRequest struct {
	// The UUID of the event.
	//
	// This parameter is required.
	//
	// example:
	//
	// 85ea4241-798f-4684-a876-65d4f0c3****
	IncidentUuid *string `json:"IncidentUuid,omitempty" xml:"IncidentUuid,omitempty"`
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
	// The ID of the account that you switch from the management account.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view. Valid values:
	//
	// - 0: the current Alibaba Cloud account
	//
	// - 1: the global account
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s DescribeCloudSiemAssetsCounterRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudSiemAssetsCounterRequest) GoString() string {
	return s.String()
}

func (s *DescribeCloudSiemAssetsCounterRequest) SetIncidentUuid(v string) *DescribeCloudSiemAssetsCounterRequest {
	s.IncidentUuid = &v
	return s
}

func (s *DescribeCloudSiemAssetsCounterRequest) SetRegionId(v string) *DescribeCloudSiemAssetsCounterRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCloudSiemAssetsCounterRequest) SetRoleFor(v int64) *DescribeCloudSiemAssetsCounterRequest {
	s.RoleFor = &v
	return s
}

func (s *DescribeCloudSiemAssetsCounterRequest) SetRoleType(v int32) *DescribeCloudSiemAssetsCounterRequest {
	s.RoleType = &v
	return s
}

type DescribeCloudSiemAssetsCounterResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	//
	// example:
	//
	// 123456
	Data []*DescribeCloudSiemAssetsCounterResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeCloudSiemAssetsCounterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudSiemAssetsCounterResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCloudSiemAssetsCounterResponseBody) SetCode(v int32) *DescribeCloudSiemAssetsCounterResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeCloudSiemAssetsCounterResponseBody) SetData(v []*DescribeCloudSiemAssetsCounterResponseBodyData) *DescribeCloudSiemAssetsCounterResponseBody {
	s.Data = v
	return s
}

func (s *DescribeCloudSiemAssetsCounterResponseBody) SetMessage(v string) *DescribeCloudSiemAssetsCounterResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeCloudSiemAssetsCounterResponseBody) SetRequestId(v string) *DescribeCloudSiemAssetsCounterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCloudSiemAssetsCounterResponseBody) SetSuccess(v bool) *DescribeCloudSiemAssetsCounterResponseBody {
	s.Success = &v
	return s
}

type DescribeCloudSiemAssetsCounterResponseBodyData struct {
	// The number of assets.
	//
	// example:
	//
	// 1
	AssetNum *int32 `json:"AssetNum,omitempty" xml:"AssetNum,omitempty"`
	// The type of the asset. Valid values:
	//
	// 	- ip
	//
	// 	- domain
	//
	// 	- url
	//
	// 	- process
	//
	// 	- file
	//
	// 	- host
	//
	// example:
	//
	// domain
	AssetType *string `json:"AssetType,omitempty" xml:"AssetType,omitempty"`
}

func (s DescribeCloudSiemAssetsCounterResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudSiemAssetsCounterResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeCloudSiemAssetsCounterResponseBodyData) SetAssetNum(v int32) *DescribeCloudSiemAssetsCounterResponseBodyData {
	s.AssetNum = &v
	return s
}

func (s *DescribeCloudSiemAssetsCounterResponseBodyData) SetAssetType(v string) *DescribeCloudSiemAssetsCounterResponseBodyData {
	s.AssetType = &v
	return s
}

type DescribeCloudSiemAssetsCounterResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCloudSiemAssetsCounterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCloudSiemAssetsCounterResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudSiemAssetsCounterResponse) GoString() string {
	return s.String()
}

func (s *DescribeCloudSiemAssetsCounterResponse) SetHeaders(v map[string]*string) *DescribeCloudSiemAssetsCounterResponse {
	s.Headers = v
	return s
}

func (s *DescribeCloudSiemAssetsCounterResponse) SetStatusCode(v int32) *DescribeCloudSiemAssetsCounterResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCloudSiemAssetsCounterResponse) SetBody(v *DescribeCloudSiemAssetsCounterResponseBody) *DescribeCloudSiemAssetsCounterResponse {
	s.Body = v
	return s
}

type DescribeCloudSiemEventDetailRequest struct {
	// The UUID of the event.
	//
	// This parameter is required.
	//
	// example:
	//
	// 85ea4241-798f-4684-a876-65d4f0c3****
	IncidentUuid *string `json:"IncidentUuid,omitempty" xml:"IncidentUuid,omitempty"`
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
	// The ID of the account that you switch from the management account.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view. Valid values:
	//
	// - 0: the current Alibaba Cloud account
	//
	// - 1: the global account
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s DescribeCloudSiemEventDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudSiemEventDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeCloudSiemEventDetailRequest) SetIncidentUuid(v string) *DescribeCloudSiemEventDetailRequest {
	s.IncidentUuid = &v
	return s
}

func (s *DescribeCloudSiemEventDetailRequest) SetRegionId(v string) *DescribeCloudSiemEventDetailRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCloudSiemEventDetailRequest) SetRoleFor(v int64) *DescribeCloudSiemEventDetailRequest {
	s.RoleFor = &v
	return s
}

func (s *DescribeCloudSiemEventDetailRequest) SetRoleType(v int32) *DescribeCloudSiemEventDetailRequest {
	s.RoleType = &v
	return s
}

type DescribeCloudSiemEventDetailResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	//
	// example:
	//
	// 123456
	Data *DescribeCloudSiemEventDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeCloudSiemEventDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudSiemEventDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCloudSiemEventDetailResponseBody) SetCode(v int32) *DescribeCloudSiemEventDetailResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeCloudSiemEventDetailResponseBody) SetData(v *DescribeCloudSiemEventDetailResponseBodyData) *DescribeCloudSiemEventDetailResponseBody {
	s.Data = v
	return s
}

func (s *DescribeCloudSiemEventDetailResponseBody) SetMessage(v string) *DescribeCloudSiemEventDetailResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeCloudSiemEventDetailResponseBody) SetRequestId(v string) *DescribeCloudSiemEventDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCloudSiemEventDetailResponseBody) SetSuccess(v bool) *DescribeCloudSiemEventDetailResponseBody {
	s.Success = &v
	return s
}

type DescribeCloudSiemEventDetailResponseBodyData struct {
	// The number of alerts that are associated with the event.
	//
	// example:
	//
	// 4
	AlertNum *int32 `json:"AlertNum,omitempty" xml:"AlertNum,omitempty"`
	// The ID of the Alibaba Cloud account to which the event belongs.
	//
	// example:
	//
	// 127608589417****
	Aliuid *int64 `json:"Aliuid,omitempty" xml:"Aliuid,omitempty"`
	// The number of assets that are associated with the event.
	//
	// example:
	//
	// 4
	AssetNum *int32 `json:"AssetNum,omitempty" xml:"AssetNum,omitempty"`
	// The tags of the ATT\\&CK attacks.
	//
	// example:
	//
	// ["T1595.002 Vulnerability Scanning"]
	AttCkLabels []*string                                                  `json:"AttCkLabels,omitempty" xml:"AttCkLabels,omitempty" type:"Repeated"`
	AttckStages []*DescribeCloudSiemEventDetailResponseBodyDataAttckStages `json:"AttckStages,omitempty" xml:"AttckStages,omitempty" type:"Repeated"`
	// The source of the alert.
	//
	// example:
	//
	// [sas,waf]
	DataSources []*string `json:"DataSources,omitempty" xml:"DataSources,omitempty" type:"Repeated"`
	// The description of the event.
	//
	// example:
	//
	// The threat event contains 13 Miner Network,1 Execute suspicious encoded commands on Linux, etc
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The description of the event in English.
	//
	// example:
	//
	// The threat event contains 13 Miner Network,1 Execute suspicious encoded commands on Linux, etc
	DescriptionEn *string `json:"DescriptionEn,omitempty" xml:"DescriptionEn,omitempty"`
	// The extended information of the event in the JSON format.
	//
	// example:
	//
	// {"event_transfer_type":"customize_rule"}
	ExtContent *string `json:"ExtContent,omitempty" xml:"ExtContent,omitempty"`
	// The time when the event occurred.
	//
	// example:
	//
	// 2021-01-06 16:37:29
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the event was last updated.
	//
	// example:
	//
	// 2021-01-06 16:37:29
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The name of the event.
	//
	// example:
	//
	// Multiple type of alerts, including Miner Network, Command line download and run malicious files, Backdoor Process, etc
	IncidentName *string `json:"IncidentName,omitempty" xml:"IncidentName,omitempty"`
	// The name of the event in English.
	//
	// example:
	//
	// Multiple type of alerts, including Miner Network, Command line download and run malicious files, Backdoor Process, etc
	IncidentNameEn *string `json:"IncidentNameEn,omitempty" xml:"IncidentNameEn,omitempty"`
	IncidentType   *string `json:"IncidentType,omitempty" xml:"IncidentType,omitempty"`
	// The UUID of the event.
	//
	// example:
	//
	// 85ea4241-798f-4684-a876-65d4f0c3****
	IncidentUuid *string `json:"IncidentUuid,omitempty" xml:"IncidentUuid,omitempty"`
	// Users associated with the event.
	ReferAccount *string `json:"ReferAccount,omitempty" xml:"ReferAccount,omitempty"`
	// The remarks of the event.
	//
	// example:
	//
	// dealed
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The status of the event. Valid values:
	//
	// 	- 0: not handled
	//
	// 	- 1: handing
	//
	// 	- 5: handling failed
	//
	// 	- 10: handled
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The risk level. Valid values:
	//
	// 	- serious: high
	//
	// 	- suspicious: medium
	//
	// 	- remind: low
	//
	// example:
	//
	// remind
	ThreatLevel *string `json:"ThreatLevel,omitempty" xml:"ThreatLevel,omitempty"`
	// The risk score of the event. The score ranges from 0 to 100. A higher score indicates a higher risk level.
	//
	// example:
	//
	// 90.2
	ThreatScore *float32 `json:"ThreatScore,omitempty" xml:"ThreatScore,omitempty"`
}

func (s DescribeCloudSiemEventDetailResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudSiemEventDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeCloudSiemEventDetailResponseBodyData) SetAlertNum(v int32) *DescribeCloudSiemEventDetailResponseBodyData {
	s.AlertNum = &v
	return s
}

func (s *DescribeCloudSiemEventDetailResponseBodyData) SetAliuid(v int64) *DescribeCloudSiemEventDetailResponseBodyData {
	s.Aliuid = &v
	return s
}

func (s *DescribeCloudSiemEventDetailResponseBodyData) SetAssetNum(v int32) *DescribeCloudSiemEventDetailResponseBodyData {
	s.AssetNum = &v
	return s
}

func (s *DescribeCloudSiemEventDetailResponseBodyData) SetAttCkLabels(v []*string) *DescribeCloudSiemEventDetailResponseBodyData {
	s.AttCkLabels = v
	return s
}

func (s *DescribeCloudSiemEventDetailResponseBodyData) SetAttckStages(v []*DescribeCloudSiemEventDetailResponseBodyDataAttckStages) *DescribeCloudSiemEventDetailResponseBodyData {
	s.AttckStages = v
	return s
}

func (s *DescribeCloudSiemEventDetailResponseBodyData) SetDataSources(v []*string) *DescribeCloudSiemEventDetailResponseBodyData {
	s.DataSources = v
	return s
}

func (s *DescribeCloudSiemEventDetailResponseBodyData) SetDescription(v string) *DescribeCloudSiemEventDetailResponseBodyData {
	s.Description = &v
	return s
}

func (s *DescribeCloudSiemEventDetailResponseBodyData) SetDescriptionEn(v string) *DescribeCloudSiemEventDetailResponseBodyData {
	s.DescriptionEn = &v
	return s
}

func (s *DescribeCloudSiemEventDetailResponseBodyData) SetExtContent(v string) *DescribeCloudSiemEventDetailResponseBodyData {
	s.ExtContent = &v
	return s
}

func (s *DescribeCloudSiemEventDetailResponseBodyData) SetGmtCreate(v string) *DescribeCloudSiemEventDetailResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *DescribeCloudSiemEventDetailResponseBodyData) SetGmtModified(v string) *DescribeCloudSiemEventDetailResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *DescribeCloudSiemEventDetailResponseBodyData) SetIncidentName(v string) *DescribeCloudSiemEventDetailResponseBodyData {
	s.IncidentName = &v
	return s
}

func (s *DescribeCloudSiemEventDetailResponseBodyData) SetIncidentNameEn(v string) *DescribeCloudSiemEventDetailResponseBodyData {
	s.IncidentNameEn = &v
	return s
}

func (s *DescribeCloudSiemEventDetailResponseBodyData) SetIncidentType(v string) *DescribeCloudSiemEventDetailResponseBodyData {
	s.IncidentType = &v
	return s
}

func (s *DescribeCloudSiemEventDetailResponseBodyData) SetIncidentUuid(v string) *DescribeCloudSiemEventDetailResponseBodyData {
	s.IncidentUuid = &v
	return s
}

func (s *DescribeCloudSiemEventDetailResponseBodyData) SetReferAccount(v string) *DescribeCloudSiemEventDetailResponseBodyData {
	s.ReferAccount = &v
	return s
}

func (s *DescribeCloudSiemEventDetailResponseBodyData) SetRemark(v string) *DescribeCloudSiemEventDetailResponseBodyData {
	s.Remark = &v
	return s
}

func (s *DescribeCloudSiemEventDetailResponseBodyData) SetRuleId(v string) *DescribeCloudSiemEventDetailResponseBodyData {
	s.RuleId = &v
	return s
}

func (s *DescribeCloudSiemEventDetailResponseBodyData) SetStatus(v int32) *DescribeCloudSiemEventDetailResponseBodyData {
	s.Status = &v
	return s
}

func (s *DescribeCloudSiemEventDetailResponseBodyData) SetThreatLevel(v string) *DescribeCloudSiemEventDetailResponseBodyData {
	s.ThreatLevel = &v
	return s
}

func (s *DescribeCloudSiemEventDetailResponseBodyData) SetThreatScore(v float32) *DescribeCloudSiemEventDetailResponseBodyData {
	s.ThreatScore = &v
	return s
}

type DescribeCloudSiemEventDetailResponseBodyDataAttckStages struct {
	AlertNum   *int32  `json:"AlertNum,omitempty" xml:"AlertNum,omitempty"`
	TacticId   *string `json:"TacticId,omitempty" xml:"TacticId,omitempty"`
	TacticName *string `json:"TacticName,omitempty" xml:"TacticName,omitempty"`
}

func (s DescribeCloudSiemEventDetailResponseBodyDataAttckStages) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudSiemEventDetailResponseBodyDataAttckStages) GoString() string {
	return s.String()
}

func (s *DescribeCloudSiemEventDetailResponseBodyDataAttckStages) SetAlertNum(v int32) *DescribeCloudSiemEventDetailResponseBodyDataAttckStages {
	s.AlertNum = &v
	return s
}

func (s *DescribeCloudSiemEventDetailResponseBodyDataAttckStages) SetTacticId(v string) *DescribeCloudSiemEventDetailResponseBodyDataAttckStages {
	s.TacticId = &v
	return s
}

func (s *DescribeCloudSiemEventDetailResponseBodyDataAttckStages) SetTacticName(v string) *DescribeCloudSiemEventDetailResponseBodyDataAttckStages {
	s.TacticName = &v
	return s
}

type DescribeCloudSiemEventDetailResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCloudSiemEventDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCloudSiemEventDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudSiemEventDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeCloudSiemEventDetailResponse) SetHeaders(v map[string]*string) *DescribeCloudSiemEventDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeCloudSiemEventDetailResponse) SetStatusCode(v int32) *DescribeCloudSiemEventDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCloudSiemEventDetailResponse) SetBody(v *DescribeCloudSiemEventDetailResponseBody) *DescribeCloudSiemEventDetailResponse {
	s.Body = v
	return s
}

type DescribeCloudSiemEventsRequest struct {
	// The ID of the asset that is associated with the event.
	//
	// example:
	//
	// 6c740667-80b2-476d-8924-2e706feb****
	AssetId *string `json:"AssetId,omitempty" xml:"AssetId,omitempty"`
	// The page number. Pages start from page 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The end of the time range to query. Unit: milliseconds.
	//
	// example:
	//
	// 1577808000000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 6c740667-80b2-476d-8924-2e706feb****
	EntityUuid *string `json:"EntityUuid,omitempty" xml:"EntityUuid,omitempty"`
	// The name of the event.
	//
	// example:
	//
	// ECS unusual log in
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// The ID of the event.
	//
	// example:
	//
	// 85ea4241-798f-4684-a876-65d4f0c3****
	IncidentUuid *string `json:"IncidentUuid,omitempty" xml:"IncidentUuid,omitempty"`
	// The sort order. Valid values:
	//
	// 	- desc: descending order
	//
	// 	- asc: ascending order
	//
	// example:
	//
	// desc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The sort field. Valid values:
	//
	// 	- GmtModified: sorts the events by creation time. This is the default value.
	//
	// 	- ThreatScore: sorts the events by risk score.
	//
	// example:
	//
	// ThreatScore
	OrderField *string `json:"OrderField,omitempty" xml:"OrderField,omitempty"`
	// The number of entries per page. Maximum value: 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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
	// The ID of the account that you switch from the management account.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view. Valid values:
	//
	// - 0: the current Alibaba Cloud account
	//
	// - 1: the global account
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// The beginning of the time range to query. Unit: milliseconds.
	//
	// example:
	//
	// 1577808000000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the event. Valid values:
	//
	// 	- 0: unhandled
	//
	// 	- 1: handling
	//
	// 	- 5: handling failed
	//
	// 	- 10: handled
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The risk levels of the events. The value is a JSON array. Valid values:
	//
	// 	- serious: high
	//
	// 	- suspicious: medium
	//
	// 	- remind: low
	//
	// example:
	//
	// ["serious","suspicious","remind"]
	ThreadLevel []*string `json:"ThreadLevel,omitempty" xml:"ThreadLevel,omitempty" type:"Repeated"`
}

func (s DescribeCloudSiemEventsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudSiemEventsRequest) GoString() string {
	return s.String()
}

func (s *DescribeCloudSiemEventsRequest) SetAssetId(v string) *DescribeCloudSiemEventsRequest {
	s.AssetId = &v
	return s
}

func (s *DescribeCloudSiemEventsRequest) SetCurrentPage(v int32) *DescribeCloudSiemEventsRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeCloudSiemEventsRequest) SetEndTime(v int64) *DescribeCloudSiemEventsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeCloudSiemEventsRequest) SetEntityUuid(v string) *DescribeCloudSiemEventsRequest {
	s.EntityUuid = &v
	return s
}

func (s *DescribeCloudSiemEventsRequest) SetEventName(v string) *DescribeCloudSiemEventsRequest {
	s.EventName = &v
	return s
}

func (s *DescribeCloudSiemEventsRequest) SetIncidentUuid(v string) *DescribeCloudSiemEventsRequest {
	s.IncidentUuid = &v
	return s
}

func (s *DescribeCloudSiemEventsRequest) SetOrder(v string) *DescribeCloudSiemEventsRequest {
	s.Order = &v
	return s
}

func (s *DescribeCloudSiemEventsRequest) SetOrderField(v string) *DescribeCloudSiemEventsRequest {
	s.OrderField = &v
	return s
}

func (s *DescribeCloudSiemEventsRequest) SetPageSize(v int32) *DescribeCloudSiemEventsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCloudSiemEventsRequest) SetRegionId(v string) *DescribeCloudSiemEventsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCloudSiemEventsRequest) SetRoleFor(v int64) *DescribeCloudSiemEventsRequest {
	s.RoleFor = &v
	return s
}

func (s *DescribeCloudSiemEventsRequest) SetRoleType(v int32) *DescribeCloudSiemEventsRequest {
	s.RoleType = &v
	return s
}

func (s *DescribeCloudSiemEventsRequest) SetStartTime(v int64) *DescribeCloudSiemEventsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeCloudSiemEventsRequest) SetStatus(v int32) *DescribeCloudSiemEventsRequest {
	s.Status = &v
	return s
}

func (s *DescribeCloudSiemEventsRequest) SetThreadLevel(v []*string) *DescribeCloudSiemEventsRequest {
	s.ThreadLevel = v
	return s
}

type DescribeCloudSiemEventsResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	//
	// example:
	//
	// 123456
	Data *DescribeCloudSiemEventsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeCloudSiemEventsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudSiemEventsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCloudSiemEventsResponseBody) SetCode(v int32) *DescribeCloudSiemEventsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeCloudSiemEventsResponseBody) SetData(v *DescribeCloudSiemEventsResponseBodyData) *DescribeCloudSiemEventsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeCloudSiemEventsResponseBody) SetMessage(v string) *DescribeCloudSiemEventsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeCloudSiemEventsResponseBody) SetRequestId(v string) *DescribeCloudSiemEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCloudSiemEventsResponseBody) SetSuccess(v bool) *DescribeCloudSiemEventsResponseBody {
	s.Success = &v
	return s
}

type DescribeCloudSiemEventsResponseBodyData struct {
	// The pagination information.
	PageInfo *DescribeCloudSiemEventsResponseBodyDataPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The detailed data.
	ResponseData []*DescribeCloudSiemEventsResponseBodyDataResponseData `json:"ResponseData,omitempty" xml:"ResponseData,omitempty" type:"Repeated"`
}

func (s DescribeCloudSiemEventsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudSiemEventsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeCloudSiemEventsResponseBodyData) SetPageInfo(v *DescribeCloudSiemEventsResponseBodyDataPageInfo) *DescribeCloudSiemEventsResponseBodyData {
	s.PageInfo = v
	return s
}

func (s *DescribeCloudSiemEventsResponseBodyData) SetResponseData(v []*DescribeCloudSiemEventsResponseBodyDataResponseData) *DescribeCloudSiemEventsResponseBodyData {
	s.ResponseData = v
	return s
}

type DescribeCloudSiemEventsResponseBodyDataPageInfo struct {
	// The current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCloudSiemEventsResponseBodyDataPageInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudSiemEventsResponseBodyDataPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeCloudSiemEventsResponseBodyDataPageInfo) SetCurrentPage(v int32) *DescribeCloudSiemEventsResponseBodyDataPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeCloudSiemEventsResponseBodyDataPageInfo) SetPageSize(v int32) *DescribeCloudSiemEventsResponseBodyDataPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeCloudSiemEventsResponseBodyDataPageInfo) SetTotalCount(v int64) *DescribeCloudSiemEventsResponseBodyDataPageInfo {
	s.TotalCount = &v
	return s
}

type DescribeCloudSiemEventsResponseBodyDataResponseData struct {
	// The number of alerts that are associated with the event.
	//
	// example:
	//
	// 4
	AlertNum *int32 `json:"AlertNum,omitempty" xml:"AlertNum,omitempty"`
	// The ID of the Alibaba Cloud account to which the event belongs.
	//
	// example:
	//
	// 127608589417****
	Aliuid *int64 `json:"Aliuid,omitempty" xml:"Aliuid,omitempty"`
	// The number of assets that are associated with the event.
	//
	// example:
	//
	// 4
	AssetNum *int32 `json:"AssetNum,omitempty" xml:"AssetNum,omitempty"`
	// The tags of the ATT\\&CK techniques.
	//
	// example:
	//
	// ["T1595.002 Vulnerability Scanning"]
	AttCkLabels []*string                                                         `json:"AttCkLabels,omitempty" xml:"AttCkLabels,omitempty" type:"Repeated"`
	AttckStages []*DescribeCloudSiemEventsResponseBodyDataResponseDataAttckStages `json:"AttckStages,omitempty" xml:"AttckStages,omitempty" type:"Repeated"`
	// The sources of the alert.
	//
	// example:
	//
	// [sas,waf]
	DataSources []*string `json:"DataSources,omitempty" xml:"DataSources,omitempty" type:"Repeated"`
	// The description of the event.
	//
	// example:
	//
	// The threat event contains 13 Miner Network,1 Execute suspicious encoded commands on Linux, etc
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The event description in English.
	//
	// example:
	//
	// The threat event contains 13 Miner Network,1 Execute suspicious encoded commands on Linux, etc
	DescriptionEn *string `json:"DescriptionEn,omitempty" xml:"DescriptionEn,omitempty"`
	// The extended event information in the JSON format.
	//
	// example:
	//
	// {"event_transfer_type":"customize_rule"}
	ExtContent *string `json:"ExtContent,omitempty" xml:"ExtContent,omitempty"`
	// The time when the event occurred.
	//
	// example:
	//
	// 2021-01-06 16:37:29
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the event was last updated.
	//
	// example:
	//
	// 2021-01-06 16:37:29
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The name of the event.
	//
	// example:
	//
	// Multiple type of alerts, including Miner Network, Command line download and run malicious files, Backdoor Process, etc
	IncidentName *string `json:"IncidentName,omitempty" xml:"IncidentName,omitempty"`
	// The event name in English.
	//
	// example:
	//
	// Multiple type of alerts, including Miner Network, Command line download and run malicious files, Backdoor Process, etc
	IncidentNameEn *string `json:"IncidentNameEn,omitempty" xml:"IncidentNameEn,omitempty"`
	IncidentType   *string `json:"IncidentType,omitempty" xml:"IncidentType,omitempty"`
	// The UUID of the event.
	//
	// example:
	//
	// 85ea4241-798f-4684-a876-65d4f0c3****
	IncidentUuid *string `json:"IncidentUuid,omitempty" xml:"IncidentUuid,omitempty"`
	// the refer account info.
	//
	// example:
	//
	// 127608589417****
	ReferAccount *string `json:"ReferAccount,omitempty" xml:"ReferAccount,omitempty"`
	// The remarks of the event.
	//
	// example:
	//
	// dealed
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The status of the event. Valid values:
	//
	// 	- 0: unhandled.
	//
	// 	- 1: handling.
	//
	// 	- 5: handling failed.
	//
	// 	- 10: handled.
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The risk level. Valid values:
	//
	// 	- serious: high.
	//
	// 	- suspicious: medium.
	//
	// 	- remind: low.
	//
	// example:
	//
	// remind
	ThreatLevel *string `json:"ThreatLevel,omitempty" xml:"ThreatLevel,omitempty"`
	// The risk score of the event. Valid values: 0 to 100. A higher value indicates a higher risk level.
	//
	// example:
	//
	// 90.2
	ThreatScore *float32 `json:"ThreatScore,omitempty" xml:"ThreatScore,omitempty"`
}

func (s DescribeCloudSiemEventsResponseBodyDataResponseData) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudSiemEventsResponseBodyDataResponseData) GoString() string {
	return s.String()
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseData) SetAlertNum(v int32) *DescribeCloudSiemEventsResponseBodyDataResponseData {
	s.AlertNum = &v
	return s
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseData) SetAliuid(v int64) *DescribeCloudSiemEventsResponseBodyDataResponseData {
	s.Aliuid = &v
	return s
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseData) SetAssetNum(v int32) *DescribeCloudSiemEventsResponseBodyDataResponseData {
	s.AssetNum = &v
	return s
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseData) SetAttCkLabels(v []*string) *DescribeCloudSiemEventsResponseBodyDataResponseData {
	s.AttCkLabels = v
	return s
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseData) SetAttckStages(v []*DescribeCloudSiemEventsResponseBodyDataResponseDataAttckStages) *DescribeCloudSiemEventsResponseBodyDataResponseData {
	s.AttckStages = v
	return s
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseData) SetDataSources(v []*string) *DescribeCloudSiemEventsResponseBodyDataResponseData {
	s.DataSources = v
	return s
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseData) SetDescription(v string) *DescribeCloudSiemEventsResponseBodyDataResponseData {
	s.Description = &v
	return s
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseData) SetDescriptionEn(v string) *DescribeCloudSiemEventsResponseBodyDataResponseData {
	s.DescriptionEn = &v
	return s
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseData) SetExtContent(v string) *DescribeCloudSiemEventsResponseBodyDataResponseData {
	s.ExtContent = &v
	return s
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseData) SetGmtCreate(v string) *DescribeCloudSiemEventsResponseBodyDataResponseData {
	s.GmtCreate = &v
	return s
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseData) SetGmtModified(v string) *DescribeCloudSiemEventsResponseBodyDataResponseData {
	s.GmtModified = &v
	return s
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseData) SetIncidentName(v string) *DescribeCloudSiemEventsResponseBodyDataResponseData {
	s.IncidentName = &v
	return s
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseData) SetIncidentNameEn(v string) *DescribeCloudSiemEventsResponseBodyDataResponseData {
	s.IncidentNameEn = &v
	return s
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseData) SetIncidentType(v string) *DescribeCloudSiemEventsResponseBodyDataResponseData {
	s.IncidentType = &v
	return s
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseData) SetIncidentUuid(v string) *DescribeCloudSiemEventsResponseBodyDataResponseData {
	s.IncidentUuid = &v
	return s
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseData) SetReferAccount(v string) *DescribeCloudSiemEventsResponseBodyDataResponseData {
	s.ReferAccount = &v
	return s
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseData) SetRemark(v string) *DescribeCloudSiemEventsResponseBodyDataResponseData {
	s.Remark = &v
	return s
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseData) SetRuleId(v string) *DescribeCloudSiemEventsResponseBodyDataResponseData {
	s.RuleId = &v
	return s
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseData) SetStatus(v int32) *DescribeCloudSiemEventsResponseBodyDataResponseData {
	s.Status = &v
	return s
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseData) SetThreatLevel(v string) *DescribeCloudSiemEventsResponseBodyDataResponseData {
	s.ThreatLevel = &v
	return s
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseData) SetThreatScore(v float32) *DescribeCloudSiemEventsResponseBodyDataResponseData {
	s.ThreatScore = &v
	return s
}

type DescribeCloudSiemEventsResponseBodyDataResponseDataAttckStages struct {
	AlertNum   *int32  `json:"AlertNum,omitempty" xml:"AlertNum,omitempty"`
	TacticId   *string `json:"TacticId,omitempty" xml:"TacticId,omitempty"`
	TacticName *string `json:"TacticName,omitempty" xml:"TacticName,omitempty"`
}

func (s DescribeCloudSiemEventsResponseBodyDataResponseDataAttckStages) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudSiemEventsResponseBodyDataResponseDataAttckStages) GoString() string {
	return s.String()
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseDataAttckStages) SetAlertNum(v int32) *DescribeCloudSiemEventsResponseBodyDataResponseDataAttckStages {
	s.AlertNum = &v
	return s
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseDataAttckStages) SetTacticId(v string) *DescribeCloudSiemEventsResponseBodyDataResponseDataAttckStages {
	s.TacticId = &v
	return s
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseDataAttckStages) SetTacticName(v string) *DescribeCloudSiemEventsResponseBodyDataResponseDataAttckStages {
	s.TacticName = &v
	return s
}

type DescribeCloudSiemEventsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCloudSiemEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCloudSiemEventsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudSiemEventsResponse) GoString() string {
	return s.String()
}

func (s *DescribeCloudSiemEventsResponse) SetHeaders(v map[string]*string) *DescribeCloudSiemEventsResponse {
	s.Headers = v
	return s
}

func (s *DescribeCloudSiemEventsResponse) SetStatusCode(v int32) *DescribeCloudSiemEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCloudSiemEventsResponse) SetBody(v *DescribeCloudSiemEventsResponseBody) *DescribeCloudSiemEventsResponse {
	s.Body = v
	return s
}

type DescribeCustomizeRuleCountRequest struct {
	// The data management center of the threat analysis feature. Specify this parameter based on the region in which your assets reside. Valid values:
	//
	// 	- cn-hangzhou: Your assets reside in regions inside China.
	//
	// 	- ap-southeast-1: Your assets reside in regions outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the account that you switch from the management account.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view. Valid values:
	//
	// - 0: the current Alibaba Cloud account
	//
	// - 1: the global account
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s DescribeCustomizeRuleCountRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCustomizeRuleCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeCustomizeRuleCountRequest) SetRegionId(v string) *DescribeCustomizeRuleCountRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCustomizeRuleCountRequest) SetRoleFor(v int64) *DescribeCustomizeRuleCountRequest {
	s.RoleFor = &v
	return s
}

func (s *DescribeCustomizeRuleCountRequest) SetRoleType(v int32) *DescribeCustomizeRuleCountRequest {
	s.RoleType = &v
	return s
}

type DescribeCustomizeRuleCountResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	//
	// example:
	//
	// 123456
	Data *DescribeCustomizeRuleCountResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeCustomizeRuleCountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCustomizeRuleCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCustomizeRuleCountResponseBody) SetCode(v int32) *DescribeCustomizeRuleCountResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeCustomizeRuleCountResponseBody) SetData(v *DescribeCustomizeRuleCountResponseBodyData) *DescribeCustomizeRuleCountResponseBody {
	s.Data = v
	return s
}

func (s *DescribeCustomizeRuleCountResponseBody) SetMessage(v string) *DescribeCustomizeRuleCountResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeCustomizeRuleCountResponseBody) SetRequestId(v string) *DescribeCustomizeRuleCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCustomizeRuleCountResponseBody) SetSuccess(v bool) *DescribeCustomizeRuleCountResponseBody {
	s.Success = &v
	return s
}

type DescribeCustomizeRuleCountResponseBodyData struct {
	// 同类聚合规则数。
	//
	// example:
	//
	// 3
	AggregationRuleNum *int32 `json:"AggregationRuleNum,omitempty" xml:"AggregationRuleNum,omitempty"`
	// 自定义规则数。
	//
	// example:
	//
	// 10
	CustomizeRuleNum *int32 `json:"CustomizeRuleNum,omitempty" xml:"CustomizeRuleNum,omitempty"`
	// 专家规则数。
	//
	// example:
	//
	// 7
	ExpertRuleNum *int32 `json:"ExpertRuleNum,omitempty" xml:"ExpertRuleNum,omitempty"`
	// 图计算规则数。
	//
	// example:
	//
	// 2
	GraphComputingRuleNum *int32 `json:"GraphComputingRuleNum,omitempty" xml:"GraphComputingRuleNum,omitempty"`
	// The number of rules that are used to identify high-risk threats.
	//
	// example:
	//
	// 12
	HighRuleNum *int32 `json:"HighRuleNum,omitempty" xml:"HighRuleNum,omitempty"`
	// The total number of rules.
	//
	// example:
	//
	// 20
	InUseRuleNum *int32 `json:"InUseRuleNum,omitempty" xml:"InUseRuleNum,omitempty"`
	// The number of rules that are used to identify low-risk threats.
	//
	// example:
	//
	// 3
	LowRuleNum *int32 `json:"LowRuleNum,omitempty" xml:"LowRuleNum,omitempty"`
	// The number of rules that are used to identify medium-risk threats.
	//
	// example:
	//
	// 5
	MediumRuleNum *int32 `json:"MediumRuleNum,omitempty" xml:"MediumRuleNum,omitempty"`
	// 预定义规则数。
	//
	// example:
	//
	// 10
	PredefinedRuleNum *int32 `json:"PredefinedRuleNum,omitempty" xml:"PredefinedRuleNum,omitempty"`
	// 告警透传规则数。
	//
	// example:
	//
	// 3
	SingleAlertRuleNum *int32 `json:"SingleAlertRuleNum,omitempty" xml:"SingleAlertRuleNum,omitempty"`
	// 总规则数。
	//
	// example:
	//
	// 10
	TotalRuleNum *int32 `json:"TotalRuleNum,omitempty" xml:"TotalRuleNum,omitempty"`
	// 不产生事件规则数。
	//
	// example:
	//
	// 3
	UnEventRuleNum *int32 `json:"UnEventRuleNum,omitempty" xml:"UnEventRuleNum,omitempty"`
}

func (s DescribeCustomizeRuleCountResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeCustomizeRuleCountResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeCustomizeRuleCountResponseBodyData) SetAggregationRuleNum(v int32) *DescribeCustomizeRuleCountResponseBodyData {
	s.AggregationRuleNum = &v
	return s
}

func (s *DescribeCustomizeRuleCountResponseBodyData) SetCustomizeRuleNum(v int32) *DescribeCustomizeRuleCountResponseBodyData {
	s.CustomizeRuleNum = &v
	return s
}

func (s *DescribeCustomizeRuleCountResponseBodyData) SetExpertRuleNum(v int32) *DescribeCustomizeRuleCountResponseBodyData {
	s.ExpertRuleNum = &v
	return s
}

func (s *DescribeCustomizeRuleCountResponseBodyData) SetGraphComputingRuleNum(v int32) *DescribeCustomizeRuleCountResponseBodyData {
	s.GraphComputingRuleNum = &v
	return s
}

func (s *DescribeCustomizeRuleCountResponseBodyData) SetHighRuleNum(v int32) *DescribeCustomizeRuleCountResponseBodyData {
	s.HighRuleNum = &v
	return s
}

func (s *DescribeCustomizeRuleCountResponseBodyData) SetInUseRuleNum(v int32) *DescribeCustomizeRuleCountResponseBodyData {
	s.InUseRuleNum = &v
	return s
}

func (s *DescribeCustomizeRuleCountResponseBodyData) SetLowRuleNum(v int32) *DescribeCustomizeRuleCountResponseBodyData {
	s.LowRuleNum = &v
	return s
}

func (s *DescribeCustomizeRuleCountResponseBodyData) SetMediumRuleNum(v int32) *DescribeCustomizeRuleCountResponseBodyData {
	s.MediumRuleNum = &v
	return s
}

func (s *DescribeCustomizeRuleCountResponseBodyData) SetPredefinedRuleNum(v int32) *DescribeCustomizeRuleCountResponseBodyData {
	s.PredefinedRuleNum = &v
	return s
}

func (s *DescribeCustomizeRuleCountResponseBodyData) SetSingleAlertRuleNum(v int32) *DescribeCustomizeRuleCountResponseBodyData {
	s.SingleAlertRuleNum = &v
	return s
}

func (s *DescribeCustomizeRuleCountResponseBodyData) SetTotalRuleNum(v int32) *DescribeCustomizeRuleCountResponseBodyData {
	s.TotalRuleNum = &v
	return s
}

func (s *DescribeCustomizeRuleCountResponseBodyData) SetUnEventRuleNum(v int32) *DescribeCustomizeRuleCountResponseBodyData {
	s.UnEventRuleNum = &v
	return s
}

type DescribeCustomizeRuleCountResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCustomizeRuleCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCustomizeRuleCountResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCustomizeRuleCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeCustomizeRuleCountResponse) SetHeaders(v map[string]*string) *DescribeCustomizeRuleCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeCustomizeRuleCountResponse) SetStatusCode(v int32) *DescribeCustomizeRuleCountResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCustomizeRuleCountResponse) SetBody(v *DescribeCustomizeRuleCountResponseBody) *DescribeCustomizeRuleCountResponse {
	s.Body = v
	return s
}

type DescribeCustomizeRuleTestRequest struct {
	// The ID of the rule.
	//
	// example:
	//
	// 123456789
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
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
	// The ID of the account that you switch from the management account.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view. Valid values:
	//
	// - 0: the current Alibaba Cloud account
	//
	// - 1: the global account
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s DescribeCustomizeRuleTestRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCustomizeRuleTestRequest) GoString() string {
	return s.String()
}

func (s *DescribeCustomizeRuleTestRequest) SetId(v int64) *DescribeCustomizeRuleTestRequest {
	s.Id = &v
	return s
}

func (s *DescribeCustomizeRuleTestRequest) SetRegionId(v string) *DescribeCustomizeRuleTestRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCustomizeRuleTestRequest) SetRoleFor(v int64) *DescribeCustomizeRuleTestRequest {
	s.RoleFor = &v
	return s
}

func (s *DescribeCustomizeRuleTestRequest) SetRoleType(v int32) *DescribeCustomizeRuleTestRequest {
	s.RoleType = &v
	return s
}

type DescribeCustomizeRuleTestResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	//
	// example:
	//
	// 123456
	Data *DescribeCustomizeRuleTestResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeCustomizeRuleTestResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCustomizeRuleTestResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCustomizeRuleTestResponseBody) SetCode(v int32) *DescribeCustomizeRuleTestResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeCustomizeRuleTestResponseBody) SetData(v *DescribeCustomizeRuleTestResponseBodyData) *DescribeCustomizeRuleTestResponseBody {
	s.Data = v
	return s
}

func (s *DescribeCustomizeRuleTestResponseBody) SetMessage(v string) *DescribeCustomizeRuleTestResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeCustomizeRuleTestResponseBody) SetRequestId(v string) *DescribeCustomizeRuleTestResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCustomizeRuleTestResponseBody) SetSuccess(v bool) *DescribeCustomizeRuleTestResponseBody {
	s.Success = &v
	return s
}

type DescribeCustomizeRuleTestResponseBodyData struct {
	// The ID of the rule.
	//
	// example:
	//
	// 123456789
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The historical data that is used in the simulation test.
	//
	// example:
	//
	// [{"key1":"value1","key2":"value2","key3":"value3","key4":"value4","key5":"value5"}]
	SimulateData *string `json:"SimulateData,omitempty" xml:"SimulateData,omitempty"`
	// The status of the rule. Valid values:
	//
	// 	- 0: The rule is in the initial state.
	//
	// 	- 10: The simulation data is tested.
	//
	// 	- 15: The business data is being tested.
	//
	// 	- 20: The business data test ends.
	//
	// 	- 100: The rule takes effect.
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeCustomizeRuleTestResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeCustomizeRuleTestResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeCustomizeRuleTestResponseBodyData) SetId(v int64) *DescribeCustomizeRuleTestResponseBodyData {
	s.Id = &v
	return s
}

func (s *DescribeCustomizeRuleTestResponseBodyData) SetSimulateData(v string) *DescribeCustomizeRuleTestResponseBodyData {
	s.SimulateData = &v
	return s
}

func (s *DescribeCustomizeRuleTestResponseBodyData) SetStatus(v int32) *DescribeCustomizeRuleTestResponseBodyData {
	s.Status = &v
	return s
}

type DescribeCustomizeRuleTestResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCustomizeRuleTestResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCustomizeRuleTestResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCustomizeRuleTestResponse) GoString() string {
	return s.String()
}

func (s *DescribeCustomizeRuleTestResponse) SetHeaders(v map[string]*string) *DescribeCustomizeRuleTestResponse {
	s.Headers = v
	return s
}

func (s *DescribeCustomizeRuleTestResponse) SetStatusCode(v int32) *DescribeCustomizeRuleTestResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCustomizeRuleTestResponse) SetBody(v *DescribeCustomizeRuleTestResponseBody) *DescribeCustomizeRuleTestResponse {
	s.Body = v
	return s
}

type DescribeCustomizeRuleTestHistogramRequest struct {
	// The ID of the rule.
	//
	// example:
	//
	// 123456789
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
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
	// The ID of the account that you switch from the management account.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view. Valid values:
	//
	// - 0: the current Alibaba Cloud account
	//
	// - 1: the global account
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s DescribeCustomizeRuleTestHistogramRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCustomizeRuleTestHistogramRequest) GoString() string {
	return s.String()
}

func (s *DescribeCustomizeRuleTestHistogramRequest) SetId(v int64) *DescribeCustomizeRuleTestHistogramRequest {
	s.Id = &v
	return s
}

func (s *DescribeCustomizeRuleTestHistogramRequest) SetRegionId(v string) *DescribeCustomizeRuleTestHistogramRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCustomizeRuleTestHistogramRequest) SetRoleFor(v int64) *DescribeCustomizeRuleTestHistogramRequest {
	s.RoleFor = &v
	return s
}

func (s *DescribeCustomizeRuleTestHistogramRequest) SetRoleType(v int32) *DescribeCustomizeRuleTestHistogramRequest {
	s.RoleType = &v
	return s
}

type DescribeCustomizeRuleTestHistogramResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The return value for the request.
	//
	// example:
	//
	// 123456
	Data []*DescribeCustomizeRuleTestHistogramResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeCustomizeRuleTestHistogramResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCustomizeRuleTestHistogramResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCustomizeRuleTestHistogramResponseBody) SetCode(v int32) *DescribeCustomizeRuleTestHistogramResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeCustomizeRuleTestHistogramResponseBody) SetData(v []*DescribeCustomizeRuleTestHistogramResponseBodyData) *DescribeCustomizeRuleTestHistogramResponseBody {
	s.Data = v
	return s
}

func (s *DescribeCustomizeRuleTestHistogramResponseBody) SetMessage(v string) *DescribeCustomizeRuleTestHistogramResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeCustomizeRuleTestHistogramResponseBody) SetRequestId(v string) *DescribeCustomizeRuleTestHistogramResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCustomizeRuleTestHistogramResponseBody) SetSuccess(v bool) *DescribeCustomizeRuleTestHistogramResponseBody {
	s.Success = &v
	return s
}

type DescribeCustomizeRuleTestHistogramResponseBodyData struct {
	// The number of alerts that are generated in the query time range.
	//
	// example:
	//
	// 125
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The start of the time range for querying alerts. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1599897188
	From *int64 `json:"From,omitempty" xml:"From,omitempty"`
	// The end of the time range for querying alerts. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1599997188
	To *int64 `json:"To,omitempty" xml:"To,omitempty"`
}

func (s DescribeCustomizeRuleTestHistogramResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeCustomizeRuleTestHistogramResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeCustomizeRuleTestHistogramResponseBodyData) SetCount(v int64) *DescribeCustomizeRuleTestHistogramResponseBodyData {
	s.Count = &v
	return s
}

func (s *DescribeCustomizeRuleTestHistogramResponseBodyData) SetFrom(v int64) *DescribeCustomizeRuleTestHistogramResponseBodyData {
	s.From = &v
	return s
}

func (s *DescribeCustomizeRuleTestHistogramResponseBodyData) SetTo(v int64) *DescribeCustomizeRuleTestHistogramResponseBodyData {
	s.To = &v
	return s
}

type DescribeCustomizeRuleTestHistogramResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCustomizeRuleTestHistogramResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCustomizeRuleTestHistogramResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCustomizeRuleTestHistogramResponse) GoString() string {
	return s.String()
}

func (s *DescribeCustomizeRuleTestHistogramResponse) SetHeaders(v map[string]*string) *DescribeCustomizeRuleTestHistogramResponse {
	s.Headers = v
	return s
}

func (s *DescribeCustomizeRuleTestHistogramResponse) SetStatusCode(v int32) *DescribeCustomizeRuleTestHistogramResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCustomizeRuleTestHistogramResponse) SetBody(v *DescribeCustomizeRuleTestHistogramResponseBody) *DescribeCustomizeRuleTestHistogramResponse {
	s.Body = v
	return s
}

type DescribeDataSourceInstanceRequest struct {
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

func (s DescribeDataSourceInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataSourceInstanceRequest) GoString() string {
	return s.String()
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

type DescribeDataSourceInstanceResponseBody struct {
	// The data returned.
	Data *DescribeDataSourceInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDataSourceInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataSourceInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDataSourceInstanceResponseBody) SetData(v *DescribeDataSourceInstanceResponseBodyData) *DescribeDataSourceInstanceResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDataSourceInstanceResponseBody) SetRequestId(v string) *DescribeDataSourceInstanceResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDataSourceInstanceResponseBodyData struct {
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
	// The parameters of the data source.
	DataSourceInstanceParams []*DescribeDataSourceInstanceResponseBodyDataDataSourceInstanceParams `json:"DataSourceInstanceParams,omitempty" xml:"DataSourceInstanceParams,omitempty" type:"Repeated"`
}

func (s DescribeDataSourceInstanceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataSourceInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDataSourceInstanceResponseBodyData) SetAccountId(v string) *DescribeDataSourceInstanceResponseBodyData {
	s.AccountId = &v
	return s
}

func (s *DescribeDataSourceInstanceResponseBodyData) SetCloudCode(v string) *DescribeDataSourceInstanceResponseBodyData {
	s.CloudCode = &v
	return s
}

func (s *DescribeDataSourceInstanceResponseBodyData) SetDataSourceInstanceId(v string) *DescribeDataSourceInstanceResponseBodyData {
	s.DataSourceInstanceId = &v
	return s
}

func (s *DescribeDataSourceInstanceResponseBodyData) SetDataSourceInstanceParams(v []*DescribeDataSourceInstanceResponseBodyDataDataSourceInstanceParams) *DescribeDataSourceInstanceResponseBodyData {
	s.DataSourceInstanceParams = v
	return s
}

type DescribeDataSourceInstanceResponseBodyDataDataSourceInstanceParams struct {
	// The code of the parameter.
	//
	// example:
	//
	// region_code
	ParaCode *string `json:"ParaCode,omitempty" xml:"ParaCode,omitempty"`
	// The value of the parameter.
	//
	// example:
	//
	// ap-guangzhou
	ParaValue *string `json:"ParaValue,omitempty" xml:"ParaValue,omitempty"`
}

func (s DescribeDataSourceInstanceResponseBodyDataDataSourceInstanceParams) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataSourceInstanceResponseBodyDataDataSourceInstanceParams) GoString() string {
	return s.String()
}

func (s *DescribeDataSourceInstanceResponseBodyDataDataSourceInstanceParams) SetParaCode(v string) *DescribeDataSourceInstanceResponseBodyDataDataSourceInstanceParams {
	s.ParaCode = &v
	return s
}

func (s *DescribeDataSourceInstanceResponseBodyDataDataSourceInstanceParams) SetParaValue(v string) *DescribeDataSourceInstanceResponseBodyDataDataSourceInstanceParams {
	s.ParaValue = &v
	return s
}

type DescribeDataSourceInstanceResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDataSourceInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDataSourceInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataSourceInstanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeDataSourceInstanceResponse) SetHeaders(v map[string]*string) *DescribeDataSourceInstanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeDataSourceInstanceResponse) SetStatusCode(v int32) *DescribeDataSourceInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDataSourceInstanceResponse) SetBody(v *DescribeDataSourceInstanceResponseBody) *DescribeDataSourceInstanceResponse {
	s.Body = v
	return s
}

type DescribeDataSourceParametersRequest struct {
	// The code of the cloud service provider.
	//
	// Valid values:
	//
	// 	- qcloud
	//
	// 	- hcloud
	//
	// 	- aliyun
	//
	// This parameter is required.
	//
	// example:
	//
	// hcloud
	CloudCode *string `json:"CloudCode,omitempty" xml:"CloudCode,omitempty"`
	// The type of the data source. Valid values:
	//
	// 	- **ckafka**: Tencent Cloud TDMQ for CKafka
	//
	// 	- **obs**: Huawei Cloud Object Storage Service (OBS)
	//
	// 	- **wafApi**: download API of Tencent Cloud Web Application Firewall (WAF)
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

func (s DescribeDataSourceParametersRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataSourceParametersRequest) GoString() string {
	return s.String()
}

func (s *DescribeDataSourceParametersRequest) SetCloudCode(v string) *DescribeDataSourceParametersRequest {
	s.CloudCode = &v
	return s
}

func (s *DescribeDataSourceParametersRequest) SetDataSourceType(v string) *DescribeDataSourceParametersRequest {
	s.DataSourceType = &v
	return s
}

func (s *DescribeDataSourceParametersRequest) SetRegionId(v string) *DescribeDataSourceParametersRequest {
	s.RegionId = &v
	return s
}

type DescribeDataSourceParametersResponseBody struct {
	// The data returned.
	Data []*DescribeDataSourceParametersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDataSourceParametersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataSourceParametersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDataSourceParametersResponseBody) SetData(v []*DescribeDataSourceParametersResponseBodyData) *DescribeDataSourceParametersResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDataSourceParametersResponseBody) SetRequestId(v string) *DescribeDataSourceParametersResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDataSourceParametersResponseBodyData struct {
	// Indicates whether the edit operation is supported. Valid values:
	//
	// 	- **0**
	//
	// 	- **1**
	//
	// example:
	//
	// wafApi
	CanEditted *int32 `json:"CanEditted,omitempty" xml:"CanEditted,omitempty"`
	// The code of the cloud service provider. Valid values:
	//
	// 	- **qcloud**: Tencent Cloud
	//
	// 	- **aliyun**: Alibaba Cloud
	//
	// 	- **hcloud**: Huawei Cloud
	//
	// example:
	//
	// hcloud
	CloudCode *string `json:"CloudCode,omitempty" xml:"CloudCode,omitempty"`
	// The type of the data source. Valid values:
	//
	// 	- **obs**: Huawei Cloud Object Storage Service (OBS)
	//
	// 	- **wafApi**: download API of Tencent Cloud Web Application Firewall (WAF)
	//
	// 	- **ckafka**: Tencent Cloud TDMQ for CKafka
	//
	// example:
	//
	// obs
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	// The default value of the parameter.
	//
	// example:
	//
	// wafApi
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	// Indicates whether the modification operation is forbidden. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// wafApi
	Disabled *bool `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	// The method that is used to check the parameter format.
	//
	// example:
	//
	// email
	FormatCheck *string `json:"FormatCheck,omitempty" xml:"FormatCheck,omitempty"`
	// The additional information.
	//
	// example:
	//
	// obs docment
	Hit *string `json:"Hit,omitempty" xml:"Hit,omitempty"`
	// The code of the parameter.
	//
	// example:
	//
	// region_code
	ParaCode *string `json:"ParaCode,omitempty" xml:"ParaCode,omitempty"`
	// The parameter level. Valid values:
	//
	// 	- **1**: the parameters of the data source
	//
	// 	- **2**: the parameters of the log
	//
	// example:
	//
	// 1
	ParaLevel *int32 `json:"ParaLevel,omitempty" xml:"ParaLevel,omitempty"`
	// The name of the parameter.
	//
	// example:
	//
	// region local
	ParaName *string `json:"ParaName,omitempty" xml:"ParaName,omitempty"`
	// The data type of the parameter.
	//
	// example:
	//
	// string
	ParaType *string `json:"ParaType,omitempty" xml:"ParaType,omitempty"`
	// The value of the parameter.
	ParamValue []*DescribeDataSourceParametersResponseBodyDataParamValue `json:"ParamValue,omitempty" xml:"ParamValue,omitempty" type:"Repeated"`
	// Indicates whether the parameter is required. Valid values:
	//
	// 	- **1**: required
	//
	// 	- **0**: optional
	//
	// example:
	//
	// string
	Required *int32 `json:"Required,omitempty" xml:"Required,omitempty"`
	// The note for the parameter value.
	//
	// example:
	//
	// obs bucket name
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s DescribeDataSourceParametersResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataSourceParametersResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDataSourceParametersResponseBodyData) SetCanEditted(v int32) *DescribeDataSourceParametersResponseBodyData {
	s.CanEditted = &v
	return s
}

func (s *DescribeDataSourceParametersResponseBodyData) SetCloudCode(v string) *DescribeDataSourceParametersResponseBodyData {
	s.CloudCode = &v
	return s
}

func (s *DescribeDataSourceParametersResponseBodyData) SetDataSourceType(v string) *DescribeDataSourceParametersResponseBodyData {
	s.DataSourceType = &v
	return s
}

func (s *DescribeDataSourceParametersResponseBodyData) SetDefaultValue(v string) *DescribeDataSourceParametersResponseBodyData {
	s.DefaultValue = &v
	return s
}

func (s *DescribeDataSourceParametersResponseBodyData) SetDisabled(v bool) *DescribeDataSourceParametersResponseBodyData {
	s.Disabled = &v
	return s
}

func (s *DescribeDataSourceParametersResponseBodyData) SetFormatCheck(v string) *DescribeDataSourceParametersResponseBodyData {
	s.FormatCheck = &v
	return s
}

func (s *DescribeDataSourceParametersResponseBodyData) SetHit(v string) *DescribeDataSourceParametersResponseBodyData {
	s.Hit = &v
	return s
}

func (s *DescribeDataSourceParametersResponseBodyData) SetParaCode(v string) *DescribeDataSourceParametersResponseBodyData {
	s.ParaCode = &v
	return s
}

func (s *DescribeDataSourceParametersResponseBodyData) SetParaLevel(v int32) *DescribeDataSourceParametersResponseBodyData {
	s.ParaLevel = &v
	return s
}

func (s *DescribeDataSourceParametersResponseBodyData) SetParaName(v string) *DescribeDataSourceParametersResponseBodyData {
	s.ParaName = &v
	return s
}

func (s *DescribeDataSourceParametersResponseBodyData) SetParaType(v string) *DescribeDataSourceParametersResponseBodyData {
	s.ParaType = &v
	return s
}

func (s *DescribeDataSourceParametersResponseBodyData) SetParamValue(v []*DescribeDataSourceParametersResponseBodyDataParamValue) *DescribeDataSourceParametersResponseBodyData {
	s.ParamValue = v
	return s
}

func (s *DescribeDataSourceParametersResponseBodyData) SetRequired(v int32) *DescribeDataSourceParametersResponseBodyData {
	s.Required = &v
	return s
}

func (s *DescribeDataSourceParametersResponseBodyData) SetTitle(v string) *DescribeDataSourceParametersResponseBodyData {
	s.Title = &v
	return s
}

type DescribeDataSourceParametersResponseBodyDataParamValue struct {
	// The display value.
	//
	// example:
	//
	// guangzhou
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The actual value.
	//
	// example:
	//
	// ap-guangzhou
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDataSourceParametersResponseBodyDataParamValue) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataSourceParametersResponseBodyDataParamValue) GoString() string {
	return s.String()
}

func (s *DescribeDataSourceParametersResponseBodyDataParamValue) SetLabel(v string) *DescribeDataSourceParametersResponseBodyDataParamValue {
	s.Label = &v
	return s
}

func (s *DescribeDataSourceParametersResponseBodyDataParamValue) SetValue(v string) *DescribeDataSourceParametersResponseBodyDataParamValue {
	s.Value = &v
	return s
}

type DescribeDataSourceParametersResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDataSourceParametersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDataSourceParametersResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataSourceParametersResponse) GoString() string {
	return s.String()
}

func (s *DescribeDataSourceParametersResponse) SetHeaders(v map[string]*string) *DescribeDataSourceParametersResponse {
	s.Headers = v
	return s
}

func (s *DescribeDataSourceParametersResponse) SetStatusCode(v int32) *DescribeDataSourceParametersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDataSourceParametersResponse) SetBody(v *DescribeDataSourceParametersResponseBody) *DescribeDataSourceParametersResponse {
	s.Body = v
	return s
}

type DescribeDisposeAndPlaybookRequest struct {
	// The page number. Pages start from page 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The entity type. Valid values:
	//
	// 	- ip
	//
	// 	- process
	//
	// 	- file
	//
	// example:
	//
	// ip
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// example:
	//
	// 85ea4241-798f-4684-a876-65d4f0c3****
	EntityUuid *string `json:"EntityUuid,omitempty" xml:"EntityUuid,omitempty"`
	// The UUID of the event.
	//
	// example:
	//
	// 85ea4241-798f-4684-a876-65d4f0c3****
	IncidentUuid *string `json:"IncidentUuid,omitempty" xml:"IncidentUuid,omitempty"`
	// The number of entries to return on each page. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The data management center of the threat analysis feature. Specify this parameter based on the region in which your assets reside. Valid values:
	//
	// 	- cn-hangzhou: Your assets reside in regions inside China.
	//
	// 	- ap-southeast-1: Your assets reside in regions outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the account that you switch from the management account.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view. Valid values:
	//
	// - 0: the current Alibaba Cloud account
	//
	// - 1: the global account
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s DescribeDisposeAndPlaybookRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDisposeAndPlaybookRequest) GoString() string {
	return s.String()
}

func (s *DescribeDisposeAndPlaybookRequest) SetCurrentPage(v int32) *DescribeDisposeAndPlaybookRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeDisposeAndPlaybookRequest) SetEntityType(v string) *DescribeDisposeAndPlaybookRequest {
	s.EntityType = &v
	return s
}

func (s *DescribeDisposeAndPlaybookRequest) SetEntityUuid(v string) *DescribeDisposeAndPlaybookRequest {
	s.EntityUuid = &v
	return s
}

func (s *DescribeDisposeAndPlaybookRequest) SetIncidentUuid(v string) *DescribeDisposeAndPlaybookRequest {
	s.IncidentUuid = &v
	return s
}

func (s *DescribeDisposeAndPlaybookRequest) SetPageSize(v int32) *DescribeDisposeAndPlaybookRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDisposeAndPlaybookRequest) SetRegionId(v string) *DescribeDisposeAndPlaybookRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDisposeAndPlaybookRequest) SetRoleFor(v int64) *DescribeDisposeAndPlaybookRequest {
	s.RoleFor = &v
	return s
}

func (s *DescribeDisposeAndPlaybookRequest) SetRoleType(v int32) *DescribeDisposeAndPlaybookRequest {
	s.RoleType = &v
	return s
}

type DescribeDisposeAndPlaybookResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	//
	// example:
	//
	// 123456
	Data *DescribeDisposeAndPlaybookResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDisposeAndPlaybookResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDisposeAndPlaybookResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDisposeAndPlaybookResponseBody) SetCode(v int32) *DescribeDisposeAndPlaybookResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeDisposeAndPlaybookResponseBody) SetData(v *DescribeDisposeAndPlaybookResponseBodyData) *DescribeDisposeAndPlaybookResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDisposeAndPlaybookResponseBody) SetMessage(v string) *DescribeDisposeAndPlaybookResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeDisposeAndPlaybookResponseBody) SetRequestId(v string) *DescribeDisposeAndPlaybookResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDisposeAndPlaybookResponseBody) SetSuccess(v bool) *DescribeDisposeAndPlaybookResponseBody {
	s.Success = &v
	return s
}

type DescribeDisposeAndPlaybookResponseBodyData struct {
	// The pagination information.
	PageInfo *DescribeDisposeAndPlaybookResponseBodyDataPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The detailed data.
	ResponseData []*DescribeDisposeAndPlaybookResponseBodyDataResponseData `json:"ResponseData,omitempty" xml:"ResponseData,omitempty" type:"Repeated"`
}

func (s DescribeDisposeAndPlaybookResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDisposeAndPlaybookResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDisposeAndPlaybookResponseBodyData) SetPageInfo(v *DescribeDisposeAndPlaybookResponseBodyDataPageInfo) *DescribeDisposeAndPlaybookResponseBodyData {
	s.PageInfo = v
	return s
}

func (s *DescribeDisposeAndPlaybookResponseBodyData) SetResponseData(v []*DescribeDisposeAndPlaybookResponseBodyDataResponseData) *DescribeDisposeAndPlaybookResponseBodyData {
	s.ResponseData = v
	return s
}

type DescribeDisposeAndPlaybookResponseBodyDataPageInfo struct {
	// The current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDisposeAndPlaybookResponseBodyDataPageInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeDisposeAndPlaybookResponseBodyDataPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeDisposeAndPlaybookResponseBodyDataPageInfo) SetCurrentPage(v int32) *DescribeDisposeAndPlaybookResponseBodyDataPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeDisposeAndPlaybookResponseBodyDataPageInfo) SetPageSize(v int32) *DescribeDisposeAndPlaybookResponseBodyDataPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeDisposeAndPlaybookResponseBodyDataPageInfo) SetTotalCount(v int64) *DescribeDisposeAndPlaybookResponseBodyDataPageInfo {
	s.TotalCount = &v
	return s
}

type DescribeDisposeAndPlaybookResponseBodyDataResponseData struct {
	// The number of alerts that are associated with the entity.
	//
	// example:
	//
	// 1
	AlertNum *int32 `json:"AlertNum,omitempty" xml:"AlertNum,omitempty"`
	// The object for handling.
	//
	// example:
	//
	// 192.168.1.1
	Dispose *string `json:"Dispose,omitempty" xml:"Dispose,omitempty"`
	// The entity ID
	//
	// example:
	//
	// 12345
	EntityId *int64 `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The entity information.
	//
	// example:
	//
	// {"file_path": "c:/www/leixi.jsp","file_hash": "aa0ca926ad948cd820e0a3d9a18c09d0","host_uuid": "efed2cf7-0b77-45d9-a97b-d2cf246bcbb3","malware_type": "${aliyun.siem.sas.alert_tag.webshell}","host_name": "launch-advisor-20230531"}
	EntityInfo map[string]interface{} `json:"EntityInfo,omitempty" xml:"EntityInfo,omitempty"`
	// example:
	//
	// ip
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// The key-value pairs each of which consists of opcode and oplevel.
	//
	// example:
	//
	// 12345
	OpcodeMap map[string]*string `json:"OpcodeMap,omitempty" xml:"OpcodeMap,omitempty"`
	// The codes of the playbooks that are recommended for entity handling.
	//
	// example:
	//
	// [1,3]
	OpcodeSet []*string `json:"OpcodeSet,omitempty" xml:"OpcodeSet,omitempty" type:"Repeated"`
	// The playbooks that can handle the entity.
	//
	// example:
	//
	// [{"name":"云安全中心-云服务器安全","code":"1"}]
	PlaybookList []*DescribeDisposeAndPlaybookResponseBodyDataResponseDataPlaybookList `json:"PlaybookList,omitempty" xml:"PlaybookList,omitempty" type:"Repeated"`
	// The IDs of the users who can handle objects.
	//
	// example:
	//
	// 176618589410****
	Scope []interface{} `json:"Scope,omitempty" xml:"Scope,omitempty" type:"Repeated"`
}

func (s DescribeDisposeAndPlaybookResponseBodyDataResponseData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDisposeAndPlaybookResponseBodyDataResponseData) GoString() string {
	return s.String()
}

func (s *DescribeDisposeAndPlaybookResponseBodyDataResponseData) SetAlertNum(v int32) *DescribeDisposeAndPlaybookResponseBodyDataResponseData {
	s.AlertNum = &v
	return s
}

func (s *DescribeDisposeAndPlaybookResponseBodyDataResponseData) SetDispose(v string) *DescribeDisposeAndPlaybookResponseBodyDataResponseData {
	s.Dispose = &v
	return s
}

func (s *DescribeDisposeAndPlaybookResponseBodyDataResponseData) SetEntityId(v int64) *DescribeDisposeAndPlaybookResponseBodyDataResponseData {
	s.EntityId = &v
	return s
}

func (s *DescribeDisposeAndPlaybookResponseBodyDataResponseData) SetEntityInfo(v map[string]interface{}) *DescribeDisposeAndPlaybookResponseBodyDataResponseData {
	s.EntityInfo = v
	return s
}

func (s *DescribeDisposeAndPlaybookResponseBodyDataResponseData) SetEntityType(v string) *DescribeDisposeAndPlaybookResponseBodyDataResponseData {
	s.EntityType = &v
	return s
}

func (s *DescribeDisposeAndPlaybookResponseBodyDataResponseData) SetOpcodeMap(v map[string]*string) *DescribeDisposeAndPlaybookResponseBodyDataResponseData {
	s.OpcodeMap = v
	return s
}

func (s *DescribeDisposeAndPlaybookResponseBodyDataResponseData) SetOpcodeSet(v []*string) *DescribeDisposeAndPlaybookResponseBodyDataResponseData {
	s.OpcodeSet = v
	return s
}

func (s *DescribeDisposeAndPlaybookResponseBodyDataResponseData) SetPlaybookList(v []*DescribeDisposeAndPlaybookResponseBodyDataResponseDataPlaybookList) *DescribeDisposeAndPlaybookResponseBodyDataResponseData {
	s.PlaybookList = v
	return s
}

func (s *DescribeDisposeAndPlaybookResponseBodyDataResponseData) SetScope(v []interface{}) *DescribeDisposeAndPlaybookResponseBodyDataResponseData {
	s.Scope = v
	return s
}

type DescribeDisposeAndPlaybookResponseBodyDataResponseDataPlaybookList struct {
	Available *string `json:"Available,omitempty" xml:"Available,omitempty"`
	// The playbook description.
	//
	// example:
	//
	// WafBlockIP
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The display name of the playbook.
	//
	// example:
	//
	// WafBlockIP
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The playbook name, which is the unique identifier of the playbook.
	//
	// example:
	//
	// kill_process_isolate_file
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The opcode of the playbook, which corresponds to the opcode of the playbook recommended for entity handling.
	//
	// example:
	//
	// 7
	OpCode *string `json:"OpCode,omitempty" xml:"OpCode,omitempty"`
	// Indicates whether quick event handling is selected by default. Valid values:
	//
	// 	- 2: Quick event handling is selected.
	//
	// 	- 1: Quick event handling is displayed but not selected.
	//
	// example:
	//
	// 2
	OpLevel *string `json:"OpLevel,omitempty" xml:"OpLevel,omitempty"`
	// The playbook parameters and the corresponding properties.
	ParamConfig []interface{} `json:"ParamConfig,omitempty" xml:"ParamConfig,omitempty" type:"Repeated"`
	// The opcode configuration.
	//
	// example:
	//
	// {"opCode":"3"}
	TaskConfig *string `json:"TaskConfig,omitempty" xml:"TaskConfig,omitempty"`
	// example:
	//
	// kill_process_isolate_file
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	// Indicates whether the playbook is intended for Web Application Firewall (WAF). Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	WafPlaybook *bool `json:"WafPlaybook,omitempty" xml:"WafPlaybook,omitempty"`
}

func (s DescribeDisposeAndPlaybookResponseBodyDataResponseDataPlaybookList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDisposeAndPlaybookResponseBodyDataResponseDataPlaybookList) GoString() string {
	return s.String()
}

func (s *DescribeDisposeAndPlaybookResponseBodyDataResponseDataPlaybookList) SetAvailable(v string) *DescribeDisposeAndPlaybookResponseBodyDataResponseDataPlaybookList {
	s.Available = &v
	return s
}

func (s *DescribeDisposeAndPlaybookResponseBodyDataResponseDataPlaybookList) SetDescription(v string) *DescribeDisposeAndPlaybookResponseBodyDataResponseDataPlaybookList {
	s.Description = &v
	return s
}

func (s *DescribeDisposeAndPlaybookResponseBodyDataResponseDataPlaybookList) SetDisplayName(v string) *DescribeDisposeAndPlaybookResponseBodyDataResponseDataPlaybookList {
	s.DisplayName = &v
	return s
}

func (s *DescribeDisposeAndPlaybookResponseBodyDataResponseDataPlaybookList) SetName(v string) *DescribeDisposeAndPlaybookResponseBodyDataResponseDataPlaybookList {
	s.Name = &v
	return s
}

func (s *DescribeDisposeAndPlaybookResponseBodyDataResponseDataPlaybookList) SetOpCode(v string) *DescribeDisposeAndPlaybookResponseBodyDataResponseDataPlaybookList {
	s.OpCode = &v
	return s
}

func (s *DescribeDisposeAndPlaybookResponseBodyDataResponseDataPlaybookList) SetOpLevel(v string) *DescribeDisposeAndPlaybookResponseBodyDataResponseDataPlaybookList {
	s.OpLevel = &v
	return s
}

func (s *DescribeDisposeAndPlaybookResponseBodyDataResponseDataPlaybookList) SetParamConfig(v []interface{}) *DescribeDisposeAndPlaybookResponseBodyDataResponseDataPlaybookList {
	s.ParamConfig = v
	return s
}

func (s *DescribeDisposeAndPlaybookResponseBodyDataResponseDataPlaybookList) SetTaskConfig(v string) *DescribeDisposeAndPlaybookResponseBodyDataResponseDataPlaybookList {
	s.TaskConfig = &v
	return s
}

func (s *DescribeDisposeAndPlaybookResponseBodyDataResponseDataPlaybookList) SetUuid(v string) *DescribeDisposeAndPlaybookResponseBodyDataResponseDataPlaybookList {
	s.Uuid = &v
	return s
}

func (s *DescribeDisposeAndPlaybookResponseBodyDataResponseDataPlaybookList) SetWafPlaybook(v bool) *DescribeDisposeAndPlaybookResponseBodyDataResponseDataPlaybookList {
	s.WafPlaybook = &v
	return s
}

type DescribeDisposeAndPlaybookResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDisposeAndPlaybookResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDisposeAndPlaybookResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDisposeAndPlaybookResponse) GoString() string {
	return s.String()
}

func (s *DescribeDisposeAndPlaybookResponse) SetHeaders(v map[string]*string) *DescribeDisposeAndPlaybookResponse {
	s.Headers = v
	return s
}

func (s *DescribeDisposeAndPlaybookResponse) SetStatusCode(v int32) *DescribeDisposeAndPlaybookResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDisposeAndPlaybookResponse) SetBody(v *DescribeDisposeAndPlaybookResponseBody) *DescribeDisposeAndPlaybookResponse {
	s.Body = v
	return s
}

type DescribeDisposeStrategyPlaybookRequest struct {
	// The end of the time range to query. Unit: milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1577808000000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The data management center of the threat analysis feature. Specify this parameter based on the region in which your assets reside. Valid values:
	//
	// 	- cn-hangzhou: Your assets reside in regions inside China.
	//
	// 	- ap-southeast-1: Your assets reside in regions outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the account that you switch from the management account.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view. Valid values:
	//
	// - 0: the current Alibaba Cloud account
	//
	// - 1: the global account
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// The beginning of the time range to query. Unit: milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1577808000000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDisposeStrategyPlaybookRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDisposeStrategyPlaybookRequest) GoString() string {
	return s.String()
}

func (s *DescribeDisposeStrategyPlaybookRequest) SetEndTime(v int64) *DescribeDisposeStrategyPlaybookRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDisposeStrategyPlaybookRequest) SetRegionId(v string) *DescribeDisposeStrategyPlaybookRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDisposeStrategyPlaybookRequest) SetRoleFor(v int64) *DescribeDisposeStrategyPlaybookRequest {
	s.RoleFor = &v
	return s
}

func (s *DescribeDisposeStrategyPlaybookRequest) SetRoleType(v int32) *DescribeDisposeStrategyPlaybookRequest {
	s.RoleType = &v
	return s
}

func (s *DescribeDisposeStrategyPlaybookRequest) SetStartTime(v int64) *DescribeDisposeStrategyPlaybookRequest {
	s.StartTime = &v
	return s
}

type DescribeDisposeStrategyPlaybookResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	//
	// example:
	//
	// 123456
	Data []*DescribeDisposeStrategyPlaybookResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDisposeStrategyPlaybookResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDisposeStrategyPlaybookResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDisposeStrategyPlaybookResponseBody) SetCode(v int32) *DescribeDisposeStrategyPlaybookResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeDisposeStrategyPlaybookResponseBody) SetData(v []*DescribeDisposeStrategyPlaybookResponseBodyData) *DescribeDisposeStrategyPlaybookResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDisposeStrategyPlaybookResponseBody) SetMessage(v string) *DescribeDisposeStrategyPlaybookResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeDisposeStrategyPlaybookResponseBody) SetRequestId(v string) *DescribeDisposeStrategyPlaybookResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDisposeStrategyPlaybookResponseBody) SetSuccess(v bool) *DescribeDisposeStrategyPlaybookResponseBody {
	s.Success = &v
	return s
}

type DescribeDisposeStrategyPlaybookResponseBodyData struct {
	// The playbook name, which is the unique identifier of the playbook.
	//
	// example:
	//
	// WafBlockIP
	PlaybookName *string `json:"PlaybookName,omitempty" xml:"PlaybookName,omitempty"`
	// The UUID of the playbook.
	//
	// example:
	//
	// system_aliyun_clb_process_book
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
}

func (s DescribeDisposeStrategyPlaybookResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDisposeStrategyPlaybookResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDisposeStrategyPlaybookResponseBodyData) SetPlaybookName(v string) *DescribeDisposeStrategyPlaybookResponseBodyData {
	s.PlaybookName = &v
	return s
}

func (s *DescribeDisposeStrategyPlaybookResponseBodyData) SetPlaybookUuid(v string) *DescribeDisposeStrategyPlaybookResponseBodyData {
	s.PlaybookUuid = &v
	return s
}

type DescribeDisposeStrategyPlaybookResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDisposeStrategyPlaybookResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDisposeStrategyPlaybookResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDisposeStrategyPlaybookResponse) GoString() string {
	return s.String()
}

func (s *DescribeDisposeStrategyPlaybookResponse) SetHeaders(v map[string]*string) *DescribeDisposeStrategyPlaybookResponse {
	s.Headers = v
	return s
}

func (s *DescribeDisposeStrategyPlaybookResponse) SetStatusCode(v int32) *DescribeDisposeStrategyPlaybookResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDisposeStrategyPlaybookResponse) SetBody(v *DescribeDisposeStrategyPlaybookResponseBody) *DescribeDisposeStrategyPlaybookResponse {
	s.Body = v
	return s
}

type DescribeEntityInfoRequest struct {
	// The logical ID of the entity.
	//
	// example:
	//
	// 12345
	EntityId *int64 `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The feature value of the entity. Fuzzy match is supported.
	//
	// example:
	//
	// test22.php
	EntityIdentity *string `json:"EntityIdentity,omitempty" xml:"EntityIdentity,omitempty"`
	// The UUID of the event.
	//
	// example:
	//
	// 85ea4241-798f-4684-a876-65d4f0c3****
	IncidentUuid *string `json:"IncidentUuid,omitempty" xml:"IncidentUuid,omitempty"`
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
	// The ID of the account that you switch from the management account.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view. Valid values:
	//
	// - 0: the current Alibaba Cloud account
	//
	// - 1: the global account
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// The ID of the SOAR handling policy.
	//
	// example:
	//
	// 577bbf90-a770-44a7-8154-586aa2d318fa
	SophonTaskId *string `json:"SophonTaskId,omitempty" xml:"SophonTaskId,omitempty"`
}

func (s DescribeEntityInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEntityInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeEntityInfoRequest) SetEntityId(v int64) *DescribeEntityInfoRequest {
	s.EntityId = &v
	return s
}

func (s *DescribeEntityInfoRequest) SetEntityIdentity(v string) *DescribeEntityInfoRequest {
	s.EntityIdentity = &v
	return s
}

func (s *DescribeEntityInfoRequest) SetIncidentUuid(v string) *DescribeEntityInfoRequest {
	s.IncidentUuid = &v
	return s
}

func (s *DescribeEntityInfoRequest) SetRegionId(v string) *DescribeEntityInfoRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeEntityInfoRequest) SetRoleFor(v int64) *DescribeEntityInfoRequest {
	s.RoleFor = &v
	return s
}

func (s *DescribeEntityInfoRequest) SetRoleType(v int32) *DescribeEntityInfoRequest {
	s.RoleType = &v
	return s
}

func (s *DescribeEntityInfoRequest) SetSophonTaskId(v string) *DescribeEntityInfoRequest {
	s.SophonTaskId = &v
	return s
}

type DescribeEntityInfoResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	//
	// example:
	//
	// 123456
	Data *DescribeEntityInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeEntityInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeEntityInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEntityInfoResponseBody) SetCode(v int32) *DescribeEntityInfoResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeEntityInfoResponseBody) SetData(v *DescribeEntityInfoResponseBodyData) *DescribeEntityInfoResponseBody {
	s.Data = v
	return s
}

func (s *DescribeEntityInfoResponseBody) SetMessage(v string) *DescribeEntityInfoResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeEntityInfoResponseBody) SetRequestId(v string) *DescribeEntityInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEntityInfoResponseBody) SetSuccess(v bool) *DescribeEntityInfoResponseBody {
	s.Success = &v
	return s
}

type DescribeEntityInfoResponseBodyData struct {
	// The logical ID of the entity.
	//
	// example:
	//
	// 12345
	EntityId *int64 `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The information about the entry.
	//
	// example:
	//
	// { location: "xian", net_connect_dir: "in", malware_type: "${aliyun.siem.sas.alert_tag.login_unusual_account}" }
	EntityInfo map[string]interface{} `json:"EntityInfo,omitempty" xml:"EntityInfo,omitempty"`
	// The type of the entity. Valid values:
	//
	// 	- ip
	//
	// 	- domain
	//
	// 	- url
	//
	// 	- process
	//
	// 	- file
	//
	// 	- host
	//
	// example:
	//
	// ip
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// The information about the risk Intelligence.
	//
	// example:
	//
	// {
	//
	//       "Ip": {
	//
	//             "queryHot": "0",
	//
	//             "country": "China",
	//
	//             "province": "shanxi",
	//
	//             "ip": "221.11.XX.XXX",
	//
	//             "asn": "4837",
	//
	//             "asn_label": "CHINAXXX-Backbone - CHINA UNICOM ChinaXXX Backbone, CN"
	//
	//       }
	//
	// }
	TipInfo map[string]interface{} `json:"TipInfo,omitempty" xml:"TipInfo,omitempty"`
}

func (s DescribeEntityInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeEntityInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeEntityInfoResponseBodyData) SetEntityId(v int64) *DescribeEntityInfoResponseBodyData {
	s.EntityId = &v
	return s
}

func (s *DescribeEntityInfoResponseBodyData) SetEntityInfo(v map[string]interface{}) *DescribeEntityInfoResponseBodyData {
	s.EntityInfo = v
	return s
}

func (s *DescribeEntityInfoResponseBodyData) SetEntityType(v string) *DescribeEntityInfoResponseBodyData {
	s.EntityType = &v
	return s
}

func (s *DescribeEntityInfoResponseBodyData) SetTipInfo(v map[string]interface{}) *DescribeEntityInfoResponseBodyData {
	s.TipInfo = v
	return s
}

type DescribeEntityInfoResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEntityInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEntityInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEntityInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeEntityInfoResponse) SetHeaders(v map[string]*string) *DescribeEntityInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeEntityInfoResponse) SetStatusCode(v int32) *DescribeEntityInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEntityInfoResponse) SetBody(v *DescribeEntityInfoResponseBody) *DescribeEntityInfoResponse {
	s.Body = v
	return s
}

type DescribeEventCountByThreatLevelRequest struct {
	// The end of the time range to query. Unit: milliseconds.
	//
	// example:
	//
	// 1577808000000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
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
	// The ID of the member in the resource directory.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view.
	//
	// 	- 0: view of the current Alibaba Cloud account.
	//
	// 	- 1: view of all accounts for the enterprise.
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// The beginning of the time range to query. Unit: milliseconds.
	//
	// example:
	//
	// 1577808000000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeEventCountByThreatLevelRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventCountByThreatLevelRequest) GoString() string {
	return s.String()
}

func (s *DescribeEventCountByThreatLevelRequest) SetEndTime(v int64) *DescribeEventCountByThreatLevelRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeEventCountByThreatLevelRequest) SetRegionId(v string) *DescribeEventCountByThreatLevelRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeEventCountByThreatLevelRequest) SetRoleFor(v int64) *DescribeEventCountByThreatLevelRequest {
	s.RoleFor = &v
	return s
}

func (s *DescribeEventCountByThreatLevelRequest) SetRoleType(v int32) *DescribeEventCountByThreatLevelRequest {
	s.RoleType = &v
	return s
}

func (s *DescribeEventCountByThreatLevelRequest) SetStartTime(v int64) *DescribeEventCountByThreatLevelRequest {
	s.StartTime = &v
	return s
}

type DescribeEventCountByThreatLevelResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	//
	// example:
	//
	// 123456
	Data *DescribeEventCountByThreatLevelResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeEventCountByThreatLevelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventCountByThreatLevelResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEventCountByThreatLevelResponseBody) SetCode(v int32) *DescribeEventCountByThreatLevelResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeEventCountByThreatLevelResponseBody) SetData(v *DescribeEventCountByThreatLevelResponseBodyData) *DescribeEventCountByThreatLevelResponseBody {
	s.Data = v
	return s
}

func (s *DescribeEventCountByThreatLevelResponseBody) SetMessage(v string) *DescribeEventCountByThreatLevelResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeEventCountByThreatLevelResponseBody) SetRequestId(v string) *DescribeEventCountByThreatLevelResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEventCountByThreatLevelResponseBody) SetSuccess(v bool) *DescribeEventCountByThreatLevelResponseBody {
	s.Success = &v
	return s
}

type DescribeEventCountByThreatLevelResponseBodyData struct {
	// The total number of events.
	//
	// example:
	//
	// 100
	EventNum *int64 `json:"EventNum,omitempty" xml:"EventNum,omitempty"`
	// The number of high-risk events.
	//
	// example:
	//
	// 20
	HighLevelEventNum *int64 `json:"HighLevelEventNum,omitempty" xml:"HighLevelEventNum,omitempty"`
	// The number of low-risk events.
	//
	// example:
	//
	// 52
	LowLevelEventNum *int64 `json:"LowLevelEventNum,omitempty" xml:"LowLevelEventNum,omitempty"`
	// The number of medium-risk events.
	//
	// example:
	//
	// 3
	MediumLevelEventNum *int64 `json:"MediumLevelEventNum,omitempty" xml:"MediumLevelEventNum,omitempty"`
	// The number of unhandled events.
	//
	// example:
	//
	// 75
	UndealEventNum *int64 `json:"UndealEventNum,omitempty" xml:"UndealEventNum,omitempty"`
}

func (s DescribeEventCountByThreatLevelResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventCountByThreatLevelResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeEventCountByThreatLevelResponseBodyData) SetEventNum(v int64) *DescribeEventCountByThreatLevelResponseBodyData {
	s.EventNum = &v
	return s
}

func (s *DescribeEventCountByThreatLevelResponseBodyData) SetHighLevelEventNum(v int64) *DescribeEventCountByThreatLevelResponseBodyData {
	s.HighLevelEventNum = &v
	return s
}

func (s *DescribeEventCountByThreatLevelResponseBodyData) SetLowLevelEventNum(v int64) *DescribeEventCountByThreatLevelResponseBodyData {
	s.LowLevelEventNum = &v
	return s
}

func (s *DescribeEventCountByThreatLevelResponseBodyData) SetMediumLevelEventNum(v int64) *DescribeEventCountByThreatLevelResponseBodyData {
	s.MediumLevelEventNum = &v
	return s
}

func (s *DescribeEventCountByThreatLevelResponseBodyData) SetUndealEventNum(v int64) *DescribeEventCountByThreatLevelResponseBodyData {
	s.UndealEventNum = &v
	return s
}

type DescribeEventCountByThreatLevelResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEventCountByThreatLevelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEventCountByThreatLevelResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventCountByThreatLevelResponse) GoString() string {
	return s.String()
}

func (s *DescribeEventCountByThreatLevelResponse) SetHeaders(v map[string]*string) *DescribeEventCountByThreatLevelResponse {
	s.Headers = v
	return s
}

func (s *DescribeEventCountByThreatLevelResponse) SetStatusCode(v int32) *DescribeEventCountByThreatLevelResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEventCountByThreatLevelResponse) SetBody(v *DescribeEventCountByThreatLevelResponseBody) *DescribeEventCountByThreatLevelResponse {
	s.Body = v
	return s
}

type DescribeEventDisposeRequest struct {
	// The page number. Pages start from page 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The UUID of the event.
	//
	// example:
	//
	// 85ea4241-798f-4684-a876-65d4f0c3****
	IncidentUuid *string `json:"IncidentUuid,omitempty" xml:"IncidentUuid,omitempty"`
	// The number of entries per page. Maximum value: 500.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The data management center of the threat analysis feature. Specify this parameter based on the region in which your assets reside. Valid values:
	//
	// 	- cn-hangzhou: Your assets reside in regions inside China.
	//
	// 	- ap-southeast-1: Your assets reside in regions outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the account that you switch from the management account.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view. Valid values:
	//
	// - 0: the current Alibaba Cloud account
	//
	// - 1: the global account
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s DescribeEventDisposeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventDisposeRequest) GoString() string {
	return s.String()
}

func (s *DescribeEventDisposeRequest) SetCurrentPage(v int32) *DescribeEventDisposeRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeEventDisposeRequest) SetIncidentUuid(v string) *DescribeEventDisposeRequest {
	s.IncidentUuid = &v
	return s
}

func (s *DescribeEventDisposeRequest) SetPageSize(v int32) *DescribeEventDisposeRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeEventDisposeRequest) SetRegionId(v string) *DescribeEventDisposeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeEventDisposeRequest) SetRoleFor(v int64) *DescribeEventDisposeRequest {
	s.RoleFor = &v
	return s
}

func (s *DescribeEventDisposeRequest) SetRoleType(v int32) *DescribeEventDisposeRequest {
	s.RoleType = &v
	return s
}

type DescribeEventDisposeResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	//
	// example:
	//
	// 123456
	Data *DescribeEventDisposeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeEventDisposeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventDisposeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEventDisposeResponseBody) SetCode(v int32) *DescribeEventDisposeResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeEventDisposeResponseBody) SetData(v *DescribeEventDisposeResponseBodyData) *DescribeEventDisposeResponseBody {
	s.Data = v
	return s
}

func (s *DescribeEventDisposeResponseBody) SetMessage(v string) *DescribeEventDisposeResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeEventDisposeResponseBody) SetRequestId(v string) *DescribeEventDisposeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEventDisposeResponseBody) SetSuccess(v bool) *DescribeEventDisposeResponseBody {
	s.Success = &v
	return s
}

type DescribeEventDisposeResponseBodyData struct {
	// An array consisting of JSON objects that are configured for event handling.
	//
	// example:
	//
	// { playbookName: "使用安全组封禁入方向IP", sophonTaskId: "400442a5-4f98-45ed-97db-5ab117eb0b8f", … }
	EventDispose []interface{} `json:"EventDispose,omitempty" xml:"EventDispose,omitempty" type:"Repeated"`
	// The JSON object that is configured for an alert recipient.
	ReceiverInfo *DescribeEventDisposeResponseBodyDataReceiverInfo `json:"ReceiverInfo,omitempty" xml:"ReceiverInfo,omitempty" type:"Struct"`
	// The description of the event.
	//
	// example:
	//
	// dealed
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The status of the event. Valid values:
	//
	// 	- 0: not handled
	//
	// 	- 1: handing
	//
	// 	- 5: handling failed
	//
	// 	- 10: handled
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeEventDisposeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventDisposeResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeEventDisposeResponseBodyData) SetEventDispose(v []interface{}) *DescribeEventDisposeResponseBodyData {
	s.EventDispose = v
	return s
}

func (s *DescribeEventDisposeResponseBodyData) SetReceiverInfo(v *DescribeEventDisposeResponseBodyDataReceiverInfo) *DescribeEventDisposeResponseBodyData {
	s.ReceiverInfo = v
	return s
}

func (s *DescribeEventDisposeResponseBodyData) SetRemark(v string) *DescribeEventDisposeResponseBodyData {
	s.Remark = &v
	return s
}

func (s *DescribeEventDisposeResponseBodyData) SetStatus(v int32) *DescribeEventDisposeResponseBodyData {
	s.Status = &v
	return s
}

type DescribeEventDisposeResponseBodyDataReceiverInfo struct {
	// The channel of the contact information. Valid values:
	//
	// 	- message
	//
	// 	- mail
	//
	// example:
	//
	// message
	Channel *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2021-01-06 16:37:29
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The modification time.
	//
	// example:
	//
	// 2021-01-06 16:37:29
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the recipient who receives the event handling result.
	//
	// example:
	//
	// 123
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The UUID of the event.
	//
	// example:
	//
	// 85ea4241-798f-4684-a876-65d4f0c3****
	IncidentUuid *string `json:"IncidentUuid,omitempty" xml:"IncidentUuid,omitempty"`
	// The message title.
	//
	// example:
	//
	// siem event dealed message
	MessageTitle *string `json:"MessageTitle,omitempty" xml:"MessageTitle,omitempty"`
	// The contact information of the recipient.
	//
	// example:
	//
	// 138xxxxxx
	Receiver *string `json:"Receiver,omitempty" xml:"Receiver,omitempty"`
	// Indicates whether the message is sent. Valid values:
	//
	// 	- 0: not sent
	//
	// 	- 1: sent
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeEventDisposeResponseBodyDataReceiverInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventDisposeResponseBodyDataReceiverInfo) GoString() string {
	return s.String()
}

func (s *DescribeEventDisposeResponseBodyDataReceiverInfo) SetChannel(v string) *DescribeEventDisposeResponseBodyDataReceiverInfo {
	s.Channel = &v
	return s
}

func (s *DescribeEventDisposeResponseBodyDataReceiverInfo) SetGmtCreate(v string) *DescribeEventDisposeResponseBodyDataReceiverInfo {
	s.GmtCreate = &v
	return s
}

func (s *DescribeEventDisposeResponseBodyDataReceiverInfo) SetGmtModified(v string) *DescribeEventDisposeResponseBodyDataReceiverInfo {
	s.GmtModified = &v
	return s
}

func (s *DescribeEventDisposeResponseBodyDataReceiverInfo) SetId(v int64) *DescribeEventDisposeResponseBodyDataReceiverInfo {
	s.Id = &v
	return s
}

func (s *DescribeEventDisposeResponseBodyDataReceiverInfo) SetIncidentUuid(v string) *DescribeEventDisposeResponseBodyDataReceiverInfo {
	s.IncidentUuid = &v
	return s
}

func (s *DescribeEventDisposeResponseBodyDataReceiverInfo) SetMessageTitle(v string) *DescribeEventDisposeResponseBodyDataReceiverInfo {
	s.MessageTitle = &v
	return s
}

func (s *DescribeEventDisposeResponseBodyDataReceiverInfo) SetReceiver(v string) *DescribeEventDisposeResponseBodyDataReceiverInfo {
	s.Receiver = &v
	return s
}

func (s *DescribeEventDisposeResponseBodyDataReceiverInfo) SetStatus(v int32) *DescribeEventDisposeResponseBodyDataReceiverInfo {
	s.Status = &v
	return s
}

type DescribeEventDisposeResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEventDisposeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEventDisposeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventDisposeResponse) GoString() string {
	return s.String()
}

func (s *DescribeEventDisposeResponse) SetHeaders(v map[string]*string) *DescribeEventDisposeResponse {
	s.Headers = v
	return s
}

func (s *DescribeEventDisposeResponse) SetStatusCode(v int32) *DescribeEventDisposeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEventDisposeResponse) SetBody(v *DescribeEventDisposeResponseBody) *DescribeEventDisposeResponse {
	s.Body = v
	return s
}

type DescribeImportedLogCountRequest struct {
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
	RoleFor  *string `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	RoleType *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s DescribeImportedLogCountRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeImportedLogCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeImportedLogCountRequest) SetRegionId(v string) *DescribeImportedLogCountRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeImportedLogCountRequest) SetRoleFor(v string) *DescribeImportedLogCountRequest {
	s.RoleFor = &v
	return s
}

func (s *DescribeImportedLogCountRequest) SetRoleType(v string) *DescribeImportedLogCountRequest {
	s.RoleType = &v
	return s
}

type DescribeImportedLogCountResponseBody struct {
	// The data returned.
	Data *DescribeImportedLogCountResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeImportedLogCountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeImportedLogCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeImportedLogCountResponseBody) SetData(v *DescribeImportedLogCountResponseBodyData) *DescribeImportedLogCountResponseBody {
	s.Data = v
	return s
}

func (s *DescribeImportedLogCountResponseBody) SetRequestId(v string) *DescribeImportedLogCountResponseBody {
	s.RequestId = &v
	return s
}

type DescribeImportedLogCountResponseBodyData struct {
	// The number of logs that are added.
	//
	// example:
	//
	// 10
	ImportedLogCount *int32 `json:"ImportedLogCount,omitempty" xml:"ImportedLogCount,omitempty"`
	// The total number of logs.
	//
	// example:
	//
	// 59
	TotalLogCount *int32 `json:"TotalLogCount,omitempty" xml:"TotalLogCount,omitempty"`
	// The number of logs that are not added.
	//
	// example:
	//
	// 49
	UnImportedLogCount *int32 `json:"UnImportedLogCount,omitempty" xml:"UnImportedLogCount,omitempty"`
}

func (s DescribeImportedLogCountResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeImportedLogCountResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeImportedLogCountResponseBodyData) SetImportedLogCount(v int32) *DescribeImportedLogCountResponseBodyData {
	s.ImportedLogCount = &v
	return s
}

func (s *DescribeImportedLogCountResponseBodyData) SetTotalLogCount(v int32) *DescribeImportedLogCountResponseBodyData {
	s.TotalLogCount = &v
	return s
}

func (s *DescribeImportedLogCountResponseBodyData) SetUnImportedLogCount(v int32) *DescribeImportedLogCountResponseBodyData {
	s.UnImportedLogCount = &v
	return s
}

type DescribeImportedLogCountResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeImportedLogCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeImportedLogCountResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeImportedLogCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeImportedLogCountResponse) SetHeaders(v map[string]*string) *DescribeImportedLogCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeImportedLogCountResponse) SetStatusCode(v int32) *DescribeImportedLogCountResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeImportedLogCountResponse) SetBody(v *DescribeImportedLogCountResponseBody) *DescribeImportedLogCountResponse {
	s.Body = v
	return s
}

type DescribeLogFieldsRequest struct {
	// The log source of the rule.
	//
	// example:
	//
	// cloud_siem_aegis_sas_alert
	LogSource *string `json:"LogSource,omitempty" xml:"LogSource,omitempty"`
	// The log type of the rule.
	//
	// example:
	//
	// cloud_siem_aegis_sas_alert
	LogType *string `json:"LogType,omitempty" xml:"LogType,omitempty"`
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
	// The ID of the account that you switch from the management account.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view. Valid values:
	//
	// - 0: the current Alibaba Cloud account
	//
	// - 1: the global account
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s DescribeLogFieldsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogFieldsRequest) GoString() string {
	return s.String()
}

func (s *DescribeLogFieldsRequest) SetLogSource(v string) *DescribeLogFieldsRequest {
	s.LogSource = &v
	return s
}

func (s *DescribeLogFieldsRequest) SetLogType(v string) *DescribeLogFieldsRequest {
	s.LogType = &v
	return s
}

func (s *DescribeLogFieldsRequest) SetRegionId(v string) *DescribeLogFieldsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLogFieldsRequest) SetRoleFor(v int64) *DescribeLogFieldsRequest {
	s.RoleFor = &v
	return s
}

func (s *DescribeLogFieldsRequest) SetRoleType(v int32) *DescribeLogFieldsRequest {
	s.RoleType = &v
	return s
}

type DescribeLogFieldsResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	//
	// example:
	//
	// 123456
	Data []*DescribeLogFieldsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeLogFieldsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogFieldsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLogFieldsResponseBody) SetCode(v int32) *DescribeLogFieldsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeLogFieldsResponseBody) SetData(v []*DescribeLogFieldsResponseBodyData) *DescribeLogFieldsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeLogFieldsResponseBody) SetMessage(v string) *DescribeLogFieldsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeLogFieldsResponseBody) SetRequestId(v string) *DescribeLogFieldsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLogFieldsResponseBody) SetSuccess(v bool) *DescribeLogFieldsResponseBody {
	s.Success = &v
	return s
}

type DescribeLogFieldsResponseBodyData struct {
	// The type of the log to which the field belongs.
	//
	// example:
	//
	// HTTP_ACTIVITY
	ActivityName *string `json:"ActivityName,omitempty" xml:"ActivityName,omitempty"`
	// The internal code of the field description.
	//
	// example:
	//
	// sas.cloudsiem.prod.activity_name
	FieldDesc *string `json:"FieldDesc,omitempty" xml:"FieldDesc,omitempty"`
	// The name of the field.
	//
	// example:
	//
	// activity_name
	FieldName *string `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
	// The data type of the field. Valid values:
	//
	// 	- varchar
	//
	// 	- bigint
	//
	// example:
	//
	// varchar
	FieldType *string `json:"FieldType,omitempty" xml:"FieldType,omitempty"`
	// The log source to which the field belongs.
	//
	// example:
	//
	// cloud_siem_aegis_sas_alert
	LogCode *string `json:"LogCode,omitempty" xml:"LogCode,omitempty"`
}

func (s DescribeLogFieldsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogFieldsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeLogFieldsResponseBodyData) SetActivityName(v string) *DescribeLogFieldsResponseBodyData {
	s.ActivityName = &v
	return s
}

func (s *DescribeLogFieldsResponseBodyData) SetFieldDesc(v string) *DescribeLogFieldsResponseBodyData {
	s.FieldDesc = &v
	return s
}

func (s *DescribeLogFieldsResponseBodyData) SetFieldName(v string) *DescribeLogFieldsResponseBodyData {
	s.FieldName = &v
	return s
}

func (s *DescribeLogFieldsResponseBodyData) SetFieldType(v string) *DescribeLogFieldsResponseBodyData {
	s.FieldType = &v
	return s
}

func (s *DescribeLogFieldsResponseBodyData) SetLogCode(v string) *DescribeLogFieldsResponseBodyData {
	s.LogCode = &v
	return s
}

type DescribeLogFieldsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLogFieldsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLogFieldsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogFieldsResponse) GoString() string {
	return s.String()
}

func (s *DescribeLogFieldsResponse) SetHeaders(v map[string]*string) *DescribeLogFieldsResponse {
	s.Headers = v
	return s
}

func (s *DescribeLogFieldsResponse) SetStatusCode(v int32) *DescribeLogFieldsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLogFieldsResponse) SetBody(v *DescribeLogFieldsResponseBody) *DescribeLogFieldsResponse {
	s.Body = v
	return s
}

type DescribeLogSourceRequest struct {
	// The log type of the rule.
	//
	// example:
	//
	// HTTP_ACTIVITY
	LogType *string `json:"LogType,omitempty" xml:"LogType,omitempty"`
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
	// The ID of the account that you switch from the management account.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view. Valid values:
	//
	// - 0: the current Alibaba Cloud account
	//
	// - 1: the global account
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s DescribeLogSourceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogSourceRequest) GoString() string {
	return s.String()
}

func (s *DescribeLogSourceRequest) SetLogType(v string) *DescribeLogSourceRequest {
	s.LogType = &v
	return s
}

func (s *DescribeLogSourceRequest) SetRegionId(v string) *DescribeLogSourceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLogSourceRequest) SetRoleFor(v int64) *DescribeLogSourceRequest {
	s.RoleFor = &v
	return s
}

func (s *DescribeLogSourceRequest) SetRoleType(v int32) *DescribeLogSourceRequest {
	s.RoleType = &v
	return s
}

type DescribeLogSourceResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	//
	// example:
	//
	// 123456
	Data []*DescribeLogSourceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeLogSourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogSourceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLogSourceResponseBody) SetCode(v int32) *DescribeLogSourceResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeLogSourceResponseBody) SetData(v []*DescribeLogSourceResponseBodyData) *DescribeLogSourceResponseBody {
	s.Data = v
	return s
}

func (s *DescribeLogSourceResponseBody) SetMessage(v string) *DescribeLogSourceResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeLogSourceResponseBody) SetRequestId(v string) *DescribeLogSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLogSourceResponseBody) SetSuccess(v bool) *DescribeLogSourceResponseBody {
	s.Success = &v
	return s
}

type DescribeLogSourceResponseBodyData struct {
	// The log source of the rule.
	//
	// example:
	//
	// cloud_siem_aegis_sas_alert
	LogSource *string `json:"LogSource,omitempty" xml:"LogSource,omitempty"`
	// The internal code of the log source.
	//
	// example:
	//
	// sas.cloudsiem.prod.cloud_siem_aegis_sas_alert
	LogSourceName *string `json:"LogSourceName,omitempty" xml:"LogSourceName,omitempty"`
}

func (s DescribeLogSourceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogSourceResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeLogSourceResponseBodyData) SetLogSource(v string) *DescribeLogSourceResponseBodyData {
	s.LogSource = &v
	return s
}

func (s *DescribeLogSourceResponseBodyData) SetLogSourceName(v string) *DescribeLogSourceResponseBodyData {
	s.LogSourceName = &v
	return s
}

type DescribeLogSourceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLogSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLogSourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogSourceResponse) GoString() string {
	return s.String()
}

func (s *DescribeLogSourceResponse) SetHeaders(v map[string]*string) *DescribeLogSourceResponse {
	s.Headers = v
	return s
}

func (s *DescribeLogSourceResponse) SetStatusCode(v int32) *DescribeLogSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLogSourceResponse) SetBody(v *DescribeLogSourceResponseBody) *DescribeLogSourceResponse {
	s.Body = v
	return s
}

type DescribeLogTypeRequest struct {
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
	// The ID of the account that you switch from the management account.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view. Valid values:
	//
	// - 0: the current Alibaba Cloud account
	//
	// - 1: the global account
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s DescribeLogTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogTypeRequest) GoString() string {
	return s.String()
}

func (s *DescribeLogTypeRequest) SetRegionId(v string) *DescribeLogTypeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLogTypeRequest) SetRoleFor(v int64) *DescribeLogTypeRequest {
	s.RoleFor = &v
	return s
}

func (s *DescribeLogTypeRequest) SetRoleType(v int32) *DescribeLogTypeRequest {
	s.RoleType = &v
	return s
}

type DescribeLogTypeResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	//
	// example:
	//
	// 123456
	Data []*DescribeLogTypeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeLogTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogTypeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLogTypeResponseBody) SetCode(v int32) *DescribeLogTypeResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeLogTypeResponseBody) SetData(v []*DescribeLogTypeResponseBodyData) *DescribeLogTypeResponseBody {
	s.Data = v
	return s
}

func (s *DescribeLogTypeResponseBody) SetMessage(v string) *DescribeLogTypeResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeLogTypeResponseBody) SetRequestId(v string) *DescribeLogTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLogTypeResponseBody) SetSuccess(v bool) *DescribeLogTypeResponseBody {
	s.Success = &v
	return s
}

type DescribeLogTypeResponseBodyData struct {
	// The log type of the rule.
	//
	// example:
	//
	// HTTP_ACTIVITY
	LogType *string `json:"LogType,omitempty" xml:"LogType,omitempty"`
	// The internal code of the log type.
	//
	// example:
	//
	// sas.cloudsiem.prod.http_activity
	LogTypeName *string `json:"LogTypeName,omitempty" xml:"LogTypeName,omitempty"`
}

func (s DescribeLogTypeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogTypeResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeLogTypeResponseBodyData) SetLogType(v string) *DescribeLogTypeResponseBodyData {
	s.LogType = &v
	return s
}

func (s *DescribeLogTypeResponseBodyData) SetLogTypeName(v string) *DescribeLogTypeResponseBodyData {
	s.LogTypeName = &v
	return s
}

type DescribeLogTypeResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLogTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLogTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogTypeResponse) GoString() string {
	return s.String()
}

func (s *DescribeLogTypeResponse) SetHeaders(v map[string]*string) *DescribeLogTypeResponse {
	s.Headers = v
	return s
}

func (s *DescribeLogTypeResponse) SetStatusCode(v int32) *DescribeLogTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLogTypeResponse) SetBody(v *DescribeLogTypeResponseBody) *DescribeLogTypeResponse {
	s.Body = v
	return s
}

type DescribeOperatorsRequest struct {
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
	// The ID of the account that you switch from the management account.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view. Valid values:
	//
	// - 0: the current Alibaba Cloud account
	//
	// - 1: the global account
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// The type of the scenario in which the operator is used. Valid values:
	//
	// 	- If you do not specify this parameter, the default scenario is used.
	//
	// 	- AGGREGATE: AGGREGATE scenario.
	//
	// example:
	//
	// AGGREGATE
	SceneType *string `json:"SceneType,omitempty" xml:"SceneType,omitempty"`
}

func (s DescribeOperatorsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeOperatorsRequest) GoString() string {
	return s.String()
}

func (s *DescribeOperatorsRequest) SetRegionId(v string) *DescribeOperatorsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeOperatorsRequest) SetRoleFor(v int64) *DescribeOperatorsRequest {
	s.RoleFor = &v
	return s
}

func (s *DescribeOperatorsRequest) SetRoleType(v int32) *DescribeOperatorsRequest {
	s.RoleType = &v
	return s
}

func (s *DescribeOperatorsRequest) SetSceneType(v string) *DescribeOperatorsRequest {
	s.SceneType = &v
	return s
}

type DescribeOperatorsResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	//
	// example:
	//
	// 123456
	Data []*DescribeOperatorsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeOperatorsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeOperatorsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOperatorsResponseBody) SetCode(v int32) *DescribeOperatorsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeOperatorsResponseBody) SetData(v []*DescribeOperatorsResponseBodyData) *DescribeOperatorsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeOperatorsResponseBody) SetMessage(v string) *DescribeOperatorsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeOperatorsResponseBody) SetRequestId(v string) *DescribeOperatorsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOperatorsResponseBody) SetSuccess(v bool) *DescribeOperatorsResponseBody {
	s.Success = &v
	return s
}

type DescribeOperatorsResponseBodyData struct {
	// The position of the operator in the operator list.
	//
	// example:
	//
	// 3
	Index *int32 `json:"Index,omitempty" xml:"Index,omitempty"`
	// The operator.
	//
	// example:
	//
	// <=
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// The description of the operator in Chinese.
	//
	// example:
	//
	// arger than or equal to
	OperatorDescCn *string `json:"OperatorDescCn,omitempty" xml:"OperatorDescCn,omitempty"`
	// The description of the operator in English.
	//
	// example:
	//
	// larger than or equal to
	OperatorDescEn *string `json:"OperatorDescEn,omitempty" xml:"OperatorDescEn,omitempty"`
	// The name of the operator.
	//
	// example:
	//
	// <=
	OperatorName *string `json:"OperatorName,omitempty" xml:"OperatorName,omitempty"`
	// The data types that are supported by the operator. The data types are separated by commas (,).
	//
	// example:
	//
	// varchar
	SupportDataType *string `json:"SupportDataType,omitempty" xml:"SupportDataType,omitempty"`
	// The scenarios that are supported by the operator. Multiple scenarios are separated by commas (,), such as AGGREGATE scenarios. By default, this parameter is empty.
	//
	// example:
	//
	// [AGGREGATE]
	SupportTag []*string `json:"SupportTag,omitempty" xml:"SupportTag,omitempty" type:"Repeated"`
}

func (s DescribeOperatorsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeOperatorsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeOperatorsResponseBodyData) SetIndex(v int32) *DescribeOperatorsResponseBodyData {
	s.Index = &v
	return s
}

func (s *DescribeOperatorsResponseBodyData) SetOperator(v string) *DescribeOperatorsResponseBodyData {
	s.Operator = &v
	return s
}

func (s *DescribeOperatorsResponseBodyData) SetOperatorDescCn(v string) *DescribeOperatorsResponseBodyData {
	s.OperatorDescCn = &v
	return s
}

func (s *DescribeOperatorsResponseBodyData) SetOperatorDescEn(v string) *DescribeOperatorsResponseBodyData {
	s.OperatorDescEn = &v
	return s
}

func (s *DescribeOperatorsResponseBodyData) SetOperatorName(v string) *DescribeOperatorsResponseBodyData {
	s.OperatorName = &v
	return s
}

func (s *DescribeOperatorsResponseBodyData) SetSupportDataType(v string) *DescribeOperatorsResponseBodyData {
	s.SupportDataType = &v
	return s
}

func (s *DescribeOperatorsResponseBodyData) SetSupportTag(v []*string) *DescribeOperatorsResponseBodyData {
	s.SupportTag = v
	return s
}

type DescribeOperatorsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeOperatorsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeOperatorsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeOperatorsResponse) GoString() string {
	return s.String()
}

func (s *DescribeOperatorsResponse) SetHeaders(v map[string]*string) *DescribeOperatorsResponse {
	s.Headers = v
	return s
}

func (s *DescribeOperatorsResponse) SetStatusCode(v int32) *DescribeOperatorsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOperatorsResponse) SetBody(v *DescribeOperatorsResponseBody) *DescribeOperatorsResponse {
	s.Body = v
	return s
}

type DescribeProdCountRequest struct {
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
	RoleFor  *int64  `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	RoleType *int32  `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s DescribeProdCountRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeProdCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeProdCountRequest) SetRegionId(v string) *DescribeProdCountRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeProdCountRequest) SetRoleFor(v int64) *DescribeProdCountRequest {
	s.RoleFor = &v
	return s
}

func (s *DescribeProdCountRequest) SetRoleType(v int32) *DescribeProdCountRequest {
	s.RoleType = &v
	return s
}

type DescribeProdCountResponseBody struct {
	// The data returned.
	Data *DescribeProdCountResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeProdCountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeProdCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeProdCountResponseBody) SetData(v *DescribeProdCountResponseBodyData) *DescribeProdCountResponseBody {
	s.Data = v
	return s
}

func (s *DescribeProdCountResponseBody) SetRequestId(v string) *DescribeProdCountResponseBody {
	s.RequestId = &v
	return s
}

type DescribeProdCountResponseBodyData struct {
	AliyunImportedCount *int32 `json:"AliyunImportedCount,omitempty" xml:"AliyunImportedCount,omitempty"`
	// The number of Alibaba Cloud services.
	//
	// example:
	//
	// 19
	AliyunProdCount     *int32 `json:"AliyunProdCount,omitempty" xml:"AliyunProdCount,omitempty"`
	HcloudImportedCount *int32 `json:"HcloudImportedCount,omitempty" xml:"HcloudImportedCount,omitempty"`
	// The number of Huawei Cloud services.
	//
	// example:
	//
	// 2
	HcloudProdCount  *int32 `json:"HcloudProdCount,omitempty" xml:"HcloudProdCount,omitempty"`
	IdcImportedCount *int32 `json:"IdcImportedCount,omitempty" xml:"IdcImportedCount,omitempty"`
	// example:
	//
	// 2
	IdcProdCount        *int32 `json:"IdcProdCount,omitempty" xml:"IdcProdCount,omitempty"`
	QcloudImportedCount *int32 `json:"QcloudImportedCount,omitempty" xml:"QcloudImportedCount,omitempty"`
	// The number of Tencent Cloud services.
	//
	// example:
	//
	// 2
	QcloudProdCount *int32 `json:"QcloudProdCount,omitempty" xml:"QcloudProdCount,omitempty"`
}

func (s DescribeProdCountResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeProdCountResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeProdCountResponseBodyData) SetAliyunImportedCount(v int32) *DescribeProdCountResponseBodyData {
	s.AliyunImportedCount = &v
	return s
}

func (s *DescribeProdCountResponseBodyData) SetAliyunProdCount(v int32) *DescribeProdCountResponseBodyData {
	s.AliyunProdCount = &v
	return s
}

func (s *DescribeProdCountResponseBodyData) SetHcloudImportedCount(v int32) *DescribeProdCountResponseBodyData {
	s.HcloudImportedCount = &v
	return s
}

func (s *DescribeProdCountResponseBodyData) SetHcloudProdCount(v int32) *DescribeProdCountResponseBodyData {
	s.HcloudProdCount = &v
	return s
}

func (s *DescribeProdCountResponseBodyData) SetIdcImportedCount(v int32) *DescribeProdCountResponseBodyData {
	s.IdcImportedCount = &v
	return s
}

func (s *DescribeProdCountResponseBodyData) SetIdcProdCount(v int32) *DescribeProdCountResponseBodyData {
	s.IdcProdCount = &v
	return s
}

func (s *DescribeProdCountResponseBodyData) SetQcloudImportedCount(v int32) *DescribeProdCountResponseBodyData {
	s.QcloudImportedCount = &v
	return s
}

func (s *DescribeProdCountResponseBodyData) SetQcloudProdCount(v int32) *DescribeProdCountResponseBodyData {
	s.QcloudProdCount = &v
	return s
}

type DescribeProdCountResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeProdCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeProdCountResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeProdCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeProdCountResponse) SetHeaders(v map[string]*string) *DescribeProdCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeProdCountResponse) SetStatusCode(v int32) *DescribeProdCountResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeProdCountResponse) SetBody(v *DescribeProdCountResponseBody) *DescribeProdCountResponse {
	s.Body = v
	return s
}

type DescribeScopeUsersRequest struct {
	// The data management center of the threat analysis feature. Specify this parameter based on the region in which your assets reside. Valid values:
	//
	// 	- cn-hangzhou: Your assets reside in regions inside China.
	//
	// 	- ap-southeast-1: Your assets reside in regions outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the account that you switch from the management account.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view. Valid values:
	//
	// - 0: the current Alibaba Cloud account
	//
	// - 1: the global account
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s DescribeScopeUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeScopeUsersRequest) GoString() string {
	return s.String()
}

func (s *DescribeScopeUsersRequest) SetRegionId(v string) *DescribeScopeUsersRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeScopeUsersRequest) SetRoleFor(v int64) *DescribeScopeUsersRequest {
	s.RoleFor = &v
	return s
}

func (s *DescribeScopeUsersRequest) SetRoleType(v int32) *DescribeScopeUsersRequest {
	s.RoleType = &v
	return s
}

type DescribeScopeUsersResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	//
	// example:
	//
	// 123456
	Data []*DescribeScopeUsersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeScopeUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScopeUsersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScopeUsersResponseBody) SetCode(v int32) *DescribeScopeUsersResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeScopeUsersResponseBody) SetData(v []*DescribeScopeUsersResponseBodyData) *DescribeScopeUsersResponseBody {
	s.Data = v
	return s
}

func (s *DescribeScopeUsersResponseBody) SetMessage(v string) *DescribeScopeUsersResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeScopeUsersResponseBody) SetRequestId(v string) *DescribeScopeUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScopeUsersResponseBody) SetSuccess(v bool) *DescribeScopeUsersResponseBody {
	s.Success = &v
	return s
}

type DescribeScopeUsersResponseBodyData struct {
	// The ID of the security information and event management (SIEM) user.
	//
	// example:
	//
	// 123456789****
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// 云code。  取值：
	//
	// - qcloud：腾讯云
	//
	// - hcloud：华为云
	//
	// example:
	//
	// qcloud
	CloudCode *string `json:"CloudCode,omitempty" xml:"CloudCode,omitempty"`
	// An array consisting of the domain names that are protected by the WAF instance.
	//
	// example:
	//
	// [123.com, 456.com]
	Domains []*string `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// example:
	//
	// waf-cn-tl123ast****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 多云用户ID。
	//
	// example:
	//
	// 123456789****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The username.
	//
	// example:
	//
	// test001
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeScopeUsersResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeScopeUsersResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeScopeUsersResponseBodyData) SetAliUid(v int64) *DescribeScopeUsersResponseBodyData {
	s.AliUid = &v
	return s
}

func (s *DescribeScopeUsersResponseBodyData) SetCloudCode(v string) *DescribeScopeUsersResponseBodyData {
	s.CloudCode = &v
	return s
}

func (s *DescribeScopeUsersResponseBodyData) SetDomains(v []*string) *DescribeScopeUsersResponseBodyData {
	s.Domains = v
	return s
}

func (s *DescribeScopeUsersResponseBodyData) SetInstanceId(v string) *DescribeScopeUsersResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *DescribeScopeUsersResponseBodyData) SetUserId(v string) *DescribeScopeUsersResponseBodyData {
	s.UserId = &v
	return s
}

func (s *DescribeScopeUsersResponseBodyData) SetUserName(v string) *DescribeScopeUsersResponseBodyData {
	s.UserName = &v
	return s
}

type DescribeScopeUsersResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeScopeUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeScopeUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScopeUsersResponse) GoString() string {
	return s.String()
}

func (s *DescribeScopeUsersResponse) SetHeaders(v map[string]*string) *DescribeScopeUsersResponse {
	s.Headers = v
	return s
}

func (s *DescribeScopeUsersResponse) SetStatusCode(v int32) *DescribeScopeUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScopeUsersResponse) SetBody(v *DescribeScopeUsersResponseBody) *DescribeScopeUsersResponse {
	s.Body = v
	return s
}

type DescribeServiceStatusRequest struct {
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

func (s DescribeServiceStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeServiceStatusRequest) SetRegionId(v string) *DescribeServiceStatusRequest {
	s.RegionId = &v
	return s
}

type DescribeServiceStatusResponseBody struct {
	// Indicates whether the threat analysis feature is authorized to access the resource directory. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeServiceStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeServiceStatusResponseBody) SetData(v bool) *DescribeServiceStatusResponseBody {
	s.Data = &v
	return s
}

func (s *DescribeServiceStatusResponseBody) SetRequestId(v string) *DescribeServiceStatusResponseBody {
	s.RequestId = &v
	return s
}

type DescribeServiceStatusResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeServiceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeServiceStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeServiceStatusResponse) SetHeaders(v map[string]*string) *DescribeServiceStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeServiceStatusResponse) SetStatusCode(v int32) *DescribeServiceStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeServiceStatusResponse) SetBody(v *DescribeServiceStatusResponseBody) *DescribeServiceStatusResponse {
	s.Body = v
	return s
}

type DescribeStorageRequest struct {
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
	// The ID of the account that you switch from the management account.
	//
	// example:
	//
	// 137820528780****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view. Valid values:
	//
	// - 0: the current Alibaba Cloud account
	//
	// - 1: the global account
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s DescribeStorageRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeStorageRequest) GoString() string {
	return s.String()
}

func (s *DescribeStorageRequest) SetRegionId(v string) *DescribeStorageRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeStorageRequest) SetRoleFor(v int64) *DescribeStorageRequest {
	s.RoleFor = &v
	return s
}

func (s *DescribeStorageRequest) SetRoleType(v int32) *DescribeStorageRequest {
	s.RoleType = &v
	return s
}

type DescribeStorageResponseBody struct {
	// Indicates whether the projects and Logstores that are created for the threat analysis feature exist in Simple Log Service. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CCEEE128-6607-503E-AAA6-C5E57D94****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeStorageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeStorageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeStorageResponseBody) SetData(v bool) *DescribeStorageResponseBody {
	s.Data = &v
	return s
}

func (s *DescribeStorageResponseBody) SetRequestId(v string) *DescribeStorageResponseBody {
	s.RequestId = &v
	return s
}

type DescribeStorageResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeStorageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeStorageResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeStorageResponse) GoString() string {
	return s.String()
}

func (s *DescribeStorageResponse) SetHeaders(v map[string]*string) *DescribeStorageResponse {
	s.Headers = v
	return s
}

func (s *DescribeStorageResponse) SetStatusCode(v int32) *DescribeStorageResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeStorageResponse) SetBody(v *DescribeStorageResponseBody) *DescribeStorageResponse {
	s.Body = v
	return s
}

type DescribeUserBuyStatusRequest struct {
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
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 123XXXXXX
	SubUserId *int64 `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
}

func (s DescribeUserBuyStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserBuyStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserBuyStatusRequest) SetRegionId(v string) *DescribeUserBuyStatusRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeUserBuyStatusRequest) SetSubUserId(v int64) *DescribeUserBuyStatusRequest {
	s.SubUserId = &v
	return s
}

type DescribeUserBuyStatusResponseBody struct {
	// The data returned.
	Data *DescribeUserBuyStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 81D8EC0C-0804-51AD-8C38-17ED0BC74892
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeUserBuyStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserBuyStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserBuyStatusResponseBody) SetData(v *DescribeUserBuyStatusResponseBodyData) *DescribeUserBuyStatusResponseBody {
	s.Data = v
	return s
}

func (s *DescribeUserBuyStatusResponseBody) SetRequestId(v string) *DescribeUserBuyStatusResponseBody {
	s.RequestId = &v
	return s
}

type DescribeUserBuyStatusResponseBodyData struct {
	// Indicates whether the logon Alibaba Cloud account can be used to place orders for the threat analysis feature, such as purchase, upgrade, and specifications change orders. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	CanBuy *bool `json:"CanBuy,omitempty" xml:"CanBuy,omitempty"`
	// The log storage capacity that is purchased for the threat analysis feature. Unit: GB.
	//
	// example:
	//
	// 1024
	Capacity *int32 `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	// The number of days before the expiration time of the threat analysis feature.
	//
	// example:
	//
	// 3
	DurationDays *int64 `json:"DurationDays,omitempty" xml:"DurationDays,omitempty"`
	// The timestamp when the threat analysis feature expires. Unit: milliseconds.
	//
	// example:
	//
	// 1669823999000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the Alibaba Cloud account that is used to purchase the threat analysis feature.
	//
	// example:
	//
	// 123XXXXXX
	MainUserId *int64 `json:"MainUserId,omitempty" xml:"MainUserId,omitempty"`
	// The username of the Alibaba Cloud account that is used to purchase the threat analysis feature.
	//
	// example:
	//
	// sas_account_xxx
	MainUserName *string `json:"MainUserName,omitempty" xml:"MainUserName,omitempty"`
	// The ID of the management account of the resource directory.
	//
	// example:
	//
	// 123XXXXXX
	MasterUserId *int64 `json:"MasterUserId,omitempty" xml:"MasterUserId,omitempty"`
	// The display name of the management account of the resource directory.
	//
	// example:
	//
	// rd_master_xxx
	MasterUserName *string `json:"MasterUserName,omitempty" xml:"MasterUserName,omitempty"`
	// example:
	//
	// 1
	RdOrder *int32 `json:"RdOrder,omitempty" xml:"RdOrder,omitempty"`
	// The instance ID of Security Center.
	//
	// example:
	//
	// sas-instance-xxxxx
	SasInstanceId *string `json:"SasInstanceId,omitempty" xml:"SasInstanceId,omitempty"`
	// The ID of the logon Alibaba Cloud account.
	//
	// example:
	//
	// 123XXXXXX
	SubUserId *int64 `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
	// The username of the logon Alibaba Cloud account.
	//
	// example:
	//
	// sas_account_xxx
	SubUserName *string `json:"SubUserName,omitempty" xml:"SubUserName,omitempty"`
}

func (s DescribeUserBuyStatusResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserBuyStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeUserBuyStatusResponseBodyData) SetCanBuy(v bool) *DescribeUserBuyStatusResponseBodyData {
	s.CanBuy = &v
	return s
}

func (s *DescribeUserBuyStatusResponseBodyData) SetCapacity(v int32) *DescribeUserBuyStatusResponseBodyData {
	s.Capacity = &v
	return s
}

func (s *DescribeUserBuyStatusResponseBodyData) SetDurationDays(v int64) *DescribeUserBuyStatusResponseBodyData {
	s.DurationDays = &v
	return s
}

func (s *DescribeUserBuyStatusResponseBodyData) SetEndTime(v int64) *DescribeUserBuyStatusResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *DescribeUserBuyStatusResponseBodyData) SetMainUserId(v int64) *DescribeUserBuyStatusResponseBodyData {
	s.MainUserId = &v
	return s
}

func (s *DescribeUserBuyStatusResponseBodyData) SetMainUserName(v string) *DescribeUserBuyStatusResponseBodyData {
	s.MainUserName = &v
	return s
}

func (s *DescribeUserBuyStatusResponseBodyData) SetMasterUserId(v int64) *DescribeUserBuyStatusResponseBodyData {
	s.MasterUserId = &v
	return s
}

func (s *DescribeUserBuyStatusResponseBodyData) SetMasterUserName(v string) *DescribeUserBuyStatusResponseBodyData {
	s.MasterUserName = &v
	return s
}

func (s *DescribeUserBuyStatusResponseBodyData) SetRdOrder(v int32) *DescribeUserBuyStatusResponseBodyData {
	s.RdOrder = &v
	return s
}

func (s *DescribeUserBuyStatusResponseBodyData) SetSasInstanceId(v string) *DescribeUserBuyStatusResponseBodyData {
	s.SasInstanceId = &v
	return s
}

func (s *DescribeUserBuyStatusResponseBodyData) SetSubUserId(v int64) *DescribeUserBuyStatusResponseBodyData {
	s.SubUserId = &v
	return s
}

func (s *DescribeUserBuyStatusResponseBodyData) SetSubUserName(v string) *DescribeUserBuyStatusResponseBodyData {
	s.SubUserName = &v
	return s
}

type DescribeUserBuyStatusResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUserBuyStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUserBuyStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserBuyStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserBuyStatusResponse) SetHeaders(v map[string]*string) *DescribeUserBuyStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserBuyStatusResponse) SetStatusCode(v int32) *DescribeUserBuyStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserBuyStatusResponse) SetBody(v *DescribeUserBuyStatusResponseBody) *DescribeUserBuyStatusResponse {
	s.Body = v
	return s
}

type DescribeWafScopeRequest struct {
	// The ID of the entity.
	//
	// example:
	//
	// 20617784
	EntityId *int64 `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
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
	// The ID of the account that you switch from the management account.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view. Valid values:
	//
	// - 0: the current Alibaba Cloud account
	//
	// - 1: the global account
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s DescribeWafScopeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeWafScopeRequest) GoString() string {
	return s.String()
}

func (s *DescribeWafScopeRequest) SetEntityId(v int64) *DescribeWafScopeRequest {
	s.EntityId = &v
	return s
}

func (s *DescribeWafScopeRequest) SetRegionId(v string) *DescribeWafScopeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeWafScopeRequest) SetRoleFor(v int64) *DescribeWafScopeRequest {
	s.RoleFor = &v
	return s
}

func (s *DescribeWafScopeRequest) SetRoleType(v int32) *DescribeWafScopeRequest {
	s.RoleType = &v
	return s
}

type DescribeWafScopeResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	//
	// example:
	//
	// 123456
	Data []*DescribeWafScopeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeWafScopeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeWafScopeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWafScopeResponseBody) SetCode(v int32) *DescribeWafScopeResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeWafScopeResponseBody) SetData(v []*DescribeWafScopeResponseBodyData) *DescribeWafScopeResponseBody {
	s.Data = v
	return s
}

func (s *DescribeWafScopeResponseBody) SetMessage(v string) *DescribeWafScopeResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeWafScopeResponseBody) SetRequestId(v string) *DescribeWafScopeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWafScopeResponseBody) SetSuccess(v bool) *DescribeWafScopeResponseBody {
	s.Success = &v
	return s
}

type DescribeWafScopeResponseBodyData struct {
	// The ID of the Alibaba Cloud account in SIEM.
	//
	// example:
	//
	// 127608589417****
	Aliuid *int64 `json:"Aliuid,omitempty" xml:"Aliuid,omitempty"`
	// The domain names that are protected by the WAF instance.
	//
	// example:
	//
	// [123.com, 456.com]
	Domains []*string `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
	// The ID of the WAF instance.
	//
	// example:
	//
	// waf-cn-tl123ast****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeWafScopeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeWafScopeResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeWafScopeResponseBodyData) SetAliuid(v int64) *DescribeWafScopeResponseBodyData {
	s.Aliuid = &v
	return s
}

func (s *DescribeWafScopeResponseBodyData) SetDomains(v []*string) *DescribeWafScopeResponseBodyData {
	s.Domains = v
	return s
}

func (s *DescribeWafScopeResponseBodyData) SetInstanceId(v string) *DescribeWafScopeResponseBodyData {
	s.InstanceId = &v
	return s
}

type DescribeWafScopeResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeWafScopeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeWafScopeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeWafScopeResponse) GoString() string {
	return s.String()
}

func (s *DescribeWafScopeResponse) SetHeaders(v map[string]*string) *DescribeWafScopeResponse {
	s.Headers = v
	return s
}

func (s *DescribeWafScopeResponse) SetStatusCode(v int32) *DescribeWafScopeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeWafScopeResponse) SetBody(v *DescribeWafScopeResponseBody) *DescribeWafScopeResponse {
	s.Body = v
	return s
}

type DescribeWhiteRuleListRequest struct {
	// The name of the alert.
	//
	// example:
	//
	// Try SNMP weak password
	AlertName *string `json:"AlertName,omitempty" xml:"AlertName,omitempty"`
	// The type of the alert.
	//
	// example:
	//
	// scan
	AlertType *string `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
	// The page number. Pages start from page 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The UUID of the event.
	//
	// example:
	//
	// 85ea4241-798f-4684-a876-65d4f0c3****
	IncidentUuid *string `json:"IncidentUuid,omitempty" xml:"IncidentUuid,omitempty"`
	// The number of entries per page. Valid values: 1 to 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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
	// The ID of the account that you switch from the management account.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view. Valid values:
	//
	// - 0: the current Alibaba Cloud account
	//
	// - 1: the global account
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s DescribeWhiteRuleListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeWhiteRuleListRequest) GoString() string {
	return s.String()
}

func (s *DescribeWhiteRuleListRequest) SetAlertName(v string) *DescribeWhiteRuleListRequest {
	s.AlertName = &v
	return s
}

func (s *DescribeWhiteRuleListRequest) SetAlertType(v string) *DescribeWhiteRuleListRequest {
	s.AlertType = &v
	return s
}

func (s *DescribeWhiteRuleListRequest) SetCurrentPage(v int32) *DescribeWhiteRuleListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeWhiteRuleListRequest) SetIncidentUuid(v string) *DescribeWhiteRuleListRequest {
	s.IncidentUuid = &v
	return s
}

func (s *DescribeWhiteRuleListRequest) SetPageSize(v int32) *DescribeWhiteRuleListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeWhiteRuleListRequest) SetRegionId(v string) *DescribeWhiteRuleListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeWhiteRuleListRequest) SetRoleFor(v int64) *DescribeWhiteRuleListRequest {
	s.RoleFor = &v
	return s
}

func (s *DescribeWhiteRuleListRequest) SetRoleType(v int32) *DescribeWhiteRuleListRequest {
	s.RoleType = &v
	return s
}

type DescribeWhiteRuleListResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	//
	// example:
	//
	// 123456
	Data *DescribeWhiteRuleListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeWhiteRuleListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeWhiteRuleListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWhiteRuleListResponseBody) SetCode(v int32) *DescribeWhiteRuleListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeWhiteRuleListResponseBody) SetData(v *DescribeWhiteRuleListResponseBodyData) *DescribeWhiteRuleListResponseBody {
	s.Data = v
	return s
}

func (s *DescribeWhiteRuleListResponseBody) SetMessage(v string) *DescribeWhiteRuleListResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeWhiteRuleListResponseBody) SetRequestId(v string) *DescribeWhiteRuleListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWhiteRuleListResponseBody) SetSuccess(v bool) *DescribeWhiteRuleListResponseBody {
	s.Success = &v
	return s
}

type DescribeWhiteRuleListResponseBodyData struct {
	// The pagination information.
	PageInfo *DescribeWhiteRuleListResponseBodyDataPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The detailed data.
	ResponseData []*DescribeWhiteRuleListResponseBodyDataResponseData `json:"ResponseData,omitempty" xml:"ResponseData,omitempty" type:"Repeated"`
}

func (s DescribeWhiteRuleListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeWhiteRuleListResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeWhiteRuleListResponseBodyData) SetPageInfo(v *DescribeWhiteRuleListResponseBodyDataPageInfo) *DescribeWhiteRuleListResponseBodyData {
	s.PageInfo = v
	return s
}

func (s *DescribeWhiteRuleListResponseBodyData) SetResponseData(v []*DescribeWhiteRuleListResponseBodyDataResponseData) *DescribeWhiteRuleListResponseBodyData {
	s.ResponseData = v
	return s
}

type DescribeWhiteRuleListResponseBodyDataPageInfo struct {
	// The current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeWhiteRuleListResponseBodyDataPageInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeWhiteRuleListResponseBodyDataPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeWhiteRuleListResponseBodyDataPageInfo) SetCurrentPage(v int32) *DescribeWhiteRuleListResponseBodyDataPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeWhiteRuleListResponseBodyDataPageInfo) SetPageSize(v int32) *DescribeWhiteRuleListResponseBodyDataPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeWhiteRuleListResponseBodyDataPageInfo) SetTotalCount(v int64) *DescribeWhiteRuleListResponseBodyDataPageInfo {
	s.TotalCount = &v
	return s
}

type DescribeWhiteRuleListResponseBodyDataResponseData struct {
	// The alert name.
	//
	// example:
	//
	// Try SNMP weak password
	AlertName *string `json:"AlertName,omitempty" xml:"AlertName,omitempty"`
	// The ID of the alert name.
	//
	// example:
	//
	// Try SNMP weak password
	AlertNameId *string `json:"AlertNameId,omitempty" xml:"AlertNameId,omitempty"`
	// The alert type.
	//
	// example:
	//
	// scan
	AlertType *string `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
	// The ID of the alert type.
	//
	// example:
	//
	// scan
	AlertTypeId *string `json:"AlertTypeId,omitempty" xml:"AlertTypeId,omitempty"`
	// The UUID of the alert.
	//
	// example:
	//
	// sas_71e24437d2797ce8fc59692905a4****
	AlertUuid *string `json:"AlertUuid,omitempty" xml:"AlertUuid,omitempty"`
	// The ID of the Alibaba Cloud account that is used to purchase the threat analysis feature.
	//
	// example:
	//
	// 127608589417****
	Aliuid *int64 `json:"Aliuid,omitempty" xml:"Aliuid,omitempty"`
	// The conditions in the rule. The value is a JSON array.
	//
	// example:
	//
	// [{"conditions":[{"isNot":false,"itemId":0,"left":{"value":"host_uuid.host_name"},"operator":"containsString","right":{"value":"Cloud-MCH"}}]}]
	Expression *DescribeWhiteRuleListResponseBodyDataResponseDataExpression `json:"Expression,omitempty" xml:"Expression,omitempty" type:"Struct"`
	// The time when the whitelist rule was created.
	//
	// example:
	//
	// 2021-01-06 16:37:29
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the whitelist rule was modified.
	//
	// example:
	//
	// 2021-01-06 16:37:29
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the whitelist rule.
	//
	// example:
	//
	// 123456789
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The UUID of the event.
	//
	// example:
	//
	// 85ea4241-798f-4684-a876-65d4f0c3****
	IncidentUuid *string `json:"IncidentUuid,omitempty" xml:"IncidentUuid,omitempty"`
	// The status of the whitelist rule. Valid values:
	//
	// 	- 1: enabled.
	//
	// 	- 0: disabled.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the Alibaba Cloud account that is used to create the whitelist rule.
	//
	// example:
	//
	// 176555323***
	SubAliuid *int64 `json:"SubAliuid,omitempty" xml:"SubAliuid,omitempty"`
}

func (s DescribeWhiteRuleListResponseBodyDataResponseData) String() string {
	return tea.Prettify(s)
}

func (s DescribeWhiteRuleListResponseBodyDataResponseData) GoString() string {
	return s.String()
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseData) SetAlertName(v string) *DescribeWhiteRuleListResponseBodyDataResponseData {
	s.AlertName = &v
	return s
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseData) SetAlertNameId(v string) *DescribeWhiteRuleListResponseBodyDataResponseData {
	s.AlertNameId = &v
	return s
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseData) SetAlertType(v string) *DescribeWhiteRuleListResponseBodyDataResponseData {
	s.AlertType = &v
	return s
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseData) SetAlertTypeId(v string) *DescribeWhiteRuleListResponseBodyDataResponseData {
	s.AlertTypeId = &v
	return s
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseData) SetAlertUuid(v string) *DescribeWhiteRuleListResponseBodyDataResponseData {
	s.AlertUuid = &v
	return s
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseData) SetAliuid(v int64) *DescribeWhiteRuleListResponseBodyDataResponseData {
	s.Aliuid = &v
	return s
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseData) SetExpression(v *DescribeWhiteRuleListResponseBodyDataResponseDataExpression) *DescribeWhiteRuleListResponseBodyDataResponseData {
	s.Expression = v
	return s
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseData) SetGmtCreate(v string) *DescribeWhiteRuleListResponseBodyDataResponseData {
	s.GmtCreate = &v
	return s
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseData) SetGmtModified(v string) *DescribeWhiteRuleListResponseBodyDataResponseData {
	s.GmtModified = &v
	return s
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseData) SetId(v int64) *DescribeWhiteRuleListResponseBodyDataResponseData {
	s.Id = &v
	return s
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseData) SetIncidentUuid(v string) *DescribeWhiteRuleListResponseBodyDataResponseData {
	s.IncidentUuid = &v
	return s
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseData) SetStatus(v int32) *DescribeWhiteRuleListResponseBodyDataResponseData {
	s.Status = &v
	return s
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseData) SetSubAliuid(v int64) *DescribeWhiteRuleListResponseBodyDataResponseData {
	s.SubAliuid = &v
	return s
}

type DescribeWhiteRuleListResponseBodyDataResponseDataExpression struct {
	// The rule conditions.
	Conditions []*DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditions `json:"Conditions,omitempty" xml:"Conditions,omitempty" type:"Repeated"`
	// The logical relationships among the rule conditions.
	//
	// example:
	//
	// (1&2)|(3&4)
	Logic *string `json:"Logic,omitempty" xml:"Logic,omitempty"`
}

func (s DescribeWhiteRuleListResponseBodyDataResponseDataExpression) String() string {
	return tea.Prettify(s)
}

func (s DescribeWhiteRuleListResponseBodyDataResponseDataExpression) GoString() string {
	return s.String()
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseDataExpression) SetConditions(v []*DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditions) *DescribeWhiteRuleListResponseBodyDataResponseDataExpression {
	s.Conditions = v
	return s
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseDataExpression) SetLogic(v string) *DescribeWhiteRuleListResponseBodyDataResponseDataExpression {
	s.Logic = &v
	return s
}

type DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditions struct {
	// Indicates whether the result is inverted. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	IsNot *bool `json:"IsNot,omitempty" xml:"IsNot,omitempty"`
	// The ID of the rule condition.
	//
	// example:
	//
	// 1
	ItemId *int32 `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	// The left operand of the rule condition.
	Left *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditionsLeft `json:"Left,omitempty" xml:"Left,omitempty" type:"Struct"`
	// The logical operator of the rule condition. Valid values:
	//
	// 	- `=`: equals to.
	//
	// 	- `<>`: does not equal to.
	//
	// 	- `in`: contains.
	//
	// 	- `not in`: does not contain.
	//
	// 	- `REGEXP`: matches a regular expression.
	//
	// 	- `NOT REGEXP`: does not match a regular expression.
	//
	// example:
	//
	// REGEXP
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// The right operand of the rule condition.
	Right *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditionsRight `json:"Right,omitempty" xml:"Right,omitempty" type:"Struct"`
}

func (s DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditions) String() string {
	return tea.Prettify(s)
}

func (s DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditions) GoString() string {
	return s.String()
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditions) SetIsNot(v bool) *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditions {
	s.IsNot = &v
	return s
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditions) SetItemId(v int32) *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditions {
	s.ItemId = &v
	return s
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditions) SetLeft(v *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditionsLeft) *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditions {
	s.Left = v
	return s
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditions) SetOperator(v string) *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditions {
	s.Operator = &v
	return s
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditions) SetRight(v *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditionsRight) *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditions {
	s.Right = v
	return s
}

type DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditionsLeft struct {
	// Indicates whether the left operand is a variable. Valid values:
	//
	// 	- true: variable.
	//
	// 	- false: constant.
	//
	// example:
	//
	// true
	IsVar *bool `json:"IsVar,omitempty" xml:"IsVar,omitempty"`
	// The remarks on the right operand.
	//
	// example:
	//
	// length
	Modifier *string `json:"Modifier,omitempty" xml:"Modifier,omitempty"`
	// The key-value pair information of the remarks.
	ModifierParam map[string]interface{} `json:"ModifierParam,omitempty" xml:"ModifierParam,omitempty"`
	// Indicates whether the left operand is a constant. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The variable of the left operand.
	//
	// example:
	//
	// ip
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditionsLeft) String() string {
	return tea.Prettify(s)
}

func (s DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditionsLeft) GoString() string {
	return s.String()
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditionsLeft) SetIsVar(v bool) *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditionsLeft {
	s.IsVar = &v
	return s
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditionsLeft) SetModifier(v string) *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditionsLeft {
	s.Modifier = &v
	return s
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditionsLeft) SetModifierParam(v map[string]interface{}) *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditionsLeft {
	s.ModifierParam = v
	return s
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditionsLeft) SetType(v string) *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditionsLeft {
	s.Type = &v
	return s
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditionsLeft) SetValue(v string) *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditionsLeft {
	s.Value = &v
	return s
}

type DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditionsRight struct {
	// Indicates whether the right operand is a constant or a runtime variable that is obtained from the runtime context. Valid values:
	//
	// 	- true: runtime variable.
	//
	// 	- false: constant.
	//
	// example:
	//
	// false
	IsVar *bool `json:"IsVar,omitempty" xml:"IsVar,omitempty"`
	// The remarks on the right operand.
	//
	// example:
	//
	// length
	Modifier *string `json:"Modifier,omitempty" xml:"Modifier,omitempty"`
	// The key-value pair information of the remarks.
	ModifierParam map[string]interface{} `json:"ModifierParam,omitempty" xml:"ModifierParam,omitempty"`
	// The data type of the right operand.
	//
	// example:
	//
	// String
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The right operand.
	//
	// example:
	//
	// 12345
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditionsRight) String() string {
	return tea.Prettify(s)
}

func (s DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditionsRight) GoString() string {
	return s.String()
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditionsRight) SetIsVar(v bool) *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditionsRight {
	s.IsVar = &v
	return s
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditionsRight) SetModifier(v string) *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditionsRight {
	s.Modifier = &v
	return s
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditionsRight) SetModifierParam(v map[string]interface{}) *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditionsRight {
	s.ModifierParam = v
	return s
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditionsRight) SetType(v string) *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditionsRight {
	s.Type = &v
	return s
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditionsRight) SetValue(v string) *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditionsRight {
	s.Value = &v
	return s
}

type DescribeWhiteRuleListResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeWhiteRuleListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeWhiteRuleListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeWhiteRuleListResponse) GoString() string {
	return s.String()
}

func (s *DescribeWhiteRuleListResponse) SetHeaders(v map[string]*string) *DescribeWhiteRuleListResponse {
	s.Headers = v
	return s
}

func (s *DescribeWhiteRuleListResponse) SetStatusCode(v int32) *DescribeWhiteRuleListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeWhiteRuleListResponse) SetBody(v *DescribeWhiteRuleListResponseBody) *DescribeWhiteRuleListResponse {
	s.Body = v
	return s
}

type EnableAccessForCloudSiemRequest struct {
	// Whether import the log of SAS alert, the log of WAF alert, the log of CFW alert or not. Valid values:
	//
	// - 0: not imported automatically
	//
	// - 1: imported automatically
	//
	// example:
	//
	// 1
	AutoSubmit *int32 `json:"AutoSubmit,omitempty" xml:"AutoSubmit,omitempty"`
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
	// The ID of the account that you switch from the management account.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view. Valid values:
	//
	// - 0: the current Alibaba Cloud account
	//
	// - 1: the global account
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s EnableAccessForCloudSiemRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableAccessForCloudSiemRequest) GoString() string {
	return s.String()
}

func (s *EnableAccessForCloudSiemRequest) SetAutoSubmit(v int32) *EnableAccessForCloudSiemRequest {
	s.AutoSubmit = &v
	return s
}

func (s *EnableAccessForCloudSiemRequest) SetRegionId(v string) *EnableAccessForCloudSiemRequest {
	s.RegionId = &v
	return s
}

func (s *EnableAccessForCloudSiemRequest) SetRoleFor(v int64) *EnableAccessForCloudSiemRequest {
	s.RoleFor = &v
	return s
}

func (s *EnableAccessForCloudSiemRequest) SetRoleType(v int32) *EnableAccessForCloudSiemRequest {
	s.RoleType = &v
	return s
}

type EnableAccessForCloudSiemResponseBody struct {
	// The data returned.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableAccessForCloudSiemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableAccessForCloudSiemResponseBody) GoString() string {
	return s.String()
}

func (s *EnableAccessForCloudSiemResponseBody) SetData(v bool) *EnableAccessForCloudSiemResponseBody {
	s.Data = &v
	return s
}

func (s *EnableAccessForCloudSiemResponseBody) SetRequestId(v string) *EnableAccessForCloudSiemResponseBody {
	s.RequestId = &v
	return s
}

type EnableAccessForCloudSiemResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnableAccessForCloudSiemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableAccessForCloudSiemResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableAccessForCloudSiemResponse) GoString() string {
	return s.String()
}

func (s *EnableAccessForCloudSiemResponse) SetHeaders(v map[string]*string) *EnableAccessForCloudSiemResponse {
	s.Headers = v
	return s
}

func (s *EnableAccessForCloudSiemResponse) SetStatusCode(v int32) *EnableAccessForCloudSiemResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableAccessForCloudSiemResponse) SetBody(v *EnableAccessForCloudSiemResponseBody) *EnableAccessForCloudSiemResponse {
	s.Body = v
	return s
}

type EnableServiceForCloudSiemRequest struct {
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

func (s EnableServiceForCloudSiemRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableServiceForCloudSiemRequest) GoString() string {
	return s.String()
}

func (s *EnableServiceForCloudSiemRequest) SetRegionId(v string) *EnableServiceForCloudSiemRequest {
	s.RegionId = &v
	return s
}

type EnableServiceForCloudSiemResponseBody struct {
	// Indicates whether the threat analysis feature is authorized to access the resource directory. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableServiceForCloudSiemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableServiceForCloudSiemResponseBody) GoString() string {
	return s.String()
}

func (s *EnableServiceForCloudSiemResponseBody) SetData(v bool) *EnableServiceForCloudSiemResponseBody {
	s.Data = &v
	return s
}

func (s *EnableServiceForCloudSiemResponseBody) SetRequestId(v string) *EnableServiceForCloudSiemResponseBody {
	s.RequestId = &v
	return s
}

type EnableServiceForCloudSiemResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnableServiceForCloudSiemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableServiceForCloudSiemResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableServiceForCloudSiemResponse) GoString() string {
	return s.String()
}

func (s *EnableServiceForCloudSiemResponse) SetHeaders(v map[string]*string) *EnableServiceForCloudSiemResponse {
	s.Headers = v
	return s
}

func (s *EnableServiceForCloudSiemResponse) SetStatusCode(v int32) *EnableServiceForCloudSiemResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableServiceForCloudSiemResponse) SetBody(v *EnableServiceForCloudSiemResponseBody) *EnableServiceForCloudSiemResponse {
	s.Body = v
	return s
}

type GetCapacityRequest struct {
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
	// The ID of the account that you switch from the management account.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view. Valid values:
	//
	// - 0: the current Alibaba Cloud account
	//
	// - 1: the global account
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s GetCapacityRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCapacityRequest) GoString() string {
	return s.String()
}

func (s *GetCapacityRequest) SetRegionId(v string) *GetCapacityRequest {
	s.RegionId = &v
	return s
}

func (s *GetCapacityRequest) SetRoleFor(v int64) *GetCapacityRequest {
	s.RoleFor = &v
	return s
}

func (s *GetCapacityRequest) SetRoleType(v int32) *GetCapacityRequest {
	s.RoleType = &v
	return s
}

type GetCapacityResponseBody struct {
	// The information about the storage capacity.
	Data *GetCapacityResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 27D27DCB-D76B-5064-8B3B-0900DEF7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetCapacityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCapacityResponseBody) GoString() string {
	return s.String()
}

func (s *GetCapacityResponseBody) SetData(v *GetCapacityResponseBodyData) *GetCapacityResponseBody {
	s.Data = v
	return s
}

func (s *GetCapacityResponseBody) SetRequestId(v string) *GetCapacityResponseBody {
	s.RequestId = &v
	return s
}

type GetCapacityResponseBodyData struct {
	// Indicates whether the Logstores for the threat analysis feature exist on the user side. Valid values:
	//
	// 	- true: The logs are in the normal state. The log analysis feature is available.
	//
	// 	- false: The logs are being cleared. The log analysis feature is unavailable.
	//
	// example:
	//
	// true
	ExistLogStore *bool `json:"ExistLogStore,omitempty" xml:"ExistLogStore,omitempty"`
	// The purchased storage capacity of the threat analysis feature. Unit: GB.
	//
	// example:
	//
	// 9000
	PreservedCapacity *int64 `json:"PreservedCapacity,omitempty" xml:"PreservedCapacity,omitempty"`
	// The billable storage capacity of the threat analysis feature. Unit: GB.
	//
	// example:
	//
	// 10
	UsedCapacity *float64 `json:"UsedCapacity,omitempty" xml:"UsedCapacity,omitempty"`
}

func (s GetCapacityResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetCapacityResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCapacityResponseBodyData) SetExistLogStore(v bool) *GetCapacityResponseBodyData {
	s.ExistLogStore = &v
	return s
}

func (s *GetCapacityResponseBodyData) SetPreservedCapacity(v int64) *GetCapacityResponseBodyData {
	s.PreservedCapacity = &v
	return s
}

func (s *GetCapacityResponseBodyData) SetUsedCapacity(v float64) *GetCapacityResponseBodyData {
	s.UsedCapacity = &v
	return s
}

type GetCapacityResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCapacityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCapacityResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCapacityResponse) GoString() string {
	return s.String()
}

func (s *GetCapacityResponse) SetHeaders(v map[string]*string) *GetCapacityResponse {
	s.Headers = v
	return s
}

func (s *GetCapacityResponse) SetStatusCode(v int32) *GetCapacityResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCapacityResponse) SetBody(v *GetCapacityResponseBody) *GetCapacityResponse {
	s.Body = v
	return s
}

type GetStorageRequest struct {
	// The data management center of the threat analysis feature. Specify this parameter based on the region where your assets reside. Valid values:
	//
	// 	- cn-hangzhou: Your assets reside in regions in China.
	//
	// 	- ap-southeast-1: Your assets reside in regions outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the account that you switch from the management account.
	//
	// example:
	//
	// 127XXXX
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view. Valid values:
	//
	// - 0: the current Alibaba Cloud account
	//
	// - 1: the global account
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s GetStorageRequest) String() string {
	return tea.Prettify(s)
}

func (s GetStorageRequest) GoString() string {
	return s.String()
}

func (s *GetStorageRequest) SetRegionId(v string) *GetStorageRequest {
	s.RegionId = &v
	return s
}

func (s *GetStorageRequest) SetRoleFor(v int64) *GetStorageRequest {
	s.RoleFor = &v
	return s
}

func (s *GetStorageRequest) SetRoleType(v int32) *GetStorageRequest {
	s.RoleType = &v
	return s
}

type GetStorageResponseBody struct {
	// The information about the storage.
	Data *GetStorageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 97A31C3A-3F9F-5866-8979-5159E3DC****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetStorageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetStorageResponseBody) GoString() string {
	return s.String()
}

func (s *GetStorageResponseBody) SetData(v *GetStorageResponseBodyData) *GetStorageResponseBody {
	s.Data = v
	return s
}

func (s *GetStorageResponseBody) SetRequestId(v string) *GetStorageResponseBody {
	s.RequestId = &v
	return s
}

type GetStorageResponseBodyData struct {
	// Indicates whether the storage region can be changed for once. Default value: false Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	CanOperate *bool `json:"CanOperate,omitempty" xml:"CanOperate,omitempty"`
	// Indicates whether the storage region can be changed. Default value: false Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	DisplayRegion *bool `json:"DisplayRegion,omitempty" xml:"DisplayRegion,omitempty"`
	// The region where the data is stored.
	//
	// If the data management center is **cn-hangzhou**, the default value of **Region*	- is cn-shanghai, which specifies the China (Shanghai) region. If the data management center is **ap-southeast-1**, the default value of **Region*	- is ap-southeast-1, which specifies the Singapore region.
	//
	// example:
	//
	// cn-shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The storage period of logs. Unit: day. Default value: 180. Valid values: 30 to 3000.
	//
	// example:
	//
	// 180
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
}

func (s GetStorageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetStorageResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetStorageResponseBodyData) SetCanOperate(v bool) *GetStorageResponseBodyData {
	s.CanOperate = &v
	return s
}

func (s *GetStorageResponseBodyData) SetDisplayRegion(v bool) *GetStorageResponseBodyData {
	s.DisplayRegion = &v
	return s
}

func (s *GetStorageResponseBodyData) SetRegion(v string) *GetStorageResponseBodyData {
	s.Region = &v
	return s
}

func (s *GetStorageResponseBodyData) SetTtl(v int32) *GetStorageResponseBodyData {
	s.Ttl = &v
	return s
}

type GetStorageResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetStorageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetStorageResponse) String() string {
	return tea.Prettify(s)
}

func (s GetStorageResponse) GoString() string {
	return s.String()
}

func (s *GetStorageResponse) SetHeaders(v map[string]*string) *GetStorageResponse {
	s.Headers = v
	return s
}

func (s *GetStorageResponse) SetStatusCode(v int32) *GetStorageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetStorageResponse) SetBody(v *GetStorageResponseBody) *GetStorageResponse {
	s.Body = v
	return s
}

type ListAccountAccessIdRequest struct {
	// The code of the cloud service provider.
	//
	// Valid values:
	//
	// 	- qcloud
	//
	// 	- hcloud
	//
	// This parameter is required.
	//
	// example:
	//
	// hcloud
	CloudCode *string `json:"CloudCode,omitempty" xml:"CloudCode,omitempty"`
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
	// The ID of the account that you switch from the management account.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view. Valid values:
	//
	// - 0: the current Alibaba Cloud account
	//
	// - 1: the global account
	//
	// example:
	//
	// 0
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s ListAccountAccessIdRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAccountAccessIdRequest) GoString() string {
	return s.String()
}

func (s *ListAccountAccessIdRequest) SetCloudCode(v string) *ListAccountAccessIdRequest {
	s.CloudCode = &v
	return s
}

func (s *ListAccountAccessIdRequest) SetRegionId(v string) *ListAccountAccessIdRequest {
	s.RegionId = &v
	return s
}

func (s *ListAccountAccessIdRequest) SetRoleFor(v int64) *ListAccountAccessIdRequest {
	s.RoleFor = &v
	return s
}

func (s *ListAccountAccessIdRequest) SetRoleType(v int32) *ListAccountAccessIdRequest {
	s.RoleType = &v
	return s
}

type ListAccountAccessIdResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data []*ListAccountAccessIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListAccountAccessIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAccountAccessIdResponseBody) GoString() string {
	return s.String()
}

func (s *ListAccountAccessIdResponseBody) SetCode(v int32) *ListAccountAccessIdResponseBody {
	s.Code = &v
	return s
}

func (s *ListAccountAccessIdResponseBody) SetData(v []*ListAccountAccessIdResponseBodyData) *ListAccountAccessIdResponseBody {
	s.Data = v
	return s
}

func (s *ListAccountAccessIdResponseBody) SetMessage(v string) *ListAccountAccessIdResponseBody {
	s.Message = &v
	return s
}

func (s *ListAccountAccessIdResponseBody) SetRequestId(v string) *ListAccountAccessIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAccountAccessIdResponseBody) SetSuccess(v bool) *ListAccountAccessIdResponseBody {
	s.Success = &v
	return s
}

type ListAccountAccessIdResponseBodyData struct {
	// The AccessKey ID of the cloud account that is added to the threat analysis feature.
	//
	// example:
	//
	// ABCXXXXXXXX
	AccessId *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	// The MD5 hash value of the AccessKey ID.
	//
	// example:
	//
	// abcXXXXXXXX
	AccessIdMd5 *string `json:"AccessIdMd5,omitempty" xml:"AccessIdMd5,omitempty"`
	// The ID of the cloud account.
	//
	// example:
	//
	// 123xxxxxxx
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The information about the cloud account to which the AccessKey ID belongs. The value is in the following format: Alibaba Cloud account ID|Alibaba Cloud account username|AccessKey ID.
	//
	// example:
	//
	// 123xxxxxx|xxxx|ABCXXXXX
	AccountStr *string `json:"AccountStr,omitempty" xml:"AccountStr,omitempty"`
	// Indicates whether the cloud account to which the AccessKey ID belongs is added to the threat analysis feature. Valid values:
	//
	// 	- 0: no
	//
	// 	- 1: yes
	//
	// example:
	//
	// 1
	Bound *int32 `json:"Bound,omitempty" xml:"Bound,omitempty"`
	// The code of the cloud service provider.
	//
	// example:
	//
	// hcloud
	CloudCode *string `json:"CloudCode,omitempty" xml:"CloudCode,omitempty"`
	// The ID of the Alibaba Cloud account that is used to add the third-party cloud account.
	//
	// example:
	//
	// ABCXXXXXXXX
	SubUserId *int64 `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
}

func (s ListAccountAccessIdResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListAccountAccessIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAccountAccessIdResponseBodyData) SetAccessId(v string) *ListAccountAccessIdResponseBodyData {
	s.AccessId = &v
	return s
}

func (s *ListAccountAccessIdResponseBodyData) SetAccessIdMd5(v string) *ListAccountAccessIdResponseBodyData {
	s.AccessIdMd5 = &v
	return s
}

func (s *ListAccountAccessIdResponseBodyData) SetAccountId(v string) *ListAccountAccessIdResponseBodyData {
	s.AccountId = &v
	return s
}

func (s *ListAccountAccessIdResponseBodyData) SetAccountStr(v string) *ListAccountAccessIdResponseBodyData {
	s.AccountStr = &v
	return s
}

func (s *ListAccountAccessIdResponseBodyData) SetBound(v int32) *ListAccountAccessIdResponseBodyData {
	s.Bound = &v
	return s
}

func (s *ListAccountAccessIdResponseBodyData) SetCloudCode(v string) *ListAccountAccessIdResponseBodyData {
	s.CloudCode = &v
	return s
}

func (s *ListAccountAccessIdResponseBodyData) SetSubUserId(v int64) *ListAccountAccessIdResponseBodyData {
	s.SubUserId = &v
	return s
}

type ListAccountAccessIdResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAccountAccessIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAccountAccessIdResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAccountAccessIdResponse) GoString() string {
	return s.String()
}

func (s *ListAccountAccessIdResponse) SetHeaders(v map[string]*string) *ListAccountAccessIdResponse {
	s.Headers = v
	return s
}

func (s *ListAccountAccessIdResponse) SetStatusCode(v int32) *ListAccountAccessIdResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAccountAccessIdResponse) SetBody(v *ListAccountAccessIdResponseBody) *ListAccountAccessIdResponse {
	s.Body = v
	return s
}

type ListAccountsByLogRequest struct {
	// The code that is used for multi-cloud environments.
	//
	// This parameter is required.
	//
	// example:
	//
	// hcloud
	CloudCode *string `json:"CloudCode,omitempty" xml:"CloudCode,omitempty"`
	// The codes of logs. The value is a JSON array.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["cloud_siem_hcloud_waf_alert_log"]
	LogCodes []*string `json:"LogCodes,omitempty" xml:"LogCodes,omitempty" type:"Repeated"`
	// The code of the service.
	//
	// This parameter is required.
	//
	// example:
	//
	// qcloud_waf
	ProdCode *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
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
	// The ID of the account that you switch from the management account.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view. Valid values:
	//
	// - 0: the current Alibaba Cloud account
	//
	// - 1: the global account
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s ListAccountsByLogRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAccountsByLogRequest) GoString() string {
	return s.String()
}

func (s *ListAccountsByLogRequest) SetCloudCode(v string) *ListAccountsByLogRequest {
	s.CloudCode = &v
	return s
}

func (s *ListAccountsByLogRequest) SetLogCodes(v []*string) *ListAccountsByLogRequest {
	s.LogCodes = v
	return s
}

func (s *ListAccountsByLogRequest) SetProdCode(v string) *ListAccountsByLogRequest {
	s.ProdCode = &v
	return s
}

func (s *ListAccountsByLogRequest) SetRegionId(v string) *ListAccountsByLogRequest {
	s.RegionId = &v
	return s
}

func (s *ListAccountsByLogRequest) SetRoleFor(v int64) *ListAccountsByLogRequest {
	s.RoleFor = &v
	return s
}

func (s *ListAccountsByLogRequest) SetRoleType(v int32) *ListAccountsByLogRequest {
	s.RoleType = &v
	return s
}

type ListAccountsByLogResponseBody struct {
	// The data returned.
	Data []*ListAccountsByLogResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAccountsByLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAccountsByLogResponseBody) GoString() string {
	return s.String()
}

func (s *ListAccountsByLogResponseBody) SetData(v []*ListAccountsByLogResponseBodyData) *ListAccountsByLogResponseBody {
	s.Data = v
	return s
}

func (s *ListAccountsByLogResponseBody) SetRequestId(v string) *ListAccountsByLogResponseBody {
	s.RequestId = &v
	return s
}

type ListAccountsByLogResponseBodyData struct {
	// The ID of the cloud account.
	//
	// example:
	//
	// 123xxxxxxx
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The name of the cloud account.
	//
	// example:
	//
	// sas_account_xxx
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// Indicates whether the account is added. Valid values: -1: yes -0: no
	//
	// example:
	//
	// 123xxxxxxx
	Imported *int32 `json:"Imported,omitempty" xml:"Imported,omitempty"`
	// The code of the log.
	//
	// example:
	//
	// cloud_siem_waf_xxxxx
	LogCode *string `json:"LogCode,omitempty" xml:"LogCode,omitempty"`
	// The ID of the Alibaba Cloud account that is used to purchase the threat analysis feature.
	//
	// example:
	//
	// 123XXXXXXXXX
	MainUserId *int64 `json:"MainUserId,omitempty" xml:"MainUserId,omitempty"`
	// The code of the service.
	//
	// example:
	//
	// qcloud_waf
	ProdCode *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	// The ID of the Alibaba Cloud account for which the threat analysis feature is enabled.
	//
	// example:
	//
	// 123XXXXXXXX
	SubUserId *int64 `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
}

func (s ListAccountsByLogResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListAccountsByLogResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAccountsByLogResponseBodyData) SetAccountId(v string) *ListAccountsByLogResponseBodyData {
	s.AccountId = &v
	return s
}

func (s *ListAccountsByLogResponseBodyData) SetAccountName(v string) *ListAccountsByLogResponseBodyData {
	s.AccountName = &v
	return s
}

func (s *ListAccountsByLogResponseBodyData) SetImported(v int32) *ListAccountsByLogResponseBodyData {
	s.Imported = &v
	return s
}

func (s *ListAccountsByLogResponseBodyData) SetLogCode(v string) *ListAccountsByLogResponseBodyData {
	s.LogCode = &v
	return s
}

func (s *ListAccountsByLogResponseBodyData) SetMainUserId(v int64) *ListAccountsByLogResponseBodyData {
	s.MainUserId = &v
	return s
}

func (s *ListAccountsByLogResponseBodyData) SetProdCode(v string) *ListAccountsByLogResponseBodyData {
	s.ProdCode = &v
	return s
}

func (s *ListAccountsByLogResponseBodyData) SetSubUserId(v int64) *ListAccountsByLogResponseBodyData {
	s.SubUserId = &v
	return s
}

type ListAccountsByLogResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAccountsByLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAccountsByLogResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAccountsByLogResponse) GoString() string {
	return s.String()
}

func (s *ListAccountsByLogResponse) SetHeaders(v map[string]*string) *ListAccountsByLogResponse {
	s.Headers = v
	return s
}

func (s *ListAccountsByLogResponse) SetStatusCode(v int32) *ListAccountsByLogResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAccountsByLogResponse) SetBody(v *ListAccountsByLogResponseBody) *ListAccountsByLogResponse {
	s.Body = v
	return s
}

type ListAllProdsRequest struct {
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
	// The ID of the account that you switch from the management account.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view. Valid values:
	//
	// - 0: the current Alibaba Cloud account
	//
	// - 1: the global account
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s ListAllProdsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAllProdsRequest) GoString() string {
	return s.String()
}

func (s *ListAllProdsRequest) SetRegionId(v string) *ListAllProdsRequest {
	s.RegionId = &v
	return s
}

func (s *ListAllProdsRequest) SetRoleFor(v int64) *ListAllProdsRequest {
	s.RoleFor = &v
	return s
}

func (s *ListAllProdsRequest) SetRoleType(v int32) *ListAllProdsRequest {
	s.RoleType = &v
	return s
}

type ListAllProdsResponseBody struct {
	// The data returned.
	Data *ListAllProdsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAllProdsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAllProdsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAllProdsResponseBody) SetData(v *ListAllProdsResponseBodyData) *ListAllProdsResponseBody {
	s.Data = v
	return s
}

func (s *ListAllProdsResponseBody) SetRequestId(v string) *ListAllProdsResponseBody {
	s.RequestId = &v
	return s
}

type ListAllProdsResponseBodyData struct {
	// The page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The cloud services.
	//
	// example:
	//
	// 1
	ProdList []*ListAllProdsResponseBodyDataProdList `json:"ProdList,omitempty" xml:"ProdList,omitempty" type:"Repeated"`
	// The total number of logs.
	//
	// example:
	//
	// 19
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAllProdsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListAllProdsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAllProdsResponseBodyData) SetCurrentPage(v int32) *ListAllProdsResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *ListAllProdsResponseBodyData) SetPageSize(v int32) *ListAllProdsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListAllProdsResponseBodyData) SetProdList(v []*ListAllProdsResponseBodyDataProdList) *ListAllProdsResponseBodyData {
	s.ProdList = v
	return s
}

func (s *ListAllProdsResponseBodyData) SetTotalCount(v int32) *ListAllProdsResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListAllProdsResponseBodyDataProdList struct {
	// The code of the cloud service provider. Valid values:
	//
	// 	- qcloud: Tencent Cloud.
	//
	// 	- aliyun: Alibaba Cloud.
	//
	// 	- hcloud: Huawei Cloud.
	//
	// example:
	//
	// hcloud
	CloudCode *string `json:"CloudCode,omitempty" xml:"CloudCode,omitempty"`
	// The number of logs within the cloud service that are added to the threat analysis feature.
	//
	// example:
	//
	// 10
	ImportedLogCount *int32 `json:"ImportedLogCount,omitempty" xml:"ImportedLogCount,omitempty"`
	// The time when the logs within the cloud service were last added to the threat analysis feature.
	//
	// example:
	//
	// 2023-11-23 12:12:12
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The code of the cloud service.
	//
	// example:
	//
	// sas
	ProdCode *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	// The total number of logs within the cloud service.
	//
	// example:
	//
	// 19
	TotalLogCount *int32 `json:"TotalLogCount,omitempty" xml:"TotalLogCount,omitempty"`
}

func (s ListAllProdsResponseBodyDataProdList) String() string {
	return tea.Prettify(s)
}

func (s ListAllProdsResponseBodyDataProdList) GoString() string {
	return s.String()
}

func (s *ListAllProdsResponseBodyDataProdList) SetCloudCode(v string) *ListAllProdsResponseBodyDataProdList {
	s.CloudCode = &v
	return s
}

func (s *ListAllProdsResponseBodyDataProdList) SetImportedLogCount(v int32) *ListAllProdsResponseBodyDataProdList {
	s.ImportedLogCount = &v
	return s
}

func (s *ListAllProdsResponseBodyDataProdList) SetModifyTime(v string) *ListAllProdsResponseBodyDataProdList {
	s.ModifyTime = &v
	return s
}

func (s *ListAllProdsResponseBodyDataProdList) SetProdCode(v string) *ListAllProdsResponseBodyDataProdList {
	s.ProdCode = &v
	return s
}

func (s *ListAllProdsResponseBodyDataProdList) SetTotalLogCount(v int32) *ListAllProdsResponseBodyDataProdList {
	s.TotalLogCount = &v
	return s
}

type ListAllProdsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAllProdsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAllProdsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAllProdsResponse) GoString() string {
	return s.String()
}

func (s *ListAllProdsResponse) SetHeaders(v map[string]*string) *ListAllProdsResponse {
	s.Headers = v
	return s
}

func (s *ListAllProdsResponse) SetStatusCode(v int32) *ListAllProdsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAllProdsResponse) SetBody(v *ListAllProdsResponseBody) *ListAllProdsResponse {
	s.Body = v
	return s
}

type ListAutomateResponseConfigsRequest struct {
	// The type of the handling action. Valid values:
	//
	// 	- doPlaybook: runs a playbook.
	//
	// 	- changeEventStatus: changes the status of an event.
	//
	// 	- changeThreatLevel: changes the risk level of an event.
	//
	// example:
	//
	// doPlaybook
	ActionType *string `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	// The type of the automated response rule. Valid values:
	//
	// 	- event
	//
	// 	- alert
	//
	// example:
	//
	// event
	AutoResponseType *string `json:"AutoResponseType,omitempty" xml:"AutoResponseType,omitempty"`
	// The page number. Pages start from page 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The ID of the automated response rule.
	//
	// example:
	//
	// 123
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The number of entries per page. Maximum value: 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The UUID of the playbook.
	//
	// example:
	//
	// system_aliyun_aegis_kill_quara_book
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
	// The region in which the data management center of the threat analysis feature resides. Specify this parameter based on the regions in which your assets reside. Valid values:
	//
	// 	- cn-hangzhou: Your assets reside in regions in China.
	//
	// 	- ap-southeast-1: Your assets reside in regions outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResponseRuleType *string `json:"ResponseRuleType,omitempty" xml:"ResponseRuleType,omitempty"`
	// The ID of the account that you switch from the management account.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view. Valid values:
	//
	// - 0: the current Alibaba Cloud account
	//
	// - 1: the global account
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// The name of the automated response rule.
	//
	// example:
	//
	// cfw kill quara book
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The status of the rule. Valid values:
	//
	// 	- 0: disabled
	//
	// 	- 100: enabled
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the user who created the rule.
	//
	// example:
	//
	// 17108579417****
	SubUserId *int64 `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
}

func (s ListAutomateResponseConfigsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAutomateResponseConfigsRequest) GoString() string {
	return s.String()
}

func (s *ListAutomateResponseConfigsRequest) SetActionType(v string) *ListAutomateResponseConfigsRequest {
	s.ActionType = &v
	return s
}

func (s *ListAutomateResponseConfigsRequest) SetAutoResponseType(v string) *ListAutomateResponseConfigsRequest {
	s.AutoResponseType = &v
	return s
}

func (s *ListAutomateResponseConfigsRequest) SetCurrentPage(v int32) *ListAutomateResponseConfigsRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListAutomateResponseConfigsRequest) SetId(v int64) *ListAutomateResponseConfigsRequest {
	s.Id = &v
	return s
}

func (s *ListAutomateResponseConfigsRequest) SetPageSize(v int32) *ListAutomateResponseConfigsRequest {
	s.PageSize = &v
	return s
}

func (s *ListAutomateResponseConfigsRequest) SetPlaybookUuid(v string) *ListAutomateResponseConfigsRequest {
	s.PlaybookUuid = &v
	return s
}

func (s *ListAutomateResponseConfigsRequest) SetRegionId(v string) *ListAutomateResponseConfigsRequest {
	s.RegionId = &v
	return s
}

func (s *ListAutomateResponseConfigsRequest) SetResponseRuleType(v string) *ListAutomateResponseConfigsRequest {
	s.ResponseRuleType = &v
	return s
}

func (s *ListAutomateResponseConfigsRequest) SetRoleFor(v int64) *ListAutomateResponseConfigsRequest {
	s.RoleFor = &v
	return s
}

func (s *ListAutomateResponseConfigsRequest) SetRoleType(v int32) *ListAutomateResponseConfigsRequest {
	s.RoleType = &v
	return s
}

func (s *ListAutomateResponseConfigsRequest) SetRuleName(v string) *ListAutomateResponseConfigsRequest {
	s.RuleName = &v
	return s
}

func (s *ListAutomateResponseConfigsRequest) SetStatus(v int32) *ListAutomateResponseConfigsRequest {
	s.Status = &v
	return s
}

func (s *ListAutomateResponseConfigsRequest) SetSubUserId(v int64) *ListAutomateResponseConfigsRequest {
	s.SubUserId = &v
	return s
}

type ListAutomateResponseConfigsResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	//
	// example:
	//
	// 123456
	Data *ListAutomateResponseConfigsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListAutomateResponseConfigsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAutomateResponseConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAutomateResponseConfigsResponseBody) SetCode(v int32) *ListAutomateResponseConfigsResponseBody {
	s.Code = &v
	return s
}

func (s *ListAutomateResponseConfigsResponseBody) SetData(v *ListAutomateResponseConfigsResponseBodyData) *ListAutomateResponseConfigsResponseBody {
	s.Data = v
	return s
}

func (s *ListAutomateResponseConfigsResponseBody) SetMessage(v string) *ListAutomateResponseConfigsResponseBody {
	s.Message = &v
	return s
}

func (s *ListAutomateResponseConfigsResponseBody) SetRequestId(v string) *ListAutomateResponseConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAutomateResponseConfigsResponseBody) SetSuccess(v bool) *ListAutomateResponseConfigsResponseBody {
	s.Success = &v
	return s
}

type ListAutomateResponseConfigsResponseBodyData struct {
	// The pagination information.
	PageInfo *ListAutomateResponseConfigsResponseBodyDataPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The detailed data.
	ResponseData []*ListAutomateResponseConfigsResponseBodyDataResponseData `json:"ResponseData,omitempty" xml:"ResponseData,omitempty" type:"Repeated"`
}

func (s ListAutomateResponseConfigsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListAutomateResponseConfigsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAutomateResponseConfigsResponseBodyData) SetPageInfo(v *ListAutomateResponseConfigsResponseBodyDataPageInfo) *ListAutomateResponseConfigsResponseBodyData {
	s.PageInfo = v
	return s
}

func (s *ListAutomateResponseConfigsResponseBodyData) SetResponseData(v []*ListAutomateResponseConfigsResponseBodyDataResponseData) *ListAutomateResponseConfigsResponseBodyData {
	s.ResponseData = v
	return s
}

type ListAutomateResponseConfigsResponseBodyDataPageInfo struct {
	// The current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAutomateResponseConfigsResponseBodyDataPageInfo) String() string {
	return tea.Prettify(s)
}

func (s ListAutomateResponseConfigsResponseBodyDataPageInfo) GoString() string {
	return s.String()
}

func (s *ListAutomateResponseConfigsResponseBodyDataPageInfo) SetCurrentPage(v int32) *ListAutomateResponseConfigsResponseBodyDataPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *ListAutomateResponseConfigsResponseBodyDataPageInfo) SetPageSize(v int32) *ListAutomateResponseConfigsResponseBodyDataPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListAutomateResponseConfigsResponseBodyDataPageInfo) SetTotalCount(v int64) *ListAutomateResponseConfigsResponseBodyDataPageInfo {
	s.TotalCount = &v
	return s
}

type ListAutomateResponseConfigsResponseBodyDataResponseData struct {
	// The configuration of the action that is performed after the automated response rule is hit. The value is in the JSON format.
	//
	// example:
	//
	// [
	//
	//       {
	//
	//             "actionType": "doPlaybook",
	//
	//             "playbookName": "WafBlockIP",
	//
	//             "playbookUuid": "bdad6220-6584-41b2-9704-fc6584568758"
	//
	//       }
	//
	// ]
	ActionConfig *string `json:"ActionConfig,omitempty" xml:"ActionConfig,omitempty"`
	// The type of the handling action. Multiple types are separated by commas (,). Valid values:
	//
	// 	- **doPlaybook**: runs the playbook.
	//
	// 	- **changeEventStatus**: changes the event status.
	//
	// 	- **changeThreatLevel**: changes the risk level of the event.
	//
	// example:
	//
	// doPlaybook,changeEventStatus
	ActionType *string `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	// The ID of the Alibaba Cloud account that is associated with the rule in SIEM.
	//
	// example:
	//
	// 127608589417****
	Aliuid *int64 `json:"Aliuid,omitempty" xml:"Aliuid,omitempty"`
	// The type of the automated response rule. Valid values:
	//
	// 	- **event**
	//
	// 	- **alert**
	//
	// example:
	//
	// event
	AutoResponseType *string `json:"AutoResponseType,omitempty" xml:"AutoResponseType,omitempty"`
	// The type of the view. Valid values:
	//
	// 0: the current Alibaba Cloud account
	//
	// 1: the global account
	//
	// example:
	//
	// 1
	DataType *int32 `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// The trigger condition of the automated response rule. The value is in the JSON format.
	//
	// example:
	//
	// [{"left":{"value":"alert_name"},"operator":"containsString","right":{"value":"webshell_online"}}]
	ExecutionCondition *string `json:"ExecutionCondition,omitempty" xml:"ExecutionCondition,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2021-01-06 16:37:29
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2021-01-06 16:37:29
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the automated response rule.
	//
	// example:
	//
	// 123
	Id               *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	ResponseRuleType *string `json:"ResponseRuleType,omitempty" xml:"ResponseRuleType,omitempty"`
	// The name of the automated response rule.
	//
	// example:
	//
	// cfw kill quara book
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The status of the rule. Valid values:
	//
	// 	- **0**: disabled.
	//
	// 	- **100**: enabled.
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the user who created the rule.
	//
	// example:
	//
	// 17108579417****
	SubUserId *int64 `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
}

func (s ListAutomateResponseConfigsResponseBodyDataResponseData) String() string {
	return tea.Prettify(s)
}

func (s ListAutomateResponseConfigsResponseBodyDataResponseData) GoString() string {
	return s.String()
}

func (s *ListAutomateResponseConfigsResponseBodyDataResponseData) SetActionConfig(v string) *ListAutomateResponseConfigsResponseBodyDataResponseData {
	s.ActionConfig = &v
	return s
}

func (s *ListAutomateResponseConfigsResponseBodyDataResponseData) SetActionType(v string) *ListAutomateResponseConfigsResponseBodyDataResponseData {
	s.ActionType = &v
	return s
}

func (s *ListAutomateResponseConfigsResponseBodyDataResponseData) SetAliuid(v int64) *ListAutomateResponseConfigsResponseBodyDataResponseData {
	s.Aliuid = &v
	return s
}

func (s *ListAutomateResponseConfigsResponseBodyDataResponseData) SetAutoResponseType(v string) *ListAutomateResponseConfigsResponseBodyDataResponseData {
	s.AutoResponseType = &v
	return s
}

func (s *ListAutomateResponseConfigsResponseBodyDataResponseData) SetDataType(v int32) *ListAutomateResponseConfigsResponseBodyDataResponseData {
	s.DataType = &v
	return s
}

func (s *ListAutomateResponseConfigsResponseBodyDataResponseData) SetExecutionCondition(v string) *ListAutomateResponseConfigsResponseBodyDataResponseData {
	s.ExecutionCondition = &v
	return s
}

func (s *ListAutomateResponseConfigsResponseBodyDataResponseData) SetGmtCreate(v string) *ListAutomateResponseConfigsResponseBodyDataResponseData {
	s.GmtCreate = &v
	return s
}

func (s *ListAutomateResponseConfigsResponseBodyDataResponseData) SetGmtModified(v string) *ListAutomateResponseConfigsResponseBodyDataResponseData {
	s.GmtModified = &v
	return s
}

func (s *ListAutomateResponseConfigsResponseBodyDataResponseData) SetId(v int64) *ListAutomateResponseConfigsResponseBodyDataResponseData {
	s.Id = &v
	return s
}

func (s *ListAutomateResponseConfigsResponseBodyDataResponseData) SetResponseRuleType(v string) *ListAutomateResponseConfigsResponseBodyDataResponseData {
	s.ResponseRuleType = &v
	return s
}

func (s *ListAutomateResponseConfigsResponseBodyDataResponseData) SetRuleName(v string) *ListAutomateResponseConfigsResponseBodyDataResponseData {
	s.RuleName = &v
	return s
}

func (s *ListAutomateResponseConfigsResponseBodyDataResponseData) SetStatus(v int32) *ListAutomateResponseConfigsResponseBodyDataResponseData {
	s.Status = &v
	return s
}

func (s *ListAutomateResponseConfigsResponseBodyDataResponseData) SetSubUserId(v int64) *ListAutomateResponseConfigsResponseBodyDataResponseData {
	s.SubUserId = &v
	return s
}

type ListAutomateResponseConfigsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAutomateResponseConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAutomateResponseConfigsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAutomateResponseConfigsResponse) GoString() string {
	return s.String()
}

func (s *ListAutomateResponseConfigsResponse) SetHeaders(v map[string]*string) *ListAutomateResponseConfigsResponse {
	s.Headers = v
	return s
}

func (s *ListAutomateResponseConfigsResponse) SetStatusCode(v int32) *ListAutomateResponseConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAutomateResponseConfigsResponse) SetBody(v *ListAutomateResponseConfigsResponseBody) *ListAutomateResponseConfigsResponse {
	s.Body = v
	return s
}

type ListBindAccountRequest struct {
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
	RoleFor  *int64  `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	RoleType *int32  `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s ListBindAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s ListBindAccountRequest) GoString() string {
	return s.String()
}

func (s *ListBindAccountRequest) SetCloudCode(v string) *ListBindAccountRequest {
	s.CloudCode = &v
	return s
}

func (s *ListBindAccountRequest) SetRegionId(v string) *ListBindAccountRequest {
	s.RegionId = &v
	return s
}

func (s *ListBindAccountRequest) SetRoleFor(v int64) *ListBindAccountRequest {
	s.RoleFor = &v
	return s
}

func (s *ListBindAccountRequest) SetRoleType(v int32) *ListBindAccountRequest {
	s.RoleType = &v
	return s
}

type ListBindAccountResponseBody struct {
	// The data returned.
	Data []*ListBindAccountResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListBindAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListBindAccountResponseBody) GoString() string {
	return s.String()
}

func (s *ListBindAccountResponseBody) SetData(v []*ListBindAccountResponseBodyData) *ListBindAccountResponseBody {
	s.Data = v
	return s
}

func (s *ListBindAccountResponseBody) SetRequestId(v string) *ListBindAccountResponseBody {
	s.RequestId = &v
	return s
}

type ListBindAccountResponseBodyData struct {
	// The AccessKey ID of the cloud account.
	//
	// example:
	//
	// ABCXXXXXXXX
	AccessId *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
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
	// sas_account_xxx
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The ID that is generated when the cloud account is added.
	//
	// example:
	//
	// 123xxxxxxx
	BindId *int64 `json:"BindId,omitempty" xml:"BindId,omitempty"`
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
	// The ID of the account that is used to add the cloud account.
	//
	// example:
	//
	// 123xxxxxxx
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// The number of data sources that are added to the threat analysis feature within the cloud account.
	//
	// example:
	//
	// 2
	DataSourceCount *int64 `json:"DataSourceCount,omitempty" xml:"DataSourceCount,omitempty"`
	// The modification time.
	//
	// example:
	//
	// 2023-11-10 12:20:35
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
}

func (s ListBindAccountResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListBindAccountResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListBindAccountResponseBodyData) SetAccessId(v string) *ListBindAccountResponseBodyData {
	s.AccessId = &v
	return s
}

func (s *ListBindAccountResponseBodyData) SetAccountId(v string) *ListBindAccountResponseBodyData {
	s.AccountId = &v
	return s
}

func (s *ListBindAccountResponseBodyData) SetAccountName(v string) *ListBindAccountResponseBodyData {
	s.AccountName = &v
	return s
}

func (s *ListBindAccountResponseBodyData) SetBindId(v int64) *ListBindAccountResponseBodyData {
	s.BindId = &v
	return s
}

func (s *ListBindAccountResponseBodyData) SetCloudCode(v string) *ListBindAccountResponseBodyData {
	s.CloudCode = &v
	return s
}

func (s *ListBindAccountResponseBodyData) SetCreateUser(v string) *ListBindAccountResponseBodyData {
	s.CreateUser = &v
	return s
}

func (s *ListBindAccountResponseBodyData) SetDataSourceCount(v int64) *ListBindAccountResponseBodyData {
	s.DataSourceCount = &v
	return s
}

func (s *ListBindAccountResponseBodyData) SetModifyTime(v string) *ListBindAccountResponseBodyData {
	s.ModifyTime = &v
	return s
}

type ListBindAccountResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBindAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBindAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s ListBindAccountResponse) GoString() string {
	return s.String()
}

func (s *ListBindAccountResponse) SetHeaders(v map[string]*string) *ListBindAccountResponse {
	s.Headers = v
	return s
}

func (s *ListBindAccountResponse) SetStatusCode(v int32) *ListBindAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBindAccountResponse) SetBody(v *ListBindAccountResponseBody) *ListBindAccountResponse {
	s.Body = v
	return s
}

type ListBindDataSourcesRequest struct {
	// The ID of the cloud account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123xxxxxxx
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The code of the cloud service provider.
	//
	// Valid values:
	//
	// 	- qcloud
	//
	// 	- hcloud
	//
	// 	- aliyun
	//
	// This parameter is required.
	//
	// example:
	//
	// hcloud
	CloudCode *string `json:"CloudCode,omitempty" xml:"CloudCode,omitempty"`
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

func (s ListBindDataSourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListBindDataSourcesRequest) GoString() string {
	return s.String()
}

func (s *ListBindDataSourcesRequest) SetAccountId(v string) *ListBindDataSourcesRequest {
	s.AccountId = &v
	return s
}

func (s *ListBindDataSourcesRequest) SetCloudCode(v string) *ListBindDataSourcesRequest {
	s.CloudCode = &v
	return s
}

func (s *ListBindDataSourcesRequest) SetRegionId(v string) *ListBindDataSourcesRequest {
	s.RegionId = &v
	return s
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
	return tea.Prettify(s)
}

func (s ListBindDataSourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListBindDataSourcesResponseBody) SetData(v []*ListBindDataSourcesResponseBodyData) *ListBindDataSourcesResponseBody {
	s.Data = v
	return s
}

func (s *ListBindDataSourcesResponseBody) SetRequestId(v string) *ListBindDataSourcesResponseBody {
	s.RequestId = &v
	return s
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
	return tea.Prettify(s)
}

func (s ListBindDataSourcesResponseBodyData) GoString() string {
	return s.String()
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

type ListBindDataSourcesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBindDataSourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBindDataSourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListBindDataSourcesResponse) GoString() string {
	return s.String()
}

func (s *ListBindDataSourcesResponse) SetHeaders(v map[string]*string) *ListBindDataSourcesResponse {
	s.Headers = v
	return s
}

func (s *ListBindDataSourcesResponse) SetStatusCode(v int32) *ListBindDataSourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBindDataSourcesResponse) SetBody(v *ListBindDataSourcesResponseBody) *ListBindDataSourcesResponse {
	s.Body = v
	return s
}

type ListCloudSiemCustomizeRulesRequest struct {
	// The alert type.
	//
	// example:
	//
	// scan
	AlertType *string `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
	// The page number. Pages start from page 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The end of the time range to query. Unit: milliseconds.
	//
	// example:
	//
	// 1577808000000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the custom rule.
	//
	// example:
	//
	// 10223
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The sort method. Valid values:
	//
	// 	- desc: descending order.
	//
	// 	- asc: ascending order.
	//
	// example:
	//
	// desc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The field that is used to sort the rules. Valid values:
	//
	// 	- GmtModified: The rules are sorted based on the modification time.
	//
	// 	- Id (default): The rules are sorted based on the rule ID.
	//
	// example:
	//
	// Id
	OrderField *string `json:"OrderField,omitempty" xml:"OrderField,omitempty"`
	// The number of entries per page. The value can be up to 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The data management center of the threat analysis feature. Specify this parameter based on the regions in which your assets reside. Valid values:
	//
	// 	- **cn-hangzhou**: Your assets reside in regions in China.
	//
	// 	- **ap-southeast-1**: Your assets reside in regions outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the destination account to which you switch the view from the management account.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view. Valid values:
	//
	// 	- 0: view of the current Alibaba Cloud account.
	//
	// 	- 1: view of all accounts for the enterprise.
	//
	// example:
	//
	// 0
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// The name of the rule. The name can contain letters, digits, underscores (_), and periods (.).
	//
	// example:
	//
	// waf_scan
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The type of the rule. Valid values:
	//
	// 	- **predefine**
	//
	// 	- **customize**
	//
	// example:
	//
	// customize
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// The beginning of the time range to query. Unit: milliseconds.
	//
	// example:
	//
	// 1577808000000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the rule. Valid values:
	//
	// 	- **0**: The rule is in the initial state.
	//
	// 	- **10**: The simulation data is tested.
	//
	// 	- **15**: The business data is being tested.
	//
	// 	- **20**: The business data test is complete.
	//
	// 	- **100**: The rule is in effect.
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The threat level. The value must be a JSON array. Valid values:
	//
	// 	- **serious**: high-risk.
	//
	// 	- **suspicious**: medium-risk.
	//
	// 	- **remind**: low-risk.
	//
	// example:
	//
	// ["serious","suspicious","remind"]
	ThreatLevel []*string `json:"ThreatLevel,omitempty" xml:"ThreatLevel,omitempty" type:"Repeated"`
}

func (s ListCloudSiemCustomizeRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCloudSiemCustomizeRulesRequest) GoString() string {
	return s.String()
}

func (s *ListCloudSiemCustomizeRulesRequest) SetAlertType(v string) *ListCloudSiemCustomizeRulesRequest {
	s.AlertType = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesRequest) SetCurrentPage(v int32) *ListCloudSiemCustomizeRulesRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesRequest) SetEndTime(v int64) *ListCloudSiemCustomizeRulesRequest {
	s.EndTime = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesRequest) SetId(v string) *ListCloudSiemCustomizeRulesRequest {
	s.Id = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesRequest) SetOrder(v string) *ListCloudSiemCustomizeRulesRequest {
	s.Order = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesRequest) SetOrderField(v string) *ListCloudSiemCustomizeRulesRequest {
	s.OrderField = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesRequest) SetPageSize(v int32) *ListCloudSiemCustomizeRulesRequest {
	s.PageSize = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesRequest) SetRegionId(v string) *ListCloudSiemCustomizeRulesRequest {
	s.RegionId = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesRequest) SetRoleFor(v int64) *ListCloudSiemCustomizeRulesRequest {
	s.RoleFor = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesRequest) SetRoleType(v int32) *ListCloudSiemCustomizeRulesRequest {
	s.RoleType = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesRequest) SetRuleName(v string) *ListCloudSiemCustomizeRulesRequest {
	s.RuleName = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesRequest) SetRuleType(v string) *ListCloudSiemCustomizeRulesRequest {
	s.RuleType = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesRequest) SetStartTime(v int64) *ListCloudSiemCustomizeRulesRequest {
	s.StartTime = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesRequest) SetStatus(v int32) *ListCloudSiemCustomizeRulesRequest {
	s.Status = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesRequest) SetThreatLevel(v []*string) *ListCloudSiemCustomizeRulesRequest {
	s.ThreatLevel = v
	return s
}

type ListCloudSiemCustomizeRulesResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	//
	// example:
	//
	// 123456
	Data *ListCloudSiemCustomizeRulesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListCloudSiemCustomizeRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCloudSiemCustomizeRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCloudSiemCustomizeRulesResponseBody) SetCode(v int32) *ListCloudSiemCustomizeRulesResponseBody {
	s.Code = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBody) SetData(v *ListCloudSiemCustomizeRulesResponseBodyData) *ListCloudSiemCustomizeRulesResponseBody {
	s.Data = v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBody) SetMessage(v string) *ListCloudSiemCustomizeRulesResponseBody {
	s.Message = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBody) SetRequestId(v string) *ListCloudSiemCustomizeRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBody) SetSuccess(v bool) *ListCloudSiemCustomizeRulesResponseBody {
	s.Success = &v
	return s
}

type ListCloudSiemCustomizeRulesResponseBodyData struct {
	// The pagination information.
	PageInfo *ListCloudSiemCustomizeRulesResponseBodyDataPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The detailed data.
	ResponseData []*ListCloudSiemCustomizeRulesResponseBodyDataResponseData `json:"ResponseData,omitempty" xml:"ResponseData,omitempty" type:"Repeated"`
}

func (s ListCloudSiemCustomizeRulesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListCloudSiemCustomizeRulesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCloudSiemCustomizeRulesResponseBodyData) SetPageInfo(v *ListCloudSiemCustomizeRulesResponseBodyDataPageInfo) *ListCloudSiemCustomizeRulesResponseBodyData {
	s.PageInfo = v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBodyData) SetResponseData(v []*ListCloudSiemCustomizeRulesResponseBodyDataResponseData) *ListCloudSiemCustomizeRulesResponseBodyData {
	s.ResponseData = v
	return s
}

type ListCloudSiemCustomizeRulesResponseBodyDataPageInfo struct {
	// The current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCloudSiemCustomizeRulesResponseBodyDataPageInfo) String() string {
	return tea.Prettify(s)
}

func (s ListCloudSiemCustomizeRulesResponseBodyDataPageInfo) GoString() string {
	return s.String()
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataPageInfo) SetCurrentPage(v int32) *ListCloudSiemCustomizeRulesResponseBodyDataPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataPageInfo) SetPageSize(v int32) *ListCloudSiemCustomizeRulesResponseBodyDataPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataPageInfo) SetTotalCount(v int64) *ListCloudSiemCustomizeRulesResponseBodyDataPageInfo {
	s.TotalCount = &v
	return s
}

type ListCloudSiemCustomizeRulesResponseBodyDataResponseData struct {
	// The type of the risk.
	//
	// example:
	//
	// WEBSHELL
	AlertType *string `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
	// The internal code of the risk type.
	//
	// example:
	//
	// ${siem_rule_type_process_abnormal_command}
	AlertTypeMds *string `json:"AlertTypeMds,omitempty" xml:"AlertTypeMds,omitempty"`
	// The ID of the Alibaba Cloud account in SIEM.
	//
	// example:
	//
	// 127608589417****
	Aliuid *int64 `json:"Aliuid,omitempty" xml:"Aliuid,omitempty"`
	// The alert additional field for ATT\\&CK.
	//
	// example:
	//
	// T1595.002 Vulnerability Scanning
	AttCk *string `json:"AttCk,omitempty" xml:"AttCk,omitempty"`
	// The type of the view. Valid values:
	//
	// 0: view of the current Alibaba Cloud account. 1: view of all accounts for the enterprise.
	//
	// example:
	//
	// 1
	DataType *int32 `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// The extended information about event generation. If the value of **eventTransferType*	- is **allToSingle**, the value of this parameter indicates the length and unit of the alert aggregation window. The HTML escape characters are reversed.
	//
	// example:
	//
	// {&quot;time&quot;:&quot;1&quot;,&quot;unit&quot;:&quot;MINUTE&quot;}
	EventTransferExt *string `json:"EventTransferExt,omitempty" xml:"EventTransferExt,omitempty"`
	// Indicates whether the system generates an event for the alert. Valid values:
	//
	// 	- **0**: no.
	//
	// 	- **1**: yes.
	//
	// example:
	//
	// 1
	EventTransferSwitch *int32 `json:"EventTransferSwitch,omitempty" xml:"EventTransferSwitch,omitempty"`
	// The method that is used to generate an event. Valid values:
	//
	// 	- **default**: built-in method.
	//
	// 	- **singleToSingle**: The system generates an event for each alert.
	//
	// 	- **allToSingle**: The system generates an event for alerts within a period of time.
	//
	// example:
	//
	// allToSingle
	EventTransferType *string `json:"EventTransferType,omitempty" xml:"EventTransferType,omitempty"`
	// The time when the custom rule was created.
	//
	// example:
	//
	// 2021-01-06 16:37:29
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the custom rule was last updated.
	//
	// example:
	//
	// 2021-01-06 16:37:29
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the custom rule.
	//
	// example:
	//
	// 123456789
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The log source of the rule.
	//
	// example:
	//
	// cloud_siem_aegis_sas_alert
	LogSource *string `json:"LogSource,omitempty" xml:"LogSource,omitempty"`
	// The internal code of the log source.
	//
	// example:
	//
	// ${sas.cloudsiem.prod.cloud_siem_aegis_sas_alert}
	LogSourceMds *string `json:"LogSourceMds,omitempty" xml:"LogSourceMds,omitempty"`
	// The log type of the rule.
	//
	// example:
	//
	// ALERT_ACTIVITY
	LogType *string `json:"LogType,omitempty" xml:"LogType,omitempty"`
	// The internal code of the log type.
	//
	// example:
	//
	// ${sas.cloudsiem.prod.alert_activity}
	LogTypeMds *string `json:"LogTypeMds,omitempty" xml:"LogTypeMds,omitempty"`
	// The window length of the rule. The HTML escape characters are reversed.
	//
	// example:
	//
	// {&quot;time&quot;:&quot;1&quot;,&quot;unit&quot;:&quot;HOUR&quot;}
	QueryCycle *string `json:"QueryCycle,omitempty" xml:"QueryCycle,omitempty"`
	// The query condition of the rule. The value is in the JSON format. The HTML escape characters are reversed.
	//
	// example:
	//
	// [[{&quot;not&quot;:false,&quot;left&quot;:&quot;alert_name&quot;,&quot;operator&quot;:&quot;=&quot;,&quot;right&quot;:&quot;WEBSHELL&quot;}]]
	RuleCondition *string `json:"RuleCondition,omitempty" xml:"RuleCondition,omitempty"`
	// The description of the rule.
	//
	// example:
	//
	// this rule is for waf scan
	RuleDesc *string `json:"RuleDesc,omitempty" xml:"RuleDesc,omitempty"`
	// The log aggregation field. The value is in the JSON format. The HTML escape characters are reversed.
	//
	// example:
	//
	// [&quot;asset_id&quot;]
	RuleGroup *string `json:"RuleGroup,omitempty" xml:"RuleGroup,omitempty"`
	// The name of the rule.
	//
	// example:
	//
	// waf_scan
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The threshold configurations of the rule in the JSON format. The HTML escape characters are reversed.
	//
	// example:
	//
	// {&quot;aggregateFunction&quot;:&quot;count&quot;,&quot;aggregateFunctionName&quot;:&quot;count&quot;,&quot;field&quot;:&quot;activity_name&quot;,&quot;operator&quot;:&quot;&lt;=&quot;,&quot;value&quot;:1}
	RuleThreshold *string `json:"RuleThreshold,omitempty" xml:"RuleThreshold,omitempty"`
	// The type of the rule. Valid values:
	//
	// 	- **predefine**
	//
	// 	- **customize**
	//
	// example:
	//
	// customize
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// The status of the rule. Valid values:
	//
	// 	- **0**: The rule is in the initial state.
	//
	// 	- **10**: The simulation data is tested.
	//
	// 	- **15**: The business data is being tested.
	//
	// 	- **20**: The business data test is complete.
	//
	// 	- **100**: The rule is in effect.
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The risk level. Valid values:
	//
	// 	- **serious**: high-risk.
	//
	// 	- **suspicious**: medium-risk.
	//
	// 	- **remind**: low-risk.
	//
	// example:
	//
	// remind
	ThreatLevel *string `json:"ThreatLevel,omitempty" xml:"ThreatLevel,omitempty"`
}

func (s ListCloudSiemCustomizeRulesResponseBodyDataResponseData) String() string {
	return tea.Prettify(s)
}

func (s ListCloudSiemCustomizeRulesResponseBodyDataResponseData) GoString() string {
	return s.String()
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) SetAlertType(v string) *ListCloudSiemCustomizeRulesResponseBodyDataResponseData {
	s.AlertType = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) SetAlertTypeMds(v string) *ListCloudSiemCustomizeRulesResponseBodyDataResponseData {
	s.AlertTypeMds = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) SetAliuid(v int64) *ListCloudSiemCustomizeRulesResponseBodyDataResponseData {
	s.Aliuid = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) SetAttCk(v string) *ListCloudSiemCustomizeRulesResponseBodyDataResponseData {
	s.AttCk = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) SetDataType(v int32) *ListCloudSiemCustomizeRulesResponseBodyDataResponseData {
	s.DataType = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) SetEventTransferExt(v string) *ListCloudSiemCustomizeRulesResponseBodyDataResponseData {
	s.EventTransferExt = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) SetEventTransferSwitch(v int32) *ListCloudSiemCustomizeRulesResponseBodyDataResponseData {
	s.EventTransferSwitch = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) SetEventTransferType(v string) *ListCloudSiemCustomizeRulesResponseBodyDataResponseData {
	s.EventTransferType = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) SetGmtCreate(v string) *ListCloudSiemCustomizeRulesResponseBodyDataResponseData {
	s.GmtCreate = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) SetGmtModified(v string) *ListCloudSiemCustomizeRulesResponseBodyDataResponseData {
	s.GmtModified = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) SetId(v int64) *ListCloudSiemCustomizeRulesResponseBodyDataResponseData {
	s.Id = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) SetLogSource(v string) *ListCloudSiemCustomizeRulesResponseBodyDataResponseData {
	s.LogSource = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) SetLogSourceMds(v string) *ListCloudSiemCustomizeRulesResponseBodyDataResponseData {
	s.LogSourceMds = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) SetLogType(v string) *ListCloudSiemCustomizeRulesResponseBodyDataResponseData {
	s.LogType = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) SetLogTypeMds(v string) *ListCloudSiemCustomizeRulesResponseBodyDataResponseData {
	s.LogTypeMds = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) SetQueryCycle(v string) *ListCloudSiemCustomizeRulesResponseBodyDataResponseData {
	s.QueryCycle = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) SetRuleCondition(v string) *ListCloudSiemCustomizeRulesResponseBodyDataResponseData {
	s.RuleCondition = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) SetRuleDesc(v string) *ListCloudSiemCustomizeRulesResponseBodyDataResponseData {
	s.RuleDesc = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) SetRuleGroup(v string) *ListCloudSiemCustomizeRulesResponseBodyDataResponseData {
	s.RuleGroup = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) SetRuleName(v string) *ListCloudSiemCustomizeRulesResponseBodyDataResponseData {
	s.RuleName = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) SetRuleThreshold(v string) *ListCloudSiemCustomizeRulesResponseBodyDataResponseData {
	s.RuleThreshold = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) SetRuleType(v string) *ListCloudSiemCustomizeRulesResponseBodyDataResponseData {
	s.RuleType = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) SetStatus(v int32) *ListCloudSiemCustomizeRulesResponseBodyDataResponseData {
	s.Status = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) SetThreatLevel(v string) *ListCloudSiemCustomizeRulesResponseBodyDataResponseData {
	s.ThreatLevel = &v
	return s
}

type ListCloudSiemCustomizeRulesResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCloudSiemCustomizeRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCloudSiemCustomizeRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCloudSiemCustomizeRulesResponse) GoString() string {
	return s.String()
}

func (s *ListCloudSiemCustomizeRulesResponse) SetHeaders(v map[string]*string) *ListCloudSiemCustomizeRulesResponse {
	s.Headers = v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponse) SetStatusCode(v int32) *ListCloudSiemCustomizeRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponse) SetBody(v *ListCloudSiemCustomizeRulesResponseBody) *ListCloudSiemCustomizeRulesResponse {
	s.Body = v
	return s
}

type ListCloudSiemPredefinedRulesRequest struct {
	// The alert type.
	//
	// example:
	//
	// scan
	AlertType *string `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
	// The ATT\\&CK information.
	//
	// example:
	//
	// T1595.002 Vulnerability Scanning
	AttCk *string `json:"AttCk,omitempty" xml:"AttCk,omitempty"`
	// The page number. Pages start from page 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The end of the time range to query. Unit: milliseconds.
	//
	// example:
	//
	// 1577808000000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The method that is used to generate an event. Valid values:
	//
	// 	- default: built-in method.
	//
	// 	- singleToSingle: The system generates an event for each alert.
	//
	// 	- allToSingle: The system generates an event for alerts within a period of time.
	//
	// example:
	//
	// allToSingle
	EventTransferType *string `json:"EventTransferType,omitempty" xml:"EventTransferType,omitempty"`
	// The ID of the rule.
	//
	// example:
	//
	// 10223
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The log source.
	//
	// example:
	//
	// cloud_siem_aegis_sas_alert
	LogSource *string `json:"LogSource,omitempty" xml:"LogSource,omitempty"`
	// The sort method. Valid values:
	//
	// 	- desc: descending order.
	//
	// 	- asc: ascending order.
	//
	// example:
	//
	// desc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The field that is used to sort the rules. Valid values:
	//
	// 	- GmtModified: The rules are sorted based on the modification time.
	//
	// 	- Id (default): The rules are sorted based on the rule ID.
	//
	// example:
	//
	// Id
	OrderField *string `json:"OrderField,omitempty" xml:"OrderField,omitempty"`
	// The number of entries per page. Maximum value: 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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
	// The ID of the destination account to which you switch the view from the management account.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view.
	//
	// 	- 0: view of the current Alibaba Cloud account.
	//
	// 	- 1: view of all accounts for the enterprise.
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// The name of the rule. The name can contain letters, digits, underscores (_), and periods (.).
	//
	// example:
	//
	// waf_scan
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The type of the rule. Valid values:
	//
	// 	- predefine
	//
	// 	- customize
	//
	// example:
	//
	// customize
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// The beginning of the time range to query. Unit: milliseconds.
	//
	// example:
	//
	// 1577808000000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the rule. Valid values:
	//
	// 	- 0: The rule is in the initial state.
	//
	// 	- 10: The simulation data is tested.
	//
	// 	- 15: The business data is being tested.
	//
	// 	- 20: The business data test ends.
	//
	// 	- 100: The rule takes effect.
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The risk level. The value is a JSON array. Valid values:
	//
	// 	- serious: high
	//
	// 	- suspicious: medium
	//
	// 	- remind: low
	//
	// example:
	//
	// ["serious","suspicious","remind"]
	ThreatLevel []*string `json:"ThreatLevel,omitempty" xml:"ThreatLevel,omitempty" type:"Repeated"`
}

func (s ListCloudSiemPredefinedRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCloudSiemPredefinedRulesRequest) GoString() string {
	return s.String()
}

func (s *ListCloudSiemPredefinedRulesRequest) SetAlertType(v string) *ListCloudSiemPredefinedRulesRequest {
	s.AlertType = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesRequest) SetAttCk(v string) *ListCloudSiemPredefinedRulesRequest {
	s.AttCk = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesRequest) SetCurrentPage(v int32) *ListCloudSiemPredefinedRulesRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesRequest) SetEndTime(v int64) *ListCloudSiemPredefinedRulesRequest {
	s.EndTime = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesRequest) SetEventTransferType(v string) *ListCloudSiemPredefinedRulesRequest {
	s.EventTransferType = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesRequest) SetId(v string) *ListCloudSiemPredefinedRulesRequest {
	s.Id = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesRequest) SetLogSource(v string) *ListCloudSiemPredefinedRulesRequest {
	s.LogSource = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesRequest) SetOrder(v string) *ListCloudSiemPredefinedRulesRequest {
	s.Order = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesRequest) SetOrderField(v string) *ListCloudSiemPredefinedRulesRequest {
	s.OrderField = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesRequest) SetPageSize(v int32) *ListCloudSiemPredefinedRulesRequest {
	s.PageSize = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesRequest) SetRegionId(v string) *ListCloudSiemPredefinedRulesRequest {
	s.RegionId = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesRequest) SetRoleFor(v int64) *ListCloudSiemPredefinedRulesRequest {
	s.RoleFor = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesRequest) SetRoleType(v int32) *ListCloudSiemPredefinedRulesRequest {
	s.RoleType = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesRequest) SetRuleName(v string) *ListCloudSiemPredefinedRulesRequest {
	s.RuleName = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesRequest) SetRuleType(v string) *ListCloudSiemPredefinedRulesRequest {
	s.RuleType = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesRequest) SetStartTime(v int64) *ListCloudSiemPredefinedRulesRequest {
	s.StartTime = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesRequest) SetStatus(v int32) *ListCloudSiemPredefinedRulesRequest {
	s.Status = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesRequest) SetThreatLevel(v []*string) *ListCloudSiemPredefinedRulesRequest {
	s.ThreatLevel = v
	return s
}

type ListCloudSiemPredefinedRulesResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	//
	// example:
	//
	// 123456
	Data *ListCloudSiemPredefinedRulesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListCloudSiemPredefinedRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCloudSiemPredefinedRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCloudSiemPredefinedRulesResponseBody) SetCode(v int32) *ListCloudSiemPredefinedRulesResponseBody {
	s.Code = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesResponseBody) SetData(v *ListCloudSiemPredefinedRulesResponseBodyData) *ListCloudSiemPredefinedRulesResponseBody {
	s.Data = v
	return s
}

func (s *ListCloudSiemPredefinedRulesResponseBody) SetMessage(v string) *ListCloudSiemPredefinedRulesResponseBody {
	s.Message = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesResponseBody) SetRequestId(v string) *ListCloudSiemPredefinedRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesResponseBody) SetSuccess(v bool) *ListCloudSiemPredefinedRulesResponseBody {
	s.Success = &v
	return s
}

type ListCloudSiemPredefinedRulesResponseBodyData struct {
	// The pagination information.
	PageInfo *ListCloudSiemPredefinedRulesResponseBodyDataPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The detailed data.
	ResponseData []*ListCloudSiemPredefinedRulesResponseBodyDataResponseData `json:"ResponseData,omitempty" xml:"ResponseData,omitempty" type:"Repeated"`
}

func (s ListCloudSiemPredefinedRulesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListCloudSiemPredefinedRulesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCloudSiemPredefinedRulesResponseBodyData) SetPageInfo(v *ListCloudSiemPredefinedRulesResponseBodyDataPageInfo) *ListCloudSiemPredefinedRulesResponseBodyData {
	s.PageInfo = v
	return s
}

func (s *ListCloudSiemPredefinedRulesResponseBodyData) SetResponseData(v []*ListCloudSiemPredefinedRulesResponseBodyDataResponseData) *ListCloudSiemPredefinedRulesResponseBodyData {
	s.ResponseData = v
	return s
}

type ListCloudSiemPredefinedRulesResponseBodyDataPageInfo struct {
	// The current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCloudSiemPredefinedRulesResponseBodyDataPageInfo) String() string {
	return tea.Prettify(s)
}

func (s ListCloudSiemPredefinedRulesResponseBodyDataPageInfo) GoString() string {
	return s.String()
}

func (s *ListCloudSiemPredefinedRulesResponseBodyDataPageInfo) SetCurrentPage(v int32) *ListCloudSiemPredefinedRulesResponseBodyDataPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesResponseBodyDataPageInfo) SetPageSize(v int32) *ListCloudSiemPredefinedRulesResponseBodyDataPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesResponseBodyDataPageInfo) SetTotalCount(v int64) *ListCloudSiemPredefinedRulesResponseBodyDataPageInfo {
	s.TotalCount = &v
	return s
}

type ListCloudSiemPredefinedRulesResponseBodyDataResponseData struct {
	// The type of the risk.
	//
	// example:
	//
	// WEBSHELL
	AlertType *string `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
	// The alert additional field for ATT\\&CK.
	//
	// example:
	//
	// T1595.002 Vulnerability Scanning
	AttCk *string `json:"AttCk,omitempty" xml:"AttCk,omitempty"`
	// The method that is used to generate an event. Valid values:
	//
	// 	- default: built-in method.
	//
	// 	- singleToSingle: The system generates an event for each alert.
	//
	// 	- allToSingle: The system generates an event for alerts within a period of time.
	//
	// example:
	//
	// allToSingle
	EventTransferType *string `json:"EventTransferType,omitempty" xml:"EventTransferType,omitempty"`
	// The time when the rule was created.
	//
	// example:
	//
	// 2021-01-06 16:37:29
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the rule was modified.
	//
	// example:
	//
	// 2021-01-06 16:37:29
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the predefined rule.
	//
	// example:
	//
	// 123456789
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The internal code of the rule description.
	//
	// example:
	//
	// ${siem_rule_description_siem_cfw-attack-count-level-up_cfw-attack}
	RuleDescMds *string `json:"RuleDescMds,omitempty" xml:"RuleDescMds,omitempty"`
	// The name of the rule.
	//
	// example:
	//
	// siem_base64-command-exec_aegis-proc
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The rule name in Chinese.
	//
	// example:
	//
	// siem_base64-command-exec_aegis-proc
	RuleNameCn *string `json:"RuleNameCn,omitempty" xml:"RuleNameCn,omitempty"`
	// The rule name in English.
	//
	// example:
	//
	// siem_base64-command-exec_aegis-proc
	RuleNameEn *string `json:"RuleNameEn,omitempty" xml:"RuleNameEn,omitempty"`
	// The internal code of the rule name.
	//
	// example:
	//
	// ${siem_rule_name_siem_cfw-attack-count-level-up_cfw-attack}
	RuleNameMds *string `json:"RuleNameMds,omitempty" xml:"RuleNameMds,omitempty"`
	// The log source of the rule.
	//
	// example:
	//
	// cloud_siem_aegis_proc
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The status of the predefined rule. Valid values:
	//
	// 	- 0: The rule is in the initial state.
	//
	// 	- 100: The rule takes effect.
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The risk level. Valid values:
	//
	// 	- serious: high.
	//
	// 	- suspicious: medium.
	//
	// 	- remind: low.
	//
	// example:
	//
	// remind
	ThreatLevel *string `json:"ThreatLevel,omitempty" xml:"ThreatLevel,omitempty"`
}

func (s ListCloudSiemPredefinedRulesResponseBodyDataResponseData) String() string {
	return tea.Prettify(s)
}

func (s ListCloudSiemPredefinedRulesResponseBodyDataResponseData) GoString() string {
	return s.String()
}

func (s *ListCloudSiemPredefinedRulesResponseBodyDataResponseData) SetAlertType(v string) *ListCloudSiemPredefinedRulesResponseBodyDataResponseData {
	s.AlertType = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesResponseBodyDataResponseData) SetAttCk(v string) *ListCloudSiemPredefinedRulesResponseBodyDataResponseData {
	s.AttCk = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesResponseBodyDataResponseData) SetEventTransferType(v string) *ListCloudSiemPredefinedRulesResponseBodyDataResponseData {
	s.EventTransferType = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesResponseBodyDataResponseData) SetGmtCreate(v string) *ListCloudSiemPredefinedRulesResponseBodyDataResponseData {
	s.GmtCreate = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesResponseBodyDataResponseData) SetGmtModified(v string) *ListCloudSiemPredefinedRulesResponseBodyDataResponseData {
	s.GmtModified = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesResponseBodyDataResponseData) SetId(v int64) *ListCloudSiemPredefinedRulesResponseBodyDataResponseData {
	s.Id = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesResponseBodyDataResponseData) SetRuleDescMds(v string) *ListCloudSiemPredefinedRulesResponseBodyDataResponseData {
	s.RuleDescMds = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesResponseBodyDataResponseData) SetRuleName(v string) *ListCloudSiemPredefinedRulesResponseBodyDataResponseData {
	s.RuleName = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesResponseBodyDataResponseData) SetRuleNameCn(v string) *ListCloudSiemPredefinedRulesResponseBodyDataResponseData {
	s.RuleNameCn = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesResponseBodyDataResponseData) SetRuleNameEn(v string) *ListCloudSiemPredefinedRulesResponseBodyDataResponseData {
	s.RuleNameEn = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesResponseBodyDataResponseData) SetRuleNameMds(v string) *ListCloudSiemPredefinedRulesResponseBodyDataResponseData {
	s.RuleNameMds = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesResponseBodyDataResponseData) SetSource(v string) *ListCloudSiemPredefinedRulesResponseBodyDataResponseData {
	s.Source = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesResponseBodyDataResponseData) SetStatus(v int32) *ListCloudSiemPredefinedRulesResponseBodyDataResponseData {
	s.Status = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesResponseBodyDataResponseData) SetThreatLevel(v string) *ListCloudSiemPredefinedRulesResponseBodyDataResponseData {
	s.ThreatLevel = &v
	return s
}

type ListCloudSiemPredefinedRulesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCloudSiemPredefinedRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCloudSiemPredefinedRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCloudSiemPredefinedRulesResponse) GoString() string {
	return s.String()
}

func (s *ListCloudSiemPredefinedRulesResponse) SetHeaders(v map[string]*string) *ListCloudSiemPredefinedRulesResponse {
	s.Headers = v
	return s
}

func (s *ListCloudSiemPredefinedRulesResponse) SetStatusCode(v int32) *ListCloudSiemPredefinedRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesResponse) SetBody(v *ListCloudSiemPredefinedRulesResponseBody) *ListCloudSiemPredefinedRulesResponse {
	s.Body = v
	return s
}

type ListCustomizeRuleTestResultRequest struct {
	// The page number. Pages start from page 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage     *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	DetectionRuleId *string `json:"DetectionRuleId,omitempty" xml:"DetectionRuleId,omitempty"`
	EndTime         *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the rule.
	//
	// example:
	//
	// 123456789
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The number of entries per page. Valid values: 1 to 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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
	// The ID of the account that you switch from the management account.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view. Valid values:
	//
	// - 0: the current Alibaba Cloud account
	//
	// - 1: the global account
	//
	// example:
	//
	// 1
	RoleType   *int32  `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	StartTime  *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	VerifyType *string `json:"VerifyType,omitempty" xml:"VerifyType,omitempty"`
}

func (s ListCustomizeRuleTestResultRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCustomizeRuleTestResultRequest) GoString() string {
	return s.String()
}

func (s *ListCustomizeRuleTestResultRequest) SetCurrentPage(v int32) *ListCustomizeRuleTestResultRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListCustomizeRuleTestResultRequest) SetDetectionRuleId(v string) *ListCustomizeRuleTestResultRequest {
	s.DetectionRuleId = &v
	return s
}

func (s *ListCustomizeRuleTestResultRequest) SetEndTime(v int64) *ListCustomizeRuleTestResultRequest {
	s.EndTime = &v
	return s
}

func (s *ListCustomizeRuleTestResultRequest) SetId(v int64) *ListCustomizeRuleTestResultRequest {
	s.Id = &v
	return s
}

func (s *ListCustomizeRuleTestResultRequest) SetPageSize(v int32) *ListCustomizeRuleTestResultRequest {
	s.PageSize = &v
	return s
}

func (s *ListCustomizeRuleTestResultRequest) SetRegionId(v string) *ListCustomizeRuleTestResultRequest {
	s.RegionId = &v
	return s
}

func (s *ListCustomizeRuleTestResultRequest) SetRoleFor(v int64) *ListCustomizeRuleTestResultRequest {
	s.RoleFor = &v
	return s
}

func (s *ListCustomizeRuleTestResultRequest) SetRoleType(v int32) *ListCustomizeRuleTestResultRequest {
	s.RoleType = &v
	return s
}

func (s *ListCustomizeRuleTestResultRequest) SetStartTime(v int64) *ListCustomizeRuleTestResultRequest {
	s.StartTime = &v
	return s
}

func (s *ListCustomizeRuleTestResultRequest) SetVerifyType(v string) *ListCustomizeRuleTestResultRequest {
	s.VerifyType = &v
	return s
}

type ListCustomizeRuleTestResultResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	//
	// example:
	//
	// 123456
	Data *ListCustomizeRuleTestResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListCustomizeRuleTestResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCustomizeRuleTestResultResponseBody) GoString() string {
	return s.String()
}

func (s *ListCustomizeRuleTestResultResponseBody) SetCode(v int32) *ListCustomizeRuleTestResultResponseBody {
	s.Code = &v
	return s
}

func (s *ListCustomizeRuleTestResultResponseBody) SetData(v *ListCustomizeRuleTestResultResponseBodyData) *ListCustomizeRuleTestResultResponseBody {
	s.Data = v
	return s
}

func (s *ListCustomizeRuleTestResultResponseBody) SetMessage(v string) *ListCustomizeRuleTestResultResponseBody {
	s.Message = &v
	return s
}

func (s *ListCustomizeRuleTestResultResponseBody) SetRequestId(v string) *ListCustomizeRuleTestResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCustomizeRuleTestResultResponseBody) SetSuccess(v bool) *ListCustomizeRuleTestResultResponseBody {
	s.Success = &v
	return s
}

type ListCustomizeRuleTestResultResponseBodyData struct {
	// The pagination information.
	PageInfo *ListCustomizeRuleTestResultResponseBodyDataPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The detailed data.
	ResponseData []*ListCustomizeRuleTestResultResponseBodyDataResponseData `json:"ResponseData,omitempty" xml:"ResponseData,omitempty" type:"Repeated"`
}

func (s ListCustomizeRuleTestResultResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListCustomizeRuleTestResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCustomizeRuleTestResultResponseBodyData) SetPageInfo(v *ListCustomizeRuleTestResultResponseBodyDataPageInfo) *ListCustomizeRuleTestResultResponseBodyData {
	s.PageInfo = v
	return s
}

func (s *ListCustomizeRuleTestResultResponseBodyData) SetResponseData(v []*ListCustomizeRuleTestResultResponseBodyDataResponseData) *ListCustomizeRuleTestResultResponseBodyData {
	s.ResponseData = v
	return s
}

type ListCustomizeRuleTestResultResponseBodyDataPageInfo struct {
	// The current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount    *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	VerifiedCount *int64 `json:"VerifiedCount,omitempty" xml:"VerifiedCount,omitempty"`
}

func (s ListCustomizeRuleTestResultResponseBodyDataPageInfo) String() string {
	return tea.Prettify(s)
}

func (s ListCustomizeRuleTestResultResponseBodyDataPageInfo) GoString() string {
	return s.String()
}

func (s *ListCustomizeRuleTestResultResponseBodyDataPageInfo) SetCurrentPage(v int32) *ListCustomizeRuleTestResultResponseBodyDataPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *ListCustomizeRuleTestResultResponseBodyDataPageInfo) SetPageSize(v int32) *ListCustomizeRuleTestResultResponseBodyDataPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListCustomizeRuleTestResultResponseBodyDataPageInfo) SetTotalCount(v int64) *ListCustomizeRuleTestResultResponseBodyDataPageInfo {
	s.TotalCount = &v
	return s
}

func (s *ListCustomizeRuleTestResultResponseBodyDataPageInfo) SetVerifiedCount(v int64) *ListCustomizeRuleTestResultResponseBodyDataPageInfo {
	s.VerifiedCount = &v
	return s
}

type ListCustomizeRuleTestResultResponseBodyDataResponseData struct {
	// The description of the alert.
	//
	// example:
	//
	// The account you logged in this time is not in the legal account category defined by you. Please confirm the legality of the login behavior.
	AlertDesc *string `json:"AlertDesc,omitempty" xml:"AlertDesc,omitempty"`
	// The alert details in the JSON format.
	//
	// example:
	//
	// {"main_user_id": "165295629792****";"log_uuid_count": "99";"attack_ip": "218.92.XX.XX"}
	AlertDetail *string `json:"AlertDetail,omitempty" xml:"AlertDetail,omitempty"`
	// The source of the alert.
	//
	// example:
	//
	// sas
	AlertSrcProd *string `json:"AlertSrcProd,omitempty" xml:"AlertSrcProd,omitempty"`
	// The sub-module of the source.
	//
	// example:
	//
	// waf
	AlertSrcProdModule *string `json:"AlertSrcProdModule,omitempty" xml:"AlertSrcProdModule,omitempty"`
	// The tag of the ATT\\&CK attack.
	//
	// example:
	//
	// T1595.002 Vulnerability Scanning
	AttCk *string `json:"AttCk,omitempty" xml:"AttCk,omitempty"`
	// The name of the alert, which corresponds to the name of the custom rule.
	//
	// example:
	//
	// waf_scan
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// The threat type, which indicates the alert type.
	//
	// example:
	//
	// WEBSHELL
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// The threat level. Valid values:
	//
	// 	- serious: high.
	//
	// 	- suspicious: medium.
	//
	// 	- remind: low.
	//
	// example:
	//
	// remind
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The log source of the rule.
	//
	// example:
	//
	// cloud_siem_aegis_sas_alert
	LogSource *string `json:"LogSource,omitempty" xml:"LogSource,omitempty"`
	// The time when the alert was recorded.
	//
	// example:
	//
	// 2021-01-06 16:37:29
	LogTime *string `json:"LogTime,omitempty" xml:"LogTime,omitempty"`
	// The log type of the rule.
	//
	// example:
	//
	// ALERT_ACTIVITY
	LogType *string `json:"LogType,omitempty" xml:"LogType,omitempty"`
	// The ID of the Alibaba Cloud account that is associated with the alert in SIEM.
	//
	// example:
	//
	// 127608589417****
	MainUserId *string `json:"MainUserId,omitempty" xml:"MainUserId,omitempty"`
	// The status of the alert data. Valid values:
	//
	// 	- test: business test data.
	//
	// 	- online: online data.
	//
	// example:
	//
	// test
	OnlineStatus *string `json:"OnlineStatus,omitempty" xml:"OnlineStatus,omitempty"`
	// The ID of the Alibaba Cloud account within which the alert is generated.
	//
	// example:
	//
	// 176555323***
	SubUserId *string `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
	// The UUID of the alert.
	//
	// example:
	//
	// sas_71e24437d2797ce8fc59692905a4****
	Uuid       *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	VerifyType *string `json:"VerifyType,omitempty" xml:"VerifyType,omitempty"`
}

func (s ListCustomizeRuleTestResultResponseBodyDataResponseData) String() string {
	return tea.Prettify(s)
}

func (s ListCustomizeRuleTestResultResponseBodyDataResponseData) GoString() string {
	return s.String()
}

func (s *ListCustomizeRuleTestResultResponseBodyDataResponseData) SetAlertDesc(v string) *ListCustomizeRuleTestResultResponseBodyDataResponseData {
	s.AlertDesc = &v
	return s
}

func (s *ListCustomizeRuleTestResultResponseBodyDataResponseData) SetAlertDetail(v string) *ListCustomizeRuleTestResultResponseBodyDataResponseData {
	s.AlertDetail = &v
	return s
}

func (s *ListCustomizeRuleTestResultResponseBodyDataResponseData) SetAlertSrcProd(v string) *ListCustomizeRuleTestResultResponseBodyDataResponseData {
	s.AlertSrcProd = &v
	return s
}

func (s *ListCustomizeRuleTestResultResponseBodyDataResponseData) SetAlertSrcProdModule(v string) *ListCustomizeRuleTestResultResponseBodyDataResponseData {
	s.AlertSrcProdModule = &v
	return s
}

func (s *ListCustomizeRuleTestResultResponseBodyDataResponseData) SetAttCk(v string) *ListCustomizeRuleTestResultResponseBodyDataResponseData {
	s.AttCk = &v
	return s
}

func (s *ListCustomizeRuleTestResultResponseBodyDataResponseData) SetEventName(v string) *ListCustomizeRuleTestResultResponseBodyDataResponseData {
	s.EventName = &v
	return s
}

func (s *ListCustomizeRuleTestResultResponseBodyDataResponseData) SetEventType(v string) *ListCustomizeRuleTestResultResponseBodyDataResponseData {
	s.EventType = &v
	return s
}

func (s *ListCustomizeRuleTestResultResponseBodyDataResponseData) SetLevel(v string) *ListCustomizeRuleTestResultResponseBodyDataResponseData {
	s.Level = &v
	return s
}

func (s *ListCustomizeRuleTestResultResponseBodyDataResponseData) SetLogSource(v string) *ListCustomizeRuleTestResultResponseBodyDataResponseData {
	s.LogSource = &v
	return s
}

func (s *ListCustomizeRuleTestResultResponseBodyDataResponseData) SetLogTime(v string) *ListCustomizeRuleTestResultResponseBodyDataResponseData {
	s.LogTime = &v
	return s
}

func (s *ListCustomizeRuleTestResultResponseBodyDataResponseData) SetLogType(v string) *ListCustomizeRuleTestResultResponseBodyDataResponseData {
	s.LogType = &v
	return s
}

func (s *ListCustomizeRuleTestResultResponseBodyDataResponseData) SetMainUserId(v string) *ListCustomizeRuleTestResultResponseBodyDataResponseData {
	s.MainUserId = &v
	return s
}

func (s *ListCustomizeRuleTestResultResponseBodyDataResponseData) SetOnlineStatus(v string) *ListCustomizeRuleTestResultResponseBodyDataResponseData {
	s.OnlineStatus = &v
	return s
}

func (s *ListCustomizeRuleTestResultResponseBodyDataResponseData) SetSubUserId(v string) *ListCustomizeRuleTestResultResponseBodyDataResponseData {
	s.SubUserId = &v
	return s
}

func (s *ListCustomizeRuleTestResultResponseBodyDataResponseData) SetUuid(v string) *ListCustomizeRuleTestResultResponseBodyDataResponseData {
	s.Uuid = &v
	return s
}

func (s *ListCustomizeRuleTestResultResponseBodyDataResponseData) SetVerifyType(v string) *ListCustomizeRuleTestResultResponseBodyDataResponseData {
	s.VerifyType = &v
	return s
}

type ListCustomizeRuleTestResultResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCustomizeRuleTestResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCustomizeRuleTestResultResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCustomizeRuleTestResultResponse) GoString() string {
	return s.String()
}

func (s *ListCustomizeRuleTestResultResponse) SetHeaders(v map[string]*string) *ListCustomizeRuleTestResultResponse {
	s.Headers = v
	return s
}

func (s *ListCustomizeRuleTestResultResponse) SetStatusCode(v int32) *ListCustomizeRuleTestResultResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCustomizeRuleTestResultResponse) SetBody(v *ListCustomizeRuleTestResultResponseBody) *ListCustomizeRuleTestResultResponse {
	s.Body = v
	return s
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
	return tea.Prettify(s)
}

func (s ListDataSourceLogsRequest) GoString() string {
	return s.String()
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

type ListDataSourceLogsResponseBody struct {
	// The data returned.
	Data *ListDataSourceLogsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDataSourceLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourceLogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataSourceLogsResponseBody) SetData(v *ListDataSourceLogsResponseBodyData) *ListDataSourceLogsResponseBody {
	s.Data = v
	return s
}

func (s *ListDataSourceLogsResponseBody) SetRequestId(v string) *ListDataSourceLogsResponseBody {
	s.RequestId = &v
	return s
}

type ListDataSourceLogsResponseBodyData struct {
	// The ID of the cloud account.
	//
	// example:
	//
	// 123xxxxxxx
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The code that is used for multi-cloud environments. Valid values:
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
	// The ID of the data source. The value is obtained after the threat analysis feature calculates the MD5 hash value of a parameter.
	//
	// example:
	//
	// 220ba97c9d1fdb0b9c7e8c7ca328d7ea
	DataSourceInstanceId *string `json:"DataSourceInstanceId,omitempty" xml:"DataSourceInstanceId,omitempty"`
	// The logs of the data source.
	DataSourceInstanceLogs []*ListDataSourceLogsResponseBodyDataDataSourceInstanceLogs `json:"DataSourceInstanceLogs,omitempty" xml:"DataSourceInstanceLogs,omitempty" type:"Repeated"`
	// The name of the data source.
	//
	// example:
	//
	// waf kafka
	DataSourceInstanceName *string `json:"DataSourceInstanceName,omitempty" xml:"DataSourceInstanceName,omitempty"`
	// The remarks of the data source.
	//
	// example:
	//
	// waf kafka
	DataSourceInstanceRemark *string `json:"DataSourceInstanceRemark,omitempty" xml:"DataSourceInstanceRemark,omitempty"`
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 123XXXXXXXX
	SubUserId *int64 `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
}

func (s ListDataSourceLogsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourceLogsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDataSourceLogsResponseBodyData) SetAccountId(v string) *ListDataSourceLogsResponseBodyData {
	s.AccountId = &v
	return s
}

func (s *ListDataSourceLogsResponseBodyData) SetCloudCode(v string) *ListDataSourceLogsResponseBodyData {
	s.CloudCode = &v
	return s
}

func (s *ListDataSourceLogsResponseBodyData) SetDataSourceInstanceId(v string) *ListDataSourceLogsResponseBodyData {
	s.DataSourceInstanceId = &v
	return s
}

func (s *ListDataSourceLogsResponseBodyData) SetDataSourceInstanceLogs(v []*ListDataSourceLogsResponseBodyDataDataSourceInstanceLogs) *ListDataSourceLogsResponseBodyData {
	s.DataSourceInstanceLogs = v
	return s
}

func (s *ListDataSourceLogsResponseBodyData) SetDataSourceInstanceName(v string) *ListDataSourceLogsResponseBodyData {
	s.DataSourceInstanceName = &v
	return s
}

func (s *ListDataSourceLogsResponseBodyData) SetDataSourceInstanceRemark(v string) *ListDataSourceLogsResponseBodyData {
	s.DataSourceInstanceRemark = &v
	return s
}

func (s *ListDataSourceLogsResponseBodyData) SetSubUserId(v int64) *ListDataSourceLogsResponseBodyData {
	s.SubUserId = &v
	return s
}

type ListDataSourceLogsResponseBodyDataDataSourceInstanceLogs struct {
	// The code of the log.
	//
	// example:
	//
	// cloud_siem_waf_xxxxx
	LogCode *string `json:"LogCode,omitempty" xml:"LogCode,omitempty"`
	// The ID of the log. The value is obtained after the threat analysis feature calculates the MD5 hash value of a parameter.
	//
	// example:
	//
	// 220ba97c9d1fdb0b9c7e8c7ca328d7ea
	LogInstanceId *string `json:"LogInstanceId,omitempty" xml:"LogInstanceId,omitempty"`
	// The display code of the log.
	//
	// example:
	//
	// ${siem.prod.cloud_siem_waf_xxxxx}
	LogMdsCode *string `json:"LogMdsCode,omitempty" xml:"LogMdsCode,omitempty"`
	// The parameters of the log.
	LogParams []*ListDataSourceLogsResponseBodyDataDataSourceInstanceLogsLogParams `json:"LogParams,omitempty" xml:"LogParams,omitempty" type:"Repeated"`
	// Indicates whether the task for which logs are collected is enabled. Valid values:
	//
	// 	- 1: yes
	//
	// 	- 0: no
	//
	// example:
	//
	// 1
	TaskStatus *int32 `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
}

func (s ListDataSourceLogsResponseBodyDataDataSourceInstanceLogs) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourceLogsResponseBodyDataDataSourceInstanceLogs) GoString() string {
	return s.String()
}

func (s *ListDataSourceLogsResponseBodyDataDataSourceInstanceLogs) SetLogCode(v string) *ListDataSourceLogsResponseBodyDataDataSourceInstanceLogs {
	s.LogCode = &v
	return s
}

func (s *ListDataSourceLogsResponseBodyDataDataSourceInstanceLogs) SetLogInstanceId(v string) *ListDataSourceLogsResponseBodyDataDataSourceInstanceLogs {
	s.LogInstanceId = &v
	return s
}

func (s *ListDataSourceLogsResponseBodyDataDataSourceInstanceLogs) SetLogMdsCode(v string) *ListDataSourceLogsResponseBodyDataDataSourceInstanceLogs {
	s.LogMdsCode = &v
	return s
}

func (s *ListDataSourceLogsResponseBodyDataDataSourceInstanceLogs) SetLogParams(v []*ListDataSourceLogsResponseBodyDataDataSourceInstanceLogsLogParams) *ListDataSourceLogsResponseBodyDataDataSourceInstanceLogs {
	s.LogParams = v
	return s
}

func (s *ListDataSourceLogsResponseBodyDataDataSourceInstanceLogs) SetTaskStatus(v int32) *ListDataSourceLogsResponseBodyDataDataSourceInstanceLogs {
	s.TaskStatus = &v
	return s
}

type ListDataSourceLogsResponseBodyDataDataSourceInstanceLogsLogParams struct {
	// The parameter code of the log.
	//
	// example:
	//
	// region_code
	ParaCode *string `json:"ParaCode,omitempty" xml:"ParaCode,omitempty"`
	// The parameter value of the log.
	//
	// example:
	//
	// ap-guangzhou
	ParaValue *string `json:"ParaValue,omitempty" xml:"ParaValue,omitempty"`
}

func (s ListDataSourceLogsResponseBodyDataDataSourceInstanceLogsLogParams) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourceLogsResponseBodyDataDataSourceInstanceLogsLogParams) GoString() string {
	return s.String()
}

func (s *ListDataSourceLogsResponseBodyDataDataSourceInstanceLogsLogParams) SetParaCode(v string) *ListDataSourceLogsResponseBodyDataDataSourceInstanceLogsLogParams {
	s.ParaCode = &v
	return s
}

func (s *ListDataSourceLogsResponseBodyDataDataSourceInstanceLogsLogParams) SetParaValue(v string) *ListDataSourceLogsResponseBodyDataDataSourceInstanceLogsLogParams {
	s.ParaValue = &v
	return s
}

type ListDataSourceLogsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataSourceLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataSourceLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourceLogsResponse) GoString() string {
	return s.String()
}

func (s *ListDataSourceLogsResponse) SetHeaders(v map[string]*string) *ListDataSourceLogsResponse {
	s.Headers = v
	return s
}

func (s *ListDataSourceLogsResponse) SetStatusCode(v int32) *ListDataSourceLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataSourceLogsResponse) SetBody(v *ListDataSourceLogsResponseBody) *ListDataSourceLogsResponse {
	s.Body = v
	return s
}

type ListDataSourceTypesRequest struct {
	// The code of the third-party cloud service.
	//
	// Valid values:
	//
	// 	- qcloud
	//
	// 	- hcloud
	//
	// This parameter is required.
	//
	// example:
	//
	// hcloud
	CloudCode *string `json:"CloudCode,omitempty" xml:"CloudCode,omitempty"`
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

func (s ListDataSourceTypesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourceTypesRequest) GoString() string {
	return s.String()
}

func (s *ListDataSourceTypesRequest) SetCloudCode(v string) *ListDataSourceTypesRequest {
	s.CloudCode = &v
	return s
}

func (s *ListDataSourceTypesRequest) SetRegionId(v string) *ListDataSourceTypesRequest {
	s.RegionId = &v
	return s
}

type ListDataSourceTypesResponseBody struct {
	// The data returned.
	Data []*ListDataSourceTypesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDataSourceTypesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourceTypesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataSourceTypesResponseBody) SetData(v []*ListDataSourceTypesResponseBodyData) *ListDataSourceTypesResponseBody {
	s.Data = v
	return s
}

func (s *ListDataSourceTypesResponseBody) SetRequestId(v string) *ListDataSourceTypesResponseBody {
	s.RequestId = &v
	return s
}

type ListDataSourceTypesResponseBodyData struct {
	// The code of the third-party cloud service.
	//
	// example:
	//
	// hcloud
	CloudCode *string `json:"CloudCode,omitempty" xml:"CloudCode,omitempty"`
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
}

func (s ListDataSourceTypesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourceTypesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDataSourceTypesResponseBodyData) SetCloudCode(v string) *ListDataSourceTypesResponseBodyData {
	s.CloudCode = &v
	return s
}

func (s *ListDataSourceTypesResponseBodyData) SetDataSourceType(v string) *ListDataSourceTypesResponseBodyData {
	s.DataSourceType = &v
	return s
}

type ListDataSourceTypesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataSourceTypesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataSourceTypesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourceTypesResponse) GoString() string {
	return s.String()
}

func (s *ListDataSourceTypesResponse) SetHeaders(v map[string]*string) *ListDataSourceTypesResponse {
	s.Headers = v
	return s
}

func (s *ListDataSourceTypesResponse) SetStatusCode(v int32) *ListDataSourceTypesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataSourceTypesResponse) SetBody(v *ListDataSourceTypesResponseBody) *ListDataSourceTypesResponse {
	s.Body = v
	return s
}

type ListDeliveryRequest struct {
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
	// The ID of the account that you switch from the management account.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view. Valid values:
	//
	// - 0: the current Alibaba Cloud account
	//
	// - 1: the global account
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s ListDeliveryRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDeliveryRequest) GoString() string {
	return s.String()
}

func (s *ListDeliveryRequest) SetRegionId(v string) *ListDeliveryRequest {
	s.RegionId = &v
	return s
}

func (s *ListDeliveryRequest) SetRoleFor(v int64) *ListDeliveryRequest {
	s.RoleFor = &v
	return s
}

func (s *ListDeliveryRequest) SetRoleType(v int32) *ListDeliveryRequest {
	s.RoleType = &v
	return s
}

type ListDeliveryResponseBody struct {
	// The response parameters.
	Data *ListDeliveryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 6276D891-58D4-55B2-87B9-74D413F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDeliveryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDeliveryResponseBody) GoString() string {
	return s.String()
}

func (s *ListDeliveryResponseBody) SetData(v *ListDeliveryResponseBodyData) *ListDeliveryResponseBody {
	s.Data = v
	return s
}

func (s *ListDeliveryResponseBody) SetRequestId(v string) *ListDeliveryResponseBody {
	s.RequestId = &v
	return s
}

type ListDeliveryResponseBodyData struct {
	// The URL that is displayed in charts.
	//
	// example:
	//
	// https://sls4service.console.aliyun.com/lognext/project/aliyun-cloudsiem-data-127608589417****-cn-shanghai
	//
	// /dashboard/cloud-siem?isShare=true&hideTopbar=true&hideSidebar=true&ignoreTabLocalStorage=true
	DashboardUrl *string `json:"DashboardUrl,omitempty" xml:"DashboardUrl,omitempty"`
	// Indicates whether the log delivery switch is displayed. Default value: true. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	DisplaySwitchOrNot *bool `json:"DisplaySwitchOrNot,omitempty" xml:"DisplaySwitchOrNot,omitempty"`
	// The name of the Logstore for the threat analysis feature on the user side. The value is in the cloud_siem format.
	//
	// example:
	//
	// cloud-siem
	LogStoreName *string `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
	// The cloud services.
	ProductList []*ListDeliveryResponseBodyDataProductList `json:"ProductList,omitempty" xml:"ProductList,omitempty" type:"Repeated"`
	// The name of the project for the threat analysis feature in Simple Log service on the user side. The value is in the aliyun-cloudsiem-data-${aliUid}-${region} format.
	//
	// example:
	//
	// aliyun-cloudsiem-data-127608589417****-cn-shanghai
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The URL that is used for log analysis.
	//
	// example:
	//
	// https://sls4service.console.aliyun.com/lognext/project/aliyun-cloudsiem-data-127608589417****-cn-shanghai
	//
	// /logsearch/cloud-siem?isShare=true&hideTopbar=true&hideSidebar=true&ignoreTabLocalStorage=true
	SearchUrl *string `json:"SearchUrl,omitempty" xml:"SearchUrl,omitempty"`
}

func (s ListDeliveryResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListDeliveryResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDeliveryResponseBodyData) SetDashboardUrl(v string) *ListDeliveryResponseBodyData {
	s.DashboardUrl = &v
	return s
}

func (s *ListDeliveryResponseBodyData) SetDisplaySwitchOrNot(v bool) *ListDeliveryResponseBodyData {
	s.DisplaySwitchOrNot = &v
	return s
}

func (s *ListDeliveryResponseBodyData) SetLogStoreName(v string) *ListDeliveryResponseBodyData {
	s.LogStoreName = &v
	return s
}

func (s *ListDeliveryResponseBodyData) SetProductList(v []*ListDeliveryResponseBodyDataProductList) *ListDeliveryResponseBodyData {
	s.ProductList = v
	return s
}

func (s *ListDeliveryResponseBodyData) SetProjectName(v string) *ListDeliveryResponseBodyData {
	s.ProjectName = &v
	return s
}

func (s *ListDeliveryResponseBodyData) SetSearchUrl(v string) *ListDeliveryResponseBodyData {
	s.SearchUrl = &v
	return s
}

type ListDeliveryResponseBodyDataProductList struct {
	// The logs of the cloud services.
	LogList []*ListDeliveryResponseBodyDataProductListLogList `json:"LogList,omitempty" xml:"LogList,omitempty" type:"Repeated"`
	// The log group. For example, in Security Center, the logs of hosts and networks are stored in different groups. Key indicates the group information, and value indicates the logs in the group.
	LogMap map[string][]*DataProductListLogMapValue `json:"LogMap,omitempty" xml:"LogMap,omitempty"`
	// The code of the cloud service. Valid values:
	//
	// 	- qcloud_waf
	//
	// 	- qlcoud_cfw
	//
	// 	- hcloud_waf
	//
	// 	- hcloud_cfw
	//
	// 	- ddos
	//
	// 	- sas
	//
	// 	- cfw
	//
	// 	- config
	//
	// 	- csk
	//
	// 	- fc
	//
	// 	- rds
	//
	// 	- nas
	//
	// 	- apigateway
	//
	// 	- cdn
	//
	// 	- mongodb
	//
	// 	- eip
	//
	// 	- slb
	//
	// 	- vpc
	//
	// 	- actiontrail
	//
	// 	- waf
	//
	// 	- bastionhost
	//
	// 	- oss
	//
	// 	- polardb
	//
	// example:
	//
	// sas
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// Security Center
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
}

func (s ListDeliveryResponseBodyDataProductList) String() string {
	return tea.Prettify(s)
}

func (s ListDeliveryResponseBodyDataProductList) GoString() string {
	return s.String()
}

func (s *ListDeliveryResponseBodyDataProductList) SetLogList(v []*ListDeliveryResponseBodyDataProductListLogList) *ListDeliveryResponseBodyDataProductList {
	s.LogList = v
	return s
}

func (s *ListDeliveryResponseBodyDataProductList) SetLogMap(v map[string][]*DataProductListLogMapValue) *ListDeliveryResponseBodyDataProductList {
	s.LogMap = v
	return s
}

func (s *ListDeliveryResponseBodyDataProductList) SetProductCode(v string) *ListDeliveryResponseBodyDataProductList {
	s.ProductCode = &v
	return s
}

func (s *ListDeliveryResponseBodyDataProductList) SetProductName(v string) *ListDeliveryResponseBodyDataProductList {
	s.ProductName = &v
	return s
}

type ListDeliveryResponseBodyDataProductListLogList struct {
	// Indicates whether the log delivery feature can be enabled or disabled. The feature can be enabled or disabled only by the administrator of the threat analysis feature. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	CanOperateOrNot *bool `json:"CanOperateOrNot,omitempty" xml:"CanOperateOrNot,omitempty"`
	// The extended parameter.
	ExtraParameters []*ListDeliveryResponseBodyDataProductListLogListExtraParameters `json:"ExtraParameters,omitempty" xml:"ExtraParameters,omitempty" type:"Repeated"`
	// The code of the log.
	//
	// example:
	//
	// cloud_siem_config_log
	LogCode *string `json:"LogCode,omitempty" xml:"LogCode,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// audit log
	LogName *string `json:"LogName,omitempty" xml:"LogName,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// audit log
	LogNameEn *string `json:"LogNameEn,omitempty" xml:"LogNameEn,omitempty"`
	// The language code of the log that is used to indicate the language in which the log is displayed.
	//
	// example:
	//
	// ${sas.cloudsiem.prod.cloud_siem_aegis_crack_from_beaver}
	LogNameKey *string `json:"LogNameKey,omitempty" xml:"LogNameKey,omitempty"`
	// The status of the log delivery. Valid values:
	//
	// 	- true: The logs are being delivered.
	//
	// 	- false: The log delivery feature is disabled.
	//
	// example:
	//
	// true
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
	// The topic of the log in the Logstore. The value is an index field in the Logstore that can be used to distinguish different logs.
	//
	// example:
	//
	// sas_login_event
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s ListDeliveryResponseBodyDataProductListLogList) String() string {
	return tea.Prettify(s)
}

func (s ListDeliveryResponseBodyDataProductListLogList) GoString() string {
	return s.String()
}

func (s *ListDeliveryResponseBodyDataProductListLogList) SetCanOperateOrNot(v bool) *ListDeliveryResponseBodyDataProductListLogList {
	s.CanOperateOrNot = &v
	return s
}

func (s *ListDeliveryResponseBodyDataProductListLogList) SetExtraParameters(v []*ListDeliveryResponseBodyDataProductListLogListExtraParameters) *ListDeliveryResponseBodyDataProductListLogList {
	s.ExtraParameters = v
	return s
}

func (s *ListDeliveryResponseBodyDataProductListLogList) SetLogCode(v string) *ListDeliveryResponseBodyDataProductListLogList {
	s.LogCode = &v
	return s
}

func (s *ListDeliveryResponseBodyDataProductListLogList) SetLogName(v string) *ListDeliveryResponseBodyDataProductListLogList {
	s.LogName = &v
	return s
}

func (s *ListDeliveryResponseBodyDataProductListLogList) SetLogNameEn(v string) *ListDeliveryResponseBodyDataProductListLogList {
	s.LogNameEn = &v
	return s
}

func (s *ListDeliveryResponseBodyDataProductListLogList) SetLogNameKey(v string) *ListDeliveryResponseBodyDataProductListLogList {
	s.LogNameKey = &v
	return s
}

func (s *ListDeliveryResponseBodyDataProductListLogList) SetStatus(v bool) *ListDeliveryResponseBodyDataProductListLogList {
	s.Status = &v
	return s
}

func (s *ListDeliveryResponseBodyDataProductListLogList) SetTopic(v string) *ListDeliveryResponseBodyDataProductListLogList {
	s.Topic = &v
	return s
}

type ListDeliveryResponseBodyDataProductListLogListExtraParameters struct {
	// The ID of the extended parameter.
	//
	// example:
	//
	// flag
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the extended parameter.
	//
	// example:
	//
	// value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDeliveryResponseBodyDataProductListLogListExtraParameters) String() string {
	return tea.Prettify(s)
}

func (s ListDeliveryResponseBodyDataProductListLogListExtraParameters) GoString() string {
	return s.String()
}

func (s *ListDeliveryResponseBodyDataProductListLogListExtraParameters) SetKey(v string) *ListDeliveryResponseBodyDataProductListLogListExtraParameters {
	s.Key = &v
	return s
}

func (s *ListDeliveryResponseBodyDataProductListLogListExtraParameters) SetValue(v string) *ListDeliveryResponseBodyDataProductListLogListExtraParameters {
	s.Value = &v
	return s
}

type ListDeliveryResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDeliveryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDeliveryResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDeliveryResponse) GoString() string {
	return s.String()
}

func (s *ListDeliveryResponse) SetHeaders(v map[string]*string) *ListDeliveryResponse {
	s.Headers = v
	return s
}

func (s *ListDeliveryResponse) SetStatusCode(v int32) *ListDeliveryResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDeliveryResponse) SetBody(v *ListDeliveryResponseBody) *ListDeliveryResponse {
	s.Body = v
	return s
}

type ListDisposeStrategyRequest struct {
	// The page number. Pages start from page 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The status of the policy. Valid values:
	//
	// 	- 0: invalid
	//
	// 	- 1: valid
	//
	// example:
	//
	// 0
	EffectiveStatus *int32 `json:"EffectiveStatus,omitempty" xml:"EffectiveStatus,omitempty"`
	// The end of the time range to query. Unit: milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1577808000000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The feature value of the entity. Fuzzy match is supported.
	//
	// example:
	//
	// test22.php
	EntityIdentity *string `json:"EntityIdentity,omitempty" xml:"EntityIdentity,omitempty"`
	// The entity type of the playbook. Valid values:
	//
	// 	- ip
	//
	// 	- process
	//
	// 	- file
	//
	// example:
	//
	// ip
	EntityType   *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	IncidentUuid *string `json:"IncidentUuid,omitempty" xml:"IncidentUuid,omitempty"`
	// The sort order. Valid values:
	//
	// 	- desc: descending order.
	//
	// 	- asc: ascending order.
	//
	// example:
	//
	// desc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The sort field. Valid values:
	//
	// 	- GmtModified: sorts the policies by update time.
	//
	// 	- GmtCreate: sorts the policies by creation time.
	//
	// 	- FinishTime: sorts the policies by end time.
	//
	// example:
	//
	// GmtModified
	OrderField *string `json:"OrderField,omitempty" xml:"OrderField,omitempty"`
	// The number of entries per page. Maximum value: 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the playbook, which is the unique identifier of the playbook.
	//
	// example:
	//
	// WafBlockIP
	PlaybookName *string `json:"PlaybookName,omitempty" xml:"PlaybookName,omitempty"`
	// The type of the playbook. Valid values:
	//
	// 	- system: user-triggered playbook
	//
	// 	- custom: event-triggered playbook
	//
	// 	- custom_alert: alert-triggered playbook
	//
	// 	- soar-manual: user-run playbook
	//
	// 	- soar-mdr: MDR-run playbook
	//
	// example:
	//
	// system
	PlaybookTypes *string `json:"PlaybookTypes,omitempty" xml:"PlaybookTypes,omitempty"`
	// The UUID of the playbook.
	//
	// example:
	//
	// system_aliyun_clb_process_book
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
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
	// The ID of the account that you switch from the management account.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view. Valid values:
	//
	// - 0: the current Alibaba Cloud account
	//
	// - 1: the global account
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// The ID of the SOAR handling policy.
	//
	// example:
	//
	// a50a49b7-6044-4593-ab15-2b46567caadd
	SophonTaskId *string `json:"SophonTaskId,omitempty" xml:"SophonTaskId,omitempty"`
	// The beginning of the time range to query. Unit: milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1577808000000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListDisposeStrategyRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDisposeStrategyRequest) GoString() string {
	return s.String()
}

func (s *ListDisposeStrategyRequest) SetCurrentPage(v int32) *ListDisposeStrategyRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListDisposeStrategyRequest) SetEffectiveStatus(v int32) *ListDisposeStrategyRequest {
	s.EffectiveStatus = &v
	return s
}

func (s *ListDisposeStrategyRequest) SetEndTime(v int64) *ListDisposeStrategyRequest {
	s.EndTime = &v
	return s
}

func (s *ListDisposeStrategyRequest) SetEntityIdentity(v string) *ListDisposeStrategyRequest {
	s.EntityIdentity = &v
	return s
}

func (s *ListDisposeStrategyRequest) SetEntityType(v string) *ListDisposeStrategyRequest {
	s.EntityType = &v
	return s
}

func (s *ListDisposeStrategyRequest) SetIncidentUuid(v string) *ListDisposeStrategyRequest {
	s.IncidentUuid = &v
	return s
}

func (s *ListDisposeStrategyRequest) SetOrder(v string) *ListDisposeStrategyRequest {
	s.Order = &v
	return s
}

func (s *ListDisposeStrategyRequest) SetOrderField(v string) *ListDisposeStrategyRequest {
	s.OrderField = &v
	return s
}

func (s *ListDisposeStrategyRequest) SetPageSize(v int32) *ListDisposeStrategyRequest {
	s.PageSize = &v
	return s
}

func (s *ListDisposeStrategyRequest) SetPlaybookName(v string) *ListDisposeStrategyRequest {
	s.PlaybookName = &v
	return s
}

func (s *ListDisposeStrategyRequest) SetPlaybookTypes(v string) *ListDisposeStrategyRequest {
	s.PlaybookTypes = &v
	return s
}

func (s *ListDisposeStrategyRequest) SetPlaybookUuid(v string) *ListDisposeStrategyRequest {
	s.PlaybookUuid = &v
	return s
}

func (s *ListDisposeStrategyRequest) SetRegionId(v string) *ListDisposeStrategyRequest {
	s.RegionId = &v
	return s
}

func (s *ListDisposeStrategyRequest) SetRoleFor(v int64) *ListDisposeStrategyRequest {
	s.RoleFor = &v
	return s
}

func (s *ListDisposeStrategyRequest) SetRoleType(v int32) *ListDisposeStrategyRequest {
	s.RoleType = &v
	return s
}

func (s *ListDisposeStrategyRequest) SetSophonTaskId(v string) *ListDisposeStrategyRequest {
	s.SophonTaskId = &v
	return s
}

func (s *ListDisposeStrategyRequest) SetStartTime(v int64) *ListDisposeStrategyRequest {
	s.StartTime = &v
	return s
}

type ListDisposeStrategyResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	//
	// example:
	//
	// 123456
	Data *ListDisposeStrategyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListDisposeStrategyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDisposeStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *ListDisposeStrategyResponseBody) SetCode(v int32) *ListDisposeStrategyResponseBody {
	s.Code = &v
	return s
}

func (s *ListDisposeStrategyResponseBody) SetData(v *ListDisposeStrategyResponseBodyData) *ListDisposeStrategyResponseBody {
	s.Data = v
	return s
}

func (s *ListDisposeStrategyResponseBody) SetMessage(v string) *ListDisposeStrategyResponseBody {
	s.Message = &v
	return s
}

func (s *ListDisposeStrategyResponseBody) SetRequestId(v string) *ListDisposeStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDisposeStrategyResponseBody) SetSuccess(v bool) *ListDisposeStrategyResponseBody {
	s.Success = &v
	return s
}

type ListDisposeStrategyResponseBodyData struct {
	// The pagination information.
	PageInfo *ListDisposeStrategyResponseBodyDataPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The detailed data.
	ResponseData []*ListDisposeStrategyResponseBodyDataResponseData `json:"ResponseData,omitempty" xml:"ResponseData,omitempty" type:"Repeated"`
}

func (s ListDisposeStrategyResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListDisposeStrategyResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDisposeStrategyResponseBodyData) SetPageInfo(v *ListDisposeStrategyResponseBodyDataPageInfo) *ListDisposeStrategyResponseBodyData {
	s.PageInfo = v
	return s
}

func (s *ListDisposeStrategyResponseBodyData) SetResponseData(v []*ListDisposeStrategyResponseBodyDataResponseData) *ListDisposeStrategyResponseBodyData {
	s.ResponseData = v
	return s
}

type ListDisposeStrategyResponseBodyDataPageInfo struct {
	// The current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDisposeStrategyResponseBodyDataPageInfo) String() string {
	return tea.Prettify(s)
}

func (s ListDisposeStrategyResponseBodyDataPageInfo) GoString() string {
	return s.String()
}

func (s *ListDisposeStrategyResponseBodyDataPageInfo) SetCurrentPage(v int32) *ListDisposeStrategyResponseBodyDataPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *ListDisposeStrategyResponseBodyDataPageInfo) SetPageSize(v int32) *ListDisposeStrategyResponseBodyDataPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListDisposeStrategyResponseBodyDataPageInfo) SetTotalCount(v int64) *ListDisposeStrategyResponseBodyDataPageInfo {
	s.TotalCount = &v
	return s
}

type ListDisposeStrategyResponseBodyDataResponseData struct {
	// The UUID of the alert.
	//
	// example:
	//
	// sas_71e24437d2797ce8fc59692905a4****
	AlertUuid *string `json:"AlertUuid,omitempty" xml:"AlertUuid,omitempty"`
	// The ID of the Alibaba Cloud account that is associated with the policy in SIEM.
	//
	// example:
	//
	// 127608589417****
	Aliuid *int64 `json:"Aliuid,omitempty" xml:"Aliuid,omitempty"`
	// The status of the policy. Valid values:
	//
	// 	- 0: invalid
	//
	// 	- 1: valid
	//
	// example:
	//
	// 0
	EffectiveStatus *int32 `json:"EffectiveStatus,omitempty" xml:"EffectiveStatus,omitempty"`
	// The details of the entity. The value is a JSON array.
	//
	// example:
	//
	// [{"ip":"1.1.1.1"}]
	Entity []interface{} `json:"Entity,omitempty" xml:"Entity,omitempty" type:"Repeated"`
	// The ID of the entity.
	//
	// example:
	//
	// 123456789
	EntityId *int64 `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The type of the entity. Valid values:
	//
	// 	- ip
	//
	// 	- process
	//
	// 	- file
	//
	// example:
	//
	// ip
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// The summary information about the failed task.
	//
	// example:
	//
	// DisposalEntity failed which description is Aegis Quarantine File , return_info failed which description is Check Aegis Process Result , [ERROR DETAIL] *******.php:file not found
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The end time of the task.
	//
	// example:
	//
	// 2021-08-10 21:34:07
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2021-01-06 16:37:29
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2021-01-06 16:37:29
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the policy.
	//
	// example:
	//
	// 123
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the event.
	//
	// example:
	//
	// Multiple type of alerts, including Miner Network, Command line download and run malicious files, Backdoor Process, etc
	IncidentName *string `json:"IncidentName,omitempty" xml:"IncidentName,omitempty"`
	// The UUID of the event.
	//
	// example:
	//
	// 85ea4241-798f-4684-a876-65d4f0c3****
	IncidentUuid *string `json:"IncidentUuid,omitempty" xml:"IncidentUuid,omitempty"`
	// The name of the playbook, which is the unique identifier of the playbook.
	//
	// example:
	//
	// WafBlockIP
	PlaybookName *string `json:"PlaybookName,omitempty" xml:"PlaybookName,omitempty"`
	// The type of the playbook. Valid values:
	//
	// 	- system: user-triggered playbook
	//
	// 	- custom: event-triggered playbook
	//
	// 	- custom_alert: alert-triggered playbook
	//
	// 	- soar-manual: user-run playbook
	//
	// 	- soar-mdr: MDR-run playbook
	//
	// example:
	//
	// system
	PlaybookType *string `json:"PlaybookType,omitempty" xml:"PlaybookType,omitempty"`
	// The UUID of the playbook.
	//
	// example:
	//
	// system_aliyun_clb_process_book
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
	// The scope of the policy.
	//
	// example:
	//
	// [{ aliUid: 1766185894104675 }]
	Scope []interface{} `json:"Scope,omitempty" xml:"Scope,omitempty" type:"Repeated"`
	// The ID of the SOAR handling policy.
	//
	// example:
	//
	// 577bbf90-a770-44a7-8154-586aa2d318fa
	SophonTaskId *string `json:"SophonTaskId,omitempty" xml:"SophonTaskId,omitempty"`
	// The running status of the playbook. Valid values:
	//
	// 	- 200: successful
	//
	// 	- 10: deleted
	//
	// 	- 5: failed
	//
	// 	- 0: initial
	//
	// example:
	//
	// 10
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the Alibaba account that is used to configure the policy.
	//
	// example:
	//
	// 176555323***
	SubAliuid *int64 `json:"SubAliuid,omitempty" xml:"SubAliuid,omitempty"`
	// The parameters that are used to trigger the playbook. The value is in the JSON format.
	//
	// example:
	//
	// {
	//
	//       "file": {
	//
	//             "op_code": "2",
	//
	//             "file_path": "/root/alert0913/a886.jsp",
	//
	//             "entity_type": "file",
	//
	//             "entity_name": "a886.jsp",
	//
	//             "file_name": "a886.jsp",
	//
	//             "file_owner": "USER:,GROUP:",
	//
	//             "hash_value": "5def10c9a4287d0920d86b42420b20b0",
	//
	//             "op_level": "2",
	//
	//             "entity_id": "/root/alert0913/a886.jsp",
	//
	//             "host_uuid": {
	//
	//                   "entity_type": "host",
	//
	//                   "entity_name": "N/A",
	//
	//                   "is_comprised": "1",
	//
	//                   "os_type": "linux",
	//
	//                   "entity_id": "5f58ef67-8803-4314-8d67-c87dc92b****",
	//
	//                   "host_uuid": "5f58ef67-8803-4314-8d67-c87dc92b****",
	//
	//                   "host_name": "N/A"
	//
	//             },
	//
	//             "malware_type": "${aliyun.siem.sas.alert_tag.webshell}"
	//
	//       },
	//
	//       "_sys_siem": {
	//
	//             "cloudCode": "aliyun",
	//
	//             "alertId": "89416745494****"
	//
	//       },
	//
	//       "scope": [
	//
	//             {
	//
	//                   "aliUid": 1766185894104****
	//
	//             }
	//
	//       ]
	//
	// }
	TaskParam *string `json:"TaskParam,omitempty" xml:"TaskParam,omitempty"`
	TaskUrl   *string `json:"TaskUrl,omitempty" xml:"TaskUrl,omitempty"`
}

func (s ListDisposeStrategyResponseBodyDataResponseData) String() string {
	return tea.Prettify(s)
}

func (s ListDisposeStrategyResponseBodyDataResponseData) GoString() string {
	return s.String()
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) SetAlertUuid(v string) *ListDisposeStrategyResponseBodyDataResponseData {
	s.AlertUuid = &v
	return s
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) SetAliuid(v int64) *ListDisposeStrategyResponseBodyDataResponseData {
	s.Aliuid = &v
	return s
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) SetEffectiveStatus(v int32) *ListDisposeStrategyResponseBodyDataResponseData {
	s.EffectiveStatus = &v
	return s
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) SetEntity(v []interface{}) *ListDisposeStrategyResponseBodyDataResponseData {
	s.Entity = v
	return s
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) SetEntityId(v int64) *ListDisposeStrategyResponseBodyDataResponseData {
	s.EntityId = &v
	return s
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) SetEntityType(v string) *ListDisposeStrategyResponseBodyDataResponseData {
	s.EntityType = &v
	return s
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) SetErrorMessage(v string) *ListDisposeStrategyResponseBodyDataResponseData {
	s.ErrorMessage = &v
	return s
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) SetFinishTime(v string) *ListDisposeStrategyResponseBodyDataResponseData {
	s.FinishTime = &v
	return s
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) SetGmtCreate(v string) *ListDisposeStrategyResponseBodyDataResponseData {
	s.GmtCreate = &v
	return s
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) SetGmtModified(v string) *ListDisposeStrategyResponseBodyDataResponseData {
	s.GmtModified = &v
	return s
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) SetId(v int64) *ListDisposeStrategyResponseBodyDataResponseData {
	s.Id = &v
	return s
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) SetIncidentName(v string) *ListDisposeStrategyResponseBodyDataResponseData {
	s.IncidentName = &v
	return s
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) SetIncidentUuid(v string) *ListDisposeStrategyResponseBodyDataResponseData {
	s.IncidentUuid = &v
	return s
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) SetPlaybookName(v string) *ListDisposeStrategyResponseBodyDataResponseData {
	s.PlaybookName = &v
	return s
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) SetPlaybookType(v string) *ListDisposeStrategyResponseBodyDataResponseData {
	s.PlaybookType = &v
	return s
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) SetPlaybookUuid(v string) *ListDisposeStrategyResponseBodyDataResponseData {
	s.PlaybookUuid = &v
	return s
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) SetScope(v []interface{}) *ListDisposeStrategyResponseBodyDataResponseData {
	s.Scope = v
	return s
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) SetSophonTaskId(v string) *ListDisposeStrategyResponseBodyDataResponseData {
	s.SophonTaskId = &v
	return s
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) SetStatus(v int32) *ListDisposeStrategyResponseBodyDataResponseData {
	s.Status = &v
	return s
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) SetSubAliuid(v int64) *ListDisposeStrategyResponseBodyDataResponseData {
	s.SubAliuid = &v
	return s
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) SetTaskParam(v string) *ListDisposeStrategyResponseBodyDataResponseData {
	s.TaskParam = &v
	return s
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) SetTaskUrl(v string) *ListDisposeStrategyResponseBodyDataResponseData {
	s.TaskUrl = &v
	return s
}

type ListDisposeStrategyResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDisposeStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDisposeStrategyResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDisposeStrategyResponse) GoString() string {
	return s.String()
}

func (s *ListDisposeStrategyResponse) SetHeaders(v map[string]*string) *ListDisposeStrategyResponse {
	s.Headers = v
	return s
}

func (s *ListDisposeStrategyResponse) SetStatusCode(v int32) *ListDisposeStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDisposeStrategyResponse) SetBody(v *ListDisposeStrategyResponseBody) *ListDisposeStrategyResponse {
	s.Body = v
	return s
}

type ListEntitiesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// host1****
	EntityName *string `json:"EntityName,omitempty" xml:"EntityName,omitempty"`
	// example:
	//
	// ip
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// example:
	//
	// 6c740667-80b2-476d-8924-2e706feb****
	EntityUuid *string `json:"EntityUuid,omitempty" xml:"EntityUuid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 85ea4241-798f-4684-a876-65d4f0c3****
	IncidentUuid *string `json:"IncidentUuid,omitempty" xml:"IncidentUuid,omitempty"`
	// example:
	//
	// 1
	IsMalwareEntity *string `json:"IsMalwareEntity,omitempty" xml:"IsMalwareEntity,omitempty"`
	// example:
	//
	// aliyun.siem.sas.alert_tag.miner_software
	MalwareType *string `json:"MalwareType,omitempty" xml:"MalwareType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// example:
	//
	// 1
	RoleType *int32  `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	Tags     *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s ListEntitiesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListEntitiesRequest) GoString() string {
	return s.String()
}

func (s *ListEntitiesRequest) SetCurrentPage(v int32) *ListEntitiesRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListEntitiesRequest) SetEntityName(v string) *ListEntitiesRequest {
	s.EntityName = &v
	return s
}

func (s *ListEntitiesRequest) SetEntityType(v string) *ListEntitiesRequest {
	s.EntityType = &v
	return s
}

func (s *ListEntitiesRequest) SetEntityUuid(v string) *ListEntitiesRequest {
	s.EntityUuid = &v
	return s
}

func (s *ListEntitiesRequest) SetIncidentUuid(v string) *ListEntitiesRequest {
	s.IncidentUuid = &v
	return s
}

func (s *ListEntitiesRequest) SetIsMalwareEntity(v string) *ListEntitiesRequest {
	s.IsMalwareEntity = &v
	return s
}

func (s *ListEntitiesRequest) SetMalwareType(v string) *ListEntitiesRequest {
	s.MalwareType = &v
	return s
}

func (s *ListEntitiesRequest) SetPageSize(v int32) *ListEntitiesRequest {
	s.PageSize = &v
	return s
}

func (s *ListEntitiesRequest) SetRegionId(v string) *ListEntitiesRequest {
	s.RegionId = &v
	return s
}

func (s *ListEntitiesRequest) SetRoleFor(v int64) *ListEntitiesRequest {
	s.RoleFor = &v
	return s
}

func (s *ListEntitiesRequest) SetRoleType(v int32) *ListEntitiesRequest {
	s.RoleType = &v
	return s
}

func (s *ListEntitiesRequest) SetTags(v string) *ListEntitiesRequest {
	s.Tags = &v
	return s
}

type ListEntitiesResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 123456
	Data *ListEntitiesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListEntitiesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEntitiesResponseBody) GoString() string {
	return s.String()
}

func (s *ListEntitiesResponseBody) SetCode(v int32) *ListEntitiesResponseBody {
	s.Code = &v
	return s
}

func (s *ListEntitiesResponseBody) SetData(v *ListEntitiesResponseBodyData) *ListEntitiesResponseBody {
	s.Data = v
	return s
}

func (s *ListEntitiesResponseBody) SetMessage(v string) *ListEntitiesResponseBody {
	s.Message = &v
	return s
}

func (s *ListEntitiesResponseBody) SetRequestId(v string) *ListEntitiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEntitiesResponseBody) SetSuccess(v bool) *ListEntitiesResponseBody {
	s.Success = &v
	return s
}

type ListEntitiesResponseBodyData struct {
	PageInfo     *ListEntitiesResponseBodyDataPageInfo       `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	ResponseData []*ListEntitiesResponseBodyDataResponseData `json:"ResponseData,omitempty" xml:"ResponseData,omitempty" type:"Repeated"`
}

func (s ListEntitiesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListEntitiesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListEntitiesResponseBodyData) SetPageInfo(v *ListEntitiesResponseBodyDataPageInfo) *ListEntitiesResponseBodyData {
	s.PageInfo = v
	return s
}

func (s *ListEntitiesResponseBodyData) SetResponseData(v []*ListEntitiesResponseBodyDataResponseData) *ListEntitiesResponseBodyData {
	s.ResponseData = v
	return s
}

type ListEntitiesResponseBodyDataPageInfo struct {
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListEntitiesResponseBodyDataPageInfo) String() string {
	return tea.Prettify(s)
}

func (s ListEntitiesResponseBodyDataPageInfo) GoString() string {
	return s.String()
}

func (s *ListEntitiesResponseBodyDataPageInfo) SetCurrentPage(v int32) *ListEntitiesResponseBodyDataPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *ListEntitiesResponseBodyDataPageInfo) SetPageSize(v int32) *ListEntitiesResponseBodyDataPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListEntitiesResponseBodyDataPageInfo) SetTotalCount(v int64) *ListEntitiesResponseBodyDataPageInfo {
	s.TotalCount = &v
	return s
}

type ListEntitiesResponseBodyDataResponseData struct {
	// example:
	//
	// 1
	AlertNum *int32 `json:"AlertNum,omitempty" xml:"AlertNum,omitempty"`
	// example:
	//
	// sas_71e24437d2797ce8fc59692905a4****
	AlertUuid *string `json:"AlertUuid,omitempty" xml:"AlertUuid,omitempty"`
	// example:
	//
	// 123456789****
	Aliuid *int64 `json:"Aliuid,omitempty" xml:"Aliuid,omitempty"`
	// example:
	//
	// aliyun
	CloudCode *string `json:"CloudCode,omitempty" xml:"CloudCode,omitempty"`
	// example:
	//
	// 12345****
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// example:
	//
	// {"file_path": "c:/www/leixi.jsp","file_hash": "aa0ca926ad948cd820e0a3d9a18c****","host_uuid": "efed2cf7-0b77-45d9-a97b-d2cf246b****","malware_type": "${aliyun.siem.sas.alert_tag.webshell}","host_name": "launch-advisor-2023****"}
	EntityInfo *string `json:"EntityInfo,omitempty" xml:"EntityInfo,omitempty"`
	// example:
	//
	// 123.123.123.123
	EntityName *string `json:"EntityName,omitempty" xml:"EntityName,omitempty"`
	// example:
	//
	// ip
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// example:
	//
	// 8087b3e4aa6862852c100c8738cf****
	EntityUuid *string `json:"EntityUuid,omitempty" xml:"EntityUuid,omitempty"`
	// example:
	//
	// 1
	EventNum *int32 `json:"EventNum,omitempty" xml:"EventNum,omitempty"`
	// example:
	//
	// 2021-01-06 16:37:29
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 2021-01-06 16:37:29
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 123456789***
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 85ea4241-798f-4684-a876-65d4f0c3****
	IncidentUuid *string `json:"IncidentUuid,omitempty" xml:"IncidentUuid,omitempty"`
	IsAsset      *string `json:"IsAsset,omitempty" xml:"IsAsset,omitempty"`
	IsMalware    *string `json:"IsMalware,omitempty" xml:"IsMalware,omitempty"`
	// example:
	//
	// aliyun.siem.sas.alert_tag.webshell
	MalwareType *string `json:"MalwareType,omitempty" xml:"MalwareType,omitempty"`
	// example:
	//
	// 113091674488****
	SubUserId *int64  `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
	Tags      *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s ListEntitiesResponseBodyDataResponseData) String() string {
	return tea.Prettify(s)
}

func (s ListEntitiesResponseBodyDataResponseData) GoString() string {
	return s.String()
}

func (s *ListEntitiesResponseBodyDataResponseData) SetAlertNum(v int32) *ListEntitiesResponseBodyDataResponseData {
	s.AlertNum = &v
	return s
}

func (s *ListEntitiesResponseBodyDataResponseData) SetAlertUuid(v string) *ListEntitiesResponseBodyDataResponseData {
	s.AlertUuid = &v
	return s
}

func (s *ListEntitiesResponseBodyDataResponseData) SetAliuid(v int64) *ListEntitiesResponseBodyDataResponseData {
	s.Aliuid = &v
	return s
}

func (s *ListEntitiesResponseBodyDataResponseData) SetCloudCode(v string) *ListEntitiesResponseBodyDataResponseData {
	s.CloudCode = &v
	return s
}

func (s *ListEntitiesResponseBodyDataResponseData) SetEntityId(v string) *ListEntitiesResponseBodyDataResponseData {
	s.EntityId = &v
	return s
}

func (s *ListEntitiesResponseBodyDataResponseData) SetEntityInfo(v string) *ListEntitiesResponseBodyDataResponseData {
	s.EntityInfo = &v
	return s
}

func (s *ListEntitiesResponseBodyDataResponseData) SetEntityName(v string) *ListEntitiesResponseBodyDataResponseData {
	s.EntityName = &v
	return s
}

func (s *ListEntitiesResponseBodyDataResponseData) SetEntityType(v string) *ListEntitiesResponseBodyDataResponseData {
	s.EntityType = &v
	return s
}

func (s *ListEntitiesResponseBodyDataResponseData) SetEntityUuid(v string) *ListEntitiesResponseBodyDataResponseData {
	s.EntityUuid = &v
	return s
}

func (s *ListEntitiesResponseBodyDataResponseData) SetEventNum(v int32) *ListEntitiesResponseBodyDataResponseData {
	s.EventNum = &v
	return s
}

func (s *ListEntitiesResponseBodyDataResponseData) SetGmtCreate(v string) *ListEntitiesResponseBodyDataResponseData {
	s.GmtCreate = &v
	return s
}

func (s *ListEntitiesResponseBodyDataResponseData) SetGmtModified(v string) *ListEntitiesResponseBodyDataResponseData {
	s.GmtModified = &v
	return s
}

func (s *ListEntitiesResponseBodyDataResponseData) SetId(v int64) *ListEntitiesResponseBodyDataResponseData {
	s.Id = &v
	return s
}

func (s *ListEntitiesResponseBodyDataResponseData) SetIncidentUuid(v string) *ListEntitiesResponseBodyDataResponseData {
	s.IncidentUuid = &v
	return s
}

func (s *ListEntitiesResponseBodyDataResponseData) SetIsAsset(v string) *ListEntitiesResponseBodyDataResponseData {
	s.IsAsset = &v
	return s
}

func (s *ListEntitiesResponseBodyDataResponseData) SetIsMalware(v string) *ListEntitiesResponseBodyDataResponseData {
	s.IsMalware = &v
	return s
}

func (s *ListEntitiesResponseBodyDataResponseData) SetMalwareType(v string) *ListEntitiesResponseBodyDataResponseData {
	s.MalwareType = &v
	return s
}

func (s *ListEntitiesResponseBodyDataResponseData) SetSubUserId(v int64) *ListEntitiesResponseBodyDataResponseData {
	s.SubUserId = &v
	return s
}

func (s *ListEntitiesResponseBodyDataResponseData) SetTags(v string) *ListEntitiesResponseBodyDataResponseData {
	s.Tags = &v
	return s
}

type ListEntitiesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEntitiesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEntitiesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEntitiesResponse) GoString() string {
	return s.String()
}

func (s *ListEntitiesResponse) SetHeaders(v map[string]*string) *ListEntitiesResponse {
	s.Headers = v
	return s
}

func (s *ListEntitiesResponse) SetStatusCode(v int32) *ListEntitiesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEntitiesResponse) SetBody(v *ListEntitiesResponseBody) *ListEntitiesResponse {
	s.Body = v
	return s
}

type ListImportedLogsByProdRequest struct {
	// The code of the cloud service provider. Valid values:
	//
	// 	- qcloud: Tencent Cloud.
	//
	// 	- aliyun: Alibaba Cloud.
	//
	// 	- hcloud: Huawei Cloud.
	//
	// This parameter is required.
	//
	// example:
	//
	// hcloud
	CloudCode *string `json:"CloudCode,omitempty" xml:"CloudCode,omitempty"`
	// The code of the cloud service.
	//
	// This parameter is required.
	//
	// example:
	//
	// qcloud_waf
	ProdCode *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
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
	// The ID of the account that you switch from the management account.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view. Valid values:
	//
	// - 0: the current Alibaba Cloud account
	//
	// - 1: the global account
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s ListImportedLogsByProdRequest) String() string {
	return tea.Prettify(s)
}

func (s ListImportedLogsByProdRequest) GoString() string {
	return s.String()
}

func (s *ListImportedLogsByProdRequest) SetCloudCode(v string) *ListImportedLogsByProdRequest {
	s.CloudCode = &v
	return s
}

func (s *ListImportedLogsByProdRequest) SetProdCode(v string) *ListImportedLogsByProdRequest {
	s.ProdCode = &v
	return s
}

func (s *ListImportedLogsByProdRequest) SetRegionId(v string) *ListImportedLogsByProdRequest {
	s.RegionId = &v
	return s
}

func (s *ListImportedLogsByProdRequest) SetRoleFor(v int64) *ListImportedLogsByProdRequest {
	s.RoleFor = &v
	return s
}

func (s *ListImportedLogsByProdRequest) SetRoleType(v int32) *ListImportedLogsByProdRequest {
	s.RoleType = &v
	return s
}

type ListImportedLogsByProdResponseBody struct {
	// The data returned.
	Data []*ListImportedLogsByProdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListImportedLogsByProdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListImportedLogsByProdResponseBody) GoString() string {
	return s.String()
}

func (s *ListImportedLogsByProdResponseBody) SetData(v []*ListImportedLogsByProdResponseBodyData) *ListImportedLogsByProdResponseBody {
	s.Data = v
	return s
}

func (s *ListImportedLogsByProdResponseBody) SetRequestId(v string) *ListImportedLogsByProdResponseBody {
	s.RequestId = &v
	return s
}

type ListImportedLogsByProdResponseBodyData struct {
	// Indicates whether the log is automatically added to the threat analysis feature within newly added accounts. Valid values:
	//
	// 	- 1: yes.
	//
	// 	- 0: no.
	//
	// example:
	//
	// 2023-11-23 12:30:00
	AutoImported *int32 `json:"AutoImported,omitempty" xml:"AutoImported,omitempty"`
	// The code of the cloud service provider. Valid values:
	//
	// 	- qcloud: Tencent Cloud.
	//
	// 	- aliyun: Alibaba Cloud.
	//
	// 	- hcloud: Huawei Cloud.
	//
	// example:
	//
	// hcloud
	CloudCode *string `json:"CloudCode,omitempty" xml:"CloudCode,omitempty"`
	// Indicates whether the log is added to the threat analysis feature. Valid values:
	//
	// 	- 1: yes.
	//
	// 	- 0: no.
	//
	// example:
	//
	// 2023-11-23 12:30:00
	Imported *int32 `json:"Imported,omitempty" xml:"Imported,omitempty"`
	// The number of users who have added the log.
	//
	// example:
	//
	// 2
	ImportedUserCount *int32 `json:"ImportedUserCount,omitempty" xml:"ImportedUserCount,omitempty"`
	// The code of the log.
	//
	// example:
	//
	// cloud_siem_waf_xxxxx
	LogCode *string `json:"LogCode,omitempty" xml:"LogCode,omitempty"`
	// The display code of the log.
	//
	// example:
	//
	// ${siem.prod. cloud_siem_waf_xxxxx}
	LogMdsCode *string `json:"LogMdsCode,omitempty" xml:"LogMdsCode,omitempty"`
	// The type of log. Valid values:
	//
	//  - 1: the log produced by other product
	//
	//  - 2: the predefined log
	//
	//  - 3: the custom log
	//
	// example:
	//
	// 1
	LogType *int32 `json:"LogType,omitempty" xml:"LogType,omitempty"`
	// The time when the log was last added.
	//
	// example:
	//
	// 2023-11-23 12:30:00
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The code of the cloud service to which the log belongs.
	//
	// example:
	//
	// qcloud_waf
	ProdCode *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	// The total number of users who have the log.
	//
	// example:
	//
	// 5
	TotalUserCount *int32 `json:"TotalUserCount,omitempty" xml:"TotalUserCount,omitempty"`
	// The number of users who have not added the log.
	//
	// example:
	//
	// 3
	UnImportedUserCount *int32 `json:"UnImportedUserCount,omitempty" xml:"UnImportedUserCount,omitempty"`
}

func (s ListImportedLogsByProdResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListImportedLogsByProdResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListImportedLogsByProdResponseBodyData) SetAutoImported(v int32) *ListImportedLogsByProdResponseBodyData {
	s.AutoImported = &v
	return s
}

func (s *ListImportedLogsByProdResponseBodyData) SetCloudCode(v string) *ListImportedLogsByProdResponseBodyData {
	s.CloudCode = &v
	return s
}

func (s *ListImportedLogsByProdResponseBodyData) SetImported(v int32) *ListImportedLogsByProdResponseBodyData {
	s.Imported = &v
	return s
}

func (s *ListImportedLogsByProdResponseBodyData) SetImportedUserCount(v int32) *ListImportedLogsByProdResponseBodyData {
	s.ImportedUserCount = &v
	return s
}

func (s *ListImportedLogsByProdResponseBodyData) SetLogCode(v string) *ListImportedLogsByProdResponseBodyData {
	s.LogCode = &v
	return s
}

func (s *ListImportedLogsByProdResponseBodyData) SetLogMdsCode(v string) *ListImportedLogsByProdResponseBodyData {
	s.LogMdsCode = &v
	return s
}

func (s *ListImportedLogsByProdResponseBodyData) SetLogType(v int32) *ListImportedLogsByProdResponseBodyData {
	s.LogType = &v
	return s
}

func (s *ListImportedLogsByProdResponseBodyData) SetModifyTime(v string) *ListImportedLogsByProdResponseBodyData {
	s.ModifyTime = &v
	return s
}

func (s *ListImportedLogsByProdResponseBodyData) SetProdCode(v string) *ListImportedLogsByProdResponseBodyData {
	s.ProdCode = &v
	return s
}

func (s *ListImportedLogsByProdResponseBodyData) SetTotalUserCount(v int32) *ListImportedLogsByProdResponseBodyData {
	s.TotalUserCount = &v
	return s
}

func (s *ListImportedLogsByProdResponseBodyData) SetUnImportedUserCount(v int32) *ListImportedLogsByProdResponseBodyData {
	s.UnImportedUserCount = &v
	return s
}

type ListImportedLogsByProdResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListImportedLogsByProdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListImportedLogsByProdResponse) String() string {
	return tea.Prettify(s)
}

func (s ListImportedLogsByProdResponse) GoString() string {
	return s.String()
}

func (s *ListImportedLogsByProdResponse) SetHeaders(v map[string]*string) *ListImportedLogsByProdResponse {
	s.Headers = v
	return s
}

func (s *ListImportedLogsByProdResponse) SetStatusCode(v int32) *ListImportedLogsByProdResponse {
	s.StatusCode = &v
	return s
}

func (s *ListImportedLogsByProdResponse) SetBody(v *ListImportedLogsByProdResponseBody) *ListImportedLogsByProdResponse {
	s.Body = v
	return s
}

type ListProjectLogStoresRequest struct {
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
	// The log code.
	//
	// This parameter is required.
	//
	// example:
	//
	// cloud_siem_aegis_proc
	SourceLogCode *string `json:"SourceLogCode,omitempty" xml:"SourceLogCode,omitempty"`
	// The code of the cloud service.
	//
	// This parameter is required.
	//
	// example:
	//
	// sas
	SourceProdCode *string `json:"SourceProdCode,omitempty" xml:"SourceProdCode,omitempty"`
	// The ID of the Alibaba Cloud account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123XXXXXXXX
	SubUserId *int64 `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
}

func (s ListProjectLogStoresRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProjectLogStoresRequest) GoString() string {
	return s.String()
}

func (s *ListProjectLogStoresRequest) SetRegionId(v string) *ListProjectLogStoresRequest {
	s.RegionId = &v
	return s
}

func (s *ListProjectLogStoresRequest) SetSourceLogCode(v string) *ListProjectLogStoresRequest {
	s.SourceLogCode = &v
	return s
}

func (s *ListProjectLogStoresRequest) SetSourceProdCode(v string) *ListProjectLogStoresRequest {
	s.SourceProdCode = &v
	return s
}

func (s *ListProjectLogStoresRequest) SetSubUserId(v int64) *ListProjectLogStoresRequest {
	s.SubUserId = &v
	return s
}

type ListProjectLogStoresResponseBody struct {
	// The data returned.
	Data []*ListProjectLogStoresResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListProjectLogStoresResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProjectLogStoresResponseBody) GoString() string {
	return s.String()
}

func (s *ListProjectLogStoresResponseBody) SetData(v []*ListProjectLogStoresResponseBodyData) *ListProjectLogStoresResponseBody {
	s.Data = v
	return s
}

func (s *ListProjectLogStoresResponseBody) SetRequestId(v string) *ListProjectLogStoresResponseBody {
	s.RequestId = &v
	return s
}

type ListProjectLogStoresResponseBodyData struct {
	// The endpoint of the Simple Log Service project.
	//
	// example:
	//
	// cn-hangzhou.log.aliyuncs.com
	EndPoint *string `json:"EndPoint,omitempty" xml:"EndPoint,omitempty"`
	// The name of the region in which the Simple Log Service project resides.
	//
	// example:
	//
	// hangzhou
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// The name of the Simple Log Service Logstore.
	//
	// example:
	//
	// cloud-siem-logstore
	LogStore *string `json:"LogStore,omitempty" xml:"LogStore,omitempty"`
	// The ID of the Alibaba Cloud account that is used to purchase the threat analysis feature.
	//
	// example:
	//
	// 123XXXXXXXXX
	MainUserId *int64 `json:"MainUserId,omitempty" xml:"MainUserId,omitempty"`
	// The name of the Simple Log Service project.
	//
	// example:
	//
	// cloud-siem-project
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
	// The ID of the region in which the Simple Log Service project resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud account that can be used to perform operations supported by the threat analysis feature.
	//
	// example:
	//
	// 123XXXXXXXX
	SubUserId *int64 `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
	// The username of the Alibaba Cloud account that can be used to perform operations supported by the threat analysis feature.
	//
	// example:
	//
	// sas_account_xxxx
	SubUserName *string `json:"SubUserName,omitempty" xml:"SubUserName,omitempty"`
}

func (s ListProjectLogStoresResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListProjectLogStoresResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListProjectLogStoresResponseBodyData) SetEndPoint(v string) *ListProjectLogStoresResponseBodyData {
	s.EndPoint = &v
	return s
}

func (s *ListProjectLogStoresResponseBodyData) SetLocalName(v string) *ListProjectLogStoresResponseBodyData {
	s.LocalName = &v
	return s
}

func (s *ListProjectLogStoresResponseBodyData) SetLogStore(v string) *ListProjectLogStoresResponseBodyData {
	s.LogStore = &v
	return s
}

func (s *ListProjectLogStoresResponseBodyData) SetMainUserId(v int64) *ListProjectLogStoresResponseBodyData {
	s.MainUserId = &v
	return s
}

func (s *ListProjectLogStoresResponseBodyData) SetProject(v string) *ListProjectLogStoresResponseBodyData {
	s.Project = &v
	return s
}

func (s *ListProjectLogStoresResponseBodyData) SetRegionId(v string) *ListProjectLogStoresResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *ListProjectLogStoresResponseBodyData) SetSubUserId(v int64) *ListProjectLogStoresResponseBodyData {
	s.SubUserId = &v
	return s
}

func (s *ListProjectLogStoresResponseBodyData) SetSubUserName(v string) *ListProjectLogStoresResponseBodyData {
	s.SubUserName = &v
	return s
}

type ListProjectLogStoresResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListProjectLogStoresResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListProjectLogStoresResponse) String() string {
	return tea.Prettify(s)
}

func (s ListProjectLogStoresResponse) GoString() string {
	return s.String()
}

func (s *ListProjectLogStoresResponse) SetHeaders(v map[string]*string) *ListProjectLogStoresResponse {
	s.Headers = v
	return s
}

func (s *ListProjectLogStoresResponse) SetStatusCode(v int32) *ListProjectLogStoresResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProjectLogStoresResponse) SetBody(v *ListProjectLogStoresResponseBody) *ListProjectLogStoresResponse {
	s.Body = v
	return s
}

type ListRdUsersRequest struct {
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

func (s ListRdUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRdUsersRequest) GoString() string {
	return s.String()
}

func (s *ListRdUsersRequest) SetRegionId(v string) *ListRdUsersRequest {
	s.RegionId = &v
	return s
}

type ListRdUsersResponseBody struct {
	// The data returned.
	Data []*ListRdUsersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListRdUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRdUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListRdUsersResponseBody) SetData(v []*ListRdUsersResponseBodyData) *ListRdUsersResponseBody {
	s.Data = v
	return s
}

func (s *ListRdUsersResponseBody) SetRequestId(v string) *ListRdUsersResponseBody {
	s.RequestId = &v
	return s
}

type ListRdUsersResponseBodyData struct {
	// Indicates whether the account can be used to view the logs and alerts within the account.
	//
	// example:
	//
	// true
	DelegatedOrNot *bool `json:"DelegatedOrNot,omitempty" xml:"DelegatedOrNot,omitempty"`
	// Indicates whether the account is added to the threat analysis feature for centralized management. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Joined *bool `json:"Joined,omitempty" xml:"Joined,omitempty"`
	// The time when the account was added to the threat analysis feature.
	//
	// example:
	//
	// 2013-10-01 00:00:00
	JoinedTime *string `json:"JoinedTime,omitempty" xml:"JoinedTime,omitempty"`
	// The ID of the Alibaba Cloud account that is used to purchase the threat analysis feature.
	//
	// example:
	//
	// 123XXXXXXXXX
	MainUserId *int64 `json:"MainUserId,omitempty" xml:"MainUserId,omitempty"`
	// The ID of the Alibaba Cloud account that is used to perform operations supported by the threat analysis feature.
	//
	// example:
	//
	// 123XXXXXXXX
	SubUserId *int64 `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
	// The username of the Alibaba Cloud account that can be used to perform operations supported by the threat analysis feature.
	//
	// example:
	//
	// sas_account_xxx
	SubUserName *string `json:"SubUserName,omitempty" xml:"SubUserName,omitempty"`
}

func (s ListRdUsersResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListRdUsersResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListRdUsersResponseBodyData) SetDelegatedOrNot(v bool) *ListRdUsersResponseBodyData {
	s.DelegatedOrNot = &v
	return s
}

func (s *ListRdUsersResponseBodyData) SetJoined(v bool) *ListRdUsersResponseBodyData {
	s.Joined = &v
	return s
}

func (s *ListRdUsersResponseBodyData) SetJoinedTime(v string) *ListRdUsersResponseBodyData {
	s.JoinedTime = &v
	return s
}

func (s *ListRdUsersResponseBodyData) SetMainUserId(v int64) *ListRdUsersResponseBodyData {
	s.MainUserId = &v
	return s
}

func (s *ListRdUsersResponseBodyData) SetSubUserId(v int64) *ListRdUsersResponseBodyData {
	s.SubUserId = &v
	return s
}

func (s *ListRdUsersResponseBodyData) SetSubUserName(v string) *ListRdUsersResponseBodyData {
	s.SubUserName = &v
	return s
}

type ListRdUsersResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRdUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRdUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRdUsersResponse) GoString() string {
	return s.String()
}

func (s *ListRdUsersResponse) SetHeaders(v map[string]*string) *ListRdUsersResponse {
	s.Headers = v
	return s
}

func (s *ListRdUsersResponse) SetStatusCode(v int32) *ListRdUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRdUsersResponse) SetBody(v *ListRdUsersResponseBody) *ListRdUsersResponse {
	s.Body = v
	return s
}

type ModifyBindAccountRequest struct {
	// The AccessKey ID of the cloud account.
	//
	// example:
	//
	// ABCXXXXXXXXX
	AccessId *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	// The ID of the cloud account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123xxxxxxx
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The username of the cloud account.
	//
	// example:
	//
	// sas_account_xxx
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The ID that is generated by the system when the account is added. You can call the ListBindAccount operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	BindId *int64 `json:"BindId,omitempty" xml:"BindId,omitempty"`
	// The code of the cloud service provider.
	//
	// This parameter is required.
	//
	// example:
	//
	// hcloud
	CloudCode *string `json:"CloudCode,omitempty" xml:"CloudCode,omitempty"`
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
	RoleFor  *int64  `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	RoleType *int32  `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s ModifyBindAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyBindAccountRequest) GoString() string {
	return s.String()
}

func (s *ModifyBindAccountRequest) SetAccessId(v string) *ModifyBindAccountRequest {
	s.AccessId = &v
	return s
}

func (s *ModifyBindAccountRequest) SetAccountId(v string) *ModifyBindAccountRequest {
	s.AccountId = &v
	return s
}

func (s *ModifyBindAccountRequest) SetAccountName(v string) *ModifyBindAccountRequest {
	s.AccountName = &v
	return s
}

func (s *ModifyBindAccountRequest) SetBindId(v int64) *ModifyBindAccountRequest {
	s.BindId = &v
	return s
}

func (s *ModifyBindAccountRequest) SetCloudCode(v string) *ModifyBindAccountRequest {
	s.CloudCode = &v
	return s
}

func (s *ModifyBindAccountRequest) SetRegionId(v string) *ModifyBindAccountRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyBindAccountRequest) SetRoleFor(v int64) *ModifyBindAccountRequest {
	s.RoleFor = &v
	return s
}

func (s *ModifyBindAccountRequest) SetRoleType(v int32) *ModifyBindAccountRequest {
	s.RoleType = &v
	return s
}

type ModifyBindAccountResponseBody struct {
	// The data returned.
	Data *ModifyBindAccountResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyBindAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyBindAccountResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyBindAccountResponseBody) SetData(v *ModifyBindAccountResponseBodyData) *ModifyBindAccountResponseBody {
	s.Data = v
	return s
}

func (s *ModifyBindAccountResponseBody) SetRequestId(v string) *ModifyBindAccountResponseBody {
	s.RequestId = &v
	return s
}

type ModifyBindAccountResponseBodyData struct {
	// The number of the accounts that are modified. The value 1 indicates that the modification is successful, and a value less than or equal to 0 indicates that the modification failed.
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s ModifyBindAccountResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ModifyBindAccountResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModifyBindAccountResponseBodyData) SetCount(v int32) *ModifyBindAccountResponseBodyData {
	s.Count = &v
	return s
}

type ModifyBindAccountResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyBindAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyBindAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyBindAccountResponse) GoString() string {
	return s.String()
}

func (s *ModifyBindAccountResponse) SetHeaders(v map[string]*string) *ModifyBindAccountResponse {
	s.Headers = v
	return s
}

func (s *ModifyBindAccountResponse) SetStatusCode(v int32) *ModifyBindAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyBindAccountResponse) SetBody(v *ModifyBindAccountResponseBody) *ModifyBindAccountResponse {
	s.Body = v
	return s
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
	return tea.Prettify(s)
}

func (s ModifyDataSourceRequest) GoString() string {
	return s.String()
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

type ModifyDataSourceResponseBody struct {
	// The data returned.
	Data *ModifyDataSourceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDataSourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDataSourceResponseBody) SetData(v *ModifyDataSourceResponseBodyData) *ModifyDataSourceResponseBody {
	s.Data = v
	return s
}

func (s *ModifyDataSourceResponseBody) SetRequestId(v string) *ModifyDataSourceResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDataSourceResponseBodyData struct {
	// The number of data sources that are modified. The value 1 indicates that the modification is successful, and a value less than or equal to 0 indicates that the modification failed.
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The ID of the data source. The ID is an MD5 hash value that is calculated by the threat analysis feature based on specific parameters.
	//
	// example:
	//
	// 220ba97c9d1fdb0b9c7e8c7ca328d7ea
	DataSourceInstanceId *string `json:"DataSourceInstanceId,omitempty" xml:"DataSourceInstanceId,omitempty"`
}

func (s ModifyDataSourceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ModifyDataSourceResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModifyDataSourceResponseBodyData) SetCount(v int32) *ModifyDataSourceResponseBodyData {
	s.Count = &v
	return s
}

func (s *ModifyDataSourceResponseBodyData) SetDataSourceInstanceId(v string) *ModifyDataSourceResponseBodyData {
	s.DataSourceInstanceId = &v
	return s
}

type ModifyDataSourceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDataSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDataSourceResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDataSourceResponse) GoString() string {
	return s.String()
}

func (s *ModifyDataSourceResponse) SetHeaders(v map[string]*string) *ModifyDataSourceResponse {
	s.Headers = v
	return s
}

func (s *ModifyDataSourceResponse) SetStatusCode(v int32) *ModifyDataSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDataSourceResponse) SetBody(v *ModifyDataSourceResponseBody) *ModifyDataSourceResponse {
	s.Body = v
	return s
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
	return tea.Prettify(s)
}

func (s ModifyDataSourceLogRequest) GoString() string {
	return s.String()
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

type ModifyDataSourceLogResponseBody struct {
	// The data returned.
	Data *ModifyDataSourceLogResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDataSourceLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDataSourceLogResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDataSourceLogResponseBody) SetData(v *ModifyDataSourceLogResponseBodyData) *ModifyDataSourceLogResponseBody {
	s.Data = v
	return s
}

func (s *ModifyDataSourceLogResponseBody) SetRequestId(v string) *ModifyDataSourceLogResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDataSourceLogResponseBodyData struct {
	// The number of logs that are modified. The value 1 indicates that the modification is successful, and a value less than or equal to 0 indicates that the modification failed.
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The ID of the log. The ID is an MD5 hash value that is calculated by the threat analysis feature based on specific parameters.
	//
	// example:
	//
	// 220ba97c9d1fdb0b9c7e8c7ca328d7ea
	LogInstanceId *string `json:"LogInstanceId,omitempty" xml:"LogInstanceId,omitempty"`
}

func (s ModifyDataSourceLogResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ModifyDataSourceLogResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModifyDataSourceLogResponseBodyData) SetCount(v int32) *ModifyDataSourceLogResponseBodyData {
	s.Count = &v
	return s
}

func (s *ModifyDataSourceLogResponseBodyData) SetLogInstanceId(v string) *ModifyDataSourceLogResponseBodyData {
	s.LogInstanceId = &v
	return s
}

type ModifyDataSourceLogResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDataSourceLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDataSourceLogResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDataSourceLogResponse) GoString() string {
	return s.String()
}

func (s *ModifyDataSourceLogResponse) SetHeaders(v map[string]*string) *ModifyDataSourceLogResponse {
	s.Headers = v
	return s
}

func (s *ModifyDataSourceLogResponse) SetStatusCode(v int32) *ModifyDataSourceLogResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDataSourceLogResponse) SetBody(v *ModifyDataSourceLogResponseBody) *ModifyDataSourceLogResponse {
	s.Body = v
	return s
}

type OpenDeliveryRequest struct {
	// The log code of the cloud service, such as the code of the process log for Security Center. This parameter is optional. If you leave this parameter empty, operations are performed on all logs of the cloud service.
	//
	// example:
	//
	// cloud_siem_cfw_flow
	LogCode *string `json:"LogCode,omitempty" xml:"LogCode,omitempty"`
	// The code of the cloud service. Valid values:
	//
	// 	- qcloud_waf
	//
	// 	- qlcoud_cfw
	//
	// 	- hcloud_waf
	//
	// 	- hcloud_cfw
	//
	// 	- ddos
	//
	// 	- sas
	//
	// 	- cfw
	//
	// 	- config
	//
	// 	- csk
	//
	// 	- fc
	//
	// 	- rds
	//
	// 	- nas
	//
	// 	- apigateway
	//
	// 	- cdn
	//
	// 	- mongodb
	//
	// 	- eip
	//
	// 	- slb
	//
	// 	- vpc
	//
	// 	- actiontrail
	//
	// 	- waf
	//
	// 	- bastionhost
	//
	// 	- oss
	//
	// 	- polardb
	//
	// This parameter is required.
	//
	// example:
	//
	// cfw
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The region in which the data management center of the threat analysis feature resides. Specify this parameter based on the region where your assets reside. Valid values:
	//
	// 	- cn-hangzhou: Your assets reside in regions in China.
	//
	// 	- ap-southeast-1: Your assets reside in regions outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the account that you switch from the management account.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view. Valid values:
	//
	// - 0: the current Alibaba Cloud account
	//
	// - 1: the global account
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s OpenDeliveryRequest) String() string {
	return tea.Prettify(s)
}

func (s OpenDeliveryRequest) GoString() string {
	return s.String()
}

func (s *OpenDeliveryRequest) SetLogCode(v string) *OpenDeliveryRequest {
	s.LogCode = &v
	return s
}

func (s *OpenDeliveryRequest) SetProductCode(v string) *OpenDeliveryRequest {
	s.ProductCode = &v
	return s
}

func (s *OpenDeliveryRequest) SetRegionId(v string) *OpenDeliveryRequest {
	s.RegionId = &v
	return s
}

func (s *OpenDeliveryRequest) SetRoleFor(v int64) *OpenDeliveryRequest {
	s.RoleFor = &v
	return s
}

func (s *OpenDeliveryRequest) SetRoleType(v int32) *OpenDeliveryRequest {
	s.RoleType = &v
	return s
}

type OpenDeliveryResponseBody struct {
	// Indicates whether the log delivery feature is enabled. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 15FD134E-D69B-51E8-B052-73F97BD8****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OpenDeliveryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OpenDeliveryResponseBody) GoString() string {
	return s.String()
}

func (s *OpenDeliveryResponseBody) SetData(v bool) *OpenDeliveryResponseBody {
	s.Data = &v
	return s
}

func (s *OpenDeliveryResponseBody) SetRequestId(v string) *OpenDeliveryResponseBody {
	s.RequestId = &v
	return s
}

type OpenDeliveryResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenDeliveryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenDeliveryResponse) String() string {
	return tea.Prettify(s)
}

func (s OpenDeliveryResponse) GoString() string {
	return s.String()
}

func (s *OpenDeliveryResponse) SetHeaders(v map[string]*string) *OpenDeliveryResponse {
	s.Headers = v
	return s
}

func (s *OpenDeliveryResponse) SetStatusCode(v int32) *OpenDeliveryResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenDeliveryResponse) SetBody(v *OpenDeliveryResponseBody) *OpenDeliveryResponse {
	s.Body = v
	return s
}

type PostAutomateResponseConfigRequest struct {
	// The action configuration of the automated response rule. The value is in the JSON format.
	//
	// example:
	//
	// [
	//
	//       {
	//
	//             "actionType": "doPlaybook",
	//
	//             "playbookName": "WafBlockIP",
	//
	//             "playbookUuid": "bdad6220-6584-41b2-9704-fc6584568758"
	//
	//       }
	//
	// ]
	ActionConfig *string `json:"ActionConfig,omitempty" xml:"ActionConfig,omitempty"`
	// The type of the handling action. Multiple types are separated by commas (,). Valid values:
	//
	// 	- **doPlaybook**: runs the playbook.
	//
	// 	- **changeEventStatus**: changes the event status.
	//
	// 	- **changeThreatLevel**: changes the threat level of the event.
	//
	// example:
	//
	// doPlaybook,changeEventStatus
	ActionType *string `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	// The type of the automated response rule. Valid values:
	//
	// 	- **event**
	//
	// 	- **alert**
	//
	// example:
	//
	// event
	AutoResponseType *string `json:"AutoResponseType,omitempty" xml:"AutoResponseType,omitempty"`
	// The trigger condition of the automated response rule. The value is in the JSON format.
	//
	// example:
	//
	// [{"left":{"value":"alert_name"},"operator":"containsString","right":{"value":"webshell_online"}}]
	ExecutionCondition *string `json:"ExecutionCondition,omitempty" xml:"ExecutionCondition,omitempty"`
	// The rule ID.
	//
	// example:
	//
	// 123
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The data management center of the threat analysis feature. Specify this parameter based on the regions in which your assets reside. Valid values:
	//
	// 	- **cn-hangzhou**: Your assets reside in regions in China.
	//
	// 	- **ap-southeast-1**: Your assets reside in regions outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the account that you switch from the management account.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view. Valid values:
	//
	// - 0: the current Alibaba Cloud account
	//
	// - 1: the global account
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// The rule name.
	//
	// example:
	//
	// cfw kill quara book
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The ID of the user who created the rule.
	//
	// example:
	//
	// 17108579417****
	SubUserId *int64 `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
}

func (s PostAutomateResponseConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s PostAutomateResponseConfigRequest) GoString() string {
	return s.String()
}

func (s *PostAutomateResponseConfigRequest) SetActionConfig(v string) *PostAutomateResponseConfigRequest {
	s.ActionConfig = &v
	return s
}

func (s *PostAutomateResponseConfigRequest) SetActionType(v string) *PostAutomateResponseConfigRequest {
	s.ActionType = &v
	return s
}

func (s *PostAutomateResponseConfigRequest) SetAutoResponseType(v string) *PostAutomateResponseConfigRequest {
	s.AutoResponseType = &v
	return s
}

func (s *PostAutomateResponseConfigRequest) SetExecutionCondition(v string) *PostAutomateResponseConfigRequest {
	s.ExecutionCondition = &v
	return s
}

func (s *PostAutomateResponseConfigRequest) SetId(v int64) *PostAutomateResponseConfigRequest {
	s.Id = &v
	return s
}

func (s *PostAutomateResponseConfigRequest) SetRegionId(v string) *PostAutomateResponseConfigRequest {
	s.RegionId = &v
	return s
}

func (s *PostAutomateResponseConfigRequest) SetRoleFor(v int64) *PostAutomateResponseConfigRequest {
	s.RoleFor = &v
	return s
}

func (s *PostAutomateResponseConfigRequest) SetRoleType(v int32) *PostAutomateResponseConfigRequest {
	s.RoleType = &v
	return s
}

func (s *PostAutomateResponseConfigRequest) SetRuleName(v string) *PostAutomateResponseConfigRequest {
	s.RuleName = &v
	return s
}

func (s *PostAutomateResponseConfigRequest) SetSubUserId(v int64) *PostAutomateResponseConfigRequest {
	s.SubUserId = &v
	return s
}

type PostAutomateResponseConfigResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	//
	// example:
	//
	// 123456
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PostAutomateResponseConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PostAutomateResponseConfigResponseBody) GoString() string {
	return s.String()
}

func (s *PostAutomateResponseConfigResponseBody) SetCode(v int32) *PostAutomateResponseConfigResponseBody {
	s.Code = &v
	return s
}

func (s *PostAutomateResponseConfigResponseBody) SetData(v string) *PostAutomateResponseConfigResponseBody {
	s.Data = &v
	return s
}

func (s *PostAutomateResponseConfigResponseBody) SetMessage(v string) *PostAutomateResponseConfigResponseBody {
	s.Message = &v
	return s
}

func (s *PostAutomateResponseConfigResponseBody) SetRequestId(v string) *PostAutomateResponseConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *PostAutomateResponseConfigResponseBody) SetSuccess(v bool) *PostAutomateResponseConfigResponseBody {
	s.Success = &v
	return s
}

type PostAutomateResponseConfigResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PostAutomateResponseConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PostAutomateResponseConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s PostAutomateResponseConfigResponse) GoString() string {
	return s.String()
}

func (s *PostAutomateResponseConfigResponse) SetHeaders(v map[string]*string) *PostAutomateResponseConfigResponse {
	s.Headers = v
	return s
}

func (s *PostAutomateResponseConfigResponse) SetStatusCode(v int32) *PostAutomateResponseConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *PostAutomateResponseConfigResponse) SetBody(v *PostAutomateResponseConfigResponseBody) *PostAutomateResponseConfigResponse {
	s.Body = v
	return s
}

type PostCustomizeRuleRequest struct {
	// The risk type.
	//
	// example:
	//
	// WEBSHELL
	AlertType *string `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
	// The internal code of the risk type.
	//
	// example:
	//
	// ${siem_rule_type_process_abnormal_command}
	AlertTypeMds *string `json:"AlertTypeMds,omitempty" xml:"AlertTypeMds,omitempty"`
	// att&ck.
	//
	// example:
	//
	// T1595.002 Vulnerability Scanning
	AttCk *string `json:"AttCk,omitempty" xml:"AttCk,omitempty"`
	// The extended information about event generation. If eventTransferType is set to allToSingle, the value of this parameter indicates the length and unit of the alert aggregation window.
	//
	// example:
	//
	// {"time":"1","unit":"MINUTE"}
	EventTransferExt *string `json:"EventTransferExt,omitempty" xml:"EventTransferExt,omitempty"`
	// Specifies whether to convert an alert to an event. Valid values:
	//
	// 	- 0: no
	//
	// 	- 1: yes
	//
	// example:
	//
	// 1
	EventTransferSwitch *int32 `json:"EventTransferSwitch,omitempty" xml:"EventTransferSwitch,omitempty"`
	// The event generation method. Valid values:
	//
	// 	- default: The default method is used.
	//
	// 	- singleToSingle: The system generates an event for each alert.
	//
	// 	- allToSingle: The system generates an event for alerts within a period of time.
	//
	// example:
	//
	// allToSingle
	EventTransferType *string `json:"EventTransferType,omitempty" xml:"EventTransferType,omitempty"`
	// The ID of the rule.
	//
	// example:
	//
	// 123456789
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The log source of the rule.
	//
	// example:
	//
	// cloud_siem_aegis_sas_alert
	LogSource *string `json:"LogSource,omitempty" xml:"LogSource,omitempty"`
	// The internal code of the log source.
	//
	// example:
	//
	// ${sas.cloudsiem.prod.cloud_siem_aegis_sas_alert}
	LogSourceMds *string `json:"LogSourceMds,omitempty" xml:"LogSourceMds,omitempty"`
	// The log type of the rule.
	//
	// example:
	//
	// ALERT_ACTIVITY
	LogType *string `json:"LogType,omitempty" xml:"LogType,omitempty"`
	// The internal code of the log type.
	//
	// example:
	//
	// ${security_event_config.event_name.webshellName_clientav}
	LogTypeMds *string `json:"LogTypeMds,omitempty" xml:"LogTypeMds,omitempty"`
	// The window length of the rule.
	//
	// example:
	//
	// {"time":"1","unit":"HOUR"}
	QueryCycle *string `json:"QueryCycle,omitempty" xml:"QueryCycle,omitempty"`
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
	// The ID of the account that you switch from the management account.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view. Valid values:
	//
	// - 0: the current Alibaba Cloud account
	//
	// - 1: the global account
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// The query condition of the rule. The value is in the JSON format.
	//
	// example:
	//
	// [[{"not":false,"left":"alert_name","operator":"=","right":"WEBSHELL"}]]
	RuleCondition *string `json:"RuleCondition,omitempty" xml:"RuleCondition,omitempty"`
	// The description of the rule.
	//
	// example:
	//
	// this rule is for waf scan
	RuleDesc *string `json:"RuleDesc,omitempty" xml:"RuleDesc,omitempty"`
	// The log aggregation field of the rule. The value is a JSON string.
	//
	// example:
	//
	// ["asset_id"]
	RuleGroup *string `json:"RuleGroup,omitempty" xml:"RuleGroup,omitempty"`
	// The name of the rule.
	//
	// example:
	//
	// waf_scan
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The threshold configuration of the rule. The value is in the JSON format.
	//
	// example:
	//
	// {"aggregateFunction":"count","aggregateFunctionName":"count","field":"activity_name","operator":"&lt;=","value":1}
	RuleThreshold *string `json:"RuleThreshold,omitempty" xml:"RuleThreshold,omitempty"`
	// The risk level. Valid values:
	//
	// 	- serious: high
	//
	// 	- suspicious: medium
	//
	// 	- remind: low
	//
	// example:
	//
	// remind
	ThreatLevel *string `json:"ThreatLevel,omitempty" xml:"ThreatLevel,omitempty"`
}

func (s PostCustomizeRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s PostCustomizeRuleRequest) GoString() string {
	return s.String()
}

func (s *PostCustomizeRuleRequest) SetAlertType(v string) *PostCustomizeRuleRequest {
	s.AlertType = &v
	return s
}

func (s *PostCustomizeRuleRequest) SetAlertTypeMds(v string) *PostCustomizeRuleRequest {
	s.AlertTypeMds = &v
	return s
}

func (s *PostCustomizeRuleRequest) SetAttCk(v string) *PostCustomizeRuleRequest {
	s.AttCk = &v
	return s
}

func (s *PostCustomizeRuleRequest) SetEventTransferExt(v string) *PostCustomizeRuleRequest {
	s.EventTransferExt = &v
	return s
}

func (s *PostCustomizeRuleRequest) SetEventTransferSwitch(v int32) *PostCustomizeRuleRequest {
	s.EventTransferSwitch = &v
	return s
}

func (s *PostCustomizeRuleRequest) SetEventTransferType(v string) *PostCustomizeRuleRequest {
	s.EventTransferType = &v
	return s
}

func (s *PostCustomizeRuleRequest) SetId(v int64) *PostCustomizeRuleRequest {
	s.Id = &v
	return s
}

func (s *PostCustomizeRuleRequest) SetLogSource(v string) *PostCustomizeRuleRequest {
	s.LogSource = &v
	return s
}

func (s *PostCustomizeRuleRequest) SetLogSourceMds(v string) *PostCustomizeRuleRequest {
	s.LogSourceMds = &v
	return s
}

func (s *PostCustomizeRuleRequest) SetLogType(v string) *PostCustomizeRuleRequest {
	s.LogType = &v
	return s
}

func (s *PostCustomizeRuleRequest) SetLogTypeMds(v string) *PostCustomizeRuleRequest {
	s.LogTypeMds = &v
	return s
}

func (s *PostCustomizeRuleRequest) SetQueryCycle(v string) *PostCustomizeRuleRequest {
	s.QueryCycle = &v
	return s
}

func (s *PostCustomizeRuleRequest) SetRegionId(v string) *PostCustomizeRuleRequest {
	s.RegionId = &v
	return s
}

func (s *PostCustomizeRuleRequest) SetRoleFor(v int64) *PostCustomizeRuleRequest {
	s.RoleFor = &v
	return s
}

func (s *PostCustomizeRuleRequest) SetRoleType(v int32) *PostCustomizeRuleRequest {
	s.RoleType = &v
	return s
}

func (s *PostCustomizeRuleRequest) SetRuleCondition(v string) *PostCustomizeRuleRequest {
	s.RuleCondition = &v
	return s
}

func (s *PostCustomizeRuleRequest) SetRuleDesc(v string) *PostCustomizeRuleRequest {
	s.RuleDesc = &v
	return s
}

func (s *PostCustomizeRuleRequest) SetRuleGroup(v string) *PostCustomizeRuleRequest {
	s.RuleGroup = &v
	return s
}

func (s *PostCustomizeRuleRequest) SetRuleName(v string) *PostCustomizeRuleRequest {
	s.RuleName = &v
	return s
}

func (s *PostCustomizeRuleRequest) SetRuleThreshold(v string) *PostCustomizeRuleRequest {
	s.RuleThreshold = &v
	return s
}

func (s *PostCustomizeRuleRequest) SetThreatLevel(v string) *PostCustomizeRuleRequest {
	s.ThreatLevel = &v
	return s
}

type PostCustomizeRuleResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	//
	// example:
	//
	// 123456
	Data *PostCustomizeRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PostCustomizeRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PostCustomizeRuleResponseBody) GoString() string {
	return s.String()
}

func (s *PostCustomizeRuleResponseBody) SetCode(v int32) *PostCustomizeRuleResponseBody {
	s.Code = &v
	return s
}

func (s *PostCustomizeRuleResponseBody) SetData(v *PostCustomizeRuleResponseBodyData) *PostCustomizeRuleResponseBody {
	s.Data = v
	return s
}

func (s *PostCustomizeRuleResponseBody) SetMessage(v string) *PostCustomizeRuleResponseBody {
	s.Message = &v
	return s
}

func (s *PostCustomizeRuleResponseBody) SetRequestId(v string) *PostCustomizeRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *PostCustomizeRuleResponseBody) SetSuccess(v bool) *PostCustomizeRuleResponseBody {
	s.Success = &v
	return s
}

type PostCustomizeRuleResponseBodyData struct {
	// The risk type.
	//
	// example:
	//
	// WEBSHELL
	AlertType *string `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
	// The internal code of the risk type.
	//
	// example:
	//
	// ${siem_rule_type_process_abnormal_command}
	AlertTypeMds *string `json:"AlertTypeMds,omitempty" xml:"AlertTypeMds,omitempty"`
	// The ID of the Alibaba Cloud account that is used to purchase the threat analysis feature.
	//
	// example:
	//
	// 127608589417****
	Aliuid *int64 `json:"Aliuid,omitempty" xml:"Aliuid,omitempty"`
	// 告警附加字段attck
	//
	// example:
	//
	// T1595.002 Vulnerability Scanning
	AttCk *string `json:"AttCk,omitempty" xml:"AttCk,omitempty"`
	// 自动化响应规则条件字段数据类型。
	//
	// example:
	//
	// varchar
	DataType *int32 `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// The extended information about event generation. If eventTransferType is set to allToSingle, the value of this parameter indicates the length and unit of the alert aggregation window. The HTML escape characters are reversed.
	//
	// example:
	//
	// {&quot;time&quot;:&quot;1&quot;,&quot;unit&quot;:&quot;MINUTE&quot;}
	EventTransferExt *string `json:"EventTransferExt,omitempty" xml:"EventTransferExt,omitempty"`
	// Indicates whether the system generates an event for the alert. Valid values:
	//
	// 	- 0: no
	//
	// 	- 1: yes
	//
	// example:
	//
	// 1
	EventTransferSwitch *int32 `json:"EventTransferSwitch,omitempty" xml:"EventTransferSwitch,omitempty"`
	// The event generation method. Valid values:
	//
	// 	- default: The default method is used.
	//
	// 	- singleToSingle: The system generates an event for each alert.
	//
	// 	- allToSingle: The system generates an event for alerts within a period of time.
	//
	// example:
	//
	// allToSingle
	EventTransferType *string `json:"EventTransferType,omitempty" xml:"EventTransferType,omitempty"`
	// The time when the custom rule was created.
	//
	// example:
	//
	// 2021-01-06 16:37:29
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the custom rule was last updated.
	//
	// example:
	//
	// 2021-01-06 16:37:29
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the custom rule.
	//
	// example:
	//
	// 123456789
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The log source of the rule.
	//
	// example:
	//
	// cloud_siem_aegis_sas_alert
	LogSource *string `json:"LogSource,omitempty" xml:"LogSource,omitempty"`
	// The internal code of the log source.
	//
	// example:
	//
	// ${sas.cloudsiem.prod.cloud_siem_aegis_sas_alert}
	LogSourceMds *string `json:"LogSourceMds,omitempty" xml:"LogSourceMds,omitempty"`
	// The log type of the rule.
	//
	// example:
	//
	// ALERT_ACTIVITY
	LogType *string `json:"LogType,omitempty" xml:"LogType,omitempty"`
	// The internal code of the log type.
	//
	// example:
	//
	// ${security_event_config.event_name.webshellName_clientav}
	LogTypeMds *string `json:"LogTypeMds,omitempty" xml:"LogTypeMds,omitempty"`
	// The window length of the rule. The HTML escape characters are reversed.
	//
	// example:
	//
	// {&quot;time&quot;:&quot;1&quot;,&quot;unit&quot;:&quot;HOUR&quot;}
	QueryCycle *string `json:"QueryCycle,omitempty" xml:"QueryCycle,omitempty"`
	// The query condition of the rule. The value is in the JSON format. The HTML escape characters are reversed.
	//
	// example:
	//
	// [[{&quot;not&quot;:false,&quot;left&quot;:&quot;alert_name&quot;,&quot;operator&quot;:&quot;=&quot;,&quot;right&quot;:&quot;WEBSHELL&quot;}]]
	RuleCondition *string `json:"RuleCondition,omitempty" xml:"RuleCondition,omitempty"`
	// The description of the rule.
	//
	// example:
	//
	// this rule is for waf scan
	RuleDesc *string `json:"RuleDesc,omitempty" xml:"RuleDesc,omitempty"`
	// The log aggregation field of the rule. The value is a JSON string. The HTML escape characters are reversed.
	//
	// example:
	//
	// [&quot;asset_id&quot;]
	RuleGroup *string `json:"RuleGroup,omitempty" xml:"RuleGroup,omitempty"`
	// The name of the rule.
	//
	// example:
	//
	// waf_scan
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The threshold configuration of the rule. The value is in the JSON format. The HTML escape characters are reversed.
	//
	// example:
	//
	// {&quot;aggregateFunction&quot;:&quot;count&quot;,&quot;aggregateFunctionName&quot;:&quot;count&quot;,&quot;field&quot;:&quot;activity_name&quot;,&quot;operator&quot;:&quot;&lt;=&quot;,&quot;value&quot;:1}
	RuleThreshold *string `json:"RuleThreshold,omitempty" xml:"RuleThreshold,omitempty"`
	// The type of the rule. Valid values:
	//
	// 	- predefine
	//
	// 	- customize
	//
	// example:
	//
	// customize
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// The rule status. Valid values:
	//
	// 	- 0: The rule is in the initial state.
	//
	// 	- 10: The simulation data is tested.
	//
	// 	- 15: The business data is being tested.
	//
	// 	- 20: The business data test ends.
	//
	// 	- 100: The rule takes effect.
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The risk level. Valid values:
	//
	// 	- serious: high
	//
	// 	- suspicious: medium
	//
	// 	- remind: low
	//
	// example:
	//
	// remind
	ThreatLevel *string `json:"ThreatLevel,omitempty" xml:"ThreatLevel,omitempty"`
}

func (s PostCustomizeRuleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s PostCustomizeRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *PostCustomizeRuleResponseBodyData) SetAlertType(v string) *PostCustomizeRuleResponseBodyData {
	s.AlertType = &v
	return s
}

func (s *PostCustomizeRuleResponseBodyData) SetAlertTypeMds(v string) *PostCustomizeRuleResponseBodyData {
	s.AlertTypeMds = &v
	return s
}

func (s *PostCustomizeRuleResponseBodyData) SetAliuid(v int64) *PostCustomizeRuleResponseBodyData {
	s.Aliuid = &v
	return s
}

func (s *PostCustomizeRuleResponseBodyData) SetAttCk(v string) *PostCustomizeRuleResponseBodyData {
	s.AttCk = &v
	return s
}

func (s *PostCustomizeRuleResponseBodyData) SetDataType(v int32) *PostCustomizeRuleResponseBodyData {
	s.DataType = &v
	return s
}

func (s *PostCustomizeRuleResponseBodyData) SetEventTransferExt(v string) *PostCustomizeRuleResponseBodyData {
	s.EventTransferExt = &v
	return s
}

func (s *PostCustomizeRuleResponseBodyData) SetEventTransferSwitch(v int32) *PostCustomizeRuleResponseBodyData {
	s.EventTransferSwitch = &v
	return s
}

func (s *PostCustomizeRuleResponseBodyData) SetEventTransferType(v string) *PostCustomizeRuleResponseBodyData {
	s.EventTransferType = &v
	return s
}

func (s *PostCustomizeRuleResponseBodyData) SetGmtCreate(v string) *PostCustomizeRuleResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *PostCustomizeRuleResponseBodyData) SetGmtModified(v string) *PostCustomizeRuleResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *PostCustomizeRuleResponseBodyData) SetId(v int64) *PostCustomizeRuleResponseBodyData {
	s.Id = &v
	return s
}

func (s *PostCustomizeRuleResponseBodyData) SetLogSource(v string) *PostCustomizeRuleResponseBodyData {
	s.LogSource = &v
	return s
}

func (s *PostCustomizeRuleResponseBodyData) SetLogSourceMds(v string) *PostCustomizeRuleResponseBodyData {
	s.LogSourceMds = &v
	return s
}

func (s *PostCustomizeRuleResponseBodyData) SetLogType(v string) *PostCustomizeRuleResponseBodyData {
	s.LogType = &v
	return s
}

func (s *PostCustomizeRuleResponseBodyData) SetLogTypeMds(v string) *PostCustomizeRuleResponseBodyData {
	s.LogTypeMds = &v
	return s
}

func (s *PostCustomizeRuleResponseBodyData) SetQueryCycle(v string) *PostCustomizeRuleResponseBodyData {
	s.QueryCycle = &v
	return s
}

func (s *PostCustomizeRuleResponseBodyData) SetRuleCondition(v string) *PostCustomizeRuleResponseBodyData {
	s.RuleCondition = &v
	return s
}

func (s *PostCustomizeRuleResponseBodyData) SetRuleDesc(v string) *PostCustomizeRuleResponseBodyData {
	s.RuleDesc = &v
	return s
}

func (s *PostCustomizeRuleResponseBodyData) SetRuleGroup(v string) *PostCustomizeRuleResponseBodyData {
	s.RuleGroup = &v
	return s
}

func (s *PostCustomizeRuleResponseBodyData) SetRuleName(v string) *PostCustomizeRuleResponseBodyData {
	s.RuleName = &v
	return s
}

func (s *PostCustomizeRuleResponseBodyData) SetRuleThreshold(v string) *PostCustomizeRuleResponseBodyData {
	s.RuleThreshold = &v
	return s
}

func (s *PostCustomizeRuleResponseBodyData) SetRuleType(v string) *PostCustomizeRuleResponseBodyData {
	s.RuleType = &v
	return s
}

func (s *PostCustomizeRuleResponseBodyData) SetStatus(v int32) *PostCustomizeRuleResponseBodyData {
	s.Status = &v
	return s
}

func (s *PostCustomizeRuleResponseBodyData) SetThreatLevel(v string) *PostCustomizeRuleResponseBodyData {
	s.ThreatLevel = &v
	return s
}

type PostCustomizeRuleResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PostCustomizeRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PostCustomizeRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s PostCustomizeRuleResponse) GoString() string {
	return s.String()
}

func (s *PostCustomizeRuleResponse) SetHeaders(v map[string]*string) *PostCustomizeRuleResponse {
	s.Headers = v
	return s
}

func (s *PostCustomizeRuleResponse) SetStatusCode(v int32) *PostCustomizeRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *PostCustomizeRuleResponse) SetBody(v *PostCustomizeRuleResponseBody) *PostCustomizeRuleResponse {
	s.Body = v
	return s
}

type PostCustomizeRuleTestRequest struct {
	// The ID of the rule.
	//
	// example:
	//
	// 123456789
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The data management center of the threat analysis feature. Specify this parameter based on the region in which your assets reside. Valid values:
	//
	// 	- cn-hangzhou: Your assets reside in regions inside China.
	//
	// 	- ap-southeast-1: Your assets reside in regions outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the account that you switch from the management account.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view. Valid values:
	//
	// - 0: the current Alibaba Cloud account
	//
	// - 1: the global account
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// The simulation data for the test. This parameter is available only when TestType is set to simulate.
	//
	// example:
	//
	// [{"key1":"value1","key2":"value2","key3":"value3","key4":"value4","key5":"value5"}]
	SimulatedData *string `json:"SimulatedData,omitempty" xml:"SimulatedData,omitempty"`
	// The test type. Valid values:
	//
	// 	- simulate: simulation data test
	//
	// 	- business: business data test
	//
	// example:
	//
	// simulate
	TestType *string `json:"TestType,omitempty" xml:"TestType,omitempty"`
}

func (s PostCustomizeRuleTestRequest) String() string {
	return tea.Prettify(s)
}

func (s PostCustomizeRuleTestRequest) GoString() string {
	return s.String()
}

func (s *PostCustomizeRuleTestRequest) SetId(v int64) *PostCustomizeRuleTestRequest {
	s.Id = &v
	return s
}

func (s *PostCustomizeRuleTestRequest) SetRegionId(v string) *PostCustomizeRuleTestRequest {
	s.RegionId = &v
	return s
}

func (s *PostCustomizeRuleTestRequest) SetRoleFor(v int64) *PostCustomizeRuleTestRequest {
	s.RoleFor = &v
	return s
}

func (s *PostCustomizeRuleTestRequest) SetRoleType(v int32) *PostCustomizeRuleTestRequest {
	s.RoleType = &v
	return s
}

func (s *PostCustomizeRuleTestRequest) SetSimulatedData(v string) *PostCustomizeRuleTestRequest {
	s.SimulatedData = &v
	return s
}

func (s *PostCustomizeRuleTestRequest) SetTestType(v string) *PostCustomizeRuleTestRequest {
	s.TestType = &v
	return s
}

type PostCustomizeRuleTestResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	//
	// example:
	//
	// 123456
	Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PostCustomizeRuleTestResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PostCustomizeRuleTestResponseBody) GoString() string {
	return s.String()
}

func (s *PostCustomizeRuleTestResponseBody) SetCode(v int32) *PostCustomizeRuleTestResponseBody {
	s.Code = &v
	return s
}

func (s *PostCustomizeRuleTestResponseBody) SetData(v interface{}) *PostCustomizeRuleTestResponseBody {
	s.Data = v
	return s
}

func (s *PostCustomizeRuleTestResponseBody) SetMessage(v string) *PostCustomizeRuleTestResponseBody {
	s.Message = &v
	return s
}

func (s *PostCustomizeRuleTestResponseBody) SetRequestId(v string) *PostCustomizeRuleTestResponseBody {
	s.RequestId = &v
	return s
}

func (s *PostCustomizeRuleTestResponseBody) SetSuccess(v bool) *PostCustomizeRuleTestResponseBody {
	s.Success = &v
	return s
}

type PostCustomizeRuleTestResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PostCustomizeRuleTestResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PostCustomizeRuleTestResponse) String() string {
	return tea.Prettify(s)
}

func (s PostCustomizeRuleTestResponse) GoString() string {
	return s.String()
}

func (s *PostCustomizeRuleTestResponse) SetHeaders(v map[string]*string) *PostCustomizeRuleTestResponse {
	s.Headers = v
	return s
}

func (s *PostCustomizeRuleTestResponse) SetStatusCode(v int32) *PostCustomizeRuleTestResponse {
	s.StatusCode = &v
	return s
}

func (s *PostCustomizeRuleTestResponse) SetBody(v *PostCustomizeRuleTestResponseBody) *PostCustomizeRuleTestResponse {
	s.Body = v
	return s
}

type PostEventDisposeAndWhiteruleListRequest struct {
	// The configuration of event handling. The value is a JSON object.
	//
	// example:
	//
	// [
	//
	//       {
	//
	//             "playbookName": "WafBlockIP",
	//
	//             "entityId": "104466118",
	//
	//             "scope": [
	//
	//                   "176618589410****"
	//
	//             ],
	//
	//             "startTime": 1604168946281,
	//
	//             "endTime": 1614168946281
	//
	//       },
	//
	//       {
	//
	//             "playbookName": "WafBlockIP",
	//
	//             "entityId": "104466118",
	//
	//             "scope": [
	//
	//                   {
	//
	//                         "instanceId": "waf-cn-n6w1oy1****",
	//
	//                         "domains": [
	//
	//                               "lmfip.wafqax.***"
	//
	//                         ]
	//
	//                   }
	//
	//             ],
	//
	//             "startTime": 1604168946281,
	//
	//             "endTime": 1614168946281
	//
	//       }
	//
	// ]
	EventDispose *string `json:"EventDispose,omitempty" xml:"EventDispose,omitempty"`
	// The UUID of the event.
	//
	// example:
	//
	// 85ea4241-798f-4684-a876-65d4f0c3****
	IncidentUuid *string `json:"IncidentUuid,omitempty" xml:"IncidentUuid,omitempty"`
	// The configuration of the alert recipient. The value is a JSON object.
	//
	// example:
	//
	// {
	//
	//       "messageTitle": "test",
	//
	//       "receiver": "xiaowang",
	//
	//       "channel": "message"
	//
	// }
	ReceiverInfo *string `json:"ReceiverInfo,omitempty" xml:"ReceiverInfo,omitempty"`
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
	// The remarks of the event.
	//
	// example:
	//
	// dealed
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The ID of the account that you switch from the management account.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view. Valid values:
	//
	// - 0: the current Alibaba Cloud account
	//
	// - 1: the global account
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// The status of the event. Valid values:
	//
	// 	- 0: unhandled
	//
	// 	- 1: handing
	//
	// 	- 5: handling failed
	//
	// 	- 10: handled
	//
	// example:
	//
	// 0
	Status      *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	ThreatLevel *string `json:"ThreatLevel,omitempty" xml:"ThreatLevel,omitempty"`
}

func (s PostEventDisposeAndWhiteruleListRequest) String() string {
	return tea.Prettify(s)
}

func (s PostEventDisposeAndWhiteruleListRequest) GoString() string {
	return s.String()
}

func (s *PostEventDisposeAndWhiteruleListRequest) SetEventDispose(v string) *PostEventDisposeAndWhiteruleListRequest {
	s.EventDispose = &v
	return s
}

func (s *PostEventDisposeAndWhiteruleListRequest) SetIncidentUuid(v string) *PostEventDisposeAndWhiteruleListRequest {
	s.IncidentUuid = &v
	return s
}

func (s *PostEventDisposeAndWhiteruleListRequest) SetReceiverInfo(v string) *PostEventDisposeAndWhiteruleListRequest {
	s.ReceiverInfo = &v
	return s
}

func (s *PostEventDisposeAndWhiteruleListRequest) SetRegionId(v string) *PostEventDisposeAndWhiteruleListRequest {
	s.RegionId = &v
	return s
}

func (s *PostEventDisposeAndWhiteruleListRequest) SetRemark(v string) *PostEventDisposeAndWhiteruleListRequest {
	s.Remark = &v
	return s
}

func (s *PostEventDisposeAndWhiteruleListRequest) SetRoleFor(v int64) *PostEventDisposeAndWhiteruleListRequest {
	s.RoleFor = &v
	return s
}

func (s *PostEventDisposeAndWhiteruleListRequest) SetRoleType(v int32) *PostEventDisposeAndWhiteruleListRequest {
	s.RoleType = &v
	return s
}

func (s *PostEventDisposeAndWhiteruleListRequest) SetStatus(v int32) *PostEventDisposeAndWhiteruleListRequest {
	s.Status = &v
	return s
}

func (s *PostEventDisposeAndWhiteruleListRequest) SetThreatLevel(v string) *PostEventDisposeAndWhiteruleListRequest {
	s.ThreatLevel = &v
	return s
}

type PostEventDisposeAndWhiteruleListResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	//
	// example:
	//
	// 123456
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PostEventDisposeAndWhiteruleListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PostEventDisposeAndWhiteruleListResponseBody) GoString() string {
	return s.String()
}

func (s *PostEventDisposeAndWhiteruleListResponseBody) SetCode(v int32) *PostEventDisposeAndWhiteruleListResponseBody {
	s.Code = &v
	return s
}

func (s *PostEventDisposeAndWhiteruleListResponseBody) SetData(v string) *PostEventDisposeAndWhiteruleListResponseBody {
	s.Data = &v
	return s
}

func (s *PostEventDisposeAndWhiteruleListResponseBody) SetMessage(v string) *PostEventDisposeAndWhiteruleListResponseBody {
	s.Message = &v
	return s
}

func (s *PostEventDisposeAndWhiteruleListResponseBody) SetRequestId(v string) *PostEventDisposeAndWhiteruleListResponseBody {
	s.RequestId = &v
	return s
}

func (s *PostEventDisposeAndWhiteruleListResponseBody) SetSuccess(v bool) *PostEventDisposeAndWhiteruleListResponseBody {
	s.Success = &v
	return s
}

type PostEventDisposeAndWhiteruleListResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PostEventDisposeAndWhiteruleListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PostEventDisposeAndWhiteruleListResponse) String() string {
	return tea.Prettify(s)
}

func (s PostEventDisposeAndWhiteruleListResponse) GoString() string {
	return s.String()
}

func (s *PostEventDisposeAndWhiteruleListResponse) SetHeaders(v map[string]*string) *PostEventDisposeAndWhiteruleListResponse {
	s.Headers = v
	return s
}

func (s *PostEventDisposeAndWhiteruleListResponse) SetStatusCode(v int32) *PostEventDisposeAndWhiteruleListResponse {
	s.StatusCode = &v
	return s
}

func (s *PostEventDisposeAndWhiteruleListResponse) SetBody(v *PostEventDisposeAndWhiteruleListResponseBody) *PostEventDisposeAndWhiteruleListResponse {
	s.Body = v
	return s
}

type PostEventWhiteruleListRequest struct {
	// The UUID of the event.
	//
	// example:
	//
	// 85ea4241-798f-4684-a876-65d4f0c3****
	IncidentUuid *string `json:"IncidentUuid,omitempty" xml:"IncidentUuid,omitempty"`
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
	// The ID of the account that you switch from the management account.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view. Valid values:
	//
	// - 0: the current Alibaba Cloud account
	//
	// - 1: the global account
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// The alert whitelist rule. The value is a JSON object.
	//
	// This parameter is required.
	//
	// example:
	//
	// [
	//
	//       {
	//
	//             "alertName": "webshell",
	//
	//             "alertNameId": "webshell",
	//
	//             "alertType": "command",
	//
	//             "alertTypeId": "command",
	//
	//             "expression": {
	//
	//                   "status": 1,
	//
	//                   "conditions": [
	//
	//                         {
	//
	//                               "isNot": false,
	//
	//                               "left": {
	//
	//                                     "value": "file_path"
	//
	//                               },
	//
	//                               "operator": "gt",
	//
	//                               "right": {
	//
	//                                     "value": "cp"
	//
	//                               }
	//
	//                         }
	//
	//                   ]
	//
	//             }
	//
	//       }
	//
	// ]
	WhiteruleList *string `json:"WhiteruleList,omitempty" xml:"WhiteruleList,omitempty"`
}

func (s PostEventWhiteruleListRequest) String() string {
	return tea.Prettify(s)
}

func (s PostEventWhiteruleListRequest) GoString() string {
	return s.String()
}

func (s *PostEventWhiteruleListRequest) SetIncidentUuid(v string) *PostEventWhiteruleListRequest {
	s.IncidentUuid = &v
	return s
}

func (s *PostEventWhiteruleListRequest) SetRegionId(v string) *PostEventWhiteruleListRequest {
	s.RegionId = &v
	return s
}

func (s *PostEventWhiteruleListRequest) SetRoleFor(v int64) *PostEventWhiteruleListRequest {
	s.RoleFor = &v
	return s
}

func (s *PostEventWhiteruleListRequest) SetRoleType(v int32) *PostEventWhiteruleListRequest {
	s.RoleType = &v
	return s
}

func (s *PostEventWhiteruleListRequest) SetWhiteruleList(v string) *PostEventWhiteruleListRequest {
	s.WhiteruleList = &v
	return s
}

type PostEventWhiteruleListResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	//
	// example:
	//
	// 123456
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PostEventWhiteruleListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PostEventWhiteruleListResponseBody) GoString() string {
	return s.String()
}

func (s *PostEventWhiteruleListResponseBody) SetCode(v int32) *PostEventWhiteruleListResponseBody {
	s.Code = &v
	return s
}

func (s *PostEventWhiteruleListResponseBody) SetData(v string) *PostEventWhiteruleListResponseBody {
	s.Data = &v
	return s
}

func (s *PostEventWhiteruleListResponseBody) SetMessage(v string) *PostEventWhiteruleListResponseBody {
	s.Message = &v
	return s
}

func (s *PostEventWhiteruleListResponseBody) SetRequestId(v string) *PostEventWhiteruleListResponseBody {
	s.RequestId = &v
	return s
}

func (s *PostEventWhiteruleListResponseBody) SetSuccess(v bool) *PostEventWhiteruleListResponseBody {
	s.Success = &v
	return s
}

type PostEventWhiteruleListResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PostEventWhiteruleListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PostEventWhiteruleListResponse) String() string {
	return tea.Prettify(s)
}

func (s PostEventWhiteruleListResponse) GoString() string {
	return s.String()
}

func (s *PostEventWhiteruleListResponse) SetHeaders(v map[string]*string) *PostEventWhiteruleListResponse {
	s.Headers = v
	return s
}

func (s *PostEventWhiteruleListResponse) SetStatusCode(v int32) *PostEventWhiteruleListResponse {
	s.StatusCode = &v
	return s
}

func (s *PostEventWhiteruleListResponse) SetBody(v *PostEventWhiteruleListResponseBody) *PostEventWhiteruleListResponse {
	s.Body = v
	return s
}

type PostFinishCustomizeRuleTestRequest struct {
	// The ID of the rule.
	//
	// example:
	//
	// 123456789
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
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
	// The ID of the account that you switch from the management account.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view. Valid values:
	//
	// - 0: the current Alibaba Cloud account
	//
	// - 1: the global account
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s PostFinishCustomizeRuleTestRequest) String() string {
	return tea.Prettify(s)
}

func (s PostFinishCustomizeRuleTestRequest) GoString() string {
	return s.String()
}

func (s *PostFinishCustomizeRuleTestRequest) SetId(v int64) *PostFinishCustomizeRuleTestRequest {
	s.Id = &v
	return s
}

func (s *PostFinishCustomizeRuleTestRequest) SetRegionId(v string) *PostFinishCustomizeRuleTestRequest {
	s.RegionId = &v
	return s
}

func (s *PostFinishCustomizeRuleTestRequest) SetRoleFor(v int64) *PostFinishCustomizeRuleTestRequest {
	s.RoleFor = &v
	return s
}

func (s *PostFinishCustomizeRuleTestRequest) SetRoleType(v int32) *PostFinishCustomizeRuleTestRequest {
	s.RoleType = &v
	return s
}

type PostFinishCustomizeRuleTestResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	//
	// example:
	//
	// 123456
	Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PostFinishCustomizeRuleTestResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PostFinishCustomizeRuleTestResponseBody) GoString() string {
	return s.String()
}

func (s *PostFinishCustomizeRuleTestResponseBody) SetCode(v int32) *PostFinishCustomizeRuleTestResponseBody {
	s.Code = &v
	return s
}

func (s *PostFinishCustomizeRuleTestResponseBody) SetData(v interface{}) *PostFinishCustomizeRuleTestResponseBody {
	s.Data = v
	return s
}

func (s *PostFinishCustomizeRuleTestResponseBody) SetMessage(v string) *PostFinishCustomizeRuleTestResponseBody {
	s.Message = &v
	return s
}

func (s *PostFinishCustomizeRuleTestResponseBody) SetRequestId(v string) *PostFinishCustomizeRuleTestResponseBody {
	s.RequestId = &v
	return s
}

func (s *PostFinishCustomizeRuleTestResponseBody) SetSuccess(v bool) *PostFinishCustomizeRuleTestResponseBody {
	s.Success = &v
	return s
}

type PostFinishCustomizeRuleTestResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PostFinishCustomizeRuleTestResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PostFinishCustomizeRuleTestResponse) String() string {
	return tea.Prettify(s)
}

func (s PostFinishCustomizeRuleTestResponse) GoString() string {
	return s.String()
}

func (s *PostFinishCustomizeRuleTestResponse) SetHeaders(v map[string]*string) *PostFinishCustomizeRuleTestResponse {
	s.Headers = v
	return s
}

func (s *PostFinishCustomizeRuleTestResponse) SetStatusCode(v int32) *PostFinishCustomizeRuleTestResponse {
	s.StatusCode = &v
	return s
}

func (s *PostFinishCustomizeRuleTestResponse) SetBody(v *PostFinishCustomizeRuleTestResponseBody) *PostFinishCustomizeRuleTestResponse {
	s.Body = v
	return s
}

type PostRuleStatusChangeRequest struct {
	// The rule IDs. The value is a JSON array.
	//
	// example:
	//
	// [123,345]
	Ids *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	// Specifies whether to enable the rule. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	InUse *bool `json:"InUse,omitempty" xml:"InUse,omitempty"`
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
	// The ID of the account that you switch from the management account.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view. Valid values:
	//
	// - 0: the current Alibaba Cloud account
	//
	// - 1: the global account
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// The type of the rule. Valid values:
	//
	// 	- predefine
	//
	// 	- customize
	//
	// example:
	//
	// customize
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
}

func (s PostRuleStatusChangeRequest) String() string {
	return tea.Prettify(s)
}

func (s PostRuleStatusChangeRequest) GoString() string {
	return s.String()
}

func (s *PostRuleStatusChangeRequest) SetIds(v string) *PostRuleStatusChangeRequest {
	s.Ids = &v
	return s
}

func (s *PostRuleStatusChangeRequest) SetInUse(v bool) *PostRuleStatusChangeRequest {
	s.InUse = &v
	return s
}

func (s *PostRuleStatusChangeRequest) SetRegionId(v string) *PostRuleStatusChangeRequest {
	s.RegionId = &v
	return s
}

func (s *PostRuleStatusChangeRequest) SetRoleFor(v int64) *PostRuleStatusChangeRequest {
	s.RoleFor = &v
	return s
}

func (s *PostRuleStatusChangeRequest) SetRoleType(v int32) *PostRuleStatusChangeRequest {
	s.RoleType = &v
	return s
}

func (s *PostRuleStatusChangeRequest) SetRuleType(v string) *PostRuleStatusChangeRequest {
	s.RuleType = &v
	return s
}

type PostRuleStatusChangeResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	//
	// example:
	//
	// 123456
	Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PostRuleStatusChangeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PostRuleStatusChangeResponseBody) GoString() string {
	return s.String()
}

func (s *PostRuleStatusChangeResponseBody) SetCode(v int32) *PostRuleStatusChangeResponseBody {
	s.Code = &v
	return s
}

func (s *PostRuleStatusChangeResponseBody) SetData(v interface{}) *PostRuleStatusChangeResponseBody {
	s.Data = v
	return s
}

func (s *PostRuleStatusChangeResponseBody) SetMessage(v string) *PostRuleStatusChangeResponseBody {
	s.Message = &v
	return s
}

func (s *PostRuleStatusChangeResponseBody) SetRequestId(v string) *PostRuleStatusChangeResponseBody {
	s.RequestId = &v
	return s
}

func (s *PostRuleStatusChangeResponseBody) SetSuccess(v bool) *PostRuleStatusChangeResponseBody {
	s.Success = &v
	return s
}

type PostRuleStatusChangeResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PostRuleStatusChangeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PostRuleStatusChangeResponse) String() string {
	return tea.Prettify(s)
}

func (s PostRuleStatusChangeResponse) GoString() string {
	return s.String()
}

func (s *PostRuleStatusChangeResponse) SetHeaders(v map[string]*string) *PostRuleStatusChangeResponse {
	s.Headers = v
	return s
}

func (s *PostRuleStatusChangeResponse) SetStatusCode(v int32) *PostRuleStatusChangeResponse {
	s.StatusCode = &v
	return s
}

func (s *PostRuleStatusChangeResponse) SetBody(v *PostRuleStatusChangeResponseBody) *PostRuleStatusChangeResponse {
	s.Body = v
	return s
}

type RestoreCapacityRequest struct {
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
	// The ID of the account that you switch from the management account.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view. Valid values:
	//
	// - 0: the current Alibaba Cloud account
	//
	// - 1: the global account
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s RestoreCapacityRequest) String() string {
	return tea.Prettify(s)
}

func (s RestoreCapacityRequest) GoString() string {
	return s.String()
}

func (s *RestoreCapacityRequest) SetRegionId(v string) *RestoreCapacityRequest {
	s.RegionId = &v
	return s
}

func (s *RestoreCapacityRequest) SetRoleFor(v int64) *RestoreCapacityRequest {
	s.RoleFor = &v
	return s
}

func (s *RestoreCapacityRequest) SetRoleType(v int32) *RestoreCapacityRequest {
	s.RoleType = &v
	return s
}

type RestoreCapacityResponseBody struct {
	// Indicates whether the release command has been sent. Valid values:
	//
	// 	- true: The command has been sent and the storage space is being released.
	//
	// 	- false: The command failed to be sent.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6276D891-58D4-55B2-87B9-74D413F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RestoreCapacityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RestoreCapacityResponseBody) GoString() string {
	return s.String()
}

func (s *RestoreCapacityResponseBody) SetData(v bool) *RestoreCapacityResponseBody {
	s.Data = &v
	return s
}

func (s *RestoreCapacityResponseBody) SetRequestId(v string) *RestoreCapacityResponseBody {
	s.RequestId = &v
	return s
}

type RestoreCapacityResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RestoreCapacityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RestoreCapacityResponse) String() string {
	return tea.Prettify(s)
}

func (s RestoreCapacityResponse) GoString() string {
	return s.String()
}

func (s *RestoreCapacityResponse) SetHeaders(v map[string]*string) *RestoreCapacityResponse {
	s.Headers = v
	return s
}

func (s *RestoreCapacityResponse) SetStatusCode(v int32) *RestoreCapacityResponse {
	s.StatusCode = &v
	return s
}

func (s *RestoreCapacityResponse) SetBody(v *RestoreCapacityResponseBody) *RestoreCapacityResponse {
	s.Body = v
	return s
}

type SetStorageRequest struct {
	// The storage region of logs.
	//
	// If the data management center is **cn-hangzhou**, the default value of **Region*	- is cn-shanghai, which specifies the China (Shanghai) region. If the data management center is **ap-southeast-1**, the default value of **Region*	- is ap-southeast-1, which specifies the Singapore region.
	//
	// The region for log storage cannot be changed. To change the region, contact the technical support of threat analysis.
	//
	// example:
	//
	// cn-shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The data management center of the threat analysis feature. Specify this parameter based on the region where your assets reside. Valid values:
	//
	// 	- cn-hangzhou: Your assets reside in regions in China.
	//
	// 	- ap-southeast-1: Your assets reside in regions outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the account that you switch from the management account.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view. Valid values:
	//
	// - 0: the current Alibaba Cloud account
	//
	// - 1: the global account
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// The storage duration of logs. Default value: 180. Minimum value: 30. Maximum value: 3000. Unit: days.
	//
	// This parameter is required.
	//
	// example:
	//
	// 180
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
}

func (s SetStorageRequest) String() string {
	return tea.Prettify(s)
}

func (s SetStorageRequest) GoString() string {
	return s.String()
}

func (s *SetStorageRequest) SetRegion(v string) *SetStorageRequest {
	s.Region = &v
	return s
}

func (s *SetStorageRequest) SetRegionId(v string) *SetStorageRequest {
	s.RegionId = &v
	return s
}

func (s *SetStorageRequest) SetRoleFor(v int64) *SetStorageRequest {
	s.RoleFor = &v
	return s
}

func (s *SetStorageRequest) SetRoleType(v int32) *SetStorageRequest {
	s.RoleType = &v
	return s
}

func (s *SetStorageRequest) SetTtl(v int32) *SetStorageRequest {
	s.Ttl = &v
	return s
}

type SetStorageResponseBody struct {
	// Indicates whether the settings are saved. Valid values:
	//
	// 	- true:
	//
	// 	- false:
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6276D891-58D4-55B2-87B9-74D413F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetStorageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetStorageResponseBody) GoString() string {
	return s.String()
}

func (s *SetStorageResponseBody) SetData(v bool) *SetStorageResponseBody {
	s.Data = &v
	return s
}

func (s *SetStorageResponseBody) SetRequestId(v string) *SetStorageResponseBody {
	s.RequestId = &v
	return s
}

type SetStorageResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetStorageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetStorageResponse) String() string {
	return tea.Prettify(s)
}

func (s SetStorageResponse) GoString() string {
	return s.String()
}

func (s *SetStorageResponse) SetHeaders(v map[string]*string) *SetStorageResponse {
	s.Headers = v
	return s
}

func (s *SetStorageResponse) SetStatusCode(v int32) *SetStorageResponse {
	s.StatusCode = &v
	return s
}

func (s *SetStorageResponse) SetBody(v *SetStorageResponseBody) *SetStorageResponse {
	s.Body = v
	return s
}

type SubmitImportLogTasksRequest struct {
	// The accounts that you want to add. The value is a JSON array. Valid values:
	//
	// 	- AccountId: the IDs of the accounts.
	//
	// 	- Imported: specifies whether to add the accounts. Valid values:
	//
	//     	- 0: no
	//
	//     	- 1: yes
	//
	// example:
	//
	// [{"AccountId":"123123","Imported":1}]
	Accounts *string `json:"Accounts,omitempty" xml:"Accounts,omitempty"`
	// Specifies whether to automatically add the account for which the logging feature is configured. Valid values:
	//
	// 	- 1: yes
	//
	// 	- 0: no
	//
	// example:
	//
	// ["cloud_siem_qcloud_cfw_alert_log"]
	AutoImported *int32 `json:"AutoImported,omitempty" xml:"AutoImported,omitempty"`
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
	// The logs that you want to collect. The value is a JSON array.
	//
	// example:
	//
	// ["cloud_siem_qcloud_cfw_alert_log"]
	LogCodes *string `json:"LogCodes,omitempty" xml:"LogCodes,omitempty"`
	// The code of the service.
	//
	// This parameter is required.
	//
	// example:
	//
	// qcloud_waf
	ProdCode *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
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
	// The ID of the account that you switch from the management account.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view. Valid values:
	//
	// - 0: the current Alibaba Cloud account
	//
	// - 1: the global account
	//
	// example:
	//
	// 0
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s SubmitImportLogTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitImportLogTasksRequest) GoString() string {
	return s.String()
}

func (s *SubmitImportLogTasksRequest) SetAccounts(v string) *SubmitImportLogTasksRequest {
	s.Accounts = &v
	return s
}

func (s *SubmitImportLogTasksRequest) SetAutoImported(v int32) *SubmitImportLogTasksRequest {
	s.AutoImported = &v
	return s
}

func (s *SubmitImportLogTasksRequest) SetCloudCode(v string) *SubmitImportLogTasksRequest {
	s.CloudCode = &v
	return s
}

func (s *SubmitImportLogTasksRequest) SetLogCodes(v string) *SubmitImportLogTasksRequest {
	s.LogCodes = &v
	return s
}

func (s *SubmitImportLogTasksRequest) SetProdCode(v string) *SubmitImportLogTasksRequest {
	s.ProdCode = &v
	return s
}

func (s *SubmitImportLogTasksRequest) SetRegionId(v string) *SubmitImportLogTasksRequest {
	s.RegionId = &v
	return s
}

func (s *SubmitImportLogTasksRequest) SetRoleFor(v int64) *SubmitImportLogTasksRequest {
	s.RoleFor = &v
	return s
}

func (s *SubmitImportLogTasksRequest) SetRoleType(v int32) *SubmitImportLogTasksRequest {
	s.RoleType = &v
	return s
}

type SubmitImportLogTasksResponseBody struct {
	// The data returned.
	Data *SubmitImportLogTasksResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitImportLogTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitImportLogTasksResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitImportLogTasksResponseBody) SetData(v *SubmitImportLogTasksResponseBodyData) *SubmitImportLogTasksResponseBody {
	s.Data = v
	return s
}

func (s *SubmitImportLogTasksResponseBody) SetRequestId(v string) *SubmitImportLogTasksResponseBody {
	s.RequestId = &v
	return s
}

type SubmitImportLogTasksResponseBodyData struct {
	// The number of log collection tasks that are submitted.
	//
	// example:
	//
	// 10
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s SubmitImportLogTasksResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SubmitImportLogTasksResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitImportLogTasksResponseBodyData) SetCount(v int32) *SubmitImportLogTasksResponseBodyData {
	s.Count = &v
	return s
}

type SubmitImportLogTasksResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitImportLogTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitImportLogTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitImportLogTasksResponse) GoString() string {
	return s.String()
}

func (s *SubmitImportLogTasksResponse) SetHeaders(v map[string]*string) *SubmitImportLogTasksResponse {
	s.Headers = v
	return s
}

func (s *SubmitImportLogTasksResponse) SetStatusCode(v int32) *SubmitImportLogTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitImportLogTasksResponse) SetBody(v *SubmitImportLogTasksResponseBody) *SubmitImportLogTasksResponse {
	s.Body = v
	return s
}

type UpdateAutomateResponseConfigStatusRequest struct {
	// The IDs of the automatic response rules. The value is a JSON array.
	//
	// example:
	//
	// [123,345]
	Ids *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	// Specifies whether the rule is enabled. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	InUse *bool `json:"InUse,omitempty" xml:"InUse,omitempty"`
	// The data management center of the threat analysis feature. Specify this parameter based on the region in which your assets reside. Valid values:
	//
	// 	- cn-hangzhou: Your assets reside in regions inside China.
	//
	// 	- ap-southeast-1: Your assets reside in regions outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the account that you switch from the management account.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view. Valid values:
	//
	// - 0: the current Alibaba Cloud account
	//
	// - 1: the global account
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s UpdateAutomateResponseConfigStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAutomateResponseConfigStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateAutomateResponseConfigStatusRequest) SetIds(v string) *UpdateAutomateResponseConfigStatusRequest {
	s.Ids = &v
	return s
}

func (s *UpdateAutomateResponseConfigStatusRequest) SetInUse(v bool) *UpdateAutomateResponseConfigStatusRequest {
	s.InUse = &v
	return s
}

func (s *UpdateAutomateResponseConfigStatusRequest) SetRegionId(v string) *UpdateAutomateResponseConfigStatusRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateAutomateResponseConfigStatusRequest) SetRoleFor(v int64) *UpdateAutomateResponseConfigStatusRequest {
	s.RoleFor = &v
	return s
}

func (s *UpdateAutomateResponseConfigStatusRequest) SetRoleType(v int32) *UpdateAutomateResponseConfigStatusRequest {
	s.RoleType = &v
	return s
}

type UpdateAutomateResponseConfigStatusResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	//
	// example:
	//
	// 123456
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateAutomateResponseConfigStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAutomateResponseConfigStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAutomateResponseConfigStatusResponseBody) SetCode(v int32) *UpdateAutomateResponseConfigStatusResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateAutomateResponseConfigStatusResponseBody) SetData(v string) *UpdateAutomateResponseConfigStatusResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateAutomateResponseConfigStatusResponseBody) SetMessage(v string) *UpdateAutomateResponseConfigStatusResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAutomateResponseConfigStatusResponseBody) SetRequestId(v string) *UpdateAutomateResponseConfigStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAutomateResponseConfigStatusResponseBody) SetSuccess(v bool) *UpdateAutomateResponseConfigStatusResponseBody {
	s.Success = &v
	return s
}

type UpdateAutomateResponseConfigStatusResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAutomateResponseConfigStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAutomateResponseConfigStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAutomateResponseConfigStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateAutomateResponseConfigStatusResponse) SetHeaders(v map[string]*string) *UpdateAutomateResponseConfigStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateAutomateResponseConfigStatusResponse) SetStatusCode(v int32) *UpdateAutomateResponseConfigStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAutomateResponseConfigStatusResponse) SetBody(v *UpdateAutomateResponseConfigStatusResponseBody) *UpdateAutomateResponseConfigStatusResponse {
	s.Body = v
	return s
}

type UpdateWhiteRuleListRequest struct {
	// The alert whitelist rule. The value is a JSON object.
	//
	// This parameter is required.
	//
	// example:
	//
	// [
	//
	//       {
	//
	//             "alertName": "webshell",
	//
	//             "alertNameId": "webshell",
	//
	//             "alertType": "command",
	//
	//             "alertTypeId": "command",
	//
	//             "expression": {
	//
	//                   "status": 1,
	//
	//                   "conditions": [
	//
	//                         {
	//
	//                               "isNot": false,
	//
	//                               "left": {
	//
	//                                     "value": "file_path"
	//
	//                               },
	//
	//                               "operator": "gt",
	//
	//                               "right": {
	//
	//                                     "value": "cp"
	//
	//                               }
	//
	//                         }
	//
	//                   ]
	//
	//             }
	//
	//       }
	//
	// ]
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// The UUID of the event.
	//
	// example:
	//
	// 85ea4241-798f-4684-a876-65d4f0c3****
	IncidentUuid *string `json:"IncidentUuid,omitempty" xml:"IncidentUuid,omitempty"`
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
	RoleFor  *int64  `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	RoleType *int32  `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// The unique ID of the whitelist rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456789
	WhiteRuleId *int64 `json:"WhiteRuleId,omitempty" xml:"WhiteRuleId,omitempty"`
}

func (s UpdateWhiteRuleListRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateWhiteRuleListRequest) GoString() string {
	return s.String()
}

func (s *UpdateWhiteRuleListRequest) SetExpression(v string) *UpdateWhiteRuleListRequest {
	s.Expression = &v
	return s
}

func (s *UpdateWhiteRuleListRequest) SetIncidentUuid(v string) *UpdateWhiteRuleListRequest {
	s.IncidentUuid = &v
	return s
}

func (s *UpdateWhiteRuleListRequest) SetRegionId(v string) *UpdateWhiteRuleListRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateWhiteRuleListRequest) SetRoleFor(v int64) *UpdateWhiteRuleListRequest {
	s.RoleFor = &v
	return s
}

func (s *UpdateWhiteRuleListRequest) SetRoleType(v int32) *UpdateWhiteRuleListRequest {
	s.RoleType = &v
	return s
}

func (s *UpdateWhiteRuleListRequest) SetWhiteRuleId(v int64) *UpdateWhiteRuleListRequest {
	s.WhiteRuleId = &v
	return s
}

type UpdateWhiteRuleListResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	//
	// example:
	//
	// 123456
	Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateWhiteRuleListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateWhiteRuleListResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWhiteRuleListResponseBody) SetCode(v int32) *UpdateWhiteRuleListResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateWhiteRuleListResponseBody) SetData(v interface{}) *UpdateWhiteRuleListResponseBody {
	s.Data = v
	return s
}

func (s *UpdateWhiteRuleListResponseBody) SetMessage(v string) *UpdateWhiteRuleListResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateWhiteRuleListResponseBody) SetRequestId(v string) *UpdateWhiteRuleListResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateWhiteRuleListResponseBody) SetSuccess(v bool) *UpdateWhiteRuleListResponseBody {
	s.Success = &v
	return s
}

type UpdateWhiteRuleListResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateWhiteRuleListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateWhiteRuleListResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateWhiteRuleListResponse) GoString() string {
	return s.String()
}

func (s *UpdateWhiteRuleListResponse) SetHeaders(v map[string]*string) *UpdateWhiteRuleListResponse {
	s.Headers = v
	return s
}

func (s *UpdateWhiteRuleListResponse) SetStatusCode(v int32) *UpdateWhiteRuleListResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWhiteRuleListResponse) SetBody(v *UpdateWhiteRuleListResponseBody) *UpdateWhiteRuleListResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("cloud-siem"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds a data source to a cloud account that is added to the threat analysis feature.
//
// @param request - AddDataSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddDataSourceResponse
func (client *Client) AddDataSourceWithOptions(request *AddDataSourceRequest, runtime *util.RuntimeOptions) (_result *AddDataSourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		body["AccountId"] = request.AccountId
	}

	if !tea.BoolValue(util.IsUnset(request.CloudCode)) {
		body["CloudCode"] = request.CloudCode
	}

	if !tea.BoolValue(util.IsUnset(request.DataSourceInstanceName)) {
		body["DataSourceInstanceName"] = request.DataSourceInstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.DataSourceInstanceParams)) {
		body["DataSourceInstanceParams"] = request.DataSourceInstanceParams
	}

	if !tea.BoolValue(util.IsUnset(request.DataSourceInstanceRemark)) {
		body["DataSourceInstanceRemark"] = request.DataSourceInstanceRemark
	}

	if !tea.BoolValue(util.IsUnset(request.DataSourceType)) {
		body["DataSourceType"] = request.DataSourceType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddDataSource"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddDataSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a data source to a cloud account that is added to the threat analysis feature.
//
// @param request - AddDataSourceRequest
//
// @return AddDataSourceResponse
func (client *Client) AddDataSource(request *AddDataSourceRequest) (_result *AddDataSourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddDataSourceResponse{}
	_body, _err := client.AddDataSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds logs of a cloud account to the threat analysis feature.
//
// @param request - AddDataSourceLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddDataSourceLogResponse
func (client *Client) AddDataSourceLogWithOptions(request *AddDataSourceLogRequest, runtime *util.RuntimeOptions) (_result *AddDataSourceLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		body["AccountId"] = request.AccountId
	}

	if !tea.BoolValue(util.IsUnset(request.CloudCode)) {
		body["CloudCode"] = request.CloudCode
	}

	if !tea.BoolValue(util.IsUnset(request.DataSourceInstanceId)) {
		body["DataSourceInstanceId"] = request.DataSourceInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.DataSourceInstanceLogs)) {
		body["DataSourceInstanceLogs"] = request.DataSourceInstanceLogs
	}

	if !tea.BoolValue(util.IsUnset(request.LogCode)) {
		body["LogCode"] = request.LogCode
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddDataSourceLog"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddDataSourceLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds logs of a cloud account to the threat analysis feature.
//
// @param request - AddDataSourceLogRequest
//
// @return AddDataSourceLogResponse
func (client *Client) AddDataSourceLog(request *AddDataSourceLogRequest) (_result *AddDataSourceLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddDataSourceLogResponse{}
	_body, _err := client.AddDataSourceLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds the logs of a cloud service within a cloud account to the threat analysis feature for alert and event anslysis.
//
// @param request - AddUserSourceLogConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddUserSourceLogConfigResponse
func (client *Client) AddUserSourceLogConfigWithOptions(request *AddUserSourceLogConfigRequest, runtime *util.RuntimeOptions) (_result *AddUserSourceLogConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Deleted)) {
		body["Deleted"] = request.Deleted
	}

	if !tea.BoolValue(util.IsUnset(request.DisPlayLine)) {
		body["DisPlayLine"] = request.DisPlayLine
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceLogCode)) {
		body["SourceLogCode"] = request.SourceLogCode
	}

	if !tea.BoolValue(util.IsUnset(request.SourceLogInfo)) {
		body["SourceLogInfo"] = request.SourceLogInfo
	}

	if !tea.BoolValue(util.IsUnset(request.SourceProdCode)) {
		body["SourceProdCode"] = request.SourceProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.SubUserId)) {
		body["SubUserId"] = request.SubUserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddUserSourceLogConfig"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddUserSourceLogConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds the logs of a cloud service within a cloud account to the threat analysis feature for alert and event anslysis.
//
// @param request - AddUserSourceLogConfigRequest
//
// @return AddUserSourceLogConfigResponse
func (client *Client) AddUserSourceLogConfig(request *AddUserSourceLogConfigRequest) (_result *AddUserSourceLogConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddUserSourceLogConfigResponse{}
	_body, _err := client.AddUserSourceLogConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds a third-party cloud account that is displayed on the Multi-cloud assets tab of the Feature Settings page to the threat analysis feature.
//
// @param request - BindAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindAccountResponse
func (client *Client) BindAccountWithOptions(request *BindAccountRequest, runtime *util.RuntimeOptions) (_result *BindAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessId)) {
		body["AccessId"] = request.AccessId
	}

	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		body["AccountId"] = request.AccountId
	}

	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		body["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.CloudCode)) {
		body["CloudCode"] = request.CloudCode
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleFor)) {
		body["RoleFor"] = request.RoleFor
	}

	if !tea.BoolValue(util.IsUnset(request.RoleType)) {
		body["RoleType"] = request.RoleType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BindAccount"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BindAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a third-party cloud account that is displayed on the Multi-cloud assets tab of the Feature Settings page to the threat analysis feature.
//
// @param request - BindAccountRequest
//
// @return BindAccountResponse
func (client *Client) BindAccount(request *BindAccountRequest) (_result *BindAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BindAccountResponse{}
	_body, _err := client.BindAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables the log delivery feature for a cloud service.
//
// @param request - CloseDeliveryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloseDeliveryResponse
func (client *Client) CloseDeliveryWithOptions(request *CloseDeliveryRequest, runtime *util.RuntimeOptions) (_result *CloseDeliveryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LogCode)) {
		body["LogCode"] = request.LogCode
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		body["ProductCode"] = request.ProductCode
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleFor)) {
		body["RoleFor"] = request.RoleFor
	}

	if !tea.BoolValue(util.IsUnset(request.RoleType)) {
		body["RoleType"] = request.RoleType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CloseDelivery"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CloseDeliveryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables the log delivery feature for a cloud service.
//
// @param request - CloseDeliveryRequest
//
// @return CloseDeliveryResponse
func (client *Client) CloseDelivery(request *CloseDeliveryRequest) (_result *CloseDeliveryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CloseDeliveryResponse{}
	_body, _err := client.CloseDeliveryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes the automated response rule with a specified ID.
//
// @param request - DeleteAutomateResponseConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAutomateResponseConfigResponse
func (client *Client) DeleteAutomateResponseConfigWithOptions(request *DeleteAutomateResponseConfigRequest, runtime *util.RuntimeOptions) (_result *DeleteAutomateResponseConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleFor)) {
		body["RoleFor"] = request.RoleFor
	}

	if !tea.BoolValue(util.IsUnset(request.RoleType)) {
		body["RoleType"] = request.RoleType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAutomateResponseConfig"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAutomateResponseConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the automated response rule with a specified ID.
//
// @param request - DeleteAutomateResponseConfigRequest
//
// @return DeleteAutomateResponseConfigResponse
func (client *Client) DeleteAutomateResponseConfig(request *DeleteAutomateResponseConfigRequest) (_result *DeleteAutomateResponseConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAutomateResponseConfigResponse{}
	_body, _err := client.DeleteAutomateResponseConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes a third-party cloud account that is added to the threat analysis feature by using its AccessKey ID. You can add another cloud account based on your business requirements.
//
// @param request - DeleteBindAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteBindAccountResponse
func (client *Client) DeleteBindAccountWithOptions(request *DeleteBindAccountRequest, runtime *util.RuntimeOptions) (_result *DeleteBindAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessId)) {
		body["AccessId"] = request.AccessId
	}

	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		body["AccountId"] = request.AccountId
	}

	if !tea.BoolValue(util.IsUnset(request.BindId)) {
		body["BindId"] = request.BindId
	}

	if !tea.BoolValue(util.IsUnset(request.CloudCode)) {
		body["CloudCode"] = request.CloudCode
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleFor)) {
		body["RoleFor"] = request.RoleFor
	}

	if !tea.BoolValue(util.IsUnset(request.RoleType)) {
		body["RoleType"] = request.RoleType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteBindAccount"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteBindAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes a third-party cloud account that is added to the threat analysis feature by using its AccessKey ID. You can add another cloud account based on your business requirements.
//
// @param request - DeleteBindAccountRequest
//
// @return DeleteBindAccountResponse
func (client *Client) DeleteBindAccount(request *DeleteBindAccountRequest) (_result *DeleteBindAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteBindAccountResponse{}
	_body, _err := client.DeleteBindAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a rule by rule ID.
//
// @param request - DeleteCustomizeRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCustomizeRuleResponse
func (client *Client) DeleteCustomizeRuleWithOptions(request *DeleteCustomizeRuleRequest, runtime *util.RuntimeOptions) (_result *DeleteCustomizeRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleFor)) {
		body["RoleFor"] = request.RoleFor
	}

	if !tea.BoolValue(util.IsUnset(request.RoleType)) {
		body["RoleType"] = request.RoleType
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		body["RuleId"] = request.RuleId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteCustomizeRule"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteCustomizeRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a rule by rule ID.
//
// @param request - DeleteCustomizeRuleRequest
//
// @return DeleteCustomizeRuleResponse
func (client *Client) DeleteCustomizeRule(request *DeleteCustomizeRuleRequest) (_result *DeleteCustomizeRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteCustomizeRuleResponse{}
	_body, _err := client.DeleteCustomizeRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes a data source that is no longer required.
//
// @param request - DeleteDataSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataSourceResponse
func (client *Client) DeleteDataSourceWithOptions(request *DeleteDataSourceRequest, runtime *util.RuntimeOptions) (_result *DeleteDataSourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		body["AccountId"] = request.AccountId
	}

	if !tea.BoolValue(util.IsUnset(request.CloudCode)) {
		body["CloudCode"] = request.CloudCode
	}

	if !tea.BoolValue(util.IsUnset(request.DataSourceInstanceId)) {
		body["DataSourceInstanceId"] = request.DataSourceInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDataSource"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDataSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes a data source that is no longer required.
//
// @param request - DeleteDataSourceRequest
//
// @return DeleteDataSourceResponse
func (client *Client) DeleteDataSource(request *DeleteDataSourceRequest) (_result *DeleteDataSourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDataSourceResponse{}
	_body, _err := client.DeleteDataSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes a log.
//
// @param request - DeleteDataSourceLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataSourceLogResponse
func (client *Client) DeleteDataSourceLogWithOptions(request *DeleteDataSourceLogRequest, runtime *util.RuntimeOptions) (_result *DeleteDataSourceLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		body["AccountId"] = request.AccountId
	}

	if !tea.BoolValue(util.IsUnset(request.CloudCode)) {
		body["CloudCode"] = request.CloudCode
	}

	if !tea.BoolValue(util.IsUnset(request.DataSourceInstanceId)) {
		body["DataSourceInstanceId"] = request.DataSourceInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.LogInstanceId)) {
		body["LogInstanceId"] = request.LogInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDataSourceLog"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDataSourceLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes a log.
//
// @param request - DeleteDataSourceLogRequest
//
// @return DeleteDataSourceLogResponse
func (client *Client) DeleteDataSourceLog(request *DeleteDataSourceLogRequest) (_result *DeleteDataSourceLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDataSourceLogResponse{}
	_body, _err := client.DeleteDataSourceLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an alert whitelist rule with a specified ID.
//
// @param request - DeleteWhiteRuleListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWhiteRuleListResponse
func (client *Client) DeleteWhiteRuleListWithOptions(request *DeleteWhiteRuleListRequest, runtime *util.RuntimeOptions) (_result *DeleteWhiteRuleListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleFor)) {
		body["RoleFor"] = request.RoleFor
	}

	if !tea.BoolValue(util.IsUnset(request.RoleType)) {
		body["RoleType"] = request.RoleType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteWhiteRuleList"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteWhiteRuleListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an alert whitelist rule with a specified ID.
//
// @param request - DeleteWhiteRuleListRequest
//
// @return DeleteWhiteRuleListResponse
func (client *Client) DeleteWhiteRuleList(request *DeleteWhiteRuleListRequest) (_result *DeleteWhiteRuleListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteWhiteRuleListResponse{}
	_body, _err := client.DeleteWhiteRuleListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the aggregate functions that are supported for a custom rule.
//
// @param request - DescribeAggregateFunctionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAggregateFunctionResponse
func (client *Client) DescribeAggregateFunctionWithOptions(request *DescribeAggregateFunctionRequest, runtime *util.RuntimeOptions) (_result *DescribeAggregateFunctionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleFor)) {
		body["RoleFor"] = request.RoleFor
	}

	if !tea.BoolValue(util.IsUnset(request.RoleType)) {
		body["RoleType"] = request.RoleType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAggregateFunction"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAggregateFunctionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the aggregate functions that are supported for a custom rule.
//
// @param request - DescribeAggregateFunctionRequest
//
// @return DescribeAggregateFunctionResponse
func (client *Client) DescribeAggregateFunction(request *DescribeAggregateFunctionRequest) (_result *DescribeAggregateFunctionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAggregateFunctionResponse{}
	_body, _err := client.DescribeAggregateFunctionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the scenarios in which an alert needs to be added to the whitelist.
//
// @param request - DescribeAlertSceneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAlertSceneResponse
func (client *Client) DescribeAlertSceneWithOptions(request *DescribeAlertSceneRequest, runtime *util.RuntimeOptions) (_result *DescribeAlertSceneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleFor)) {
		body["RoleFor"] = request.RoleFor
	}

	if !tea.BoolValue(util.IsUnset(request.RoleType)) {
		body["RoleType"] = request.RoleType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAlertScene"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAlertSceneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the scenarios in which an alert needs to be added to the whitelist.
//
// @param request - DescribeAlertSceneRequest
//
// @return DescribeAlertSceneResponse
func (client *Client) DescribeAlertScene(request *DescribeAlertSceneRequest) (_result *DescribeAlertSceneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAlertSceneResponse{}
	_body, _err := client.DescribeAlertSceneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the scenarios and objects that can be added to an alert whitelist rule.
//
// @param request - DescribeAlertSceneByEventRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAlertSceneByEventResponse
func (client *Client) DescribeAlertSceneByEventWithOptions(request *DescribeAlertSceneByEventRequest, runtime *util.RuntimeOptions) (_result *DescribeAlertSceneByEventResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IncidentUuid)) {
		body["IncidentUuid"] = request.IncidentUuid
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleFor)) {
		body["RoleFor"] = request.RoleFor
	}

	if !tea.BoolValue(util.IsUnset(request.RoleType)) {
		body["RoleType"] = request.RoleType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAlertSceneByEvent"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAlertSceneByEventResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the scenarios and objects that can be added to an alert whitelist rule.
//
// @param request - DescribeAlertSceneByEventRequest
//
// @return DescribeAlertSceneByEventResponse
func (client *Client) DescribeAlertSceneByEvent(request *DescribeAlertSceneByEventRequest) (_result *DescribeAlertSceneByEventResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAlertSceneByEventResponse{}
	_body, _err := client.DescribeAlertSceneByEventWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries alert data sources.
//
// @param request - DescribeAlertSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAlertSourceResponse
func (client *Client) DescribeAlertSourceWithOptions(request *DescribeAlertSourceRequest, runtime *util.RuntimeOptions) (_result *DescribeAlertSourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Level)) {
		body["Level"] = request.Level
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleFor)) {
		body["RoleFor"] = request.RoleFor
	}

	if !tea.BoolValue(util.IsUnset(request.RoleType)) {
		body["RoleType"] = request.RoleType
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAlertSource"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAlertSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries alert data sources.
//
// @param request - DescribeAlertSourceRequest
//
// @return DescribeAlertSourceResponse
func (client *Client) DescribeAlertSource(request *DescribeAlertSourceRequest) (_result *DescribeAlertSourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAlertSourceResponse{}
	_body, _err := client.DescribeAlertSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the data sources of the alert that is associated with an event.
//
// @param request - DescribeAlertSourceWithEventRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAlertSourceWithEventResponse
func (client *Client) DescribeAlertSourceWithEventWithOptions(request *DescribeAlertSourceWithEventRequest, runtime *util.RuntimeOptions) (_result *DescribeAlertSourceWithEventResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IncidentUuid)) {
		body["IncidentUuid"] = request.IncidentUuid
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleFor)) {
		body["RoleFor"] = request.RoleFor
	}

	if !tea.BoolValue(util.IsUnset(request.RoleType)) {
		body["RoleType"] = request.RoleType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAlertSourceWithEvent"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAlertSourceWithEventResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the data sources of the alert that is associated with an event.
//
// @param request - DescribeAlertSourceWithEventRequest
//
// @return DescribeAlertSourceWithEventResponse
func (client *Client) DescribeAlertSourceWithEvent(request *DescribeAlertSourceWithEventRequest) (_result *DescribeAlertSourceWithEventResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAlertSourceWithEventResponse{}
	_body, _err := client.DescribeAlertSourceWithEventWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the threat types that you can select when you create a custom rule.
//
// @param request - DescribeAlertTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAlertTypeResponse
func (client *Client) DescribeAlertTypeWithOptions(request *DescribeAlertTypeRequest, runtime *util.RuntimeOptions) (_result *DescribeAlertTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleFor)) {
		body["RoleFor"] = request.RoleFor
	}

	if !tea.BoolValue(util.IsUnset(request.RoleType)) {
		body["RoleType"] = request.RoleType
	}

	if !tea.BoolValue(util.IsUnset(request.RuleType)) {
		body["RuleType"] = request.RuleType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAlertType"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAlertTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the threat types that you can select when you create a custom rule.
//
// @param request - DescribeAlertTypeRequest
//
// @return DescribeAlertTypeResponse
func (client *Client) DescribeAlertType(request *DescribeAlertTypeRequest) (_result *DescribeAlertTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAlertTypeResponse{}
	_body, _err := client.DescribeAlertTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries alerts within your account.
//
// @param request - DescribeAlertsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAlertsResponse
func (client *Client) DescribeAlertsWithOptions(request *DescribeAlertsRequest, runtime *util.RuntimeOptions) (_result *DescribeAlertsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlertName)) {
		body["AlertName"] = request.AlertName
	}

	if !tea.BoolValue(util.IsUnset(request.AlertTitle)) {
		body["AlertTitle"] = request.AlertTitle
	}

	if !tea.BoolValue(util.IsUnset(request.AlertType)) {
		body["AlertType"] = request.AlertType
	}

	if !tea.BoolValue(util.IsUnset(request.AlertUuid)) {
		body["AlertUuid"] = request.AlertUuid
	}

	if !tea.BoolValue(util.IsUnset(request.AssetId)) {
		body["AssetId"] = request.AssetId
	}

	if !tea.BoolValue(util.IsUnset(request.AssetName)) {
		body["AssetName"] = request.AssetName
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.EntityId)) {
		body["EntityId"] = request.EntityId
	}

	if !tea.BoolValue(util.IsUnset(request.EntityName)) {
		body["EntityName"] = request.EntityName
	}

	if !tea.BoolValue(util.IsUnset(request.IsDefend)) {
		body["IsDefend"] = request.IsDefend
	}

	if !tea.BoolValue(util.IsUnset(request.LabelType)) {
		body["LabelType"] = request.LabelType
	}

	if !tea.BoolValue(util.IsUnset(request.Level)) {
		body["Level"] = request.Level
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleFor)) {
		body["RoleFor"] = request.RoleFor
	}

	if !tea.BoolValue(util.IsUnset(request.RoleType)) {
		body["RoleType"] = request.RoleType
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		body["Source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.SubUserId)) {
		body["SubUserId"] = request.SubUserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAlerts"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAlertsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries alerts within your account.
//
// @param request - DescribeAlertsRequest
//
// @return DescribeAlertsResponse
func (client *Client) DescribeAlerts(request *DescribeAlertsRequest) (_result *DescribeAlertsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAlertsResponse{}
	_body, _err := client.DescribeAlertsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the number of alerts of different severities.
//
// @param request - DescribeAlertsCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAlertsCountResponse
func (client *Client) DescribeAlertsCountWithOptions(request *DescribeAlertsCountRequest, runtime *util.RuntimeOptions) (_result *DescribeAlertsCountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.QueryType)) {
		body["QueryType"] = request.QueryType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleFor)) {
		body["RoleFor"] = request.RoleFor
	}

	if !tea.BoolValue(util.IsUnset(request.RoleType)) {
		body["RoleType"] = request.RoleType
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAlertsCount"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAlertsCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of alerts of different severities.
//
// @param request - DescribeAlertsCountRequest
//
// @return DescribeAlertsCountResponse
func (client *Client) DescribeAlertsCount(request *DescribeAlertsCountRequest) (_result *DescribeAlertsCountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAlertsCountResponse{}
	_body, _err := client.DescribeAlertsCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the alerts that are associated with an entity.
//
// @param request - DescribeAlertsWithEntityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAlertsWithEntityResponse
func (client *Client) DescribeAlertsWithEntityWithOptions(request *DescribeAlertsWithEntityRequest, runtime *util.RuntimeOptions) (_result *DescribeAlertsWithEntityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.EntityId)) {
		body["EntityId"] = request.EntityId
	}

	if !tea.BoolValue(util.IsUnset(request.EntityUuid)) {
		body["EntityUuid"] = request.EntityUuid
	}

	if !tea.BoolValue(util.IsUnset(request.IncidentUuid)) {
		body["IncidentUuid"] = request.IncidentUuid
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleFor)) {
		body["RoleFor"] = request.RoleFor
	}

	if !tea.BoolValue(util.IsUnset(request.RoleType)) {
		body["RoleType"] = request.RoleType
	}

	if !tea.BoolValue(util.IsUnset(request.SophonTaskId)) {
		body["SophonTaskId"] = request.SophonTaskId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAlertsWithEntity"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAlertsWithEntityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the alerts that are associated with an entity.
//
// @param request - DescribeAlertsWithEntityRequest
//
// @return DescribeAlertsWithEntityResponse
func (client *Client) DescribeAlertsWithEntity(request *DescribeAlertsWithEntityRequest) (_result *DescribeAlertsWithEntityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAlertsWithEntityResponse{}
	_body, _err := client.DescribeAlertsWithEntityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the alerts that are associated with an event.
//
// @param request - DescribeAlertsWithEventRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAlertsWithEventResponse
func (client *Client) DescribeAlertsWithEventWithOptions(request *DescribeAlertsWithEventRequest, runtime *util.RuntimeOptions) (_result *DescribeAlertsWithEventResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlertName)) {
		body["AlertName"] = request.AlertName
	}

	if !tea.BoolValue(util.IsUnset(request.AlertTitle)) {
		body["AlertTitle"] = request.AlertTitle
	}

	if !tea.BoolValue(util.IsUnset(request.AlertType)) {
		body["AlertType"] = request.AlertType
	}

	if !tea.BoolValue(util.IsUnset(request.AssetId)) {
		body["AssetId"] = request.AssetId
	}

	if !tea.BoolValue(util.IsUnset(request.AssetName)) {
		body["AssetName"] = request.AssetName
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.EntityId)) {
		body["EntityId"] = request.EntityId
	}

	if !tea.BoolValue(util.IsUnset(request.EntityName)) {
		body["EntityName"] = request.EntityName
	}

	if !tea.BoolValue(util.IsUnset(request.IncidentUuid)) {
		body["IncidentUuid"] = request.IncidentUuid
	}

	if !tea.BoolValue(util.IsUnset(request.IsDefend)) {
		body["IsDefend"] = request.IsDefend
	}

	if !tea.BoolValue(util.IsUnset(request.Level)) {
		body["Level"] = request.Level
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleFor)) {
		body["RoleFor"] = request.RoleFor
	}

	if !tea.BoolValue(util.IsUnset(request.RoleType)) {
		body["RoleType"] = request.RoleType
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		body["Source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.SubUserId)) {
		body["SubUserId"] = request.SubUserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAlertsWithEvent"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAlertsWithEventResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the alerts that are associated with an event.
//
// @param request - DescribeAlertsWithEventRequest
//
// @return DescribeAlertsWithEventResponse
func (client *Client) DescribeAlertsWithEvent(request *DescribeAlertsWithEventRequest) (_result *DescribeAlertsWithEventResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAlertsWithEventResponse{}
	_body, _err := client.DescribeAlertsWithEventWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Checks whether the security information and event management (SIEM) system is granted the required permissions to access other cloud resources within your Alibaba Cloud account and whether the AliyunServiceRoleForSasCloudSiem service-linked role is created.
//
// @param request - DescribeAuthRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAuthResponse
func (client *Client) DescribeAuthWithOptions(request *DescribeAuthRequest, runtime *util.RuntimeOptions) (_result *DescribeAuthResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAuth"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAuthResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks whether the security information and event management (SIEM) system is granted the required permissions to access other cloud resources within your Alibaba Cloud account and whether the AliyunServiceRoleForSasCloudSiem service-linked role is created.
//
// @param request - DescribeAuthRequest
//
// @return DescribeAuthResponse
func (client *Client) DescribeAuth(request *DescribeAuthRequest) (_result *DescribeAuthResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAuthResponse{}
	_body, _err := client.DescribeAuthWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the number of automated response rules.
//
// @param request - DescribeAutomateResponseConfigCounterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAutomateResponseConfigCounterResponse
func (client *Client) DescribeAutomateResponseConfigCounterWithOptions(request *DescribeAutomateResponseConfigCounterRequest, runtime *util.RuntimeOptions) (_result *DescribeAutomateResponseConfigCounterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleFor)) {
		body["RoleFor"] = request.RoleFor
	}

	if !tea.BoolValue(util.IsUnset(request.RoleType)) {
		body["RoleType"] = request.RoleType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAutomateResponseConfigCounter"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAutomateResponseConfigCounterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of automated response rules.
//
// @param request - DescribeAutomateResponseConfigCounterRequest
//
// @return DescribeAutomateResponseConfigCounterResponse
func (client *Client) DescribeAutomateResponseConfigCounter(request *DescribeAutomateResponseConfigCounterRequest) (_result *DescribeAutomateResponseConfigCounterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAutomateResponseConfigCounterResponse{}
	_body, _err := client.DescribeAutomateResponseConfigCounterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the configurable fields and operators of an automated response rule.
//
// @param request - DescribeAutomateResponseConfigFeatureRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAutomateResponseConfigFeatureResponse
func (client *Client) DescribeAutomateResponseConfigFeatureWithOptions(request *DescribeAutomateResponseConfigFeatureRequest, runtime *util.RuntimeOptions) (_result *DescribeAutomateResponseConfigFeatureResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoResponseType)) {
		body["AutoResponseType"] = request.AutoResponseType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleFor)) {
		body["RoleFor"] = request.RoleFor
	}

	if !tea.BoolValue(util.IsUnset(request.RoleType)) {
		body["RoleType"] = request.RoleType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAutomateResponseConfigFeature"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAutomateResponseConfigFeatureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configurable fields and operators of an automated response rule.
//
// @param request - DescribeAutomateResponseConfigFeatureRequest
//
// @return DescribeAutomateResponseConfigFeatureResponse
func (client *Client) DescribeAutomateResponseConfigFeature(request *DescribeAutomateResponseConfigFeatureRequest) (_result *DescribeAutomateResponseConfigFeatureResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAutomateResponseConfigFeatureResponse{}
	_body, _err := client.DescribeAutomateResponseConfigFeatureWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries user-defined playbooks.
//
// @param request - DescribeAutomateResponseConfigPlayBooksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAutomateResponseConfigPlayBooksResponse
func (client *Client) DescribeAutomateResponseConfigPlayBooksWithOptions(request *DescribeAutomateResponseConfigPlayBooksRequest, runtime *util.RuntimeOptions) (_result *DescribeAutomateResponseConfigPlayBooksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoResponseType)) {
		body["AutoResponseType"] = request.AutoResponseType
	}

	if !tea.BoolValue(util.IsUnset(request.EntityType)) {
		body["EntityType"] = request.EntityType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleFor)) {
		body["RoleFor"] = request.RoleFor
	}

	if !tea.BoolValue(util.IsUnset(request.RoleType)) {
		body["RoleType"] = request.RoleType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAutomateResponseConfigPlayBooks"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAutomateResponseConfigPlayBooksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries user-defined playbooks.
//
// @param request - DescribeAutomateResponseConfigPlayBooksRequest
//
// @return DescribeAutomateResponseConfigPlayBooksResponse
func (client *Client) DescribeAutomateResponseConfigPlayBooks(request *DescribeAutomateResponseConfigPlayBooksRequest) (_result *DescribeAutomateResponseConfigPlayBooksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAutomateResponseConfigPlayBooksResponse{}
	_body, _err := client.DescribeAutomateResponseConfigPlayBooksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the assets that are associated with an event.
//
// @param request - DescribeCloudSiemAssetsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCloudSiemAssetsResponse
func (client *Client) DescribeCloudSiemAssetsWithOptions(request *DescribeCloudSiemAssetsRequest, runtime *util.RuntimeOptions) (_result *DescribeCloudSiemAssetsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AssetName)) {
		body["AssetName"] = request.AssetName
	}

	if !tea.BoolValue(util.IsUnset(request.AssetType)) {
		body["AssetType"] = request.AssetType
	}

	if !tea.BoolValue(util.IsUnset(request.AssetUuid)) {
		body["AssetUuid"] = request.AssetUuid
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.IncidentUuid)) {
		body["IncidentUuid"] = request.IncidentUuid
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleFor)) {
		body["RoleFor"] = request.RoleFor
	}

	if !tea.BoolValue(util.IsUnset(request.RoleType)) {
		body["RoleType"] = request.RoleType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCloudSiemAssets"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCloudSiemAssetsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the assets that are associated with an event.
//
// @param request - DescribeCloudSiemAssetsRequest
//
// @return DescribeCloudSiemAssetsResponse
func (client *Client) DescribeCloudSiemAssets(request *DescribeCloudSiemAssetsRequest) (_result *DescribeCloudSiemAssetsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCloudSiemAssetsResponse{}
	_body, _err := client.DescribeCloudSiemAssetsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the number of assets that are associated with an event by asset type.
//
// @param request - DescribeCloudSiemAssetsCounterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCloudSiemAssetsCounterResponse
func (client *Client) DescribeCloudSiemAssetsCounterWithOptions(request *DescribeCloudSiemAssetsCounterRequest, runtime *util.RuntimeOptions) (_result *DescribeCloudSiemAssetsCounterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IncidentUuid)) {
		body["IncidentUuid"] = request.IncidentUuid
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleFor)) {
		body["RoleFor"] = request.RoleFor
	}

	if !tea.BoolValue(util.IsUnset(request.RoleType)) {
		body["RoleType"] = request.RoleType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCloudSiemAssetsCounter"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCloudSiemAssetsCounterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of assets that are associated with an event by asset type.
//
// @param request - DescribeCloudSiemAssetsCounterRequest
//
// @return DescribeCloudSiemAssetsCounterResponse
func (client *Client) DescribeCloudSiemAssetsCounter(request *DescribeCloudSiemAssetsCounterRequest) (_result *DescribeCloudSiemAssetsCounterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCloudSiemAssetsCounterResponse{}
	_body, _err := client.DescribeCloudSiemAssetsCounterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of an event.
//
// @param request - DescribeCloudSiemEventDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCloudSiemEventDetailResponse
func (client *Client) DescribeCloudSiemEventDetailWithOptions(request *DescribeCloudSiemEventDetailRequest, runtime *util.RuntimeOptions) (_result *DescribeCloudSiemEventDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IncidentUuid)) {
		body["IncidentUuid"] = request.IncidentUuid
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleFor)) {
		body["RoleFor"] = request.RoleFor
	}

	if !tea.BoolValue(util.IsUnset(request.RoleType)) {
		body["RoleType"] = request.RoleType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCloudSiemEventDetail"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCloudSiemEventDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an event.
//
// @param request - DescribeCloudSiemEventDetailRequest
//
// @return DescribeCloudSiemEventDetailResponse
func (client *Client) DescribeCloudSiemEventDetail(request *DescribeCloudSiemEventDetailRequest) (_result *DescribeCloudSiemEventDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCloudSiemEventDetailResponse{}
	_body, _err := client.DescribeCloudSiemEventDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries events in SIEM.
//
// @param request - DescribeCloudSiemEventsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCloudSiemEventsResponse
func (client *Client) DescribeCloudSiemEventsWithOptions(request *DescribeCloudSiemEventsRequest, runtime *util.RuntimeOptions) (_result *DescribeCloudSiemEventsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AssetId)) {
		body["AssetId"] = request.AssetId
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.EntityUuid)) {
		body["EntityUuid"] = request.EntityUuid
	}

	if !tea.BoolValue(util.IsUnset(request.EventName)) {
		body["EventName"] = request.EventName
	}

	if !tea.BoolValue(util.IsUnset(request.IncidentUuid)) {
		body["IncidentUuid"] = request.IncidentUuid
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		body["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.OrderField)) {
		body["OrderField"] = request.OrderField
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleFor)) {
		body["RoleFor"] = request.RoleFor
	}

	if !tea.BoolValue(util.IsUnset(request.RoleType)) {
		body["RoleType"] = request.RoleType
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.ThreadLevel)) {
		body["ThreadLevel"] = request.ThreadLevel
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCloudSiemEvents"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCloudSiemEventsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries events in SIEM.
//
// @param request - DescribeCloudSiemEventsRequest
//
// @return DescribeCloudSiemEventsResponse
func (client *Client) DescribeCloudSiemEvents(request *DescribeCloudSiemEventsRequest) (_result *DescribeCloudSiemEventsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCloudSiemEventsResponse{}
	_body, _err := client.DescribeCloudSiemEventsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the number of custom rules.
//
// @param request - DescribeCustomizeRuleCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCustomizeRuleCountResponse
func (client *Client) DescribeCustomizeRuleCountWithOptions(request *DescribeCustomizeRuleCountRequest, runtime *util.RuntimeOptions) (_result *DescribeCustomizeRuleCountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleFor)) {
		body["RoleFor"] = request.RoleFor
	}

	if !tea.BoolValue(util.IsUnset(request.RoleType)) {
		body["RoleType"] = request.RoleType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCustomizeRuleCount"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCustomizeRuleCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of custom rules.
//
// @param request - DescribeCustomizeRuleCountRequest
//
// @return DescribeCustomizeRuleCountResponse
func (client *Client) DescribeCustomizeRuleCount(request *DescribeCustomizeRuleCountRequest) (_result *DescribeCustomizeRuleCountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCustomizeRuleCountResponse{}
	_body, _err := client.DescribeCustomizeRuleCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the historical simulation data that is used in a simulation test scenario.
//
// @param request - DescribeCustomizeRuleTestRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCustomizeRuleTestResponse
func (client *Client) DescribeCustomizeRuleTestWithOptions(request *DescribeCustomizeRuleTestRequest, runtime *util.RuntimeOptions) (_result *DescribeCustomizeRuleTestResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleFor)) {
		body["RoleFor"] = request.RoleFor
	}

	if !tea.BoolValue(util.IsUnset(request.RoleType)) {
		body["RoleType"] = request.RoleType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCustomizeRuleTest"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCustomizeRuleTestResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the historical simulation data that is used in a simulation test scenario.
//
// @param request - DescribeCustomizeRuleTestRequest
//
// @return DescribeCustomizeRuleTestResponse
func (client *Client) DescribeCustomizeRuleTest(request *DescribeCustomizeRuleTestRequest) (_result *DescribeCustomizeRuleTestResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCustomizeRuleTestResponse{}
	_body, _err := client.DescribeCustomizeRuleTestWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the chart that displays the test results of business data for a custom rule.
//
// @param request - DescribeCustomizeRuleTestHistogramRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCustomizeRuleTestHistogramResponse
func (client *Client) DescribeCustomizeRuleTestHistogramWithOptions(request *DescribeCustomizeRuleTestHistogramRequest, runtime *util.RuntimeOptions) (_result *DescribeCustomizeRuleTestHistogramResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleFor)) {
		body["RoleFor"] = request.RoleFor
	}

	if !tea.BoolValue(util.IsUnset(request.RoleType)) {
		body["RoleType"] = request.RoleType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCustomizeRuleTestHistogram"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCustomizeRuleTestHistogramResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the chart that displays the test results of business data for a custom rule.
//
// @param request - DescribeCustomizeRuleTestHistogramRequest
//
// @return DescribeCustomizeRuleTestHistogramResponse
func (client *Client) DescribeCustomizeRuleTestHistogram(request *DescribeCustomizeRuleTestHistogramRequest) (_result *DescribeCustomizeRuleTestHistogramResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCustomizeRuleTestHistogramResponse{}
	_body, _err := client.DescribeCustomizeRuleTestHistogramWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a data source.
//
// @param request - DescribeDataSourceInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDataSourceInstanceResponse
func (client *Client) DescribeDataSourceInstanceWithOptions(request *DescribeDataSourceInstanceRequest, runtime *util.RuntimeOptions) (_result *DescribeDataSourceInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		body["AccountId"] = request.AccountId
	}

	if !tea.BoolValue(util.IsUnset(request.CloudCode)) {
		body["CloudCode"] = request.CloudCode
	}

	if !tea.BoolValue(util.IsUnset(request.DataSourceInstanceId)) {
		body["DataSourceInstanceId"] = request.DataSourceInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDataSourceInstance"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDataSourceInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a data source.
//
// @param request - DescribeDataSourceInstanceRequest
//
// @return DescribeDataSourceInstanceResponse
func (client *Client) DescribeDataSourceInstance(request *DescribeDataSourceInstanceRequest) (_result *DescribeDataSourceInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDataSourceInstanceResponse{}
	_body, _err := client.DescribeDataSourceInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the parameters of a data source.
//
// @param request - DescribeDataSourceParametersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDataSourceParametersResponse
func (client *Client) DescribeDataSourceParametersWithOptions(request *DescribeDataSourceParametersRequest, runtime *util.RuntimeOptions) (_result *DescribeDataSourceParametersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CloudCode)) {
		body["CloudCode"] = request.CloudCode
	}

	if !tea.BoolValue(util.IsUnset(request.DataSourceType)) {
		body["DataSourceType"] = request.DataSourceType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDataSourceParameters"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDataSourceParametersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the parameters of a data source.
//
// @param request - DescribeDataSourceParametersRequest
//
// @return DescribeDataSourceParametersResponse
func (client *Client) DescribeDataSourceParameters(request *DescribeDataSourceParametersRequest) (_result *DescribeDataSourceParametersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDataSourceParametersResponse{}
	_body, _err := client.DescribeDataSourceParametersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of entities and playbooks that need to be handled.
//
// @param request - DescribeDisposeAndPlaybookRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDisposeAndPlaybookResponse
func (client *Client) DescribeDisposeAndPlaybookWithOptions(request *DescribeDisposeAndPlaybookRequest, runtime *util.RuntimeOptions) (_result *DescribeDisposeAndPlaybookResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.EntityType)) {
		body["EntityType"] = request.EntityType
	}

	if !tea.BoolValue(util.IsUnset(request.EntityUuid)) {
		body["EntityUuid"] = request.EntityUuid
	}

	if !tea.BoolValue(util.IsUnset(request.IncidentUuid)) {
		body["IncidentUuid"] = request.IncidentUuid
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleFor)) {
		body["RoleFor"] = request.RoleFor
	}

	if !tea.BoolValue(util.IsUnset(request.RoleType)) {
		body["RoleType"] = request.RoleType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDisposeAndPlaybook"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDisposeAndPlaybookResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of entities and playbooks that need to be handled.
//
// @param request - DescribeDisposeAndPlaybookRequest
//
// @return DescribeDisposeAndPlaybookResponse
func (client *Client) DescribeDisposeAndPlaybook(request *DescribeDisposeAndPlaybookRequest) (_result *DescribeDisposeAndPlaybookResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDisposeAndPlaybookResponse{}
	_body, _err := client.DescribeDisposeAndPlaybookWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of playbooks that are used by a handling policy.
//
// @param request - DescribeDisposeStrategyPlaybookRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDisposeStrategyPlaybookResponse
func (client *Client) DescribeDisposeStrategyPlaybookWithOptions(request *DescribeDisposeStrategyPlaybookRequest, runtime *util.RuntimeOptions) (_result *DescribeDisposeStrategyPlaybookResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleFor)) {
		body["RoleFor"] = request.RoleFor
	}

	if !tea.BoolValue(util.IsUnset(request.RoleType)) {
		body["RoleType"] = request.RoleType
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDisposeStrategyPlaybook"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDisposeStrategyPlaybookResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of playbooks that are used by a handling policy.
//
// @param request - DescribeDisposeStrategyPlaybookRequest
//
// @return DescribeDisposeStrategyPlaybookResponse
func (client *Client) DescribeDisposeStrategyPlaybook(request *DescribeDisposeStrategyPlaybookRequest) (_result *DescribeDisposeStrategyPlaybookResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDisposeStrategyPlaybookResponse{}
	_body, _err := client.DescribeDisposeStrategyPlaybookWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of an entity.
//
// @param request - DescribeEntityInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEntityInfoResponse
func (client *Client) DescribeEntityInfoWithOptions(request *DescribeEntityInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeEntityInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EntityId)) {
		body["EntityId"] = request.EntityId
	}

	if !tea.BoolValue(util.IsUnset(request.EntityIdentity)) {
		body["EntityIdentity"] = request.EntityIdentity
	}

	if !tea.BoolValue(util.IsUnset(request.IncidentUuid)) {
		body["IncidentUuid"] = request.IncidentUuid
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleFor)) {
		body["RoleFor"] = request.RoleFor
	}

	if !tea.BoolValue(util.IsUnset(request.RoleType)) {
		body["RoleType"] = request.RoleType
	}

	if !tea.BoolValue(util.IsUnset(request.SophonTaskId)) {
		body["SophonTaskId"] = request.SophonTaskId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeEntityInfo"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeEntityInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an entity.
//
// @param request - DescribeEntityInfoRequest
//
// @return DescribeEntityInfoResponse
func (client *Client) DescribeEntityInfo(request *DescribeEntityInfoRequest) (_result *DescribeEntityInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeEntityInfoResponse{}
	_body, _err := client.DescribeEntityInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the number of events by type.
//
// @param request - DescribeEventCountByThreatLevelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEventCountByThreatLevelResponse
func (client *Client) DescribeEventCountByThreatLevelWithOptions(request *DescribeEventCountByThreatLevelRequest, runtime *util.RuntimeOptions) (_result *DescribeEventCountByThreatLevelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleFor)) {
		body["RoleFor"] = request.RoleFor
	}

	if !tea.BoolValue(util.IsUnset(request.RoleType)) {
		body["RoleType"] = request.RoleType
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeEventCountByThreatLevel"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeEventCountByThreatLevelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of events by type.
//
// @param request - DescribeEventCountByThreatLevelRequest
//
// @return DescribeEventCountByThreatLevelResponse
func (client *Client) DescribeEventCountByThreatLevel(request *DescribeEventCountByThreatLevelRequest) (_result *DescribeEventCountByThreatLevelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeEventCountByThreatLevelResponse{}
	_body, _err := client.DescribeEventCountByThreatLevelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the handling policies of a historical event.
//
// @param request - DescribeEventDisposeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEventDisposeResponse
func (client *Client) DescribeEventDisposeWithOptions(request *DescribeEventDisposeRequest, runtime *util.RuntimeOptions) (_result *DescribeEventDisposeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.IncidentUuid)) {
		body["IncidentUuid"] = request.IncidentUuid
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleFor)) {
		body["RoleFor"] = request.RoleFor
	}

	if !tea.BoolValue(util.IsUnset(request.RoleType)) {
		body["RoleType"] = request.RoleType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeEventDispose"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeEventDisposeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the handling policies of a historical event.
//
// @param request - DescribeEventDisposeRequest
//
// @return DescribeEventDisposeResponse
func (client *Client) DescribeEventDispose(request *DescribeEventDisposeRequest) (_result *DescribeEventDisposeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeEventDisposeResponse{}
	_body, _err := client.DescribeEventDisposeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the number of logs that are added to the threat analysis feature.
//
// @param request - DescribeImportedLogCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeImportedLogCountResponse
func (client *Client) DescribeImportedLogCountWithOptions(request *DescribeImportedLogCountRequest, runtime *util.RuntimeOptions) (_result *DescribeImportedLogCountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleFor)) {
		body["RoleFor"] = request.RoleFor
	}

	if !tea.BoolValue(util.IsUnset(request.RoleType)) {
		body["RoleType"] = request.RoleType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeImportedLogCount"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeImportedLogCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of logs that are added to the threat analysis feature.
//
// @param request - DescribeImportedLogCountRequest
//
// @return DescribeImportedLogCountResponse
func (client *Client) DescribeImportedLogCount(request *DescribeImportedLogCountRequest) (_result *DescribeImportedLogCountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeImportedLogCountResponse{}
	_body, _err := client.DescribeImportedLogCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the fields that can be configured for a custom rule.
//
// @param request - DescribeLogFieldsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLogFieldsResponse
func (client *Client) DescribeLogFieldsWithOptions(request *DescribeLogFieldsRequest, runtime *util.RuntimeOptions) (_result *DescribeLogFieldsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LogSource)) {
		body["LogSource"] = request.LogSource
	}

	if !tea.BoolValue(util.IsUnset(request.LogType)) {
		body["LogType"] = request.LogType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleFor)) {
		body["RoleFor"] = request.RoleFor
	}

	if !tea.BoolValue(util.IsUnset(request.RoleType)) {
		body["RoleType"] = request.RoleType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeLogFields"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeLogFieldsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the fields that can be configured for a custom rule.
//
// @param request - DescribeLogFieldsRequest
//
// @return DescribeLogFieldsResponse
func (client *Client) DescribeLogFields(request *DescribeLogFieldsRequest) (_result *DescribeLogFieldsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLogFieldsResponse{}
	_body, _err := client.DescribeLogFieldsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the log sources that can be configured for a custom rule.
//
// @param request - DescribeLogSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLogSourceResponse
func (client *Client) DescribeLogSourceWithOptions(request *DescribeLogSourceRequest, runtime *util.RuntimeOptions) (_result *DescribeLogSourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LogType)) {
		body["LogType"] = request.LogType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleFor)) {
		body["RoleFor"] = request.RoleFor
	}

	if !tea.BoolValue(util.IsUnset(request.RoleType)) {
		body["RoleType"] = request.RoleType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeLogSource"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeLogSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the log sources that can be configured for a custom rule.
//
// @param request - DescribeLogSourceRequest
//
// @return DescribeLogSourceResponse
func (client *Client) DescribeLogSource(request *DescribeLogSourceRequest) (_result *DescribeLogSourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLogSourceResponse{}
	_body, _err := client.DescribeLogSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the log types that can be configured for a custom rule.
//
// @param request - DescribeLogTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLogTypeResponse
func (client *Client) DescribeLogTypeWithOptions(request *DescribeLogTypeRequest, runtime *util.RuntimeOptions) (_result *DescribeLogTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleFor)) {
		body["RoleFor"] = request.RoleFor
	}

	if !tea.BoolValue(util.IsUnset(request.RoleType)) {
		body["RoleType"] = request.RoleType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeLogType"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeLogTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the log types that can be configured for a custom rule.
//
// @param request - DescribeLogTypeRequest
//
// @return DescribeLogTypeResponse
func (client *Client) DescribeLogType(request *DescribeLogTypeRequest) (_result *DescribeLogTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLogTypeResponse{}
	_body, _err := client.DescribeLogTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the operator of a custom rule.
//
// @param request - DescribeOperatorsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeOperatorsResponse
func (client *Client) DescribeOperatorsWithOptions(request *DescribeOperatorsRequest, runtime *util.RuntimeOptions) (_result *DescribeOperatorsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleFor)) {
		body["RoleFor"] = request.RoleFor
	}

	if !tea.BoolValue(util.IsUnset(request.RoleType)) {
		body["RoleType"] = request.RoleType
	}

	if !tea.BoolValue(util.IsUnset(request.SceneType)) {
		body["SceneType"] = request.SceneType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeOperators"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeOperatorsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the operator of a custom rule.
//
// @param request - DescribeOperatorsRequest
//
// @return DescribeOperatorsResponse
func (client *Client) DescribeOperators(request *DescribeOperatorsRequest) (_result *DescribeOperatorsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeOperatorsResponse{}
	_body, _err := client.DescribeOperatorsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the number of services that can be added to the threat analysis feature in Alibaba Cloud, Tenant Cloud, and Huawei Cloud.
//
// @param request - DescribeProdCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeProdCountResponse
func (client *Client) DescribeProdCountWithOptions(request *DescribeProdCountRequest, runtime *util.RuntimeOptions) (_result *DescribeProdCountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleFor)) {
		body["RoleFor"] = request.RoleFor
	}

	if !tea.BoolValue(util.IsUnset(request.RoleType)) {
		body["RoleType"] = request.RoleType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeProdCount"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeProdCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of services that can be added to the threat analysis feature in Alibaba Cloud, Tenant Cloud, and Huawei Cloud.
//
// @param request - DescribeProdCountRequest
//
// @return DescribeProdCountResponse
func (client *Client) DescribeProdCount(request *DescribeProdCountRequest) (_result *DescribeProdCountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeProdCountResponse{}
	_body, _err := client.DescribeProdCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of users in the playbook scope.
//
// @param request - DescribeScopeUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeScopeUsersResponse
func (client *Client) DescribeScopeUsersWithOptions(request *DescribeScopeUsersRequest, runtime *util.RuntimeOptions) (_result *DescribeScopeUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleFor)) {
		body["RoleFor"] = request.RoleFor
	}

	if !tea.BoolValue(util.IsUnset(request.RoleType)) {
		body["RoleType"] = request.RoleType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeScopeUsers"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeScopeUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of users in the playbook scope.
//
// @param request - DescribeScopeUsersRequest
//
// @return DescribeScopeUsersResponse
func (client *Client) DescribeScopeUsers(request *DescribeScopeUsersRequest) (_result *DescribeScopeUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScopeUsersResponse{}
	_body, _err := client.DescribeScopeUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Checks whether the threat analysis feature is authorized to access a resource directory.
//
// @param request - DescribeServiceStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeServiceStatusResponse
func (client *Client) DescribeServiceStatusWithOptions(request *DescribeServiceStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeServiceStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeServiceStatus"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeServiceStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks whether the threat analysis feature is authorized to access a resource directory.
//
// @param request - DescribeServiceStatusRequest
//
// @return DescribeServiceStatusResponse
func (client *Client) DescribeServiceStatus(request *DescribeServiceStatusRequest) (_result *DescribeServiceStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeServiceStatusResponse{}
	_body, _err := client.DescribeServiceStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status of the Logstores for the threat analysis feature in Simple Log Service on the user side.
//
// @param request - DescribeStorageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeStorageResponse
func (client *Client) DescribeStorageWithOptions(request *DescribeStorageRequest, runtime *util.RuntimeOptions) (_result *DescribeStorageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleFor)) {
		body["RoleFor"] = request.RoleFor
	}

	if !tea.BoolValue(util.IsUnset(request.RoleType)) {
		body["RoleType"] = request.RoleType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeStorage"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeStorageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of the Logstores for the threat analysis feature in Simple Log Service on the user side.
//
// @param request - DescribeStorageRequest
//
// @return DescribeStorageResponse
func (client *Client) DescribeStorage(request *DescribeStorageRequest) (_result *DescribeStorageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeStorageResponse{}
	_body, _err := client.DescribeStorageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Checks whether the current Alibaba Cloud account or the management account of a resource directory is used to purchase the threat analysis feature.
//
// @param request - DescribeUserBuyStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserBuyStatusResponse
func (client *Client) DescribeUserBuyStatusWithOptions(request *DescribeUserBuyStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeUserBuyStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SubUserId)) {
		body["SubUserId"] = request.SubUserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeUserBuyStatus"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeUserBuyStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks whether the current Alibaba Cloud account or the management account of a resource directory is used to purchase the threat analysis feature.
//
// @param request - DescribeUserBuyStatusRequest
//
// @return DescribeUserBuyStatusResponse
func (client *Client) DescribeUserBuyStatus(request *DescribeUserBuyStatusRequest) (_result *DescribeUserBuyStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUserBuyStatusResponse{}
	_body, _err := client.DescribeUserBuyStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the protected domain names of the WAF instance for a user to which an entity belongs.
//
// @param request - DescribeWafScopeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeWafScopeResponse
func (client *Client) DescribeWafScopeWithOptions(request *DescribeWafScopeRequest, runtime *util.RuntimeOptions) (_result *DescribeWafScopeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EntityId)) {
		body["EntityId"] = request.EntityId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleFor)) {
		body["RoleFor"] = request.RoleFor
	}

	if !tea.BoolValue(util.IsUnset(request.RoleType)) {
		body["RoleType"] = request.RoleType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeWafScope"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeWafScopeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the protected domain names of the WAF instance for a user to which an entity belongs.
//
// @param request - DescribeWafScopeRequest
//
// @return DescribeWafScopeResponse
func (client *Client) DescribeWafScope(request *DescribeWafScopeRequest) (_result *DescribeWafScopeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeWafScopeResponse{}
	_body, _err := client.DescribeWafScopeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of whitelist rules for alerts.
//
// @param request - DescribeWhiteRuleListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeWhiteRuleListResponse
func (client *Client) DescribeWhiteRuleListWithOptions(request *DescribeWhiteRuleListRequest, runtime *util.RuntimeOptions) (_result *DescribeWhiteRuleListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlertName)) {
		body["AlertName"] = request.AlertName
	}

	if !tea.BoolValue(util.IsUnset(request.AlertType)) {
		body["AlertType"] = request.AlertType
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.IncidentUuid)) {
		body["IncidentUuid"] = request.IncidentUuid
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleFor)) {
		body["RoleFor"] = request.RoleFor
	}

	if !tea.BoolValue(util.IsUnset(request.RoleType)) {
		body["RoleType"] = request.RoleType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeWhiteRuleList"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeWhiteRuleListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of whitelist rules for alerts.
//
// @param request - DescribeWhiteRuleListRequest
//
// @return DescribeWhiteRuleListResponse
func (client *Client) DescribeWhiteRuleList(request *DescribeWhiteRuleListRequest) (_result *DescribeWhiteRuleListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeWhiteRuleListResponse{}
	_body, _err := client.DescribeWhiteRuleListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a service-linked role named AliyunServiceRoleForSasCloudSiem for the threat analysis feature. The feature can assume this role to access cloud services.
//
// @param request - EnableAccessForCloudSiemRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableAccessForCloudSiemResponse
func (client *Client) EnableAccessForCloudSiemWithOptions(request *EnableAccessForCloudSiemRequest, runtime *util.RuntimeOptions) (_result *EnableAccessForCloudSiemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoSubmit)) {
		body["AutoSubmit"] = request.AutoSubmit
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleFor)) {
		body["RoleFor"] = request.RoleFor
	}

	if !tea.BoolValue(util.IsUnset(request.RoleType)) {
		body["RoleType"] = request.RoleType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableAccessForCloudSiem"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableAccessForCloudSiemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a service-linked role named AliyunServiceRoleForSasCloudSiem for the threat analysis feature. The feature can assume this role to access cloud services.
//
// @param request - EnableAccessForCloudSiemRequest
//
// @return EnableAccessForCloudSiemResponse
func (client *Client) EnableAccessForCloudSiem(request *EnableAccessForCloudSiemRequest) (_result *EnableAccessForCloudSiemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableAccessForCloudSiemResponse{}
	_body, _err := client.EnableAccessForCloudSiemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Authorizes the threat analysis feature to access a resource directory. This operation must be called by the management account of the resource directory.
//
// @param request - EnableServiceForCloudSiemRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableServiceForCloudSiemResponse
func (client *Client) EnableServiceForCloudSiemWithOptions(request *EnableServiceForCloudSiemRequest, runtime *util.RuntimeOptions) (_result *EnableServiceForCloudSiemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableServiceForCloudSiem"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableServiceForCloudSiemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Authorizes the threat analysis feature to access a resource directory. This operation must be called by the management account of the resource directory.
//
// @param request - EnableServiceForCloudSiemRequest
//
// @return EnableServiceForCloudSiemResponse
func (client *Client) EnableServiceForCloudSiem(request *EnableServiceForCloudSiemRequest) (_result *EnableServiceForCloudSiemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableServiceForCloudSiemResponse{}
	_body, _err := client.EnableServiceForCloudSiemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Queries the storage capacity usage of the threat analysis feature and the purchased storage capacity
//
// @param request - GetCapacityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCapacityResponse
func (client *Client) GetCapacityWithOptions(request *GetCapacityRequest, runtime *util.RuntimeOptions) (_result *GetCapacityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleFor)) {
		body["RoleFor"] = request.RoleFor
	}

	if !tea.BoolValue(util.IsUnset(request.RoleType)) {
		body["RoleType"] = request.RoleType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCapacity"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCapacityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Queries the storage capacity usage of the threat analysis feature and the purchased storage capacity
//
// @param request - GetCapacityRequest
//
// @return GetCapacityResponse
func (client *Client) GetCapacity(request *GetCapacityRequest) (_result *GetCapacityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCapacityResponse{}
	_body, _err := client.GetCapacityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the storage configurations for the threat analysis feature on the user side.
//
// @param request - GetStorageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetStorageResponse
func (client *Client) GetStorageWithOptions(request *GetStorageRequest, runtime *util.RuntimeOptions) (_result *GetStorageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleFor)) {
		body["RoleFor"] = request.RoleFor
	}

	if !tea.BoolValue(util.IsUnset(request.RoleType)) {
		body["RoleType"] = request.RoleType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetStorage"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetStorageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the storage configurations for the threat analysis feature on the user side.
//
// @param request - GetStorageRequest
//
// @return GetStorageResponse
func (client *Client) GetStorage(request *GetStorageRequest) (_result *GetStorageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetStorageResponse{}
	_body, _err := client.GetStorageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of AccessKey IDs of third-party cloud accounts that are added to the threat analysis feature.
//
// @param request - ListAccountAccessIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAccountAccessIdResponse
func (client *Client) ListAccountAccessIdWithOptions(request *ListAccountAccessIdRequest, runtime *util.RuntimeOptions) (_result *ListAccountAccessIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CloudCode)) {
		body["CloudCode"] = request.CloudCode
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleFor)) {
		body["RoleFor"] = request.RoleFor
	}

	if !tea.BoolValue(util.IsUnset(request.RoleType)) {
		body["RoleType"] = request.RoleType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAccountAccessId"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAccountAccessIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of AccessKey IDs of third-party cloud accounts that are added to the threat analysis feature.
//
// @param request - ListAccountAccessIdRequest
//
// @return ListAccountAccessIdResponse
func (client *Client) ListAccountAccessId(request *ListAccountAccessIdRequest) (_result *ListAccountAccessIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAccountAccessIdResponse{}
	_body, _err := client.ListAccountAccessIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query accounts by log.
//
// @param request - ListAccountsByLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAccountsByLogResponse
func (client *Client) ListAccountsByLogWithOptions(request *ListAccountsByLogRequest, runtime *util.RuntimeOptions) (_result *ListAccountsByLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CloudCode)) {
		body["CloudCode"] = request.CloudCode
	}

	if !tea.BoolValue(util.IsUnset(request.LogCodes)) {
		body["LogCodes"] = request.LogCodes
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		body["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleFor)) {
		body["RoleFor"] = request.RoleFor
	}

	if !tea.BoolValue(util.IsUnset(request.RoleType)) {
		body["RoleType"] = request.RoleType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAccountsByLog"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAccountsByLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query accounts by log.
//
// @param request - ListAccountsByLogRequest
//
// @return ListAccountsByLogResponse
func (client *Client) ListAccountsByLog(request *ListAccountsByLogRequest) (_result *ListAccountsByLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAccountsByLogResponse{}
	_body, _err := client.ListAccountsByLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of cloud services that can be added to the threat analysis feature.
//
// @param request - ListAllProdsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAllProdsResponse
func (client *Client) ListAllProdsWithOptions(request *ListAllProdsRequest, runtime *util.RuntimeOptions) (_result *ListAllProdsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleFor)) {
		body["RoleFor"] = request.RoleFor
	}

	if !tea.BoolValue(util.IsUnset(request.RoleType)) {
		body["RoleType"] = request.RoleType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAllProds"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAllProdsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of cloud services that can be added to the threat analysis feature.
//
// @param request - ListAllProdsRequest
//
// @return ListAllProdsResponse
func (client *Client) ListAllProds(request *ListAllProdsRequest) (_result *ListAllProdsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAllProdsResponse{}
	_body, _err := client.ListAllProdsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries automated response rules.
//
// @param request - ListAutomateResponseConfigsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAutomateResponseConfigsResponse
func (client *Client) ListAutomateResponseConfigsWithOptions(request *ListAutomateResponseConfigsRequest, runtime *util.RuntimeOptions) (_result *ListAutomateResponseConfigsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActionType)) {
		body["ActionType"] = request.ActionType
	}

	if !tea.BoolValue(util.IsUnset(request.AutoResponseType)) {
		body["AutoResponseType"] = request.AutoResponseType
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PlaybookUuid)) {
		body["PlaybookUuid"] = request.PlaybookUuid
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResponseRuleType)) {
		body["ResponseRuleType"] = request.ResponseRuleType
	}

	if !tea.BoolValue(util.IsUnset(request.RoleFor)) {
		body["RoleFor"] = request.RoleFor
	}

	if !tea.BoolValue(util.IsUnset(request.RoleType)) {
		body["RoleType"] = request.RoleType
	}

	if !tea.BoolValue(util.IsUnset(request.RuleName)) {
		body["RuleName"] = request.RuleName
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.SubUserId)) {
		body["SubUserId"] = request.SubUserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAutomateResponseConfigs"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAutomateResponseConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries automated response rules.
//
// @param request - ListAutomateResponseConfigsRequest
//
// @return ListAutomateResponseConfigsResponse
func (client *Client) ListAutomateResponseConfigs(request *ListAutomateResponseConfigsRequest) (_result *ListAutomateResponseConfigsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAutomateResponseConfigsResponse{}
	_body, _err := client.ListAutomateResponseConfigsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of cloud accounts that are added to the threat analysis feature.
//
// @param request - ListBindAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListBindAccountResponse
func (client *Client) ListBindAccountWithOptions(request *ListBindAccountRequest, runtime *util.RuntimeOptions) (_result *ListBindAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CloudCode)) {
		body["CloudCode"] = request.CloudCode
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleFor)) {
		body["RoleFor"] = request.RoleFor
	}

	if !tea.BoolValue(util.IsUnset(request.RoleType)) {
		body["RoleType"] = request.RoleType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListBindAccount"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListBindAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of cloud accounts that are added to the threat analysis feature.
//
// @param request - ListBindAccountRequest
//
// @return ListBindAccountResponse
func (client *Client) ListBindAccount(request *ListBindAccountRequest) (_result *ListBindAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListBindAccountResponse{}
	_body, _err := client.ListBindAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of data sources that are added to the threat analysis feature.
//
// @param request - ListBindDataSourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListBindDataSourcesResponse
func (client *Client) ListBindDataSourcesWithOptions(request *ListBindDataSourcesRequest, runtime *util.RuntimeOptions) (_result *ListBindDataSourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		body["AccountId"] = request.AccountId
	}

	if !tea.BoolValue(util.IsUnset(request.CloudCode)) {
		body["CloudCode"] = request.CloudCode
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListBindDataSources"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListBindDataSourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of data sources that are added to the threat analysis feature.
//
// @param request - ListBindDataSourcesRequest
//
// @return ListBindDataSourcesResponse
func (client *Client) ListBindDataSources(request *ListBindDataSourcesRequest) (_result *ListBindDataSourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListBindDataSourcesResponse{}
	_body, _err := client.ListBindDataSourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries custom rules.
//
// @param request - ListCloudSiemCustomizeRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCloudSiemCustomizeRulesResponse
func (client *Client) ListCloudSiemCustomizeRulesWithOptions(request *ListCloudSiemCustomizeRulesRequest, runtime *util.RuntimeOptions) (_result *ListCloudSiemCustomizeRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlertType)) {
		body["AlertType"] = request.AlertType
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		body["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.OrderField)) {
		body["OrderField"] = request.OrderField
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleFor)) {
		body["RoleFor"] = request.RoleFor
	}

	if !tea.BoolValue(util.IsUnset(request.RoleType)) {
		body["RoleType"] = request.RoleType
	}

	if !tea.BoolValue(util.IsUnset(request.RuleName)) {
		body["RuleName"] = request.RuleName
	}

	if !tea.BoolValue(util.IsUnset(request.RuleType)) {
		body["RuleType"] = request.RuleType
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.ThreatLevel)) {
		body["ThreatLevel"] = request.ThreatLevel
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCloudSiemCustomizeRules"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCloudSiemCustomizeRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries custom rules.
//
// @param request - ListCloudSiemCustomizeRulesRequest
//
// @return ListCloudSiemCustomizeRulesResponse
func (client *Client) ListCloudSiemCustomizeRules(request *ListCloudSiemCustomizeRulesRequest) (_result *ListCloudSiemCustomizeRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCloudSiemCustomizeRulesResponse{}
	_body, _err := client.ListCloudSiemCustomizeRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries predefined rules.
//
// @param request - ListCloudSiemPredefinedRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCloudSiemPredefinedRulesResponse
func (client *Client) ListCloudSiemPredefinedRulesWithOptions(request *ListCloudSiemPredefinedRulesRequest, runtime *util.RuntimeOptions) (_result *ListCloudSiemPredefinedRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlertType)) {
		body["AlertType"] = request.AlertType
	}

	if !tea.BoolValue(util.IsUnset(request.AttCk)) {
		body["AttCk"] = request.AttCk
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.EventTransferType)) {
		body["EventTransferType"] = request.EventTransferType
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.LogSource)) {
		body["LogSource"] = request.LogSource
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		body["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.OrderField)) {
		body["OrderField"] = request.OrderField
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleFor)) {
		body["RoleFor"] = request.RoleFor
	}

	if !tea.BoolValue(util.IsUnset(request.RoleType)) {
		body["RoleType"] = request.RoleType
	}

	if !tea.BoolValue(util.IsUnset(request.RuleName)) {
		body["RuleName"] = request.RuleName
	}

	if !tea.BoolValue(util.IsUnset(request.RuleType)) {
		body["RuleType"] = request.RuleType
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.ThreatLevel)) {
		body["ThreatLevel"] = request.ThreatLevel
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCloudSiemPredefinedRules"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCloudSiemPredefinedRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries predefined rules.
//
// @param request - ListCloudSiemPredefinedRulesRequest
//
// @return ListCloudSiemPredefinedRulesResponse
func (client *Client) ListCloudSiemPredefinedRules(request *ListCloudSiemPredefinedRulesRequest) (_result *ListCloudSiemPredefinedRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCloudSiemPredefinedRulesResponse{}
	_body, _err := client.ListCloudSiemPredefinedRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the test results of a custom rule.
//
// @param request - ListCustomizeRuleTestResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCustomizeRuleTestResultResponse
func (client *Client) ListCustomizeRuleTestResultWithOptions(request *ListCustomizeRuleTestResultRequest, runtime *util.RuntimeOptions) (_result *ListCustomizeRuleTestResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.DetectionRuleId)) {
		body["DetectionRuleId"] = request.DetectionRuleId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleFor)) {
		body["RoleFor"] = request.RoleFor
	}

	if !tea.BoolValue(util.IsUnset(request.RoleType)) {
		body["RoleType"] = request.RoleType
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.VerifyType)) {
		body["VerifyType"] = request.VerifyType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCustomizeRuleTestResult"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCustomizeRuleTestResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the test results of a custom rule.
//
// @param request - ListCustomizeRuleTestResultRequest
//
// @return ListCustomizeRuleTestResultResponse
func (client *Client) ListCustomizeRuleTestResult(request *ListCustomizeRuleTestResultRequest) (_result *ListCustomizeRuleTestResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCustomizeRuleTestResultResponse{}
	_body, _err := client.ListCustomizeRuleTestResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the logs of a data source.
//
// @param request - ListDataSourceLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataSourceLogsResponse
func (client *Client) ListDataSourceLogsWithOptions(request *ListDataSourceLogsRequest, runtime *util.RuntimeOptions) (_result *ListDataSourceLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		body["AccountId"] = request.AccountId
	}

	if !tea.BoolValue(util.IsUnset(request.CloudCode)) {
		body["CloudCode"] = request.CloudCode
	}

	if !tea.BoolValue(util.IsUnset(request.DataSourceInstanceId)) {
		body["DataSourceInstanceId"] = request.DataSourceInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDataSourceLogs"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDataSourceLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the logs of a data source.
//
// @param request - ListDataSourceLogsRequest
//
// @return ListDataSourceLogsResponse
func (client *Client) ListDataSourceLogs(request *ListDataSourceLogsRequest) (_result *ListDataSourceLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDataSourceLogsResponse{}
	_body, _err := client.ListDataSourceLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of data source types in third-party cloud services that can be added to the threat analysis feature.
//
// @param request - ListDataSourceTypesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataSourceTypesResponse
func (client *Client) ListDataSourceTypesWithOptions(request *ListDataSourceTypesRequest, runtime *util.RuntimeOptions) (_result *ListDataSourceTypesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CloudCode)) {
		body["CloudCode"] = request.CloudCode
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDataSourceTypes"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDataSourceTypesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of data source types in third-party cloud services that can be added to the threat analysis feature.
//
// @param request - ListDataSourceTypesRequest
//
// @return ListDataSourceTypesResponse
func (client *Client) ListDataSourceTypes(request *ListDataSourceTypesRequest) (_result *ListDataSourceTypesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDataSourceTypesResponse{}
	_body, _err := client.ListDataSourceTypesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about the cloud services that are integrated with the threat analysis feature, the logs of the cloud services, and the delivery of the logs.
//
// @param request - ListDeliveryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDeliveryResponse
func (client *Client) ListDeliveryWithOptions(request *ListDeliveryRequest, runtime *util.RuntimeOptions) (_result *ListDeliveryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleFor)) {
		body["RoleFor"] = request.RoleFor
	}

	if !tea.BoolValue(util.IsUnset(request.RoleType)) {
		body["RoleType"] = request.RoleType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDelivery"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDeliveryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about the cloud services that are integrated with the threat analysis feature, the logs of the cloud services, and the delivery of the logs.
//
// @param request - ListDeliveryRequest
//
// @return ListDeliveryResponse
func (client *Client) ListDelivery(request *ListDeliveryRequest) (_result *ListDeliveryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDeliveryResponse{}
	_body, _err := client.ListDeliveryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries handling policies.
//
// @param request - ListDisposeStrategyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDisposeStrategyResponse
func (client *Client) ListDisposeStrategyWithOptions(request *ListDisposeStrategyRequest, runtime *util.RuntimeOptions) (_result *ListDisposeStrategyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.EffectiveStatus)) {
		body["EffectiveStatus"] = request.EffectiveStatus
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.EntityIdentity)) {
		body["EntityIdentity"] = request.EntityIdentity
	}

	if !tea.BoolValue(util.IsUnset(request.EntityType)) {
		body["EntityType"] = request.EntityType
	}

	if !tea.BoolValue(util.IsUnset(request.IncidentUuid)) {
		body["IncidentUuid"] = request.IncidentUuid
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		body["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.OrderField)) {
		body["OrderField"] = request.OrderField
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PlaybookName)) {
		body["PlaybookName"] = request.PlaybookName
	}

	if !tea.BoolValue(util.IsUnset(request.PlaybookTypes)) {
		body["PlaybookTypes"] = request.PlaybookTypes
	}

	if !tea.BoolValue(util.IsUnset(request.PlaybookUuid)) {
		body["PlaybookUuid"] = request.PlaybookUuid
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleFor)) {
		body["RoleFor"] = request.RoleFor
	}

	if !tea.BoolValue(util.IsUnset(request.RoleType)) {
		body["RoleType"] = request.RoleType
	}

	if !tea.BoolValue(util.IsUnset(request.SophonTaskId)) {
		body["SophonTaskId"] = request.SophonTaskId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDisposeStrategy"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDisposeStrategyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries handling policies.
//
// @param request - ListDisposeStrategyRequest
//
// @return ListDisposeStrategyResponse
func (client *Client) ListDisposeStrategy(request *ListDisposeStrategyRequest) (_result *ListDisposeStrategyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDisposeStrategyResponse{}
	_body, _err := client.ListDisposeStrategyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询实体列表
//
// @param request - ListEntitiesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEntitiesResponse
func (client *Client) ListEntitiesWithOptions(request *ListEntitiesRequest, runtime *util.RuntimeOptions) (_result *ListEntitiesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.EntityName)) {
		body["EntityName"] = request.EntityName
	}

	if !tea.BoolValue(util.IsUnset(request.EntityType)) {
		body["EntityType"] = request.EntityType
	}

	if !tea.BoolValue(util.IsUnset(request.EntityUuid)) {
		body["EntityUuid"] = request.EntityUuid
	}

	if !tea.BoolValue(util.IsUnset(request.IncidentUuid)) {
		body["IncidentUuid"] = request.IncidentUuid
	}

	if !tea.BoolValue(util.IsUnset(request.IsMalwareEntity)) {
		body["IsMalwareEntity"] = request.IsMalwareEntity
	}

	if !tea.BoolValue(util.IsUnset(request.MalwareType)) {
		body["MalwareType"] = request.MalwareType
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleFor)) {
		body["RoleFor"] = request.RoleFor
	}

	if !tea.BoolValue(util.IsUnset(request.RoleType)) {
		body["RoleType"] = request.RoleType
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		body["Tags"] = request.Tags
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListEntities"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListEntitiesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询实体列表
//
// @param request - ListEntitiesRequest
//
// @return ListEntitiesResponse
func (client *Client) ListEntities(request *ListEntitiesRequest) (_result *ListEntitiesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListEntitiesResponse{}
	_body, _err := client.ListEntitiesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of the logs in a cloud service that is added to the threat analysis feature.
//
// @param request - ListImportedLogsByProdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListImportedLogsByProdResponse
func (client *Client) ListImportedLogsByProdWithOptions(request *ListImportedLogsByProdRequest, runtime *util.RuntimeOptions) (_result *ListImportedLogsByProdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CloudCode)) {
		body["CloudCode"] = request.CloudCode
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		body["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleFor)) {
		body["RoleFor"] = request.RoleFor
	}

	if !tea.BoolValue(util.IsUnset(request.RoleType)) {
		body["RoleType"] = request.RoleType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListImportedLogsByProd"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListImportedLogsByProdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of the logs in a cloud service that is added to the threat analysis feature.
//
// @param request - ListImportedLogsByProdRequest
//
// @return ListImportedLogsByProdResponse
func (client *Client) ListImportedLogsByProd(request *ListImportedLogsByProdRequest) (_result *ListImportedLogsByProdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListImportedLogsByProdResponse{}
	_body, _err := client.ListImportedLogsByProdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the dedicated Simple Log Service project and Logstore for a cloud service based on the patterns of the project and Logstore names.
//
// @param request - ListProjectLogStoresRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProjectLogStoresResponse
func (client *Client) ListProjectLogStoresWithOptions(request *ListProjectLogStoresRequest, runtime *util.RuntimeOptions) (_result *ListProjectLogStoresResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceLogCode)) {
		body["SourceLogCode"] = request.SourceLogCode
	}

	if !tea.BoolValue(util.IsUnset(request.SourceProdCode)) {
		body["SourceProdCode"] = request.SourceProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.SubUserId)) {
		body["SubUserId"] = request.SubUserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListProjectLogStores"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListProjectLogStoresResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the dedicated Simple Log Service project and Logstore for a cloud service based on the patterns of the project and Logstore names.
//
// @param request - ListProjectLogStoresRequest
//
// @return ListProjectLogStoresResponse
func (client *Client) ListProjectLogStores(request *ListProjectLogStoresRequest) (_result *ListProjectLogStoresResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListProjectLogStoresResponse{}
	_body, _err := client.ListProjectLogStoresWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of Alibaba Cloud accounts that are added to the threat analysis feature for centralized management. These accounts can be used to perform operations supported by the threat analysis feature, such as adding logs and handling events.
//
// @param request - ListRdUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRdUsersResponse
func (client *Client) ListRdUsersWithOptions(request *ListRdUsersRequest, runtime *util.RuntimeOptions) (_result *ListRdUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListRdUsers"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRdUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of Alibaba Cloud accounts that are added to the threat analysis feature for centralized management. These accounts can be used to perform operations supported by the threat analysis feature, such as adding logs and handling events.
//
// @param request - ListRdUsersRequest
//
// @return ListRdUsersResponse
func (client *Client) ListRdUsers(request *ListRdUsersRequest) (_result *ListRdUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRdUsersResponse{}
	_body, _err := client.ListRdUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a third-party cloud account that is added to the threat analysis feature.
//
// @param request - ModifyBindAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyBindAccountResponse
func (client *Client) ModifyBindAccountWithOptions(request *ModifyBindAccountRequest, runtime *util.RuntimeOptions) (_result *ModifyBindAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessId)) {
		body["AccessId"] = request.AccessId
	}

	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		body["AccountId"] = request.AccountId
	}

	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		body["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.BindId)) {
		body["BindId"] = request.BindId
	}

	if !tea.BoolValue(util.IsUnset(request.CloudCode)) {
		body["CloudCode"] = request.CloudCode
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleFor)) {
		body["RoleFor"] = request.RoleFor
	}

	if !tea.BoolValue(util.IsUnset(request.RoleType)) {
		body["RoleType"] = request.RoleType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyBindAccount"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyBindAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a third-party cloud account that is added to the threat analysis feature.
//
// @param request - ModifyBindAccountRequest
//
// @return ModifyBindAccountResponse
func (client *Client) ModifyBindAccount(request *ModifyBindAccountRequest) (_result *ModifyBindAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyBindAccountResponse{}
	_body, _err := client.ModifyBindAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a data source that is added to the threat analysis feature.
//
// @param request - ModifyDataSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDataSourceResponse
func (client *Client) ModifyDataSourceWithOptions(request *ModifyDataSourceRequest, runtime *util.RuntimeOptions) (_result *ModifyDataSourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		body["AccountId"] = request.AccountId
	}

	if !tea.BoolValue(util.IsUnset(request.CloudCode)) {
		body["CloudCode"] = request.CloudCode
	}

	if !tea.BoolValue(util.IsUnset(request.DataSourceInstanceId)) {
		body["DataSourceInstanceId"] = request.DataSourceInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.DataSourceInstanceName)) {
		body["DataSourceInstanceName"] = request.DataSourceInstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.DataSourceInstanceParams)) {
		body["DataSourceInstanceParams"] = request.DataSourceInstanceParams
	}

	if !tea.BoolValue(util.IsUnset(request.DataSourceInstanceRemark)) {
		body["DataSourceInstanceRemark"] = request.DataSourceInstanceRemark
	}

	if !tea.BoolValue(util.IsUnset(request.DataSourceType)) {
		body["DataSourceType"] = request.DataSourceType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDataSource"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDataSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a data source that is added to the threat analysis feature.
//
// @param request - ModifyDataSourceRequest
//
// @return ModifyDataSourceResponse
func (client *Client) ModifyDataSource(request *ModifyDataSourceRequest) (_result *ModifyDataSourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDataSourceResponse{}
	_body, _err := client.ModifyDataSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the description of the logs that are added to the threat analysis feature for a data source within a cloud account.
//
// @param request - ModifyDataSourceLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDataSourceLogResponse
func (client *Client) ModifyDataSourceLogWithOptions(request *ModifyDataSourceLogRequest, runtime *util.RuntimeOptions) (_result *ModifyDataSourceLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		body["AccountId"] = request.AccountId
	}

	if !tea.BoolValue(util.IsUnset(request.CloudCode)) {
		body["CloudCode"] = request.CloudCode
	}

	if !tea.BoolValue(util.IsUnset(request.DataSourceInstanceId)) {
		body["DataSourceInstanceId"] = request.DataSourceInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.DataSourceInstanceLogs)) {
		body["DataSourceInstanceLogs"] = request.DataSourceInstanceLogs
	}

	if !tea.BoolValue(util.IsUnset(request.DataSourceType)) {
		body["DataSourceType"] = request.DataSourceType
	}

	if !tea.BoolValue(util.IsUnset(request.LogCode)) {
		body["LogCode"] = request.LogCode
	}

	if !tea.BoolValue(util.IsUnset(request.LogInstanceId)) {
		body["LogInstanceId"] = request.LogInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDataSourceLog"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDataSourceLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the description of the logs that are added to the threat analysis feature for a data source within a cloud account.
//
// @param request - ModifyDataSourceLogRequest
//
// @return ModifyDataSourceLogResponse
func (client *Client) ModifyDataSourceLog(request *ModifyDataSourceLogRequest) (_result *ModifyDataSourceLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDataSourceLogResponse{}
	_body, _err := client.ModifyDataSourceLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables the log delivery feature for a cloud service that is integrated with Simple Log Service.
//
// @param request - OpenDeliveryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OpenDeliveryResponse
func (client *Client) OpenDeliveryWithOptions(request *OpenDeliveryRequest, runtime *util.RuntimeOptions) (_result *OpenDeliveryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LogCode)) {
		body["LogCode"] = request.LogCode
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		body["ProductCode"] = request.ProductCode
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleFor)) {
		body["RoleFor"] = request.RoleFor
	}

	if !tea.BoolValue(util.IsUnset(request.RoleType)) {
		body["RoleType"] = request.RoleType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("OpenDelivery"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &OpenDeliveryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables the log delivery feature for a cloud service that is integrated with Simple Log Service.
//
// @param request - OpenDeliveryRequest
//
// @return OpenDeliveryResponse
func (client *Client) OpenDelivery(request *OpenDeliveryRequest) (_result *OpenDeliveryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OpenDeliveryResponse{}
	_body, _err := client.OpenDeliveryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates or updates an automatic response rule.
//
// @param request - PostAutomateResponseConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PostAutomateResponseConfigResponse
func (client *Client) PostAutomateResponseConfigWithOptions(request *PostAutomateResponseConfigRequest, runtime *util.RuntimeOptions) (_result *PostAutomateResponseConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActionConfig)) {
		body["ActionConfig"] = request.ActionConfig
	}

	if !tea.BoolValue(util.IsUnset(request.ActionType)) {
		body["ActionType"] = request.ActionType
	}

	if !tea.BoolValue(util.IsUnset(request.AutoResponseType)) {
		body["AutoResponseType"] = request.AutoResponseType
	}

	if !tea.BoolValue(util.IsUnset(request.ExecutionCondition)) {
		body["ExecutionCondition"] = request.ExecutionCondition
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleFor)) {
		body["RoleFor"] = request.RoleFor
	}

	if !tea.BoolValue(util.IsUnset(request.RoleType)) {
		body["RoleType"] = request.RoleType
	}

	if !tea.BoolValue(util.IsUnset(request.RuleName)) {
		body["RuleName"] = request.RuleName
	}

	if !tea.BoolValue(util.IsUnset(request.SubUserId)) {
		body["SubUserId"] = request.SubUserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PostAutomateResponseConfig"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PostAutomateResponseConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates or updates an automatic response rule.
//
// @param request - PostAutomateResponseConfigRequest
//
// @return PostAutomateResponseConfigResponse
func (client *Client) PostAutomateResponseConfig(request *PostAutomateResponseConfigRequest) (_result *PostAutomateResponseConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PostAutomateResponseConfigResponse{}
	_body, _err := client.PostAutomateResponseConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates or updates a custom rule.
//
// @param request - PostCustomizeRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PostCustomizeRuleResponse
func (client *Client) PostCustomizeRuleWithOptions(request *PostCustomizeRuleRequest, runtime *util.RuntimeOptions) (_result *PostCustomizeRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlertType)) {
		body["AlertType"] = request.AlertType
	}

	if !tea.BoolValue(util.IsUnset(request.AlertTypeMds)) {
		body["AlertTypeMds"] = request.AlertTypeMds
	}

	if !tea.BoolValue(util.IsUnset(request.AttCk)) {
		body["AttCk"] = request.AttCk
	}

	if !tea.BoolValue(util.IsUnset(request.EventTransferExt)) {
		body["EventTransferExt"] = request.EventTransferExt
	}

	if !tea.BoolValue(util.IsUnset(request.EventTransferSwitch)) {
		body["EventTransferSwitch"] = request.EventTransferSwitch
	}

	if !tea.BoolValue(util.IsUnset(request.EventTransferType)) {
		body["EventTransferType"] = request.EventTransferType
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.LogSource)) {
		body["LogSource"] = request.LogSource
	}

	if !tea.BoolValue(util.IsUnset(request.LogSourceMds)) {
		body["LogSourceMds"] = request.LogSourceMds
	}

	if !tea.BoolValue(util.IsUnset(request.LogType)) {
		body["LogType"] = request.LogType
	}

	if !tea.BoolValue(util.IsUnset(request.LogTypeMds)) {
		body["LogTypeMds"] = request.LogTypeMds
	}

	if !tea.BoolValue(util.IsUnset(request.QueryCycle)) {
		body["QueryCycle"] = request.QueryCycle
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleFor)) {
		body["RoleFor"] = request.RoleFor
	}

	if !tea.BoolValue(util.IsUnset(request.RoleType)) {
		body["RoleType"] = request.RoleType
	}

	if !tea.BoolValue(util.IsUnset(request.RuleCondition)) {
		body["RuleCondition"] = request.RuleCondition
	}

	if !tea.BoolValue(util.IsUnset(request.RuleDesc)) {
		body["RuleDesc"] = request.RuleDesc
	}

	if !tea.BoolValue(util.IsUnset(request.RuleGroup)) {
		body["RuleGroup"] = request.RuleGroup
	}

	if !tea.BoolValue(util.IsUnset(request.RuleName)) {
		body["RuleName"] = request.RuleName
	}

	if !tea.BoolValue(util.IsUnset(request.RuleThreshold)) {
		body["RuleThreshold"] = request.RuleThreshold
	}

	if !tea.BoolValue(util.IsUnset(request.ThreatLevel)) {
		body["ThreatLevel"] = request.ThreatLevel
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PostCustomizeRule"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PostCustomizeRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates or updates a custom rule.
//
// @param request - PostCustomizeRuleRequest
//
// @return PostCustomizeRuleResponse
func (client *Client) PostCustomizeRule(request *PostCustomizeRuleRequest) (_result *PostCustomizeRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PostCustomizeRuleResponse{}
	_body, _err := client.PostCustomizeRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Submits a custom rule for testing.
//
// @param request - PostCustomizeRuleTestRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PostCustomizeRuleTestResponse
func (client *Client) PostCustomizeRuleTestWithOptions(request *PostCustomizeRuleTestRequest, runtime *util.RuntimeOptions) (_result *PostCustomizeRuleTestResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleFor)) {
		body["RoleFor"] = request.RoleFor
	}

	if !tea.BoolValue(util.IsUnset(request.RoleType)) {
		body["RoleType"] = request.RoleType
	}

	if !tea.BoolValue(util.IsUnset(request.SimulatedData)) {
		body["SimulatedData"] = request.SimulatedData
	}

	if !tea.BoolValue(util.IsUnset(request.TestType)) {
		body["TestType"] = request.TestType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PostCustomizeRuleTest"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PostCustomizeRuleTestResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits a custom rule for testing.
//
// @param request - PostCustomizeRuleTestRequest
//
// @return PostCustomizeRuleTestResponse
func (client *Client) PostCustomizeRuleTest(request *PostCustomizeRuleTestRequest) (_result *PostCustomizeRuleTestResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PostCustomizeRuleTestResponse{}
	_body, _err := client.PostCustomizeRuleTestWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Submits event handling information.
//
// @param request - PostEventDisposeAndWhiteruleListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PostEventDisposeAndWhiteruleListResponse
func (client *Client) PostEventDisposeAndWhiteruleListWithOptions(request *PostEventDisposeAndWhiteruleListRequest, runtime *util.RuntimeOptions) (_result *PostEventDisposeAndWhiteruleListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EventDispose)) {
		body["EventDispose"] = request.EventDispose
	}

	if !tea.BoolValue(util.IsUnset(request.IncidentUuid)) {
		body["IncidentUuid"] = request.IncidentUuid
	}

	if !tea.BoolValue(util.IsUnset(request.ReceiverInfo)) {
		body["ReceiverInfo"] = request.ReceiverInfo
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		body["Remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.RoleFor)) {
		body["RoleFor"] = request.RoleFor
	}

	if !tea.BoolValue(util.IsUnset(request.RoleType)) {
		body["RoleType"] = request.RoleType
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.ThreatLevel)) {
		body["ThreatLevel"] = request.ThreatLevel
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PostEventDisposeAndWhiteruleList"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PostEventDisposeAndWhiteruleListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits event handling information.
//
// @param request - PostEventDisposeAndWhiteruleListRequest
//
// @return PostEventDisposeAndWhiteruleListResponse
func (client *Client) PostEventDisposeAndWhiteruleList(request *PostEventDisposeAndWhiteruleListRequest) (_result *PostEventDisposeAndWhiteruleListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PostEventDisposeAndWhiteruleListResponse{}
	_body, _err := client.PostEventDisposeAndWhiteruleListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Submits an alert whitelist rule.
//
// @param request - PostEventWhiteruleListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PostEventWhiteruleListResponse
func (client *Client) PostEventWhiteruleListWithOptions(request *PostEventWhiteruleListRequest, runtime *util.RuntimeOptions) (_result *PostEventWhiteruleListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IncidentUuid)) {
		body["IncidentUuid"] = request.IncidentUuid
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleFor)) {
		body["RoleFor"] = request.RoleFor
	}

	if !tea.BoolValue(util.IsUnset(request.RoleType)) {
		body["RoleType"] = request.RoleType
	}

	if !tea.BoolValue(util.IsUnset(request.WhiteruleList)) {
		body["WhiteruleList"] = request.WhiteruleList
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PostEventWhiteruleList"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PostEventWhiteruleListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits an alert whitelist rule.
//
// @param request - PostEventWhiteruleListRequest
//
// @return PostEventWhiteruleListResponse
func (client *Client) PostEventWhiteruleList(request *PostEventWhiteruleListRequest) (_result *PostEventWhiteruleListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PostEventWhiteruleListResponse{}
	_body, _err := client.PostEventWhiteruleListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Ends the test of a custom rule.
//
// @param request - PostFinishCustomizeRuleTestRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PostFinishCustomizeRuleTestResponse
func (client *Client) PostFinishCustomizeRuleTestWithOptions(request *PostFinishCustomizeRuleTestRequest, runtime *util.RuntimeOptions) (_result *PostFinishCustomizeRuleTestResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleFor)) {
		body["RoleFor"] = request.RoleFor
	}

	if !tea.BoolValue(util.IsUnset(request.RoleType)) {
		body["RoleType"] = request.RoleType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PostFinishCustomizeRuleTest"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PostFinishCustomizeRuleTestResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Ends the test of a custom rule.
//
// @param request - PostFinishCustomizeRuleTestRequest
//
// @return PostFinishCustomizeRuleTestResponse
func (client *Client) PostFinishCustomizeRuleTest(request *PostFinishCustomizeRuleTestRequest) (_result *PostFinishCustomizeRuleTestResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PostFinishCustomizeRuleTestResponse{}
	_body, _err := client.PostFinishCustomizeRuleTestWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the status of a custom rule.
//
// @param request - PostRuleStatusChangeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PostRuleStatusChangeResponse
func (client *Client) PostRuleStatusChangeWithOptions(request *PostRuleStatusChangeRequest, runtime *util.RuntimeOptions) (_result *PostRuleStatusChangeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Ids)) {
		body["Ids"] = request.Ids
	}

	if !tea.BoolValue(util.IsUnset(request.InUse)) {
		body["InUse"] = request.InUse
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleFor)) {
		body["RoleFor"] = request.RoleFor
	}

	if !tea.BoolValue(util.IsUnset(request.RoleType)) {
		body["RoleType"] = request.RoleType
	}

	if !tea.BoolValue(util.IsUnset(request.RuleType)) {
		body["RuleType"] = request.RuleType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PostRuleStatusChange"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PostRuleStatusChangeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the status of a custom rule.
//
// @param request - PostRuleStatusChangeRequest
//
// @return PostRuleStatusChangeResponse
func (client *Client) PostRuleStatusChange(request *PostRuleStatusChangeRequest) (_result *PostRuleStatusChangeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PostRuleStatusChangeResponse{}
	_body, _err := client.PostRuleStatusChangeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Releases storage to reduce the storage usage. The release operation is irreversible and may cause data loss. Proceed with caution.
//
// @param request - RestoreCapacityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestoreCapacityResponse
func (client *Client) RestoreCapacityWithOptions(request *RestoreCapacityRequest, runtime *util.RuntimeOptions) (_result *RestoreCapacityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleFor)) {
		body["RoleFor"] = request.RoleFor
	}

	if !tea.BoolValue(util.IsUnset(request.RoleType)) {
		body["RoleType"] = request.RoleType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RestoreCapacity"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RestoreCapacityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Releases storage to reduce the storage usage. The release operation is irreversible and may cause data loss. Proceed with caution.
//
// @param request - RestoreCapacityRequest
//
// @return RestoreCapacityResponse
func (client *Client) RestoreCapacity(request *RestoreCapacityRequest) (_result *RestoreCapacityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RestoreCapacityResponse{}
	_body, _err := client.RestoreCapacityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures the settings of log storage, such as the storage duration and storage region.
//
// @param request - SetStorageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetStorageResponse
func (client *Client) SetStorageWithOptions(request *SetStorageRequest, runtime *util.RuntimeOptions) (_result *SetStorageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Region)) {
		body["Region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleFor)) {
		body["RoleFor"] = request.RoleFor
	}

	if !tea.BoolValue(util.IsUnset(request.RoleType)) {
		body["RoleType"] = request.RoleType
	}

	if !tea.BoolValue(util.IsUnset(request.Ttl)) {
		body["Ttl"] = request.Ttl
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SetStorage"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetStorageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures the settings of log storage, such as the storage duration and storage region.
//
// @param request - SetStorageRequest
//
// @return SetStorageResponse
func (client *Client) SetStorage(request *SetStorageRequest) (_result *SetStorageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetStorageResponse{}
	_body, _err := client.SetStorageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Submits log collection tasks at a time.
//
// @param request - SubmitImportLogTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitImportLogTasksResponse
func (client *Client) SubmitImportLogTasksWithOptions(request *SubmitImportLogTasksRequest, runtime *util.RuntimeOptions) (_result *SubmitImportLogTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Accounts)) {
		body["Accounts"] = request.Accounts
	}

	if !tea.BoolValue(util.IsUnset(request.AutoImported)) {
		body["AutoImported"] = request.AutoImported
	}

	if !tea.BoolValue(util.IsUnset(request.CloudCode)) {
		body["CloudCode"] = request.CloudCode
	}

	if !tea.BoolValue(util.IsUnset(request.LogCodes)) {
		body["LogCodes"] = request.LogCodes
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		body["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleFor)) {
		body["RoleFor"] = request.RoleFor
	}

	if !tea.BoolValue(util.IsUnset(request.RoleType)) {
		body["RoleType"] = request.RoleType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitImportLogTasks"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitImportLogTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits log collection tasks at a time.
//
// @param request - SubmitImportLogTasksRequest
//
// @return SubmitImportLogTasksResponse
func (client *Client) SubmitImportLogTasks(request *SubmitImportLogTasksRequest) (_result *SubmitImportLogTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitImportLogTasksResponse{}
	_body, _err := client.SubmitImportLogTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the status of an automatic response rule.
//
// @param request - UpdateAutomateResponseConfigStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAutomateResponseConfigStatusResponse
func (client *Client) UpdateAutomateResponseConfigStatusWithOptions(request *UpdateAutomateResponseConfigStatusRequest, runtime *util.RuntimeOptions) (_result *UpdateAutomateResponseConfigStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Ids)) {
		body["Ids"] = request.Ids
	}

	if !tea.BoolValue(util.IsUnset(request.InUse)) {
		body["InUse"] = request.InUse
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleFor)) {
		body["RoleFor"] = request.RoleFor
	}

	if !tea.BoolValue(util.IsUnset(request.RoleType)) {
		body["RoleType"] = request.RoleType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAutomateResponseConfigStatus"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAutomateResponseConfigStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the status of an automatic response rule.
//
// @param request - UpdateAutomateResponseConfigStatusRequest
//
// @return UpdateAutomateResponseConfigStatusResponse
func (client *Client) UpdateAutomateResponseConfigStatus(request *UpdateAutomateResponseConfigStatusRequest) (_result *UpdateAutomateResponseConfigStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAutomateResponseConfigStatusResponse{}
	_body, _err := client.UpdateAutomateResponseConfigStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates or updates an alert whitelist rule.
//
// @param request - UpdateWhiteRuleListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateWhiteRuleListResponse
func (client *Client) UpdateWhiteRuleListWithOptions(request *UpdateWhiteRuleListRequest, runtime *util.RuntimeOptions) (_result *UpdateWhiteRuleListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Expression)) {
		body["Expression"] = request.Expression
	}

	if !tea.BoolValue(util.IsUnset(request.IncidentUuid)) {
		body["IncidentUuid"] = request.IncidentUuid
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleFor)) {
		body["RoleFor"] = request.RoleFor
	}

	if !tea.BoolValue(util.IsUnset(request.RoleType)) {
		body["RoleType"] = request.RoleType
	}

	if !tea.BoolValue(util.IsUnset(request.WhiteRuleId)) {
		body["WhiteRuleId"] = request.WhiteRuleId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateWhiteRuleList"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateWhiteRuleListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates or updates an alert whitelist rule.
//
// @param request - UpdateWhiteRuleListRequest
//
// @return UpdateWhiteRuleListResponse
func (client *Client) UpdateWhiteRuleList(request *UpdateWhiteRuleListRequest) (_result *UpdateWhiteRuleListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateWhiteRuleListResponse{}
	_body, _err := client.UpdateWhiteRuleListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
