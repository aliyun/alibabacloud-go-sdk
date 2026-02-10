// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHoneypotAlarmEventsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListHoneypotAlarmEventsRequest
	GetCurrentPage() *int32
	SetDealed(v string) *ListHoneypotAlarmEventsRequest
	GetDealed() *string
	SetDstIp(v string) *ListHoneypotAlarmEventsRequest
	GetDstIp() *string
	SetPageSize(v int32) *ListHoneypotAlarmEventsRequest
	GetPageSize() *int32
	SetRiskLevelList(v []*string) *ListHoneypotAlarmEventsRequest
	GetRiskLevelList() []*string
	SetSrcIp(v string) *ListHoneypotAlarmEventsRequest
	GetSrcIp() *string
}

type ListHoneypotAlarmEventsRequest struct {
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 10
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The status of the alert event. Default value: **a**. Valid values:
	//
	// 	- **a**: all states
	//
	// 	- **y**: handled
	//
	// 	- **n**: unhandled
	//
	// example:
	//
	// n
	Dealed *string `json:"Dealed,omitempty" xml:"Dealed,omitempty"`
	// The destination IP address.
	//
	// example:
	//
	// 172.20.XX.XX
	DstIp *string `json:"DstIp,omitempty" xml:"DstIp,omitempty"`
	// The number of entries per page. Default value: 100. If you leave this parameter empty, 100 entries are returned on each page.
	//
	// >  We recommend that you do not leave this parameter empty. We recommend that you set the value to a value no greater than 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The risk levels.
	RiskLevelList []*string `json:"RiskLevelList,omitempty" xml:"RiskLevelList,omitempty" type:"Repeated"`
	// The source IP address.
	//
	// example:
	//
	// 101.132.XX.XX
	SrcIp *string `json:"SrcIp,omitempty" xml:"SrcIp,omitempty"`
}

func (s ListHoneypotAlarmEventsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListHoneypotAlarmEventsRequest) GoString() string {
	return s.String()
}

func (s *ListHoneypotAlarmEventsRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListHoneypotAlarmEventsRequest) GetDealed() *string {
	return s.Dealed
}

func (s *ListHoneypotAlarmEventsRequest) GetDstIp() *string {
	return s.DstIp
}

func (s *ListHoneypotAlarmEventsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListHoneypotAlarmEventsRequest) GetRiskLevelList() []*string {
	return s.RiskLevelList
}

func (s *ListHoneypotAlarmEventsRequest) GetSrcIp() *string {
	return s.SrcIp
}

func (s *ListHoneypotAlarmEventsRequest) SetCurrentPage(v int32) *ListHoneypotAlarmEventsRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListHoneypotAlarmEventsRequest) SetDealed(v string) *ListHoneypotAlarmEventsRequest {
	s.Dealed = &v
	return s
}

func (s *ListHoneypotAlarmEventsRequest) SetDstIp(v string) *ListHoneypotAlarmEventsRequest {
	s.DstIp = &v
	return s
}

func (s *ListHoneypotAlarmEventsRequest) SetPageSize(v int32) *ListHoneypotAlarmEventsRequest {
	s.PageSize = &v
	return s
}

func (s *ListHoneypotAlarmEventsRequest) SetRiskLevelList(v []*string) *ListHoneypotAlarmEventsRequest {
	s.RiskLevelList = v
	return s
}

func (s *ListHoneypotAlarmEventsRequest) SetSrcIp(v string) *ListHoneypotAlarmEventsRequest {
	s.SrcIp = &v
	return s
}

func (s *ListHoneypotAlarmEventsRequest) Validate() error {
	return dara.Validate(s)
}
