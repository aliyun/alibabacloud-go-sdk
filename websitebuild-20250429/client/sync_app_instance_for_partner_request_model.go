// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncAppInstanceForPartnerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppInstance(v *SyncAppInstanceForPartnerRequestAppInstance) *SyncAppInstanceForPartnerRequest
	GetAppInstance() *SyncAppInstanceForPartnerRequestAppInstance
	SetEventType(v string) *SyncAppInstanceForPartnerRequest
	GetEventType() *string
	SetOperator(v string) *SyncAppInstanceForPartnerRequest
	GetOperator() *string
	SetSourceBizId(v string) *SyncAppInstanceForPartnerRequest
	GetSourceBizId() *string
	SetSourceType(v string) *SyncAppInstanceForPartnerRequest
	GetSourceType() *string
}

type SyncAppInstanceForPartnerRequest struct {
	// Application instance object data
	AppInstance *SyncAppInstanceForPartnerRequestAppInstance `json:"AppInstance,omitempty" xml:"AppInstance,omitempty" type:"Struct"`
	// Type of system event. CREATE, UPDATE, COMPLETE
	//
	// example:
	//
	// CREATE
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// Operator ID
	//
	// example:
	//
	// system
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// Source business ID.
	//
	// example:
	//
	// 31104757
	SourceBizId *string `json:"SourceBizId,omitempty" xml:"SourceBizId,omitempty"`
	// Source: MARKET_CLOUD_DREAM
	//
	// example:
	//
	// MARKET_CLOUD_DREAM
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s SyncAppInstanceForPartnerRequest) String() string {
	return dara.Prettify(s)
}

func (s SyncAppInstanceForPartnerRequest) GoString() string {
	return s.String()
}

func (s *SyncAppInstanceForPartnerRequest) GetAppInstance() *SyncAppInstanceForPartnerRequestAppInstance {
	return s.AppInstance
}

func (s *SyncAppInstanceForPartnerRequest) GetEventType() *string {
	return s.EventType
}

func (s *SyncAppInstanceForPartnerRequest) GetOperator() *string {
	return s.Operator
}

func (s *SyncAppInstanceForPartnerRequest) GetSourceBizId() *string {
	return s.SourceBizId
}

func (s *SyncAppInstanceForPartnerRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *SyncAppInstanceForPartnerRequest) SetAppInstance(v *SyncAppInstanceForPartnerRequestAppInstance) *SyncAppInstanceForPartnerRequest {
	s.AppInstance = v
	return s
}

func (s *SyncAppInstanceForPartnerRequest) SetEventType(v string) *SyncAppInstanceForPartnerRequest {
	s.EventType = &v
	return s
}

func (s *SyncAppInstanceForPartnerRequest) SetOperator(v string) *SyncAppInstanceForPartnerRequest {
	s.Operator = &v
	return s
}

func (s *SyncAppInstanceForPartnerRequest) SetSourceBizId(v string) *SyncAppInstanceForPartnerRequest {
	s.SourceBizId = &v
	return s
}

func (s *SyncAppInstanceForPartnerRequest) SetSourceType(v string) *SyncAppInstanceForPartnerRequest {
	s.SourceType = &v
	return s
}

