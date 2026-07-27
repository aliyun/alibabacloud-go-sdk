// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBranchName(v string) *StopInstanceRequest
	GetBranchName() *string
	SetForce(v bool) *StopInstanceRequest
	GetForce() *bool
	SetInstanceName(v string) *StopInstanceRequest
	GetInstanceName() *string
	SetRegionId(v string) *StopInstanceRequest
	GetRegionId() *string
}

type StopInstanceRequest struct {
	BranchName *string `json:"BranchName,omitempty" xml:"BranchName,omitempty"`
	Force      *bool   `json:"Force,omitempty" xml:"Force,omitempty"`
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

func (s StopInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s StopInstanceRequest) GoString() string {
	return s.String()
}

func (s *StopInstanceRequest) GetBranchName() *string {
	return s.BranchName
}

func (s *StopInstanceRequest) GetForce() *bool {
	return s.Force
}

func (s *StopInstanceRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *StopInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StopInstanceRequest) SetBranchName(v string) *StopInstanceRequest {
	s.BranchName = &v
	return s
}

func (s *StopInstanceRequest) SetForce(v bool) *StopInstanceRequest {
	s.Force = &v
	return s
}

func (s *StopInstanceRequest) SetInstanceName(v string) *StopInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *StopInstanceRequest) SetRegionId(v string) *StopInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *StopInstanceRequest) Validate() error {
	return dara.Validate(s)
}
