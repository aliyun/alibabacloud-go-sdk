// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInfiniteCanvasRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCoverUrl(v string) *CreateInfiniteCanvasRequest
	GetCoverUrl() *string
	SetProductionId(v string) *CreateInfiniteCanvasRequest
	GetProductionId() *string
	SetTitle(v string) *CreateInfiniteCanvasRequest
	GetTitle() *string
	SetWorkspaceId(v string) *CreateInfiniteCanvasRequest
	GetWorkspaceId() *string
}

type CreateInfiniteCanvasRequest struct {
	// The cover URL.
	//
	// example:
	//
	// https://tagvvcloud-material-center-prod.oss-cn-hangzhou.aliyuncs.com/sumvideo/utils_image/sumvideo-video-cover.png
	CoverUrl *string `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	// The project ID.
	//
	// example:
	//
	// pd_061716****
	ProductionId *string `json:"ProductionId,omitempty" xml:"ProductionId,omitempty"`
	// The project title.
	//
	// This parameter is required.
	//
	// example:
	//
	// title
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// ws_zna577pdximv***
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateInfiniteCanvasRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateInfiniteCanvasRequest) GoString() string {
	return s.String()
}

func (s *CreateInfiniteCanvasRequest) GetCoverUrl() *string {
	return s.CoverUrl
}

func (s *CreateInfiniteCanvasRequest) GetProductionId() *string {
	return s.ProductionId
}

func (s *CreateInfiniteCanvasRequest) GetTitle() *string {
	return s.Title
}

func (s *CreateInfiniteCanvasRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateInfiniteCanvasRequest) SetCoverUrl(v string) *CreateInfiniteCanvasRequest {
	s.CoverUrl = &v
	return s
}

func (s *CreateInfiniteCanvasRequest) SetProductionId(v string) *CreateInfiniteCanvasRequest {
	s.ProductionId = &v
	return s
}

func (s *CreateInfiniteCanvasRequest) SetTitle(v string) *CreateInfiniteCanvasRequest {
	s.Title = &v
	return s
}

func (s *CreateInfiniteCanvasRequest) SetWorkspaceId(v string) *CreateInfiniteCanvasRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateInfiniteCanvasRequest) Validate() error {
	return dara.Validate(s)
}
