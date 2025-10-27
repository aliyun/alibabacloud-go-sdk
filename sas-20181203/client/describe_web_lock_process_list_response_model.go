// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebLockProcessListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeWebLockProcessListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeWebLockProcessListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeWebLockProcessListResponseBody) *DescribeWebLockProcessListResponse
	GetBody() *DescribeWebLockProcessListResponseBody
}

type DescribeWebLockProcessListResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeWebLockProcessListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeWebLockProcessListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebLockProcessListResponse) GoString() string {
	return s.String()
}

func (s *DescribeWebLockProcessListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeWebLockProcessListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeWebLockProcessListResponse) GetBody() *DescribeWebLockProcessListResponseBody {
	return s.Body
}

func (s *DescribeWebLockProcessListResponse) SetHeaders(v map[string]*string) *DescribeWebLockProcessListResponse {
	s.Headers = v
	return s
}

func (s *DescribeWebLockProcessListResponse) SetStatusCode(v int32) *DescribeWebLockProcessListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeWebLockProcessListResponse) SetBody(v *DescribeWebLockProcessListResponseBody) *DescribeWebLockProcessListResponse {
	s.Body = v
	return s
}

func (s *DescribeWebLockProcessListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
