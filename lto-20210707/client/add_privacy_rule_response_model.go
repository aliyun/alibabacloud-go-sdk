// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddPrivacyRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddPrivacyRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddPrivacyRuleResponse
	GetStatusCode() *int32
	SetBody(v *AddPrivacyRuleResponseBody) *AddPrivacyRuleResponse
	GetBody() *AddPrivacyRuleResponseBody
}

type AddPrivacyRuleResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddPrivacyRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddPrivacyRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s AddPrivacyRuleResponse) GoString() string {
	return s.String()
}

func (s *AddPrivacyRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddPrivacyRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddPrivacyRuleResponse) GetBody() *AddPrivacyRuleResponseBody {
	return s.Body
}

func (s *AddPrivacyRuleResponse) SetHeaders(v map[string]*string) *AddPrivacyRuleResponse {
	s.Headers = v
	return s
}

func (s *AddPrivacyRuleResponse) SetStatusCode(v int32) *AddPrivacyRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *AddPrivacyRuleResponse) SetBody(v *AddPrivacyRuleResponseBody) *AddPrivacyRuleResponse {
	s.Body = v
	return s
}

func (s *AddPrivacyRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
