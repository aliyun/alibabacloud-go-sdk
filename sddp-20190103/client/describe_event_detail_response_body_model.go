// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEvent(v *DescribeEventDetailResponseBodyEvent) *DescribeEventDetailResponseBody
	GetEvent() *DescribeEventDetailResponseBodyEvent
	SetRequestId(v string) *DescribeEventDetailResponseBody
	GetRequestId() *string
}

type DescribeEventDetailResponseBody struct {
	// The details of the anomalous event.
	Event *DescribeEventDetailResponseBodyEvent `json:"Event,omitempty" xml:"Event,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 69FB3C1-F4C9-42DF-9B72-7077A8989C13
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeEventDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEventDetailResponseBody) GetEvent() *DescribeEventDetailResponseBodyEvent {
	return s.Event
}

func (s *DescribeEventDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEventDetailResponseBody) SetEvent(v *DescribeEventDetailResponseBodyEvent) *DescribeEventDetailResponseBody {
	s.Event = v
	return s
}

func (s *DescribeEventDetailResponseBody) SetRequestId(v string) *DescribeEventDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEventDetailResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeEventDetailResponseBodyEvent struct {
	// The time when the alert for the anomalous event was generated. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1545829129000
	AlertTime *int64 `json:"AlertTime,omitempty" xml:"AlertTime,omitempty"`
	// Indicates whether the handling result of the anomalous event is used to enhance the detection of anomalous events. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// > If you enhance the detection of anomalous events, the detection accuracy and the rate of triggering alerts for anomalous events are improved.
	//
	// example:
	//
	// false
	Backed *bool `json:"Backed,omitempty" xml:"Backed,omitempty"`
	// The instance name of the service in which the anomalous event was detected.
	//
	// example:
	//
	// in-222***
	DataInstance *string `json:"DataInstance,omitempty" xml:"DataInstance,omitempty"`
	// The display name of the account that is used to handle the anomalous event.
	//
	// example:
	//
	// yundunsr
	DealDisplayName *string `json:"DealDisplayName,omitempty" xml:"DealDisplayName,omitempty"`
	// The username of the account that is used to handle the anomalous event.
	//
	// example:
	//
	// det1111
	DealLoginName *string `json:"DealLoginName,omitempty" xml:"DealLoginName,omitempty"`
	// The reason why the anomalous event is handled.
	//
	// example:
	//
	// Anomaly confirmed
	DealReason *string `json:"DealReason,omitempty" xml:"DealReason,omitempty"`
	// The time when the anomalous event was handled. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1230000
	DealTime *int64 `json:"DealTime,omitempty" xml:"DealTime,omitempty"`
	// The ID of the account that is used to handle the anomalous event.
	//
	// example:
	//
	// 229157443385014***
	DealUserId *int64 `json:"DealUserId,omitempty" xml:"DealUserId,omitempty"`
	// The content in the details of the anomalous event.
	Detail *DescribeEventDetailResponseBodyEventDetail `json:"Detail,omitempty" xml:"Detail,omitempty" type:"Struct"`
	// The display name of the account that triggered the anomalous event.
	//
	// example:
	//
	// yundunsr
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The time when the anomalous event occurred. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1545829129000
	EventTime *int64 `json:"EventTime,omitempty" xml:"EventTime,omitempty"`
	// An array that consists of the handling records of the anomalous event.
	HandleInfoList []*DescribeEventDetailResponseBodyEventHandleInfoList `json:"HandleInfoList,omitempty" xml:"HandleInfoList,omitempty" type:"Repeated"`
	// The unique ID of the anomalous event.
	//
	// example:
	//
	// 52234
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The details of the alert logs.
	//
	// example:
	//
	// {"client_ip": ["106.11.XX.XX", "106.11.XX.XX", "106.11.XX.XX", "106.11.XX.XX", "106.11.XX.XX", "106.11.XX.XX", "106.11.XX.XX", "106.11.XX.XX", "106.11.XX.XX"], "start_time": "2020-05-10 00:00:01", "instance": ["omniscience-data", "punish-beaver-data"], "end_time": "2020-05-10 00:21:22", "client_ua": ["Java/1.8.0_152", "Java/1.8.0_92", "aliyun-sdk-java/2.0.0", "aliyun-sdk-java/2.8.0(Linux/4.9.151-015.ali3000.alios7.x86_64/amd64;1.8.0_152)"], "user_name": 1512222261295262}
	LogDetail *string `json:"LogDetail,omitempty" xml:"LogDetail,omitempty"`
	// The username of the account that triggered the anomalous event.
	//
	// example:
	//
	// det1111
	LoginName *string `json:"LoginName,omitempty" xml:"LoginName,omitempty"`
	// Whether it is a new version of the alarm. Value:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// example:
	//
	// true
	NewAlarm *bool `json:"NewAlarm,omitempty" xml:"NewAlarm,omitempty"`
	// The name of the service in which the anomalous event was detected. Valid values include **MaxCompute, OSS, ADS, OTS, and RDS**.
	//
	// example:
	//
	// MaxCompute
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The handling status for the anomalous event. Valid values:
	//
	// 	- **0**: unhandled
	//
	// 	- **1**: confirmed
	//
	// 	- **2**: marked as false positive
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The name of the handling status for the anomalous event.
	//
	// example:
	//
	// Pending
	StatusName *string `json:"StatusName,omitempty" xml:"StatusName,omitempty"`
	// The code of the anomalous event subtype.
	//
	// example:
	//
	// 020008
	SubTypeCode *string `json:"SubTypeCode,omitempty" xml:"SubTypeCode,omitempty"`
	// The name of the anomalous event subtype.
	//
	// example:
	//
	// Anomalous volume of downloaded data
	SubTypeName *string `json:"SubTypeName,omitempty" xml:"SubTypeName,omitempty"`
	// The code of the anomalous event type.
	//
	// example:
	//
	// 02
	TypeCode *string `json:"TypeCode,omitempty" xml:"TypeCode,omitempty"`
	// The name of the anomalous event type. Valid values:
	//
	// 	- **01**: anomalous permission usage
	//
	// 	- **02**: anomalous data flow
	//
	// 	- **03**: anomalous data operation
	//
	// example:
	//
	// Anomalous data flow
	TypeName *string `json:"TypeName,omitempty" xml:"TypeName,omitempty"`
	// The ID of the account that triggered the anomalous event.
	//
	// example:
	//
	// 229157443385014***
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeEventDetailResponseBodyEvent) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventDetailResponseBodyEvent) GoString() string {
	return s.String()
}

func (s *DescribeEventDetailResponseBodyEvent) GetAlertTime() *int64 {
	return s.AlertTime
}

func (s *DescribeEventDetailResponseBodyEvent) GetBacked() *bool {
	return s.Backed
}

func (s *DescribeEventDetailResponseBodyEvent) GetDataInstance() *string {
	return s.DataInstance
}

func (s *DescribeEventDetailResponseBodyEvent) GetDealDisplayName() *string {
	return s.DealDisplayName
}

func (s *DescribeEventDetailResponseBodyEvent) GetDealLoginName() *string {
	return s.DealLoginName
}

func (s *DescribeEventDetailResponseBodyEvent) GetDealReason() *string {
	return s.DealReason
}

func (s *DescribeEventDetailResponseBodyEvent) GetDealTime() *int64 {
	return s.DealTime
}

func (s *DescribeEventDetailResponseBodyEvent) GetDealUserId() *int64 {
	return s.DealUserId
}

func (s *DescribeEventDetailResponseBodyEvent) GetDetail() *DescribeEventDetailResponseBodyEventDetail {
	return s.Detail
}

func (s *DescribeEventDetailResponseBodyEvent) GetDisplayName() *string {
	return s.DisplayName
}

func (s *DescribeEventDetailResponseBodyEvent) GetEventTime() *int64 {
	return s.EventTime
}

func (s *DescribeEventDetailResponseBodyEvent) GetHandleInfoList() []*DescribeEventDetailResponseBodyEventHandleInfoList {
	return s.HandleInfoList
}

func (s *DescribeEventDetailResponseBodyEvent) GetId() *int64 {
	return s.Id
}

func (s *DescribeEventDetailResponseBodyEvent) GetLogDetail() *string {
	return s.LogDetail
}

func (s *DescribeEventDetailResponseBodyEvent) GetLoginName() *string {
	return s.LoginName
}

func (s *DescribeEventDetailResponseBodyEvent) GetNewAlarm() *bool {
	return s.NewAlarm
}

func (s *DescribeEventDetailResponseBodyEvent) GetProductCode() *string {
	return s.ProductCode
}

func (s *DescribeEventDetailResponseBodyEvent) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeEventDetailResponseBodyEvent) GetStatusName() *string {
	return s.StatusName
}

func (s *DescribeEventDetailResponseBodyEvent) GetSubTypeCode() *string {
	return s.SubTypeCode
}

func (s *DescribeEventDetailResponseBodyEvent) GetSubTypeName() *string {
	return s.SubTypeName
}

func (s *DescribeEventDetailResponseBodyEvent) GetTypeCode() *string {
	return s.TypeCode
}

func (s *DescribeEventDetailResponseBodyEvent) GetTypeName() *string {
	return s.TypeName
}

func (s *DescribeEventDetailResponseBodyEvent) GetUserId() *int64 {
	return s.UserId
}

func (s *DescribeEventDetailResponseBodyEvent) SetAlertTime(v int64) *DescribeEventDetailResponseBodyEvent {
	s.AlertTime = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEvent) SetBacked(v bool) *DescribeEventDetailResponseBodyEvent {
	s.Backed = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEvent) SetDataInstance(v string) *DescribeEventDetailResponseBodyEvent {
	s.DataInstance = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEvent) SetDealDisplayName(v string) *DescribeEventDetailResponseBodyEvent {
	s.DealDisplayName = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEvent) SetDealLoginName(v string) *DescribeEventDetailResponseBodyEvent {
	s.DealLoginName = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEvent) SetDealReason(v string) *DescribeEventDetailResponseBodyEvent {
	s.DealReason = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEvent) SetDealTime(v int64) *DescribeEventDetailResponseBodyEvent {
	s.DealTime = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEvent) SetDealUserId(v int64) *DescribeEventDetailResponseBodyEvent {
	s.DealUserId = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEvent) SetDetail(v *DescribeEventDetailResponseBodyEventDetail) *DescribeEventDetailResponseBodyEvent {
	s.Detail = v
	return s
}

func (s *DescribeEventDetailResponseBodyEvent) SetDisplayName(v string) *DescribeEventDetailResponseBodyEvent {
	s.DisplayName = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEvent) SetEventTime(v int64) *DescribeEventDetailResponseBodyEvent {
	s.EventTime = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEvent) SetHandleInfoList(v []*DescribeEventDetailResponseBodyEventHandleInfoList) *DescribeEventDetailResponseBodyEvent {
	s.HandleInfoList = v
	return s
}

func (s *DescribeEventDetailResponseBodyEvent) SetId(v int64) *DescribeEventDetailResponseBodyEvent {
	s.Id = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEvent) SetLogDetail(v string) *DescribeEventDetailResponseBodyEvent {
	s.LogDetail = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEvent) SetLoginName(v string) *DescribeEventDetailResponseBodyEvent {
	s.LoginName = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEvent) SetNewAlarm(v bool) *DescribeEventDetailResponseBodyEvent {
	s.NewAlarm = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEvent) SetProductCode(v string) *DescribeEventDetailResponseBodyEvent {
	s.ProductCode = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEvent) SetStatus(v int32) *DescribeEventDetailResponseBodyEvent {
	s.Status = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEvent) SetStatusName(v string) *DescribeEventDetailResponseBodyEvent {
	s.StatusName = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEvent) SetSubTypeCode(v string) *DescribeEventDetailResponseBodyEvent {
	s.SubTypeCode = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEvent) SetSubTypeName(v string) *DescribeEventDetailResponseBodyEvent {
	s.SubTypeName = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEvent) SetTypeCode(v string) *DescribeEventDetailResponseBodyEvent {
	s.TypeCode = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEvent) SetTypeName(v string) *DescribeEventDetailResponseBodyEvent {
	s.TypeName = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEvent) SetUserId(v int64) *DescribeEventDetailResponseBodyEvent {
	s.UserId = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEvent) Validate() error {
	return dara.Validate(s)
}

type DescribeEventDetailResponseBodyEventDetail struct {
	// The baseline behavior chart of the anomalous event.
	Chart []*DescribeEventDetailResponseBodyEventDetailChart `json:"Chart,omitempty" xml:"Chart,omitempty" type:"Repeated"`
	// The content in the anomalous event.
	Content []*DescribeEventDetailResponseBodyEventDetailContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Repeated"`
	// An array that consists of the source from which the information of the anomalous event is recorded.
	ResourceInfo []*DescribeEventDetailResponseBodyEventDetailResourceInfo `json:"ResourceInfo,omitempty" xml:"ResourceInfo,omitempty" type:"Repeated"`
}

func (s DescribeEventDetailResponseBodyEventDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventDetailResponseBodyEventDetail) GoString() string {
	return s.String()
}

func (s *DescribeEventDetailResponseBodyEventDetail) GetChart() []*DescribeEventDetailResponseBodyEventDetailChart {
	return s.Chart
}

func (s *DescribeEventDetailResponseBodyEventDetail) GetContent() []*DescribeEventDetailResponseBodyEventDetailContent {
	return s.Content
}

func (s *DescribeEventDetailResponseBodyEventDetail) GetResourceInfo() []*DescribeEventDetailResponseBodyEventDetailResourceInfo {
	return s.ResourceInfo
}

func (s *DescribeEventDetailResponseBodyEventDetail) SetChart(v []*DescribeEventDetailResponseBodyEventDetailChart) *DescribeEventDetailResponseBodyEventDetail {
	s.Chart = v
	return s
}

func (s *DescribeEventDetailResponseBodyEventDetail) SetContent(v []*DescribeEventDetailResponseBodyEventDetailContent) *DescribeEventDetailResponseBodyEventDetail {
	s.Content = v
	return s
}

func (s *DescribeEventDetailResponseBodyEventDetail) SetResourceInfo(v []*DescribeEventDetailResponseBodyEventDetailResourceInfo) *DescribeEventDetailResponseBodyEventDetail {
	s.ResourceInfo = v
	return s
}

