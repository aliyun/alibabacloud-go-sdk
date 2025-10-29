// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPolicyAttachmentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPolicyAttachmentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPolicyAttachmentResponse
	GetStatusCode() *int32
	SetBody(v *GetPolicyAttachmentResponseBody) *GetPolicyAttachmentResponse
	GetBody() *GetPolicyAttachmentResponseBody
}

type GetPolicyAttachmentResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPolicyAttachmentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPolicyAttachmentResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPolicyAttachmentResponse) GoString() string {
	return s.String()
}

func (s *GetPolicyAttachmentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPolicyAttachmentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPolicyAttachmentResponse) GetBody() *GetPolicyAttachmentResponseBody {
	return s.Body
}

func (s *GetPolicyAttachmentResponse) SetHeaders(v map[string]*string) *GetPolicyAttachmentResponse {
	s.Headers = v
	return s
}

func (s *GetPolicyAttachmentResponse) SetStatusCode(v int32) *GetPolicyAttachmentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPolicyAttachmentResponse) SetBody(v *GetPolicyAttachmentResponseBody) *GetPolicyAttachmentResponse {
	s.Body = v
	return s
}

func (s *GetPolicyAttachmentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
