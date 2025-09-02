// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFolderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFolderId(v string) *GetFolderRequest
	GetFolderId() *string
	SetFolderPath(v string) *GetFolderRequest
	GetFolderPath() *string
	SetProjectId(v int64) *GetFolderRequest
	GetProjectId() *int64
	SetProjectIdentifier(v string) *GetFolderRequest
	GetProjectIdentifier() *string
}

type GetFolderRequest struct {
	// The ID of the folder. You must configure either this parameter or the FolderPath parameter. You can call the [ListFolders](https://help.aliyun.com/document_detail/173955.html) operation to query the ID.
	//
	// example:
	//
	// 273****
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	// The path of the folder. You must configure either this parameter or the FolderId parameter. You can call the [ListFolders](https://help.aliyun.com/document_detail/173955.html) operation to query the path.
	//
	// example:
	//
	// Business_process/my_first_business_process/MaxCompute/ods_layer
	FolderPath *string `json:"FolderPath,omitempty" xml:"FolderPath,omitempty"`
	// The DataWorks workspace ID. You can log on to the DataWorks console and go to the Workspace page to query the ID. You must configure either this parameter or the ProjectIdentifier parameter to determine the DataWorks workspace to which the operation is applied.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The name of the DataWorks workspace. You can log on to the DataWorks console and go to the Workspace Settings panel to obtain the name. You must specify either this parameter or ProjectId to determine the DataWorks workspace to which the operation is applied.
	//
	// example:
	//
	// dw_project
	ProjectIdentifier *string `json:"ProjectIdentifier,omitempty" xml:"ProjectIdentifier,omitempty"`
}

func (s GetFolderRequest) String() string {
	return dara.Prettify(s)
}

func (s GetFolderRequest) GoString() string {
	return s.String()
}

func (s *GetFolderRequest) GetFolderId() *string {
	return s.FolderId
}

func (s *GetFolderRequest) GetFolderPath() *string {
	return s.FolderPath
}

func (s *GetFolderRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetFolderRequest) GetProjectIdentifier() *string {
	return s.ProjectIdentifier
}

func (s *GetFolderRequest) SetFolderId(v string) *GetFolderRequest {
	s.FolderId = &v
	return s
}

func (s *GetFolderRequest) SetFolderPath(v string) *GetFolderRequest {
	s.FolderPath = &v
	return s
}

func (s *GetFolderRequest) SetProjectId(v int64) *GetFolderRequest {
	s.ProjectId = &v
	return s
}

func (s *GetFolderRequest) SetProjectIdentifier(v string) *GetFolderRequest {
	s.ProjectIdentifier = &v
	return s
}

func (s *GetFolderRequest) Validate() error {
	return dara.Validate(s)
}
