// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScriptTrendingReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetScriptTrendingReportResponseBody
	GetCode() *string
	SetData(v []*GetScriptTrendingReportResponseBodyData) *GetScriptTrendingReportResponseBody
	GetData() []*GetScriptTrendingReportResponseBodyData
	SetHttpStatusCode(v int32) *GetScriptTrendingReportResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetScriptTrendingReportResponseBody
	GetMessage() *string
	SetParams(v []*string) *GetScriptTrendingReportResponseBody
	GetParams() []*string
	SetRequestId(v string) *GetScriptTrendingReportResponseBody
	GetRequestId() *string
}

type GetScriptTrendingReportResponseBody struct {
	// The service-specific error code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// An array of scenario template objects.
	Data []*GetScriptTrendingReportResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// get upload tool url success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// An array of error parameters.
	Params []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 9DB8BA95-4513-54B9-9C67-A28909CFB4AD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetScriptTrendingReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetScriptTrendingReportResponseBody) GoString() string {
	return s.String()
}

func (s *GetScriptTrendingReportResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetScriptTrendingReportResponseBody) GetData() []*GetScriptTrendingReportResponseBodyData {
	return s.Data
}

func (s *GetScriptTrendingReportResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetScriptTrendingReportResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetScriptTrendingReportResponseBody) GetParams() []*string {
	return s.Params
}

func (s *GetScriptTrendingReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetScriptTrendingReportResponseBody) SetCode(v string) *GetScriptTrendingReportResponseBody {
	s.Code = &v
	return s
}

func (s *GetScriptTrendingReportResponseBody) SetData(v []*GetScriptTrendingReportResponseBodyData) *GetScriptTrendingReportResponseBody {
	s.Data = v
	return s
}

func (s *GetScriptTrendingReportResponseBody) SetHttpStatusCode(v int32) *GetScriptTrendingReportResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetScriptTrendingReportResponseBody) SetMessage(v string) *GetScriptTrendingReportResponseBody {
	s.Message = &v
	return s
}

func (s *GetScriptTrendingReportResponseBody) SetParams(v []*string) *GetScriptTrendingReportResponseBody {
	s.Params = v
	return s
}

func (s *GetScriptTrendingReportResponseBody) SetRequestId(v string) *GetScriptTrendingReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetScriptTrendingReportResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetScriptTrendingReportResponseBodyData struct {
	// The number of handled calls.
	//
	// example:
	//
	// 7
	CallsHandled *int64 `json:"CallsHandled,omitempty" xml:"CallsHandled,omitempty"`
	// The number of offered calls.
	//
	// example:
	//
	// 10
	CallsOffered *int64 `json:"CallsOffered,omitempty" xml:"CallsOffered,omitempty"`
	// The number of rejected calls.
	//
	// example:
	//
	// 3
	CallsRejected *int64 `json:"CallsRejected,omitempty" xml:"CallsRejected,omitempty"`
	// The number of resolved calls, based on the large model\\"s evaluation.
	//
	// example:
	//
	// 3
	CallsResolved *int64 `json:"CallsResolved,omitempty" xml:"CallsResolved,omitempty"`
	// The configured concurrency limit.
	//
	// example:
	//
	// 13
	ConfiguredConcurrency *int64 `json:"ConfiguredConcurrency,omitempty" xml:"ConfiguredConcurrency,omitempty"`
	// A JSON string that provides a breakdown of label distribution.
	//
	// example:
	//
	// {
	//
	//   "label1":{
	//
	//      "label1_sub1":10,
	//
	//      "label1_sub2":2
	//
	// }
	//
	// }
	LabelBreakDown *string `json:"LabelBreakDown,omitempty" xml:"LabelBreakDown,omitempty"`
	// The maximum talk time, in milliseconds.
	//
	// example:
	//
	// 20
	MaxTalkTime *int64 `json:"MaxTalkTime,omitempty" xml:"MaxTalkTime,omitempty"`
	// A JSON string that provides a breakdown of rejection reasons.
	//
	// example:
	//
	// {
	//
	//   "ConcurrentLimitExceeded":3
	//
	// }
	RejectionBreakdown *string `json:"RejectionBreakdown,omitempty" xml:"RejectionBreakdown,omitempty"`
	// A JSON string that provides a breakdown of call disconnection reasons.
	//
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
	// The start time of the interval. The value is a Unix timestamp in milliseconds.
	//
	// example:
	//
	// 1774881208361
	StatsTime *int64 `json:"StatsTime,omitempty" xml:"StatsTime,omitempty"`
	// A JSON string that provides a breakdown of conversation turn distribution.
	//
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
	// This parameter is reserved for future use.
	//
	// example:
	//
	// 0
	TotalInputTokens *int64 `json:"TotalInputTokens,omitempty" xml:"TotalInputTokens,omitempty"`
	// This parameter is reserved for future use.
	//
	// example:
	//
	// 0
	TotalOutputTokens *int64 `json:"TotalOutputTokens,omitempty" xml:"TotalOutputTokens,omitempty"`
	// The total talk time, in milliseconds.
	//
	// example:
	//
	// 100
	TotalTalkTime *int64 `json:"TotalTalkTime,omitempty" xml:"TotalTalkTime,omitempty"`
	// This parameter is reserved for future use.
	//
	// example:
	//
	// 0
	TotalTokens *int64 `json:"TotalTokens,omitempty" xml:"TotalTokens,omitempty"`
	// The peak concurrency for the interval.
	//
	// example:
	//
	// 10
	UsedConcurrency *int64 `json:"UsedConcurrency,omitempty" xml:"UsedConcurrency,omitempty"`
}

