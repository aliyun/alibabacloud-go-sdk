// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeliveryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListDeliveryResponseBodyData) *ListDeliveryResponseBody
	GetData() *ListDeliveryResponseBodyData
	SetRequestId(v string) *ListDeliveryResponseBody
	GetRequestId() *string
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
	return dara.Prettify(s)
}

func (s ListDeliveryResponseBody) GoString() string {
	return s.String()
}

func (s *ListDeliveryResponseBody) GetData() *ListDeliveryResponseBodyData {
	return s.Data
}

func (s *ListDeliveryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDeliveryResponseBody) SetData(v *ListDeliveryResponseBodyData) *ListDeliveryResponseBody {
	s.Data = v
	return s
}

func (s *ListDeliveryResponseBody) SetRequestId(v string) *ListDeliveryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDeliveryResponseBody) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s ListDeliveryResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDeliveryResponseBodyData) GetDashboardUrl() *string {
	return s.DashboardUrl
}

func (s *ListDeliveryResponseBodyData) GetDisplaySwitchOrNot() *bool {
	return s.DisplaySwitchOrNot
}

func (s *ListDeliveryResponseBodyData) GetLogStoreName() *string {
	return s.LogStoreName
}

func (s *ListDeliveryResponseBodyData) GetProductList() []*ListDeliveryResponseBodyDataProductList {
	return s.ProductList
}

func (s *ListDeliveryResponseBodyData) GetProjectName() *string {
	return s.ProjectName
}

func (s *ListDeliveryResponseBodyData) GetSearchUrl() *string {
	return s.SearchUrl
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

func (s *ListDeliveryResponseBodyData) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s ListDeliveryResponseBodyDataProductList) GoString() string {
	return s.String()
}

func (s *ListDeliveryResponseBodyDataProductList) GetLogList() []*ListDeliveryResponseBodyDataProductListLogList {
	return s.LogList
}

func (s *ListDeliveryResponseBodyDataProductList) GetLogMap() map[string][]*DataProductListLogMapValue {
	return s.LogMap
}

func (s *ListDeliveryResponseBodyDataProductList) GetProductCode() *string {
	return s.ProductCode
}

func (s *ListDeliveryResponseBodyDataProductList) GetProductName() *string {
	return s.ProductName
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

func (s *ListDeliveryResponseBodyDataProductList) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s ListDeliveryResponseBodyDataProductListLogList) GoString() string {
	return s.String()
}

func (s *ListDeliveryResponseBodyDataProductListLogList) GetCanOperateOrNot() *bool {
	return s.CanOperateOrNot
}

func (s *ListDeliveryResponseBodyDataProductListLogList) GetExtraParameters() []*ListDeliveryResponseBodyDataProductListLogListExtraParameters {
	return s.ExtraParameters
}

func (s *ListDeliveryResponseBodyDataProductListLogList) GetLogCode() *string {
	return s.LogCode
}

func (s *ListDeliveryResponseBodyDataProductListLogList) GetLogName() *string {
	return s.LogName
}

func (s *ListDeliveryResponseBodyDataProductListLogList) GetLogNameEn() *string {
	return s.LogNameEn
}

func (s *ListDeliveryResponseBodyDataProductListLogList) GetLogNameKey() *string {
	return s.LogNameKey
}

func (s *ListDeliveryResponseBodyDataProductListLogList) GetStatus() *bool {
	return s.Status
}

func (s *ListDeliveryResponseBodyDataProductListLogList) GetTopic() *string {
	return s.Topic
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

func (s *ListDeliveryResponseBodyDataProductListLogList) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s ListDeliveryResponseBodyDataProductListLogListExtraParameters) GoString() string {
	return s.String()
}

func (s *ListDeliveryResponseBodyDataProductListLogListExtraParameters) GetKey() *string {
	return s.Key
}

func (s *ListDeliveryResponseBodyDataProductListLogListExtraParameters) GetValue() *string {
	return s.Value
}

func (s *ListDeliveryResponseBodyDataProductListLogListExtraParameters) SetKey(v string) *ListDeliveryResponseBodyDataProductListLogListExtraParameters {
	s.Key = &v
	return s
}

func (s *ListDeliveryResponseBodyDataProductListLogListExtraParameters) SetValue(v string) *ListDeliveryResponseBodyDataProductListLogListExtraParameters {
	s.Value = &v
	return s
}

func (s *ListDeliveryResponseBodyDataProductListLogListExtraParameters) Validate() error {
	return dara.Validate(s)
}
