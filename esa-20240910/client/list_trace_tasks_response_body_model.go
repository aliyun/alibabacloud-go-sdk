// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTraceTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int64) *ListTraceTasksResponseBody
	GetCount() *int64
	SetList(v []*ListTraceTasksResponseBodyList) *ListTraceTasksResponseBody
	GetList() []*ListTraceTasksResponseBodyList
	SetPageNumber(v int64) *ListTraceTasksResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *ListTraceTasksResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListTraceTasksResponseBody
	GetRequestId() *string
}

type ListTraceTasksResponseBody struct {
	// example:
	//
	// 10
	Count *int64                            `json:"Count,omitempty" xml:"Count,omitempty"`
	List  []*ListTraceTasksResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 6
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 4E09C5D7-E1CF-4CAA-A45E-8727F4C8FD70
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListTraceTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTraceTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListTraceTasksResponseBody) GetCount() *int64 {
	return s.Count
}

func (s *ListTraceTasksResponseBody) GetList() []*ListTraceTasksResponseBodyList {
	return s.List
}

func (s *ListTraceTasksResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListTraceTasksResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListTraceTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTraceTasksResponseBody) SetCount(v int64) *ListTraceTasksResponseBody {
	s.Count = &v
	return s
}

func (s *ListTraceTasksResponseBody) SetList(v []*ListTraceTasksResponseBodyList) *ListTraceTasksResponseBody {
	s.List = v
	return s
}

func (s *ListTraceTasksResponseBody) SetPageNumber(v int64) *ListTraceTasksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListTraceTasksResponseBody) SetPageSize(v int64) *ListTraceTasksResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListTraceTasksResponseBody) SetRequestId(v string) *ListTraceTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTraceTasksResponseBody) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTraceTasksResponseBodyList struct {
	// example:
	//
	// 1077********7468
	Aliuid *string `json:"Aliuid,omitempty" xml:"Aliuid,omitempty"`
	// example:
	//
	// xx.xx.xx.xx
	ClientAddr *string `json:"ClientAddr,omitempty" xml:"ClientAddr,omitempty"`
	// example:
	//
	// xx.xx.xx.xx
	ClientIp *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	// example:
	//
	// 2022-12-10 15:11:47
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// f2a18ad5
	DiagnoseId *string `json:"DiagnoseId,omitempty" xml:"DiagnoseId,omitempty"`
	// example:
	//
	// http://cdn.dns-detect.alicdn.com/diagnose/xxxxxx
	DiagnoseUrl *string `json:"DiagnoseUrl,omitempty" xml:"DiagnoseUrl,omitempty"`
	// example:
	//
	// http://www.example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// 1669285111
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// 10
	RemainDiagnoseTimes *int64 `json:"RemainDiagnoseTimes,omitempty" xml:"RemainDiagnoseTimes,omitempty"`
	// example:
	//
	// 0
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// example:
	//
	// 1
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 000000xxxxxxxxxxxxxxxxxxxxxxxxxxxx475e
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// 1
	TimeConsuming *int64 `json:"TimeConsuming,omitempty" xml:"TimeConsuming,omitempty"`
	// example:
	//
	// 000000xxxxxxxxxxxxxxxxxxxxxx25941e
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s ListTraceTasksResponseBodyList) String() string {
	return dara.Prettify(s)
}

func (s ListTraceTasksResponseBodyList) GoString() string {
	return s.String()
}

func (s *ListTraceTasksResponseBodyList) GetAliuid() *string {
	return s.Aliuid
}

func (s *ListTraceTasksResponseBodyList) GetClientAddr() *string {
	return s.ClientAddr
}

func (s *ListTraceTasksResponseBodyList) GetClientIp() *string {
	return s.ClientIp
}

func (s *ListTraceTasksResponseBodyList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListTraceTasksResponseBodyList) GetDiagnoseId() *string {
	return s.DiagnoseId
}

func (s *ListTraceTasksResponseBodyList) GetDiagnoseUrl() *string {
	return s.DiagnoseUrl
}

func (s *ListTraceTasksResponseBodyList) GetDomain() *string {
	return s.Domain
}

func (s *ListTraceTasksResponseBodyList) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *ListTraceTasksResponseBodyList) GetRemainDiagnoseTimes() *int64 {
	return s.RemainDiagnoseTimes
}

func (s *ListTraceTasksResponseBodyList) GetState() *string {
	return s.State
}

func (s *ListTraceTasksResponseBodyList) GetStatus() *int64 {
	return s.Status
}

func (s *ListTraceTasksResponseBodyList) GetTaskId() *string {
	return s.TaskId
}

func (s *ListTraceTasksResponseBodyList) GetTimeConsuming() *int64 {
	return s.TimeConsuming
}

func (s *ListTraceTasksResponseBodyList) GetTraceId() *string {
	return s.TraceId
}

func (s *ListTraceTasksResponseBodyList) SetAliuid(v string) *ListTraceTasksResponseBodyList {
	s.Aliuid = &v
	return s
}

func (s *ListTraceTasksResponseBodyList) SetClientAddr(v string) *ListTraceTasksResponseBodyList {
	s.ClientAddr = &v
	return s
}

func (s *ListTraceTasksResponseBodyList) SetClientIp(v string) *ListTraceTasksResponseBodyList {
	s.ClientIp = &v
	return s
}

func (s *ListTraceTasksResponseBodyList) SetCreateTime(v string) *ListTraceTasksResponseBodyList {
	s.CreateTime = &v
	return s
}

func (s *ListTraceTasksResponseBodyList) SetDiagnoseId(v string) *ListTraceTasksResponseBodyList {
	s.DiagnoseId = &v
	return s
}

func (s *ListTraceTasksResponseBodyList) SetDiagnoseUrl(v string) *ListTraceTasksResponseBodyList {
	s.DiagnoseUrl = &v
	return s
}

func (s *ListTraceTasksResponseBodyList) SetDomain(v string) *ListTraceTasksResponseBodyList {
	s.Domain = &v
	return s
}

func (s *ListTraceTasksResponseBodyList) SetExpireTime(v int64) *ListTraceTasksResponseBodyList {
	s.ExpireTime = &v
	return s
}

func (s *ListTraceTasksResponseBodyList) SetRemainDiagnoseTimes(v int64) *ListTraceTasksResponseBodyList {
	s.RemainDiagnoseTimes = &v
	return s
}

func (s *ListTraceTasksResponseBodyList) SetState(v string) *ListTraceTasksResponseBodyList {
	s.State = &v
	return s
}

func (s *ListTraceTasksResponseBodyList) SetStatus(v int64) *ListTraceTasksResponseBodyList {
	s.Status = &v
	return s
}

func (s *ListTraceTasksResponseBodyList) SetTaskId(v string) *ListTraceTasksResponseBodyList {
	s.TaskId = &v
	return s
}

func (s *ListTraceTasksResponseBodyList) SetTimeConsuming(v int64) *ListTraceTasksResponseBodyList {
	s.TimeConsuming = &v
	return s
}

func (s *ListTraceTasksResponseBodyList) SetTraceId(v string) *ListTraceTasksResponseBodyList {
	s.TraceId = &v
	return s
}

func (s *ListTraceTasksResponseBodyList) Validate() error {
	return dara.Validate(s)
}
