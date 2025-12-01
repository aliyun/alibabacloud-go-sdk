// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAlarmDetailByIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetAlarmDetailByIdResponseBody
	GetCode() *string
	SetData(v *GetAlarmDetailByIdResponseBodyData) *GetAlarmDetailByIdResponseBody
	GetData() *GetAlarmDetailByIdResponseBodyData
	SetHttpStatusCode(v int32) *GetAlarmDetailByIdResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetAlarmDetailByIdResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAlarmDetailByIdResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAlarmDetailByIdResponseBody
	GetSuccess() *bool
}

type GetAlarmDetailByIdResponseBody struct {
	// API response code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Data returned by the interface.
	Data *GetAlarmDetailByIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Return message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5C1B0668-442C-57AE-9668-D894B0B012EB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the operation was successful: - true: Success. - false: Failure.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAlarmDetailByIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAlarmDetailByIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetAlarmDetailByIdResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAlarmDetailByIdResponseBody) GetData() *GetAlarmDetailByIdResponseBodyData {
	return s.Data
}

func (s *GetAlarmDetailByIdResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetAlarmDetailByIdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAlarmDetailByIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAlarmDetailByIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAlarmDetailByIdResponseBody) SetCode(v string) *GetAlarmDetailByIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetAlarmDetailByIdResponseBody) SetData(v *GetAlarmDetailByIdResponseBodyData) *GetAlarmDetailByIdResponseBody {
	s.Data = v
	return s
}

func (s *GetAlarmDetailByIdResponseBody) SetHttpStatusCode(v int32) *GetAlarmDetailByIdResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetAlarmDetailByIdResponseBody) SetMessage(v string) *GetAlarmDetailByIdResponseBody {
	s.Message = &v
	return s
}

func (s *GetAlarmDetailByIdResponseBody) SetRequestId(v string) *GetAlarmDetailByIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAlarmDetailByIdResponseBody) SetSuccess(v bool) *GetAlarmDetailByIdResponseBody {
	s.Success = &v
	return s
}

func (s *GetAlarmDetailByIdResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAlarmDetailByIdResponseBodyData struct {
	// Alarm event type.
	//
	// example:
	//
	// Unusual Logon
	AlarmEventType *string `json:"AlarmEventType,omitempty" xml:"AlarmEventType,omitempty"`
	// Alarm event type.
	//
	// example:
	//
	// Login with unusual location
	AlarmEventTypeDisplay *string `json:"AlarmEventTypeDisplay,omitempty" xml:"AlarmEventTypeDisplay,omitempty"`
	// Alarm ID.
	//
	// example:
	//
	// 202427220
	AlarmId *int64 `json:"AlarmId,omitempty" xml:"AlarmId,omitempty"`
	// Alarm name.
	//
	// example:
	//
	// 负载均衡可挂载服务器数量告警
	AlarmName *string `json:"AlarmName,omitempty" xml:"AlarmName,omitempty"`
	// Alarm source.
	//
	// example:
	//
	// SUSP_EVENT
	AlarmSource *string `json:"AlarmSource,omitempty" xml:"AlarmSource,omitempty"`
	// Latest alarm time.
	//
	// example:
	//
	// 2018-09-26 01:51:01
	AlarmTime *string `json:"AlarmTime,omitempty" xml:"AlarmTime,omitempty"`
	// Analysis process.
	//
	// example:
	//
	// [{"value":"服务器可能已被黑客攻击，存在恶意进程在运行。 分析过程：告警显示，服务端存在一个名为”dns.exe”的进程在访问”polling.burpcollaborator.net”，这是一个被黑名单列出的恶意域名。在正常情况下,”dns.exe”不应该单独存在于系统的路径下，并且也不应该访问这类恶意域名。因此，这个进程可能是黑客留下的恶意进程。","key":"结论"},{"value":"尽快对服务器进行全面扫描，清除恶意进程。同时，联系网络安全专家进行深入调查，以确定是否有其他潜在的安全威胁。","key":"处置建议"}]
	AnalysisResult *string `json:"AnalysisResult,omitempty" xml:"AnalysisResult,omitempty"`
	// Whether high-protection mode is enabled. true means enabled, false means not enabled.
	//
	// example:
	//
	// false
	ContainHwMode *bool `json:"ContainHwMode,omitempty" xml:"ContainHwMode,omitempty"`
	// Alarm handling time.
	//
	// example:
	//
	// 2018-09-26 01:51:01
	DealTime *string `json:"DealTime,omitempty" xml:"DealTime,omitempty"`
	// Description.
	//
	// example:
	//
	// webshell
	Desc *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	// Event details information.
	EventDetails []*GetAlarmDetailByIdResponseBodyDataEventDetails `json:"EventDetails,omitempty" xml:"EventDetails,omitempty" type:"Repeated"`
	// Alarm level.
	//
	// example:
	//
	// suspicious
	EventLevel *string `json:"EventLevel,omitempty" xml:"EventLevel,omitempty"`
	// Primary key ID of the work order.
	//
	// example:
	//
	// 9772
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Affected asset.
	//
	// example:
	//
	// nginx
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// Public IP.
	//
	// example:
	//
	// 47.116.126.79
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// Private IP.
	//
	// example:
	//
	// 172.19.195.176
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// First occurrence time
	//
	// example:
	//
	// 2018-09-26 01:51:01
	OccurrenceTime *string `json:"OccurrenceTime,omitempty" xml:"OccurrenceTime,omitempty"`
	// Owner.
	//
	// example:
	//
	// 324546
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Disposal method.
	//
	// example:
	//
	// 192.168.XX.XX
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// Handling status.
	//
	// example:
	//
	// 要查询的告警事件状态。取值：
	//
	// 0：全部
	//
	// 1：待处理
	//
	// 2：已忽略
	//
	// 4：已确认
	//
	// 8：已标记为误报
	//
	// 16：处理中
	//
	// 32：处理完毕
	//
	// 64：已经过期
	//
	// 128：已经删除
	//
	// 512：自动拦截中
	//
	// 513：自动拦截完毕
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// ATT&CK tactic name.
	//
	// example:
	//
	// Malicious scripts-Malicious script code execution
	TacticDisplayName *string `json:"TacticDisplayName,omitempty" xml:"TacticDisplayName,omitempty"`
}

func (s GetAlarmDetailByIdResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAlarmDetailByIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAlarmDetailByIdResponseBodyData) GetAlarmEventType() *string {
	return s.AlarmEventType
}

func (s *GetAlarmDetailByIdResponseBodyData) GetAlarmEventTypeDisplay() *string {
	return s.AlarmEventTypeDisplay
}

func (s *GetAlarmDetailByIdResponseBodyData) GetAlarmId() *int64 {
	return s.AlarmId
}

func (s *GetAlarmDetailByIdResponseBodyData) GetAlarmName() *string {
	return s.AlarmName
}

func (s *GetAlarmDetailByIdResponseBodyData) GetAlarmSource() *string {
	return s.AlarmSource
}

func (s *GetAlarmDetailByIdResponseBodyData) GetAlarmTime() *string {
	return s.AlarmTime
}

func (s *GetAlarmDetailByIdResponseBodyData) GetAnalysisResult() *string {
	return s.AnalysisResult
}

func (s *GetAlarmDetailByIdResponseBodyData) GetContainHwMode() *bool {
	return s.ContainHwMode
}

func (s *GetAlarmDetailByIdResponseBodyData) GetDealTime() *string {
	return s.DealTime
}

func (s *GetAlarmDetailByIdResponseBodyData) GetDesc() *string {
	return s.Desc
}

func (s *GetAlarmDetailByIdResponseBodyData) GetEventDetails() []*GetAlarmDetailByIdResponseBodyDataEventDetails {
	return s.EventDetails
}

func (s *GetAlarmDetailByIdResponseBodyData) GetEventLevel() *string {
	return s.EventLevel
}

func (s *GetAlarmDetailByIdResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *GetAlarmDetailByIdResponseBodyData) GetInstanceName() *string {
	return s.InstanceName
}

func (s *GetAlarmDetailByIdResponseBodyData) GetInternetIp() *string {
	return s.InternetIp
}

func (s *GetAlarmDetailByIdResponseBodyData) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *GetAlarmDetailByIdResponseBodyData) GetOccurrenceTime() *string {
	return s.OccurrenceTime
}

func (s *GetAlarmDetailByIdResponseBodyData) GetOwnerId() *string {
	return s.OwnerId
}

func (s *GetAlarmDetailByIdResponseBodyData) GetRemark() *string {
	return s.Remark
}

func (s *GetAlarmDetailByIdResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetAlarmDetailByIdResponseBodyData) GetTacticDisplayName() *string {
	return s.TacticDisplayName
}

func (s *GetAlarmDetailByIdResponseBodyData) SetAlarmEventType(v string) *GetAlarmDetailByIdResponseBodyData {
	s.AlarmEventType = &v
	return s
}

func (s *GetAlarmDetailByIdResponseBodyData) SetAlarmEventTypeDisplay(v string) *GetAlarmDetailByIdResponseBodyData {
	s.AlarmEventTypeDisplay = &v
	return s
}

func (s *GetAlarmDetailByIdResponseBodyData) SetAlarmId(v int64) *GetAlarmDetailByIdResponseBodyData {
	s.AlarmId = &v
	return s
}

func (s *GetAlarmDetailByIdResponseBodyData) SetAlarmName(v string) *GetAlarmDetailByIdResponseBodyData {
	s.AlarmName = &v
	return s
}

func (s *GetAlarmDetailByIdResponseBodyData) SetAlarmSource(v string) *GetAlarmDetailByIdResponseBodyData {
	s.AlarmSource = &v
	return s
}

func (s *GetAlarmDetailByIdResponseBodyData) SetAlarmTime(v string) *GetAlarmDetailByIdResponseBodyData {
	s.AlarmTime = &v
	return s
}

func (s *GetAlarmDetailByIdResponseBodyData) SetAnalysisResult(v string) *GetAlarmDetailByIdResponseBodyData {
	s.AnalysisResult = &v
	return s
}

