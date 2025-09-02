// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBaselineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetBaselineResponseBodyData) *GetBaselineResponseBody
	GetData() *GetBaselineResponseBodyData
	SetErrorCode(v string) *GetBaselineResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetBaselineResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *GetBaselineResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *GetBaselineResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetBaselineResponseBody
	GetSuccess() *bool
}

type GetBaselineResponseBody struct {
	// The data returned.
	Data *GetBaselineResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// 1031203110005
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The specified parameters are invalid
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ecb967ec-c137-48a5-860****
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

func (s GetBaselineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBaselineResponseBody) GoString() string {
	return s.String()
}

func (s *GetBaselineResponseBody) GetData() *GetBaselineResponseBodyData {
	return s.Data
}

func (s *GetBaselineResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetBaselineResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetBaselineResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetBaselineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetBaselineResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetBaselineResponseBody) SetData(v *GetBaselineResponseBodyData) *GetBaselineResponseBody {
	s.Data = v
	return s
}

func (s *GetBaselineResponseBody) SetErrorCode(v string) *GetBaselineResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetBaselineResponseBody) SetErrorMessage(v string) *GetBaselineResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetBaselineResponseBody) SetHttpStatusCode(v int32) *GetBaselineResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetBaselineResponseBody) SetRequestId(v string) *GetBaselineResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBaselineResponseBody) SetSuccess(v bool) *GetBaselineResponseBody {
	s.Success = &v
	return s
}

func (s *GetBaselineResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetBaselineResponseBodyData struct {
	// Indicates whether the alerting feature is enabled. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	AlertEnabled *bool `json:"AlertEnabled,omitempty" xml:"AlertEnabled,omitempty"`
	// The alert margin threshold. Unit: minutes.
	//
	// example:
	//
	// 60
	AlertMarginThreshold *int32 `json:"AlertMarginThreshold,omitempty" xml:"AlertMarginThreshold,omitempty"`
	// The alert settings.
	AlertSettings []*GetBaselineResponseBodyDataAlertSettings `json:"AlertSettings,omitempty" xml:"AlertSettings,omitempty" type:"Repeated"`
	// The baseline ID.
	//
	// example:
	//
	// 1001
	BaselineId *int64 `json:"BaselineId,omitempty" xml:"BaselineId,omitempty"`
	// The name of the baseline.
	//
	// example:
	//
	// Test baseline
	BaselineName *string `json:"BaselineName,omitempty" xml:"BaselineName,omitempty"`
	// The type of the baseline. Valid values:
	//
	// 	- DAILY
	//
	// 	- HOURLY
	//
	// example:
	//
	// DAILY
	BaselineType *string `json:"BaselineType,omitempty" xml:"BaselineType,omitempty"`
	// Indicates whether the baseline is enabled.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The node IDs.
	NodeIds []*int64 `json:"NodeIds,omitempty" xml:"NodeIds,omitempty" type:"Repeated"`
	// The settings of the committed completion time of the baseline.
	OverTimeSettings []*GetBaselineResponseBodyDataOverTimeSettings `json:"OverTimeSettings,omitempty" xml:"OverTimeSettings,omitempty" type:"Repeated"`
	// The owner.
	//
	// example:
	//
	// 9527952****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The priority of the baseline. Valid values: 1, 3, 5, 7, and 8.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetBaselineResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetBaselineResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetBaselineResponseBodyData) GetAlertEnabled() *bool {
	return s.AlertEnabled
}

func (s *GetBaselineResponseBodyData) GetAlertMarginThreshold() *int32 {
	return s.AlertMarginThreshold
}

func (s *GetBaselineResponseBodyData) GetAlertSettings() []*GetBaselineResponseBodyDataAlertSettings {
	return s.AlertSettings
}

func (s *GetBaselineResponseBodyData) GetBaselineId() *int64 {
	return s.BaselineId
}

func (s *GetBaselineResponseBodyData) GetBaselineName() *string {
	return s.BaselineName
}

func (s *GetBaselineResponseBodyData) GetBaselineType() *string {
	return s.BaselineType
}

func (s *GetBaselineResponseBodyData) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetBaselineResponseBodyData) GetNodeIds() []*int64 {
	return s.NodeIds
}

