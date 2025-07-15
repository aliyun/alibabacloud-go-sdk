// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetForceDeleteInstance(v bool) *ReleaseInstanceRequest
	GetForceDeleteInstance() *bool
	SetInstanceId(v string) *ReleaseInstanceRequest
	GetInstanceId() *string
	SetRegionId(v string) *ReleaseInstanceRequest
	GetRegionId() *string
}

type ReleaseInstanceRequest struct {
	// Specifies whether to immediately release the physical resources of the instance. Valid values:
	//
	// 	- **true**: The physical resources of the instance are immediately released.
	//
	// 	- **false**: The physical resources of the instance are retained for a period of time before they are released.
	//
	// example:
	//
	// false
	ForceDeleteInstance *bool `json:"ForceDeleteInstance,omitempty" xml:"ForceDeleteInstance,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_post-cn-mp919o4v****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ReleaseInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ReleaseInstanceRequest) GoString() string {
	return s.String()
}

func (s *ReleaseInstanceRequest) GetForceDeleteInstance() *bool {
	return s.ForceDeleteInstance
}

func (s *ReleaseInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ReleaseInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ReleaseInstanceRequest) SetForceDeleteInstance(v bool) *ReleaseInstanceRequest {
	s.ForceDeleteInstance = &v
	return s
}

func (s *ReleaseInstanceRequest) SetInstanceId(v string) *ReleaseInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *ReleaseInstanceRequest) SetRegionId(v string) *ReleaseInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *ReleaseInstanceRequest) Validate() error {
	return dara.Validate(s)
}
