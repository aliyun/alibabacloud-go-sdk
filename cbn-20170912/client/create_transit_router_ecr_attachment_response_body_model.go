// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTransitRouterEcrAttachmentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateTransitRouterEcrAttachmentResponseBody
	GetRequestId() *string
	SetTransitRouterAttachmentId(v string) *CreateTransitRouterEcrAttachmentResponseBody
	GetTransitRouterAttachmentId() *string
}

type CreateTransitRouterEcrAttachmentResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// C0245BEF-52AC-44A8-A776-EF96FD26A5CA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the ECR connection.
	//
	// example:
	//
	// tr-attach-qieks13jnt1cchy***
	TransitRouterAttachmentId *string `json:"TransitRouterAttachmentId,omitempty" xml:"TransitRouterAttachmentId,omitempty"`
}

func (s CreateTransitRouterEcrAttachmentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTransitRouterEcrAttachmentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTransitRouterEcrAttachmentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTransitRouterEcrAttachmentResponseBody) GetTransitRouterAttachmentId() *string {
	return s.TransitRouterAttachmentId
}

func (s *CreateTransitRouterEcrAttachmentResponseBody) SetRequestId(v string) *CreateTransitRouterEcrAttachmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTransitRouterEcrAttachmentResponseBody) SetTransitRouterAttachmentId(v string) *CreateTransitRouterEcrAttachmentResponseBody {
	s.TransitRouterAttachmentId = &v
	return s
}

func (s *CreateTransitRouterEcrAttachmentResponseBody) Validate() error {
	return dara.Validate(s)
}
