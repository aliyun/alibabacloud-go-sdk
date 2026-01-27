// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessType(v string) *DescribeUserResourcesRequest
	GetAccessType() *string
	SetAutoRefresh(v bool) *DescribeUserResourcesRequest
	GetAutoRefresh() *bool
	SetCategoryId(v int32) *DescribeUserResourcesRequest
	GetCategoryId() *int32
	SetCategoryType(v int32) *DescribeUserResourcesRequest
	GetCategoryType() *int32
	SetClientId(v string) *DescribeUserResourcesRequest
	GetClientId() *string
	SetClientType(v string) *DescribeUserResourcesRequest
	GetClientType() *string
	SetClientVersion(v string) *DescribeUserResourcesRequest
	GetClientVersion() *string
	SetDualCenterForward(v bool) *DescribeUserResourcesRequest
	GetDualCenterForward() *bool
	SetLanguage(v string) *DescribeUserResourcesRequest
	GetLanguage() *string
	SetLoginRegionId(v string) *DescribeUserResourcesRequest
	GetLoginRegionId() *string
	SetLoginToken(v string) *DescribeUserResourcesRequest
	GetLoginToken() *string
	SetMaxResults(v int32) *DescribeUserResourcesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeUserResourcesRequest
	GetNextToken() *string
	SetOfficeSiteIds(v []*string) *DescribeUserResourcesRequest
	GetOfficeSiteIds() []*string
	SetOrderBy(v string) *DescribeUserResourcesRequest
	GetOrderBy() *string
	SetProductTypes(v []*string) *DescribeUserResourcesRequest
	GetProductTypes() []*string
	SetProtocolType(v string) *DescribeUserResourcesRequest
	GetProtocolType() *string
	SetQueryDesktopDurationList(v bool) *DescribeUserResourcesRequest
	GetQueryDesktopDurationList() *bool
	SetQueryDesktopTimers(v bool) *DescribeUserResourcesRequest
	GetQueryDesktopTimers() *bool
	SetQueryFotaUpdate(v bool) *DescribeUserResourcesRequest
	GetQueryFotaUpdate() *bool
	SetRefreshFotaUpdate(v bool) *DescribeUserResourcesRequest
	GetRefreshFotaUpdate() *bool
	SetResourceIds(v []*string) *DescribeUserResourcesRequest
	GetResourceIds() []*string
	SetResourceName(v string) *DescribeUserResourcesRequest
	GetResourceName() *string
	SetResourceTypes(v []*string) *DescribeUserResourcesRequest
	GetResourceTypes() []*string
	SetScene(v string) *DescribeUserResourcesRequest
	GetScene() *string
	SetSearchRegionId(v string) *DescribeUserResourcesRequest
	GetSearchRegionId() *string
	SetSessionId(v string) *DescribeUserResourcesRequest
	GetSessionId() *string
	SetSortType(v string) *DescribeUserResourcesRequest
	GetSortType() *string
}

