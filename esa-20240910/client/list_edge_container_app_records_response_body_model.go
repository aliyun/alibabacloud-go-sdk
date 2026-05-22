// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEdgeContainerAppRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListEdgeContainerAppRecordsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListEdgeContainerAppRecordsResponseBody
	GetPageSize() *int32
	SetRecords(v []*ListEdgeContainerAppRecordsResponseBodyRecords) *ListEdgeContainerAppRecordsResponseBody
	GetRecords() []*ListEdgeContainerAppRecordsResponseBodyRecords
	SetRequestId(v string) *ListEdgeContainerAppRecordsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListEdgeContainerAppRecordsResponseBody
	GetTotalCount() *int32
}

type ListEdgeContainerAppRecordsResponseBody struct {
	PageNumber *int32                                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Records    []*ListEdgeContainerAppRecordsResponseBodyRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	RequestId  *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListEdgeContainerAppRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEdgeContainerAppRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *ListEdgeContainerAppRecordsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListEdgeContainerAppRecordsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListEdgeContainerAppRecordsResponseBody) GetRecords() []*ListEdgeContainerAppRecordsResponseBodyRecords {
	return s.Records
}

func (s *ListEdgeContainerAppRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEdgeContainerAppRecordsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListEdgeContainerAppRecordsResponseBody) SetPageNumber(v int32) *ListEdgeContainerAppRecordsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListEdgeContainerAppRecordsResponseBody) SetPageSize(v int32) *ListEdgeContainerAppRecordsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListEdgeContainerAppRecordsResponseBody) SetRecords(v []*ListEdgeContainerAppRecordsResponseBodyRecords) *ListEdgeContainerAppRecordsResponseBody {
	s.Records = v
	return s
}

func (s *ListEdgeContainerAppRecordsResponseBody) SetRequestId(v string) *ListEdgeContainerAppRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEdgeContainerAppRecordsResponseBody) SetTotalCount(v int32) *ListEdgeContainerAppRecordsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListEdgeContainerAppRecordsResponseBody) Validate() error {
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

type ListEdgeContainerAppRecordsResponseBodyRecords struct {
	AppId      *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Cname      *string `json:"Cname,omitempty" xml:"Cname,omitempty"`
	ConfigId   *int64  `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	RecordId   *int64  `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	RecordName *string `json:"RecordName,omitempty" xml:"RecordName,omitempty"`
	SchemdId   *int32  `json:"SchemdId,omitempty" xml:"SchemdId,omitempty"`
	SiteId     *int64  `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListEdgeContainerAppRecordsResponseBodyRecords) String() string {
	return dara.Prettify(s)
}

func (s ListEdgeContainerAppRecordsResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *ListEdgeContainerAppRecordsResponseBodyRecords) GetAppId() *string {
	return s.AppId
}

func (s *ListEdgeContainerAppRecordsResponseBodyRecords) GetCname() *string {
	return s.Cname
}

func (s *ListEdgeContainerAppRecordsResponseBodyRecords) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *ListEdgeContainerAppRecordsResponseBodyRecords) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListEdgeContainerAppRecordsResponseBodyRecords) GetRecordId() *int64 {
	return s.RecordId
}

func (s *ListEdgeContainerAppRecordsResponseBodyRecords) GetRecordName() *string {
	return s.RecordName
}

func (s *ListEdgeContainerAppRecordsResponseBodyRecords) GetSchemdId() *int32 {
	return s.SchemdId
}

func (s *ListEdgeContainerAppRecordsResponseBodyRecords) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListEdgeContainerAppRecordsResponseBodyRecords) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListEdgeContainerAppRecordsResponseBodyRecords) SetAppId(v string) *ListEdgeContainerAppRecordsResponseBodyRecords {
	s.AppId = &v
	return s
}

func (s *ListEdgeContainerAppRecordsResponseBodyRecords) SetCname(v string) *ListEdgeContainerAppRecordsResponseBodyRecords {
	s.Cname = &v
	return s
}

func (s *ListEdgeContainerAppRecordsResponseBodyRecords) SetConfigId(v int64) *ListEdgeContainerAppRecordsResponseBodyRecords {
	s.ConfigId = &v
	return s
}

func (s *ListEdgeContainerAppRecordsResponseBodyRecords) SetCreateTime(v string) *ListEdgeContainerAppRecordsResponseBodyRecords {
	s.CreateTime = &v
	return s
}

func (s *ListEdgeContainerAppRecordsResponseBodyRecords) SetRecordId(v int64) *ListEdgeContainerAppRecordsResponseBodyRecords {
	s.RecordId = &v
	return s
}

func (s *ListEdgeContainerAppRecordsResponseBodyRecords) SetRecordName(v string) *ListEdgeContainerAppRecordsResponseBodyRecords {
	s.RecordName = &v
	return s
}

func (s *ListEdgeContainerAppRecordsResponseBodyRecords) SetSchemdId(v int32) *ListEdgeContainerAppRecordsResponseBodyRecords {
	s.SchemdId = &v
	return s
}

func (s *ListEdgeContainerAppRecordsResponseBodyRecords) SetSiteId(v int64) *ListEdgeContainerAppRecordsResponseBodyRecords {
	s.SiteId = &v
	return s
}

func (s *ListEdgeContainerAppRecordsResponseBodyRecords) SetUpdateTime(v string) *ListEdgeContainerAppRecordsResponseBodyRecords {
	s.UpdateTime = &v
	return s
}

func (s *ListEdgeContainerAppRecordsResponseBodyRecords) Validate() error {
	return dara.Validate(s)
}
