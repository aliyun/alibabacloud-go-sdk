// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceMissingIndexListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAvgTotalUserCost(v string) *GetInstanceMissingIndexListRequest
	GetAvgTotalUserCost() *string
	SetAvgUserImpact(v string) *GetInstanceMissingIndexListRequest
	GetAvgUserImpact() *string
	SetEndTime(v string) *GetInstanceMissingIndexListRequest
	GetEndTime() *string
	SetIndexCount(v string) *GetInstanceMissingIndexListRequest
	GetIndexCount() *string
	SetInstanceId(v string) *GetInstanceMissingIndexListRequest
	GetInstanceId() *string
	SetObjectName(v string) *GetInstanceMissingIndexListRequest
	GetObjectName() *string
	SetPageNo(v string) *GetInstanceMissingIndexListRequest
	GetPageNo() *string
	SetPageSize(v string) *GetInstanceMissingIndexListRequest
	GetPageSize() *string
	SetReservedPages(v string) *GetInstanceMissingIndexListRequest
	GetReservedPages() *string
	SetReservedSize(v string) *GetInstanceMissingIndexListRequest
	GetReservedSize() *string
	SetRowCount(v string) *GetInstanceMissingIndexListRequest
	GetRowCount() *string
	SetStartTime(v string) *GetInstanceMissingIndexListRequest
	GetStartTime() *string
	SetUniqueCompiles(v string) *GetInstanceMissingIndexListRequest
	GetUniqueCompiles() *string
	SetUserScans(v string) *GetInstanceMissingIndexListRequest
	GetUserScans() *string
	SetUserSeeks(v string) *GetInstanceMissingIndexListRequest
	GetUserSeeks() *string
}

type GetInstanceMissingIndexListRequest struct {
	// The query condition based on the average cost savings.
	//
	// example:
	//
	// <=|8
	AvgTotalUserCost *string `json:"AvgTotalUserCost,omitempty" xml:"AvgTotalUserCost,omitempty"`
	// The query condition based on the performance improvement.
	//
	// example:
	//
	// >|10000
	AvgUserImpact *string `json:"AvgUserImpact,omitempty" xml:"AvgUserImpact,omitempty"`
	// The end time of the last seek.
	//
	// example:
	//
	// 1681869544000
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The query condition based on the number of indexes.
	//
	// example:
	//
	// >=|8
	IndexCount *string `json:"IndexCount,omitempty" xml:"IndexCount,omitempty"`
	// The database instance ID.
	//
	// >  Only ApsaraDB RDS for SQL Server instances are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-************
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The object name.
	//
	// example:
	//
	// bas_customer
	ObjectName *string `json:"ObjectName,omitempty" xml:"ObjectName,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *string `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The query condition based on the total number of pages.
	//
	// example:
	//
	// >=|100
	ReservedPages *string `json:"ReservedPages,omitempty" xml:"ReservedPages,omitempty"`
	// The query condition based on the table size.
	//
	// example:
	//
	// >=|100
	ReservedSize *string `json:"ReservedSize,omitempty" xml:"ReservedSize,omitempty"`
	// The query condition based on the number of table rows.
	//
	// example:
	//
	// >=|100000
	RowCount *string `json:"RowCount,omitempty" xml:"RowCount,omitempty"`
	// The start time of the last seek.
	//
	// example:
	//
	// 1679414400000
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The query condition based on the number of compilations.
	//
	// example:
	//
	// >=|10000
	UniqueCompiles *string `json:"UniqueCompiles,omitempty" xml:"UniqueCompiles,omitempty"`
	// The query condition based on the number of scans.
	//
	// example:
	//
	// >=|10000
	UserScans *string `json:"UserScans,omitempty" xml:"UserScans,omitempty"`
	// The query condition based on the number of seeks.
	//
	// example:
	//
	// >=|1000
	UserSeeks *string `json:"UserSeeks,omitempty" xml:"UserSeeks,omitempty"`
}

func (s GetInstanceMissingIndexListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceMissingIndexListRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceMissingIndexListRequest) GetAvgTotalUserCost() *string {
	return s.AvgTotalUserCost
}

