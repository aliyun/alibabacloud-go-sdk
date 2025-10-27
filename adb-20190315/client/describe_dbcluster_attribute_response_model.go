// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBClusterAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBClusterAttributeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBClusterAttributeResponseBody) *DescribeDBClusterAttributeResponse
	GetBody() *DescribeDBClusterAttributeResponseBody
}

type DescribeDBClusterAttributeResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBClusterAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBClusterAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBClusterAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBClusterAttributeResponse) GetBody() *DescribeDBClusterAttributeResponseBody {
	return s.Body
}

func (s *DescribeDBClusterAttributeResponse) SetHeaders(v map[string]*string) *DescribeDBClusterAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBClusterAttributeResponse) SetStatusCode(v int32) *DescribeDBClusterAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBClusterAttributeResponse) SetBody(v *DescribeDBClusterAttributeResponseBody) *DescribeDBClusterAttributeResponse {
	s.Body = v
	return s
}

func (s *DescribeDBClusterAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
