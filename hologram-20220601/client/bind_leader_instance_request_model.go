// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindLeaderInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *BindLeaderInstanceRequest
	GetRegionId() *string
	SetLeaderInstanceId(v string) *BindLeaderInstanceRequest
	GetLeaderInstanceId() *string
}

type BindLeaderInstanceRequest struct {
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The leader instance ID. The region ID, zone ID, and version of the follower instance must be the same as those of the leader instance. Select the leader instance ID from the list of instances returned by the ListFollowerInstances operation.
	//
	// example:
	//
	// hgpostcn-cn-uqm3316l1004
	LeaderInstanceId *string `json:"leaderInstanceId,omitempty" xml:"leaderInstanceId,omitempty"`
}

func (s BindLeaderInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s BindLeaderInstanceRequest) GoString() string {
	return s.String()
}

func (s *BindLeaderInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *BindLeaderInstanceRequest) GetLeaderInstanceId() *string {
	return s.LeaderInstanceId
}

func (s *BindLeaderInstanceRequest) SetRegionId(v string) *BindLeaderInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *BindLeaderInstanceRequest) SetLeaderInstanceId(v string) *BindLeaderInstanceRequest {
	s.LeaderInstanceId = &v
	return s
}

func (s *BindLeaderInstanceRequest) Validate() error {
	return dara.Validate(s)
}
