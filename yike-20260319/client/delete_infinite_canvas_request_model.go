// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInfiniteCanvasRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCanvasId(v string) *DeleteInfiniteCanvasRequest
	GetCanvasId() *string
}

type DeleteInfiniteCanvasRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// canvas_adaasd*
	CanvasId *string `json:"CanvasId,omitempty" xml:"CanvasId,omitempty"`
}

func (s DeleteInfiniteCanvasRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteInfiniteCanvasRequest) GoString() string {
	return s.String()
}

func (s *DeleteInfiniteCanvasRequest) GetCanvasId() *string {
	return s.CanvasId
}

func (s *DeleteInfiniteCanvasRequest) SetCanvasId(v string) *DeleteInfiniteCanvasRequest {
	s.CanvasId = &v
	return s
}

func (s *DeleteInfiniteCanvasRequest) Validate() error {
	return dara.Validate(s)
}
