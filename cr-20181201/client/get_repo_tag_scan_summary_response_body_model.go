// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRepoTagScanSummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetRepoTagScanSummaryResponseBody
	GetCode() *string
	SetHighSeverity(v int32) *GetRepoTagScanSummaryResponseBody
	GetHighSeverity() *int32
	SetIsSuccess(v bool) *GetRepoTagScanSummaryResponseBody
	GetIsSuccess() *bool
	SetLowSeverity(v int32) *GetRepoTagScanSummaryResponseBody
	GetLowSeverity() *int32
	SetMediumSeverity(v int32) *GetRepoTagScanSummaryResponseBody
	GetMediumSeverity() *int32
	SetRequestId(v string) *GetRepoTagScanSummaryResponseBody
	GetRequestId() *string
	SetTotalSeverity(v int32) *GetRepoTagScanSummaryResponseBody
	GetTotalSeverity() *int32
	SetUnknownSeverity(v int32) *GetRepoTagScanSummaryResponseBody
	GetUnknownSeverity() *int32
}

type GetRepoTagScanSummaryResponseBody struct {
	// The number of medium-severity vulnerabilities.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The number of low-severity vulnerabilities.
	//
	// example:
	//
	// 22
	HighSeverity *int32 `json:"HighSeverity,omitempty" xml:"HighSeverity,omitempty"`
	// The number of high-severity vulnerabilities.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// example:
	//
	// 89
	LowSeverity *int32 `json:"LowSeverity,omitempty" xml:"LowSeverity,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- `true`: The request is successful.
	//
	// 	- `false`: The request fails.
	//
	// example:
	//
	// 81
	MediumSeverity *int32 `json:"MediumSeverity,omitempty" xml:"MediumSeverity,omitempty"`
	// The total number of vulnerabilities detected on images.
	//
	// example:
	//
	// BC648259-91A7-4502-BED3-EDF64361FA83
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The return value.
	//
	// example:
	//
	// 196
	TotalSeverity *int32 `json:"TotalSeverity,omitempty" xml:"TotalSeverity,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4
	UnknownSeverity *int32 `json:"UnknownSeverity,omitempty" xml:"UnknownSeverity,omitempty"`
}

func (s GetRepoTagScanSummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRepoTagScanSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *GetRepoTagScanSummaryResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetRepoTagScanSummaryResponseBody) GetHighSeverity() *int32 {
	return s.HighSeverity
}

func (s *GetRepoTagScanSummaryResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *GetRepoTagScanSummaryResponseBody) GetLowSeverity() *int32 {
	return s.LowSeverity
}

func (s *GetRepoTagScanSummaryResponseBody) GetMediumSeverity() *int32 {
	return s.MediumSeverity
}

func (s *GetRepoTagScanSummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRepoTagScanSummaryResponseBody) GetTotalSeverity() *int32 {
	return s.TotalSeverity
}

func (s *GetRepoTagScanSummaryResponseBody) GetUnknownSeverity() *int32 {
	return s.UnknownSeverity
}

func (s *GetRepoTagScanSummaryResponseBody) SetCode(v string) *GetRepoTagScanSummaryResponseBody {
	s.Code = &v
	return s
}

func (s *GetRepoTagScanSummaryResponseBody) SetHighSeverity(v int32) *GetRepoTagScanSummaryResponseBody {
	s.HighSeverity = &v
	return s
}

func (s *GetRepoTagScanSummaryResponseBody) SetIsSuccess(v bool) *GetRepoTagScanSummaryResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *GetRepoTagScanSummaryResponseBody) SetLowSeverity(v int32) *GetRepoTagScanSummaryResponseBody {
	s.LowSeverity = &v
	return s
}

func (s *GetRepoTagScanSummaryResponseBody) SetMediumSeverity(v int32) *GetRepoTagScanSummaryResponseBody {
	s.MediumSeverity = &v
	return s
}

func (s *GetRepoTagScanSummaryResponseBody) SetRequestId(v string) *GetRepoTagScanSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRepoTagScanSummaryResponseBody) SetTotalSeverity(v int32) *GetRepoTagScanSummaryResponseBody {
	s.TotalSeverity = &v
	return s
}

func (s *GetRepoTagScanSummaryResponseBody) SetUnknownSeverity(v int32) *GetRepoTagScanSummaryResponseBody {
	s.UnknownSeverity = &v
	return s
}

func (s *GetRepoTagScanSummaryResponseBody) Validate() error {
	return dara.Validate(s)
}
