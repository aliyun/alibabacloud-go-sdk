// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateResponseRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateResponseRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateResponseRuleResponse
	GetStatusCode() *int32
	SetBody(v *UpdateResponseRuleResponseBody) *UpdateResponseRuleResponse
	GetBody() *UpdateResponseRuleResponseBody
}

type UpdateResponseRuleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateResponseRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateResponseRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateResponseRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateResponseRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateResponseRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateResponseRuleResponse) GetBody() *UpdateResponseRuleResponseBody {
	return s.Body
}

func (s *UpdateResponseRuleResponse) SetHeaders(v map[string]*string) *UpdateResponseRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateResponseRuleResponse) SetStatusCode(v int32) *UpdateResponseRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateResponseRuleResponse) SetBody(v *UpdateResponseRuleResponseBody) *UpdateResponseRuleResponse {
	s.Body = v
	return s
}

func (s *UpdateResponseRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
