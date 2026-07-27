// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBranchName(v string) *RestartInstanceRequest
	GetBranchName() *string
	SetInstanceName(v string) *RestartInstanceRequest
	GetInstanceName() *string
	SetRegionId(v string) *RestartInstanceRequest
	GetRegionId() *string
}

type RestartInstanceRequest struct {
	BranchName *string `json:"BranchName,omitempty" xml:"BranchName,omitempty"`
	// The instance ID of the AI application.
	//
	// This parameter is required.
	//
	// example:
	//
	// ra-supabase-8moov5lxba****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s RestartInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RestartInstanceRequest) GoString() string {
	return s.String()
}

func (s *RestartInstanceRequest) GetBranchName() *string {
	return s.BranchName
}

func (s *RestartInstanceRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *RestartInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RestartInstanceRequest) SetBranchName(v string) *RestartInstanceRequest {
	s.BranchName = &v
	return s
}

func (s *RestartInstanceRequest) SetInstanceName(v string) *RestartInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *RestartInstanceRequest) SetRegionId(v string) *RestartInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *RestartInstanceRequest) Validate() error {
	return dara.Validate(s)
}
