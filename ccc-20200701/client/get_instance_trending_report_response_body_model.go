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
	SetData(v *GetInstanceTrendingReportResponseBodyData) *GetInstanceTrendingReportResponseBody
	GetData() *GetInstanceTrendingReportResponseBodyData
	SetHttpStatusCode(v int32) *GetInstanceTrendingReportResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetInstanceTrendingReportResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetInstanceTrendingReportResponseBody
	GetRequestId() *string
}

type GetInstanceTrendingReportResponseBody struct {
	// example:
	//
	// OK
	Code *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetInstanceTrendingReportResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 943D8EF3-3321-471F-A104-51C96FCA94D6
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

func (s *GetInstanceTrendingReportResponseBody) GetData() *GetInstanceTrendingReportResponseBodyData {
	return s.Data
}

func (s *GetInstanceTrendingReportResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetInstanceTrendingReportResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetInstanceTrendingReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstanceTrendingReportResponseBody) SetCode(v string) *GetInstanceTrendingReportResponseBody {
	s.Code = &v
	return s
}

func (s *GetInstanceTrendingReportResponseBody) SetData(v *GetInstanceTrendingReportResponseBodyData) *GetInstanceTrendingReportResponseBody {
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

func (s *GetInstanceTrendingReportResponseBody) SetRequestId(v string) *GetInstanceTrendingReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceTrendingReportResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetInstanceTrendingReportResponseBodyData struct {
	Inbound  []*GetInstanceTrendingReportResponseBodyDataInbound  `json:"Inbound,omitempty" xml:"Inbound,omitempty" type:"Repeated"`
	Outbound []*GetInstanceTrendingReportResponseBodyDataOutbound `json:"Outbound,omitempty" xml:"Outbound,omitempty" type:"Repeated"`
	Overall  []*GetInstanceTrendingReportResponseBodyDataOverall  `json:"Overall,omitempty" xml:"Overall,omitempty" type:"Repeated"`
}

func (s GetInstanceTrendingReportResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceTrendingReportResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetInstanceTrendingReportResponseBodyData) GetInbound() []*GetInstanceTrendingReportResponseBodyDataInbound {
	return s.Inbound
}

func (s *GetInstanceTrendingReportResponseBodyData) GetOutbound() []*GetInstanceTrendingReportResponseBodyDataOutbound {
	return s.Outbound
}

func (s *GetInstanceTrendingReportResponseBodyData) GetOverall() []*GetInstanceTrendingReportResponseBodyDataOverall {
	return s.Overall
}

func (s *GetInstanceTrendingReportResponseBodyData) SetInbound(v []*GetInstanceTrendingReportResponseBodyDataInbound) *GetInstanceTrendingReportResponseBodyData {
	s.Inbound = v
	return s
}

func (s *GetInstanceTrendingReportResponseBodyData) SetOutbound(v []*GetInstanceTrendingReportResponseBodyDataOutbound) *GetInstanceTrendingReportResponseBodyData {
	s.Outbound = v
	return s
}

func (s *GetInstanceTrendingReportResponseBodyData) SetOverall(v []*GetInstanceTrendingReportResponseBodyDataOverall) *GetInstanceTrendingReportResponseBodyData {
	s.Overall = v
	return s
}

func (s *GetInstanceTrendingReportResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetInstanceTrendingReportResponseBodyDataInbound struct {
	// example:
	//
	// 0
	CallsAbandonedInIVR *int64 `json:"CallsAbandonedInIVR,omitempty" xml:"CallsAbandonedInIVR,omitempty"`
	// example:
	//
	// 0
	CallsAbandonedInQueue *int64 `json:"CallsAbandonedInQueue,omitempty" xml:"CallsAbandonedInQueue,omitempty"`
	// example:
	//
	// 0
	CallsAbandonedInRing *int64 `json:"CallsAbandonedInRing,omitempty" xml:"CallsAbandonedInRing,omitempty"`
	// example:
	//
	// 0
	CallsHandled *int64 `json:"CallsHandled,omitempty" xml:"CallsHandled,omitempty"`
	// example:
	//
	// 0
	CallsQueued *int64 `json:"CallsQueued,omitempty" xml:"CallsQueued,omitempty"`
	// example:
	//
	// 1604639129000
	StatsTime *int64 `json:"StatsTime,omitempty" xml:"StatsTime,omitempty"`
	// example:
	//
	// 0
	TotalCalls *int64 `json:"TotalCalls,omitempty" xml:"TotalCalls,omitempty"`
}

func (s GetInstanceTrendingReportResponseBodyDataInbound) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceTrendingReportResponseBodyDataInbound) GoString() string {
	return s.String()
}

func (s *GetInstanceTrendingReportResponseBodyDataInbound) GetCallsAbandonedInIVR() *int64 {
	return s.CallsAbandonedInIVR
}

func (s *GetInstanceTrendingReportResponseBodyDataInbound) GetCallsAbandonedInQueue() *int64 {
	return s.CallsAbandonedInQueue
}

func (s *GetInstanceTrendingReportResponseBodyDataInbound) GetCallsAbandonedInRing() *int64 {
	return s.CallsAbandonedInRing
}

func (s *GetInstanceTrendingReportResponseBodyDataInbound) GetCallsHandled() *int64 {
	return s.CallsHandled
}

func (s *GetInstanceTrendingReportResponseBodyDataInbound) GetCallsQueued() *int64 {
	return s.CallsQueued
}

func (s *GetInstanceTrendingReportResponseBodyDataInbound) GetStatsTime() *int64 {
	return s.StatsTime
}

func (s *GetInstanceTrendingReportResponseBodyDataInbound) GetTotalCalls() *int64 {
	return s.TotalCalls
}

func (s *GetInstanceTrendingReportResponseBodyDataInbound) SetCallsAbandonedInIVR(v int64) *GetInstanceTrendingReportResponseBodyDataInbound {
	s.CallsAbandonedInIVR = &v
	return s
}

func (s *GetInstanceTrendingReportResponseBodyDataInbound) SetCallsAbandonedInQueue(v int64) *GetInstanceTrendingReportResponseBodyDataInbound {
	s.CallsAbandonedInQueue = &v
	return s
}

func (s *GetInstanceTrendingReportResponseBodyDataInbound) SetCallsAbandonedInRing(v int64) *GetInstanceTrendingReportResponseBodyDataInbound {
	s.CallsAbandonedInRing = &v
	return s
}

func (s *GetInstanceTrendingReportResponseBodyDataInbound) SetCallsHandled(v int64) *GetInstanceTrendingReportResponseBodyDataInbound {
	s.CallsHandled = &v
	return s
}

func (s *GetInstanceTrendingReportResponseBodyDataInbound) SetCallsQueued(v int64) *GetInstanceTrendingReportResponseBodyDataInbound {
	s.CallsQueued = &v
	return s
}

func (s *GetInstanceTrendingReportResponseBodyDataInbound) SetStatsTime(v int64) *GetInstanceTrendingReportResponseBodyDataInbound {
	s.StatsTime = &v
	return s
}

func (s *GetInstanceTrendingReportResponseBodyDataInbound) SetTotalCalls(v int64) *GetInstanceTrendingReportResponseBodyDataInbound {
	s.TotalCalls = &v
	return s
}

func (s *GetInstanceTrendingReportResponseBodyDataInbound) Validate() error {
	return dara.Validate(s)
}

type GetInstanceTrendingReportResponseBodyDataOutbound struct {
	// example:
	//
	// 0
	CallsAnswered *int64 `json:"CallsAnswered,omitempty" xml:"CallsAnswered,omitempty"`
	// example:
	//
	// 1604639129000
	StatsTime *int64 `json:"StatsTime,omitempty" xml:"StatsTime,omitempty"`
	// example:
	//
	// 0
	TotalCalls *int64 `json:"TotalCalls,omitempty" xml:"TotalCalls,omitempty"`
}

func (s GetInstanceTrendingReportResponseBodyDataOutbound) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceTrendingReportResponseBodyDataOutbound) GoString() string {
	return s.String()
}

func (s *GetInstanceTrendingReportResponseBodyDataOutbound) GetCallsAnswered() *int64 {
	return s.CallsAnswered
}

func (s *GetInstanceTrendingReportResponseBodyDataOutbound) GetStatsTime() *int64 {
	return s.StatsTime
}

func (s *GetInstanceTrendingReportResponseBodyDataOutbound) GetTotalCalls() *int64 {
	return s.TotalCalls
}

func (s *GetInstanceTrendingReportResponseBodyDataOutbound) SetCallsAnswered(v int64) *GetInstanceTrendingReportResponseBodyDataOutbound {
	s.CallsAnswered = &v
	return s
}

func (s *GetInstanceTrendingReportResponseBodyDataOutbound) SetStatsTime(v int64) *GetInstanceTrendingReportResponseBodyDataOutbound {
	s.StatsTime = &v
	return s
}

func (s *GetInstanceTrendingReportResponseBodyDataOutbound) SetTotalCalls(v int64) *GetInstanceTrendingReportResponseBodyDataOutbound {
	s.TotalCalls = &v
	return s
}

func (s *GetInstanceTrendingReportResponseBodyDataOutbound) Validate() error {
	return dara.Validate(s)
}

type GetInstanceTrendingReportResponseBodyDataOverall struct {
	MaxLoggedInAgents *int64 `json:"MaxLoggedInAgents,omitempty" xml:"MaxLoggedInAgents,omitempty"`
	StatsTime         *int64 `json:"StatsTime,omitempty" xml:"StatsTime,omitempty"`
}

func (s GetInstanceTrendingReportResponseBodyDataOverall) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceTrendingReportResponseBodyDataOverall) GoString() string {
	return s.String()
}

func (s *GetInstanceTrendingReportResponseBodyDataOverall) GetMaxLoggedInAgents() *int64 {
	return s.MaxLoggedInAgents
}

func (s *GetInstanceTrendingReportResponseBodyDataOverall) GetStatsTime() *int64 {
	return s.StatsTime
}

func (s *GetInstanceTrendingReportResponseBodyDataOverall) SetMaxLoggedInAgents(v int64) *GetInstanceTrendingReportResponseBodyDataOverall {
	s.MaxLoggedInAgents = &v
	return s
}

func (s *GetInstanceTrendingReportResponseBodyDataOverall) SetStatsTime(v int64) *GetInstanceTrendingReportResponseBodyDataOverall {
	s.StatsTime = &v
	return s
}

func (s *GetInstanceTrendingReportResponseBodyDataOverall) Validate() error {
	return dara.Validate(s)
}
