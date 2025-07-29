// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTriggerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAction(v string) *CreateTriggerRequest
	GetAction() *string
	SetClusterId(v string) *CreateTriggerRequest
	GetClusterId() *string
	SetProjectId(v string) *CreateTriggerRequest
	GetProjectId() *string
	SetType(v string) *CreateTriggerRequest
	GetType() *string
}

type CreateTriggerRequest struct {
	// The action that the trigger performs. Set the value to redeploy.
	//
	// `redeploy`: redeploys the resources specified by `project_id`.
	//
	// This parameter is required.
	//
	// example:
	//
	// redeploy
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c5cdf7e3938bc4f8eb0e44b21a80f****
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	// The name of the trigger project.
	//
	// The name consists of the namespace where the application is deployed and the name of the application. The format is `${namespace}/${name}`.
	//
	// Example: `default/test-app`.
	//
	// This parameter is required.
	//
	// example:
	//
	// default/test-app
	ProjectId *string `json:"project_id,omitempty" xml:"project_id,omitempty"`
	// The type of trigger. Valid values:
	//
	// 	- `deployment`: performs actions on Deployments.
	//
	// 	- `application`: performs actions on applications that are deployed in Application Center.
	//
	// Default value: `deployment`.
	//
	// example:
	//
	// deployment
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateTriggerRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTriggerRequest) GoString() string {
	return s.String()
}

func (s *CreateTriggerRequest) GetAction() *string {
	return s.Action
}

func (s *CreateTriggerRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateTriggerRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *CreateTriggerRequest) GetType() *string {
	return s.Type
}

func (s *CreateTriggerRequest) SetAction(v string) *CreateTriggerRequest {
	s.Action = &v
	return s
}

func (s *CreateTriggerRequest) SetClusterId(v string) *CreateTriggerRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateTriggerRequest) SetProjectId(v string) *CreateTriggerRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateTriggerRequest) SetType(v string) *CreateTriggerRequest {
	s.Type = &v
	return s
}

func (s *CreateTriggerRequest) Validate() error {
	return dara.Validate(s)
}
