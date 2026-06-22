// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBatchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActions(v []*CreateBatchRequestActions) *CreateBatchRequest
	GetActions() []*CreateBatchRequestActions
	SetInput(v *Input) *CreateBatchRequest
	GetInput() *Input
	SetNotification(v *CreateBatchRequestNotification) *CreateBatchRequest
	GetNotification() *CreateBatchRequestNotification
	SetProjectName(v string) *CreateBatchRequest
	GetProjectName() *string
	SetServiceRole(v string) *CreateBatchRequest
	GetServiceRole() *string
	SetTags(v map[string]interface{}) *CreateBatchRequest
	GetTags() map[string]interface{}
}

type CreateBatchRequest struct {
	// A list of processing templates.
	//
	// This parameter is required.
	Actions []*CreateBatchRequestActions `json:"Actions,omitempty" xml:"Actions,omitempty" type:"Repeated"`
	// The input data source configuration.
	//
	// This parameter is required.
	Input *Input `json:"Input,omitempty" xml:"Input,omitempty"`
	// The notification recipient. Multiple middleware services are supported. For details about the message format, see Asynchronous notification message. Select one of the following methods to receive messages:
	//
	// Activate and access EventBridge in the same region as Intelligent Media Management (IMM) to promptly receive task notifications. For more information, see IMM events.
	//
	// Activate MNS in the same region as IMM and configure a subscription.
	Notification *CreateBatchRequestNotification `json:"Notification,omitempty" xml:"Notification,omitempty" type:"Struct"`
	// The project name. For more information, see [Create a project](https://help.aliyun.com/document_detail/478153.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The service role that grants IMM the permissions to access other cloud resources, such as OSS. The default value is AliyunIMMBatchTriggerRole.
	//
	// To use a custom service role, create a service role and grant permissions to the role in the RAM console. For more information, see [Create a service role](https://help.aliyun.com/document_detail/116800.html) and [Grant permissions to a RAM role](https://help.aliyun.com/document_detail/116147.html).
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
	// {"key": "val"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s CreateBatchRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateBatchRequest) GoString() string {
	return s.String()
}

func (s *CreateBatchRequest) GetActions() []*CreateBatchRequestActions {
	return s.Actions
}

func (s *CreateBatchRequest) GetInput() *Input {
	return s.Input
}

func (s *CreateBatchRequest) GetNotification() *CreateBatchRequestNotification {
	return s.Notification
}

func (s *CreateBatchRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *CreateBatchRequest) GetServiceRole() *string {
	return s.ServiceRole
}

func (s *CreateBatchRequest) GetTags() map[string]interface{} {
	return s.Tags
}

func (s *CreateBatchRequest) SetActions(v []*CreateBatchRequestActions) *CreateBatchRequest {
	s.Actions = v
	return s
}

func (s *CreateBatchRequest) SetInput(v *Input) *CreateBatchRequest {
	s.Input = v
	return s
}

func (s *CreateBatchRequest) SetNotification(v *CreateBatchRequestNotification) *CreateBatchRequest {
	s.Notification = v
	return s
}

func (s *CreateBatchRequest) SetProjectName(v string) *CreateBatchRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateBatchRequest) SetServiceRole(v string) *CreateBatchRequest {
	s.ServiceRole = &v
	return s
}

func (s *CreateBatchRequest) SetTags(v map[string]interface{}) *CreateBatchRequest {
	s.Tags = v
	return s
}

func (s *CreateBatchRequest) Validate() error {
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

type CreateBatchRequestActions struct {
	// The fast-fail policy configuration.
	FastFailPolicy *FastFailPolicy `json:"FastFailPolicy,omitempty" xml:"FastFailPolicy,omitempty"`
	// The template name.
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

func (s CreateBatchRequestActions) String() string {
	return dara.Prettify(s)
}

func (s CreateBatchRequestActions) GoString() string {
	return s.String()
}

func (s *CreateBatchRequestActions) GetFastFailPolicy() *FastFailPolicy {
	return s.FastFailPolicy
}

func (s *CreateBatchRequestActions) GetName() *string {
	return s.Name
}

func (s *CreateBatchRequestActions) GetParameters() []*string {
	return s.Parameters
}

func (s *CreateBatchRequestActions) SetFastFailPolicy(v *FastFailPolicy) *CreateBatchRequestActions {
	s.FastFailPolicy = v
	return s
}

func (s *CreateBatchRequestActions) SetName(v string) *CreateBatchRequestActions {
	s.Name = &v
	return s
}

func (s *CreateBatchRequestActions) SetParameters(v []*string) *CreateBatchRequestActions {
	s.Parameters = v
	return s
}

func (s *CreateBatchRequestActions) Validate() error {
	if s.FastFailPolicy != nil {
		if err := s.FastFailPolicy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateBatchRequestNotification struct {
	// The parameters for MNS notifications.
	MNS *MNS `json:"MNS,omitempty" xml:"MNS,omitempty"`
}

func (s CreateBatchRequestNotification) String() string {
	return dara.Prettify(s)
}

func (s CreateBatchRequestNotification) GoString() string {
	return s.String()
}

func (s *CreateBatchRequestNotification) GetMNS() *MNS {
	return s.MNS
}

func (s *CreateBatchRequestNotification) SetMNS(v *MNS) *CreateBatchRequestNotification {
	s.MNS = v
	return s
}

func (s *CreateBatchRequestNotification) Validate() error {
	if s.MNS != nil {
		if err := s.MNS.Validate(); err != nil {
			return err
		}
	}
	return nil
}
