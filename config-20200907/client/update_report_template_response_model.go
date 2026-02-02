// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateReportTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateReportTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateReportTemplateResponse
	GetStatusCode() *int32
	SetBody(v *UpdateReportTemplateResponseBody) *UpdateReportTemplateResponse
	GetBody() *UpdateReportTemplateResponseBody
}

type UpdateReportTemplateResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateReportTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateReportTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateReportTemplateResponse) GoString() string {
	return s.String()
}

func (s *UpdateReportTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateReportTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateReportTemplateResponse) GetBody() *UpdateReportTemplateResponseBody {
	return s.Body
}

func (s *UpdateReportTemplateResponse) SetHeaders(v map[string]*string) *UpdateReportTemplateResponse {
	s.Headers = v
	return s
}

func (s *UpdateReportTemplateResponse) SetStatusCode(v int32) *UpdateReportTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateReportTemplateResponse) SetBody(v *UpdateReportTemplateResponseBody) *UpdateReportTemplateResponse {
	s.Body = v
	return s
}

func (s *UpdateReportTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
