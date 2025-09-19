// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHoneypotAttackerSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListHoneypotAttackerSourceRequest
	GetCurrentPage() *int32
	SetEndTimeStamp(v int64) *ListHoneypotAttackerSourceRequest
	GetEndTimeStamp() *int64
	SetLang(v string) *ListHoneypotAttackerSourceRequest
	GetLang() *string
	SetPageSize(v int32) *ListHoneypotAttackerSourceRequest
	GetPageSize() *int32
	SetRiskLevelList(v []*string) *ListHoneypotAttackerSourceRequest
	GetRiskLevelList() []*string
	SetSrcIp(v string) *ListHoneypotAttackerSourceRequest
	GetSrcIp() *string
	SetStartTimeStamp(v int64) *ListHoneypotAttackerSourceRequest
	GetStartTimeStamp() *int64
}

type ListHoneypotAttackerSourceRequest struct {
	// The page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The end of the time range to query. This value is a timestamp.
	//
	// example:
	//
	// 1676945366221
	EndTimeStamp *int64 `json:"EndTimeStamp,omitempty" xml:"EndTimeStamp,omitempty"`
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
	// The number of entries to return on each page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// An array that consists of risk levels.
	RiskLevelList []*string `json:"RiskLevelList,omitempty" xml:"RiskLevelList,omitempty" type:"Repeated"`
	// The source IP address of the attack.
	//
	// example:
	//
	// 175.136.230.***
	SrcIp *string `json:"SrcIp,omitempty" xml:"SrcIp,omitempty"`
	// The beginning of the time range to query. This value is a timestamp.
	//
	// example:
	//
	// 1674007632124
	StartTimeStamp *int64 `json:"StartTimeStamp,omitempty" xml:"StartTimeStamp,omitempty"`
}

func (s ListHoneypotAttackerSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s ListHoneypotAttackerSourceRequest) GoString() string {
	return s.String()
}

func (s *ListHoneypotAttackerSourceRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListHoneypotAttackerSourceRequest) GetEndTimeStamp() *int64 {
	return s.EndTimeStamp
}

func (s *ListHoneypotAttackerSourceRequest) GetLang() *string {
	return s.Lang
}

func (s *ListHoneypotAttackerSourceRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListHoneypotAttackerSourceRequest) GetRiskLevelList() []*string {
	return s.RiskLevelList
}

func (s *ListHoneypotAttackerSourceRequest) GetSrcIp() *string {
	return s.SrcIp
}

func (s *ListHoneypotAttackerSourceRequest) GetStartTimeStamp() *int64 {
	return s.StartTimeStamp
}

func (s *ListHoneypotAttackerSourceRequest) SetCurrentPage(v int32) *ListHoneypotAttackerSourceRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListHoneypotAttackerSourceRequest) SetEndTimeStamp(v int64) *ListHoneypotAttackerSourceRequest {
	s.EndTimeStamp = &v
	return s
}

func (s *ListHoneypotAttackerSourceRequest) SetLang(v string) *ListHoneypotAttackerSourceRequest {
	s.Lang = &v
	return s
}

func (s *ListHoneypotAttackerSourceRequest) SetPageSize(v int32) *ListHoneypotAttackerSourceRequest {
	s.PageSize = &v
	return s
}

func (s *ListHoneypotAttackerSourceRequest) SetRiskLevelList(v []*string) *ListHoneypotAttackerSourceRequest {
	s.RiskLevelList = v
	return s
}

func (s *ListHoneypotAttackerSourceRequest) SetSrcIp(v string) *ListHoneypotAttackerSourceRequest {
	s.SrcIp = &v
	return s
}

func (s *ListHoneypotAttackerSourceRequest) SetStartTimeStamp(v int64) *ListHoneypotAttackerSourceRequest {
	s.StartTimeStamp = &v
	return s
}

func (s *ListHoneypotAttackerSourceRequest) Validate() error {
	return dara.Validate(s)
}
