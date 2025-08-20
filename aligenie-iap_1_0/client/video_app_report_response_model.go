// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVideoAppReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *VideoAppReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *VideoAppReportResponse
	GetStatusCode() *int32
	SetBody(v *VideoAppReportResponseBody) *VideoAppReportResponse
	GetBody() *VideoAppReportResponseBody
}

type VideoAppReportResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VideoAppReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VideoAppReportResponse) String() string {
	return dara.Prettify(s)
}

func (s VideoAppReportResponse) GoString() string {
	return s.String()
}

func (s *VideoAppReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *VideoAppReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *VideoAppReportResponse) GetBody() *VideoAppReportResponseBody {
	return s.Body
}

func (s *VideoAppReportResponse) SetHeaders(v map[string]*string) *VideoAppReportResponse {
	s.Headers = v
	return s
}

func (s *VideoAppReportResponse) SetStatusCode(v int32) *VideoAppReportResponse {
	s.StatusCode = &v
	return s
}

func (s *VideoAppReportResponse) SetBody(v *VideoAppReportResponseBody) *VideoAppReportResponse {
	s.Body = v
	return s
}

func (s *VideoAppReportResponse) Validate() error {
	return dara.Validate(s)
}
