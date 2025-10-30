// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRuleBaseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateRuleBaseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateRuleBaseResponse
	GetStatusCode() *int32
	SetBody(v *UpdateRuleBaseResponseBody) *UpdateRuleBaseResponse
	GetBody() *UpdateRuleBaseResponseBody
}

type UpdateRuleBaseResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRuleBaseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRuleBaseResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateRuleBaseResponse) GoString() string {
	return s.String()
}

func (s *UpdateRuleBaseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateRuleBaseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateRuleBaseResponse) GetBody() *UpdateRuleBaseResponseBody {
	return s.Body
}

func (s *UpdateRuleBaseResponse) SetHeaders(v map[string]*string) *UpdateRuleBaseResponse {
	s.Headers = v
	return s
}

func (s *UpdateRuleBaseResponse) SetStatusCode(v int32) *UpdateRuleBaseResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRuleBaseResponse) SetBody(v *UpdateRuleBaseResponseBody) *UpdateRuleBaseResponse {
	s.Body = v
	return s
}

func (s *UpdateRuleBaseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
