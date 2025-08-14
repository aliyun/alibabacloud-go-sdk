// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTransitRouterVpnAttachmentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteTransitRouterVpnAttachmentResponseBody
	GetRequestId() *string
}

type DeleteTransitRouterVpnAttachmentResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// FA43C571-E88B-56C0-8FF8-5646D9B96297
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTransitRouterVpnAttachmentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTransitRouterVpnAttachmentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTransitRouterVpnAttachmentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTransitRouterVpnAttachmentResponseBody) SetRequestId(v string) *DeleteTransitRouterVpnAttachmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTransitRouterVpnAttachmentResponseBody) Validate() error {
	return dara.Validate(s)
}
