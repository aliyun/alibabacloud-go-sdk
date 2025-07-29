// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTriggerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *DescribeTriggerRequest
	GetName() *string
	SetNamespace(v string) *DescribeTriggerRequest
	GetNamespace() *string
	SetType(v string) *DescribeTriggerRequest
	GetType() *string
	SetAction(v string) *DescribeTriggerRequest
	GetAction() *string
}

type DescribeTriggerRequest struct {
	// The application name.
	//
	// This parameter is required.
	//
	// example:
	//
	// web-server
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The namespace to which the application belongs.
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

func (s DescribeTriggerRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTriggerRequest) GoString() string {
	return s.String()
}

func (s *DescribeTriggerRequest) GetName() *string {
	return s.Name
}

func (s *DescribeTriggerRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeTriggerRequest) GetType() *string {
	return s.Type
}

func (s *DescribeTriggerRequest) GetAction() *string {
	return s.Action
}

func (s *DescribeTriggerRequest) SetName(v string) *DescribeTriggerRequest {
	s.Name = &v
	return s
}

func (s *DescribeTriggerRequest) SetNamespace(v string) *DescribeTriggerRequest {
	s.Namespace = &v
	return s
}

func (s *DescribeTriggerRequest) SetType(v string) *DescribeTriggerRequest {
	s.Type = &v
	return s
}

func (s *DescribeTriggerRequest) SetAction(v string) *DescribeTriggerRequest {
	s.Action = &v
	return s
}

func (s *DescribeTriggerRequest) Validate() error {
	return dara.Validate(s)
}
