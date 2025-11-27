// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFileCompressionTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCompressedFormat(v string) *CreateFileCompressionTaskRequest
	GetCompressedFormat() *string
	SetCredentialConfig(v *CredentialConfig) *CreateFileCompressionTaskRequest
	GetCredentialConfig() *CredentialConfig
	SetNotification(v *Notification) *CreateFileCompressionTaskRequest
	GetNotification() *Notification
	SetProjectName(v string) *CreateFileCompressionTaskRequest
	GetProjectName() *string
	SetSourceManifestURI(v string) *CreateFileCompressionTaskRequest
	GetSourceManifestURI() *string
	SetSources(v []*CreateFileCompressionTaskRequestSources) *CreateFileCompressionTaskRequest
	GetSources() []*CreateFileCompressionTaskRequestSources
	SetTargetURI(v string) *CreateFileCompressionTaskRequest
	GetTargetURI() *string
	SetUserData(v string) *CreateFileCompressionTaskRequest
	GetUserData() *string
}

type CreateFileCompressionTaskRequest struct {
	// The format of the output file.
	//
	// > Only the ZIP format is supported.
	//
	// example:
	//
	// zip
	CompressedFormat *string `json:"CompressedFormat,omitempty" xml:"CompressedFormat,omitempty"`
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
	// immtest
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The OSS URI of the inventory object that contains the objects to compress. The inventory object stores the objects to compress by using the same data structure of the Sources parameter in the JSON format. This parameter is suitable for specifying a large number of objects to compress.
	//
	// >  You must specify this parameter or the `Sources` parameter. The `URI` parameter is required and the `Alias` parameter is optional. You can specify up to 80,000 compression rule by using SourceManifestURI in one single call to the operation. The following line provides an example of the content within an inventory object.
	//
	//     [{"URI":"oss://<bucket>/<object>", "Alias":"/new-dir/new-name"}]
	//
	// example:
	//
	// oss://test-bucket/test-object.json
	SourceManifestURI *string `json:"SourceManifestURI,omitempty" xml:"SourceManifestURI,omitempty"`
	// The objects to be packed and packing rules.
	//
	// >  You must specify this parameter or the SourceManifestURI parameter. The Sources parameter can hold up to 100 packing rules. If you want to include more than 100 packing rules, use the SourceManifestURI parameter.
	Sources []*CreateFileCompressionTaskRequestSources `json:"Sources,omitempty" xml:"Sources,omitempty" type:"Repeated"`
	// The OSS URI of the package. The object name part in the URI is used as the name of the package. Example: `name.zip`.
	//
	// Specify the OSS URI in the oss://${Bucket}/${Object} format, where `${Bucket}` is the name of the bucket in the same region as the current project and `${Object}` is the path of the object with the extension included.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://test-bucket/test-target-object.zip
	TargetURI *string `json:"TargetURI,omitempty" xml:"TargetURI,omitempty"`
	// The custom information, which is returned in an asynchronous notification and facilitates notification management. The maximum length of the value is 2,048 bytes.
	//
	// example:
	//
	// {"ID": "testuid","Name": "test-user","Avatar": "http://test.com/testuid"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s CreateFileCompressionTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFileCompressionTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateFileCompressionTaskRequest) GetCompressedFormat() *string {
	return s.CompressedFormat
}

func (s *CreateFileCompressionTaskRequest) GetCredentialConfig() *CredentialConfig {
	return s.CredentialConfig
}

func (s *CreateFileCompressionTaskRequest) GetNotification() *Notification {
	return s.Notification
}

func (s *CreateFileCompressionTaskRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *CreateFileCompressionTaskRequest) GetSourceManifestURI() *string {
	return s.SourceManifestURI
}

func (s *CreateFileCompressionTaskRequest) GetSources() []*CreateFileCompressionTaskRequestSources {
	return s.Sources
}

func (s *CreateFileCompressionTaskRequest) GetTargetURI() *string {
	return s.TargetURI
}

func (s *CreateFileCompressionTaskRequest) GetUserData() *string {
	return s.UserData
}

func (s *CreateFileCompressionTaskRequest) SetCompressedFormat(v string) *CreateFileCompressionTaskRequest {
	s.CompressedFormat = &v
	return s
}

