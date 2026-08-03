// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTrailDeliveryMetricDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTrailDeliveryMetricDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTrailDeliveryMetricDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTrailDeliveryMetricDataResponseBody) *DescribeTrailDeliveryMetricDataResponse
	GetBody() *DescribeTrailDeliveryMetricDataResponseBody
}

type DescribeTrailDeliveryMetricDataResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTrailDeliveryMetricDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTrailDeliveryMetricDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTrailDeliveryMetricDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeTrailDeliveryMetricDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTrailDeliveryMetricDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTrailDeliveryMetricDataResponse) GetBody() *DescribeTrailDeliveryMetricDataResponseBody {
	return s.Body
}

func (s *DescribeTrailDeliveryMetricDataResponse) SetHeaders(v map[string]*string) *DescribeTrailDeliveryMetricDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeTrailDeliveryMetricDataResponse) SetStatusCode(v int32) *DescribeTrailDeliveryMetricDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTrailDeliveryMetricDataResponse) SetBody(v *DescribeTrailDeliveryMetricDataResponseBody) *DescribeTrailDeliveryMetricDataResponse {
	s.Body = v
	return s
}

func (s *DescribeTrailDeliveryMetricDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
