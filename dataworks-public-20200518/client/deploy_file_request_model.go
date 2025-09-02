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
	// The description of the deployment operation.
	//
	// example:
	//
	// First release task
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The file ID. You can call the [ListFiles](https://help.aliyun.com/document_detail/173942.html) operation to query the file ID. You must configure either this parameter or the NodeId parameter.
	//
	// example:
	//
	// 10000001
	FileId *int64 `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The ID of the node in the scheduling system that corresponds to the file that you want to deploy. You must configure either the NodeId parameter or the FileId parameter.
	//
	// example:
	//
	// 2000001
	NodeId *int64 `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The ID of the DataWorks workspace. You can log on to the DataWorks console and go to the Workspace Management page to obtain the workspace ID. You must configure either this parameter or the ProjectIdentifier parameter to determine the DataWorks workspace to which the operation is applied.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The name of the DataWorks workspace. You can log on to the DataWorks console and go to the Workspace Management page to obtain the workspace name. You must configure either this parameter or the ProjectId parameter to determine the DataWorks workspace to which the operation is applied.
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
