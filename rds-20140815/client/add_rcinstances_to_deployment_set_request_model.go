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
	// The group number of the instance in the deployment set. This parameter is valid only when the deployment set strategy of the ECS instance to be added or adjusted is AvailabilityGroup. Valid values: 1 to 7.
	//
	// example:
	//
	// 1
	DeploymentSetGroupNo *string `json:"DeploymentSetGroupNo,omitempty" xml:"DeploymentSetGroupNo,omitempty"`
	// The ID of the deployment set.
	//
	// This parameter is required.
	//
	// example:
	//
	// ds-uf6c8qerk019bj1l****
	DeploymentSetId *string `json:"DeploymentSetId,omitempty" xml:"DeploymentSetId,omitempty"`
	// Specifies whether to forcibly release the running instance. Valid values:
	//
	// 	- **true**: yes.
	//
	// 	- **false*	- (default): no.
	//
	// example:
	//
	// false
	Force *bool `json:"Force,omitempty" xml:"Force,omitempty"`
	// The instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// rc-aaaa,rc-bbb
	RCInstanceIds *string `json:"RCInstanceIds,omitempty" xml:"RCInstanceIds,omitempty"`
	// The region ID. You can call the DescribeRegions operation to query the region ID.
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
