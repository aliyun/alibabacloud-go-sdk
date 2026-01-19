// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetReportFromTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetReportFromTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetReportFromTemplateResponse
	GetStatusCode() *int32
	SetBody(v *GetReportFromTemplateResponseBody) *GetReportFromTemplateResponse
	GetBody() *GetReportFromTemplateResponseBody
}

type GetReportFromTemplateResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetReportFromTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetReportFromTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s GetReportFromTemplateResponse) GoString() string {
	return s.String()
}

func (s *GetReportFromTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetReportFromTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetReportFromTemplateResponse) GetBody() *GetReportFromTemplateResponseBody {
	return s.Body
}

func (s *GetReportFromTemplateResponse) SetHeaders(v map[string]*string) *GetReportFromTemplateResponse {
	s.Headers = v
	return s
}

func (s *GetReportFromTemplateResponse) SetStatusCode(v int32) *GetReportFromTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetReportFromTemplateResponse) SetBody(v *GetReportFromTemplateResponseBody) *GetReportFromTemplateResponse {
	s.Body = v
	return s
}

func (s *GetReportFromTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
