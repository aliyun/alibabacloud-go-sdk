// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSLSLogDestination interface {
	dara.Model
	String() string
	GoString() string
	SetLogStore(v string) *SLSLogDestination
	GetLogStore() *string
	SetProject(v string) *SLSLogDestination
	GetProject() *string
}

type SLSLogDestination struct {
	// SLS日志库名称
	//
	// example:
	//
	// my-logstore
	LogStore *string `json:"logStore,omitempty" xml:"logStore,omitempty"`
	// SLS项目名称
	//
	// example:
	//
	// my-sls-project
	Project *string `json:"project,omitempty" xml:"project,omitempty"`
}

func (s SLSLogDestination) String() string {
	return dara.Prettify(s)
}

func (s SLSLogDestination) GoString() string {
	return s.String()
}

func (s *SLSLogDestination) GetLogStore() *string {
	return s.LogStore
}

func (s *SLSLogDestination) GetProject() *string {
	return s.Project
}

func (s *SLSLogDestination) SetLogStore(v string) *SLSLogDestination {
	s.LogStore = &v
	return s
}

func (s *SLSLogDestination) SetProject(v string) *SLSLogDestination {
	s.Project = &v
	return s
}

func (s *SLSLogDestination) Validate() error {
	return dara.Validate(s)
}