func (s *DescribeEventDetailResponseBodyEventDetail) Validate() error {
	return dara.Validate(s)
}

type DescribeEventDetailResponseBodyEventDetailChart struct {
	// The type of the chart. Valid values:
	//
	// 	- **1**: column chart
	//
	// 	- **2**: line chart
	//
	// >This field will be returned only when NewAlarm is true.
	//
	// example:
	//
	// 1
	ChatType *int32 `json:"ChatType,omitempty" xml:"ChatType,omitempty"`
	// The data in the baseline behavior profile of the anomalous event.
	Data *DescribeEventDetailResponseBodyEventDetailChartData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The name of the baseline behavior chart of the anomalous event.
	//
	// example:
	//
	// Baseline behavior chart
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// Icon title.
	//
	// >This field will be returned only when NewAlarm is true.
	//
	// example:
	//
	// misskingm
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the chart. Valid values:
	//
	// 	- **1**: column chart
	//
	// 	- **2**: line chart
	//
	// example:
	//
	// 1
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The descriptive label of data items on the X axis.
	//
	// example:
	//
	// Number of days
	XLabel *string `json:"XLabel,omitempty" xml:"XLabel,omitempty"`
	// The descriptive label of data items on the Y axis.
	//
	// example:
	//
	// Value
	YLabel *string `json:"YLabel,omitempty" xml:"YLabel,omitempty"`
	// The descriptive label of data items on the Z axis.
	//
	// >This field will be returned only when NewAlarm is true.
	//
	// example:
	//
	// chart description
	ZLabel *string `json:"ZLabel,omitempty" xml:"ZLabel,omitempty"`
}