func (s *CreateFileCompressionTaskRequest) SetCredentialConfig(v *CredentialConfig) *CreateFileCompressionTaskRequest {
	s.CredentialConfig = v
	return s
}

func (s *CreateFileCompressionTaskRequest) SetNotification(v *Notification) *CreateFileCompressionTaskRequest {
	s.Notification = v
	return s
}

func (s *CreateFileCompressionTaskRequest) SetProjectName(v string) *CreateFileCompressionTaskRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateFileCompressionTaskRequest) SetSourceManifestURI(v string) *CreateFileCompressionTaskRequest {
	s.SourceManifestURI = &v
	return s
}

func (s *CreateFileCompressionTaskRequest) SetSources(v []*CreateFileCompressionTaskRequestSources) *CreateFileCompressionTaskRequest {
	s.Sources = v
	return s
}

func (s *CreateFileCompressionTaskRequest) SetTargetURI(v string) *CreateFileCompressionTaskRequest {
	s.TargetURI = &v
	return s
}

func (s *CreateFileCompressionTaskRequest) SetUserData(v string) *CreateFileCompressionTaskRequest {
	s.UserData = &v
	return s
}

func (s *CreateFileCompressionTaskRequest) Validate() error {
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

type CreateFileCompressionTaskRequestSources struct {
	// Specifies the path of the object in the package, or renames the object in the package.
	//
	// 	- Leave this parameter empty to retain the original directory structure of the object in the package. For example, if the object is stored at `oss://test-bucket/test-dir/test-object.doc` and you do not specify this parameter, the path of the object in the package is `/test-dir/test-object.doc`.
	//
	// 	- Rename the object. For example, if the object is stored at `oss://test-bucket/test-object.jpg` and you set the **Alias*	- parameter to `test-rename-object.jpg`, the name of the object in the package is `test-rename-object.jpg`.
	//
	// 	- Specify a different path for the object in the package. For example, if the directory to be packed is `oss://test-bucket/test-dir/` and you set the **Alias*	- parameter to `/new-dir/`, all objects in the directory are placed in the `/new-dir/` path in the package.
	//
	// 	- Set the parameter to `/` to remove the original directory structure.
	//
	// >  Duplicate object names may cause a failure in extracting the objects from the package, depending on the packing tool that you use. We recommend that you avoid using duplicate object names when you rename objects in the packing task.
	//
	// example:
	//
	// /new-dir/
	Alias *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	// The object matching rule. Valid values: `fullname` and `prefix`. Default value: `prefix`
	//
	// 	- `prefix`: matches objects by object name prefix.
	//
	// 	- `fullname`: exactly matches one single object by its full object name.
	//
	// example:
	//
	// fullname
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The OSS URI of the object or directory.
	//
	// Specify the OSS URI in the oss://${Bucket}/${Object} format, where `${Bucket}` is the name of the bucket in the same region as the current project and `${Object}` is a directory or object:
	//
	// When you pack a directory, `${Object}` is the path of the directory.
	//
	// 	- When you pack an object, `${Object}` is the path of the object with the extension included.
	//
	// example:
	//
	// oss://test-bucket/test-object
	URI *string `json:"URI,omitempty" xml:"URI,omitempty"`
}

func (s CreateFileCompressionTaskRequestSources) String() string {
	return dara.Prettify(s)
}

func (s CreateFileCompressionTaskRequestSources) GoString() string {
	return s.String()
}

func (s *CreateFileCompressionTaskRequestSources) GetAlias() *string {
	return s.Alias
}

func (s *CreateFileCompressionTaskRequestSources) GetMode() *string {
	return s.Mode
}

func (s *CreateFileCompressionTaskRequestSources) GetURI() *string {
	return s.URI
}

func (s *CreateFileCompressionTaskRequestSources) SetAlias(v string) *CreateFileCompressionTaskRequestSources {
	s.Alias = &v
	return s
}

func (s *CreateFileCompressionTaskRequestSources) SetMode(v string) *CreateFileCompressionTaskRequestSources {
	s.Mode = &v
	return s
}

func (s *CreateFileCompressionTaskRequestSources) SetURI(v string) *CreateFileCompressionTaskRequestSources {
	s.URI = &v
	return s
}

func (s *CreateFileCompressionTaskRequestSources) Validate() error {
	return dara.Validate(s)
}
