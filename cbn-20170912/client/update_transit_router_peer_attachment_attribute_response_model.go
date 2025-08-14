// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTransitRouterPeerAttachmentAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateTransitRouterPeerAttachmentAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateTransitRouterPeerAttachmentAttributeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateTransitRouterPeerAttachmentAttributeResponseBody) *UpdateTransitRouterPeerAttachmentAttributeResponse
	GetBody() *UpdateTransitRouterPeerAttachmentAttributeResponseBody
}

type UpdateTransitRouterPeerAttachmentAttributeResponse struct {
	Headers    map[string]*string                                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTransitRouterPeerAttachmentAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTransitRouterPeerAttachmentAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateTransitRouterPeerAttachmentAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateTransitRouterPeerAttachmentAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateTransitRouterPeerAttachmentAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTransitRouterPeerAttachmentAttributeResponse) GetBody() *UpdateTransitRouterPeerAttachmentAttributeResponseBody {
	return s.Body
}

func (s *UpdateTransitRouterPeerAttachmentAttributeResponse) SetHeaders(v map[string]*string) *UpdateTransitRouterPeerAttachmentAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateTransitRouterPeerAttachmentAttributeResponse) SetStatusCode(v int32) *UpdateTransitRouterPeerAttachmentAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTransitRouterPeerAttachmentAttributeResponse) SetBody(v *UpdateTransitRouterPeerAttachmentAttributeResponseBody) *UpdateTransitRouterPeerAttachmentAttributeResponse {
	s.Body = v
	return s
}

func (s *UpdateTransitRouterPeerAttachmentAttributeResponse) Validate() error {
	return dara.Validate(s)
}
