// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePluginAttachmentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePluginAttachmentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePluginAttachmentResponse
	GetStatusCode() *int32
	SetBody(v *CreatePluginAttachmentResponseBody) *CreatePluginAttachmentResponse
	GetBody() *CreatePluginAttachmentResponseBody
}

type CreatePluginAttachmentResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePluginAttachmentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePluginAttachmentResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePluginAttachmentResponse) GoString() string {
	return s.String()
}

func (s *CreatePluginAttachmentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePluginAttachmentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePluginAttachmentResponse) GetBody() *CreatePluginAttachmentResponseBody {
	return s.Body
}

func (s *CreatePluginAttachmentResponse) SetHeaders(v map[string]*string) *CreatePluginAttachmentResponse {
	s.Headers = v
	return s
}

func (s *CreatePluginAttachmentResponse) SetStatusCode(v int32) *CreatePluginAttachmentResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePluginAttachmentResponse) SetBody(v *CreatePluginAttachmentResponseBody) *CreatePluginAttachmentResponse {
	s.Body = v
	return s
}

func (s *CreatePluginAttachmentResponse) Validate() error {
	return dara.Validate(s)
}
