// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDatabaseInstanceMetricDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDatabaseInstanceMetricDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDatabaseInstanceMetricDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDatabaseInstanceMetricDataResponseBody) *DescribeDatabaseInstanceMetricDataResponse
	GetBody() *DescribeDatabaseInstanceMetricDataResponseBody
}

type DescribeDatabaseInstanceMetricDataResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDatabaseInstanceMetricDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDatabaseInstanceMetricDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDatabaseInstanceMetricDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDatabaseInstanceMetricDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDatabaseInstanceMetricDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDatabaseInstanceMetricDataResponse) GetBody() *DescribeDatabaseInstanceMetricDataResponseBody {
	return s.Body
}

func (s *DescribeDatabaseInstanceMetricDataResponse) SetHeaders(v map[string]*string) *DescribeDatabaseInstanceMetricDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDatabaseInstanceMetricDataResponse) SetStatusCode(v int32) *DescribeDatabaseInstanceMetricDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDatabaseInstanceMetricDataResponse) SetBody(v *DescribeDatabaseInstanceMetricDataResponseBody) *DescribeDatabaseInstanceMetricDataResponse {
	s.Body = v
	return s
}

func (s *DescribeDatabaseInstanceMetricDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
