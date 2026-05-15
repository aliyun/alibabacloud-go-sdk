// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSkillGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOuterGroupId(v string) *DeleteSkillGroupRequest
	GetOuterGroupId() *string
	SetOuterGroupType(v string) *DeleteSkillGroupRequest
	GetOuterGroupType() *string
}

type DeleteSkillGroupRequest struct {
	// example:
	//
	// 123456
	OuterGroupId *string `json:"OuterGroupId,omitempty" xml:"OuterGroupId,omitempty"`
	// example:
	//
	// 2
	OuterGroupType *string `json:"OuterGroupType,omitempty" xml:"OuterGroupType,omitempty"`
}

func (s DeleteSkillGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSkillGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteSkillGroupRequest) GetOuterGroupId() *string {
	return s.OuterGroupId
}

func (s *DeleteSkillGroupRequest) GetOuterGroupType() *string {
	return s.OuterGroupType
}

func (s *DeleteSkillGroupRequest) SetOuterGroupId(v string) *DeleteSkillGroupRequest {
	s.OuterGroupId = &v
	return s
}

func (s *DeleteSkillGroupRequest) SetOuterGroupType(v string) *DeleteSkillGroupRequest {
	s.OuterGroupType = &v
	return s
}

func (s *DeleteSkillGroupRequest) Validate() error {
	return dara.Validate(s)
}
