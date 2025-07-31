// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlertEventsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListAlertEventsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListAlertEventsResponseBody
	GetHttpStatusCode() *int32
	SetListResult(v *ListAlertEventsResponseBodyListResult) *ListAlertEventsResponseBody
	GetListResult() *ListAlertEventsResponseBodyListResult
	SetMessage(v string) *ListAlertEventsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListAlertEventsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListAlertEventsResponseBody
	GetSuccess() *bool
}

type ListAlertEventsResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32                                 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	ListResult     *ListAlertEventsResponseBodyListResult `json:"ListResult,omitempty" xml:"ListResult,omitempty" type:"Struct"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListAlertEventsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAlertEventsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAlertEventsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListAlertEventsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListAlertEventsResponseBody) GetListResult() *ListAlertEventsResponseBodyListResult {
	return s.ListResult
}

func (s *ListAlertEventsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAlertEventsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAlertEventsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAlertEventsResponseBody) SetCode(v string) *ListAlertEventsResponseBody {
	s.Code = &v
	return s
}

func (s *ListAlertEventsResponseBody) SetHttpStatusCode(v int32) *ListAlertEventsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListAlertEventsResponseBody) SetListResult(v *ListAlertEventsResponseBodyListResult) *ListAlertEventsResponseBody {
	s.ListResult = v
	return s
}

func (s *ListAlertEventsResponseBody) SetMessage(v string) *ListAlertEventsResponseBody {
	s.Message = &v
	return s
}

func (s *ListAlertEventsResponseBody) SetRequestId(v string) *ListAlertEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAlertEventsResponseBody) SetSuccess(v bool) *ListAlertEventsResponseBody {
	s.Success = &v
	return s
}

func (s *ListAlertEventsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListAlertEventsResponseBodyListResult struct {
	Data []*ListAlertEventsResponseBodyListResultData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 5
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAlertEventsResponseBodyListResult) String() string {
	return dara.Prettify(s)
}

func (s ListAlertEventsResponseBodyListResult) GoString() string {
	return s.String()
}

func (s *ListAlertEventsResponseBodyListResult) GetData() []*ListAlertEventsResponseBodyListResultData {
	return s.Data
}

func (s *ListAlertEventsResponseBodyListResult) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListAlertEventsResponseBodyListResult) SetData(v []*ListAlertEventsResponseBodyListResultData) *ListAlertEventsResponseBodyListResult {
	s.Data = v
	return s
}

func (s *ListAlertEventsResponseBodyListResult) SetTotalCount(v int32) *ListAlertEventsResponseBodyListResult {
	s.TotalCount = &v
	return s
}

func (s *ListAlertEventsResponseBodyListResult) Validate() error {
	return dara.Validate(s)
}

type ListAlertEventsResponseBodyListResultData struct {
	// example:
	//
	// ONCE
	AlertFrequency    *string                                                       `json:"AlertFrequency,omitempty" xml:"AlertFrequency,omitempty"`
	AlertObject       *ListAlertEventsResponseBodyListResultDataAlertObject         `json:"AlertObject,omitempty" xml:"AlertObject,omitempty" type:"Struct"`
	AlertReason       *ListAlertEventsResponseBodyListResultDataAlertReason         `json:"AlertReason,omitempty" xml:"AlertReason,omitempty" type:"Struct"`
	AlertReceiverList []*ListAlertEventsResponseBodyListResultDataAlertReceiverList `json:"AlertReceiverList,omitempty" xml:"AlertReceiverList,omitempty" type:"Repeated"`
	BelongProject     *ListAlertEventsResponseBodyListResultDataBelongProject       `json:"BelongProject,omitempty" xml:"BelongProject,omitempty" type:"Struct"`
	// example:
	//
	// 2024-11-25 00:00:00
	DoNotDisturbEndTime *string `json:"DoNotDisturbEndTime,omitempty" xml:"DoNotDisturbEndTime,omitempty"`
	// example:
	//
	// 2024-11-25 10:02:47
	FirstAlertTime *string `json:"FirstAlertTime,omitempty" xml:"FirstAlertTime,omitempty"`
	// example:
	//
	// 12345
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 2024-11-25 10:02:47
	LatestAlertTime *string `json:"LatestAlertTime,omitempty" xml:"LatestAlertTime,omitempty"`
	// example:
	//
	// FINISH
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 1
	TotalAlertTimes *int64                                              `json:"TotalAlertTimes,omitempty" xml:"TotalAlertTimes,omitempty"`
	UrlConfig       *ListAlertEventsResponseBodyListResultDataUrlConfig `json:"UrlConfig,omitempty" xml:"UrlConfig,omitempty" type:"Struct"`
}

func (s ListAlertEventsResponseBodyListResultData) String() string {
	return dara.Prettify(s)
}

func (s ListAlertEventsResponseBodyListResultData) GoString() string {
	return s.String()
}

func (s *ListAlertEventsResponseBodyListResultData) GetAlertFrequency() *string {
	return s.AlertFrequency
}

func (s *ListAlertEventsResponseBodyListResultData) GetAlertObject() *ListAlertEventsResponseBodyListResultDataAlertObject {
	return s.AlertObject
}

func (s *ListAlertEventsResponseBodyListResultData) GetAlertReason() *ListAlertEventsResponseBodyListResultDataAlertReason {
	return s.AlertReason
}

func (s *ListAlertEventsResponseBodyListResultData) GetAlertReceiverList() []*ListAlertEventsResponseBodyListResultDataAlertReceiverList {
	return s.AlertReceiverList
}

func (s *ListAlertEventsResponseBodyListResultData) GetBelongProject() *ListAlertEventsResponseBodyListResultDataBelongProject {
	return s.BelongProject
}

func (s *ListAlertEventsResponseBodyListResultData) GetDoNotDisturbEndTime() *string {
	return s.DoNotDisturbEndTime
}

func (s *ListAlertEventsResponseBodyListResultData) GetFirstAlertTime() *string {
	return s.FirstAlertTime
}

func (s *ListAlertEventsResponseBodyListResultData) GetId() *string {
	return s.Id
}

func (s *ListAlertEventsResponseBodyListResultData) GetLatestAlertTime() *string {
	return s.LatestAlertTime
}

func (s *ListAlertEventsResponseBodyListResultData) GetStatus() *string {
	return s.Status
}

func (s *ListAlertEventsResponseBodyListResultData) GetTotalAlertTimes() *int64 {
	return s.TotalAlertTimes
}

func (s *ListAlertEventsResponseBodyListResultData) GetUrlConfig() *ListAlertEventsResponseBodyListResultDataUrlConfig {
	return s.UrlConfig
}

func (s *ListAlertEventsResponseBodyListResultData) SetAlertFrequency(v string) *ListAlertEventsResponseBodyListResultData {
	s.AlertFrequency = &v
	return s
}

func (s *ListAlertEventsResponseBodyListResultData) SetAlertObject(v *ListAlertEventsResponseBodyListResultDataAlertObject) *ListAlertEventsResponseBodyListResultData {
	s.AlertObject = v
	return s
}

func (s *ListAlertEventsResponseBodyListResultData) SetAlertReason(v *ListAlertEventsResponseBodyListResultDataAlertReason) *ListAlertEventsResponseBodyListResultData {
	s.AlertReason = v
	return s
}

func (s *ListAlertEventsResponseBodyListResultData) SetAlertReceiverList(v []*ListAlertEventsResponseBodyListResultDataAlertReceiverList) *ListAlertEventsResponseBodyListResultData {
	s.AlertReceiverList = v
	return s
}

func (s *ListAlertEventsResponseBodyListResultData) SetBelongProject(v *ListAlertEventsResponseBodyListResultDataBelongProject) *ListAlertEventsResponseBodyListResultData {
	s.BelongProject = v
	return s
}

func (s *ListAlertEventsResponseBodyListResultData) SetDoNotDisturbEndTime(v string) *ListAlertEventsResponseBodyListResultData {
	s.DoNotDisturbEndTime = &v
	return s
}

func (s *ListAlertEventsResponseBodyListResultData) SetFirstAlertTime(v string) *ListAlertEventsResponseBodyListResultData {
	s.FirstAlertTime = &v
	return s
}

func (s *ListAlertEventsResponseBodyListResultData) SetId(v string) *ListAlertEventsResponseBodyListResultData {
	s.Id = &v
	return s
}

func (s *ListAlertEventsResponseBodyListResultData) SetLatestAlertTime(v string) *ListAlertEventsResponseBodyListResultData {
	s.LatestAlertTime = &v
	return s
}

func (s *ListAlertEventsResponseBodyListResultData) SetStatus(v string) *ListAlertEventsResponseBodyListResultData {
	s.Status = &v
	return s
}

func (s *ListAlertEventsResponseBodyListResultData) SetTotalAlertTimes(v int64) *ListAlertEventsResponseBodyListResultData {
	s.TotalAlertTimes = &v
	return s
}

func (s *ListAlertEventsResponseBodyListResultData) SetUrlConfig(v *ListAlertEventsResponseBodyListResultDataUrlConfig) *ListAlertEventsResponseBodyListResultData {
	s.UrlConfig = v
	return s
}

func (s *ListAlertEventsResponseBodyListResultData) Validate() error {
	return dara.Validate(s)
}

type ListAlertEventsResponseBodyListResultDataAlertObject struct {
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// VDM_BATCH
	SourceSystemType *string `json:"SourceSystemType,omitempty" xml:"SourceSystemType,omitempty"`
	// example:
	//
	// STREAM_TASK
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListAlertEventsResponseBodyListResultDataAlertObject) String() string {
	return dara.Prettify(s)
}

func (s ListAlertEventsResponseBodyListResultDataAlertObject) GoString() string {
	return s.String()
}

func (s *ListAlertEventsResponseBodyListResultDataAlertObject) GetName() *string {
	return s.Name
}

func (s *ListAlertEventsResponseBodyListResultDataAlertObject) GetSourceSystemType() *string {
	return s.SourceSystemType
}

func (s *ListAlertEventsResponseBodyListResultDataAlertObject) GetType() *string {
	return s.Type
}

func (s *ListAlertEventsResponseBodyListResultDataAlertObject) SetName(v string) *ListAlertEventsResponseBodyListResultDataAlertObject {
	s.Name = &v
	return s
}

func (s *ListAlertEventsResponseBodyListResultDataAlertObject) SetSourceSystemType(v string) *ListAlertEventsResponseBodyListResultDataAlertObject {
	s.SourceSystemType = &v
	return s
}

func (s *ListAlertEventsResponseBodyListResultDataAlertObject) SetType(v string) *ListAlertEventsResponseBodyListResultDataAlertObject {
	s.Type = &v
	return s
}

func (s *ListAlertEventsResponseBodyListResultDataAlertObject) Validate() error {
	return dara.Validate(s)
}

type ListAlertEventsResponseBodyListResultDataAlertReason struct {
	AlertReasonParamList []*ListAlertEventsResponseBodyListResultDataAlertReasonAlertReasonParamList `json:"AlertReasonParamList,omitempty" xml:"AlertReasonParamList,omitempty" type:"Repeated"`
	// example:
	//
	// 2024-11-25 10:02:47
	BizDate *string `json:"BizDate,omitempty" xml:"BizDate,omitempty"`
	// example:
	//
	// VDM_BATCH_FINISH
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// t_6340134343289405440_20241124_639873707610
	UniqueKey *string `json:"UniqueKey,omitempty" xml:"UniqueKey,omitempty"`
}

func (s ListAlertEventsResponseBodyListResultDataAlertReason) String() string {
	return dara.Prettify(s)
}

func (s ListAlertEventsResponseBodyListResultDataAlertReason) GoString() string {
	return s.String()
}

func (s *ListAlertEventsResponseBodyListResultDataAlertReason) GetAlertReasonParamList() []*ListAlertEventsResponseBodyListResultDataAlertReasonAlertReasonParamList {
	return s.AlertReasonParamList
}

func (s *ListAlertEventsResponseBodyListResultDataAlertReason) GetBizDate() *string {
	return s.BizDate
}

func (s *ListAlertEventsResponseBodyListResultDataAlertReason) GetType() *string {
	return s.Type
}

func (s *ListAlertEventsResponseBodyListResultDataAlertReason) GetUniqueKey() *string {
	return s.UniqueKey
}

func (s *ListAlertEventsResponseBodyListResultDataAlertReason) SetAlertReasonParamList(v []*ListAlertEventsResponseBodyListResultDataAlertReasonAlertReasonParamList) *ListAlertEventsResponseBodyListResultDataAlertReason {
	s.AlertReasonParamList = v
	return s
}

func (s *ListAlertEventsResponseBodyListResultDataAlertReason) SetBizDate(v string) *ListAlertEventsResponseBodyListResultDataAlertReason {
	s.BizDate = &v
	return s
}

func (s *ListAlertEventsResponseBodyListResultDataAlertReason) SetType(v string) *ListAlertEventsResponseBodyListResultDataAlertReason {
	s.Type = &v
	return s
}

func (s *ListAlertEventsResponseBodyListResultDataAlertReason) SetUniqueKey(v string) *ListAlertEventsResponseBodyListResultDataAlertReason {
	s.UniqueKey = &v
	return s
}

func (s *ListAlertEventsResponseBodyListResultDataAlertReason) Validate() error {
	return dara.Validate(s)
}

type ListAlertEventsResponseBodyListResultDataAlertReasonAlertReasonParamList struct {
	// example:
	//
	// biz_date
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// 2024-11-24 00:00:00
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListAlertEventsResponseBodyListResultDataAlertReasonAlertReasonParamList) String() string {
	return dara.Prettify(s)
}

func (s ListAlertEventsResponseBodyListResultDataAlertReasonAlertReasonParamList) GoString() string {
	return s.String()
}

func (s *ListAlertEventsResponseBodyListResultDataAlertReasonAlertReasonParamList) GetKey() *string {
	return s.Key
}

func (s *ListAlertEventsResponseBodyListResultDataAlertReasonAlertReasonParamList) GetValue() *string {
	return s.Value
}

func (s *ListAlertEventsResponseBodyListResultDataAlertReasonAlertReasonParamList) SetKey(v string) *ListAlertEventsResponseBodyListResultDataAlertReasonAlertReasonParamList {
	s.Key = &v
	return s
}

func (s *ListAlertEventsResponseBodyListResultDataAlertReasonAlertReasonParamList) SetValue(v string) *ListAlertEventsResponseBodyListResultDataAlertReasonAlertReasonParamList {
	s.Value = &v
	return s
}

func (s *ListAlertEventsResponseBodyListResultDataAlertReasonAlertReasonParamList) Validate() error {
	return dara.Validate(s)
}

type ListAlertEventsResponseBodyListResultDataAlertReceiverList struct {
	AlertChannelTypeList     []*string `json:"AlertChannelTypeList,omitempty" xml:"AlertChannelTypeList,omitempty" type:"Repeated"`
	CustomAlertChannelIdList []*string `json:"CustomAlertChannelIdList,omitempty" xml:"CustomAlertChannelIdList,omitempty" type:"Repeated"`
	// example:
	//
	// test
	OnCallTableName *string `json:"OnCallTableName,omitempty" xml:"OnCallTableName,omitempty"`
	// example:
	//
	// OWNER
	Type     *string                                                               `json:"Type,omitempty" xml:"Type,omitempty"`
	UserList []*ListAlertEventsResponseBodyListResultDataAlertReceiverListUserList `json:"UserList,omitempty" xml:"UserList,omitempty" type:"Repeated"`
}

func (s ListAlertEventsResponseBodyListResultDataAlertReceiverList) String() string {
	return dara.Prettify(s)
}

func (s ListAlertEventsResponseBodyListResultDataAlertReceiverList) GoString() string {
	return s.String()
}

func (s *ListAlertEventsResponseBodyListResultDataAlertReceiverList) GetAlertChannelTypeList() []*string {
	return s.AlertChannelTypeList
}

func (s *ListAlertEventsResponseBodyListResultDataAlertReceiverList) GetCustomAlertChannelIdList() []*string {
	return s.CustomAlertChannelIdList
}

func (s *ListAlertEventsResponseBodyListResultDataAlertReceiverList) GetOnCallTableName() *string {
	return s.OnCallTableName
}

func (s *ListAlertEventsResponseBodyListResultDataAlertReceiverList) GetType() *string {
	return s.Type
}

func (s *ListAlertEventsResponseBodyListResultDataAlertReceiverList) GetUserList() []*ListAlertEventsResponseBodyListResultDataAlertReceiverListUserList {
	return s.UserList
}

func (s *ListAlertEventsResponseBodyListResultDataAlertReceiverList) SetAlertChannelTypeList(v []*string) *ListAlertEventsResponseBodyListResultDataAlertReceiverList {
	s.AlertChannelTypeList = v
	return s
}

func (s *ListAlertEventsResponseBodyListResultDataAlertReceiverList) SetCustomAlertChannelIdList(v []*string) *ListAlertEventsResponseBodyListResultDataAlertReceiverList {
	s.CustomAlertChannelIdList = v
	return s
}

func (s *ListAlertEventsResponseBodyListResultDataAlertReceiverList) SetOnCallTableName(v string) *ListAlertEventsResponseBodyListResultDataAlertReceiverList {
	s.OnCallTableName = &v
	return s
}

func (s *ListAlertEventsResponseBodyListResultDataAlertReceiverList) SetType(v string) *ListAlertEventsResponseBodyListResultDataAlertReceiverList {
	s.Type = &v
	return s
}

func (s *ListAlertEventsResponseBodyListResultDataAlertReceiverList) SetUserList(v []*ListAlertEventsResponseBodyListResultDataAlertReceiverListUserList) *ListAlertEventsResponseBodyListResultDataAlertReceiverList {
	s.UserList = v
	return s
}

func (s *ListAlertEventsResponseBodyListResultDataAlertReceiverList) Validate() error {
	return dara.Validate(s)
}

type ListAlertEventsResponseBodyListResultDataAlertReceiverListUserList struct {
	// example:
	//
	// ADMIN
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListAlertEventsResponseBodyListResultDataAlertReceiverListUserList) String() string {
	return dara.Prettify(s)
}

func (s ListAlertEventsResponseBodyListResultDataAlertReceiverListUserList) GoString() string {
	return s.String()
}

func (s *ListAlertEventsResponseBodyListResultDataAlertReceiverListUserList) GetName() *string {
	return s.Name
}

func (s *ListAlertEventsResponseBodyListResultDataAlertReceiverListUserList) SetName(v string) *ListAlertEventsResponseBodyListResultDataAlertReceiverListUserList {
	s.Name = &v
	return s
}

func (s *ListAlertEventsResponseBodyListResultDataAlertReceiverListUserList) Validate() error {
	return dara.Validate(s)
}

type ListAlertEventsResponseBodyListResultDataBelongProject struct {
	// example:
	//
	// biz_1
	BizName *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	// example:
	//
	// project_1
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s ListAlertEventsResponseBodyListResultDataBelongProject) String() string {
	return dara.Prettify(s)
}

func (s ListAlertEventsResponseBodyListResultDataBelongProject) GoString() string {
	return s.String()
}

func (s *ListAlertEventsResponseBodyListResultDataBelongProject) GetBizName() *string {
	return s.BizName
}

func (s *ListAlertEventsResponseBodyListResultDataBelongProject) GetProjectName() *string {
	return s.ProjectName
}

func (s *ListAlertEventsResponseBodyListResultDataBelongProject) SetBizName(v string) *ListAlertEventsResponseBodyListResultDataBelongProject {
	s.BizName = &v
	return s
}

func (s *ListAlertEventsResponseBodyListResultDataBelongProject) SetProjectName(v string) *ListAlertEventsResponseBodyListResultDataBelongProject {
	s.ProjectName = &v
	return s
}

func (s *ListAlertEventsResponseBodyListResultDataBelongProject) Validate() error {
	return dara.Validate(s)
}

type ListAlertEventsResponseBodyListResultDataUrlConfig struct {
	// example:
	//
	// https://dataphin.com/ops/test3
	AlertConfigUrl *string `json:"AlertConfigUrl,omitempty" xml:"AlertConfigUrl,omitempty"`
	// example:
	//
	// https://dataphin.com/ops/test2
	LogUrl *string `json:"LogUrl,omitempty" xml:"LogUrl,omitempty"`
	// example:
	//
	// https://dataphin.com/ops/test1
	ObjectUrl *string `json:"ObjectUrl,omitempty" xml:"ObjectUrl,omitempty"`
}

func (s ListAlertEventsResponseBodyListResultDataUrlConfig) String() string {
	return dara.Prettify(s)
}

func (s ListAlertEventsResponseBodyListResultDataUrlConfig) GoString() string {
	return s.String()
}

func (s *ListAlertEventsResponseBodyListResultDataUrlConfig) GetAlertConfigUrl() *string {
	return s.AlertConfigUrl
}

func (s *ListAlertEventsResponseBodyListResultDataUrlConfig) GetLogUrl() *string {
	return s.LogUrl
}

func (s *ListAlertEventsResponseBodyListResultDataUrlConfig) GetObjectUrl() *string {
	return s.ObjectUrl
}

func (s *ListAlertEventsResponseBodyListResultDataUrlConfig) SetAlertConfigUrl(v string) *ListAlertEventsResponseBodyListResultDataUrlConfig {
	s.AlertConfigUrl = &v
	return s
}

func (s *ListAlertEventsResponseBodyListResultDataUrlConfig) SetLogUrl(v string) *ListAlertEventsResponseBodyListResultDataUrlConfig {
	s.LogUrl = &v
	return s
}

func (s *ListAlertEventsResponseBodyListResultDataUrlConfig) SetObjectUrl(v string) *ListAlertEventsResponseBodyListResultDataUrlConfig {
	s.ObjectUrl = &v
	return s
}

func (s *ListAlertEventsResponseBodyListResultDataUrlConfig) Validate() error {
	return dara.Validate(s)
}
