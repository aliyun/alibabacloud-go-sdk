// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeElasticDailyPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeElasticDailyPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeElasticDailyPlanResponse
	GetStatusCode() *int32
	SetBody(v *DescribeElasticDailyPlanResponseBody) *DescribeElasticDailyPlanResponse
	GetBody() *DescribeElasticDailyPlanResponseBody
}

type DescribeElasticDailyPlanResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeElasticDailyPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeElasticDailyPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeElasticDailyPlanResponse) GoString() string {
	return s.String()
}

func (s *DescribeElasticDailyPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeElasticDailyPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeElasticDailyPlanResponse) GetBody() *DescribeElasticDailyPlanResponseBody {
	return s.Body
}

func (s *DescribeElasticDailyPlanResponse) SetHeaders(v map[string]*string) *DescribeElasticDailyPlanResponse {
	s.Headers = v
	return s
}

func (s *DescribeElasticDailyPlanResponse) SetStatusCode(v int32) *DescribeElasticDailyPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeElasticDailyPlanResponse) SetBody(v *DescribeElasticDailyPlanResponseBody) *DescribeElasticDailyPlanResponse {
	s.Body = v
	return s
}

func (s *DescribeElasticDailyPlanResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
