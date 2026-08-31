// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublicTemplateLogConfig interface {
	dara.Model
	String() string
	GoString() string
	SetLogstore(v string) *PublicTemplateLogConfig
	GetLogstore() *string
	SetProject(v string) *PublicTemplateLogConfig
	GetProject() *string
}

type PublicTemplateLogConfig struct {
	// The name of the SLS Logstore.
	//
	// example:
	//
	// my-logstore
	Logstore *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	// The name of the Simple Log Service (SLS) project.
	//
	// example:
	//
	// my-sls-project
	Project *string `json:"project,omitempty" xml:"project,omitempty"`
}

func (s PublicTemplateLogConfig) String() string {
	return dara.Prettify(s)
}

func (s PublicTemplateLogConfig) GoString() string {
	return s.String()
}

func (s *PublicTemplateLogConfig) GetLogstore() *string {
	return s.Logstore
}

func (s *PublicTemplateLogConfig) GetProject() *string {
	return s.Project
}

func (s *PublicTemplateLogConfig) SetLogstore(v string) *PublicTemplateLogConfig {
	s.Logstore = &v
	return s
}

func (s *PublicTemplateLogConfig) SetProject(v string) *PublicTemplateLogConfig {
	s.Project = &v
	return s
}

func (s *PublicTemplateLogConfig) Validate() error {
	return dara.Validate(s)
}
