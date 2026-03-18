// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePackIpListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePackIpListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePackIpListResponse
	GetStatusCode() *int32
	SetBody(v *DescribePackIpListResponseBody) *DescribePackIpListResponse
	GetBody() *DescribePackIpListResponseBody
}

type DescribePackIpListResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePackIpListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePackIpListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePackIpListResponse) GoString() string {
	return s.String()
}

func (s *DescribePackIpListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePackIpListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePackIpListResponse) GetBody() *DescribePackIpListResponseBody {
	return s.Body
}

func (s *DescribePackIpListResponse) SetHeaders(v map[string]*string) *DescribePackIpListResponse {
	s.Headers = v
	return s
}

func (s *DescribePackIpListResponse) SetStatusCode(v int32) *DescribePackIpListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePackIpListResponse) SetBody(v *DescribePackIpListResponseBody) *DescribePackIpListResponse {
	s.Body = v
	return s
}

func (s *DescribePackIpListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
