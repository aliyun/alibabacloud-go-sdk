// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeArchiveTableListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeArchiveTableListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeArchiveTableListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeArchiveTableListResponseBody) *DescribeArchiveTableListResponse
	GetBody() *DescribeArchiveTableListResponseBody
}

type DescribeArchiveTableListResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeArchiveTableListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeArchiveTableListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeArchiveTableListResponse) GoString() string {
	return s.String()
}

func (s *DescribeArchiveTableListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeArchiveTableListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeArchiveTableListResponse) GetBody() *DescribeArchiveTableListResponseBody {
	return s.Body
}

func (s *DescribeArchiveTableListResponse) SetHeaders(v map[string]*string) *DescribeArchiveTableListResponse {
	s.Headers = v
	return s
}

func (s *DescribeArchiveTableListResponse) SetStatusCode(v int32) *DescribeArchiveTableListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeArchiveTableListResponse) SetBody(v *DescribeArchiveTableListResponseBody) *DescribeArchiveTableListResponse {
	s.Body = v
	return s
}

func (s *DescribeArchiveTableListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
