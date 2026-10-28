// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBrowserInstanceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppPackageType(v string) *CreateBrowserInstanceGroupRequest
	GetAppPackageType() *string
	SetAuthNotificationEnabled(v bool) *CreateBrowserInstanceGroupRequest
	GetAuthNotificationEnabled() *bool
	SetAutoPay(v bool) *CreateBrowserInstanceGroupRequest
	GetAutoPay() *bool
	SetAutoRenew(v bool) *CreateBrowserInstanceGroupRequest
	GetAutoRenew() *bool
	SetBizRegionId(v string) *CreateBrowserInstanceGroupRequest
	GetBizRegionId() *string
	SetBrowserConfig(v *CreateBrowserInstanceGroupRequestBrowserConfig) *CreateBrowserInstanceGroupRequest
	GetBrowserConfig() *CreateBrowserInstanceGroupRequestBrowserConfig
	SetChargeResourceMode(v string) *CreateBrowserInstanceGroupRequest
	GetChargeResourceMode() *string
	SetChargeType(v string) *CreateBrowserInstanceGroupRequest
	GetChargeType() *string
	SetCloudBrowserName(v string) *CreateBrowserInstanceGroupRequest
	GetCloudBrowserName() *string
	SetImageId(v string) *CreateBrowserInstanceGroupRequest
	GetImageId() *string
	SetInstanceType(v string) *CreateBrowserInstanceGroupRequest
	GetInstanceType() *string
	SetMaxAmount(v int32) *CreateBrowserInstanceGroupRequest
	GetMaxAmount() *int32
	SetNetwork(v *CreateBrowserInstanceGroupRequestNetwork) *CreateBrowserInstanceGroupRequest
	GetNetwork() *CreateBrowserInstanceGroupRequestNetwork
	SetNodePool(v *CreateBrowserInstanceGroupRequestNodePool) *CreateBrowserInstanceGroupRequest
	GetNodePool() *CreateBrowserInstanceGroupRequestNodePool
	SetOsType(v string) *CreateBrowserInstanceGroupRequest
	GetOsType() *string
	SetPeriod(v int32) *CreateBrowserInstanceGroupRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *CreateBrowserInstanceGroupRequest
	GetPeriodUnit() *string
	SetPolicy(v *CreateBrowserInstanceGroupRequestPolicy) *CreateBrowserInstanceGroupRequest
	GetPolicy() *CreateBrowserInstanceGroupRequestPolicy
	SetPromotionId(v string) *CreateBrowserInstanceGroupRequest
	GetPromotionId() *string
	SetSecurityPolicy(v *CreateBrowserInstanceGroupRequestSecurityPolicy) *CreateBrowserInstanceGroupRequest
	GetSecurityPolicy() *CreateBrowserInstanceGroupRequestSecurityPolicy
	SetStoragePolicy(v *CreateBrowserInstanceGroupRequestStoragePolicy) *CreateBrowserInstanceGroupRequest
	GetStoragePolicy() *CreateBrowserInstanceGroupRequestStoragePolicy
	SetSubPayType(v string) *CreateBrowserInstanceGroupRequest
	GetSubPayType() *string
	SetTag(v []*CreateBrowserInstanceGroupRequestTag) *CreateBrowserInstanceGroupRequest
	GetTag() []*CreateBrowserInstanceGroupRequestTag
	SetTimers(v []*CreateBrowserInstanceGroupRequestTimers) *CreateBrowserInstanceGroupRequest
	GetTimers() []*CreateBrowserInstanceGroupRequestTimers
	SetUserGroupIds(v []*string) *CreateBrowserInstanceGroupRequest
	GetUserGroupIds() []*string
	SetUserInfo(v *CreateBrowserInstanceGroupRequestUserInfo) *CreateBrowserInstanceGroupRequest
	GetUserInfo() *CreateBrowserInstanceGroupRequestUserInfo
	SetUsers(v []*CreateBrowserInstanceGroupRequestUsers) *CreateBrowserInstanceGroupRequest
	GetUsers() []*CreateBrowserInstanceGroupRequestUsers
}

type CreateBrowserInstanceGroupRequest struct {
	// The package identifier.
	//
	// You do not need to specify this parameter.
	//
	// example:
	//
	// -
	AppPackageType *string `json:"AppPackageType,omitempty" xml:"AppPackageType,omitempty"`
	// Specifies whether to send notification emails for authorization and deauthorization.
	//
	// - `true`: Sends notification emails.
	//
	// - `false`: Does not send notification emails.
	//
	// example:
	//
	// true
	AuthNotificationEnabled *bool `json:"AuthNotificationEnabled,omitempty" xml:"AuthNotificationEnabled,omitempty"`
	// The automatic payment parameter.
	//
	// You do not need to specify this parameter.
	//
	// example:
	//
	// -
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The auto-renewal parameter.
	//
	// You do not need to specify this parameter.
	//
	// example:
	//
	// -
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The identifier of the business region. This parameter is required.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// The browser configuration.
	BrowserConfig *CreateBrowserInstanceGroupRequestBrowserConfig `json:"BrowserConfig,omitempty" xml:"BrowserConfig,omitempty" type:"Struct"`
	// The resource billing method.
	//
	// **For MAU scenarios:*	- Set this parameter to `AppInstance`, which indicates billing by instance resources.
	//
	// example:
	//
	// AppInstance
	ChargeResourceMode *string `json:"ChargeResourceMode,omitempty" xml:"ChargeResourceMode,omitempty"`
	// The billing method.
	//
	// **For MAU scenarios:*	- Set this parameter to PostPaid, which indicates pay-as-you-go.
	//
	// example:
	//
	// PostPaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The name of the cloud browser group. This parameter cannot be empty. It is used to distinguish different browser groups in business management scenarios.
	//
	// This parameter is required.
	//
	// example:
	//
	// Business office browser
	CloudBrowserName *string `json:"CloudBrowserName,omitempty" xml:"CloudBrowserName,omitempty"`
	// The image ID used by the cloud browser. The image must be compatible with the operating system.
	//
	// If this parameter is omitted, the default image available for the account is used. If no default image is available, the creation may fail.
	//
	// **Usage conditions:*	- When CookiesSync is enabled, explicitly specify an image that supports cookie synchronization.
	//
	// example:
	//
	// img-bp13mu****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The instance type identifier. Select an instance type that matches the target region, operating system, and inventory conditions.
	//
	// If this parameter is omitted, the default instance type is used.
	//
	// example:
	//
	// appstreaming.general.basic
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The capacity configuration in MAU billing scenarios.
	//
	// example:
	//
	// 5
	MaxAmount *int32 `json:"MaxAmount,omitempty" xml:"MaxAmount,omitempty"`
	// The configurations for the office network and website access restrictions. The selected office network must belong to the current account and reside in the region specified by BizRegionId.
	Network *CreateBrowserInstanceGroupRequestNetwork `json:"Network,omitempty" xml:"Network,omitempty" type:"Struct"`
	// The node pool configuration.
	//
	// You do not need to specify this parameter.
	//
	// example:
	//
	// -
	NodePool *CreateBrowserInstanceGroupRequestNodePool `json:"NodePool,omitempty" xml:"NodePool,omitempty" type:"Struct"`
	// The operating system type. This parameter is required.
	//
	// Only `Windows` is supported. Other operating systems are not supported.
	//
	// example:
	//
	// Windows
	OsType *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	// The number of billing cycles for the purchase.
	//
	// You do not need to specify this parameter.
	//
	// example:
	//
	// -
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the billing cycle for the purchase.
	//
	// You do not need to specify this parameter.
	//
	// example:
	//
	// -
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// Configures policies for the clipboard, video, watermark, session, and client access.
	Policy *CreateBrowserInstanceGroupRequestPolicy `json:"Policy,omitempty" xml:"Policy,omitempty" type:"Struct"`
	// The promotion identifier, which is used to specify the promotional offer applied to the order.
	//
	// Applicability and specific discounts are subject to the promotion rules. You do not need to specify this parameter if no promotional offer is applied.
	//
	// example:
	//
	// 17440009****
	PromotionId *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	// The connection security policy of the browser group.
	SecurityPolicy *CreateBrowserInstanceGroupRequestSecurityPolicy `json:"SecurityPolicy,omitempty" xml:"SecurityPolicy,omitempty" type:"Struct"`
	// The user data storage configuration of the browser group.
	StoragePolicy *CreateBrowserInstanceGroupRequestStoragePolicy `json:"StoragePolicy,omitempty" xml:"StoragePolicy,omitempty" type:"Struct"`
	// The billing sub-type.
	//
	// **You must explicitly set this parameter to mau, which indicates billing by monthly active users (MAUs).*	- Omitting this parameter does not enable MAU billing.
	//
	// example:
	//
	// mau
	SubPayType *string `json:"SubPayType,omitempty" xml:"SubPayType,omitempty"`
	// This parameter is not supported and does not need to be specified.
	//
	// example:
	//
	// -
	Tag []*CreateBrowserInstanceGroupRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// This parameter is not supported and does not need to be specified.
	//
	// example:
	//
	// -
	Timers []*CreateBrowserInstanceGroupRequestTimers `json:"Timers,omitempty" xml:"Timers,omitempty" type:"Repeated"`
	// The list of authorized user group IDs. You can specify up to 10 IDs. The user groups must belong to the current account and match the account type of the office network.
	//
	// **Limits:*	- This parameter cannot be specified together with a non-empty Users parameter.
	//
	// if can be null:
	// true
	UserGroupIds []*string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty" type:"Repeated"`
	// The authorized user account information. The information must match the user and office network types.
	UserInfo *CreateBrowserInstanceGroupRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
	// The list of authorized users. You can specify up to 200 users. The users must be created in advance and match the account type.
	//
	// **Limit:*	- This parameter cannot be specified together with a non-empty UserGroupIds parameter.
	Users []*CreateBrowserInstanceGroupRequestUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s CreateBrowserInstanceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateBrowserInstanceGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateBrowserInstanceGroupRequest) GetAppPackageType() *string {
	return s.AppPackageType
}

