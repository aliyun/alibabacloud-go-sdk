// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceIpWhitelistResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstanceIpWhitelistResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstanceIpWhitelistResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstanceIpWhitelistResponseBody) *DescribeInstanceIpWhitelistResponse
	GetBody() *DescribeInstanceIpWhitelistResponseBody
}

type DescribeInstanceIpWhitelistResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceIpWhitelistResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceIpWhitelistResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceIpWhitelistResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceIpWhitelistResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstanceIpWhitelistResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstanceIpWhitelistResponse) GetBody() *DescribeInstanceIpWhitelistResponseBody {
	return s.Body
}

func (s *DescribeInstanceIpWhitelistResponse) SetHeaders(v map[string]*string) *DescribeInstanceIpWhitelistResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceIpWhitelistResponse) SetStatusCode(v int32) *DescribeInstanceIpWhitelistResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceIpWhitelistResponse) SetBody(v *DescribeInstanceIpWhitelistResponseBody) *DescribeInstanceIpWhitelistResponse {
	s.Body = v
	return s
}

func (s *DescribeInstanceIpWhitelistResponse) Validate() error {
	return dara.Validate(s)
}
