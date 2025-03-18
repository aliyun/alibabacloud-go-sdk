// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CreateScreenSettingRequest struct {
	// example:
	//
	// 123
	Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// false
	LogoPower *bool `json:"LogoPower,omitempty" xml:"LogoPower,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://img.alicdn.com/tfs/xxxx.png
	LogoUrl *string `json:"LogoUrl,omitempty" xml:"LogoUrl,omitempty"`
	// example:
	//
	// https://monitor.xxxxxxx
	MonitorUrl *string `json:"MonitorUrl,omitempty" xml:"MonitorUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// [{"positionId":1,"componentId":3,"date":"7-day"},{"positionId":2,"componentId":1,"date":"0-second"},{"positionId":3,"componentId":8,"date":"15-day"},{"positionId":4,"componentId":11,"date":"15-day"},{"positionId":5,"componentId":23,"date":"24-hour"},{"positionId":6,"componentId":17,"date":"24-hour"},{"positionId":7,"componentId":13,"date":"24-hour"},{"positionId":8,"componentId":25,"date":"0-second"}]
	ScreenDataMap *string `json:"ScreenDataMap,omitempty" xml:"ScreenDataMap,omitempty"`
	// example:
	//
	// 0
	ScreenDefault *int32 `json:"ScreenDefault,omitempty" xml:"ScreenDefault,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s CreateScreenSettingRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateScreenSettingRequest) GoString() string {
	return s.String()
}

func (s *CreateScreenSettingRequest) SetId(v int32) *CreateScreenSettingRequest {
	s.Id = &v
	return s
}

func (s *CreateScreenSettingRequest) SetLogoPower(v bool) *CreateScreenSettingRequest {
	s.LogoPower = &v
	return s
}

func (s *CreateScreenSettingRequest) SetLogoUrl(v string) *CreateScreenSettingRequest {
	s.LogoUrl = &v
	return s
}

func (s *CreateScreenSettingRequest) SetMonitorUrl(v string) *CreateScreenSettingRequest {
	s.MonitorUrl = &v
	return s
}

func (s *CreateScreenSettingRequest) SetScreenDataMap(v string) *CreateScreenSettingRequest {
	s.ScreenDataMap = &v
	return s
}

func (s *CreateScreenSettingRequest) SetScreenDefault(v int32) *CreateScreenSettingRequest {
	s.ScreenDefault = &v
	return s
}

func (s *CreateScreenSettingRequest) SetTitle(v string) *CreateScreenSettingRequest {
	s.Title = &v
	return s
}

type CreateScreenSettingResponseBody struct {
	// example:
	//
	// 123
	Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 898F7AA7-CECD-5EC7-AF4D-664C601B****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateScreenSettingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateScreenSettingResponseBody) GoString() string {
	return s.String()
}

func (s *CreateScreenSettingResponseBody) SetId(v int32) *CreateScreenSettingResponseBody {
	s.Id = &v
	return s
}

func (s *CreateScreenSettingResponseBody) SetRequestId(v string) *CreateScreenSettingResponseBody {
	s.RequestId = &v
	return s
}

type CreateScreenSettingResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateScreenSettingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateScreenSettingResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateScreenSettingResponse) GoString() string {
	return s.String()
}

func (s *CreateScreenSettingResponse) SetHeaders(v map[string]*string) *CreateScreenSettingResponse {
	s.Headers = v
	return s
}

func (s *CreateScreenSettingResponse) SetStatusCode(v int32) *CreateScreenSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateScreenSettingResponse) SetBody(v *CreateScreenSettingResponseBody) *CreateScreenSettingResponse {
	s.Body = v
	return s
}

type DeleteScreenSettingRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteScreenSettingRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteScreenSettingRequest) GoString() string {
	return s.String()
}

func (s *DeleteScreenSettingRequest) SetId(v int64) *DeleteScreenSettingRequest {
	s.Id = &v
	return s
}

type DeleteScreenSettingResponseBody struct {
	// example:
	//
	// CE500770-42D3-442E-9DDD-156E0F9F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteScreenSettingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteScreenSettingResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteScreenSettingResponseBody) SetRequestId(v string) *DeleteScreenSettingResponseBody {
	s.RequestId = &v
	return s
}

type DeleteScreenSettingResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteScreenSettingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteScreenSettingResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteScreenSettingResponse) GoString() string {
	return s.String()
}

func (s *DeleteScreenSettingResponse) SetHeaders(v map[string]*string) *DeleteScreenSettingResponse {
	s.Headers = v
	return s
}

func (s *DeleteScreenSettingResponse) SetStatusCode(v int32) *DeleteScreenSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteScreenSettingResponse) SetBody(v *DeleteScreenSettingResponseBody) *DeleteScreenSettingResponse {
	s.Body = v
	return s
}

type DescribeScreenAlarmEventListRequest struct {
	AlarmEventName *string `json:"AlarmEventName,omitempty" xml:"AlarmEventName,omitempty"`
	AlarmEventType *string `json:"AlarmEventType,omitempty" xml:"AlarmEventType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// Y
	Dealed *string `json:"Dealed,omitempty" xml:"Dealed,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sas
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// serious
	Levels *string `json:"Levels,omitempty" xml:"Levels,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 222.185.XX.XX
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// example:
	//
	// 1683862286000
	TimeEnd *string `json:"TimeEnd,omitempty" xml:"TimeEnd,omitempty"`
	// example:
	//
	// 1687104000000
	TimeStart *string `json:"TimeStart,omitempty" xml:"TimeStart,omitempty"`
}

func (s DescribeScreenAlarmEventListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeScreenAlarmEventListRequest) GoString() string {
	return s.String()
}

func (s *DescribeScreenAlarmEventListRequest) SetAlarmEventName(v string) *DescribeScreenAlarmEventListRequest {
	s.AlarmEventName = &v
	return s
}

func (s *DescribeScreenAlarmEventListRequest) SetAlarmEventType(v string) *DescribeScreenAlarmEventListRequest {
	s.AlarmEventType = &v
	return s
}

func (s *DescribeScreenAlarmEventListRequest) SetCurrentPage(v int32) *DescribeScreenAlarmEventListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeScreenAlarmEventListRequest) SetDealed(v string) *DescribeScreenAlarmEventListRequest {
	s.Dealed = &v
	return s
}

func (s *DescribeScreenAlarmEventListRequest) SetFrom(v string) *DescribeScreenAlarmEventListRequest {
	s.From = &v
	return s
}

func (s *DescribeScreenAlarmEventListRequest) SetLang(v string) *DescribeScreenAlarmEventListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeScreenAlarmEventListRequest) SetLevels(v string) *DescribeScreenAlarmEventListRequest {
	s.Levels = &v
	return s
}

func (s *DescribeScreenAlarmEventListRequest) SetPageSize(v string) *DescribeScreenAlarmEventListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeScreenAlarmEventListRequest) SetRemark(v string) *DescribeScreenAlarmEventListRequest {
	s.Remark = &v
	return s
}

func (s *DescribeScreenAlarmEventListRequest) SetTimeEnd(v string) *DescribeScreenAlarmEventListRequest {
	s.TimeEnd = &v
	return s
}

func (s *DescribeScreenAlarmEventListRequest) SetTimeStart(v string) *DescribeScreenAlarmEventListRequest {
	s.TimeStart = &v
	return s
}

type DescribeScreenAlarmEventListResponseBody struct {
	PageInfo *DescribeScreenAlarmEventListResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// example:
	//
	// 09969D2C-4FAD-429E-BFBF-9A60DEF8BF6F
	RequestId  *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SuspEvents []*DescribeScreenAlarmEventListResponseBodySuspEvents `json:"SuspEvents,omitempty" xml:"SuspEvents,omitempty" type:"Repeated"`
}

func (s DescribeScreenAlarmEventListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScreenAlarmEventListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScreenAlarmEventListResponseBody) SetPageInfo(v *DescribeScreenAlarmEventListResponseBodyPageInfo) *DescribeScreenAlarmEventListResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeScreenAlarmEventListResponseBody) SetRequestId(v string) *DescribeScreenAlarmEventListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScreenAlarmEventListResponseBody) SetSuspEvents(v []*DescribeScreenAlarmEventListResponseBodySuspEvents) *DescribeScreenAlarmEventListResponseBody {
	s.SuspEvents = v
	return s
}

type DescribeScreenAlarmEventListResponseBodyPageInfo struct {
	// example:
	//
	// 2
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeScreenAlarmEventListResponseBodyPageInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeScreenAlarmEventListResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeScreenAlarmEventListResponseBodyPageInfo) SetCount(v int32) *DescribeScreenAlarmEventListResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *DescribeScreenAlarmEventListResponseBodyPageInfo) SetCurrentPage(v int32) *DescribeScreenAlarmEventListResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeScreenAlarmEventListResponseBodyPageInfo) SetPageSize(v int32) *DescribeScreenAlarmEventListResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeScreenAlarmEventListResponseBodyPageInfo) SetTotalCount(v int32) *DescribeScreenAlarmEventListResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

type DescribeScreenAlarmEventListResponseBodySuspEvents struct {
	AlarmEventName *string `json:"AlarmEventName,omitempty" xml:"AlarmEventName,omitempty"`
	AlarmEventType *string `json:"AlarmEventType,omitempty" xml:"AlarmEventType,omitempty"`
	// example:
	//
	// 8df914418f4211fbf756efe7a6f4****
	AlarmUniqueInfo *string `json:"AlarmUniqueInfo,omitempty" xml:"AlarmUniqueInfo,omitempty"`
	// example:
	//
	// true
	CanBeDealOnLine *bool `json:"CanBeDealOnLine,omitempty" xml:"CanBeDealOnLine,omitempty"`
	// example:
	//
	// false
	CanCancelFault *bool `json:"CanCancelFault,omitempty" xml:"CanCancelFault,omitempty"`
	// example:
	//
	// sas
	DataSource *string `json:"DataSource,omitempty" xml:"DataSource,omitempty"`
	// example:
	//
	// false
	Dealed *bool `json:"Dealed,omitempty" xml:"Dealed,omitempty"`
	// example:
	//
	// {\\"Type\\": \\"text\\", \\"Value\\": u\\"\\u5efa\\u8bae\\u8fdb\\u884c\\u79c1\\u7f51\\u767d\\u540d\\u5355\\u914d\\u7f6e\\uff0c\\u786e\\u4fdd\\u8bbf\\u95ee\\u5b89\\u5168\\u3002\\"}
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 1543740301000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// fzerp-dev
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// 123.21.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// example:
	//
	// 100.100.XX.XX
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// example:
	//
	// serious
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// example:
	//
	// 1
	SaleVersion *string `json:"SaleVersion,omitempty" xml:"SaleVersion,omitempty"`
	// example:
	//
	// {\\"Type\\": \\"text\\", \\"Value\\": \\"Enter NAS console - monitoring and auditing - log analysis - log management - new log dump to create a log recording service for the file system.\\"}
	Solution *string `json:"Solution,omitempty" xml:"Solution,omitempty"`
	// example:
	//
	// 1543740301000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 1
	SuspiciousEventCount *int32 `json:"SuspiciousEventCount,omitempty" xml:"SuspiciousEventCount,omitempty"`
	// example:
	//
	// bf6b30d3-eea8-4924-9f0a-****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeScreenAlarmEventListResponseBodySuspEvents) String() string {
	return tea.Prettify(s)
}

func (s DescribeScreenAlarmEventListResponseBodySuspEvents) GoString() string {
	return s.String()
}

func (s *DescribeScreenAlarmEventListResponseBodySuspEvents) SetAlarmEventName(v string) *DescribeScreenAlarmEventListResponseBodySuspEvents {
	s.AlarmEventName = &v
	return s
}

func (s *DescribeScreenAlarmEventListResponseBodySuspEvents) SetAlarmEventType(v string) *DescribeScreenAlarmEventListResponseBodySuspEvents {
	s.AlarmEventType = &v
	return s
}

func (s *DescribeScreenAlarmEventListResponseBodySuspEvents) SetAlarmUniqueInfo(v string) *DescribeScreenAlarmEventListResponseBodySuspEvents {
	s.AlarmUniqueInfo = &v
	return s
}

func (s *DescribeScreenAlarmEventListResponseBodySuspEvents) SetCanBeDealOnLine(v bool) *DescribeScreenAlarmEventListResponseBodySuspEvents {
	s.CanBeDealOnLine = &v
	return s
}

