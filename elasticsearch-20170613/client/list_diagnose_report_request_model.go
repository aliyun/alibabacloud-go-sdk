// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDiagnoseReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDetail(v bool) *ListDiagnoseReportRequest
	GetDetail() *bool
	SetEndTime(v int64) *ListDiagnoseReportRequest
	GetEndTime() *int64
	SetLang(v string) *ListDiagnoseReportRequest
	GetLang() *string
	SetPage(v int32) *ListDiagnoseReportRequest
	GetPage() *int32
	SetSize(v int32) *ListDiagnoseReportRequest
	GetSize() *int32
	SetStartTime(v int64) *ListDiagnoseReportRequest
	GetStartTime() *int64
	SetTrigger(v string) *ListDiagnoseReportRequest
	GetTrigger() *string
}

type ListDiagnoseReportRequest struct {
	// SYSTEM
	//
	// example:
	//
	// true
	Detail *bool `json:"detail,omitempty" xml:"detail,omitempty"`
	// 1
	//
	// This parameter is required.
	//
	// example:
	//
	// 1595174399999
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// 1594569600000
	//
	// example:
	//
	// spanish
	Lang *string `json:"lang,omitempty" xml:"lang,omitempty"`
	// 20
	//
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// true
	//
	// example:
	//
	// 20
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
	// 1595174399999
	//
	// This parameter is required.
	//
	// example:
	//
	// 1594569600000
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// SYSTEM
	Trigger *string `json:"trigger,omitempty" xml:"trigger,omitempty"`
}

func (s ListDiagnoseReportRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDiagnoseReportRequest) GoString() string {
	return s.String()
}

func (s *ListDiagnoseReportRequest) GetDetail() *bool {
	return s.Detail
}

func (s *ListDiagnoseReportRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListDiagnoseReportRequest) GetLang() *string {
	return s.Lang
}

func (s *ListDiagnoseReportRequest) GetPage() *int32 {
	return s.Page
}

func (s *ListDiagnoseReportRequest) GetSize() *int32 {
	return s.Size
}

func (s *ListDiagnoseReportRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListDiagnoseReportRequest) GetTrigger() *string {
	return s.Trigger
}

func (s *ListDiagnoseReportRequest) SetDetail(v bool) *ListDiagnoseReportRequest {
	s.Detail = &v
	return s
}

func (s *ListDiagnoseReportRequest) SetEndTime(v int64) *ListDiagnoseReportRequest {
	s.EndTime = &v
	return s
}

func (s *ListDiagnoseReportRequest) SetLang(v string) *ListDiagnoseReportRequest {
	s.Lang = &v
	return s
}

func (s *ListDiagnoseReportRequest) SetPage(v int32) *ListDiagnoseReportRequest {
	s.Page = &v
	return s
}

func (s *ListDiagnoseReportRequest) SetSize(v int32) *ListDiagnoseReportRequest {
	s.Size = &v
	return s
}

func (s *ListDiagnoseReportRequest) SetStartTime(v int64) *ListDiagnoseReportRequest {
	s.StartTime = &v
	return s
}

func (s *ListDiagnoseReportRequest) SetTrigger(v string) *ListDiagnoseReportRequest {
	s.Trigger = &v
	return s
}

func (s *ListDiagnoseReportRequest) Validate() error {
	return dara.Validate(s)
}
