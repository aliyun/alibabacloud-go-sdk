// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePolarClawSkillRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DeletePolarClawSkillRequest
	GetRegionId() *string
	SetSkillName(v string) *DeletePolarClawSkillRequest
	GetSkillName() *string
}

type DeletePolarClawSkillRequest struct {
	// The region ID of the region where the parameter template resides.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/98041.html) operation to query region IDs.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the skill.
	//
	// This parameter is required.
	//
	// example:
	//
	// polarclawhub
	SkillName *string `json:"SkillName,omitempty" xml:"SkillName,omitempty"`
}

func (s DeletePolarClawSkillRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePolarClawSkillRequest) GoString() string {
	return s.String()
}

func (s *DeletePolarClawSkillRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeletePolarClawSkillRequest) GetSkillName() *string {
	return s.SkillName
}

func (s *DeletePolarClawSkillRequest) SetRegionId(v string) *DeletePolarClawSkillRequest {
	s.RegionId = &v
	return s
}

func (s *DeletePolarClawSkillRequest) SetSkillName(v string) *DeletePolarClawSkillRequest {
	s.SkillName = &v
	return s
}

func (s *DeletePolarClawSkillRequest) Validate() error {
	return dara.Validate(s)
}
