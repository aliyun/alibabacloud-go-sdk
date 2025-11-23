// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDesensitizationRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddDesensitizationRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddDesensitizationRuleResponse
	GetStatusCode() *int32
	SetBody(v *AddDesensitizationRuleResponseBody) *AddDesensitizationRuleResponse
	GetBody() *AddDesensitizationRuleResponseBody
}

type AddDesensitizationRuleResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddDesensitizationRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddDesensitizationRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s AddDesensitizationRuleResponse) GoString() string {
	return s.String()
}

func (s *AddDesensitizationRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddDesensitizationRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddDesensitizationRuleResponse) GetBody() *AddDesensitizationRuleResponseBody {
	return s.Body
}

func (s *AddDesensitizationRuleResponse) SetHeaders(v map[string]*string) *AddDesensitizationRuleResponse {
	s.Headers = v
	return s
}

func (s *AddDesensitizationRuleResponse) SetStatusCode(v int32) *AddDesensitizationRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *AddDesensitizationRuleResponse) SetBody(v *AddDesensitizationRuleResponseBody) *AddDesensitizationRuleResponse {
	s.Body = v
	return s
}

func (s *AddDesensitizationRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
