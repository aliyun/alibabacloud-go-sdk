// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRealtimeCampaignStatsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetRealtimeCampaignStatsResponseBody
	GetCode() *string
	SetData(v *GetRealtimeCampaignStatsResponseBodyData) *GetRealtimeCampaignStatsResponseBody
	GetData() *GetRealtimeCampaignStatsResponseBodyData
	SetHttpStatusCode(v int32) *GetRealtimeCampaignStatsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetRealtimeCampaignStatsResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetRealtimeCampaignStatsResponseBody
	GetRequestId() *string
}

type GetRealtimeCampaignStatsResponseBody struct {
	// example:
	//
	// OK
	Code *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetRealtimeCampaignStatsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 42970829-E2C8-515A-8F42-5A6B59F852A7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetRealtimeCampaignStatsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRealtimeCampaignStatsResponseBody) GoString() string {
	return s.String()
}

func (s *GetRealtimeCampaignStatsResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetRealtimeCampaignStatsResponseBody) GetData() *GetRealtimeCampaignStatsResponseBodyData {
	return s.Data
}

func (s *GetRealtimeCampaignStatsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetRealtimeCampaignStatsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetRealtimeCampaignStatsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRealtimeCampaignStatsResponseBody) SetCode(v string) *GetRealtimeCampaignStatsResponseBody {
	s.Code = &v
	return s
}

func (s *GetRealtimeCampaignStatsResponseBody) SetData(v *GetRealtimeCampaignStatsResponseBodyData) *GetRealtimeCampaignStatsResponseBody {
	s.Data = v
	return s
}

func (s *GetRealtimeCampaignStatsResponseBody) SetHttpStatusCode(v int32) *GetRealtimeCampaignStatsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetRealtimeCampaignStatsResponseBody) SetMessage(v string) *GetRealtimeCampaignStatsResponseBody {
	s.Message = &v
	return s
}

func (s *GetRealtimeCampaignStatsResponseBody) SetRequestId(v string) *GetRealtimeCampaignStatsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRealtimeCampaignStatsResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetRealtimeCampaignStatsResponseBodyData struct {
	// example:
	//
	// 1
	BreakingAgents *int64 `json:"BreakingAgents,omitempty" xml:"BreakingAgents,omitempty"`
	// example:
	//
	// 5
	Caps *int64 `json:"Caps,omitempty" xml:"Caps,omitempty"`
	// example:
	//
	// 10
	LoggedInAgents                 *int64 `json:"LoggedInAgents,omitempty" xml:"LoggedInAgents,omitempty"`
	OutboundScenarioBreakingAgents *int64 `json:"OutboundScenarioBreakingAgents,omitempty" xml:"OutboundScenarioBreakingAgents,omitempty"`
	OutboundScenarioReadyAgents    *int64 `json:"OutboundScenarioReadyAgents,omitempty" xml:"OutboundScenarioReadyAgents,omitempty"`
	OutboundScenarioTalkingAgents  *int64 `json:"OutboundScenarioTalkingAgents,omitempty" xml:"OutboundScenarioTalkingAgents,omitempty"`
	OutboundScenarioWorkingAgents  *int64 `json:"OutboundScenarioWorkingAgents,omitempty" xml:"OutboundScenarioWorkingAgents,omitempty"`
	// example:
	//
	// 3
	ReadyAgents *int64 `json:"ReadyAgents,omitempty" xml:"ReadyAgents,omitempty"`
	// example:
	//
	// 4
	TalkingAgents *int64 `json:"TalkingAgents,omitempty" xml:"TalkingAgents,omitempty"`
	// example:
	//
	// 10
	TotalAgents *int64 `json:"TotalAgents,omitempty" xml:"TotalAgents,omitempty"`
	// example:
	//
	// 2
	WorkingAgents *int64 `json:"WorkingAgents,omitempty" xml:"WorkingAgents,omitempty"`
}

func (s GetRealtimeCampaignStatsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetRealtimeCampaignStatsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRealtimeCampaignStatsResponseBodyData) GetBreakingAgents() *int64 {
	return s.BreakingAgents
}

