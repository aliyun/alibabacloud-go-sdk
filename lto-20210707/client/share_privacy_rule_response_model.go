// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSharePrivacyRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SharePrivacyRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SharePrivacyRuleResponse
	GetStatusCode() *int32
	SetBody(v *SharePrivacyRuleResponseBody) *SharePrivacyRuleResponse
	GetBody() *SharePrivacyRuleResponseBody
}

type SharePrivacyRuleResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SharePrivacyRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SharePrivacyRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s SharePrivacyRuleResponse) GoString() string {
	return s.String()
}

func (s *SharePrivacyRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SharePrivacyRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SharePrivacyRuleResponse) GetBody() *SharePrivacyRuleResponseBody {
	return s.Body
}

func (s *SharePrivacyRuleResponse) SetHeaders(v map[string]*string) *SharePrivacyRuleResponse {
	s.Headers = v
	return s
}

func (s *SharePrivacyRuleResponse) SetStatusCode(v int32) *SharePrivacyRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *SharePrivacyRuleResponse) SetBody(v *SharePrivacyRuleResponseBody) *SharePrivacyRuleResponse {
	s.Body = v
	return s
}

func (s *SharePrivacyRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
