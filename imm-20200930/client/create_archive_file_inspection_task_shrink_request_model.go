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
	// **Leave this parameter empty if you do not have special requirements.**
	//
	// The configuration for chained authorization. This parameter is not required. For more information, see [Use chained authorization to access resources of other entities](https://help.aliyun.com/document_detail/465340.html).
	CredentialConfigShrink *string `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	// The notification configuration. For more information, see Notification. For the format of asynchronous notification messages, see [Asynchronous notification message format](https://help.aliyun.com/document_detail/2743997.html).
	//
	// > Currently, API callbacks in IMM do not support custom webhook addresses. Use MNS instead.
	NotificationShrink *string `json:"Notification,omitempty" xml:"Notification,omitempty"`
	// The password of the compressed file. If the file is encrypted, provide the password to inspect its contents.
	//
	// example:
	//
	// 123456
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The project name. For more information, see [Create a project](https://help.aliyun.com/document_detail/478153.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The location of the compressed file.
	//
	// The Object Storage Service (OSS) URI must be in the oss\\://${Bucket}/${Object} format. In this format, `${Bucket}` is the name of the OSS bucket that is in the same region as the current project, and `${Object}` is the full path of the file, including the file name extension.
	//
	// example:
	//
	// oss://bucket/test-object.zip
	SourceURI *string `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
	// Custom information that is returned in the asynchronous notification message. You can use this information to associate the notification message with your services. The maximum length is 2,048 bytes.
	//
	// example:
	//
	// test-data
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
