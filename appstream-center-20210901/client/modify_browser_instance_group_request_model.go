// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBrowserInstanceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBrowserConfig(v *ModifyBrowserInstanceGroupRequestBrowserConfig) *ModifyBrowserInstanceGroupRequest
	GetBrowserConfig() *ModifyBrowserInstanceGroupRequestBrowserConfig
	SetBrowserInstanceGroupId(v string) *ModifyBrowserInstanceGroupRequest
	GetBrowserInstanceGroupId() *string
	SetCloudBrowserName(v string) *ModifyBrowserInstanceGroupRequest
	GetCloudBrowserName() *string
	SetMaxAmount(v int32) *ModifyBrowserInstanceGroupRequest
	GetMaxAmount() *int32
	SetNetwork(v *ModifyBrowserInstanceGroupRequestNetwork) *ModifyBrowserInstanceGroupRequest
	GetNetwork() *ModifyBrowserInstanceGroupRequestNetwork
	SetPolicy(v *ModifyBrowserInstanceGroupRequestPolicy) *ModifyBrowserInstanceGroupRequest
	GetPolicy() *ModifyBrowserInstanceGroupRequestPolicy
	SetStoragePolicy(v *ModifyBrowserInstanceGroupRequestStoragePolicy) *ModifyBrowserInstanceGroupRequest
	GetStoragePolicy() *ModifyBrowserInstanceGroupRequestStoragePolicy
	SetTimers(v []*ModifyBrowserInstanceGroupRequestTimers) *ModifyBrowserInstanceGroupRequest
	GetTimers() []*ModifyBrowserInstanceGroupRequestTimers
}

type ModifyBrowserInstanceGroupRequest struct {
	// The browser configuration.
	BrowserConfig *ModifyBrowserInstanceGroupRequestBrowserConfig `json:"BrowserConfig,omitempty" xml:"BrowserConfig,omitempty" type:"Struct"`
	// The ID of the cloud browser to modify.
	//
	// This parameter is required.
	//
	// example:
	//
	// big-0bz55ixxxxx9xi9w9
	BrowserInstanceGroupId *string `json:"BrowserInstanceGroupId,omitempty" xml:"BrowserInstanceGroupId,omitempty"`
	// The name of the cloud browser.
	//
	// example:
	//
	// BrowserTest
	CloudBrowserName *string `json:"CloudBrowserName,omitempty" xml:"CloudBrowserName,omitempty"`
	// The maximum resource count. This parameter takes effect for monthly active pay-as-you-go billing.
	//
	// example:
	//
	// 5
	MaxAmount *int32 `json:"MaxAmount,omitempty" xml:"MaxAmount,omitempty"`
	// The network configuration.
	Network *ModifyBrowserInstanceGroupRequestNetwork `json:"Network,omitempty" xml:"Network,omitempty" type:"Struct"`
	// The access policy.
	Policy *ModifyBrowserInstanceGroupRequestPolicy `json:"Policy,omitempty" xml:"Policy,omitempty" type:"Struct"`
	// The storage-related policy.
	StoragePolicy *ModifyBrowserInstanceGroupRequestStoragePolicy `json:"StoragePolicy,omitempty" xml:"StoragePolicy,omitempty" type:"Struct"`
	// The timers.
	Timers []*ModifyBrowserInstanceGroupRequestTimers `json:"Timers,omitempty" xml:"Timers,omitempty" type:"Repeated"`
}

func (s ModifyBrowserInstanceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyBrowserInstanceGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyBrowserInstanceGroupRequest) GetBrowserConfig() *ModifyBrowserInstanceGroupRequestBrowserConfig {
	return s.BrowserConfig
}

func (s *ModifyBrowserInstanceGroupRequest) GetBrowserInstanceGroupId() *string {
	return s.BrowserInstanceGroupId
}

func (s *ModifyBrowserInstanceGroupRequest) GetCloudBrowserName() *string {
	return s.CloudBrowserName
}

func (s *ModifyBrowserInstanceGroupRequest) GetMaxAmount() *int32 {
	return s.MaxAmount
}

func (s *ModifyBrowserInstanceGroupRequest) GetNetwork() *ModifyBrowserInstanceGroupRequestNetwork {
	return s.Network
}

func (s *ModifyBrowserInstanceGroupRequest) GetPolicy() *ModifyBrowserInstanceGroupRequestPolicy {
	return s.Policy
}

func (s *ModifyBrowserInstanceGroupRequest) GetStoragePolicy() *ModifyBrowserInstanceGroupRequestStoragePolicy {
	return s.StoragePolicy
}

func (s *ModifyBrowserInstanceGroupRequest) GetTimers() []*ModifyBrowserInstanceGroupRequestTimers {
	return s.Timers
}

func (s *ModifyBrowserInstanceGroupRequest) SetBrowserConfig(v *ModifyBrowserInstanceGroupRequestBrowserConfig) *ModifyBrowserInstanceGroupRequest {
	s.BrowserConfig = v
	return s
}

func (s *ModifyBrowserInstanceGroupRequest) SetBrowserInstanceGroupId(v string) *ModifyBrowserInstanceGroupRequest {
	s.BrowserInstanceGroupId = &v
	return s
}

func (s *ModifyBrowserInstanceGroupRequest) SetCloudBrowserName(v string) *ModifyBrowserInstanceGroupRequest {
	s.CloudBrowserName = &v
	return s
}

func (s *ModifyBrowserInstanceGroupRequest) SetMaxAmount(v int32) *ModifyBrowserInstanceGroupRequest {
	s.MaxAmount = &v
	return s
}

func (s *ModifyBrowserInstanceGroupRequest) SetNetwork(v *ModifyBrowserInstanceGroupRequestNetwork) *ModifyBrowserInstanceGroupRequest {
	s.Network = v
	return s
}

func (s *ModifyBrowserInstanceGroupRequest) SetPolicy(v *ModifyBrowserInstanceGroupRequestPolicy) *ModifyBrowserInstanceGroupRequest {
	s.Policy = v
	return s
}

func (s *ModifyBrowserInstanceGroupRequest) SetStoragePolicy(v *ModifyBrowserInstanceGroupRequestStoragePolicy) *ModifyBrowserInstanceGroupRequest {
	s.StoragePolicy = v
	return s
}

