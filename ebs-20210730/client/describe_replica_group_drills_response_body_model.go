// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeReplicaGroupDrillsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDrills(v []*DescribeReplicaGroupDrillsResponseBodyDrills) *DescribeReplicaGroupDrillsResponseBody
	GetDrills() []*DescribeReplicaGroupDrillsResponseBodyDrills
	SetNextToken(v string) *DescribeReplicaGroupDrillsResponseBody
	GetNextToken() *string
	SetPageNumber(v int32) *DescribeReplicaGroupDrillsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeReplicaGroupDrillsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeReplicaGroupDrillsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeReplicaGroupDrillsResponseBody
	GetTotalCount() *int64
}

type DescribeReplicaGroupDrillsResponseBody struct {
	// The information of disaster recovery drills that were performed on the replication pair-consistent group.
	Drills []*DescribeReplicaGroupDrillsResponseBodyDrills `json:"Drills,omitempty" xml:"Drills,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.
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
	// C123F94F-4E38-19AE-942A-A8D6F44F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeReplicaGroupDrillsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeReplicaGroupDrillsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeReplicaGroupDrillsResponseBody) GetDrills() []*DescribeReplicaGroupDrillsResponseBodyDrills {
	return s.Drills
}

func (s *DescribeReplicaGroupDrillsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeReplicaGroupDrillsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeReplicaGroupDrillsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeReplicaGroupDrillsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeReplicaGroupDrillsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeReplicaGroupDrillsResponseBody) SetDrills(v []*DescribeReplicaGroupDrillsResponseBodyDrills) *DescribeReplicaGroupDrillsResponseBody {
	s.Drills = v
	return s
}

func (s *DescribeReplicaGroupDrillsResponseBody) SetNextToken(v string) *DescribeReplicaGroupDrillsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeReplicaGroupDrillsResponseBody) SetPageNumber(v int32) *DescribeReplicaGroupDrillsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeReplicaGroupDrillsResponseBody) SetPageSize(v int32) *DescribeReplicaGroupDrillsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeReplicaGroupDrillsResponseBody) SetRequestId(v string) *DescribeReplicaGroupDrillsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeReplicaGroupDrillsResponseBody) SetTotalCount(v int64) *DescribeReplicaGroupDrillsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeReplicaGroupDrillsResponseBody) Validate() error {
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

type DescribeReplicaGroupDrillsResponseBodyDrills struct {
	// The ID of the drill.
	//
	// example:
	//
	// pg-drill-xxx
	DrillId *string `json:"DrillId,omitempty" xml:"DrillId,omitempty"`
	// The ID of the replication pair-consistent group.
	//
	// example:
	//
	// pg-xxx
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The information of replication pairs.
	PairsInfo []*DescribeReplicaGroupDrillsResponseBodyDrillsPairsInfo `json:"PairsInfo,omitempty" xml:"PairsInfo,omitempty" type:"Repeated"`
	// The recovery point of the drill. The value of this parameter is a timestamp. Unit: seconds.
	//
	// example:
	//
	// 1691114995
	RecoverPoint *int64 `json:"RecoverPoint,omitempty" xml:"RecoverPoint,omitempty"`
	// The beginning time of the drill. The value of this parameter is a timestamp. Unit: seconds.
	//
	// example:
	//
	// 1649750977
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
	// executed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The error message that appears if the drill fails to be executed.
	//
	// example:
	//
	// GROUP_SYNCPOINT_NOT_FOUND
	StatusMessage *string `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
}

func (s DescribeReplicaGroupDrillsResponseBodyDrills) String() string {
	return dara.Prettify(s)
}

func (s DescribeReplicaGroupDrillsResponseBodyDrills) GoString() string {
	return s.String()
}

func (s *DescribeReplicaGroupDrillsResponseBodyDrills) GetDrillId() *string {
	return s.DrillId
}

func (s *DescribeReplicaGroupDrillsResponseBodyDrills) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeReplicaGroupDrillsResponseBodyDrills) GetPairsInfo() []*DescribeReplicaGroupDrillsResponseBodyDrillsPairsInfo {
	return s.PairsInfo
}

