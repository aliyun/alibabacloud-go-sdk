// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateReportTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetReportTemplateId(v string) *UpdateReportTemplateResponseBody
	GetReportTemplateId() *string
	SetRequestId(v string) *UpdateReportTemplateResponseBody
	GetRequestId() *string
}

type UpdateReportTemplateResponseBody struct {
	// example:
	//
	// crt-xxx
	ReportTemplateId *string `json:"ReportTemplateId,omitempty" xml:"ReportTemplateId,omitempty"`
	// example:
	//
	// A7A0FFF8-0B44-40C6-8BBF-3A185EFDF3F7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateReportTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateReportTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateReportTemplateResponseBody) GetReportTemplateId() *string {
	return s.ReportTemplateId
}

func (s *UpdateReportTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateReportTemplateResponseBody) SetReportTemplateId(v string) *UpdateReportTemplateResponseBody {
	s.ReportTemplateId = &v
	return s
}

func (s *UpdateReportTemplateResponseBody) SetRequestId(v string) *UpdateReportTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateReportTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
