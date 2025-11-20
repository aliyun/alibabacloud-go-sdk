// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStatisticsListByTypeReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StatisticsListByTypeReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StatisticsListByTypeReportResponse
	GetStatusCode() *int32
	SetBody(v *StatisticsListByTypeReportResponseBody) *StatisticsListByTypeReportResponse
	GetBody() *StatisticsListByTypeReportResponseBody
}

type StatisticsListByTypeReportResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StatisticsListByTypeReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StatisticsListByTypeReportResponse) String() string {
	return dara.Prettify(s)
}

func (s StatisticsListByTypeReportResponse) GoString() string {
	return s.String()
}

func (s *StatisticsListByTypeReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StatisticsListByTypeReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StatisticsListByTypeReportResponse) GetBody() *StatisticsListByTypeReportResponseBody {
	return s.Body
}

func (s *StatisticsListByTypeReportResponse) SetHeaders(v map[string]*string) *StatisticsListByTypeReportResponse {
	s.Headers = v
	return s
}

func (s *StatisticsListByTypeReportResponse) SetStatusCode(v int32) *StatisticsListByTypeReportResponse {
	s.StatusCode = &v
	return s
}

func (s *StatisticsListByTypeReportResponse) SetBody(v *StatisticsListByTypeReportResponseBody) *StatisticsListByTypeReportResponse {
	s.Body = v
	return s
}

func (s *StatisticsListByTypeReportResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
