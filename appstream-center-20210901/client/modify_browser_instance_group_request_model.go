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
	SetNetwork(v *ModifyBrowserInstanceGroupRequestNetwork) *ModifyBrowserInstanceGroupRequest
	GetNetwork() *ModifyBrowserInstanceGroupRequestNetwork
	SetPolicy(v *ModifyBrowserInstanceGroupRequestPolicy) *ModifyBrowserInstanceGroupRequest
	GetPolicy() *ModifyBrowserInstanceGroupRequestPolicy
	SetTimers(v []*ModifyBrowserInstanceGroupRequestTimers) *ModifyBrowserInstanceGroupRequest
	GetTimers() []*ModifyBrowserInstanceGroupRequestTimers
}

type ModifyBrowserInstanceGroupRequest struct {
	// The browser settings.
	BrowserConfig *ModifyBrowserInstanceGroupRequestBrowserConfig `json:"BrowserConfig,omitempty" xml:"BrowserConfig,omitempty" type:"Struct"`
	// The ID of the cloud browser to be modified.
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
	// The network configurations.
	Network *ModifyBrowserInstanceGroupRequestNetwork `json:"Network,omitempty" xml:"Network,omitempty" type:"Struct"`
	// The access policy.
	Policy *ModifyBrowserInstanceGroupRequestPolicy `json:"Policy,omitempty" xml:"Policy,omitempty" type:"Struct"`
	// The timer.
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

func (s *ModifyBrowserInstanceGroupRequest) GetNetwork() *ModifyBrowserInstanceGroupRequestNetwork {
	return s.Network
}

func (s *ModifyBrowserInstanceGroupRequest) GetPolicy() *ModifyBrowserInstanceGroupRequestPolicy {
	return s.Policy
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

func (s *ModifyBrowserInstanceGroupRequest) SetNetwork(v *ModifyBrowserInstanceGroupRequestNetwork) *ModifyBrowserInstanceGroupRequest {
	s.Network = v
	return s
}

func (s *ModifyBrowserInstanceGroupRequest) SetPolicy(v *ModifyBrowserInstanceGroupRequestPolicy) *ModifyBrowserInstanceGroupRequest {
	s.Policy = v
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
	// The bookmark.
	Bookmarks []*ModifyBrowserInstanceGroupRequestBrowserConfigBookmarks `json:"Bookmarks,omitempty" xml:"Bookmarks,omitempty" type:"Repeated"`
	// The startup parameter.
	//
	// example:
	//
	// --incognito
	BrowserParam *string `json:"BrowserParam,omitempty" xml:"BrowserParam,omitempty"`
	// The home page.
	//
	// example:
	//
	// https://www.aliyun.com
	Homepage *string `json:"Homepage,omitempty" xml:"Homepage,omitempty"`
	// The removed bookmarks.
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

func (s *ModifyBrowserInstanceGroupRequestBrowserConfig) GetBrowserParam() *string {
	return s.BrowserParam
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

func (s *ModifyBrowserInstanceGroupRequestBrowserConfig) SetBrowserParam(v string) *ModifyBrowserInstanceGroupRequestBrowserConfig {
	s.BrowserParam = &v
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
	// The folder where the bookmark is located.
	//
	// example:
	//
	// test
	BookmarkFolder *string `json:"BookmarkFolder,omitempty" xml:"BookmarkFolder,omitempty"`
	// The ID of the bookmark. This parameter needs to be specified only to modify the bookmark.
	//
	// example:
	//
	// bm-12345
	BookmarkId *string `json:"BookmarkId,omitempty" xml:"BookmarkId,omitempty"`
	// The name of the bookmark.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	BookmarkName *string `json:"BookmarkName,omitempty" xml:"BookmarkName,omitempty"`
	// The URL of the bookmark.
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
	// The type of the access control list.
	//
	// Valid value:
	//
	// 	- ALLOW_LIST: The whitelist.
	//
	// example:
	//
	// ALLOW_LIST
	AccessRestriction *string `json:"AccessRestriction,omitempty" xml:"AccessRestriction,omitempty"`
	// The domain names to be removed.
	RemoveRestrictedURLIds []*string `json:"RemoveRestrictedURLIds,omitempty" xml:"RemoveRestrictedURLIds,omitempty" type:"Repeated"`
	// The domain restriction configurations.
	RestrictedURLs []*ModifyBrowserInstanceGroupRequestNetworkRestrictedURLs `json:"RestrictedURLs,omitempty" xml:"RestrictedURLs,omitempty" type:"Repeated"`
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
	// The ID of the domain name. This parameter is required only when you want to modify the domain restriction configuration.
	//
	// example:
	//
	// ru-12345
	RestrictedURLId *string `json:"RestrictedURLId,omitempty" xml:"RestrictedURLId,omitempty"`
	// The restricted domain name.
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
	// The settings related to clipboard control.
	ClipboardPolicy *ModifyBrowserInstanceGroupRequestPolicyClipboardPolicy `json:"ClipboardPolicy,omitempty" xml:"ClipboardPolicy,omitempty" type:"Struct"`
	// Defines what happens to a session when a user disconnects.
	//
	// Valid values:
	//
	// 	- customTime: The session will be terminated after a custom-defined timeout.
	//
	// 	- persistent: The session will never be automatically terminated..
	//
	// example:
	//
	// customTime
	DisconnectKeepSession *string `json:"DisconnectKeepSession,omitempty" xml:"DisconnectKeepSession,omitempty"`
	// The session persistence duration.
	//
	// example:
	//
	// 15
	DisconnectKeepSessionTime *int32 `json:"DisconnectKeepSessionTime,omitempty" xml:"DisconnectKeepSessionTime,omitempty"`
	// The file transfer policy on the web client.
	//
	// example:
	//
	// off
	Html5FileTransfer         *string `json:"Html5FileTransfer,omitempty" xml:"Html5FileTransfer,omitempty"`
	NoOperationDisconnect     *string `json:"NoOperationDisconnect,omitempty" xml:"NoOperationDisconnect,omitempty"`
	NoOperationDisconnectTime *int32  `json:"NoOperationDisconnectTime,omitempty" xml:"NoOperationDisconnectTime,omitempty"`
	// The ID of the policy.
	//
	// example:
	//
	// pg-12345
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The policy version.
	//
	// Valid value:
	//
	// 	- Center: center policy
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

func (s *ModifyBrowserInstanceGroupRequestPolicy) GetClipboardPolicy() *ModifyBrowserInstanceGroupRequestPolicyClipboardPolicy {
	return s.ClipboardPolicy
}

func (s *ModifyBrowserInstanceGroupRequestPolicy) GetDisconnectKeepSession() *string {
	return s.DisconnectKeepSession
}

func (s *ModifyBrowserInstanceGroupRequestPolicy) GetDisconnectKeepSessionTime() *int32 {
	return s.DisconnectKeepSessionTime
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

type ModifyBrowserInstanceGroupRequestPolicyClipboardPolicy struct {
	// The clipboard policy.
	//
	// Valid values:
	//
	// 	- read: Allows copying from the local device to the cloud browser.
	//
	// 	- readwrite: Allows copying in both directions.
	//
	// 	- write: Allows copying from the cloud browser to the local device.
	//
	// 	- off: Blocks copying in both directions.
	//
	// example:
	//
	// off
	Clipboard *string `json:"Clipboard,omitempty" xml:"Clipboard,omitempty"`
	// The maximum number of characters allowed when copying from the clipboard.
	//
	// example:
	//
	// 1000
	ClipboardReadLimit *int32 `json:"ClipboardReadLimit,omitempty" xml:"ClipboardReadLimit,omitempty"`
	// The clipboard control scope.
	//
	// Valid values:
	//
	// 	- grained: fine-grained control
	//
	// 	- global: global control
	//
	// example:
	//
	// global
	ClipboardScope    *string `json:"ClipboardScope,omitempty" xml:"ClipboardScope,omitempty"`
	ClipboardSizeUnit *string `json:"ClipboardSizeUnit,omitempty" xml:"ClipboardSizeUnit,omitempty"`
	// The maximum number of characters allowed when copying to the clipboard.
	//
	// example:
	//
	// 1000
	ClipboardWriteLimit *int32 `json:"ClipboardWriteLimit,omitempty" xml:"ClipboardWriteLimit,omitempty"`
	// The file clipboard policy.
	//
	// Valid values:
	//
	// 	- read: Allows copying from the local device to the cloud browser.
	//
	// 	- readwrite: Allows copying in both directions.
	//
	// 	- write: Allows copying from the cloud browser to the local device.
	//
	// 	- off: Blocks copying in both directions.
	//
	// example:
	//
	// off
	FileClipboard *string `json:"FileClipboard,omitempty" xml:"FileClipboard,omitempty"`
	// The rich text clipboard policy.
	//
	// Valid values:
	//
	// 	- read: Allows copying from the local device to the cloud browser.
	//
	// 	- readwrite: Allows copying in both directions.
	//
	// 	- write: Allows copying from the cloud browser to the local device.
	//
	// 	- off: Blocks copying in both directions.
	//
	// example:
	//
	// off
	RichTextClipboard *string `json:"RichTextClipboard,omitempty" xml:"RichTextClipboard,omitempty"`
	// The text clipboard policy.
	//
	// Valid values:
	//
	// 	- read: Allows copying from the local device to the cloud browser.
	//
	// 	- readwrite: Allows copying in both directions.
	//
	// 	- write: Allows copying from the cloud browser to the local device.
	//
	// 	- off: Blocks copying in both directions.
	//
	// example:
	//
	// off
	TextClipboard *string `json:"TextClipboard,omitempty" xml:"TextClipboard,omitempty"`
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

func (s *ModifyBrowserInstanceGroupRequestPolicyClipboardPolicy) GetTextClipboard() *string {
	return s.TextClipboard
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

func (s *ModifyBrowserInstanceGroupRequestPolicyClipboardPolicy) SetTextClipboard(v string) *ModifyBrowserInstanceGroupRequestPolicyClipboardPolicy {
	s.TextClipboard = &v
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
	// Specifies whether to enable the watermark.
	//
	// Valid values:
	//
	// 	- off
	//
	// 	- on
	//
	// example:
	//
	// off
	WatermarkSwitch *string `json:"WatermarkSwitch,omitempty" xml:"WatermarkSwitch,omitempty"`
	// The watermark types.
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

type ModifyBrowserInstanceGroupRequestTimers struct {
	// The interval.
	//
	// example:
	//
	// 15
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The timer type:
	//
	// Valid value:
	//
	// 	- SESSION_TIMEOUT: Defines the timeout period before a disconnected session is terminated.
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
