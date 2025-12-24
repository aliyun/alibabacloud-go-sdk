// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceSummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLockingCount(v int32) *GetInstanceSummaryResponseBody
	GetLockingCount() *int32
	SetRegionalSummary(v []*GetInstanceSummaryResponseBodyRegionalSummary) *GetInstanceSummaryResponseBody
	GetRegionalSummary() []*GetInstanceSummaryResponseBodyRegionalSummary
	SetRequestId(v string) *GetInstanceSummaryResponseBody
	GetRequestId() *string
	SetRunningCount(v int32) *GetInstanceSummaryResponseBody
	GetRunningCount() *int32
	SetTotal(v int32) *GetInstanceSummaryResponseBody
	GetTotal() *int32
}

type GetInstanceSummaryResponseBody struct {
	// example:
	//
	// 1
	LockingCount    *int32                                           `json:"LockingCount,omitempty" xml:"LockingCount,omitempty"`
	RegionalSummary []*GetInstanceSummaryResponseBodyRegionalSummary `json:"RegionalSummary,omitempty" xml:"RegionalSummary,omitempty" type:"Repeated"`
	// example:
	//
	// 1556DCB0-043A-4444-8BD9-CF4A68E7EE64
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 9
	RunningCount *int32 `json:"RunningCount,omitempty" xml:"RunningCount,omitempty"`
	// example:
	//
	// 10
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetInstanceSummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceSummaryResponseBody) GetLockingCount() *int32 {
	return s.LockingCount
}

func (s *GetInstanceSummaryResponseBody) GetRegionalSummary() []*GetInstanceSummaryResponseBodyRegionalSummary {
	return s.RegionalSummary
}

func (s *GetInstanceSummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstanceSummaryResponseBody) GetRunningCount() *int32 {
	return s.RunningCount
}

func (s *GetInstanceSummaryResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *GetInstanceSummaryResponseBody) SetLockingCount(v int32) *GetInstanceSummaryResponseBody {
	s.LockingCount = &v
	return s
}

func (s *GetInstanceSummaryResponseBody) SetRegionalSummary(v []*GetInstanceSummaryResponseBodyRegionalSummary) *GetInstanceSummaryResponseBody {
	s.RegionalSummary = v
	return s
}

func (s *GetInstanceSummaryResponseBody) SetRequestId(v string) *GetInstanceSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceSummaryResponseBody) SetRunningCount(v int32) *GetInstanceSummaryResponseBody {
	s.RunningCount = &v
	return s
}

func (s *GetInstanceSummaryResponseBody) SetTotal(v int32) *GetInstanceSummaryResponseBody {
	s.Total = &v
	return s
}

func (s *GetInstanceSummaryResponseBody) Validate() error {
	if s.RegionalSummary != nil {
		for _, item := range s.RegionalSummary {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetInstanceSummaryResponseBodyRegionalSummary struct {
	// example:
	//
	// 1
	LockingCount *int32 `json:"LockingCount,omitempty" xml:"LockingCount,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 5
	RunningCount *int32 `json:"RunningCount,omitempty" xml:"RunningCount,omitempty"`
	// example:
	//
	// 6
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetInstanceSummaryResponseBodyRegionalSummary) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceSummaryResponseBodyRegionalSummary) GoString() string {
	return s.String()
}

func (s *GetInstanceSummaryResponseBodyRegionalSummary) GetLockingCount() *int32 {
	return s.LockingCount
}

func (s *GetInstanceSummaryResponseBodyRegionalSummary) GetRegionId() *string {
	return s.RegionId
}

func (s *GetInstanceSummaryResponseBodyRegionalSummary) GetRunningCount() *int32 {
	return s.RunningCount
}

func (s *GetInstanceSummaryResponseBodyRegionalSummary) GetTotal() *int32 {
	return s.Total
}

func (s *GetInstanceSummaryResponseBodyRegionalSummary) SetLockingCount(v int32) *GetInstanceSummaryResponseBodyRegionalSummary {
	s.LockingCount = &v
	return s
}

func (s *GetInstanceSummaryResponseBodyRegionalSummary) SetRegionId(v string) *GetInstanceSummaryResponseBodyRegionalSummary {
	s.RegionId = &v
	return s
}

func (s *GetInstanceSummaryResponseBodyRegionalSummary) SetRunningCount(v int32) *GetInstanceSummaryResponseBodyRegionalSummary {
	s.RunningCount = &v
	return s
}

func (s *GetInstanceSummaryResponseBodyRegionalSummary) SetTotal(v int32) *GetInstanceSummaryResponseBodyRegionalSummary {
	s.Total = &v
	return s
}

func (s *GetInstanceSummaryResponseBodyRegionalSummary) Validate() error {
	return dara.Validate(s)
}
