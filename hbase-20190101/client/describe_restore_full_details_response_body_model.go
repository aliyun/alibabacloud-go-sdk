// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRestoreFullDetailsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeRestoreFullDetailsResponseBody
	GetRequestId() *string
	SetRestoreFull(v *DescribeRestoreFullDetailsResponseBodyRestoreFull) *DescribeRestoreFullDetailsResponseBody
	GetRestoreFull() *DescribeRestoreFullDetailsResponseBodyRestoreFull
}

type DescribeRestoreFullDetailsResponseBody struct {
	// example:
	//
	// CFE525CF-C691-4140-A981-D004DAA7A840
	RequestId   *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RestoreFull *DescribeRestoreFullDetailsResponseBodyRestoreFull `json:"RestoreFull,omitempty" xml:"RestoreFull,omitempty" type:"Struct"`
}

func (s DescribeRestoreFullDetailsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRestoreFullDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRestoreFullDetailsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRestoreFullDetailsResponseBody) GetRestoreFull() *DescribeRestoreFullDetailsResponseBodyRestoreFull {
	return s.RestoreFull
}

func (s *DescribeRestoreFullDetailsResponseBody) SetRequestId(v string) *DescribeRestoreFullDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRestoreFullDetailsResponseBody) SetRestoreFull(v *DescribeRestoreFullDetailsResponseBodyRestoreFull) *DescribeRestoreFullDetailsResponseBody {
	s.RestoreFull = v
	return s
}

