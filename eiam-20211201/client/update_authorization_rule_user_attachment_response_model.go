// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAuthorizationRuleUserAttachmentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAuthorizationRuleUserAttachmentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAuthorizationRuleUserAttachmentResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAuthorizationRuleUserAttachmentResponseBody) *UpdateAuthorizationRuleUserAttachmentResponse
	GetBody() *UpdateAuthorizationRuleUserAttachmentResponseBody
}

type UpdateAuthorizationRuleUserAttachmentResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAuthorizationRuleUserAttachmentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAuthorizationRuleUserAttachmentResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAuthorizationRuleUserAttachmentResponse) GoString() string {
	return s.String()
}

func (s *UpdateAuthorizationRuleUserAttachmentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAuthorizationRuleUserAttachmentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAuthorizationRuleUserAttachmentResponse) GetBody() *UpdateAuthorizationRuleUserAttachmentResponseBody {
	return s.Body
}

func (s *UpdateAuthorizationRuleUserAttachmentResponse) SetHeaders(v map[string]*string) *UpdateAuthorizationRuleUserAttachmentResponse {
	s.Headers = v
	return s
}

func (s *UpdateAuthorizationRuleUserAttachmentResponse) SetStatusCode(v int32) *UpdateAuthorizationRuleUserAttachmentResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAuthorizationRuleUserAttachmentResponse) SetBody(v *UpdateAuthorizationRuleUserAttachmentResponseBody) *UpdateAuthorizationRuleUserAttachmentResponse {
	s.Body = v
	return s
}

func (s *UpdateAuthorizationRuleUserAttachmentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
