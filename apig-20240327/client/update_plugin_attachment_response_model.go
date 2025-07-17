// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePluginAttachmentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdatePluginAttachmentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdatePluginAttachmentResponse
	GetStatusCode() *int32
	SetBody(v *UpdatePluginAttachmentResponseBody) *UpdatePluginAttachmentResponse
	GetBody() *UpdatePluginAttachmentResponseBody
}

type UpdatePluginAttachmentResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePluginAttachmentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePluginAttachmentResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdatePluginAttachmentResponse) GoString() string {
	return s.String()
}

func (s *UpdatePluginAttachmentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdatePluginAttachmentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdatePluginAttachmentResponse) GetBody() *UpdatePluginAttachmentResponseBody {
	return s.Body
}

func (s *UpdatePluginAttachmentResponse) SetHeaders(v map[string]*string) *UpdatePluginAttachmentResponse {
	s.Headers = v
	return s
}

func (s *UpdatePluginAttachmentResponse) SetStatusCode(v int32) *UpdatePluginAttachmentResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePluginAttachmentResponse) SetBody(v *UpdatePluginAttachmentResponseBody) *UpdatePluginAttachmentResponse {
	s.Body = v
	return s
}

func (s *UpdatePluginAttachmentResponse) Validate() error {
	return dara.Validate(s)
}
