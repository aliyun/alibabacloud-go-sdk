// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTransitRouterPeerAttachmentAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateTransitRouterPeerAttachmentAttributeResponseBody
	GetRequestId() *string
}

type UpdateTransitRouterPeerAttachmentAttributeResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// A01FEDD7-7D69-4EB3-996D-CF79F6F885CA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateTransitRouterPeerAttachmentAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTransitRouterPeerAttachmentAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTransitRouterPeerAttachmentAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTransitRouterPeerAttachmentAttributeResponseBody) SetRequestId(v string) *UpdateTransitRouterPeerAttachmentAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTransitRouterPeerAttachmentAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