func (s DescribeEventDetailResponseBodyEventDetailChart) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventDetailResponseBodyEventDetailChart) GoString() string {
	return s.String()
}

func (s *DescribeEventDetailResponseBodyEventDetailChart) GetChatType() *int32 {
	return s.ChatType
}

func (s *DescribeEventDetailResponseBodyEventDetailChart) GetData() *DescribeEventDetailResponseBodyEventDetailChartData {
	return s.Data
}

func (s *DescribeEventDetailResponseBodyEventDetailChart) GetLabel() *string {
	return s.Label
}

func (s *DescribeEventDetailResponseBodyEventDetailChart) GetName() *string {
	return s.Name
}

func (s *DescribeEventDetailResponseBodyEventDetailChart) GetType() *string {
	return s.Type
}

func (s *DescribeEventDetailResponseBodyEventDetailChart) GetXLabel() *string {
	return s.XLabel
}

func (s *DescribeEventDetailResponseBodyEventDetailChart) GetYLabel() *string {
	return s.YLabel
}

func (s *DescribeEventDetailResponseBodyEventDetailChart) GetZLabel() *string {
	return s.ZLabel
}

func (s *DescribeEventDetailResponseBodyEventDetailChart) SetChatType(v int32) *DescribeEventDetailResponseBodyEventDetailChart {
	s.ChatType = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEventDetailChart) SetData(v *DescribeEventDetailResponseBodyEventDetailChartData) *DescribeEventDetailResponseBodyEventDetailChart {
	s.Data = v
	return s
}

