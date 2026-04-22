// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePvtzStatisticsGlobalOverviewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePvtzStatisticsGlobalOverviewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePvtzStatisticsGlobalOverviewResponse
	GetStatusCode() *int32
	SetBody(v *DescribePvtzStatisticsGlobalOverviewResponseBody) *DescribePvtzStatisticsGlobalOverviewResponse
	GetBody() *DescribePvtzStatisticsGlobalOverviewResponseBody
}

type DescribePvtzStatisticsGlobalOverviewResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePvtzStatisticsGlobalOverviewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePvtzStatisticsGlobalOverviewResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePvtzStatisticsGlobalOverviewResponse) GoString() string {
	return s.String()
}

func (s *DescribePvtzStatisticsGlobalOverviewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePvtzStatisticsGlobalOverviewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePvtzStatisticsGlobalOverviewResponse) GetBody() *DescribePvtzStatisticsGlobalOverviewResponseBody {
	return s.Body
}

func (s *DescribePvtzStatisticsGlobalOverviewResponse) SetHeaders(v map[string]*string) *DescribePvtzStatisticsGlobalOverviewResponse {
	s.Headers = v
	return s
}

func (s *DescribePvtzStatisticsGlobalOverviewResponse) SetStatusCode(v int32) *DescribePvtzStatisticsGlobalOverviewResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePvtzStatisticsGlobalOverviewResponse) SetBody(v *DescribePvtzStatisticsGlobalOverviewResponseBody) *DescribePvtzStatisticsGlobalOverviewResponse {
	s.Body = v
	return s
}

func (s *DescribePvtzStatisticsGlobalOverviewResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
