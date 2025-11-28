// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAllPrivacyRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAllPrivacyRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAllPrivacyRuleResponse
	GetStatusCode() *int32
	SetBody(v *ListAllPrivacyRuleResponseBody) *ListAllPrivacyRuleResponse
	GetBody() *ListAllPrivacyRuleResponseBody
}

type ListAllPrivacyRuleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAllPrivacyRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAllPrivacyRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAllPrivacyRuleResponse) GoString() string {
	return s.String()
}

func (s *ListAllPrivacyRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAllPrivacyRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAllPrivacyRuleResponse) GetBody() *ListAllPrivacyRuleResponseBody {
	return s.Body
}

func (s *ListAllPrivacyRuleResponse) SetHeaders(v map[string]*string) *ListAllPrivacyRuleResponse {
	s.Headers = v
	return s
}

func (s *ListAllPrivacyRuleResponse) SetStatusCode(v int32) *ListAllPrivacyRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAllPrivacyRuleResponse) SetBody(v *ListAllPrivacyRuleResponseBody) *ListAllPrivacyRuleResponse {
	s.Body = v
	return s
}

func (s *ListAllPrivacyRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
