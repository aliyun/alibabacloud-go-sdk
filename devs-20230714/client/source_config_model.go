// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSourceConfig interface {
	dara.Model
	String() string
	GoString() string
	SetOss(v *OssSourceConfig) *SourceConfig
	GetOss() *OssSourceConfig
	SetRepository(v *RepositorySourceConfig) *SourceConfig
	GetRepository() *RepositorySourceConfig
	SetTemplate(v *TemplateSourceConfig) *SourceConfig
	GetTemplate() *TemplateSourceConfig
}

type SourceConfig struct {
	Oss        *OssSourceConfig        `json:"oss,omitempty" xml:"oss,omitempty"`
	Repository *RepositorySourceConfig `json:"repository,omitempty" xml:"repository,omitempty"`
	Template   *TemplateSourceConfig   `json:"template,omitempty" xml:"template,omitempty"`
}

func (s SourceConfig) String() string {
	return dara.Prettify(s)
}

func (s SourceConfig) GoString() string {
	return s.String()
}

func (s *SourceConfig) GetOss() *OssSourceConfig {
	return s.Oss
}

func (s *SourceConfig) GetRepository() *RepositorySourceConfig {
	return s.Repository
}

func (s *SourceConfig) GetTemplate() *TemplateSourceConfig {
	return s.Template
}

func (s *SourceConfig) SetOss(v *OssSourceConfig) *SourceConfig {
	s.Oss = v
	return s
}

func (s *SourceConfig) SetRepository(v *RepositorySourceConfig) *SourceConfig {
	s.Repository = v
	return s
}

func (s *SourceConfig) SetTemplate(v *TemplateSourceConfig) *SourceConfig {
	s.Template = v
	return s
}

func (s *SourceConfig) Validate() error {
	return dara.Validate(s)
}
