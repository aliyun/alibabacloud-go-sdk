// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInvadeEventListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInvadeEventListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInvadeEventListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInvadeEventListResponseBody) *DescribeInvadeEventListResponse
	GetBody() *DescribeInvadeEventListResponseBody
}

type DescribeInvadeEventListResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInvadeEventListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInvadeEventListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInvadeEventListResponse) GoString() string {
	return s.String()
}

func (s *DescribeInvadeEventListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInvadeEventListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInvadeEventListResponse) GetBody() *DescribeInvadeEventListResponseBody {
	return s.Body
}

func (s *DescribeInvadeEventListResponse) SetHeaders(v map[string]*string) *DescribeInvadeEventListResponse {
	s.Headers = v
	return s
}

func (s *DescribeInvadeEventListResponse) SetStatusCode(v int32) *DescribeInvadeEventListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInvadeEventListResponse) SetBody(v *DescribeInvadeEventListResponseBody) *DescribeInvadeEventListResponse {
	s.Body = v
	return s
}

func (s *DescribeInvadeEventListResponse) Validate() error {
	return dara.Validate(s)
}
