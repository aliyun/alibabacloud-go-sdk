// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetReportTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetReportTemplateId(v string) *GetReportTemplateRequest
	GetReportTemplateId() *string
}

type GetReportTemplateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// crt-xxx
	ReportTemplateId *string `json:"ReportTemplateId,omitempty" xml:"ReportTemplateId,omitempty"`
}

func (s GetReportTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetReportTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetReportTemplateRequest) GetReportTemplateId() *string {
	return s.ReportTemplateId
}

func (s *GetReportTemplateRequest) SetReportTemplateId(v string) *GetReportTemplateRequest {
	s.ReportTemplateId = &v
	return s
}

func (s *GetReportTemplateRequest) Validate() error {
	return dara.Validate(s)
}
