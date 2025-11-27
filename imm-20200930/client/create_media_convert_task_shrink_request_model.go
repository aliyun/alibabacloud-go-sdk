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
	// When performing media concatenation, the index of the primary media file (which provides the default transcoding parameters for `Video` and `Audio`, including resolution, frame rate, etc.) in the concatenation list. The default value is 0 (aligning with the first media file in the concatenation list).
	//
	// example:
	//
	// 0
	AlignmentIndex *int32 `json:"AlignmentIndex,omitempty" xml:"AlignmentIndex,omitempty"`
	// **If there are no special requirements, please leave this blank.**
	//
	// Chain authorization configuration. For more information, see [Using Chain Authorization to Access Other Entity Resources](https://help.aliyun.com/document_detail/465340.html).
	CredentialConfigShrink *string `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	// Notification configuration. For details, click Notification. The format of asynchronous notification messages can be found in [Asynchronous Notification Message Format](https://help.aliyun.com/document_detail/2743997.html).
	NotificationShrink *string `json:"Notification,omitempty" xml:"Notification,omitempty"`
	// The name of the project. For how to obtain it, see [Creating a Project](https://help.aliyun.com/document_detail/478153.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// A list of media files. If the list contains more than one element, it indicates that the Concat (concatenation) function is enabled. The Concat order follows the sequence of the input video file URIs.
	//
	// This parameter is required.
	SourcesShrink *string `json:"Sources,omitempty" xml:"Sources,omitempty"`
	// Custom tags used for searching and filtering asynchronous tasks.
	//
	// example:
	//
	// {"test":"val1"}
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// List of media processing tasks, supporting multiple task configurations.
	//
	// This parameter is required.
	TargetsShrink *string `json:"Targets,omitempty" xml:"Targets,omitempty"`
	// User-defined information that will be returned in asynchronous message notifications, used for convenient association and processing within your system. The maximum length is 2048 bytes.
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
