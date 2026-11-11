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
	// The business scenario for record acceleration. This parameter is used as a filter condition for the query. Valid values:
	//
	// - **image_video**: images and videos
	//
	// - **api**: APIs
	//
	// - **web**: web pages
	//
	// example:
	//
	// web
	BizName *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	// Specifies whether a custom port is enabled for the record. Valid values:
	//
	// - on: enabled
	//
	// - off: disabled
	//
	// example:
	//
	// off
	CustomPort *string `json:"CustomPort,omitempty" xml:"CustomPort,omitempty"`
	// The page number for paging. The value must be greater than or equal to 1. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page for paging. Valid values: 1 to 500. Default value: **500**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Specifies whether proxy acceleration is enabled for the record. This parameter is used as a filter condition for the query. Valid values:
	//
	// - **true**: Proxy acceleration is enabled.
	//
	// - **false**: Proxy acceleration is disabled.
	//
	// example:
	//
	// true
	Proxied *bool `json:"Proxied,omitempty" xml:"Proxied,omitempty"`
	// The search match mode for the record name. Default value: exact. Valid values:
	//
	// - **prefix**: prefix match
	//
	// - **suffix**: suffix match
	//
	// - **exact**: exact match
	//
	// - **fuzzy**: fuzzy match
	//
	// example:
	//
	// fuzzy
	RecordMatchType *string `json:"RecordMatchType,omitempty" xml:"RecordMatchType,omitempty"`
	// The record name. This parameter is used as a filter condition for the query.
	//
	// example:
	//
	// www.example.com
	RecordName *string `json:"RecordName,omitempty" xml:"RecordName,omitempty"`
	// The site ID. You can obtain the site ID by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234567890123
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The origin server type of the record. This parameter is used as a filter condition for the query. Only CNAME records can be filtered by this field. Valid values:
	//
	// - **OSS**: OSS origin server
	//
	// - **S3**: S3 origin server
	//
	// - **LB**: load balancing origin server
	//
	// - **OP**: IPAM pool origin server
	//
	// - **Domain**: regular domain name origin server
	//
	// example:
	//
	// OSS
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The DNS record type of the record. This parameter is used as a filter condition for the query.
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
