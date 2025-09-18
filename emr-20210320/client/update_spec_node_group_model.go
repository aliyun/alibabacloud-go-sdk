// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSpecNodeGroup interface {
	dara.Model
	String() string
	GoString() string
	SetModifyType(v string) *UpdateSpecNodeGroup
	GetModifyType() *string
	SetNewInstanceType(v string) *UpdateSpecNodeGroup
	GetNewInstanceType() *string
	SetNodeGroupId(v string) *UpdateSpecNodeGroup
	GetNodeGroupId() *string
}

type UpdateSpecNodeGroup struct {
	ModifyType *string `json:"ModifyType,omitempty" xml:"ModifyType,omitempty"`
	// 新实例类型。
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs.g7.xlarge
	NewInstanceType *string `json:"NewInstanceType,omitempty" xml:"NewInstanceType,omitempty"`
	// 节点组ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// ng-869471354ecd****
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
}

func (s UpdateSpecNodeGroup) String() string {
	return dara.Prettify(s)
}

func (s UpdateSpecNodeGroup) GoString() string {
	return s.String()
}

func (s *UpdateSpecNodeGroup) GetModifyType() *string {
	return s.ModifyType
}

func (s *UpdateSpecNodeGroup) GetNewInstanceType() *string {
	return s.NewInstanceType
}

func (s *UpdateSpecNodeGroup) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *UpdateSpecNodeGroup) SetModifyType(v string) *UpdateSpecNodeGroup {
	s.ModifyType = &v
	return s
}

func (s *UpdateSpecNodeGroup) SetNewInstanceType(v string) *UpdateSpecNodeGroup {
	s.NewInstanceType = &v
	return s
}

func (s *UpdateSpecNodeGroup) SetNodeGroupId(v string) *UpdateSpecNodeGroup {
	s.NodeGroupId = &v
	return s
}

func (s *UpdateSpecNodeGroup) Validate() error {
	return dara.Validate(s)
}