func (s *DescribeScreenAlarmEventListResponseBodySuspEvents) SetCanCancelFault(v bool) *DescribeScreenAlarmEventListResponseBodySuspEvents {
	s.CanCancelFault = &v
	return s
}

func (s *DescribeScreenAlarmEventListResponseBodySuspEvents) SetDataSource(v string) *DescribeScreenAlarmEventListResponseBodySuspEvents {
	s.DataSource = &v
	return s
}

func (s *DescribeScreenAlarmEventListResponseBodySuspEvents) SetDealed(v bool) *DescribeScreenAlarmEventListResponseBodySuspEvents {
	s.Dealed = &v
	return s
}

func (s *DescribeScreenAlarmEventListResponseBodySuspEvents) SetDescription(v string) *DescribeScreenAlarmEventListResponseBodySuspEvents {
	s.Description = &v
	return s
}

func (s *DescribeScreenAlarmEventListResponseBodySuspEvents) SetEndTime(v int64) *DescribeScreenAlarmEventListResponseBodySuspEvents {
	s.EndTime = &v
	return s
}

func (s *DescribeScreenAlarmEventListResponseBodySuspEvents) SetInstanceName(v string) *DescribeScreenAlarmEventListResponseBodySuspEvents {
	s.InstanceName = &v
	return s
}

func (s *DescribeScreenAlarmEventListResponseBodySuspEvents) SetInternetIp(v string) *DescribeScreenAlarmEventListResponseBodySuspEvents {
	s.InternetIp = &v
	return s
}

func (s *DescribeScreenAlarmEventListResponseBodySuspEvents) SetIntranetIp(v string) *DescribeScreenAlarmEventListResponseBodySuspEvents {
	s.IntranetIp = &v
	return s
}

func (s *DescribeScreenAlarmEventListResponseBodySuspEvents) SetLevel(v string) *DescribeScreenAlarmEventListResponseBodySuspEvents {
	s.Level = &v
	return s
}

func (s *DescribeScreenAlarmEventListResponseBodySuspEvents) SetSaleVersion(v string) *DescribeScreenAlarmEventListResponseBodySuspEvents {
	s.SaleVersion = &v
	return s
}

func (s *DescribeScreenAlarmEventListResponseBodySuspEvents) SetSolution(v string) *DescribeScreenAlarmEventListResponseBodySuspEvents {
	s.Solution = &v
	return s
}

func (s *DescribeScreenAlarmEventListResponseBodySuspEvents) SetStartTime(v int64) *DescribeScreenAlarmEventListResponseBodySuspEvents {
	s.StartTime = &v
	return s
}

func (s *DescribeScreenAlarmEventListResponseBodySuspEvents) SetSuspiciousEventCount(v int32) *DescribeScreenAlarmEventListResponseBodySuspEvents {
	s.SuspiciousEventCount = &v
	return s
}

func (s *DescribeScreenAlarmEventListResponseBodySuspEvents) SetUuid(v string) *DescribeScreenAlarmEventListResponseBodySuspEvents {
	s.Uuid = &v
	return s
}

type DescribeScreenAlarmEventListResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeScreenAlarmEventListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeScreenAlarmEventListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScreenAlarmEventListResponse) GoString() string {
	return s.String()
}

func (s *DescribeScreenAlarmEventListResponse) SetHeaders(v map[string]*string) *DescribeScreenAlarmEventListResponse {
	s.Headers = v
	return s
}

func (s *DescribeScreenAlarmEventListResponse) SetStatusCode(v int32) *DescribeScreenAlarmEventListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScreenAlarmEventListResponse) SetBody(v *DescribeScreenAlarmEventListResponseBody) *DescribeScreenAlarmEventListResponse {
	s.Body = v
	return s
}

type DescribeScreenAttackAnalysisDataRequest struct {
	// example:
	//
	// true
	Base64 *string `json:"Base64,omitempty" xml:"Base64,omitempty"`
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// {"crack_type":"9"}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1668064495000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1644027670
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// DETAILS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeScreenAttackAnalysisDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeScreenAttackAnalysisDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeScreenAttackAnalysisDataRequest) SetBase64(v string) *DescribeScreenAttackAnalysisDataRequest {
	s.Base64 = &v
	return s
}

func (s *DescribeScreenAttackAnalysisDataRequest) SetCurrentPage(v int32) *DescribeScreenAttackAnalysisDataRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeScreenAttackAnalysisDataRequest) SetData(v string) *DescribeScreenAttackAnalysisDataRequest {
	s.Data = &v
	return s
}

func (s *DescribeScreenAttackAnalysisDataRequest) SetEndTime(v int64) *DescribeScreenAttackAnalysisDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeScreenAttackAnalysisDataRequest) SetPageSize(v int32) *DescribeScreenAttackAnalysisDataRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeScreenAttackAnalysisDataRequest) SetStartTime(v int64) *DescribeScreenAttackAnalysisDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeScreenAttackAnalysisDataRequest) SetType(v string) *DescribeScreenAttackAnalysisDataRequest {
	s.Type = &v
	return s
}

type DescribeScreenAttackAnalysisDataResponseBody struct {
	// example:
	//
	// [{\\"crack_hour\\":1662480000000,\\"crack_cnt\\":471},{\\"crack_hour\\":1662483600000,\\"crack_cnt\\":461},{\\"crack_hour\\":1662487200000,\\"crack_cnt\\":445},{\\"crack_hour\\":1662490800000,\\"crack_cnt\\":471},{\\"crack_hour\\":1662494400000,\\"crack_cnt\\":534},{\\"crack_hour\\":1662498000000,\\"crack_cnt\\":652},{\\"crack_hour\\":1662501600000,\\"crack_cnt\\":706},{\\"crack_hour\\":1662505200000,\\"crack_cnt\\":613},{\\"crack_hour\\":1662508800000,\\"crack_cnt\\":578},{\\"crack_hour\\":1662512400000,\\"crack_cnt\\":577},{\\"crack_hour\\":1662516000000,\\"crack_cnt\\":616},{\\"crack_hour\\":1662519600000,\\"crack_cnt\\":597},{\\"crack_hour\\":1662523200000,\\"crack_cnt\\":575},{\\"crack_hour\\":1662526800000,\\"crack_cnt\\":507}]
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 7532B7EE-7CE7-5F4D-BF04-Bxxxxxxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 11
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeScreenAttackAnalysisDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScreenAttackAnalysisDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScreenAttackAnalysisDataResponseBody) SetData(v string) *DescribeScreenAttackAnalysisDataResponseBody {
	s.Data = &v
	return s
}

func (s *DescribeScreenAttackAnalysisDataResponseBody) SetPage(v int32) *DescribeScreenAttackAnalysisDataResponseBody {
	s.Page = &v
	return s
}

func (s *DescribeScreenAttackAnalysisDataResponseBody) SetPageSize(v int32) *DescribeScreenAttackAnalysisDataResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeScreenAttackAnalysisDataResponseBody) SetRequestId(v string) *DescribeScreenAttackAnalysisDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScreenAttackAnalysisDataResponseBody) SetTotal(v int32) *DescribeScreenAttackAnalysisDataResponseBody {
	s.Total = &v
	return s
}

type DescribeScreenAttackAnalysisDataResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeScreenAttackAnalysisDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeScreenAttackAnalysisDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScreenAttackAnalysisDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeScreenAttackAnalysisDataResponse) SetHeaders(v map[string]*string) *DescribeScreenAttackAnalysisDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeScreenAttackAnalysisDataResponse) SetStatusCode(v int32) *DescribeScreenAttackAnalysisDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScreenAttackAnalysisDataResponse) SetBody(v *DescribeScreenAttackAnalysisDataResponseBody) *DescribeScreenAttackAnalysisDataResponse {
	s.Body = v
	return s
}

type DescribeScreenCloudHcRiskResponseBody struct {
	CloudHcRiskItems []*DescribeScreenCloudHcRiskResponseBodyCloudHcRiskItems `json:"CloudHcRiskItems,omitempty" xml:"CloudHcRiskItems,omitempty" type:"Repeated"`
	// example:
	//
	// 0C8487EF-50C2-54BB-8634-10F8C35D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeScreenCloudHcRiskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScreenCloudHcRiskResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScreenCloudHcRiskResponseBody) SetCloudHcRiskItems(v []*DescribeScreenCloudHcRiskResponseBodyCloudHcRiskItems) *DescribeScreenCloudHcRiskResponseBody {
	s.CloudHcRiskItems = v
	return s
}

func (s *DescribeScreenCloudHcRiskResponseBody) SetRequestId(v string) *DescribeScreenCloudHcRiskResponseBody {
	s.RequestId = &v
	return s
}

type DescribeScreenCloudHcRiskResponseBodyCloudHcRiskItems struct {
	// example:
	//
	// 5
	AffectCount *int32 `json:"AffectCount,omitempty" xml:"AffectCount,omitempty"`
	// example:
	//
	// OSS-PublicReadOpenManifestFileWithoutEncryption
	CheckItem *string `json:"CheckItem,omitempty" xml:"CheckItem,omitempty"`
	// example:
	//
	// HIGH
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	Pass  *bool   `json:"Pass,omitempty" xml:"Pass,omitempty"`
}

func (s DescribeScreenCloudHcRiskResponseBodyCloudHcRiskItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeScreenCloudHcRiskResponseBodyCloudHcRiskItems) GoString() string {
	return s.String()
}

func (s *DescribeScreenCloudHcRiskResponseBodyCloudHcRiskItems) SetAffectCount(v int32) *DescribeScreenCloudHcRiskResponseBodyCloudHcRiskItems {
	s.AffectCount = &v
	return s
}

func (s *DescribeScreenCloudHcRiskResponseBodyCloudHcRiskItems) SetCheckItem(v string) *DescribeScreenCloudHcRiskResponseBodyCloudHcRiskItems {
	s.CheckItem = &v
	return s
}

func (s *DescribeScreenCloudHcRiskResponseBodyCloudHcRiskItems) SetLevel(v string) *DescribeScreenCloudHcRiskResponseBodyCloudHcRiskItems {
	s.Level = &v
	return s
}

func (s *DescribeScreenCloudHcRiskResponseBodyCloudHcRiskItems) SetPass(v bool) *DescribeScreenCloudHcRiskResponseBodyCloudHcRiskItems {
	s.Pass = &v
	return s
}

type DescribeScreenCloudHcRiskResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeScreenCloudHcRiskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeScreenCloudHcRiskResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScreenCloudHcRiskResponse) GoString() string {
	return s.String()
}

func (s *DescribeScreenCloudHcRiskResponse) SetHeaders(v map[string]*string) *DescribeScreenCloudHcRiskResponse {
	s.Headers = v
	return s
}

func (s *DescribeScreenCloudHcRiskResponse) SetStatusCode(v int32) *DescribeScreenCloudHcRiskResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScreenCloudHcRiskResponse) SetBody(v *DescribeScreenCloudHcRiskResponseBody) *DescribeScreenCloudHcRiskResponse {
	s.Body = v
	return s
}

type DescribeScreenDataMapResponseBody struct {
	// example:
	//
	// 7532B7EE-7CE7-5F4D-BF04-XXXXXXXX
	RequestId         *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SasScreenTypeList []*DescribeScreenDataMapResponseBodySasScreenTypeList `json:"SasScreenTypeList,omitempty" xml:"SasScreenTypeList,omitempty" type:"Repeated"`
}

func (s DescribeScreenDataMapResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScreenDataMapResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScreenDataMapResponseBody) SetRequestId(v string) *DescribeScreenDataMapResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScreenDataMapResponseBody) SetSasScreenTypeList(v []*DescribeScreenDataMapResponseBodySasScreenTypeList) *DescribeScreenDataMapResponseBody {
	s.SasScreenTypeList = v
	return s
}

type DescribeScreenDataMapResponseBodySasScreenTypeList struct {
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// ASSETS
	TypeCode *string                                                       `json:"TypeCode,omitempty" xml:"TypeCode,omitempty"`
	TypeData []*DescribeScreenDataMapResponseBodySasScreenTypeListTypeData `json:"TypeData,omitempty" xml:"TypeData,omitempty" type:"Repeated"`
}

func (s DescribeScreenDataMapResponseBodySasScreenTypeList) String() string {
	return tea.Prettify(s)
}

