// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFolderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFolderId(v string) *DeleteFolderRequest
	GetFolderId() *string
	SetProjectId(v int64) *DeleteFolderRequest
	GetProjectId() *int64
	SetProjectIdentifier(v string) *DeleteFolderRequest
	GetProjectIdentifier() *string
}

type DeleteFolderRequest struct {
	// The ID of the folder. For more information about how to obtain the folder ID, see [ListFolders](https://help.aliyun.com/document_detail/173955.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 2eb6f9****
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	// The ID of the DataWorks workspace. To obtain the workspace ID, log on to the DataWorks console and go to the Workspace Management page. You must set this parameter or ProjectIdentifier to specify the DataWorks workspace for the API call.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The name of the DataWorks workspace. To obtain the workspace name, log on to the DataWorks console and go to the Workspace Management page. You must set this parameter or ProjectId to specify the DataWorks workspace for the API call.
	//
	// example:
	//
	// dw_project
	ProjectIdentifier *string `json:"ProjectIdentifier,omitempty" xml:"ProjectIdentifier,omitempty"`
}

func (s DeleteFolderRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteFolderRequest) GoString() string {
	return s.String()
}

func (s *DeleteFolderRequest) GetFolderId() *string {
	return s.FolderId
}

func (s *DeleteFolderRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *DeleteFolderRequest) GetProjectIdentifier() *string {
	return s.ProjectIdentifier
}

func (s *DeleteFolderRequest) SetFolderId(v string) *DeleteFolderRequest {
	s.FolderId = &v
	return s
}

func (s *DeleteFolderRequest) SetProjectId(v int64) *DeleteFolderRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteFolderRequest) SetProjectIdentifier(v string) *DeleteFolderRequest {
	s.ProjectIdentifier = &v
	return s
}

func (s *DeleteFolderRequest) Validate() error {
	return dara.Validate(s)
}
