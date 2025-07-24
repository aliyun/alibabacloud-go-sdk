// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScalingAnalyzeResult interface {
	dara.Model
	String() string
	GoString() string
	SetActualUsage(v float32) *ScalingAnalyzeResult
	GetActualUsage() *float32
	SetIdealUsage(v float32) *ScalingAnalyzeResult
	GetIdealUsage() *float32
	SetReleaseCores(v int32) *ScalingAnalyzeResult
	GetReleaseCores() *int32
	SetReservedCores(v int32) *ScalingAnalyzeResult
	GetReservedCores() *int32
}

type ScalingAnalyzeResult struct {
	// 实际资源利用率。
	//
	// example:
	//
	// 0.12
	ActualUsage *float32 `json:"ActualUsage,omitempty" xml:"ActualUsage,omitempty"`
	// 理想资源用量。
	//
	// example:
	//
	// 0.12
	IdealUsage *float32 `json:"IdealUsage,omitempty" xml:"IdealUsage,omitempty"`
	// 固定资源释放核数（非master）core。
	//
	// example:
	//
	// 1
	ReleaseCores *int32 `json:"ReleaseCores,omitempty" xml:"ReleaseCores,omitempty"`
	// 固定资源保留核数（非master）core。
	//
	// example:
	//
	// 1
	ReservedCores *int32 `json:"ReservedCores,omitempty" xml:"ReservedCores,omitempty"`
}

func (s ScalingAnalyzeResult) String() string {
	return dara.Prettify(s)
}

func (s ScalingAnalyzeResult) GoString() string {
	return s.String()
}

func (s *ScalingAnalyzeResult) GetActualUsage() *float32 {
	return s.ActualUsage
}

func (s *ScalingAnalyzeResult) GetIdealUsage() *float32 {
	return s.IdealUsage
}

func (s *ScalingAnalyzeResult) GetReleaseCores() *int32 {
	return s.ReleaseCores
}

func (s *ScalingAnalyzeResult) GetReservedCores() *int32 {
	return s.ReservedCores
}

func (s *ScalingAnalyzeResult) SetActualUsage(v float32) *ScalingAnalyzeResult {
	s.ActualUsage = &v
	return s
}

func (s *ScalingAnalyzeResult) SetIdealUsage(v float32) *ScalingAnalyzeResult {
	s.IdealUsage = &v
	return s
}

func (s *ScalingAnalyzeResult) SetReleaseCores(v int32) *ScalingAnalyzeResult {
	s.ReleaseCores = &v
	return s
}

func (s *ScalingAnalyzeResult) SetReservedCores(v int32) *ScalingAnalyzeResult {
	s.ReservedCores = &v
	return s
}

func (s *ScalingAnalyzeResult) Validate() error {
	return dara.Validate(s)
}
