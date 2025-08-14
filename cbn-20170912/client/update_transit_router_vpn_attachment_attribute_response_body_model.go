// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTransitRouterVpnAttachmentAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateTransitRouterVpnAttachmentAttributeResponseBody
	GetRequestId() *string
}

type UpdateTransitRouterVpnAttachmentAttributeResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 1AB038F2-96E5-560B-9F6E-734311D466FA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateTransitRouterVpnAttachmentAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTransitRouterVpnAttachmentAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTransitRouterVpnAttachmentAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTransitRouterVpnAttachmentAttributeResponseBody) SetRequestId(v string) *UpdateTransitRouterVpnAttachmentAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTransitRouterVpnAttachmentAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
