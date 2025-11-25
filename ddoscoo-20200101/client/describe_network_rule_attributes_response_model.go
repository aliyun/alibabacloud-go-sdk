// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNetworkRuleAttributesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeNetworkRuleAttributesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeNetworkRuleAttributesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeNetworkRuleAttributesResponseBody) *DescribeNetworkRuleAttributesResponse
	GetBody() *DescribeNetworkRuleAttributesResponseBody
}

type DescribeNetworkRuleAttributesResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNetworkRuleAttributesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNetworkRuleAttributesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkRuleAttributesResponse) GoString() string {
	return s.String()
}

func (s *DescribeNetworkRuleAttributesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeNetworkRuleAttributesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeNetworkRuleAttributesResponse) GetBody() *DescribeNetworkRuleAttributesResponseBody {
	return s.Body
}

func (s *DescribeNetworkRuleAttributesResponse) SetHeaders(v map[string]*string) *DescribeNetworkRuleAttributesResponse {
	s.Headers = v
	return s
}

func (s *DescribeNetworkRuleAttributesResponse) SetStatusCode(v int32) *DescribeNetworkRuleAttributesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNetworkRuleAttributesResponse) SetBody(v *DescribeNetworkRuleAttributesResponseBody) *DescribeNetworkRuleAttributesResponse {
	s.Body = v
	return s
}

func (s *DescribeNetworkRuleAttributesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
