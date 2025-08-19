// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSLSTriggerLogConfig interface {
	dara.Model
	String() string
	GoString() string
	SetLogstore(v string) *SLSTriggerLogConfig
	GetLogstore() *string
	SetProject(v string) *SLSTriggerLogConfig
	GetProject() *string
}

type SLSTriggerLogConfig struct {
	// example:
	//
	// my-sls-logstore-name
	Logstore *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	// example:
	//
	// my-sls-project-name
	Project *string `json:"project,omitempty" xml:"project,omitempty"`
}

func (s SLSTriggerLogConfig) String() string {
	return dara.Prettify(s)
}

func (s SLSTriggerLogConfig) GoString() string {
	return s.String()
}

func (s *SLSTriggerLogConfig) GetLogstore() *string {
	return s.Logstore
}

func (s *SLSTriggerLogConfig) GetProject() *string {
	return s.Project
}

func (s *SLSTriggerLogConfig) SetLogstore(v string) *SLSTriggerLogConfig {
	s.Logstore = &v
	return s
}

func (s *SLSTriggerLogConfig) SetProject(v string) *SLSTriggerLogConfig {
	s.Project = &v
	return s
}

func (s *SLSTriggerLogConfig) Validate() error {
	return dara.Validate(s)
}
