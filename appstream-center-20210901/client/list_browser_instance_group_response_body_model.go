// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBrowserInstanceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBrowserInstanceGroupModels(v []*ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModels) *ListBrowserInstanceGroupResponseBody
	GetBrowserInstanceGroupModels() []*ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModels
	SetPageNumber(v int32) *ListBrowserInstanceGroupResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListBrowserInstanceGroupResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListBrowserInstanceGroupResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListBrowserInstanceGroupResponseBody
	GetTotalCount() *int32
}

type ListBrowserInstanceGroupResponseBody struct {
	// The list of browser groups on the current page.
	BrowserInstanceGroupModels []*ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModels `json:"BrowserInstanceGroupModels,omitempty" xml:"BrowserInstanceGroupModels,omitempty" type:"Repeated"`
	// The current page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID, which is used for troubleshooting.
	//
	// example:
	//
	// 01A0C2ED-95F2-1A37-9FC6-4A395179****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of browser groups that match the filter conditions.
	//
	// example:
	//
	// 5
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListBrowserInstanceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListBrowserInstanceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListBrowserInstanceGroupResponseBody) GetBrowserInstanceGroupModels() []*ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModels {
	return s.BrowserInstanceGroupModels
}

func (s *ListBrowserInstanceGroupResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListBrowserInstanceGroupResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListBrowserInstanceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListBrowserInstanceGroupResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListBrowserInstanceGroupResponseBody) SetBrowserInstanceGroupModels(v []*ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModels) *ListBrowserInstanceGroupResponseBody {
	s.BrowserInstanceGroupModels = v
	return s
}

