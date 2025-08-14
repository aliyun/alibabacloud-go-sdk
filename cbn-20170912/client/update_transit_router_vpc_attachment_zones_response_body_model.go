// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTransitRouterVpcAttachmentZonesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateTransitRouterVpcAttachmentZonesResponseBody
	GetRequestId() *string
}

type UpdateTransitRouterVpcAttachmentZonesResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 595AE52F-49FF-5788-A677-0DD1467941A1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateTransitRouterVpcAttachmentZonesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTransitRouterVpcAttachmentZonesResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTransitRouterVpcAttachmentZonesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTransitRouterVpcAttachmentZonesResponseBody) SetRequestId(v string) *UpdateTransitRouterVpcAttachmentZonesResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTransitRouterVpcAttachmentZonesResponseBody) Validate() error {
	return dara.Validate(s)
}
