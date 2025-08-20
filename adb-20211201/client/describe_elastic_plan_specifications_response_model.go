// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeElasticPlanSpecificationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeElasticPlanSpecificationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeElasticPlanSpecificationsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeElasticPlanSpecificationsResponseBody) *DescribeElasticPlanSpecificationsResponse
	GetBody() *DescribeElasticPlanSpecificationsResponseBody
}

type DescribeElasticPlanSpecificationsResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeElasticPlanSpecificationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeElasticPlanSpecificationsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeElasticPlanSpecificationsResponse) GoString() string {
	return s.String()
}

func (s *DescribeElasticPlanSpecificationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeElasticPlanSpecificationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeElasticPlanSpecificationsResponse) GetBody() *DescribeElasticPlanSpecificationsResponseBody {
	return s.Body
}

func (s *DescribeElasticPlanSpecificationsResponse) SetHeaders(v map[string]*string) *DescribeElasticPlanSpecificationsResponse {
	s.Headers = v
	return s
}

func (s *DescribeElasticPlanSpecificationsResponse) SetStatusCode(v int32) *DescribeElasticPlanSpecificationsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeElasticPlanSpecificationsResponse) SetBody(v *DescribeElasticPlanSpecificationsResponseBody) *DescribeElasticPlanSpecificationsResponse {
	s.Body = v
	return s
}

func (s *DescribeElasticPlanSpecificationsResponse) Validate() error {
	return dara.Validate(s)
}
