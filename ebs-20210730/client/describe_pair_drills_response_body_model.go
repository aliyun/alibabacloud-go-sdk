// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePairDrillsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDrills(v []*DescribePairDrillsResponseBodyDrills) *DescribePairDrillsResponseBody
	GetDrills() []*DescribePairDrillsResponseBodyDrills
	SetNextToken(v string) *DescribePairDrillsResponseBody
	GetNextToken() *string
	SetPageNumber(v int32) *DescribePairDrillsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribePairDrillsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribePairDrillsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribePairDrillsResponseBody
	GetTotalCount() *int64
}

type DescribePairDrillsResponseBody struct {
	// The information of disaster recovery drills that were performed on the replication pair.
	Drills []*DescribePairDrillsResponseBodyDrills `json:"Drills,omitempty" xml:"Drills,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// AAAAAdDWBF2****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C46FF5A8-C5F0-4024-8262-B16B6392****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribePairDrillsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePairDrillsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePairDrillsResponseBody) GetDrills() []*DescribePairDrillsResponseBodyDrills {
	return s.Drills
}

func (s *DescribePairDrillsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribePairDrillsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribePairDrillsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribePairDrillsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePairDrillsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribePairDrillsResponseBody) SetDrills(v []*DescribePairDrillsResponseBodyDrills) *DescribePairDrillsResponseBody {
	s.Drills = v
	return s
}

func (s *DescribePairDrillsResponseBody) SetNextToken(v string) *DescribePairDrillsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribePairDrillsResponseBody) SetPageNumber(v int32) *DescribePairDrillsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribePairDrillsResponseBody) SetPageSize(v int32) *DescribePairDrillsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribePairDrillsResponseBody) SetRequestId(v string) *DescribePairDrillsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePairDrillsResponseBody) SetTotalCount(v int64) *DescribePairDrillsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribePairDrillsResponseBody) Validate() error {
	if s.Drills != nil {
		for _, item := range s.Drills {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePairDrillsResponseBodyDrills struct {
	// The ID of the drill disk.
	//
	// example:
	//
	// d-xxx
	DrillDiskId *string `json:"DrillDiskId,omitempty" xml:"DrillDiskId,omitempty"`
	// The status of the drill disk. Valid values:
	//
	// 	- created
	//
	// 	- deleted
	//
	// 	- creating
	//
	// 	- deleting
	//
	// >  This parameter can also display error code details if your drill disk fails to be created or deleted.
	//
	// example:
	//
	// created
	DrillDiskStatus *string `json:"DrillDiskStatus,omitempty" xml:"DrillDiskStatus,omitempty"`
	// The ID of the drill.
	//
	// example:
	//
	// drill-xxx
	DrillId *string `json:"DrillId,omitempty" xml:"DrillId,omitempty"`
	// The recovery point of the drill. The value of this parameter is a timestamp. Unit: seconds.
	//
	// example:
	//
	// 1690855931
	RecoverPoint *int64 `json:"RecoverPoint,omitempty" xml:"RecoverPoint,omitempty"`
	// The beginning time of the drill. The value of this parameter is a timestamp. Unit: seconds.
	//
	// example:
	//
	// 1690855888
	StartAt *int64 `json:"StartAt,omitempty" xml:"StartAt,omitempty"`
	// The status of the drill. Valid values:
	//
	// 	- execute_failed
	//
	// 	- executed
	//
	// 	- executing
	//
	// 	- clear_failed
	//
	// 	- clearing
	//
	// example:
	//
	// executing
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The error message that was displayed if the drill failed to be executed.
	//
	// example:
	//
	// PAIR_SYNCPOINT_NOT_FOUND
	StatusMessage *string `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
}

func (s DescribePairDrillsResponseBodyDrills) String() string {
	return dara.Prettify(s)
}

func (s DescribePairDrillsResponseBodyDrills) GoString() string {
	return s.String()
}

func (s *DescribePairDrillsResponseBodyDrills) GetDrillDiskId() *string {
	return s.DrillDiskId
}

func (s *DescribePairDrillsResponseBodyDrills) GetDrillDiskStatus() *string {
	return s.DrillDiskStatus
}

func (s *DescribePairDrillsResponseBodyDrills) GetDrillId() *string {
	return s.DrillId
}

func (s *DescribePairDrillsResponseBodyDrills) GetRecoverPoint() *int64 {
	return s.RecoverPoint
}

func (s *DescribePairDrillsResponseBodyDrills) GetStartAt() *int64 {
	return s.StartAt
}

func (s *DescribePairDrillsResponseBodyDrills) GetStatus() *string {
	return s.Status
}

func (s *DescribePairDrillsResponseBodyDrills) GetStatusMessage() *string {
	return s.StatusMessage
}

func (s *DescribePairDrillsResponseBodyDrills) SetDrillDiskId(v string) *DescribePairDrillsResponseBodyDrills {
	s.DrillDiskId = &v
	return s
}

func (s *DescribePairDrillsResponseBodyDrills) SetDrillDiskStatus(v string) *DescribePairDrillsResponseBodyDrills {
	s.DrillDiskStatus = &v
	return s
}

func (s *DescribePairDrillsResponseBodyDrills) SetDrillId(v string) *DescribePairDrillsResponseBodyDrills {
	s.DrillId = &v
	return s
}

func (s *DescribePairDrillsResponseBodyDrills) SetRecoverPoint(v int64) *DescribePairDrillsResponseBodyDrills {
	s.RecoverPoint = &v
	return s
}

func (s *DescribePairDrillsResponseBodyDrills) SetStartAt(v int64) *DescribePairDrillsResponseBodyDrills {
	s.StartAt = &v
	return s
}

func (s *DescribePairDrillsResponseBodyDrills) SetStatus(v string) *DescribePairDrillsResponseBodyDrills {
	s.Status = &v
	return s
}

func (s *DescribePairDrillsResponseBodyDrills) SetStatusMessage(v string) *DescribePairDrillsResponseBodyDrills {
	s.StatusMessage = &v
	return s
}

func (s *DescribePairDrillsResponseBodyDrills) Validate() error {
	return dara.Validate(s)
}
