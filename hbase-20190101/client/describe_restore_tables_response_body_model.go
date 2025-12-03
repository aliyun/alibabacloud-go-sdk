// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRestoreTablesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeRestoreTablesResponseBody
	GetRequestId() *string
	SetRestoreFull(v *DescribeRestoreTablesResponseBodyRestoreFull) *DescribeRestoreTablesResponseBody
	GetRestoreFull() *DescribeRestoreTablesResponseBodyRestoreFull
	SetRestoreIncrDetail(v *DescribeRestoreTablesResponseBodyRestoreIncrDetail) *DescribeRestoreTablesResponseBody
	GetRestoreIncrDetail() *DescribeRestoreTablesResponseBodyRestoreIncrDetail
	SetRestoreSchema(v *DescribeRestoreTablesResponseBodyRestoreSchema) *DescribeRestoreTablesResponseBody
	GetRestoreSchema() *DescribeRestoreTablesResponseBodyRestoreSchema
	SetRestoreSummary(v *DescribeRestoreTablesResponseBodyRestoreSummary) *DescribeRestoreTablesResponseBody
	GetRestoreSummary() *DescribeRestoreTablesResponseBodyRestoreSummary
	SetTables(v *DescribeRestoreTablesResponseBodyTables) *DescribeRestoreTablesResponseBody
	GetTables() *DescribeRestoreTablesResponseBodyTables
}

type DescribeRestoreTablesResponseBody struct {
	// example:
	//
	// 18D9CC47-D913-48BF-AB6B-4FA9B28FBDB1
	RequestId         *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RestoreFull       *DescribeRestoreTablesResponseBodyRestoreFull       `json:"RestoreFull,omitempty" xml:"RestoreFull,omitempty" type:"Struct"`
	RestoreIncrDetail *DescribeRestoreTablesResponseBodyRestoreIncrDetail `json:"RestoreIncrDetail,omitempty" xml:"RestoreIncrDetail,omitempty" type:"Struct"`
	RestoreSchema     *DescribeRestoreTablesResponseBodyRestoreSchema     `json:"RestoreSchema,omitempty" xml:"RestoreSchema,omitempty" type:"Struct"`
	RestoreSummary    *DescribeRestoreTablesResponseBodyRestoreSummary    `json:"RestoreSummary,omitempty" xml:"RestoreSummary,omitempty" type:"Struct"`
	Tables            *DescribeRestoreTablesResponseBodyTables            `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Struct"`
}

func (s DescribeRestoreTablesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRestoreTablesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRestoreTablesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRestoreTablesResponseBody) GetRestoreFull() *DescribeRestoreTablesResponseBodyRestoreFull {
	return s.RestoreFull
}

func (s *DescribeRestoreTablesResponseBody) GetRestoreIncrDetail() *DescribeRestoreTablesResponseBodyRestoreIncrDetail {
	return s.RestoreIncrDetail
}

func (s *DescribeRestoreTablesResponseBody) GetRestoreSchema() *DescribeRestoreTablesResponseBodyRestoreSchema {
	return s.RestoreSchema
}

func (s *DescribeRestoreTablesResponseBody) GetRestoreSummary() *DescribeRestoreTablesResponseBodyRestoreSummary {
	return s.RestoreSummary
}

func (s *DescribeRestoreTablesResponseBody) GetTables() *DescribeRestoreTablesResponseBodyTables {
	return s.Tables
}

func (s *DescribeRestoreTablesResponseBody) SetRequestId(v string) *DescribeRestoreTablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRestoreTablesResponseBody) SetRestoreFull(v *DescribeRestoreTablesResponseBodyRestoreFull) *DescribeRestoreTablesResponseBody {
	s.RestoreFull = v
	return s
}

func (s *DescribeRestoreTablesResponseBody) SetRestoreIncrDetail(v *DescribeRestoreTablesResponseBodyRestoreIncrDetail) *DescribeRestoreTablesResponseBody {
	s.RestoreIncrDetail = v
	return s
}

func (s *DescribeRestoreTablesResponseBody) SetRestoreSchema(v *DescribeRestoreTablesResponseBodyRestoreSchema) *DescribeRestoreTablesResponseBody {
	s.RestoreSchema = v
	return s
}

