// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHistoricalScriptReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListHistoricalScriptReportResponseBody
	GetCode() *string
	SetData(v *ListHistoricalScriptReportResponseBodyData) *ListHistoricalScriptReportResponseBody
	GetData() *ListHistoricalScriptReportResponseBodyData
	SetHttpStatusCode(v int32) *ListHistoricalScriptReportResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListHistoricalScriptReportResponseBody
	GetMessage() *string
	SetParams(v []*string) *ListHistoricalScriptReportResponseBody
	GetParams() []*string
	SetRequestId(v string) *ListHistoricalScriptReportResponseBody
	GetRequestId() *string
}

type ListHistoricalScriptReportResponseBody struct {
	// example:
	//
	// OK
	Code *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListHistoricalScriptReportResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Success
	Message *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params  []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 9ab43460-c0b9-40e2-8447-48d82c97fc67
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListHistoricalScriptReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListHistoricalScriptReportResponseBody) GoString() string {
	return s.String()
}

func (s *ListHistoricalScriptReportResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListHistoricalScriptReportResponseBody) GetData() *ListHistoricalScriptReportResponseBodyData {
	return s.Data
}

func (s *ListHistoricalScriptReportResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListHistoricalScriptReportResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListHistoricalScriptReportResponseBody) GetParams() []*string {
	return s.Params
}

func (s *ListHistoricalScriptReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListHistoricalScriptReportResponseBody) SetCode(v string) *ListHistoricalScriptReportResponseBody {
	s.Code = &v
	return s
}

func (s *ListHistoricalScriptReportResponseBody) SetData(v *ListHistoricalScriptReportResponseBodyData) *ListHistoricalScriptReportResponseBody {
	s.Data = v
	return s
}

func (s *ListHistoricalScriptReportResponseBody) SetHttpStatusCode(v int32) *ListHistoricalScriptReportResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListHistoricalScriptReportResponseBody) SetMessage(v string) *ListHistoricalScriptReportResponseBody {
	s.Message = &v
	return s
}

func (s *ListHistoricalScriptReportResponseBody) SetParams(v []*string) *ListHistoricalScriptReportResponseBody {
	s.Params = v
	return s
}

func (s *ListHistoricalScriptReportResponseBody) SetRequestId(v string) *ListHistoricalScriptReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHistoricalScriptReportResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListHistoricalScriptReportResponseBodyData struct {
	List []*ListHistoricalScriptReportResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 20
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListHistoricalScriptReportResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListHistoricalScriptReportResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListHistoricalScriptReportResponseBodyData) GetList() []*ListHistoricalScriptReportResponseBodyDataList {
	return s.List
}

func (s *ListHistoricalScriptReportResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListHistoricalScriptReportResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListHistoricalScriptReportResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListHistoricalScriptReportResponseBodyData) SetList(v []*ListHistoricalScriptReportResponseBodyDataList) *ListHistoricalScriptReportResponseBodyData {
	s.List = v
	return s
}