func (s *GetAlarmDetailByIdResponseBodyData) SetContainHwMode(v bool) *GetAlarmDetailByIdResponseBodyData {
	s.ContainHwMode = &v
	return s
}

func (s *GetAlarmDetailByIdResponseBodyData) SetDealTime(v string) *GetAlarmDetailByIdResponseBodyData {
	s.DealTime = &v
	return s
}

func (s *GetAlarmDetailByIdResponseBodyData) SetDesc(v string) *GetAlarmDetailByIdResponseBodyData {
	s.Desc = &v
	return s
}

func (s *GetAlarmDetailByIdResponseBodyData) SetEventDetails(v []*GetAlarmDetailByIdResponseBodyDataEventDetails) *GetAlarmDetailByIdResponseBodyData {
	s.EventDetails = v
	return s
}

func (s *GetAlarmDetailByIdResponseBodyData) SetEventLevel(v string) *GetAlarmDetailByIdResponseBodyData {
	s.EventLevel = &v
	return s
}

func (s *GetAlarmDetailByIdResponseBodyData) SetId(v int64) *GetAlarmDetailByIdResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetAlarmDetailByIdResponseBodyData) SetInstanceName(v string) *GetAlarmDetailByIdResponseBodyData {
	s.InstanceName = &v
	return s
}

func (s *GetAlarmDetailByIdResponseBodyData) SetInternetIp(v string) *GetAlarmDetailByIdResponseBodyData {
	s.InternetIp = &v
	return s
}

func (s *GetAlarmDetailByIdResponseBodyData) SetIntranetIp(v string) *GetAlarmDetailByIdResponseBodyData {
	s.IntranetIp = &v
	return s
}

func (s *GetAlarmDetailByIdResponseBodyData) SetOccurrenceTime(v string) *GetAlarmDetailByIdResponseBodyData {
	s.OccurrenceTime = &v
	return s
}

func (s *GetAlarmDetailByIdResponseBodyData) SetOwnerId(v string) *GetAlarmDetailByIdResponseBodyData {
	s.OwnerId = &v
	return s
}

func (s *GetAlarmDetailByIdResponseBodyData) SetRemark(v string) *GetAlarmDetailByIdResponseBodyData {
	s.Remark = &v
	return s
}

func (s *GetAlarmDetailByIdResponseBodyData) SetStatus(v string) *GetAlarmDetailByIdResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetAlarmDetailByIdResponseBodyData) SetTacticDisplayName(v string) *GetAlarmDetailByIdResponseBodyData {
	s.TacticDisplayName = &v
	return s
}

func (s *GetAlarmDetailByIdResponseBodyData) Validate() error {
	if s.EventDetails != nil {
		for _, item := range s.EventDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAlarmDetailByIdResponseBodyDataEventDetails struct {
	// Alarm event display name.
	//
	// example:
	//
	// Login with unusual location
	NameDisplay *string `json:"NameDisplay,omitempty" xml:"NameDisplay,omitempty"`
	// Alarm event type.
	//
	// example:
	//
	// text
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// Path where the alarm event occurred.
	//
	// example:
	//
	// /etc/crontab
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// Path where the alarm event occurred.
	//
	// example:
	//
	// /etc/crontab
	ValueDisplay *string `json:"ValueDisplay,omitempty" xml:"ValueDisplay,omitempty"`
}

func (s GetAlarmDetailByIdResponseBodyDataEventDetails) String() string {
	return dara.Prettify(s)
}

func (s GetAlarmDetailByIdResponseBodyDataEventDetails) GoString() string {
	return s.String()
}

func (s *GetAlarmDetailByIdResponseBodyDataEventDetails) GetNameDisplay() *string {
	return s.NameDisplay
}

func (s *GetAlarmDetailByIdResponseBodyDataEventDetails) GetType() *string {
	return s.Type
}

func (s *GetAlarmDetailByIdResponseBodyDataEventDetails) GetValue() *string {
	return s.Value
}

func (s *GetAlarmDetailByIdResponseBodyDataEventDetails) GetValueDisplay() *string {
	return s.ValueDisplay
}

func (s *GetAlarmDetailByIdResponseBodyDataEventDetails) SetNameDisplay(v string) *GetAlarmDetailByIdResponseBodyDataEventDetails {
	s.NameDisplay = &v
	return s
}

func (s *GetAlarmDetailByIdResponseBodyDataEventDetails) SetType(v string) *GetAlarmDetailByIdResponseBodyDataEventDetails {
	s.Type = &v
	return s
}

func (s *GetAlarmDetailByIdResponseBodyDataEventDetails) SetValue(v string) *GetAlarmDetailByIdResponseBodyDataEventDetails {
	s.Value = &v
	return s
}

func (s *GetAlarmDetailByIdResponseBodyDataEventDetails) SetValueDisplay(v string) *GetAlarmDetailByIdResponseBodyDataEventDetails {
	s.ValueDisplay = &v
	return s
}

func (s *GetAlarmDetailByIdResponseBodyDataEventDetails) Validate() error {
	return dara.Validate(s)
}
