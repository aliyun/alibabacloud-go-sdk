// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebootRCInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDryRun(v bool) *RebootRCInstanceRequest
	GetDryRun() *bool
	SetForceStop(v bool) *RebootRCInstanceRequest
	GetForceStop() *bool
	SetInstanceId(v string) *RebootRCInstanceRequest
	GetInstanceId() *string
	SetRebootTime(v string) *RebootRCInstanceRequest
	GetRebootTime() *string
	SetRegionId(v string) *RebootRCInstanceRequest
	GetRegionId() *string
}

type RebootRCInstanceRequest struct {
	DryRun    *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	ForceStop *bool `json:"ForceStop,omitempty" xml:"ForceStop,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 2018-01-01T12:05Z
	RebootTime *string `json:"RebootTime,omitempty" xml:"RebootTime,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s RebootRCInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RebootRCInstanceRequest) GoString() string {
	return s.String()
}

func (s *RebootRCInstanceRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *RebootRCInstanceRequest) GetForceStop() *bool {
	return s.ForceStop
}

func (s *RebootRCInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RebootRCInstanceRequest) GetRebootTime() *string {
	return s.RebootTime
}

func (s *RebootRCInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RebootRCInstanceRequest) SetDryRun(v bool) *RebootRCInstanceRequest {
	s.DryRun = &v
	return s
}

func (s *RebootRCInstanceRequest) SetForceStop(v bool) *RebootRCInstanceRequest {
	s.ForceStop = &v
	return s
}

func (s *RebootRCInstanceRequest) SetInstanceId(v string) *RebootRCInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *RebootRCInstanceRequest) SetRebootTime(v string) *RebootRCInstanceRequest {
	s.RebootTime = &v
	return s
}

func (s *RebootRCInstanceRequest) SetRegionId(v string) *RebootRCInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *RebootRCInstanceRequest) Validate() error {
	return dara.Validate(s)
}
