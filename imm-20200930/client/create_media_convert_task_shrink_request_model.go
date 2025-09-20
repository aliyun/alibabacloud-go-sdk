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
	SetTargetsShrink(v string) *CreateMediaConvertTaskShrinkRequest
	GetTargetsShrink() *string
	SetUserData(v string) *CreateMediaConvertTaskShrinkRequest
	GetUserData() *string
}

type CreateMediaConvertTaskShrinkRequest struct {
	// The sequence number of the main media file in the concatenation list of media files. The main media file provides the default transcoding settings, such as the resolution and the frame rate, for videos and audios. Default value: `0`. A value of `0` specifies that the main media file is aligned with the first media file in the concatenation list.
	AlignmentIndex *int32 `json:"AlignmentIndex,omitempty" xml:"AlignmentIndex,omitempty"`
	// **If you have no special requirements, leave this parameter empty.**
	//
	// The authorization chain. For more information, see [Use authorization chains to access resources of other entities](https://help.aliyun.com/document_detail/465340.html).
	CredentialConfigShrink *string `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	// The notification settings. For more information, see "Notification". For information about the asynchronous notification format, see [Asynchronous notification format](https://help.aliyun.com/document_detail/2743997.html).
	NotificationShrink *string `json:"Notification,omitempty" xml:"Notification,omitempty"`
	// The name of the project. You can obtain the name of the project from the response of the [CreateProject](https://help.aliyun.com/document_detail/478153.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// immtest
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The source media files. If multiple files exist at the same time, the Concat feature is enabled. The video files are concatenated in the order of their URI inputs.
	//
	// This parameter is required.
	SourcesShrink *string `json:"Sources,omitempty" xml:"Sources,omitempty"`
	// The custom tags. You can search for or filter asynchronous tasks by custom tag.
	//
	// example:
	//
	// {"test":"val1"}
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The media processing tasks. You can specify multiple values for this parameter.
	//
	// This parameter is required.
	TargetsShrink *string `json:"Targets,omitempty" xml:"Targets,omitempty"`
	// The custom information, which is returned as asynchronous notifications to facilitate notification management in your system. The maximum information length is 2,048 bytes.
	//
	// example:
	//
	// {"ID": "user1","Name": "test-user1","Avatar": "http://example.com?id=user1"}
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
