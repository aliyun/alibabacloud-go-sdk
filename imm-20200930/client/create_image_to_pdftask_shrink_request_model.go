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
	// The configurations of authorization chains. For more information, see [Use authorization chains to access resources of other entities](https://help.aliyun.com/document_detail/465340.html).
	CredentialConfigShrink *string `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	// The notification settings. For information about the asynchronous notification format, see [Asynchronous message examples](https://help.aliyun.com/document_detail/2743997.html).
	NotificationShrink *string `json:"Notification,omitempty" xml:"Notification,omitempty"`
	// The name of the project.[](~~478153~~)
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The list of images. The sequence of image URIs in the list determines the order in which they are converted.
	//
	// This parameter is required.
	SourcesShrink *string `json:"Sources,omitempty" xml:"Sources,omitempty"`
	// The custom tags. You can search for or filter asynchronous tasks by custom tag.
	//
	// example:
	//
	// {
	//
	//       "User": "Jane"
	//
	// }
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The OSS URI of the output file.
	//
	// Specify the OSS URI in the oss://${bucketname}/${objectname} format, where ${bucketname} is the name of the bucket in the same region as the current project and ${objectname} is the path of the object with the extension included.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://examplebucket/outputDocument.pdf
	TargetURI *string `json:"TargetURI,omitempty" xml:"TargetURI,omitempty"`
	// The custom information, which is returned in an asynchronous notification and facilitates notification management. The maximum length of the value is 2,048 bytes.
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
