// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceLinkedRoleStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthType(v string) *GetServiceLinkedRoleStatusRequest
	GetAuthType() *string
}

type GetServiceLinkedRoleStatusRequest struct {
	// The authorization type. Valid values:
	//
	// - **SecuritySense**: security check authorization
	//
	// - **DisposalTool**: threat removal tool authorization
	//
	// - **SensePosture**: security posture authorization
	//
	// This parameter is required.
	//
	// example:
	//
	// DisposalTool
	AuthType *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
}

func (s GetServiceLinkedRoleStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetServiceLinkedRoleStatusRequest) GoString() string {
	return s.String()
}

func (s *GetServiceLinkedRoleStatusRequest) GetAuthType() *string {
	return s.AuthType
}

func (s *GetServiceLinkedRoleStatusRequest) SetAuthType(v string) *GetServiceLinkedRoleStatusRequest {
	s.AuthType = &v
	return s
}

func (s *GetServiceLinkedRoleStatusRequest) Validate() error {
	return dara.Validate(s)
}
