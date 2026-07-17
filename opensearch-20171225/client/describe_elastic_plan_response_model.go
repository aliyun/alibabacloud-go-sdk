// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeElasticPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeElasticPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeElasticPlanResponse
	GetStatusCode() *int32
	SetBody(v *DescribeElasticPlanResponseBody) *DescribeElasticPlanResponse
	GetBody() *DescribeElasticPlanResponseBody
}

type DescribeElasticPlanResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeElasticPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeElasticPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeElasticPlanResponse) GoString() string {
	return s.String()
}

func (s *DescribeElasticPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeElasticPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeElasticPlanResponse) GetBody() *DescribeElasticPlanResponseBody {
	return s.Body
}

func (s *DescribeElasticPlanResponse) SetHeaders(v map[string]*string) *DescribeElasticPlanResponse {
	s.Headers = v
	return s
}

func (s *DescribeElasticPlanResponse) SetStatusCode(v int32) *DescribeElasticPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeElasticPlanResponse) SetBody(v *DescribeElasticPlanResponseBody) *DescribeElasticPlanResponse {
	s.Body = v
	return s
}

func (s *DescribeElasticPlanResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
