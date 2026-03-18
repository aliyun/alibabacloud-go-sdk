// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPolicyAttachmentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPolicyAttachmentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPolicyAttachmentResponse
	GetStatusCode() *int32
	SetBody(v *ListPolicyAttachmentResponseBody) *ListPolicyAttachmentResponse
	GetBody() *ListPolicyAttachmentResponseBody
}

type ListPolicyAttachmentResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPolicyAttachmentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPolicyAttachmentResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPolicyAttachmentResponse) GoString() string {
	return s.String()
}

func (s *ListPolicyAttachmentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPolicyAttachmentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPolicyAttachmentResponse) GetBody() *ListPolicyAttachmentResponseBody {
	return s.Body
}

func (s *ListPolicyAttachmentResponse) SetHeaders(v map[string]*string) *ListPolicyAttachmentResponse {
	s.Headers = v
	return s
}

func (s *ListPolicyAttachmentResponse) SetStatusCode(v int32) *ListPolicyAttachmentResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPolicyAttachmentResponse) SetBody(v *ListPolicyAttachmentResponseBody) *ListPolicyAttachmentResponse {
	s.Body = v
	return s
}

func (s *ListPolicyAttachmentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