func (s *DescribeRestoreFullDetailsResponseBody) Validate() error {
	if s.RestoreFull != nil {
		if err := s.RestoreFull.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRestoreFullDetailsResponseBodyRestoreFull struct {
	// example:
	//
	// 1.2 kB
	DataSize *string `json:"DataSize,omitempty" xml:"DataSize,omitempty"`
	// example:
	//
	// 0
	Fail *int32 `json:"Fail,omitempty" xml:"Fail,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize           *int32                                                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RestoreFullDetails *DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetails `json:"RestoreFullDetails,omitempty" xml:"RestoreFullDetails,omitempty" type:"Struct"`
	// example:
	//
	// 0.00 MB/s
	Speed *string `json:"Speed,omitempty" xml:"Speed,omitempty"`
	// example:
	//
	// 1
	Succeed *int32 `json:"Succeed,omitempty" xml:"Succeed,omitempty"`
	// example:
	//
	// 1
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeRestoreFullDetailsResponseBodyRestoreFull) String() string {
	return dara.Prettify(s)
}

func (s DescribeRestoreFullDetailsResponseBodyRestoreFull) GoString() string {
	return s.String()
}

func (s *DescribeRestoreFullDetailsResponseBodyRestoreFull) GetDataSize() *string {
	return s.DataSize
}

func (s *DescribeRestoreFullDetailsResponseBodyRestoreFull) GetFail() *int32 {
	return s.Fail
}

func (s *DescribeRestoreFullDetailsResponseBodyRestoreFull) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeRestoreFullDetailsResponseBodyRestoreFull) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeRestoreFullDetailsResponseBodyRestoreFull) GetRestoreFullDetails() *DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetails {
	return s.RestoreFullDetails
}

func (s *DescribeRestoreFullDetailsResponseBodyRestoreFull) GetSpeed() *string {
	return s.Speed
}

func (s *DescribeRestoreFullDetailsResponseBodyRestoreFull) GetSucceed() *int32 {
	return s.Succeed
}

func (s *DescribeRestoreFullDetailsResponseBodyRestoreFull) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeRestoreFullDetailsResponseBodyRestoreFull) SetDataSize(v string) *DescribeRestoreFullDetailsResponseBodyRestoreFull {
	s.DataSize = &v
	return s
}

func (s *DescribeRestoreFullDetailsResponseBodyRestoreFull) SetFail(v int32) *DescribeRestoreFullDetailsResponseBodyRestoreFull {
	s.Fail = &v
	return s
}

func (s *DescribeRestoreFullDetailsResponseBodyRestoreFull) SetPageNumber(v int32) *DescribeRestoreFullDetailsResponseBodyRestoreFull {
	s.PageNumber = &v
	return s
}

func (s *DescribeRestoreFullDetailsResponseBodyRestoreFull) SetPageSize(v int32) *DescribeRestoreFullDetailsResponseBodyRestoreFull {
	s.PageSize = &v
	return s
}

func (s *DescribeRestoreFullDetailsResponseBodyRestoreFull) SetRestoreFullDetails(v *DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetails) *DescribeRestoreFullDetailsResponseBodyRestoreFull {
	s.RestoreFullDetails = v
	return s
}

func (s *DescribeRestoreFullDetailsResponseBodyRestoreFull) SetSpeed(v string) *DescribeRestoreFullDetailsResponseBodyRestoreFull {
	s.Speed = &v
	return s
}

func (s *DescribeRestoreFullDetailsResponseBodyRestoreFull) SetSucceed(v int32) *DescribeRestoreFullDetailsResponseBodyRestoreFull {
	s.Succeed = &v
	return s
}

func (s *DescribeRestoreFullDetailsResponseBodyRestoreFull) SetTotal(v int64) *DescribeRestoreFullDetailsResponseBodyRestoreFull {
	s.Total = &v
	return s
}

func (s *DescribeRestoreFullDetailsResponseBodyRestoreFull) Validate() error {
	if s.RestoreFullDetails != nil {
		if err := s.RestoreFullDetails.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetails struct {
	RestoreFullDetail []*DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail `json:"RestoreFullDetail,omitempty" xml:"RestoreFullDetail,omitempty" type:"Repeated"`
}

func (s DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetails) String() string {
	return dara.Prettify(s)
}

func (s DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetails) GoString() string {
	return s.String()
}

func (s *DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetails) GetRestoreFullDetail() []*DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail {
	return s.RestoreFullDetail
}

func (s *DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetails) SetRestoreFullDetail(v []*DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail) *DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetails {
	s.RestoreFullDetail = v
	return s
}

func (s *DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetails) Validate() error {
	if s.RestoreFullDetail != nil {
		for _, item := range s.RestoreFullDetail {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail struct {
	// example:
	//
	// 1.2 kB
	DataSize *string `json:"DataSize,omitempty" xml:"DataSize,omitempty"`
	// example:
	//
	// 2020-11-05T06:45:51Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// null
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 14/14
	Process *string `json:"Process,omitempty" xml:"Process,omitempty"`
	// example:
	//
	// 0.00 MB/s
	Speed *string `json:"Speed,omitempty" xml:"Speed,omitempty"`
	// example:
	//
	// 2020-11-05T06:45:45Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// SUCCEEDED
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// example:
	//
	// default:test1
	Table *string `json:"Table,omitempty" xml:"Table,omitempty"`
}

func (s DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail) GoString() string {
	return s.String()
}

func (s *DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail) GetDataSize() *string {
	return s.DataSize
}

func (s *DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail) GetMessage() *string {
	return s.Message
}

func (s *DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail) GetProcess() *string {
	return s.Process
}

func (s *DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail) GetSpeed() *string {
	return s.Speed
}

func (s *DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail) GetState() *string {
	return s.State
}

func (s *DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail) GetTable() *string {
	return s.Table
}

func (s *DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail) SetDataSize(v string) *DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail {
	s.DataSize = &v
	return s
}

func (s *DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail) SetEndTime(v string) *DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail {
	s.EndTime = &v
	return s
}

func (s *DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail) SetMessage(v string) *DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail {
	s.Message = &v
	return s
}

func (s *DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail) SetProcess(v string) *DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail {
	s.Process = &v
	return s
}

func (s *DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail) SetSpeed(v string) *DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail {
	s.Speed = &v
	return s
}

func (s *DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail) SetStartTime(v string) *DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail {
	s.StartTime = &v
	return s
}

func (s *DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail) SetState(v string) *DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail {
	s.State = &v
	return s
}

func (s *DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail) SetTable(v string) *DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail {
	s.Table = &v
	return s
}

func (s *DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail) Validate() error {
	return dara.Validate(s)
}
