// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAiReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallId(v string) *GetAiReportRequest
	GetCallId() *string
	SetReportId(v string) *GetAiReportRequest
	GetReportId() *string
}

type GetAiReportRequest struct {
	// example:
	//
	// call-3f2b8c1e-6d4a-4b7e-9a1c-2e5f7d8b9c0a
	CallId *string `json:"callId,omitempty" xml:"callId,omitempty"`
	// example:
	//
	// report-3f2b8c1e-6d4a-4b7e-9a1c-2e5f7d8b9c0a
	ReportId *string `json:"reportId,omitempty" xml:"reportId,omitempty"`
}

func (s GetAiReportRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAiReportRequest) GoString() string {
	return s.String()
}

func (s *GetAiReportRequest) GetCallId() *string {
	return s.CallId
}

func (s *GetAiReportRequest) GetReportId() *string {
	return s.ReportId
}

func (s *GetAiReportRequest) SetCallId(v string) *GetAiReportRequest {
	s.CallId = &v
	return s
}

func (s *GetAiReportRequest) SetReportId(v string) *GetAiReportRequest {
	s.ReportId = &v
	return s
}

func (s *GetAiReportRequest) Validate() error {
	return dara.Validate(s)
}
