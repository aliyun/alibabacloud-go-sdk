// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceCronScalerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeServiceCronScalerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeServiceCronScalerResponse
	GetStatusCode() *int32
	SetBody(v *DescribeServiceCronScalerResponseBody) *DescribeServiceCronScalerResponse
	GetBody() *DescribeServiceCronScalerResponseBody
}

type DescribeServiceCronScalerResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeServiceCronScalerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeServiceCronScalerResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceCronScalerResponse) GoString() string {
	return s.String()
}

func (s *DescribeServiceCronScalerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeServiceCronScalerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeServiceCronScalerResponse) GetBody() *DescribeServiceCronScalerResponseBody {
	return s.Body
}

func (s *DescribeServiceCronScalerResponse) SetHeaders(v map[string]*string) *DescribeServiceCronScalerResponse {
	s.Headers = v
	return s
}

func (s *DescribeServiceCronScalerResponse) SetStatusCode(v int32) *DescribeServiceCronScalerResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeServiceCronScalerResponse) SetBody(v *DescribeServiceCronScalerResponseBody) *DescribeServiceCronScalerResponse {
	s.Body = v
	return s
}

func (s *DescribeServiceCronScalerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