func (s *ModifyBrowserInstanceGroupRequest) SetTimers(v []*ModifyBrowserInstanceGroupRequestTimers) *ModifyBrowserInstanceGroupRequest {
	s.Timers = v
	return s
}

func (s *ModifyBrowserInstanceGroupRequest) Validate() error {
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
	if s.Policy != nil {
		if err := s.Policy.Validate(); err != nil {
			return err
		}
	}
	if s.StoragePolicy != nil {
		if err := s.StoragePolicy.Validate(); err != nil {
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
	return nil
}

type ModifyBrowserInstanceGroupRequestBrowserConfig struct {
	// The bookmarks.
	Bookmarks []*ModifyBrowserInstanceGroupRequestBrowserConfigBookmarks `json:"Bookmarks,omitempty" xml:"Bookmarks,omitempty" type:"Repeated"`
	// The file path of the bookmark list.
	//
	// example:
	//
	// cn-hangzhou/aig_upm/xxx/temp/BrowserBookmarks/BrowserBookmarksTemplate.csv
	BookmarksFilePath *string `json:"BookmarksFilePath,omitempty" xml:"BookmarksFilePath,omitempty"`
	// The startup parameters.
	//
	// example:
	//
	// --incognito
	BrowserParam *string `json:"BrowserParam,omitempty" xml:"BrowserParam,omitempty"`
	// Specifies whether to enable cookies synchronization.
	CookiesSync *bool `json:"CookiesSync,omitempty" xml:"CookiesSync,omitempty"`
	// The homepage.
	//
	// example:
	//
	// https://www.aliyun.com
	Homepage *string `json:"Homepage,omitempty" xml:"Homepage,omitempty"`
	// The list of bookmarks to remove.
	RemoveBookmarks []*string `json:"RemoveBookmarks,omitempty" xml:"RemoveBookmarks,omitempty" type:"Repeated"`
}

func (s ModifyBrowserInstanceGroupRequestBrowserConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyBrowserInstanceGroupRequestBrowserConfig) GoString() string {
	return s.String()
}

func (s *ModifyBrowserInstanceGroupRequestBrowserConfig) GetBookmarks() []*ModifyBrowserInstanceGroupRequestBrowserConfigBookmarks {
	return s.Bookmarks
}

func (s *ModifyBrowserInstanceGroupRequestBrowserConfig) GetBookmarksFilePath() *string {
	return s.BookmarksFilePath
}

func (s *ModifyBrowserInstanceGroupRequestBrowserConfig) GetBrowserParam() *string {
	return s.BrowserParam
}

func (s *ModifyBrowserInstanceGroupRequestBrowserConfig) GetCookiesSync() *bool {
	return s.CookiesSync
}

func (s *ModifyBrowserInstanceGroupRequestBrowserConfig) GetHomepage() *string {
	return s.Homepage
}

func (s *ModifyBrowserInstanceGroupRequestBrowserConfig) GetRemoveBookmarks() []*string {
	return s.RemoveBookmarks
}

func (s *ModifyBrowserInstanceGroupRequestBrowserConfig) SetBookmarks(v []*ModifyBrowserInstanceGroupRequestBrowserConfigBookmarks) *ModifyBrowserInstanceGroupRequestBrowserConfig {
	s.Bookmarks = v
	return s
}

func (s *ModifyBrowserInstanceGroupRequestBrowserConfig) SetBookmarksFilePath(v string) *ModifyBrowserInstanceGroupRequestBrowserConfig {
	s.BookmarksFilePath = &v
	return s
}

func (s *ModifyBrowserInstanceGroupRequestBrowserConfig) SetBrowserParam(v string) *ModifyBrowserInstanceGroupRequestBrowserConfig {
	s.BrowserParam = &v
	return s
}

func (s *ModifyBrowserInstanceGroupRequestBrowserConfig) SetCookiesSync(v bool) *ModifyBrowserInstanceGroupRequestBrowserConfig {
	s.CookiesSync = &v
	return s
}

func (s *ModifyBrowserInstanceGroupRequestBrowserConfig) SetHomepage(v string) *ModifyBrowserInstanceGroupRequestBrowserConfig {
	s.Homepage = &v
	return s
}

func (s *ModifyBrowserInstanceGroupRequestBrowserConfig) SetRemoveBookmarks(v []*string) *ModifyBrowserInstanceGroupRequestBrowserConfig {
	s.RemoveBookmarks = v
	return s
}

func (s *ModifyBrowserInstanceGroupRequestBrowserConfig) Validate() error {
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

type ModifyBrowserInstanceGroupRequestBrowserConfigBookmarks struct {
	// The folder to which the bookmark belongs.
	//
	// example:
	//
	// test
	BookmarkFolder *string `json:"BookmarkFolder,omitempty" xml:"BookmarkFolder,omitempty"`
	// The bookmark ID. This parameter is required only for modification scenarios.
	//
	// example:
	//
	// bm-12345
	BookmarkId *string `json:"BookmarkId,omitempty" xml:"BookmarkId,omitempty"`
	// The bookmark name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	BookmarkName *string `json:"BookmarkName,omitempty" xml:"BookmarkName,omitempty"`
	// The bookmark URL.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://www.aliyun.com
	BookmarkURL *string `json:"BookmarkURL,omitempty" xml:"BookmarkURL,omitempty"`
}

func (s ModifyBrowserInstanceGroupRequestBrowserConfigBookmarks) String() string {
	return dara.Prettify(s)
}

func (s ModifyBrowserInstanceGroupRequestBrowserConfigBookmarks) GoString() string {
	return s.String()
}

func (s *ModifyBrowserInstanceGroupRequestBrowserConfigBookmarks) GetBookmarkFolder() *string {
	return s.BookmarkFolder
}

func (s *ModifyBrowserInstanceGroupRequestBrowserConfigBookmarks) GetBookmarkId() *string {
	return s.BookmarkId
}

func (s *ModifyBrowserInstanceGroupRequestBrowserConfigBookmarks) GetBookmarkName() *string {
	return s.BookmarkName
}

func (s *ModifyBrowserInstanceGroupRequestBrowserConfigBookmarks) GetBookmarkURL() *string {
	return s.BookmarkURL
}

func (s *ModifyBrowserInstanceGroupRequestBrowserConfigBookmarks) SetBookmarkFolder(v string) *ModifyBrowserInstanceGroupRequestBrowserConfigBookmarks {
	s.BookmarkFolder = &v
	return s
}

func (s *ModifyBrowserInstanceGroupRequestBrowserConfigBookmarks) SetBookmarkId(v string) *ModifyBrowserInstanceGroupRequestBrowserConfigBookmarks {
	s.BookmarkId = &v
	return s
}

func (s *ModifyBrowserInstanceGroupRequestBrowserConfigBookmarks) SetBookmarkName(v string) *ModifyBrowserInstanceGroupRequestBrowserConfigBookmarks {
	s.BookmarkName = &v
	return s
}

func (s *ModifyBrowserInstanceGroupRequestBrowserConfigBookmarks) SetBookmarkURL(v string) *ModifyBrowserInstanceGroupRequestBrowserConfigBookmarks {
	s.BookmarkURL = &v
	return s
}

func (s *ModifyBrowserInstanceGroupRequestBrowserConfigBookmarks) Validate() error {
	return dara.Validate(s)
}

type ModifyBrowserInstanceGroupRequestNetwork struct {
	// The access restriction type.
	//
	// example:
	//
	// ALLOW_LIST
	AccessRestriction *string `json:"AccessRestriction,omitempty" xml:"AccessRestriction,omitempty"`
	// The list of domain names to remove.
	RemoveRestrictedURLIds []*string `json:"RemoveRestrictedURLIds,omitempty" xml:"RemoveRestrictedURLIds,omitempty" type:"Repeated"`
	// The restricted domain name configurations.
	RestrictedURLs []*ModifyBrowserInstanceGroupRequestNetworkRestrictedURLs `json:"RestrictedURLs,omitempty" xml:"RestrictedURLs,omitempty" type:"Repeated"`
	// The file path of the restricted URLs.
	//
	// example:
	//
	// cn-hangzhou/aig_upm/xxx/temp/BrowserRestrictionUrls/URL白名单模版.csv
	RestrictedURLsFilePath *string `json:"RestrictedURLsFilePath,omitempty" xml:"RestrictedURLsFilePath,omitempty"`
}

func (s ModifyBrowserInstanceGroupRequestNetwork) String() string {
	return dara.Prettify(s)
}

func (s ModifyBrowserInstanceGroupRequestNetwork) GoString() string {
	return s.String()
}

func (s *ModifyBrowserInstanceGroupRequestNetwork) GetAccessRestriction() *string {
	return s.AccessRestriction
}

func (s *ModifyBrowserInstanceGroupRequestNetwork) GetRemoveRestrictedURLIds() []*string {
	return s.RemoveRestrictedURLIds
}

func (s *ModifyBrowserInstanceGroupRequestNetwork) GetRestrictedURLs() []*ModifyBrowserInstanceGroupRequestNetworkRestrictedURLs {
	return s.RestrictedURLs
}

func (s *ModifyBrowserInstanceGroupRequestNetwork) GetRestrictedURLsFilePath() *string {
	return s.RestrictedURLsFilePath
}

func (s *ModifyBrowserInstanceGroupRequestNetwork) SetAccessRestriction(v string) *ModifyBrowserInstanceGroupRequestNetwork {
	s.AccessRestriction = &v
	return s
}

func (s *ModifyBrowserInstanceGroupRequestNetwork) SetRemoveRestrictedURLIds(v []*string) *ModifyBrowserInstanceGroupRequestNetwork {
	s.RemoveRestrictedURLIds = v
	return s
}

func (s *ModifyBrowserInstanceGroupRequestNetwork) SetRestrictedURLs(v []*ModifyBrowserInstanceGroupRequestNetworkRestrictedURLs) *ModifyBrowserInstanceGroupRequestNetwork {
	s.RestrictedURLs = v
	return s
}

func (s *ModifyBrowserInstanceGroupRequestNetwork) SetRestrictedURLsFilePath(v string) *ModifyBrowserInstanceGroupRequestNetwork {
	s.RestrictedURLsFilePath = &v
	return s
}

func (s *ModifyBrowserInstanceGroupRequestNetwork) Validate() error {
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

type ModifyBrowserInstanceGroupRequestNetworkRestrictedURLs struct {
	// The domain name configuration ID. This parameter is required only for modification.
	//
	// example:
	//
	// ru-12345
	RestrictedURLId *string `json:"RestrictedURLId,omitempty" xml:"RestrictedURLId,omitempty"`
	// The domain name.
	//
	// example:
	//
	// aliyun.com
	URL *string `json:"URL,omitempty" xml:"URL,omitempty"`
}

func (s ModifyBrowserInstanceGroupRequestNetworkRestrictedURLs) String() string {
	return dara.Prettify(s)
}

func (s ModifyBrowserInstanceGroupRequestNetworkRestrictedURLs) GoString() string {
	return s.String()
}

func (s *ModifyBrowserInstanceGroupRequestNetworkRestrictedURLs) GetRestrictedURLId() *string {
	return s.RestrictedURLId
}

func (s *ModifyBrowserInstanceGroupRequestNetworkRestrictedURLs) GetURL() *string {
	return s.URL
}

func (s *ModifyBrowserInstanceGroupRequestNetworkRestrictedURLs) SetRestrictedURLId(v string) *ModifyBrowserInstanceGroupRequestNetworkRestrictedURLs {
	s.RestrictedURLId = &v
	return s
}

func (s *ModifyBrowserInstanceGroupRequestNetworkRestrictedURLs) SetURL(v string) *ModifyBrowserInstanceGroupRequestNetworkRestrictedURLs {
	s.URL = &v
	return s
}

func (s *ModifyBrowserInstanceGroupRequestNetworkRestrictedURLs) Validate() error {
	return dara.Validate(s)
}

type ModifyBrowserInstanceGroupRequestPolicy struct {
	// Specifies whether to enable screenshot protection.
	AppContentProtection *string `json:"AppContentProtection,omitempty" xml:"AppContentProtection,omitempty"`
	// The server-side access IP address whitelist.
	AuthorizeAccessPolicyRules []*ModifyBrowserInstanceGroupRequestPolicyAuthorizeAccessPolicyRules `json:"AuthorizeAccessPolicyRules,omitempty" xml:"AuthorizeAccessPolicyRules,omitempty" type:"Repeated"`
	// The logon client type control settings.
	ClientTypes []*ModifyBrowserInstanceGroupRequestPolicyClientTypes `json:"ClientTypes,omitempty" xml:"ClientTypes,omitempty" type:"Repeated"`
	// The clipboard-related policy.
	ClipboardPolicy *ModifyBrowserInstanceGroupRequestPolicyClipboardPolicy `json:"ClipboardPolicy,omitempty" xml:"ClipboardPolicy,omitempty" type:"Struct"`
	// The data retention policy upon disconnection.
	//
	// example:
	//
	// customTime
	DisconnectKeepSession *string `json:"DisconnectKeepSession,omitempty" xml:"DisconnectKeepSession,omitempty"`
	// The session retention duration upon disconnection.
	//
	// example:
	//
	// 15
	DisconnectKeepSessionTime *int32 `json:"DisconnectKeepSessionTime,omitempty" xml:"DisconnectKeepSessionTime,omitempty"`
	// Specifies whether to enable the floating ball file manager.
	//
	// example:
	//
	// off
	FileManager *string `json:"FileManager,omitempty" xml:"FileManager,omitempty"`
	// The file transfer policy for the web client.
	//
	// example:
	//
	// off
	Html5FileTransfer *string `json:"Html5FileTransfer,omitempty" xml:"Html5FileTransfer,omitempty"`
	// The policy for disconnecting sessions when no operation is performed.
	//
	// example:
	//
	// on
	NoOperationDisconnect *string `json:"NoOperationDisconnect,omitempty" xml:"NoOperationDisconnect,omitempty"`
	// The time in seconds before a session is disconnected when no operation is performed.
	//
	// example:
	//
	// 1
	NoOperationDisconnectTime *int32 `json:"NoOperationDisconnectTime,omitempty" xml:"NoOperationDisconnectTime,omitempty"`
	// The policy ID.
	//
	// example:
	//
	// pg-12345
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The policy version.
	//
	// example:
	//
	// Center
	PolicyVersion *string `json:"PolicyVersion,omitempty" xml:"PolicyVersion,omitempty"`
	// The display policy.
	VideoPolicy *ModifyBrowserInstanceGroupRequestPolicyVideoPolicy `json:"VideoPolicy,omitempty" xml:"VideoPolicy,omitempty" type:"Struct"`
	// The watermark configuration.
	WatermarkPolicy *ModifyBrowserInstanceGroupRequestPolicyWatermarkPolicy `json:"WatermarkPolicy,omitempty" xml:"WatermarkPolicy,omitempty" type:"Struct"`
}

func (s ModifyBrowserInstanceGroupRequestPolicy) String() string {
	return dara.Prettify(s)
}

func (s ModifyBrowserInstanceGroupRequestPolicy) GoString() string {
	return s.String()
}

func (s *ModifyBrowserInstanceGroupRequestPolicy) GetAppContentProtection() *string {
	return s.AppContentProtection
}

func (s *ModifyBrowserInstanceGroupRequestPolicy) GetAuthorizeAccessPolicyRules() []*ModifyBrowserInstanceGroupRequestPolicyAuthorizeAccessPolicyRules {
	return s.AuthorizeAccessPolicyRules
}

func (s *ModifyBrowserInstanceGroupRequestPolicy) GetClientTypes() []*ModifyBrowserInstanceGroupRequestPolicyClientTypes {
	return s.ClientTypes
}

func (s *ModifyBrowserInstanceGroupRequestPolicy) GetClipboardPolicy() *ModifyBrowserInstanceGroupRequestPolicyClipboardPolicy {
	return s.ClipboardPolicy
}

func (s *ModifyBrowserInstanceGroupRequestPolicy) GetDisconnectKeepSession() *string {
	return s.DisconnectKeepSession
}

func (s *ModifyBrowserInstanceGroupRequestPolicy) GetDisconnectKeepSessionTime() *int32 {
	return s.DisconnectKeepSessionTime
}

func (s *ModifyBrowserInstanceGroupRequestPolicy) GetFileManager() *string {
	return s.FileManager
}

func (s *ModifyBrowserInstanceGroupRequestPolicy) GetHtml5FileTransfer() *string {
	return s.Html5FileTransfer
}

func (s *ModifyBrowserInstanceGroupRequestPolicy) GetNoOperationDisconnect() *string {
	return s.NoOperationDisconnect
}

func (s *ModifyBrowserInstanceGroupRequestPolicy) GetNoOperationDisconnectTime() *int32 {
	return s.NoOperationDisconnectTime
}

func (s *ModifyBrowserInstanceGroupRequestPolicy) GetPolicyId() *string {
	return s.PolicyId
}

func (s *ModifyBrowserInstanceGroupRequestPolicy) GetPolicyVersion() *string {
	return s.PolicyVersion
}

func (s *ModifyBrowserInstanceGroupRequestPolicy) GetVideoPolicy() *ModifyBrowserInstanceGroupRequestPolicyVideoPolicy {
	return s.VideoPolicy
}

func (s *ModifyBrowserInstanceGroupRequestPolicy) GetWatermarkPolicy() *ModifyBrowserInstanceGroupRequestPolicyWatermarkPolicy {
	return s.WatermarkPolicy
}

func (s *ModifyBrowserInstanceGroupRequestPolicy) SetAppContentProtection(v string) *ModifyBrowserInstanceGroupRequestPolicy {
	s.AppContentProtection = &v
	return s
}

func (s *ModifyBrowserInstanceGroupRequestPolicy) SetAuthorizeAccessPolicyRules(v []*ModifyBrowserInstanceGroupRequestPolicyAuthorizeAccessPolicyRules) *ModifyBrowserInstanceGroupRequestPolicy {
	s.AuthorizeAccessPolicyRules = v
	return s
}

func (s *ModifyBrowserInstanceGroupRequestPolicy) SetClientTypes(v []*ModifyBrowserInstanceGroupRequestPolicyClientTypes) *ModifyBrowserInstanceGroupRequestPolicy {
	s.ClientTypes = v
	return s
}

func (s *ModifyBrowserInstanceGroupRequestPolicy) SetClipboardPolicy(v *ModifyBrowserInstanceGroupRequestPolicyClipboardPolicy) *ModifyBrowserInstanceGroupRequestPolicy {
	s.ClipboardPolicy = v
	return s
}

func (s *ModifyBrowserInstanceGroupRequestPolicy) SetDisconnectKeepSession(v string) *ModifyBrowserInstanceGroupRequestPolicy {
	s.DisconnectKeepSession = &v
	return s
}

func (s *ModifyBrowserInstanceGroupRequestPolicy) SetDisconnectKeepSessionTime(v int32) *ModifyBrowserInstanceGroupRequestPolicy {
	s.DisconnectKeepSessionTime = &v
	return s
}

func (s *ModifyBrowserInstanceGroupRequestPolicy) SetFileManager(v string) *ModifyBrowserInstanceGroupRequestPolicy {
	s.FileManager = &v
	return s
}

func (s *ModifyBrowserInstanceGroupRequestPolicy) SetHtml5FileTransfer(v string) *ModifyBrowserInstanceGroupRequestPolicy {
	s.Html5FileTransfer = &v
	return s
}

func (s *ModifyBrowserInstanceGroupRequestPolicy) SetNoOperationDisconnect(v string) *ModifyBrowserInstanceGroupRequestPolicy {
	s.NoOperationDisconnect = &v
	return s
}

func (s *ModifyBrowserInstanceGroupRequestPolicy) SetNoOperationDisconnectTime(v int32) *ModifyBrowserInstanceGroupRequestPolicy {
	s.NoOperationDisconnectTime = &v
	return s
}

func (s *ModifyBrowserInstanceGroupRequestPolicy) SetPolicyId(v string) *ModifyBrowserInstanceGroupRequestPolicy {
	s.PolicyId = &v
	return s
}

func (s *ModifyBrowserInstanceGroupRequestPolicy) SetPolicyVersion(v string) *ModifyBrowserInstanceGroupRequestPolicy {
	s.PolicyVersion = &v
	return s
}

func (s *ModifyBrowserInstanceGroupRequestPolicy) SetVideoPolicy(v *ModifyBrowserInstanceGroupRequestPolicyVideoPolicy) *ModifyBrowserInstanceGroupRequestPolicy {
	s.VideoPolicy = v
	return s
}

func (s *ModifyBrowserInstanceGroupRequestPolicy) SetWatermarkPolicy(v *ModifyBrowserInstanceGroupRequestPolicyWatermarkPolicy) *ModifyBrowserInstanceGroupRequestPolicy {
	s.WatermarkPolicy = v
	return s
}

func (s *ModifyBrowserInstanceGroupRequestPolicy) Validate() error {
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

type ModifyBrowserInstanceGroupRequestPolicyAuthorizeAccessPolicyRules struct {
	CidrIp      *string `json:"CidrIp,omitempty" xml:"CidrIp,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s ModifyBrowserInstanceGroupRequestPolicyAuthorizeAccessPolicyRules) String() string {
	return dara.Prettify(s)
}

func (s ModifyBrowserInstanceGroupRequestPolicyAuthorizeAccessPolicyRules) GoString() string {
	return s.String()
}

func (s *ModifyBrowserInstanceGroupRequestPolicyAuthorizeAccessPolicyRules) GetCidrIp() *string {
	return s.CidrIp
}

func (s *ModifyBrowserInstanceGroupRequestPolicyAuthorizeAccessPolicyRules) GetDescription() *string {
	return s.Description
}

func (s *ModifyBrowserInstanceGroupRequestPolicyAuthorizeAccessPolicyRules) SetCidrIp(v string) *ModifyBrowserInstanceGroupRequestPolicyAuthorizeAccessPolicyRules {
	s.CidrIp = &v
	return s
}

func (s *ModifyBrowserInstanceGroupRequestPolicyAuthorizeAccessPolicyRules) SetDescription(v string) *ModifyBrowserInstanceGroupRequestPolicyAuthorizeAccessPolicyRules {
	s.Description = &v
	return s
}

func (s *ModifyBrowserInstanceGroupRequestPolicyAuthorizeAccessPolicyRules) Validate() error {
	return dara.Validate(s)
}

type ModifyBrowserInstanceGroupRequestPolicyClientTypes struct {
	ClientType *string `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ModifyBrowserInstanceGroupRequestPolicyClientTypes) String() string {
	return dara.Prettify(s)
}

func (s ModifyBrowserInstanceGroupRequestPolicyClientTypes) GoString() string {
	return s.String()
}

func (s *ModifyBrowserInstanceGroupRequestPolicyClientTypes) GetClientType() *string {
	return s.ClientType
}

func (s *ModifyBrowserInstanceGroupRequestPolicyClientTypes) GetStatus() *string {
	return s.Status
}

func (s *ModifyBrowserInstanceGroupRequestPolicyClientTypes) SetClientType(v string) *ModifyBrowserInstanceGroupRequestPolicyClientTypes {
	s.ClientType = &v
	return s
}

func (s *ModifyBrowserInstanceGroupRequestPolicyClientTypes) SetStatus(v string) *ModifyBrowserInstanceGroupRequestPolicyClientTypes {
	s.Status = &v
	return s
}

func (s *ModifyBrowserInstanceGroupRequestPolicyClientTypes) Validate() error {
	return dara.Validate(s)
}

type ModifyBrowserInstanceGroupRequestPolicyClipboardPolicy struct {
	// The clipboard policy.
	//
	// example:
	//
	// off
	Clipboard *string `json:"Clipboard,omitempty" xml:"Clipboard,omitempty"`
	// The clipboard read length limit.
	//
	// example:
	//
	// 1000
	ClipboardReadLimit *int32 `json:"ClipboardReadLimit,omitempty" xml:"ClipboardReadLimit,omitempty"`
	// The clipboard control scope.
	//
	// example:
	//
	// global
	ClipboardScope *string `json:"ClipboardScope,omitempty" xml:"ClipboardScope,omitempty"`
	// The clipboard size unit.
	//
	// example:
	//
	// B
	ClipboardSizeUnit *string `json:"ClipboardSizeUnit,omitempty" xml:"ClipboardSizeUnit,omitempty"`
	// The clipboard write length limit.
	//
	// example:
	//
	// 1000
	ClipboardWriteLimit *int32 `json:"ClipboardWriteLimit,omitempty" xml:"ClipboardWriteLimit,omitempty"`
	// The file clipboard policy.
	//
	// example:
	//
	// off
	FileClipboard *string `json:"FileClipboard,omitempty" xml:"FileClipboard,omitempty"`
	// The rich text clipboard policy.
	//
	// example:
	//
	// off
	RichTextClipboard *string `json:"RichTextClipboard,omitempty" xml:"RichTextClipboard,omitempty"`
	// The rich text clipboard limit.
	//
	// example:
	//
	// 1
	RichTextClipboardLimit *int32 `json:"RichTextClipboardLimit,omitempty" xml:"RichTextClipboardLimit,omitempty"`
	// The maximum size of rich text that can be downloaded from the cloud via the clipboard.
	//
	// example:
	//
	// 1
	RichTextClipboardReadLimit *int32 `json:"RichTextClipboardReadLimit,omitempty" xml:"RichTextClipboardReadLimit,omitempty"`
	// The size unit for rich text clipboard downloads from the cloud.
	//
	// example:
	//
	// KB
	RichTextClipboardReadSizeUnit *string `json:"RichTextClipboardReadSizeUnit,omitempty" xml:"RichTextClipboardReadSizeUnit,omitempty"`
	// The rich text clipboard size unit.
	//
	// example:
	//
	// B
	RichTextClipboardSizeUnit *string `json:"RichTextClipboardSizeUnit,omitempty" xml:"RichTextClipboardSizeUnit,omitempty"`
	// The maximum size of rich text that can be uploaded to the cloud via the clipboard.
	//
	// example:
	//
	// 1
	RichTextClipboardWriteLimit *int32 `json:"RichTextClipboardWriteLimit,omitempty" xml:"RichTextClipboardWriteLimit,omitempty"`
	// The size unit for rich text clipboard uploads to the cloud.
	//
	// example:
	//
	// KB
	RichTextClipboardWriteSizeUnit *string `json:"RichTextClipboardWriteSizeUnit,omitempty" xml:"RichTextClipboardWriteSizeUnit,omitempty"`
	// The text clipboard policy.
	//
	// example:
	//
	// off
	TextClipboard *string `json:"TextClipboard,omitempty" xml:"TextClipboard,omitempty"`
	// The maximum size of text that can be downloaded from the cloud via the clipboard.
	//
	// example:
	//
	// 1
	TextClipboardReadLimit *int32 `json:"TextClipboardReadLimit,omitempty" xml:"TextClipboardReadLimit,omitempty"`
	// The size unit for text clipboard downloads from the cloud.
	//
	// example:
	//
	// KB
	TextClipboardReadSizeUnit *string `json:"TextClipboardReadSizeUnit,omitempty" xml:"TextClipboardReadSizeUnit,omitempty"`
	// The maximum size of text that can be uploaded to the cloud via the clipboard.
	//
	// example:
	//
	// 1
	TextClipboardWriteLimit *int32 `json:"TextClipboardWriteLimit,omitempty" xml:"TextClipboardWriteLimit,omitempty"`
	// The size unit for text clipboard uploads to the cloud.
	//
	// example:
	//
	// KB
	TextClipboardWriteSizeUnit *string `json:"TextClipboardWriteSizeUnit,omitempty" xml:"TextClipboardWriteSizeUnit,omitempty"`
}

func (s ModifyBrowserInstanceGroupRequestPolicyClipboardPolicy) String() string {
	return dara.Prettify(s)
}

func (s ModifyBrowserInstanceGroupRequestPolicyClipboardPolicy) GoString() string {
	return s.String()
}

func (s *ModifyBrowserInstanceGroupRequestPolicyClipboardPolicy) GetClipboard() *string {
	return s.Clipboard
}

func (s *ModifyBrowserInstanceGroupRequestPolicyClipboardPolicy) GetClipboardReadLimit() *int32 {
	return s.ClipboardReadLimit
}

func (s *ModifyBrowserInstanceGroupRequestPolicyClipboardPolicy) GetClipboardScope() *string {
	return s.ClipboardScope
}

func (s *ModifyBrowserInstanceGroupRequestPolicyClipboardPolicy) GetClipboardSizeUnit() *string {
	return s.ClipboardSizeUnit
}

func (s *ModifyBrowserInstanceGroupRequestPolicyClipboardPolicy) GetClipboardWriteLimit() *int32 {
	return s.ClipboardWriteLimit
}

func (s *ModifyBrowserInstanceGroupRequestPolicyClipboardPolicy) GetFileClipboard() *string {
	return s.FileClipboard
}

func (s *ModifyBrowserInstanceGroupRequestPolicyClipboardPolicy) GetRichTextClipboard() *string {
	return s.RichTextClipboard
}

func (s *ModifyBrowserInstanceGroupRequestPolicyClipboardPolicy) GetRichTextClipboardLimit() *int32 {
	return s.RichTextClipboardLimit
}

func (s *ModifyBrowserInstanceGroupRequestPolicyClipboardPolicy) GetRichTextClipboardReadLimit() *int32 {
	return s.RichTextClipboardReadLimit
}

func (s *ModifyBrowserInstanceGroupRequestPolicyClipboardPolicy) GetRichTextClipboardReadSizeUnit() *string {
	return s.RichTextClipboardReadSizeUnit
}

func (s *ModifyBrowserInstanceGroupRequestPolicyClipboardPolicy) GetRichTextClipboardSizeUnit() *string {
	return s.RichTextClipboardSizeUnit
}

func (s *ModifyBrowserInstanceGroupRequestPolicyClipboardPolicy) GetRichTextClipboardWriteLimit() *int32 {
	return s.RichTextClipboardWriteLimit
}

func (s *ModifyBrowserInstanceGroupRequestPolicyClipboardPolicy) GetRichTextClipboardWriteSizeUnit() *string {
	return s.RichTextClipboardWriteSizeUnit
}

func (s *ModifyBrowserInstanceGroupRequestPolicyClipboardPolicy) GetTextClipboard() *string {
	return s.TextClipboard
}

func (s *ModifyBrowserInstanceGroupRequestPolicyClipboardPolicy) GetTextClipboardReadLimit() *int32 {
	return s.TextClipboardReadLimit
}

func (s *ModifyBrowserInstanceGroupRequestPolicyClipboardPolicy) GetTextClipboardReadSizeUnit() *string {
	return s.TextClipboardReadSizeUnit
}

func (s *ModifyBrowserInstanceGroupRequestPolicyClipboardPolicy) GetTextClipboardWriteLimit() *int32 {
	return s.TextClipboardWriteLimit
}

func (s *ModifyBrowserInstanceGroupRequestPolicyClipboardPolicy) GetTextClipboardWriteSizeUnit() *string {
	return s.TextClipboardWriteSizeUnit
}

func (s *ModifyBrowserInstanceGroupRequestPolicyClipboardPolicy) SetClipboard(v string) *ModifyBrowserInstanceGroupRequestPolicyClipboardPolicy {
	s.Clipboard = &v
	return s
}

func (s *ModifyBrowserInstanceGroupRequestPolicyClipboardPolicy) SetClipboardReadLimit(v int32) *ModifyBrowserInstanceGroupRequestPolicyClipboardPolicy {
	s.ClipboardReadLimit = &v
	return s
}

func (s *ModifyBrowserInstanceGroupRequestPolicyClipboardPolicy) SetClipboardScope(v string) *ModifyBrowserInstanceGroupRequestPolicyClipboardPolicy {
	s.ClipboardScope = &v
	return s
}

func (s *ModifyBrowserInstanceGroupRequestPolicyClipboardPolicy) SetClipboardSizeUnit(v string) *ModifyBrowserInstanceGroupRequestPolicyClipboardPolicy {
	s.ClipboardSizeUnit = &v
	return s
}

func (s *ModifyBrowserInstanceGroupRequestPolicyClipboardPolicy) SetClipboardWriteLimit(v int32) *ModifyBrowserInstanceGroupRequestPolicyClipboardPolicy {
	s.ClipboardWriteLimit = &v
	return s
}

func (s *ModifyBrowserInstanceGroupRequestPolicyClipboardPolicy) SetFileClipboard(v string) *ModifyBrowserInstanceGroupRequestPolicyClipboardPolicy {
	s.FileClipboard = &v
	return s
}

func (s *ModifyBrowserInstanceGroupRequestPolicyClipboardPolicy) SetRichTextClipboard(v string) *ModifyBrowserInstanceGroupRequestPolicyClipboardPolicy {
	s.RichTextClipboard = &v
	return s
}

func (s *ModifyBrowserInstanceGroupRequestPolicyClipboardPolicy) SetRichTextClipboardLimit(v int32) *ModifyBrowserInstanceGroupRequestPolicyClipboardPolicy {
	s.RichTextClipboardLimit = &v
	return s
}

func (s *ModifyBrowserInstanceGroupRequestPolicyClipboardPolicy) SetRichTextClipboardReadLimit(v int32) *ModifyBrowserInstanceGroupRequestPolicyClipboardPolicy {
	s.RichTextClipboardReadLimit = &v
	return s
}

func (s *ModifyBrowserInstanceGroupRequestPolicyClipboardPolicy) SetRichTextClipboardReadSizeUnit(v string) *ModifyBrowserInstanceGroupRequestPolicyClipboardPolicy {
	s.RichTextClipboardReadSizeUnit = &v
	return s
}

func (s *ModifyBrowserInstanceGroupRequestPolicyClipboardPolicy) SetRichTextClipboardSizeUnit(v string) *ModifyBrowserInstanceGroupRequestPolicyClipboardPolicy {
	s.RichTextClipboardSizeUnit = &v
	return s
}

func (s *ModifyBrowserInstanceGroupRequestPolicyClipboardPolicy) SetRichTextClipboardWriteLimit(v int32) *ModifyBrowserInstanceGroupRequestPolicyClipboardPolicy {
	s.RichTextClipboardWriteLimit = &v
	return s
}

func (s *ModifyBrowserInstanceGroupRequestPolicyClipboardPolicy) SetRichTextClipboardWriteSizeUnit(v string) *ModifyBrowserInstanceGroupRequestPolicyClipboardPolicy {
	s.RichTextClipboardWriteSizeUnit = &v
	return s
}

func (s *ModifyBrowserInstanceGroupRequestPolicyClipboardPolicy) SetTextClipboard(v string) *ModifyBrowserInstanceGroupRequestPolicyClipboardPolicy {
	s.TextClipboard = &v
	return s
}

func (s *ModifyBrowserInstanceGroupRequestPolicyClipboardPolicy) SetTextClipboardReadLimit(v int32) *ModifyBrowserInstanceGroupRequestPolicyClipboardPolicy {
	s.TextClipboardReadLimit = &v
	return s
}

func (s *ModifyBrowserInstanceGroupRequestPolicyClipboardPolicy) SetTextClipboardReadSizeUnit(v string) *ModifyBrowserInstanceGroupRequestPolicyClipboardPolicy {
	s.TextClipboardReadSizeUnit = &v
	return s
}

func (s *ModifyBrowserInstanceGroupRequestPolicyClipboardPolicy) SetTextClipboardWriteLimit(v int32) *ModifyBrowserInstanceGroupRequestPolicyClipboardPolicy {
	s.TextClipboardWriteLimit = &v
	return s
}

func (s *ModifyBrowserInstanceGroupRequestPolicyClipboardPolicy) SetTextClipboardWriteSizeUnit(v string) *ModifyBrowserInstanceGroupRequestPolicyClipboardPolicy {
	s.TextClipboardWriteSizeUnit = &v
	return s
}

func (s *ModifyBrowserInstanceGroupRequestPolicyClipboardPolicy) Validate() error {
	return dara.Validate(s)
}

type ModifyBrowserInstanceGroupRequestPolicyVideoPolicy struct {
	// The frame rate.
	//
	// example:
	//
	// 60
	FrameRate *int32 `json:"FrameRate,omitempty" xml:"FrameRate,omitempty"`
}

func (s ModifyBrowserInstanceGroupRequestPolicyVideoPolicy) String() string {
	return dara.Prettify(s)
}

func (s ModifyBrowserInstanceGroupRequestPolicyVideoPolicy) GoString() string {
	return s.String()
}

func (s *ModifyBrowserInstanceGroupRequestPolicyVideoPolicy) GetFrameRate() *int32 {
	return s.FrameRate
}

func (s *ModifyBrowserInstanceGroupRequestPolicyVideoPolicy) SetFrameRate(v int32) *ModifyBrowserInstanceGroupRequestPolicyVideoPolicy {
	s.FrameRate = &v
	return s
}

func (s *ModifyBrowserInstanceGroupRequestPolicyVideoPolicy) Validate() error {
	return dara.Validate(s)
}

type ModifyBrowserInstanceGroupRequestPolicyWatermarkPolicy struct {
	// The watermark switch.
	//
	// example:
	//
	// off
	WatermarkSwitch *string `json:"WatermarkSwitch,omitempty" xml:"WatermarkSwitch,omitempty"`
	// The list of watermark types.
	WatermarkTypes []*string `json:"WatermarkTypes,omitempty" xml:"WatermarkTypes,omitempty" type:"Repeated"`
}

func (s ModifyBrowserInstanceGroupRequestPolicyWatermarkPolicy) String() string {
	return dara.Prettify(s)
}

func (s ModifyBrowserInstanceGroupRequestPolicyWatermarkPolicy) GoString() string {
	return s.String()
}

func (s *ModifyBrowserInstanceGroupRequestPolicyWatermarkPolicy) GetWatermarkSwitch() *string {
	return s.WatermarkSwitch
}

func (s *ModifyBrowserInstanceGroupRequestPolicyWatermarkPolicy) GetWatermarkTypes() []*string {
	return s.WatermarkTypes
}

func (s *ModifyBrowserInstanceGroupRequestPolicyWatermarkPolicy) SetWatermarkSwitch(v string) *ModifyBrowserInstanceGroupRequestPolicyWatermarkPolicy {
	s.WatermarkSwitch = &v
	return s
}

func (s *ModifyBrowserInstanceGroupRequestPolicyWatermarkPolicy) SetWatermarkTypes(v []*string) *ModifyBrowserInstanceGroupRequestPolicyWatermarkPolicy {
	s.WatermarkTypes = v
	return s
}

func (s *ModifyBrowserInstanceGroupRequestPolicyWatermarkPolicy) Validate() error {
	return dara.Validate(s)
}

type ModifyBrowserInstanceGroupRequestStoragePolicy struct {
	// The user roaming policy.
	UserProfile *ModifyBrowserInstanceGroupRequestStoragePolicyUserProfile `json:"UserProfile,omitempty" xml:"UserProfile,omitempty" type:"Struct"`
}

func (s ModifyBrowserInstanceGroupRequestStoragePolicy) String() string {
	return dara.Prettify(s)
}

func (s ModifyBrowserInstanceGroupRequestStoragePolicy) GoString() string {
	return s.String()
}

func (s *ModifyBrowserInstanceGroupRequestStoragePolicy) GetUserProfile() *ModifyBrowserInstanceGroupRequestStoragePolicyUserProfile {
	return s.UserProfile
}

func (s *ModifyBrowserInstanceGroupRequestStoragePolicy) SetUserProfile(v *ModifyBrowserInstanceGroupRequestStoragePolicyUserProfile) *ModifyBrowserInstanceGroupRequestStoragePolicy {
	s.UserProfile = v
	return s
}

func (s *ModifyBrowserInstanceGroupRequestStoragePolicy) Validate() error {
	if s.UserProfile != nil {
		if err := s.UserProfile.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyBrowserInstanceGroupRequestStoragePolicyUserProfile struct {
	// Specifies whether to enable user roaming.
	UserProfileSwitch *bool `json:"UserProfileSwitch,omitempty" xml:"UserProfileSwitch,omitempty"`
}

func (s ModifyBrowserInstanceGroupRequestStoragePolicyUserProfile) String() string {
	return dara.Prettify(s)
}

func (s ModifyBrowserInstanceGroupRequestStoragePolicyUserProfile) GoString() string {
	return s.String()
}

func (s *ModifyBrowserInstanceGroupRequestStoragePolicyUserProfile) GetUserProfileSwitch() *bool {
	return s.UserProfileSwitch
}

func (s *ModifyBrowserInstanceGroupRequestStoragePolicyUserProfile) SetUserProfileSwitch(v bool) *ModifyBrowserInstanceGroupRequestStoragePolicyUserProfile {
	s.UserProfileSwitch = &v
	return s
}

func (s *ModifyBrowserInstanceGroupRequestStoragePolicyUserProfile) Validate() error {
	return dara.Validate(s)
}

type ModifyBrowserInstanceGroupRequestTimers struct {
	// The interval.
	//
	// example:
	//
	// 15
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The timer type.
	//
	// example:
	//
	// SESSION_TIMEOUT
	TimerType *string `json:"TimerType,omitempty" xml:"TimerType,omitempty"`
}

func (s ModifyBrowserInstanceGroupRequestTimers) String() string {
	return dara.Prettify(s)
}

func (s ModifyBrowserInstanceGroupRequestTimers) GoString() string {
	return s.String()
}

func (s *ModifyBrowserInstanceGroupRequestTimers) GetInterval() *int32 {
	return s.Interval
}

func (s *ModifyBrowserInstanceGroupRequestTimers) GetTimerType() *string {
	return s.TimerType
}

func (s *ModifyBrowserInstanceGroupRequestTimers) SetInterval(v int32) *ModifyBrowserInstanceGroupRequestTimers {
	s.Interval = &v
	return s
}

func (s *ModifyBrowserInstanceGroupRequestTimers) SetTimerType(v string) *ModifyBrowserInstanceGroupRequestTimers {
	s.TimerType = &v
	return s
}

func (s *ModifyBrowserInstanceGroupRequestTimers) Validate() error {
	return dara.Validate(s)
}
