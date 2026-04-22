// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceTrendingReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetInstanceTrendingReportResponseBody
	GetCode() *string
	SetData(v []*GetInstanceTrendingReportResponseBodyData) *GetInstanceTrendingReportResponseBody
	GetData() []*GetInstanceTrendingReportResponseBodyData
	SetHttpStatusCode(v int32) *GetInstanceTrendingReportResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetInstanceTrendingReportResponseBody
	GetMessage() *string
	SetParams(v []*string) *GetInstanceTrendingReportResponseBody
	GetParams() []*string
	SetRequestId(v string) *GetInstanceTrendingReportResponseBody
	GetRequestId() *string
}

type GetInstanceTrendingReportResponseBody struct {
	// example:
	//
	// OK
	Code *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*GetInstanceTrendingReportResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Instance af81a389-91f0-4157-8d82-720edd02b66a
	//
	//  does not exist.
	Message *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params  []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 14C39896-AE6D-4643-9C9A-E0566B2C2DDD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetInstanceTrendingReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceTrendingReportResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceTrendingReportResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetInstanceTrendingReportResponseBody) GetData() []*GetInstanceTrendingReportResponseBodyData {
	return s.Data
}

func (s *GetInstanceTrendingReportResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetInstanceTrendingReportResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetInstanceTrendingReportResponseBody) GetParams() []*string {
	return s.Params
}

func (s *GetInstanceTrendingReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstanceTrendingReportResponseBody) SetCode(v string) *GetInstanceTrendingReportResponseBody {
	s.Code = &v
	return s
}

func (s *GetInstanceTrendingReportResponseBody) SetData(v []*GetInstanceTrendingReportResponseBodyData) *GetInstanceTrendingReportResponseBody {
	s.Data = v
	return s
}

func (s *GetInstanceTrendingReportResponseBody) SetHttpStatusCode(v int32) *GetInstanceTrendingReportResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetInstanceTrendingReportResponseBody) SetMessage(v string) *GetInstanceTrendingReportResponseBody {
	s.Message = &v
	return s
}

func (s *GetInstanceTrendingReportResponseBody) SetParams(v []*string) *GetInstanceTrendingReportResponseBody {
	s.Params = v
	return s
}

func (s *GetInstanceTrendingReportResponseBody) SetRequestId(v string) *GetInstanceTrendingReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceTrendingReportResponseBody) Validate() error {
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

type GetInstanceTrendingReportResponseBodyData struct {
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
	// 13
	ConfiguredConcurrency *int64 `json:"ConfiguredConcurrency,omitempty" xml:"ConfiguredConcurrency,omitempty"`
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
	LabelBreakdown *string `json:"LabelBreakdown,omitempty" xml:"LabelBreakdown,omitempty"`
	// example:
	//
	// 20
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
	// 1774881208361
	StatsTime *int64 `json:"StatsTime,omitempty" xml:"StatsTime,omitempty"`
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
	// 10
	UsedConcurrency *int64 `json:"UsedConcurrency,omitempty" xml:"UsedConcurrency,omitempty"`
}

func (s GetInstanceTrendingReportResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceTrendingReportResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetInstanceTrendingReportResponseBodyData) GetCallsHandled() *int64 {
	return s.CallsHandled
}

func (s *GetInstanceTrendingReportResponseBodyData) GetCallsOffered() *int64 {
	return s.CallsOffered
}

func (s *GetInstanceTrendingReportResponseBodyData) GetCallsRejected() *int64 {
	return s.CallsRejected
}

func (s *GetInstanceTrendingReportResponseBodyData) GetCallsResolved() *int64 {
	return s.CallsResolved
}

func (s *GetInstanceTrendingReportResponseBodyData) GetConfiguredConcurrency() *int64 {
	return s.ConfiguredConcurrency
}

func (s *GetInstanceTrendingReportResponseBodyData) GetLabelBreakdown() *string {
	return s.LabelBreakdown
}

func (s *GetInstanceTrendingReportResponseBodyData) GetMaxTalkTime() *int64 {
	return s.MaxTalkTime
}

func (s *GetInstanceTrendingReportResponseBodyData) GetRejectionBreakdown() *string {
	return s.RejectionBreakdown
}

func (s *GetInstanceTrendingReportResponseBodyData) GetReleaseReasonBreakdown() *string {
	return s.ReleaseReasonBreakdown
}

func (s *GetInstanceTrendingReportResponseBodyData) GetStatsTime() *int64 {
	return s.StatsTime
}

func (s *GetInstanceTrendingReportResponseBodyData) GetTalkTurnsDistribution() *string {
	return s.TalkTurnsDistribution
}

func (s *GetInstanceTrendingReportResponseBodyData) GetTotalInputTokens() *int64 {
	return s.TotalInputTokens
}

func (s *GetInstanceTrendingReportResponseBodyData) GetTotalOutputTokens() *int64 {
	return s.TotalOutputTokens
}

func (s *GetInstanceTrendingReportResponseBodyData) GetTotalTalkTime() *int64 {
	return s.TotalTalkTime
}

func (s *GetInstanceTrendingReportResponseBodyData) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *GetInstanceTrendingReportResponseBodyData) GetUsedConcurrency() *int64 {
	return s.UsedConcurrency
}

