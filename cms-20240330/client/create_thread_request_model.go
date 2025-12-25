// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateThreadRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTitle(v string) *CreateThreadRequest
	GetTitle() *string
	SetVariables(v *CreateThreadRequestVariables) *CreateThreadRequest
	GetVariables() *CreateThreadRequestVariables
}

type CreateThreadRequest struct {
	// example:
	//
	// test
	Title     *string                       `json:"title,omitempty" xml:"title,omitempty"`
	Variables *CreateThreadRequestVariables `json:"variables,omitempty" xml:"variables,omitempty" type:"Struct"`
}

func (s CreateThreadRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateThreadRequest) GoString() string {
	return s.String()
}

func (s *CreateThreadRequest) GetTitle() *string {
	return s.Title
}

func (s *CreateThreadRequest) GetVariables() *CreateThreadRequestVariables {
	return s.Variables
}

func (s *CreateThreadRequest) SetTitle(v string) *CreateThreadRequest {
	s.Title = &v
	return s
}

func (s *CreateThreadRequest) SetVariables(v *CreateThreadRequestVariables) *CreateThreadRequest {
	s.Variables = v
	return s
}

func (s *CreateThreadRequest) Validate() error {
	if s.Variables != nil {
		if err := s.Variables.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateThreadRequestVariables struct {
	// example:
	//
	// az_alipay
	Project *string `json:"project,omitempty" xml:"project,omitempty"`
	// example:
	//
	// rum-monitor-test-aysls-pub-cn-heyuan-monitor
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s CreateThreadRequestVariables) String() string {
	return dara.Prettify(s)
}

func (s CreateThreadRequestVariables) GoString() string {
	return s.String()
}

func (s *CreateThreadRequestVariables) GetProject() *string {
	return s.Project
}

func (s *CreateThreadRequestVariables) GetWorkspace() *string {
	return s.Workspace
}

func (s *CreateThreadRequestVariables) SetProject(v string) *CreateThreadRequestVariables {
	s.Project = &v
	return s
}

func (s *CreateThreadRequestVariables) SetWorkspace(v string) *CreateThreadRequestVariables {
	s.Workspace = &v
	return s
}

func (s *CreateThreadRequestVariables) Validate() error {
	return dara.Validate(s)
}
