// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileId(v int64) *GetFileRequest
	GetFileId() *int64
	SetNodeId(v int64) *GetFileRequest
	GetNodeId() *int64
	SetProjectId(v int64) *GetFileRequest
	GetProjectId() *int64
	SetProjectIdentifier(v string) *GetFileRequest
	GetProjectIdentifier() *string
}

type GetFileRequest struct {
	// The file ID. You can call the [ListFiles](https://help.aliyun.com/document_detail/173942.html) operation to query the ID.
	//
	// example:
	//
	// 100000001
	FileId *int64 `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The ID of the node that is scheduled. You can call the [ListFiles](https://help.aliyun.com/document_detail/173942.html) operation to query the ID.
	//
	// example:
	//
	// 200000001
	NodeId *int64 `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to query the ID.
	//
	// You must configure either this parameter or the ProjectIdentifier parameter to determine the DataWorks workspace to which the operation is applied.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The name of the DataWorks workspace. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to query the name.
	//
	// You must configure either this parameter or the ProjectId parameter to determine the DataWorks workspace to which the operation is applied.
	//
	// example:
	//
	// dw_project
	ProjectIdentifier *string `json:"ProjectIdentifier,omitempty" xml:"ProjectIdentifier,omitempty"`
}

func (s GetFileRequest) String() string {
	return dara.Prettify(s)
}

func (s GetFileRequest) GoString() string {
	return s.String()
}

func (s *GetFileRequest) GetFileId() *int64 {
	return s.FileId
}

func (s *GetFileRequest) GetNodeId() *int64 {
	return s.NodeId
}

func (s *GetFileRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetFileRequest) GetProjectIdentifier() *string {
	return s.ProjectIdentifier
}

func (s *GetFileRequest) SetFileId(v int64) *GetFileRequest {
	s.FileId = &v
	return s
}

func (s *GetFileRequest) SetNodeId(v int64) *GetFileRequest {
	s.NodeId = &v
	return s
}

func (s *GetFileRequest) SetProjectId(v int64) *GetFileRequest {
	s.ProjectId = &v
	return s
}

func (s *GetFileRequest) SetProjectIdentifier(v string) *GetFileRequest {
	s.ProjectIdentifier = &v
	return s
}

func (s *GetFileRequest) Validate() error {
	return dara.Validate(s)
}