func (s *GetInstanceTrendingReportResponseBodyData) SetCallsHandled(v int64) *GetInstanceTrendingReportResponseBodyData {
	s.CallsHandled = &v
	return s
}

func (s *GetInstanceTrendingReportResponseBodyData) SetCallsOffered(v int64) *GetInstanceTrendingReportResponseBodyData {
	s.CallsOffered = &v
	return s
}

func (s *GetInstanceTrendingReportResponseBodyData) SetCallsRejected(v int64) *GetInstanceTrendingReportResponseBodyData {
	s.CallsRejected = &v
	return s
}

func (s *GetInstanceTrendingReportResponseBodyData) SetCallsResolved(v int64) *GetInstanceTrendingReportResponseBodyData {
	s.CallsResolved = &v
	return s
}

func (s *GetInstanceTrendingReportResponseBodyData) SetConfiguredConcurrency(v int64) *GetInstanceTrendingReportResponseBodyData {
	s.ConfiguredConcurrency = &v
	return s
}

func (s *GetInstanceTrendingReportResponseBodyData) SetLabelBreakdown(v string) *GetInstanceTrendingReportResponseBodyData {
	s.LabelBreakdown = &v
	return s
}

func (s *GetInstanceTrendingReportResponseBodyData) SetMaxTalkTime(v int64) *GetInstanceTrendingReportResponseBodyData {
	s.MaxTalkTime = &v
	return s
}

func (s *GetInstanceTrendingReportResponseBodyData) SetRejectionBreakdown(v string) *GetInstanceTrendingReportResponseBodyData {
	s.RejectionBreakdown = &v
	return s
}

func (s *GetInstanceTrendingReportResponseBodyData) SetReleaseReasonBreakdown(v string) *GetInstanceTrendingReportResponseBodyData {
	s.ReleaseReasonBreakdown = &v
	return s
}

func (s *GetInstanceTrendingReportResponseBodyData) SetStatsTime(v int64) *GetInstanceTrendingReportResponseBodyData {
	s.StatsTime = &v
	return s
}

func (s *GetInstanceTrendingReportResponseBodyData) SetTalkTurnsDistribution(v string) *GetInstanceTrendingReportResponseBodyData {
	s.TalkTurnsDistribution = &v
	return s
}

func (s *GetInstanceTrendingReportResponseBodyData) SetTotalInputTokens(v int64) *GetInstanceTrendingReportResponseBodyData {
	s.TotalInputTokens = &v
	return s
}

func (s *GetInstanceTrendingReportResponseBodyData) SetTotalOutputTokens(v int64) *GetInstanceTrendingReportResponseBodyData {
	s.TotalOutputTokens = &v
	return s
}

func (s *GetInstanceTrendingReportResponseBodyData) SetTotalTalkTime(v int64) *GetInstanceTrendingReportResponseBodyData {
	s.TotalTalkTime = &v
	return s
}

func (s *GetInstanceTrendingReportResponseBodyData) SetTotalTokens(v int64) *GetInstanceTrendingReportResponseBodyData {
	s.TotalTokens = &v
	return s
}

func (s *GetInstanceTrendingReportResponseBodyData) SetUsedConcurrency(v int64) *GetInstanceTrendingReportResponseBodyData {
	s.UsedConcurrency = &v
	return s
}

func (s *GetInstanceTrendingReportResponseBodyData) Validate() error {
	return dara.Validate(s)
}
