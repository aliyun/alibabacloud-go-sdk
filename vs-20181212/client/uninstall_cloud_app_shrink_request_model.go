// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUninstallCloudAppShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *UninstallCloudAppShrinkRequest
	GetAppId() *string
	SetPageNumber(v int32) *UninstallCloudAppShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *UninstallCloudAppShrinkRequest
	GetPageSize() *int32
	SetPatchId(v string) *UninstallCloudAppShrinkRequest
	GetPatchId() *string
	SetProjectId(v string) *UninstallCloudAppShrinkRequest
	GetProjectId() *string
	SetRenderingInstanceId(v string) *UninstallCloudAppShrinkRequest
	GetRenderingInstanceId() *string
	SetRenderingInstanceIdsShrink(v string) *UninstallCloudAppShrinkRequest
	GetRenderingInstanceIdsShrink() *string
}

type UninstallCloudAppShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cap-b06b26edfhytbn b94a75ae1a79efc90eb
	AppId      *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// patch-03fa76e8e13a49b6a966b063d9d309b4
	PatchId   *string `json:"PatchId,omitempty" xml:"PatchId,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// render-9f8c57355d224ad7beaf95e145f22111
	RenderingInstanceId        *string `json:"RenderingInstanceId,omitempty" xml:"RenderingInstanceId,omitempty"`
	RenderingInstanceIdsShrink *string `json:"RenderingInstanceIds,omitempty" xml:"RenderingInstanceIds,omitempty"`
}

func (s UninstallCloudAppShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UninstallCloudAppShrinkRequest) GoString() string {
	return s.String()
}

func (s *UninstallCloudAppShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *UninstallCloudAppShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *UninstallCloudAppShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *UninstallCloudAppShrinkRequest) GetPatchId() *string {
	return s.PatchId
}

func (s *UninstallCloudAppShrinkRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *UninstallCloudAppShrinkRequest) GetRenderingInstanceId() *string {
	return s.RenderingInstanceId
}

func (s *UninstallCloudAppShrinkRequest) GetRenderingInstanceIdsShrink() *string {
	return s.RenderingInstanceIdsShrink
}

func (s *UninstallCloudAppShrinkRequest) SetAppId(v string) *UninstallCloudAppShrinkRequest {
	s.AppId = &v
	return s
}

func (s *UninstallCloudAppShrinkRequest) SetPageNumber(v int32) *UninstallCloudAppShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *UninstallCloudAppShrinkRequest) SetPageSize(v int32) *UninstallCloudAppShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *UninstallCloudAppShrinkRequest) SetPatchId(v string) *UninstallCloudAppShrinkRequest {
	s.PatchId = &v
	return s
}

func (s *UninstallCloudAppShrinkRequest) SetProjectId(v string) *UninstallCloudAppShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *UninstallCloudAppShrinkRequest) SetRenderingInstanceId(v string) *UninstallCloudAppShrinkRequest {
	s.RenderingInstanceId = &v
	return s
}

func (s *UninstallCloudAppShrinkRequest) SetRenderingInstanceIdsShrink(v string) *UninstallCloudAppShrinkRequest {
	s.RenderingInstanceIdsShrink = &v
	return s
}

func (s *UninstallCloudAppShrinkRequest) Validate() error {
	return dara.Validate(s)
}
