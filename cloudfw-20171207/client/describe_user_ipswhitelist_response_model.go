// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserIPSWhitelistResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUserIPSWhitelistResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUserIPSWhitelistResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUserIPSWhitelistResponseBody) *DescribeUserIPSWhitelistResponse
	GetBody() *DescribeUserIPSWhitelistResponseBody
}

type DescribeUserIPSWhitelistResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUserIPSWhitelistResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUserIPSWhitelistResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserIPSWhitelistResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserIPSWhitelistResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUserIPSWhitelistResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUserIPSWhitelistResponse) GetBody() *DescribeUserIPSWhitelistResponseBody {
	return s.Body
}

func (s *DescribeUserIPSWhitelistResponse) SetHeaders(v map[string]*string) *DescribeUserIPSWhitelistResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserIPSWhitelistResponse) SetStatusCode(v int32) *DescribeUserIPSWhitelistResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserIPSWhitelistResponse) SetBody(v *DescribeUserIPSWhitelistResponseBody) *DescribeUserIPSWhitelistResponse {
	s.Body = v
	return s
}

func (s *DescribeUserIPSWhitelistResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
