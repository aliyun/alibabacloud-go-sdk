// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPermissionCheckCmd interface {
	dara.Model
	String() string
	GoString() string
	SetAction(v string) *PermissionCheckCmd
	GetAction() *string
	SetOperatorId(v string) *PermissionCheckCmd
	GetOperatorId() *string
	SetOperatorType(v string) *PermissionCheckCmd
	GetOperatorType() *string
	SetResource(v string) *PermissionCheckCmd
	GetResource() *string
}

type PermissionCheckCmd struct {
	// This parameter is required.
	//
	// example:
	//
	// write
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// developer
	OperatorType *string `json:"operatorType,omitempty" xml:"operatorType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// neuron:catalog:company/1:pbc/2
	Resource *string `json:"resource,omitempty" xml:"resource,omitempty"`
}

func (s PermissionCheckCmd) String() string {
	return dara.Prettify(s)
}

func (s PermissionCheckCmd) GoString() string {
	return s.String()
}

func (s *PermissionCheckCmd) GetAction() *string {
	return s.Action
}

func (s *PermissionCheckCmd) GetOperatorId() *string {
	return s.OperatorId
}

func (s *PermissionCheckCmd) GetOperatorType() *string {
	return s.OperatorType
}

func (s *PermissionCheckCmd) GetResource() *string {
	return s.Resource
}

func (s *PermissionCheckCmd) SetAction(v string) *PermissionCheckCmd {
	s.Action = &v
	return s
}

func (s *PermissionCheckCmd) SetOperatorId(v string) *PermissionCheckCmd {
	s.OperatorId = &v
	return s
}

func (s *PermissionCheckCmd) SetOperatorType(v string) *PermissionCheckCmd {
	s.OperatorType = &v
	return s
}

func (s *PermissionCheckCmd) SetResource(v string) *PermissionCheckCmd {
	s.Resource = &v
	return s
}

func (s *PermissionCheckCmd) Validate() error {
	return dara.Validate(s)
}
