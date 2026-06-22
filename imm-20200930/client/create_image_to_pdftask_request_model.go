// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateImageToPDFTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCredentialConfig(v *CredentialConfig) *CreateImageToPDFTaskRequest
	GetCredentialConfig() *CredentialConfig
	SetNotification(v *Notification) *CreateImageToPDFTaskRequest
	GetNotification() *Notification
	SetProjectName(v string) *CreateImageToPDFTaskRequest
	GetProjectName() *string
	SetSources(v []*CreateImageToPDFTaskRequestSources) *CreateImageToPDFTaskRequest
	GetSources() []*CreateImageToPDFTaskRequestSources
	SetTags(v map[string]interface{}) *CreateImageToPDFTaskRequest
	GetTags() map[string]interface{}
	SetTargetURI(v string) *CreateImageToPDFTaskRequest
	GetTargetURI() *string
	SetUserData(v string) *CreateImageToPDFTaskRequest
	GetUserData() *string
}

type CreateImageToPDFTaskRequest struct {
	// **If you have no special requirements, leave this parameter empty.**
	//
	// The chained authorization configuration. For more information, see [Use chained authorization to access resources of other entities](https://help.aliyun.com/document_detail/465340.html).
	CredentialConfig *CredentialConfig `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	// The message notification configuration. For more information, click Notification. For the format of asynchronous notification messages, see [Asynchronous notification message format](https://help.aliyun.com/document_detail/2743997.html).
	Notification *Notification `json:"Notification,omitempty" xml:"Notification,omitempty"`
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
	Sources []*CreateImageToPDFTaskRequestSources `json:"Sources,omitempty" xml:"Sources,omitempty" type:"Repeated"`
	// Custom tags used to search for and filter asynchronous tasks.
	//
	// example:
	//
	// {
	//
	//       "User": "Jane"
	//
	// }
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
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

func (s CreateImageToPDFTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateImageToPDFTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateImageToPDFTaskRequest) GetCredentialConfig() *CredentialConfig {
	return s.CredentialConfig
}

func (s *CreateImageToPDFTaskRequest) GetNotification() *Notification {
	return s.Notification
}

func (s *CreateImageToPDFTaskRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *CreateImageToPDFTaskRequest) GetSources() []*CreateImageToPDFTaskRequestSources {
	return s.Sources
}

func (s *CreateImageToPDFTaskRequest) GetTags() map[string]interface{} {
	return s.Tags
}

func (s *CreateImageToPDFTaskRequest) GetTargetURI() *string {
	return s.TargetURI
}

func (s *CreateImageToPDFTaskRequest) GetUserData() *string {
	return s.UserData
}

func (s *CreateImageToPDFTaskRequest) SetCredentialConfig(v *CredentialConfig) *CreateImageToPDFTaskRequest {
	s.CredentialConfig = v
	return s
}

func (s *CreateImageToPDFTaskRequest) SetNotification(v *Notification) *CreateImageToPDFTaskRequest {
	s.Notification = v
	return s
}

func (s *CreateImageToPDFTaskRequest) SetProjectName(v string) *CreateImageToPDFTaskRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateImageToPDFTaskRequest) SetSources(v []*CreateImageToPDFTaskRequestSources) *CreateImageToPDFTaskRequest {
	s.Sources = v
	return s
}

func (s *CreateImageToPDFTaskRequest) SetTags(v map[string]interface{}) *CreateImageToPDFTaskRequest {
	s.Tags = v
	return s
}

func (s *CreateImageToPDFTaskRequest) SetTargetURI(v string) *CreateImageToPDFTaskRequest {
	s.TargetURI = &v
	return s
}

func (s *CreateImageToPDFTaskRequest) SetUserData(v string) *CreateImageToPDFTaskRequest {
	s.UserData = &v
	return s
}

func (s *CreateImageToPDFTaskRequest) Validate() error {
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
	if s.Sources != nil {
		for _, item := range s.Sources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateImageToPDFTaskRequestSources struct {
	// The rotation angle of the image in degrees. Valid values:
	//
	// - 0 (default)
	//
	// - 90
	//
	// - 180
	//
	// - 270
	//
	// example:
	//
	// 90
	Rotate *int64 `json:"Rotate,omitempty" xml:"Rotate,omitempty"`
	// The OSS address of the source image.
	//
	// The address must be in the \\`oss\\://${Bucket}/${Object}\\` format. \\``${Bucket}`\\` must be an OSS bucket in the same region as the project. \\``${Object}`\\` must be the full path of the file, including its file name extension.
	//
	// Supported formats: JPG, JP2, PNG, TIFF, WebP, BMP, and SVG.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://examplebucket/sampleobject.jpg
	URI *string `json:"URI,omitempty" xml:"URI,omitempty"`
}

func (s CreateImageToPDFTaskRequestSources) String() string {
	return dara.Prettify(s)
}

func (s CreateImageToPDFTaskRequestSources) GoString() string {
	return s.String()
}

func (s *CreateImageToPDFTaskRequestSources) GetRotate() *int64 {
	return s.Rotate
}

func (s *CreateImageToPDFTaskRequestSources) GetURI() *string {
	return s.URI
}

func (s *CreateImageToPDFTaskRequestSources) SetRotate(v int64) *CreateImageToPDFTaskRequestSources {
	s.Rotate = &v
	return s
}

func (s *CreateImageToPDFTaskRequestSources) SetURI(v string) *CreateImageToPDFTaskRequestSources {
	s.URI = &v
	return s
}

func (s *CreateImageToPDFTaskRequestSources) Validate() error {
	return dara.Validate(s)
}
