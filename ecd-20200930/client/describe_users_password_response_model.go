// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUsersPasswordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUsersPasswordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUsersPasswordResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUsersPasswordResponseBody) *DescribeUsersPasswordResponse
	GetBody() *DescribeUsersPasswordResponseBody
}

type DescribeUsersPasswordResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUsersPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUsersPasswordResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUsersPasswordResponse) GoString() string {
	return s.String()
}

func (s *DescribeUsersPasswordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUsersPasswordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUsersPasswordResponse) GetBody() *DescribeUsersPasswordResponseBody {
	return s.Body
}

func (s *DescribeUsersPasswordResponse) SetHeaders(v map[string]*string) *DescribeUsersPasswordResponse {
	s.Headers = v
	return s
}

func (s *DescribeUsersPasswordResponse) SetStatusCode(v int32) *DescribeUsersPasswordResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUsersPasswordResponse) SetBody(v *DescribeUsersPasswordResponseBody) *DescribeUsersPasswordResponse {
	s.Body = v
	return s
}

func (s *DescribeUsersPasswordResponse) Validate() error {
	return dara.Validate(s)
}
