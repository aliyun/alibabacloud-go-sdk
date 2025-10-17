// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *DescribeUserResourcesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeUserResourcesResponseBody
	GetNextToken() *string
	SetQueryFailedResourceTypes(v []*string) *DescribeUserResourcesResponseBody
	GetQueryFailedResourceTypes() []*string
	SetRankVersion(v int64) *DescribeUserResourcesResponseBody
	GetRankVersion() *int64
	SetRequestId(v string) *DescribeUserResourcesResponseBody
	GetRequestId() *string
	SetResources(v []*DescribeUserResourcesResponseBodyResources) *DescribeUserResourcesResponseBody
	GetResources() []*DescribeUserResourcesResponseBodyResources
	SetTotalCount(v int32) *DescribeUserResourcesResponseBody
	GetTotalCount() *int32
}

type DescribeUserResourcesResponseBody struct {
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6nmB7qrRFJ8vmttjxPL****
	NextToken                *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	QueryFailedResourceTypes []*string `json:"QueryFailedResourceTypes,omitempty" xml:"QueryFailedResourceTypes,omitempty" type:"Repeated"`
	// example:
	//
	// 1732869815062
	RankVersion *int64 `json:"RankVersion,omitempty" xml:"RankVersion,omitempty"`
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Resources []*DescribeUserResourcesResponseBodyResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeUserResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserResourcesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeUserResourcesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeUserResourcesResponseBody) GetQueryFailedResourceTypes() []*string {
	return s.QueryFailedResourceTypes
}

func (s *DescribeUserResourcesResponseBody) GetRankVersion() *int64 {
	return s.RankVersion
}

func (s *DescribeUserResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUserResourcesResponseBody) GetResources() []*DescribeUserResourcesResponseBodyResources {
	return s.Resources
}

func (s *DescribeUserResourcesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeUserResourcesResponseBody) SetMaxResults(v int32) *DescribeUserResourcesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeUserResourcesResponseBody) SetNextToken(v string) *DescribeUserResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeUserResourcesResponseBody) SetQueryFailedResourceTypes(v []*string) *DescribeUserResourcesResponseBody {
	s.QueryFailedResourceTypes = v
	return s
}

func (s *DescribeUserResourcesResponseBody) SetRankVersion(v int64) *DescribeUserResourcesResponseBody {
	s.RankVersion = &v
	return s
}

func (s *DescribeUserResourcesResponseBody) SetRequestId(v string) *DescribeUserResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserResourcesResponseBody) SetResources(v []*DescribeUserResourcesResponseBodyResources) *DescribeUserResourcesResponseBody {
	s.Resources = v
	return s
}

