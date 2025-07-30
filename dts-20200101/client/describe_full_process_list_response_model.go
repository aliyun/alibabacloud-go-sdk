// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFullProcessListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeFullProcessListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeFullProcessListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeFullProcessListResponseBody) *DescribeFullProcessListResponse
	GetBody() *DescribeFullProcessListResponseBody
}

type DescribeFullProcessListResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFullProcessListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFullProcessListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeFullProcessListResponse) GoString() string {
	return s.String()
}

func (s *DescribeFullProcessListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeFullProcessListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeFullProcessListResponse) GetBody() *DescribeFullProcessListResponseBody {
	return s.Body
}

func (s *DescribeFullProcessListResponse) SetHeaders(v map[string]*string) *DescribeFullProcessListResponse {
	s.Headers = v
	return s
}

func (s *DescribeFullProcessListResponse) SetStatusCode(v int32) *DescribeFullProcessListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFullProcessListResponse) SetBody(v *DescribeFullProcessListResponseBody) *DescribeFullProcessListResponse {
	s.Body = v
	return s
}

func (s *DescribeFullProcessListResponse) Validate() error {
	return dara.Validate(s)
}
