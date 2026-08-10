// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInfiniteCanvasResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInfiniteCanvas(v *GetInfiniteCanvasResponseBodyInfiniteCanvas) *GetInfiniteCanvasResponseBody
	GetInfiniteCanvas() *GetInfiniteCanvasResponseBodyInfiniteCanvas
	SetRequestId(v string) *GetInfiniteCanvasResponseBody
	GetRequestId() *string
}

type GetInfiniteCanvasResponseBody struct {
	// The infinite canvas details.
	InfiniteCanvas *GetInfiniteCanvasResponseBodyInfiniteCanvas `json:"InfiniteCanvas,omitempty" xml:"InfiniteCanvas,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetInfiniteCanvasResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInfiniteCanvasResponseBody) GoString() string {
	return s.String()
}

func (s *GetInfiniteCanvasResponseBody) GetInfiniteCanvas() *GetInfiniteCanvasResponseBodyInfiniteCanvas {
	return s.InfiniteCanvas
}

func (s *GetInfiniteCanvasResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInfiniteCanvasResponseBody) SetInfiniteCanvas(v *GetInfiniteCanvasResponseBodyInfiniteCanvas) *GetInfiniteCanvasResponseBody {
	s.InfiniteCanvas = v
	return s
}

func (s *GetInfiniteCanvasResponseBody) SetRequestId(v string) *GetInfiniteCanvasResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInfiniteCanvasResponseBody) Validate() error {
	if s.InfiniteCanvas != nil {
		if err := s.InfiniteCanvas.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetInfiniteCanvasResponseBodyInfiniteCanvas struct {
	// The ID of the infinite canvas.
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
	// The creation time in UTC.
	//
	// example:
	//
	// 2026-07-01T08:42:16Z
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The modification time.
	//
	// example:
	//
	// 2026-07-01T08:42:16Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The thumbnail URL.
	//
	// example:
	//
	// https://*uncs.com/cover.png
	Thumbnail *string `json:"Thumbnail,omitempty" xml:"Thumbnail,omitempty"`
	// The title of the infinite canvas.
	//
	// example:
	//
	// test infinite canvas
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetInfiniteCanvasResponseBodyInfiniteCanvas) String() string {
	return dara.Prettify(s)
}

func (s GetInfiniteCanvasResponseBodyInfiniteCanvas) GoString() string {
	return s.String()
}

func (s *GetInfiniteCanvasResponseBodyInfiniteCanvas) GetCanvasId() *string {
	return s.CanvasId
}

func (s *GetInfiniteCanvasResponseBodyInfiniteCanvas) GetCoverUrl() *string {
	return s.CoverUrl
}

func (s *GetInfiniteCanvasResponseBodyInfiniteCanvas) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetInfiniteCanvasResponseBodyInfiniteCanvas) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetInfiniteCanvasResponseBodyInfiniteCanvas) GetThumbnail() *string {
	return s.Thumbnail
}

func (s *GetInfiniteCanvasResponseBodyInfiniteCanvas) GetTitle() *string {
	return s.Title
}

func (s *GetInfiniteCanvasResponseBodyInfiniteCanvas) SetCanvasId(v string) *GetInfiniteCanvasResponseBodyInfiniteCanvas {
	s.CanvasId = &v
	return s
}

func (s *GetInfiniteCanvasResponseBodyInfiniteCanvas) SetCoverUrl(v string) *GetInfiniteCanvasResponseBodyInfiniteCanvas {
	s.CoverUrl = &v
	return s
}

func (s *GetInfiniteCanvasResponseBodyInfiniteCanvas) SetGmtCreate(v string) *GetInfiniteCanvasResponseBodyInfiniteCanvas {
	s.GmtCreate = &v
	return s
}

func (s *GetInfiniteCanvasResponseBodyInfiniteCanvas) SetGmtModified(v string) *GetInfiniteCanvasResponseBodyInfiniteCanvas {
	s.GmtModified = &v
	return s
}

func (s *GetInfiniteCanvasResponseBodyInfiniteCanvas) SetThumbnail(v string) *GetInfiniteCanvasResponseBodyInfiniteCanvas {
	s.Thumbnail = &v
	return s
}

func (s *GetInfiniteCanvasResponseBodyInfiniteCanvas) SetTitle(v string) *GetInfiniteCanvasResponseBodyInfiniteCanvas {
	s.Title = &v
	return s
}

func (s *GetInfiniteCanvasResponseBodyInfiniteCanvas) Validate() error {
	return dara.Validate(s)
}
