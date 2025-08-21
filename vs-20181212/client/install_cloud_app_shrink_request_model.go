// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallCloudAppShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *InstallCloudAppShrinkRequest
	GetAppId() *string
	SetPageNumber(v int32) *InstallCloudAppShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *InstallCloudAppShrinkRequest
	GetPageSize() *int32
	SetPatchId(v string) *InstallCloudAppShrinkRequest
	GetPatchId() *string
	SetProjectId(v string) *InstallCloudAppShrinkRequest
	GetProjectId() *string
	SetRenderingInstanceId(v string) *InstallCloudAppShrinkRequest
	GetRenderingInstanceId() *string
	SetRenderingInstanceIdsShrink(v string) *InstallCloudAppShrinkRequest
	GetRenderingInstanceIdsShrink() *string
}

type InstallCloudAppShrinkRequest struct {
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
	// patch-7bdf679812484df08a956b73e0b3bdf6
	PatchId   *string `json:"PatchId,omitempty" xml:"PatchId,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// render-9f8c57355d224ad7beaf95e145f22111
	RenderingInstanceId        *string `json:"RenderingInstanceId,omitempty" xml:"RenderingInstanceId,omitempty"`
	RenderingInstanceIdsShrink *string `json:"RenderingInstanceIds,omitempty" xml:"RenderingInstanceIds,omitempty"`
}

func (s InstallCloudAppShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s InstallCloudAppShrinkRequest) GoString() string {
	return s.String()
}

func (s *InstallCloudAppShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *InstallCloudAppShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *InstallCloudAppShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *InstallCloudAppShrinkRequest) GetPatchId() *string {
	return s.PatchId
}

func (s *InstallCloudAppShrinkRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *InstallCloudAppShrinkRequest) GetRenderingInstanceId() *string {
	return s.RenderingInstanceId
}

func (s *InstallCloudAppShrinkRequest) GetRenderingInstanceIdsShrink() *string {
	return s.RenderingInstanceIdsShrink
}

func (s *InstallCloudAppShrinkRequest) SetAppId(v string) *InstallCloudAppShrinkRequest {
	s.AppId = &v
	return s
}

func (s *InstallCloudAppShrinkRequest) SetPageNumber(v int32) *InstallCloudAppShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *InstallCloudAppShrinkRequest) SetPageSize(v int32) *InstallCloudAppShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *InstallCloudAppShrinkRequest) SetPatchId(v string) *InstallCloudAppShrinkRequest {
	s.PatchId = &v
	return s
}

func (s *InstallCloudAppShrinkRequest) SetProjectId(v string) *InstallCloudAppShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *InstallCloudAppShrinkRequest) SetRenderingInstanceId(v string) *InstallCloudAppShrinkRequest {
	s.RenderingInstanceId = &v
	return s
}

func (s *InstallCloudAppShrinkRequest) SetRenderingInstanceIdsShrink(v string) *InstallCloudAppShrinkRequest {
	s.RenderingInstanceIdsShrink = &v
	return s
}

func (s *InstallCloudAppShrinkRequest) Validate() error {
	return dara.Validate(s)
}
