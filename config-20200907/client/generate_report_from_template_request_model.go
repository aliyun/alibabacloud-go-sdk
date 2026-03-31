// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateReportFromTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetReportTemplateId(v string) *GenerateReportFromTemplateRequest
	GetReportTemplateId() *string
}

type GenerateReportFromTemplateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// crt-xxx
	ReportTemplateId *string `json:"ReportTemplateId,omitempty" xml:"ReportTemplateId,omitempty"`
}

func (s GenerateReportFromTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateReportFromTemplateRequest) GoString() string {
	return s.String()
}

func (s *GenerateReportFromTemplateRequest) GetReportTemplateId() *string {
	return s.ReportTemplateId
}

func (s *GenerateReportFromTemplateRequest) SetReportTemplateId(v string) *GenerateReportFromTemplateRequest {
	s.ReportTemplateId = &v
	return s
}

func (s *GenerateReportFromTemplateRequest) Validate() error {
	return dara.Validate(s)
}
