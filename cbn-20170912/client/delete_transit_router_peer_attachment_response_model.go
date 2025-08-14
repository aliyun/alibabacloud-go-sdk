// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTransitRouterPeerAttachmentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteTransitRouterPeerAttachmentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteTransitRouterPeerAttachmentResponse
	GetStatusCode() *int32
	SetBody(v *DeleteTransitRouterPeerAttachmentResponseBody) *DeleteTransitRouterPeerAttachmentResponse
	GetBody() *DeleteTransitRouterPeerAttachmentResponseBody
}

type DeleteTransitRouterPeerAttachmentResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTransitRouterPeerAttachmentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTransitRouterPeerAttachmentResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteTransitRouterPeerAttachmentResponse) GoString() string {
	return s.String()
}

func (s *DeleteTransitRouterPeerAttachmentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteTransitRouterPeerAttachmentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteTransitRouterPeerAttachmentResponse) GetBody() *DeleteTransitRouterPeerAttachmentResponseBody {
	return s.Body
}

func (s *DeleteTransitRouterPeerAttachmentResponse) SetHeaders(v map[string]*string) *DeleteTransitRouterPeerAttachmentResponse {
	s.Headers = v
	return s
}

func (s *DeleteTransitRouterPeerAttachmentResponse) SetStatusCode(v int32) *DeleteTransitRouterPeerAttachmentResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTransitRouterPeerAttachmentResponse) SetBody(v *DeleteTransitRouterPeerAttachmentResponseBody) *DeleteTransitRouterPeerAttachmentResponse {
	s.Body = v
	return s
}

func (s *DeleteTransitRouterPeerAttachmentResponse) Validate() error {
	return dara.Validate(s)
}
