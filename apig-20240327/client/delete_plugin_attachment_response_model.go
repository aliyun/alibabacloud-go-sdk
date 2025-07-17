// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePluginAttachmentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePluginAttachmentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePluginAttachmentResponse
	GetStatusCode() *int32
	SetBody(v *DeletePluginAttachmentResponseBody) *DeletePluginAttachmentResponse
	GetBody() *DeletePluginAttachmentResponseBody
}

type DeletePluginAttachmentResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePluginAttachmentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePluginAttachmentResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePluginAttachmentResponse) GoString() string {
	return s.String()
}

func (s *DeletePluginAttachmentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePluginAttachmentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePluginAttachmentResponse) GetBody() *DeletePluginAttachmentResponseBody {
	return s.Body
}

func (s *DeletePluginAttachmentResponse) SetHeaders(v map[string]*string) *DeletePluginAttachmentResponse {
	s.Headers = v
	return s
}

func (s *DeletePluginAttachmentResponse) SetStatusCode(v int32) *DeletePluginAttachmentResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePluginAttachmentResponse) SetBody(v *DeletePluginAttachmentResponseBody) *DeletePluginAttachmentResponse {
	s.Body = v
	return s
}

func (s *DeletePluginAttachmentResponse) Validate() error {
	return dara.Validate(s)
}
