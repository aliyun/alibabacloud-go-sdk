// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddRCInstancesToDeploymentSetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeploymentSetGroupNo(v string) *AddRCInstancesToDeploymentSetRequest
	GetDeploymentSetGroupNo() *string
	SetDeploymentSetId(v string) *AddRCInstancesToDeploymentSetRequest
	GetDeploymentSetId() *string
	SetForce(v bool) *AddRCInstancesToDeploymentSetRequest
	GetForce() *bool
	SetRCInstanceIds(v string) *AddRCInstancesToDeploymentSetRequest
	GetRCInstanceIds() *string
	SetRegionId(v string) *AddRCInstancesToDeploymentSetRequest
	GetRegionId() *string
}

type AddRCInstancesToDeploymentSetRequest struct {
	// The group number of the ECS instance in the deployment set when the deployment set policy is high availability group (AvailabilityGroup). You can use this parameter to specify the group number. Valid values: 1 to 7. If no value is specified, the system automatically assigns an active group.
	//
	// example:
	//
	// 1
	DeploymentSetGroupNo *string `json:"DeploymentSetGroupNo,omitempty" xml:"DeploymentSetGroupNo,omitempty"`
	// The deployment set ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ds-uf6c8qerk019bj1l****
	DeploymentSetId *string `json:"DeploymentSetId,omitempty" xml:"DeploymentSetId,omitempty"`
	// Specifies whether to forcibly release running instances. Valid values:
	//
	// 	- **true**: Forcibly release.
	//
	// 	- **false*	- (default): Do not forcibly release.
	//
	// example:
	//
	// false
	Force *bool `json:"Force,omitempty" xml:"Force,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rc-aaaa,rc-bbb
	RCInstanceIds *string `json:"RCInstanceIds,omitempty" xml:"RCInstanceIds,omitempty"`
	// The region ID. You can call DescribeRegions to query available regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AddRCInstancesToDeploymentSetRequest) String() string {
	return dara.Prettify(s)
}

func (s AddRCInstancesToDeploymentSetRequest) GoString() string {
	return s.String()
}

func (s *AddRCInstancesToDeploymentSetRequest) GetDeploymentSetGroupNo() *string {
	return s.DeploymentSetGroupNo
}

func (s *AddRCInstancesToDeploymentSetRequest) GetDeploymentSetId() *string {
	return s.DeploymentSetId
}

func (s *AddRCInstancesToDeploymentSetRequest) GetForce() *bool {
	return s.Force
}

func (s *AddRCInstancesToDeploymentSetRequest) GetRCInstanceIds() *string {
	return s.RCInstanceIds
}

func (s *AddRCInstancesToDeploymentSetRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddRCInstancesToDeploymentSetRequest) SetDeploymentSetGroupNo(v string) *AddRCInstancesToDeploymentSetRequest {
	s.DeploymentSetGroupNo = &v
	return s
}

func (s *AddRCInstancesToDeploymentSetRequest) SetDeploymentSetId(v string) *AddRCInstancesToDeploymentSetRequest {
	s.DeploymentSetId = &v
	return s
}

func (s *AddRCInstancesToDeploymentSetRequest) SetForce(v bool) *AddRCInstancesToDeploymentSetRequest {
	s.Force = &v
	return s
}

func (s *AddRCInstancesToDeploymentSetRequest) SetRCInstanceIds(v string) *AddRCInstancesToDeploymentSetRequest {
	s.RCInstanceIds = &v
	return s
}

func (s *AddRCInstancesToDeploymentSetRequest) SetRegionId(v string) *AddRCInstancesToDeploymentSetRequest {
	s.RegionId = &v
	return s
}

func (s *AddRCInstancesToDeploymentSetRequest) Validate() error {
	return dara.Validate(s)
}
