// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRCVClusterShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ModifyRCVClusterShrinkRequest
	GetClusterId() *string
	SetRegionId(v string) *ModifyRCVClusterShrinkRequest
	GetRegionId() *string
	SetSupportDiskPerformanceLevelShrink(v string) *ModifyRCVClusterShrinkRequest
	GetSupportDiskPerformanceLevelShrink() *string
}

type ModifyRCVClusterShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cd21387ea640145bab79a78276c1a****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId                          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SupportDiskPerformanceLevelShrink *string `json:"SupportDiskPerformanceLevel,omitempty" xml:"SupportDiskPerformanceLevel,omitempty"`
}

func (s ModifyRCVClusterShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyRCVClusterShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyRCVClusterShrinkRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ModifyRCVClusterShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyRCVClusterShrinkRequest) GetSupportDiskPerformanceLevelShrink() *string {
	return s.SupportDiskPerformanceLevelShrink
}

func (s *ModifyRCVClusterShrinkRequest) SetClusterId(v string) *ModifyRCVClusterShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyRCVClusterShrinkRequest) SetRegionId(v string) *ModifyRCVClusterShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyRCVClusterShrinkRequest) SetSupportDiskPerformanceLevelShrink(v string) *ModifyRCVClusterShrinkRequest {
	s.SupportDiskPerformanceLevelShrink = &v
	return s
}

func (s *ModifyRCVClusterShrinkRequest) Validate() error {
	return dara.Validate(s)
}
