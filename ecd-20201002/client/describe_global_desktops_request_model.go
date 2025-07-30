// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGlobalDesktopsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientId(v string) *DescribeGlobalDesktopsRequest
	GetClientId() *string
	SetDesktopAccessType(v string) *DescribeGlobalDesktopsRequest
	GetDesktopAccessType() *string
	SetDesktopId(v []*string) *DescribeGlobalDesktopsRequest
	GetDesktopId() []*string
	SetDesktopName(v string) *DescribeGlobalDesktopsRequest
	GetDesktopName() *string
	SetDesktopStatus(v string) *DescribeGlobalDesktopsRequest
	GetDesktopStatus() *string
	SetDirectoryId(v string) *DescribeGlobalDesktopsRequest
	GetDirectoryId() *string
	SetKeyword(v string) *DescribeGlobalDesktopsRequest
	GetKeyword() *string
	SetLanguage(v string) *DescribeGlobalDesktopsRequest
	GetLanguage() *string
	SetLoginRegionId(v string) *DescribeGlobalDesktopsRequest
	GetLoginRegionId() *string
	SetLoginToken(v string) *DescribeGlobalDesktopsRequest
	GetLoginToken() *string
	SetMaxResults(v int32) *DescribeGlobalDesktopsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeGlobalDesktopsRequest
	GetNextToken() *string
	SetOfficeSiteId(v string) *DescribeGlobalDesktopsRequest
	GetOfficeSiteId() *string
	SetOrderBy(v string) *DescribeGlobalDesktopsRequest
	GetOrderBy() *string
	SetQueryFotaUpdate(v bool) *DescribeGlobalDesktopsRequest
	GetQueryFotaUpdate() *bool
	SetRegionId(v string) *DescribeGlobalDesktopsRequest
	GetRegionId() *string
	SetSearchRegionId(v string) *DescribeGlobalDesktopsRequest
	GetSearchRegionId() *string
	SetSessionId(v string) *DescribeGlobalDesktopsRequest
	GetSessionId() *string
	SetSortType(v string) *DescribeGlobalDesktopsRequest
	GetSortType() *string
	SetWithoutLatency(v bool) *DescribeGlobalDesktopsRequest
	GetWithoutLatency() *bool
}

type DescribeGlobalDesktopsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// c213150d-7ac3-432c-b749-6e1e090b****
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// example:
	//
	// INTERNET
	DesktopAccessType *string   `json:"DesktopAccessType,omitempty" xml:"DesktopAccessType,omitempty"`
	DesktopId         []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
	// example:
	//
	// DesktopTest
	DesktopName *string `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
	// example:
	//
	// Running
	DesktopStatus *string `json:"DesktopStatus,omitempty" xml:"DesktopStatus,omitempty"`
	// example:
	//
	// cn-hangzhou+dir-880841****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// 关键字。支持模糊搜索桌面ID、云桌面名称和终端用户自定义的桌面名称。
	//
	// example:
	//
	// ecd
	Keyword  *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// example:
	//
	// cn-hangzhou
	LoginRegionId *string `json:"LoginRegionId,omitempty" xml:"LoginRegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// v147c9114a180489f89691663893169****
	LoginToken *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	// example:
	//
	// 500
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// eyJkZWZhdWx0IjpbIjk2MjEy****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// cn-hangzhou+dir-880841****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// example:
	//
	// AssignTime
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// example:
	//
	// true
	QueryFotaUpdate *bool `json:"QueryFotaUpdate,omitempty" xml:"QueryFotaUpdate,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// cn-shanghai
	SearchRegionId *string `json:"SearchRegionId,omitempty" xml:"SearchRegionId,omitempty"`
	// example:
	//
	// 5c456a41-1e65-4e72-ab4d-5dcfff52****
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// ASC
	SortType *string `json:"SortType,omitempty" xml:"SortType,omitempty"`
	// example:
	//
	// true
	WithoutLatency *bool `json:"WithoutLatency,omitempty" xml:"WithoutLatency,omitempty"`
}

func (s DescribeGlobalDesktopsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeGlobalDesktopsRequest) GoString() string {
	return s.String()
}

func (s *DescribeGlobalDesktopsRequest) GetClientId() *string {
	return s.ClientId
}

func (s *DescribeGlobalDesktopsRequest) GetDesktopAccessType() *string {
	return s.DesktopAccessType
}

func (s *DescribeGlobalDesktopsRequest) GetDesktopId() []*string {
	return s.DesktopId
}

func (s *DescribeGlobalDesktopsRequest) GetDesktopName() *string {
	return s.DesktopName
}

func (s *DescribeGlobalDesktopsRequest) GetDesktopStatus() *string {
	return s.DesktopStatus
}

func (s *DescribeGlobalDesktopsRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *DescribeGlobalDesktopsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *DescribeGlobalDesktopsRequest) GetLanguage() *string {
	return s.Language
}

func (s *DescribeGlobalDesktopsRequest) GetLoginRegionId() *string {
	return s.LoginRegionId
}

func (s *DescribeGlobalDesktopsRequest) GetLoginToken() *string {
	return s.LoginToken
}

func (s *DescribeGlobalDesktopsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeGlobalDesktopsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeGlobalDesktopsRequest) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *DescribeGlobalDesktopsRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *DescribeGlobalDesktopsRequest) GetQueryFotaUpdate() *bool {
	return s.QueryFotaUpdate
}

func (s *DescribeGlobalDesktopsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeGlobalDesktopsRequest) GetSearchRegionId() *string {
	return s.SearchRegionId
}

func (s *DescribeGlobalDesktopsRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *DescribeGlobalDesktopsRequest) GetSortType() *string {
	return s.SortType
}

func (s *DescribeGlobalDesktopsRequest) GetWithoutLatency() *bool {
	return s.WithoutLatency
}

func (s *DescribeGlobalDesktopsRequest) SetClientId(v string) *DescribeGlobalDesktopsRequest {
	s.ClientId = &v
	return s
}

func (s *DescribeGlobalDesktopsRequest) SetDesktopAccessType(v string) *DescribeGlobalDesktopsRequest {
	s.DesktopAccessType = &v
	return s
}

func (s *DescribeGlobalDesktopsRequest) SetDesktopId(v []*string) *DescribeGlobalDesktopsRequest {
	s.DesktopId = v
	return s
}

func (s *DescribeGlobalDesktopsRequest) SetDesktopName(v string) *DescribeGlobalDesktopsRequest {
	s.DesktopName = &v
	return s
}

func (s *DescribeGlobalDesktopsRequest) SetDesktopStatus(v string) *DescribeGlobalDesktopsRequest {
	s.DesktopStatus = &v
	return s
}

func (s *DescribeGlobalDesktopsRequest) SetDirectoryId(v string) *DescribeGlobalDesktopsRequest {
	s.DirectoryId = &v
	return s
}

func (s *DescribeGlobalDesktopsRequest) SetKeyword(v string) *DescribeGlobalDesktopsRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeGlobalDesktopsRequest) SetLanguage(v string) *DescribeGlobalDesktopsRequest {
	s.Language = &v
	return s
}

func (s *DescribeGlobalDesktopsRequest) SetLoginRegionId(v string) *DescribeGlobalDesktopsRequest {
	s.LoginRegionId = &v
	return s
}

func (s *DescribeGlobalDesktopsRequest) SetLoginToken(v string) *DescribeGlobalDesktopsRequest {
	s.LoginToken = &v
	return s
}

func (s *DescribeGlobalDesktopsRequest) SetMaxResults(v int32) *DescribeGlobalDesktopsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeGlobalDesktopsRequest) SetNextToken(v string) *DescribeGlobalDesktopsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeGlobalDesktopsRequest) SetOfficeSiteId(v string) *DescribeGlobalDesktopsRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeGlobalDesktopsRequest) SetOrderBy(v string) *DescribeGlobalDesktopsRequest {
	s.OrderBy = &v
	return s
}

func (s *DescribeGlobalDesktopsRequest) SetQueryFotaUpdate(v bool) *DescribeGlobalDesktopsRequest {
	s.QueryFotaUpdate = &v
	return s
}

func (s *DescribeGlobalDesktopsRequest) SetRegionId(v string) *DescribeGlobalDesktopsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeGlobalDesktopsRequest) SetSearchRegionId(v string) *DescribeGlobalDesktopsRequest {
	s.SearchRegionId = &v
	return s
}

func (s *DescribeGlobalDesktopsRequest) SetSessionId(v string) *DescribeGlobalDesktopsRequest {
	s.SessionId = &v
	return s
}

func (s *DescribeGlobalDesktopsRequest) SetSortType(v string) *DescribeGlobalDesktopsRequest {
	s.SortType = &v
	return s
}

func (s *DescribeGlobalDesktopsRequest) SetWithoutLatency(v bool) *DescribeGlobalDesktopsRequest {
	s.WithoutLatency = &v
	return s
}

func (s *DescribeGlobalDesktopsRequest) Validate() error {
	return dara.Validate(s)
}
