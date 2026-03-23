// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRCInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDryRun(v bool) *DeleteRCInstancesRequest
	GetDryRun() *bool
	SetForce(v bool) *DeleteRCInstancesRequest
	GetForce() *bool
	SetInstanceId(v []*string) *DeleteRCInstancesRequest
	GetInstanceId() []*string
	SetRegionId(v string) *DeleteRCInstancesRequest
	GetRegionId() *string
	SetTerminateSubscription(v bool) *DeleteRCInstancesRequest
	GetTerminateSubscription() *bool
}

type DeleteRCInstancesRequest struct {
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	Force  *bool `json:"Force,omitempty" xml:"Force,omitempty"`
	// This parameter is required.
	InstanceId            []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
	RegionId              *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TerminateSubscription *bool     `json:"TerminateSubscription,omitempty" xml:"TerminateSubscription,omitempty"`
}

func (s DeleteRCInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRCInstancesRequest) GoString() string {
	return s.String()
}

func (s *DeleteRCInstancesRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteRCInstancesRequest) GetForce() *bool {
	return s.Force
}

func (s *DeleteRCInstancesRequest) GetInstanceId() []*string {
	return s.InstanceId
}

func (s *DeleteRCInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteRCInstancesRequest) GetTerminateSubscription() *bool {
	return s.TerminateSubscription
}

func (s *DeleteRCInstancesRequest) SetDryRun(v bool) *DeleteRCInstancesRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteRCInstancesRequest) SetForce(v bool) *DeleteRCInstancesRequest {
	s.Force = &v
	return s
}

func (s *DeleteRCInstancesRequest) SetInstanceId(v []*string) *DeleteRCInstancesRequest {
	s.InstanceId = v
	return s
}

func (s *DeleteRCInstancesRequest) SetRegionId(v string) *DeleteRCInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteRCInstancesRequest) SetTerminateSubscription(v bool) *DeleteRCInstancesRequest {
	s.TerminateSubscription = &v
	return s
}

func (s *DeleteRCInstancesRequest) Validate() error {
	return dara.Validate(s)
}