func (s DescribeScreenDataMapResponseBodySasScreenTypeList) GoString() string {
	return s.String()
}

func (s *DescribeScreenDataMapResponseBodySasScreenTypeList) SetType(v string) *DescribeScreenDataMapResponseBodySasScreenTypeList {
	s.Type = &v
	return s
}

func (s *DescribeScreenDataMapResponseBodySasScreenTypeList) SetTypeCode(v string) *DescribeScreenDataMapResponseBodySasScreenTypeList {
	s.TypeCode = &v
	return s
}

func (s *DescribeScreenDataMapResponseBodySasScreenTypeList) SetTypeData(v []*DescribeScreenDataMapResponseBodySasScreenTypeListTypeData) *DescribeScreenDataMapResponseBodySasScreenTypeList {
	s.TypeData = v
	return s
}

type DescribeScreenDataMapResponseBodySasScreenTypeListTypeData struct {
	// example:
	//
	// VUL_VUL
	Code *string                                                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Date []*DescribeScreenDataMapResponseBodySasScreenTypeListTypeDataDate `json:"Date,omitempty" xml:"Date,omitempty" type:"Repeated"`
	// example:
	//
	// 25
	Id    *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s DescribeScreenDataMapResponseBodySasScreenTypeListTypeData) String() string {
	return tea.Prettify(s)
}

func (s DescribeScreenDataMapResponseBodySasScreenTypeListTypeData) GoString() string {
	return s.String()
}

func (s *DescribeScreenDataMapResponseBodySasScreenTypeListTypeData) SetCode(v string) *DescribeScreenDataMapResponseBodySasScreenTypeListTypeData {
	s.Code = &v
	return s
}

func (s *DescribeScreenDataMapResponseBodySasScreenTypeListTypeData) SetDate(v []*DescribeScreenDataMapResponseBodySasScreenTypeListTypeDataDate) *DescribeScreenDataMapResponseBodySasScreenTypeListTypeData {
	s.Date = v
	return s
}

func (s *DescribeScreenDataMapResponseBodySasScreenTypeListTypeData) SetId(v string) *DescribeScreenDataMapResponseBodySasScreenTypeListTypeData {
	s.Id = &v
	return s
}

func (s *DescribeScreenDataMapResponseBodySasScreenTypeListTypeData) SetTitle(v string) *DescribeScreenDataMapResponseBodySasScreenTypeListTypeData {
	s.Title = &v
	return s
}

type DescribeScreenDataMapResponseBodySasScreenTypeListTypeDataDate struct {
	// example:
	//
	// second
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// example:
	//
	// 1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeScreenDataMapResponseBodySasScreenTypeListTypeDataDate) String() string {
	return tea.Prettify(s)
}

func (s DescribeScreenDataMapResponseBodySasScreenTypeListTypeDataDate) GoString() string {
	return s.String()
}

func (s *DescribeScreenDataMapResponseBodySasScreenTypeListTypeDataDate) SetUnit(v string) *DescribeScreenDataMapResponseBodySasScreenTypeListTypeDataDate {
	s.Unit = &v
	return s
}

func (s *DescribeScreenDataMapResponseBodySasScreenTypeListTypeDataDate) SetValue(v string) *DescribeScreenDataMapResponseBodySasScreenTypeListTypeDataDate {
	s.Value = &v
	return s
}

type DescribeScreenDataMapResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeScreenDataMapResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeScreenDataMapResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScreenDataMapResponse) GoString() string {
	return s.String()
}

func (s *DescribeScreenDataMapResponse) SetHeaders(v map[string]*string) *DescribeScreenDataMapResponse {
	s.Headers = v
	return s
}

func (s *DescribeScreenDataMapResponse) SetStatusCode(v int32) *DescribeScreenDataMapResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScreenDataMapResponse) SetBody(v *DescribeScreenDataMapResponseBody) *DescribeScreenDataMapResponse {
	s.Body = v
	return s
}

type DescribeScreenEmerRiskResponseBody struct {
	CloudHcRiskItems []*DescribeScreenEmerRiskResponseBodyCloudHcRiskItems `json:"CloudHcRiskItems,omitempty" xml:"CloudHcRiskItems,omitempty" type:"Repeated"`
	// example:
	//
	// 23AD0BD2-8771-5647-819E-6xxxxxxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeScreenEmerRiskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScreenEmerRiskResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScreenEmerRiskResponseBody) SetCloudHcRiskItems(v []*DescribeScreenEmerRiskResponseBodyCloudHcRiskItems) *DescribeScreenEmerRiskResponseBody {
	s.CloudHcRiskItems = v
	return s
}

func (s *DescribeScreenEmerRiskResponseBody) SetRequestId(v string) *DescribeScreenEmerRiskResponseBody {
	s.RequestId = &v
	return s
}

type DescribeScreenEmerRiskResponseBodyCloudHcRiskItems struct {
	// example:
	//
	// 3
	AffectCount *int32 `json:"AffectCount,omitempty" xml:"AffectCount,omitempty"`
	// example:
	//
	// ASAP
	Level   *string `json:"Level,omitempty" xml:"Level,omitempty"`
	VulName *string `json:"VulName,omitempty" xml:"VulName,omitempty"`
}

func (s DescribeScreenEmerRiskResponseBodyCloudHcRiskItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeScreenEmerRiskResponseBodyCloudHcRiskItems) GoString() string {
	return s.String()
}

func (s *DescribeScreenEmerRiskResponseBodyCloudHcRiskItems) SetAffectCount(v int32) *DescribeScreenEmerRiskResponseBodyCloudHcRiskItems {
	s.AffectCount = &v
	return s
}

func (s *DescribeScreenEmerRiskResponseBodyCloudHcRiskItems) SetLevel(v string) *DescribeScreenEmerRiskResponseBodyCloudHcRiskItems {
	s.Level = &v
	return s
}

func (s *DescribeScreenEmerRiskResponseBodyCloudHcRiskItems) SetVulName(v string) *DescribeScreenEmerRiskResponseBodyCloudHcRiskItems {
	s.VulName = &v
	return s
}

type DescribeScreenEmerRiskResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeScreenEmerRiskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeScreenEmerRiskResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScreenEmerRiskResponse) GoString() string {
	return s.String()
}

func (s *DescribeScreenEmerRiskResponse) SetHeaders(v map[string]*string) *DescribeScreenEmerRiskResponse {
	s.Headers = v
	return s
}

func (s *DescribeScreenEmerRiskResponse) SetStatusCode(v int32) *DescribeScreenEmerRiskResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScreenEmerRiskResponse) SetBody(v *DescribeScreenEmerRiskResponseBody) *DescribeScreenEmerRiskResponse {
	s.Body = v
	return s
}

type DescribeScreenHostStatisticsResponseBody struct {
	Data *DescribeScreenHostStatisticsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// D65AADFC-1D20-5A6A-8F6A-9FA53C0DC1F8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeScreenHostStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScreenHostStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScreenHostStatisticsResponseBody) SetData(v *DescribeScreenHostStatisticsResponseBodyData) *DescribeScreenHostStatisticsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeScreenHostStatisticsResponseBody) SetRequestId(v string) *DescribeScreenHostStatisticsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeScreenHostStatisticsResponseBodyData struct {
	SafeCount             []*string `json:"SafeCount,omitempty" xml:"SafeCount,omitempty" type:"Repeated"`
	SuspEventMachineNames []*string `json:"SuspEventMachineNames,omitempty" xml:"SuspEventMachineNames,omitempty" type:"Repeated"`
	SuspEventUuids        []*string `json:"SuspEventUuids,omitempty" xml:"SuspEventUuids,omitempty" type:"Repeated"`
	WeaknessMachineNames  []*string `json:"WeaknessMachineNames,omitempty" xml:"WeaknessMachineNames,omitempty" type:"Repeated"`
	WeaknessUuids         []*string `json:"WeaknessUuids,omitempty" xml:"WeaknessUuids,omitempty" type:"Repeated"`
}

func (s DescribeScreenHostStatisticsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeScreenHostStatisticsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeScreenHostStatisticsResponseBodyData) SetSafeCount(v []*string) *DescribeScreenHostStatisticsResponseBodyData {
	s.SafeCount = v
	return s
}

func (s *DescribeScreenHostStatisticsResponseBodyData) SetSuspEventMachineNames(v []*string) *DescribeScreenHostStatisticsResponseBodyData {
	s.SuspEventMachineNames = v
	return s
}

func (s *DescribeScreenHostStatisticsResponseBodyData) SetSuspEventUuids(v []*string) *DescribeScreenHostStatisticsResponseBodyData {
	s.SuspEventUuids = v
	return s
}

func (s *DescribeScreenHostStatisticsResponseBodyData) SetWeaknessMachineNames(v []*string) *DescribeScreenHostStatisticsResponseBodyData {
	s.WeaknessMachineNames = v
	return s
}

func (s *DescribeScreenHostStatisticsResponseBodyData) SetWeaknessUuids(v []*string) *DescribeScreenHostStatisticsResponseBodyData {
	s.WeaknessUuids = v
	return s
}

type DescribeScreenHostStatisticsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeScreenHostStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeScreenHostStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScreenHostStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeScreenHostStatisticsResponse) SetHeaders(v map[string]*string) *DescribeScreenHostStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeScreenHostStatisticsResponse) SetStatusCode(v int32) *DescribeScreenHostStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScreenHostStatisticsResponse) SetBody(v *DescribeScreenHostStatisticsResponseBody) *DescribeScreenHostStatisticsResponse {
	s.Body = v
	return s
}

type DescribeScreenOperateInfoRequest struct {
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1634725571000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeScreenOperateInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeScreenOperateInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeScreenOperateInfoRequest) SetLang(v string) *DescribeScreenOperateInfoRequest {
	s.Lang = &v
	return s
}

func (s *DescribeScreenOperateInfoRequest) SetStartTime(v int64) *DescribeScreenOperateInfoRequest {
	s.StartTime = &v
	return s
}

type DescribeScreenOperateInfoResponseBody struct {
	DateArray []*string `json:"DateArray,omitempty" xml:"DateArray,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	HealthCheckDealedCount *int32    `json:"HealthCheckDealedCount,omitempty" xml:"HealthCheckDealedCount,omitempty"`
	HealthCheckValueArray  []*string `json:"HealthCheckValueArray,omitempty" xml:"HealthCheckValueArray,omitempty" type:"Repeated"`
	// example:
	//
	// 23AD0BD2-8771-5647-819E-6xxxxxxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	SecurityEventDealedCount *int32    `json:"SecurityEventDealedCount,omitempty" xml:"SecurityEventDealedCount,omitempty"`
	SuspEventValueArray      []*string `json:"SuspEventValueArray,omitempty" xml:"SuspEventValueArray,omitempty" type:"Repeated"`
	VulValueArray            []*string `json:"VulValueArray,omitempty" xml:"VulValueArray,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	VulnerabilityDealedCount *int32 `json:"VulnerabilityDealedCount,omitempty" xml:"VulnerabilityDealedCount,omitempty"`
}

func (s DescribeScreenOperateInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScreenOperateInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScreenOperateInfoResponseBody) SetDateArray(v []*string) *DescribeScreenOperateInfoResponseBody {
	s.DateArray = v
	return s
}

func (s *DescribeScreenOperateInfoResponseBody) SetHealthCheckDealedCount(v int32) *DescribeScreenOperateInfoResponseBody {
	s.HealthCheckDealedCount = &v
	return s
}

func (s *DescribeScreenOperateInfoResponseBody) SetHealthCheckValueArray(v []*string) *DescribeScreenOperateInfoResponseBody {
	s.HealthCheckValueArray = v
	return s
}

func (s *DescribeScreenOperateInfoResponseBody) SetRequestId(v string) *DescribeScreenOperateInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScreenOperateInfoResponseBody) SetSecurityEventDealedCount(v int32) *DescribeScreenOperateInfoResponseBody {
	s.SecurityEventDealedCount = &v
	return s
}

func (s *DescribeScreenOperateInfoResponseBody) SetSuspEventValueArray(v []*string) *DescribeScreenOperateInfoResponseBody {
	s.SuspEventValueArray = v
	return s
}

func (s *DescribeScreenOperateInfoResponseBody) SetVulValueArray(v []*string) *DescribeScreenOperateInfoResponseBody {
	s.VulValueArray = v
	return s
}

