// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTemplateLogConfig interface {
	dara.Model
	String() string
	GoString() string
	SetLogstore(v string) *CreateTemplateLogConfig
	GetLogstore() *string
	SetProject(v string) *CreateTemplateLogConfig
	GetProject() *string
}

type CreateTemplateLogConfig struct {
	// example:
	//
	// sandbox-runtime
	Logstore *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	// example:
	//
	// sandbox-log-cn-hangzhou
	Project *string `json:"project,omitempty" xml:"project,omitempty"`
}

func (s CreateTemplateLogConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateTemplateLogConfig) GoString() string {
	return s.String()
}

func (s *CreateTemplateLogConfig) GetLogstore() *string {
	return s.Logstore
}

func (s *CreateTemplateLogConfig) GetProject() *string {
	return s.Project
}

func (s *CreateTemplateLogConfig) SetLogstore(v string) *CreateTemplateLogConfig {
	s.Logstore = &v
	return s
}

func (s *CreateTemplateLogConfig) SetProject(v string) *CreateTemplateLogConfig {
	s.Project = &v
	return s
}

func (s *CreateTemplateLogConfig) Validate() error {
	return dara.Validate(s)
}
