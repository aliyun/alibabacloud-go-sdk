// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTransformStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTransformStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTransformStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTransformStatusResponseBody) *DescribeTransformStatusResponse
	GetBody() *DescribeTransformStatusResponseBody
}

type DescribeTransformStatusResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTransformStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTransformStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTransformStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeTransformStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTransformStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTransformStatusResponse) GetBody() *DescribeTransformStatusResponseBody {
	return s.Body
}

func (s *DescribeTransformStatusResponse) SetHeaders(v map[string]*string) *DescribeTransformStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeTransformStatusResponse) SetStatusCode(v int32) *DescribeTransformStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTransformStatusResponse) SetBody(v *DescribeTransformStatusResponseBody) *DescribeTransformStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeTransformStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
