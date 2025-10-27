// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebLockTotalFileChangeCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeWebLockTotalFileChangeCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeWebLockTotalFileChangeCountResponse
	GetStatusCode() *int32
	SetBody(v *DescribeWebLockTotalFileChangeCountResponseBody) *DescribeWebLockTotalFileChangeCountResponse
	GetBody() *DescribeWebLockTotalFileChangeCountResponseBody
}

type DescribeWebLockTotalFileChangeCountResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeWebLockTotalFileChangeCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeWebLockTotalFileChangeCountResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebLockTotalFileChangeCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeWebLockTotalFileChangeCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeWebLockTotalFileChangeCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeWebLockTotalFileChangeCountResponse) GetBody() *DescribeWebLockTotalFileChangeCountResponseBody {
	return s.Body
}

func (s *DescribeWebLockTotalFileChangeCountResponse) SetHeaders(v map[string]*string) *DescribeWebLockTotalFileChangeCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeWebLockTotalFileChangeCountResponse) SetStatusCode(v int32) *DescribeWebLockTotalFileChangeCountResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeWebLockTotalFileChangeCountResponse) SetBody(v *DescribeWebLockTotalFileChangeCountResponseBody) *DescribeWebLockTotalFileChangeCountResponse {
	s.Body = v
	return s
}

func (s *DescribeWebLockTotalFileChangeCountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
