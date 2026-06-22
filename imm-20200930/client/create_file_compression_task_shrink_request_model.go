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
	// The compression format for file packaging.
	//
	// > Currently, only the zip format is supported.
	//
	// example:
	//
	// zip
	CompressedFormat *string `json:"CompressedFormat,omitempty" xml:"CompressedFormat,omitempty"`
	// **If you do not have special requirements, leave this parameter empty.**
	//
	// The chained authorization configuration. This parameter is not required. For more information, see [Use chained authorization to access resources of other entities](https://help.aliyun.com/document_detail/465340.html).
	CredentialConfigShrink *string `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	// The message notification configuration. For more information, see the Notification data type. For the format of asynchronous notification messages, see [Asynchronous notification message format](https://help.aliyun.com/document_detail/2743997.html).
	//
	// > IMM API callbacks do not currently support specifying a webhook address. Use MNS instead.
	NotificationShrink *string `json:"Notification,omitempty" xml:"Notification,omitempty"`
	// The name of the project. For more information, see [Create a project](https://help.aliyun.com/document_detail/478153.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The address where the file manifest is stored. The file manifest stores the \\`Sources\\` structure in JSON format on OSS. This is suitable for scenarios with many files to package.
	//
	// > Specify either this parameter or `Sources`. In the manifest file, the `URI` parameter is required and the `Alias` parameter is optional. \\`SourceManifestURI\\` supports up to 80,000 packaging rules.
	//
	// > 	Warning: When you save the content to OSS, specify the OSS address of the file for this parameter.
	//
	// The following is an example of the file structure:
	//
	// ```
	//
	// [{"URI":"oss://<bucket>/<object>", "Alias":"/new-dir/new-name"}]
	//
	// ```
	//
	// example:
	//
	// oss://test-bucket/test-object.json
	SourceManifestURI *string `json:"SourceManifestURI,omitempty" xml:"SourceManifestURI,omitempty"`
	// A list of files to package and their packaging rules.
	//
	// > Specify either this parameter or \\`SourceManifestURI\\`. \\`Sources\\` supports a maximum of 100 packaging rules.
	//
	// > 	Warning: If you have more than 100 packaging rules, use the \\`SourceManifestURI\\` parameter.
	SourcesShrink *string `json:"Sources,omitempty" xml:"Sources,omitempty"`
	// The OSS address of the output file. The compressed file is named after the file name in this path, such as `name.zip`.
	//
	// The OSS address must be in the \\`oss\\://${Bucket}/${Object}\\` format. \\`${Bucket}\\` is the name of the OSS bucket that is in the same region as the current project. \\`${Object}\\` is the full path of the file, including the file name extension.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://test-bucket/test-target-object.zip
	TargetURI *string `json:"TargetURI,omitempty" xml:"TargetURI,omitempty"`
	// Custom user data. This data is returned in the asynchronous notification message, which helps you associate the notification with your internal system. The maximum length is 2,048 bytes.
	//
	// example:
	//
	// test-data
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
