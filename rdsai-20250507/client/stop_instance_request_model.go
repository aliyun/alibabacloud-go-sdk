// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceName(v string) *StopInstanceRequest
	GetInstanceName() *string
	SetRegionId(v string) *StopInstanceRequest
	GetRegionId() *string
}

type StopInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ra-supabase-8moov5lxba****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
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

func (s *StopInstanceRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *StopInstanceRequest) GetRegionId() *string {
	return s.RegionId
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
