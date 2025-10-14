// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRoutineRelatedRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int64) *ListRoutineRelatedRecordsResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *ListRoutineRelatedRecordsResponseBody
	GetPageSize() *int64
	SetRelatedRecords(v []*ListRoutineRelatedRecordsResponseBodyRelatedRecords) *ListRoutineRelatedRecordsResponseBody
	GetRelatedRecords() []*ListRoutineRelatedRecordsResponseBodyRelatedRecords
	SetRequestId(v string) *ListRoutineRelatedRecordsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListRoutineRelatedRecordsResponseBody
	GetTotalCount() *int64
}

type ListRoutineRelatedRecordsResponseBody struct {
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The list of records associated with the edge function.
	RelatedRecords []*ListRoutineRelatedRecordsResponseBodyRelatedRecords `json:"RelatedRecords,omitempty" xml:"RelatedRecords,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 16
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListRoutineRelatedRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRoutineRelatedRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRoutineRelatedRecordsResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListRoutineRelatedRecordsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListRoutineRelatedRecordsResponseBody) GetRelatedRecords() []*ListRoutineRelatedRecordsResponseBodyRelatedRecords {
	return s.RelatedRecords
}

func (s *ListRoutineRelatedRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRoutineRelatedRecordsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListRoutineRelatedRecordsResponseBody) SetPageNumber(v int64) *ListRoutineRelatedRecordsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListRoutineRelatedRecordsResponseBody) SetPageSize(v int64) *ListRoutineRelatedRecordsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListRoutineRelatedRecordsResponseBody) SetRelatedRecords(v []*ListRoutineRelatedRecordsResponseBodyRelatedRecords) *ListRoutineRelatedRecordsResponseBody {
	s.RelatedRecords = v
	return s
}

func (s *ListRoutineRelatedRecordsResponseBody) SetRequestId(v string) *ListRoutineRelatedRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRoutineRelatedRecordsResponseBody) SetTotalCount(v int64) *ListRoutineRelatedRecordsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListRoutineRelatedRecordsResponseBody) Validate() error {
	if s.RelatedRecords != nil {
		for _, item := range s.RelatedRecords {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListRoutineRelatedRecordsResponseBodyRelatedRecords struct {
	// The record ID of the domain name.
	//
	// example:
	//
	// 509348423011904
	RecordId *int64 `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// The name of the record.
	//
	// example:
	//
	// test-record-1.example.com
	RecordName *string `json:"RecordName,omitempty" xml:"RecordName,omitempty"`
	// The website ID.
	//
	// example:
	//
	// 54362329990032
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The website name.
	//
	// example:
	//
	// example.com
	SiteName *string `json:"SiteName,omitempty" xml:"SiteName,omitempty"`
}

func (s ListRoutineRelatedRecordsResponseBodyRelatedRecords) String() string {
	return dara.Prettify(s)
}

func (s ListRoutineRelatedRecordsResponseBodyRelatedRecords) GoString() string {
	return s.String()
}

func (s *ListRoutineRelatedRecordsResponseBodyRelatedRecords) GetRecordId() *int64 {
	return s.RecordId
}

func (s *ListRoutineRelatedRecordsResponseBodyRelatedRecords) GetRecordName() *string {
	return s.RecordName
}

func (s *ListRoutineRelatedRecordsResponseBodyRelatedRecords) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListRoutineRelatedRecordsResponseBodyRelatedRecords) GetSiteName() *string {
	return s.SiteName
}

func (s *ListRoutineRelatedRecordsResponseBodyRelatedRecords) SetRecordId(v int64) *ListRoutineRelatedRecordsResponseBodyRelatedRecords {
	s.RecordId = &v
	return s
}

func (s *ListRoutineRelatedRecordsResponseBodyRelatedRecords) SetRecordName(v string) *ListRoutineRelatedRecordsResponseBodyRelatedRecords {
	s.RecordName = &v
	return s
}

func (s *ListRoutineRelatedRecordsResponseBodyRelatedRecords) SetSiteId(v int64) *ListRoutineRelatedRecordsResponseBodyRelatedRecords {
	s.SiteId = &v
	return s
}

func (s *ListRoutineRelatedRecordsResponseBodyRelatedRecords) SetSiteName(v string) *ListRoutineRelatedRecordsResponseBodyRelatedRecords {
	s.SiteName = &v
	return s
}

func (s *ListRoutineRelatedRecordsResponseBodyRelatedRecords) Validate() error {
	return dara.Validate(s)
}
