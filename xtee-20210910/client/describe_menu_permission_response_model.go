// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMenuPermissionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMenuPermissionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMenuPermissionResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMenuPermissionResponseBody) *DescribeMenuPermissionResponse
	GetBody() *DescribeMenuPermissionResponseBody
}

type DescribeMenuPermissionResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMenuPermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMenuPermissionResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMenuPermissionResponse) GoString() string {
	return s.String()
}

func (s *DescribeMenuPermissionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMenuPermissionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMenuPermissionResponse) GetBody() *DescribeMenuPermissionResponseBody {
	return s.Body
}

func (s *DescribeMenuPermissionResponse) SetHeaders(v map[string]*string) *DescribeMenuPermissionResponse {
	s.Headers = v
	return s
}

func (s *DescribeMenuPermissionResponse) SetStatusCode(v int32) *DescribeMenuPermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMenuPermissionResponse) SetBody(v *DescribeMenuPermissionResponseBody) *DescribeMenuPermissionResponse {
	s.Body = v
	return s
}

func (s *DescribeMenuPermissionResponse) Validate() error {
	return dara.Validate(s)
}
