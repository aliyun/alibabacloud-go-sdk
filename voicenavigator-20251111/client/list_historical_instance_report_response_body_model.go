// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHistoricalInstanceReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListHistoricalInstanceReportResponseBody
	GetCode() *string
	SetData(v *ListHistoricalInstanceReportResponseBodyData) *ListHistoricalInstanceReportResponseBody
	GetData() *ListHistoricalInstanceReportResponseBodyData
	SetHttpStatusCode(v int32) *ListHistoricalInstanceReportResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListHistoricalInstanceReportResponseBody
	GetMessage() *string
	SetParams(v []*string) *ListHistoricalInstanceReportResponseBody
	GetParams() []*string
	SetRequestId(v string) *ListHistoricalInstanceReportResponseBody
	GetRequestId() *string
}

type ListHistoricalInstanceReportResponseBody struct {
	// example:
	//
	// OK
	Code *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListHistoricalInstanceReportResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// SUCCESS
	Message *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params  []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 9DB8BA95-4513-54B9-9C67-A28909CFB4AD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListHistoricalInstanceReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListHistoricalInstanceReportResponseBody) GoString() string {
	return s.String()
}

func (s *ListHistoricalInstanceReportResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListHistoricalInstanceReportResponseBody) GetData() *ListHistoricalInstanceReportResponseBodyData {
	return s.Data
}

func (s *ListHistoricalInstanceReportResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListHistoricalInstanceReportResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListHistoricalInstanceReportResponseBody) GetParams() []*string {
	return s.Params
}

func (s *ListHistoricalInstanceReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListHistoricalInstanceReportResponseBody) SetCode(v string) *ListHistoricalInstanceReportResponseBody {
	s.Code = &v
	return s
}

func (s *ListHistoricalInstanceReportResponseBody) SetData(v *ListHistoricalInstanceReportResponseBodyData) *ListHistoricalInstanceReportResponseBody {
	s.Data = v
	return s
}

func (s *ListHistoricalInstanceReportResponseBody) SetHttpStatusCode(v int32) *ListHistoricalInstanceReportResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListHistoricalInstanceReportResponseBody) SetMessage(v string) *ListHistoricalInstanceReportResponseBody {
	s.Message = &v
	return s
}

func (s *ListHistoricalInstanceReportResponseBody) SetParams(v []*string) *ListHistoricalInstanceReportResponseBody {
	s.Params = v
	return s
}

func (s *ListHistoricalInstanceReportResponseBody) SetRequestId(v string) *ListHistoricalInstanceReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHistoricalInstanceReportResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListHistoricalInstanceReportResponseBodyData struct {
	List []*ListHistoricalInstanceReportResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
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

func (s ListHistoricalInstanceReportResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListHistoricalInstanceReportResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListHistoricalInstanceReportResponseBodyData) GetList() []*ListHistoricalInstanceReportResponseBodyDataList {
	return s.List
}

func (s *ListHistoricalInstanceReportResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListHistoricalInstanceReportResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListHistoricalInstanceReportResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListHistoricalInstanceReportResponseBodyData) SetList(v []*ListHistoricalInstanceReportResponseBodyDataList) *ListHistoricalInstanceReportResponseBodyData {
	s.List = v
	return s
}

func (s *ListHistoricalInstanceReportResponseBodyData) SetPageNumber(v int32) *ListHistoricalInstanceReportResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListHistoricalInstanceReportResponseBodyData) SetPageSize(v int32) *ListHistoricalInstanceReportResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListHistoricalInstanceReportResponseBodyData) SetTotalCount(v int32) *ListHistoricalInstanceReportResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListHistoricalInstanceReportResponseBodyData) Validate() error {
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

type ListHistoricalInstanceReportResponseBodyDataList struct {
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
	// 2
	ConfiguredConcurrency *int64 `json:"ConfiguredConcurrency,omitempty" xml:"ConfiguredConcurrency,omitempty"`
	// example:
	//
	// 38d5fb05092b469e86b6061ad7945437
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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
	// 11
	UsedConcurrency *int64 `json:"UsedConcurrency,omitempty" xml:"UsedConcurrency,omitempty"`
}

func (s ListHistoricalInstanceReportResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s ListHistoricalInstanceReportResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListHistoricalInstanceReportResponseBodyDataList) GetCallsHandled() *int64 {
	return s.CallsHandled
}

func (s *ListHistoricalInstanceReportResponseBodyDataList) GetCallsOffered() *int64 {
	return s.CallsOffered
}

func (s *ListHistoricalInstanceReportResponseBodyDataList) GetCallsRejected() *int64 {
	return s.CallsRejected
}

func (s *ListHistoricalInstanceReportResponseBodyDataList) GetCallsResolved() *int64 {
	return s.CallsResolved
}

func (s *ListHistoricalInstanceReportResponseBodyDataList) GetConfiguredConcurrency() *int64 {
	return s.ConfiguredConcurrency
}

func (s *ListHistoricalInstanceReportResponseBodyDataList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListHistoricalInstanceReportResponseBodyDataList) GetLabelBreakdown() *string {
	return s.LabelBreakdown
}

func (s *ListHistoricalInstanceReportResponseBodyDataList) GetMaxTalkTime() *int64 {
	return s.MaxTalkTime
}

func (s *ListHistoricalInstanceReportResponseBodyDataList) GetRejectionBreakdown() *string {
	return s.RejectionBreakdown
}

func (s *ListHistoricalInstanceReportResponseBodyDataList) GetReleaseReasonBreakdown() *string {
	return s.ReleaseReasonBreakdown
}

func (s *ListHistoricalInstanceReportResponseBodyDataList) GetTalkTurnsDistribution() *string {
	return s.TalkTurnsDistribution
}

func (s *ListHistoricalInstanceReportResponseBodyDataList) GetTotalInputTokens() *int64 {
	return s.TotalInputTokens
}

func (s *ListHistoricalInstanceReportResponseBodyDataList) GetTotalOutputTokens() *int64 {
	return s.TotalOutputTokens
}

func (s *ListHistoricalInstanceReportResponseBodyDataList) GetTotalTalkTime() *int64 {
	return s.TotalTalkTime
}

func (s *ListHistoricalInstanceReportResponseBodyDataList) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *ListHistoricalInstanceReportResponseBodyDataList) GetUsedConcurrency() *int64 {
	return s.UsedConcurrency
}

