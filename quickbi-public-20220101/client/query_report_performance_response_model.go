// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryReportPerformanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryReportPerformanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryReportPerformanceResponse
	GetStatusCode() *int32
	SetBody(v *QueryReportPerformanceResponseBody) *QueryReportPerformanceResponse
	GetBody() *QueryReportPerformanceResponseBody
}

type QueryReportPerformanceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryReportPerformanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryReportPerformanceResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryReportPerformanceResponse) GoString() string {
	return s.String()
}

func (s *QueryReportPerformanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryReportPerformanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryReportPerformanceResponse) GetBody() *QueryReportPerformanceResponseBody {
	return s.Body
}

func (s *QueryReportPerformanceResponse) SetHeaders(v map[string]*string) *QueryReportPerformanceResponse {
	s.Headers = v
	return s
}

func (s *QueryReportPerformanceResponse) SetStatusCode(v int32) *QueryReportPerformanceResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryReportPerformanceResponse) SetBody(v *QueryReportPerformanceResponseBody) *QueryReportPerformanceResponse {
	s.Body = v
	return s
}

func (s *QueryReportPerformanceResponse) Validate() error {
	return dara.Validate(s)
}
