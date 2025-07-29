// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKubernetesTriggerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAction(v string) *CreateKubernetesTriggerRequest
	GetAction() *string
	SetClusterId(v string) *CreateKubernetesTriggerRequest
	GetClusterId() *string
	SetProjectId(v string) *CreateKubernetesTriggerRequest
	GetProjectId() *string
	SetType(v string) *CreateKubernetesTriggerRequest
	GetType() *string
}

type CreateKubernetesTriggerRequest struct {
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

func (s CreateKubernetesTriggerRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateKubernetesTriggerRequest) GoString() string {
	return s.String()
}

func (s *CreateKubernetesTriggerRequest) GetAction() *string {
	return s.Action
}

func (s *CreateKubernetesTriggerRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateKubernetesTriggerRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *CreateKubernetesTriggerRequest) GetType() *string {
	return s.Type
}

func (s *CreateKubernetesTriggerRequest) SetAction(v string) *CreateKubernetesTriggerRequest {
	s.Action = &v
	return s
}

func (s *CreateKubernetesTriggerRequest) SetClusterId(v string) *CreateKubernetesTriggerRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateKubernetesTriggerRequest) SetProjectId(v string) *CreateKubernetesTriggerRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateKubernetesTriggerRequest) SetType(v string) *CreateKubernetesTriggerRequest {
	s.Type = &v
	return s
}

func (s *CreateKubernetesTriggerRequest) Validate() error {
	return dara.Validate(s)
}
