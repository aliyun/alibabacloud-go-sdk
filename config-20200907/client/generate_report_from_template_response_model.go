// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateReportFromTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateReportFromTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateReportFromTemplateResponse
	GetStatusCode() *int32
	SetBody(v *GenerateReportFromTemplateResponseBody) *GenerateReportFromTemplateResponse
	GetBody() *GenerateReportFromTemplateResponseBody
}

type GenerateReportFromTemplateResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateReportFromTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateReportFromTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateReportFromTemplateResponse) GoString() string {
	return s.String()
}

func (s *GenerateReportFromTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateReportFromTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateReportFromTemplateResponse) GetBody() *GenerateReportFromTemplateResponseBody {
	return s.Body
}

func (s *GenerateReportFromTemplateResponse) SetHeaders(v map[string]*string) *GenerateReportFromTemplateResponse {
	s.Headers = v
	return s
}

func (s *GenerateReportFromTemplateResponse) SetStatusCode(v int32) *GenerateReportFromTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateReportFromTemplateResponse) SetBody(v *GenerateReportFromTemplateResponseBody) *GenerateReportFromTemplateResponse {
	s.Body = v
	return s
}

func (s *GenerateReportFromTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