func (s *DescribeScreenOperateInfoResponseBody) SetVulnerabilityDealedCount(v int32) *DescribeScreenOperateInfoResponseBody {
	s.VulnerabilityDealedCount = &v
	return s
}

type DescribeScreenOperateInfoResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeScreenOperateInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeScreenOperateInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScreenOperateInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeScreenOperateInfoResponse) SetHeaders(v map[string]*string) *DescribeScreenOperateInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeScreenOperateInfoResponse) SetStatusCode(v int32) *DescribeScreenOperateInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScreenOperateInfoResponse) SetBody(v *DescribeScreenOperateInfoResponseBody) *DescribeScreenOperateInfoResponse {
	s.Body = v
	return s
}

type DescribeScreenOssUploadInfoResponseBody struct {
	// example:
	//
	// LTAI5txxxxxxx
	AccessId *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	// example:
	//
	// 1719919893
	Expire *int32 `json:"Expire,omitempty" xml:"Expire,omitempty"`
	// example:
	//
	// https://oss-cipxxxxxxxxxliyuncs.com
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// example:
	//
	// DegradePool_Offset_****
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// eyJleHBpcmF0aW9uIjoiMjAyNC0wOC0xNVQwOToxMTo1My40MDVaIiwiY29uZGl0aW9ucyI6W1siY29udGVudC1sZW5ndGgtcmFuZ2UiLDAsMTA0ODU3NjAwMF0sWyJzdGFydHMtd2l0aCIsIiRrZXkiLCJzY3JlZW5Mb2dvXC8xNzY2MTg1ODkxxxx
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// example:
	//
	// 30CBF632-109F-596F-97F2-451C8B2A****
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// example:
	//
	// wBiwkhd5LGcLzijtc3FhI****
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
}

func (s DescribeScreenOssUploadInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScreenOssUploadInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScreenOssUploadInfoResponseBody) SetAccessId(v string) *DescribeScreenOssUploadInfoResponseBody {
	s.AccessId = &v
	return s
}

func (s *DescribeScreenOssUploadInfoResponseBody) SetExpire(v int32) *DescribeScreenOssUploadInfoResponseBody {
	s.Expire = &v
	return s
}

func (s *DescribeScreenOssUploadInfoResponseBody) SetHost(v string) *DescribeScreenOssUploadInfoResponseBody {
	s.Host = &v
	return s
}

func (s *DescribeScreenOssUploadInfoResponseBody) SetKey(v string) *DescribeScreenOssUploadInfoResponseBody {
	s.Key = &v
	return s
}

func (s *DescribeScreenOssUploadInfoResponseBody) SetPolicy(v string) *DescribeScreenOssUploadInfoResponseBody {
	s.Policy = &v
	return s
}

func (s *DescribeScreenOssUploadInfoResponseBody) SetRequestId(v string) *DescribeScreenOssUploadInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScreenOssUploadInfoResponseBody) SetSecurityToken(v string) *DescribeScreenOssUploadInfoResponseBody {
	s.SecurityToken = &v
	return s
}

func (s *DescribeScreenOssUploadInfoResponseBody) SetSignature(v string) *DescribeScreenOssUploadInfoResponseBody {
	s.Signature = &v
	return s
}

type DescribeScreenOssUploadInfoResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeScreenOssUploadInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeScreenOssUploadInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScreenOssUploadInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeScreenOssUploadInfoResponse) SetHeaders(v map[string]*string) *DescribeScreenOssUploadInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeScreenOssUploadInfoResponse) SetStatusCode(v int32) *DescribeScreenOssUploadInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScreenOssUploadInfoResponse) SetBody(v *DescribeScreenOssUploadInfoResponseBody) *DescribeScreenOssUploadInfoResponse {
	s.Body = v
	return s
}

type DescribeScreenScoreThreadRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1723445464501
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1722840664501
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeScreenScoreThreadRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeScreenScoreThreadRequest) GoString() string {
	return s.String()
}

func (s *DescribeScreenScoreThreadRequest) SetEndTime(v int64) *DescribeScreenScoreThreadRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeScreenScoreThreadRequest) SetStartTime(v int64) *DescribeScreenScoreThreadRequest {
	s.StartTime = &v
	return s
}

type DescribeScreenScoreThreadResponseBody struct {
	Data *DescribeScreenScoreThreadResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// D03DD0FD-6041-5107-AC00-383E28F1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeScreenScoreThreadResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScreenScoreThreadResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScreenScoreThreadResponseBody) SetData(v *DescribeScreenScoreThreadResponseBodyData) *DescribeScreenScoreThreadResponseBody {
	s.Data = v
	return s
}

func (s *DescribeScreenScoreThreadResponseBody) SetRequestId(v string) *DescribeScreenScoreThreadResponseBody {
	s.RequestId = &v
	return s
}

type DescribeScreenScoreThreadResponseBodyData struct {
	SocreThread     []*string `json:"SocreThread,omitempty" xml:"SocreThread,omitempty" type:"Repeated"`
	SocreThreadDate []*string `json:"SocreThreadDate,omitempty" xml:"SocreThreadDate,omitempty" type:"Repeated"`
}

func (s DescribeScreenScoreThreadResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeScreenScoreThreadResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeScreenScoreThreadResponseBodyData) SetSocreThread(v []*string) *DescribeScreenScoreThreadResponseBodyData {
	s.SocreThread = v
	return s
}

func (s *DescribeScreenScoreThreadResponseBodyData) SetSocreThreadDate(v []*string) *DescribeScreenScoreThreadResponseBodyData {
	s.SocreThreadDate = v
	return s
}

type DescribeScreenScoreThreadResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeScreenScoreThreadResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeScreenScoreThreadResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScreenScoreThreadResponse) GoString() string {
	return s.String()
}

func (s *DescribeScreenScoreThreadResponse) SetHeaders(v map[string]*string) *DescribeScreenScoreThreadResponse {
	s.Headers = v
	return s
}

func (s *DescribeScreenScoreThreadResponse) SetStatusCode(v int32) *DescribeScreenScoreThreadResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScreenScoreThreadResponse) SetBody(v *DescribeScreenScoreThreadResponseBody) *DescribeScreenScoreThreadResponse {
	s.Body = v
	return s
}

type DescribeScreenSecurityStatInfoResponseBody struct {
	AttackEvent *DescribeScreenSecurityStatInfoResponseBodyAttackEvent `json:"AttackEvent,omitempty" xml:"AttackEvent,omitempty" type:"Struct"`
	HealthCheck *DescribeScreenSecurityStatInfoResponseBodyHealthCheck `json:"HealthCheck,omitempty" xml:"HealthCheck,omitempty" type:"Struct"`
	// example:
	//
	// F8B6F758-BCD4-597A-8A2C-DA5A552C****
	RequestId     *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SecurityEvent *DescribeScreenSecurityStatInfoResponseBodySecurityEvent `json:"SecurityEvent,omitempty" xml:"SecurityEvent,omitempty" type:"Struct"`
	Vulnerability *DescribeScreenSecurityStatInfoResponseBodyVulnerability `json:"Vulnerability,omitempty" xml:"Vulnerability,omitempty" type:"Struct"`
}

func (s DescribeScreenSecurityStatInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScreenSecurityStatInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScreenSecurityStatInfoResponseBody) SetAttackEvent(v *DescribeScreenSecurityStatInfoResponseBodyAttackEvent) *DescribeScreenSecurityStatInfoResponseBody {
	s.AttackEvent = v
	return s
}

func (s *DescribeScreenSecurityStatInfoResponseBody) SetHealthCheck(v *DescribeScreenSecurityStatInfoResponseBodyHealthCheck) *DescribeScreenSecurityStatInfoResponseBody {
	s.HealthCheck = v
	return s
}

func (s *DescribeScreenSecurityStatInfoResponseBody) SetRequestId(v string) *DescribeScreenSecurityStatInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScreenSecurityStatInfoResponseBody) SetSecurityEvent(v *DescribeScreenSecurityStatInfoResponseBodySecurityEvent) *DescribeScreenSecurityStatInfoResponseBody {
	s.SecurityEvent = v
	return s
}

func (s *DescribeScreenSecurityStatInfoResponseBody) SetVulnerability(v *DescribeScreenSecurityStatInfoResponseBodyVulnerability) *DescribeScreenSecurityStatInfoResponseBody {
	s.Vulnerability = v
	return s
}

type DescribeScreenSecurityStatInfoResponseBodyAttackEvent struct {
	DateArray []*string `json:"DateArray,omitempty" xml:"DateArray,omitempty" type:"Repeated"`
	// example:
	//
	// 1096
	TotalCount *int32    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	ValueArray []*string `json:"ValueArray,omitempty" xml:"ValueArray,omitempty" type:"Repeated"`
}

func (s DescribeScreenSecurityStatInfoResponseBodyAttackEvent) String() string {
	return tea.Prettify(s)
}

func (s DescribeScreenSecurityStatInfoResponseBodyAttackEvent) GoString() string {
	return s.String()
}

func (s *DescribeScreenSecurityStatInfoResponseBodyAttackEvent) SetDateArray(v []*string) *DescribeScreenSecurityStatInfoResponseBodyAttackEvent {
	s.DateArray = v
	return s
}

func (s *DescribeScreenSecurityStatInfoResponseBodyAttackEvent) SetTotalCount(v int32) *DescribeScreenSecurityStatInfoResponseBodyAttackEvent {
	s.TotalCount = &v
	return s
}

func (s *DescribeScreenSecurityStatInfoResponseBodyAttackEvent) SetValueArray(v []*string) *DescribeScreenSecurityStatInfoResponseBodyAttackEvent {
	s.ValueArray = v
	return s
}

type DescribeScreenSecurityStatInfoResponseBodyHealthCheck struct {
	DateArray []*string `json:"DateArray,omitempty" xml:"DateArray,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	HighCount *int32    `json:"HighCount,omitempty" xml:"HighCount,omitempty"`
	HighList  []*string `json:"HighList,omitempty" xml:"HighList,omitempty" type:"Repeated"`
	LevelsOn  []*string `json:"LevelsOn,omitempty" xml:"LevelsOn,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	LowCount *int32    `json:"LowCount,omitempty" xml:"LowCount,omitempty"`
	LowList  []*string `json:"LowList,omitempty" xml:"LowList,omitempty" type:"Repeated"`
	// example:
	//
	// 21
	MediumCount *int32    `json:"MediumCount,omitempty" xml:"MediumCount,omitempty"`
	MediumList  []*string `json:"MediumList,omitempty" xml:"MediumList,omitempty" type:"Repeated"`
	// example:
	//
	// 32
	TotalCount *int32    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	ValueArray []*string `json:"ValueArray,omitempty" xml:"ValueArray,omitempty" type:"Repeated"`
}

func (s DescribeScreenSecurityStatInfoResponseBodyHealthCheck) String() string {
	return tea.Prettify(s)
}

func (s DescribeScreenSecurityStatInfoResponseBodyHealthCheck) GoString() string {
	return s.String()
}

func (s *DescribeScreenSecurityStatInfoResponseBodyHealthCheck) SetDateArray(v []*string) *DescribeScreenSecurityStatInfoResponseBodyHealthCheck {
	s.DateArray = v
	return s
}

func (s *DescribeScreenSecurityStatInfoResponseBodyHealthCheck) SetHighCount(v int32) *DescribeScreenSecurityStatInfoResponseBodyHealthCheck {
	s.HighCount = &v
	return s
}

func (s *DescribeScreenSecurityStatInfoResponseBodyHealthCheck) SetHighList(v []*string) *DescribeScreenSecurityStatInfoResponseBodyHealthCheck {
	s.HighList = v
	return s
}

func (s *DescribeScreenSecurityStatInfoResponseBodyHealthCheck) SetLevelsOn(v []*string) *DescribeScreenSecurityStatInfoResponseBodyHealthCheck {
	s.LevelsOn = v
	return s
}

func (s *DescribeScreenSecurityStatInfoResponseBodyHealthCheck) SetLowCount(v int32) *DescribeScreenSecurityStatInfoResponseBodyHealthCheck {
	s.LowCount = &v
	return s
}

func (s *DescribeScreenSecurityStatInfoResponseBodyHealthCheck) SetLowList(v []*string) *DescribeScreenSecurityStatInfoResponseBodyHealthCheck {
	s.LowList = v
	return s
}

