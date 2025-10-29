// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLogConfiguration interface {
	dara.Model
	String() string
	GoString() string
	SetLogstore(v string) *LogConfiguration
	GetLogstore() *string
	SetProject(v string) *LogConfiguration
	GetProject() *string
}

type LogConfiguration struct {
	// SLS日志库名称
	//
	// example:
	//
	// agent-runtime-logs
	Logstore *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	// SLS项目名称
	//
	// example:
	//
	// agent-runtime-logs
	Project *string `json:"project,omitempty" xml:"project,omitempty"`
}

func (s LogConfiguration) String() string {
	return dara.Prettify(s)
}

func (s LogConfiguration) GoString() string {
	return s.String()
}

func (s *LogConfiguration) GetLogstore() *string {
	return s.Logstore
}

func (s *LogConfiguration) GetProject() *string {
	return s.Project
}

func (s *LogConfiguration) SetLogstore(v string) *LogConfiguration {
	s.Logstore = &v
	return s
}

func (s *LogConfiguration) SetProject(v string) *LogConfiguration {
	s.Project = &v
	return s
}

func (s *LogConfiguration) Validate() error {
	return dara.Validate(s)
}
