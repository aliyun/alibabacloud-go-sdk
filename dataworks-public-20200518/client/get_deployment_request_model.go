// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeploymentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeploymentId(v int64) *GetDeploymentRequest
	GetDeploymentId() *int64
	SetProjectId(v int64) *GetDeploymentRequest
	GetProjectId() *int64
	SetProjectIdentifier(v string) *GetDeploymentRequest
	GetProjectIdentifier() *string
}

type GetDeploymentRequest struct {
	// The ID of the deployment task. A deployment task ID is generated when you call the [SubmitFile](https://help.aliyun.com/document_detail/173944.html) or [DeployFile](https://help.aliyun.com/document_detail/173956.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3000001
	DeploymentId *int64 `json:"DeploymentId,omitempty" xml:"DeploymentId,omitempty"`
	// The ID of the DataWorks workspace. You can click the Workspace Manage icon in the upper-right corner of the DataStudio page to go to the Workspace Management page and view the workspace ID.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The GUID of the DataWorks workspace. You can view the GUID in the upper part of the DataStudio page. You can also select another GUID to switch to another workspace.
	//
	// You must specify either this parameter or the ProjectId parameter to determine the DataWorks workspace to which the operation is applied.
	//
	// example:
	//
	// dw_project
	ProjectIdentifier *string `json:"ProjectIdentifier,omitempty" xml:"ProjectIdentifier,omitempty"`
}

func (s GetDeploymentRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDeploymentRequest) GoString() string {
	return s.String()
}

func (s *GetDeploymentRequest) GetDeploymentId() *int64 {
	return s.DeploymentId
}

func (s *GetDeploymentRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetDeploymentRequest) GetProjectIdentifier() *string {
	return s.ProjectIdentifier
}

func (s *GetDeploymentRequest) SetDeploymentId(v int64) *GetDeploymentRequest {
	s.DeploymentId = &v
	return s
}

func (s *GetDeploymentRequest) SetProjectId(v int64) *GetDeploymentRequest {
	s.ProjectId = &v
	return s
}

func (s *GetDeploymentRequest) SetProjectIdentifier(v string) *GetDeploymentRequest {
	s.ProjectIdentifier = &v
	return s
}

func (s *GetDeploymentRequest) Validate() error {
	return dara.Validate(s)
}
