// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHttpIncomingResponseHeaderModificationRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateHttpIncomingResponseHeaderModificationRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateHttpIncomingResponseHeaderModificationRuleResponse
	GetStatusCode() *int32
	SetBody(v *UpdateHttpIncomingResponseHeaderModificationRuleResponseBody) *UpdateHttpIncomingResponseHeaderModificationRuleResponse
	GetBody() *UpdateHttpIncomingResponseHeaderModificationRuleResponseBody
}

type UpdateHttpIncomingResponseHeaderModificationRuleResponse struct {
	Headers    map[string]*string                                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateHttpIncomingResponseHeaderModificationRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateHttpIncomingResponseHeaderModificationRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateHttpIncomingResponseHeaderModificationRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateHttpIncomingResponseHeaderModificationRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateHttpIncomingResponseHeaderModificationRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateHttpIncomingResponseHeaderModificationRuleResponse) GetBody() *UpdateHttpIncomingResponseHeaderModificationRuleResponseBody {
	return s.Body
}

func (s *UpdateHttpIncomingResponseHeaderModificationRuleResponse) SetHeaders(v map[string]*string) *UpdateHttpIncomingResponseHeaderModificationRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateHttpIncomingResponseHeaderModificationRuleResponse) SetStatusCode(v int32) *UpdateHttpIncomingResponseHeaderModificationRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateHttpIncomingResponseHeaderModificationRuleResponse) SetBody(v *UpdateHttpIncomingResponseHeaderModificationRuleResponseBody) *UpdateHttpIncomingResponseHeaderModificationRuleResponse {
	s.Body = v
	return s
}

func (s *UpdateHttpIncomingResponseHeaderModificationRuleResponse) Validate() error {
	return dara.Validate(s)
}
