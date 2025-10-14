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
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The details about the associated domain names.
	Records []*ListEdgeContainerAppRecordsResponseBodyRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// CB1A380B-09F0-41BB-A198-72F8FD6DA2FE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of domain names that are associated with the specified application.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	// The application ID.
	//
	// example:
	//
	// app-88068867578379****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The CNAME of the associated domain name.
	//
	// example:
	//
	// kdxceo****.yun****.com
	Cname *string `json:"Cname,omitempty" xml:"Cname,omitempty"`
	// The configuration ID of the associated domain name.
	//
	// example:
	//
	// 27522948436****
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The time when the domain name was added. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ss format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-12-24T02:01:11Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The record ID of the associated domain name.
	//
	// example:
	//
	// 266****
	RecordId *int64 `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// The associated domain name.
	//
	// example:
	//
	// a.example.com
	RecordName *string `json:"RecordName,omitempty" xml:"RecordName,omitempty"`
	// The scheduling domain ID of the associated domain name.
	//
	// example:
	//
	// 123456
	SchemdId *int32 `json:"SchemdId,omitempty" xml:"SchemdId,omitempty"`
	// The website ID.
	//
	// example:
	//
	// 5407498413****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The time when the scheduling domain ID or CNAME was last modified. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ss format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-12-22T08:32:02Z
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
