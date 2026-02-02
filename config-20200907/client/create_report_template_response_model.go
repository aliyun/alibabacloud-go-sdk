// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateReportTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateReportTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateReportTemplateResponse
	GetStatusCode() *int32
	SetBody(v *CreateReportTemplateResponseBody) *CreateReportTemplateResponse
	GetBody() *CreateReportTemplateResponseBody
}

type CreateReportTemplateResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateReportTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateReportTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateReportTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateReportTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateReportTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateReportTemplateResponse) GetBody() *CreateReportTemplateResponseBody {
	return s.Body
}

func (s *CreateReportTemplateResponse) SetHeaders(v map[string]*string) *CreateReportTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateReportTemplateResponse) SetStatusCode(v int32) *CreateReportTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateReportTemplateResponse) SetBody(v *CreateReportTemplateResponseBody) *CreateReportTemplateResponse {
	s.Body = v
	return s
}

func (s *CreateReportTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
