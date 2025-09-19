// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendCustomizeReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SendCustomizeReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SendCustomizeReportResponse
	GetStatusCode() *int32
	SetBody(v *SendCustomizeReportResponseBody) *SendCustomizeReportResponse
	GetBody() *SendCustomizeReportResponseBody
}

type SendCustomizeReportResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendCustomizeReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendCustomizeReportResponse) String() string {
	return dara.Prettify(s)
}

func (s SendCustomizeReportResponse) GoString() string {
	return s.String()
}

func (s *SendCustomizeReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SendCustomizeReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SendCustomizeReportResponse) GetBody() *SendCustomizeReportResponseBody {
	return s.Body
}

func (s *SendCustomizeReportResponse) SetHeaders(v map[string]*string) *SendCustomizeReportResponse {
	s.Headers = v
	return s
}

func (s *SendCustomizeReportResponse) SetStatusCode(v int32) *SendCustomizeReportResponse {
	s.StatusCode = &v
	return s
}

func (s *SendCustomizeReportResponse) SetBody(v *SendCustomizeReportResponseBody) *SendCustomizeReportResponse {
	s.Body = v
	return s
}

func (s *SendCustomizeReportResponse) Validate() error {
	return dara.Validate(s)
}
