// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRuleAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRuleAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRuleAttributeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRuleAttributeResponseBody) *DescribeRuleAttributeResponse
	GetBody() *DescribeRuleAttributeResponseBody
}

type DescribeRuleAttributeResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRuleAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRuleAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRuleAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeRuleAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRuleAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRuleAttributeResponse) GetBody() *DescribeRuleAttributeResponseBody {
	return s.Body
}

func (s *DescribeRuleAttributeResponse) SetHeaders(v map[string]*string) *DescribeRuleAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeRuleAttributeResponse) SetStatusCode(v int32) *DescribeRuleAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRuleAttributeResponse) SetBody(v *DescribeRuleAttributeResponseBody) *DescribeRuleAttributeResponse {
	s.Body = v
	return s
}

func (s *DescribeRuleAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