func (s *DescribeEventDetailResponseBodyEventDetailChart) SetLabel(v string) *DescribeEventDetailResponseBodyEventDetailChart {
	s.Label = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEventDetailChart) SetName(v string) *DescribeEventDetailResponseBodyEventDetailChart {
	s.Name = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEventDetailChart) SetType(v string) *DescribeEventDetailResponseBodyEventDetailChart {
	s.Type = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEventDetailChart) SetXLabel(v string) *DescribeEventDetailResponseBodyEventDetailChart {
	s.XLabel = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEventDetailChart) SetYLabel(v string) *DescribeEventDetailResponseBodyEventDetailChart {
	s.YLabel = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEventDetailChart) SetZLabel(v string) *DescribeEventDetailResponseBodyEventDetailChart {
	s.ZLabel = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEventDetailChart) Validate() error {
	return dara.Validate(s)
}

type DescribeEventDetailResponseBodyEventDetailChartData struct {
	// The value of the data item on the X axis.
	//
	// example:
	//
	// [test1,test2,...]
	X []*string `json:"X,omitempty" xml:"X,omitempty" type:"Repeated"`
	// The value of the data item on the Y axis.
	//
	// example:
	//
	// [1,2,3,...]
	Y []*string `json:"Y,omitempty" xml:"Y,omitempty" type:"Repeated"`
	// The value of the data item for the Z axis.
	Z []*string `json:"Z,omitempty" xml:"Z,omitempty" type:"Repeated"`
}

func (s DescribeEventDetailResponseBodyEventDetailChartData) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventDetailResponseBodyEventDetailChartData) GoString() string {
	return s.String()
}

func (s *DescribeEventDetailResponseBodyEventDetailChartData) GetX() []*string {
	return s.X
}

func (s *DescribeEventDetailResponseBodyEventDetailChartData) GetY() []*string {
	return s.Y
}

func (s *DescribeEventDetailResponseBodyEventDetailChartData) GetZ() []*string {
	return s.Z
}

func (s *DescribeEventDetailResponseBodyEventDetailChartData) SetX(v []*string) *DescribeEventDetailResponseBodyEventDetailChartData {
	s.X = v
	return s
}

func (s *DescribeEventDetailResponseBodyEventDetailChartData) SetY(v []*string) *DescribeEventDetailResponseBodyEventDetailChartData {
	s.Y = v
	return s
}

func (s *DescribeEventDetailResponseBodyEventDetailChartData) SetZ(v []*string) *DescribeEventDetailResponseBodyEventDetailChartData {
	s.Z = v
	return s
}

func (s *DescribeEventDetailResponseBodyEventDetailChartData) Validate() error {
	return dara.Validate(s)
}

type DescribeEventDetailResponseBodyEventDetailContent struct {
	// The title of the content in the anomalous event.
	//
	// example:
	//
	// Anomaly description
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// Exception event name.
	//
	// example:
	//
	// daliaoyuncom
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The description of the content in the anomalous event.
	//
	// example:
	//
	// The account was used to access OSS from an unusual terminal whose IP address is 1.2.3.4 from 00:06:45 on September 9, 2019 to 00:57:37 on September 9, 2019.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeEventDetailResponseBodyEventDetailContent) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventDetailResponseBodyEventDetailContent) GoString() string {
	return s.String()
}

func (s *DescribeEventDetailResponseBodyEventDetailContent) GetLabel() *string {
	return s.Label
}

func (s *DescribeEventDetailResponseBodyEventDetailContent) GetName() *string {
	return s.Name
}

func (s *DescribeEventDetailResponseBodyEventDetailContent) GetValue() *string {
	return s.Value
}

func (s *DescribeEventDetailResponseBodyEventDetailContent) SetLabel(v string) *DescribeEventDetailResponseBodyEventDetailContent {
	s.Label = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEventDetailContent) SetName(v string) *DescribeEventDetailResponseBodyEventDetailContent {
	s.Name = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEventDetailContent) SetValue(v string) *DescribeEventDetailResponseBodyEventDetailContent {
	s.Value = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEventDetailContent) Validate() error {
	return dara.Validate(s)
}

type DescribeEventDetailResponseBodyEventDetailResourceInfo struct {
	// The source title.
	//
	// example:
	//
	// Risk
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The source description.
	//
	// example:
	//
	// Based on the record of authentication by using an unusual terminal, an attacker may have obtained the access permission of the account, or an employee accessed data from a personal terminal.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeEventDetailResponseBodyEventDetailResourceInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventDetailResponseBodyEventDetailResourceInfo) GoString() string {
	return s.String()
}

func (s *DescribeEventDetailResponseBodyEventDetailResourceInfo) GetLabel() *string {
	return s.Label
}

