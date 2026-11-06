// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPolarClawSkillRegionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *GetPolarClawSkillRegionRequest
	GetRegionId() *string
}

type GetPolarClawSkillRegionRequest struct {
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetPolarClawSkillRegionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPolarClawSkillRegionRequest) GoString() string {
	return s.String()
}

func (s *GetPolarClawSkillRegionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetPolarClawSkillRegionRequest) SetRegionId(v string) *GetPolarClawSkillRegionRequest {
	s.RegionId = &v
	return s
}

func (s *GetPolarClawSkillRegionRequest) Validate() error {
	return dara.Validate(s)
}
