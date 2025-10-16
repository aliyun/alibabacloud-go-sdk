// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVpcFirewallTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *CreateVpcFirewallTaskRequest
	GetContent() *string
	SetLang(v string) *CreateVpcFirewallTaskRequest
	GetLang() *string
	SetPriority(v string) *CreateVpcFirewallTaskRequest
	GetPriority() *string
	SetTaskAction(v string) *CreateVpcFirewallTaskRequest
	GetTaskAction() *string
}

type CreateVpcFirewallTaskRequest struct {
	// example:
	//
	// test
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 1
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sync
	TaskAction *string `json:"TaskAction,omitempty" xml:"TaskAction,omitempty"`
}

func (s CreateVpcFirewallTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVpcFirewallTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateVpcFirewallTaskRequest) GetContent() *string {
	return s.Content
}

func (s *CreateVpcFirewallTaskRequest) GetLang() *string {
	return s.Lang
}

func (s *CreateVpcFirewallTaskRequest) GetPriority() *string {
	return s.Priority
}

func (s *CreateVpcFirewallTaskRequest) GetTaskAction() *string {
	return s.TaskAction
}

func (s *CreateVpcFirewallTaskRequest) SetContent(v string) *CreateVpcFirewallTaskRequest {
	s.Content = &v
	return s
}

func (s *CreateVpcFirewallTaskRequest) SetLang(v string) *CreateVpcFirewallTaskRequest {
	s.Lang = &v
	return s
}

func (s *CreateVpcFirewallTaskRequest) SetPriority(v string) *CreateVpcFirewallTaskRequest {
	s.Priority = &v
	return s
}

func (s *CreateVpcFirewallTaskRequest) SetTaskAction(v string) *CreateVpcFirewallTaskRequest {
	s.TaskAction = &v
	return s
}

func (s *CreateVpcFirewallTaskRequest) Validate() error {
	return dara.Validate(s)
}