func (s *DescribeRestoreTablesResponseBody) SetRestoreSummary(v *DescribeRestoreTablesResponseBodyRestoreSummary) *DescribeRestoreTablesResponseBody {
	s.RestoreSummary = v
	return s
}

func (s *DescribeRestoreTablesResponseBody) SetTables(v *DescribeRestoreTablesResponseBodyTables) *DescribeRestoreTablesResponseBody {
	s.Tables = v
	return s
}

func (s *DescribeRestoreTablesResponseBody) Validate() error {
	if s.RestoreFull != nil {
		if err := s.RestoreFull.Validate(); err != nil {
			return err
		}
	}
	if s.RestoreIncrDetail != nil {
		if err := s.RestoreIncrDetail.Validate(); err != nil {
			return err
		}
	}
	if s.RestoreSchema != nil {
		if err := s.RestoreSchema.Validate(); err != nil {
			return err
		}
	}
	if s.RestoreSummary != nil {
		if err := s.RestoreSummary.Validate(); err != nil {
			return err
		}
	}
	if s.Tables != nil {
		if err := s.Tables.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRestoreTablesResponseBodyRestoreFull struct {
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
	PageSize           *int32                                                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RestoreFullDetails *DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetails `json:"RestoreFullDetails,omitempty" xml:"RestoreFullDetails,omitempty" type:"Struct"`
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

func (s DescribeRestoreTablesResponseBodyRestoreFull) String() string {
	return dara.Prettify(s)
}

func (s DescribeRestoreTablesResponseBodyRestoreFull) GoString() string {
	return s.String()
}

func (s *DescribeRestoreTablesResponseBodyRestoreFull) GetDataSize() *string {
	return s.DataSize
}

func (s *DescribeRestoreTablesResponseBodyRestoreFull) GetFail() *int32 {
	return s.Fail
}

func (s *DescribeRestoreTablesResponseBodyRestoreFull) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeRestoreTablesResponseBodyRestoreFull) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeRestoreTablesResponseBodyRestoreFull) GetRestoreFullDetails() *DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetails {
	return s.RestoreFullDetails
}

func (s *DescribeRestoreTablesResponseBodyRestoreFull) GetSpeed() *string {
	return s.Speed
}

func (s *DescribeRestoreTablesResponseBodyRestoreFull) GetSucceed() *int32 {
	return s.Succeed
}

func (s *DescribeRestoreTablesResponseBodyRestoreFull) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeRestoreTablesResponseBodyRestoreFull) SetDataSize(v string) *DescribeRestoreTablesResponseBodyRestoreFull {
	s.DataSize = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreFull) SetFail(v int32) *DescribeRestoreTablesResponseBodyRestoreFull {
	s.Fail = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreFull) SetPageNumber(v int32) *DescribeRestoreTablesResponseBodyRestoreFull {
	s.PageNumber = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreFull) SetPageSize(v int32) *DescribeRestoreTablesResponseBodyRestoreFull {
	s.PageSize = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreFull) SetRestoreFullDetails(v *DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetails) *DescribeRestoreTablesResponseBodyRestoreFull {
	s.RestoreFullDetails = v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreFull) SetSpeed(v string) *DescribeRestoreTablesResponseBodyRestoreFull {
	s.Speed = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreFull) SetSucceed(v int32) *DescribeRestoreTablesResponseBodyRestoreFull {
	s.Succeed = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreFull) SetTotal(v int64) *DescribeRestoreTablesResponseBodyRestoreFull {
	s.Total = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreFull) Validate() error {
	if s.RestoreFullDetails != nil {
		if err := s.RestoreFullDetails.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetails struct {
	RestoreFullDetail []*DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail `json:"RestoreFullDetail,omitempty" xml:"RestoreFullDetail,omitempty" type:"Repeated"`
}

func (s DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetails) String() string {
	return dara.Prettify(s)
}

func (s DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetails) GoString() string {
	return s.String()
}

func (s *DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetails) GetRestoreFullDetail() []*DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail {
	return s.RestoreFullDetail
}

func (s *DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetails) SetRestoreFullDetail(v []*DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail) *DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetails {
	s.RestoreFullDetail = v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetails) Validate() error {
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

type DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail struct {
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
	// “”
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

func (s DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail) GoString() string {
	return s.String()
}

func (s *DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail) GetDataSize() *string {
	return s.DataSize
}

func (s *DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail) GetMessage() *string {
	return s.Message
}

func (s *DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail) GetProcess() *string {
	return s.Process
}

func (s *DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail) GetSpeed() *string {
	return s.Speed
}

func (s *DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail) GetState() *string {
	return s.State
}

func (s *DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail) GetTable() *string {
	return s.Table
}

func (s *DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail) SetDataSize(v string) *DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail {
	s.DataSize = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail) SetEndTime(v string) *DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail {
	s.EndTime = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail) SetMessage(v string) *DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail {
	s.Message = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail) SetProcess(v string) *DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail {
	s.Process = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail) SetSpeed(v string) *DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail {
	s.Speed = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail) SetStartTime(v string) *DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail {
	s.StartTime = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail) SetState(v string) *DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail {
	s.State = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail) SetTable(v string) *DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail {
	s.Table = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail) Validate() error {
	return dara.Validate(s)
}

type DescribeRestoreTablesResponseBodyRestoreIncrDetail struct {
	// example:
	//
	// 2020-11-05T06:45:44Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 0/0
	Process *string `json:"Process,omitempty" xml:"Process,omitempty"`
	// example:
	//
	// 0 ms
	RestoreDelay *string `json:"RestoreDelay,omitempty" xml:"RestoreDelay,omitempty"`
	// example:
	//
	// 2020-11-02T18:00:00Z
	RestoreStartTs *string `json:"RestoreStartTs,omitempty" xml:"RestoreStartTs,omitempty"`
	// example:
	//
	// “”
	RestoredTs *string `json:"RestoredTs,omitempty" xml:"RestoredTs,omitempty"`
	// example:
	//
	// 2020-11-05T06:45:44Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// SUCCEEDED
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s DescribeRestoreTablesResponseBodyRestoreIncrDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeRestoreTablesResponseBodyRestoreIncrDetail) GoString() string {
	return s.String()
}

func (s *DescribeRestoreTablesResponseBodyRestoreIncrDetail) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeRestoreTablesResponseBodyRestoreIncrDetail) GetProcess() *string {
	return s.Process
}

func (s *DescribeRestoreTablesResponseBodyRestoreIncrDetail) GetRestoreDelay() *string {
	return s.RestoreDelay
}

func (s *DescribeRestoreTablesResponseBodyRestoreIncrDetail) GetRestoreStartTs() *string {
	return s.RestoreStartTs
}

func (s *DescribeRestoreTablesResponseBodyRestoreIncrDetail) GetRestoredTs() *string {
	return s.RestoredTs
}

func (s *DescribeRestoreTablesResponseBodyRestoreIncrDetail) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeRestoreTablesResponseBodyRestoreIncrDetail) GetState() *string {
	return s.State
}

func (s *DescribeRestoreTablesResponseBodyRestoreIncrDetail) SetEndTime(v string) *DescribeRestoreTablesResponseBodyRestoreIncrDetail {
	s.EndTime = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreIncrDetail) SetProcess(v string) *DescribeRestoreTablesResponseBodyRestoreIncrDetail {
	s.Process = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreIncrDetail) SetRestoreDelay(v string) *DescribeRestoreTablesResponseBodyRestoreIncrDetail {
	s.RestoreDelay = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreIncrDetail) SetRestoreStartTs(v string) *DescribeRestoreTablesResponseBodyRestoreIncrDetail {
	s.RestoreStartTs = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreIncrDetail) SetRestoredTs(v string) *DescribeRestoreTablesResponseBodyRestoreIncrDetail {
	s.RestoredTs = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreIncrDetail) SetStartTime(v string) *DescribeRestoreTablesResponseBodyRestoreIncrDetail {
	s.StartTime = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreIncrDetail) SetState(v string) *DescribeRestoreTablesResponseBodyRestoreIncrDetail {
	s.State = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreIncrDetail) Validate() error {
	return dara.Validate(s)
}

type DescribeRestoreTablesResponseBodyRestoreSchema struct {
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
	PageSize             *int32                                                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RestoreSchemaDetails *DescribeRestoreTablesResponseBodyRestoreSchemaRestoreSchemaDetails `json:"RestoreSchemaDetails,omitempty" xml:"RestoreSchemaDetails,omitempty" type:"Struct"`
	// example:
	//
	// 1
	Succeed *int32 `json:"Succeed,omitempty" xml:"Succeed,omitempty"`
	// example:
	//
	// 1
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeRestoreTablesResponseBodyRestoreSchema) String() string {
	return dara.Prettify(s)
}

func (s DescribeRestoreTablesResponseBodyRestoreSchema) GoString() string {
	return s.String()
}

func (s *DescribeRestoreTablesResponseBodyRestoreSchema) GetFail() *int32 {
	return s.Fail
}

func (s *DescribeRestoreTablesResponseBodyRestoreSchema) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeRestoreTablesResponseBodyRestoreSchema) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeRestoreTablesResponseBodyRestoreSchema) GetRestoreSchemaDetails() *DescribeRestoreTablesResponseBodyRestoreSchemaRestoreSchemaDetails {
	return s.RestoreSchemaDetails
}

