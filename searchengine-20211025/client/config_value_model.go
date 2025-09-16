// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigValue interface {
	dara.Model
	String() string
	GoString() string
	SetDesc(v string) *ConfigValue
	GetDesc() *string
	SetFiles(v []*ConfigValueFiles) *ConfigValue
	GetFiles() []*ConfigValueFiles
}

type ConfigValue struct {
	// The description of the offline configuration.
	//
	// example:
	//
	// test
	Desc *string `json:"desc,omitempty" xml:"desc,omitempty"`
	// The files to be modified.
	Files []*ConfigValueFiles `json:"files,omitempty" xml:"files,omitempty" type:"Repeated"`
}

func (s ConfigValue) String() string {
	return dara.Prettify(s)
}

func (s ConfigValue) GoString() string {
	return s.String()
}

func (s *ConfigValue) GetDesc() *string {
	return s.Desc
}

func (s *ConfigValue) GetFiles() []*ConfigValueFiles {
	return s.Files
}

func (s *ConfigValue) SetDesc(v string) *ConfigValue {
	s.Desc = &v
	return s
}

func (s *ConfigValue) SetFiles(v []*ConfigValueFiles) *ConfigValue {
	s.Files = v
	return s
}

func (s *ConfigValue) Validate() error {
	return dara.Validate(s)
}

type ConfigValueFiles struct {
	// The operation type. Valid values: UPDATE and DELETE. Default value: UPDATE.
	//
	// example:
	//
	// UPDATE
	OperateType *string `json:"operateType,omitempty" xml:"operateType,omitempty"`
	// The path of the parent directory.
	//
	// example:
	//
	// /
	ParentFullPath *string `json:"parentFullPath,omitempty" xml:"parentFullPath,omitempty"`
	// The file name.
	//
	// example:
	//
	// /intervene_dict/中文-通用分析器.dict
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// The configuration to be modified.
	Config *ConfigValueFilesConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
	// The directory name.
	//
	// example:
	//
	// /test
	DirName *string `json:"dirName,omitempty" xml:"dirName,omitempty"`
}

func (s ConfigValueFiles) String() string {
	return dara.Prettify(s)
}

func (s ConfigValueFiles) GoString() string {
	return s.String()
}

func (s *ConfigValueFiles) GetOperateType() *string {
	return s.OperateType
}

func (s *ConfigValueFiles) GetParentFullPath() *string {
	return s.ParentFullPath
}

func (s *ConfigValueFiles) GetFileName() *string {
	return s.FileName
}

func (s *ConfigValueFiles) GetConfig() *ConfigValueFilesConfig {
	return s.Config
}

func (s *ConfigValueFiles) GetDirName() *string {
	return s.DirName
}

func (s *ConfigValueFiles) SetOperateType(v string) *ConfigValueFiles {
	s.OperateType = &v
	return s
}

func (s *ConfigValueFiles) SetParentFullPath(v string) *ConfigValueFiles {
	s.ParentFullPath = &v
	return s
}

func (s *ConfigValueFiles) SetFileName(v string) *ConfigValueFiles {
	s.FileName = &v
	return s
}

func (s *ConfigValueFiles) SetConfig(v *ConfigValueFilesConfig) *ConfigValueFiles {
	s.Config = v
	return s
}

func (s *ConfigValueFiles) SetDirName(v string) *ConfigValueFiles {
	s.DirName = &v
	return s
}

func (s *ConfigValueFiles) Validate() error {
	return dara.Validate(s)
}

type ConfigValueFilesConfig struct {
	// The file content.
	//
	// example:
	//
	// $dictContent
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// The variables.
	Variables map[string]*ConfigValueFilesConfigVariablesValue `json:"variables,omitempty" xml:"variables,omitempty"`
}

func (s ConfigValueFilesConfig) String() string {
	return dara.Prettify(s)
}

func (s ConfigValueFilesConfig) GoString() string {
	return s.String()
}

func (s *ConfigValueFilesConfig) GetContent() *string {
	return s.Content
}

func (s *ConfigValueFilesConfig) GetVariables() map[string]*ConfigValueFilesConfigVariablesValue {
	return s.Variables
}

func (s *ConfigValueFilesConfig) SetContent(v string) *ConfigValueFilesConfig {
	s.Content = &v
	return s
}

func (s *ConfigValueFilesConfig) SetVariables(v map[string]*ConfigValueFilesConfigVariablesValue) *ConfigValueFilesConfig {
	s.Variables = v
	return s
}

func (s *ConfigValueFilesConfig) Validate() error {
	return dara.Validate(s)
}
