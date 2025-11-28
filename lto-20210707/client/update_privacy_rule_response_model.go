// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePrivacyRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdatePrivacyRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdatePrivacyRuleResponse
	GetStatusCode() *int32
	SetBody(v *UpdatePrivacyRuleResponseBody) *UpdatePrivacyRuleResponse
	GetBody() *UpdatePrivacyRuleResponseBody
}

type UpdatePrivacyRuleResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePrivacyRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePrivacyRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdatePrivacyRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdatePrivacyRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdatePrivacyRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdatePrivacyRuleResponse) GetBody() *UpdatePrivacyRuleResponseBody {
	return s.Body
}

func (s *UpdatePrivacyRuleResponse) SetHeaders(v map[string]*string) *UpdatePrivacyRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdatePrivacyRuleResponse) SetStatusCode(v int32) *UpdatePrivacyRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePrivacyRuleResponse) SetBody(v *UpdatePrivacyRuleResponseBody) *UpdatePrivacyRuleResponse {
	s.Body = v
	return s
}

func (s *UpdatePrivacyRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
