// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebLockStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeWebLockStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeWebLockStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeWebLockStatusResponseBody) *DescribeWebLockStatusResponse
	GetBody() *DescribeWebLockStatusResponseBody
}

type DescribeWebLockStatusResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeWebLockStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeWebLockStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebLockStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeWebLockStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeWebLockStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeWebLockStatusResponse) GetBody() *DescribeWebLockStatusResponseBody {
	return s.Body
}

func (s *DescribeWebLockStatusResponse) SetHeaders(v map[string]*string) *DescribeWebLockStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeWebLockStatusResponse) SetStatusCode(v int32) *DescribeWebLockStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeWebLockStatusResponse) SetBody(v *DescribeWebLockStatusResponseBody) *DescribeWebLockStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeWebLockStatusResponse) Validate() error {
	return dara.Validate(s)
}
