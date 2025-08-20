// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAppUseTimeReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AppUseTimeReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AppUseTimeReportResponse
	GetStatusCode() *int32
	SetBody(v *AppUseTimeReportResponseBody) *AppUseTimeReportResponse
	GetBody() *AppUseTimeReportResponseBody
}

type AppUseTimeReportResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AppUseTimeReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AppUseTimeReportResponse) String() string {
	return dara.Prettify(s)
}

func (s AppUseTimeReportResponse) GoString() string {
	return s.String()
}

func (s *AppUseTimeReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AppUseTimeReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AppUseTimeReportResponse) GetBody() *AppUseTimeReportResponseBody {
	return s.Body
}

func (s *AppUseTimeReportResponse) SetHeaders(v map[string]*string) *AppUseTimeReportResponse {
	s.Headers = v
	return s
}

func (s *AppUseTimeReportResponse) SetStatusCode(v int32) *AppUseTimeReportResponse {
	s.StatusCode = &v
	return s
}

func (s *AppUseTimeReportResponse) SetBody(v *AppUseTimeReportResponseBody) *AppUseTimeReportResponse {
	s.Body = v
	return s
}

func (s *AppUseTimeReportResponse) Validate() error {
	return dara.Validate(s)
}
