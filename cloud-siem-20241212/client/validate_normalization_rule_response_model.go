// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateNormalizationRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ValidateNormalizationRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ValidateNormalizationRuleResponse
	GetStatusCode() *int32
	SetBody(v *ValidateNormalizationRuleResponseBody) *ValidateNormalizationRuleResponse
	GetBody() *ValidateNormalizationRuleResponseBody
}

type ValidateNormalizationRuleResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ValidateNormalizationRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ValidateNormalizationRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s ValidateNormalizationRuleResponse) GoString() string {
	return s.String()
}

func (s *ValidateNormalizationRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ValidateNormalizationRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ValidateNormalizationRuleResponse) GetBody() *ValidateNormalizationRuleResponseBody {
	return s.Body
}

func (s *ValidateNormalizationRuleResponse) SetHeaders(v map[string]*string) *ValidateNormalizationRuleResponse {
	s.Headers = v
	return s
}

func (s *ValidateNormalizationRuleResponse) SetStatusCode(v int32) *ValidateNormalizationRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ValidateNormalizationRuleResponse) SetBody(v *ValidateNormalizationRuleResponseBody) *ValidateNormalizationRuleResponse {
	s.Body = v
	return s
}

func (s *ValidateNormalizationRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
