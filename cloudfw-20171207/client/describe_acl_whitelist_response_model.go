// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAclWhitelistResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAclWhitelistResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAclWhitelistResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAclWhitelistResponseBody) *DescribeAclWhitelistResponse
	GetBody() *DescribeAclWhitelistResponseBody
}

type DescribeAclWhitelistResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAclWhitelistResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAclWhitelistResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAclWhitelistResponse) GoString() string {
	return s.String()
}

func (s *DescribeAclWhitelistResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAclWhitelistResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAclWhitelistResponse) GetBody() *DescribeAclWhitelistResponseBody {
	return s.Body
}

func (s *DescribeAclWhitelistResponse) SetHeaders(v map[string]*string) *DescribeAclWhitelistResponse {
	s.Headers = v
	return s
}

func (s *DescribeAclWhitelistResponse) SetStatusCode(v int32) *DescribeAclWhitelistResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAclWhitelistResponse) SetBody(v *DescribeAclWhitelistResponseBody) *DescribeAclWhitelistResponse {
	s.Body = v
	return s
}

func (s *DescribeAclWhitelistResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
