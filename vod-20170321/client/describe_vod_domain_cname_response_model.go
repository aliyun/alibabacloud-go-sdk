// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainCnameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVodDomainCnameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVodDomainCnameResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVodDomainCnameResponseBody) *DescribeVodDomainCnameResponse
	GetBody() *DescribeVodDomainCnameResponseBody
}

type DescribeVodDomainCnameResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVodDomainCnameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVodDomainCnameResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainCnameResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainCnameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVodDomainCnameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVodDomainCnameResponse) GetBody() *DescribeVodDomainCnameResponseBody {
	return s.Body
}

func (s *DescribeVodDomainCnameResponse) SetHeaders(v map[string]*string) *DescribeVodDomainCnameResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodDomainCnameResponse) SetStatusCode(v int32) *DescribeVodDomainCnameResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVodDomainCnameResponse) SetBody(v *DescribeVodDomainCnameResponseBody) *DescribeVodDomainCnameResponse {
	s.Body = v
	return s
}

func (s *DescribeVodDomainCnameResponse) Validate() error {
	return dara.Validate(s)
}
