// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImagePermissionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeImagePermissionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeImagePermissionResponse
	GetStatusCode() *int32
	SetBody(v *DescribeImagePermissionResponseBody) *DescribeImagePermissionResponse
	GetBody() *DescribeImagePermissionResponseBody
}

type DescribeImagePermissionResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeImagePermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeImagePermissionResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeImagePermissionResponse) GoString() string {
	return s.String()
}

func (s *DescribeImagePermissionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeImagePermissionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeImagePermissionResponse) GetBody() *DescribeImagePermissionResponseBody {
	return s.Body
}

func (s *DescribeImagePermissionResponse) SetHeaders(v map[string]*string) *DescribeImagePermissionResponse {
	s.Headers = v
	return s
}

func (s *DescribeImagePermissionResponse) SetStatusCode(v int32) *DescribeImagePermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeImagePermissionResponse) SetBody(v *DescribeImagePermissionResponseBody) *DescribeImagePermissionResponse {
	s.Body = v
	return s
}

func (s *DescribeImagePermissionResponse) Validate() error {
	return dara.Validate(s)
}
