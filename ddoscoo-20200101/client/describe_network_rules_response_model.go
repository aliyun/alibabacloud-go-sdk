// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNetworkRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeNetworkRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeNetworkRulesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeNetworkRulesResponseBody) *DescribeNetworkRulesResponse
	GetBody() *DescribeNetworkRulesResponseBody
}

type DescribeNetworkRulesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNetworkRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNetworkRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkRulesResponse) GoString() string {
	return s.String()
}

func (s *DescribeNetworkRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeNetworkRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeNetworkRulesResponse) GetBody() *DescribeNetworkRulesResponseBody {
	return s.Body
}

func (s *DescribeNetworkRulesResponse) SetHeaders(v map[string]*string) *DescribeNetworkRulesResponse {
	s.Headers = v
	return s
}

func (s *DescribeNetworkRulesResponse) SetStatusCode(v int32) *DescribeNetworkRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNetworkRulesResponse) SetBody(v *DescribeNetworkRulesResponseBody) *DescribeNetworkRulesResponse {
	s.Body = v
	return s
}

func (s *DescribeNetworkRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
