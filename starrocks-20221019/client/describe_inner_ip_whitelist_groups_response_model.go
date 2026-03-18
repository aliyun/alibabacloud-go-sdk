// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInnerIpWhitelistGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInnerIpWhitelistGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInnerIpWhitelistGroupsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInnerIpWhitelistGroupsResponseBody) *DescribeInnerIpWhitelistGroupsResponse
	GetBody() *DescribeInnerIpWhitelistGroupsResponseBody
}

type DescribeInnerIpWhitelistGroupsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInnerIpWhitelistGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInnerIpWhitelistGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInnerIpWhitelistGroupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeInnerIpWhitelistGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInnerIpWhitelistGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInnerIpWhitelistGroupsResponse) GetBody() *DescribeInnerIpWhitelistGroupsResponseBody {
	return s.Body
}

func (s *DescribeInnerIpWhitelistGroupsResponse) SetHeaders(v map[string]*string) *DescribeInnerIpWhitelistGroupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeInnerIpWhitelistGroupsResponse) SetStatusCode(v int32) *DescribeInnerIpWhitelistGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInnerIpWhitelistGroupsResponse) SetBody(v *DescribeInnerIpWhitelistGroupsResponseBody) *DescribeInnerIpWhitelistGroupsResponse {
	s.Body = v
	return s
}

func (s *DescribeInnerIpWhitelistGroupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
