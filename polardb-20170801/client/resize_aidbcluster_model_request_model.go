// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResizeAIDBClusterModelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ResizeAIDBClusterModelRequest
	GetDBClusterId() *string
	SetDryRun(v bool) *ResizeAIDBClusterModelRequest
	GetDryRun() *bool
	SetRegionId(v string) *ResizeAIDBClusterModelRequest
	GetRegionId() *string
	SetTargetCapacity(v int64) *ResizeAIDBClusterModelRequest
	GetTargetCapacity() *int64
}

type ResizeAIDBClusterModelRequest struct {
	// The ID of the PolarDB AI 3.0 logical instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// pm-2ze4x2mwo81knj08a
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Specifies whether to perform only a read-only capacity dry run.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The region ID. Requests are routed to the region where the instance resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The target total capacity. An abstract capacity unit of 1000 is equivalent to one L20N card. The value must be a positive integer.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12000
	TargetCapacity *int64 `json:"TargetCapacity,omitempty" xml:"TargetCapacity,omitempty"`
}

func (s ResizeAIDBClusterModelRequest) String() string {
	return dara.Prettify(s)
}

func (s ResizeAIDBClusterModelRequest) GoString() string {
	return s.String()
}

func (s *ResizeAIDBClusterModelRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ResizeAIDBClusterModelRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ResizeAIDBClusterModelRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ResizeAIDBClusterModelRequest) GetTargetCapacity() *int64 {
	return s.TargetCapacity
}

func (s *ResizeAIDBClusterModelRequest) SetDBClusterId(v string) *ResizeAIDBClusterModelRequest {
	s.DBClusterId = &v
	return s
}

func (s *ResizeAIDBClusterModelRequest) SetDryRun(v bool) *ResizeAIDBClusterModelRequest {
	s.DryRun = &v
	return s
}

func (s *ResizeAIDBClusterModelRequest) SetRegionId(v string) *ResizeAIDBClusterModelRequest {
	s.RegionId = &v
	return s
}

func (s *ResizeAIDBClusterModelRequest) SetTargetCapacity(v int64) *ResizeAIDBClusterModelRequest {
	s.TargetCapacity = &v
	return s
}

func (s *ResizeAIDBClusterModelRequest) Validate() error {
	return dara.Validate(s)
}
