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
	// The trigger action. Valid values:
	//
	// `redeploy`: redeploys the resources defined in project_id.
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
	// The trigger project name.
	//
	// The value consists of the namespace and application name in the format of `${namespace}/${name}`.
	//
	// Example: `default/test-app`.
	//
	// This parameter is required.
	//
	// example:
	//
	// default/test-app
	ProjectId *string `json:"project_id,omitempty" xml:"project_id,omitempty"`
	// The trigger type. Valid values:
	//
	// - `deployment`: a trigger for stateless applications.
	//
	// - `application`: a trigger for Application Center applications.
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
