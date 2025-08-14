// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTransitRouterVbrAttachmentAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateTransitRouterVbrAttachmentAttributeResponseBody
	GetRequestId() *string
}

type UpdateTransitRouterVbrAttachmentAttributeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 91B36DC3-FF8A-45C3-AC1E-456B1789136D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateTransitRouterVbrAttachmentAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTransitRouterVbrAttachmentAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTransitRouterVbrAttachmentAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTransitRouterVbrAttachmentAttributeResponseBody) SetRequestId(v string) *UpdateTransitRouterVbrAttachmentAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTransitRouterVbrAttachmentAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
