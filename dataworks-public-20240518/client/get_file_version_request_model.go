// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFileVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileId(v int64) *GetFileVersionRequest
	GetFileId() *int64
	SetFileVersion(v int32) *GetFileVersionRequest
	GetFileVersion() *int32
	SetProjectId(v int64) *GetFileVersionRequest
	GetProjectId() *int64
	SetProjectIdentifier(v string) *GetFileVersionRequest
	GetProjectIdentifier() *string
}

type GetFileVersionRequest struct {
	// The ID of the file. You can obtain the ID of the desired file by calling [ListFiles](https://help.aliyun.com/document_detail/173942.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 1000001
	FileId *int64 `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The version of the file to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	FileVersion *int32 `json:"FileVersion,omitempty" xml:"FileVersion,omitempty"`
	// The ID of the DataWorks workspace. You can click the wrench icon in the upper-right corner of the page to go to the Workspace Management page and view it.
	//
	// example:
	//
	// 1000011
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The unique identifier of the DataWorks workspace, which is the English identifier displayed at the top of the Data Development page when you switch workspaces.
	//
	// You must specify either this parameter or the ProjectId parameter to identify the DataWorks workspace for this API call.
	//
	// example:
	//
	// dw_project
	ProjectIdentifier *string `json:"ProjectIdentifier,omitempty" xml:"ProjectIdentifier,omitempty"`
}

func (s GetFileVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetFileVersionRequest) GoString() string {
	return s.String()
}

func (s *GetFileVersionRequest) GetFileId() *int64 {
	return s.FileId
}

func (s *GetFileVersionRequest) GetFileVersion() *int32 {
	return s.FileVersion
}

func (s *GetFileVersionRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetFileVersionRequest) GetProjectIdentifier() *string {
	return s.ProjectIdentifier
}

func (s *GetFileVersionRequest) SetFileId(v int64) *GetFileVersionRequest {
	s.FileId = &v
	return s
}

func (s *GetFileVersionRequest) SetFileVersion(v int32) *GetFileVersionRequest {
	s.FileVersion = &v
	return s
}

func (s *GetFileVersionRequest) SetProjectId(v int64) *GetFileVersionRequest {
	s.ProjectId = &v
	return s
}

func (s *GetFileVersionRequest) SetProjectIdentifier(v string) *GetFileVersionRequest {
	s.ProjectIdentifier = &v
	return s
}

func (s *GetFileVersionRequest) Validate() error {
	return dara.Validate(s)
}
