// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFirewallTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeFirewallTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeFirewallTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeFirewallTemplatesResponseBody) *DescribeFirewallTemplatesResponse
	GetBody() *DescribeFirewallTemplatesResponseBody
}

type DescribeFirewallTemplatesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFirewallTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFirewallTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeFirewallTemplatesResponse) GoString() string {
	return s.String()
}

func (s *DescribeFirewallTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeFirewallTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeFirewallTemplatesResponse) GetBody() *DescribeFirewallTemplatesResponseBody {
	return s.Body
}

func (s *DescribeFirewallTemplatesResponse) SetHeaders(v map[string]*string) *DescribeFirewallTemplatesResponse {
	s.Headers = v
	return s
}

func (s *DescribeFirewallTemplatesResponse) SetStatusCode(v int32) *DescribeFirewallTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFirewallTemplatesResponse) SetBody(v *DescribeFirewallTemplatesResponseBody) *DescribeFirewallTemplatesResponse {
	s.Body = v
	return s
}

func (s *DescribeFirewallTemplatesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
