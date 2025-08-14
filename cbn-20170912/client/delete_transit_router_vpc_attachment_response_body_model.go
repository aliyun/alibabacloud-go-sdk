// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTransitRouterVpcAttachmentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteTransitRouterVpcAttachmentResponseBody
	GetRequestId() *string
}

type DeleteTransitRouterVpcAttachmentResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 7E61D334-4025-41EF-9145-FC327B35301D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTransitRouterVpcAttachmentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTransitRouterVpcAttachmentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTransitRouterVpcAttachmentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTransitRouterVpcAttachmentResponseBody) SetRequestId(v string) *DeleteTransitRouterVpcAttachmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTransitRouterVpcAttachmentResponseBody) Validate() error {
	return dara.Validate(s)
}
