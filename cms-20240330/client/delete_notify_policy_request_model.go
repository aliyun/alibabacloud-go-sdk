// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNotifyPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUuid(v string) *DeleteNotifyPolicyRequest
	GetUuid() *string
	SetWorkspace(v string) *DeleteNotifyPolicyRequest
	GetWorkspace() *string
}

type DeleteNotifyPolicyRequest struct {
	// The unique identifier of the notification policy, returned by the creation operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// np-12345678-1234-1234-1234-123456789012
	Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// The workspace ID. Used to isolate notification policy resources across different business spaces.
	//
	// This parameter is required.
	//
	// example:
	//
	// default-cms-xxxx-cn-hangzhou
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s DeleteNotifyPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteNotifyPolicyRequest) GoString() string {
	return s.String()
}

func (s *DeleteNotifyPolicyRequest) GetUuid() *string {
	return s.Uuid
}

func (s *DeleteNotifyPolicyRequest) GetWorkspace() *string {
	return s.Workspace
}

func (s *DeleteNotifyPolicyRequest) SetUuid(v string) *DeleteNotifyPolicyRequest {
	s.Uuid = &v
	return s
}

func (s *DeleteNotifyPolicyRequest) SetWorkspace(v string) *DeleteNotifyPolicyRequest {
	s.Workspace = &v
	return s
}

func (s *DeleteNotifyPolicyRequest) Validate() error {
	return dara.Validate(s)
}