type DescribeUserResourcesRequest struct {
	// The access type. If you leave this parameter empty, both types will be displayed.
	//
	// Valid values:
	//
	// 	- INTERNET: access over the Internet.
	//
	// 	- VPC: access over an enterprise virtual private cloud (VPC).
	//
	// example:
	//
	// INTERNET
	AccessType *string `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	// Specifies whether to enable the immediate refresh feature.
	//
	// >  To ensure the operation response speed, we recommend that you set the value to `false`.
	//
	// Valid values:
	//
	// 	- false
	//
	// 	- true
	//
	// example:
	//
	// false
	AutoRefresh *bool `json:"AutoRefresh,omitempty" xml:"AutoRefresh,omitempty"`
	// The level-2 resource category.
	//
	// example:
	//
	// 0
	CategoryId *int32 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// The level-1 resource category.
	//
	// example:
	//
	// 1
	CategoryType *int32 `json:"CategoryType,omitempty" xml:"CategoryType,omitempty"`
	// The client ID. The system generates a unique ID for each client. This parameter is non-sensitive and does not need encryption.
	//
	// This parameter is required.
	//
	// example:
	//
	// 54c17e1d-2d72-4b87-aa33-25f3b3f2****
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The client type.
	//
	// Valid values:
	//
	// 	- html5: the web client.
	//
	// 	- android: the Android client.
	//
	// 	- windows: the Windows client.
	//
	// 	- ios: the iOS client.
	//
	// 	- macos: the macOS client.
	//
	// example:
	//
	// windows
	ClientType *string `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	// The client version.
	//
	// example:
	//
	// 7.6.0-R-20241112.222305
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	// Deprecated
	//
	// Specifies whether to enable geo-redundant forwarding. Set the value to `false`.
	//
	// Valid value:
	//
	// 	- false: disables geo-redundant forwarding.
	//
	// example:
	//
	// false
	DualCenterForward *bool `json:"DualCenterForward,omitempty" xml:"DualCenterForward,omitempty"`
	// The client language.
	//
	// Valid values:
	//
	// 	- en_US: English.
	//
	// 	- zh_CN: Simplified Chinese.
	//
	// 	- ja_JP: Japanese.
	//
	// example:
	//
	// zh-CN
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The ID of the region where end users log on to clients. This parameter applies to office network ID-based logons. For organization ID-based logons, you can leave this parameter empty.
	//
	// example:
	//
	// cn-hangzhou
	LoginRegionId *string `json:"LoginRegionId,omitempty" xml:"LoginRegionId,omitempty"`
	// The logon token. You can call the `GetLoginToken` or `RefreshLoginToken` operation to retrieve the logon token.
	//
	// This parameter is required.
	//
	// example:
	//
	// v18101ac6a9e69c66b04a163031680463660b4b216cd758f34b60b9ad6a7c7f7334b83dd8f75eef4209c68f9f1080b****
	LoginToken *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	// The number of entries per page. Default value: 500.
	//
	// example:
	//
	// 500
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The office network IDs. This parameter is required for logons by using enterprise Active Directory (AD) accounts.
	OfficeSiteIds []*string `json:"OfficeSiteIds,omitempty" xml:"OfficeSiteIds,omitempty" type:"Repeated"`
	// The sorting field. If you do not specify this parameter, resources are sorted by creation time by default.
	//
	// Valid values:
	//
	// 	- CreateTime: the time when resources are created.
	//
	// 	- AssignTime: the time when resources are assigned.
	//
	// example:
	//
	// AssignTime
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The service types. If you leave this parameter empty, services of all supported types will be queried.
	ProductTypes []*string `json:"ProductTypes,omitempty" xml:"ProductTypes,omitempty" type:"Repeated"`
	// The protocol type. You can specify this parameter to filter cloud computers.
	//
	// Valid values:
	//
	// 	- HDX: High-definition Experience (HDX).
	//
	// 	- ASP: Adaptive Streaming Protocol (ASP).
	//
	// example:
	//
	// ASP
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	// 是否查询云桌面套餐包信息，默认为true。
	//
	// if can be null:
	// true
	//
	// example:
	//
	// true
	QueryDesktopDurationList *bool `json:"QueryDesktopDurationList,omitempty" xml:"QueryDesktopDurationList,omitempty"`
	// 是否查询云电脑定时任务信息，默认为true。
	//
	// if can be null:
	// true
	//
	// example:
	//
	// true
	QueryDesktopTimers *bool `json:"QueryDesktopTimers,omitempty" xml:"QueryDesktopTimers,omitempty"`
	// Specifies whether to return the image version information of cloud computers.
	//
	// example:
	//
	// false
	QueryFotaUpdate *bool `json:"QueryFotaUpdate,omitempty" xml:"QueryFotaUpdate,omitempty"`
	// Specifies whether to refresh over-the-air (OTA) information in real time.
	//
	// example:
	//
	// false
	RefreshFotaUpdate *bool `json:"RefreshFotaUpdate,omitempty" xml:"RefreshFotaUpdate,omitempty"`
	// The resource IDs. You can specify up to 100 resource IDs.
	ResourceIds []*string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty" type:"Repeated"`
	// The resource name. Fuzzy search is supported.
	//
	// example:
	//
	// testName
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The resource types. If you leave this parameter empty, resources of all supported types will be queried.
	ResourceTypes []*string `json:"ResourceTypes,omitempty" xml:"ResourceTypes,omitempty" type:"Repeated"`
	// The client usage scenario. Set the value to `desktop`.
	//
	// Valid value:
	//
	// 	- desktop: cloud computers.
	//
	// example:
	//
	// desktop
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// The ID of the searched region. You can specify this parameter to filter cloud resources in specific regions.
	//
	// example:
	//
	// cn-hangzhou
	SearchRegionId *string `json:"SearchRegionId,omitempty" xml:"SearchRegionId,omitempty"`
	// The session ID. You can call the `GetLoginToken` operation to retrieve the session ID.
	//
	// example:
	//
	// cd45e873-650d-4d70-acb9-f996187a****
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The sorting method.
	//
	// Valid values:
	//
	// 	- ASC (default): the ascending order.
	//
	// 	- DESC: the descending order.
	//
	// example:
	//
	// ASC
	SortType *string `json:"SortType,omitempty" xml:"SortType,omitempty"`
}

