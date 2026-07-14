// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableNotifyPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetWorkspace(v string) *DisableNotifyPolicyRequest
	GetWorkspace() *string
}

type DisableNotifyPolicyRequest struct {
	// The workspace ID. Used to isolate notification policy resources across different business spaces.
	//
	// This parameter is required.
	//
	// example:
	//
	// default-cms-xxxx-cn-hangzhou
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s DisableNotifyPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableNotifyPolicyRequest) GoString() string {
	return s.String()
}

func (s *DisableNotifyPolicyRequest) GetWorkspace() *string {
	return s.Workspace
}

func (s *DisableNotifyPolicyRequest) SetWorkspace(v string) *DisableNotifyPolicyRequest {
	s.Workspace = &v
	return s
}

func (s *DisableNotifyPolicyRequest) Validate() error {
	return dara.Validate(s)
}
