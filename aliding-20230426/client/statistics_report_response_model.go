// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStatisticsReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StatisticsReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StatisticsReportResponse
	GetStatusCode() *int32
	SetBody(v *StatisticsReportResponseBody) *StatisticsReportResponse
	GetBody() *StatisticsReportResponseBody
}

type StatisticsReportResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StatisticsReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StatisticsReportResponse) String() string {
	return dara.Prettify(s)
}

func (s StatisticsReportResponse) GoString() string {
	return s.String()
}

func (s *StatisticsReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StatisticsReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StatisticsReportResponse) GetBody() *StatisticsReportResponseBody {
	return s.Body
}

func (s *StatisticsReportResponse) SetHeaders(v map[string]*string) *StatisticsReportResponse {
	s.Headers = v
	return s
}

func (s *StatisticsReportResponse) SetStatusCode(v int32) *StatisticsReportResponse {
	s.StatusCode = &v
	return s
}

func (s *StatisticsReportResponse) SetBody(v *StatisticsReportResponseBody) *StatisticsReportResponse {
	s.Body = v
	return s
}

func (s *StatisticsReportResponse) Validate() error {
	return dara.Validate(s)
}
