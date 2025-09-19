// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebLockExclusiveFileTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeWebLockExclusiveFileTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeWebLockExclusiveFileTypeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeWebLockExclusiveFileTypeResponseBody) *DescribeWebLockExclusiveFileTypeResponse
	GetBody() *DescribeWebLockExclusiveFileTypeResponseBody
}

type DescribeWebLockExclusiveFileTypeResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeWebLockExclusiveFileTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeWebLockExclusiveFileTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebLockExclusiveFileTypeResponse) GoString() string {
	return s.String()
}

func (s *DescribeWebLockExclusiveFileTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeWebLockExclusiveFileTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeWebLockExclusiveFileTypeResponse) GetBody() *DescribeWebLockExclusiveFileTypeResponseBody {
	return s.Body
}

func (s *DescribeWebLockExclusiveFileTypeResponse) SetHeaders(v map[string]*string) *DescribeWebLockExclusiveFileTypeResponse {
	s.Headers = v
	return s
}

func (s *DescribeWebLockExclusiveFileTypeResponse) SetStatusCode(v int32) *DescribeWebLockExclusiveFileTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeWebLockExclusiveFileTypeResponse) SetBody(v *DescribeWebLockExclusiveFileTypeResponseBody) *DescribeWebLockExclusiveFileTypeResponse {
	s.Body = v
	return s
}

func (s *DescribeWebLockExclusiveFileTypeResponse) Validate() error {
	return dara.Validate(s)
}
