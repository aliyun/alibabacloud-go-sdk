// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTransitRouterVpnAttachmentAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateTransitRouterVpnAttachmentAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateTransitRouterVpnAttachmentAttributeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateTransitRouterVpnAttachmentAttributeResponseBody) *UpdateTransitRouterVpnAttachmentAttributeResponse
	GetBody() *UpdateTransitRouterVpnAttachmentAttributeResponseBody
}

type UpdateTransitRouterVpnAttachmentAttributeResponse struct {
	Headers    map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTransitRouterVpnAttachmentAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTransitRouterVpnAttachmentAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateTransitRouterVpnAttachmentAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateTransitRouterVpnAttachmentAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateTransitRouterVpnAttachmentAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTransitRouterVpnAttachmentAttributeResponse) GetBody() *UpdateTransitRouterVpnAttachmentAttributeResponseBody {
	return s.Body
}

func (s *UpdateTransitRouterVpnAttachmentAttributeResponse) SetHeaders(v map[string]*string) *UpdateTransitRouterVpnAttachmentAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateTransitRouterVpnAttachmentAttributeResponse) SetStatusCode(v int32) *UpdateTransitRouterVpnAttachmentAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTransitRouterVpnAttachmentAttributeResponse) SetBody(v *UpdateTransitRouterVpnAttachmentAttributeResponseBody) *UpdateTransitRouterVpnAttachmentAttributeResponse {
	s.Body = v
	return s
}

func (s *UpdateTransitRouterVpnAttachmentAttributeResponse) Validate() error {
	return dara.Validate(s)
}
