// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBaseSystemRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeBaseSystemRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeBaseSystemRulesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeBaseSystemRulesResponseBody) *DescribeBaseSystemRulesResponse
	GetBody() *DescribeBaseSystemRulesResponseBody
}

type DescribeBaseSystemRulesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBaseSystemRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBaseSystemRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeBaseSystemRulesResponse) GoString() string {
	return s.String()
}

func (s *DescribeBaseSystemRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeBaseSystemRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeBaseSystemRulesResponse) GetBody() *DescribeBaseSystemRulesResponseBody {
	return s.Body
}

func (s *DescribeBaseSystemRulesResponse) SetHeaders(v map[string]*string) *DescribeBaseSystemRulesResponse {
	s.Headers = v
	return s
}

func (s *DescribeBaseSystemRulesResponse) SetStatusCode(v int32) *DescribeBaseSystemRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBaseSystemRulesResponse) SetBody(v *DescribeBaseSystemRulesResponseBody) *DescribeBaseSystemRulesResponse {
	s.Body = v
	return s
}

func (s *DescribeBaseSystemRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