func (s *DescribeScreenSecurityStatInfoResponseBodyHealthCheck) SetMediumCount(v int32) *DescribeScreenSecurityStatInfoResponseBodyHealthCheck {
	s.MediumCount = &v
	return s
}

func (s *DescribeScreenSecurityStatInfoResponseBodyHealthCheck) SetMediumList(v []*string) *DescribeScreenSecurityStatInfoResponseBodyHealthCheck {
	s.MediumList = v
	return s
}

func (s *DescribeScreenSecurityStatInfoResponseBodyHealthCheck) SetTotalCount(v int32) *DescribeScreenSecurityStatInfoResponseBodyHealthCheck {
	s.TotalCount = &v
	return s
}

func (s *DescribeScreenSecurityStatInfoResponseBodyHealthCheck) SetValueArray(v []*string) *DescribeScreenSecurityStatInfoResponseBodyHealthCheck {
	s.ValueArray = v
	return s
}

type DescribeScreenSecurityStatInfoResponseBodySecurityEvent struct {
	DateArray []*string `json:"DateArray,omitempty" xml:"DateArray,omitempty" type:"Repeated"`
	LevelsOn  []*string `json:"LevelsOn,omitempty" xml:"LevelsOn,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	RemindCount *int32    `json:"RemindCount,omitempty" xml:"RemindCount,omitempty"`
	RemindList  []*string `json:"RemindList,omitempty" xml:"RemindList,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	SeriousCount *int32    `json:"SeriousCount,omitempty" xml:"SeriousCount,omitempty"`
	SeriousList  []*string `json:"SeriousList,omitempty" xml:"SeriousList,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	SuspiciousCount *int32    `json:"SuspiciousCount,omitempty" xml:"SuspiciousCount,omitempty"`
	SuspiciousList  []*string `json:"SuspiciousList,omitempty" xml:"SuspiciousList,omitempty" type:"Repeated"`
	// example:
	//
	// 8
	TotalCount *int32    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	ValueArray []*string `json:"ValueArray,omitempty" xml:"ValueArray,omitempty" type:"Repeated"`
}

func (s DescribeScreenSecurityStatInfoResponseBodySecurityEvent) String() string {
	return tea.Prettify(s)
}

func (s DescribeScreenSecurityStatInfoResponseBodySecurityEvent) GoString() string {
	return s.String()
}

func (s *DescribeScreenSecurityStatInfoResponseBodySecurityEvent) SetDateArray(v []*string) *DescribeScreenSecurityStatInfoResponseBodySecurityEvent {
	s.DateArray = v
	return s
}

func (s *DescribeScreenSecurityStatInfoResponseBodySecurityEvent) SetLevelsOn(v []*string) *DescribeScreenSecurityStatInfoResponseBodySecurityEvent {
	s.LevelsOn = v
	return s
}

func (s *DescribeScreenSecurityStatInfoResponseBodySecurityEvent) SetRemindCount(v int32) *DescribeScreenSecurityStatInfoResponseBodySecurityEvent {
	s.RemindCount = &v
	return s
}

func (s *DescribeScreenSecurityStatInfoResponseBodySecurityEvent) SetRemindList(v []*string) *DescribeScreenSecurityStatInfoResponseBodySecurityEvent {
	s.RemindList = v
	return s
}

func (s *DescribeScreenSecurityStatInfoResponseBodySecurityEvent) SetSeriousCount(v int32) *DescribeScreenSecurityStatInfoResponseBodySecurityEvent {
	s.SeriousCount = &v
	return s
}

func (s *DescribeScreenSecurityStatInfoResponseBodySecurityEvent) SetSeriousList(v []*string) *DescribeScreenSecurityStatInfoResponseBodySecurityEvent {
	s.SeriousList = v
	return s
}

func (s *DescribeScreenSecurityStatInfoResponseBodySecurityEvent) SetSuspiciousCount(v int32) *DescribeScreenSecurityStatInfoResponseBodySecurityEvent {
	s.SuspiciousCount = &v
	return s
}

func (s *DescribeScreenSecurityStatInfoResponseBodySecurityEvent) SetSuspiciousList(v []*string) *DescribeScreenSecurityStatInfoResponseBodySecurityEvent {
	s.SuspiciousList = v
	return s
}

func (s *DescribeScreenSecurityStatInfoResponseBodySecurityEvent) SetTotalCount(v int32) *DescribeScreenSecurityStatInfoResponseBodySecurityEvent {
	s.TotalCount = &v
	return s
}

func (s *DescribeScreenSecurityStatInfoResponseBodySecurityEvent) SetValueArray(v []*string) *DescribeScreenSecurityStatInfoResponseBodySecurityEvent {
	s.ValueArray = v
	return s
}

type DescribeScreenSecurityStatInfoResponseBodyVulnerability struct {
	// example:
	//
	// 109
	AsapCount *int32    `json:"AsapCount,omitempty" xml:"AsapCount,omitempty"`
	AsapList  []*string `json:"AsapList,omitempty" xml:"AsapList,omitempty" type:"Repeated"`
	DateArray []*string `json:"DateArray,omitempty" xml:"DateArray,omitempty" type:"Repeated"`
	// example:
	//
	// 275
	LaterCount *int32    `json:"LaterCount,omitempty" xml:"LaterCount,omitempty"`
	LaterList  []*string `json:"LaterList,omitempty" xml:"LaterList,omitempty" type:"Repeated"`
	LevelsOn   []*string `json:"LevelsOn,omitempty" xml:"LevelsOn,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	NntfCount *int32    `json:"NntfCount,omitempty" xml:"NntfCount,omitempty"`
	NntfList  []*string `json:"NntfList,omitempty" xml:"NntfList,omitempty" type:"Repeated"`
	// example:
	//
	// 384
	TotalCount *int32    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	ValueArray []*string `json:"ValueArray,omitempty" xml:"ValueArray,omitempty" type:"Repeated"`
}

func (s DescribeScreenSecurityStatInfoResponseBodyVulnerability) String() string {
	return tea.Prettify(s)
}

func (s DescribeScreenSecurityStatInfoResponseBodyVulnerability) GoString() string {
	return s.String()
}

func (s *DescribeScreenSecurityStatInfoResponseBodyVulnerability) SetAsapCount(v int32) *DescribeScreenSecurityStatInfoResponseBodyVulnerability {
	s.AsapCount = &v
	return s
}

func (s *DescribeScreenSecurityStatInfoResponseBodyVulnerability) SetAsapList(v []*string) *DescribeScreenSecurityStatInfoResponseBodyVulnerability {
	s.AsapList = v
	return s
}

func (s *DescribeScreenSecurityStatInfoResponseBodyVulnerability) SetDateArray(v []*string) *DescribeScreenSecurityStatInfoResponseBodyVulnerability {
	s.DateArray = v
	return s
}

func (s *DescribeScreenSecurityStatInfoResponseBodyVulnerability) SetLaterCount(v int32) *DescribeScreenSecurityStatInfoResponseBodyVulnerability {
	s.LaterCount = &v
	return s
}

func (s *DescribeScreenSecurityStatInfoResponseBodyVulnerability) SetLaterList(v []*string) *DescribeScreenSecurityStatInfoResponseBodyVulnerability {
	s.LaterList = v
	return s
}

func (s *DescribeScreenSecurityStatInfoResponseBodyVulnerability) SetLevelsOn(v []*string) *DescribeScreenSecurityStatInfoResponseBodyVulnerability {
	s.LevelsOn = v
	return s
}

func (s *DescribeScreenSecurityStatInfoResponseBodyVulnerability) SetNntfCount(v int32) *DescribeScreenSecurityStatInfoResponseBodyVulnerability {
	s.NntfCount = &v
	return s
}

func (s *DescribeScreenSecurityStatInfoResponseBodyVulnerability) SetNntfList(v []*string) *DescribeScreenSecurityStatInfoResponseBodyVulnerability {
	s.NntfList = v
	return s
}

func (s *DescribeScreenSecurityStatInfoResponseBodyVulnerability) SetTotalCount(v int32) *DescribeScreenSecurityStatInfoResponseBodyVulnerability {
	s.TotalCount = &v
	return s
}

func (s *DescribeScreenSecurityStatInfoResponseBodyVulnerability) SetValueArray(v []*string) *DescribeScreenSecurityStatInfoResponseBodyVulnerability {
	s.ValueArray = v
	return s
}

type DescribeScreenSecurityStatInfoResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeScreenSecurityStatInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeScreenSecurityStatInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScreenSecurityStatInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeScreenSecurityStatInfoResponse) SetHeaders(v map[string]*string) *DescribeScreenSecurityStatInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeScreenSecurityStatInfoResponse) SetStatusCode(v int32) *DescribeScreenSecurityStatInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScreenSecurityStatInfoResponse) SetBody(v *DescribeScreenSecurityStatInfoResponseBody) *DescribeScreenSecurityStatInfoResponse {
	s.Body = v
	return s
}

type DescribeScreenSettingRequest struct {
	// example:
	//
	// 101786
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeScreenSettingRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeScreenSettingRequest) GoString() string {
	return s.String()
}

func (s *DescribeScreenSettingRequest) SetId(v string) *DescribeScreenSettingRequest {
	s.Id = &v
	return s
}

type DescribeScreenSettingResponseBody struct {
	// example:
	//
	// false
	LogoPower *bool `json:"LogoPower,omitempty" xml:"LogoPower,omitempty"`
	// example:
	//
	// https://img.alicdn.XXXXXXXXXXX.jpg
	LogoUrl *string `json:"LogoUrl,omitempty" xml:"LogoUrl,omitempty"`
	// example:
	//
	// https://XXX.monitor.XXXXcom
	MonitorUrl *string `json:"MonitorUrl,omitempty" xml:"MonitorUrl,omitempty"`
	// example:
	//
	// B9A68671-BD84-55CD-807A-XXXXXXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// [{\\"positionId\\":XX,\\"componentId\\":XX,\\"date\\":\\"XXX\\"},{\\"positionId\\":X,\\"componentId\\":X,\\"date\\":\\"XXX\\"},{\\"positionId\\":X,\\"componentId\\":X,\\"date\\":\\"XX\\"},{\\"positionId\\":X,\\"componentId\\":XX,\\"date\\":\\"XXX\\"},{\\"positionId\\":X,\\"componentId\\":XX,\\"date\\":\\"XX\\"},{\\"positionId\\":X,\\"componentId\\":XX,\\"date\\":\\"XX\\"},{\\"positionId\\":X,\\"componentId\\":XX,\\"date\\":\\"XXX\\"},{\\"positionId\\":X,\\"componentId\\":,\\"date\\":\\"XXXX\\"}]
	ScreenDataMap *string `json:"ScreenDataMap,omitempty" xml:"ScreenDataMap,omitempty"`
	// example:
	//
	// 7849
	ScreenDefault *int32 `json:"ScreenDefault,omitempty" xml:"ScreenDefault,omitempty"`
	// example:
	//
	// 1004770
	ScreenId *int32 `json:"ScreenId,omitempty" xml:"ScreenId,omitempty"`
	// example:
	//
	// Daily Report
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s DescribeScreenSettingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScreenSettingResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScreenSettingResponseBody) SetLogoPower(v bool) *DescribeScreenSettingResponseBody {
	s.LogoPower = &v
	return s
}

func (s *DescribeScreenSettingResponseBody) SetLogoUrl(v string) *DescribeScreenSettingResponseBody {
	s.LogoUrl = &v
	return s
}

func (s *DescribeScreenSettingResponseBody) SetMonitorUrl(v string) *DescribeScreenSettingResponseBody {
	s.MonitorUrl = &v
	return s
}

func (s *DescribeScreenSettingResponseBody) SetRequestId(v string) *DescribeScreenSettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScreenSettingResponseBody) SetScreenDataMap(v string) *DescribeScreenSettingResponseBody {
	s.ScreenDataMap = &v
	return s
}

func (s *DescribeScreenSettingResponseBody) SetScreenDefault(v int32) *DescribeScreenSettingResponseBody {
	s.ScreenDefault = &v
	return s
}

func (s *DescribeScreenSettingResponseBody) SetScreenId(v int32) *DescribeScreenSettingResponseBody {
	s.ScreenId = &v
	return s
}

func (s *DescribeScreenSettingResponseBody) SetTitle(v string) *DescribeScreenSettingResponseBody {
	s.Title = &v
	return s
}

type DescribeScreenSettingResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeScreenSettingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeScreenSettingResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScreenSettingResponse) GoString() string {
	return s.String()
}

