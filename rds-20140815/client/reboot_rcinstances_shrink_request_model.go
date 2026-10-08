// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebootRCInstancesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBatchOptimization(v string) *RebootRCInstancesShrinkRequest
	GetBatchOptimization() *string
	SetForceReboot(v bool) *RebootRCInstancesShrinkRequest
	GetForceReboot() *bool
	SetInstanceIdsShrink(v string) *RebootRCInstancesShrinkRequest
	GetInstanceIdsShrink() *string
	SetRebootTime(v string) *RebootRCInstancesShrinkRequest
	GetRebootTime() *string
	SetRegionId(v string) *RebootRCInstancesShrinkRequest
	GetRegionId() *string
}

type RebootRCInstancesShrinkRequest struct {
	BatchOptimization *string `json:"BatchOptimization,omitempty" xml:"BatchOptimization,omitempty"`
	ForceReboot       *bool   `json:"ForceReboot,omitempty" xml:"ForceReboot,omitempty"`
	InstanceIdsShrink *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// example:
	//
	// 2018-01-01T12:05Z
	RebootTime *string `json:"RebootTime,omitempty" xml:"RebootTime,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s RebootRCInstancesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RebootRCInstancesShrinkRequest) GoString() string {
	return s.String()
}

func (s *RebootRCInstancesShrinkRequest) GetBatchOptimization() *string {
	return s.BatchOptimization
}

func (s *RebootRCInstancesShrinkRequest) GetForceReboot() *bool {
	return s.ForceReboot
}

func (s *RebootRCInstancesShrinkRequest) GetInstanceIdsShrink() *string {
	return s.InstanceIdsShrink
}

func (s *RebootRCInstancesShrinkRequest) GetRebootTime() *string {
	return s.RebootTime
}

func (s *RebootRCInstancesShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RebootRCInstancesShrinkRequest) SetBatchOptimization(v string) *RebootRCInstancesShrinkRequest {
	s.BatchOptimization = &v
	return s
}

func (s *RebootRCInstancesShrinkRequest) SetForceReboot(v bool) *RebootRCInstancesShrinkRequest {
	s.ForceReboot = &v
	return s
}

func (s *RebootRCInstancesShrinkRequest) SetInstanceIdsShrink(v string) *RebootRCInstancesShrinkRequest {
	s.InstanceIdsShrink = &v
	return s
}

func (s *RebootRCInstancesShrinkRequest) SetRebootTime(v string) *RebootRCInstancesShrinkRequest {
	s.RebootTime = &v
	return s
}

func (s *RebootRCInstancesShrinkRequest) SetRegionId(v string) *RebootRCInstancesShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *RebootRCInstancesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
