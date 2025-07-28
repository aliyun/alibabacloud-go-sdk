// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserSlsLogRegionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUserSlsLogRegionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUserSlsLogRegionsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUserSlsLogRegionsResponseBody) *DescribeUserSlsLogRegionsResponse
	GetBody() *DescribeUserSlsLogRegionsResponseBody
}

type DescribeUserSlsLogRegionsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUserSlsLogRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUserSlsLogRegionsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserSlsLogRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserSlsLogRegionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUserSlsLogRegionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUserSlsLogRegionsResponse) GetBody() *DescribeUserSlsLogRegionsResponseBody {
	return s.Body
}

func (s *DescribeUserSlsLogRegionsResponse) SetHeaders(v map[string]*string) *DescribeUserSlsLogRegionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserSlsLogRegionsResponse) SetStatusCode(v int32) *DescribeUserSlsLogRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserSlsLogRegionsResponse) SetBody(v *DescribeUserSlsLogRegionsResponseBody) *DescribeUserSlsLogRegionsResponse {
	s.Body = v
	return s
}

func (s *DescribeUserSlsLogRegionsResponse) Validate() error {
	return dara.Validate(s)
}
