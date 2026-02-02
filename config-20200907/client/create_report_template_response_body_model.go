// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateReportTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetReportTemplateId(v string) *CreateReportTemplateResponseBody
	GetReportTemplateId() *string
	SetRequestId(v string) *CreateReportTemplateResponseBody
	GetRequestId() *string
}

type CreateReportTemplateResponseBody struct {
	// example:
	//
	// crt-xxx
	ReportTemplateId *string `json:"ReportTemplateId,omitempty" xml:"ReportTemplateId,omitempty"`
	// example:
	//
	// A7A0FFF8-0B44-40C6-8BBF-3A185EFDFXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateReportTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateReportTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateReportTemplateResponseBody) GetReportTemplateId() *string {
	return s.ReportTemplateId
}

func (s *CreateReportTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateReportTemplateResponseBody) SetReportTemplateId(v string) *CreateReportTemplateResponseBody {
	s.ReportTemplateId = &v
	return s
}

func (s *CreateReportTemplateResponseBody) SetRequestId(v string) *CreateReportTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateReportTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
