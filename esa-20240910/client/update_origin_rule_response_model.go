// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOriginRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateOriginRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateOriginRuleResponse
	GetStatusCode() *int32
	SetBody(v *UpdateOriginRuleResponseBody) *UpdateOriginRuleResponse
	GetBody() *UpdateOriginRuleResponseBody
}

type UpdateOriginRuleResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateOriginRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateOriginRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateOriginRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateOriginRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateOriginRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateOriginRuleResponse) GetBody() *UpdateOriginRuleResponseBody {
	return s.Body
}

func (s *UpdateOriginRuleResponse) SetHeaders(v map[string]*string) *UpdateOriginRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateOriginRuleResponse) SetStatusCode(v int32) *UpdateOriginRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateOriginRuleResponse) SetBody(v *UpdateOriginRuleResponseBody) *UpdateOriginRuleResponse {
	s.Body = v
	return s
}

func (s *UpdateOriginRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
