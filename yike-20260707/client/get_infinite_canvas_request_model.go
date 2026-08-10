// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInfiniteCanvasRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCanvasId(v string) *GetInfiniteCanvasRequest
	GetCanvasId() *string
}

type GetInfiniteCanvasRequest struct {
	// The ID of the infinite canvas.
	//
	// This parameter is required.
	//
	// example:
	//
	// canvas_***
	CanvasId *string `json:"CanvasId,omitempty" xml:"CanvasId,omitempty"`
}

func (s GetInfiniteCanvasRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInfiniteCanvasRequest) GoString() string {
	return s.String()
}

func (s *GetInfiniteCanvasRequest) GetCanvasId() *string {
	return s.CanvasId
}

func (s *GetInfiniteCanvasRequest) SetCanvasId(v string) *GetInfiniteCanvasRequest {
	s.CanvasId = &v
	return s
}

func (s *GetInfiniteCanvasRequest) Validate() error {
	return dara.Validate(s)
}
