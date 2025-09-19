// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebLockConfigListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeWebLockConfigListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeWebLockConfigListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeWebLockConfigListResponseBody) *DescribeWebLockConfigListResponse
	GetBody() *DescribeWebLockConfigListResponseBody
}

type DescribeWebLockConfigListResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeWebLockConfigListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeWebLockConfigListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebLockConfigListResponse) GoString() string {
	return s.String()
}

func (s *DescribeWebLockConfigListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeWebLockConfigListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeWebLockConfigListResponse) GetBody() *DescribeWebLockConfigListResponseBody {
	return s.Body
}

func (s *DescribeWebLockConfigListResponse) SetHeaders(v map[string]*string) *DescribeWebLockConfigListResponse {
	s.Headers = v
	return s
}

func (s *DescribeWebLockConfigListResponse) SetStatusCode(v int32) *DescribeWebLockConfigListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeWebLockConfigListResponse) SetBody(v *DescribeWebLockConfigListResponseBody) *DescribeWebLockConfigListResponse {
	s.Body = v
	return s
}

func (s *DescribeWebLockConfigListResponse) Validate() error {
	return dara.Validate(s)
}
