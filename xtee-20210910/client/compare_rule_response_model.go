// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCompareRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CompareRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CompareRuleResponse
	GetStatusCode() *int32
	SetBody(v *CompareRuleResponseBody) *CompareRuleResponse
	GetBody() *CompareRuleResponseBody
}

type CompareRuleResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CompareRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CompareRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CompareRuleResponse) GoString() string {
	return s.String()
}

func (s *CompareRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CompareRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CompareRuleResponse) GetBody() *CompareRuleResponseBody {
	return s.Body
}

func (s *CompareRuleResponse) SetHeaders(v map[string]*string) *CompareRuleResponse {
	s.Headers = v
	return s
}

func (s *CompareRuleResponse) SetStatusCode(v int32) *CompareRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CompareRuleResponse) SetBody(v *CompareRuleResponseBody) *CompareRuleResponse {
	s.Body = v
	return s
}

func (s *CompareRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
