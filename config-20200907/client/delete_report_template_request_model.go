// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteReportTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetReportTemplateId(v string) *DeleteReportTemplateRequest
	GetReportTemplateId() *string
}

type DeleteReportTemplateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// crt-xxx
	ReportTemplateId *string `json:"ReportTemplateId,omitempty" xml:"ReportTemplateId,omitempty"`
}

func (s DeleteReportTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteReportTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteReportTemplateRequest) GetReportTemplateId() *string {
	return s.ReportTemplateId
}

func (s *DeleteReportTemplateRequest) SetReportTemplateId(v string) *DeleteReportTemplateRequest {
	s.ReportTemplateId = &v
	return s
}

func (s *DeleteReportTemplateRequest) Validate() error {
	return dara.Validate(s)
}
