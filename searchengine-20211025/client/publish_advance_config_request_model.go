// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishAdvanceConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDesc(v string) *PublishAdvanceConfigRequest
	GetDesc() *string
	SetFiles(v []*PublishAdvanceConfigRequestFiles) *PublishAdvanceConfigRequest
	GetFiles() []*PublishAdvanceConfigRequestFiles
}

type PublishAdvanceConfigRequest struct {
	// The description of the advanced configuration.
	//
	// example:
	//
	// Custom configuration
	Desc *string `json:"desc,omitempty" xml:"desc,omitempty"`
	// The files.
	Files []*PublishAdvanceConfigRequestFiles `json:"files,omitempty" xml:"files,omitempty" type:"Repeated"`
}

func (s PublishAdvanceConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s PublishAdvanceConfigRequest) GoString() string {
	return s.String()
}

func (s *PublishAdvanceConfigRequest) GetDesc() *string {
	return s.Desc
}

func (s *PublishAdvanceConfigRequest) GetFiles() []*PublishAdvanceConfigRequestFiles {
	return s.Files
}

func (s *PublishAdvanceConfigRequest) SetDesc(v string) *PublishAdvanceConfigRequest {
	s.Desc = &v
	return s
}

func (s *PublishAdvanceConfigRequest) SetFiles(v []*PublishAdvanceConfigRequestFiles) *PublishAdvanceConfigRequest {
	s.Files = v
	return s
}

func (s *PublishAdvanceConfigRequest) Validate() error {
	return dara.Validate(s)
}

type PublishAdvanceConfigRequestFiles struct {
	// The information about the advanced configuration.
	Config *PublishAdvanceConfigRequestFilesConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
	// The directory name.
	//
	// example:
	//
	// /clusters
	DirName *string `json:"dirName,omitempty" xml:"dirName,omitempty"`
	// The file name.
	//
	// example:
	//
	// vector_question_schema.json
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// The operation type. Valid values: UPDATE and DELETE. Default value: UPDATE.
	//
	// example:
	//
	// UPDATE
	OperateType *string `json:"operateType,omitempty" xml:"operateType,omitempty"`
	// The path of the Object Storage Service (OSS) object.
	//
	// example:
	//
	// oss://opensearch/test.json
	OssPath *string `json:"ossPath,omitempty" xml:"ossPath,omitempty"`
	// The path of the parent directory.
	//
	// example:
	//
	// /
	ParentFullPath *string `json:"parentFullPath,omitempty" xml:"parentFullPath,omitempty"`
}

func (s PublishAdvanceConfigRequestFiles) String() string {
	return dara.Prettify(s)
}

func (s PublishAdvanceConfigRequestFiles) GoString() string {
	return s.String()
}

func (s *PublishAdvanceConfigRequestFiles) GetConfig() *PublishAdvanceConfigRequestFilesConfig {
	return s.Config
}

func (s *PublishAdvanceConfigRequestFiles) GetDirName() *string {
	return s.DirName
}

func (s *PublishAdvanceConfigRequestFiles) GetFileName() *string {
	return s.FileName
}

func (s *PublishAdvanceConfigRequestFiles) GetOperateType() *string {
	return s.OperateType
}

func (s *PublishAdvanceConfigRequestFiles) GetOssPath() *string {
	return s.OssPath
}

func (s *PublishAdvanceConfigRequestFiles) GetParentFullPath() *string {
	return s.ParentFullPath
}

func (s *PublishAdvanceConfigRequestFiles) SetConfig(v *PublishAdvanceConfigRequestFilesConfig) *PublishAdvanceConfigRequestFiles {
	s.Config = v
	return s
}

func (s *PublishAdvanceConfigRequestFiles) SetDirName(v string) *PublishAdvanceConfigRequestFiles {
	s.DirName = &v
	return s
}

func (s *PublishAdvanceConfigRequestFiles) SetFileName(v string) *PublishAdvanceConfigRequestFiles {
	s.FileName = &v
	return s
}

func (s *PublishAdvanceConfigRequestFiles) SetOperateType(v string) *PublishAdvanceConfigRequestFiles {
	s.OperateType = &v
	return s
}

func (s *PublishAdvanceConfigRequestFiles) SetOssPath(v string) *PublishAdvanceConfigRequestFiles {
	s.OssPath = &v
	return s
}

func (s *PublishAdvanceConfigRequestFiles) SetParentFullPath(v string) *PublishAdvanceConfigRequestFiles {
	s.ParentFullPath = &v
	return s
}

func (s *PublishAdvanceConfigRequestFiles) Validate() error {
	return dara.Validate(s)
}

type PublishAdvanceConfigRequestFilesConfig struct {
	// The file content.
	//
	// example:
	//
	// {\\"url\\":\\"http://xxxxxx.aliyuncs.com/outnet_hz/packages/xxxxx/opensearch_offline_plugins_xxxxx.tar\\"}
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// The variables.
	Variables map[string]*FilesConfigVariablesValue `json:"variables,omitempty" xml:"variables,omitempty"`
}

func (s PublishAdvanceConfigRequestFilesConfig) String() string {
	return dara.Prettify(s)
}

func (s PublishAdvanceConfigRequestFilesConfig) GoString() string {
	return s.String()
}

func (s *PublishAdvanceConfigRequestFilesConfig) GetContent() *string {
	return s.Content
}

func (s *PublishAdvanceConfigRequestFilesConfig) GetVariables() map[string]*FilesConfigVariablesValue {
	return s.Variables
}

func (s *PublishAdvanceConfigRequestFilesConfig) SetContent(v string) *PublishAdvanceConfigRequestFilesConfig {
	s.Content = &v
	return s
}

func (s *PublishAdvanceConfigRequestFilesConfig) SetVariables(v map[string]*FilesConfigVariablesValue) *PublishAdvanceConfigRequestFilesConfig {
	s.Variables = v
	return s
}

func (s *PublishAdvanceConfigRequestFilesConfig) Validate() error {
	return dara.Validate(s)
}