func (s *ListBrowserInstanceGroupResponseBody) SetPageNumber(v int32) *ListBrowserInstanceGroupResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListBrowserInstanceGroupResponseBody) SetPageSize(v int32) *ListBrowserInstanceGroupResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListBrowserInstanceGroupResponseBody) SetRequestId(v string) *ListBrowserInstanceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBrowserInstanceGroupResponseBody) SetTotalCount(v int32) *ListBrowserInstanceGroupResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListBrowserInstanceGroupResponseBody) Validate() error {
	if s.BrowserInstanceGroupModels != nil {
		for _, item := range s.BrowserInstanceGroupModels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModels struct {
	// Indicates whether authorization and deauthorization notification emails are enabled. `true` indicates enabled. `false` indicates disabled.
	//
	// example:
	//
	// true
	AuthNotificationEnabled *bool `json:"AuthNotificationEnabled,omitempty" xml:"AuthNotificationEnabled,omitempty"`
	// The authorized user statistics of the browser group.
	AuthorizedUserInfo *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsAuthorizedUserInfo `json:"AuthorizedUserInfo,omitempty" xml:"AuthorizedUserInfo,omitempty" type:"Struct"`
	// The business region where the browser group is located.
	//
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// The current browser configuration.
	BrowserConfig *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsBrowserConfig `json:"BrowserConfig,omitempty" xml:"BrowserConfig,omitempty" type:"Struct"`
	// The ID of the cloud browser group.
	//
	// example:
	//
	// big-0c7loey7fzjq****
	BrowserInstanceGroupId *string `json:"BrowserInstanceGroupId,omitempty" xml:"BrowserInstanceGroupId,omitempty"`
	// The name of the cloud browser group.
	//
	// example:
	//
	// OfficeBrowser
	BrowserInstanceGroupName *string `json:"BrowserInstanceGroupName,omitempty" xml:"BrowserInstanceGroupName,omitempty"`
	// The ID of the browser group set to which the browser group belongs.
	//
	// example:
	//
	// set-3jm9d0abc00example
	BrowserInstanceGroupSetId *string `json:"BrowserInstanceGroupSetId,omitempty" xml:"BrowserInstanceGroupSetId,omitempty"`
	// The billing type. In MAU scenarios, `PostPaid` is returned, which indicates pay-as-you-go.
	//
	// example:
	//
	// PostPaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The default access URL of the browser group. Use the URL returned by the API for access. The resource identifiers in the example must be replaced.
	//
	// example:
	//
	// https://wuying.aliyun.com/integration?appId=browser&appInstanceGroupId=big-0c7loey7fzjq****
	DefaultAccessUrl *string `json:"DefaultAccessUrl,omitempty" xml:"DefaultAccessUrl,omitempty"`
	// The plan duration information. In MAU scenarios, plan duration does not apply, and an empty object may be returned.
	Duration *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsDuration `json:"Duration,omitempty" xml:"Duration,omitempty" type:"Struct"`
	// The expiration time of the browser group. Not applicable in MAU scenarios. This field is not returned.
	//
	// example:
	//
	// -
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The creation time of the browser group.
	//
	// The time is in RFC 3339 format: `yyyy-MM-dd\\"T\\"HH:mm:ss.SSSXXX`, which includes milliseconds and a time zone offset. The actual POP response uses the UTC offset `+00:00`.
	//
	// example:
	//
	// 2026-09-21T07:00:39.000+00:00
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The homepage URL of the browser group.
	//
	// example:
	//
	// https://www.aliyun.com
	Homepage *string `json:"Homepage,omitempty" xml:"Homepage,omitempty"`
	// The image ID used by the browser group.
	//
	// example:
	//
	// imgc-070qhs8oeju4****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The instance type used by the browser group.
	//
	// example:
	//
	// appstreaming.general.basic
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The workspace network and website access restriction configuration.
	Network *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsNetwork `json:"Network,omitempty" xml:"Network,omitempty" type:"Struct"`
	// The operating system type of the browser group. The current MAU product scenario uses Windows.
	//
	// example:
	//
	// Windows
	OsType *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	// The policy configuration returned for the browser group. Policy fields are used to view existing settings and do not indicate that all corresponding creation parameters are configurable.
	Policy *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsPolicy `json:"Policy,omitempty" xml:"Policy,omitempty" type:"Struct"`
	// The browser group status.
	//
	// - `DEPLOYING`: Being deployed.
	//
	// - `PUBLISHED`: Deployed.
	//
	// - `FAILED`: Deployment failed.
	//
	// - `EXPIRED`: Expired.
	//
	// - `CEASED`: Suspended due to overdue payment.
	//
	// - `MAINTAINING`: Being updated.
	//
	// - `MAINTAIN_FAILED`: Update failed.
	//
	// - `DELETING`: Being deleted.
	//
	// - `UNAVAILABLE`: Unavailable.
	//
	// example:
	//
	// PUBLISHED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The sub-payment type. In MAU scenarios, the actual returned value is `mau`, which indicates billing by monthly active users.
	//
	// example:
	//
	// mau
	SubPayType *string `json:"SubPayType,omitempty" xml:"SubPayType,omitempty"`
	// Indicates whether mixed authorization of users and user groups is supported. `true` indicates supported, and `false` indicates not supported. Evaluate this value based on the current authorization mode.
	//
	// example:
	//
	// false
	SupportUserGroupMixedAuth *bool `json:"SupportUserGroupMixedAuth,omitempty" xml:"SupportUserGroupMixedAuth,omitempty"`
	// The list of resource tags.
	Tags []*ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The version of the browser. Valid values:
	//
	// - `Basic`: Basic Edition.
	//
	// - `Pro`: Premium Edition.
	//
	// In MAU scenarios, the value is `Pro`.
	//
	// example:
	//
	// Pro
	Tier *string `json:"Tier,omitempty" xml:"Tier,omitempty"`
	// The session timer configurations currently returned. This is used to view the effective settings and does not indicate that the create operation supports setting this parameter.
	Timers []*ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsTimers `json:"Timers,omitempty" xml:"Timers,omitempty" type:"Repeated"`
	// The current authorization mode. Valid values:
	//
	// - `Mixed`: Mixed authorization of users and user groups.
	//
	// - `User`: User authorization.
	//
	// - `UserGroup`: User group authorization.
	//
	// example:
	//
	// Mixed
	UserGroupAuthMode *string `json:"UserGroupAuthMode,omitempty" xml:"UserGroupAuthMode,omitempty"`
	// The user quota information.
	UserLimit *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsUserLimit `json:"UserLimit,omitempty" xml:"UserLimit,omitempty" type:"Struct"`
}

func (s ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModels) String() string {
	return dara.Prettify(s)
}

func (s ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModels) GoString() string {
	return s.String()
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModels) GetAuthNotificationEnabled() *bool {
	return s.AuthNotificationEnabled
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModels) GetAuthorizedUserInfo() *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsAuthorizedUserInfo {
	return s.AuthorizedUserInfo
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModels) GetBizRegionId() *string {
	return s.BizRegionId
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModels) GetBrowserConfig() *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsBrowserConfig {
	return s.BrowserConfig
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModels) GetBrowserInstanceGroupId() *string {
	return s.BrowserInstanceGroupId
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModels) GetBrowserInstanceGroupName() *string {
	return s.BrowserInstanceGroupName
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModels) GetBrowserInstanceGroupSetId() *string {
	return s.BrowserInstanceGroupSetId
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModels) GetChargeType() *string {
	return s.ChargeType
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModels) GetDefaultAccessUrl() *string {
	return s.DefaultAccessUrl
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModels) GetDuration() *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsDuration {
	return s.Duration
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModels) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModels) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModels) GetHomepage() *string {
	return s.Homepage
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModels) GetImageId() *string {
	return s.ImageId
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModels) GetInstanceType() *string {
	return s.InstanceType
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModels) GetNetwork() *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsNetwork {
	return s.Network
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModels) GetOsType() *string {
	return s.OsType
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModels) GetPolicy() *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsPolicy {
	return s.Policy
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModels) GetStatus() *string {
	return s.Status
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModels) GetSubPayType() *string {
	return s.SubPayType
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModels) GetSupportUserGroupMixedAuth() *bool {
	return s.SupportUserGroupMixedAuth
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModels) GetTags() []*ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsTags {
	return s.Tags
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModels) GetTier() *string {
	return s.Tier
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModels) GetTimers() []*ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsTimers {
	return s.Timers
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModels) GetUserGroupAuthMode() *string {
	return s.UserGroupAuthMode
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModels) GetUserLimit() *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsUserLimit {
	return s.UserLimit
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModels) SetAuthNotificationEnabled(v bool) *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModels {
	s.AuthNotificationEnabled = &v
	return s
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModels) SetAuthorizedUserInfo(v *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsAuthorizedUserInfo) *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModels {
	s.AuthorizedUserInfo = v
	return s
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModels) SetBizRegionId(v string) *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModels {
	s.BizRegionId = &v
	return s
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModels) SetBrowserConfig(v *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsBrowserConfig) *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModels {
	s.BrowserConfig = v
	return s
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModels) SetBrowserInstanceGroupId(v string) *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModels {
	s.BrowserInstanceGroupId = &v
	return s
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModels) SetBrowserInstanceGroupName(v string) *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModels {
	s.BrowserInstanceGroupName = &v
	return s
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModels) SetBrowserInstanceGroupSetId(v string) *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModels {
	s.BrowserInstanceGroupSetId = &v
	return s
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModels) SetChargeType(v string) *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModels {
	s.ChargeType = &v
	return s
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModels) SetDefaultAccessUrl(v string) *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModels {
	s.DefaultAccessUrl = &v
	return s
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModels) SetDuration(v *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsDuration) *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModels {
	s.Duration = v
	return s
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModels) SetExpiredTime(v string) *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModels {
	s.ExpiredTime = &v
	return s
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModels) SetGmtCreate(v string) *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModels {
	s.GmtCreate = &v
	return s
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModels) SetHomepage(v string) *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModels {
	s.Homepage = &v
	return s
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModels) SetImageId(v string) *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModels {
	s.ImageId = &v
	return s
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModels) SetInstanceType(v string) *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModels {
	s.InstanceType = &v
	return s
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModels) SetNetwork(v *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsNetwork) *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModels {
	s.Network = v
	return s
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModels) SetOsType(v string) *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModels {
	s.OsType = &v
	return s
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModels) SetPolicy(v *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsPolicy) *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModels {
	s.Policy = v
	return s
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModels) SetStatus(v string) *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModels {
	s.Status = &v
	return s
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModels) SetSubPayType(v string) *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModels {
	s.SubPayType = &v
	return s
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModels) SetSupportUserGroupMixedAuth(v bool) *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModels {
	s.SupportUserGroupMixedAuth = &v
	return s
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModels) SetTags(v []*ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsTags) *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModels {
	s.Tags = v
	return s
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModels) SetTier(v string) *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModels {
	s.Tier = &v
	return s
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModels) SetTimers(v []*ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsTimers) *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModels {
	s.Timers = v
	return s
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModels) SetUserGroupAuthMode(v string) *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModels {
	s.UserGroupAuthMode = &v
	return s
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModels) SetUserLimit(v *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsUserLimit) *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModels {
	s.UserLimit = v
	return s
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModels) Validate() error {
	if s.AuthorizedUserInfo != nil {
		if err := s.AuthorizedUserInfo.Validate(); err != nil {
			return err
		}
	}
	if s.BrowserConfig != nil {
		if err := s.BrowserConfig.Validate(); err != nil {
			return err
		}
	}
	if s.Duration != nil {
		if err := s.Duration.Validate(); err != nil {
			return err
		}
	}
	if s.Network != nil {
		if err := s.Network.Validate(); err != nil {
			return err
		}
	}
	if s.Policy != nil {
		if err := s.Policy.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Timers != nil {
		for _, item := range s.Timers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.UserLimit != nil {
		if err := s.UserLimit.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsAuthorizedUserInfo struct {
	// The total number of authorized users.
	//
	// example:
	//
	// 0
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The total number of authorized user groups.
	//
	// example:
	//
	// 0
	TotalUserGroupCount *int32 `json:"TotalUserGroupCount,omitempty" xml:"TotalUserGroupCount,omitempty"`
}

func (s ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsAuthorizedUserInfo) String() string {
	return dara.Prettify(s)
}

func (s ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsAuthorizedUserInfo) GoString() string {
	return s.String()
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsAuthorizedUserInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsAuthorizedUserInfo) GetTotalUserGroupCount() *int32 {
	return s.TotalUserGroupCount
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsAuthorizedUserInfo) SetTotalCount(v int32) *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsAuthorizedUserInfo {
	s.TotalCount = &v
	return s
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsAuthorizedUserInfo) SetTotalUserGroupCount(v int32) *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsAuthorizedUserInfo {
	s.TotalUserGroupCount = &v
	return s
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsAuthorizedUserInfo) Validate() error {
	return dara.Validate(s)
}

type ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsBrowserConfig struct {
	// The browser startup parameters. For example, `--incognito` indicates an incognito window.
	//
	// example:
	//
	// --incognito
	BrowserParam *string `json:"BrowserParam,omitempty" xml:"BrowserParam,omitempty"`
	// The cookie synchronization configuration. The string `true` indicates enabled. `false` indicates disabled.
	//
	// example:
	//
	// true
	CookiesSync *string `json:"CookiesSync,omitempty" xml:"CookiesSync,omitempty"`
	// The homepage URL that opens when the browser starts.
	//
	// example:
	//
	// https://www.aliyun.com
	Homepage *string `json:"Homepage,omitempty" xml:"Homepage,omitempty"`
}

func (s ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsBrowserConfig) String() string {
	return dara.Prettify(s)
}

func (s ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsBrowserConfig) GoString() string {
	return s.String()
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsBrowserConfig) GetBrowserParam() *string {
	return s.BrowserParam
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsBrowserConfig) GetCookiesSync() *string {
	return s.CookiesSync
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsBrowserConfig) GetHomepage() *string {
	return s.Homepage
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsBrowserConfig) SetBrowserParam(v string) *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsBrowserConfig {
	s.BrowserParam = &v
	return s
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsBrowserConfig) SetCookiesSync(v string) *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsBrowserConfig {
	s.CookiesSync = &v
	return s
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsBrowserConfig) SetHomepage(v string) *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsBrowserConfig {
	s.Homepage = &v
	return s
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsBrowserConfig) Validate() error {
	return dara.Validate(s)
}

type ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsDuration struct {
	// The current payment stage of the plan. Not applicable in MAU scenarios.
	//
	// example:
	//
	// -
	CurrentPayStage *string `json:"CurrentPayStage,omitempty" xml:"CurrentPayStage,omitempty"`
	// The end time of the plan period. Not applicable in MAU scenarios. This field is not returned.
	//
	// example:
	//
	// -
	PeriodEndTime *string `json:"PeriodEndTime,omitempty" xml:"PeriodEndTime,omitempty"`
	// The start time of the plan period. Not applicable in MAU scenarios. This field is not returned.
	//
	// example:
	//
	// -
	PeriodStartTime *string `json:"PeriodStartTime,omitempty" xml:"PeriodStartTime,omitempty"`
	// The total duration of the plan, in seconds. Not applicable in MAU scenarios.
	//
	// example:
	//
	// -
	TotalDuration *int32 `json:"TotalDuration,omitempty" xml:"TotalDuration,omitempty"`
	// The used duration of the plan, in seconds. Not applicable in MAU scenarios.
	//
	// example:
	//
	// -
	UsedDuration *int32 `json:"UsedDuration,omitempty" xml:"UsedDuration,omitempty"`
}

func (s ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsDuration) String() string {
	return dara.Prettify(s)
}

func (s ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsDuration) GoString() string {
	return s.String()
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsDuration) GetCurrentPayStage() *string {
	return s.CurrentPayStage
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsDuration) GetPeriodEndTime() *string {
	return s.PeriodEndTime
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsDuration) GetPeriodStartTime() *string {
	return s.PeriodStartTime
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsDuration) GetTotalDuration() *int32 {
	return s.TotalDuration
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsDuration) GetUsedDuration() *int32 {
	return s.UsedDuration
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsDuration) SetCurrentPayStage(v string) *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsDuration {
	s.CurrentPayStage = &v
	return s
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsDuration) SetPeriodEndTime(v string) *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsDuration {
	s.PeriodEndTime = &v
	return s
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsDuration) SetPeriodStartTime(v string) *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsDuration {
	s.PeriodStartTime = &v
	return s
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsDuration) SetTotalDuration(v int32) *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsDuration {
	s.TotalDuration = &v
	return s
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsDuration) SetUsedDuration(v int32) *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsDuration {
	s.UsedDuration = &v
	return s
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsDuration) Validate() error {
	return dara.Validate(s)
}

type ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsNetwork struct {
	// The website access restriction mode.
	//
	// - `ALLOW_ALL`: All domain names are allowed.
	//
	// - `ALLOW_LIST`: Only websites in the allowlist are allowed.
	//
	// The returned value reflects the current configuration of the browser group.
	//
	// example:
	//
	// ALLOW_ALL
	AccessRestriction *string `json:"AccessRestriction,omitempty" xml:"AccessRestriction,omitempty"`
	// The ID of the workspace to which the browser group belongs.
	//
	// example:
	//
	// cn-hangzhou+dir-843734****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The list of vSwitch IDs used by the browser group, available for scenarios with custom network configurations.
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
}

func (s ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsNetwork) String() string {
	return dara.Prettify(s)
}

func (s ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsNetwork) GoString() string {
	return s.String()
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsNetwork) GetAccessRestriction() *string {
	return s.AccessRestriction
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsNetwork) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsNetwork) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsNetwork) SetAccessRestriction(v string) *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsNetwork {
	s.AccessRestriction = &v
	return s
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsNetwork) SetOfficeSiteId(v string) *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsNetwork {
	s.OfficeSiteId = &v
	return s
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsNetwork) SetVSwitchIds(v []*string) *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsNetwork {
	s.VSwitchIds = v
	return s
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsNetwork) Validate() error {
	return dara.Validate(s)
}

type ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsPolicy struct {
	// The screenshot protection switch.
	//
	// - `on`: Screenshot protection is enabled.
	//
	// - `off`: Screenshot protection is disabled.
	//
	// example:
	//
	// on
	AppContentProtection *string `json:"AppContentProtection,omitempty" xml:"AppContentProtection,omitempty"`
	// The client access IP address whitelist rules.
	AuthorizeAccessPolicyRules []*ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsPolicyAuthorizeAccessPolicyRules `json:"AuthorizeAccessPolicyRules,omitempty" xml:"AuthorizeAccessPolicyRules,omitempty" type:"Repeated"`
	// The client access control configuration list.
	ClientTypes []*ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsPolicyClientTypes `json:"ClientTypes,omitempty" xml:"ClientTypes,omitempty" type:"Repeated"`
	// The clipboard transfer direction, content type, and size limit settings. read indicates transfer from the local PC to the cloud browser. write indicates transfer from the cloud browser to the local PC.
	ClipboardPolicy *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsPolicyClipboardPolicy `json:"ClipboardPolicy,omitempty" xml:"ClipboardPolicy,omitempty" type:"Struct"`
	// The video display policy for browser sessions.
	VideoPolicy *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsPolicyVideoPolicy `json:"VideoPolicy,omitempty" xml:"VideoPolicy,omitempty" type:"Struct"`
	// The watermark display configuration for browser sessions.
	WatermarkPolicy *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsPolicyWatermarkPolicy `json:"WatermarkPolicy,omitempty" xml:"WatermarkPolicy,omitempty" type:"Struct"`
}

func (s ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsPolicy) String() string {
	return dara.Prettify(s)
}

func (s ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsPolicy) GoString() string {
	return s.String()
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsPolicy) GetAppContentProtection() *string {
	return s.AppContentProtection
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsPolicy) GetAuthorizeAccessPolicyRules() []*ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsPolicyAuthorizeAccessPolicyRules {
	return s.AuthorizeAccessPolicyRules
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsPolicy) GetClientTypes() []*ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsPolicyClientTypes {
	return s.ClientTypes
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsPolicy) GetClipboardPolicy() *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsPolicyClipboardPolicy {
	return s.ClipboardPolicy
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsPolicy) GetVideoPolicy() *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsPolicyVideoPolicy {
	return s.VideoPolicy
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsPolicy) GetWatermarkPolicy() *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsPolicyWatermarkPolicy {
	return s.WatermarkPolicy
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsPolicy) SetAppContentProtection(v string) *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsPolicy {
	s.AppContentProtection = &v
	return s
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsPolicy) SetAuthorizeAccessPolicyRules(v []*ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsPolicyAuthorizeAccessPolicyRules) *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsPolicy {
	s.AuthorizeAccessPolicyRules = v
	return s
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsPolicy) SetClientTypes(v []*ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsPolicyClientTypes) *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsPolicy {
	s.ClientTypes = v
	return s
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsPolicy) SetClipboardPolicy(v *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsPolicyClipboardPolicy) *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsPolicy {
	s.ClipboardPolicy = v
	return s
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsPolicy) SetVideoPolicy(v *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsPolicyVideoPolicy) *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsPolicy {
	s.VideoPolicy = v
	return s
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsPolicy) SetWatermarkPolicy(v *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsPolicyWatermarkPolicy) *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsPolicy {
	s.WatermarkPolicy = v
	return s
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsPolicy) Validate() error {
	if s.AuthorizeAccessPolicyRules != nil {
		for _, item := range s.AuthorizeAccessPolicyRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ClientTypes != nil {
		for _, item := range s.ClientTypes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ClipboardPolicy != nil {
		if err := s.ClipboardPolicy.Validate(); err != nil {
			return err
		}
	}
	if s.VideoPolicy != nil {
		if err := s.VideoPolicy.Validate(); err != nil {
			return err
		}
	}
	if s.WatermarkPolicy != nil {
		if err := s.WatermarkPolicy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsPolicyAuthorizeAccessPolicyRules struct {
	// The source CIDR block of the client that is allowed to access.
	//
	// example:
	//
	// 192.168.1.0/24
	CidrIp *string `json:"CidrIp,omitempty" xml:"CidrIp,omitempty"`
	// The description of the client access rule.
	//
	// example:
	//
	// Office network access
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The action of the client access rule.
	//
	// - `allow`: Access is allowed.
	//
	// - `deny`: Access is denied.
	//
	// example:
	//
	// allow
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
}

func (s ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsPolicyAuthorizeAccessPolicyRules) String() string {
	return dara.Prettify(s)
}

func (s ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsPolicyAuthorizeAccessPolicyRules) GoString() string {
	return s.String()
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsPolicyAuthorizeAccessPolicyRules) GetCidrIp() *string {
	return s.CidrIp
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsPolicyAuthorizeAccessPolicyRules) GetDescription() *string {
	return s.Description
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsPolicyAuthorizeAccessPolicyRules) GetPolicy() *string {
	return s.Policy
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsPolicyAuthorizeAccessPolicyRules) SetCidrIp(v string) *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsPolicyAuthorizeAccessPolicyRules {
	s.CidrIp = &v
	return s
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsPolicyAuthorizeAccessPolicyRules) SetDescription(v string) *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsPolicyAuthorizeAccessPolicyRules {
	s.Description = &v
	return s
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsPolicyAuthorizeAccessPolicyRules) SetPolicy(v string) *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsPolicyAuthorizeAccessPolicyRules {
	s.Policy = &v
	return s
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsPolicyAuthorizeAccessPolicyRules) Validate() error {
	return dara.Validate(s)
}

type ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsPolicyClientTypes struct {
	// The client type.
	//
	// - `windows`: Windows client.
	//
	// - `macos`: macOS client.
	//
	// - `html5`: Web client.
	//
	// - `linux`: Linux client.
	//
	// - `android`: Android client.
	//
	// - `ios`: iOS client.
	//
	// This field reflects the existing access configuration and does not indicate that all client types are available for the current product.
	//
	// example:
	//
	// windows
	ClientType *string `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	// The access policy switch for this client type.
	//
	// - `on`: Access from this client type is allowed.
	//
	// - `off`: Access from this client type is denied.
	//
	// example:
	//
	// on
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsPolicyClientTypes) String() string {
	return dara.Prettify(s)
}

func (s ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsPolicyClientTypes) GoString() string {
	return s.String()
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsPolicyClientTypes) GetClientType() *string {
	return s.ClientType
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsPolicyClientTypes) GetStatus() *string {
	return s.Status
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsPolicyClientTypes) SetClientType(v string) *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsPolicyClientTypes {
	s.ClientType = &v
	return s
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsPolicyClientTypes) SetStatus(v string) *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsPolicyClientTypes {
	s.Status = &v
	return s
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsPolicyClientTypes) Validate() error {
	return dara.Validate(s)
}

type ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsPolicyClipboardPolicy struct {
	// The clipboard transfer direction. The value is case-insensitive.
	//
	// - `off`: Bidirectional transfer is disabled.
	//
	// - `read`: Copy and paste from the local PC to the cloud browser is allowed.
	//
	// - `write`: Copy and paste from the cloud browser to the local PC is allowed.
	//
	// - `readwrite`: Bidirectional transfer is allowed.
	//
	// example:
	//
	// readwrite
	Clipboard *string `json:"Clipboard,omitempty" xml:"Clipboard,omitempty"`
}

func (s ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsPolicyClipboardPolicy) String() string {
	return dara.Prettify(s)
}

func (s ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsPolicyClipboardPolicy) GoString() string {
	return s.String()
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsPolicyClipboardPolicy) GetClipboard() *string {
	return s.Clipboard
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsPolicyClipboardPolicy) SetClipboard(v string) *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsPolicyClipboardPolicy {
	s.Clipboard = &v
	return s
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsPolicyClipboardPolicy) Validate() error {
	return dara.Validate(s)
}

type ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsPolicyVideoPolicy struct {
	// The frame rate of browser sessions.
	//
	// example:
	//
	// 30
	FrameRate *int32 `json:"FrameRate,omitempty" xml:"FrameRate,omitempty"`
}

func (s ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsPolicyVideoPolicy) String() string {
	return dara.Prettify(s)
}

func (s ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsPolicyVideoPolicy) GoString() string {
	return s.String()
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsPolicyVideoPolicy) GetFrameRate() *int32 {
	return s.FrameRate
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsPolicyVideoPolicy) SetFrameRate(v int32) *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsPolicyVideoPolicy {
	s.FrameRate = &v
	return s
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsPolicyVideoPolicy) Validate() error {
	return dara.Validate(s)
}

type ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsPolicyWatermarkPolicy struct {
	// The watermark switch. The value is case-insensitive. Valid values:
	//
	// - `ON`: Watermark is enabled.
	//
	// - `OFF`: Watermark is disabled.
	//
	// When disabled, the watermark content type list is not used.
	//
	// example:
	//
	// ON
	WatermarkSwitch *string `json:"WatermarkSwitch,omitempty" xml:"WatermarkSwitch,omitempty"`
	// The list of watermark content types. Valid values:
	//
	// - `EndUserId`: The user identifier.
	//
	// - `InstanceGroupId`: The delivery group identifier.
	//
	// - `ClientTime`: The current time on the client.
	//
	// Use watermark types that are supported by the browser and client.
	WatermarkTypes []*string `json:"WatermarkTypes,omitempty" xml:"WatermarkTypes,omitempty" type:"Repeated"`
}

func (s ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsPolicyWatermarkPolicy) String() string {
	return dara.Prettify(s)
}

func (s ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsPolicyWatermarkPolicy) GoString() string {
	return s.String()
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsPolicyWatermarkPolicy) GetWatermarkSwitch() *string {
	return s.WatermarkSwitch
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsPolicyWatermarkPolicy) GetWatermarkTypes() []*string {
	return s.WatermarkTypes
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsPolicyWatermarkPolicy) SetWatermarkSwitch(v string) *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsPolicyWatermarkPolicy {
	s.WatermarkSwitch = &v
	return s
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsPolicyWatermarkPolicy) SetWatermarkTypes(v []*string) *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsPolicyWatermarkPolicy {
	s.WatermarkTypes = v
	return s
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsPolicyWatermarkPolicy) Validate() error {
	return dara.Validate(s)
}

type ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsTags struct {
	// The tag key.
	//
	// example:
	//
	// usage
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// office
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsTags) String() string {
	return dara.Prettify(s)
}

func (s ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsTags) GoString() string {
	return s.String()
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsTags) GetKey() *string {
	return s.Key
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsTags) GetValue() *string {
	return s.Value
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsTags) SetKey(v string) *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsTags {
	s.Key = &v
	return s
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsTags) SetValue(v string) *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsTags {
	s.Value = &v
	return s
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsTags) Validate() error {
	return dara.Validate(s)
}

type ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsTimers struct {
	// The session retention duration after disconnection, in minutes. `-1` indicates that the session is not unbound due to this timeout. The session is still subject to authorization and other session release policies.
	//
	// example:
	//
	// 15
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The timer configuration type. `SESSION_TIMEOUT` indicates the session retention duration after disconnection.
	//
	// example:
	//
	// SESSION_TIMEOUT
	TimerType *string `json:"TimerType,omitempty" xml:"TimerType,omitempty"`
}

func (s ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsTimers) String() string {
	return dara.Prettify(s)
}

func (s ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsTimers) GoString() string {
	return s.String()
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsTimers) GetInterval() *int32 {
	return s.Interval
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsTimers) GetTimerType() *string {
	return s.TimerType
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsTimers) SetInterval(v int32) *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsTimers {
	s.Interval = &v
	return s
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsTimers) SetTimerType(v string) *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsTimers {
	s.TimerType = &v
	return s
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsTimers) Validate() error {
	return dara.Validate(s)
}

type ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsUserLimit struct {
	// The user quota.
	//
	// example:
	//
	// 100
	UserQuota *int32 `json:"UserQuota,omitempty" xml:"UserQuota,omitempty"`
}

func (s ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsUserLimit) String() string {
	return dara.Prettify(s)
}

func (s ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsUserLimit) GoString() string {
	return s.String()
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsUserLimit) GetUserQuota() *int32 {
	return s.UserQuota
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsUserLimit) SetUserQuota(v int32) *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsUserLimit {
	s.UserQuota = &v
	return s
}

func (s *ListBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelsUserLimit) Validate() error {
	return dara.Validate(s)
}