func (s *GetBaselineResponseBodyData) GetOverTimeSettings() []*GetBaselineResponseBodyDataOverTimeSettings {
	return s.OverTimeSettings
}

func (s *GetBaselineResponseBodyData) GetOwner() *string {
	return s.Owner
}

func (s *GetBaselineResponseBodyData) GetPriority() *int32 {
	return s.Priority
}

func (s *GetBaselineResponseBodyData) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetBaselineResponseBodyData) SetAlertEnabled(v bool) *GetBaselineResponseBodyData {
	s.AlertEnabled = &v
	return s
}

func (s *GetBaselineResponseBodyData) SetAlertMarginThreshold(v int32) *GetBaselineResponseBodyData {
	s.AlertMarginThreshold = &v
	return s
}

func (s *GetBaselineResponseBodyData) SetAlertSettings(v []*GetBaselineResponseBodyDataAlertSettings) *GetBaselineResponseBodyData {
	s.AlertSettings = v
	return s
}

func (s *GetBaselineResponseBodyData) SetBaselineId(v int64) *GetBaselineResponseBodyData {
	s.BaselineId = &v
	return s
}

func (s *GetBaselineResponseBodyData) SetBaselineName(v string) *GetBaselineResponseBodyData {
	s.BaselineName = &v
	return s
}

func (s *GetBaselineResponseBodyData) SetBaselineType(v string) *GetBaselineResponseBodyData {
	s.BaselineType = &v
	return s
}

func (s *GetBaselineResponseBodyData) SetEnabled(v bool) *GetBaselineResponseBodyData {
	s.Enabled = &v
	return s
}

func (s *GetBaselineResponseBodyData) SetNodeIds(v []*int64) *GetBaselineResponseBodyData {
	s.NodeIds = v
	return s
}

func (s *GetBaselineResponseBodyData) SetOverTimeSettings(v []*GetBaselineResponseBodyDataOverTimeSettings) *GetBaselineResponseBodyData {
	s.OverTimeSettings = v
	return s
}

func (s *GetBaselineResponseBodyData) SetOwner(v string) *GetBaselineResponseBodyData {
	s.Owner = &v
	return s
}

func (s *GetBaselineResponseBodyData) SetPriority(v int32) *GetBaselineResponseBodyData {
	s.Priority = &v
	return s
}

func (s *GetBaselineResponseBodyData) SetProjectId(v int64) *GetBaselineResponseBodyData {
	s.ProjectId = &v
	return s
}

