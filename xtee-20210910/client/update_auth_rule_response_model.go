// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAuthRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAuthRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAuthRuleResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAuthRuleResponseBody) *UpdateAuthRuleResponse
	GetBody() *UpdateAuthRuleResponseBody
}

type UpdateAuthRuleResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAuthRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAuthRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAuthRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateAuthRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAuthRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAuthRuleResponse) GetBody() *UpdateAuthRuleResponseBody {
	return s.Body
}

func (s *UpdateAuthRuleResponse) SetHeaders(v map[string]*string) *UpdateAuthRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateAuthRuleResponse) SetStatusCode(v int32) *UpdateAuthRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAuthRuleResponse) SetBody(v *UpdateAuthRuleResponseBody) *UpdateAuthRuleResponse {
	s.Body = v
	return s
}

func (s *UpdateAuthRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
