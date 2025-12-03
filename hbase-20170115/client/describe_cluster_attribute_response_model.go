// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeClusterAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeClusterAttributeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeClusterAttributeResponseBody) *DescribeClusterAttributeResponse
	GetBody() *DescribeClusterAttributeResponseBody
}

type DescribeClusterAttributeResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeClusterAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeClusterAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeClusterAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeClusterAttributeResponse) GetBody() *DescribeClusterAttributeResponseBody {
	return s.Body
}

func (s *DescribeClusterAttributeResponse) SetHeaders(v map[string]*string) *DescribeClusterAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterAttributeResponse) SetStatusCode(v int32) *DescribeClusterAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeClusterAttributeResponse) SetBody(v *DescribeClusterAttributeResponseBody) *DescribeClusterAttributeResponse {
	s.Body = v
	return s
}

func (s *DescribeClusterAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