func (s *DescribeScreenSettingResponse) SetHeaders(v map[string]*string) *DescribeScreenSettingResponse {
	s.Headers = v
	return s
}

func (s *DescribeScreenSettingResponse) SetStatusCode(v int32) *DescribeScreenSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScreenSettingResponse) SetBody(v *DescribeScreenSettingResponseBody) *DescribeScreenSettingResponse {
	s.Body = v
	return s
}

type DescribeScreenSummaryInfoResponseBody struct {
	// example:
	//
	// 12
	AegisClientOfflineCount *int32 `json:"AegisClientOfflineCount,omitempty" xml:"AegisClientOfflineCount,omitempty"`
	// example:
	//
	// 127
	AegisClientOnlineCount *int32 `json:"AegisClientOnlineCount,omitempty" xml:"AegisClientOnlineCount,omitempty"`
	// example:
	//
	// 23AD0BD2-8771-5647-819E-XXXXXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 100
	SecurityScore *int32 `json:"SecurityScore,omitempty" xml:"SecurityScore,omitempty"`
}

func (s DescribeScreenSummaryInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScreenSummaryInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScreenSummaryInfoResponseBody) SetAegisClientOfflineCount(v int32) *DescribeScreenSummaryInfoResponseBody {
	s.AegisClientOfflineCount = &v
	return s
}

func (s *DescribeScreenSummaryInfoResponseBody) SetAegisClientOnlineCount(v int32) *DescribeScreenSummaryInfoResponseBody {
	s.AegisClientOnlineCount = &v
	return s
}

func (s *DescribeScreenSummaryInfoResponseBody) SetRequestId(v string) *DescribeScreenSummaryInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScreenSummaryInfoResponseBody) SetSecurityScore(v int32) *DescribeScreenSummaryInfoResponseBody {
	s.SecurityScore = &v
	return s
}

type DescribeScreenSummaryInfoResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeScreenSummaryInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeScreenSummaryInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScreenSummaryInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeScreenSummaryInfoResponse) SetHeaders(v map[string]*string) *DescribeScreenSummaryInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeScreenSummaryInfoResponse) SetStatusCode(v int32) *DescribeScreenSummaryInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScreenSummaryInfoResponse) SetBody(v *DescribeScreenSummaryInfoResponseBody) *DescribeScreenSummaryInfoResponse {
	s.Body = v
	return s
}

type DescribeScreenTitlesResponseBody struct {
	// example:
	//
	// 09969D2C-4FAD-429E-BFBF-XXXXXXXXXXX
	RequestId            *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SasScreenSettingList []*DescribeScreenTitlesResponseBodySasScreenSettingList `json:"SasScreenSettingList,omitempty" xml:"SasScreenSettingList,omitempty" type:"Repeated"`
}

func (s DescribeScreenTitlesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScreenTitlesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScreenTitlesResponseBody) SetRequestId(v string) *DescribeScreenTitlesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScreenTitlesResponseBody) SetSasScreenSettingList(v []*DescribeScreenTitlesResponseBodySasScreenSettingList) *DescribeScreenTitlesResponseBody {
	s.SasScreenSettingList = v
	return s
}

type DescribeScreenTitlesResponseBodySasScreenSettingList struct {
	// example:
	//
	// 3267
	ScreenID *int64 `json:"ScreenID,omitempty" xml:"ScreenID,omitempty"`
	// example:
	//
	// titlexxx
	ScreenTitle *string `json:"ScreenTitle,omitempty" xml:"ScreenTitle,omitempty"`
}

func (s DescribeScreenTitlesResponseBodySasScreenSettingList) String() string {
	return tea.Prettify(s)
}

func (s DescribeScreenTitlesResponseBodySasScreenSettingList) GoString() string {
	return s.String()
}

func (s *DescribeScreenTitlesResponseBodySasScreenSettingList) SetScreenID(v int64) *DescribeScreenTitlesResponseBodySasScreenSettingList {
	s.ScreenID = &v
	return s
}

func (s *DescribeScreenTitlesResponseBodySasScreenSettingList) SetScreenTitle(v string) *DescribeScreenTitlesResponseBodySasScreenSettingList {
	s.ScreenTitle = &v
	return s
}

type DescribeScreenTitlesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeScreenTitlesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeScreenTitlesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScreenTitlesResponse) GoString() string {
	return s.String()
}

func (s *DescribeScreenTitlesResponse) SetHeaders(v map[string]*string) *DescribeScreenTitlesResponse {
	s.Headers = v
	return s
}

func (s *DescribeScreenTitlesResponse) SetStatusCode(v int32) *DescribeScreenTitlesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScreenTitlesResponse) SetBody(v *DescribeScreenTitlesResponseBody) *DescribeScreenTitlesResponse {
	s.Body = v
	return s
}

type DescribeScreenUploadPictureRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// https://security-pic.oss-cn-hangzhou.aliyuncs.com/screenLogo/1766185894104675/c28bd4d2-c5c1-43f8-9ef5-de41d762xxxx
	LogoUrl *string `json:"LogoUrl,omitempty" xml:"LogoUrl,omitempty"`
}

func (s DescribeScreenUploadPictureRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeScreenUploadPictureRequest) GoString() string {
	return s.String()
}

func (s *DescribeScreenUploadPictureRequest) SetLogoUrl(v string) *DescribeScreenUploadPictureRequest {
	s.LogoUrl = &v
	return s
}

type DescribeScreenUploadPictureResponseBody struct {
	// example:
	//
	// D65AADFC-1D20-5A6A-8F6A-9FA53C0Dxxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// http://security-pic.oss-cn-hangzhou.aliyuncs.com/screenLogo/1766185894104675/c28bd4d2-c5c1-43f8-9ef5-de41d76218eb?Expires=1723720214&OSSAccessKeyId=LTAI4G1mgPbjvGQuiV1Xxxxx&Signature=4o3xxxx
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DescribeScreenUploadPictureResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScreenUploadPictureResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScreenUploadPictureResponseBody) SetRequestId(v string) *DescribeScreenUploadPictureResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScreenUploadPictureResponseBody) SetUrl(v string) *DescribeScreenUploadPictureResponseBody {
	s.Url = &v
	return s
}

type DescribeScreenUploadPictureResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeScreenUploadPictureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeScreenUploadPictureResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScreenUploadPictureResponse) GoString() string {
	return s.String()
}

func (s *DescribeScreenUploadPictureResponse) SetHeaders(v map[string]*string) *DescribeScreenUploadPictureResponse {
	s.Headers = v
	return s
}

func (s *DescribeScreenUploadPictureResponse) SetStatusCode(v int32) *DescribeScreenUploadPictureResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScreenUploadPictureResponse) SetBody(v *DescribeScreenUploadPictureResponseBody) *DescribeScreenUploadPictureResponse {
	s.Body = v
	return s
}

type DescribeScreenVersionConfigResponseBody struct {
	// example:
	//
	// 30
	AssetLevel *int32 `json:"AssetLevel,omitempty" xml:"AssetLevel,omitempty"`
	// example:
	//
	// sas-b5***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 0
	IsTrialVersion *int32 `json:"IsTrialVersion,omitempty" xml:"IsTrialVersion,omitempty"`
	// example:
	//
	// 1625846400000
	ReleaseTime *int64 `json:"ReleaseTime,omitempty" xml:"ReleaseTime,omitempty"`
	// example:
	//
	// CE500770-42D3-442E-9DDD-1XXXXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	SasLog *int32 `json:"SasLog,omitempty" xml:"SasLog,omitempty"`
	// example:
	//
	// 0
	SasScreen *int32 `json:"SasScreen,omitempty" xml:"SasScreen,omitempty"`
	// example:
	//
	// 3
	Version *int32 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeScreenVersionConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScreenVersionConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScreenVersionConfigResponseBody) SetAssetLevel(v int32) *DescribeScreenVersionConfigResponseBody {
	s.AssetLevel = &v
	return s
}

func (s *DescribeScreenVersionConfigResponseBody) SetInstanceId(v string) *DescribeScreenVersionConfigResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeScreenVersionConfigResponseBody) SetIsTrialVersion(v int32) *DescribeScreenVersionConfigResponseBody {
	s.IsTrialVersion = &v
	return s
}

func (s *DescribeScreenVersionConfigResponseBody) SetReleaseTime(v int64) *DescribeScreenVersionConfigResponseBody {
	s.ReleaseTime = &v
	return s
}

func (s *DescribeScreenVersionConfigResponseBody) SetRequestId(v string) *DescribeScreenVersionConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScreenVersionConfigResponseBody) SetSasLog(v int32) *DescribeScreenVersionConfigResponseBody {
	s.SasLog = &v
	return s
}

func (s *DescribeScreenVersionConfigResponseBody) SetSasScreen(v int32) *DescribeScreenVersionConfigResponseBody {
	s.SasScreen = &v
	return s
}

func (s *DescribeScreenVersionConfigResponseBody) SetVersion(v int32) *DescribeScreenVersionConfigResponseBody {
	s.Version = &v
	return s
}

type DescribeScreenVersionConfigResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeScreenVersionConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeScreenVersionConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScreenVersionConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeScreenVersionConfigResponse) SetHeaders(v map[string]*string) *DescribeScreenVersionConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeScreenVersionConfigResponse) SetStatusCode(v int32) *DescribeScreenVersionConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScreenVersionConfigResponse) SetBody(v *DescribeScreenVersionConfigResponseBody) *DescribeScreenVersionConfigResponse {
	s.Body = v
	return s
}

