// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTriggerShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActionsShrink(v string) *CreateTriggerShrinkRequest
	GetActionsShrink() *string
	SetInputShrink(v string) *CreateTriggerShrinkRequest
	GetInputShrink() *string
	SetNotificationShrink(v string) *CreateTriggerShrinkRequest
	GetNotificationShrink() *string
	SetProjectName(v string) *CreateTriggerShrinkRequest
	GetProjectName() *string
	SetServiceRole(v string) *CreateTriggerShrinkRequest
	GetServiceRole() *string
	SetTagsShrink(v string) *CreateTriggerShrinkRequest
	GetTagsShrink() *string
}

type CreateTriggerShrinkRequest struct {
	// The templates.
	//
	// This parameter is required.
	ActionsShrink *string `json:"Actions,omitempty" xml:"Actions,omitempty"`
	// The data source configurations.
	//
	// This parameter is required.
	InputShrink *string `json:"Input,omitempty" xml:"Input,omitempty"`
	// The notification settings. The operation supports multiple messaging middleware options. For more information about notification messages, see Asynchronous message examples. You can use one of the following methods to receive notification messages:
	//
	// In the region in which the IMM project is located, use EventBridge to receive task notifications. For more information, see IMM events. In the region in which the IMM project is located, configure a Simple Message Queue (SMQ) subscription to receive task notifications.
	NotificationShrink *string `json:"Notification,omitempty" xml:"Notification,omitempty"`
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
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s CreateTriggerShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTriggerShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateTriggerShrinkRequest) GetActionsShrink() *string {
	return s.ActionsShrink
}

func (s *CreateTriggerShrinkRequest) GetInputShrink() *string {
	return s.InputShrink
}

func (s *CreateTriggerShrinkRequest) GetNotificationShrink() *string {
	return s.NotificationShrink
}

func (s *CreateTriggerShrinkRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *CreateTriggerShrinkRequest) GetServiceRole() *string {
	return s.ServiceRole
}

func (s *CreateTriggerShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *CreateTriggerShrinkRequest) SetActionsShrink(v string) *CreateTriggerShrinkRequest {
	s.ActionsShrink = &v
	return s
}

func (s *CreateTriggerShrinkRequest) SetInputShrink(v string) *CreateTriggerShrinkRequest {
	s.InputShrink = &v
	return s
}

func (s *CreateTriggerShrinkRequest) SetNotificationShrink(v string) *CreateTriggerShrinkRequest {
	s.NotificationShrink = &v
	return s
}

func (s *CreateTriggerShrinkRequest) SetProjectName(v string) *CreateTriggerShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateTriggerShrinkRequest) SetServiceRole(v string) *CreateTriggerShrinkRequest {
	s.ServiceRole = &v
	return s
}

func (s *CreateTriggerShrinkRequest) SetTagsShrink(v string) *CreateTriggerShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *CreateTriggerShrinkRequest) Validate() error {
	return dara.Validate(s)
}
