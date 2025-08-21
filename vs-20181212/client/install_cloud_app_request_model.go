// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallCloudAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *InstallCloudAppRequest
	GetAppId() *string
	SetPageNumber(v int32) *InstallCloudAppRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *InstallCloudAppRequest
	GetPageSize() *int32
	SetPatchId(v string) *InstallCloudAppRequest
	GetPatchId() *string
	SetProjectId(v string) *InstallCloudAppRequest
	GetProjectId() *string
	SetRenderingInstanceId(v string) *InstallCloudAppRequest
	GetRenderingInstanceId() *string
	SetRenderingInstanceIds(v []*string) *InstallCloudAppRequest
	GetRenderingInstanceIds() []*string
}

type InstallCloudAppRequest struct {
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
	RenderingInstanceId  *string   `json:"RenderingInstanceId,omitempty" xml:"RenderingInstanceId,omitempty"`
	RenderingInstanceIds []*string `json:"RenderingInstanceIds,omitempty" xml:"RenderingInstanceIds,omitempty" type:"Repeated"`
}

func (s InstallCloudAppRequest) String() string {
	return dara.Prettify(s)
}

func (s InstallCloudAppRequest) GoString() string {
	return s.String()
}

func (s *InstallCloudAppRequest) GetAppId() *string {
	return s.AppId
}

func (s *InstallCloudAppRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *InstallCloudAppRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *InstallCloudAppRequest) GetPatchId() *string {
	return s.PatchId
}

func (s *InstallCloudAppRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *InstallCloudAppRequest) GetRenderingInstanceId() *string {
	return s.RenderingInstanceId
}

func (s *InstallCloudAppRequest) GetRenderingInstanceIds() []*string {
	return s.RenderingInstanceIds
}

func (s *InstallCloudAppRequest) SetAppId(v string) *InstallCloudAppRequest {
	s.AppId = &v
	return s
}

func (s *InstallCloudAppRequest) SetPageNumber(v int32) *InstallCloudAppRequest {
	s.PageNumber = &v
	return s
}

func (s *InstallCloudAppRequest) SetPageSize(v int32) *InstallCloudAppRequest {
	s.PageSize = &v
	return s
}

func (s *InstallCloudAppRequest) SetPatchId(v string) *InstallCloudAppRequest {
	s.PatchId = &v
	return s
}

func (s *InstallCloudAppRequest) SetProjectId(v string) *InstallCloudAppRequest {
	s.ProjectId = &v
	return s
}

func (s *InstallCloudAppRequest) SetRenderingInstanceId(v string) *InstallCloudAppRequest {
	s.RenderingInstanceId = &v
	return s
}

func (s *InstallCloudAppRequest) SetRenderingInstanceIds(v []*string) *InstallCloudAppRequest {
	s.RenderingInstanceIds = v
	return s
}

func (s *InstallCloudAppRequest) Validate() error {
	return dara.Validate(s)
}
