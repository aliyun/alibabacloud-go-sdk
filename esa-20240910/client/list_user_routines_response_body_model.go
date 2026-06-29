// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserRoutinesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int64) *ListUserRoutinesResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *ListUserRoutinesResponseBody
	GetPageSize() *int64
	SetQuotaRoutineNumber(v int64) *ListUserRoutinesResponseBody
	GetQuotaRoutineNumber() *int64
	SetRequestId(v string) *ListUserRoutinesResponseBody
	GetRequestId() *string
	SetRoutines(v []*ListUserRoutinesResponseBodyRoutines) *ListUserRoutinesResponseBody
	GetRoutines() []*ListUserRoutinesResponseBodyRoutines
	SetTotalCount(v int64) *ListUserRoutinesResponseBody
	GetTotalCount() *int64
	SetUsedRoutineNumber(v int64) *ListUserRoutinesResponseBody
	GetUsedRoutineNumber() *int64
}

type ListUserRoutinesResponseBody struct {
	// The current page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The Routine quota for the current plan.
	//
	// example:
	//
	// 20
	QuotaRoutineNumber *int64 `json:"QuotaRoutineNumber,omitempty" xml:"QuotaRoutineNumber,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1234567890ABCDEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of Routines.
	Routines []*ListUserRoutinesResponseBodyRoutines `json:"Routines,omitempty" xml:"Routines,omitempty" type:"Repeated"`
	// The total number of entries.
	//
	// example:
	//
	// 3
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The number of Routines already created.
	//
	// example:
	//
	// 5
	UsedRoutineNumber *int64 `json:"UsedRoutineNumber,omitempty" xml:"UsedRoutineNumber,omitempty"`
}

func (s ListUserRoutinesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUserRoutinesResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserRoutinesResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListUserRoutinesResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListUserRoutinesResponseBody) GetQuotaRoutineNumber() *int64 {
	return s.QuotaRoutineNumber
}

func (s *ListUserRoutinesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUserRoutinesResponseBody) GetRoutines() []*ListUserRoutinesResponseBodyRoutines {
	return s.Routines
}

func (s *ListUserRoutinesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListUserRoutinesResponseBody) GetUsedRoutineNumber() *int64 {
	return s.UsedRoutineNumber
}

func (s *ListUserRoutinesResponseBody) SetPageNumber(v int64) *ListUserRoutinesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListUserRoutinesResponseBody) SetPageSize(v int64) *ListUserRoutinesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListUserRoutinesResponseBody) SetQuotaRoutineNumber(v int64) *ListUserRoutinesResponseBody {
	s.QuotaRoutineNumber = &v
	return s
}

func (s *ListUserRoutinesResponseBody) SetRequestId(v string) *ListUserRoutinesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserRoutinesResponseBody) SetRoutines(v []*ListUserRoutinesResponseBodyRoutines) *ListUserRoutinesResponseBody {
	s.Routines = v
	return s
}

func (s *ListUserRoutinesResponseBody) SetTotalCount(v int64) *ListUserRoutinesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListUserRoutinesResponseBody) SetUsedRoutineNumber(v int64) *ListUserRoutinesResponseBody {
	s.UsedRoutineNumber = &v
	return s
}

func (s *ListUserRoutinesResponseBody) Validate() error {
	if s.Routines != nil {
		for _, item := range s.Routines {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListUserRoutinesResponseBodyRoutines struct {
	// The time when the Edge Routine was created. The time follows the RFC 3339 standard in the UTC time zone.
	//
	// example:
	//
	// 2024-03-11T01:23:02.883361712Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The default access record associated with the Routine.
	//
	// example:
	//
	// serverless-test-2.154edaf6.er.aliyun-esa.net
	DefaultRelatedRecord *string `json:"DefaultRelatedRecord,omitempty" xml:"DefaultRelatedRecord,omitempty"`
	// The Routine description.
	//
	// example:
	//
	// ZWRpdCByb3V0aW5lIGNvbmZpZyBkZXNjcmlwdGlvbg==
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether the Routine contains asset files.
	//
	// example:
	//
	// false
	HasAssets *bool `json:"HasAssets,omitempty" xml:"HasAssets,omitempty"`
	// The Routine name.
	//
	// example:
	//
	// hello
	RoutineName *string `json:"RoutineName,omitempty" xml:"RoutineName,omitempty"`
}

func (s ListUserRoutinesResponseBodyRoutines) String() string {
	return dara.Prettify(s)
}

func (s ListUserRoutinesResponseBodyRoutines) GoString() string {
	return s.String()
}

func (s *ListUserRoutinesResponseBodyRoutines) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListUserRoutinesResponseBodyRoutines) GetDefaultRelatedRecord() *string {
	return s.DefaultRelatedRecord
}

func (s *ListUserRoutinesResponseBodyRoutines) GetDescription() *string {
	return s.Description
}

func (s *ListUserRoutinesResponseBodyRoutines) GetHasAssets() *bool {
	return s.HasAssets
}

func (s *ListUserRoutinesResponseBodyRoutines) GetRoutineName() *string {
	return s.RoutineName
}

func (s *ListUserRoutinesResponseBodyRoutines) SetCreateTime(v string) *ListUserRoutinesResponseBodyRoutines {
	s.CreateTime = &v
	return s
}

func (s *ListUserRoutinesResponseBodyRoutines) SetDefaultRelatedRecord(v string) *ListUserRoutinesResponseBodyRoutines {
	s.DefaultRelatedRecord = &v
	return s
}

func (s *ListUserRoutinesResponseBodyRoutines) SetDescription(v string) *ListUserRoutinesResponseBodyRoutines {
	s.Description = &v
	return s
}

func (s *ListUserRoutinesResponseBodyRoutines) SetHasAssets(v bool) *ListUserRoutinesResponseBodyRoutines {
	s.HasAssets = &v
	return s
}

func (s *ListUserRoutinesResponseBodyRoutines) SetRoutineName(v string) *ListUserRoutinesResponseBodyRoutines {
	s.RoutineName = &v
	return s
}

func (s *ListUserRoutinesResponseBodyRoutines) Validate() error {
	return dara.Validate(s)
}
