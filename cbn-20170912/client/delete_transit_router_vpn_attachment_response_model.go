// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTransitRouterVpnAttachmentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteTransitRouterVpnAttachmentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteTransitRouterVpnAttachmentResponse
	GetStatusCode() *int32
	SetBody(v *DeleteTransitRouterVpnAttachmentResponseBody) *DeleteTransitRouterVpnAttachmentResponse
	GetBody() *DeleteTransitRouterVpnAttachmentResponseBody
}

type DeleteTransitRouterVpnAttachmentResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTransitRouterVpnAttachmentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTransitRouterVpnAttachmentResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteTransitRouterVpnAttachmentResponse) GoString() string {
	return s.String()
}

func (s *DeleteTransitRouterVpnAttachmentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteTransitRouterVpnAttachmentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteTransitRouterVpnAttachmentResponse) GetBody() *DeleteTransitRouterVpnAttachmentResponseBody {
	return s.Body
}

func (s *DeleteTransitRouterVpnAttachmentResponse) SetHeaders(v map[string]*string) *DeleteTransitRouterVpnAttachmentResponse {
	s.Headers = v
	return s
}

func (s *DeleteTransitRouterVpnAttachmentResponse) SetStatusCode(v int32) *DeleteTransitRouterVpnAttachmentResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTransitRouterVpnAttachmentResponse) SetBody(v *DeleteTransitRouterVpnAttachmentResponseBody) *DeleteTransitRouterVpnAttachmentResponse {
	s.Body = v
	return s
}

func (s *DeleteTransitRouterVpnAttachmentResponse) Validate() error {
	return dara.Validate(s)
}
