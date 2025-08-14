// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTransitRouterVpcAttachmentAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateTransitRouterVpcAttachmentAttributeResponseBody
	GetRequestId() *string
}

type UpdateTransitRouterVpcAttachmentAttributeResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 7E61D334-4025-41EF-9145-FC327B35301D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateTransitRouterVpcAttachmentAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTransitRouterVpcAttachmentAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTransitRouterVpcAttachmentAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTransitRouterVpcAttachmentAttributeResponseBody) SetRequestId(v string) *UpdateTransitRouterVpcAttachmentAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTransitRouterVpcAttachmentAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
