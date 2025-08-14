// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTransitRouterPeerAttachmentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteTransitRouterPeerAttachmentResponseBody
	GetRequestId() *string
}

type DeleteTransitRouterPeerAttachmentResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// A01FEDD7-7D69-4EB3-996D-CF79F6F885CA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTransitRouterPeerAttachmentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTransitRouterPeerAttachmentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTransitRouterPeerAttachmentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTransitRouterPeerAttachmentResponseBody) SetRequestId(v string) *DeleteTransitRouterPeerAttachmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTransitRouterPeerAttachmentResponseBody) Validate() error {
	return dara.Validate(s)
}
