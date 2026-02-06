// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomResponseCodeRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCustomResponseCodeRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCustomResponseCodeRuleResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCustomResponseCodeRuleResponseBody) *UpdateCustomResponseCodeRuleResponse
	GetBody() *UpdateCustomResponseCodeRuleResponseBody
}

type UpdateCustomResponseCodeRuleResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCustomResponseCodeRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCustomResponseCodeRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomResponseCodeRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateCustomResponseCodeRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCustomResponseCodeRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCustomResponseCodeRuleResponse) GetBody() *UpdateCustomResponseCodeRuleResponseBody {
	return s.Body
}

func (s *UpdateCustomResponseCodeRuleResponse) SetHeaders(v map[string]*string) *UpdateCustomResponseCodeRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateCustomResponseCodeRuleResponse) SetStatusCode(v int32) *UpdateCustomResponseCodeRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCustomResponseCodeRuleResponse) SetBody(v *UpdateCustomResponseCodeRuleResponseBody) *UpdateCustomResponseCodeRuleResponse {
	s.Body = v
	return s
}

func (s *UpdateCustomResponseCodeRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
