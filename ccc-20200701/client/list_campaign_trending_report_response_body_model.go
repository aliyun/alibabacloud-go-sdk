// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCampaignTrendingReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListCampaignTrendingReportResponseBody
	GetCode() *string
	SetData(v []*ListCampaignTrendingReportResponseBodyData) *ListCampaignTrendingReportResponseBody
	GetData() []*ListCampaignTrendingReportResponseBodyData
	SetHttpStatusCode(v int32) *ListCampaignTrendingReportResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListCampaignTrendingReportResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListCampaignTrendingReportResponseBody
	GetRequestId() *string
}

type ListCampaignTrendingReportResponseBody struct {
	// example:
	//
	// OK
	Code *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*ListCampaignTrendingReportResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 6CCEF32F-8614-535F-A1D9-D85B8C0DC4F0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListCampaignTrendingReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCampaignTrendingReportResponseBody) GoString() string {
	return s.String()
}

func (s *ListCampaignTrendingReportResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListCampaignTrendingReportResponseBody) GetData() []*ListCampaignTrendingReportResponseBodyData {
	return s.Data
}

func (s *ListCampaignTrendingReportResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListCampaignTrendingReportResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListCampaignTrendingReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCampaignTrendingReportResponseBody) SetCode(v string) *ListCampaignTrendingReportResponseBody {
	s.Code = &v
	return s
}

func (s *ListCampaignTrendingReportResponseBody) SetData(v []*ListCampaignTrendingReportResponseBodyData) *ListCampaignTrendingReportResponseBody {
	s.Data = v
	return s
}

func (s *ListCampaignTrendingReportResponseBody) SetHttpStatusCode(v int32) *ListCampaignTrendingReportResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListCampaignTrendingReportResponseBody) SetMessage(v string) *ListCampaignTrendingReportResponseBody {
	s.Message = &v
	return s
}

func (s *ListCampaignTrendingReportResponseBody) SetRequestId(v string) *ListCampaignTrendingReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCampaignTrendingReportResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListCampaignTrendingReportResponseBodyData struct {
	// example:
	//
	// 0
	BreakAgents    *int64 `json:"BreakAgents,omitempty" xml:"BreakAgents,omitempty"`
	BreakingAgents *int64 `json:"BreakingAgents,omitempty" xml:"BreakingAgents,omitempty"`
	// example:
	//
	// 1
	Concurrency *int64 `json:"Concurrency,omitempty" xml:"Concurrency,omitempty"`
	// example:
	//
	// 1634037840000
	Datetime *int64 `json:"Datetime,omitempty" xml:"Datetime,omitempty"`
	// example:
	//
	// 2
	LoggedInAgents                 *int64 `json:"LoggedInAgents,omitempty" xml:"LoggedInAgents,omitempty"`
	OutboundScenarioBreakingAgents *int64 `json:"OutboundScenarioBreakingAgents,omitempty" xml:"OutboundScenarioBreakingAgents,omitempty"`
	OutboundScenarioReadyAgents    *int64 `json:"OutboundScenarioReadyAgents,omitempty" xml:"OutboundScenarioReadyAgents,omitempty"`
	OutboundScenarioTalkingAgents  *int64 `json:"OutboundScenarioTalkingAgents,omitempty" xml:"OutboundScenarioTalkingAgents,omitempty"`
	OutboundScenarioWorkingAgents  *int64 `json:"OutboundScenarioWorkingAgents,omitempty" xml:"OutboundScenarioWorkingAgents,omitempty"`
	// example:
	//
	// 2
	ReadyAgents *int64 `json:"ReadyAgents,omitempty" xml:"ReadyAgents,omitempty"`
	StatsTime   *int64 `json:"StatsTime,omitempty" xml:"StatsTime,omitempty"`
	// example:
	//
	// 0
	TalkAgents    *int64 `json:"TalkAgents,omitempty" xml:"TalkAgents,omitempty"`
	TalkingAgents *int64 `json:"TalkingAgents,omitempty" xml:"TalkingAgents,omitempty"`
	// example:
	//
	// 0
	WorkAgents    *int64 `json:"WorkAgents,omitempty" xml:"WorkAgents,omitempty"`
	WorkingAgents *int64 `json:"WorkingAgents,omitempty" xml:"WorkingAgents,omitempty"`
}

func (s ListCampaignTrendingReportResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListCampaignTrendingReportResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCampaignTrendingReportResponseBodyData) GetBreakAgents() *int64 {
	return s.BreakAgents
}

func (s *ListCampaignTrendingReportResponseBodyData) GetBreakingAgents() *int64 {
	return s.BreakingAgents
}

func (s *ListCampaignTrendingReportResponseBodyData) GetConcurrency() *int64 {
	return s.Concurrency
}

func (s *ListCampaignTrendingReportResponseBodyData) GetDatetime() *int64 {
	return s.Datetime
}

