// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizName(v string) *ListRecordsRequest
	GetBizName() *string
	SetPageNumber(v int32) *ListRecordsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListRecordsRequest
	GetPageSize() *int32
	SetProxied(v bool) *ListRecordsRequest
	GetProxied() *bool
	SetRecordMatchType(v string) *ListRecordsRequest
	GetRecordMatchType() *string
	SetRecordName(v string) *ListRecordsRequest
	GetRecordName() *string
	SetSiteId(v int64) *ListRecordsRequest
	GetSiteId() *int64
	SetSourceType(v string) *ListRecordsRequest
	GetSourceType() *string
	SetType(v string) *ListRecordsRequest
	GetType() *string
}

type ListRecordsRequest struct {
	BizName         *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	PageNumber      *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Proxied         *bool   `json:"Proxied,omitempty" xml:"Proxied,omitempty"`
	RecordMatchType *string `json:"RecordMatchType,omitempty" xml:"RecordMatchType,omitempty"`
	RecordName      *string `json:"RecordName,omitempty" xml:"RecordName,omitempty"`
	// This parameter is required.
	SiteId     *int64  `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	Type       *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRecordsRequest) GoString() string {
	return s.String()
}

func (s *ListRecordsRequest) GetBizName() *string {
	return s.BizName
}

func (s *ListRecordsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListRecordsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRecordsRequest) GetProxied() *bool {
	return s.Proxied
}

func (s *ListRecordsRequest) GetRecordMatchType() *string {
	return s.RecordMatchType
}

func (s *ListRecordsRequest) GetRecordName() *string {
	return s.RecordName
}

func (s *ListRecordsRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListRecordsRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *ListRecordsRequest) GetType() *string {
	return s.Type
}

func (s *ListRecordsRequest) SetBizName(v string) *ListRecordsRequest {
	s.BizName = &v
	return s
}

func (s *ListRecordsRequest) SetPageNumber(v int32) *ListRecordsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListRecordsRequest) SetPageSize(v int32) *ListRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *ListRecordsRequest) SetProxied(v bool) *ListRecordsRequest {
	s.Proxied = &v
	return s
}

func (s *ListRecordsRequest) SetRecordMatchType(v string) *ListRecordsRequest {
	s.RecordMatchType = &v
	return s
}

func (s *ListRecordsRequest) SetRecordName(v string) *ListRecordsRequest {
	s.RecordName = &v
	return s
}

func (s *ListRecordsRequest) SetSiteId(v int64) *ListRecordsRequest {
	s.SiteId = &v
	return s
}

func (s *ListRecordsRequest) SetSourceType(v string) *ListRecordsRequest {
	s.SourceType = &v
	return s
}

func (s *ListRecordsRequest) SetType(v string) *ListRecordsRequest {
	s.Type = &v
	return s
}

func (s *ListRecordsRequest) Validate() error {
	return dara.Validate(s)
}
