// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteReportTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetReportTemplateId(v string) *DeleteReportTemplateResponseBody
	GetReportTemplateId() *string
	SetRequestId(v string) *DeleteReportTemplateResponseBody
	GetRequestId() *string
}

type DeleteReportTemplateResponseBody struct {
	// example:
	//
	// crt-xxx
	ReportTemplateId *string `json:"ReportTemplateId,omitempty" xml:"ReportTemplateId,omitempty"`
	// example:
	//
	// DE9FFFE5-FCAD-4B24-9546-BF49273C562B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteReportTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteReportTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteReportTemplateResponseBody) GetReportTemplateId() *string {
	return s.ReportTemplateId
}

func (s *DeleteReportTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteReportTemplateResponseBody) SetReportTemplateId(v string) *DeleteReportTemplateResponseBody {
	s.ReportTemplateId = &v
	return s
}

func (s *DeleteReportTemplateResponseBody) SetRequestId(v string) *DeleteReportTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteReportTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
