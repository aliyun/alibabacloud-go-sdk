// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGrantRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeGrantRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeGrantRulesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeGrantRulesResponseBody) *DescribeGrantRulesResponse
	GetBody() *DescribeGrantRulesResponseBody
}

type DescribeGrantRulesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeGrantRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeGrantRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeGrantRulesResponse) GoString() string {
	return s.String()
}

func (s *DescribeGrantRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeGrantRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeGrantRulesResponse) GetBody() *DescribeGrantRulesResponseBody {
	return s.Body
}

func (s *DescribeGrantRulesResponse) SetHeaders(v map[string]*string) *DescribeGrantRulesResponse {
	s.Headers = v
	return s
}

func (s *DescribeGrantRulesResponse) SetStatusCode(v int32) *DescribeGrantRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGrantRulesResponse) SetBody(v *DescribeGrantRulesResponseBody) *DescribeGrantRulesResponse {
	s.Body = v
	return s
}

func (s *DescribeGrantRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
