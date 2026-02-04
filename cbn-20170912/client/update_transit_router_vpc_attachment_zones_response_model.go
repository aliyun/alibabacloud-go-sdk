// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTransitRouterVpcAttachmentZonesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateTransitRouterVpcAttachmentZonesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateTransitRouterVpcAttachmentZonesResponse
	GetStatusCode() *int32
	SetBody(v *UpdateTransitRouterVpcAttachmentZonesResponseBody) *UpdateTransitRouterVpcAttachmentZonesResponse
	GetBody() *UpdateTransitRouterVpcAttachmentZonesResponseBody
}

type UpdateTransitRouterVpcAttachmentZonesResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTransitRouterVpcAttachmentZonesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTransitRouterVpcAttachmentZonesResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateTransitRouterVpcAttachmentZonesResponse) GoString() string {
	return s.String()
}

func (s *UpdateTransitRouterVpcAttachmentZonesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateTransitRouterVpcAttachmentZonesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTransitRouterVpcAttachmentZonesResponse) GetBody() *UpdateTransitRouterVpcAttachmentZonesResponseBody {
	return s.Body
}

func (s *UpdateTransitRouterVpcAttachmentZonesResponse) SetHeaders(v map[string]*string) *UpdateTransitRouterVpcAttachmentZonesResponse {
	s.Headers = v
	return s
}

func (s *UpdateTransitRouterVpcAttachmentZonesResponse) SetStatusCode(v int32) *UpdateTransitRouterVpcAttachmentZonesResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTransitRouterVpcAttachmentZonesResponse) SetBody(v *UpdateTransitRouterVpcAttachmentZonesResponseBody) *UpdateTransitRouterVpcAttachmentZonesResponse {
	s.Body = v
	return s
}

func (s *UpdateTransitRouterVpcAttachmentZonesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
