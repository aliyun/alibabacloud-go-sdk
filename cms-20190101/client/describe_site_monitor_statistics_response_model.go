// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSiteMonitorStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSiteMonitorStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSiteMonitorStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSiteMonitorStatisticsResponseBody) *DescribeSiteMonitorStatisticsResponse
	GetBody() *DescribeSiteMonitorStatisticsResponseBody
}

type DescribeSiteMonitorStatisticsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSiteMonitorStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSiteMonitorStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteMonitorStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSiteMonitorStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSiteMonitorStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSiteMonitorStatisticsResponse) GetBody() *DescribeSiteMonitorStatisticsResponseBody {
	return s.Body
}

func (s *DescribeSiteMonitorStatisticsResponse) SetHeaders(v map[string]*string) *DescribeSiteMonitorStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSiteMonitorStatisticsResponse) SetStatusCode(v int32) *DescribeSiteMonitorStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSiteMonitorStatisticsResponse) SetBody(v *DescribeSiteMonitorStatisticsResponseBody) *DescribeSiteMonitorStatisticsResponse {
	s.Body = v
	return s
}

func (s *DescribeSiteMonitorStatisticsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
