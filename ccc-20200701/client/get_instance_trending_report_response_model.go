// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceTrendingReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInstanceTrendingReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInstanceTrendingReportResponse
	GetStatusCode() *int32
	SetBody(v *GetInstanceTrendingReportResponseBody) *GetInstanceTrendingReportResponse
	GetBody() *GetInstanceTrendingReportResponseBody
}

type GetInstanceTrendingReportResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceTrendingReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceTrendingReportResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceTrendingReportResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceTrendingReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInstanceTrendingReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInstanceTrendingReportResponse) GetBody() *GetInstanceTrendingReportResponseBody {
	return s.Body
}

func (s *GetInstanceTrendingReportResponse) SetHeaders(v map[string]*string) *GetInstanceTrendingReportResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceTrendingReportResponse) SetStatusCode(v int32) *GetInstanceTrendingReportResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceTrendingReportResponse) SetBody(v *GetInstanceTrendingReportResponseBody) *GetInstanceTrendingReportResponse {
	s.Body = v
	return s
}

func (s *GetInstanceTrendingReportResponse) Validate() error {
	return dara.Validate(s)
}
