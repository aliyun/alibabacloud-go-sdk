// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePvtzStatisticsZoneOverviewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePvtzStatisticsZoneOverviewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePvtzStatisticsZoneOverviewResponse
	GetStatusCode() *int32
	SetBody(v *DescribePvtzStatisticsZoneOverviewResponseBody) *DescribePvtzStatisticsZoneOverviewResponse
	GetBody() *DescribePvtzStatisticsZoneOverviewResponseBody
}

type DescribePvtzStatisticsZoneOverviewResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePvtzStatisticsZoneOverviewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePvtzStatisticsZoneOverviewResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePvtzStatisticsZoneOverviewResponse) GoString() string {
	return s.String()
}

func (s *DescribePvtzStatisticsZoneOverviewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePvtzStatisticsZoneOverviewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePvtzStatisticsZoneOverviewResponse) GetBody() *DescribePvtzStatisticsZoneOverviewResponseBody {
	return s.Body
}

func (s *DescribePvtzStatisticsZoneOverviewResponse) SetHeaders(v map[string]*string) *DescribePvtzStatisticsZoneOverviewResponse {
	s.Headers = v
	return s
}

func (s *DescribePvtzStatisticsZoneOverviewResponse) SetStatusCode(v int32) *DescribePvtzStatisticsZoneOverviewResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePvtzStatisticsZoneOverviewResponse) SetBody(v *DescribePvtzStatisticsZoneOverviewResponseBody) *DescribePvtzStatisticsZoneOverviewResponse {
	s.Body = v
	return s
}

func (s *DescribePvtzStatisticsZoneOverviewResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
