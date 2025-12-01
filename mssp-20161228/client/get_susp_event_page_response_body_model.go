// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSuspEventPageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetSuspEventPageResponseBody
	GetCode() *string
	SetData(v []*GetSuspEventPageResponseBodyData) *GetSuspEventPageResponseBody
	GetData() []*GetSuspEventPageResponseBodyData
	SetHttpStatusCode(v int32) *GetSuspEventPageResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetSuspEventPageResponseBody
	GetMessage() *string
	SetPageInfo(v *GetSuspEventPageResponseBodyPageInfo) *GetSuspEventPageResponseBody
	GetPageInfo() *GetSuspEventPageResponseBodyPageInfo
	SetRequestId(v string) *GetSuspEventPageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetSuspEventPageResponseBody
	GetSuccess() *bool
}

type GetSuspEventPageResponseBody struct {
	// API response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Data returned by the interface.
	Data []*GetSuspEventPageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Prompt message of the returned result.
	//
	// example:
	//
	// system error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Pagination information.
	PageInfo *GetSuspEventPageResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// Request ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// AFA6F7B7-7C4B-58BB-B8FB-E0FFA4483561
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the call was successful.
	//
	// - **true**: The call was successful. - **false**: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetSuspEventPageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSuspEventPageResponseBody) GoString() string {
	return s.String()
}

func (s *GetSuspEventPageResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetSuspEventPageResponseBody) GetData() []*GetSuspEventPageResponseBodyData {
	return s.Data
}

func (s *GetSuspEventPageResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetSuspEventPageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetSuspEventPageResponseBody) GetPageInfo() *GetSuspEventPageResponseBodyPageInfo {
	return s.PageInfo
}

func (s *GetSuspEventPageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSuspEventPageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetSuspEventPageResponseBody) SetCode(v string) *GetSuspEventPageResponseBody {
	s.Code = &v
	return s
}

func (s *GetSuspEventPageResponseBody) SetData(v []*GetSuspEventPageResponseBodyData) *GetSuspEventPageResponseBody {
	s.Data = v
	return s
}

func (s *GetSuspEventPageResponseBody) SetHttpStatusCode(v int32) *GetSuspEventPageResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetSuspEventPageResponseBody) SetMessage(v string) *GetSuspEventPageResponseBody {
	s.Message = &v
	return s
}

func (s *GetSuspEventPageResponseBody) SetPageInfo(v *GetSuspEventPageResponseBodyPageInfo) *GetSuspEventPageResponseBody {
	s.PageInfo = v
	return s
}

func (s *GetSuspEventPageResponseBody) SetRequestId(v string) *GetSuspEventPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSuspEventPageResponseBody) SetSuccess(v bool) *GetSuspEventPageResponseBody {
	s.Success = &v
	return s
}

