// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAttackProtectionCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBlockHighCount(v int64) *DescribeAttackProtectionCountResponseBody
	GetBlockHighCount() *int64
	SetBlockLowCount(v int64) *DescribeAttackProtectionCountResponseBody
	GetBlockLowCount() *int64
	SetBlockMediumCount(v int64) *DescribeAttackProtectionCountResponseBody
	GetBlockMediumCount() *int64
	SetMonitorHighCount(v int64) *DescribeAttackProtectionCountResponseBody
	GetMonitorHighCount() *int64
	SetMonitorLowCount(v int64) *DescribeAttackProtectionCountResponseBody
	GetMonitorLowCount() *int64
	SetMonitorMediumCount(v int64) *DescribeAttackProtectionCountResponseBody
	GetMonitorMediumCount() *int64
	SetRequestId(v string) *DescribeAttackProtectionCountResponseBody
	GetRequestId() *string
	SetTotalRequestCount(v int64) *DescribeAttackProtectionCountResponseBody
	GetTotalRequestCount() *int64
}

type DescribeAttackProtectionCountResponseBody struct {
	// example:
	//
	// 0
	BlockHighCount *int64 `json:"BlockHighCount,omitempty" xml:"BlockHighCount,omitempty"`
	// example:
	//
	// 0
	BlockLowCount *int64 `json:"BlockLowCount,omitempty" xml:"BlockLowCount,omitempty"`
	// example:
	//
	// 0
	BlockMediumCount *int64 `json:"BlockMediumCount,omitempty" xml:"BlockMediumCount,omitempty"`
	// example:
	//
	// 0
	MonitorHighCount *int64 `json:"MonitorHighCount,omitempty" xml:"MonitorHighCount,omitempty"`
	// example:
	//
	// 0
	MonitorLowCount *int64 `json:"MonitorLowCount,omitempty" xml:"MonitorLowCount,omitempty"`
	// example:
	//
	// 0
	MonitorMediumCount *int64 `json:"MonitorMediumCount,omitempty" xml:"MonitorMediumCount,omitempty"`
	// example:
	//
	// 400E8C8C-ADD3-5F25-9038-BDC057841D20
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 61134279
	TotalRequestCount *int64 `json:"TotalRequestCount,omitempty" xml:"TotalRequestCount,omitempty"`
}

func (s DescribeAttackProtectionCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAttackProtectionCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAttackProtectionCountResponseBody) GetBlockHighCount() *int64 {
	return s.BlockHighCount
}

func (s *DescribeAttackProtectionCountResponseBody) GetBlockLowCount() *int64 {
	return s.BlockLowCount
}

func (s *DescribeAttackProtectionCountResponseBody) GetBlockMediumCount() *int64 {
	return s.BlockMediumCount
}

func (s *DescribeAttackProtectionCountResponseBody) GetMonitorHighCount() *int64 {
	return s.MonitorHighCount
}

func (s *DescribeAttackProtectionCountResponseBody) GetMonitorLowCount() *int64 {
	return s.MonitorLowCount
}

func (s *DescribeAttackProtectionCountResponseBody) GetMonitorMediumCount() *int64 {
	return s.MonitorMediumCount
}

func (s *DescribeAttackProtectionCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAttackProtectionCountResponseBody) GetTotalRequestCount() *int64 {
	return s.TotalRequestCount
}

func (s *DescribeAttackProtectionCountResponseBody) SetBlockHighCount(v int64) *DescribeAttackProtectionCountResponseBody {
	s.BlockHighCount = &v
	return s
}

func (s *DescribeAttackProtectionCountResponseBody) SetBlockLowCount(v int64) *DescribeAttackProtectionCountResponseBody {
	s.BlockLowCount = &v
	return s
}

func (s *DescribeAttackProtectionCountResponseBody) SetBlockMediumCount(v int64) *DescribeAttackProtectionCountResponseBody {
	s.BlockMediumCount = &v
	return s
}

func (s *DescribeAttackProtectionCountResponseBody) SetMonitorHighCount(v int64) *DescribeAttackProtectionCountResponseBody {
	s.MonitorHighCount = &v
	return s
}

func (s *DescribeAttackProtectionCountResponseBody) SetMonitorLowCount(v int64) *DescribeAttackProtectionCountResponseBody {
	s.MonitorLowCount = &v
	return s
}

func (s *DescribeAttackProtectionCountResponseBody) SetMonitorMediumCount(v int64) *DescribeAttackProtectionCountResponseBody {
	s.MonitorMediumCount = &v
	return s
}

func (s *DescribeAttackProtectionCountResponseBody) SetRequestId(v string) *DescribeAttackProtectionCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAttackProtectionCountResponseBody) SetTotalRequestCount(v int64) *DescribeAttackProtectionCountResponseBody {
	s.TotalRequestCount = &v
	return s
}

func (s *DescribeAttackProtectionCountResponseBody) Validate() error {
	return dara.Validate(s)
}
