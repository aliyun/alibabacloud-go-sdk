// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublicUpdateTemplateLogConfig interface {
	dara.Model
	String() string
	GoString() string
	SetLogstore(v string) *PublicUpdateTemplateLogConfig
	GetLogstore() *string
	SetProject(v string) *PublicUpdateTemplateLogConfig
	GetProject() *string
}

type PublicUpdateTemplateLogConfig struct {
	Logstore *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	Project  *string `json:"project,omitempty" xml:"project,omitempty"`
}

func (s PublicUpdateTemplateLogConfig) String() string {
	return dara.Prettify(s)
}

func (s PublicUpdateTemplateLogConfig) GoString() string {
	return s.String()
}

func (s *PublicUpdateTemplateLogConfig) GetLogstore() *string {
	return s.Logstore
}

func (s *PublicUpdateTemplateLogConfig) GetProject() *string {
	return s.Project
}

func (s *PublicUpdateTemplateLogConfig) SetLogstore(v string) *PublicUpdateTemplateLogConfig {
	s.Logstore = &v
	return s
}

func (s *PublicUpdateTemplateLogConfig) SetProject(v string) *PublicUpdateTemplateLogConfig {
	s.Project = &v
	return s
}

func (s *PublicUpdateTemplateLogConfig) Validate() error {
	return dara.Validate(s)
}
