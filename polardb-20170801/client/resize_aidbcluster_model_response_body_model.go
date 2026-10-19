// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResizeAIDBClusterModelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentCapacity(v int64) *ResizeAIDBClusterModelResponseBody
	GetCurrentCapacity() *int64
	SetDBClusterId(v string) *ResizeAIDBClusterModelResponseBody
	GetDBClusterId() *string
	SetDryRun(v bool) *ResizeAIDBClusterModelResponseBody
	GetDryRun() *bool
	SetPoolHitCapacity(v int64) *ResizeAIDBClusterModelResponseBody
	GetPoolHitCapacity() *int64
	SetPoolHitRatio(v float64) *ResizeAIDBClusterModelResponseBody
	GetPoolHitRatio() *float64
	SetRequestId(v string) *ResizeAIDBClusterModelResponseBody
	GetRequestId() *string
	SetRequiredDelta(v int64) *ResizeAIDBClusterModelResponseBody
	GetRequiredDelta() *int64
	SetTargetCapacity(v int64) *ResizeAIDBClusterModelResponseBody
	GetTargetCapacity() *int64
}

type ResizeAIDBClusterModelResponseBody struct {
	// The current measured capacity of the instance. This parameter is returned only when DryRun is set to true.
	//
	// example:
	//
	// 8000
	CurrentCapacity *int64 `json:"CurrentCapacity,omitempty" xml:"CurrentCapacity,omitempty"`
	// The ID of the PolarDB AI 3.0 logical instance.
	//
	// example:
	//
	// pm-2ze4x2mwo81knj08a
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Indicates whether the current request performed only a capacity dry run.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The capacity that can be directly borrowed from the global pool without reshaping. This parameter is returned only when DryRun is set to true.
	//
	// example:
	//
	// 3000
	PoolHitCapacity *int64 `json:"PoolHitCapacity,omitempty" xml:"PoolHitCapacity,omitempty"`
	// The direct hit ratio of the global pool. Valid values: 0 to 1. This parameter is returned only when DryRun is set to true.
	//
	// example:
	//
	// 0.75
	PoolHitRatio *float64 `json:"PoolHitRatio,omitempty" xml:"PoolHitRatio,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 3AA6E0E4-1234-5678-90AB-1234567890AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The positive difference between the target capacity and the current capacity. This parameter is returned only when DryRun is set to true.
	//
	// example:
	//
	// 4000
	RequiredDelta *int64 `json:"RequiredDelta,omitempty" xml:"RequiredDelta,omitempty"`
	// The target total capacity.
	//
	// example:
	//
	// 12000
	TargetCapacity *int64 `json:"TargetCapacity,omitempty" xml:"TargetCapacity,omitempty"`
}

func (s ResizeAIDBClusterModelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResizeAIDBClusterModelResponseBody) GoString() string {
	return s.String()
}

func (s *ResizeAIDBClusterModelResponseBody) GetCurrentCapacity() *int64 {
	return s.CurrentCapacity
}

func (s *ResizeAIDBClusterModelResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ResizeAIDBClusterModelResponseBody) GetDryRun() *bool {
	return s.DryRun
}

func (s *ResizeAIDBClusterModelResponseBody) GetPoolHitCapacity() *int64 {
	return s.PoolHitCapacity
}

func (s *ResizeAIDBClusterModelResponseBody) GetPoolHitRatio() *float64 {
	return s.PoolHitRatio
}

func (s *ResizeAIDBClusterModelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResizeAIDBClusterModelResponseBody) GetRequiredDelta() *int64 {
	return s.RequiredDelta
}

func (s *ResizeAIDBClusterModelResponseBody) GetTargetCapacity() *int64 {
	return s.TargetCapacity
}

func (s *ResizeAIDBClusterModelResponseBody) SetCurrentCapacity(v int64) *ResizeAIDBClusterModelResponseBody {
	s.CurrentCapacity = &v
	return s
}

func (s *ResizeAIDBClusterModelResponseBody) SetDBClusterId(v string) *ResizeAIDBClusterModelResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *ResizeAIDBClusterModelResponseBody) SetDryRun(v bool) *ResizeAIDBClusterModelResponseBody {
	s.DryRun = &v
	return s
}

func (s *ResizeAIDBClusterModelResponseBody) SetPoolHitCapacity(v int64) *ResizeAIDBClusterModelResponseBody {
	s.PoolHitCapacity = &v
	return s
}

func (s *ResizeAIDBClusterModelResponseBody) SetPoolHitRatio(v float64) *ResizeAIDBClusterModelResponseBody {
	s.PoolHitRatio = &v
	return s
}

func (s *ResizeAIDBClusterModelResponseBody) SetRequestId(v string) *ResizeAIDBClusterModelResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResizeAIDBClusterModelResponseBody) SetRequiredDelta(v int64) *ResizeAIDBClusterModelResponseBody {
	s.RequiredDelta = &v
	return s
}

func (s *ResizeAIDBClusterModelResponseBody) SetTargetCapacity(v int64) *ResizeAIDBClusterModelResponseBody {
	s.TargetCapacity = &v
	return s
}

func (s *ResizeAIDBClusterModelResponseBody) Validate() error {
	return dara.Validate(s)
}
