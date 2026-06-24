// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDiagnoseReportIdsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *ListDiagnoseReportIdsRequest
	GetEndTime() *int64
	SetLang(v string) *ListDiagnoseReportIdsRequest
	GetLang() *string
	SetPage(v int32) *ListDiagnoseReportIdsRequest
	GetPage() *int32
	SetSize(v int32) *ListDiagnoseReportIdsRequest
	GetSize() *int32
	SetStartTime(v int64) *ListDiagnoseReportIdsRequest
	GetStartTime() *int64
	SetTrigger(v string) *ListDiagnoseReportIdsRequest
	GetTrigger() *string
}

type ListDiagnoseReportIdsRequest struct {
	// The end timestamp of the query.
	//
	// Minimum value: 1000000000000 ms
	//
	// Maximum value: 2000000000000 ms.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1595174399999
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The language of the reports to retrieve. Default value: the browser language. Valid values:
	//
	// - en: English
	//
	// - zh: Simplified Chinese
	//
	// - zt: Traditional Chinese
	//
	// - es: Spanish
	//
	// - fr: French.
	//
	// example:
	//
	// en
	Lang *string `json:"lang,omitempty" xml:"lang,omitempty"`
	// The page number. Default value: 1. Minimum value: 1. Maximum value: 200.
	//
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// The number of report IDs per page. Default value: 10. Minimum value: 1. Maximum value: 500.
	//
	// example:
	//
	// 15
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
	// The start timestamp of the query.
	//
	// Minimum value: 1000000000000 ms
	//
	// Maximum value: 2000000000000 ms.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1595088000000
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// The trigger method of the health diagnostics. Valid values: SYSTEM (automatically triggered by the system), INNER (internally triggered), and USER (manually triggered by the user).
	//
	// example:
	//
	// SYSTEM
	Trigger *string `json:"trigger,omitempty" xml:"trigger,omitempty"`
}

func (s ListDiagnoseReportIdsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDiagnoseReportIdsRequest) GoString() string {
	return s.String()
}

func (s *ListDiagnoseReportIdsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListDiagnoseReportIdsRequest) GetLang() *string {
	return s.Lang
}

func (s *ListDiagnoseReportIdsRequest) GetPage() *int32 {
	return s.Page
}

func (s *ListDiagnoseReportIdsRequest) GetSize() *int32 {
	return s.Size
}

func (s *ListDiagnoseReportIdsRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListDiagnoseReportIdsRequest) GetTrigger() *string {
	return s.Trigger
}

func (s *ListDiagnoseReportIdsRequest) SetEndTime(v int64) *ListDiagnoseReportIdsRequest {
	s.EndTime = &v
	return s
}

func (s *ListDiagnoseReportIdsRequest) SetLang(v string) *ListDiagnoseReportIdsRequest {
	s.Lang = &v
	return s
}

func (s *ListDiagnoseReportIdsRequest) SetPage(v int32) *ListDiagnoseReportIdsRequest {
	s.Page = &v
	return s
}

func (s *ListDiagnoseReportIdsRequest) SetSize(v int32) *ListDiagnoseReportIdsRequest {
	s.Size = &v
	return s
}

func (s *ListDiagnoseReportIdsRequest) SetStartTime(v int64) *ListDiagnoseReportIdsRequest {
	s.StartTime = &v
	return s
}

func (s *ListDiagnoseReportIdsRequest) SetTrigger(v string) *ListDiagnoseReportIdsRequest {
	s.Trigger = &v
	return s
}

func (s *ListDiagnoseReportIdsRequest) Validate() error {
	return dara.Validate(s)
}