func (s *SyncAppInstanceForPartnerRequest) Validate() error {
	if s.AppInstance != nil {
		if err := s.AppInstance.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SyncAppInstanceForPartnerRequestAppInstance struct {
	// Application Type: WEBSITE
	//
	// example:
	//
	// WEBSITE
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// Website business ID
	//
	// example:
	//
	// WD20250711094503000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// Logical deletion
	//
	// example:
	//
	// fase
	Deleted *string `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	// Domain name
	//
	// example:
	//
	// alliveout.xntv.tv
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// Event end time (UNIX timestamp).
	//
	// example:
	//
	// 2024-08-23T02:14:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Deletion time
	//
	// example:
	//
	// 2025-01-01 00:00:00
	GmtDelete *string `json:"GmtDelete,omitempty" xml:"GmtDelete,omitempty"`
	// Vulnerability published UNIX timestamp, in milliseconds.
	//
	// example:
	//
	// 2025-01-01 00:00:00
	GmtPublish *string `json:"GmtPublish,omitempty" xml:"GmtPublish,omitempty"`
	// Application icon URL.
	//
	// example:
	//
	// icon/WS20250626112715000001/thumbnail.jpg
	IconUrl *string `json:"IconUrl,omitempty" xml:"IconUrl,omitempty"`
	// Website name
	//
	// example:
	//
	// AI建站
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Website configuration information
	//
	// example:
	//
	// 网站配置信息
	Profile *SyncAppInstanceForPartnerRequestAppInstanceProfile `json:"Profile,omitempty" xml:"Profile,omitempty" type:"Struct"`
	// siteId
	//
	// example:
	//
	// xxxx.scd.wezhan.cn
	SiteHost *string `json:"SiteHost,omitempty" xml:"SiteHost,omitempty"`
	// Website SiteID
	//
	// example:
	//
	// 31104757
	Slug *string `json:"Slug,omitempty" xml:"Slug,omitempty"`
	// Query start time. If no start and end times are provided, all historical deployment records of the instance are queried.
	//
	// example:
	//
	// 2025-07-15T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// Instance running status.
	//
	// - NotRun: Not running
	//
	// - Running: Running
	//
	// - WaitTime: Waiting for TriggerTime
	//
	// - CheckingCondition: Checking branch conditions
	//
	// - WaitResource: Waiting for resources
	//
	// - Failure: Execution failed
	//
	// - Success: Execution succeeded
	//
	// - Checking: Sent to Data Quality check
	//
	// example:
	//
	// {\\"Phase\\": \\"Running\\", \\"SlotNum\\": 1, \\"UsedCapacity\\": \\"500.0Gi\\"}
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Thumbnail URL.
	//
	// example:
	//
	// thumbnails/WS20250626112715000001/thumbnail.jpg
	ThumbnailUrl *string `json:"ThumbnailUrl,omitempty" xml:"ThumbnailUrl,omitempty"`
	// 123123123131232
	//
	// example:
	//
	// 12313213131
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s SyncAppInstanceForPartnerRequestAppInstance) String() string {
	return dara.Prettify(s)
}

func (s SyncAppInstanceForPartnerRequestAppInstance) GoString() string {
	return s.String()
}

func (s *SyncAppInstanceForPartnerRequestAppInstance) GetAppType() *string {
	return s.AppType
}

func (s *SyncAppInstanceForPartnerRequestAppInstance) GetBizId() *string {
	return s.BizId
}

func (s *SyncAppInstanceForPartnerRequestAppInstance) GetDeleted() *string {
	return s.Deleted
}

func (s *SyncAppInstanceForPartnerRequestAppInstance) GetDomain() *string {
	return s.Domain
}

func (s *SyncAppInstanceForPartnerRequestAppInstance) GetEndTime() *string {
	return s.EndTime
}

func (s *SyncAppInstanceForPartnerRequestAppInstance) GetGmtDelete() *string {
	return s.GmtDelete
}

func (s *SyncAppInstanceForPartnerRequestAppInstance) GetGmtPublish() *string {
	return s.GmtPublish
}

func (s *SyncAppInstanceForPartnerRequestAppInstance) GetIconUrl() *string {
	return s.IconUrl
}

func (s *SyncAppInstanceForPartnerRequestAppInstance) GetName() *string {
	return s.Name
}

func (s *SyncAppInstanceForPartnerRequestAppInstance) GetProfile() *SyncAppInstanceForPartnerRequestAppInstanceProfile {
	return s.Profile
}

func (s *SyncAppInstanceForPartnerRequestAppInstance) GetSiteHost() *string {
	return s.SiteHost
}

func (s *SyncAppInstanceForPartnerRequestAppInstance) GetSlug() *string {
	return s.Slug
}

func (s *SyncAppInstanceForPartnerRequestAppInstance) GetStartTime() *string {
	return s.StartTime
}

func (s *SyncAppInstanceForPartnerRequestAppInstance) GetStatus() *string {
	return s.Status
}

func (s *SyncAppInstanceForPartnerRequestAppInstance) GetThumbnailUrl() *string {
	return s.ThumbnailUrl
}

func (s *SyncAppInstanceForPartnerRequestAppInstance) GetUserId() *string {
	return s.UserId
}

func (s *SyncAppInstanceForPartnerRequestAppInstance) SetAppType(v string) *SyncAppInstanceForPartnerRequestAppInstance {
	s.AppType = &v
	return s
}

func (s *SyncAppInstanceForPartnerRequestAppInstance) SetBizId(v string) *SyncAppInstanceForPartnerRequestAppInstance {
	s.BizId = &v
	return s
}

func (s *SyncAppInstanceForPartnerRequestAppInstance) SetDeleted(v string) *SyncAppInstanceForPartnerRequestAppInstance {
	s.Deleted = &v
	return s
}

func (s *SyncAppInstanceForPartnerRequestAppInstance) SetDomain(v string) *SyncAppInstanceForPartnerRequestAppInstance {
	s.Domain = &v
	return s
}

func (s *SyncAppInstanceForPartnerRequestAppInstance) SetEndTime(v string) *SyncAppInstanceForPartnerRequestAppInstance {
	s.EndTime = &v
	return s
}

func (s *SyncAppInstanceForPartnerRequestAppInstance) SetGmtDelete(v string) *SyncAppInstanceForPartnerRequestAppInstance {
	s.GmtDelete = &v
	return s
}

func (s *SyncAppInstanceForPartnerRequestAppInstance) SetGmtPublish(v string) *SyncAppInstanceForPartnerRequestAppInstance {
	s.GmtPublish = &v
	return s
}

func (s *SyncAppInstanceForPartnerRequestAppInstance) SetIconUrl(v string) *SyncAppInstanceForPartnerRequestAppInstance {
	s.IconUrl = &v
	return s
}

func (s *SyncAppInstanceForPartnerRequestAppInstance) SetName(v string) *SyncAppInstanceForPartnerRequestAppInstance {
	s.Name = &v
	return s
}

func (s *SyncAppInstanceForPartnerRequestAppInstance) SetProfile(v *SyncAppInstanceForPartnerRequestAppInstanceProfile) *SyncAppInstanceForPartnerRequestAppInstance {
	s.Profile = v
	return s
}

func (s *SyncAppInstanceForPartnerRequestAppInstance) SetSiteHost(v string) *SyncAppInstanceForPartnerRequestAppInstance {
	s.SiteHost = &v
	return s
}

func (s *SyncAppInstanceForPartnerRequestAppInstance) SetSlug(v string) *SyncAppInstanceForPartnerRequestAppInstance {
	s.Slug = &v
	return s
}

func (s *SyncAppInstanceForPartnerRequestAppInstance) SetStartTime(v string) *SyncAppInstanceForPartnerRequestAppInstance {
	s.StartTime = &v
	return s
}

func (s *SyncAppInstanceForPartnerRequestAppInstance) SetStatus(v string) *SyncAppInstanceForPartnerRequestAppInstance {
	s.Status = &v
	return s
}

func (s *SyncAppInstanceForPartnerRequestAppInstance) SetThumbnailUrl(v string) *SyncAppInstanceForPartnerRequestAppInstance {
	s.ThumbnailUrl = &v
	return s
}

func (s *SyncAppInstanceForPartnerRequestAppInstance) SetUserId(v string) *SyncAppInstanceForPartnerRequestAppInstance {
	s.UserId = &v
	return s
}

func (s *SyncAppInstanceForPartnerRequestAppInstance) Validate() error {
	if s.Profile != nil {
		if err := s.Profile.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SyncAppInstanceForPartnerRequestAppInstanceProfile struct {
	// Deployment region
	//
	// example:
	//
	// ChineseMainland
	DeployArea *string `json:"DeployArea,omitempty" xml:"DeployArea,omitempty"`
	// Lingxiao instance ID
	//
	// example:
	//
	// 12313213
	LxInstanceId *string `json:"LxInstanceId,omitempty" xml:"LxInstanceId,omitempty"`
	// Order ID
	//
	// example:
	//
	// 222217928591
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// Version
	//
	// example:
	//
	// Basic_Edition
	SiteVersion *string `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
	// Template ID
	//
	// example:
	//
	// DC4D30B7BADDAFE9928A6C36416A2A4C
	TemplateEtag *string `json:"TemplateEtag,omitempty" xml:"TemplateEtag,omitempty"`
	// Model template ID
	//
	// example:
	//
	// ST20211231160247sYG4
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s SyncAppInstanceForPartnerRequestAppInstanceProfile) String() string {
	return dara.Prettify(s)
}

func (s SyncAppInstanceForPartnerRequestAppInstanceProfile) GoString() string {
	return s.String()
}

func (s *SyncAppInstanceForPartnerRequestAppInstanceProfile) GetDeployArea() *string {
	return s.DeployArea
}

func (s *SyncAppInstanceForPartnerRequestAppInstanceProfile) GetLxInstanceId() *string {
	return s.LxInstanceId
}

func (s *SyncAppInstanceForPartnerRequestAppInstanceProfile) GetOrderId() *string {
	return s.OrderId
}

func (s *SyncAppInstanceForPartnerRequestAppInstanceProfile) GetSiteVersion() *string {
	return s.SiteVersion
}

func (s *SyncAppInstanceForPartnerRequestAppInstanceProfile) GetTemplateEtag() *string {
	return s.TemplateEtag
}

func (s *SyncAppInstanceForPartnerRequestAppInstanceProfile) GetTemplateId() *string {
	return s.TemplateId
}

func (s *SyncAppInstanceForPartnerRequestAppInstanceProfile) SetDeployArea(v string) *SyncAppInstanceForPartnerRequestAppInstanceProfile {
	s.DeployArea = &v
	return s
}

func (s *SyncAppInstanceForPartnerRequestAppInstanceProfile) SetLxInstanceId(v string) *SyncAppInstanceForPartnerRequestAppInstanceProfile {
	s.LxInstanceId = &v
	return s
}

func (s *SyncAppInstanceForPartnerRequestAppInstanceProfile) SetOrderId(v string) *SyncAppInstanceForPartnerRequestAppInstanceProfile {
	s.OrderId = &v
	return s
}

func (s *SyncAppInstanceForPartnerRequestAppInstanceProfile) SetSiteVersion(v string) *SyncAppInstanceForPartnerRequestAppInstanceProfile {
	s.SiteVersion = &v
	return s
}

func (s *SyncAppInstanceForPartnerRequestAppInstanceProfile) SetTemplateEtag(v string) *SyncAppInstanceForPartnerRequestAppInstanceProfile {
	s.TemplateEtag = &v
	return s
}

func (s *SyncAppInstanceForPartnerRequestAppInstanceProfile) SetTemplateId(v string) *SyncAppInstanceForPartnerRequestAppInstanceProfile {
	s.TemplateId = &v
	return s
}

func (s *SyncAppInstanceForPartnerRequestAppInstanceProfile) Validate() error {
	return dara.Validate(s)
}
