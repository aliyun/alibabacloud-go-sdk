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
	// The configurations of authorization chains. For more information, see [Use authorization chains to access resources of other entities](https://help.aliyun.com/document_detail/465340.html).
	CredentialConfig *CredentialConfig `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	// The notification settings. For information about the asynchronous notification format, see [Asynchronous message examples](https://help.aliyun.com/document_detail/2743997.html).
	Notification *Notification `json:"Notification,omitempty" xml:"Notification,omitempty"`
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
	Sources []*CreateImageToPDFTaskRequestSources `json:"Sources,omitempty" xml:"Sources,omitempty" type:"Repeated"`
	// The custom tags. You can search for or filter asynchronous tasks by custom tag.
	//
	// example:
	//
	// {
	//
	//       "User": "Jane"
	//
	// }
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
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
	return dara.Validate(s)
}

type CreateImageToPDFTaskRequestSources struct {
	// The rotation angle. Valid values:
	//
	// 	- 0 (default)
	//
	// 	- 90
	//
	// 	- 180
	//
	// 	- 270
	//
	// example:
	//
	// 90
	Rotate *int64 `json:"Rotate,omitempty" xml:"Rotate,omitempty"`
	// The OSS URI of the input image.
	//
	// Specify the OSS URI in the oss://${Bucket}/${Object} format, where `${Bucket}` is the name of the bucket in the same region as the current project and `${Object}` is the path of the object with the extension included.
	//
	// The operation supports the following image formats: JPG, JP2, PNG, TIFF, WebP, BMP, and SVG.
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
