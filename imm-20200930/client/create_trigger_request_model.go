// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTriggerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActions(v []*CreateTriggerRequestActions) *CreateTriggerRequest
	GetActions() []*CreateTriggerRequestActions
	SetInput(v *Input) *CreateTriggerRequest
	GetInput() *Input
	SetNotification(v *CreateTriggerRequestNotification) *CreateTriggerRequest
	GetNotification() *CreateTriggerRequestNotification
	SetProjectName(v string) *CreateTriggerRequest
	GetProjectName() *string
	SetServiceRole(v string) *CreateTriggerRequest
	GetServiceRole() *string
	SetTags(v map[string]interface{}) *CreateTriggerRequest
	GetTags() map[string]interface{}
}

type CreateTriggerRequest struct {
	// The templates.
	//
	// This parameter is required.
	Actions []*CreateTriggerRequestActions `json:"Actions,omitempty" xml:"Actions,omitempty" type:"Repeated"`
	// The data source configurations.
	//
	// This parameter is required.
	Input *Input `json:"Input,omitempty" xml:"Input,omitempty"`
	// The notification settings. The operation supports multiple messaging middleware options. For more information about notification messages, see Asynchronous message examples. You can use one of the following methods to receive notification messages:
	//
	// In the region in which the IMM project is located, use EventBridge to receive task notifications. For more information, see IMM events. In the region in which the IMM project is located, configure a Simple Message Queue (SMQ) subscription to receive task notifications.
	Notification *CreateTriggerRequestNotification `json:"Notification,omitempty" xml:"Notification,omitempty" type:"Struct"`
	// The name of the project.[](~~478153~~)
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The service role. IMM assumes the service role so that it can access resources in other cloud services, such as OSS. Default value: AliyunIMMBatchTriggerRole.
	//
	// You can also create a custom service role in the RAM console and grant the required permissions to the role based on your business requirements. For more information, see [Create a regular service role](https://help.aliyun.com/document_detail/116800.html) and [Grant permissions to a role](https://help.aliyun.com/document_detail/116147.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// AliyunIMMDefaultRole
	ServiceRole *string `json:"ServiceRole,omitempty" xml:"ServiceRole,omitempty"`
	// The custom tags. You can search for or filter asynchronous tasks by custom tag.
	//
	// example:
	//
	// {"key":"val"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s CreateTriggerRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTriggerRequest) GoString() string {
	return s.String()
}

func (s *CreateTriggerRequest) GetActions() []*CreateTriggerRequestActions {
	return s.Actions
}

func (s *CreateTriggerRequest) GetInput() *Input {
	return s.Input
}

func (s *CreateTriggerRequest) GetNotification() *CreateTriggerRequestNotification {
	return s.Notification
}

func (s *CreateTriggerRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *CreateTriggerRequest) GetServiceRole() *string {
	return s.ServiceRole
}

func (s *CreateTriggerRequest) GetTags() map[string]interface{} {
	return s.Tags
}

func (s *CreateTriggerRequest) SetActions(v []*CreateTriggerRequestActions) *CreateTriggerRequest {
	s.Actions = v
	return s
}

func (s *CreateTriggerRequest) SetInput(v *Input) *CreateTriggerRequest {
	s.Input = v
	return s
}

func (s *CreateTriggerRequest) SetNotification(v *CreateTriggerRequestNotification) *CreateTriggerRequest {
	s.Notification = v
	return s
}

func (s *CreateTriggerRequest) SetProjectName(v string) *CreateTriggerRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateTriggerRequest) SetServiceRole(v string) *CreateTriggerRequest {
	s.ServiceRole = &v
	return s
}

func (s *CreateTriggerRequest) SetTags(v map[string]interface{}) *CreateTriggerRequest {
	s.Tags = v
	return s
}

func (s *CreateTriggerRequest) Validate() error {
	if s.Actions != nil {
		for _, item := range s.Actions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Input != nil {
		if err := s.Input.Validate(); err != nil {
			return err
		}
	}
	if s.Notification != nil {
		if err := s.Notification.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateTriggerRequestActions struct {
	// The policy configurations for handling failures.
	FastFailPolicy *FastFailPolicy `json:"FastFailPolicy,omitempty" xml:"FastFailPolicy,omitempty"`
	// The name of the template.
	//
	// This parameter is required.
	//
	// example:
	//
	// doc/convert
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The template parameters.
	Parameters []*string `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
}

func (s CreateTriggerRequestActions) String() string {
	return dara.Prettify(s)
}

func (s CreateTriggerRequestActions) GoString() string {
	return s.String()
}

func (s *CreateTriggerRequestActions) GetFastFailPolicy() *FastFailPolicy {
	return s.FastFailPolicy
}

func (s *CreateTriggerRequestActions) GetName() *string {
	return s.Name
}

func (s *CreateTriggerRequestActions) GetParameters() []*string {
	return s.Parameters
}

func (s *CreateTriggerRequestActions) SetFastFailPolicy(v *FastFailPolicy) *CreateTriggerRequestActions {
	s.FastFailPolicy = v
	return s
}

func (s *CreateTriggerRequestActions) SetName(v string) *CreateTriggerRequestActions {
	s.Name = &v
	return s
}

func (s *CreateTriggerRequestActions) SetParameters(v []*string) *CreateTriggerRequestActions {
	s.Parameters = v
	return s
}

func (s *CreateTriggerRequestActions) Validate() error {
	if s.FastFailPolicy != nil {
		if err := s.FastFailPolicy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateTriggerRequestNotification struct {
	// The SMQ notification settings.
	MNS *MNS `json:"MNS,omitempty" xml:"MNS,omitempty"`
}

func (s CreateTriggerRequestNotification) String() string {
	return dara.Prettify(s)
}

func (s CreateTriggerRequestNotification) GoString() string {
	return s.String()
}

func (s *CreateTriggerRequestNotification) GetMNS() *MNS {
	return s.MNS
}

func (s *CreateTriggerRequestNotification) SetMNS(v *MNS) *CreateTriggerRequestNotification {
	s.MNS = v
	return s
}

func (s *CreateTriggerRequestNotification) Validate() error {
	if s.MNS != nil {
		if err := s.MNS.Validate(); err != nil {
			return err
		}
	}
	return nil
}
