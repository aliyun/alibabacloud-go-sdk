// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSpecNodeGroupParam interface {
	dara.Model
	String() string
	GoString() string
	SetNewInstanceType(v string) *UpdateSpecNodeGroupParam
	GetNewInstanceType() *string
	SetNodeGroupId(v string) *UpdateSpecNodeGroupParam
	GetNodeGroupId() *string
}

type UpdateSpecNodeGroupParam struct {
	NewInstanceType *string `json:"NewInstanceType,omitempty" xml:"NewInstanceType,omitempty"`
	NodeGroupId     *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
}

func (s UpdateSpecNodeGroupParam) String() string {
	return dara.Prettify(s)
}

func (s UpdateSpecNodeGroupParam) GoString() string {
	return s.String()
}

func (s *UpdateSpecNodeGroupParam) GetNewInstanceType() *string {
	return s.NewInstanceType
}

func (s *UpdateSpecNodeGroupParam) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *UpdateSpecNodeGroupParam) SetNewInstanceType(v string) *UpdateSpecNodeGroupParam {
	s.NewInstanceType = &v
	return s
}

func (s *UpdateSpecNodeGroupParam) SetNodeGroupId(v string) *UpdateSpecNodeGroupParam {
	s.NodeGroupId = &v
	return s
}

func (s *UpdateSpecNodeGroupParam) Validate() error {
	return dara.Validate(s)
}
