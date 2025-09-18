// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplicationConfigFile interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationName(v string) *ApplicationConfigFile
	GetApplicationName() *string
	SetConfigFileName(v string) *ApplicationConfigFile
	GetConfigFileName() *string
}

type ApplicationConfigFile struct {
	// 应用名称。
	//
	// This parameter is required.
	//
	// example:
	//
	// HDFS
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// 配置文件名称。
	//
	// example:
	//
	// core-site.xml
	ConfigFileName *string `json:"ConfigFileName,omitempty" xml:"ConfigFileName,omitempty"`
}

func (s ApplicationConfigFile) String() string {
	return dara.Prettify(s)
}

func (s ApplicationConfigFile) GoString() string {
	return s.String()
}

func (s *ApplicationConfigFile) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *ApplicationConfigFile) GetConfigFileName() *string {
	return s.ConfigFileName
}

func (s *ApplicationConfigFile) SetApplicationName(v string) *ApplicationConfigFile {
	s.ApplicationName = &v
	return s
}

func (s *ApplicationConfigFile) SetConfigFileName(v string) *ApplicationConfigFile {
	s.ConfigFileName = &v
	return s
}

func (s *ApplicationConfigFile) Validate() error {
	return dara.Validate(s)
}
