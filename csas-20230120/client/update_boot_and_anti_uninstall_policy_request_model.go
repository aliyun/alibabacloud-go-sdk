// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBootAndAntiUninstallPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllowReport(v bool) *UpdateBootAndAntiUninstallPolicyRequest
	GetAllowReport() *bool
	SetBlockContent(v *UpdateBootAndAntiUninstallPolicyRequestBlockContent) *UpdateBootAndAntiUninstallPolicyRequest
	GetBlockContent() *UpdateBootAndAntiUninstallPolicyRequestBlockContent
	SetIsAntiUninstall(v bool) *UpdateBootAndAntiUninstallPolicyRequest
	GetIsAntiUninstall() *bool
	SetIsBoot(v bool) *UpdateBootAndAntiUninstallPolicyRequest
	GetIsBoot() *bool
	SetUserGroupIds(v []*string) *UpdateBootAndAntiUninstallPolicyRequest
	GetUserGroupIds() []*string
	SetWhitelistUsers(v []*string) *UpdateBootAndAntiUninstallPolicyRequest
	GetWhitelistUsers() []*string
}

type UpdateBootAndAntiUninstallPolicyRequest struct {
	// example:
	//
	// true
	AllowReport  *bool                                                `json:"AllowReport,omitempty" xml:"AllowReport,omitempty"`
	BlockContent *UpdateBootAndAntiUninstallPolicyRequestBlockContent `json:"BlockContent,omitempty" xml:"BlockContent,omitempty" type:"Struct"`
	// example:
	//
	// true
	IsAntiUninstall *bool `json:"IsAntiUninstall,omitempty" xml:"IsAntiUninstall,omitempty"`
	// example:
	//
	// true
	IsBoot         *bool     `json:"IsBoot,omitempty" xml:"IsBoot,omitempty"`
	UserGroupIds   []*string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty" type:"Repeated"`
	WhitelistUsers []*string `json:"WhitelistUsers,omitempty" xml:"WhitelistUsers,omitempty" type:"Repeated"`
}

func (s UpdateBootAndAntiUninstallPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateBootAndAntiUninstallPolicyRequest) GoString() string {
	return s.String()
}

func (s *UpdateBootAndAntiUninstallPolicyRequest) GetAllowReport() *bool {
	return s.AllowReport
}

func (s *UpdateBootAndAntiUninstallPolicyRequest) GetBlockContent() *UpdateBootAndAntiUninstallPolicyRequestBlockContent {
	return s.BlockContent
}

func (s *UpdateBootAndAntiUninstallPolicyRequest) GetIsAntiUninstall() *bool {
	return s.IsAntiUninstall
}

func (s *UpdateBootAndAntiUninstallPolicyRequest) GetIsBoot() *bool {
	return s.IsBoot
}

func (s *UpdateBootAndAntiUninstallPolicyRequest) GetUserGroupIds() []*string {
	return s.UserGroupIds
}

func (s *UpdateBootAndAntiUninstallPolicyRequest) GetWhitelistUsers() []*string {
	return s.WhitelistUsers
}

func (s *UpdateBootAndAntiUninstallPolicyRequest) SetAllowReport(v bool) *UpdateBootAndAntiUninstallPolicyRequest {
	s.AllowReport = &v
	return s
}

func (s *UpdateBootAndAntiUninstallPolicyRequest) SetBlockContent(v *UpdateBootAndAntiUninstallPolicyRequestBlockContent) *UpdateBootAndAntiUninstallPolicyRequest {
	s.BlockContent = v
	return s
}

func (s *UpdateBootAndAntiUninstallPolicyRequest) SetIsAntiUninstall(v bool) *UpdateBootAndAntiUninstallPolicyRequest {
	s.IsAntiUninstall = &v
	return s
}

func (s *UpdateBootAndAntiUninstallPolicyRequest) SetIsBoot(v bool) *UpdateBootAndAntiUninstallPolicyRequest {
	s.IsBoot = &v
	return s
}

