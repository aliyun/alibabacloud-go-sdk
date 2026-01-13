// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeThreatEventTopMetricResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeThreatEventTopMetricResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeThreatEventTopMetricResponse
	GetStatusCode() *int32
	SetBody(v *DescribeThreatEventTopMetricResponseBody) *DescribeThreatEventTopMetricResponse
	GetBody() *DescribeThreatEventTopMetricResponseBody
}

type DescribeThreatEventTopMetricResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeThreatEventTopMetricResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeThreatEventTopMetricResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeThreatEventTopMetricResponse) GoString() string {
	return s.String()
}

func (s *DescribeThreatEventTopMetricResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeThreatEventTopMetricResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeThreatEventTopMetricResponse) GetBody() *DescribeThreatEventTopMetricResponseBody {
	return s.Body
}

func (s *DescribeThreatEventTopMetricResponse) SetHeaders(v map[string]*string) *DescribeThreatEventTopMetricResponse {
	s.Headers = v
	return s
}

func (s *DescribeThreatEventTopMetricResponse) SetStatusCode(v int32) *DescribeThreatEventTopMetricResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeThreatEventTopMetricResponse) SetBody(v *DescribeThreatEventTopMetricResponseBody) *DescribeThreatEventTopMetricResponse {
	s.Body = v
	return s
}

func (s *DescribeThreatEventTopMetricResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
