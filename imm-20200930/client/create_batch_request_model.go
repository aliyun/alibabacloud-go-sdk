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
	// The templates.
	//
	// This parameter is required.
	Actions []*CreateBatchRequestActions `json:"Actions,omitempty" xml:"Actions,omitempty" type:"Repeated"`
	// The data source configurations.
	//
	// This parameter is required.
	Input *Input `json:"Input,omitempty" xml:"Input,omitempty"`
	// The notification settings. The operation supports multiple messaging middleware options. For more information about notification messages, see Asynchronous message examples. You can use one of the following methods to receive notification messages:
	//
	// In the region in which the IMM project is located, use EventBridge to receive task notifications. For more information, see IMM events. In the region in which the IMM project is located, configure a Simple Message Queue (SMQ) subscription to receive task notifications.
	Notification *CreateBatchRequestNotification `json:"Notification,omitempty" xml:"Notification,omitempty" type:"Struct"`
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
	return dara.Validate(s)
}

type CreateBatchRequestActions struct {
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
	return dara.Validate(s)
}

type CreateBatchRequestNotification struct {
	// The SMQ notification settings.
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
	return dara.Validate(s)
}
