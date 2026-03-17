// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGrantSagRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeGrantSagRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeGrantSagRulesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeGrantSagRulesResponseBody) *DescribeGrantSagRulesResponse
	GetBody() *DescribeGrantSagRulesResponseBody
}

type DescribeGrantSagRulesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeGrantSagRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeGrantSagRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeGrantSagRulesResponse) GoString() string {
	return s.String()
}

func (s *DescribeGrantSagRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeGrantSagRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeGrantSagRulesResponse) GetBody() *DescribeGrantSagRulesResponseBody {
	return s.Body
}

func (s *DescribeGrantSagRulesResponse) SetHeaders(v map[string]*string) *DescribeGrantSagRulesResponse {
	s.Headers = v
	return s
}

func (s *DescribeGrantSagRulesResponse) SetStatusCode(v int32) *DescribeGrantSagRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGrantSagRulesResponse) SetBody(v *DescribeGrantSagRulesResponseBody) *DescribeGrantSagRulesResponse {
	s.Body = v
	return s
}

func (s *DescribeGrantSagRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
