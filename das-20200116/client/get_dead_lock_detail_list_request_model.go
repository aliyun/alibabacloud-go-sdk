// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeadLockDetailListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbNameList(v string) *GetDeadLockDetailListRequest
	GetDbNameList() *string
	SetEndTime(v string) *GetDeadLockDetailListRequest
	GetEndTime() *string
	SetInstanceId(v string) *GetDeadLockDetailListRequest
	GetInstanceId() *string
	SetPageNo(v string) *GetDeadLockDetailListRequest
	GetPageNo() *string
	SetPageSize(v string) *GetDeadLockDetailListRequest
	GetPageSize() *string
	SetStartTime(v string) *GetDeadLockDetailListRequest
	GetStartTime() *string
}

type GetDeadLockDetailListRequest struct {
	// The name of the database. When you specify multiple databases, you must separate the database names with commas (,).
	//
	// example:
	//
	// school1,school2
	DbNameList *string `json:"DbNameList,omitempty" xml:"DbNameList,omitempty"`
	// The end of the time range to query. Set this parameter to a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1702360530292
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the database instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-2ze2016723b328gs2
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number. The value must be an integer that is greater than 0. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *string `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 5
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The beginning of the time range to query. Set this parameter to a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1701755730292
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetDeadLockDetailListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDeadLockDetailListRequest) GoString() string {
	return s.String()
}

func (s *GetDeadLockDetailListRequest) GetDbNameList() *string {
	return s.DbNameList
}

func (s *GetDeadLockDetailListRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *GetDeadLockDetailListRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetDeadLockDetailListRequest) GetPageNo() *string {
	return s.PageNo
}

func (s *GetDeadLockDetailListRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *GetDeadLockDetailListRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *GetDeadLockDetailListRequest) SetDbNameList(v string) *GetDeadLockDetailListRequest {
	s.DbNameList = &v
	return s
}

func (s *GetDeadLockDetailListRequest) SetEndTime(v string) *GetDeadLockDetailListRequest {
	s.EndTime = &v
	return s
}

func (s *GetDeadLockDetailListRequest) SetInstanceId(v string) *GetDeadLockDetailListRequest {
	s.InstanceId = &v
	return s
}

func (s *GetDeadLockDetailListRequest) SetPageNo(v string) *GetDeadLockDetailListRequest {
	s.PageNo = &v
	return s
}

func (s *GetDeadLockDetailListRequest) SetPageSize(v string) *GetDeadLockDetailListRequest {
	s.PageSize = &v
	return s
}

func (s *GetDeadLockDetailListRequest) SetStartTime(v string) *GetDeadLockDetailListRequest {
	s.StartTime = &v
	return s
}

func (s *GetDeadLockDetailListRequest) Validate() error {
	return dara.Validate(s)
}
