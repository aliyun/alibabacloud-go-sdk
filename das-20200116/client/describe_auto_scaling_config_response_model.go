// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAutoScalingConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAutoScalingConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAutoScalingConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAutoScalingConfigResponseBody) *DescribeAutoScalingConfigResponse
	GetBody() *DescribeAutoScalingConfigResponseBody
}

type DescribeAutoScalingConfigResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAutoScalingConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAutoScalingConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoScalingConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeAutoScalingConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAutoScalingConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAutoScalingConfigResponse) GetBody() *DescribeAutoScalingConfigResponseBody {
	return s.Body
}

func (s *DescribeAutoScalingConfigResponse) SetHeaders(v map[string]*string) *DescribeAutoScalingConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeAutoScalingConfigResponse) SetStatusCode(v int32) *DescribeAutoScalingConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAutoScalingConfigResponse) SetBody(v *DescribeAutoScalingConfigResponseBody) *DescribeAutoScalingConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeAutoScalingConfigResponse) Validate() error {
	return dara.Validate(s)
}
