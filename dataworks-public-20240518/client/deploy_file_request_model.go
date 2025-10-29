// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeployFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *DeployFileRequest
	GetComment() *string
	SetFileId(v int64) *DeployFileRequest
	GetFileId() *int64
	SetNodeId(v int64) *DeployFileRequest
	GetNodeId() *int64
	SetProjectId(v int64) *DeployFileRequest
	GetProjectId() *int64
	SetProjectIdentifier(v string) *DeployFileRequest
	GetProjectIdentifier() *string
}

type DeployFileRequest struct {
	// The description of the deployment.
	//
	// example:
	//
	// First release task
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The file ID. You can call the [ListFiles](https://help.aliyun.com/document_detail/173942.html) operation to obtain the ID. You need to configure either this parameter or the NodeId parameter.
	//
	// example:
	//
	// 10000001
	FileId *int64 `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The task ID of the file to be deployed in the scheduling system. You need to configure either this parameter or the FileId parameter.
	//
	// example:
	//
	// 2000001
	NodeId *int64 `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The DataWorks workspace ID. You can log on to the DataWorks console and go to the Workspace page to query the ID. You must specify either this parameter or the ProjectIdentifier parameter to identify the DataWorks workspace when you call this operation.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The name of the DataWorks workspace. You can log on to the DataWorks console and go to the Workspace page to query the workspace name. You must specify either this parameter or the ProjectId parameter to identify the DataWorks workspace when you call this operation.
	//
	// example:
	//
	// dw_project
	ProjectIdentifier *string `json:"ProjectIdentifier,omitempty" xml:"ProjectIdentifier,omitempty"`
}

func (s DeployFileRequest) String() string {
	return dara.Prettify(s)
}

func (s DeployFileRequest) GoString() string {
	return s.String()
}

func (s *DeployFileRequest) GetComment() *string {
	return s.Comment
}

func (s *DeployFileRequest) GetFileId() *int64 {
	return s.FileId
}

func (s *DeployFileRequest) GetNodeId() *int64 {
	return s.NodeId
}

func (s *DeployFileRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *DeployFileRequest) GetProjectIdentifier() *string {
	return s.ProjectIdentifier
}

func (s *DeployFileRequest) SetComment(v string) *DeployFileRequest {
	s.Comment = &v
	return s
}

func (s *DeployFileRequest) SetFileId(v int64) *DeployFileRequest {
	s.FileId = &v
	return s
}

func (s *DeployFileRequest) SetNodeId(v int64) *DeployFileRequest {
	s.NodeId = &v
	return s
}

func (s *DeployFileRequest) SetProjectId(v int64) *DeployFileRequest {
	s.ProjectId = &v
	return s
}

func (s *DeployFileRequest) SetProjectIdentifier(v string) *DeployFileRequest {
	s.ProjectIdentifier = &v
	return s
}

func (s *DeployFileRequest) Validate() error {
	return dara.Validate(s)
}