func (s *ListHistoricalScriptReportResponseBodyData) SetPageNumber(v int32) *ListHistoricalScriptReportResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListHistoricalScriptReportResponseBodyData) SetPageSize(v int32) *ListHistoricalScriptReportResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListHistoricalScriptReportResponseBodyData) SetTotalCount(v int32) *ListHistoricalScriptReportResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListHistoricalScriptReportResponseBodyData) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListHistoricalScriptReportResponseBodyDataList struct {
	// example:
	//
	// 7
	CallsHandled *int64 `json:"CallsHandled,omitempty" xml:"CallsHandled,omitempty"`
	// example:
	//
	// 10
	CallsOffered *int64 `json:"CallsOffered,omitempty" xml:"CallsOffered,omitempty"`
	// example:
	//
	// 3
	CallsRejected *int64 `json:"CallsRejected,omitempty" xml:"CallsRejected,omitempty"`
	// example:
	//
	// 4
	CallsResolved *int64 `json:"CallsResolved,omitempty" xml:"CallsResolved,omitempty"`
	// example:
	//
	// 3
	ConfiguredConcurrency *int64 `json:"ConfiguredConcurrency,omitempty" xml:"ConfiguredConcurrency,omitempty"`
	// example:
	//
	// {
	//
	//   "1":2,
	//
	//   "10":3
	//
	// }
	LabelDistribution *string `json:"LabelDistribution,omitempty" xml:"LabelDistribution,omitempty"`
	// example:
	//
	// 30
	MaxTalkTime *int64 `json:"MaxTalkTime,omitempty" xml:"MaxTalkTime,omitempty"`
	// example:
	//
	// {
	//
	//   "ConcurrentLimitExceeded":3
	//
	// }
	RejectionBreakdown *string `json:"RejectionBreakdown,omitempty" xml:"RejectionBreakdown,omitempty"`
	// example:
	//
	// {
	//
	//   "SilenceTimeout":1,
	//
	//   "TurnsLimitExceeded":1,
	//
	//    "Exception":1,
	//
	//    "Transferred":1,
	//
	//    "Normal":1
	//
	// }
	ReleaseReasonBreakdown *string `json:"ReleaseReasonBreakdown,omitempty" xml:"ReleaseReasonBreakdown,omitempty"`
	// example:
	//
	// 891264b9-5883-4068-94a6-241ceb8d51e4
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// example:
	//
	// {
	//
	//   "1":2,
	//
	//   "10":3
	//
	// }
	TalkTurnsDistribution *string `json:"TalkTurnsDistribution,omitempty" xml:"TalkTurnsDistribution,omitempty"`
	// example:
	//
	// 0
	TotalInputTokens *int64 `json:"TotalInputTokens,omitempty" xml:"TotalInputTokens,omitempty"`
	// example:
	//
	// 0
	TotalOutputTokens *int64 `json:"TotalOutputTokens,omitempty" xml:"TotalOutputTokens,omitempty"`
	// example:
	//
	// 100
	TotalTalkTime *int64 `json:"TotalTalkTime,omitempty" xml:"TotalTalkTime,omitempty"`
	// example:
	//
	// 0
	TotalTokens *int64 `json:"TotalTokens,omitempty" xml:"TotalTokens,omitempty"`
	// example:
	//
	// 2
	UsedConcurrency *int64 `json:"UsedConcurrency,omitempty" xml:"UsedConcurrency,omitempty"`
}

func (s ListHistoricalScriptReportResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s ListHistoricalScriptReportResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListHistoricalScriptReportResponseBodyDataList) GetCallsHandled() *int64 {
	return s.CallsHandled
}

func (s *ListHistoricalScriptReportResponseBodyDataList) GetCallsOffered() *int64 {
	return s.CallsOffered
}

func (s *ListHistoricalScriptReportResponseBodyDataList) GetCallsRejected() *int64 {
	return s.CallsRejected
}

func (s *ListHistoricalScriptReportResponseBodyDataList) GetCallsResolved() *int64 {
	return s.CallsResolved
}

func (s *ListHistoricalScriptReportResponseBodyDataList) GetConfiguredConcurrency() *int64 {
	return s.ConfiguredConcurrency
}

func (s *ListHistoricalScriptReportResponseBodyDataList) GetLabelDistribution() *string {
	return s.LabelDistribution
}

func (s *ListHistoricalScriptReportResponseBodyDataList) GetMaxTalkTime() *int64 {
	return s.MaxTalkTime
}

func (s *ListHistoricalScriptReportResponseBodyDataList) GetRejectionBreakdown() *string {
	return s.RejectionBreakdown
}

func (s *ListHistoricalScriptReportResponseBodyDataList) GetReleaseReasonBreakdown() *string {
	return s.ReleaseReasonBreakdown
}

func (s *ListHistoricalScriptReportResponseBodyDataList) GetScriptId() *string {
	return s.ScriptId
}

func (s *ListHistoricalScriptReportResponseBodyDataList) GetTalkTurnsDistribution() *string {
	return s.TalkTurnsDistribution
}

