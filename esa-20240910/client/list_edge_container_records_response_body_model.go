// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEdgeContainerRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListEdgeContainerRecordsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListEdgeContainerRecordsResponseBody
	GetPageSize() *int32
	SetRecords(v []*ListEdgeContainerRecordsResponseBodyRecords) *ListEdgeContainerRecordsResponseBody
	GetRecords() []*ListEdgeContainerRecordsResponseBodyRecords
	SetRequestId(v string) *ListEdgeContainerRecordsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListEdgeContainerRecordsResponseBody
	GetTotalCount() *int32
}

type ListEdgeContainerRecordsResponseBody struct {
	PageNumber *int32                                         `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                         `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Records    []*ListEdgeContainerRecordsResponseBodyRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	RequestId  *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                         `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListEdgeContainerRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEdgeContainerRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *ListEdgeContainerRecordsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListEdgeContainerRecordsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListEdgeContainerRecordsResponseBody) GetRecords() []*ListEdgeContainerRecordsResponseBodyRecords {
	return s.Records
}

func (s *ListEdgeContainerRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEdgeContainerRecordsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListEdgeContainerRecordsResponseBody) SetPageNumber(v int32) *ListEdgeContainerRecordsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListEdgeContainerRecordsResponseBody) SetPageSize(v int32) *ListEdgeContainerRecordsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListEdgeContainerRecordsResponseBody) SetRecords(v []*ListEdgeContainerRecordsResponseBodyRecords) *ListEdgeContainerRecordsResponseBody {
	s.Records = v
	return s
}

func (s *ListEdgeContainerRecordsResponseBody) SetRequestId(v string) *ListEdgeContainerRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEdgeContainerRecordsResponseBody) SetTotalCount(v int32) *ListEdgeContainerRecordsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListEdgeContainerRecordsResponseBody) Validate() error {
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

type ListEdgeContainerRecordsResponseBodyRecords struct {
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	RecordCname *string `json:"RecordCname,omitempty" xml:"RecordCname,omitempty"`
	RecordName  *string `json:"RecordName,omitempty" xml:"RecordName,omitempty"`
	SiteId      *int64  `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	SiteName    *string `json:"SiteName,omitempty" xml:"SiteName,omitempty"`
	UpdateTime  *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListEdgeContainerRecordsResponseBodyRecords) String() string {
	return dara.Prettify(s)
}

func (s ListEdgeContainerRecordsResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *ListEdgeContainerRecordsResponseBodyRecords) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListEdgeContainerRecordsResponseBodyRecords) GetRecordCname() *string {
	return s.RecordCname
}

func (s *ListEdgeContainerRecordsResponseBodyRecords) GetRecordName() *string {
	return s.RecordName
}

func (s *ListEdgeContainerRecordsResponseBodyRecords) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListEdgeContainerRecordsResponseBodyRecords) GetSiteName() *string {
	return s.SiteName
}

func (s *ListEdgeContainerRecordsResponseBodyRecords) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListEdgeContainerRecordsResponseBodyRecords) SetCreateTime(v string) *ListEdgeContainerRecordsResponseBodyRecords {
	s.CreateTime = &v
	return s
}

func (s *ListEdgeContainerRecordsResponseBodyRecords) SetRecordCname(v string) *ListEdgeContainerRecordsResponseBodyRecords {
	s.RecordCname = &v
	return s
}

func (s *ListEdgeContainerRecordsResponseBodyRecords) SetRecordName(v string) *ListEdgeContainerRecordsResponseBodyRecords {
	s.RecordName = &v
	return s
}

func (s *ListEdgeContainerRecordsResponseBodyRecords) SetSiteId(v int64) *ListEdgeContainerRecordsResponseBodyRecords {
	s.SiteId = &v
	return s
}

func (s *ListEdgeContainerRecordsResponseBodyRecords) SetSiteName(v string) *ListEdgeContainerRecordsResponseBodyRecords {
	s.SiteName = &v
	return s
}

func (s *ListEdgeContainerRecordsResponseBodyRecords) SetUpdateTime(v string) *ListEdgeContainerRecordsResponseBodyRecords {
	s.UpdateTime = &v
	return s
}

func (s *ListEdgeContainerRecordsResponseBodyRecords) Validate() error {
	return dara.Validate(s)
}
