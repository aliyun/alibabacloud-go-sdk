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
	// The returned data.
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
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDeliveryResponseBodyData struct {
	// The URL of the dashboard on the log analysis page.
	//
	// example:
	//
	// https://sls4service.console.aliyun.com/lognext/project/aliyun-cloudsiem-data-127608589417****-cn-shanghai
	//
	// /dashboard/cloud-siem?isShare=true&hideTopbar=true&hideSidebar=true&ignoreTabLocalStorage=true
	DashboardUrl *string `json:"DashboardUrl,omitempty" xml:"DashboardUrl,omitempty"`
	// Indicates whether to display the delivery switch. The default value is true. Valid values:
	//
	// - true: The delivery switch is displayed.
	//
	// - false: The delivery switch is hidden.
	//
	// example:
	//
	// true
	DisplaySwitchOrNot *bool `json:"DisplaySwitchOrNot,omitempty" xml:"DisplaySwitchOrNot,omitempty"`
	// The name of your LogStore for threat analysis. The format is \\`cloud_siem\\`.
	//
	// example:
	//
	// cloud-siem
	LogStoreName *string `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
	// A list of products.
	ProductList []*ListDeliveryResponseBodyDataProductList `json:"ProductList,omitempty" xml:"ProductList,omitempty" type:"Repeated"`
	// The name of your Simple Log Service (SLS) project for threat analysis. The format is \\`aliyun-cloudsiem-data-${aliUid}-${region}\\`.
	//
	// example:
	//
	// aliyun-cloudsiem-data-127608589417****-cn-shanghai
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The URL of the Search & Analysis page in the SLS console.
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
	if s.ProductList != nil {
		for _, item := range s.ProductList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDeliveryResponseBodyDataProductList struct {
	// A list of logs for cloud products that do not have subcategories.
	LogList []*ListDeliveryResponseBodyDataProductListLogList `json:"LogList,omitempty" xml:"LogList,omitempty" type:"Repeated"`
	// A list of logs that are categorized. For example, Security Center logs are categorized into groups such as Host and Network. The group is the key, and the logs in the group are the value.
	LogMap map[string][]*DataProductListLogMapValue `json:"LogMap,omitempty" xml:"LogMap,omitempty"`
	// The code of the cloud product. Valid values:
	//
	// - qcloud_waf
	//
	// - qcloud_cfw
	//
	// - hcloud_waf
	//
	// - hcloud_cfw
	//
	// - ddos
	//
	// - sas
	//
	// - cfw
	//
	// - config
	//
	// - csk
	//
	// - fc
	//
	// - rds
	//
	// - nas
	//
	// - apigateway
	//
	// - cdn
	//
	// - mongodb
	//
	// - eip
	//
	// - slb
	//
	// - vpc
	//
	// - actiontrail
	//
	// - waf
	//
	// - bastionhost
	//
	// - oss
	//
	// - polardb
	//
	// example:
	//
	// sas
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// This parameter is deprecated. You can ignore it.
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
	if s.LogList != nil {
		for _, item := range s.LogList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDeliveryResponseBodyDataProductListLogList struct {
	// Indicates whether the log delivery switch can be operated. Only the delegated administrator for threat analysis can operate the switch. Valid values:
	//
	// - true: The switch can be operated.
	//
	// - false: The switch cannot be operated.
	//
	// example:
	//
	// true
	CanOperateOrNot *bool `json:"CanOperateOrNot,omitempty" xml:"CanOperateOrNot,omitempty"`
	// Additional parameters.
	ExtraParameters []*ListDeliveryResponseBodyDataProductListLogListExtraParameters `json:"ExtraParameters,omitempty" xml:"ExtraParameters,omitempty" type:"Repeated"`
	// The log code.
	//
	// example:
	//
	// cloud_siem_config_log
	LogCode *string `json:"LogCode,omitempty" xml:"LogCode,omitempty"`
	// This parameter is deprecated. You can ignore it.
	//
	// example:
	//
	// audit log
	LogName *string `json:"LogName,omitempty" xml:"LogName,omitempty"`
	// This parameter is deprecated. You can ignore it.
	//
	// example:
	//
	// audit log
	LogNameEn *string `json:"LogNameEn,omitempty" xml:"LogNameEn,omitempty"`
	// The language key of the log name. This key is used to display the log name in different languages.
	//
	// example:
	//
	// ${sas.cloudsiem.prod.cloud_siem_aegis_crack_from_beaver}
	LogNameKey *string `json:"LogNameKey,omitempty" xml:"LogNameKey,omitempty"`
	// The log delivery status. Valid values:
	//
	// - true: Delivery is in progress.
	//
	// - false: Delivery is disabled.
	//
	// example:
	//
	// true
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
	// The topic of the log in the LogStore. This parameter is an index field in the LogStore and is used to differentiate logs.
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
	if s.ExtraParameters != nil {
		for _, item := range s.ExtraParameters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDeliveryResponseBodyDataProductListLogListExtraParameters struct {
	// The key of the additional parameter.
	//
	// example:
	//
	// flag
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the additional parameter.
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
