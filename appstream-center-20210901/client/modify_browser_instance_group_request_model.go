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
	BrowserConfig *ModifyBrowserInstanceGroupRequestBrowserConfig `json:"BrowserConfig,omitempty" xml:"BrowserConfig,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// big-0bz55ixxxxx9xi9w9
	BrowserInstanceGroupId *string `json:"BrowserInstanceGroupId,omitempty" xml:"BrowserInstanceGroupId,omitempty"`
	// example:
	//
	// BrowserTest
	CloudBrowserName *string                                    `json:"CloudBrowserName,omitempty" xml:"CloudBrowserName,omitempty"`
	Network          *ModifyBrowserInstanceGroupRequestNetwork  `json:"Network,omitempty" xml:"Network,omitempty" type:"Struct"`
	Policy           *ModifyBrowserInstanceGroupRequestPolicy   `json:"Policy,omitempty" xml:"Policy,omitempty" type:"Struct"`
	Timers           []*ModifyBrowserInstanceGroupRequestTimers `json:"Timers,omitempty" xml:"Timers,omitempty" type:"Repeated"`
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
	Bookmarks []*ModifyBrowserInstanceGroupRequestBrowserConfigBookmarks `json:"Bookmarks,omitempty" xml:"Bookmarks,omitempty" type:"Repeated"`
	// example:
	//
	// --incognito
	BrowserParam *string `json:"BrowserParam,omitempty" xml:"BrowserParam,omitempty"`
	// example:
	//
	// https://www.aliyun.com
	Homepage        *string   `json:"Homepage,omitempty" xml:"Homepage,omitempty"`
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
	// example:
	//
	// test
	BookmarkFolder *string `json:"BookmarkFolder,omitempty" xml:"BookmarkFolder,omitempty"`
	// example:
	//
	// bm-12345
	BookmarkId *string `json:"BookmarkId,omitempty" xml:"BookmarkId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test
	BookmarkName *string `json:"BookmarkName,omitempty" xml:"BookmarkName,omitempty"`
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
	// example:
	//
	// ALLOW_LIST
	AccessRestriction      *string                                                   `json:"AccessRestriction,omitempty" xml:"AccessRestriction,omitempty"`
	RemoveRestrictedURLIds []*string                                                 `json:"RemoveRestrictedURLIds,omitempty" xml:"RemoveRestrictedURLIds,omitempty" type:"Repeated"`
	RestrictedURLs         []*ModifyBrowserInstanceGroupRequestNetworkRestrictedURLs `json:"RestrictedURLs,omitempty" xml:"RestrictedURLs,omitempty" type:"Repeated"`
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
	// example:
	//
	// ru-12345
	RestrictedURLId *string `json:"RestrictedURLId,omitempty" xml:"RestrictedURLId,omitempty"`
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
	ClipboardPolicy *ModifyBrowserInstanceGroupRequestPolicyClipboardPolicy `json:"ClipboardPolicy,omitempty" xml:"ClipboardPolicy,omitempty" type:"Struct"`
	// example:
	//
	// customTime
	DisconnectKeepSession *string `json:"DisconnectKeepSession,omitempty" xml:"DisconnectKeepSession,omitempty"`
	// example:
	//
	// 15
	DisconnectKeepSessionTime *int32 `json:"DisconnectKeepSessionTime,omitempty" xml:"DisconnectKeepSessionTime,omitempty"`
	// example:
	//
	// off
	Html5FileTransfer *string `json:"Html5FileTransfer,omitempty" xml:"Html5FileTransfer,omitempty"`
	// example:
	//
	// pg-12345
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// example:
	//
	// Center
	PolicyVersion   *string                                                 `json:"PolicyVersion,omitempty" xml:"PolicyVersion,omitempty"`
	VideoPolicy     *ModifyBrowserInstanceGroupRequestPolicyVideoPolicy     `json:"VideoPolicy,omitempty" xml:"VideoPolicy,omitempty" type:"Struct"`
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
	// example:
	//
	// off
	Clipboard *string `json:"Clipboard,omitempty" xml:"Clipboard,omitempty"`
	// example:
	//
	// 1000
	ClipboardReadLimit *int32 `json:"ClipboardReadLimit,omitempty" xml:"ClipboardReadLimit,omitempty"`
	// example:
	//
	// global
	ClipboardScope *string `json:"ClipboardScope,omitempty" xml:"ClipboardScope,omitempty"`
	// example:
	//
	// 1000
	ClipboardWriteLimit *int32 `json:"ClipboardWriteLimit,omitempty" xml:"ClipboardWriteLimit,omitempty"`
	// example:
	//
	// off
	FileClipboard *string `json:"FileClipboard,omitempty" xml:"FileClipboard,omitempty"`
	// example:
	//
	// off
	RichTextClipboard *string `json:"RichTextClipboard,omitempty" xml:"RichTextClipboard,omitempty"`
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
	// example:
	//
	// off
	WatermarkSwitch *string   `json:"WatermarkSwitch,omitempty" xml:"WatermarkSwitch,omitempty"`
	WatermarkTypes  []*string `json:"WatermarkTypes,omitempty" xml:"WatermarkTypes,omitempty" type:"Repeated"`
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
	// example:
	//
	// 15
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
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
