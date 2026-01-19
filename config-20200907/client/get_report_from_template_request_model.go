// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetReportFromTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetReportTemplateId(v string) *GetReportFromTemplateRequest
	GetReportTemplateId() *string
}

type GetReportFromTemplateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// crt-xxx
	ReportTemplateId *string `json:"ReportTemplateId,omitempty" xml:"ReportTemplateId,omitempty"`
}

func (s GetReportFromTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetReportFromTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetReportFromTemplateRequest) GetReportTemplateId() *string {
	return s.ReportTemplateId
}

func (s *GetReportFromTemplateRequest) SetReportTemplateId(v string) *GetReportFromTemplateRequest {
	s.ReportTemplateId = &v
	return s
}

func (s *GetReportFromTemplateRequest) Validate() error {
	return dara.Validate(s)
}