func (s *UpdateBootAndAntiUninstallPolicyRequest) SetUserGroupIds(v []*string) *UpdateBootAndAntiUninstallPolicyRequest {
	s.UserGroupIds = v
	return s
}

func (s *UpdateBootAndAntiUninstallPolicyRequest) SetWhitelistUsers(v []*string) *UpdateBootAndAntiUninstallPolicyRequest {
	s.WhitelistUsers = v
	return s
}

func (s *UpdateBootAndAntiUninstallPolicyRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateBootAndAntiUninstallPolicyRequestBlockContent struct {
	BlockTextEn *UpdateBootAndAntiUninstallPolicyRequestBlockContentBlockTextEn `json:"BlockTextEn,omitempty" xml:"BlockTextEn,omitempty" type:"Struct"`
	BlockTextZh *UpdateBootAndAntiUninstallPolicyRequestBlockContentBlockTextZh `json:"BlockTextZh,omitempty" xml:"BlockTextZh,omitempty" type:"Struct"`
}

func (s UpdateBootAndAntiUninstallPolicyRequestBlockContent) String() string {
	return dara.Prettify(s)
}

func (s UpdateBootAndAntiUninstallPolicyRequestBlockContent) GoString() string {
	return s.String()
}

func (s *UpdateBootAndAntiUninstallPolicyRequestBlockContent) GetBlockTextEn() *UpdateBootAndAntiUninstallPolicyRequestBlockContentBlockTextEn {
	return s.BlockTextEn
}

func (s *UpdateBootAndAntiUninstallPolicyRequestBlockContent) GetBlockTextZh() *UpdateBootAndAntiUninstallPolicyRequestBlockContentBlockTextZh {
	return s.BlockTextZh
}

func (s *UpdateBootAndAntiUninstallPolicyRequestBlockContent) SetBlockTextEn(v *UpdateBootAndAntiUninstallPolicyRequestBlockContentBlockTextEn) *UpdateBootAndAntiUninstallPolicyRequestBlockContent {
	s.BlockTextEn = v
	return s
}

func (s *UpdateBootAndAntiUninstallPolicyRequestBlockContent) SetBlockTextZh(v *UpdateBootAndAntiUninstallPolicyRequestBlockContentBlockTextZh) *UpdateBootAndAntiUninstallPolicyRequestBlockContent {
	s.BlockTextZh = v
	return s
}

func (s *UpdateBootAndAntiUninstallPolicyRequestBlockContent) Validate() error {
	return dara.Validate(s)
}

type UpdateBootAndAntiUninstallPolicyRequestBlockContentBlockTextEn struct {
	// example:
	//
	// After uninstalling, the device can no longer be used for company work, and it will lose access to the company\\"s intranet!
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// Report
	MainButtonText *string `json:"MainButtonText,omitempty" xml:"MainButtonText,omitempty"`
	// example:
	//
	// Ignore
	MinorButtonText *string `json:"MinorButtonText,omitempty" xml:"MinorButtonText,omitempty"`
	// example:
	//
	// Anti-Uninstall Warning
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s UpdateBootAndAntiUninstallPolicyRequestBlockContentBlockTextEn) String() string {
	return dara.Prettify(s)
}

func (s UpdateBootAndAntiUninstallPolicyRequestBlockContentBlockTextEn) GoString() string {
	return s.String()
}

func (s *UpdateBootAndAntiUninstallPolicyRequestBlockContentBlockTextEn) GetContent() *string {
	return s.Content
}

func (s *UpdateBootAndAntiUninstallPolicyRequestBlockContentBlockTextEn) GetMainButtonText() *string {
	return s.MainButtonText
}

func (s *UpdateBootAndAntiUninstallPolicyRequestBlockContentBlockTextEn) GetMinorButtonText() *string {
	return s.MinorButtonText
}

func (s *UpdateBootAndAntiUninstallPolicyRequestBlockContentBlockTextEn) GetTitle() *string {
	return s.Title
}

func (s *UpdateBootAndAntiUninstallPolicyRequestBlockContentBlockTextEn) SetContent(v string) *UpdateBootAndAntiUninstallPolicyRequestBlockContentBlockTextEn {
	s.Content = &v
	return s
}

func (s *UpdateBootAndAntiUninstallPolicyRequestBlockContentBlockTextEn) SetMainButtonText(v string) *UpdateBootAndAntiUninstallPolicyRequestBlockContentBlockTextEn {
	s.MainButtonText = &v
	return s
}

func (s *UpdateBootAndAntiUninstallPolicyRequestBlockContentBlockTextEn) SetMinorButtonText(v string) *UpdateBootAndAntiUninstallPolicyRequestBlockContentBlockTextEn {
	s.MinorButtonText = &v
	return s
}

func (s *UpdateBootAndAntiUninstallPolicyRequestBlockContentBlockTextEn) SetTitle(v string) *UpdateBootAndAntiUninstallPolicyRequestBlockContentBlockTextEn {
	s.Title = &v
	return s
}

func (s *UpdateBootAndAntiUninstallPolicyRequestBlockContentBlockTextEn) Validate() error {
	return dara.Validate(s)
}

type UpdateBootAndAntiUninstallPolicyRequestBlockContentBlockTextZh struct {
	Content         *string `json:"Content,omitempty" xml:"Content,omitempty"`
	MainButtonText  *string `json:"MainButtonText,omitempty" xml:"MainButtonText,omitempty"`
	MinorButtonText *string `json:"MinorButtonText,omitempty" xml:"MinorButtonText,omitempty"`
	Title           *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s UpdateBootAndAntiUninstallPolicyRequestBlockContentBlockTextZh) String() string {
	return dara.Prettify(s)
}

func (s UpdateBootAndAntiUninstallPolicyRequestBlockContentBlockTextZh) GoString() string {
	return s.String()
}

func (s *UpdateBootAndAntiUninstallPolicyRequestBlockContentBlockTextZh) GetContent() *string {
	return s.Content
}

func (s *UpdateBootAndAntiUninstallPolicyRequestBlockContentBlockTextZh) GetMainButtonText() *string {
	return s.MainButtonText
}

func (s *UpdateBootAndAntiUninstallPolicyRequestBlockContentBlockTextZh) GetMinorButtonText() *string {
	return s.MinorButtonText
}

func (s *UpdateBootAndAntiUninstallPolicyRequestBlockContentBlockTextZh) GetTitle() *string {
	return s.Title
}

func (s *UpdateBootAndAntiUninstallPolicyRequestBlockContentBlockTextZh) SetContent(v string) *UpdateBootAndAntiUninstallPolicyRequestBlockContentBlockTextZh {
	s.Content = &v
	return s
}

func (s *UpdateBootAndAntiUninstallPolicyRequestBlockContentBlockTextZh) SetMainButtonText(v string) *UpdateBootAndAntiUninstallPolicyRequestBlockContentBlockTextZh {
	s.MainButtonText = &v
	return s
}

func (s *UpdateBootAndAntiUninstallPolicyRequestBlockContentBlockTextZh) SetMinorButtonText(v string) *UpdateBootAndAntiUninstallPolicyRequestBlockContentBlockTextZh {
	s.MinorButtonText = &v
	return s
}

func (s *UpdateBootAndAntiUninstallPolicyRequestBlockContentBlockTextZh) SetTitle(v string) *UpdateBootAndAntiUninstallPolicyRequestBlockContentBlockTextZh {
	s.Title = &v
	return s
}

func (s *UpdateBootAndAntiUninstallPolicyRequestBlockContentBlockTextZh) Validate() error {
	return dara.Validate(s)
}
