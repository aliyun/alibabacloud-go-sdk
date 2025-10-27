// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebLockInclusiveFileTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeWebLockInclusiveFileTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeWebLockInclusiveFileTypeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeWebLockInclusiveFileTypeResponseBody) *DescribeWebLockInclusiveFileTypeResponse
	GetBody() *DescribeWebLockInclusiveFileTypeResponseBody
}

type DescribeWebLockInclusiveFileTypeResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeWebLockInclusiveFileTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeWebLockInclusiveFileTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebLockInclusiveFileTypeResponse) GoString() string {
	return s.String()
}

func (s *DescribeWebLockInclusiveFileTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeWebLockInclusiveFileTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeWebLockInclusiveFileTypeResponse) GetBody() *DescribeWebLockInclusiveFileTypeResponseBody {
	return s.Body
}

func (s *DescribeWebLockInclusiveFileTypeResponse) SetHeaders(v map[string]*string) *DescribeWebLockInclusiveFileTypeResponse {
	s.Headers = v
	return s
}

func (s *DescribeWebLockInclusiveFileTypeResponse) SetStatusCode(v int32) *DescribeWebLockInclusiveFileTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeWebLockInclusiveFileTypeResponse) SetBody(v *DescribeWebLockInclusiveFileTypeResponseBody) *DescribeWebLockInclusiveFileTypeResponse {
	s.Body = v
	return s
}

func (s *DescribeWebLockInclusiveFileTypeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
