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
	// A list of data processing templates.
	//
	// This parameter is required.
	Actions []*CreateTriggerRequestActions `json:"Actions,omitempty" xml:"Actions,omitempty" type:"Repeated"`
	// The data source configuration.
	//
	// This parameter is required.
	Input *Input `json:"Input,omitempty" xml:"Input,omitempty"`
	// The notification recipient. Various message intermediaries are supported. For details about the message format, see Asynchronous notification message. Choose one of the following methods to receive messages:
	//
	// Activate and connect to EventBridge in the same region as Intelligent Media Management (IMM) to receive task notifications. For more information, see IMM events. Activate Message Service (MNS) in the same region as IMM and configure a subscription.
	Notification *CreateTriggerRequestNotification `json:"Notification,omitempty" xml:"Notification,omitempty" type:"Struct"`
	// The project name. For more information, see [Create a project](https://help.aliyun.com/document_detail/478153.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The service role that grants Intelligent Media Management (IMM) permissions to access other cloud resources, such as Object Storage Service (OSS). The default value is AliyunIMMBatchTriggerRole.
	//
	// To use a custom service role, create a service role and grant permissions to the role in the Resource Access Management (RAM) console. For more information, see [Create a service role](https://help.aliyun.com/document_detail/116800.html) and [Grant permissions to a RAM role](https://help.aliyun.com/document_detail/116147.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// AliyunIMMBatchTriggerRole
	ServiceRole *string `json:"ServiceRole,omitempty" xml:"ServiceRole,omitempty"`
	// Custom tags used to search and filter asynchronous tasks.
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
	// The fast-fail policy configuration.
	FastFailPolicy *FastFailPolicy `json:"FastFailPolicy,omitempty" xml:"FastFailPolicy,omitempty"`
	// The name of the template.
	//
	// This parameter is required.
	//
	// example:
	//
	// doc/convert
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// A list of template parameters.
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
	// The parameter object for MNS notifications.
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
