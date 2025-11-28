// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePrivacyRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePrivacyRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePrivacyRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeletePrivacyRuleResponseBody) *DeletePrivacyRuleResponse
	GetBody() *DeletePrivacyRuleResponseBody
}

type DeletePrivacyRuleResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePrivacyRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePrivacyRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePrivacyRuleResponse) GoString() string {
	return s.String()
}

func (s *DeletePrivacyRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePrivacyRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePrivacyRuleResponse) GetBody() *DeletePrivacyRuleResponseBody {
	return s.Body
}

func (s *DeletePrivacyRuleResponse) SetHeaders(v map[string]*string) *DeletePrivacyRuleResponse {
	s.Headers = v
	return s
}

func (s *DeletePrivacyRuleResponse) SetStatusCode(v int32) *DeletePrivacyRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePrivacyRuleResponse) SetBody(v *DeletePrivacyRuleResponseBody) *DeletePrivacyRuleResponse {
	s.Body = v
	return s
}

func (s *DeletePrivacyRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
