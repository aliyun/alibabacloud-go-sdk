// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKubernetesTriggerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAction(v string) *CreateKubernetesTriggerResponseBody
	GetAction() *string
	SetClusterId(v string) *CreateKubernetesTriggerResponseBody
	GetClusterId() *string
	SetId(v string) *CreateKubernetesTriggerResponseBody
	GetId() *string
	SetProjectId(v string) *CreateKubernetesTriggerResponseBody
	GetProjectId() *string
	SetType(v string) *CreateKubernetesTriggerResponseBody
	GetType() *string
}

type CreateKubernetesTriggerResponseBody struct {
	// The action that the trigger performs. For example, a value of `redeploy` indicates that the trigger redeploys the application.
	//
	// example:
	//
	// redeploy
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// The ID of the cluster.
	//
	// example:
	//
	// c5cdf7e3938bc4f8eb0e44b21a80f****
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	// The ID of the trigger.
	//
	// example:
	//
	// 111
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The name of the trigger project.
	//
	// example:
	//
	// default/test-app
	ProjectId *string `json:"project_id,omitempty" xml:"project_id,omitempty"`
	// The type of trigger.
	//
	// Valid values:
	//
	// 	- `deployment`: performs actions on Deployments.
	//
	// 	- `application`: performs actions on applications that are deployed in Application Center.
	//
	// example:
	//
	// deployment
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateKubernetesTriggerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateKubernetesTriggerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateKubernetesTriggerResponseBody) GetAction() *string {
	return s.Action
}

func (s *CreateKubernetesTriggerResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateKubernetesTriggerResponseBody) GetId() *string {
	return s.Id
}

func (s *CreateKubernetesTriggerResponseBody) GetProjectId() *string {
	return s.ProjectId
}

func (s *CreateKubernetesTriggerResponseBody) GetType() *string {
	return s.Type
}

func (s *CreateKubernetesTriggerResponseBody) SetAction(v string) *CreateKubernetesTriggerResponseBody {
	s.Action = &v
	return s
}

func (s *CreateKubernetesTriggerResponseBody) SetClusterId(v string) *CreateKubernetesTriggerResponseBody {
	s.ClusterId = &v
	return s
}

func (s *CreateKubernetesTriggerResponseBody) SetId(v string) *CreateKubernetesTriggerResponseBody {
	s.Id = &v
	return s
}

func (s *CreateKubernetesTriggerResponseBody) SetProjectId(v string) *CreateKubernetesTriggerResponseBody {
	s.ProjectId = &v
	return s
}

func (s *CreateKubernetesTriggerResponseBody) SetType(v string) *CreateKubernetesTriggerResponseBody {
	s.Type = &v
	return s
}

func (s *CreateKubernetesTriggerResponseBody) Validate() error {
	return dara.Validate(s)
}
