// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetReportTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetReportTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetReportTemplateResponse
	GetStatusCode() *int32
	SetBody(v *GetReportTemplateResponseBody) *GetReportTemplateResponse
	GetBody() *GetReportTemplateResponseBody
}

type GetReportTemplateResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetReportTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetReportTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s GetReportTemplateResponse) GoString() string {
	return s.String()
}

func (s *GetReportTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetReportTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetReportTemplateResponse) GetBody() *GetReportTemplateResponseBody {
	return s.Body
}

func (s *GetReportTemplateResponse) SetHeaders(v map[string]*string) *GetReportTemplateResponse {
	s.Headers = v
	return s
}

func (s *GetReportTemplateResponse) SetStatusCode(v int32) *GetReportTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetReportTemplateResponse) SetBody(v *GetReportTemplateResponseBody) *GetReportTemplateResponse {
	s.Body = v
	return s
}

func (s *GetReportTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
