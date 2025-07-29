// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTriggerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAction(v string) *CreateTriggerResponseBody
	GetAction() *string
	SetClusterId(v string) *CreateTriggerResponseBody
	GetClusterId() *string
	SetId(v string) *CreateTriggerResponseBody
	GetId() *string
	SetProjectId(v string) *CreateTriggerResponseBody
	GetProjectId() *string
	SetType(v string) *CreateTriggerResponseBody
	GetType() *string
}

type CreateTriggerResponseBody struct {
	// The actions performed by the trigger.
	//
	// example:
	//
	// redeploy
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// The ID of the cluster.
	//
	// example:
	//
	// c93095129fc41463aa455d89444fd****
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	// The trigger ID.
	//
	// example:
	//
	// 102536
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The name of the trigger project.
	//
	// example:
	//
	// default/test-app
	ProjectId *string `json:"project_id,omitempty" xml:"project_id,omitempty"`
	// The trigger type.
	//
	// example:
	//
	// deployment
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateTriggerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTriggerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTriggerResponseBody) GetAction() *string {
	return s.Action
}

func (s *CreateTriggerResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateTriggerResponseBody) GetId() *string {
	return s.Id
}

func (s *CreateTriggerResponseBody) GetProjectId() *string {
	return s.ProjectId
}

func (s *CreateTriggerResponseBody) GetType() *string {
	return s.Type
}

func (s *CreateTriggerResponseBody) SetAction(v string) *CreateTriggerResponseBody {
	s.Action = &v
	return s
}

func (s *CreateTriggerResponseBody) SetClusterId(v string) *CreateTriggerResponseBody {
	s.ClusterId = &v
	return s
}

func (s *CreateTriggerResponseBody) SetId(v string) *CreateTriggerResponseBody {
	s.Id = &v
	return s
}

func (s *CreateTriggerResponseBody) SetProjectId(v string) *CreateTriggerResponseBody {
	s.ProjectId = &v
	return s
}

func (s *CreateTriggerResponseBody) SetType(v string) *CreateTriggerResponseBody {
	s.Type = &v
	return s
}

func (s *CreateTriggerResponseBody) Validate() error {
	return dara.Validate(s)
}
