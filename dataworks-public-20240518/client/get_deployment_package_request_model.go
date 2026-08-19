// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeploymentPackageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeploymentId(v int64) *GetDeploymentPackageRequest
	GetDeploymentId() *int64
	SetProjectId(v int64) *GetDeploymentPackageRequest
	GetProjectId() *int64
	SetProjectIdentifier(v string) *GetDeploymentPackageRequest
	GetProjectIdentifier() *string
}

type GetDeploymentPackageRequest struct {
	// The ID of the deployment package.
	//
	// A deployment package ID is generated when you call [SubmitFile](https://help.aliyun.com/document_detail/173944.html) or [DeployFile](https://help.aliyun.com/document_detail/173956.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 1000000001
	DeploymentId *int64 `json:"DeploymentId,omitempty" xml:"DeploymentId,omitempty"`
	// The ID of the DataWorks workspace.
	//
	// You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the workspace configuration page to obtain the workspace ID.
	//
	// This parameter specifies the DataWorks workspace for this API call.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The unique identifier of the DataWorks workspace, which is the English identifier displayed at the top of the DataStudio page for switching workspaces.
	//
	// You must specify either this parameter or ProjectId to determine the DataWorks workspace for this API call.
	//
	// example:
	//
	// dw_project
	ProjectIdentifier *string `json:"ProjectIdentifier,omitempty" xml:"ProjectIdentifier,omitempty"`
}

func (s GetDeploymentPackageRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDeploymentPackageRequest) GoString() string {
	return s.String()
}

func (s *GetDeploymentPackageRequest) GetDeploymentId() *int64 {
	return s.DeploymentId
}

func (s *GetDeploymentPackageRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetDeploymentPackageRequest) GetProjectIdentifier() *string {
	return s.ProjectIdentifier
}

func (s *GetDeploymentPackageRequest) SetDeploymentId(v int64) *GetDeploymentPackageRequest {
	s.DeploymentId = &v
	return s
}

func (s *GetDeploymentPackageRequest) SetProjectId(v int64) *GetDeploymentPackageRequest {
	s.ProjectId = &v
	return s
}

func (s *GetDeploymentPackageRequest) SetProjectIdentifier(v string) *GetDeploymentPackageRequest {
	s.ProjectIdentifier = &v
	return s
}

func (s *GetDeploymentPackageRequest) Validate() error {
	return dara.Validate(s)
}
