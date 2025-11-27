// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRCInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetForce(v bool) *DeleteRCInstanceRequest
	GetForce() *bool
	SetInstanceId(v string) *DeleteRCInstanceRequest
	GetInstanceId() *string
	SetRegionId(v string) *DeleteRCInstanceRequest
	GetRegionId() *string
}

type DeleteRCInstanceRequest struct {
	Force *bool `json:"Force,omitempty" xml:"Force,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteRCInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRCInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteRCInstanceRequest) GetForce() *bool {
	return s.Force
}

func (s *DeleteRCInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteRCInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteRCInstanceRequest) SetForce(v bool) *DeleteRCInstanceRequest {
	s.Force = &v
	return s
}

func (s *DeleteRCInstanceRequest) SetInstanceId(v string) *DeleteRCInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteRCInstanceRequest) SetRegionId(v string) *DeleteRCInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteRCInstanceRequest) Validate() error {
	return dara.Validate(s)
}
