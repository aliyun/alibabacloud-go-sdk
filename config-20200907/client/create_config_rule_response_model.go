// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateConfigRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateConfigRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateConfigRuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateConfigRuleResponseBody) *CreateConfigRuleResponse
	GetBody() *CreateConfigRuleResponseBody
}

type CreateConfigRuleResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateConfigRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateConfigRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateConfigRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateConfigRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateConfigRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateConfigRuleResponse) GetBody() *CreateConfigRuleResponseBody {
	return s.Body
}

func (s *CreateConfigRuleResponse) SetHeaders(v map[string]*string) *CreateConfigRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateConfigRuleResponse) SetStatusCode(v int32) *CreateConfigRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateConfigRuleResponse) SetBody(v *CreateConfigRuleResponseBody) *CreateConfigRuleResponse {
	s.Body = v
	return s
}

func (s *CreateConfigRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
