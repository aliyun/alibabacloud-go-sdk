// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAiReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAiReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAiReportResponse
	GetStatusCode() *int32
	SetBody(v *GetAiReportResponseBody) *GetAiReportResponse
	GetBody() *GetAiReportResponseBody
}

type GetAiReportResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAiReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAiReportResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAiReportResponse) GoString() string {
	return s.String()
}

func (s *GetAiReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAiReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAiReportResponse) GetBody() *GetAiReportResponseBody {
	return s.Body
}

func (s *GetAiReportResponse) SetHeaders(v map[string]*string) *GetAiReportResponse {
	s.Headers = v
	return s
}

func (s *GetAiReportResponse) SetStatusCode(v int32) *GetAiReportResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAiReportResponse) SetBody(v *GetAiReportResponseBody) *GetAiReportResponse {
	s.Body = v
	return s
}

func (s *GetAiReportResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
