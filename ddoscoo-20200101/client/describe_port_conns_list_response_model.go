// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePortConnsListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePortConnsListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePortConnsListResponse
	GetStatusCode() *int32
	SetBody(v *DescribePortConnsListResponseBody) *DescribePortConnsListResponse
	GetBody() *DescribePortConnsListResponseBody
}

type DescribePortConnsListResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePortConnsListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePortConnsListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePortConnsListResponse) GoString() string {
	return s.String()
}

func (s *DescribePortConnsListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePortConnsListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePortConnsListResponse) GetBody() *DescribePortConnsListResponseBody {
	return s.Body
}

func (s *DescribePortConnsListResponse) SetHeaders(v map[string]*string) *DescribePortConnsListResponse {
	s.Headers = v
	return s
}

func (s *DescribePortConnsListResponse) SetStatusCode(v int32) *DescribePortConnsListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePortConnsListResponse) SetBody(v *DescribePortConnsListResponseBody) *DescribePortConnsListResponse {
	s.Body = v
	return s
}

func (s *DescribePortConnsListResponse) Validate() error {
	return dara.Validate(s)
}
