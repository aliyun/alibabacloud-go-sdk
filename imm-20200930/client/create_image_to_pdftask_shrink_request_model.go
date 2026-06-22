// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateImageToPDFTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCredentialConfigShrink(v string) *CreateImageToPDFTaskShrinkRequest
	GetCredentialConfigShrink() *string
	SetNotificationShrink(v string) *CreateImageToPDFTaskShrinkRequest
	GetNotificationShrink() *string
	SetProjectName(v string) *CreateImageToPDFTaskShrinkRequest
	GetProjectName() *string
	SetSourcesShrink(v string) *CreateImageToPDFTaskShrinkRequest
	GetSourcesShrink() *string
	SetTagsShrink(v string) *CreateImageToPDFTaskShrinkRequest
	GetTagsShrink() *string
	SetTargetURI(v string) *CreateImageToPDFTaskShrinkRequest
	GetTargetURI() *string
	SetUserData(v string) *CreateImageToPDFTaskShrinkRequest
	GetUserData() *string
}

type CreateImageToPDFTaskShrinkRequest struct {
	// **If you have no special requirements, leave this parameter empty.**
	//
	// The chained authorization configuration. For more information, see [Use chained authorization to access resources of other entities](https://help.aliyun.com/document_detail/465340.html).
	CredentialConfigShrink *string `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	// The message notification configuration. For more information, click Notification. For the format of asynchronous notification messages, see [Asynchronous notification message format](https://help.aliyun.com/document_detail/2743997.html).
	NotificationShrink *string `json:"Notification,omitempty" xml:"Notification,omitempty"`
	// The name of the project. For more information, see [Create a project](https://help.aliyun.com/document_detail/478153.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// A list of input images. The images are converted in the order of their URIs in this list.
	//
	// This parameter is required.
	SourcesShrink *string `json:"Sources,omitempty" xml:"Sources,omitempty"`
	// Custom tags used to search for and filter asynchronous tasks.
	//
	// example:
	//
	// {
	//
	//       "User": "Jane"
	//
	// }
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The OSS address where the output PDF file is stored.
	//
	// The address must be in the \\`oss\\://${bucketname}/${objectname}\\` format. \\`${bucketname}\\` must be an OSS bucket in the same region as the project. \\`${objectname}\\` must be the path of the file, including the file name.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://examplebucket/outputDocument.pdf
	TargetURI *string `json:"TargetURI,omitempty" xml:"TargetURI,omitempty"`
	// Custom user information that is returned in the asynchronous notification message. This helps you associate the notification message with your system. The maximum length is 2048 bytes.
	//
	// example:
	//
	// test-data
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s CreateImageToPDFTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateImageToPDFTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateImageToPDFTaskShrinkRequest) GetCredentialConfigShrink() *string {
	return s.CredentialConfigShrink
}

func (s *CreateImageToPDFTaskShrinkRequest) GetNotificationShrink() *string {
	return s.NotificationShrink
}

func (s *CreateImageToPDFTaskShrinkRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *CreateImageToPDFTaskShrinkRequest) GetSourcesShrink() *string {
	return s.SourcesShrink
}

func (s *CreateImageToPDFTaskShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *CreateImageToPDFTaskShrinkRequest) GetTargetURI() *string {
	return s.TargetURI
}

func (s *CreateImageToPDFTaskShrinkRequest) GetUserData() *string {
	return s.UserData
}

func (s *CreateImageToPDFTaskShrinkRequest) SetCredentialConfigShrink(v string) *CreateImageToPDFTaskShrinkRequest {
	s.CredentialConfigShrink = &v
	return s
}

func (s *CreateImageToPDFTaskShrinkRequest) SetNotificationShrink(v string) *CreateImageToPDFTaskShrinkRequest {
	s.NotificationShrink = &v
	return s
}

func (s *CreateImageToPDFTaskShrinkRequest) SetProjectName(v string) *CreateImageToPDFTaskShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateImageToPDFTaskShrinkRequest) SetSourcesShrink(v string) *CreateImageToPDFTaskShrinkRequest {
	s.SourcesShrink = &v
	return s
}

func (s *CreateImageToPDFTaskShrinkRequest) SetTagsShrink(v string) *CreateImageToPDFTaskShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *CreateImageToPDFTaskShrinkRequest) SetTargetURI(v string) *CreateImageToPDFTaskShrinkRequest {
	s.TargetURI = &v
	return s
}

func (s *CreateImageToPDFTaskShrinkRequest) SetUserData(v string) *CreateImageToPDFTaskShrinkRequest {
	s.UserData = &v
	return s
}

func (s *CreateImageToPDFTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