func (s *CreateBrowserInstanceGroupRequest) GetAuthNotificationEnabled() *bool {
	return s.AuthNotificationEnabled
}

func (s *CreateBrowserInstanceGroupRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *CreateBrowserInstanceGroupRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *CreateBrowserInstanceGroupRequest) GetBizRegionId() *string {
	return s.BizRegionId
}

func (s *CreateBrowserInstanceGroupRequest) GetBrowserConfig() *CreateBrowserInstanceGroupRequestBrowserConfig {
	return s.BrowserConfig
}

func (s *CreateBrowserInstanceGroupRequest) GetChargeResourceMode() *string {
	return s.ChargeResourceMode
}

func (s *CreateBrowserInstanceGroupRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *CreateBrowserInstanceGroupRequest) GetCloudBrowserName() *string {
	return s.CloudBrowserName
}

func (s *CreateBrowserInstanceGroupRequest) GetImageId() *string {
	return s.ImageId
}

func (s *CreateBrowserInstanceGroupRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *CreateBrowserInstanceGroupRequest) GetMaxAmount() *int32 {
	return s.MaxAmount
}

func (s *CreateBrowserInstanceGroupRequest) GetNetwork() *CreateBrowserInstanceGroupRequestNetwork {
	return s.Network
}

func (s *CreateBrowserInstanceGroupRequest) GetNodePool() *CreateBrowserInstanceGroupRequestNodePool {
	return s.NodePool
}

func (s *CreateBrowserInstanceGroupRequest) GetOsType() *string {
	return s.OsType
}

func (s *CreateBrowserInstanceGroupRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *CreateBrowserInstanceGroupRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *CreateBrowserInstanceGroupRequest) GetPolicy() *CreateBrowserInstanceGroupRequestPolicy {
	return s.Policy
}

func (s *CreateBrowserInstanceGroupRequest) GetPromotionId() *string {
	return s.PromotionId
}

func (s *CreateBrowserInstanceGroupRequest) GetSecurityPolicy() *CreateBrowserInstanceGroupRequestSecurityPolicy {
	return s.SecurityPolicy
}

func (s *CreateBrowserInstanceGroupRequest) GetStoragePolicy() *CreateBrowserInstanceGroupRequestStoragePolicy {
	return s.StoragePolicy
}

func (s *CreateBrowserInstanceGroupRequest) GetSubPayType() *string {
	return s.SubPayType
}

func (s *CreateBrowserInstanceGroupRequest) GetTag() []*CreateBrowserInstanceGroupRequestTag {
	return s.Tag
}

func (s *CreateBrowserInstanceGroupRequest) GetTimers() []*CreateBrowserInstanceGroupRequestTimers {
	return s.Timers
}

func (s *CreateBrowserInstanceGroupRequest) GetUserGroupIds() []*string {
	return s.UserGroupIds
}

func (s *CreateBrowserInstanceGroupRequest) GetUserInfo() *CreateBrowserInstanceGroupRequestUserInfo {
	return s.UserInfo
}

func (s *CreateBrowserInstanceGroupRequest) GetUsers() []*CreateBrowserInstanceGroupRequestUsers {
	return s.Users
}

func (s *CreateBrowserInstanceGroupRequest) SetAppPackageType(v string) *CreateBrowserInstanceGroupRequest {
	s.AppPackageType = &v
	return s
}

func (s *CreateBrowserInstanceGroupRequest) SetAuthNotificationEnabled(v bool) *CreateBrowserInstanceGroupRequest {
	s.AuthNotificationEnabled = &v
	return s
}

func (s *CreateBrowserInstanceGroupRequest) SetAutoPay(v bool) *CreateBrowserInstanceGroupRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateBrowserInstanceGroupRequest) SetAutoRenew(v bool) *CreateBrowserInstanceGroupRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateBrowserInstanceGroupRequest) SetBizRegionId(v string) *CreateBrowserInstanceGroupRequest {
	s.BizRegionId = &v
	return s
}

func (s *CreateBrowserInstanceGroupRequest) SetBrowserConfig(v *CreateBrowserInstanceGroupRequestBrowserConfig) *CreateBrowserInstanceGroupRequest {
	s.BrowserConfig = v
	return s
}

func (s *CreateBrowserInstanceGroupRequest) SetChargeResourceMode(v string) *CreateBrowserInstanceGroupRequest {
	s.ChargeResourceMode = &v
	return s
}

func (s *CreateBrowserInstanceGroupRequest) SetChargeType(v string) *CreateBrowserInstanceGroupRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateBrowserInstanceGroupRequest) SetCloudBrowserName(v string) *CreateBrowserInstanceGroupRequest {
	s.CloudBrowserName = &v
	return s
}

func (s *CreateBrowserInstanceGroupRequest) SetImageId(v string) *CreateBrowserInstanceGroupRequest {
	s.ImageId = &v
	return s
}

func (s *CreateBrowserInstanceGroupRequest) SetInstanceType(v string) *CreateBrowserInstanceGroupRequest {
	s.InstanceType = &v
	return s
}

func (s *CreateBrowserInstanceGroupRequest) SetMaxAmount(v int32) *CreateBrowserInstanceGroupRequest {
	s.MaxAmount = &v
	return s
}

func (s *CreateBrowserInstanceGroupRequest) SetNetwork(v *CreateBrowserInstanceGroupRequestNetwork) *CreateBrowserInstanceGroupRequest {
	s.Network = v
	return s
}

func (s *CreateBrowserInstanceGroupRequest) SetNodePool(v *CreateBrowserInstanceGroupRequestNodePool) *CreateBrowserInstanceGroupRequest {
	s.NodePool = v
	return s
}

func (s *CreateBrowserInstanceGroupRequest) SetOsType(v string) *CreateBrowserInstanceGroupRequest {
	s.OsType = &v
	return s
}

func (s *CreateBrowserInstanceGroupRequest) SetPeriod(v int32) *CreateBrowserInstanceGroupRequest {
	s.Period = &v
	return s
}

func (s *CreateBrowserInstanceGroupRequest) SetPeriodUnit(v string) *CreateBrowserInstanceGroupRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateBrowserInstanceGroupRequest) SetPolicy(v *CreateBrowserInstanceGroupRequestPolicy) *CreateBrowserInstanceGroupRequest {
	s.Policy = v
	return s
}

func (s *CreateBrowserInstanceGroupRequest) SetPromotionId(v string) *CreateBrowserInstanceGroupRequest {
	s.PromotionId = &v
	return s
}

func (s *CreateBrowserInstanceGroupRequest) SetSecurityPolicy(v *CreateBrowserInstanceGroupRequestSecurityPolicy) *CreateBrowserInstanceGroupRequest {
	s.SecurityPolicy = v
	return s
}

func (s *CreateBrowserInstanceGroupRequest) SetStoragePolicy(v *CreateBrowserInstanceGroupRequestStoragePolicy) *CreateBrowserInstanceGroupRequest {
	s.StoragePolicy = v
	return s
}

func (s *CreateBrowserInstanceGroupRequest) SetSubPayType(v string) *CreateBrowserInstanceGroupRequest {
	s.SubPayType = &v
	return s
}

func (s *CreateBrowserInstanceGroupRequest) SetTag(v []*CreateBrowserInstanceGroupRequestTag) *CreateBrowserInstanceGroupRequest {
	s.Tag = v
	return s
}

func (s *CreateBrowserInstanceGroupRequest) SetTimers(v []*CreateBrowserInstanceGroupRequestTimers) *CreateBrowserInstanceGroupRequest {
	s.Timers = v
	return s
}

func (s *CreateBrowserInstanceGroupRequest) SetUserGroupIds(v []*string) *CreateBrowserInstanceGroupRequest {
	s.UserGroupIds = v
	return s
}

func (s *CreateBrowserInstanceGroupRequest) SetUserInfo(v *CreateBrowserInstanceGroupRequestUserInfo) *CreateBrowserInstanceGroupRequest {
	s.UserInfo = v
	return s
}

func (s *CreateBrowserInstanceGroupRequest) SetUsers(v []*CreateBrowserInstanceGroupRequestUsers) *CreateBrowserInstanceGroupRequest {
	s.Users = v
	return s
}

