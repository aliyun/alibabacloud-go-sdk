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
	PageNumber         *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize           *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	QuotaRoutineNumber *int64 `json:"QuotaRoutineNumber,omitempty" xml:"QuotaRoutineNumber,omitempty"`
	// Id of the request
	RequestId         *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Routines          []*ListUserRoutinesResponseBodyRoutines `json:"Routines,omitempty" xml:"Routines,omitempty" type:"Repeated"`
	TotalCount        *int64                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	UsedRoutineNumber *int64                                  `json:"UsedRoutineNumber,omitempty" xml:"UsedRoutineNumber,omitempty"`
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
	return dara.Validate(s)
}

type ListUserRoutinesResponseBodyRoutines struct {
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
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

func (s *ListUserRoutinesResponseBodyRoutines) GetDescription() *string {
	return s.Description
}

func (s *ListUserRoutinesResponseBodyRoutines) GetRoutineName() *string {
	return s.RoutineName
}

func (s *ListUserRoutinesResponseBodyRoutines) SetCreateTime(v string) *ListUserRoutinesResponseBodyRoutines {
	s.CreateTime = &v
	return s
}

func (s *ListUserRoutinesResponseBodyRoutines) SetDescription(v string) *ListUserRoutinesResponseBodyRoutines {
	s.Description = &v
	return s
}

func (s *ListUserRoutinesResponseBodyRoutines) SetRoutineName(v string) *ListUserRoutinesResponseBodyRoutines {
	s.RoutineName = &v
	return s
}

func (s *ListUserRoutinesResponseBodyRoutines) Validate() error {
	return dara.Validate(s)
}