func (s *GetRealtimeCampaignStatsResponseBodyData) GetCaps() *int64 {
	return s.Caps
}

func (s *GetRealtimeCampaignStatsResponseBodyData) GetLoggedInAgents() *int64 {
	return s.LoggedInAgents
}

func (s *GetRealtimeCampaignStatsResponseBodyData) GetOutboundScenarioBreakingAgents() *int64 {
	return s.OutboundScenarioBreakingAgents
}

func (s *GetRealtimeCampaignStatsResponseBodyData) GetOutboundScenarioReadyAgents() *int64 {
	return s.OutboundScenarioReadyAgents
}

func (s *GetRealtimeCampaignStatsResponseBodyData) GetOutboundScenarioTalkingAgents() *int64 {
	return s.OutboundScenarioTalkingAgents
}

func (s *GetRealtimeCampaignStatsResponseBodyData) GetOutboundScenarioWorkingAgents() *int64 {
	return s.OutboundScenarioWorkingAgents
}

func (s *GetRealtimeCampaignStatsResponseBodyData) GetReadyAgents() *int64 {
	return s.ReadyAgents
}

func (s *GetRealtimeCampaignStatsResponseBodyData) GetTalkingAgents() *int64 {
	return s.TalkingAgents
}

func (s *GetRealtimeCampaignStatsResponseBodyData) GetTotalAgents() *int64 {
	return s.TotalAgents
}

func (s *GetRealtimeCampaignStatsResponseBodyData) GetWorkingAgents() *int64 {
	return s.WorkingAgents
}

func (s *GetRealtimeCampaignStatsResponseBodyData) SetBreakingAgents(v int64) *GetRealtimeCampaignStatsResponseBodyData {
	s.BreakingAgents = &v
	return s
}

func (s *GetRealtimeCampaignStatsResponseBodyData) SetCaps(v int64) *GetRealtimeCampaignStatsResponseBodyData {
	s.Caps = &v
	return s
}

func (s *GetRealtimeCampaignStatsResponseBodyData) SetLoggedInAgents(v int64) *GetRealtimeCampaignStatsResponseBodyData {
	s.LoggedInAgents = &v
	return s
}

func (s *GetRealtimeCampaignStatsResponseBodyData) SetOutboundScenarioBreakingAgents(v int64) *GetRealtimeCampaignStatsResponseBodyData {
	s.OutboundScenarioBreakingAgents = &v
	return s
}

func (s *GetRealtimeCampaignStatsResponseBodyData) SetOutboundScenarioReadyAgents(v int64) *GetRealtimeCampaignStatsResponseBodyData {
	s.OutboundScenarioReadyAgents = &v
	return s
}

func (s *GetRealtimeCampaignStatsResponseBodyData) SetOutboundScenarioTalkingAgents(v int64) *GetRealtimeCampaignStatsResponseBodyData {
	s.OutboundScenarioTalkingAgents = &v
	return s
}

func (s *GetRealtimeCampaignStatsResponseBodyData) SetOutboundScenarioWorkingAgents(v int64) *GetRealtimeCampaignStatsResponseBodyData {
	s.OutboundScenarioWorkingAgents = &v
	return s
}

func (s *GetRealtimeCampaignStatsResponseBodyData) SetReadyAgents(v int64) *GetRealtimeCampaignStatsResponseBodyData {
	s.ReadyAgents = &v
	return s
}

func (s *GetRealtimeCampaignStatsResponseBodyData) SetTalkingAgents(v int64) *GetRealtimeCampaignStatsResponseBodyData {
	s.TalkingAgents = &v
	return s
}

func (s *GetRealtimeCampaignStatsResponseBodyData) SetTotalAgents(v int64) *GetRealtimeCampaignStatsResponseBodyData {
	s.TotalAgents = &v
	return s
}

func (s *GetRealtimeCampaignStatsResponseBodyData) SetWorkingAgents(v int64) *GetRealtimeCampaignStatsResponseBodyData {
	s.WorkingAgents = &v
	return s
}

func (s *GetRealtimeCampaignStatsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
