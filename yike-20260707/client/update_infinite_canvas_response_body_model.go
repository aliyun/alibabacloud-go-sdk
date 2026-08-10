// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInfiniteCanvasResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCanvasId(v string) *UpdateInfiniteCanvasResponseBody
	GetCanvasId() *string
	SetRequestId(v string) *UpdateInfiniteCanvasResponseBody
	GetRequestId() *string
}

type UpdateInfiniteCanvasResponseBody struct {
	// The ID of the infinite canvas.
	//
	// example:
	//
	// canvas_gesad*
	CanvasId *string `json:"CanvasId,omitempty" xml:"CanvasId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ****63E8B7C7-4812-46AD-0FA56029AC86****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateInfiniteCanvasResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateInfiniteCanvasResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInfiniteCanvasResponseBody) GetCanvasId() *string {
	return s.CanvasId
}

func (s *UpdateInfiniteCanvasResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateInfiniteCanvasResponseBody) SetCanvasId(v string) *UpdateInfiniteCanvasResponseBody {
	s.CanvasId = &v
	return s
}

func (s *UpdateInfiniteCanvasResponseBody) SetRequestId(v string) *UpdateInfiniteCanvasResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateInfiniteCanvasResponseBody) Validate() error {
	return dara.Validate(s)
}
