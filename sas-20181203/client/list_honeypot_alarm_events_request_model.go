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
	CurrentPage   *int32    `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Dealed        *string   `json:"Dealed,omitempty" xml:"Dealed,omitempty"`
	DstIp         *string   `json:"DstIp,omitempty" xml:"DstIp,omitempty"`
	PageSize      *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RiskLevelList []*string `json:"RiskLevelList,omitempty" xml:"RiskLevelList,omitempty" type:"Repeated"`
	SrcIp         *string   `json:"SrcIp,omitempty" xml:"SrcIp,omitempty"`
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
