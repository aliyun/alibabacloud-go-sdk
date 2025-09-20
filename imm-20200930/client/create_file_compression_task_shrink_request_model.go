// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFileCompressionTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCompressedFormat(v string) *CreateFileCompressionTaskShrinkRequest
	GetCompressedFormat() *string
	SetCredentialConfigShrink(v string) *CreateFileCompressionTaskShrinkRequest
	GetCredentialConfigShrink() *string
	SetNotificationShrink(v string) *CreateFileCompressionTaskShrinkRequest
	GetNotificationShrink() *string
	SetProjectName(v string) *CreateFileCompressionTaskShrinkRequest
	GetProjectName() *string
	SetSourceManifestURI(v string) *CreateFileCompressionTaskShrinkRequest
	GetSourceManifestURI() *string
	SetSourcesShrink(v string) *CreateFileCompressionTaskShrinkRequest
	GetSourcesShrink() *string
	SetTargetURI(v string) *CreateFileCompressionTaskShrinkRequest
	GetTargetURI() *string
	SetUserData(v string) *CreateFileCompressionTaskShrinkRequest
	GetUserData() *string
}

type CreateFileCompressionTaskShrinkRequest struct {
	// The format of the package. Default value: zip.
	//
	// >  Only the ZIP format is supported.
	//
	// example:
	//
	// zip
	CompressedFormat *string `json:"CompressedFormat,omitempty" xml:"CompressedFormat,omitempty"`
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
	SourcesShrink *string `json:"Sources,omitempty" xml:"Sources,omitempty"`
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

func (s CreateFileCompressionTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFileCompressionTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateFileCompressionTaskShrinkRequest) GetCompressedFormat() *string {
	return s.CompressedFormat
}

func (s *CreateFileCompressionTaskShrinkRequest) GetCredentialConfigShrink() *string {
	return s.CredentialConfigShrink
}

func (s *CreateFileCompressionTaskShrinkRequest) GetNotificationShrink() *string {
	return s.NotificationShrink
}

func (s *CreateFileCompressionTaskShrinkRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *CreateFileCompressionTaskShrinkRequest) GetSourceManifestURI() *string {
	return s.SourceManifestURI
}

func (s *CreateFileCompressionTaskShrinkRequest) GetSourcesShrink() *string {
	return s.SourcesShrink
}

func (s *CreateFileCompressionTaskShrinkRequest) GetTargetURI() *string {
	return s.TargetURI
}

func (s *CreateFileCompressionTaskShrinkRequest) GetUserData() *string {
	return s.UserData
}

func (s *CreateFileCompressionTaskShrinkRequest) SetCompressedFormat(v string) *CreateFileCompressionTaskShrinkRequest {
	s.CompressedFormat = &v
	return s
}

func (s *CreateFileCompressionTaskShrinkRequest) SetCredentialConfigShrink(v string) *CreateFileCompressionTaskShrinkRequest {
	s.CredentialConfigShrink = &v
	return s
}

func (s *CreateFileCompressionTaskShrinkRequest) SetNotificationShrink(v string) *CreateFileCompressionTaskShrinkRequest {
	s.NotificationShrink = &v
	return s
}

func (s *CreateFileCompressionTaskShrinkRequest) SetProjectName(v string) *CreateFileCompressionTaskShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateFileCompressionTaskShrinkRequest) SetSourceManifestURI(v string) *CreateFileCompressionTaskShrinkRequest {
	s.SourceManifestURI = &v
	return s
}

func (s *CreateFileCompressionTaskShrinkRequest) SetSourcesShrink(v string) *CreateFileCompressionTaskShrinkRequest {
	s.SourcesShrink = &v
	return s
}

func (s *CreateFileCompressionTaskShrinkRequest) SetTargetURI(v string) *CreateFileCompressionTaskShrinkRequest {
	s.TargetURI = &v
	return s
}

func (s *CreateFileCompressionTaskShrinkRequest) SetUserData(v string) *CreateFileCompressionTaskShrinkRequest {
	s.UserData = &v
	return s
}

func (s *CreateFileCompressionTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