func (s *DescribeReplicaGroupDrillsResponseBodyDrills) GetRecoverPoint() *int64 {
	return s.RecoverPoint
}

func (s *DescribeReplicaGroupDrillsResponseBodyDrills) GetStartAt() *int64 {
	return s.StartAt
}

func (s *DescribeReplicaGroupDrillsResponseBodyDrills) GetStatus() *string {
	return s.Status
}

func (s *DescribeReplicaGroupDrillsResponseBodyDrills) GetStatusMessage() *string {
	return s.StatusMessage
}

func (s *DescribeReplicaGroupDrillsResponseBodyDrills) SetDrillId(v string) *DescribeReplicaGroupDrillsResponseBodyDrills {
	s.DrillId = &v
	return s
}

func (s *DescribeReplicaGroupDrillsResponseBodyDrills) SetGroupId(v string) *DescribeReplicaGroupDrillsResponseBodyDrills {
	s.GroupId = &v
	return s
}

func (s *DescribeReplicaGroupDrillsResponseBodyDrills) SetPairsInfo(v []*DescribeReplicaGroupDrillsResponseBodyDrillsPairsInfo) *DescribeReplicaGroupDrillsResponseBodyDrills {
	s.PairsInfo = v
	return s
}

func (s *DescribeReplicaGroupDrillsResponseBodyDrills) SetRecoverPoint(v int64) *DescribeReplicaGroupDrillsResponseBodyDrills {
	s.RecoverPoint = &v
	return s
}

func (s *DescribeReplicaGroupDrillsResponseBodyDrills) SetStartAt(v int64) *DescribeReplicaGroupDrillsResponseBodyDrills {
	s.StartAt = &v
	return s
}

func (s *DescribeReplicaGroupDrillsResponseBodyDrills) SetStatus(v string) *DescribeReplicaGroupDrillsResponseBodyDrills {
	s.Status = &v
	return s
}

func (s *DescribeReplicaGroupDrillsResponseBodyDrills) SetStatusMessage(v string) *DescribeReplicaGroupDrillsResponseBodyDrills {
	s.StatusMessage = &v
	return s
}

func (s *DescribeReplicaGroupDrillsResponseBodyDrills) Validate() error {
	if s.PairsInfo != nil {
		for _, item := range s.PairsInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeReplicaGroupDrillsResponseBodyDrillsPairsInfo struct {
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
	// The ID of the replication pair.
	//
	// example:
	//
	// pair-xxx
	PairId *string `json:"PairId,omitempty" xml:"PairId,omitempty"`
}

func (s DescribeReplicaGroupDrillsResponseBodyDrillsPairsInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeReplicaGroupDrillsResponseBodyDrillsPairsInfo) GoString() string {
	return s.String()
}

func (s *DescribeReplicaGroupDrillsResponseBodyDrillsPairsInfo) GetDrillDiskId() *string {
	return s.DrillDiskId
}

func (s *DescribeReplicaGroupDrillsResponseBodyDrillsPairsInfo) GetDrillDiskStatus() *string {
	return s.DrillDiskStatus
}

func (s *DescribeReplicaGroupDrillsResponseBodyDrillsPairsInfo) GetPairId() *string {
	return s.PairId
}

func (s *DescribeReplicaGroupDrillsResponseBodyDrillsPairsInfo) SetDrillDiskId(v string) *DescribeReplicaGroupDrillsResponseBodyDrillsPairsInfo {
	s.DrillDiskId = &v
	return s
}

func (s *DescribeReplicaGroupDrillsResponseBodyDrillsPairsInfo) SetDrillDiskStatus(v string) *DescribeReplicaGroupDrillsResponseBodyDrillsPairsInfo {
	s.DrillDiskStatus = &v
	return s
}

func (s *DescribeReplicaGroupDrillsResponseBodyDrillsPairsInfo) SetPairId(v string) *DescribeReplicaGroupDrillsResponseBodyDrillsPairsInfo {
	s.PairId = &v
	return s
}

func (s *DescribeReplicaGroupDrillsResponseBodyDrillsPairsInfo) Validate() error {
	return dara.Validate(s)
}
