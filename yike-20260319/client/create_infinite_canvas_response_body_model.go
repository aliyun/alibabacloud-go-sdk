// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInfiniteCanvasResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCanvasId(v string) *CreateInfiniteCanvasResponseBody
	GetCanvasId() *string
	SetRequestId(v string) *CreateInfiniteCanvasResponseBody
	GetRequestId() *string
}

type CreateInfiniteCanvasResponseBody struct {
	// The canvas ID.
	//
	// example:
	//
	// canvas_***
	CanvasId *string `json:"CanvasId,omitempty" xml:"CanvasId,omitempty"`
	// RequestId
	//
	// example:
	//
	// ******3B-0E1A-586A-AC29-742247******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateInfiniteCanvasResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateInfiniteCanvasResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInfiniteCanvasResponseBody) GetCanvasId() *string {
	return s.CanvasId
}

func (s *CreateInfiniteCanvasResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateInfiniteCanvasResponseBody) SetCanvasId(v string) *CreateInfiniteCanvasResponseBody {
	s.CanvasId = &v
	return s
}

func (s *CreateInfiniteCanvasResponseBody) SetRequestId(v string) *CreateInfiniteCanvasResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateInfiniteCanvasResponseBody) Validate() error {
	return dara.Validate(s)
}
