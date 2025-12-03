// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIpWhitelistResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeIpWhitelistResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeIpWhitelistResponse
	GetStatusCode() *int32
	SetBody(v *DescribeIpWhitelistResponseBody) *DescribeIpWhitelistResponse
	GetBody() *DescribeIpWhitelistResponseBody
}

type DescribeIpWhitelistResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeIpWhitelistResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeIpWhitelistResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeIpWhitelistResponse) GoString() string {
	return s.String()
}

func (s *DescribeIpWhitelistResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeIpWhitelistResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeIpWhitelistResponse) GetBody() *DescribeIpWhitelistResponseBody {
	return s.Body
}

func (s *DescribeIpWhitelistResponse) SetHeaders(v map[string]*string) *DescribeIpWhitelistResponse {
	s.Headers = v
	return s
}

func (s *DescribeIpWhitelistResponse) SetStatusCode(v int32) *DescribeIpWhitelistResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeIpWhitelistResponse) SetBody(v *DescribeIpWhitelistResponseBody) *DescribeIpWhitelistResponse {
	s.Body = v
	return s
}

func (s *DescribeIpWhitelistResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
