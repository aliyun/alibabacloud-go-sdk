// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePolicyAttachmentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePolicyAttachmentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePolicyAttachmentResponse
	GetStatusCode() *int32
	SetBody(v *DeletePolicyAttachmentResponseBody) *DeletePolicyAttachmentResponse
	GetBody() *DeletePolicyAttachmentResponseBody
}

type DeletePolicyAttachmentResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePolicyAttachmentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePolicyAttachmentResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePolicyAttachmentResponse) GoString() string {
	return s.String()
}

func (s *DeletePolicyAttachmentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePolicyAttachmentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePolicyAttachmentResponse) GetBody() *DeletePolicyAttachmentResponseBody {
	return s.Body
}

func (s *DeletePolicyAttachmentResponse) SetHeaders(v map[string]*string) *DeletePolicyAttachmentResponse {
	s.Headers = v
	return s
}

func (s *DeletePolicyAttachmentResponse) SetStatusCode(v int32) *DeletePolicyAttachmentResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePolicyAttachmentResponse) SetBody(v *DeletePolicyAttachmentResponseBody) *DeletePolicyAttachmentResponse {
	s.Body = v
	return s
}

func (s *DeletePolicyAttachmentResponse) Validate() error {
	return dara.Validate(s)
}