func (s *GetInstanceMissingIndexListRequest) GetAvgUserImpact() *string {
	return s.AvgUserImpact
}

func (s *GetInstanceMissingIndexListRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *GetInstanceMissingIndexListRequest) GetIndexCount() *string {
	return s.IndexCount
}

func (s *GetInstanceMissingIndexListRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetInstanceMissingIndexListRequest) GetObjectName() *string {
	return s.ObjectName
}

func (s *GetInstanceMissingIndexListRequest) GetPageNo() *string {
	return s.PageNo
}

func (s *GetInstanceMissingIndexListRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *GetInstanceMissingIndexListRequest) GetReservedPages() *string {
	return s.ReservedPages
}

func (s *GetInstanceMissingIndexListRequest) GetReservedSize() *string {
	return s.ReservedSize
}

func (s *GetInstanceMissingIndexListRequest) GetRowCount() *string {
	return s.RowCount
}

func (s *GetInstanceMissingIndexListRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *GetInstanceMissingIndexListRequest) GetUniqueCompiles() *string {
	return s.UniqueCompiles
}

func (s *GetInstanceMissingIndexListRequest) GetUserScans() *string {
	return s.UserScans
}

func (s *GetInstanceMissingIndexListRequest) GetUserSeeks() *string {
	return s.UserSeeks
}

func (s *GetInstanceMissingIndexListRequest) SetAvgTotalUserCost(v string) *GetInstanceMissingIndexListRequest {
	s.AvgTotalUserCost = &v
	return s
}

func (s *GetInstanceMissingIndexListRequest) SetAvgUserImpact(v string) *GetInstanceMissingIndexListRequest {
	s.AvgUserImpact = &v
	return s
}

func (s *GetInstanceMissingIndexListRequest) SetEndTime(v string) *GetInstanceMissingIndexListRequest {
	s.EndTime = &v
	return s
}

func (s *GetInstanceMissingIndexListRequest) SetIndexCount(v string) *GetInstanceMissingIndexListRequest {
	s.IndexCount = &v
	return s
}

func (s *GetInstanceMissingIndexListRequest) SetInstanceId(v string) *GetInstanceMissingIndexListRequest {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceMissingIndexListRequest) SetObjectName(v string) *GetInstanceMissingIndexListRequest {
	s.ObjectName = &v
	return s
}

func (s *GetInstanceMissingIndexListRequest) SetPageNo(v string) *GetInstanceMissingIndexListRequest {
	s.PageNo = &v
	return s
}

func (s *GetInstanceMissingIndexListRequest) SetPageSize(v string) *GetInstanceMissingIndexListRequest {
	s.PageSize = &v
	return s
}

func (s *GetInstanceMissingIndexListRequest) SetReservedPages(v string) *GetInstanceMissingIndexListRequest {
	s.ReservedPages = &v
	return s
}

func (s *GetInstanceMissingIndexListRequest) SetReservedSize(v string) *GetInstanceMissingIndexListRequest {
	s.ReservedSize = &v
	return s
}

func (s *GetInstanceMissingIndexListRequest) SetRowCount(v string) *GetInstanceMissingIndexListRequest {
	s.RowCount = &v
	return s
}

func (s *GetInstanceMissingIndexListRequest) SetStartTime(v string) *GetInstanceMissingIndexListRequest {
	s.StartTime = &v
	return s
}

func (s *GetInstanceMissingIndexListRequest) SetUniqueCompiles(v string) *GetInstanceMissingIndexListRequest {
	s.UniqueCompiles = &v
	return s
}

func (s *GetInstanceMissingIndexListRequest) SetUserScans(v string) *GetInstanceMissingIndexListRequest {
	s.UserScans = &v
	return s
}

func (s *GetInstanceMissingIndexListRequest) SetUserSeeks(v string) *GetInstanceMissingIndexListRequest {
	s.UserSeeks = &v
	return s
}

func (s *GetInstanceMissingIndexListRequest) Validate() error {
	return dara.Validate(s)
}