func (s DescribeUserResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserResourcesRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserResourcesRequest) GetAccessType() *string {
	return s.AccessType
}

func (s *DescribeUserResourcesRequest) GetAutoRefresh() *bool {
	return s.AutoRefresh
}

func (s *DescribeUserResourcesRequest) GetCategoryId() *int32 {
	return s.CategoryId
}

func (s *DescribeUserResourcesRequest) GetCategoryType() *int32 {
	return s.CategoryType
}

func (s *DescribeUserResourcesRequest) GetClientId() *string {
	return s.ClientId
}

func (s *DescribeUserResourcesRequest) GetClientType() *string {
	return s.ClientType
}

func (s *DescribeUserResourcesRequest) GetClientVersion() *string {
	return s.ClientVersion
}

func (s *DescribeUserResourcesRequest) GetDualCenterForward() *bool {
	return s.DualCenterForward
}

func (s *DescribeUserResourcesRequest) GetLanguage() *string {
	return s.Language
}

func (s *DescribeUserResourcesRequest) GetLoginRegionId() *string {
	return s.LoginRegionId
}

func (s *DescribeUserResourcesRequest) GetLoginToken() *string {
	return s.LoginToken
}

func (s *DescribeUserResourcesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeUserResourcesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeUserResourcesRequest) GetOfficeSiteIds() []*string {
	return s.OfficeSiteIds
}

func (s *DescribeUserResourcesRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *DescribeUserResourcesRequest) GetProductTypes() []*string {
	return s.ProductTypes
}

func (s *DescribeUserResourcesRequest) GetProtocolType() *string {
	return s.ProtocolType
}

func (s *DescribeUserResourcesRequest) GetQueryDesktopDurationList() *bool {
	return s.QueryDesktopDurationList
}

func (s *DescribeUserResourcesRequest) GetQueryDesktopTimers() *bool {
	return s.QueryDesktopTimers
}

func (s *DescribeUserResourcesRequest) GetQueryFotaUpdate() *bool {
	return s.QueryFotaUpdate
}

func (s *DescribeUserResourcesRequest) GetRefreshFotaUpdate() *bool {
	return s.RefreshFotaUpdate
}

func (s *DescribeUserResourcesRequest) GetResourceIds() []*string {
	return s.ResourceIds
}

func (s *DescribeUserResourcesRequest) GetResourceName() *string {
	return s.ResourceName
}

func (s *DescribeUserResourcesRequest) GetResourceTypes() []*string {
	return s.ResourceTypes
}

func (s *DescribeUserResourcesRequest) GetScene() *string {
	return s.Scene
}

func (s *DescribeUserResourcesRequest) GetSearchRegionId() *string {
	return s.SearchRegionId
}

func (s *DescribeUserResourcesRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *DescribeUserResourcesRequest) GetSortType() *string {
	return s.SortType
}

func (s *DescribeUserResourcesRequest) SetAccessType(v string) *DescribeUserResourcesRequest {
	s.AccessType = &v
	return s
}

func (s *DescribeUserResourcesRequest) SetAutoRefresh(v bool) *DescribeUserResourcesRequest {
	s.AutoRefresh = &v
	return s
}

