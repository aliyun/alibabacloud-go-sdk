// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDissociateTransitRouterAttachmentFromRouteTableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DissociateTransitRouterAttachmentFromRouteTableResponseBody
	GetRequestId() *string
}

type DissociateTransitRouterAttachmentFromRouteTableResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 4EC47282-1B74-4534-BD0E-403F3EE64CAF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DissociateTransitRouterAttachmentFromRouteTableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DissociateTransitRouterAttachmentFromRouteTableResponseBody) GoString() string {
	return s.String()
}

func (s *DissociateTransitRouterAttachmentFromRouteTableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DissociateTransitRouterAttachmentFromRouteTableResponseBody) SetRequestId(v string) *DissociateTransitRouterAttachmentFromRouteTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *DissociateTransitRouterAttachmentFromRouteTableResponseBody) Validate() error {
	return dara.Validate(s)
}