func (s *DescribeUserResourcesResponseBody) SetTotalCount(v int32) *DescribeUserResourcesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeUserResourcesResponseBody) Validate() error {
	if s.Resources != nil {
		for _, item := range s.Resources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeUserResourcesResponseBodyResources struct {
	// example:
	//
	// INTERNET
	AccessType *string `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	// example:
	//
	// 194101959****
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// example:
	//
	// app-0001
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// App
	AuthMode *string `json:"AuthMode,omitempty" xml:"AuthMode,omitempty"`
	// example:
	//
	// 0
	CategoryId *int32 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// example:
	//
	// 1
	CategoryType *int32 `json:"CategoryType,omitempty" xml:"CategoryType,omitempty"`
	// example:
	//
	// cn-shanghai+cds-695277****
	CdsName *string `json:"CdsName,omitempty" xml:"CdsName,omitempty"`
	// example:
	//
	// ecds-0****
	CenterResourceId *string `json:"CenterResourceId,omitempty" xml:"CenterResourceId,omitempty"`
	// example:
	//
	// PrePaid
	ChargeType *string                                              `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	Clients    []*DescribeUserResourcesResponseBodyResourcesClients `json:"Clients,omitempty" xml:"Clients,omitempty" type:"Repeated"`
	// example:
	//
	// {"authMode":"App"}
	ConnectionProperties *string `json:"ConnectionProperties,omitempty" xml:"ConnectionProperties,omitempty"`
	// example:
	//
	// 2024-12-11T07:12:12Z
	CreateTime          *string                                                          `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DesktopDurationList []*DescribeUserResourcesResponseBodyResourcesDesktopDurationList `json:"DesktopDurationList,omitempty" xml:"DesktopDurationList,omitempty" type:"Repeated"`
	DesktopTimers       []*DescribeUserResourcesResponseBodyResourcesDesktopTimers       `json:"DesktopTimers,omitempty" xml:"DesktopTimers,omitempty" type:"Repeated"`
	// example:
	//
	// 2025-02-22T16:00:00Z
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// example:
	//
	// stg114510
	ExternalDomainId *string `json:"ExternalDomainId,omitempty" xml:"ExternalDomainId,omitempty"`
	// example:
	//
	// test001
	ExternalUserId *string                                               `json:"ExternalUserId,omitempty" xml:"ExternalUserId,omitempty"`
	FotaUpdate     *DescribeUserResourcesResponseBodyResourcesFotaUpdate `json:"FotaUpdate,omitempty" xml:"FotaUpdate,omitempty" type:"Struct"`
	// example:
	//
	// true
	GlobalStatus *bool `json:"GlobalStatus,omitempty" xml:"GlobalStatus,omitempty"`
	HasUpgrade   *bool `json:"HasUpgrade,omitempty" xml:"HasUpgrade,omitempty"`
	// example:
	//
	// false
	HibernationBeta *bool `json:"HibernationBeta,omitempty" xml:"HibernationBeta,omitempty"`
	// example:
	//
	// http://example.com/icon.png
	Icon *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	// example:
	//
	// 2025-01-24T03:12:04Z
	LastStartTime      *string   `json:"LastStartTime,omitempty" xml:"LastStartTime,omitempty"`
	LocalName          *string   `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	ManagementStatuses []*string `json:"ManagementStatuses,omitempty" xml:"ManagementStatuses,omitempty" type:"Repeated"`
	// example:
	//
	// cn-shanghai+dir-3367****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// example:
	//
	// Normal
	OrderStatus *string `json:"OrderStatus,omitempty" xml:"OrderStatus,omitempty"`
	// example:
	//
	// Windows Server 2022
	Os            *string `json:"Os,omitempty" xml:"Os,omitempty"`
	OsDescription *string `json:"OsDescription,omitempty" xml:"OsDescription,omitempty"`
	// example:
	//
	// Windows
	OsType   *string                                             `json:"OsType,omitempty" xml:"OsType,omitempty"`
	OsUpdate *DescribeUserResourcesResponseBodyResourcesOsUpdate `json:"OsUpdate,omitempty" xml:"OsUpdate,omitempty" type:"Struct"`
	// example:
	//
	// AndroidCloud
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// example:
	//
	// ASP
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	// example:
	//
	// ecd-0001
	RealDesktopId *string `json:"RealDesktopId,omitempty" xml:"RealDesktopId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// Mainland
	RegionLocation *string `json:"RegionLocation,omitempty" xml:"RegionLocation,omitempty"`
	// example:
	//
	// dg-0****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// ecd-d19tya8zi4****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// Center
	ResourceLevel *string `json:"ResourceLevel,omitempty" xml:"ResourceLevel,omitempty"`
	// example:
	//
	// testName01
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// example:
	//
	// Connected
	ResourceSessionStatus *string `json:"ResourceSessionStatus,omitempty" xml:"ResourceSessionStatus,omitempty"`
	// example:
	//
	// Running
	ResourceStatus *string `json:"ResourceStatus,omitempty" xml:"ResourceStatus,omitempty"`
	// example:
	//
	// Desktop
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// SINGLE_SESSION
	SessionType *string                                               `json:"SessionType,omitempty" xml:"SessionType,omitempty"`
	Sessions    []*DescribeUserResourcesResponseBodyResourcesSessions `json:"Sessions,omitempty" xml:"Sessions,omitempty" type:"Repeated"`
	// example:
	//
	// PrePaid
	SubPayType *string `json:"SubPayType,omitempty" xml:"SubPayType,omitempty"`
	// example:
	//
	// true
	SupportHibernation *bool     `json:"SupportHibernation,omitempty" xml:"SupportHibernation,omitempty"`
	SupportedActions   []*string `json:"SupportedActions,omitempty" xml:"SupportedActions,omitempty" type:"Repeated"`
	// example:
	//
	// #FFFFFF
	ThemeColor     *string `json:"ThemeColor,omitempty" xml:"ThemeColor,omitempty"`
	UserCustomName *string `json:"UserCustomName,omitempty" xml:"UserCustomName,omitempty"`
	Version        *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeUserResourcesResponseBodyResources) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserResourcesResponseBodyResources) GoString() string {
	return s.String()
}

func (s *DescribeUserResourcesResponseBodyResources) GetAccessType() *string {
	return s.AccessType
}

func (s *DescribeUserResourcesResponseBodyResources) GetAliUid() *int64 {
	return s.AliUid
}

func (s *DescribeUserResourcesResponseBodyResources) GetAppId() *string {
	return s.AppId
}

func (s *DescribeUserResourcesResponseBodyResources) GetAuthMode() *string {
	return s.AuthMode
}

func (s *DescribeUserResourcesResponseBodyResources) GetCategoryId() *int32 {
	return s.CategoryId
}

func (s *DescribeUserResourcesResponseBodyResources) GetCategoryType() *int32 {
	return s.CategoryType
}

func (s *DescribeUserResourcesResponseBodyResources) GetCdsName() *string {
	return s.CdsName
}

func (s *DescribeUserResourcesResponseBodyResources) GetCenterResourceId() *string {
	return s.CenterResourceId
}

func (s *DescribeUserResourcesResponseBodyResources) GetChargeType() *string {
	return s.ChargeType
}

func (s *DescribeUserResourcesResponseBodyResources) GetClients() []*DescribeUserResourcesResponseBodyResourcesClients {
	return s.Clients
}

func (s *DescribeUserResourcesResponseBodyResources) GetConnectionProperties() *string {
	return s.ConnectionProperties
}

func (s *DescribeUserResourcesResponseBodyResources) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeUserResourcesResponseBodyResources) GetDesktopDurationList() []*DescribeUserResourcesResponseBodyResourcesDesktopDurationList {
	return s.DesktopDurationList
}

func (s *DescribeUserResourcesResponseBodyResources) GetDesktopTimers() []*DescribeUserResourcesResponseBodyResourcesDesktopTimers {
	return s.DesktopTimers
}

func (s *DescribeUserResourcesResponseBodyResources) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *DescribeUserResourcesResponseBodyResources) GetExternalDomainId() *string {
	return s.ExternalDomainId
}

func (s *DescribeUserResourcesResponseBodyResources) GetExternalUserId() *string {
	return s.ExternalUserId
}

func (s *DescribeUserResourcesResponseBodyResources) GetFotaUpdate() *DescribeUserResourcesResponseBodyResourcesFotaUpdate {
	return s.FotaUpdate
}

func (s *DescribeUserResourcesResponseBodyResources) GetGlobalStatus() *bool {
	return s.GlobalStatus
}

func (s *DescribeUserResourcesResponseBodyResources) GetHasUpgrade() *bool {
	return s.HasUpgrade
}

func (s *DescribeUserResourcesResponseBodyResources) GetHibernationBeta() *bool {
	return s.HibernationBeta
}

func (s *DescribeUserResourcesResponseBodyResources) GetIcon() *string {
	return s.Icon
}

func (s *DescribeUserResourcesResponseBodyResources) GetLastStartTime() *string {
	return s.LastStartTime
}

func (s *DescribeUserResourcesResponseBodyResources) GetLocalName() *string {
	return s.LocalName
}

func (s *DescribeUserResourcesResponseBodyResources) GetManagementStatuses() []*string {
	return s.ManagementStatuses
}

func (s *DescribeUserResourcesResponseBodyResources) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *DescribeUserResourcesResponseBodyResources) GetOrderStatus() *string {
	return s.OrderStatus
}

func (s *DescribeUserResourcesResponseBodyResources) GetOs() *string {
	return s.Os
}

func (s *DescribeUserResourcesResponseBodyResources) GetOsDescription() *string {
	return s.OsDescription
}

func (s *DescribeUserResourcesResponseBodyResources) GetOsType() *string {
	return s.OsType
}

func (s *DescribeUserResourcesResponseBodyResources) GetOsUpdate() *DescribeUserResourcesResponseBodyResourcesOsUpdate {
	return s.OsUpdate
}

func (s *DescribeUserResourcesResponseBodyResources) GetProductType() *string {
	return s.ProductType
}

func (s *DescribeUserResourcesResponseBodyResources) GetProtocolType() *string {
	return s.ProtocolType
}

func (s *DescribeUserResourcesResponseBodyResources) GetRealDesktopId() *string {
	return s.RealDesktopId
}

func (s *DescribeUserResourcesResponseBodyResources) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeUserResourcesResponseBodyResources) GetRegionLocation() *string {
	return s.RegionLocation
}

func (s *DescribeUserResourcesResponseBodyResources) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeUserResourcesResponseBodyResources) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeUserResourcesResponseBodyResources) GetResourceLevel() *string {
	return s.ResourceLevel
}

func (s *DescribeUserResourcesResponseBodyResources) GetResourceName() *string {
	return s.ResourceName
}

func (s *DescribeUserResourcesResponseBodyResources) GetResourceSessionStatus() *string {
	return s.ResourceSessionStatus
}

func (s *DescribeUserResourcesResponseBodyResources) GetResourceStatus() *string {
	return s.ResourceStatus
}

func (s *DescribeUserResourcesResponseBodyResources) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeUserResourcesResponseBodyResources) GetSessionType() *string {
	return s.SessionType
}

func (s *DescribeUserResourcesResponseBodyResources) GetSessions() []*DescribeUserResourcesResponseBodyResourcesSessions {
	return s.Sessions
}

func (s *DescribeUserResourcesResponseBodyResources) GetSubPayType() *string {
	return s.SubPayType
}

func (s *DescribeUserResourcesResponseBodyResources) GetSupportHibernation() *bool {
	return s.SupportHibernation
}

func (s *DescribeUserResourcesResponseBodyResources) GetSupportedActions() []*string {
	return s.SupportedActions
}

func (s *DescribeUserResourcesResponseBodyResources) GetThemeColor() *string {
	return s.ThemeColor
}

func (s *DescribeUserResourcesResponseBodyResources) GetUserCustomName() *string {
	return s.UserCustomName
}

func (s *DescribeUserResourcesResponseBodyResources) GetVersion() *string {
	return s.Version
}

func (s *DescribeUserResourcesResponseBodyResources) SetAccessType(v string) *DescribeUserResourcesResponseBodyResources {
	s.AccessType = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetAliUid(v int64) *DescribeUserResourcesResponseBodyResources {
	s.AliUid = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetAppId(v string) *DescribeUserResourcesResponseBodyResources {
	s.AppId = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetAuthMode(v string) *DescribeUserResourcesResponseBodyResources {
	s.AuthMode = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetCategoryId(v int32) *DescribeUserResourcesResponseBodyResources {
	s.CategoryId = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetCategoryType(v int32) *DescribeUserResourcesResponseBodyResources {
	s.CategoryType = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetCdsName(v string) *DescribeUserResourcesResponseBodyResources {
	s.CdsName = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetCenterResourceId(v string) *DescribeUserResourcesResponseBodyResources {
	s.CenterResourceId = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetChargeType(v string) *DescribeUserResourcesResponseBodyResources {
	s.ChargeType = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetClients(v []*DescribeUserResourcesResponseBodyResourcesClients) *DescribeUserResourcesResponseBodyResources {
	s.Clients = v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetConnectionProperties(v string) *DescribeUserResourcesResponseBodyResources {
	s.ConnectionProperties = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetCreateTime(v string) *DescribeUserResourcesResponseBodyResources {
	s.CreateTime = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetDesktopDurationList(v []*DescribeUserResourcesResponseBodyResourcesDesktopDurationList) *DescribeUserResourcesResponseBodyResources {
	s.DesktopDurationList = v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetDesktopTimers(v []*DescribeUserResourcesResponseBodyResourcesDesktopTimers) *DescribeUserResourcesResponseBodyResources {
	s.DesktopTimers = v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetExpiredTime(v string) *DescribeUserResourcesResponseBodyResources {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetExternalDomainId(v string) *DescribeUserResourcesResponseBodyResources {
	s.ExternalDomainId = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetExternalUserId(v string) *DescribeUserResourcesResponseBodyResources {
	s.ExternalUserId = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetFotaUpdate(v *DescribeUserResourcesResponseBodyResourcesFotaUpdate) *DescribeUserResourcesResponseBodyResources {
	s.FotaUpdate = v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetGlobalStatus(v bool) *DescribeUserResourcesResponseBodyResources {
	s.GlobalStatus = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetHasUpgrade(v bool) *DescribeUserResourcesResponseBodyResources {
	s.HasUpgrade = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetHibernationBeta(v bool) *DescribeUserResourcesResponseBodyResources {
	s.HibernationBeta = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetIcon(v string) *DescribeUserResourcesResponseBodyResources {
	s.Icon = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetLastStartTime(v string) *DescribeUserResourcesResponseBodyResources {
	s.LastStartTime = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetLocalName(v string) *DescribeUserResourcesResponseBodyResources {
	s.LocalName = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetManagementStatuses(v []*string) *DescribeUserResourcesResponseBodyResources {
	s.ManagementStatuses = v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetOfficeSiteId(v string) *DescribeUserResourcesResponseBodyResources {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetOrderStatus(v string) *DescribeUserResourcesResponseBodyResources {
	s.OrderStatus = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetOs(v string) *DescribeUserResourcesResponseBodyResources {
	s.Os = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetOsDescription(v string) *DescribeUserResourcesResponseBodyResources {
	s.OsDescription = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetOsType(v string) *DescribeUserResourcesResponseBodyResources {
	s.OsType = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetOsUpdate(v *DescribeUserResourcesResponseBodyResourcesOsUpdate) *DescribeUserResourcesResponseBodyResources {
	s.OsUpdate = v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetProductType(v string) *DescribeUserResourcesResponseBodyResources {
	s.ProductType = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetProtocolType(v string) *DescribeUserResourcesResponseBodyResources {
	s.ProtocolType = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetRealDesktopId(v string) *DescribeUserResourcesResponseBodyResources {
	s.RealDesktopId = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetRegionId(v string) *DescribeUserResourcesResponseBodyResources {
	s.RegionId = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetRegionLocation(v string) *DescribeUserResourcesResponseBodyResources {
	s.RegionLocation = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetResourceGroupId(v string) *DescribeUserResourcesResponseBodyResources {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetResourceId(v string) *DescribeUserResourcesResponseBodyResources {
	s.ResourceId = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetResourceLevel(v string) *DescribeUserResourcesResponseBodyResources {
	s.ResourceLevel = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetResourceName(v string) *DescribeUserResourcesResponseBodyResources {
	s.ResourceName = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetResourceSessionStatus(v string) *DescribeUserResourcesResponseBodyResources {
	s.ResourceSessionStatus = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetResourceStatus(v string) *DescribeUserResourcesResponseBodyResources {
	s.ResourceStatus = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetResourceType(v string) *DescribeUserResourcesResponseBodyResources {
	s.ResourceType = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetSessionType(v string) *DescribeUserResourcesResponseBodyResources {
	s.SessionType = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetSessions(v []*DescribeUserResourcesResponseBodyResourcesSessions) *DescribeUserResourcesResponseBodyResources {
	s.Sessions = v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetSubPayType(v string) *DescribeUserResourcesResponseBodyResources {
	s.SubPayType = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetSupportHibernation(v bool) *DescribeUserResourcesResponseBodyResources {
	s.SupportHibernation = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetSupportedActions(v []*string) *DescribeUserResourcesResponseBodyResources {
	s.SupportedActions = v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetThemeColor(v string) *DescribeUserResourcesResponseBodyResources {
	s.ThemeColor = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetUserCustomName(v string) *DescribeUserResourcesResponseBodyResources {
	s.UserCustomName = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) SetVersion(v string) *DescribeUserResourcesResponseBodyResources {
	s.Version = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResources) Validate() error {
	if s.Clients != nil {
		for _, item := range s.Clients {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.DesktopDurationList != nil {
		for _, item := range s.DesktopDurationList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.DesktopTimers != nil {
		for _, item := range s.DesktopTimers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.FotaUpdate != nil {
		if err := s.FotaUpdate.Validate(); err != nil {
			return err
		}
	}
	if s.OsUpdate != nil {
		if err := s.OsUpdate.Validate(); err != nil {
			return err
		}
	}
	if s.Sessions != nil {
		for _, item := range s.Sessions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeUserResourcesResponseBodyResourcesClients struct {
	// example:
	//
	// windows
	ClientType *string `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	// example:
	//
	// ON
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeUserResourcesResponseBodyResourcesClients) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserResourcesResponseBodyResourcesClients) GoString() string {
	return s.String()
}

func (s *DescribeUserResourcesResponseBodyResourcesClients) GetClientType() *string {
	return s.ClientType
}

func (s *DescribeUserResourcesResponseBodyResourcesClients) GetStatus() *string {
	return s.Status
}

func (s *DescribeUserResourcesResponseBodyResourcesClients) SetClientType(v string) *DescribeUserResourcesResponseBodyResourcesClients {
	s.ClientType = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResourcesClients) SetStatus(v string) *DescribeUserResourcesResponseBodyResourcesClients {
	s.Status = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResourcesClients) Validate() error {
	return dara.Validate(s)
}

type DescribeUserResourcesResponseBodyResourcesDesktopDurationList struct {
	// example:
	//
	// mdp-0bxls4qpi6bl6****
	OrderInstanceId *string `json:"OrderInstanceId,omitempty" xml:"OrderInstanceId,omitempty"`
	// example:
	//
	// 2025-01-17T07:01Z
	PackageCreationTime *string `json:"PackageCreationTime,omitempty" xml:"PackageCreationTime,omitempty"`
	// example:
	//
	// 2025-02-17T15:59Z
	PackageExpiredTime *string `json:"PackageExpiredTime,omitempty" xml:"PackageExpiredTime,omitempty"`
	// example:
	//
	// mdp-0bxls4qpi6bl6****
	PackageId *string `json:"PackageId,omitempty" xml:"PackageId,omitempty"`
	// example:
	//
	// Available
	PackageStatus *string `json:"PackageStatus,omitempty" xml:"PackageStatus,omitempty"`
	// example:
	//
	// NORMAL_PACKAGE
	PackageType *string `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	// example:
	//
	// Postpaid
	PackageUsedUpStrategy *string `json:"PackageUsedUpStrategy,omitempty" xml:"PackageUsedUpStrategy,omitempty"`
	// example:
	//
	// 2025-02-17T15:59Z
	PeriodEndTime *string `json:"PeriodEndTime,omitempty" xml:"PeriodEndTime,omitempty"`
	// example:
	//
	// 2025-01-17T07:01Z
	PeriodStartTime *string `json:"PeriodStartTime,omitempty" xml:"PeriodStartTime,omitempty"`
	// example:
	//
	// 199
	PostPaidLimitFee *float32 `json:"PostPaidLimitFee,omitempty" xml:"PostPaidLimitFee,omitempty"`
	// example:
	//
	// 432000
	TotalDuration *int64 `json:"TotalDuration,omitempty" xml:"TotalDuration,omitempty"`
	// example:
	//
	// 16850
	UsedDuration *int64 `json:"UsedDuration,omitempty" xml:"UsedDuration,omitempty"`
}

func (s DescribeUserResourcesResponseBodyResourcesDesktopDurationList) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserResourcesResponseBodyResourcesDesktopDurationList) GoString() string {
	return s.String()
}

func (s *DescribeUserResourcesResponseBodyResourcesDesktopDurationList) GetOrderInstanceId() *string {
	return s.OrderInstanceId
}

func (s *DescribeUserResourcesResponseBodyResourcesDesktopDurationList) GetPackageCreationTime() *string {
	return s.PackageCreationTime
}

func (s *DescribeUserResourcesResponseBodyResourcesDesktopDurationList) GetPackageExpiredTime() *string {
	return s.PackageExpiredTime
}

func (s *DescribeUserResourcesResponseBodyResourcesDesktopDurationList) GetPackageId() *string {
	return s.PackageId
}

func (s *DescribeUserResourcesResponseBodyResourcesDesktopDurationList) GetPackageStatus() *string {
	return s.PackageStatus
}

func (s *DescribeUserResourcesResponseBodyResourcesDesktopDurationList) GetPackageType() *string {
	return s.PackageType
}

func (s *DescribeUserResourcesResponseBodyResourcesDesktopDurationList) GetPackageUsedUpStrategy() *string {
	return s.PackageUsedUpStrategy
}

func (s *DescribeUserResourcesResponseBodyResourcesDesktopDurationList) GetPeriodEndTime() *string {
	return s.PeriodEndTime
}

func (s *DescribeUserResourcesResponseBodyResourcesDesktopDurationList) GetPeriodStartTime() *string {
	return s.PeriodStartTime
}

func (s *DescribeUserResourcesResponseBodyResourcesDesktopDurationList) GetPostPaidLimitFee() *float32 {
	return s.PostPaidLimitFee
}

func (s *DescribeUserResourcesResponseBodyResourcesDesktopDurationList) GetTotalDuration() *int64 {
	return s.TotalDuration
}

func (s *DescribeUserResourcesResponseBodyResourcesDesktopDurationList) GetUsedDuration() *int64 {
	return s.UsedDuration
}

func (s *DescribeUserResourcesResponseBodyResourcesDesktopDurationList) SetOrderInstanceId(v string) *DescribeUserResourcesResponseBodyResourcesDesktopDurationList {
	s.OrderInstanceId = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResourcesDesktopDurationList) SetPackageCreationTime(v string) *DescribeUserResourcesResponseBodyResourcesDesktopDurationList {
	s.PackageCreationTime = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResourcesDesktopDurationList) SetPackageExpiredTime(v string) *DescribeUserResourcesResponseBodyResourcesDesktopDurationList {
	s.PackageExpiredTime = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResourcesDesktopDurationList) SetPackageId(v string) *DescribeUserResourcesResponseBodyResourcesDesktopDurationList {
	s.PackageId = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResourcesDesktopDurationList) SetPackageStatus(v string) *DescribeUserResourcesResponseBodyResourcesDesktopDurationList {
	s.PackageStatus = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResourcesDesktopDurationList) SetPackageType(v string) *DescribeUserResourcesResponseBodyResourcesDesktopDurationList {
	s.PackageType = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResourcesDesktopDurationList) SetPackageUsedUpStrategy(v string) *DescribeUserResourcesResponseBodyResourcesDesktopDurationList {
	s.PackageUsedUpStrategy = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResourcesDesktopDurationList) SetPeriodEndTime(v string) *DescribeUserResourcesResponseBodyResourcesDesktopDurationList {
	s.PeriodEndTime = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResourcesDesktopDurationList) SetPeriodStartTime(v string) *DescribeUserResourcesResponseBodyResourcesDesktopDurationList {
	s.PeriodStartTime = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResourcesDesktopDurationList) SetPostPaidLimitFee(v float32) *DescribeUserResourcesResponseBodyResourcesDesktopDurationList {
	s.PostPaidLimitFee = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResourcesDesktopDurationList) SetTotalDuration(v int64) *DescribeUserResourcesResponseBodyResourcesDesktopDurationList {
	s.TotalDuration = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResourcesDesktopDurationList) SetUsedDuration(v int64) *DescribeUserResourcesResponseBodyResourcesDesktopDurationList {
	s.UsedDuration = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResourcesDesktopDurationList) Validate() error {
	return dara.Validate(s)
}

type DescribeUserResourcesResponseBodyResourcesDesktopTimers struct {
	// example:
	//
	// false
	AllowClientSetting *string `json:"AllowClientSetting,omitempty" xml:"AllowClientSetting,omitempty"`
	// example:
	//
	// 0 30 13 ? 	- 1-7
	CronExpression *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	// example:
	//
	// false
	Enforce *bool `json:"Enforce,omitempty" xml:"Enforce,omitempty"`
	// example:
	//
	// 2025-01-21T11:37Z
	ExecutionTime *string `json:"ExecutionTime,omitempty" xml:"ExecutionTime,omitempty"`
	// example:
	//
	// 15
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// example:
	//
	// Hibernate
	OperationType *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	// example:
	//
	// RESET_TYPE_SYSTEM
	ResetType *string `json:"ResetType,omitempty" xml:"ResetType,omitempty"`
	// example:
	//
	// TimerBoot
	TimerType *string `json:"TimerType,omitempty" xml:"TimerType,omitempty"`
}

func (s DescribeUserResourcesResponseBodyResourcesDesktopTimers) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserResourcesResponseBodyResourcesDesktopTimers) GoString() string {
	return s.String()
}

func (s *DescribeUserResourcesResponseBodyResourcesDesktopTimers) GetAllowClientSetting() *string {
	return s.AllowClientSetting
}

func (s *DescribeUserResourcesResponseBodyResourcesDesktopTimers) GetCronExpression() *string {
	return s.CronExpression
}

func (s *DescribeUserResourcesResponseBodyResourcesDesktopTimers) GetEnforce() *bool {
	return s.Enforce
}

func (s *DescribeUserResourcesResponseBodyResourcesDesktopTimers) GetExecutionTime() *string {
	return s.ExecutionTime
}

func (s *DescribeUserResourcesResponseBodyResourcesDesktopTimers) GetInterval() *int32 {
	return s.Interval
}

func (s *DescribeUserResourcesResponseBodyResourcesDesktopTimers) GetOperationType() *string {
	return s.OperationType
}

func (s *DescribeUserResourcesResponseBodyResourcesDesktopTimers) GetResetType() *string {
	return s.ResetType
}

func (s *DescribeUserResourcesResponseBodyResourcesDesktopTimers) GetTimerType() *string {
	return s.TimerType
}

func (s *DescribeUserResourcesResponseBodyResourcesDesktopTimers) SetAllowClientSetting(v string) *DescribeUserResourcesResponseBodyResourcesDesktopTimers {
	s.AllowClientSetting = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResourcesDesktopTimers) SetCronExpression(v string) *DescribeUserResourcesResponseBodyResourcesDesktopTimers {
	s.CronExpression = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResourcesDesktopTimers) SetEnforce(v bool) *DescribeUserResourcesResponseBodyResourcesDesktopTimers {
	s.Enforce = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResourcesDesktopTimers) SetExecutionTime(v string) *DescribeUserResourcesResponseBodyResourcesDesktopTimers {
	s.ExecutionTime = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResourcesDesktopTimers) SetInterval(v int32) *DescribeUserResourcesResponseBodyResourcesDesktopTimers {
	s.Interval = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResourcesDesktopTimers) SetOperationType(v string) *DescribeUserResourcesResponseBodyResourcesDesktopTimers {
	s.OperationType = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResourcesDesktopTimers) SetResetType(v string) *DescribeUserResourcesResponseBodyResourcesDesktopTimers {
	s.ResetType = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResourcesDesktopTimers) SetTimerType(v string) *DescribeUserResourcesResponseBodyResourcesDesktopTimers {
	s.TimerType = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResourcesDesktopTimers) Validate() error {
	return dara.Validate(s)
}

type DescribeUserResourcesResponseBodyResourcesFotaUpdate struct {
	// example:
	//
	// aliyun
	Channel *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	// example:
	//
	// 2.7.0-R-20250122.154826
	CurrentAppVersion *string `json:"CurrentAppVersion,omitempty" xml:"CurrentAppVersion,omitempty"`
	// example:
	//
	// false
	Force *bool `json:"Force,omitempty" xml:"Force,omitempty"`
	// example:
	//
	// 2.7.0-R-20250125.154826
	NewAppVersion *string `json:"NewAppVersion,omitempty" xml:"NewAppVersion,omitempty"`
	// example:
	//
	// 2.6.9-R-20250123.153415
	NewDcdVersion *string `json:"NewDcdVersion,omitempty" xml:"NewDcdVersion,omitempty"`
	// example:
	//
	// wuying-asp_single_session_desktop_win_x64
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
	// example:
	//
	// up
	ReleaseNote *string `json:"ReleaseNote,omitempty" xml:"ReleaseNote,omitempty"`
	// example:
	//
	// up
	ReleaseNoteEn *string `json:"ReleaseNoteEn,omitempty" xml:"ReleaseNoteEn,omitempty"`
	// example:
	//
	// up
	ReleaseNoteJp *string `json:"ReleaseNoteJp,omitempty" xml:"ReleaseNoteJp,omitempty"`
	// example:
	//
	// 474981930
	Size *string `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s DescribeUserResourcesResponseBodyResourcesFotaUpdate) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserResourcesResponseBodyResourcesFotaUpdate) GoString() string {
	return s.String()
}

