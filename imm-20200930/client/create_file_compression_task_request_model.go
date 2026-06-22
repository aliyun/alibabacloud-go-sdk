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
	CredentialConfig *CredentialConfig `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	// The message notification configuration. For more information, see the Notification data type. For the format of asynchronous notification messages, see [Asynchronous notification message format](https://help.aliyun.com/document_detail/2743997.html).
	//
	// > IMM API callbacks do not currently support specifying a webhook address. Use MNS instead.
	Notification *Notification `json:"Notification,omitempty" xml:"Notification,omitempty"`
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
	Sources []*CreateFileCompressionTaskRequestSources `json:"Sources,omitempty" xml:"Sources,omitempty" type:"Repeated"`
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
	// Specifies a new path or name for the source file within the output compressed file.
	//
	// - If you do not specify this parameter, the source directory structure is preserved. For example, if the source file is at `oss://test-bucket/test-dir/test-object.doc`, the path of the file in the compressed file is `/test-dir/test-object.doc`.
	//
	// - Rename the file. For example, if the source file is at `oss://test-bucket/test-object.jpg` and you set this parameter to `/test-rename-object.jpg`, the file in the compressed file is named `test-rename-object.jpg`.
	//
	// - Specify a new path for the source file in the compressed file. For example, if the source directory is `oss://test-bucket/test-dir/` and you set this parameter to `/new-dir/`, all files in the source directory are compressed into the `/new-dir/` path.
	//
	// - Set the value to `/` to remove the source directory structure. All files are placed directly in the root directory of the compressed file, and the original directory structure is not preserved.
	//
	// - Specify both a path and a file name. The file is renamed and moved to the specified path. For example, if you set this parameter to `/new-dir/alias.doc`, the file is renamed to `alias.doc` and placed in the `/new-dir/` path of the compressed file.
	//
	// > 	- Avoid creating files with duplicate names during the renaming process. If duplicate names exist, you may not be able to decompress the file in the compressed package. This depends on the decompression program you use.
	//
	// >
	//
	// > 	- Format requirement: The value must start with a forward slash (\\`/\\`), such as `/new-dir/alias.doc`.
	//
	// example:
	//
	// /new-dir/
	Alias *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	// The pattern matching mode for the packaging rule. Valid values include `prefix` (prefix matching) and `fullname` (exact matching). The default value is `prefix`.
	//
	// - `prefix`: In this mode, all files that match the prefix are packaged.
	//
	// - `fullname`: In this mode, only the file that exactly matches the rule is packaged.
	//
	// example:
	//
	// fullname
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The OSS address of the directory or file to package.
	//
	// The OSS address must be in the \\`oss\\://${Bucket}/${Object}\\` format. \\`${Bucket}\\` is the name of the OSS bucket that is in the same region as the current project. \\`${Object}\\` is described as follows:
	//
	// - To package a directory, \\`${Object}\\` is the directory name.
	//
	// - To package a file, \\`${Object}\\` is the full path of the file, including the file name extension.
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
