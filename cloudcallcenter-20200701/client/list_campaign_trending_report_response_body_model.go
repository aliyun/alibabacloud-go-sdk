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
	Code           *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           []*ListCampaignTrendingReportResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	HttpStatusCode *int32                                        `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

type ListCampaignTrendingReportResponseBodyData struct {
	BreakAgents                    *int64  `json:"BreakAgents,omitempty" xml:"BreakAgents,omitempty"`
	Concurrency                    *int64  `json:"Concurrency,omitempty" xml:"Concurrency,omitempty"`
	Datetime                       *int64  `json:"Datetime,omitempty" xml:"Datetime,omitempty"`
	LoggedInAgents                 *int64  `json:"LoggedInAgents,omitempty" xml:"LoggedInAgents,omitempty"`
	OutboundScenarioBreakingAgents *string `json:"OutboundScenarioBreakingAgents,omitempty" xml:"OutboundScenarioBreakingAgents,omitempty"`
	OutboundScenarioReadyAgents    *string `json:"OutboundScenarioReadyAgents,omitempty" xml:"OutboundScenarioReadyAgents,omitempty"`
	OutboundScenarioTalkingAgents  *string `json:"OutboundScenarioTalkingAgents,omitempty" xml:"OutboundScenarioTalkingAgents,omitempty"`
	OutboundScenarioWorkingAgents  *string `json:"OutboundScenarioWorkingAgents,omitempty" xml:"OutboundScenarioWorkingAgents,omitempty"`
	ReadyAgents                    *int64  `json:"ReadyAgents,omitempty" xml:"ReadyAgents,omitempty"`
	TalkAgents                     *int64  `json:"TalkAgents,omitempty" xml:"TalkAgents,omitempty"`
	WorkAgents                     *int64  `json:"WorkAgents,omitempty" xml:"WorkAgents,omitempty"`
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

func (s *ListCampaignTrendingReportResponseBodyData) GetConcurrency() *int64 {
	return s.Concurrency
}

func (s *ListCampaignTrendingReportResponseBodyData) GetDatetime() *int64 {
	return s.Datetime
}

func (s *ListCampaignTrendingReportResponseBodyData) GetLoggedInAgents() *int64 {
	return s.LoggedInAgents
}

func (s *ListCampaignTrendingReportResponseBodyData) GetOutboundScenarioBreakingAgents() *string {
	return s.OutboundScenarioBreakingAgents
}

func (s *ListCampaignTrendingReportResponseBodyData) GetOutboundScenarioReadyAgents() *string {
	return s.OutboundScenarioReadyAgents
}

func (s *ListCampaignTrendingReportResponseBodyData) GetOutboundScenarioTalkingAgents() *string {
	return s.OutboundScenarioTalkingAgents
}

func (s *ListCampaignTrendingReportResponseBodyData) GetOutboundScenarioWorkingAgents() *string {
	return s.OutboundScenarioWorkingAgents
}

func (s *ListCampaignTrendingReportResponseBodyData) GetReadyAgents() *int64 {
	return s.ReadyAgents
}

func (s *ListCampaignTrendingReportResponseBodyData) GetTalkAgents() *int64 {
	return s.TalkAgents
}

func (s *ListCampaignTrendingReportResponseBodyData) GetWorkAgents() *int64 {
	return s.WorkAgents
}

func (s *ListCampaignTrendingReportResponseBodyData) SetBreakAgents(v int64) *ListCampaignTrendingReportResponseBodyData {
	s.BreakAgents = &v
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

func (s *ListCampaignTrendingReportResponseBodyData) SetOutboundScenarioBreakingAgents(v string) *ListCampaignTrendingReportResponseBodyData {
	s.OutboundScenarioBreakingAgents = &v
	return s
}

func (s *ListCampaignTrendingReportResponseBodyData) SetOutboundScenarioReadyAgents(v string) *ListCampaignTrendingReportResponseBodyData {
	s.OutboundScenarioReadyAgents = &v
	return s
}

func (s *ListCampaignTrendingReportResponseBodyData) SetOutboundScenarioTalkingAgents(v string) *ListCampaignTrendingReportResponseBodyData {
	s.OutboundScenarioTalkingAgents = &v
	return s
}

func (s *ListCampaignTrendingReportResponseBodyData) SetOutboundScenarioWorkingAgents(v string) *ListCampaignTrendingReportResponseBodyData {
	s.OutboundScenarioWorkingAgents = &v
	return s
}

func (s *ListCampaignTrendingReportResponseBodyData) SetReadyAgents(v int64) *ListCampaignTrendingReportResponseBodyData {
	s.ReadyAgents = &v
	return s
}

func (s *ListCampaignTrendingReportResponseBodyData) SetTalkAgents(v int64) *ListCampaignTrendingReportResponseBodyData {
	s.TalkAgents = &v
	return s
}

func (s *ListCampaignTrendingReportResponseBodyData) SetWorkAgents(v int64) *ListCampaignTrendingReportResponseBodyData {
	s.WorkAgents = &v
	return s
}

func (s *ListCampaignTrendingReportResponseBodyData) Validate() error {
	return dara.Validate(s)
}
