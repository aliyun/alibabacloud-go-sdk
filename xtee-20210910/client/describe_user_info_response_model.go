// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUserInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUserInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUserInfoResponseBody) *DescribeUserInfoResponse
	GetBody() *DescribeUserInfoResponseBody
}

type DescribeUserInfoResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUserInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUserInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUserInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUserInfoResponse) GetBody() *DescribeUserInfoResponseBody {
	return s.Body
}

func (s *DescribeUserInfoResponse) SetHeaders(v map[string]*string) *DescribeUserInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserInfoResponse) SetStatusCode(v int32) *DescribeUserInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserInfoResponse) SetBody(v *DescribeUserInfoResponseBody) *DescribeUserInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeUserInfoResponse) Validate() error {
	return dara.Validate(s)
}
