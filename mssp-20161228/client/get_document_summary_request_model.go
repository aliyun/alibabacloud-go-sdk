// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDocumentSummaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetReportType(v string) *GetDocumentSummaryRequest
	GetReportType() *string
}

type GetDocumentSummaryRequest struct {
	// Type of service report.
	//
	// example:
	//
	// 1
	ReportType *string `json:"ReportType,omitempty" xml:"ReportType,omitempty"`
}

func (s GetDocumentSummaryRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDocumentSummaryRequest) GoString() string {
	return s.String()
}

func (s *GetDocumentSummaryRequest) GetReportType() *string {
	return s.ReportType
}

func (s *GetDocumentSummaryRequest) SetReportType(v string) *GetDocumentSummaryRequest {
	s.ReportType = &v
	return s
}

func (s *GetDocumentSummaryRequest) Validate() error {
	return dara.Validate(s)
}
