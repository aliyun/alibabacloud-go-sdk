// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventRuleAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEventRuleAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEventRuleAttributeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEventRuleAttributeResponseBody) *DescribeEventRuleAttributeResponse
	GetBody() *DescribeEventRuleAttributeResponseBody
}

type DescribeEventRuleAttributeResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEventRuleAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEventRuleAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventRuleAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeEventRuleAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEventRuleAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEventRuleAttributeResponse) GetBody() *DescribeEventRuleAttributeResponseBody {
	return s.Body
}

func (s *DescribeEventRuleAttributeResponse) SetHeaders(v map[string]*string) *DescribeEventRuleAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeEventRuleAttributeResponse) SetStatusCode(v int32) *DescribeEventRuleAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEventRuleAttributeResponse) SetBody(v *DescribeEventRuleAttributeResponseBody) *DescribeEventRuleAttributeResponse {
	s.Body = v
	return s
}

func (s *DescribeEventRuleAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
