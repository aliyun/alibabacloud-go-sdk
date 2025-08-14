// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateTransitRouterAttachmentWithRouteTableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AssociateTransitRouterAttachmentWithRouteTableResponseBody
	GetRequestId() *string
}

type AssociateTransitRouterAttachmentWithRouteTableResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 4EC47282-1B74-4534-BD0E-403F3EE64CAF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AssociateTransitRouterAttachmentWithRouteTableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AssociateTransitRouterAttachmentWithRouteTableResponseBody) GoString() string {
	return s.String()
}

func (s *AssociateTransitRouterAttachmentWithRouteTableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AssociateTransitRouterAttachmentWithRouteTableResponseBody) SetRequestId(v string) *AssociateTransitRouterAttachmentWithRouteTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssociateTransitRouterAttachmentWithRouteTableResponseBody) Validate() error {
	return dara.Validate(s)
}
