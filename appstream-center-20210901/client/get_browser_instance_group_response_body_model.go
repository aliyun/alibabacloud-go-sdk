// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBrowserInstanceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBrowserInstanceGroupModel(v *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel) *GetBrowserInstanceGroupResponseBody
	GetBrowserInstanceGroupModel() *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel
	SetRequestId(v string) *GetBrowserInstanceGroupResponseBody
	GetRequestId() *string
}

type GetBrowserInstanceGroupResponseBody struct {
	// The details of the browser group.
	BrowserInstanceGroupModel *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel `json:"BrowserInstanceGroupModel,omitempty" xml:"BrowserInstanceGroupModel,omitempty" type:"Struct"`
	// The request ID, which is used for troubleshooting.
	//
	// example:
	//
	// 01A0C2ED-95F2-1A37-9FC6-4A395179****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetBrowserInstanceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBrowserInstanceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetBrowserInstanceGroupResponseBody) GetBrowserInstanceGroupModel() *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel {
	return s.BrowserInstanceGroupModel
}

func (s *GetBrowserInstanceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetBrowserInstanceGroupResponseBody) SetBrowserInstanceGroupModel(v *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel) *GetBrowserInstanceGroupResponseBody {
	s.BrowserInstanceGroupModel = v
	return s
}

func (s *GetBrowserInstanceGroupResponseBody) SetRequestId(v string) *GetBrowserInstanceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBrowserInstanceGroupResponseBody) Validate() error {
	if s.BrowserInstanceGroupModel != nil {
		if err := s.BrowserInstanceGroupModel.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel struct {
	// Indicates whether authorization and deauthorization notification emails are enabled. `true` indicates that the feature is enabled. `false` indicates that the feature is disabled.
	//
	// example:
	//
	// true
	AuthNotificationEnabled *bool `json:"AuthNotificationEnabled,omitempty" xml:"AuthNotificationEnabled,omitempty"`
	// The statistics of authorized users for the browser group.
	AuthorizedUserInfo *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelAuthorizedUserInfo `json:"AuthorizedUserInfo,omitempty" xml:"AuthorizedUserInfo,omitempty" type:"Struct"`
	// The business region where the browser group resides.
	//
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// The current browser configuration.
	BrowserConfig *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelBrowserConfig `json:"BrowserConfig,omitempty" xml:"BrowserConfig,omitempty" type:"Struct"`
	// The cloud browser group ID.
	//
	// example:
	//
	// big-0c7loey7fzjq****
	BrowserInstanceGroupId *string `json:"BrowserInstanceGroupId,omitempty" xml:"BrowserInstanceGroupId,omitempty"`
	// The cloud browser group name.
	//
	// example:
	//
	// Office Browser
	BrowserInstanceGroupName *string `json:"BrowserInstanceGroupName,omitempty" xml:"BrowserInstanceGroupName,omitempty"`
	// The ID of the browser group set to which the browser group belongs.
	//
	// example:
	//
	// set-3jm9d0abc00example
	BrowserInstanceGroupSetId *string `json:"BrowserInstanceGroupSetId,omitempty" xml:"BrowserInstanceGroupSetId,omitempty"`
	// The billing type. In MAU scenarios, `PostPaid` is returned, which indicates the pay-as-you-go billing method.
	//
	// example:
	//
	// PostPaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The default access URL of the browser group. Use the URL returned by the API to access the browser group. Replace the resource ID in the example with your actual value.
	//
	// example:
	//
	// https://wuying.aliyun.com/integration?appId=browser&appInstanceGroupId=big-0c7loey7fzjq****
	DefaultAccessUrl *string `json:"DefaultAccessUrl,omitempty" xml:"DefaultAccessUrl,omitempty"`
	// The plan duration information. In MAU scenarios, plan duration does not apply, and an empty object may be returned.
	Duration *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelDuration `json:"Duration,omitempty" xml:"Duration,omitempty" type:"Struct"`
	// The expiration time of the browser group. This field does not apply to MAU scenarios and is not returned.
	//
	// example:
	//
	// -
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The creation time of the browser group. The value is an RFC 3339 time string in the `yyyy-MM-ddTHH:mm:ss.SSSXXX` format, which includes milliseconds and a time zone offset. The `+00:00` in the example indicates the UTC time zone.
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
	// The maximum number of instances configured for the MAU scenario.
	//
	// example:
	//
	// 100
	MaxAmount *int32 `json:"MaxAmount,omitempty" xml:"MaxAmount,omitempty"`
	// The office network and website access restriction configuration.
	Network *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelNetwork `json:"Network,omitempty" xml:"Network,omitempty" type:"Struct"`
	// The node specifications information. This field does not apply to MAU scenarios.
	NodeInstanceType *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelNodeInstanceType `json:"NodeInstanceType,omitempty" xml:"NodeInstanceType,omitempty" type:"Struct"`
	// The list of node pool information. In MAU scenarios, this field does not apply and an empty list may be returned.
	NodePool []*GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelNodePool `json:"NodePool,omitempty" xml:"NodePool,omitempty" type:"Repeated"`
	// The operating system type of the browser group. The current MAU product scenario uses Windows.
	//
	// example:
	//
	// Windows
	OsType *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	// The policy configuration returned for the browser group. The policy fields are used to view existing settings and do not indicate that all corresponding creation parameters are configurable.
	Policy *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicy `json:"Policy,omitempty" xml:"Policy,omitempty" type:"Struct"`
	// The status of the browser instance group.
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
	// The sub-billing type. In MAU scenarios, the actual returned value is `mau`, which indicates billing by monthly active users.
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
	// The version of the browser.
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
	// The session timer configurations currently returned. These are for viewing the effective settings and do not indicate that the creation API supports setting this parameter.
	Timers []*GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelTimers `json:"Timers,omitempty" xml:"Timers,omitempty" type:"Repeated"`
	// The current authorization mode.
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
	UserLimit *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelUserLimit `json:"UserLimit,omitempty" xml:"UserLimit,omitempty" type:"Struct"`
}

func (s GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel) String() string {
	return dara.Prettify(s)
}

func (s GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel) GoString() string {
	return s.String()
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel) GetAuthNotificationEnabled() *bool {
	return s.AuthNotificationEnabled
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel) GetAuthorizedUserInfo() *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelAuthorizedUserInfo {
	return s.AuthorizedUserInfo
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel) GetBizRegionId() *string {
	return s.BizRegionId
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel) GetBrowserConfig() *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelBrowserConfig {
	return s.BrowserConfig
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel) GetBrowserInstanceGroupId() *string {
	return s.BrowserInstanceGroupId
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel) GetBrowserInstanceGroupName() *string {
	return s.BrowserInstanceGroupName
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel) GetBrowserInstanceGroupSetId() *string {
	return s.BrowserInstanceGroupSetId
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel) GetChargeType() *string {
	return s.ChargeType
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel) GetDefaultAccessUrl() *string {
	return s.DefaultAccessUrl
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel) GetDuration() *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelDuration {
	return s.Duration
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel) GetHomepage() *string {
	return s.Homepage
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel) GetImageId() *string {
	return s.ImageId
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel) GetInstanceType() *string {
	return s.InstanceType
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel) GetMaxAmount() *int32 {
	return s.MaxAmount
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel) GetNetwork() *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelNetwork {
	return s.Network
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel) GetNodeInstanceType() *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelNodeInstanceType {
	return s.NodeInstanceType
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel) GetNodePool() []*GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelNodePool {
	return s.NodePool
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel) GetOsType() *string {
	return s.OsType
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel) GetPolicy() *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicy {
	return s.Policy
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel) GetStatus() *string {
	return s.Status
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel) GetSubPayType() *string {
	return s.SubPayType
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel) GetSupportUserGroupMixedAuth() *bool {
	return s.SupportUserGroupMixedAuth
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel) GetTier() *string {
	return s.Tier
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel) GetTimers() []*GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelTimers {
	return s.Timers
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel) GetUserGroupAuthMode() *string {
	return s.UserGroupAuthMode
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel) GetUserLimit() *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelUserLimit {
	return s.UserLimit
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel) SetAuthNotificationEnabled(v bool) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel {
	s.AuthNotificationEnabled = &v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel) SetAuthorizedUserInfo(v *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelAuthorizedUserInfo) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel {
	s.AuthorizedUserInfo = v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel) SetBizRegionId(v string) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel {
	s.BizRegionId = &v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel) SetBrowserConfig(v *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelBrowserConfig) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel {
	s.BrowserConfig = v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel) SetBrowserInstanceGroupId(v string) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel {
	s.BrowserInstanceGroupId = &v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel) SetBrowserInstanceGroupName(v string) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel {
	s.BrowserInstanceGroupName = &v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel) SetBrowserInstanceGroupSetId(v string) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel {
	s.BrowserInstanceGroupSetId = &v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel) SetChargeType(v string) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel {
	s.ChargeType = &v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel) SetDefaultAccessUrl(v string) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel {
	s.DefaultAccessUrl = &v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel) SetDuration(v *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelDuration) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel {
	s.Duration = v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel) SetExpiredTime(v string) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel {
	s.ExpiredTime = &v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel) SetGmtCreate(v string) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel {
	s.GmtCreate = &v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel) SetHomepage(v string) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel {
	s.Homepage = &v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel) SetImageId(v string) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel {
	s.ImageId = &v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel) SetInstanceType(v string) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel {
	s.InstanceType = &v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel) SetMaxAmount(v int32) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel {
	s.MaxAmount = &v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel) SetNetwork(v *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelNetwork) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel {
	s.Network = v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel) SetNodeInstanceType(v *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelNodeInstanceType) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel {
	s.NodeInstanceType = v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel) SetNodePool(v []*GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelNodePool) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel {
	s.NodePool = v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel) SetOsType(v string) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel {
	s.OsType = &v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel) SetPolicy(v *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicy) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel {
	s.Policy = v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel) SetStatus(v string) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel {
	s.Status = &v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel) SetSubPayType(v string) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel {
	s.SubPayType = &v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel) SetSupportUserGroupMixedAuth(v bool) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel {
	s.SupportUserGroupMixedAuth = &v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel) SetTier(v string) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel {
	s.Tier = &v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel) SetTimers(v []*GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelTimers) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel {
	s.Timers = v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel) SetUserGroupAuthMode(v string) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel {
	s.UserGroupAuthMode = &v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel) SetUserLimit(v *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelUserLimit) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel {
	s.UserLimit = v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModel) Validate() error {
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
	if s.NodeInstanceType != nil {
		if err := s.NodeInstanceType.Validate(); err != nil {
			return err
		}
	}
	if s.NodePool != nil {
		for _, item := range s.NodePool {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Policy != nil {
		if err := s.Policy.Validate(); err != nil {
			return err
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

type GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelAuthorizedUserInfo struct {
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

func (s GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelAuthorizedUserInfo) String() string {
	return dara.Prettify(s)
}

func (s GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelAuthorizedUserInfo) GoString() string {
	return s.String()
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelAuthorizedUserInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelAuthorizedUserInfo) GetTotalUserGroupCount() *int32 {
	return s.TotalUserGroupCount
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelAuthorizedUserInfo) SetTotalCount(v int32) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelAuthorizedUserInfo {
	s.TotalCount = &v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelAuthorizedUserInfo) SetTotalUserGroupCount(v int32) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelAuthorizedUserInfo {
	s.TotalUserGroupCount = &v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelAuthorizedUserInfo) Validate() error {
	return dara.Validate(s)
}

type GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelBrowserConfig struct {
	// The list of browser bookmarks. A maximum of 20 entries are returned. To query the complete bookmark list, call `ListBrowserBookmarks`.
	Bookmarks []*GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelBrowserConfigBookmarks `json:"Bookmarks,omitempty" xml:"Bookmarks,omitempty" type:"Repeated"`
	// The browser startup parameters. For example, `--incognito` specifies the incognito window mode.
	//
	// example:
	//
	// --incognito
	BrowserParam *string `json:"BrowserParam,omitempty" xml:"BrowserParam,omitempty"`
	// The cookie synchronization configuration. The string `true` indicates that synchronization is enabled. The string `false` indicates that synchronization is disabled.
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

func (s GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelBrowserConfig) String() string {
	return dara.Prettify(s)
}

func (s GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelBrowserConfig) GoString() string {
	return s.String()
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelBrowserConfig) GetBookmarks() []*GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelBrowserConfigBookmarks {
	return s.Bookmarks
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelBrowserConfig) GetBrowserParam() *string {
	return s.BrowserParam
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelBrowserConfig) GetCookiesSync() *string {
	return s.CookiesSync
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelBrowserConfig) GetHomepage() *string {
	return s.Homepage
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelBrowserConfig) SetBookmarks(v []*GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelBrowserConfigBookmarks) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelBrowserConfig {
	s.Bookmarks = v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelBrowserConfig) SetBrowserParam(v string) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelBrowserConfig {
	s.BrowserParam = &v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelBrowserConfig) SetCookiesSync(v string) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelBrowserConfig {
	s.CookiesSync = &v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelBrowserConfig) SetHomepage(v string) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelBrowserConfig {
	s.Homepage = &v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelBrowserConfig) Validate() error {
	if s.Bookmarks != nil {
		for _, item := range s.Bookmarks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelBrowserConfigBookmarks struct {
	// The folder in which the bookmark resides.
	//
	// example:
	//
	// Work Portal
	BookmarkFolder *string `json:"BookmarkFolder,omitempty" xml:"BookmarkFolder,omitempty"`
	// The bookmark ID.
	//
	// example:
	//
	// bm-12345
	BookmarkId *string `json:"BookmarkId,omitempty" xml:"BookmarkId,omitempty"`
	// The bookmark name.
	//
	// example:
	//
	// Alibaba Cloud Official Website
	BookmarkName *string `json:"BookmarkName,omitempty" xml:"BookmarkName,omitempty"`
	// The URL of the bookmark.
	//
	// example:
	//
	// https://www.aliyun.com
	BookmarkURL *string `json:"BookmarkURL,omitempty" xml:"BookmarkURL,omitempty"`
}

func (s GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelBrowserConfigBookmarks) String() string {
	return dara.Prettify(s)
}

func (s GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelBrowserConfigBookmarks) GoString() string {
	return s.String()
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelBrowserConfigBookmarks) GetBookmarkFolder() *string {
	return s.BookmarkFolder
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelBrowserConfigBookmarks) GetBookmarkId() *string {
	return s.BookmarkId
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelBrowserConfigBookmarks) GetBookmarkName() *string {
	return s.BookmarkName
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelBrowserConfigBookmarks) GetBookmarkURL() *string {
	return s.BookmarkURL
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelBrowserConfigBookmarks) SetBookmarkFolder(v string) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelBrowserConfigBookmarks {
	s.BookmarkFolder = &v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelBrowserConfigBookmarks) SetBookmarkId(v string) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelBrowserConfigBookmarks {
	s.BookmarkId = &v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelBrowserConfigBookmarks) SetBookmarkName(v string) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelBrowserConfigBookmarks {
	s.BookmarkName = &v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelBrowserConfigBookmarks) SetBookmarkURL(v string) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelBrowserConfigBookmarks {
	s.BookmarkURL = &v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelBrowserConfigBookmarks) Validate() error {
	return dara.Validate(s)
}

type GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelDuration struct {
	// The current payment stage of the plan. This field does not apply to MAU scenarios.
	//
	// example:
	//
	// -
	CurrentPayStage *string `json:"CurrentPayStage,omitempty" xml:"CurrentPayStage,omitempty"`
	// The end time of the plan period. This field does not apply to MAU scenarios and is not returned.
	//
	// example:
	//
	// -
	PeriodEndTime *string `json:"PeriodEndTime,omitempty" xml:"PeriodEndTime,omitempty"`
	// The start time of the plan period. This field does not apply to MAU scenarios and is not returned.
	//
	// example:
	//
	// -
	PeriodStartTime *string `json:"PeriodStartTime,omitempty" xml:"PeriodStartTime,omitempty"`
	// The total duration of the plan, in seconds. This field does not apply to MAU scenarios.
	//
	// example:
	//
	// -
	TotalDuration *int32 `json:"TotalDuration,omitempty" xml:"TotalDuration,omitempty"`
	// The used duration of the plan, in seconds. This field does not apply to MAU scenarios.
	//
	// example:
	//
	// -
	UsedDuration *int32 `json:"UsedDuration,omitempty" xml:"UsedDuration,omitempty"`
}

func (s GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelDuration) String() string {
	return dara.Prettify(s)
}

func (s GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelDuration) GoString() string {
	return s.String()
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelDuration) GetCurrentPayStage() *string {
	return s.CurrentPayStage
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelDuration) GetPeriodEndTime() *string {
	return s.PeriodEndTime
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelDuration) GetPeriodStartTime() *string {
	return s.PeriodStartTime
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelDuration) GetTotalDuration() *int32 {
	return s.TotalDuration
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelDuration) GetUsedDuration() *int32 {
	return s.UsedDuration
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelDuration) SetCurrentPayStage(v string) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelDuration {
	s.CurrentPayStage = &v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelDuration) SetPeriodEndTime(v string) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelDuration {
	s.PeriodEndTime = &v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelDuration) SetPeriodStartTime(v string) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelDuration {
	s.PeriodStartTime = &v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelDuration) SetTotalDuration(v int32) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelDuration {
	s.TotalDuration = &v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelDuration) SetUsedDuration(v int32) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelDuration {
	s.UsedDuration = &v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelDuration) Validate() error {
	return dara.Validate(s)
}

type GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelNetwork struct {
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
	// The office network ID to which the browser group belongs.
	//
	// example:
	//
	// cn-hangzhou+dir-843734****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The website access restriction list. A maximum of 20 entries are returned. To query the complete list, call `ListBrowserRestrictedURLs`.
	RestrictedURLs []*GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelNetworkRestrictedURLs `json:"RestrictedURLs,omitempty" xml:"RestrictedURLs,omitempty" type:"Repeated"`
	// The list of vSwitch IDs used by the browser group. This is available for scenarios with custom network configurations.
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
}

func (s GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelNetwork) String() string {
	return dara.Prettify(s)
}

func (s GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelNetwork) GoString() string {
	return s.String()
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelNetwork) GetAccessRestriction() *string {
	return s.AccessRestriction
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelNetwork) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelNetwork) GetRestrictedURLs() []*GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelNetworkRestrictedURLs {
	return s.RestrictedURLs
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelNetwork) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelNetwork) SetAccessRestriction(v string) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelNetwork {
	s.AccessRestriction = &v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelNetwork) SetOfficeSiteId(v string) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelNetwork {
	s.OfficeSiteId = &v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelNetwork) SetRestrictedURLs(v []*GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelNetworkRestrictedURLs) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelNetwork {
	s.RestrictedURLs = v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelNetwork) SetVSwitchIds(v []*string) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelNetwork {
	s.VSwitchIds = v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelNetwork) Validate() error {
	if s.RestrictedURLs != nil {
		for _, item := range s.RestrictedURLs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelNetworkRestrictedURLs struct {
	// The ID of the website access restriction entry.
	//
	// example:
	//
	// ru-12345
	RestrictedURLId *string `json:"RestrictedURLId,omitempty" xml:"RestrictedURLId,omitempty"`
	// The website URL in the access restriction entry.
	//
	// example:
	//
	// aliyun.com
	URL *string `json:"URL,omitempty" xml:"URL,omitempty"`
}

func (s GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelNetworkRestrictedURLs) String() string {
	return dara.Prettify(s)
}

func (s GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelNetworkRestrictedURLs) GoString() string {
	return s.String()
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelNetworkRestrictedURLs) GetRestrictedURLId() *string {
	return s.RestrictedURLId
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelNetworkRestrictedURLs) GetURL() *string {
	return s.URL
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelNetworkRestrictedURLs) SetRestrictedURLId(v string) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelNetworkRestrictedURLs {
	s.RestrictedURLId = &v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelNetworkRestrictedURLs) SetURL(v string) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelNetworkRestrictedURLs {
	s.URL = &v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelNetworkRestrictedURLs) Validate() error {
	return dara.Validate(s)
}

type GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelNodeInstanceType struct {
	// The CPU configuration of the node. This field does not apply to MAU scenarios.
	//
	// example:
	//
	// -
	Cpu *string `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The memory configuration of the node. This field does not apply to MAU scenarios.
	//
	// example:
	//
	// -
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
}

func (s GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelNodeInstanceType) String() string {
	return dara.Prettify(s)
}

func (s GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelNodeInstanceType) GoString() string {
	return s.String()
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelNodeInstanceType) GetCpu() *string {
	return s.Cpu
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelNodeInstanceType) GetMemory() *int32 {
	return s.Memory
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelNodeInstanceType) SetCpu(v string) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelNodeInstanceType {
	s.Cpu = &v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelNodeInstanceType) SetMemory(v int32) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelNodeInstanceType {
	s.Memory = &v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelNodeInstanceType) Validate() error {
	return dara.Validate(s)
}

type GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelNodePool struct {
	// The total number of nodes. This field does not apply to MAU scenarios.
	//
	// example:
	//
	// -
	NodeAmount *string `json:"NodeAmount,omitempty" xml:"NodeAmount,omitempty"`
	// The number of used nodes. This field does not apply to MAU scenarios.
	//
	// example:
	//
	// -
	NodeUsed *string `json:"NodeUsed,omitempty" xml:"NodeUsed,omitempty"`
}

func (s GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelNodePool) String() string {
	return dara.Prettify(s)
}

func (s GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelNodePool) GoString() string {
	return s.String()
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelNodePool) GetNodeAmount() *string {
	return s.NodeAmount
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelNodePool) GetNodeUsed() *string {
	return s.NodeUsed
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelNodePool) SetNodeAmount(v string) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelNodePool {
	s.NodeAmount = &v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelNodePool) SetNodeUsed(v string) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelNodePool {
	s.NodeUsed = &v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelNodePool) Validate() error {
	return dara.Validate(s)
}

type GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicy struct {
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
	AuthorizeAccessPolicyRules []*GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyAuthorizeAccessPolicyRules `json:"AuthorizeAccessPolicyRules,omitempty" xml:"AuthorizeAccessPolicyRules,omitempty" type:"Repeated"`
	// The list of client access control configurations.
	ClientTypes []*GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyClientTypes `json:"ClientTypes,omitempty" xml:"ClientTypes,omitempty" type:"Repeated"`
	// The clipboard transfer direction, content type, and size limit settings. read indicates transfer from the local PC to the cloud browser. write indicates transfer from the cloud browser to the local PC.
	ClipboardPolicy *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyClipboardPolicy `json:"ClipboardPolicy,omitempty" xml:"ClipboardPolicy,omitempty" type:"Struct"`
	// The data retention policy for sessions after disconnection.
	//
	// - `customTime`: The session is retained for the duration specified by `DisconnectKeepSessionTime`.
	//
	// - `persistent`: The session is not subject to automatic release based on disconnection duration.
	//
	// **Note:*	- The `persistent` option is still subject to authorization and other release policies.
	//
	// example:
	//
	// customTime
	DisconnectKeepSession *string `json:"DisconnectKeepSession,omitempty" xml:"DisconnectKeepSession,omitempty"`
	// The session retention duration after disconnection. Unit: seconds. This value is for viewing the configuration only and does not indicate that this parameter can be set through the create operation.
	//
	// example:
	//
	// 1800
	DisconnectKeepSessionTime *int32 `json:"DisconnectKeepSessionTime,omitempty" xml:"DisconnectKeepSessionTime,omitempty"`
	// The floating ball file manager switch.
	//
	// - `on`: Enabled.
	//
	// - `off`: Disabled.
	//
	// example:
	//
	// off
	FileManager *string `json:"FileManager,omitempty" xml:"FileManager,omitempty"`
	// The file transfer policy for the web client.
	//
	// - `off`: Transfer is denied.
	//
	// - `upload`: Only upload is allowed.
	//
	// - `download`: Only download is allowed.
	//
	// - `full`: Both upload and download are allowed.
	//
	// Configure this parameter together with the clipboard policy.
	//
	// example:
	//
	// full
	Html5FileTransfer *string `json:"Html5FileTransfer,omitempty" xml:"Html5FileTransfer,omitempty"`
	// The switch for automatic disconnection upon no operation. The value is case-insensitive.
	//
	// - `on`: Enabled.
	//
	// - `off`: Disabled.
	//
	// When enabled, use `NoOperationDisconnectTime` to set the wait duration.
	//
	// example:
	//
	// on
	NoOperationDisconnect *string `json:"NoOperationDisconnect,omitempty" xml:"NoOperationDisconnect,omitempty"`
	// The wait duration before disconnection is triggered after no operation, in seconds. Whether this feature is enabled is indicated by `NoOperationDisconnect`.
	//
	// example:
	//
	// 600
	NoOperationDisconnectTime *int32 `json:"NoOperationDisconnectTime,omitempty" xml:"NoOperationDisconnectTime,omitempty"`
	// The ID of the policy associated with the browser instance group.
	//
	// example:
	//
	// pg-0bf5d87epuq5****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The policy version.
	//
	// - `DEFAULT`: Legacy policy.
	//
	// - `CENTER`: Centralized policy.
	//
	// example:
	//
	// CENTER
	PolicyVersion *string `json:"PolicyVersion,omitempty" xml:"PolicyVersion,omitempty"`
	// The video display policy for browser sessions.
	VideoPolicy *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyVideoPolicy `json:"VideoPolicy,omitempty" xml:"VideoPolicy,omitempty" type:"Struct"`
	// The watermark display configuration for browser sessions.
	WatermarkPolicy *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyWatermarkPolicy `json:"WatermarkPolicy,omitempty" xml:"WatermarkPolicy,omitempty" type:"Struct"`
}

func (s GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicy) String() string {
	return dara.Prettify(s)
}

func (s GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicy) GoString() string {
	return s.String()
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicy) GetAppContentProtection() *string {
	return s.AppContentProtection
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicy) GetAuthorizeAccessPolicyRules() []*GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyAuthorizeAccessPolicyRules {
	return s.AuthorizeAccessPolicyRules
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicy) GetClientTypes() []*GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyClientTypes {
	return s.ClientTypes
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicy) GetClipboardPolicy() *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyClipboardPolicy {
	return s.ClipboardPolicy
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicy) GetDisconnectKeepSession() *string {
	return s.DisconnectKeepSession
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicy) GetDisconnectKeepSessionTime() *int32 {
	return s.DisconnectKeepSessionTime
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicy) GetFileManager() *string {
	return s.FileManager
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicy) GetHtml5FileTransfer() *string {
	return s.Html5FileTransfer
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicy) GetNoOperationDisconnect() *string {
	return s.NoOperationDisconnect
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicy) GetNoOperationDisconnectTime() *int32 {
	return s.NoOperationDisconnectTime
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicy) GetPolicyId() *string {
	return s.PolicyId
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicy) GetPolicyVersion() *string {
	return s.PolicyVersion
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicy) GetVideoPolicy() *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyVideoPolicy {
	return s.VideoPolicy
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicy) GetWatermarkPolicy() *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyWatermarkPolicy {
	return s.WatermarkPolicy
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicy) SetAppContentProtection(v string) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicy {
	s.AppContentProtection = &v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicy) SetAuthorizeAccessPolicyRules(v []*GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyAuthorizeAccessPolicyRules) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicy {
	s.AuthorizeAccessPolicyRules = v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicy) SetClientTypes(v []*GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyClientTypes) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicy {
	s.ClientTypes = v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicy) SetClipboardPolicy(v *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyClipboardPolicy) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicy {
	s.ClipboardPolicy = v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicy) SetDisconnectKeepSession(v string) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicy {
	s.DisconnectKeepSession = &v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicy) SetDisconnectKeepSessionTime(v int32) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicy {
	s.DisconnectKeepSessionTime = &v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicy) SetFileManager(v string) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicy {
	s.FileManager = &v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicy) SetHtml5FileTransfer(v string) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicy {
	s.Html5FileTransfer = &v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicy) SetNoOperationDisconnect(v string) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicy {
	s.NoOperationDisconnect = &v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicy) SetNoOperationDisconnectTime(v int32) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicy {
	s.NoOperationDisconnectTime = &v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicy) SetPolicyId(v string) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicy {
	s.PolicyId = &v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicy) SetPolicyVersion(v string) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicy {
	s.PolicyVersion = &v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicy) SetVideoPolicy(v *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyVideoPolicy) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicy {
	s.VideoPolicy = v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicy) SetWatermarkPolicy(v *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyWatermarkPolicy) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicy {
	s.WatermarkPolicy = v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicy) Validate() error {
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

type GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyAuthorizeAccessPolicyRules struct {
	// The client source CIDR block that is allowed to access the browser group.
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

func (s GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyAuthorizeAccessPolicyRules) String() string {
	return dara.Prettify(s)
}

func (s GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyAuthorizeAccessPolicyRules) GoString() string {
	return s.String()
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyAuthorizeAccessPolicyRules) GetCidrIp() *string {
	return s.CidrIp
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyAuthorizeAccessPolicyRules) GetDescription() *string {
	return s.Description
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyAuthorizeAccessPolicyRules) GetPolicy() *string {
	return s.Policy
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyAuthorizeAccessPolicyRules) SetCidrIp(v string) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyAuthorizeAccessPolicyRules {
	s.CidrIp = &v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyAuthorizeAccessPolicyRules) SetDescription(v string) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyAuthorizeAccessPolicyRules {
	s.Description = &v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyAuthorizeAccessPolicyRules) SetPolicy(v string) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyAuthorizeAccessPolicyRules {
	s.Policy = &v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyAuthorizeAccessPolicyRules) Validate() error {
	return dara.Validate(s)
}

type GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyClientTypes struct {
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
	// The access policy switch for the client type.
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

func (s GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyClientTypes) String() string {
	return dara.Prettify(s)
}

func (s GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyClientTypes) GoString() string {
	return s.String()
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyClientTypes) GetClientType() *string {
	return s.ClientType
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyClientTypes) GetStatus() *string {
	return s.Status
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyClientTypes) SetClientType(v string) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyClientTypes {
	s.ClientType = &v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyClientTypes) SetStatus(v string) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyClientTypes {
	s.Status = &v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyClientTypes) Validate() error {
	return dara.Validate(s)
}

type GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyClipboardPolicy struct {
	// The clipboard transfer direction. The value is case-insensitive.
	//
	// - `off`: Bidirectional transfer is denied.
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
	// The clipboard size limit for inbound transfer (from the local PC to the cloud browser).
	//
	// example:
	//
	// 1024
	ClipboardReadLimit *int32 `json:"ClipboardReadLimit,omitempty" xml:"ClipboardReadLimit,omitempty"`
	// The clipboard control granularity.
	//
	// - `global`: Unified control.
	//
	// - `grained`: Separate control by text, rich text, and file.
	//
	// example:
	//
	// grained
	ClipboardScope *string `json:"ClipboardScope,omitempty" xml:"ClipboardScope,omitempty"`
	// The unit of the clipboard size.
	//
	// - `B`: bytes.
	//
	// - `KB`: 1024 bytes.
	//
	// example:
	//
	// KB
	ClipboardSizeUnit *string `json:"ClipboardSizeUnit,omitempty" xml:"ClipboardSizeUnit,omitempty"`
	// The clipboard size limit for outbound transfer (from the cloud browser to the local PC).
	//
	// example:
	//
	// 1024
	ClipboardWriteLimit *int32 `json:"ClipboardWriteLimit,omitempty" xml:"ClipboardWriteLimit,omitempty"`
	// The clipboard transfer direction. The value is case-insensitive.
	//
	// - `off`: Bidirectional transfer is denied.
	//
	// - `read`: Copy and paste from the local PC to the cloud browser is allowed.
	//
	// - `write`: Copy and paste from the cloud browser to the local PC is allowed.
	//
	// - `readwrite`: Bidirectional transfer is allowed.
	//
	// example:
	//
	// off
	FileClipboard *string `json:"FileClipboard,omitempty" xml:"FileClipboard,omitempty"`
	// The clipboard transfer direction. The value is case-insensitive.
	//
	// - `off`: Bidirectional transfer is denied.
	//
	// - `read`: Copy and paste from the local PC to the cloud browser is allowed.
	//
	// - `write`: Copy and paste from the cloud browser to the local PC is allowed.
	//
	// - `readwrite`: Bidirectional transfer is allowed.
	//
	// example:
	//
	// off
	RichTextClipboard *string `json:"RichTextClipboard,omitempty" xml:"RichTextClipboard,omitempty"`
	// The rich text clipboard size limit.
	//
	// example:
	//
	// 1024
	RichTextClipboardLimit *int32 `json:"RichTextClipboardLimit,omitempty" xml:"RichTextClipboardLimit,omitempty"`
	// The clipboard size limit for inbound transfer (from the local PC to the cloud browser).
	//
	// example:
	//
	// 1
	RichTextClipboardReadLimit *int32 `json:"RichTextClipboardReadLimit,omitempty" xml:"RichTextClipboardReadLimit,omitempty"`
	// The unit of the clipboard size.
	//
	// - `B`: bytes.
	//
	// - `KB`: 1024 bytes.
	//
	// example:
	//
	// KB
	RichTextClipboardReadSizeUnit *string `json:"RichTextClipboardReadSizeUnit,omitempty" xml:"RichTextClipboardReadSizeUnit,omitempty"`
	// The unit of the clipboard size.
	//
	// - `B`: bytes.
	//
	// - `KB`: 1024 bytes.
	//
	// example:
	//
	// KB
	RichTextClipboardSizeUnit *string `json:"RichTextClipboardSizeUnit,omitempty" xml:"RichTextClipboardSizeUnit,omitempty"`
	// The clipboard size limit for outbound transfer (from the cloud browser to the local PC).
	//
	// example:
	//
	// 1
	RichTextClipboardWriteLimit *int32 `json:"RichTextClipboardWriteLimit,omitempty" xml:"RichTextClipboardWriteLimit,omitempty"`
	// The unit of the clipboard size.
	//
	// - `B`: bytes.
	//
	// - `KB`: 1024 bytes.
	//
	// example:
	//
	// KB
	RichTextClipboardWriteSizeUnit *string `json:"RichTextClipboardWriteSizeUnit,omitempty" xml:"RichTextClipboardWriteSizeUnit,omitempty"`
	// The clipboard transfer direction. The value is case-insensitive.
	//
	// - `off`: Bidirectional transfer is denied.
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
	TextClipboard *string `json:"TextClipboard,omitempty" xml:"TextClipboard,omitempty"`
	// The clipboard size limit for inbound transfer (from the local PC to the cloud browser).
	//
	// example:
	//
	// 1
	TextClipboardReadLimit *int32 `json:"TextClipboardReadLimit,omitempty" xml:"TextClipboardReadLimit,omitempty"`
	// The unit of the clipboard size.
	//
	// - `B`: bytes.
	//
	// - `KB`: 1024 bytes.
	//
	// example:
	//
	// KB
	TextClipboardReadSizeUnit *string `json:"TextClipboardReadSizeUnit,omitempty" xml:"TextClipboardReadSizeUnit,omitempty"`
	// The clipboard size limit for outbound transfer (from the cloud browser to the local PC).
	//
	// example:
	//
	// 1
	TextClipboardWriteLimit *int32 `json:"TextClipboardWriteLimit,omitempty" xml:"TextClipboardWriteLimit,omitempty"`
	// The unit of the clipboard size.
	//
	// - `B`: bytes.
	//
	// - `KB`: 1024 bytes.
	//
	// example:
	//
	// KB
	TextClipboardWriteSizeUnit *string `json:"TextClipboardWriteSizeUnit,omitempty" xml:"TextClipboardWriteSizeUnit,omitempty"`
}

func (s GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyClipboardPolicy) String() string {
	return dara.Prettify(s)
}

func (s GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyClipboardPolicy) GoString() string {
	return s.String()
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyClipboardPolicy) GetClipboard() *string {
	return s.Clipboard
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyClipboardPolicy) GetClipboardReadLimit() *int32 {
	return s.ClipboardReadLimit
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyClipboardPolicy) GetClipboardScope() *string {
	return s.ClipboardScope
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyClipboardPolicy) GetClipboardSizeUnit() *string {
	return s.ClipboardSizeUnit
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyClipboardPolicy) GetClipboardWriteLimit() *int32 {
	return s.ClipboardWriteLimit
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyClipboardPolicy) GetFileClipboard() *string {
	return s.FileClipboard
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyClipboardPolicy) GetRichTextClipboard() *string {
	return s.RichTextClipboard
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyClipboardPolicy) GetRichTextClipboardLimit() *int32 {
	return s.RichTextClipboardLimit
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyClipboardPolicy) GetRichTextClipboardReadLimit() *int32 {
	return s.RichTextClipboardReadLimit
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyClipboardPolicy) GetRichTextClipboardReadSizeUnit() *string {
	return s.RichTextClipboardReadSizeUnit
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyClipboardPolicy) GetRichTextClipboardSizeUnit() *string {
	return s.RichTextClipboardSizeUnit
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyClipboardPolicy) GetRichTextClipboardWriteLimit() *int32 {
	return s.RichTextClipboardWriteLimit
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyClipboardPolicy) GetRichTextClipboardWriteSizeUnit() *string {
	return s.RichTextClipboardWriteSizeUnit
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyClipboardPolicy) GetTextClipboard() *string {
	return s.TextClipboard
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyClipboardPolicy) GetTextClipboardReadLimit() *int32 {
	return s.TextClipboardReadLimit
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyClipboardPolicy) GetTextClipboardReadSizeUnit() *string {
	return s.TextClipboardReadSizeUnit
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyClipboardPolicy) GetTextClipboardWriteLimit() *int32 {
	return s.TextClipboardWriteLimit
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyClipboardPolicy) GetTextClipboardWriteSizeUnit() *string {
	return s.TextClipboardWriteSizeUnit
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyClipboardPolicy) SetClipboard(v string) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyClipboardPolicy {
	s.Clipboard = &v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyClipboardPolicy) SetClipboardReadLimit(v int32) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyClipboardPolicy {
	s.ClipboardReadLimit = &v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyClipboardPolicy) SetClipboardScope(v string) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyClipboardPolicy {
	s.ClipboardScope = &v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyClipboardPolicy) SetClipboardSizeUnit(v string) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyClipboardPolicy {
	s.ClipboardSizeUnit = &v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyClipboardPolicy) SetClipboardWriteLimit(v int32) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyClipboardPolicy {
	s.ClipboardWriteLimit = &v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyClipboardPolicy) SetFileClipboard(v string) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyClipboardPolicy {
	s.FileClipboard = &v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyClipboardPolicy) SetRichTextClipboard(v string) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyClipboardPolicy {
	s.RichTextClipboard = &v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyClipboardPolicy) SetRichTextClipboardLimit(v int32) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyClipboardPolicy {
	s.RichTextClipboardLimit = &v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyClipboardPolicy) SetRichTextClipboardReadLimit(v int32) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyClipboardPolicy {
	s.RichTextClipboardReadLimit = &v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyClipboardPolicy) SetRichTextClipboardReadSizeUnit(v string) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyClipboardPolicy {
	s.RichTextClipboardReadSizeUnit = &v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyClipboardPolicy) SetRichTextClipboardSizeUnit(v string) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyClipboardPolicy {
	s.RichTextClipboardSizeUnit = &v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyClipboardPolicy) SetRichTextClipboardWriteLimit(v int32) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyClipboardPolicy {
	s.RichTextClipboardWriteLimit = &v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyClipboardPolicy) SetRichTextClipboardWriteSizeUnit(v string) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyClipboardPolicy {
	s.RichTextClipboardWriteSizeUnit = &v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyClipboardPolicy) SetTextClipboard(v string) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyClipboardPolicy {
	s.TextClipboard = &v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyClipboardPolicy) SetTextClipboardReadLimit(v int32) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyClipboardPolicy {
	s.TextClipboardReadLimit = &v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyClipboardPolicy) SetTextClipboardReadSizeUnit(v string) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyClipboardPolicy {
	s.TextClipboardReadSizeUnit = &v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyClipboardPolicy) SetTextClipboardWriteLimit(v int32) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyClipboardPolicy {
	s.TextClipboardWriteLimit = &v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyClipboardPolicy) SetTextClipboardWriteSizeUnit(v string) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyClipboardPolicy {
	s.TextClipboardWriteSizeUnit = &v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyClipboardPolicy) Validate() error {
	return dara.Validate(s)
}

type GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyVideoPolicy struct {
	// The frame rate of browser sessions.
	//
	// example:
	//
	// 30
	FrameRate *int32 `json:"FrameRate,omitempty" xml:"FrameRate,omitempty"`
}

func (s GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyVideoPolicy) String() string {
	return dara.Prettify(s)
}

func (s GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyVideoPolicy) GoString() string {
	return s.String()
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyVideoPolicy) GetFrameRate() *int32 {
	return s.FrameRate
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyVideoPolicy) SetFrameRate(v int32) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyVideoPolicy {
	s.FrameRate = &v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyVideoPolicy) Validate() error {
	return dara.Validate(s)
}

type GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyWatermarkPolicy struct {
	// The watermark switch. The value is case-insensitive.
	//
	// - `ON`: Watermark enabled.
	//
	// - `OFF`: Watermark disabled.
	//
	// When disabled, the watermark content type list is not used.
	//
	// example:
	//
	// ON
	WatermarkSwitch *string `json:"WatermarkSwitch,omitempty" xml:"WatermarkSwitch,omitempty"`
	// The list of watermark content types.
	//
	// - `EndUserId`: User ID.
	//
	// - `InstanceGroupId`: Delivery group ID.
	//
	// - `ClientTime`: Current time on the client.
	//
	// Use watermark types supported by the browser and client.
	WatermarkTypes []*string `json:"WatermarkTypes,omitempty" xml:"WatermarkTypes,omitempty" type:"Repeated"`
}

func (s GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyWatermarkPolicy) String() string {
	return dara.Prettify(s)
}

func (s GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyWatermarkPolicy) GoString() string {
	return s.String()
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyWatermarkPolicy) GetWatermarkSwitch() *string {
	return s.WatermarkSwitch
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyWatermarkPolicy) GetWatermarkTypes() []*string {
	return s.WatermarkTypes
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyWatermarkPolicy) SetWatermarkSwitch(v string) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyWatermarkPolicy {
	s.WatermarkSwitch = &v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyWatermarkPolicy) SetWatermarkTypes(v []*string) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyWatermarkPolicy {
	s.WatermarkTypes = v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelPolicyWatermarkPolicy) Validate() error {
	return dara.Validate(s)
}

type GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelTimers struct {
	// The session retention duration after disconnection, in minutes. A value of `-1` indicates that the session is not unbound due to this timeout, but is still subject to authorization and other session release policies.
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

func (s GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelTimers) String() string {
	return dara.Prettify(s)
}

func (s GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelTimers) GoString() string {
	return s.String()
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelTimers) GetInterval() *int32 {
	return s.Interval
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelTimers) GetTimerType() *string {
	return s.TimerType
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelTimers) SetInterval(v int32) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelTimers {
	s.Interval = &v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelTimers) SetTimerType(v string) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelTimers {
	s.TimerType = &v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelTimers) Validate() error {
	return dara.Validate(s)
}

type GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelUserLimit struct {
	// The user quota.
	//
	// example:
	//
	// 100
	UserQuota *int32 `json:"UserQuota,omitempty" xml:"UserQuota,omitempty"`
}

func (s GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelUserLimit) String() string {
	return dara.Prettify(s)
}

func (s GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelUserLimit) GoString() string {
	return s.String()
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelUserLimit) GetUserQuota() *int32 {
	return s.UserQuota
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelUserLimit) SetUserQuota(v int32) *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelUserLimit {
	s.UserQuota = &v
	return s
}

func (s *GetBrowserInstanceGroupResponseBodyBrowserInstanceGroupModelUserLimit) Validate() error {
	return dara.Validate(s)
}
