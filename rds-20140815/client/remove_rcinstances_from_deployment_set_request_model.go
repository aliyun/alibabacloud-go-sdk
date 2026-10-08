// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveRCInstancesFromDeploymentSetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeploymentSetId(v string) *RemoveRCInstancesFromDeploymentSetRequest
	GetDeploymentSetId() *string
	SetRCInstanceIds(v string) *RemoveRCInstancesFromDeploymentSetRequest
	GetRCInstanceIds() *string
	SetRegionId(v string) *RemoveRCInstancesFromDeploymentSetRequest
	GetRegionId() *string
}

type RemoveRCInstancesFromDeploymentSetRequest struct {
	// The deployment set ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ds-uf6c8qerk019bj1l****
	DeploymentSetId *string `json:"DeploymentSetId,omitempty" xml:"DeploymentSetId,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rc-sff,rc-err
	RCInstanceIds *string `json:"RCInstanceIds,omitempty" xml:"RCInstanceIds,omitempty"`
	// The region ID. You can call DescribeRegions to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s RemoveRCInstancesFromDeploymentSetRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveRCInstancesFromDeploymentSetRequest) GoString() string {
	return s.String()
}

func (s *RemoveRCInstancesFromDeploymentSetRequest) GetDeploymentSetId() *string {
	return s.DeploymentSetId
}

func (s *RemoveRCInstancesFromDeploymentSetRequest) GetRCInstanceIds() *string {
	return s.RCInstanceIds
}

func (s *RemoveRCInstancesFromDeploymentSetRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RemoveRCInstancesFromDeploymentSetRequest) SetDeploymentSetId(v string) *RemoveRCInstancesFromDeploymentSetRequest {
	s.DeploymentSetId = &v
	return s
}

func (s *RemoveRCInstancesFromDeploymentSetRequest) SetRCInstanceIds(v string) *RemoveRCInstancesFromDeploymentSetRequest {
	s.RCInstanceIds = &v
	return s
}

func (s *RemoveRCInstancesFromDeploymentSetRequest) SetRegionId(v string) *RemoveRCInstancesFromDeploymentSetRequest {
	s.RegionId = &v
	return s
}

func (s *RemoveRCInstancesFromDeploymentSetRequest) Validate() error {
	return dara.Validate(s)
}
