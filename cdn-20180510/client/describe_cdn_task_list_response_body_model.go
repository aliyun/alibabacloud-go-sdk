// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnTaskListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v *DescribeCdnTaskListResponseBodyContent) *DescribeCdnTaskListResponseBody
	GetContent() *DescribeCdnTaskListResponseBodyContent
	SetRequestId(v string) *DescribeCdnTaskListResponseBody
	GetRequestId() *string
}

type DescribeCdnTaskListResponseBody struct {
	Content *DescribeCdnTaskListResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// example:
	//
	// 4E09C5D7-E1CF-4CAA-A45E-8727F4C8FD70
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCdnTaskListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnTaskListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCdnTaskListResponseBody) GetContent() *DescribeCdnTaskListResponseBodyContent {
	return s.Content
}

func (s *DescribeCdnTaskListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCdnTaskListResponseBody) SetContent(v *DescribeCdnTaskListResponseBodyContent) *DescribeCdnTaskListResponseBody {
	s.Content = v
	return s
}

func (s *DescribeCdnTaskListResponseBody) SetRequestId(v string) *DescribeCdnTaskListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCdnTaskListResponseBody) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCdnTaskListResponseBodyContent struct {
	// example:
	//
	// 10
	Count *int64                                        `json:"Count,omitempty" xml:"Count,omitempty"`
	List  []*DescribeCdnTaskListResponseBodyContentList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 6
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeCdnTaskListResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnTaskListResponseBodyContent) GoString() string {
	return s.String()
}

func (s *DescribeCdnTaskListResponseBodyContent) GetCount() *int64 {
	return s.Count
}

func (s *DescribeCdnTaskListResponseBodyContent) GetList() []*DescribeCdnTaskListResponseBodyContentList {
	return s.List
}

func (s *DescribeCdnTaskListResponseBodyContent) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeCdnTaskListResponseBodyContent) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeCdnTaskListResponseBodyContent) SetCount(v int64) *DescribeCdnTaskListResponseBodyContent {
	s.Count = &v
	return s
}

func (s *DescribeCdnTaskListResponseBodyContent) SetList(v []*DescribeCdnTaskListResponseBodyContentList) *DescribeCdnTaskListResponseBodyContent {
	s.List = v
	return s
}

func (s *DescribeCdnTaskListResponseBodyContent) SetPageNumber(v int64) *DescribeCdnTaskListResponseBodyContent {
	s.PageNumber = &v
	return s
}

func (s *DescribeCdnTaskListResponseBodyContent) SetPageSize(v int64) *DescribeCdnTaskListResponseBodyContent {
	s.PageSize = &v
	return s
}

func (s *DescribeCdnTaskListResponseBodyContent) Validate() error {
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

type DescribeCdnTaskListResponseBodyContentList struct {
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
	// 1s
	TimeConsuming *int64 `json:"TimeConsuming,omitempty" xml:"TimeConsuming,omitempty"`
	// example:
	//
	// 000000xxxxxxxxxxxxxxxxxxxxxx25941e
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DescribeCdnTaskListResponseBodyContentList) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnTaskListResponseBodyContentList) GoString() string {
	return s.String()
}

func (s *DescribeCdnTaskListResponseBodyContentList) GetAliuid() *string {
	return s.Aliuid
}

func (s *DescribeCdnTaskListResponseBodyContentList) GetClientAddr() *string {
	return s.ClientAddr
}

func (s *DescribeCdnTaskListResponseBodyContentList) GetClientIp() *string {
	return s.ClientIp
}

func (s *DescribeCdnTaskListResponseBodyContentList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeCdnTaskListResponseBodyContentList) GetDiagnoseId() *string {
	return s.DiagnoseId
}

func (s *DescribeCdnTaskListResponseBodyContentList) GetDiagnoseUrl() *string {
	return s.DiagnoseUrl
}

func (s *DescribeCdnTaskListResponseBodyContentList) GetDomain() *string {
	return s.Domain
}

func (s *DescribeCdnTaskListResponseBodyContentList) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *DescribeCdnTaskListResponseBodyContentList) GetRemainDiagnoseTimes() *int64 {
	return s.RemainDiagnoseTimes
}

func (s *DescribeCdnTaskListResponseBodyContentList) GetState() *string {
	return s.State
}

func (s *DescribeCdnTaskListResponseBodyContentList) GetStatus() *int64 {
	return s.Status
}

func (s *DescribeCdnTaskListResponseBodyContentList) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeCdnTaskListResponseBodyContentList) GetTimeConsuming() *int64 {
	return s.TimeConsuming
}

func (s *DescribeCdnTaskListResponseBodyContentList) GetTraceId() *string {
	return s.TraceId
}

func (s *DescribeCdnTaskListResponseBodyContentList) SetAliuid(v string) *DescribeCdnTaskListResponseBodyContentList {
	s.Aliuid = &v
	return s
}

func (s *DescribeCdnTaskListResponseBodyContentList) SetClientAddr(v string) *DescribeCdnTaskListResponseBodyContentList {
	s.ClientAddr = &v
	return s
}

func (s *DescribeCdnTaskListResponseBodyContentList) SetClientIp(v string) *DescribeCdnTaskListResponseBodyContentList {
	s.ClientIp = &v
	return s
}

func (s *DescribeCdnTaskListResponseBodyContentList) SetCreateTime(v string) *DescribeCdnTaskListResponseBodyContentList {
	s.CreateTime = &v
	return s
}

func (s *DescribeCdnTaskListResponseBodyContentList) SetDiagnoseId(v string) *DescribeCdnTaskListResponseBodyContentList {
	s.DiagnoseId = &v
	return s
}

func (s *DescribeCdnTaskListResponseBodyContentList) SetDiagnoseUrl(v string) *DescribeCdnTaskListResponseBodyContentList {
	s.DiagnoseUrl = &v
	return s
}

func (s *DescribeCdnTaskListResponseBodyContentList) SetDomain(v string) *DescribeCdnTaskListResponseBodyContentList {
	s.Domain = &v
	return s
}

func (s *DescribeCdnTaskListResponseBodyContentList) SetExpireTime(v int64) *DescribeCdnTaskListResponseBodyContentList {
	s.ExpireTime = &v
	return s
}

func (s *DescribeCdnTaskListResponseBodyContentList) SetRemainDiagnoseTimes(v int64) *DescribeCdnTaskListResponseBodyContentList {
	s.RemainDiagnoseTimes = &v
	return s
}

func (s *DescribeCdnTaskListResponseBodyContentList) SetState(v string) *DescribeCdnTaskListResponseBodyContentList {
	s.State = &v
	return s
}

func (s *DescribeCdnTaskListResponseBodyContentList) SetStatus(v int64) *DescribeCdnTaskListResponseBodyContentList {
	s.Status = &v
	return s
}

func (s *DescribeCdnTaskListResponseBodyContentList) SetTaskId(v string) *DescribeCdnTaskListResponseBodyContentList {
	s.TaskId = &v
	return s
}

func (s *DescribeCdnTaskListResponseBodyContentList) SetTimeConsuming(v int64) *DescribeCdnTaskListResponseBodyContentList {
	s.TimeConsuming = &v
	return s
}

func (s *DescribeCdnTaskListResponseBodyContentList) SetTraceId(v string) *DescribeCdnTaskListResponseBodyContentList {
	s.TraceId = &v
	return s
}

func (s *DescribeCdnTaskListResponseBodyContentList) Validate() error {
	return dara.Validate(s)
}
