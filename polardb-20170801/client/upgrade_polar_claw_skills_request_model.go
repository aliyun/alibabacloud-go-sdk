// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradePolarClawSkillsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *UpgradePolarClawSkillsRequest
	GetApplicationId() *string
	SetApplicationType(v string) *UpgradePolarClawSkillsRequest
	GetApplicationType() *string
	SetRegionId(v string) *UpgradePolarClawSkillsRequest
	GetRegionId() *string
	SetUpgradeMethod(v string) *UpgradePolarClawSkillsRequest
	GetUpgradeMethod() *string
}

type UpgradePolarClawSkillsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// example:
	//
	// polarclaw
	ApplicationType *string `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UpgradeMethod *string `json:"UpgradeMethod,omitempty" xml:"UpgradeMethod,omitempty"`
}

func (s UpgradePolarClawSkillsRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradePolarClawSkillsRequest) GoString() string {
	return s.String()
}

func (s *UpgradePolarClawSkillsRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *UpgradePolarClawSkillsRequest) GetApplicationType() *string {
	return s.ApplicationType
}

func (s *UpgradePolarClawSkillsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpgradePolarClawSkillsRequest) GetUpgradeMethod() *string {
	return s.UpgradeMethod
}

func (s *UpgradePolarClawSkillsRequest) SetApplicationId(v string) *UpgradePolarClawSkillsRequest {
	s.ApplicationId = &v
	return s
}

func (s *UpgradePolarClawSkillsRequest) SetApplicationType(v string) *UpgradePolarClawSkillsRequest {
	s.ApplicationType = &v
	return s
}

func (s *UpgradePolarClawSkillsRequest) SetRegionId(v string) *UpgradePolarClawSkillsRequest {
	s.RegionId = &v
	return s
}

func (s *UpgradePolarClawSkillsRequest) SetUpgradeMethod(v string) *UpgradePolarClawSkillsRequest {
	s.UpgradeMethod = &v
	return s
}

func (s *UpgradePolarClawSkillsRequest) Validate() error {
	return dara.Validate(s)
}
