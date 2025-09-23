// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExtensionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeExtensionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeExtensionsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeExtensionsResponseBody) *DescribeExtensionsResponse
	GetBody() *DescribeExtensionsResponseBody
}

type DescribeExtensionsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeExtensionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeExtensionsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeExtensionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeExtensionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeExtensionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeExtensionsResponse) GetBody() *DescribeExtensionsResponseBody {
	return s.Body
}

func (s *DescribeExtensionsResponse) SetHeaders(v map[string]*string) *DescribeExtensionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeExtensionsResponse) SetStatusCode(v int32) *DescribeExtensionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeExtensionsResponse) SetBody(v *DescribeExtensionsResponseBody) *DescribeExtensionsResponse {
	s.Body = v
	return s
}

func (s *DescribeExtensionsResponse) Validate() error {
	return dara.Validate(s)
}
