// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLocalityRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateLocalityRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateLocalityRuleResponse
	GetStatusCode() *int32
	SetBody(v *UpdateLocalityRuleResponseBody) *UpdateLocalityRuleResponse
	GetBody() *UpdateLocalityRuleResponseBody
}

type UpdateLocalityRuleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLocalityRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLocalityRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateLocalityRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateLocalityRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateLocalityRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateLocalityRuleResponse) GetBody() *UpdateLocalityRuleResponseBody {
	return s.Body
}

func (s *UpdateLocalityRuleResponse) SetHeaders(v map[string]*string) *UpdateLocalityRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateLocalityRuleResponse) SetStatusCode(v int32) *UpdateLocalityRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLocalityRuleResponse) SetBody(v *UpdateLocalityRuleResponseBody) *UpdateLocalityRuleResponse {
	s.Body = v
	return s
}

func (s *UpdateLocalityRuleResponse) Validate() error {
	return dara.Validate(s)
}
