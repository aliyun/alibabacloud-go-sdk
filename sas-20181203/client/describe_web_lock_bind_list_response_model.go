// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebLockBindListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeWebLockBindListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeWebLockBindListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeWebLockBindListResponseBody) *DescribeWebLockBindListResponse
	GetBody() *DescribeWebLockBindListResponseBody
}

type DescribeWebLockBindListResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeWebLockBindListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeWebLockBindListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebLockBindListResponse) GoString() string {
	return s.String()
}

func (s *DescribeWebLockBindListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeWebLockBindListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeWebLockBindListResponse) GetBody() *DescribeWebLockBindListResponseBody {
	return s.Body
}

func (s *DescribeWebLockBindListResponse) SetHeaders(v map[string]*string) *DescribeWebLockBindListResponse {
	s.Headers = v
	return s
}

func (s *DescribeWebLockBindListResponse) SetStatusCode(v int32) *DescribeWebLockBindListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeWebLockBindListResponse) SetBody(v *DescribeWebLockBindListResponseBody) *DescribeWebLockBindListResponse {
	s.Body = v
	return s
}

func (s *DescribeWebLockBindListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
