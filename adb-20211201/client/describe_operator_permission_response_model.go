// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOperatorPermissionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeOperatorPermissionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeOperatorPermissionResponse
	GetStatusCode() *int32
	SetBody(v *DescribeOperatorPermissionResponseBody) *DescribeOperatorPermissionResponse
	GetBody() *DescribeOperatorPermissionResponseBody
}

type DescribeOperatorPermissionResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeOperatorPermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeOperatorPermissionResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeOperatorPermissionResponse) GoString() string {
	return s.String()
}

func (s *DescribeOperatorPermissionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeOperatorPermissionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeOperatorPermissionResponse) GetBody() *DescribeOperatorPermissionResponseBody {
	return s.Body
}

func (s *DescribeOperatorPermissionResponse) SetHeaders(v map[string]*string) *DescribeOperatorPermissionResponse {
	s.Headers = v
	return s
}

func (s *DescribeOperatorPermissionResponse) SetStatusCode(v int32) *DescribeOperatorPermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOperatorPermissionResponse) SetBody(v *DescribeOperatorPermissionResponseBody) *DescribeOperatorPermissionResponse {
	s.Body = v
	return s
}

func (s *DescribeOperatorPermissionResponse) Validate() error {
	return dara.Validate(s)
}
