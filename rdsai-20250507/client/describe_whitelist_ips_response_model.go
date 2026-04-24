// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWhitelistIpsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeWhitelistIpsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeWhitelistIpsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeWhitelistIpsResponseBody) *DescribeWhitelistIpsResponse
	GetBody() *DescribeWhitelistIpsResponseBody
}

type DescribeWhitelistIpsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeWhitelistIpsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeWhitelistIpsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeWhitelistIpsResponse) GoString() string {
	return s.String()
}

func (s *DescribeWhitelistIpsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeWhitelistIpsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeWhitelistIpsResponse) GetBody() *DescribeWhitelistIpsResponseBody {
	return s.Body
}

func (s *DescribeWhitelistIpsResponse) SetHeaders(v map[string]*string) *DescribeWhitelistIpsResponse {
	s.Headers = v
	return s
}

func (s *DescribeWhitelistIpsResponse) SetStatusCode(v int32) *DescribeWhitelistIpsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeWhitelistIpsResponse) SetBody(v *DescribeWhitelistIpsResponseBody) *DescribeWhitelistIpsResponse {
	s.Body = v
	return s
}

func (s *DescribeWhitelistIpsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