func (s *DescribeRestoreTablesResponseBodyRestoreSchema) GetSucceed() *int32 {
	return s.Succeed
}

func (s *DescribeRestoreTablesResponseBodyRestoreSchema) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeRestoreTablesResponseBodyRestoreSchema) SetFail(v int32) *DescribeRestoreTablesResponseBodyRestoreSchema {
	s.Fail = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreSchema) SetPageNumber(v int32) *DescribeRestoreTablesResponseBodyRestoreSchema {
	s.PageNumber = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreSchema) SetPageSize(v int32) *DescribeRestoreTablesResponseBodyRestoreSchema {
	s.PageSize = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreSchema) SetRestoreSchemaDetails(v *DescribeRestoreTablesResponseBodyRestoreSchemaRestoreSchemaDetails) *DescribeRestoreTablesResponseBodyRestoreSchema {
	s.RestoreSchemaDetails = v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreSchema) SetSucceed(v int32) *DescribeRestoreTablesResponseBodyRestoreSchema {
	s.Succeed = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreSchema) SetTotal(v int64) *DescribeRestoreTablesResponseBodyRestoreSchema {
	s.Total = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreSchema) Validate() error {
	if s.RestoreSchemaDetails != nil {
		if err := s.RestoreSchemaDetails.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRestoreTablesResponseBodyRestoreSchemaRestoreSchemaDetails struct {
	RestoreSchemaDetail []*DescribeRestoreTablesResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail `json:"RestoreSchemaDetail,omitempty" xml:"RestoreSchemaDetail,omitempty" type:"Repeated"`
}

func (s DescribeRestoreTablesResponseBodyRestoreSchemaRestoreSchemaDetails) String() string {
	return dara.Prettify(s)
}

func (s DescribeRestoreTablesResponseBodyRestoreSchemaRestoreSchemaDetails) GoString() string {
	return s.String()
}

func (s *DescribeRestoreTablesResponseBodyRestoreSchemaRestoreSchemaDetails) GetRestoreSchemaDetail() []*DescribeRestoreTablesResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail {
	return s.RestoreSchemaDetail
}

func (s *DescribeRestoreTablesResponseBodyRestoreSchemaRestoreSchemaDetails) SetRestoreSchemaDetail(v []*DescribeRestoreTablesResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail) *DescribeRestoreTablesResponseBodyRestoreSchemaRestoreSchemaDetails {
	s.RestoreSchemaDetail = v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreSchemaRestoreSchemaDetails) Validate() error {
	if s.RestoreSchemaDetail != nil {
		for _, item := range s.RestoreSchemaDetail {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRestoreTablesResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail struct {
	// example:
	//
	// 2020-11-05T06:45:18Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// null
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 2020-11-05T06:45:14Z
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

func (s DescribeRestoreTablesResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeRestoreTablesResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail) GoString() string {
	return s.String()
}

func (s *DescribeRestoreTablesResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeRestoreTablesResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail) GetMessage() *string {
	return s.Message
}

func (s *DescribeRestoreTablesResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeRestoreTablesResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail) GetState() *string {
	return s.State
}

func (s *DescribeRestoreTablesResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail) GetTable() *string {
	return s.Table
}

func (s *DescribeRestoreTablesResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail) SetEndTime(v string) *DescribeRestoreTablesResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail {
	s.EndTime = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail) SetMessage(v string) *DescribeRestoreTablesResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail {
	s.Message = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail) SetStartTime(v string) *DescribeRestoreTablesResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail {
	s.StartTime = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail) SetState(v string) *DescribeRestoreTablesResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail {
	s.State = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail) SetTable(v string) *DescribeRestoreTablesResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail {
	s.Table = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail) Validate() error {
	return dara.Validate(s)
}

type DescribeRestoreTablesResponseBodyRestoreSummary struct {
	// example:
	//
	// 2020-11-05T06:45:51Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 20201105144514
	RecordId *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// example:
	//
	// 2020-11-04T05:15:00Z
	RestoreToDate *string `json:"RestoreToDate,omitempty" xml:"RestoreToDate,omitempty"`
	// example:
	//
	// 2020-11-05T06:45:14Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// SUCCEEDED
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// example:
	//
	// ld-m5e2t34kr54wgxxxx
	TargetCluster *string `json:"TargetCluster,omitempty" xml:"TargetCluster,omitempty"`
}

func (s DescribeRestoreTablesResponseBodyRestoreSummary) String() string {
	return dara.Prettify(s)
}

func (s DescribeRestoreTablesResponseBodyRestoreSummary) GoString() string {
	return s.String()
}

func (s *DescribeRestoreTablesResponseBodyRestoreSummary) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeRestoreTablesResponseBodyRestoreSummary) GetRecordId() *string {
	return s.RecordId
}

func (s *DescribeRestoreTablesResponseBodyRestoreSummary) GetRestoreToDate() *string {
	return s.RestoreToDate
}

func (s *DescribeRestoreTablesResponseBodyRestoreSummary) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeRestoreTablesResponseBodyRestoreSummary) GetState() *string {
	return s.State
}

func (s *DescribeRestoreTablesResponseBodyRestoreSummary) GetTargetCluster() *string {
	return s.TargetCluster
}

func (s *DescribeRestoreTablesResponseBodyRestoreSummary) SetEndTime(v string) *DescribeRestoreTablesResponseBodyRestoreSummary {
	s.EndTime = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreSummary) SetRecordId(v string) *DescribeRestoreTablesResponseBodyRestoreSummary {
	s.RecordId = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreSummary) SetRestoreToDate(v string) *DescribeRestoreTablesResponseBodyRestoreSummary {
	s.RestoreToDate = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreSummary) SetStartTime(v string) *DescribeRestoreTablesResponseBodyRestoreSummary {
	s.StartTime = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreSummary) SetState(v string) *DescribeRestoreTablesResponseBodyRestoreSummary {
	s.State = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreSummary) SetTargetCluster(v string) *DescribeRestoreTablesResponseBodyRestoreSummary {
	s.TargetCluster = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreSummary) Validate() error {
	return dara.Validate(s)
}

type DescribeRestoreTablesResponseBodyTables struct {
	Table []*string `json:"Table,omitempty" xml:"Table,omitempty" type:"Repeated"`
}

func (s DescribeRestoreTablesResponseBodyTables) String() string {
	return dara.Prettify(s)
}

func (s DescribeRestoreTablesResponseBodyTables) GoString() string {
	return s.String()
}

func (s *DescribeRestoreTablesResponseBodyTables) GetTable() []*string {
	return s.Table
}

func (s *DescribeRestoreTablesResponseBodyTables) SetTable(v []*string) *DescribeRestoreTablesResponseBodyTables {
	s.Table = v
	return s
}

func (s *DescribeRestoreTablesResponseBodyTables) Validate() error {
	return dara.Validate(s)
}
