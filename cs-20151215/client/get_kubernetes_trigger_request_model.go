// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKubernetesTriggerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *GetKubernetesTriggerRequest
	GetName() *string
	SetNamespace(v string) *GetKubernetesTriggerRequest
	GetNamespace() *string
	SetType(v string) *GetKubernetesTriggerRequest
	GetType() *string
	SetAction(v string) *GetKubernetesTriggerRequest
	GetAction() *string
}

type GetKubernetesTriggerRequest struct {
	// The application name.
	//
	// This parameter is required.
	//
	// example:
	//
	// web-server
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The namespace name.
	//
	// This parameter is required.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The type of trigger. Valid values:
	//
	// 	- `deployment`: performs actions on Deployments.
	//
	// 	- `application`: performs actions on applications that are deployed in Application Center.
	//
	// Default value: `deployment`.
	//
	// If you do not set this parameter, triggers are not filtered by type.
	//
	// example:
	//
	// deployment
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The action that the trigger performs. Set the value to redeploy.
	//
	// `redeploy`: redeploys the resources specified by `project_id`.
	//
	// If you do not specify this parameter, triggers are not filtered by action.
	//
	// example:
	//
	// redeploy
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
}

func (s GetKubernetesTriggerRequest) String() string {
	return dara.Prettify(s)
}

func (s GetKubernetesTriggerRequest) GoString() string {
	return s.String()
}

func (s *GetKubernetesTriggerRequest) GetName() *string {
	return s.Name
}

func (s *GetKubernetesTriggerRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *GetKubernetesTriggerRequest) GetType() *string {
	return s.Type
}

func (s *GetKubernetesTriggerRequest) GetAction() *string {
	return s.Action
}

func (s *GetKubernetesTriggerRequest) SetName(v string) *GetKubernetesTriggerRequest {
	s.Name = &v
	return s
}

func (s *GetKubernetesTriggerRequest) SetNamespace(v string) *GetKubernetesTriggerRequest {
	s.Namespace = &v
	return s
}

func (s *GetKubernetesTriggerRequest) SetType(v string) *GetKubernetesTriggerRequest {
	s.Type = &v
	return s
}

func (s *GetKubernetesTriggerRequest) SetAction(v string) *GetKubernetesTriggerRequest {
	s.Action = &v
	return s
}

func (s *GetKubernetesTriggerRequest) Validate() error {
	return dara.Validate(s)
}