func (s *ListHistoricalInstanceReportResponseBodyDataList) SetCallsHandled(v int64) *ListHistoricalInstanceReportResponseBodyDataList {
	s.CallsHandled = &v
	return s
}

func (s *ListHistoricalInstanceReportResponseBodyDataList) SetCallsOffered(v int64) *ListHistoricalInstanceReportResponseBodyDataList {
	s.CallsOffered = &v
	return s
}

func (s *ListHistoricalInstanceReportResponseBodyDataList) SetCallsRejected(v int64) *ListHistoricalInstanceReportResponseBodyDataList {
	s.CallsRejected = &v
	return s
}

func (s *ListHistoricalInstanceReportResponseBodyDataList) SetCallsResolved(v int64) *ListHistoricalInstanceReportResponseBodyDataList {
	s.CallsResolved = &v
	return s
}

func (s *ListHistoricalInstanceReportResponseBodyDataList) SetConfiguredConcurrency(v int64) *ListHistoricalInstanceReportResponseBodyDataList {
	s.ConfiguredConcurrency = &v
	return s
}

func (s *ListHistoricalInstanceReportResponseBodyDataList) SetInstanceId(v string) *ListHistoricalInstanceReportResponseBodyDataList {
	s.InstanceId = &v
	return s
}

func (s *ListHistoricalInstanceReportResponseBodyDataList) SetLabelBreakdown(v string) *ListHistoricalInstanceReportResponseBodyDataList {
	s.LabelBreakdown = &v
	return s
}

func (s *ListHistoricalInstanceReportResponseBodyDataList) SetMaxTalkTime(v int64) *ListHistoricalInstanceReportResponseBodyDataList {
	s.MaxTalkTime = &v
	return s
}

func (s *ListHistoricalInstanceReportResponseBodyDataList) SetRejectionBreakdown(v string) *ListHistoricalInstanceReportResponseBodyDataList {
	s.RejectionBreakdown = &v
	return s
}

func (s *ListHistoricalInstanceReportResponseBodyDataList) SetReleaseReasonBreakdown(v string) *ListHistoricalInstanceReportResponseBodyDataList {
	s.ReleaseReasonBreakdown = &v
	return s
}

func (s *ListHistoricalInstanceReportResponseBodyDataList) SetTalkTurnsDistribution(v string) *ListHistoricalInstanceReportResponseBodyDataList {
	s.TalkTurnsDistribution = &v
	return s
}

func (s *ListHistoricalInstanceReportResponseBodyDataList) SetTotalInputTokens(v int64) *ListHistoricalInstanceReportResponseBodyDataList {
	s.TotalInputTokens = &v
	return s
}

func (s *ListHistoricalInstanceReportResponseBodyDataList) SetTotalOutputTokens(v int64) *ListHistoricalInstanceReportResponseBodyDataList {
	s.TotalOutputTokens = &v
	return s
}

func (s *ListHistoricalInstanceReportResponseBodyDataList) SetTotalTalkTime(v int64) *ListHistoricalInstanceReportResponseBodyDataList {
	s.TotalTalkTime = &v
	return s
}

func (s *ListHistoricalInstanceReportResponseBodyDataList) SetTotalTokens(v int64) *ListHistoricalInstanceReportResponseBodyDataList {
	s.TotalTokens = &v
	return s
}

func (s *ListHistoricalInstanceReportResponseBodyDataList) SetUsedConcurrency(v int64) *ListHistoricalInstanceReportResponseBodyDataList {
	s.UsedConcurrency = &v
	return s
}

func (s *ListHistoricalInstanceReportResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
