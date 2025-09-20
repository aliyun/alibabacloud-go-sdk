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
