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
	// The folder ID. Either this parameter or FolderPath must be specified. You can call the [ListFolders](https://help.aliyun.com/document_detail/173955.html) operation to obtain the folder ID.
	//
	// example:
	//
	// 273****
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	// The folder path. Either this parameter or FolderId must be specified. You can call the [ListFolders](https://help.aliyun.com/document_detail/173955.html) operation to obtain the folder path.
	//
	// example:
	//
	// Business_process/my_first_business_process/MaxCompute/ods_layer
	FolderPath *string `json:"FolderPath,omitempty" xml:"FolderPath,omitempty"`
	// The ID of the DataWorks workspace. You can obtain the workspace ID from the workspace configuration page in the DataWorks console. Either this parameter or ProjectIdentifier must be specified to determine which DataWorks workspace this API call operates on.
	//
	// example:
	//
	// 1000011
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The name of the DataWorks workspace. You can obtain the workspace name from the workspace configuration page in the DataWorks console. Either this parameter or ProjectId must be specified to determine which DataWorks workspace this API call operates on.
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
