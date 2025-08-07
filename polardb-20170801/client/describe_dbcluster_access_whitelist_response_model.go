// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterAccessWhitelistResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBClusterAccessWhitelistResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBClusterAccessWhitelistResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBClusterAccessWhitelistResponseBody) *DescribeDBClusterAccessWhitelistResponse
	GetBody() *DescribeDBClusterAccessWhitelistResponseBody
}

type DescribeDBClusterAccessWhitelistResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBClusterAccessWhitelistResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBClusterAccessWhitelistResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterAccessWhitelistResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAccessWhitelistResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBClusterAccessWhitelistResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBClusterAccessWhitelistResponse) GetBody() *DescribeDBClusterAccessWhitelistResponseBody {
	return s.Body
}

func (s *DescribeDBClusterAccessWhitelistResponse) SetHeaders(v map[string]*string) *DescribeDBClusterAccessWhitelistResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBClusterAccessWhitelistResponse) SetStatusCode(v int32) *DescribeDBClusterAccessWhitelistResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBClusterAccessWhitelistResponse) SetBody(v *DescribeDBClusterAccessWhitelistResponseBody) *DescribeDBClusterAccessWhitelistResponse {
	s.Body = v
	return s
}

func (s *DescribeDBClusterAccessWhitelistResponse) Validate() error {
	return dara.Validate(s)
}
