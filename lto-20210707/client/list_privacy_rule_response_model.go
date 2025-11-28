// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrivacyRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPrivacyRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPrivacyRuleResponse
	GetStatusCode() *int32
	SetBody(v *ListPrivacyRuleResponseBody) *ListPrivacyRuleResponse
	GetBody() *ListPrivacyRuleResponseBody
}

type ListPrivacyRuleResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPrivacyRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPrivacyRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPrivacyRuleResponse) GoString() string {
	return s.String()
}

func (s *ListPrivacyRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPrivacyRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPrivacyRuleResponse) GetBody() *ListPrivacyRuleResponseBody {
	return s.Body
}

func (s *ListPrivacyRuleResponse) SetHeaders(v map[string]*string) *ListPrivacyRuleResponse {
	s.Headers = v
	return s
}

func (s *ListPrivacyRuleResponse) SetStatusCode(v int32) *ListPrivacyRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPrivacyRuleResponse) SetBody(v *ListPrivacyRuleResponseBody) *ListPrivacyRuleResponse {
	s.Body = v
	return s
}

func (s *ListPrivacyRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
