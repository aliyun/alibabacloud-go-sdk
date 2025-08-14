// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTransitRouterVbrAttachmentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteTransitRouterVbrAttachmentResponseBody
	GetRequestId() *string
}

type DeleteTransitRouterVbrAttachmentResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 3EDA94DE-0AE5-41FC-A91E-7170E408E0FD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTransitRouterVbrAttachmentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTransitRouterVbrAttachmentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTransitRouterVbrAttachmentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTransitRouterVbrAttachmentResponseBody) SetRequestId(v string) *DeleteTransitRouterVbrAttachmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTransitRouterVbrAttachmentResponseBody) Validate() error {
	return dara.Validate(s)
}