type GetFileDetectResultInnerRequest struct {
	DnaHashKeyList []*string `json:"DnaHashKeyList,omitempty" xml:"DnaHashKeyList,omitempty" type:"Repeated"`
	// This parameter is required.
	HashKeyList []*string `json:"HashKeyList,omitempty" xml:"HashKeyList,omitempty" type:"Repeated"`
	Level       *int32    `json:"Level,omitempty" xml:"Level,omitempty"`
	SourceIp    *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Type        *int32    `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetFileDetectResultInnerRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFileDetectResultInnerRequest) GoString() string {
	return s.String()
}

func (s *GetFileDetectResultInnerRequest) SetDnaHashKeyList(v []*string) *GetFileDetectResultInnerRequest {
	s.DnaHashKeyList = v
	return s
}

func (s *GetFileDetectResultInnerRequest) SetHashKeyList(v []*string) *GetFileDetectResultInnerRequest {
	s.HashKeyList = v
	return s
}

func (s *GetFileDetectResultInnerRequest) SetLevel(v int32) *GetFileDetectResultInnerRequest {
	s.Level = &v
	return s
}

func (s *GetFileDetectResultInnerRequest) SetSourceIp(v string) *GetFileDetectResultInnerRequest {
	s.SourceIp = &v
	return s
}

func (s *GetFileDetectResultInnerRequest) SetType(v int32) *GetFileDetectResultInnerRequest {
	s.Type = &v
	return s
}

type GetFileDetectResultInnerResponseBody struct {
	RequestId  *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultList []*GetFileDetectResultInnerResponseBodyResultList `json:"ResultList,omitempty" xml:"ResultList,omitempty" type:"Repeated"`
}

func (s GetFileDetectResultInnerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFileDetectResultInnerResponseBody) GoString() string {
	return s.String()
}

func (s *GetFileDetectResultInnerResponseBody) SetRequestId(v string) *GetFileDetectResultInnerResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFileDetectResultInnerResponseBody) SetResultList(v []*GetFileDetectResultInnerResponseBodyResultList) *GetFileDetectResultInnerResponseBody {
	s.ResultList = v
	return s
}

type GetFileDetectResultInnerResponseBodyResultList struct {
	Code       *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	Ext        *string `json:"Ext,omitempty" xml:"Ext,omitempty"`
	HashKey    *string `json:"HashKey,omitempty" xml:"HashKey,omitempty"`
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Result     *int32  `json:"Result,omitempty" xml:"Result,omitempty"`
	Score      *int32  `json:"Score,omitempty" xml:"Score,omitempty"`
	VirusType  *string `json:"VirusType,omitempty" xml:"VirusType,omitempty"`
}

func (s GetFileDetectResultInnerResponseBodyResultList) String() string {
	return tea.Prettify(s)
}

func (s GetFileDetectResultInnerResponseBodyResultList) GoString() string {
	return s.String()
}

func (s *GetFileDetectResultInnerResponseBodyResultList) SetCode(v string) *GetFileDetectResultInnerResponseBodyResultList {
	s.Code = &v
	return s
}

func (s *GetFileDetectResultInnerResponseBodyResultList) SetExpireTime(v string) *GetFileDetectResultInnerResponseBodyResultList {
	s.ExpireTime = &v
	return s
}

func (s *GetFileDetectResultInnerResponseBodyResultList) SetExt(v string) *GetFileDetectResultInnerResponseBodyResultList {
	s.Ext = &v
	return s
}

func (s *GetFileDetectResultInnerResponseBodyResultList) SetHashKey(v string) *GetFileDetectResultInnerResponseBodyResultList {
	s.HashKey = &v
	return s
}

func (s *GetFileDetectResultInnerResponseBodyResultList) SetMessage(v string) *GetFileDetectResultInnerResponseBodyResultList {
	s.Message = &v
	return s
}

func (s *GetFileDetectResultInnerResponseBodyResultList) SetResult(v int32) *GetFileDetectResultInnerResponseBodyResultList {
	s.Result = &v
	return s
}

func (s *GetFileDetectResultInnerResponseBodyResultList) SetScore(v int32) *GetFileDetectResultInnerResponseBodyResultList {
	s.Score = &v
	return s
}

func (s *GetFileDetectResultInnerResponseBodyResultList) SetVirusType(v string) *GetFileDetectResultInnerResponseBodyResultList {
	s.VirusType = &v
	return s
}

type GetFileDetectResultInnerResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFileDetectResultInnerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFileDetectResultInnerResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFileDetectResultInnerResponse) GoString() string {
	return s.String()
}

func (s *GetFileDetectResultInnerResponse) SetHeaders(v map[string]*string) *GetFileDetectResultInnerResponse {
	s.Headers = v
	return s
}

func (s *GetFileDetectResultInnerResponse) SetStatusCode(v int32) *GetFileDetectResultInnerResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFileDetectResultInnerResponse) SetBody(v *GetFileDetectResultInnerResponseBody) *GetFileDetectResultInnerResponse {
	s.Body = v
	return s
}

type ListGlobalUserConfigRequest struct {
	ModuleList []*string `json:"ModuleList,omitempty" xml:"ModuleList,omitempty" type:"Repeated"`
}

func (s ListGlobalUserConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ListGlobalUserConfigRequest) GoString() string {
	return s.String()
}

func (s *ListGlobalUserConfigRequest) SetModuleList(v []*string) *ListGlobalUserConfigRequest {
	s.ModuleList = v
	return s
}

type ListGlobalUserConfigShrinkRequest struct {
	ModuleListShrink *string `json:"ModuleList,omitempty" xml:"ModuleList,omitempty"`
}

func (s ListGlobalUserConfigShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListGlobalUserConfigShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListGlobalUserConfigShrinkRequest) SetModuleListShrink(v string) *ListGlobalUserConfigShrinkRequest {
	s.ModuleListShrink = &v
	return s
}

type ListGlobalUserConfigResponseBody struct {
	Data []*ListGlobalUserConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// D81DD78E-E006-5C65-A171-C8CB09XXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListGlobalUserConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListGlobalUserConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ListGlobalUserConfigResponseBody) SetData(v []*ListGlobalUserConfigResponseBodyData) *ListGlobalUserConfigResponseBody {
	s.Data = v
	return s
}

func (s *ListGlobalUserConfigResponseBody) SetRequestId(v string) *ListGlobalUserConfigResponseBody {
	s.RequestId = &v
	return s
}

type ListGlobalUserConfigResponseBodyData struct {
	// example:
	//
	// true
	GlobalConfigSwitch *bool `json:"GlobalConfigSwitch,omitempty" xml:"GlobalConfigSwitch,omitempty"`
	// example:
	//
	// ransomware_breaking
	ModuleName *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
}

func (s ListGlobalUserConfigResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListGlobalUserConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListGlobalUserConfigResponseBodyData) SetGlobalConfigSwitch(v bool) *ListGlobalUserConfigResponseBodyData {
	s.GlobalConfigSwitch = &v
	return s
}

func (s *ListGlobalUserConfigResponseBodyData) SetModuleName(v string) *ListGlobalUserConfigResponseBodyData {
	s.ModuleName = &v
	return s
}

type ListGlobalUserConfigResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGlobalUserConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGlobalUserConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ListGlobalUserConfigResponse) GoString() string {
	return s.String()
}

func (s *ListGlobalUserConfigResponse) SetHeaders(v map[string]*string) *ListGlobalUserConfigResponse {
	s.Headers = v
	return s
}

func (s *ListGlobalUserConfigResponse) SetStatusCode(v int32) *ListGlobalUserConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGlobalUserConfigResponse) SetBody(v *ListGlobalUserConfigResponseBody) *ListGlobalUserConfigResponse {
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
	client.SignatureAlgorithm = tea.String("v2")
	client.EndpointRule = tea.String("regional")
	client.EndpointMap = map[string]*string{
		"cn-qingdao":            tea.String("tds.cn-shanghai.aliyuncs.com"),
		"cn-beijing":            tea.String("tds.cn-shanghai.aliyuncs.com"),
		"cn-zhangjiakou":        tea.String("tds.cn-shanghai.aliyuncs.com"),
		"cn-huhehaote":          tea.String("tds.cn-shanghai.aliyuncs.com"),
		"cn-wulanchabu":         tea.String("tds.cn-shanghai.aliyuncs.com"),
		"cn-hangzhou":           tea.String("tds.cn-shanghai.aliyuncs.com"),
		"cn-shanghai":           tea.String("tds.cn-shanghai.aliyuncs.com"),
		"cn-nanjing":            tea.String("tds.cn-shanghai.aliyuncs.com"),
		"cn-fuzhou":             tea.String("tds.cn-shanghai.aliyuncs.com"),
		"cn-shenzhen":           tea.String("tds.cn-shanghai.aliyuncs.com"),
		"cn-heyuan":             tea.String("tds.cn-shanghai.aliyuncs.com"),
		"cn-guangzhou":          tea.String("tds.cn-shanghai.aliyuncs.com"),
		"ap-southeast-2":        tea.String("tds.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-6":        tea.String("tds.ap-southeast-1.aliyuncs.com"),
		"ap-northeast-2":        tea.String("tds.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-3":        tea.String("tds.ap-southeast-1.aliyuncs.com"),
		"ap-northeast-1":        tea.String("tds.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-7":        tea.String("tds.ap-southeast-1.aliyuncs.com"),
		"cn-chengdu":            tea.String("tds.cn-shanghai.aliyuncs.com"),
		"ap-southeast-1":        tea.String("tds.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-5":        tea.String("tds.ap-southeast-1.aliyuncs.com"),
		"cn-hongkong":           tea.String("tds.cn-shanghai.aliyuncs.com"),
		"eu-central-1":          tea.String("tds.ap-southeast-1.aliyuncs.com"),
		"us-east-1":             tea.String("tds.ap-southeast-1.aliyuncs.com"),
		"us-west-1":             tea.String("tds.ap-southeast-1.aliyuncs.com"),
		"eu-west-1":             tea.String("tds.ap-southeast-1.aliyuncs.com"),
		"me-east-1":             tea.String("tds.ap-southeast-1.aliyuncs.com"),
		"me-central-1":          tea.String("tds.ap-southeast-1.aliyuncs.com"),
		"ap-south-1":            tea.String("tds.ap-southeast-1.aliyuncs.com"),
		"cn-beijing-finance-1":  tea.String("tds.cn-shanghai.aliyuncs.com"),
		"cn-hangzhou-finance":   tea.String("tds.cn-shanghai.aliyuncs.com"),
		"cn-shanghai-finance-1": tea.String("tds.cn-shanghai.aliyuncs.com"),
		"cn-shenzhen-finance-1": tea.String("tds.cn-shanghai.aliyuncs.com"),
		"cn-heyuan-acdr-1":      tea.String("tds.cn-shanghai.aliyuncs.com"),
		"cn-north-2-gov-1":      tea.String("tds.cn-shanghai.aliyuncs.com"),
		"cn-qingdao-acdr-ut-1":  tea.String("tds.cn-shanghai.aliyuncs.com"),
		"cn-shanghai-mybk":      tea.String("tds.cn-shanghai.aliyuncs.com"),
		"cn-wuhan-lr":           tea.String("tds.cn-shanghai.aliyuncs.com"),
		"cn-zhengzhou-jva":      tea.String("tds.cn-shanghai.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("sas"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 新增或者修改用户大屏设置
//
// @param request - CreateScreenSettingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateScreenSettingResponse
func (client *Client) CreateScreenSettingWithOptions(request *CreateScreenSettingRequest, runtime *util.RuntimeOptions) (_result *CreateScreenSettingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.LogoPower)) {
		query["LogoPower"] = request.LogoPower
	}

	if !tea.BoolValue(util.IsUnset(request.LogoUrl)) {
		query["LogoUrl"] = request.LogoUrl
	}

	if !tea.BoolValue(util.IsUnset(request.MonitorUrl)) {
		query["MonitorUrl"] = request.MonitorUrl
	}

	if !tea.BoolValue(util.IsUnset(request.ScreenDataMap)) {
		query["ScreenDataMap"] = request.ScreenDataMap
	}

	if !tea.BoolValue(util.IsUnset(request.ScreenDefault)) {
		query["ScreenDefault"] = request.ScreenDefault
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		query["Title"] = request.Title
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateScreenSetting"),
		Version:     tea.String("2021-01-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateScreenSettingResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateScreenSettingResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 新增或者修改用户大屏设置
//
// @param request - CreateScreenSettingRequest
//
// @return CreateScreenSettingResponse
func (client *Client) CreateScreenSetting(request *CreateScreenSettingRequest) (_result *CreateScreenSettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateScreenSettingResponse{}
	_body, _err := client.CreateScreenSettingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除大屏设置
//
// @param request - DeleteScreenSettingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteScreenSettingResponse
func (client *Client) DeleteScreenSettingWithOptions(request *DeleteScreenSettingRequest, runtime *util.RuntimeOptions) (_result *DeleteScreenSettingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteScreenSetting"),
		Version:     tea.String("2021-01-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteScreenSettingResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteScreenSettingResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 删除大屏设置
//
// @param request - DeleteScreenSettingRequest
//
// @return DeleteScreenSettingResponse
func (client *Client) DeleteScreenSetting(request *DeleteScreenSettingRequest) (_result *DeleteScreenSettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteScreenSettingResponse{}
	_body, _err := client.DeleteScreenSettingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询安全大屏告警事件
//
// @param request - DescribeScreenAlarmEventListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeScreenAlarmEventListResponse
func (client *Client) DescribeScreenAlarmEventListWithOptions(request *DescribeScreenAlarmEventListRequest, runtime *util.RuntimeOptions) (_result *DescribeScreenAlarmEventListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlarmEventName)) {
		query["AlarmEventName"] = request.AlarmEventName
	}

	if !tea.BoolValue(util.IsUnset(request.AlarmEventType)) {
		query["AlarmEventType"] = request.AlarmEventType
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.Dealed)) {
		query["Dealed"] = request.Dealed
	}

	if !tea.BoolValue(util.IsUnset(request.From)) {
		query["From"] = request.From
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.Levels)) {
		query["Levels"] = request.Levels
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["Remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.TimeEnd)) {
		query["TimeEnd"] = request.TimeEnd
	}

	if !tea.BoolValue(util.IsUnset(request.TimeStart)) {
		query["TimeStart"] = request.TimeStart
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeScreenAlarmEventList"),
		Version:     tea.String("2021-01-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeScreenAlarmEventListResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeScreenAlarmEventListResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 查询安全大屏告警事件
//
// @param request - DescribeScreenAlarmEventListRequest
//
// @return DescribeScreenAlarmEventListResponse
func (client *Client) DescribeScreenAlarmEventList(request *DescribeScreenAlarmEventListRequest) (_result *DescribeScreenAlarmEventListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScreenAlarmEventListResponse{}
	_body, _err := client.DescribeScreenAlarmEventListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询大屏攻击防御事件
//
// @param request - DescribeScreenAttackAnalysisDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeScreenAttackAnalysisDataResponse
func (client *Client) DescribeScreenAttackAnalysisDataWithOptions(request *DescribeScreenAttackAnalysisDataRequest, runtime *util.RuntimeOptions) (_result *DescribeScreenAttackAnalysisDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Base64)) {
		query["Base64"] = request.Base64
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.Data)) {
		query["Data"] = request.Data
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeScreenAttackAnalysisData"),
		Version:     tea.String("2021-01-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeScreenAttackAnalysisDataResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeScreenAttackAnalysisDataResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 查询大屏攻击防御事件
//
// @param request - DescribeScreenAttackAnalysisDataRequest
//
// @return DescribeScreenAttackAnalysisDataResponse
func (client *Client) DescribeScreenAttackAnalysisData(request *DescribeScreenAttackAnalysisDataRequest) (_result *DescribeScreenAttackAnalysisDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScreenAttackAnalysisDataResponse{}
	_body, _err := client.DescribeScreenAttackAnalysisDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询云产品基线问题
//
// @param request - DescribeScreenCloudHcRiskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeScreenCloudHcRiskResponse
func (client *Client) DescribeScreenCloudHcRiskWithOptions(runtime *util.RuntimeOptions) (_result *DescribeScreenCloudHcRiskResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("DescribeScreenCloudHcRisk"),
		Version:     tea.String("2021-01-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeScreenCloudHcRiskResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeScreenCloudHcRiskResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 查询云产品基线问题
//
// @return DescribeScreenCloudHcRiskResponse
func (client *Client) DescribeScreenCloudHcRisk() (_result *DescribeScreenCloudHcRiskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScreenCloudHcRiskResponse{}
	_body, _err := client.DescribeScreenCloudHcRiskWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取大屏可展示数据列表
//
// @param request - DescribeScreenDataMapRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeScreenDataMapResponse
func (client *Client) DescribeScreenDataMapWithOptions(runtime *util.RuntimeOptions) (_result *DescribeScreenDataMapResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("DescribeScreenDataMap"),
		Version:     tea.String("2021-01-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeScreenDataMapResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeScreenDataMapResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取大屏可展示数据列表
//
// @return DescribeScreenDataMapResponse
func (client *Client) DescribeScreenDataMap() (_result *DescribeScreenDataMapResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScreenDataMapResponse{}
	_body, _err := client.DescribeScreenDataMapWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询云产品漏洞风险
//
// @param request - DescribeScreenEmerRiskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeScreenEmerRiskResponse
func (client *Client) DescribeScreenEmerRiskWithOptions(runtime *util.RuntimeOptions) (_result *DescribeScreenEmerRiskResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("DescribeScreenEmerRisk"),
		Version:     tea.String("2021-01-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeScreenEmerRiskResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeScreenEmerRiskResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 查询云产品漏洞风险
//
// @return DescribeScreenEmerRiskResponse
func (client *Client) DescribeScreenEmerRisk() (_result *DescribeScreenEmerRiskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScreenEmerRiskResponse{}
	_body, _err := client.DescribeScreenEmerRiskWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询大屏主机统计数据
//
// @param request - DescribeScreenHostStatisticsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeScreenHostStatisticsResponse
func (client *Client) DescribeScreenHostStatisticsWithOptions(runtime *util.RuntimeOptions) (_result *DescribeScreenHostStatisticsResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("DescribeScreenHostStatistics"),
		Version:     tea.String("2021-01-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeScreenHostStatisticsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeScreenHostStatisticsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 查询大屏主机统计数据
//
// @return DescribeScreenHostStatisticsResponse
func (client *Client) DescribeScreenHostStatistics() (_result *DescribeScreenHostStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScreenHostStatisticsResponse{}
	_body, _err := client.DescribeScreenHostStatisticsWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查看运营信息
//
// @param request - DescribeScreenOperateInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeScreenOperateInfoResponse
func (client *Client) DescribeScreenOperateInfoWithOptions(request *DescribeScreenOperateInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeScreenOperateInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeScreenOperateInfo"),
		Version:     tea.String("2021-01-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeScreenOperateInfoResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeScreenOperateInfoResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 查看运营信息
//
// @param request - DescribeScreenOperateInfoRequest
//
// @return DescribeScreenOperateInfoResponse
func (client *Client) DescribeScreenOperateInfo(request *DescribeScreenOperateInfoRequest) (_result *DescribeScreenOperateInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScreenOperateInfoResponse{}
	_body, _err := client.DescribeScreenOperateInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询大屏上传信息
//
// @param request - DescribeScreenOssUploadInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeScreenOssUploadInfoResponse
func (client *Client) DescribeScreenOssUploadInfoWithOptions(runtime *util.RuntimeOptions) (_result *DescribeScreenOssUploadInfoResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("DescribeScreenOssUploadInfo"),
		Version:     tea.String("2021-01-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeScreenOssUploadInfoResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeScreenOssUploadInfoResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 查询大屏上传信息
//
// @return DescribeScreenOssUploadInfoResponse
func (client *Client) DescribeScreenOssUploadInfo() (_result *DescribeScreenOssUploadInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScreenOssUploadInfoResponse{}
	_body, _err := client.DescribeScreenOssUploadInfoWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询安全大屏分数趋势
//
// @param request - DescribeScreenScoreThreadRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeScreenScoreThreadResponse
func (client *Client) DescribeScreenScoreThreadWithOptions(request *DescribeScreenScoreThreadRequest, runtime *util.RuntimeOptions) (_result *DescribeScreenScoreThreadResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeScreenScoreThread"),
		Version:     tea.String("2021-01-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeScreenScoreThreadResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeScreenScoreThreadResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 查询安全大屏分数趋势
//
// @param request - DescribeScreenScoreThreadRequest
//
// @return DescribeScreenScoreThreadResponse
func (client *Client) DescribeScreenScoreThread(request *DescribeScreenScoreThreadRequest) (_result *DescribeScreenScoreThreadResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScreenScoreThreadResponse{}
	_body, _err := client.DescribeScreenScoreThreadWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询已处理的风险
//
// @param request - DescribeScreenSecurityStatInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeScreenSecurityStatInfoResponse
func (client *Client) DescribeScreenSecurityStatInfoWithOptions(runtime *util.RuntimeOptions) (_result *DescribeScreenSecurityStatInfoResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("DescribeScreenSecurityStatInfo"),
		Version:     tea.String("2021-01-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeScreenSecurityStatInfoResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeScreenSecurityStatInfoResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 查询已处理的风险
//
// @return DescribeScreenSecurityStatInfoResponse
func (client *Client) DescribeScreenSecurityStatInfo() (_result *DescribeScreenSecurityStatInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScreenSecurityStatInfoResponse{}
	_body, _err := client.DescribeScreenSecurityStatInfoWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询大屏配置
//
// @param request - DescribeScreenSettingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeScreenSettingResponse
func (client *Client) DescribeScreenSettingWithOptions(request *DescribeScreenSettingRequest, runtime *util.RuntimeOptions) (_result *DescribeScreenSettingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeScreenSetting"),
		Version:     tea.String("2021-01-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeScreenSettingResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeScreenSettingResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 查询大屏配置
//
// @param request - DescribeScreenSettingRequest
//
// @return DescribeScreenSettingResponse
func (client *Client) DescribeScreenSetting(request *DescribeScreenSettingRequest) (_result *DescribeScreenSettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScreenSettingResponse{}
	_body, _err := client.DescribeScreenSettingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询大屏统计信息
//
// @param request - DescribeScreenSummaryInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeScreenSummaryInfoResponse
func (client *Client) DescribeScreenSummaryInfoWithOptions(runtime *util.RuntimeOptions) (_result *DescribeScreenSummaryInfoResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("DescribeScreenSummaryInfo"),
		Version:     tea.String("2021-01-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeScreenSummaryInfoResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeScreenSummaryInfoResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 查询大屏统计信息
//
// @return DescribeScreenSummaryInfoResponse
func (client *Client) DescribeScreenSummaryInfo() (_result *DescribeScreenSummaryInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScreenSummaryInfoResponse{}
	_body, _err := client.DescribeScreenSummaryInfoWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取大屏幕设置全部列表
//
// @param request - DescribeScreenTitlesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeScreenTitlesResponse
func (client *Client) DescribeScreenTitlesWithOptions(runtime *util.RuntimeOptions) (_result *DescribeScreenTitlesResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("DescribeScreenTitles"),
		Version:     tea.String("2021-01-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeScreenTitlesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeScreenTitlesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取大屏幕设置全部列表
//
// @return DescribeScreenTitlesResponse
func (client *Client) DescribeScreenTitles() (_result *DescribeScreenTitlesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScreenTitlesResponse{}
	_body, _err := client.DescribeScreenTitlesWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询上传之后的图片显示地址
//
// @param request - DescribeScreenUploadPictureRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeScreenUploadPictureResponse
func (client *Client) DescribeScreenUploadPictureWithOptions(request *DescribeScreenUploadPictureRequest, runtime *util.RuntimeOptions) (_result *DescribeScreenUploadPictureResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LogoUrl)) {
		query["LogoUrl"] = request.LogoUrl
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeScreenUploadPicture"),
		Version:     tea.String("2021-01-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeScreenUploadPictureResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeScreenUploadPictureResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 查询上传之后的图片显示地址
//
// @param request - DescribeScreenUploadPictureRequest
//
// @return DescribeScreenUploadPictureResponse
func (client *Client) DescribeScreenUploadPicture(request *DescribeScreenUploadPictureRequest) (_result *DescribeScreenUploadPictureResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScreenUploadPictureResponse{}
	_body, _err := client.DescribeScreenUploadPictureWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询安全大屏版本配置
//
// @param request - DescribeScreenVersionConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeScreenVersionConfigResponse
func (client *Client) DescribeScreenVersionConfigWithOptions(runtime *util.RuntimeOptions) (_result *DescribeScreenVersionConfigResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("DescribeScreenVersionConfig"),
		Version:     tea.String("2021-01-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeScreenVersionConfigResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeScreenVersionConfigResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 查询安全大屏版本配置
//
// @return DescribeScreenVersionConfigResponse
func (client *Client) DescribeScreenVersionConfig() (_result *DescribeScreenVersionConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScreenVersionConfigResponse{}
	_body, _err := client.DescribeScreenVersionConfigWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取文件检测结果。
//
// @param request - GetFileDetectResultInnerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFileDetectResultInnerResponse
func (client *Client) GetFileDetectResultInnerWithOptions(request *GetFileDetectResultInnerRequest, runtime *util.RuntimeOptions) (_result *GetFileDetectResultInnerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DnaHashKeyList)) {
		query["DnaHashKeyList"] = request.DnaHashKeyList
	}

	if !tea.BoolValue(util.IsUnset(request.HashKeyList)) {
		query["HashKeyList"] = request.HashKeyList
	}

	if !tea.BoolValue(util.IsUnset(request.Level)) {
		query["Level"] = request.Level
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetFileDetectResultInner"),
		Version:     tea.String("2021-01-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetFileDetectResultInnerResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetFileDetectResultInnerResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取文件检测结果。
//
// @param request - GetFileDetectResultInnerRequest
//
// @return GetFileDetectResultInnerResponse
func (client *Client) GetFileDetectResultInner(request *GetFileDetectResultInnerRequest) (_result *GetFileDetectResultInnerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetFileDetectResultInnerResponse{}
	_body, _err := client.GetFileDetectResultInnerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 首页配置情况汇总接口
//
// @param tmpReq - ListGlobalUserConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGlobalUserConfigResponse
func (client *Client) ListGlobalUserConfigWithOptions(tmpReq *ListGlobalUserConfigRequest, runtime *util.RuntimeOptions) (_result *ListGlobalUserConfigResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListGlobalUserConfigShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ModuleList)) {
		request.ModuleListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ModuleList, tea.String("ModuleList"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ModuleListShrink)) {
		query["ModuleList"] = request.ModuleListShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListGlobalUserConfig"),
		Version:     tea.String("2021-01-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListGlobalUserConfigResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListGlobalUserConfigResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 首页配置情况汇总接口
//
// @param request - ListGlobalUserConfigRequest
//
// @return ListGlobalUserConfigResponse
func (client *Client) ListGlobalUserConfig(request *ListGlobalUserConfigRequest) (_result *ListGlobalUserConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListGlobalUserConfigResponse{}
	_body, _err := client.ListGlobalUserConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
