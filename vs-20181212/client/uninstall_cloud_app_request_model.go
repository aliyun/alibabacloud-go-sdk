// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUninstallCloudAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *UninstallCloudAppRequest
	GetAppId() *string
	SetPageNumber(v int32) *UninstallCloudAppRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *UninstallCloudAppRequest
	GetPageSize() *int32
	SetPatchId(v string) *UninstallCloudAppRequest
	GetPatchId() *string
	SetProjectId(v string) *UninstallCloudAppRequest
	GetProjectId() *string
	SetRenderingInstanceId(v string) *UninstallCloudAppRequest
	GetRenderingInstanceId() *string
	SetRenderingInstanceIds(v []*string) *UninstallCloudAppRequest
	GetRenderingInstanceIds() []*string
}

type UninstallCloudAppRequest struct {
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
	RenderingInstanceId  *string   `json:"RenderingInstanceId,omitempty" xml:"RenderingInstanceId,omitempty"`
	RenderingInstanceIds []*string `json:"RenderingInstanceIds,omitempty" xml:"RenderingInstanceIds,omitempty" type:"Repeated"`
}

func (s UninstallCloudAppRequest) String() string {
	return dara.Prettify(s)
}

func (s UninstallCloudAppRequest) GoString() string {
	return s.String()
}

func (s *UninstallCloudAppRequest) GetAppId() *string {
	return s.AppId
}

func (s *UninstallCloudAppRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *UninstallCloudAppRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *UninstallCloudAppRequest) GetPatchId() *string {
	return s.PatchId
}

func (s *UninstallCloudAppRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *UninstallCloudAppRequest) GetRenderingInstanceId() *string {
	return s.RenderingInstanceId
}

func (s *UninstallCloudAppRequest) GetRenderingInstanceIds() []*string {
	return s.RenderingInstanceIds
}

func (s *UninstallCloudAppRequest) SetAppId(v string) *UninstallCloudAppRequest {
	s.AppId = &v
	return s
}

func (s *UninstallCloudAppRequest) SetPageNumber(v int32) *UninstallCloudAppRequest {
	s.PageNumber = &v
	return s
}

func (s *UninstallCloudAppRequest) SetPageSize(v int32) *UninstallCloudAppRequest {
	s.PageSize = &v
	return s
}

func (s *UninstallCloudAppRequest) SetPatchId(v string) *UninstallCloudAppRequest {
	s.PatchId = &v
	return s
}

func (s *UninstallCloudAppRequest) SetProjectId(v string) *UninstallCloudAppRequest {
	s.ProjectId = &v
	return s
}

func (s *UninstallCloudAppRequest) SetRenderingInstanceId(v string) *UninstallCloudAppRequest {
	s.RenderingInstanceId = &v
	return s
}

func (s *UninstallCloudAppRequest) SetRenderingInstanceIds(v []*string) *UninstallCloudAppRequest {
	s.RenderingInstanceIds = v
	return s
}

func (s *UninstallCloudAppRequest) Validate() error {
	return dara.Validate(s)
}
