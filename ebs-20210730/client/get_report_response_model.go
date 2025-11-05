// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetReportResponse
	GetStatusCode() *int32
	SetBody(v *GetReportResponseBody) *GetReportResponse
	GetBody() *GetReportResponseBody
}

type GetReportResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetReportResponse) String() string {
	return dara.Prettify(s)
}

func (s GetReportResponse) GoString() string {
	return s.String()
}

func (s *GetReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetReportResponse) GetBody() *GetReportResponseBody {
	return s.Body
}

func (s *GetReportResponse) SetHeaders(v map[string]*string) *GetReportResponse {
	s.Headers = v
	return s
}

func (s *GetReportResponse) SetStatusCode(v int32) *GetReportResponse {
	s.StatusCode = &v
	return s
}

func (s *GetReportResponse) SetBody(v *GetReportResponseBody) *GetReportResponse {
	s.Body = v
	return s
}

func (s *GetReportResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
