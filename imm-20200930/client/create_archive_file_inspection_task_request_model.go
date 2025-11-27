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
	// **If you have no special requirements, leave this parameter empty.**
	//
	// The configurations of authorization chains. For more information, see [Use authorization chains to access resources of other entities](https://help.aliyun.com/document_detail/465340.html).
	CredentialConfig *CredentialConfig `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	// The notification settings. For information about the asynchronous notification format, see [Asynchronous message examples](https://help.aliyun.com/document_detail/2743997.html).
	//
	// >  The IMM operation does not support a callback URL. We recommend that you use Simple Message Queue (SMQ) to receive notifications.
	Notification *Notification `json:"Notification,omitempty" xml:"Notification,omitempty"`
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
