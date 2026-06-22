// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateArchiveFileInspectionTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCredentialConfig(v *CredentialConfig) *CreateArchiveFileInspectionTaskRequest
	GetCredentialConfig() *CredentialConfig
	SetNotification(v *Notification) *CreateArchiveFileInspectionTaskRequest
	GetNotification() *Notification
	SetPassword(v string) *CreateArchiveFileInspectionTaskRequest
	GetPassword() *string
	SetProjectName(v string) *CreateArchiveFileInspectionTaskRequest
	GetProjectName() *string
	SetSourceURI(v string) *CreateArchiveFileInspectionTaskRequest
	GetSourceURI() *string
	SetUserData(v string) *CreateArchiveFileInspectionTaskRequest
	GetUserData() *string
}

type CreateArchiveFileInspectionTaskRequest struct {
	// **Leave this parameter empty if you do not have special requirements.**
	//
	// The configuration for chained authorization. This parameter is not required. For more information, see [Use chained authorization to access resources of other entities](https://help.aliyun.com/document_detail/465340.html).
	CredentialConfig *CredentialConfig `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	// The notification configuration. For more information, see Notification. For the format of asynchronous notification messages, see [Asynchronous notification message format](https://help.aliyun.com/document_detail/2743997.html).
	//
	// > Currently, API callbacks in IMM do not support custom webhook addresses. Use MNS instead.
	Notification *Notification `json:"Notification,omitempty" xml:"Notification,omitempty"`
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

func (s CreateArchiveFileInspectionTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateArchiveFileInspectionTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateArchiveFileInspectionTaskRequest) GetCredentialConfig() *CredentialConfig {
	return s.CredentialConfig
}

func (s *CreateArchiveFileInspectionTaskRequest) GetNotification() *Notification {
	return s.Notification
}

func (s *CreateArchiveFileInspectionTaskRequest) GetPassword() *string {
	return s.Password
}

func (s *CreateArchiveFileInspectionTaskRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *CreateArchiveFileInspectionTaskRequest) GetSourceURI() *string {
	return s.SourceURI
}

func (s *CreateArchiveFileInspectionTaskRequest) GetUserData() *string {
	return s.UserData
}

func (s *CreateArchiveFileInspectionTaskRequest) SetCredentialConfig(v *CredentialConfig) *CreateArchiveFileInspectionTaskRequest {
	s.CredentialConfig = v
	return s
}

func (s *CreateArchiveFileInspectionTaskRequest) SetNotification(v *Notification) *CreateArchiveFileInspectionTaskRequest {
	s.Notification = v
	return s
}

func (s *CreateArchiveFileInspectionTaskRequest) SetPassword(v string) *CreateArchiveFileInspectionTaskRequest {
	s.Password = &v
	return s
}

func (s *CreateArchiveFileInspectionTaskRequest) SetProjectName(v string) *CreateArchiveFileInspectionTaskRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateArchiveFileInspectionTaskRequest) SetSourceURI(v string) *CreateArchiveFileInspectionTaskRequest {
	s.SourceURI = &v
	return s
}

func (s *CreateArchiveFileInspectionTaskRequest) SetUserData(v string) *CreateArchiveFileInspectionTaskRequest {
	s.UserData = &v
	return s
}

func (s *CreateArchiveFileInspectionTaskRequest) Validate() error {
	if s.CredentialConfig != nil {
		if err := s.CredentialConfig.Validate(); err != nil {
			return err
		}
	}
	if s.Notification != nil {
		if err := s.Notification.Validate(); err != nil {
			return err
		}
	}
	return nil
}
