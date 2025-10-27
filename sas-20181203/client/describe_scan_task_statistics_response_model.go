// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScanTaskStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeScanTaskStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeScanTaskStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeScanTaskStatisticsResponseBody) *DescribeScanTaskStatisticsResponse
	GetBody() *DescribeScanTaskStatisticsResponseBody
}

type DescribeScanTaskStatisticsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeScanTaskStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeScanTaskStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeScanTaskStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeScanTaskStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeScanTaskStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeScanTaskStatisticsResponse) GetBody() *DescribeScanTaskStatisticsResponseBody {
	return s.Body
}

func (s *DescribeScanTaskStatisticsResponse) SetHeaders(v map[string]*string) *DescribeScanTaskStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeScanTaskStatisticsResponse) SetStatusCode(v int32) *DescribeScanTaskStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScanTaskStatisticsResponse) SetBody(v *DescribeScanTaskStatisticsResponseBody) *DescribeScanTaskStatisticsResponse {
	s.Body = v
	return s
}

func (s *DescribeScanTaskStatisticsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
