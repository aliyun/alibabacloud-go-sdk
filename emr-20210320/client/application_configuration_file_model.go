// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplicationConfigurationFile interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationName(v string) *ApplicationConfigurationFile
	GetApplicationName() *string
	SetClusterId(v string) *ApplicationConfigurationFile
	GetClusterId() *string
	SetConfigFileFormat(v string) *ApplicationConfigurationFile
	GetConfigFileFormat() *string
	SetConfigFileGroup(v string) *ApplicationConfigurationFile
	GetConfigFileGroup() *string
	SetConfigFileLink(v string) *ApplicationConfigurationFile
	GetConfigFileLink() *string
	SetConfigFileMode(v string) *ApplicationConfigurationFile
	GetConfigFileMode() *string
	SetConfigFileName(v string) *ApplicationConfigurationFile
	GetConfigFileName() *string
	SetConfigFileOwner(v string) *ApplicationConfigurationFile
	GetConfigFileOwner() *string
	SetConfigFilePath(v string) *ApplicationConfigurationFile
	GetConfigFilePath() *string
	SetDescription(v string) *ApplicationConfigurationFile
	GetDescription() *string
	SetNodeGroupId(v string) *ApplicationConfigurationFile
	GetNodeGroupId() *string
	SetNodeId(v string) *ApplicationConfigurationFile
	GetNodeId() *string
}

type ApplicationConfigurationFile struct {
	ApplicationName  *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	ClusterId        *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ConfigFileFormat *string `json:"ConfigFileFormat,omitempty" xml:"ConfigFileFormat,omitempty"`
	ConfigFileGroup  *string `json:"ConfigFileGroup,omitempty" xml:"ConfigFileGroup,omitempty"`
	ConfigFileLink   *string `json:"ConfigFileLink,omitempty" xml:"ConfigFileLink,omitempty"`
	ConfigFileMode   *string `json:"ConfigFileMode,omitempty" xml:"ConfigFileMode,omitempty"`
	ConfigFileName   *string `json:"ConfigFileName,omitempty" xml:"ConfigFileName,omitempty"`
	ConfigFileOwner  *string `json:"ConfigFileOwner,omitempty" xml:"ConfigFileOwner,omitempty"`
	ConfigFilePath   *string `json:"ConfigFilePath,omitempty" xml:"ConfigFilePath,omitempty"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	NodeGroupId      *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	NodeId           *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s ApplicationConfigurationFile) String() string {
	return dara.Prettify(s)
}

func (s ApplicationConfigurationFile) GoString() string {
	return s.String()
}

func (s *ApplicationConfigurationFile) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *ApplicationConfigurationFile) GetClusterId() *string {
	return s.ClusterId
}

func (s *ApplicationConfigurationFile) GetConfigFileFormat() *string {
	return s.ConfigFileFormat
}

func (s *ApplicationConfigurationFile) GetConfigFileGroup() *string {
	return s.ConfigFileGroup
}

func (s *ApplicationConfigurationFile) GetConfigFileLink() *string {
	return s.ConfigFileLink
}

func (s *ApplicationConfigurationFile) GetConfigFileMode() *string {
	return s.ConfigFileMode
}

func (s *ApplicationConfigurationFile) GetConfigFileName() *string {
	return s.ConfigFileName
}

func (s *ApplicationConfigurationFile) GetConfigFileOwner() *string {
	return s.ConfigFileOwner
}

func (s *ApplicationConfigurationFile) GetConfigFilePath() *string {
	return s.ConfigFilePath
}

func (s *ApplicationConfigurationFile) GetDescription() *string {
	return s.Description
}

func (s *ApplicationConfigurationFile) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *ApplicationConfigurationFile) GetNodeId() *string {
	return s.NodeId
}

func (s *ApplicationConfigurationFile) SetApplicationName(v string) *ApplicationConfigurationFile {
	s.ApplicationName = &v
	return s
}

func (s *ApplicationConfigurationFile) SetClusterId(v string) *ApplicationConfigurationFile {
	s.ClusterId = &v
	return s
}

func (s *ApplicationConfigurationFile) SetConfigFileFormat(v string) *ApplicationConfigurationFile {
	s.ConfigFileFormat = &v
	return s
}

func (s *ApplicationConfigurationFile) SetConfigFileGroup(v string) *ApplicationConfigurationFile {
	s.ConfigFileGroup = &v
	return s
}

func (s *ApplicationConfigurationFile) SetConfigFileLink(v string) *ApplicationConfigurationFile {
	s.ConfigFileLink = &v
	return s
}

func (s *ApplicationConfigurationFile) SetConfigFileMode(v string) *ApplicationConfigurationFile {
	s.ConfigFileMode = &v
	return s
}

func (s *ApplicationConfigurationFile) SetConfigFileName(v string) *ApplicationConfigurationFile {
	s.ConfigFileName = &v
	return s
}

func (s *ApplicationConfigurationFile) SetConfigFileOwner(v string) *ApplicationConfigurationFile {
	s.ConfigFileOwner = &v
	return s
}

func (s *ApplicationConfigurationFile) SetConfigFilePath(v string) *ApplicationConfigurationFile {
	s.ConfigFilePath = &v
	return s
}

func (s *ApplicationConfigurationFile) SetDescription(v string) *ApplicationConfigurationFile {
	s.Description = &v
	return s
}

func (s *ApplicationConfigurationFile) SetNodeGroupId(v string) *ApplicationConfigurationFile {
	s.NodeGroupId = &v
	return s
}

func (s *ApplicationConfigurationFile) SetNodeId(v string) *ApplicationConfigurationFile {
	s.NodeId = &v
	return s
}

func (s *ApplicationConfigurationFile) Validate() error {
	return dara.Validate(s)
}
