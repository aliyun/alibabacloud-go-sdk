// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceName(v string) *RestartInstanceRequest
	GetInstanceName() *string
	SetRegionId(v string) *RestartInstanceRequest
	GetRegionId() *string
}

type RestartInstanceRequest struct {
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ra-supabase-8moov5lxba****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The operation that you want to perform. Set the value to **RestartInstance**.
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

func (s *RestartInstanceRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *RestartInstanceRequest) GetRegionId() *string {
	return s.RegionId
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
