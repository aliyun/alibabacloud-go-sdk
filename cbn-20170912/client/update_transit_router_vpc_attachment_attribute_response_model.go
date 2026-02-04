// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTransitRouterVpcAttachmentAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateTransitRouterVpcAttachmentAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateTransitRouterVpcAttachmentAttributeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateTransitRouterVpcAttachmentAttributeResponseBody) *UpdateTransitRouterVpcAttachmentAttributeResponse
	GetBody() *UpdateTransitRouterVpcAttachmentAttributeResponseBody
}

type UpdateTransitRouterVpcAttachmentAttributeResponse struct {
	Headers    map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTransitRouterVpcAttachmentAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTransitRouterVpcAttachmentAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateTransitRouterVpcAttachmentAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateTransitRouterVpcAttachmentAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateTransitRouterVpcAttachmentAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTransitRouterVpcAttachmentAttributeResponse) GetBody() *UpdateTransitRouterVpcAttachmentAttributeResponseBody {
	return s.Body
}

func (s *UpdateTransitRouterVpcAttachmentAttributeResponse) SetHeaders(v map[string]*string) *UpdateTransitRouterVpcAttachmentAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateTransitRouterVpcAttachmentAttributeResponse) SetStatusCode(v int32) *UpdateTransitRouterVpcAttachmentAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTransitRouterVpcAttachmentAttributeResponse) SetBody(v *UpdateTransitRouterVpcAttachmentAttributeResponseBody) *UpdateTransitRouterVpcAttachmentAttributeResponse {
	s.Body = v
	return s
}

func (s *UpdateTransitRouterVpcAttachmentAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
