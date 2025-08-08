// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTemplateSourceConfig interface {
	dara.Model
	String() string
	GoString() string
	SetDownloadUrl(v string) *TemplateSourceConfig
	GetDownloadUrl() *string
	SetName(v string) *TemplateSourceConfig
	GetName() *string
}

type TemplateSourceConfig struct {
	// example:
	//
	// https://api.devsapp.cn/v3/packages/start-modelscope-v3/zipball/0.1.6
	DownloadUrl *string `json:"downloadUrl,omitempty" xml:"downloadUrl,omitempty"`
	// example:
	//
	// start-springboot-cap
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s TemplateSourceConfig) String() string {
	return dara.Prettify(s)
}

func (s TemplateSourceConfig) GoString() string {
	return s.String()
}

func (s *TemplateSourceConfig) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *TemplateSourceConfig) GetName() *string {
	return s.Name
}

func (s *TemplateSourceConfig) SetDownloadUrl(v string) *TemplateSourceConfig {
	s.DownloadUrl = &v
	return s
}

func (s *TemplateSourceConfig) SetName(v string) *TemplateSourceConfig {
	s.Name = &v
	return s
}

func (s *TemplateSourceConfig) Validate() error {
	return dara.Validate(s)
}