func (s GetScriptTrendingReportResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetScriptTrendingReportResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetScriptTrendingReportResponseBodyData) GetCallsHandled() *int64 {
	return s.CallsHandled
}

func (s *GetScriptTrendingReportResponseBodyData) GetCallsOffered() *int64 {
	return s.CallsOffered
}

func (s *GetScriptTrendingReportResponseBodyData) GetCallsRejected() *int64 {
	return s.CallsRejected
}

func (s *GetScriptTrendingReportResponseBodyData) GetCallsResolved() *int64 {
	return s.CallsResolved
}

func (s *GetScriptTrendingReportResponseBodyData) GetConfiguredConcurrency() *int64 {
	return s.ConfiguredConcurrency
}

func (s *GetScriptTrendingReportResponseBodyData) GetLabelBreakDown() *string {
	return s.LabelBreakDown
}

func (s *GetScriptTrendingReportResponseBodyData) GetMaxTalkTime() *int64 {
	return s.MaxTalkTime
}

func (s *GetScriptTrendingReportResponseBodyData) GetRejectionBreakdown() *string {
	return s.RejectionBreakdown
}

func (s *GetScriptTrendingReportResponseBodyData) GetReleaseReasonBreakdown() *string {
	return s.ReleaseReasonBreakdown
}

func (s *GetScriptTrendingReportResponseBodyData) GetStatsTime() *int64 {
	return s.StatsTime
}

func (s *GetScriptTrendingReportResponseBodyData) GetTalkTurnsDistribution() *string {
	return s.TalkTurnsDistribution
}

func (s *GetScriptTrendingReportResponseBodyData) GetTotalInputTokens() *int64 {
	return s.TotalInputTokens
}

func (s *GetScriptTrendingReportResponseBodyData) GetTotalOutputTokens() *int64 {
	return s.TotalOutputTokens
}

func (s *GetScriptTrendingReportResponseBodyData) GetTotalTalkTime() *int64 {
	return s.TotalTalkTime
}

func (s *GetScriptTrendingReportResponseBodyData) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *GetScriptTrendingReportResponseBodyData) GetUsedConcurrency() *int64 {
	return s.UsedConcurrency
}

func (s *GetScriptTrendingReportResponseBodyData) SetCallsHandled(v int64) *GetScriptTrendingReportResponseBodyData {
	s.CallsHandled = &v
	return s
}

func (s *GetScriptTrendingReportResponseBodyData) SetCallsOffered(v int64) *GetScriptTrendingReportResponseBodyData {
	s.CallsOffered = &v
	return s
}

func (s *GetScriptTrendingReportResponseBodyData) SetCallsRejected(v int64) *GetScriptTrendingReportResponseBodyData {
	s.CallsRejected = &v
	return s
}

func (s *GetScriptTrendingReportResponseBodyData) SetCallsResolved(v int64) *GetScriptTrendingReportResponseBodyData {
	s.CallsResolved = &v
	return s
}

func (s *GetScriptTrendingReportResponseBodyData) SetConfiguredConcurrency(v int64) *GetScriptTrendingReportResponseBodyData {
	s.ConfiguredConcurrency = &v
	return s
}

func (s *GetScriptTrendingReportResponseBodyData) SetLabelBreakDown(v string) *GetScriptTrendingReportResponseBodyData {
	s.LabelBreakDown = &v
	return s
}

func (s *GetScriptTrendingReportResponseBodyData) SetMaxTalkTime(v int64) *GetScriptTrendingReportResponseBodyData {
	s.MaxTalkTime = &v
	return s
}

func (s *GetScriptTrendingReportResponseBodyData) SetRejectionBreakdown(v string) *GetScriptTrendingReportResponseBodyData {
	s.RejectionBreakdown = &v
	return s
}

func (s *GetScriptTrendingReportResponseBodyData) SetReleaseReasonBreakdown(v string) *GetScriptTrendingReportResponseBodyData {
	s.ReleaseReasonBreakdown = &v
	return s
}

func (s *GetScriptTrendingReportResponseBodyData) SetStatsTime(v int64) *GetScriptTrendingReportResponseBodyData {
	s.StatsTime = &v
	return s
}

func (s *GetScriptTrendingReportResponseBodyData) SetTalkTurnsDistribution(v string) *GetScriptTrendingReportResponseBodyData {
	s.TalkTurnsDistribution = &v
	return s
}

func (s *GetScriptTrendingReportResponseBodyData) SetTotalInputTokens(v int64) *GetScriptTrendingReportResponseBodyData {
	s.TotalInputTokens = &v
	return s
}

func (s *GetScriptTrendingReportResponseBodyData) SetTotalOutputTokens(v int64) *GetScriptTrendingReportResponseBodyData {
	s.TotalOutputTokens = &v
	return s
}

func (s *GetScriptTrendingReportResponseBodyData) SetTotalTalkTime(v int64) *GetScriptTrendingReportResponseBodyData {
	s.TotalTalkTime = &v
	return s
}

func (s *GetScriptTrendingReportResponseBodyData) SetTotalTokens(v int64) *GetScriptTrendingReportResponseBodyData {
	s.TotalTokens = &v
	return s
}

func (s *GetScriptTrendingReportResponseBodyData) SetUsedConcurrency(v int64) *GetScriptTrendingReportResponseBodyData {
	s.UsedConcurrency = &v
	return s
}

func (s *GetScriptTrendingReportResponseBodyData) Validate() error {
	return dara.Validate(s)
}
