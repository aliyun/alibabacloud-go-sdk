// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PushReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PushReportResponse
	GetStatusCode() *int32
	SetBody(v *PushReportResponseBody) *PushReportResponse
	GetBody() *PushReportResponseBody
}

type PushReportResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PushReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PushReportResponse) String() string {
	return dara.Prettify(s)
}

func (s PushReportResponse) GoString() string {
	return s.String()
}

func (s *PushReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PushReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PushReportResponse) GetBody() *PushReportResponseBody {
	return s.Body
}

func (s *PushReportResponse) SetHeaders(v map[string]*string) *PushReportResponse {
	s.Headers = v
	return s
}

func (s *PushReportResponse) SetStatusCode(v int32) *PushReportResponse {
	s.StatusCode = &v
	return s
}

func (s *PushReportResponse) SetBody(v *PushReportResponseBody) *PushReportResponse {
	s.Body = v
	return s
}

func (s *PushReportResponse) Validate() error {
	return dara.Validate(s)
}
