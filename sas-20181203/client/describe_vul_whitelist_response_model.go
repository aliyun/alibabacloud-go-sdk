// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVulWhitelistResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVulWhitelistResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVulWhitelistResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVulWhitelistResponseBody) *DescribeVulWhitelistResponse
	GetBody() *DescribeVulWhitelistResponseBody
}

type DescribeVulWhitelistResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVulWhitelistResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVulWhitelistResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVulWhitelistResponse) GoString() string {
	return s.String()
}

func (s *DescribeVulWhitelistResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVulWhitelistResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVulWhitelistResponse) GetBody() *DescribeVulWhitelistResponseBody {
	return s.Body
}

func (s *DescribeVulWhitelistResponse) SetHeaders(v map[string]*string) *DescribeVulWhitelistResponse {
	s.Headers = v
	return s
}

func (s *DescribeVulWhitelistResponse) SetStatusCode(v int32) *DescribeVulWhitelistResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVulWhitelistResponse) SetBody(v *DescribeVulWhitelistResponseBody) *DescribeVulWhitelistResponse {
	s.Body = v
	return s
}

func (s *DescribeVulWhitelistResponse) Validate() error {
	return dara.Validate(s)
}
