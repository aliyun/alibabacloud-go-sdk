// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBatchShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActionsShrink(v string) *CreateBatchShrinkRequest
	GetActionsShrink() *string
	SetInputShrink(v string) *CreateBatchShrinkRequest
	GetInputShrink() *string
	SetNotificationShrink(v string) *CreateBatchShrinkRequest
	GetNotificationShrink() *string
	SetProjectName(v string) *CreateBatchShrinkRequest
	GetProjectName() *string
	SetServiceRole(v string) *CreateBatchShrinkRequest
	GetServiceRole() *string
	SetTagsShrink(v string) *CreateBatchShrinkRequest
	GetTagsShrink() *string
}

type CreateBatchShrinkRequest struct {
	// A list of processing templates.
	//
	// This parameter is required.
	ActionsShrink *string `json:"Actions,omitempty" xml:"Actions,omitempty"`
	// The input data source configuration.
	//
	// This parameter is required.
	InputShrink *string `json:"Input,omitempty" xml:"Input,omitempty"`
	// The notification recipient. Multiple middleware services are supported. For details about the message format, see Asynchronous notification message. Select one of the following methods to receive messages:
	//
	// Activate and access EventBridge in the same region as Intelligent Media Management (IMM) to promptly receive task notifications. For more information, see IMM events.
	//
	// Activate MNS in the same region as IMM and configure a subscription.
	NotificationShrink *string `json:"Notification,omitempty" xml:"Notification,omitempty"`
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
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s CreateBatchShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateBatchShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateBatchShrinkRequest) GetActionsShrink() *string {
	return s.ActionsShrink
}

func (s *CreateBatchShrinkRequest) GetInputShrink() *string {
	return s.InputShrink
}

func (s *CreateBatchShrinkRequest) GetNotificationShrink() *string {
	return s.NotificationShrink
}

func (s *CreateBatchShrinkRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *CreateBatchShrinkRequest) GetServiceRole() *string {
	return s.ServiceRole
}

func (s *CreateBatchShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *CreateBatchShrinkRequest) SetActionsShrink(v string) *CreateBatchShrinkRequest {
	s.ActionsShrink = &v
	return s
}

func (s *CreateBatchShrinkRequest) SetInputShrink(v string) *CreateBatchShrinkRequest {
	s.InputShrink = &v
	return s
}

func (s *CreateBatchShrinkRequest) SetNotificationShrink(v string) *CreateBatchShrinkRequest {
	s.NotificationShrink = &v
	return s
}

func (s *CreateBatchShrinkRequest) SetProjectName(v string) *CreateBatchShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateBatchShrinkRequest) SetServiceRole(v string) *CreateBatchShrinkRequest {
	s.ServiceRole = &v
	return s
}

func (s *CreateBatchShrinkRequest) SetTagsShrink(v string) *CreateBatchShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *CreateBatchShrinkRequest) Validate() error {
	return dara.Validate(s)
}