func (s *DescribeUserResourcesResponseBodyResourcesFotaUpdate) GetChannel() *string {
	return s.Channel
}

func (s *DescribeUserResourcesResponseBodyResourcesFotaUpdate) GetCurrentAppVersion() *string {
	return s.CurrentAppVersion
}

func (s *DescribeUserResourcesResponseBodyResourcesFotaUpdate) GetForce() *bool {
	return s.Force
}

func (s *DescribeUserResourcesResponseBodyResourcesFotaUpdate) GetNewAppVersion() *string {
	return s.NewAppVersion
}

func (s *DescribeUserResourcesResponseBodyResourcesFotaUpdate) GetNewDcdVersion() *string {
	return s.NewDcdVersion
}

func (s *DescribeUserResourcesResponseBodyResourcesFotaUpdate) GetProject() *string {
	return s.Project
}

func (s *DescribeUserResourcesResponseBodyResourcesFotaUpdate) GetReleaseNote() *string {
	return s.ReleaseNote
}

func (s *DescribeUserResourcesResponseBodyResourcesFotaUpdate) GetReleaseNoteEn() *string {
	return s.ReleaseNoteEn
}

func (s *DescribeUserResourcesResponseBodyResourcesFotaUpdate) GetReleaseNoteJp() *string {
	return s.ReleaseNoteJp
}

