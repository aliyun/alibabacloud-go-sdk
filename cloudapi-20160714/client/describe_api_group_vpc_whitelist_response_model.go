// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApiGroupVpcWhitelistResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApiGroupVpcWhitelistResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApiGroupVpcWhitelistResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApiGroupVpcWhitelistResponseBody) *DescribeApiGroupVpcWhitelistResponse
	GetBody() *DescribeApiGroupVpcWhitelistResponseBody
}

type DescribeApiGroupVpcWhitelistResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApiGroupVpcWhitelistResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApiGroupVpcWhitelistResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiGroupVpcWhitelistResponse) GoString() string {
	return s.String()
}

func (s *DescribeApiGroupVpcWhitelistResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApiGroupVpcWhitelistResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApiGroupVpcWhitelistResponse) GetBody() *DescribeApiGroupVpcWhitelistResponseBody {
	return s.Body
}

func (s *DescribeApiGroupVpcWhitelistResponse) SetHeaders(v map[string]*string) *DescribeApiGroupVpcWhitelistResponse {
	s.Headers = v
	return s
}

func (s *DescribeApiGroupVpcWhitelistResponse) SetStatusCode(v int32) *DescribeApiGroupVpcWhitelistResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApiGroupVpcWhitelistResponse) SetBody(v *DescribeApiGroupVpcWhitelistResponseBody) *DescribeApiGroupVpcWhitelistResponse {
	s.Body = v
	return s
}

func (s *DescribeApiGroupVpcWhitelistResponse) Validate() error {
	return dara.Validate(s)
}
