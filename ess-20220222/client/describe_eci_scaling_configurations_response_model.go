// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEciScalingConfigurationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEciScalingConfigurationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEciScalingConfigurationsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEciScalingConfigurationsResponseBody) *DescribeEciScalingConfigurationsResponse
	GetBody() *DescribeEciScalingConfigurationsResponseBody
}

type DescribeEciScalingConfigurationsResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEciScalingConfigurationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEciScalingConfigurationsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEciScalingConfigurationsResponse) GoString() string {
	return s.String()
}

func (s *DescribeEciScalingConfigurationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEciScalingConfigurationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEciScalingConfigurationsResponse) GetBody() *DescribeEciScalingConfigurationsResponseBody {
	return s.Body
}

func (s *DescribeEciScalingConfigurationsResponse) SetHeaders(v map[string]*string) *DescribeEciScalingConfigurationsResponse {
	s.Headers = v
	return s
}

func (s *DescribeEciScalingConfigurationsResponse) SetStatusCode(v int32) *DescribeEciScalingConfigurationsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponse) SetBody(v *DescribeEciScalingConfigurationsResponseBody) *DescribeEciScalingConfigurationsResponse {
	s.Body = v
	return s
}

func (s *DescribeEciScalingConfigurationsResponse) Validate() error {
	return dara.Validate(s)
}
