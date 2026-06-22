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
	// A list of data processing templates.
	//
	// This parameter is required.
	ActionsShrink *string `json:"Actions,omitempty" xml:"Actions,omitempty"`
	// The data source configuration.
	//
	// This parameter is required.
	InputShrink *string `json:"Input,omitempty" xml:"Input,omitempty"`
	// The notification recipient. Various message intermediaries are supported. For details about the message format, see Asynchronous notification message. Choose one of the following methods to receive messages:
	//
	// Activate and connect to EventBridge in the same region as Intelligent Media Management (IMM) to receive task notifications. For more information, see IMM events. Activate Message Service (MNS) in the same region as IMM and configure a subscription.
	NotificationShrink *string `json:"Notification,omitempty" xml:"Notification,omitempty"`
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
