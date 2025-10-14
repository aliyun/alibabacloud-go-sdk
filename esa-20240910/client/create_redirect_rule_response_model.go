// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRedirectRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRedirectRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRedirectRuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateRedirectRuleResponseBody) *CreateRedirectRuleResponse
	GetBody() *CreateRedirectRuleResponseBody
}

type CreateRedirectRuleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRedirectRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRedirectRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRedirectRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateRedirectRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRedirectRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRedirectRuleResponse) GetBody() *CreateRedirectRuleResponseBody {
	return s.Body
}

func (s *CreateRedirectRuleResponse) SetHeaders(v map[string]*string) *CreateRedirectRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateRedirectRuleResponse) SetStatusCode(v int32) *CreateRedirectRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRedirectRuleResponse) SetBody(v *CreateRedirectRuleResponseBody) *CreateRedirectRuleResponse {
	s.Body = v
	return s
}

func (s *CreateRedirectRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
