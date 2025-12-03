// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePushRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdatePushRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdatePushRuleResponse
	GetStatusCode() *int32
	SetBody(v *UpdatePushRuleResponseBody) *UpdatePushRuleResponse
	GetBody() *UpdatePushRuleResponseBody
}

type UpdatePushRuleResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePushRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePushRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdatePushRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdatePushRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdatePushRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdatePushRuleResponse) GetBody() *UpdatePushRuleResponseBody {
	return s.Body
}

func (s *UpdatePushRuleResponse) SetHeaders(v map[string]*string) *UpdatePushRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdatePushRuleResponse) SetStatusCode(v int32) *UpdatePushRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePushRuleResponse) SetBody(v *UpdatePushRuleResponseBody) *UpdatePushRuleResponse {
	s.Body = v
	return s
}

func (s *UpdatePushRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
