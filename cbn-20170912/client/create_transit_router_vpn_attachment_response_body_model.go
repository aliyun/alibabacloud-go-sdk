// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTransitRouterVpnAttachmentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateTransitRouterVpnAttachmentResponseBody
	GetRequestId() *string
	SetTransitRouterAttachmentId(v string) *CreateTransitRouterVpnAttachmentResponseBody
	GetTransitRouterAttachmentId() *string
}

type CreateTransitRouterVpnAttachmentResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 8B5DD40A-3A29-5AC0-B8DA-05FD10D5C893
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the VPN attachment.
	//
	// example:
	//
	// tr-attach-y5dup2qwfyh9lu****
	TransitRouterAttachmentId *string `json:"TransitRouterAttachmentId,omitempty" xml:"TransitRouterAttachmentId,omitempty"`
}

func (s CreateTransitRouterVpnAttachmentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTransitRouterVpnAttachmentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTransitRouterVpnAttachmentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTransitRouterVpnAttachmentResponseBody) GetTransitRouterAttachmentId() *string {
	return s.TransitRouterAttachmentId
}

func (s *CreateTransitRouterVpnAttachmentResponseBody) SetRequestId(v string) *CreateTransitRouterVpnAttachmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTransitRouterVpnAttachmentResponseBody) SetTransitRouterAttachmentId(v string) *CreateTransitRouterVpnAttachmentResponseBody {
	s.TransitRouterAttachmentId = &v
	return s
}

func (s *CreateTransitRouterVpnAttachmentResponseBody) Validate() error {
	return dara.Validate(s)
}