func (s *DescribeUserResourcesResponseBodyResourcesFotaUpdate) GetSize() *string {
	return s.Size
}

func (s *DescribeUserResourcesResponseBodyResourcesFotaUpdate) SetChannel(v string) *DescribeUserResourcesResponseBodyResourcesFotaUpdate {
	s.Channel = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResourcesFotaUpdate) SetCurrentAppVersion(v string) *DescribeUserResourcesResponseBodyResourcesFotaUpdate {
	s.CurrentAppVersion = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResourcesFotaUpdate) SetForce(v bool) *DescribeUserResourcesResponseBodyResourcesFotaUpdate {
	s.Force = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResourcesFotaUpdate) SetNewAppVersion(v string) *DescribeUserResourcesResponseBodyResourcesFotaUpdate {
	s.NewAppVersion = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResourcesFotaUpdate) SetNewDcdVersion(v string) *DescribeUserResourcesResponseBodyResourcesFotaUpdate {
	s.NewDcdVersion = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResourcesFotaUpdate) SetProject(v string) *DescribeUserResourcesResponseBodyResourcesFotaUpdate {
	s.Project = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResourcesFotaUpdate) SetReleaseNote(v string) *DescribeUserResourcesResponseBodyResourcesFotaUpdate {
	s.ReleaseNote = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResourcesFotaUpdate) SetReleaseNoteEn(v string) *DescribeUserResourcesResponseBodyResourcesFotaUpdate {
	s.ReleaseNoteEn = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResourcesFotaUpdate) SetReleaseNoteJp(v string) *DescribeUserResourcesResponseBodyResourcesFotaUpdate {
	s.ReleaseNoteJp = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResourcesFotaUpdate) SetSize(v string) *DescribeUserResourcesResponseBodyResourcesFotaUpdate {
	s.Size = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResourcesFotaUpdate) Validate() error {
	return dara.Validate(s)
}