func (s *ListCampaignTrendingReportResponseBodyData) GetLoggedInAgents() *int64 {
	return s.LoggedInAgents
}

func (s *ListCampaignTrendingReportResponseBodyData) GetOutboundScenarioBreakingAgents() *int64 {
	return s.OutboundScenarioBreakingAgents
}

func (s *ListCampaignTrendingReportResponseBodyData) GetOutboundScenarioReadyAgents() *int64 {
	return s.OutboundScenarioReadyAgents
}

func (s *ListCampaignTrendingReportResponseBodyData) GetOutboundScenarioTalkingAgents() *int64 {
	return s.OutboundScenarioTalkingAgents
}

func (s *ListCampaignTrendingReportResponseBodyData) GetOutboundScenarioWorkingAgents() *int64 {
	return s.OutboundScenarioWorkingAgents
}

func (s *ListCampaignTrendingReportResponseBodyData) GetReadyAgents() *int64 {
	return s.ReadyAgents
}

func (s *ListCampaignTrendingReportResponseBodyData) GetStatsTime() *int64 {
	return s.StatsTime
}

func (s *ListCampaignTrendingReportResponseBodyData) GetTalkAgents() *int64 {
	return s.TalkAgents
}

func (s *ListCampaignTrendingReportResponseBodyData) GetTalkingAgents() *int64 {
	return s.TalkingAgents
}

func (s *ListCampaignTrendingReportResponseBodyData) GetWorkAgents() *int64 {
	return s.WorkAgents
}

func (s *ListCampaignTrendingReportResponseBodyData) GetWorkingAgents() *int64 {
	return s.WorkingAgents
}

func (s *ListCampaignTrendingReportResponseBodyData) SetBreakAgents(v int64) *ListCampaignTrendingReportResponseBodyData {
	s.BreakAgents = &v
	return s
}

func (s *ListCampaignTrendingReportResponseBodyData) SetBreakingAgents(v int64) *ListCampaignTrendingReportResponseBodyData {
	s.BreakingAgents = &v
	return s
}

func (s *ListCampaignTrendingReportResponseBodyData) SetConcurrency(v int64) *ListCampaignTrendingReportResponseBodyData {
	s.Concurrency = &v
	return s
}

func (s *ListCampaignTrendingReportResponseBodyData) SetDatetime(v int64) *ListCampaignTrendingReportResponseBodyData {
	s.Datetime = &v
	return s
}

func (s *ListCampaignTrendingReportResponseBodyData) SetLoggedInAgents(v int64) *ListCampaignTrendingReportResponseBodyData {
	s.LoggedInAgents = &v
	return s
}

func (s *ListCampaignTrendingReportResponseBodyData) SetOutboundScenarioBreakingAgents(v int64) *ListCampaignTrendingReportResponseBodyData {
	s.OutboundScenarioBreakingAgents = &v
	return s
}

func (s *ListCampaignTrendingReportResponseBodyData) SetOutboundScenarioReadyAgents(v int64) *ListCampaignTrendingReportResponseBodyData {
	s.OutboundScenarioReadyAgents = &v
	return s
}

func (s *ListCampaignTrendingReportResponseBodyData) SetOutboundScenarioTalkingAgents(v int64) *ListCampaignTrendingReportResponseBodyData {
	s.OutboundScenarioTalkingAgents = &v
	return s
}

func (s *ListCampaignTrendingReportResponseBodyData) SetOutboundScenarioWorkingAgents(v int64) *ListCampaignTrendingReportResponseBodyData {
	s.OutboundScenarioWorkingAgents = &v
	return s
}

func (s *ListCampaignTrendingReportResponseBodyData) SetReadyAgents(v int64) *ListCampaignTrendingReportResponseBodyData {
	s.ReadyAgents = &v
	return s
}

func (s *ListCampaignTrendingReportResponseBodyData) SetStatsTime(v int64) *ListCampaignTrendingReportResponseBodyData {
	s.StatsTime = &v
	return s
}

func (s *ListCampaignTrendingReportResponseBodyData) SetTalkAgents(v int64) *ListCampaignTrendingReportResponseBodyData {
	s.TalkAgents = &v
	return s
}

func (s *ListCampaignTrendingReportResponseBodyData) SetTalkingAgents(v int64) *ListCampaignTrendingReportResponseBodyData {
	s.TalkingAgents = &v
	return s
}

func (s *ListCampaignTrendingReportResponseBodyData) SetWorkAgents(v int64) *ListCampaignTrendingReportResponseBodyData {
	s.WorkAgents = &v
	return s
}

func (s *ListCampaignTrendingReportResponseBodyData) SetWorkingAgents(v int64) *ListCampaignTrendingReportResponseBodyData {
	s.WorkingAgents = &v
	return s
}

func (s *ListCampaignTrendingReportResponseBodyData) Validate() error {
	return dara.Validate(s)
}
