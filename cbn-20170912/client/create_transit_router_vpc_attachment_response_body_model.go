// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTransitRouterVpcAttachmentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateTransitRouterVpcAttachmentResponseBody
	GetRequestId() *string
	SetTransitRouterAttachmentId(v string) *CreateTransitRouterVpcAttachmentResponseBody
	GetTransitRouterAttachmentId() *string
}

type CreateTransitRouterVpcAttachmentResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// C087A369-82B9-43EF-91F4-4B63A9C6E6B6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the VPC connection.
	//
	// example:
	//
	// tr-attach-ia340z7xis7t5s****
	TransitRouterAttachmentId *string `json:"TransitRouterAttachmentId,omitempty" xml:"TransitRouterAttachmentId,omitempty"`
}

func (s CreateTransitRouterVpcAttachmentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTransitRouterVpcAttachmentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTransitRouterVpcAttachmentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTransitRouterVpcAttachmentResponseBody) GetTransitRouterAttachmentId() *string {
	return s.TransitRouterAttachmentId
}

func (s *CreateTransitRouterVpcAttachmentResponseBody) SetRequestId(v string) *CreateTransitRouterVpcAttachmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTransitRouterVpcAttachmentResponseBody) SetTransitRouterAttachmentId(v string) *CreateTransitRouterVpcAttachmentResponseBody {
	s.TransitRouterAttachmentId = &v
	return s
}

func (s *CreateTransitRouterVpcAttachmentResponseBody) Validate() error {
	return dara.Validate(s)
}