func (s *GetSuspEventPageResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PageInfo != nil {
		if err := s.PageInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSuspEventPageResponseBodyData struct {
	// Alarm event type.
	//
	// example:
	//
	// 精准防御
	AlarmEventType *string `json:"AlarmEventType,omitempty" xml:"AlarmEventType,omitempty"`
	// Alarm ID.
	//
	// example:
	//
	// 5b1eeebe4f22daa2b177298234214fa3
	AlarmId *int64 `json:"AlarmId,omitempty" xml:"AlarmId,omitempty"`
	// Alarm name.
	//
	// example:
	//
	// Web服务漏洞利用
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
	// 1722515522000
	AlarmTime *string `json:"AlarmTime,omitempty" xml:"AlarmTime,omitempty"`
	// Analysis process.
	//
	// example:
	//
	// [{"value":"服务器可能已被黑客攻击，存在恶意进程在运行。 分析过程：告警显示，服务端存在一个名为”dns.exe”的进程在访问”polling.burpcollaborator.net”，这是一个被黑名单列出的恶意域名。在正常情况下,”dns.exe”不应该单独存在于系统的路径下，并且也不应该访问这类恶意域名。因此，这个进程可能是黑客留下的恶意进程。","key":"结论"},{"value":"尽快对服务器进行全面扫描，清除恶意进程。同时，联系网络安全专家进行深入调查，以确定是否有其他潜在的安全威胁。","key":"处置建议"}]
	AnalysisResult *string `json:"AnalysisResult,omitempty" xml:"AnalysisResult,omitempty"`
	// Alarm handling time.
	//
	// example:
	//
	// 1732515522000
	DealTime *string `json:"DealTime,omitempty" xml:"DealTime,omitempty"`
	// Alarm level.
	//
	// example:
	//
	// suspicious
	EventLevel *string `json:"EventLevel,omitempty" xml:"EventLevel,omitempty"`
	// Ticket primary key id.
	//
	// example:
	//
	// 9947
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Affected asset.
	//
	// example:
	//
	// shells-azhou
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// Public IP address.
	//
	// example:
	//
	// 47.99.188.31
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// Private IP address.
	//
	// example:
	//
	// 172.16.109.130
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// First occurrence time.
	//
	// example:
	//
	// 该字段暂未使用，有问题请联系管理员
	OccurrenceTime *string `json:"OccurrenceTime,omitempty" xml:"OccurrenceTime,omitempty"`
	// Owner ID.
	//
	// example:
	//
	// 张三
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Disposal method.
	//
	// example:
	//
	// 处理完成
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// Handling status.
	//
	// example:
	//
	// 未处理
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetSuspEventPageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetSuspEventPageResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSuspEventPageResponseBodyData) GetAlarmEventType() *string {
	return s.AlarmEventType
}

func (s *GetSuspEventPageResponseBodyData) GetAlarmId() *int64 {
	return s.AlarmId
}

func (s *GetSuspEventPageResponseBodyData) GetAlarmName() *string {
	return s.AlarmName
}

func (s *GetSuspEventPageResponseBodyData) GetAlarmSource() *string {
	return s.AlarmSource
}

func (s *GetSuspEventPageResponseBodyData) GetAlarmTime() *string {
	return s.AlarmTime
}

func (s *GetSuspEventPageResponseBodyData) GetAnalysisResult() *string {
	return s.AnalysisResult
}

func (s *GetSuspEventPageResponseBodyData) GetDealTime() *string {
	return s.DealTime
}

func (s *GetSuspEventPageResponseBodyData) GetEventLevel() *string {
	return s.EventLevel
}

func (s *GetSuspEventPageResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *GetSuspEventPageResponseBodyData) GetInstanceName() *string {
	return s.InstanceName
}

func (s *GetSuspEventPageResponseBodyData) GetInternetIp() *string {
	return s.InternetIp
}

func (s *GetSuspEventPageResponseBodyData) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *GetSuspEventPageResponseBodyData) GetOccurrenceTime() *string {
	return s.OccurrenceTime
}

func (s *GetSuspEventPageResponseBodyData) GetOwnerId() *string {
	return s.OwnerId
}

func (s *GetSuspEventPageResponseBodyData) GetRemark() *string {
	return s.Remark
}

func (s *GetSuspEventPageResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetSuspEventPageResponseBodyData) SetAlarmEventType(v string) *GetSuspEventPageResponseBodyData {
	s.AlarmEventType = &v
	return s
}

func (s *GetSuspEventPageResponseBodyData) SetAlarmId(v int64) *GetSuspEventPageResponseBodyData {
	s.AlarmId = &v
	return s
}

func (s *GetSuspEventPageResponseBodyData) SetAlarmName(v string) *GetSuspEventPageResponseBodyData {
	s.AlarmName = &v
	return s
}

func (s *GetSuspEventPageResponseBodyData) SetAlarmSource(v string) *GetSuspEventPageResponseBodyData {
	s.AlarmSource = &v
	return s
}

func (s *GetSuspEventPageResponseBodyData) SetAlarmTime(v string) *GetSuspEventPageResponseBodyData {
	s.AlarmTime = &v
	return s
}

func (s *GetSuspEventPageResponseBodyData) SetAnalysisResult(v string) *GetSuspEventPageResponseBodyData {
	s.AnalysisResult = &v
	return s
}

func (s *GetSuspEventPageResponseBodyData) SetDealTime(v string) *GetSuspEventPageResponseBodyData {
	s.DealTime = &v
	return s
}

func (s *GetSuspEventPageResponseBodyData) SetEventLevel(v string) *GetSuspEventPageResponseBodyData {
	s.EventLevel = &v
	return s
}

func (s *GetSuspEventPageResponseBodyData) SetId(v int64) *GetSuspEventPageResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetSuspEventPageResponseBodyData) SetInstanceName(v string) *GetSuspEventPageResponseBodyData {
	s.InstanceName = &v
	return s
}

func (s *GetSuspEventPageResponseBodyData) SetInternetIp(v string) *GetSuspEventPageResponseBodyData {
	s.InternetIp = &v
	return s
}

func (s *GetSuspEventPageResponseBodyData) SetIntranetIp(v string) *GetSuspEventPageResponseBodyData {
	s.IntranetIp = &v
	return s
}

func (s *GetSuspEventPageResponseBodyData) SetOccurrenceTime(v string) *GetSuspEventPageResponseBodyData {
	s.OccurrenceTime = &v
	return s
}

func (s *GetSuspEventPageResponseBodyData) SetOwnerId(v string) *GetSuspEventPageResponseBodyData {
	s.OwnerId = &v
	return s
}

func (s *GetSuspEventPageResponseBodyData) SetRemark(v string) *GetSuspEventPageResponseBodyData {
	s.Remark = &v
	return s
}

func (s *GetSuspEventPageResponseBodyData) SetStatus(v string) *GetSuspEventPageResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetSuspEventPageResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetSuspEventPageResponseBodyPageInfo struct {
	// The current page number in pagination queries.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of items displayed per page in the returned data.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Total number of query results.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetSuspEventPageResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s GetSuspEventPageResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *GetSuspEventPageResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetSuspEventPageResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetSuspEventPageResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetSuspEventPageResponseBodyPageInfo) SetCurrentPage(v int32) *GetSuspEventPageResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *GetSuspEventPageResponseBodyPageInfo) SetPageSize(v int32) *GetSuspEventPageResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *GetSuspEventPageResponseBodyPageInfo) SetTotalCount(v int32) *GetSuspEventPageResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *GetSuspEventPageResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