type DescribeUserResourcesResponseBodyResourcesOsUpdate struct {
	CheckId          *string                                                       `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	KbListString     *string                                                       `json:"KbListString,omitempty" xml:"KbListString,omitempty"`
	PackageCount     *int32                                                        `json:"PackageCount,omitempty" xml:"PackageCount,omitempty"`
	Packages         []*DescribeUserResourcesResponseBodyResourcesOsUpdatePackages `json:"Packages,omitempty" xml:"Packages,omitempty" type:"Repeated"`
	UpdateCatalogUrl *string                                                       `json:"UpdateCatalogUrl,omitempty" xml:"UpdateCatalogUrl,omitempty"`
}

func (s DescribeUserResourcesResponseBodyResourcesOsUpdate) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserResourcesResponseBodyResourcesOsUpdate) GoString() string {
	return s.String()
}

func (s *DescribeUserResourcesResponseBodyResourcesOsUpdate) GetCheckId() *string {
	return s.CheckId
}

func (s *DescribeUserResourcesResponseBodyResourcesOsUpdate) GetKbListString() *string {
	return s.KbListString
}

func (s *DescribeUserResourcesResponseBodyResourcesOsUpdate) GetPackageCount() *int32 {
	return s.PackageCount
}

func (s *DescribeUserResourcesResponseBodyResourcesOsUpdate) GetPackages() []*DescribeUserResourcesResponseBodyResourcesOsUpdatePackages {
	return s.Packages
}

func (s *DescribeUserResourcesResponseBodyResourcesOsUpdate) GetUpdateCatalogUrl() *string {
	return s.UpdateCatalogUrl
}

func (s *DescribeUserResourcesResponseBodyResourcesOsUpdate) SetCheckId(v string) *DescribeUserResourcesResponseBodyResourcesOsUpdate {
	s.CheckId = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResourcesOsUpdate) SetKbListString(v string) *DescribeUserResourcesResponseBodyResourcesOsUpdate {
	s.KbListString = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResourcesOsUpdate) SetPackageCount(v int32) *DescribeUserResourcesResponseBodyResourcesOsUpdate {
	s.PackageCount = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResourcesOsUpdate) SetPackages(v []*DescribeUserResourcesResponseBodyResourcesOsUpdatePackages) *DescribeUserResourcesResponseBodyResourcesOsUpdate {
	s.Packages = v
	return s
}

func (s *DescribeUserResourcesResponseBodyResourcesOsUpdate) SetUpdateCatalogUrl(v string) *DescribeUserResourcesResponseBodyResourcesOsUpdate {
	s.UpdateCatalogUrl = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResourcesOsUpdate) Validate() error {
	if s.Packages != nil {
		for _, item := range s.Packages {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeUserResourcesResponseBodyResourcesOsUpdatePackages struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Kb          *string `json:"Kb,omitempty" xml:"Kb,omitempty"`
	Title       *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s DescribeUserResourcesResponseBodyResourcesOsUpdatePackages) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserResourcesResponseBodyResourcesOsUpdatePackages) GoString() string {
	return s.String()
}

func (s *DescribeUserResourcesResponseBodyResourcesOsUpdatePackages) GetDescription() *string {
	return s.Description
}

func (s *DescribeUserResourcesResponseBodyResourcesOsUpdatePackages) GetKb() *string {
	return s.Kb
}

func (s *DescribeUserResourcesResponseBodyResourcesOsUpdatePackages) GetTitle() *string {
	return s.Title
}

func (s *DescribeUserResourcesResponseBodyResourcesOsUpdatePackages) SetDescription(v string) *DescribeUserResourcesResponseBodyResourcesOsUpdatePackages {
	s.Description = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResourcesOsUpdatePackages) SetKb(v string) *DescribeUserResourcesResponseBodyResourcesOsUpdatePackages {
	s.Kb = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResourcesOsUpdatePackages) SetTitle(v string) *DescribeUserResourcesResponseBodyResourcesOsUpdatePackages {
	s.Title = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResourcesOsUpdatePackages) Validate() error {
	return dara.Validate(s)
}

type DescribeUserResourcesResponseBodyResourcesSessions struct {
	NickName *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
	// example:
	//
	// 2025-01-22T11:03:36Z
	ResourceSessionStartTime *string `json:"ResourceSessionStartTime,omitempty" xml:"ResourceSessionStartTime,omitempty"`
	// example:
	//
	// user001
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// benchmark_test@test.shenzhen
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s DescribeUserResourcesResponseBodyResourcesSessions) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserResourcesResponseBodyResourcesSessions) GoString() string {
	return s.String()
}

func (s *DescribeUserResourcesResponseBodyResourcesSessions) GetNickName() *string {
	return s.NickName
}

func (s *DescribeUserResourcesResponseBodyResourcesSessions) GetResourceSessionStartTime() *string {
	return s.ResourceSessionStartTime
}

func (s *DescribeUserResourcesResponseBodyResourcesSessions) GetUserId() *string {
	return s.UserId
}

func (s *DescribeUserResourcesResponseBodyResourcesSessions) GetUserPrincipalName() *string {
	return s.UserPrincipalName
}

func (s *DescribeUserResourcesResponseBodyResourcesSessions) SetNickName(v string) *DescribeUserResourcesResponseBodyResourcesSessions {
	s.NickName = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResourcesSessions) SetResourceSessionStartTime(v string) *DescribeUserResourcesResponseBodyResourcesSessions {
	s.ResourceSessionStartTime = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResourcesSessions) SetUserId(v string) *DescribeUserResourcesResponseBodyResourcesSessions {
	s.UserId = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResourcesSessions) SetUserPrincipalName(v string) *DescribeUserResourcesResponseBodyResourcesSessions {
	s.UserPrincipalName = &v
	return s
}

func (s *DescribeUserResourcesResponseBodyResourcesSessions) Validate() error {
	return dara.Validate(s)
}
