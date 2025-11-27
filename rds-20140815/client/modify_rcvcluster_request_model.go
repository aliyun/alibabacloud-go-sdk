// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRCVClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ModifyRCVClusterRequest
	GetClusterId() *string
	SetRegionId(v string) *ModifyRCVClusterRequest
	GetRegionId() *string
	SetSupportDiskPerformanceLevel(v []*string) *ModifyRCVClusterRequest
	GetSupportDiskPerformanceLevel() []*string
}

type ModifyRCVClusterRequest struct {
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
	RegionId                    *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SupportDiskPerformanceLevel []*string `json:"SupportDiskPerformanceLevel,omitempty" xml:"SupportDiskPerformanceLevel,omitempty" type:"Repeated"`
}

func (s ModifyRCVClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyRCVClusterRequest) GoString() string {
	return s.String()
}

func (s *ModifyRCVClusterRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ModifyRCVClusterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyRCVClusterRequest) GetSupportDiskPerformanceLevel() []*string {
	return s.SupportDiskPerformanceLevel
}

func (s *ModifyRCVClusterRequest) SetClusterId(v string) *ModifyRCVClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyRCVClusterRequest) SetRegionId(v string) *ModifyRCVClusterRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyRCVClusterRequest) SetSupportDiskPerformanceLevel(v []*string) *ModifyRCVClusterRequest {
	s.SupportDiskPerformanceLevel = v
	return s
}

func (s *ModifyRCVClusterRequest) Validate() error {
	return dara.Validate(s)
}
