// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetRuleResponse
	GetStatusCode() *int32
	SetBody(v *SetRuleResponseBody) *SetRuleResponse
	GetBody() *SetRuleResponseBody
}

type SetRuleResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s SetRuleResponse) GoString() string {
	return s.String()
}

func (s *SetRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetRuleResponse) GetBody() *SetRuleResponseBody {
	return s.Body
}

func (s *SetRuleResponse) SetHeaders(v map[string]*string) *SetRuleResponse {
	s.Headers = v
	return s
}

func (s *SetRuleResponse) SetStatusCode(v int32) *SetRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *SetRuleResponse) SetBody(v *SetRuleResponseBody) *SetRuleResponse {
	s.Body = v
	return s
}

func (s *SetRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
