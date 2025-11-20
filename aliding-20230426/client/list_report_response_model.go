// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListReportResponse
	GetStatusCode() *int32
	SetBody(v *ListReportResponseBody) *ListReportResponse
	GetBody() *ListReportResponseBody
}

type ListReportResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListReportResponse) String() string {
	return dara.Prettify(s)
}

func (s ListReportResponse) GoString() string {
	return s.String()
}

func (s *ListReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListReportResponse) GetBody() *ListReportResponseBody {
	return s.Body
}

func (s *ListReportResponse) SetHeaders(v map[string]*string) *ListReportResponse {
	s.Headers = v
	return s
}

func (s *ListReportResponse) SetStatusCode(v int32) *ListReportResponse {
	s.StatusCode = &v
	return s
}

func (s *ListReportResponse) SetBody(v *ListReportResponseBody) *ListReportResponse {
	s.Body = v
	return s
}

func (s *ListReportResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
