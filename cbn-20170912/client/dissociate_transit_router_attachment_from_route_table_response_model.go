// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDissociateTransitRouterAttachmentFromRouteTableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DissociateTransitRouterAttachmentFromRouteTableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DissociateTransitRouterAttachmentFromRouteTableResponse
	GetStatusCode() *int32
	SetBody(v *DissociateTransitRouterAttachmentFromRouteTableResponseBody) *DissociateTransitRouterAttachmentFromRouteTableResponse
	GetBody() *DissociateTransitRouterAttachmentFromRouteTableResponseBody
}

type DissociateTransitRouterAttachmentFromRouteTableResponse struct {
	Headers    map[string]*string                                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DissociateTransitRouterAttachmentFromRouteTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DissociateTransitRouterAttachmentFromRouteTableResponse) String() string {
	return dara.Prettify(s)
}

func (s DissociateTransitRouterAttachmentFromRouteTableResponse) GoString() string {
	return s.String()
}

func (s *DissociateTransitRouterAttachmentFromRouteTableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DissociateTransitRouterAttachmentFromRouteTableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DissociateTransitRouterAttachmentFromRouteTableResponse) GetBody() *DissociateTransitRouterAttachmentFromRouteTableResponseBody {
	return s.Body
}

func (s *DissociateTransitRouterAttachmentFromRouteTableResponse) SetHeaders(v map[string]*string) *DissociateTransitRouterAttachmentFromRouteTableResponse {
	s.Headers = v
	return s
}

func (s *DissociateTransitRouterAttachmentFromRouteTableResponse) SetStatusCode(v int32) *DissociateTransitRouterAttachmentFromRouteTableResponse {
	s.StatusCode = &v
	return s
}

func (s *DissociateTransitRouterAttachmentFromRouteTableResponse) SetBody(v *DissociateTransitRouterAttachmentFromRouteTableResponseBody) *DissociateTransitRouterAttachmentFromRouteTableResponse {
	s.Body = v
	return s
}

func (s *DissociateTransitRouterAttachmentFromRouteTableResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
