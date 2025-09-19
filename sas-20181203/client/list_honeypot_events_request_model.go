// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHoneypotEventsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentIdList(v []*string) *ListHoneypotEventsRequest
	GetAgentIdList() []*string
	SetAlarmEventId(v int64) *ListHoneypotEventsRequest
	GetAlarmEventId() *int64
	SetCurrentPage(v int32) *ListHoneypotEventsRequest
	GetCurrentPage() *int32
	SetDealed(v string) *ListHoneypotEventsRequest
	GetDealed() *string
	SetHoneypotIdList(v []*string) *ListHoneypotEventsRequest
	GetHoneypotIdList() []*string
	SetLang(v string) *ListHoneypotEventsRequest
	GetLang() *string
	SetPageSize(v int32) *ListHoneypotEventsRequest
	GetPageSize() *int32
	SetPortraitId(v string) *ListHoneypotEventsRequest
	GetPortraitId() *string
	SetRequestId(v string) *ListHoneypotEventsRequest
	GetRequestId() *string
	SetRiskLevelList(v []*string) *ListHoneypotEventsRequest
	GetRiskLevelList() []*string
	SetSrcIp(v string) *ListHoneypotEventsRequest
	GetSrcIp() *string
}

type ListHoneypotEventsRequest struct {
	// The probe IDs.
	AgentIdList []*string `json:"AgentIdList,omitempty" xml:"AgentIdList,omitempty" type:"Repeated"`
	// The ID of the alert.
	//
	// example:
	//
	// 1259925
	AlarmEventId *int64 `json:"AlarmEventId,omitempty" xml:"AlarmEventId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The status of the event. Valid values:
	//
	// 	- **y**: handled
	//
	// 	- **n**: unhandled
	//
	// 	- **a**: all statuses
	//
	// example:
	//
	// y
	Dealed *string `json:"Dealed,omitempty" xml:"Dealed,omitempty"`
	// The honeypot IDs.
	HoneypotIdList []*string `json:"HoneypotIdList,omitempty" xml:"HoneypotIdList,omitempty" type:"Repeated"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of entries per page. Default value: **20**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the attacker profile.
	//
	// example:
	//
	// cd48604a-1694-4f03-ade0-ec6994c3*****
	PortraitId *string `json:"PortraitId,omitempty" xml:"PortraitId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2F9CE167-58D5-5DA6-AA3B-923EED02****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The risk levels.
	RiskLevelList []*string `json:"RiskLevelList,omitempty" xml:"RiskLevelList,omitempty" type:"Repeated"`
	// The source IP address of the attack.
	//
	// example:
	//
	// 185.237.96.***
	SrcIp *string `json:"SrcIp,omitempty" xml:"SrcIp,omitempty"`
}

func (s ListHoneypotEventsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListHoneypotEventsRequest) GoString() string {
	return s.String()
}

func (s *ListHoneypotEventsRequest) GetAgentIdList() []*string {
	return s.AgentIdList
}

func (s *ListHoneypotEventsRequest) GetAlarmEventId() *int64 {
	return s.AlarmEventId
}

func (s *ListHoneypotEventsRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListHoneypotEventsRequest) GetDealed() *string {
	return s.Dealed
}

func (s *ListHoneypotEventsRequest) GetHoneypotIdList() []*string {
	return s.HoneypotIdList
}

func (s *ListHoneypotEventsRequest) GetLang() *string {
	return s.Lang
}

func (s *ListHoneypotEventsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListHoneypotEventsRequest) GetPortraitId() *string {
	return s.PortraitId
}

func (s *ListHoneypotEventsRequest) GetRequestId() *string {
	return s.RequestId
}

func (s *ListHoneypotEventsRequest) GetRiskLevelList() []*string {
	return s.RiskLevelList
}

func (s *ListHoneypotEventsRequest) GetSrcIp() *string {
	return s.SrcIp
}

func (s *ListHoneypotEventsRequest) SetAgentIdList(v []*string) *ListHoneypotEventsRequest {
	s.AgentIdList = v
	return s
}

func (s *ListHoneypotEventsRequest) SetAlarmEventId(v int64) *ListHoneypotEventsRequest {
	s.AlarmEventId = &v
	return s
}

func (s *ListHoneypotEventsRequest) SetCurrentPage(v int32) *ListHoneypotEventsRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListHoneypotEventsRequest) SetDealed(v string) *ListHoneypotEventsRequest {
	s.Dealed = &v
	return s
}

func (s *ListHoneypotEventsRequest) SetHoneypotIdList(v []*string) *ListHoneypotEventsRequest {
	s.HoneypotIdList = v
	return s
}

func (s *ListHoneypotEventsRequest) SetLang(v string) *ListHoneypotEventsRequest {
	s.Lang = &v
	return s
}

func (s *ListHoneypotEventsRequest) SetPageSize(v int32) *ListHoneypotEventsRequest {
	s.PageSize = &v
	return s
}

func (s *ListHoneypotEventsRequest) SetPortraitId(v string) *ListHoneypotEventsRequest {
	s.PortraitId = &v
	return s
}

func (s *ListHoneypotEventsRequest) SetRequestId(v string) *ListHoneypotEventsRequest {
	s.RequestId = &v
	return s
}

func (s *ListHoneypotEventsRequest) SetRiskLevelList(v []*string) *ListHoneypotEventsRequest {
	s.RiskLevelList = v
	return s
}

func (s *ListHoneypotEventsRequest) SetSrcIp(v string) *ListHoneypotEventsRequest {
	s.SrcIp = &v
	return s
}

func (s *ListHoneypotEventsRequest) Validate() error {
	return dara.Validate(s)
}
