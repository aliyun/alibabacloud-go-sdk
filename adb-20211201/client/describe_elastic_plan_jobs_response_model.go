// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeElasticPlanJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeElasticPlanJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeElasticPlanJobsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeElasticPlanJobsResponseBody) *DescribeElasticPlanJobsResponse
	GetBody() *DescribeElasticPlanJobsResponseBody
}

type DescribeElasticPlanJobsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeElasticPlanJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeElasticPlanJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeElasticPlanJobsResponse) GoString() string {
	return s.String()
}

func (s *DescribeElasticPlanJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeElasticPlanJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeElasticPlanJobsResponse) GetBody() *DescribeElasticPlanJobsResponseBody {
	return s.Body
}

func (s *DescribeElasticPlanJobsResponse) SetHeaders(v map[string]*string) *DescribeElasticPlanJobsResponse {
	s.Headers = v
	return s
}

func (s *DescribeElasticPlanJobsResponse) SetStatusCode(v int32) *DescribeElasticPlanJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeElasticPlanJobsResponse) SetBody(v *DescribeElasticPlanJobsResponseBody) *DescribeElasticPlanJobsResponse {
	s.Body = v
	return s
}

func (s *DescribeElasticPlanJobsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
