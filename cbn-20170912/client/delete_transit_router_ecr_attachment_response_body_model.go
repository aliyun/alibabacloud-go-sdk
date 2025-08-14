// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTransitRouterEcrAttachmentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteTransitRouterEcrAttachmentResponseBody
	GetRequestId() *string
}

type DeleteTransitRouterEcrAttachmentResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F7DDDC17-FA06-4AC2-8F35-59D2470FCFC1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTransitRouterEcrAttachmentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTransitRouterEcrAttachmentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTransitRouterEcrAttachmentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTransitRouterEcrAttachmentResponseBody) SetRequestId(v string) *DeleteTransitRouterEcrAttachmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTransitRouterEcrAttachmentResponseBody) Validate() error {
	return dara.Validate(s)
}
