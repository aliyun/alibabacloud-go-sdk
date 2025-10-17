// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeElasticPlanAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeElasticPlanAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeElasticPlanAttributeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeElasticPlanAttributeResponseBody) *DescribeElasticPlanAttributeResponse
	GetBody() *DescribeElasticPlanAttributeResponseBody
}

type DescribeElasticPlanAttributeResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeElasticPlanAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeElasticPlanAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeElasticPlanAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeElasticPlanAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeElasticPlanAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeElasticPlanAttributeResponse) GetBody() *DescribeElasticPlanAttributeResponseBody {
	return s.Body
}

func (s *DescribeElasticPlanAttributeResponse) SetHeaders(v map[string]*string) *DescribeElasticPlanAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeElasticPlanAttributeResponse) SetStatusCode(v int32) *DescribeElasticPlanAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeElasticPlanAttributeResponse) SetBody(v *DescribeElasticPlanAttributeResponseBody) *DescribeElasticPlanAttributeResponse {
	s.Body = v
	return s
}

func (s *DescribeElasticPlanAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
