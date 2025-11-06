// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddMockRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddMockRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddMockRuleResponse
	GetStatusCode() *int32
	SetBody(v *AddMockRuleResponseBody) *AddMockRuleResponse
	GetBody() *AddMockRuleResponseBody
}

type AddMockRuleResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddMockRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddMockRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s AddMockRuleResponse) GoString() string {
	return s.String()
}

func (s *AddMockRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddMockRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddMockRuleResponse) GetBody() *AddMockRuleResponseBody {
	return s.Body
}

func (s *AddMockRuleResponse) SetHeaders(v map[string]*string) *AddMockRuleResponse {
	s.Headers = v
	return s
}

func (s *AddMockRuleResponse) SetStatusCode(v int32) *AddMockRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *AddMockRuleResponse) SetBody(v *AddMockRuleResponseBody) *AddMockRuleResponse {
	s.Body = v
	return s
}

func (s *AddMockRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
