// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourceDLinkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeResourceDLinkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeResourceDLinkResponse
	GetStatusCode() *int32
	SetBody(v *DescribeResourceDLinkResponseBody) *DescribeResourceDLinkResponse
	GetBody() *DescribeResourceDLinkResponseBody
}

type DescribeResourceDLinkResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeResourceDLinkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeResourceDLinkResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceDLinkResponse) GoString() string {
	return s.String()
}

func (s *DescribeResourceDLinkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeResourceDLinkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeResourceDLinkResponse) GetBody() *DescribeResourceDLinkResponseBody {
	return s.Body
}

func (s *DescribeResourceDLinkResponse) SetHeaders(v map[string]*string) *DescribeResourceDLinkResponse {
	s.Headers = v
	return s
}

func (s *DescribeResourceDLinkResponse) SetStatusCode(v int32) *DescribeResourceDLinkResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeResourceDLinkResponse) SetBody(v *DescribeResourceDLinkResponseBody) *DescribeResourceDLinkResponse {
	s.Body = v
	return s
}

func (s *DescribeResourceDLinkResponse) Validate() error {
	return dara.Validate(s)
}
