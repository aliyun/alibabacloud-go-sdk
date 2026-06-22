// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMediaConvertTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlignmentIndex(v int32) *CreateMediaConvertTaskShrinkRequest
	GetAlignmentIndex() *int32
	SetCredentialConfigShrink(v string) *CreateMediaConvertTaskShrinkRequest
	GetCredentialConfigShrink() *string
	SetNotificationShrink(v string) *CreateMediaConvertTaskShrinkRequest
	GetNotificationShrink() *string
	SetProjectName(v string) *CreateMediaConvertTaskShrinkRequest
	GetProjectName() *string
	SetSourcesShrink(v string) *CreateMediaConvertTaskShrinkRequest
	GetSourcesShrink() *string
	SetTagsShrink(v string) *CreateMediaConvertTaskShrinkRequest
	GetTagsShrink() *string
	SetTargetGroupsShrink(v string) *CreateMediaConvertTaskShrinkRequest
	GetTargetGroupsShrink() *string
	SetTargetsShrink(v string) *CreateMediaConvertTaskShrinkRequest
	GetTargetsShrink() *string
	SetUserData(v string) *CreateMediaConvertTaskShrinkRequest
	GetUserData() *string
}

type CreateMediaConvertTaskShrinkRequest struct {
	// When concatenating media files, this specifies the index of the primary file in the Sources list. The default transcoding parameters (such as resolution and frame rate from the `Video` and `Audio` objects) are taken from this primary file. The default index is 0.
	//
	// example:
	//
	// 0
	AlignmentIndex *int32 `json:"AlignmentIndex,omitempty" xml:"AlignmentIndex,omitempty"`
	// **You can leave this parameter empty if you do not have special requirements.**
	//
	// The chained authorization configuration. For more information, see [Use chained authorization to access resources of other entities](https://help.aliyun.com/document_detail/465340.html).
	CredentialConfigShrink *string `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	// The message notification settings. For more information, click Notification. For information about the format of asynchronous notifications, see [Asynchronous notification format](https://help.aliyun.com/document_detail/2743997.html).
	NotificationShrink *string `json:"Notification,omitempty" xml:"Notification,omitempty"`
	// The name of the project. For more information about how to obtain the project name, see [Create a project](https://help.aliyun.com/document_detail/478153.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// A list of media files. If you provide more than one file, they are concatenated in the order of their URIs.
	//
	// This parameter is required.
	SourcesShrink *string `json:"Sources,omitempty" xml:"Sources,omitempty"`
	// Custom tags for searching and filtering asynchronous tasks.
	//
	// example:
	//
	// {"test":"val1"}
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// A list of media packaging tasks to convert and package the input media into HLS outputs. Each TargetGroup corresponds to one HLS master playlist.
	TargetGroupsShrink *string `json:"TargetGroups,omitempty" xml:"TargetGroups,omitempty"`
	// A list of media processing tasks.
	TargetsShrink *string `json:"Targets,omitempty" xml:"Targets,omitempty"`
	// The custom user data. This data is returned in the asynchronous notification, allowing you to associate the notification with your internal system. The maximum length is 2,048 bytes.
	//
	// example:
	//
	// {"ID": "testuid","Name": "test-user","Avatar": "http://test.com/testuid"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s CreateMediaConvertTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMediaConvertTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateMediaConvertTaskShrinkRequest) GetAlignmentIndex() *int32 {
	return s.AlignmentIndex
}

func (s *CreateMediaConvertTaskShrinkRequest) GetCredentialConfigShrink() *string {
	return s.CredentialConfigShrink
}

func (s *CreateMediaConvertTaskShrinkRequest) GetNotificationShrink() *string {
	return s.NotificationShrink
}

func (s *CreateMediaConvertTaskShrinkRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *CreateMediaConvertTaskShrinkRequest) GetSourcesShrink() *string {
	return s.SourcesShrink
}

func (s *CreateMediaConvertTaskShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *CreateMediaConvertTaskShrinkRequest) GetTargetGroupsShrink() *string {
	return s.TargetGroupsShrink
}

func (s *CreateMediaConvertTaskShrinkRequest) GetTargetsShrink() *string {
	return s.TargetsShrink
}

func (s *CreateMediaConvertTaskShrinkRequest) GetUserData() *string {
	return s.UserData
}

func (s *CreateMediaConvertTaskShrinkRequest) SetAlignmentIndex(v int32) *CreateMediaConvertTaskShrinkRequest {
	s.AlignmentIndex = &v
	return s
}

func (s *CreateMediaConvertTaskShrinkRequest) SetCredentialConfigShrink(v string) *CreateMediaConvertTaskShrinkRequest {
	s.CredentialConfigShrink = &v
	return s
}

func (s *CreateMediaConvertTaskShrinkRequest) SetNotificationShrink(v string) *CreateMediaConvertTaskShrinkRequest {
	s.NotificationShrink = &v
	return s
}

func (s *CreateMediaConvertTaskShrinkRequest) SetProjectName(v string) *CreateMediaConvertTaskShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateMediaConvertTaskShrinkRequest) SetSourcesShrink(v string) *CreateMediaConvertTaskShrinkRequest {
	s.SourcesShrink = &v
	return s
}

func (s *CreateMediaConvertTaskShrinkRequest) SetTagsShrink(v string) *CreateMediaConvertTaskShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *CreateMediaConvertTaskShrinkRequest) SetTargetGroupsShrink(v string) *CreateMediaConvertTaskShrinkRequest {
	s.TargetGroupsShrink = &v
	return s
}

func (s *CreateMediaConvertTaskShrinkRequest) SetTargetsShrink(v string) *CreateMediaConvertTaskShrinkRequest {
	s.TargetsShrink = &v
	return s
}

func (s *CreateMediaConvertTaskShrinkRequest) SetUserData(v string) *CreateMediaConvertTaskShrinkRequest {
	s.UserData = &v
	return s
}

func (s *CreateMediaConvertTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