func (s *CreateBrowserInstanceGroupRequest) Validate() error {
	if s.BrowserConfig != nil {
		if err := s.BrowserConfig.Validate(); err != nil {
			return err
		}
	}
	if s.Network != nil {
		if err := s.Network.Validate(); err != nil {
			return err
		}
	}
	if s.NodePool != nil {
		if err := s.NodePool.Validate(); err != nil {
			return err
		}
	}
	if s.Policy != nil {
		if err := s.Policy.Validate(); err != nil {
			return err
		}
	}
	if s.SecurityPolicy != nil {
		if err := s.SecurityPolicy.Validate(); err != nil {
			return err
		}
	}
	if s.StoragePolicy != nil {
		if err := s.StoragePolicy.Validate(); err != nil {
			return err
		}
	}
	if s.Tag != nil {
		for _, item := range s.Tag {
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
	if s.UserInfo != nil {
		if err := s.UserInfo.Validate(); err != nil {
			return err
		}
	}
	if s.Users != nil {
		for _, item := range s.Users {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateBrowserInstanceGroupRequestBrowserConfig struct {
	// The list of browser bookmarks.
	//
	// **Limits:*	- This parameter cannot be specified together with a non-empty BookmarksFilePath parameter, and bookmark URLs cannot be duplicated.
	Bookmarks []*CreateBrowserInstanceGroupRequestBrowserConfigBookmarks `json:"Bookmarks,omitempty" xml:"Bookmarks,omitempty" type:"Repeated"`
	// The path of the uploaded bookmark file. This parameter cannot be specified together with a non-empty Bookmarks parameter.
	//
	// **File format:*	- A CSV file without a header. Specify the following four columns in order:
	//
	// 1. Bookmark name.
	//
	// 2. URL.
	//
	// 3. Folder.
	//
	// 4. Root directory type: `bookmark_bar` indicates the bookmarks bar, and `other` indicates other bookmarks.
	//
	// **Limits:**
	//
	// - Fields are separated by commas. Field content cannot contain commas or line breaks, and quote escaping is not supported.
	//
	// - The file path must belong to the upload directory specified for the current account and cannot contain `..`.
	//
	// example:
	//
	// cn-hangzhou/aig_upm/xxx/temp/BrowserBookmarks/BrowserBookmarkTemplate.csv
	BookmarksFilePath *string `json:"BookmarksFilePath,omitempty" xml:"BookmarksFilePath,omitempty"`
	// The browser startup parameters. For example, --incognito indicates that an incognito window is used.
	//
	// example:
	//
	// --incognito
	BrowserParam *string `json:"BrowserParam,omitempty" xml:"BrowserParam,omitempty"`
	// Specifies whether to synchronize cookies.
	//
	// example:
	//
	// false
	CookiesSync *bool `json:"CookiesSync,omitempty" xml:"CookiesSync,omitempty"`
	// The homepage URL opened when the browser starts. The URL must conform to the URI syntax.
	//
	// example:
	//
	// https://www.aliyun.com
	Homepage *string `json:"Homepage,omitempty" xml:"Homepage,omitempty"`
}

func (s CreateBrowserInstanceGroupRequestBrowserConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateBrowserInstanceGroupRequestBrowserConfig) GoString() string {
	return s.String()
}

func (s *CreateBrowserInstanceGroupRequestBrowserConfig) GetBookmarks() []*CreateBrowserInstanceGroupRequestBrowserConfigBookmarks {
	return s.Bookmarks
}

func (s *CreateBrowserInstanceGroupRequestBrowserConfig) GetBookmarksFilePath() *string {
	return s.BookmarksFilePath
}

func (s *CreateBrowserInstanceGroupRequestBrowserConfig) GetBrowserParam() *string {
	return s.BrowserParam
}

func (s *CreateBrowserInstanceGroupRequestBrowserConfig) GetCookiesSync() *bool {
	return s.CookiesSync
}

func (s *CreateBrowserInstanceGroupRequestBrowserConfig) GetHomepage() *string {
	return s.Homepage
}

func (s *CreateBrowserInstanceGroupRequestBrowserConfig) SetBookmarks(v []*CreateBrowserInstanceGroupRequestBrowserConfigBookmarks) *CreateBrowserInstanceGroupRequestBrowserConfig {
	s.Bookmarks = v
	return s
}

func (s *CreateBrowserInstanceGroupRequestBrowserConfig) SetBookmarksFilePath(v string) *CreateBrowserInstanceGroupRequestBrowserConfig {
	s.BookmarksFilePath = &v
	return s
}

func (s *CreateBrowserInstanceGroupRequestBrowserConfig) SetBrowserParam(v string) *CreateBrowserInstanceGroupRequestBrowserConfig {
	s.BrowserParam = &v
	return s
}

func (s *CreateBrowserInstanceGroupRequestBrowserConfig) SetCookiesSync(v bool) *CreateBrowserInstanceGroupRequestBrowserConfig {
	s.CookiesSync = &v
	return s
}

func (s *CreateBrowserInstanceGroupRequestBrowserConfig) SetHomepage(v string) *CreateBrowserInstanceGroupRequestBrowserConfig {
	s.Homepage = &v
	return s
}

func (s *CreateBrowserInstanceGroupRequestBrowserConfig) Validate() error {
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

type CreateBrowserInstanceGroupRequestBrowserConfigBookmarks struct {
	// The folder where the bookmark is located. The length cannot exceed 64 characters after leading and trailing spaces are removed.
	//
	// example:
	//
	// Work portal
	BookmarkFolder *string `json:"BookmarkFolder,omitempty" xml:"BookmarkFolder,omitempty"`
	// The bookmark name. This parameter is required and cannot be empty when you create a bookmark. The length cannot exceed 64 characters after leading and trailing spaces are removed.
	//
	// This parameter is required.
	//
	// example:
	//
	// Alibaba Cloud website
	BookmarkName *string `json:"BookmarkName,omitempty" xml:"BookmarkName,omitempty"`
	// The URL of the bookmark. This parameter is required when you create a bookmark. The length cannot exceed 1024 characters after leading and trailing spaces are removed.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://www.aliyun.com
	BookmarkURL *string `json:"BookmarkURL,omitempty" xml:"BookmarkURL,omitempty"`
}

func (s CreateBrowserInstanceGroupRequestBrowserConfigBookmarks) String() string {
	return dara.Prettify(s)
}

func (s CreateBrowserInstanceGroupRequestBrowserConfigBookmarks) GoString() string {
	return s.String()
}

func (s *CreateBrowserInstanceGroupRequestBrowserConfigBookmarks) GetBookmarkFolder() *string {
	return s.BookmarkFolder
}

func (s *CreateBrowserInstanceGroupRequestBrowserConfigBookmarks) GetBookmarkName() *string {
	return s.BookmarkName
}

func (s *CreateBrowserInstanceGroupRequestBrowserConfigBookmarks) GetBookmarkURL() *string {
	return s.BookmarkURL
}

func (s *CreateBrowserInstanceGroupRequestBrowserConfigBookmarks) SetBookmarkFolder(v string) *CreateBrowserInstanceGroupRequestBrowserConfigBookmarks {
	s.BookmarkFolder = &v
	return s
}

func (s *CreateBrowserInstanceGroupRequestBrowserConfigBookmarks) SetBookmarkName(v string) *CreateBrowserInstanceGroupRequestBrowserConfigBookmarks {
	s.BookmarkName = &v
	return s
}

func (s *CreateBrowserInstanceGroupRequestBrowserConfigBookmarks) SetBookmarkURL(v string) *CreateBrowserInstanceGroupRequestBrowserConfigBookmarks {
	s.BookmarkURL = &v
	return s
}

func (s *CreateBrowserInstanceGroupRequestBrowserConfigBookmarks) Validate() error {
	return dara.Validate(s)
}

type CreateBrowserInstanceGroupRequestNetwork struct {
	// The website access restriction mode.
	//
	// - ALLOW_ALL: Allows access to all domain names.
	//
	// - ALLOW_LIST: Allows access only to websites in the whitelist.
	//
	// example:
	//
	// ALLOW_ALL
	AccessRestriction *string `json:"AccessRestriction,omitempty" xml:"AccessRestriction,omitempty"`
	// The identifier of the created office network. The office network must belong to the current account and reside in the target region specified by BizRegionId. You can view the office network ID on the Network & Storage > Office Networks page in the Elastic Desktop Service Enterprise Edition console, or call the DescribeOfficeSites operation of Elastic Desktop Service Enterprise Edition to query the ID.
	//
	// example:
	//
	// cn-hangzhou+dir-643067****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The list of websites that are allowed for access. This parameter is used in ALLOW_LIST mode.
	//
	// **Limits:**
	//
	// - You can specify up to 20 URLs directly. To specify more than 20 URLs, use RestrictedURLsFilePath to import them from a file.
	//
	// - This parameter cannot be specified together with RestrictedURLsFilePath.
	//
	// - The URLs in the list cannot be duplicated.
	RestrictedURLs []*CreateBrowserInstanceGroupRequestNetworkRestrictedURLs `json:"RestrictedURLs,omitempty" xml:"RestrictedURLs,omitempty" type:"Repeated"`
	// The path of the uploaded website whitelist file. This parameter is used in ALLOW_LIST mode.
	//
	// If you need to specify more than 20 URLs, import them from a file. By default, you can configure up to 1,000 URLs.
	//
	// This parameter cannot be specified together with RestrictedURLs.
	//
	// example:
	//
	// cn-hangzhou/aig_upm/xxx/temp/BrowserRestrictionUrls/URL_whitelist_template.csv
	RestrictedURLsFilePath *string `json:"RestrictedURLsFilePath,omitempty" xml:"RestrictedURLsFilePath,omitempty"`
	// The list of vSwitch identifiers.
	//
	// **Usage condition:*	- This parameter is required only when you use a custom office network. Do not specify this parameter for other types of office networks.
	//
	// Select vSwitches that match the target business region and the custom office network.
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
}

func (s CreateBrowserInstanceGroupRequestNetwork) String() string {
	return dara.Prettify(s)
}

func (s CreateBrowserInstanceGroupRequestNetwork) GoString() string {
	return s.String()
}

func (s *CreateBrowserInstanceGroupRequestNetwork) GetAccessRestriction() *string {
	return s.AccessRestriction
}

func (s *CreateBrowserInstanceGroupRequestNetwork) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *CreateBrowserInstanceGroupRequestNetwork) GetRestrictedURLs() []*CreateBrowserInstanceGroupRequestNetworkRestrictedURLs {
	return s.RestrictedURLs
}

func (s *CreateBrowserInstanceGroupRequestNetwork) GetRestrictedURLsFilePath() *string {
	return s.RestrictedURLsFilePath
}

func (s *CreateBrowserInstanceGroupRequestNetwork) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *CreateBrowserInstanceGroupRequestNetwork) SetAccessRestriction(v string) *CreateBrowserInstanceGroupRequestNetwork {
	s.AccessRestriction = &v
	return s
}

func (s *CreateBrowserInstanceGroupRequestNetwork) SetOfficeSiteId(v string) *CreateBrowserInstanceGroupRequestNetwork {
	s.OfficeSiteId = &v
	return s
}

func (s *CreateBrowserInstanceGroupRequestNetwork) SetRestrictedURLs(v []*CreateBrowserInstanceGroupRequestNetworkRestrictedURLs) *CreateBrowserInstanceGroupRequestNetwork {
	s.RestrictedURLs = v
	return s
}

func (s *CreateBrowserInstanceGroupRequestNetwork) SetRestrictedURLsFilePath(v string) *CreateBrowserInstanceGroupRequestNetwork {
	s.RestrictedURLsFilePath = &v
	return s
}

func (s *CreateBrowserInstanceGroupRequestNetwork) SetVSwitchIds(v []*string) *CreateBrowserInstanceGroupRequestNetwork {
	s.VSwitchIds = v
	return s
}

func (s *CreateBrowserInstanceGroupRequestNetwork) Validate() error {
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

type CreateBrowserInstanceGroupRequestNetworkRestrictedURLs struct {
	// The URL of the website that is allowed for access.
	//
	// example:
	//
	// aliyun.com
	URL *string `json:"URL,omitempty" xml:"URL,omitempty"`
}

func (s CreateBrowserInstanceGroupRequestNetworkRestrictedURLs) String() string {
	return dara.Prettify(s)
}

func (s CreateBrowserInstanceGroupRequestNetworkRestrictedURLs) GoString() string {
	return s.String()
}

func (s *CreateBrowserInstanceGroupRequestNetworkRestrictedURLs) GetURL() *string {
	return s.URL
}

func (s *CreateBrowserInstanceGroupRequestNetworkRestrictedURLs) SetURL(v string) *CreateBrowserInstanceGroupRequestNetworkRestrictedURLs {
	s.URL = &v
	return s
}

func (s *CreateBrowserInstanceGroupRequestNetworkRestrictedURLs) Validate() error {
	return dara.Validate(s)
}

type CreateBrowserInstanceGroupRequestNodePool struct {
	// The number of nodes.
	//
	// You do not need to specify this parameter.
	//
	// example:
	//
	// -
	NodeAmount *int32 `json:"NodeAmount,omitempty" xml:"NodeAmount,omitempty"`
	// The identity of the node specifications.
	//
	// You do not need to specify this parameter.
	//
	// example:
	//
	// -
	NodeInstanceType *string `json:"NodeInstanceType,omitempty" xml:"NodeInstanceType,omitempty"`
	// The node scaling policy.
	//
	// You do not need to specify this parameter.
	//
	// example:
	//
	// -
	StrategyType *string `json:"StrategyType,omitempty" xml:"StrategyType,omitempty"`
}

func (s CreateBrowserInstanceGroupRequestNodePool) String() string {
	return dara.Prettify(s)
}

func (s CreateBrowserInstanceGroupRequestNodePool) GoString() string {
	return s.String()
}

func (s *CreateBrowserInstanceGroupRequestNodePool) GetNodeAmount() *int32 {
	return s.NodeAmount
}

func (s *CreateBrowserInstanceGroupRequestNodePool) GetNodeInstanceType() *string {
	return s.NodeInstanceType
}

func (s *CreateBrowserInstanceGroupRequestNodePool) GetStrategyType() *string {
	return s.StrategyType
}

func (s *CreateBrowserInstanceGroupRequestNodePool) SetNodeAmount(v int32) *CreateBrowserInstanceGroupRequestNodePool {
	s.NodeAmount = &v
	return s
}

func (s *CreateBrowserInstanceGroupRequestNodePool) SetNodeInstanceType(v string) *CreateBrowserInstanceGroupRequestNodePool {
	s.NodeInstanceType = &v
	return s
}

func (s *CreateBrowserInstanceGroupRequestNodePool) SetStrategyType(v string) *CreateBrowserInstanceGroupRequestNodePool {
	s.StrategyType = &v
	return s
}

func (s *CreateBrowserInstanceGroupRequestNodePool) Validate() error {
	return dara.Validate(s)
}

type CreateBrowserInstanceGroupRequestPolicy struct {
	// The screen capture prevention switch.
	//
	// - on: Enables screen capture prevention.
	//
	// - off: Disables screen capture prevention.
	//
	// example:
	//
	// on
	AppContentProtection *string `json:"AppContentProtection,omitempty" xml:"AppContentProtection,omitempty"`
	// The client access IP address whitelist, which is used to restrict the source IP addresses of clients that can access the cloud browser.
	AuthorizeAccessPolicyRules []*CreateBrowserInstanceGroupRequestPolicyAuthorizeAccessPolicyRules `json:"AuthorizeAccessPolicyRules,omitempty" xml:"AuthorizeAccessPolicyRules,omitempty" type:"Repeated"`
	// The client access control list.
	ClientTypes []*CreateBrowserInstanceGroupRequestPolicyClientTypes `json:"ClientTypes,omitempty" xml:"ClientTypes,omitempty" type:"Repeated"`
	// Specifies the clipboard transfer direction, content types, and size limits. read indicates transfer from the local PC to the cloud browser, and write indicates transfer from the cloud browser to the local PC.
	ClipboardPolicy *CreateBrowserInstanceGroupRequestPolicyClipboardPolicy `json:"ClipboardPolicy,omitempty" xml:"ClipboardPolicy,omitempty" type:"Struct"`
	// Specifies the data retention policy for sessions after disconnection.
	//
	// - `customTime`: Retains the session according to the data retention policy. Custom duration through `DisconnectKeepSessionTime` is not supported.
	//
	// - `persistent`: The session is not subject to automatic release based on disconnection duration.
	//
	// **Note:*	- `persistent` is still subject to authorization and other release policies.
	//
	// example:
	//
	// customTime
	DisconnectKeepSession *string `json:"DisconnectKeepSession,omitempty" xml:"DisconnectKeepSession,omitempty"`
	// This parameter is not supported and does not need to be specified.
	//
	// example:
	//
	// -
	DisconnectKeepSessionTime *int32 `json:"DisconnectKeepSessionTime,omitempty" xml:"DisconnectKeepSessionTime,omitempty"`
	// Specifies whether to enable the file manager floating button.
	//
	// - `on`: Enabled.
	//
	// - `off`: Disabled.
	//
	// **Default value:*	- `off`.
	//
	// example:
	//
	// off
	FileManager *string `json:"FileManager,omitempty" xml:"FileManager,omitempty"`
	// Specifies the file transfer policy for the web client.
	//
	// - `off`: Prohibits file transfer.
	//
	// - `upload`: Allows upload only.
	//
	// - `download`: Allows download only.
	//
	// - `full`: Allows both upload and download.
	//
	// Configure this parameter in conjunction with the clipboard policy.
	//
	// example:
	//
	// full
	Html5FileTransfer *string `json:"Html5FileTransfer,omitempty" xml:"Html5FileTransfer,omitempty"`
	// Specifies whether to enable the dual-screen mode for the cloud browser.
	//
	// example:
	//
	// off
	MultiScreen *string `json:"MultiScreen,omitempty" xml:"MultiScreen,omitempty"`
	// Specifies whether to enable automatic disconnection on inactivity. The value is not case-sensitive.
	//
	// - `on`: Enabled.
	//
	// - `off`: Disabled.
	//
	// When enabled, use `NoOperationDisconnectTime` to set the idle timeout.
	//
	// example:
	//
	// on
	NoOperationDisconnect *string `json:"NoOperationDisconnect,omitempty" xml:"NoOperationDisconnect,omitempty"`
	// Specifies the idle timeout in seconds before automatic disconnection is triggered.
	//
	// **Prerequisite:*	- When `NoOperationDisconnect` is enabled, you must specify a value greater than 0.
	//
	// example:
	//
	// 300
	NoOperationDisconnectTime *int32 `json:"NoOperationDisconnectTime,omitempty" xml:"NoOperationDisconnectTime,omitempty"`
	// Specifies the policy version. The value is not case-sensitive.
	//
	// - `DEFAULT`: Legacy policy.
	//
	// - `CENTER`: Centralized policy.
	//
	// **Default value:*	- `DEFAULT`. The effective policy version is subject to the policy configuration available for the account.
	//
	// example:
	//
	// CENTER
	PolicyVersion *string `json:"PolicyVersion,omitempty" xml:"PolicyVersion,omitempty"`
	// Specifies the video display policy for the browser session.
	VideoPolicy *CreateBrowserInstanceGroupRequestPolicyVideoPolicy `json:"VideoPolicy,omitempty" xml:"VideoPolicy,omitempty" type:"Struct"`
	// The watermark display configuration for browser sessions.
	WatermarkPolicy *CreateBrowserInstanceGroupRequestPolicyWatermarkPolicy `json:"WatermarkPolicy,omitempty" xml:"WatermarkPolicy,omitempty" type:"Struct"`
}

func (s CreateBrowserInstanceGroupRequestPolicy) String() string {
	return dara.Prettify(s)
}

func (s CreateBrowserInstanceGroupRequestPolicy) GoString() string {
	return s.String()
}

func (s *CreateBrowserInstanceGroupRequestPolicy) GetAppContentProtection() *string {
	return s.AppContentProtection
}

func (s *CreateBrowserInstanceGroupRequestPolicy) GetAuthorizeAccessPolicyRules() []*CreateBrowserInstanceGroupRequestPolicyAuthorizeAccessPolicyRules {
	return s.AuthorizeAccessPolicyRules
}

func (s *CreateBrowserInstanceGroupRequestPolicy) GetClientTypes() []*CreateBrowserInstanceGroupRequestPolicyClientTypes {
	return s.ClientTypes
}

func (s *CreateBrowserInstanceGroupRequestPolicy) GetClipboardPolicy() *CreateBrowserInstanceGroupRequestPolicyClipboardPolicy {
	return s.ClipboardPolicy
}

func (s *CreateBrowserInstanceGroupRequestPolicy) GetDisconnectKeepSession() *string {
	return s.DisconnectKeepSession
}

func (s *CreateBrowserInstanceGroupRequestPolicy) GetDisconnectKeepSessionTime() *int32 {
	return s.DisconnectKeepSessionTime
}

func (s *CreateBrowserInstanceGroupRequestPolicy) GetFileManager() *string {
	return s.FileManager
}

func (s *CreateBrowserInstanceGroupRequestPolicy) GetHtml5FileTransfer() *string {
	return s.Html5FileTransfer
}

func (s *CreateBrowserInstanceGroupRequestPolicy) GetMultiScreen() *string {
	return s.MultiScreen
}

func (s *CreateBrowserInstanceGroupRequestPolicy) GetNoOperationDisconnect() *string {
	return s.NoOperationDisconnect
}

func (s *CreateBrowserInstanceGroupRequestPolicy) GetNoOperationDisconnectTime() *int32 {
	return s.NoOperationDisconnectTime
}

func (s *CreateBrowserInstanceGroupRequestPolicy) GetPolicyVersion() *string {
	return s.PolicyVersion
}

func (s *CreateBrowserInstanceGroupRequestPolicy) GetVideoPolicy() *CreateBrowserInstanceGroupRequestPolicyVideoPolicy {
	return s.VideoPolicy
}

func (s *CreateBrowserInstanceGroupRequestPolicy) GetWatermarkPolicy() *CreateBrowserInstanceGroupRequestPolicyWatermarkPolicy {
	return s.WatermarkPolicy
}

func (s *CreateBrowserInstanceGroupRequestPolicy) SetAppContentProtection(v string) *CreateBrowserInstanceGroupRequestPolicy {
	s.AppContentProtection = &v
	return s
}

func (s *CreateBrowserInstanceGroupRequestPolicy) SetAuthorizeAccessPolicyRules(v []*CreateBrowserInstanceGroupRequestPolicyAuthorizeAccessPolicyRules) *CreateBrowserInstanceGroupRequestPolicy {
	s.AuthorizeAccessPolicyRules = v
	return s
}

func (s *CreateBrowserInstanceGroupRequestPolicy) SetClientTypes(v []*CreateBrowserInstanceGroupRequestPolicyClientTypes) *CreateBrowserInstanceGroupRequestPolicy {
	s.ClientTypes = v
	return s
}

func (s *CreateBrowserInstanceGroupRequestPolicy) SetClipboardPolicy(v *CreateBrowserInstanceGroupRequestPolicyClipboardPolicy) *CreateBrowserInstanceGroupRequestPolicy {
	s.ClipboardPolicy = v
	return s
}

func (s *CreateBrowserInstanceGroupRequestPolicy) SetDisconnectKeepSession(v string) *CreateBrowserInstanceGroupRequestPolicy {
	s.DisconnectKeepSession = &v
	return s
}

func (s *CreateBrowserInstanceGroupRequestPolicy) SetDisconnectKeepSessionTime(v int32) *CreateBrowserInstanceGroupRequestPolicy {
	s.DisconnectKeepSessionTime = &v
	return s
}

func (s *CreateBrowserInstanceGroupRequestPolicy) SetFileManager(v string) *CreateBrowserInstanceGroupRequestPolicy {
	s.FileManager = &v
	return s
}

func (s *CreateBrowserInstanceGroupRequestPolicy) SetHtml5FileTransfer(v string) *CreateBrowserInstanceGroupRequestPolicy {
	s.Html5FileTransfer = &v
	return s
}

func (s *CreateBrowserInstanceGroupRequestPolicy) SetMultiScreen(v string) *CreateBrowserInstanceGroupRequestPolicy {
	s.MultiScreen = &v
	return s
}

func (s *CreateBrowserInstanceGroupRequestPolicy) SetNoOperationDisconnect(v string) *CreateBrowserInstanceGroupRequestPolicy {
	s.NoOperationDisconnect = &v
	return s
}

func (s *CreateBrowserInstanceGroupRequestPolicy) SetNoOperationDisconnectTime(v int32) *CreateBrowserInstanceGroupRequestPolicy {
	s.NoOperationDisconnectTime = &v
	return s
}

func (s *CreateBrowserInstanceGroupRequestPolicy) SetPolicyVersion(v string) *CreateBrowserInstanceGroupRequestPolicy {
	s.PolicyVersion = &v
	return s
}

func (s *CreateBrowserInstanceGroupRequestPolicy) SetVideoPolicy(v *CreateBrowserInstanceGroupRequestPolicyVideoPolicy) *CreateBrowserInstanceGroupRequestPolicy {
	s.VideoPolicy = v
	return s
}

func (s *CreateBrowserInstanceGroupRequestPolicy) SetWatermarkPolicy(v *CreateBrowserInstanceGroupRequestPolicyWatermarkPolicy) *CreateBrowserInstanceGroupRequestPolicy {
	s.WatermarkPolicy = v
	return s
}

func (s *CreateBrowserInstanceGroupRequestPolicy) Validate() error {
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

type CreateBrowserInstanceGroupRequestPolicyAuthorizeAccessPolicyRules struct {
	// The source CIDR block of the client that is allowed to access the cloud browser. Use the CIDR format.
	//
	// example:
	//
	// 192.168.1.0/24
	CidrIp *string `json:"CidrIp,omitempty" xml:"CidrIp,omitempty"`
	// The description of the client access IP address whitelist rule.
	//
	// example:
	//
	// Office network access
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s CreateBrowserInstanceGroupRequestPolicyAuthorizeAccessPolicyRules) String() string {
	return dara.Prettify(s)
}

func (s CreateBrowserInstanceGroupRequestPolicyAuthorizeAccessPolicyRules) GoString() string {
	return s.String()
}

func (s *CreateBrowserInstanceGroupRequestPolicyAuthorizeAccessPolicyRules) GetCidrIp() *string {
	return s.CidrIp
}

func (s *CreateBrowserInstanceGroupRequestPolicyAuthorizeAccessPolicyRules) GetDescription() *string {
	return s.Description
}

func (s *CreateBrowserInstanceGroupRequestPolicyAuthorizeAccessPolicyRules) SetCidrIp(v string) *CreateBrowserInstanceGroupRequestPolicyAuthorizeAccessPolicyRules {
	s.CidrIp = &v
	return s
}

func (s *CreateBrowserInstanceGroupRequestPolicyAuthorizeAccessPolicyRules) SetDescription(v string) *CreateBrowserInstanceGroupRequestPolicyAuthorizeAccessPolicyRules {
	s.Description = &v
	return s
}

func (s *CreateBrowserInstanceGroupRequestPolicyAuthorizeAccessPolicyRules) Validate() error {
	return dara.Validate(s)
}

type CreateBrowserInstanceGroupRequestPolicyClientTypes struct {
	// The type of the client for which you want to configure an access policy.
	//
	// - windows: Windows client.
	//
	// - macos: macOS client.
	//
	// - html5: Web client.
	//
	// - android: Android client.
	//
	// - ios: iOS client.
	//
	// example:
	//
	// windows
	ClientType *string `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	// The access policy switch for the client type.
	//
	// - on: Allows access from this type of client.
	//
	// - off: Denies access from this type of client.
	//
	// example:
	//
	// on
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateBrowserInstanceGroupRequestPolicyClientTypes) String() string {
	return dara.Prettify(s)
}

func (s CreateBrowserInstanceGroupRequestPolicyClientTypes) GoString() string {
	return s.String()
}

func (s *CreateBrowserInstanceGroupRequestPolicyClientTypes) GetClientType() *string {
	return s.ClientType
}

func (s *CreateBrowserInstanceGroupRequestPolicyClientTypes) GetStatus() *string {
	return s.Status
}

func (s *CreateBrowserInstanceGroupRequestPolicyClientTypes) SetClientType(v string) *CreateBrowserInstanceGroupRequestPolicyClientTypes {
	s.ClientType = &v
	return s
}

func (s *CreateBrowserInstanceGroupRequestPolicyClientTypes) SetStatus(v string) *CreateBrowserInstanceGroupRequestPolicyClientTypes {
	s.Status = &v
	return s
}

func (s *CreateBrowserInstanceGroupRequestPolicyClientTypes) Validate() error {
	return dara.Validate(s)
}

type CreateBrowserInstanceGroupRequestPolicyClipboardPolicy struct {
	// Specifies the clipboard transfer direction. The value is not case-sensitive.
	//
	// - `off`: Prohibits bidirectional transfer.
	//
	// - `read`: Allows copy and paste from the local PC to the cloud browser.
	//
	// - `write`: Allows copy and paste from the cloud browser to the local PC.
	//
	// - `readwrite`: Allows bidirectional copy and paste.
	//
	// example:
	//
	// readwrite
	Clipboard *string `json:"Clipboard,omitempty" xml:"Clipboard,omitempty"`
	// Specifies the clipboard size limit for inbound transfer (from the local PC to the cloud browser).
	//
	// **Value range:*	- 1 to 102400. The unit is specified by ClipboardSizeUnit.
	//
	// This value range does not change with unit conversion.
	//
	// example:
	//
	// 1024
	ClipboardReadLimit *int32 `json:"ClipboardReadLimit,omitempty" xml:"ClipboardReadLimit,omitempty"`
	// Specifies the clipboard control granularity.
	//
	// - `global`: Unified control.
	//
	// - `grained`: Separate control for text, rich text, and files.
	//
	// example:
	//
	// grained
	ClipboardScope *string `json:"ClipboardScope,omitempty" xml:"ClipboardScope,omitempty"`
	// Specifies the clipboard size unit.
	//
	// - `B`: Bytes.
	//
	// - `KB`: 1024 bytes.
	//
	// example:
	//
	// KB
	ClipboardSizeUnit *string `json:"ClipboardSizeUnit,omitempty" xml:"ClipboardSizeUnit,omitempty"`
	// Specifies the clipboard size limit for outbound transfer (from the cloud browser to the local PC).
	//
	// **Value range:*	- 1 to 102400. The unit is specified by ClipboardSizeUnit.
	//
	// This value range does not change with unit conversion.
	//
	// example:
	//
	// 1024
	ClipboardWriteLimit *int32 `json:"ClipboardWriteLimit,omitempty" xml:"ClipboardWriteLimit,omitempty"`
	// Specifies the clipboard transfer direction. The value is not case-sensitive.
	//
	// - `off`: Prohibits bidirectional transfer.
	//
	// - `read`: Allows copy and paste from the local PC to the cloud browser.
	//
	// - `write`: Allows copy and paste from the cloud browser to the local PC.
	//
	// - `readwrite`: Allows bidirectional copy and paste.
	//
	// example:
	//
	// off
	FileClipboard *string `json:"FileClipboard,omitempty" xml:"FileClipboard,omitempty"`
	// Specifies the clipboard transfer direction. The value is not case-sensitive.
	//
	// - `off`: Prohibits bidirectional transfer.
	//
	// - `read`: Allows copy and paste from the local PC to the cloud browser.
	//
	// - `write`: Allows copy and paste from the cloud browser to the local PC.
	//
	// - `readwrite`: Allows bidirectional copy and paste.
	//
	// example:
	//
	// off
	RichTextClipboard *string `json:"RichTextClipboard,omitempty" xml:"RichTextClipboard,omitempty"`
	// Specifies the clipboard size limit for rich text.
	//
	// **Value range:*	- 1 to 204800. The unit is specified by RichTextClipboardSizeUnit.
	//
	// This value range does not change with unit conversion.
	//
	// example:
	//
	// 1024
	RichTextClipboardLimit *int32 `json:"RichTextClipboardLimit,omitempty" xml:"RichTextClipboardLimit,omitempty"`
	// Specifies the clipboard size limit for inbound transfer (from the local PC to the cloud browser).
	//
	// **Value range:*	- 1 to 204800. The unit is specified by RichTextClipboardReadSizeUnit.
	//
	// This value range does not change with unit conversion.
	//
	// example:
	//
	// 1
	RichTextClipboardReadLimit *int32 `json:"RichTextClipboardReadLimit,omitempty" xml:"RichTextClipboardReadLimit,omitempty"`
	// Specifies the clipboard size unit.
	//
	// - `B`: Bytes.
	//
	// - `KB`: 1024 bytes.
	//
	// **Default value:*	- `KB`.
	//
	// example:
	//
	// KB
	RichTextClipboardReadSizeUnit *string `json:"RichTextClipboardReadSizeUnit,omitempty" xml:"RichTextClipboardReadSizeUnit,omitempty"`
	// Specifies the clipboard size unit.
	//
	// - `B`: Bytes.
	//
	// - `KB`: 1024 bytes.
	//
	// example:
	//
	// KB
	RichTextClipboardSizeUnit *string `json:"RichTextClipboardSizeUnit,omitempty" xml:"RichTextClipboardSizeUnit,omitempty"`
	// Specifies the clipboard size limit for outbound transfer (from the cloud browser to the local PC).
	//
	// **Value range:*	- 1 to 204800. The unit is specified by RichTextClipboardWriteSizeUnit.
	//
	// This value range does not change with unit conversion.
	//
	// example:
	//
	// 1
	RichTextClipboardWriteLimit *int32 `json:"RichTextClipboardWriteLimit,omitempty" xml:"RichTextClipboardWriteLimit,omitempty"`
	// Specifies the clipboard size unit.
	//
	// - `B`: Bytes.
	//
	// - `KB`: 1024 bytes.
	//
	// **Default value:*	- `KB`.
	//
	// example:
	//
	// KB
	RichTextClipboardWriteSizeUnit *string `json:"RichTextClipboardWriteSizeUnit,omitempty" xml:"RichTextClipboardWriteSizeUnit,omitempty"`
	// Specifies the clipboard transfer direction. The value is not case-sensitive.
	//
	// - `off`: Prohibits bidirectional transfer.
	//
	// - `read`: Allows copy and paste from the local PC to the cloud browser.
	//
	// - `write`: Allows copy and paste from the cloud browser to the local PC.
	//
	// - `readwrite`: Allows bidirectional copy and paste.
	//
	// example:
	//
	// readwrite
	TextClipboard *string `json:"TextClipboard,omitempty" xml:"TextClipboard,omitempty"`
	// Specifies the clipboard size limit for inbound transfer (from the local PC to the cloud browser).
	//
	// **Value range:*	- 1 to 102400. The unit is specified by TextClipboardReadSizeUnit.
	//
	// This value range does not change with unit conversion.
	//
	// example:
	//
	// 1
	TextClipboardReadLimit *int32 `json:"TextClipboardReadLimit,omitempty" xml:"TextClipboardReadLimit,omitempty"`
	// Specifies the clipboard size unit.
	//
	// - `B`: Bytes.
	//
	// - `KB`: 1024 bytes.
	//
	// **Default value:*	- `KB`.
	//
	// example:
	//
	// KB
	TextClipboardReadSizeUnit *string `json:"TextClipboardReadSizeUnit,omitempty" xml:"TextClipboardReadSizeUnit,omitempty"`
	// Specifies the clipboard size limit for outbound transfer (from the cloud browser to the local PC).
	//
	// **Value range:*	- 1 to 102400. The unit is specified by TextClipboardWriteSizeUnit.
	//
	// This value range does not change with unit conversion.
	//
	// example:
	//
	// 1
	TextClipboardWriteLimit *int32 `json:"TextClipboardWriteLimit,omitempty" xml:"TextClipboardWriteLimit,omitempty"`
	// Specifies the clipboard size unit.
	//
	// - `B`: Bytes.
	//
	// - `KB`: 1024 bytes.
	//
	// **Default value:*	- `KB`.
	//
	// example:
	//
	// KB
	TextClipboardWriteSizeUnit *string `json:"TextClipboardWriteSizeUnit,omitempty" xml:"TextClipboardWriteSizeUnit,omitempty"`
}

func (s CreateBrowserInstanceGroupRequestPolicyClipboardPolicy) String() string {
	return dara.Prettify(s)
}

func (s CreateBrowserInstanceGroupRequestPolicyClipboardPolicy) GoString() string {
	return s.String()
}

func (s *CreateBrowserInstanceGroupRequestPolicyClipboardPolicy) GetClipboard() *string {
	return s.Clipboard
}

func (s *CreateBrowserInstanceGroupRequestPolicyClipboardPolicy) GetClipboardReadLimit() *int32 {
	return s.ClipboardReadLimit
}

func (s *CreateBrowserInstanceGroupRequestPolicyClipboardPolicy) GetClipboardScope() *string {
	return s.ClipboardScope
}

func (s *CreateBrowserInstanceGroupRequestPolicyClipboardPolicy) GetClipboardSizeUnit() *string {
	return s.ClipboardSizeUnit
}

func (s *CreateBrowserInstanceGroupRequestPolicyClipboardPolicy) GetClipboardWriteLimit() *int32 {
	return s.ClipboardWriteLimit
}

func (s *CreateBrowserInstanceGroupRequestPolicyClipboardPolicy) GetFileClipboard() *string {
	return s.FileClipboard
}

func (s *CreateBrowserInstanceGroupRequestPolicyClipboardPolicy) GetRichTextClipboard() *string {
	return s.RichTextClipboard
}

func (s *CreateBrowserInstanceGroupRequestPolicyClipboardPolicy) GetRichTextClipboardLimit() *int32 {
	return s.RichTextClipboardLimit
}

func (s *CreateBrowserInstanceGroupRequestPolicyClipboardPolicy) GetRichTextClipboardReadLimit() *int32 {
	return s.RichTextClipboardReadLimit
}

func (s *CreateBrowserInstanceGroupRequestPolicyClipboardPolicy) GetRichTextClipboardReadSizeUnit() *string {
	return s.RichTextClipboardReadSizeUnit
}

func (s *CreateBrowserInstanceGroupRequestPolicyClipboardPolicy) GetRichTextClipboardSizeUnit() *string {
	return s.RichTextClipboardSizeUnit
}

func (s *CreateBrowserInstanceGroupRequestPolicyClipboardPolicy) GetRichTextClipboardWriteLimit() *int32 {
	return s.RichTextClipboardWriteLimit
}

func (s *CreateBrowserInstanceGroupRequestPolicyClipboardPolicy) GetRichTextClipboardWriteSizeUnit() *string {
	return s.RichTextClipboardWriteSizeUnit
}

func (s *CreateBrowserInstanceGroupRequestPolicyClipboardPolicy) GetTextClipboard() *string {
	return s.TextClipboard
}

func (s *CreateBrowserInstanceGroupRequestPolicyClipboardPolicy) GetTextClipboardReadLimit() *int32 {
	return s.TextClipboardReadLimit
}

func (s *CreateBrowserInstanceGroupRequestPolicyClipboardPolicy) GetTextClipboardReadSizeUnit() *string {
	return s.TextClipboardReadSizeUnit
}

func (s *CreateBrowserInstanceGroupRequestPolicyClipboardPolicy) GetTextClipboardWriteLimit() *int32 {
	return s.TextClipboardWriteLimit
}

func (s *CreateBrowserInstanceGroupRequestPolicyClipboardPolicy) GetTextClipboardWriteSizeUnit() *string {
	return s.TextClipboardWriteSizeUnit
}

func (s *CreateBrowserInstanceGroupRequestPolicyClipboardPolicy) SetClipboard(v string) *CreateBrowserInstanceGroupRequestPolicyClipboardPolicy {
	s.Clipboard = &v
	return s
}

func (s *CreateBrowserInstanceGroupRequestPolicyClipboardPolicy) SetClipboardReadLimit(v int32) *CreateBrowserInstanceGroupRequestPolicyClipboardPolicy {
	s.ClipboardReadLimit = &v
	return s
}

func (s *CreateBrowserInstanceGroupRequestPolicyClipboardPolicy) SetClipboardScope(v string) *CreateBrowserInstanceGroupRequestPolicyClipboardPolicy {
	s.ClipboardScope = &v
	return s
}

func (s *CreateBrowserInstanceGroupRequestPolicyClipboardPolicy) SetClipboardSizeUnit(v string) *CreateBrowserInstanceGroupRequestPolicyClipboardPolicy {
	s.ClipboardSizeUnit = &v
	return s
}

func (s *CreateBrowserInstanceGroupRequestPolicyClipboardPolicy) SetClipboardWriteLimit(v int32) *CreateBrowserInstanceGroupRequestPolicyClipboardPolicy {
	s.ClipboardWriteLimit = &v
	return s
}

func (s *CreateBrowserInstanceGroupRequestPolicyClipboardPolicy) SetFileClipboard(v string) *CreateBrowserInstanceGroupRequestPolicyClipboardPolicy {
	s.FileClipboard = &v
	return s
}

func (s *CreateBrowserInstanceGroupRequestPolicyClipboardPolicy) SetRichTextClipboard(v string) *CreateBrowserInstanceGroupRequestPolicyClipboardPolicy {
	s.RichTextClipboard = &v
	return s
}

func (s *CreateBrowserInstanceGroupRequestPolicyClipboardPolicy) SetRichTextClipboardLimit(v int32) *CreateBrowserInstanceGroupRequestPolicyClipboardPolicy {
	s.RichTextClipboardLimit = &v
	return s
}

func (s *CreateBrowserInstanceGroupRequestPolicyClipboardPolicy) SetRichTextClipboardReadLimit(v int32) *CreateBrowserInstanceGroupRequestPolicyClipboardPolicy {
	s.RichTextClipboardReadLimit = &v
	return s
}

func (s *CreateBrowserInstanceGroupRequestPolicyClipboardPolicy) SetRichTextClipboardReadSizeUnit(v string) *CreateBrowserInstanceGroupRequestPolicyClipboardPolicy {
	s.RichTextClipboardReadSizeUnit = &v
	return s
}

func (s *CreateBrowserInstanceGroupRequestPolicyClipboardPolicy) SetRichTextClipboardSizeUnit(v string) *CreateBrowserInstanceGroupRequestPolicyClipboardPolicy {
	s.RichTextClipboardSizeUnit = &v
	return s
}

func (s *CreateBrowserInstanceGroupRequestPolicyClipboardPolicy) SetRichTextClipboardWriteLimit(v int32) *CreateBrowserInstanceGroupRequestPolicyClipboardPolicy {
	s.RichTextClipboardWriteLimit = &v
	return s
}

func (s *CreateBrowserInstanceGroupRequestPolicyClipboardPolicy) SetRichTextClipboardWriteSizeUnit(v string) *CreateBrowserInstanceGroupRequestPolicyClipboardPolicy {
	s.RichTextClipboardWriteSizeUnit = &v
	return s
}

func (s *CreateBrowserInstanceGroupRequestPolicyClipboardPolicy) SetTextClipboard(v string) *CreateBrowserInstanceGroupRequestPolicyClipboardPolicy {
	s.TextClipboard = &v
	return s
}

func (s *CreateBrowserInstanceGroupRequestPolicyClipboardPolicy) SetTextClipboardReadLimit(v int32) *CreateBrowserInstanceGroupRequestPolicyClipboardPolicy {
	s.TextClipboardReadLimit = &v
	return s
}

func (s *CreateBrowserInstanceGroupRequestPolicyClipboardPolicy) SetTextClipboardReadSizeUnit(v string) *CreateBrowserInstanceGroupRequestPolicyClipboardPolicy {
	s.TextClipboardReadSizeUnit = &v
	return s
}

func (s *CreateBrowserInstanceGroupRequestPolicyClipboardPolicy) SetTextClipboardWriteLimit(v int32) *CreateBrowserInstanceGroupRequestPolicyClipboardPolicy {
	s.TextClipboardWriteLimit = &v
	return s
}

func (s *CreateBrowserInstanceGroupRequestPolicyClipboardPolicy) SetTextClipboardWriteSizeUnit(v string) *CreateBrowserInstanceGroupRequestPolicyClipboardPolicy {
	s.TextClipboardWriteSizeUnit = &v
	return s
}

func (s *CreateBrowserInstanceGroupRequestPolicyClipboardPolicy) Validate() error {
	return dara.Validate(s)
}

type CreateBrowserInstanceGroupRequestPolicyVideoPolicy struct {
	// Specifies the frame rate for the browser session.
	//
	// example:
	//
	// 60
	FrameRate *int32 `json:"FrameRate,omitempty" xml:"FrameRate,omitempty"`
}

func (s CreateBrowserInstanceGroupRequestPolicyVideoPolicy) String() string {
	return dara.Prettify(s)
}

func (s CreateBrowserInstanceGroupRequestPolicyVideoPolicy) GoString() string {
	return s.String()
}

func (s *CreateBrowserInstanceGroupRequestPolicyVideoPolicy) GetFrameRate() *int32 {
	return s.FrameRate
}

func (s *CreateBrowserInstanceGroupRequestPolicyVideoPolicy) SetFrameRate(v int32) *CreateBrowserInstanceGroupRequestPolicyVideoPolicy {
	s.FrameRate = &v
	return s
}

func (s *CreateBrowserInstanceGroupRequestPolicyVideoPolicy) Validate() error {
	return dara.Validate(s)
}

type CreateBrowserInstanceGroupRequestPolicyWatermarkPolicy struct {
	// The watermark switch. The value is case-insensitive.
	//
	// - `ON`: Enables the watermark.
	//
	// - `OFF`: Disables the watermark.
	//
	// The watermark content type list is not used when the watermark is disabled.
	//
	// example:
	//
	// ON
	WatermarkSwitch *string `json:"WatermarkSwitch,omitempty" xml:"WatermarkSwitch,omitempty"`
	// The list of watermark content types.
	//
	// - `EndUserId`: The user identifier.
	//
	// - `InstanceGroupId`: The delivery group identifier.
	//
	// - `ClientTime`: The current time of the client.
	//
	// Use the watermark types supported by the browser and the client.
	WatermarkTypes []*string `json:"WatermarkTypes,omitempty" xml:"WatermarkTypes,omitempty" type:"Repeated"`
}

func (s CreateBrowserInstanceGroupRequestPolicyWatermarkPolicy) String() string {
	return dara.Prettify(s)
}

func (s CreateBrowserInstanceGroupRequestPolicyWatermarkPolicy) GoString() string {
	return s.String()
}

func (s *CreateBrowserInstanceGroupRequestPolicyWatermarkPolicy) GetWatermarkSwitch() *string {
	return s.WatermarkSwitch
}

func (s *CreateBrowserInstanceGroupRequestPolicyWatermarkPolicy) GetWatermarkTypes() []*string {
	return s.WatermarkTypes
}

func (s *CreateBrowserInstanceGroupRequestPolicyWatermarkPolicy) SetWatermarkSwitch(v string) *CreateBrowserInstanceGroupRequestPolicyWatermarkPolicy {
	s.WatermarkSwitch = &v
	return s
}

func (s *CreateBrowserInstanceGroupRequestPolicyWatermarkPolicy) SetWatermarkTypes(v []*string) *CreateBrowserInstanceGroupRequestPolicyWatermarkPolicy {
	s.WatermarkTypes = v
	return s
}

func (s *CreateBrowserInstanceGroupRequestPolicyWatermarkPolicy) Validate() error {
	return dara.Validate(s)
}

type CreateBrowserInstanceGroupRequestSecurityPolicy struct {
	// Specifies whether to skip the user authorization check when connecting to the application.
	//
	// - `true`: Skips the check.
	//
	// - `false`: Performs the check.
	//
	// If SecurityPolicy is configured but this field is omitted, the user authorization check is performed.
	//
	// **Note:*	- This field cannot be used to skip OpenAPI authentication and RAM permission verification.
	//
	// example:
	//
	// false
	SkipUserAuthCheck *bool `json:"SkipUserAuthCheck,omitempty" xml:"SkipUserAuthCheck,omitempty"`
}

func (s CreateBrowserInstanceGroupRequestSecurityPolicy) String() string {
	return dara.Prettify(s)
}

func (s CreateBrowserInstanceGroupRequestSecurityPolicy) GoString() string {
	return s.String()
}

func (s *CreateBrowserInstanceGroupRequestSecurityPolicy) GetSkipUserAuthCheck() *bool {
	return s.SkipUserAuthCheck
}

func (s *CreateBrowserInstanceGroupRequestSecurityPolicy) SetSkipUserAuthCheck(v bool) *CreateBrowserInstanceGroupRequestSecurityPolicy {
	s.SkipUserAuthCheck = &v
	return s
}

func (s *CreateBrowserInstanceGroupRequestSecurityPolicy) Validate() error {
	return dara.Validate(s)
}

type CreateBrowserInstanceGroupRequestStoragePolicy struct {
	// The user data roaming configuration, which is used to retain user configuration data.
	UserProfile *CreateBrowserInstanceGroupRequestStoragePolicyUserProfile `json:"UserProfile,omitempty" xml:"UserProfile,omitempty" type:"Struct"`
}

func (s CreateBrowserInstanceGroupRequestStoragePolicy) String() string {
	return dara.Prettify(s)
}

func (s CreateBrowserInstanceGroupRequestStoragePolicy) GoString() string {
	return s.String()
}

func (s *CreateBrowserInstanceGroupRequestStoragePolicy) GetUserProfile() *CreateBrowserInstanceGroupRequestStoragePolicyUserProfile {
	return s.UserProfile
}

func (s *CreateBrowserInstanceGroupRequestStoragePolicy) SetUserProfile(v *CreateBrowserInstanceGroupRequestStoragePolicyUserProfile) *CreateBrowserInstanceGroupRequestStoragePolicy {
	s.UserProfile = v
	return s
}

func (s *CreateBrowserInstanceGroupRequestStoragePolicy) Validate() error {
	if s.UserProfile != nil {
		if err := s.UserProfile.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateBrowserInstanceGroupRequestStoragePolicyUserProfile struct {
	// The size of the user data roaming disk, in GB.
	//
	// example:
	//
	// 30
	UserProfileSize *int64 `json:"UserProfileSize,omitempty" xml:"UserProfileSize,omitempty"`
	// Specifies whether to enable user data roaming.
	//
	// - `true`: Enables user data roaming.
	//
	// - `false`: Disables user data roaming.
	//
	// In Windows scenarios, if this parameter is explicitly specified, the specified value is used. If this parameter is omitted, the user roaming configuration of the current account is used.
	//
	// example:
	//
	// true
	UserProfileSwitch *bool `json:"UserProfileSwitch,omitempty" xml:"UserProfileSwitch,omitempty"`
}

func (s CreateBrowserInstanceGroupRequestStoragePolicyUserProfile) String() string {
	return dara.Prettify(s)
}

func (s CreateBrowserInstanceGroupRequestStoragePolicyUserProfile) GoString() string {
	return s.String()
}

func (s *CreateBrowserInstanceGroupRequestStoragePolicyUserProfile) GetUserProfileSize() *int64 {
	return s.UserProfileSize
}

func (s *CreateBrowserInstanceGroupRequestStoragePolicyUserProfile) GetUserProfileSwitch() *bool {
	return s.UserProfileSwitch
}

func (s *CreateBrowserInstanceGroupRequestStoragePolicyUserProfile) SetUserProfileSize(v int64) *CreateBrowserInstanceGroupRequestStoragePolicyUserProfile {
	s.UserProfileSize = &v
	return s
}

func (s *CreateBrowserInstanceGroupRequestStoragePolicyUserProfile) SetUserProfileSwitch(v bool) *CreateBrowserInstanceGroupRequestStoragePolicyUserProfile {
	s.UserProfileSwitch = &v
	return s
}

func (s *CreateBrowserInstanceGroupRequestStoragePolicyUserProfile) Validate() error {
	return dara.Validate(s)
}

type CreateBrowserInstanceGroupRequestTag struct {
	// This parameter is not supported and does not need to be specified.
	//
	// example:
	//
	// -
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// This parameter is not supported and does not need to be specified.
	//
	// example:
	//
	// -
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateBrowserInstanceGroupRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateBrowserInstanceGroupRequestTag) GoString() string {
	return s.String()
}

func (s *CreateBrowserInstanceGroupRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateBrowserInstanceGroupRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateBrowserInstanceGroupRequestTag) SetKey(v string) *CreateBrowserInstanceGroupRequestTag {
	s.Key = &v
	return s
}

func (s *CreateBrowserInstanceGroupRequestTag) SetValue(v string) *CreateBrowserInstanceGroupRequestTag {
	s.Value = &v
	return s
}

func (s *CreateBrowserInstanceGroupRequestTag) Validate() error {
	return dara.Validate(s)
}

type CreateBrowserInstanceGroupRequestTimers struct {
	// This parameter is not supported and does not need to be specified.
	//
	// example:
	//
	// -
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// This parameter is not supported and does not need to be specified.
	//
	// example:
	//
	// -
	TimerType *string `json:"TimerType,omitempty" xml:"TimerType,omitempty"`
}

func (s CreateBrowserInstanceGroupRequestTimers) String() string {
	return dara.Prettify(s)
}

func (s CreateBrowserInstanceGroupRequestTimers) GoString() string {
	return s.String()
}

func (s *CreateBrowserInstanceGroupRequestTimers) GetInterval() *int32 {
	return s.Interval
}

func (s *CreateBrowserInstanceGroupRequestTimers) GetTimerType() *string {
	return s.TimerType
}

func (s *CreateBrowserInstanceGroupRequestTimers) SetInterval(v int32) *CreateBrowserInstanceGroupRequestTimers {
	s.Interval = &v
	return s
}

func (s *CreateBrowserInstanceGroupRequestTimers) SetTimerType(v string) *CreateBrowserInstanceGroupRequestTimers {
	s.TimerType = &v
	return s
}

func (s *CreateBrowserInstanceGroupRequestTimers) Validate() error {
	return dara.Validate(s)
}

type CreateBrowserInstanceGroupRequestUserInfo struct {
	// The account type of the authorized user.
	//
	// - `simple`: Convenience account.
	//
	// - `ad`: AD domain account.
	//
	// The account type must match the account types of the user and the office network.
	//
	// example:
	//
	// simple
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateBrowserInstanceGroupRequestUserInfo) String() string {
	return dara.Prettify(s)
}

func (s CreateBrowserInstanceGroupRequestUserInfo) GoString() string {
	return s.String()
}

func (s *CreateBrowserInstanceGroupRequestUserInfo) GetType() *string {
	return s.Type
}

func (s *CreateBrowserInstanceGroupRequestUserInfo) SetType(v string) *CreateBrowserInstanceGroupRequestUserInfo {
	s.Type = &v
	return s
}

func (s *CreateBrowserInstanceGroupRequestUserInfo) Validate() error {
	return dara.Validate(s)
}

type CreateBrowserInstanceGroupRequestUsers struct {
	// The identity of the authorized user for authorization.
	//
	// example:
	//
	// alice
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
}

func (s CreateBrowserInstanceGroupRequestUsers) String() string {
	return dara.Prettify(s)
}

func (s CreateBrowserInstanceGroupRequestUsers) GoString() string {
	return s.String()
}

func (s *CreateBrowserInstanceGroupRequestUsers) GetEndUserId() *string {
	return s.EndUserId
}

func (s *CreateBrowserInstanceGroupRequestUsers) SetEndUserId(v string) *CreateBrowserInstanceGroupRequestUsers {
	s.EndUserId = &v
	return s
}

func (s *CreateBrowserInstanceGroupRequestUsers) Validate() error {
	return dara.Validate(s)
}
