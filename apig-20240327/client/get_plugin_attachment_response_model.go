// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPluginAttachmentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPluginAttachmentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPluginAttachmentResponse
	GetStatusCode() *int32
	SetBody(v *GetPluginAttachmentResponseBody) *GetPluginAttachmentResponse
	GetBody() *GetPluginAttachmentResponseBody
}

type GetPluginAttachmentResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPluginAttachmentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPluginAttachmentResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPluginAttachmentResponse) GoString() string {
	return s.String()
}

func (s *GetPluginAttachmentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPluginAttachmentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPluginAttachmentResponse) GetBody() *GetPluginAttachmentResponseBody {
	return s.Body
}

func (s *GetPluginAttachmentResponse) SetHeaders(v map[string]*string) *GetPluginAttachmentResponse {
	s.Headers = v
	return s
}

func (s *GetPluginAttachmentResponse) SetStatusCode(v int32) *GetPluginAttachmentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPluginAttachmentResponse) SetBody(v *GetPluginAttachmentResponseBody) *GetPluginAttachmentResponse {
	s.Body = v
	return s
}

func (s *GetPluginAttachmentResponse) Validate() error {
	return dara.Validate(s)
}