func (s *ListHistoricalScriptReportResponseBodyDataList) GetTotalInputTokens() *int64 {
	return s.TotalInputTokens
}

func (s *ListHistoricalScriptReportResponseBodyDataList) GetTotalOutputTokens() *int64 {
	return s.TotalOutputTokens
}

func (s *ListHistoricalScriptReportResponseBodyDataList) GetTotalTalkTime() *int64 {
	return s.TotalTalkTime
}

func (s *ListHistoricalScriptReportResponseBodyDataList) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *ListHistoricalScriptReportResponseBodyDataList) GetUsedConcurrency() *int64 {
	return s.UsedConcurrency
}

func (s *ListHistoricalScriptReportResponseBodyDataList) SetCallsHandled(v int64) *ListHistoricalScriptReportResponseBodyDataList {
	s.CallsHandled = &v
	return s
}

func (s *ListHistoricalScriptReportResponseBodyDataList) SetCallsOffered(v int64) *ListHistoricalScriptReportResponseBodyDataList {
	s.CallsOffered = &v
	return s
}

func (s *ListHistoricalScriptReportResponseBodyDataList) SetCallsRejected(v int64) *ListHistoricalScriptReportResponseBodyDataList {
	s.CallsRejected = &v
	return s
}

func (s *ListHistoricalScriptReportResponseBodyDataList) SetCallsResolved(v int64) *ListHistoricalScriptReportResponseBodyDataList {
	s.CallsResolved = &v
	return s
}

func (s *ListHistoricalScriptReportResponseBodyDataList) SetConfiguredConcurrency(v int64) *ListHistoricalScriptReportResponseBodyDataList {
	s.ConfiguredConcurrency = &v
	return s
}

func (s *ListHistoricalScriptReportResponseBodyDataList) SetLabelDistribution(v string) *ListHistoricalScriptReportResponseBodyDataList {
	s.LabelDistribution = &v
	return s
}

func (s *ListHistoricalScriptReportResponseBodyDataList) SetMaxTalkTime(v int64) *ListHistoricalScriptReportResponseBodyDataList {
	s.MaxTalkTime = &v
	return s
}

func (s *ListHistoricalScriptReportResponseBodyDataList) SetRejectionBreakdown(v string) *ListHistoricalScriptReportResponseBodyDataList {
	s.RejectionBreakdown = &v
	return s
}

func (s *ListHistoricalScriptReportResponseBodyDataList) SetReleaseReasonBreakdown(v string) *ListHistoricalScriptReportResponseBodyDataList {
	s.ReleaseReasonBreakdown = &v
	return s
}

func (s *ListHistoricalScriptReportResponseBodyDataList) SetScriptId(v string) *ListHistoricalScriptReportResponseBodyDataList {
	s.ScriptId = &v
	return s
}

func (s *ListHistoricalScriptReportResponseBodyDataList) SetTalkTurnsDistribution(v string) *ListHistoricalScriptReportResponseBodyDataList {
	s.TalkTurnsDistribution = &v
	return s
}

func (s *ListHistoricalScriptReportResponseBodyDataList) SetTotalInputTokens(v int64) *ListHistoricalScriptReportResponseBodyDataList {
	s.TotalInputTokens = &v
	return s
}

func (s *ListHistoricalScriptReportResponseBodyDataList) SetTotalOutputTokens(v int64) *ListHistoricalScriptReportResponseBodyDataList {
	s.TotalOutputTokens = &v
	return s
}

func (s *ListHistoricalScriptReportResponseBodyDataList) SetTotalTalkTime(v int64) *ListHistoricalScriptReportResponseBodyDataList {
	s.TotalTalkTime = &v
	return s
}

func (s *ListHistoricalScriptReportResponseBodyDataList) SetTotalTokens(v int64) *ListHistoricalScriptReportResponseBodyDataList {
	s.TotalTokens = &v
	return s
}

func (s *ListHistoricalScriptReportResponseBodyDataList) SetUsedConcurrency(v int64) *ListHistoricalScriptReportResponseBodyDataList {
	s.UsedConcurrency = &v
	return s
}

func (s *ListHistoricalScriptReportResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
