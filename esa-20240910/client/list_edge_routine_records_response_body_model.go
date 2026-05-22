// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEdgeRoutineRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListEdgeRoutineRecordsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListEdgeRoutineRecordsResponseBody
	GetPageSize() *int32
	SetRecords(v []*ListEdgeRoutineRecordsResponseBodyRecords) *ListEdgeRoutineRecordsResponseBody
	GetRecords() []*ListEdgeRoutineRecordsResponseBodyRecords
	SetRequestId(v string) *ListEdgeRoutineRecordsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListEdgeRoutineRecordsResponseBody
	GetTotalCount() *int32
}

type ListEdgeRoutineRecordsResponseBody struct {
	PageNumber *int32                                       `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Records    []*ListEdgeRoutineRecordsResponseBodyRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	RequestId  *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                       `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListEdgeRoutineRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEdgeRoutineRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *ListEdgeRoutineRecordsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListEdgeRoutineRecordsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListEdgeRoutineRecordsResponseBody) GetRecords() []*ListEdgeRoutineRecordsResponseBodyRecords {
	return s.Records
}

func (s *ListEdgeRoutineRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEdgeRoutineRecordsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListEdgeRoutineRecordsResponseBody) SetPageNumber(v int32) *ListEdgeRoutineRecordsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListEdgeRoutineRecordsResponseBody) SetPageSize(v int32) *ListEdgeRoutineRecordsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListEdgeRoutineRecordsResponseBody) SetRecords(v []*ListEdgeRoutineRecordsResponseBodyRecords) *ListEdgeRoutineRecordsResponseBody {
	s.Records = v
	return s
}

func (s *ListEdgeRoutineRecordsResponseBody) SetRequestId(v string) *ListEdgeRoutineRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEdgeRoutineRecordsResponseBody) SetTotalCount(v int32) *ListEdgeRoutineRecordsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListEdgeRoutineRecordsResponseBody) Validate() error {
	if s.Records != nil {
		for _, item := range s.Records {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListEdgeRoutineRecordsResponseBodyRecords struct {
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	RecordCname *string `json:"RecordCname,omitempty" xml:"RecordCname,omitempty"`
	RecordName  *string `json:"RecordName,omitempty" xml:"RecordName,omitempty"`
	SiteId      *int64  `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	SiteName    *string `json:"SiteName,omitempty" xml:"SiteName,omitempty"`
	UpdateTime  *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListEdgeRoutineRecordsResponseBodyRecords) String() string {
	return dara.Prettify(s)
}

func (s ListEdgeRoutineRecordsResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *ListEdgeRoutineRecordsResponseBodyRecords) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListEdgeRoutineRecordsResponseBodyRecords) GetRecordCname() *string {
	return s.RecordCname
}

func (s *ListEdgeRoutineRecordsResponseBodyRecords) GetRecordName() *string {
	return s.RecordName
}

func (s *ListEdgeRoutineRecordsResponseBodyRecords) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListEdgeRoutineRecordsResponseBodyRecords) GetSiteName() *string {
	return s.SiteName
}

func (s *ListEdgeRoutineRecordsResponseBodyRecords) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListEdgeRoutineRecordsResponseBodyRecords) SetCreateTime(v string) *ListEdgeRoutineRecordsResponseBodyRecords {
	s.CreateTime = &v
	return s
}

func (s *ListEdgeRoutineRecordsResponseBodyRecords) SetRecordCname(v string) *ListEdgeRoutineRecordsResponseBodyRecords {
	s.RecordCname = &v
	return s
}

func (s *ListEdgeRoutineRecordsResponseBodyRecords) SetRecordName(v string) *ListEdgeRoutineRecordsResponseBodyRecords {
	s.RecordName = &v
	return s
}

func (s *ListEdgeRoutineRecordsResponseBodyRecords) SetSiteId(v int64) *ListEdgeRoutineRecordsResponseBodyRecords {
	s.SiteId = &v
	return s
}

func (s *ListEdgeRoutineRecordsResponseBodyRecords) SetSiteName(v string) *ListEdgeRoutineRecordsResponseBodyRecords {
	s.SiteName = &v
	return s
}

func (s *ListEdgeRoutineRecordsResponseBodyRecords) SetUpdateTime(v string) *ListEdgeRoutineRecordsResponseBodyRecords {
	s.UpdateTime = &v
	return s
}

func (s *ListEdgeRoutineRecordsResponseBodyRecords) Validate() error {
	return dara.Validate(s)
}
