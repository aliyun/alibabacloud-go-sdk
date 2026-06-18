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
	SetCustomPort(v string) *ListRecordsRequest
	GetCustomPort() *string
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
	// The business scenario for acceleration. Use this parameter to filter results. Valid values:
	//
	// - **image_video**: Images and videos.
	//
	// - **api**: API.
	//
	// - **web**: Web page.
	//
	// example:
	//
	// web
	BizName    *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	CustomPort *string `json:"CustomPort,omitempty" xml:"CustomPort,omitempty"`
	// The page number. Defaults to **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The page size. Defaults to **500**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Filters the results based on whether the record is proxied. Valid values:
	//
	// - **true**: The record is proxied.
	//
	// - **false**: The record is not proxied.
	//
	// example:
	//
	// true
	Proxied *bool `json:"Proxied,omitempty" xml:"Proxied,omitempty"`
	// The match type for the record name search. Defaults to **exact**. Valid values:
	//
	// - **prefix**: Prefix match.
	//
	// - **suffix**: Suffix match.
	//
	// - **exact**: Exact match.
	//
	// - **fuzzy**: Fuzzy match.
	//
	// example:
	//
	// fuzzy
	RecordMatchType *string `json:"RecordMatchType,omitempty" xml:"RecordMatchType,omitempty"`
	// The record name. Use this parameter to filter query results.
	//
	// example:
	//
	// www.example.com
	RecordName *string `json:"RecordName,omitempty" xml:"RecordName,omitempty"`
	// The site ID. You can get this ID by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234567890123
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// Filters the results by the record\\"s origin type. This filter applies only to CNAME records. Valid values:
	//
	// - **OSS**: OSS origin.
	//
	// - **S3**: S3 origin.
	//
	// - **LB**: Load balancer origin.
	//
	// - **OP**: Origin pool.
	//
	// - **Domain**: Domain origin.
	//
	// example:
	//
	// OSS
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The DNS record type. Use this parameter to filter results.
	//
	// example:
	//
	// CNAME
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
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

func (s *ListRecordsRequest) GetCustomPort() *string {
	return s.CustomPort
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

func (s *ListRecordsRequest) SetCustomPort(v string) *ListRecordsRequest {
	s.CustomPort = &v
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
