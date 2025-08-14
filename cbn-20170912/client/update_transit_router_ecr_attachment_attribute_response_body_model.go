// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTransitRouterEcrAttachmentAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateTransitRouterEcrAttachmentAttributeResponseBody
	GetRequestId() *string
}

type UpdateTransitRouterEcrAttachmentAttributeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 54B48E3D-DF70-471B-AA93-08E683A1B457
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateTransitRouterEcrAttachmentAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTransitRouterEcrAttachmentAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTransitRouterEcrAttachmentAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTransitRouterEcrAttachmentAttributeResponseBody) SetRequestId(v string) *UpdateTransitRouterEcrAttachmentAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTransitRouterEcrAttachmentAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
