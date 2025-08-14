// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTransitRouterVbrAttachmentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateTransitRouterVbrAttachmentResponseBody
	GetRequestId() *string
	SetTransitRouterAttachmentId(v string) *CreateTransitRouterVbrAttachmentResponseBody
	GetTransitRouterAttachmentId() *string
}

type CreateTransitRouterVbrAttachmentResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// C087A369-82B9-43EF-91F4-4B63A9C6E6B6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the VBR connection.
	//
	// example:
	//
	// tr-attach-ia340z7xis7t5s****
	TransitRouterAttachmentId *string `json:"TransitRouterAttachmentId,omitempty" xml:"TransitRouterAttachmentId,omitempty"`
}

func (s CreateTransitRouterVbrAttachmentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTransitRouterVbrAttachmentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTransitRouterVbrAttachmentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTransitRouterVbrAttachmentResponseBody) GetTransitRouterAttachmentId() *string {
	return s.TransitRouterAttachmentId
}

func (s *CreateTransitRouterVbrAttachmentResponseBody) SetRequestId(v string) *CreateTransitRouterVbrAttachmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTransitRouterVbrAttachmentResponseBody) SetTransitRouterAttachmentId(v string) *CreateTransitRouterVbrAttachmentResponseBody {
	s.TransitRouterAttachmentId = &v
	return s
}

func (s *CreateTransitRouterVbrAttachmentResponseBody) Validate() error {
	return dara.Validate(s)
}
