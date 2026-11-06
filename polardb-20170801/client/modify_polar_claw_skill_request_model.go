// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPolarClawSkillRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *ModifyPolarClawSkillRequest
	GetApplicationId() *string
	SetModifyType(v string) *ModifyPolarClawSkillRequest
	GetModifyType() *string
	SetPath(v string) *ModifyPolarClawSkillRequest
	GetPath() *string
	SetRegionId(v string) *ModifyPolarClawSkillRequest
	GetRegionId() *string
	SetSkillName(v string) *ModifyPolarClawSkillRequest
	GetSkillName() *string
	SetSkillVersion(v string) *ModifyPolarClawSkillRequest
	GetSkillVersion() *string
}

type ModifyPolarClawSkillRequest struct {
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The modification type. Valid values:
	//
	// 	- **evaluate**: Evaluates the skill.
	//
	// This parameter is required.
	//
	// example:
	//
	// evaluate
	ModifyType *string `json:"ModifyType,omitempty" xml:"ModifyType,omitempty"`
	// The destination path.
	//
	// example:
	//
	// /home/node/.openclaw/workspace/skills/polarclawhub
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The skill name.
	//
	// This parameter is required.
	//
	// example:
	//
	// polarclawhub
	SkillName *string `json:"SkillName,omitempty" xml:"SkillName,omitempty"`
	// The skill version.
	//
	// example:
	//
	// 0.0.0
	SkillVersion *string `json:"SkillVersion,omitempty" xml:"SkillVersion,omitempty"`
}

func (s ModifyPolarClawSkillRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyPolarClawSkillRequest) GoString() string {
	return s.String()
}

func (s *ModifyPolarClawSkillRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ModifyPolarClawSkillRequest) GetModifyType() *string {
	return s.ModifyType
}

func (s *ModifyPolarClawSkillRequest) GetPath() *string {
	return s.Path
}

func (s *ModifyPolarClawSkillRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyPolarClawSkillRequest) GetSkillName() *string {
	return s.SkillName
}

func (s *ModifyPolarClawSkillRequest) GetSkillVersion() *string {
	return s.SkillVersion
}

func (s *ModifyPolarClawSkillRequest) SetApplicationId(v string) *ModifyPolarClawSkillRequest {
	s.ApplicationId = &v
	return s
}

func (s *ModifyPolarClawSkillRequest) SetModifyType(v string) *ModifyPolarClawSkillRequest {
	s.ModifyType = &v
	return s
}

func (s *ModifyPolarClawSkillRequest) SetPath(v string) *ModifyPolarClawSkillRequest {
	s.Path = &v
	return s
}

func (s *ModifyPolarClawSkillRequest) SetRegionId(v string) *ModifyPolarClawSkillRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyPolarClawSkillRequest) SetSkillName(v string) *ModifyPolarClawSkillRequest {
	s.SkillName = &v
	return s
}

func (s *ModifyPolarClawSkillRequest) SetSkillVersion(v string) *ModifyPolarClawSkillRequest {
	s.SkillVersion = &v
	return s
}

func (s *ModifyPolarClawSkillRequest) Validate() error {
	return dara.Validate(s)
}
