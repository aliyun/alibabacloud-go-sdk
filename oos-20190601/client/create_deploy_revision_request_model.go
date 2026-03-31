// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDeployRevisionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationName(v string) *CreateDeployRevisionRequest
	GetApplicationName() *string
	SetDeployResourceType(v string) *CreateDeployRevisionRequest
	GetDeployResourceType() *string
	SetDescription(v string) *CreateDeployRevisionRequest
	GetDescription() *string
	SetHooks(v string) *CreateDeployRevisionRequest
	GetHooks() *string
	SetLocation(v string) *CreateDeployRevisionRequest
	GetLocation() *string
	SetRevisionType(v string) *CreateDeployRevisionRequest
	GetRevisionType() *string
}

type CreateDeployRevisionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// AgentColin3
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// example:
	//
	// Kubernetes
	DeployResourceType *string `json:"DeployResourceType,omitempty" xml:"DeployResourceType,omitempty"`
	// example:
	//
	// 2026-01-03
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// {"applicationStart":"# Execute build process\\nbuild() {\\n  REPO_URL=\\"https://github.com/aldinokemal/go-whatsapp-web-multidevice.git\\"\\n  BRANCH=\\"main\\"\\n  DOCKERFILE_PATH=\\"./dockerfile\\"\\n  "}
	Hooks *string `json:"Hooks,omitempty" xml:"Hooks,omitempty"`
	// example:
	//
	// {"bucketName":"ecs-application-ui-test","objectName":"319137376-pipeline-run-319137376-task-1-cmd-exec.log","regionId":"cn-hangzhou"}
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// example:
	//
	// Command
	RevisionType *string `json:"RevisionType,omitempty" xml:"RevisionType,omitempty"`
}

func (s CreateDeployRevisionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDeployRevisionRequest) GoString() string {
	return s.String()
}

func (s *CreateDeployRevisionRequest) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *CreateDeployRevisionRequest) GetDeployResourceType() *string {
	return s.DeployResourceType
}

func (s *CreateDeployRevisionRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateDeployRevisionRequest) GetHooks() *string {
	return s.Hooks
}

func (s *CreateDeployRevisionRequest) GetLocation() *string {
	return s.Location
}

func (s *CreateDeployRevisionRequest) GetRevisionType() *string {
	return s.RevisionType
}

func (s *CreateDeployRevisionRequest) SetApplicationName(v string) *CreateDeployRevisionRequest {
	s.ApplicationName = &v
	return s
}

func (s *CreateDeployRevisionRequest) SetDeployResourceType(v string) *CreateDeployRevisionRequest {
	s.DeployResourceType = &v
	return s
}

func (s *CreateDeployRevisionRequest) SetDescription(v string) *CreateDeployRevisionRequest {
	s.Description = &v
	return s
}

func (s *CreateDeployRevisionRequest) SetHooks(v string) *CreateDeployRevisionRequest {
	s.Hooks = &v
	return s
}

func (s *CreateDeployRevisionRequest) SetLocation(v string) *CreateDeployRevisionRequest {
	s.Location = &v
	return s
}

func (s *CreateDeployRevisionRequest) SetRevisionType(v string) *CreateDeployRevisionRequest {
	s.RevisionType = &v
	return s
}

func (s *CreateDeployRevisionRequest) Validate() error {
	return dara.Validate(s)
}
