// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSDGSharedDisksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSDGSharedDisksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSDGSharedDisksResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSDGSharedDisksResponseBody) *DescribeSDGSharedDisksResponse
	GetBody() *DescribeSDGSharedDisksResponseBody
}

type DescribeSDGSharedDisksResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSDGSharedDisksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSDGSharedDisksResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSDGSharedDisksResponse) GoString() string {
	return s.String()
}

func (s *DescribeSDGSharedDisksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSDGSharedDisksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSDGSharedDisksResponse) GetBody() *DescribeSDGSharedDisksResponseBody {
	return s.Body
}

func (s *DescribeSDGSharedDisksResponse) SetHeaders(v map[string]*string) *DescribeSDGSharedDisksResponse {
	s.Headers = v
	return s
}

func (s *DescribeSDGSharedDisksResponse) SetStatusCode(v int32) *DescribeSDGSharedDisksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSDGSharedDisksResponse) SetBody(v *DescribeSDGSharedDisksResponseBody) *DescribeSDGSharedDisksResponse {
	s.Body = v
	return s
}

func (s *DescribeSDGSharedDisksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
