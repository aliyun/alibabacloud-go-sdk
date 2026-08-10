// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInfiniteCanvasRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCanvasId(v string) *UpdateInfiniteCanvasRequest
	GetCanvasId() *string
	SetCoverUrl(v string) *UpdateInfiniteCanvasRequest
	GetCoverUrl() *string
	SetTitle(v string) *UpdateInfiniteCanvasRequest
	GetTitle() *string
}

type UpdateInfiniteCanvasRequest struct {
	// The ID of the infinite canvas.
	//
	// This parameter is required.
	//
	// example:
	//
	// canvas_gesad*
	CanvasId *string `json:"CanvasId,omitempty" xml:"CanvasId,omitempty"`
	// The cover URL.
	//
	// example:
	//
	// https://*uncs.com/cover.png
	CoverUrl *string `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	// The title of the infinite canvas.
	//
	// example:
	//
	// example
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s UpdateInfiniteCanvasRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateInfiniteCanvasRequest) GoString() string {
	return s.String()
}

func (s *UpdateInfiniteCanvasRequest) GetCanvasId() *string {
	return s.CanvasId
}

func (s *UpdateInfiniteCanvasRequest) GetCoverUrl() *string {
	return s.CoverUrl
}

func (s *UpdateInfiniteCanvasRequest) GetTitle() *string {
	return s.Title
}

func (s *UpdateInfiniteCanvasRequest) SetCanvasId(v string) *UpdateInfiniteCanvasRequest {
	s.CanvasId = &v
	return s
}

func (s *UpdateInfiniteCanvasRequest) SetCoverUrl(v string) *UpdateInfiniteCanvasRequest {
	s.CoverUrl = &v
	return s
}

func (s *UpdateInfiniteCanvasRequest) SetTitle(v string) *UpdateInfiniteCanvasRequest {
	s.Title = &v
	return s
}

func (s *UpdateInfiniteCanvasRequest) Validate() error {
	return dara.Validate(s)
}