func (s *DescribeUserResourcesRequest) SetCategoryId(v int32) *DescribeUserResourcesRequest {
	s.CategoryId = &v
	return s
}

func (s *DescribeUserResourcesRequest) SetCategoryType(v int32) *DescribeUserResourcesRequest {
	s.CategoryType = &v
	return s
}

func (s *DescribeUserResourcesRequest) SetClientId(v string) *DescribeUserResourcesRequest {
	s.ClientId = &v
	return s
}

func (s *DescribeUserResourcesRequest) SetClientType(v string) *DescribeUserResourcesRequest {
	s.ClientType = &v
	return s
}

func (s *DescribeUserResourcesRequest) SetClientVersion(v string) *DescribeUserResourcesRequest {
	s.ClientVersion = &v
	return s
}

func (s *DescribeUserResourcesRequest) SetDualCenterForward(v bool) *DescribeUserResourcesRequest {
	s.DualCenterForward = &v
	return s
}

func (s *DescribeUserResourcesRequest) SetLanguage(v string) *DescribeUserResourcesRequest {
	s.Language = &v
	return s
}

func (s *DescribeUserResourcesRequest) SetLoginRegionId(v string) *DescribeUserResourcesRequest {
	s.LoginRegionId = &v
	return s
}

func (s *DescribeUserResourcesRequest) SetLoginToken(v string) *DescribeUserResourcesRequest {
	s.LoginToken = &v
	return s
}

func (s *DescribeUserResourcesRequest) SetMaxResults(v int32) *DescribeUserResourcesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeUserResourcesRequest) SetNextToken(v string) *DescribeUserResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeUserResourcesRequest) SetOfficeSiteIds(v []*string) *DescribeUserResourcesRequest {
	s.OfficeSiteIds = v
	return s
}

func (s *DescribeUserResourcesRequest) SetOrderBy(v string) *DescribeUserResourcesRequest {
	s.OrderBy = &v
	return s
}

func (s *DescribeUserResourcesRequest) SetProductTypes(v []*string) *DescribeUserResourcesRequest {
	s.ProductTypes = v
	return s
}

func (s *DescribeUserResourcesRequest) SetProtocolType(v string) *DescribeUserResourcesRequest {
	s.ProtocolType = &v
	return s
}

func (s *DescribeUserResourcesRequest) SetQueryDesktopDurationList(v bool) *DescribeUserResourcesRequest {
	s.QueryDesktopDurationList = &v
	return s
}

func (s *DescribeUserResourcesRequest) SetQueryDesktopTimers(v bool) *DescribeUserResourcesRequest {
	s.QueryDesktopTimers = &v
	return s
}

func (s *DescribeUserResourcesRequest) SetQueryFotaUpdate(v bool) *DescribeUserResourcesRequest {
	s.QueryFotaUpdate = &v
	return s
}

func (s *DescribeUserResourcesRequest) SetRefreshFotaUpdate(v bool) *DescribeUserResourcesRequest {
	s.RefreshFotaUpdate = &v
	return s
}

func (s *DescribeUserResourcesRequest) SetResourceIds(v []*string) *DescribeUserResourcesRequest {
	s.ResourceIds = v
	return s
}

func (s *DescribeUserResourcesRequest) SetResourceName(v string) *DescribeUserResourcesRequest {
	s.ResourceName = &v
	return s
}

func (s *DescribeUserResourcesRequest) SetResourceTypes(v []*string) *DescribeUserResourcesRequest {
	s.ResourceTypes = v
	return s
}

func (s *DescribeUserResourcesRequest) SetScene(v string) *DescribeUserResourcesRequest {
	s.Scene = &v
	return s
}

func (s *DescribeUserResourcesRequest) SetSearchRegionId(v string) *DescribeUserResourcesRequest {
	s.SearchRegionId = &v
	return s
}

func (s *DescribeUserResourcesRequest) SetSessionId(v string) *DescribeUserResourcesRequest {
	s.SessionId = &v
	return s
}

func (s *DescribeUserResourcesRequest) SetSortType(v string) *DescribeUserResourcesRequest {
	s.SortType = &v
	return s
}

func (s *DescribeUserResourcesRequest) Validate() error {
	return dara.Validate(s)
}