func (s *GetBaselineResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetBaselineResponseBodyDataAlertSettings struct {
	// The event alert interval, in seconds.
	//
	// example:
	//
	// 900
	AlertInterval *int32 `json:"AlertInterval,omitempty" xml:"AlertInterval,omitempty"`
	// The maximum number of event alerts.
	//
	// example:
	//
	// 1
	AlertMaximum *int32 `json:"AlertMaximum,omitempty" xml:"AlertMaximum,omitempty"`
	// Alert method list
	AlertMethods []*string `json:"AlertMethods,omitempty" xml:"AlertMethods,omitempty" type:"Repeated"`
	// Alert recipient details.
	//
	// AlertRecipientType is OWNER: empty
	//
	// AlertRecipientType is SHIFT_SCHEDULE: duty table uid
	//
	// AlertRecipientType is OTHER: uid list, multiple UIDs are in English, split
	//
	// example:
	//
	// 123123
	AlertRecipient *string `json:"AlertRecipient,omitempty" xml:"AlertRecipient,omitempty"`
	// The type of alert recipient.
	//
	// - OWNER: task owner
	//
	// - OTHER: designated person
	//
	// - SHIFT: SCHEDULE-duty table
	//
	// example:
	//
	// OWNER
	AlertRecipientType *string `json:"AlertRecipientType,omitempty" xml:"AlertRecipientType,omitempty"`
	// Alert type
	//
	// - BASELINE: baseline
	//
	// - TOPIC: event
	//
	// example:
	//
	// BASELINE
	AlertType *string `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
	// The baseline alarm switch.
	//
	// - true
	//
	// - false
	//
	// example:
	//
	// false
	BaselineAlertEnabled *bool `json:"BaselineAlertEnabled,omitempty" xml:"BaselineAlertEnabled,omitempty"`
	// DingTalk robot list.
	DingRobots []*GetBaselineResponseBodyDataAlertSettingsDingRobots `json:"DingRobots,omitempty" xml:"DingRobots,omitempty" type:"Repeated"`
	// The end time of the silence. The format is HH:mm:ss.
	//
	// example:
	//
	// 00:00:00
	SilenceEndTime *string `json:"SilenceEndTime,omitempty" xml:"SilenceEndTime,omitempty"`
	// The start time of the silence. Format: HH:mm:ss
	//
	// example:
	//
	// 00:00:00
	SilenceStartTime *string `json:"SilenceStartTime,omitempty" xml:"SilenceStartTime,omitempty"`
	// The list of Event Alert types.
	TopicTypes []*string `json:"TopicTypes,omitempty" xml:"TopicTypes,omitempty" type:"Repeated"`
	// webhook list.
	Webhooks []*string `json:"Webhooks,omitempty" xml:"Webhooks,omitempty" type:"Repeated"`
}

func (s GetBaselineResponseBodyDataAlertSettings) String() string {
	return dara.Prettify(s)
}

func (s GetBaselineResponseBodyDataAlertSettings) GoString() string {
	return s.String()
}

func (s *GetBaselineResponseBodyDataAlertSettings) GetAlertInterval() *int32 {
	return s.AlertInterval
}

func (s *GetBaselineResponseBodyDataAlertSettings) GetAlertMaximum() *int32 {
	return s.AlertMaximum
}

func (s *GetBaselineResponseBodyDataAlertSettings) GetAlertMethods() []*string {
	return s.AlertMethods
}

func (s *GetBaselineResponseBodyDataAlertSettings) GetAlertRecipient() *string {
	return s.AlertRecipient
}

func (s *GetBaselineResponseBodyDataAlertSettings) GetAlertRecipientType() *string {
	return s.AlertRecipientType
}

func (s *GetBaselineResponseBodyDataAlertSettings) GetAlertType() *string {
	return s.AlertType
}

func (s *GetBaselineResponseBodyDataAlertSettings) GetBaselineAlertEnabled() *bool {
	return s.BaselineAlertEnabled
}

func (s *GetBaselineResponseBodyDataAlertSettings) GetDingRobots() []*GetBaselineResponseBodyDataAlertSettingsDingRobots {
	return s.DingRobots
}

func (s *GetBaselineResponseBodyDataAlertSettings) GetSilenceEndTime() *string {
	return s.SilenceEndTime
}

func (s *GetBaselineResponseBodyDataAlertSettings) GetSilenceStartTime() *string {
	return s.SilenceStartTime
}

func (s *GetBaselineResponseBodyDataAlertSettings) GetTopicTypes() []*string {
	return s.TopicTypes
}

func (s *GetBaselineResponseBodyDataAlertSettings) GetWebhooks() []*string {
	return s.Webhooks
}

func (s *GetBaselineResponseBodyDataAlertSettings) SetAlertInterval(v int32) *GetBaselineResponseBodyDataAlertSettings {
	s.AlertInterval = &v
	return s
}

func (s *GetBaselineResponseBodyDataAlertSettings) SetAlertMaximum(v int32) *GetBaselineResponseBodyDataAlertSettings {
	s.AlertMaximum = &v
	return s
}

func (s *GetBaselineResponseBodyDataAlertSettings) SetAlertMethods(v []*string) *GetBaselineResponseBodyDataAlertSettings {
	s.AlertMethods = v
	return s
}

func (s *GetBaselineResponseBodyDataAlertSettings) SetAlertRecipient(v string) *GetBaselineResponseBodyDataAlertSettings {
	s.AlertRecipient = &v
	return s
}

func (s *GetBaselineResponseBodyDataAlertSettings) SetAlertRecipientType(v string) *GetBaselineResponseBodyDataAlertSettings {
	s.AlertRecipientType = &v
	return s
}

func (s *GetBaselineResponseBodyDataAlertSettings) SetAlertType(v string) *GetBaselineResponseBodyDataAlertSettings {
	s.AlertType = &v
	return s
}

func (s *GetBaselineResponseBodyDataAlertSettings) SetBaselineAlertEnabled(v bool) *GetBaselineResponseBodyDataAlertSettings {
	s.BaselineAlertEnabled = &v
	return s
}

func (s *GetBaselineResponseBodyDataAlertSettings) SetDingRobots(v []*GetBaselineResponseBodyDataAlertSettingsDingRobots) *GetBaselineResponseBodyDataAlertSettings {
	s.DingRobots = v
	return s
}

func (s *GetBaselineResponseBodyDataAlertSettings) SetSilenceEndTime(v string) *GetBaselineResponseBodyDataAlertSettings {
	s.SilenceEndTime = &v
	return s
}

func (s *GetBaselineResponseBodyDataAlertSettings) SetSilenceStartTime(v string) *GetBaselineResponseBodyDataAlertSettings {
	s.SilenceStartTime = &v
	return s
}

func (s *GetBaselineResponseBodyDataAlertSettings) SetTopicTypes(v []*string) *GetBaselineResponseBodyDataAlertSettings {
	s.TopicTypes = v
	return s
}

func (s *GetBaselineResponseBodyDataAlertSettings) SetWebhooks(v []*string) *GetBaselineResponseBodyDataAlertSettings {
	s.Webhooks = v
	return s
}

func (s *GetBaselineResponseBodyDataAlertSettings) Validate() error {
	return dara.Validate(s)
}

type GetBaselineResponseBodyDataAlertSettingsDingRobots struct {
	// Whether @ everyone.
	//
	// example:
	//
	// true
	AtAll *bool `json:"AtAll,omitempty" xml:"AtAll,omitempty"`
	// DingTalk robot address
	//
	// example:
	//
	// https://oapi.dingtalk.com/robot/send?access_token=xxx
	WebUrl *string `json:"WebUrl,omitempty" xml:"WebUrl,omitempty"`
}

func (s GetBaselineResponseBodyDataAlertSettingsDingRobots) String() string {
	return dara.Prettify(s)
}

func (s GetBaselineResponseBodyDataAlertSettingsDingRobots) GoString() string {
	return s.String()
}

func (s *GetBaselineResponseBodyDataAlertSettingsDingRobots) GetAtAll() *bool {
	return s.AtAll
}

func (s *GetBaselineResponseBodyDataAlertSettingsDingRobots) GetWebUrl() *string {
	return s.WebUrl
}

func (s *GetBaselineResponseBodyDataAlertSettingsDingRobots) SetAtAll(v bool) *GetBaselineResponseBodyDataAlertSettingsDingRobots {
	s.AtAll = &v
	return s
}

func (s *GetBaselineResponseBodyDataAlertSettingsDingRobots) SetWebUrl(v string) *GetBaselineResponseBodyDataAlertSettingsDingRobots {
	s.WebUrl = &v
	return s
}

func (s *GetBaselineResponseBodyDataAlertSettingsDingRobots) Validate() error {
	return dara.Validate(s)
}

type GetBaselineResponseBodyDataOverTimeSettings struct {
	// The period corresponding to the commitment time. The space-based line is 1, and the hourly baseline can be configured for up to 24 cycles.
	//
	// example:
	//
	// 1
	Cycle *int32 `json:"Cycle,omitempty" xml:"Cycle,omitempty"`
	// Commitment time, hh:mm format, hh value range is [0,47],mm value range is [0,59].
	//
	// example:
	//
	// 00:00
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s GetBaselineResponseBodyDataOverTimeSettings) String() string {
	return dara.Prettify(s)
}

func (s GetBaselineResponseBodyDataOverTimeSettings) GoString() string {
	return s.String()
}

func (s *GetBaselineResponseBodyDataOverTimeSettings) GetCycle() *int32 {
	return s.Cycle
}

func (s *GetBaselineResponseBodyDataOverTimeSettings) GetTime() *string {
	return s.Time
}

func (s *GetBaselineResponseBodyDataOverTimeSettings) SetCycle(v int32) *GetBaselineResponseBodyDataOverTimeSettings {
	s.Cycle = &v
	return s
}

func (s *GetBaselineResponseBodyDataOverTimeSettings) SetTime(v string) *GetBaselineResponseBodyDataOverTimeSettings {
	s.Time = &v
	return s
}

func (s *GetBaselineResponseBodyDataOverTimeSettings) Validate() error {
	return dara.Validate(s)
}
