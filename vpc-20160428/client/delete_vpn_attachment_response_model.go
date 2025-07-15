// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVpnAttachmentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteVpnAttachmentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteVpnAttachmentResponse
	GetStatusCode() *int32
	SetBody(v *DeleteVpnAttachmentResponseBody) *DeleteVpnAttachmentResponse
	GetBody() *DeleteVpnAttachmentResponseBody
}

type DeleteVpnAttachmentResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVpnAttachmentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteVpnAttachmentResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteVpnAttachmentResponse) GoString() string {
	return s.String()
}

func (s *DeleteVpnAttachmentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteVpnAttachmentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteVpnAttachmentResponse) GetBody() *DeleteVpnAttachmentResponseBody {
	return s.Body
}

func (s *DeleteVpnAttachmentResponse) SetHeaders(v map[string]*string) *DeleteVpnAttachmentResponse {
	s.Headers = v
	return s
}

func (s *DeleteVpnAttachmentResponse) SetStatusCode(v int32) *DeleteVpnAttachmentResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVpnAttachmentResponse) SetBody(v *DeleteVpnAttachmentResponseBody) *DeleteVpnAttachmentResponse {
	s.Body = v
	return s
}

func (s *DeleteVpnAttachmentResponse) Validate() error {
	return dara.Validate(s)
}
