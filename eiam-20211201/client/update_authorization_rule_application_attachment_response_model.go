// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAuthorizationRuleApplicationAttachmentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAuthorizationRuleApplicationAttachmentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAuthorizationRuleApplicationAttachmentResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAuthorizationRuleApplicationAttachmentResponseBody) *UpdateAuthorizationRuleApplicationAttachmentResponse
	GetBody() *UpdateAuthorizationRuleApplicationAttachmentResponseBody
}

type UpdateAuthorizationRuleApplicationAttachmentResponse struct {
	Headers    map[string]*string                                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAuthorizationRuleApplicationAttachmentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAuthorizationRuleApplicationAttachmentResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAuthorizationRuleApplicationAttachmentResponse) GoString() string {
	return s.String()
}

func (s *UpdateAuthorizationRuleApplicationAttachmentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAuthorizationRuleApplicationAttachmentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAuthorizationRuleApplicationAttachmentResponse) GetBody() *UpdateAuthorizationRuleApplicationAttachmentResponseBody {
	return s.Body
}

func (s *UpdateAuthorizationRuleApplicationAttachmentResponse) SetHeaders(v map[string]*string) *UpdateAuthorizationRuleApplicationAttachmentResponse {
	s.Headers = v
	return s
}

func (s *UpdateAuthorizationRuleApplicationAttachmentResponse) SetStatusCode(v int32) *UpdateAuthorizationRuleApplicationAttachmentResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAuthorizationRuleApplicationAttachmentResponse) SetBody(v *UpdateAuthorizationRuleApplicationAttachmentResponseBody) *UpdateAuthorizationRuleApplicationAttachmentResponse {
	s.Body = v
	return s
}

func (s *UpdateAuthorizationRuleApplicationAttachmentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
