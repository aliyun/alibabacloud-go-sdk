// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePolicyAttachmentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePolicyAttachmentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePolicyAttachmentResponse
	GetStatusCode() *int32
	SetBody(v *CreatePolicyAttachmentResponseBody) *CreatePolicyAttachmentResponse
	GetBody() *CreatePolicyAttachmentResponseBody
}

type CreatePolicyAttachmentResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePolicyAttachmentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePolicyAttachmentResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePolicyAttachmentResponse) GoString() string {
	return s.String()
}

func (s *CreatePolicyAttachmentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePolicyAttachmentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePolicyAttachmentResponse) GetBody() *CreatePolicyAttachmentResponseBody {
	return s.Body
}

func (s *CreatePolicyAttachmentResponse) SetHeaders(v map[string]*string) *CreatePolicyAttachmentResponse {
	s.Headers = v
	return s
}

func (s *CreatePolicyAttachmentResponse) SetStatusCode(v int32) *CreatePolicyAttachmentResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePolicyAttachmentResponse) SetBody(v *CreatePolicyAttachmentResponseBody) *CreatePolicyAttachmentResponse {
	s.Body = v
	return s
}

func (s *CreatePolicyAttachmentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
