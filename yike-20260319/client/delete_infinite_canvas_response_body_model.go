// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInfiniteCanvasResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCanvasId(v string) *DeleteInfiniteCanvasResponseBody
	GetCanvasId() *string
	SetRequestId(v string) *DeleteInfiniteCanvasResponseBody
	GetRequestId() *string
}

type DeleteInfiniteCanvasResponseBody struct {
	// example:
	//
	// canvas_adaasd*
	CanvasId *string `json:"CanvasId,omitempty" xml:"CanvasId,omitempty"`
	// example:
	//
	// ***F88A3-AC51-5588-859A-03144F082***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteInfiniteCanvasResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteInfiniteCanvasResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInfiniteCanvasResponseBody) GetCanvasId() *string {
	return s.CanvasId
}

func (s *DeleteInfiniteCanvasResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteInfiniteCanvasResponseBody) SetCanvasId(v string) *DeleteInfiniteCanvasResponseBody {
	s.CanvasId = &v
	return s
}

func (s *DeleteInfiniteCanvasResponseBody) SetRequestId(v string) *DeleteInfiniteCanvasResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteInfiniteCanvasResponseBody) Validate() error {
	return dara.Validate(s)
}
