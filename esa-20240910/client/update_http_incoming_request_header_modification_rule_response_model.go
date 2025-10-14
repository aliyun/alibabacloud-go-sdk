// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHttpIncomingRequestHeaderModificationRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateHttpIncomingRequestHeaderModificationRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateHttpIncomingRequestHeaderModificationRuleResponse
	GetStatusCode() *int32
	SetBody(v *UpdateHttpIncomingRequestHeaderModificationRuleResponseBody) *UpdateHttpIncomingRequestHeaderModificationRuleResponse
	GetBody() *UpdateHttpIncomingRequestHeaderModificationRuleResponseBody
}

type UpdateHttpIncomingRequestHeaderModificationRuleResponse struct {
	Headers    map[string]*string                                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateHttpIncomingRequestHeaderModificationRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateHttpIncomingRequestHeaderModificationRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateHttpIncomingRequestHeaderModificationRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateHttpIncomingRequestHeaderModificationRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateHttpIncomingRequestHeaderModificationRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateHttpIncomingRequestHeaderModificationRuleResponse) GetBody() *UpdateHttpIncomingRequestHeaderModificationRuleResponseBody {
	return s.Body
}

func (s *UpdateHttpIncomingRequestHeaderModificationRuleResponse) SetHeaders(v map[string]*string) *UpdateHttpIncomingRequestHeaderModificationRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateHttpIncomingRequestHeaderModificationRuleResponse) SetStatusCode(v int32) *UpdateHttpIncomingRequestHeaderModificationRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateHttpIncomingRequestHeaderModificationRuleResponse) SetBody(v *UpdateHttpIncomingRequestHeaderModificationRuleResponseBody) *UpdateHttpIncomingRequestHeaderModificationRuleResponse {
	s.Body = v
	return s
}

func (s *UpdateHttpIncomingRequestHeaderModificationRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
