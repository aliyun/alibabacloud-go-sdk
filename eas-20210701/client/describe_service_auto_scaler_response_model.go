// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceAutoScalerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeServiceAutoScalerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeServiceAutoScalerResponse
	GetStatusCode() *int32
	SetBody(v *DescribeServiceAutoScalerResponseBody) *DescribeServiceAutoScalerResponse
	GetBody() *DescribeServiceAutoScalerResponseBody
}

type DescribeServiceAutoScalerResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeServiceAutoScalerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeServiceAutoScalerResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceAutoScalerResponse) GoString() string {
	return s.String()
}

func (s *DescribeServiceAutoScalerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeServiceAutoScalerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeServiceAutoScalerResponse) GetBody() *DescribeServiceAutoScalerResponseBody {
	return s.Body
}

func (s *DescribeServiceAutoScalerResponse) SetHeaders(v map[string]*string) *DescribeServiceAutoScalerResponse {
	s.Headers = v
	return s
}

func (s *DescribeServiceAutoScalerResponse) SetStatusCode(v int32) *DescribeServiceAutoScalerResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeServiceAutoScalerResponse) SetBody(v *DescribeServiceAutoScalerResponseBody) *DescribeServiceAutoScalerResponse {
	s.Body = v
	return s
}

func (s *DescribeServiceAutoScalerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
