// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterAddonsVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeClusterAddonsVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeClusterAddonsVersionResponse
	GetStatusCode() *int32
	SetBody(v map[string]interface{}) *DescribeClusterAddonsVersionResponse
	GetBody() map[string]interface{}
}

type DescribeClusterAddonsVersionResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       map[string]interface{} `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeClusterAddonsVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterAddonsVersionResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterAddonsVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeClusterAddonsVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeClusterAddonsVersionResponse) GetBody() map[string]interface{} {
	return s.Body
}

func (s *DescribeClusterAddonsVersionResponse) SetHeaders(v map[string]*string) *DescribeClusterAddonsVersionResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterAddonsVersionResponse) SetStatusCode(v int32) *DescribeClusterAddonsVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeClusterAddonsVersionResponse) SetBody(v map[string]interface{}) *DescribeClusterAddonsVersionResponse {
	s.Body = v
	return s
}

func (s *DescribeClusterAddonsVersionResponse) Validate() error {
	return dara.Validate(s)
}
