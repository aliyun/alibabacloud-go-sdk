// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNotifyPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUuid(v string) *GetNotifyPolicyRequest
	GetUuid() *string
	SetWorkspace(v string) *GetNotifyPolicyRequest
	GetWorkspace() *string
}

type GetNotifyPolicyRequest struct {
	// The unique identifier of the notification policy, returned by the create operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// a1b2c3d4-e5f6-7890-abcd-ef1234567890
	Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// The workspace ID. Used to isolate notification policy resources across different business spaces. Example: `default-cms-xxxx-ap-southeast-1`.
	//
	// This parameter is required.
	//
	// example:
	//
	// default-cms-xxxx-cn-hangzhou
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s GetNotifyPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s GetNotifyPolicyRequest) GoString() string {
	return s.String()
}

func (s *GetNotifyPolicyRequest) GetUuid() *string {
	return s.Uuid
}

func (s *GetNotifyPolicyRequest) GetWorkspace() *string {
	return s.Workspace
}

func (s *GetNotifyPolicyRequest) SetUuid(v string) *GetNotifyPolicyRequest {
	s.Uuid = &v
	return s
}

func (s *GetNotifyPolicyRequest) SetWorkspace(v string) *GetNotifyPolicyRequest {
	s.Workspace = &v
	return s
}

func (s *GetNotifyPolicyRequest) Validate() error {
	return dara.Validate(s)
}
