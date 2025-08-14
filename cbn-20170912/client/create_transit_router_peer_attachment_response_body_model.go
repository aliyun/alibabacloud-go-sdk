// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTransitRouterPeerAttachmentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateTransitRouterPeerAttachmentResponseBody
	GetRequestId() *string
	SetTransitRouterAttachmentId(v string) *CreateTransitRouterPeerAttachmentResponseBody
	GetTransitRouterAttachmentId() *string
}

type CreateTransitRouterPeerAttachmentResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 59422BF5-BAAD-4CFD-9019-9557BD3ACFA3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the inter-region connection.
	//
	// example:
	//
	// tr-attach-nwkiqfvw22qesz****
	TransitRouterAttachmentId *string `json:"TransitRouterAttachmentId,omitempty" xml:"TransitRouterAttachmentId,omitempty"`
}

func (s CreateTransitRouterPeerAttachmentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTransitRouterPeerAttachmentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTransitRouterPeerAttachmentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTransitRouterPeerAttachmentResponseBody) GetTransitRouterAttachmentId() *string {
	return s.TransitRouterAttachmentId
}

func (s *CreateTransitRouterPeerAttachmentResponseBody) SetRequestId(v string) *CreateTransitRouterPeerAttachmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTransitRouterPeerAttachmentResponseBody) SetTransitRouterAttachmentId(v string) *CreateTransitRouterPeerAttachmentResponseBody {
	s.TransitRouterAttachmentId = &v
	return s
}

func (s *CreateTransitRouterPeerAttachmentResponseBody) Validate() error {
	return dara.Validate(s)
}