func (s *DescribeEventDetailResponseBodyEventDetailResourceInfo) GetValue() *string {
	return s.Value
}

func (s *DescribeEventDetailResponseBodyEventDetailResourceInfo) SetLabel(v string) *DescribeEventDetailResponseBodyEventDetailResourceInfo {
	s.Label = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEventDetailResourceInfo) SetValue(v string) *DescribeEventDetailResponseBodyEventDetailResourceInfo {
	s.Value = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEventDetailResourceInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeEventDetailResponseBodyEventHandleInfoList struct {
	// The account that is used to handle the anomalous event.
	//
	// example:
	//
	// sddp-test2
	CurrentValue *string `json:"CurrentValue,omitempty" xml:"CurrentValue,omitempty"`
	// The time when the account is disabled. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1611139155000
	DisableTime *int64 `json:"DisableTime,omitempty" xml:"DisableTime,omitempty"`
	// The time when the disabled account is enabled. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1611139155000
	EnableTime *int64 `json:"EnableTime,omitempty" xml:"EnableTime,omitempty"`
	// The handling method.
	//
	// example:
	//
	// Remove from the whitelist
	HandlerName *string `json:"HandlerName,omitempty" xml:"HandlerName,omitempty"`
	// The type of the handling method.
	//
	// example:
	//
	// rds_security_ip
	HandlerType *string `json:"HandlerType,omitempty" xml:"HandlerType,omitempty"`
	// The duration for which the handling operation takes effect. If you leave this parameter empty, the handling operation is permanently valid. Unit: minutes.
	//
	// example:
	//
	// 10
	HandlerValue *int32 `json:"HandlerValue,omitempty" xml:"HandlerValue,omitempty"`
	// The ID of the handling rule.
	//
	// example:
	//
	// 11
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The status of the account that triggered the anomalous event. Valid values:
	//
	// 	- **0**: disabled
	//
	// 	- **1**: enabled
	//
	// 	- **-1**: failed to disable the account
	//
	// 	- **-2**: failed to enable the account
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeEventDetailResponseBodyEventHandleInfoList) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventDetailResponseBodyEventHandleInfoList) GoString() string {
	return s.String()
}

func (s *DescribeEventDetailResponseBodyEventHandleInfoList) GetCurrentValue() *string {
	return s.CurrentValue
}

func (s *DescribeEventDetailResponseBodyEventHandleInfoList) GetDisableTime() *int64 {
	return s.DisableTime
}

func (s *DescribeEventDetailResponseBodyEventHandleInfoList) GetEnableTime() *int64 {
	return s.EnableTime
}

func (s *DescribeEventDetailResponseBodyEventHandleInfoList) GetHandlerName() *string {
	return s.HandlerName
}

func (s *DescribeEventDetailResponseBodyEventHandleInfoList) GetHandlerType() *string {
	return s.HandlerType
}

func (s *DescribeEventDetailResponseBodyEventHandleInfoList) GetHandlerValue() *int32 {
	return s.HandlerValue
}

func (s *DescribeEventDetailResponseBodyEventHandleInfoList) GetId() *int64 {
	return s.Id
}

func (s *DescribeEventDetailResponseBodyEventHandleInfoList) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeEventDetailResponseBodyEventHandleInfoList) SetCurrentValue(v string) *DescribeEventDetailResponseBodyEventHandleInfoList {
	s.CurrentValue = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEventHandleInfoList) SetDisableTime(v int64) *DescribeEventDetailResponseBodyEventHandleInfoList {
	s.DisableTime = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEventHandleInfoList) SetEnableTime(v int64) *DescribeEventDetailResponseBodyEventHandleInfoList {
	s.EnableTime = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEventHandleInfoList) SetHandlerName(v string) *DescribeEventDetailResponseBodyEventHandleInfoList {
	s.HandlerName = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEventHandleInfoList) SetHandlerType(v string) *DescribeEventDetailResponseBodyEventHandleInfoList {
	s.HandlerType = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEventHandleInfoList) SetHandlerValue(v int32) *DescribeEventDetailResponseBodyEventHandleInfoList {
	s.HandlerValue = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEventHandleInfoList) SetId(v int64) *DescribeEventDetailResponseBodyEventHandleInfoList {
	s.Id = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEventHandleInfoList) SetStatus(v int32) *DescribeEventDetailResponseBodyEventHandleInfoList {
	s.Status = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEventHandleInfoList) Validate() error {
	return dara.Validate(s)
}
