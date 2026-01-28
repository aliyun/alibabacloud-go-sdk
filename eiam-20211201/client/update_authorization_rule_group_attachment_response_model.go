// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAuthorizationRuleGroupAttachmentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAuthorizationRuleGroupAttachmentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAuthorizationRuleGroupAttachmentResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAuthorizationRuleGroupAttachmentResponseBody) *UpdateAuthorizationRuleGroupAttachmentResponse
	GetBody() *UpdateAuthorizationRuleGroupAttachmentResponseBody
}

type UpdateAuthorizationRuleGroupAttachmentResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAuthorizationRuleGroupAttachmentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAuthorizationRuleGroupAttachmentResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAuthorizationRuleGroupAttachmentResponse) GoString() string {
	return s.String()
}

func (s *UpdateAuthorizationRuleGroupAttachmentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAuthorizationRuleGroupAttachmentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAuthorizationRuleGroupAttachmentResponse) GetBody() *UpdateAuthorizationRuleGroupAttachmentResponseBody {
	return s.Body
}

func (s *UpdateAuthorizationRuleGroupAttachmentResponse) SetHeaders(v map[string]*string) *UpdateAuthorizationRuleGroupAttachmentResponse {
	s.Headers = v
	return s
}

func (s *UpdateAuthorizationRuleGroupAttachmentResponse) SetStatusCode(v int32) *UpdateAuthorizationRuleGroupAttachmentResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAuthorizationRuleGroupAttachmentResponse) SetBody(v *UpdateAuthorizationRuleGroupAttachmentResponseBody) *UpdateAuthorizationRuleGroupAttachmentResponse {
	s.Body = v
	return s
}

func (s *UpdateAuthorizationRuleGroupAttachmentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
