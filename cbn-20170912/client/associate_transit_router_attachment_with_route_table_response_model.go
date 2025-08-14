// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateTransitRouterAttachmentWithRouteTableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AssociateTransitRouterAttachmentWithRouteTableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AssociateTransitRouterAttachmentWithRouteTableResponse
	GetStatusCode() *int32
	SetBody(v *AssociateTransitRouterAttachmentWithRouteTableResponseBody) *AssociateTransitRouterAttachmentWithRouteTableResponse
	GetBody() *AssociateTransitRouterAttachmentWithRouteTableResponseBody
}

type AssociateTransitRouterAttachmentWithRouteTableResponse struct {
	Headers    map[string]*string                                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssociateTransitRouterAttachmentWithRouteTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssociateTransitRouterAttachmentWithRouteTableResponse) String() string {
	return dara.Prettify(s)
}

func (s AssociateTransitRouterAttachmentWithRouteTableResponse) GoString() string {
	return s.String()
}

func (s *AssociateTransitRouterAttachmentWithRouteTableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AssociateTransitRouterAttachmentWithRouteTableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AssociateTransitRouterAttachmentWithRouteTableResponse) GetBody() *AssociateTransitRouterAttachmentWithRouteTableResponseBody {
	return s.Body
}

func (s *AssociateTransitRouterAttachmentWithRouteTableResponse) SetHeaders(v map[string]*string) *AssociateTransitRouterAttachmentWithRouteTableResponse {
	s.Headers = v
	return s
}

func (s *AssociateTransitRouterAttachmentWithRouteTableResponse) SetStatusCode(v int32) *AssociateTransitRouterAttachmentWithRouteTableResponse {
	s.StatusCode = &v
	return s
}

func (s *AssociateTransitRouterAttachmentWithRouteTableResponse) SetBody(v *AssociateTransitRouterAttachmentWithRouteTableResponseBody) *AssociateTransitRouterAttachmentWithRouteTableResponse {
	s.Body = v
	return s
}

func (s *AssociateTransitRouterAttachmentWithRouteTableResponse) Validate() error {
	return dara.Validate(s)
}
