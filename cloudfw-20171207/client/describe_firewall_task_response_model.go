// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFirewallTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeFirewallTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeFirewallTaskResponse
	GetStatusCode() *int32
	SetBody(v *DescribeFirewallTaskResponseBody) *DescribeFirewallTaskResponse
	GetBody() *DescribeFirewallTaskResponseBody
}

type DescribeFirewallTaskResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFirewallTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFirewallTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeFirewallTaskResponse) GoString() string {
	return s.String()
}

func (s *DescribeFirewallTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeFirewallTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeFirewallTaskResponse) GetBody() *DescribeFirewallTaskResponseBody {
	return s.Body
}

func (s *DescribeFirewallTaskResponse) SetHeaders(v map[string]*string) *DescribeFirewallTaskResponse {
	s.Headers = v
	return s
}

func (s *DescribeFirewallTaskResponse) SetStatusCode(v int32) *DescribeFirewallTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFirewallTaskResponse) SetBody(v *DescribeFirewallTaskResponseBody) *DescribeFirewallTaskResponse {
	s.Body = v
	return s
}

func (s *DescribeFirewallTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
