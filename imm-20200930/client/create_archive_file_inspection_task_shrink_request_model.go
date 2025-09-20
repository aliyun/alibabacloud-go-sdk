// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateArchiveFileInspectionTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCredentialConfigShrink(v string) *CreateArchiveFileInspectionTaskShrinkRequest
	GetCredentialConfigShrink() *string
	SetNotificationShrink(v string) *CreateArchiveFileInspectionTaskShrinkRequest
	GetNotificationShrink() *string
	SetPassword(v string) *CreateArchiveFileInspectionTaskShrinkRequest
	GetPassword() *string
	SetProjectName(v string) *CreateArchiveFileInspectionTaskShrinkRequest
	GetProjectName() *string
	SetSourceURI(v string) *CreateArchiveFileInspectionTaskShrinkRequest
	GetSourceURI() *string
	SetUserData(v string) *CreateArchiveFileInspectionTaskShrinkRequest
	GetUserData() *string
}

type CreateArchiveFileInspectionTaskShrinkRequest struct {
	// **If you have no special requirements, leave this parameter empty.**
	//
	// The configurations of authorization chains. For more information, see [Use authorization chains to access resources of other entities](https://help.aliyun.com/document_detail/465340.html).
	CredentialConfigShrink *string `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	// The notification settings. For information about the asynchronous notification format, see [Asynchronous message examples](https://help.aliyun.com/document_detail/2743997.html).
	NotificationShrink *string `json:"Notification,omitempty" xml:"Notification,omitempty"`
	// The password that protects the package. If the package is password-protected, you must provide the password to view the contents of the package.
	//
	// example:
	//
	// 123456
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The name of the project.[](~~478153~~)
	//
	// This parameter is required.
	//
	// example:
	//
	// immtest
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The URI of the package.
	//
	// Specify the OSS URI in the oss://${Bucket}/${Object} format, where `${Bucket}` is the name of the bucket in the same region as the current project and `${Object}` is the path of the object with the extension included.
	//
	// example:
	//
	// oss://imm-apitest-fxf2/name.zip
	SourceURI *string `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
	// The custom information, which is returned in an asynchronous notification and facilitates notification management. The maximum length of the value is 2,048 bytes.
	//
	// example:
	//
	// {"ID": "user1","Name": "test-user1","Avatar": "http://example.com?id=user1"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s CreateArchiveFileInspectionTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateArchiveFileInspectionTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateArchiveFileInspectionTaskShrinkRequest) GetCredentialConfigShrink() *string {
	return s.CredentialConfigShrink
}

func (s *CreateArchiveFileInspectionTaskShrinkRequest) GetNotificationShrink() *string {
	return s.NotificationShrink
}

func (s *CreateArchiveFileInspectionTaskShrinkRequest) GetPassword() *string {
	return s.Password
}

func (s *CreateArchiveFileInspectionTaskShrinkRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *CreateArchiveFileInspectionTaskShrinkRequest) GetSourceURI() *string {
	return s.SourceURI
}

func (s *CreateArchiveFileInspectionTaskShrinkRequest) GetUserData() *string {
	return s.UserData
}

func (s *CreateArchiveFileInspectionTaskShrinkRequest) SetCredentialConfigShrink(v string) *CreateArchiveFileInspectionTaskShrinkRequest {
	s.CredentialConfigShrink = &v
	return s
}

func (s *CreateArchiveFileInspectionTaskShrinkRequest) SetNotificationShrink(v string) *CreateArchiveFileInspectionTaskShrinkRequest {
	s.NotificationShrink = &v
	return s
}

func (s *CreateArchiveFileInspectionTaskShrinkRequest) SetPassword(v string) *CreateArchiveFileInspectionTaskShrinkRequest {
	s.Password = &v
	return s
}

func (s *CreateArchiveFileInspectionTaskShrinkRequest) SetProjectName(v string) *CreateArchiveFileInspectionTaskShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateArchiveFileInspectionTaskShrinkRequest) SetSourceURI(v string) *CreateArchiveFileInspectionTaskShrinkRequest {
	s.SourceURI = &v
	return s
}

func (s *CreateArchiveFileInspectionTaskShrinkRequest) SetUserData(v string) *CreateArchiveFileInspectionTaskShrinkRequest {
	s.UserData = &v
	return s
}

func (s *CreateArchiveFileInspectionTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
