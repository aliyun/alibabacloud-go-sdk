// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPAApplicationUnauthorizedAccessConfig interface {
	dara.Model
	String() string
	GoString() string
	SetAllowReport(v bool) *PAApplicationUnauthorizedAccessConfig
	GetAllowReport() *bool
	SetBlockContent(v *PAApplicationUnauthorizedAccessConfigBlockContent) *PAApplicationUnauthorizedAccessConfig
	GetBlockContent() *PAApplicationUnauthorizedAccessConfigBlockContent
	SetEnabled(v bool) *PAApplicationUnauthorizedAccessConfig
	GetEnabled() *bool
	SetReportProcessId(v string) *PAApplicationUnauthorizedAccessConfig
	GetReportProcessId() *string
}

type PAApplicationUnauthorizedAccessConfig struct {
	// Specifies whether end users are allowed to submit approval requests.
	//
	// example:
	//
	// true
	AllowReport *bool `json:"AllowReport,omitempty" xml:"AllowReport,omitempty"`
	// The content displayed in the client interception pop-up window.
	BlockContent *PAApplicationUnauthorizedAccessConfigBlockContent `json:"BlockContent,omitempty" xml:"BlockContent,omitempty" type:"Struct"`
	// Specifies whether the feature is enabled. Valid values:
	//
	// - **true**: Enabled. Users are redirected to an interception page when they access an unauthorized application.
	//
	// - **false**: Disabled. An error message is returned by default when users access an unauthorized application.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The ID of the approval flow associated with the policy.
	//
	// example:
	//
	// approval-process-0ee84ac4f9c31bc5
	ReportProcessId *string `json:"ReportProcessId,omitempty" xml:"ReportProcessId,omitempty"`
}

func (s PAApplicationUnauthorizedAccessConfig) String() string {
	return dara.Prettify(s)
}

func (s PAApplicationUnauthorizedAccessConfig) GoString() string {
	return s.String()
}

func (s *PAApplicationUnauthorizedAccessConfig) GetAllowReport() *bool {
	return s.AllowReport
}

func (s *PAApplicationUnauthorizedAccessConfig) GetBlockContent() *PAApplicationUnauthorizedAccessConfigBlockContent {
	return s.BlockContent
}

func (s *PAApplicationUnauthorizedAccessConfig) GetEnabled() *bool {
	return s.Enabled
}

func (s *PAApplicationUnauthorizedAccessConfig) GetReportProcessId() *string {
	return s.ReportProcessId
}

func (s *PAApplicationUnauthorizedAccessConfig) SetAllowReport(v bool) *PAApplicationUnauthorizedAccessConfig {
	s.AllowReport = &v
	return s
}

func (s *PAApplicationUnauthorizedAccessConfig) SetBlockContent(v *PAApplicationUnauthorizedAccessConfigBlockContent) *PAApplicationUnauthorizedAccessConfig {
	s.BlockContent = v
	return s
}

func (s *PAApplicationUnauthorizedAccessConfig) SetEnabled(v bool) *PAApplicationUnauthorizedAccessConfig {
	s.Enabled = &v
	return s
}

func (s *PAApplicationUnauthorizedAccessConfig) SetReportProcessId(v string) *PAApplicationUnauthorizedAccessConfig {
	s.ReportProcessId = &v
	return s
}

func (s *PAApplicationUnauthorizedAccessConfig) Validate() error {
	if s.BlockContent != nil {
		if err := s.BlockContent.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PAApplicationUnauthorizedAccessConfigBlockContent struct {
	// The English content.
	BlockTextEn *PAApplicationUnauthorizedAccessConfigBlockContentBlockTextEn `json:"BlockTextEn,omitempty" xml:"BlockTextEn,omitempty" type:"Struct"`
	// The Chinese content.
	BlockTextZh *PAApplicationUnauthorizedAccessConfigBlockContentBlockTextZh `json:"BlockTextZh,omitempty" xml:"BlockTextZh,omitempty" type:"Struct"`
}

func (s PAApplicationUnauthorizedAccessConfigBlockContent) String() string {
	return dara.Prettify(s)
}

func (s PAApplicationUnauthorizedAccessConfigBlockContent) GoString() string {
	return s.String()
}

func (s *PAApplicationUnauthorizedAccessConfigBlockContent) GetBlockTextEn() *PAApplicationUnauthorizedAccessConfigBlockContentBlockTextEn {
	return s.BlockTextEn
}

func (s *PAApplicationUnauthorizedAccessConfigBlockContent) GetBlockTextZh() *PAApplicationUnauthorizedAccessConfigBlockContentBlockTextZh {
	return s.BlockTextZh
}

func (s *PAApplicationUnauthorizedAccessConfigBlockContent) SetBlockTextEn(v *PAApplicationUnauthorizedAccessConfigBlockContentBlockTextEn) *PAApplicationUnauthorizedAccessConfigBlockContent {
	s.BlockTextEn = v
	return s
}

func (s *PAApplicationUnauthorizedAccessConfigBlockContent) SetBlockTextZh(v *PAApplicationUnauthorizedAccessConfigBlockContentBlockTextZh) *PAApplicationUnauthorizedAccessConfigBlockContent {
	s.BlockTextZh = v
	return s
}

func (s *PAApplicationUnauthorizedAccessConfigBlockContent) Validate() error {
	if s.BlockTextEn != nil {
		if err := s.BlockTextEn.Validate(); err != nil {
			return err
		}
	}
	if s.BlockTextZh != nil {
		if err := s.BlockTextZh.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PAApplicationUnauthorizedAccessConfigBlockContentBlockTextEn struct {
	// The prompt content of the English block page.
	//
	// example:
	//
	// You do not have permission to access this system. If you need to do so, please submit a permission request.
	BrowserAlertContent *string `json:"BrowserAlertContent,omitempty" xml:"BrowserAlertContent,omitempty"`
	// The background pattern of the English block page.
	//
	// example:
	//
	// https://img.alicdn.com/xxx.png
	BrowserAlertStyle *string `json:"BrowserAlertStyle,omitempty" xml:"BrowserAlertStyle,omitempty"`
	// The title of the English block page.
	//
	// example:
	//
	// No Permission Access
	BrowserAlertTitle *string `json:"BrowserAlertTitle,omitempty" xml:"BrowserAlertTitle,omitempty"`
	// The text of the English report approval button.
	//
	// example:
	//
	// Report
	ReportButtonText *string `json:"ReportButtonText,omitempty" xml:"ReportButtonText,omitempty"`
}

func (s PAApplicationUnauthorizedAccessConfigBlockContentBlockTextEn) String() string {
	return dara.Prettify(s)
}

func (s PAApplicationUnauthorizedAccessConfigBlockContentBlockTextEn) GoString() string {
	return s.String()
}

func (s *PAApplicationUnauthorizedAccessConfigBlockContentBlockTextEn) GetBrowserAlertContent() *string {
	return s.BrowserAlertContent
}

func (s *PAApplicationUnauthorizedAccessConfigBlockContentBlockTextEn) GetBrowserAlertStyle() *string {
	return s.BrowserAlertStyle
}

func (s *PAApplicationUnauthorizedAccessConfigBlockContentBlockTextEn) GetBrowserAlertTitle() *string {
	return s.BrowserAlertTitle
}

func (s *PAApplicationUnauthorizedAccessConfigBlockContentBlockTextEn) GetReportButtonText() *string {
	return s.ReportButtonText
}

func (s *PAApplicationUnauthorizedAccessConfigBlockContentBlockTextEn) SetBrowserAlertContent(v string) *PAApplicationUnauthorizedAccessConfigBlockContentBlockTextEn {
	s.BrowserAlertContent = &v
	return s
}

func (s *PAApplicationUnauthorizedAccessConfigBlockContentBlockTextEn) SetBrowserAlertStyle(v string) *PAApplicationUnauthorizedAccessConfigBlockContentBlockTextEn {
	s.BrowserAlertStyle = &v
	return s
}

func (s *PAApplicationUnauthorizedAccessConfigBlockContentBlockTextEn) SetBrowserAlertTitle(v string) *PAApplicationUnauthorizedAccessConfigBlockContentBlockTextEn {
	s.BrowserAlertTitle = &v
	return s
}

func (s *PAApplicationUnauthorizedAccessConfigBlockContentBlockTextEn) SetReportButtonText(v string) *PAApplicationUnauthorizedAccessConfigBlockContentBlockTextEn {
	s.ReportButtonText = &v
	return s
}

func (s *PAApplicationUnauthorizedAccessConfigBlockContentBlockTextEn) Validate() error {
	return dara.Validate(s)
}

type PAApplicationUnauthorizedAccessConfigBlockContentBlockTextZh struct {
	// The prompt content of the block page.
	//
	// example:
	//
	// 您暂无权限访问该系统。如有工作需要，请提交权限申请。
	BrowserAlertContent *string `json:"BrowserAlertContent,omitempty" xml:"BrowserAlertContent,omitempty"`
	// The background pattern of the block page.
	//
	// example:
	//
	// https://img.alicdn.com/xxx.png
	BrowserAlertStyle *string `json:"BrowserAlertStyle,omitempty" xml:"BrowserAlertStyle,omitempty"`
	// The title of the block page.
	//
	// example:
	//
	// 无权限访问
	BrowserAlertTitle *string `json:"BrowserAlertTitle,omitempty" xml:"BrowserAlertTitle,omitempty"`
	// The text of the report approval button.
	//
	// example:
	//
	// 前往报备
	ReportButtonText *string `json:"ReportButtonText,omitempty" xml:"ReportButtonText,omitempty"`
}

func (s PAApplicationUnauthorizedAccessConfigBlockContentBlockTextZh) String() string {
	return dara.Prettify(s)
}

func (s PAApplicationUnauthorizedAccessConfigBlockContentBlockTextZh) GoString() string {
	return s.String()
}

func (s *PAApplicationUnauthorizedAccessConfigBlockContentBlockTextZh) GetBrowserAlertContent() *string {
	return s.BrowserAlertContent
}

func (s *PAApplicationUnauthorizedAccessConfigBlockContentBlockTextZh) GetBrowserAlertStyle() *string {
	return s.BrowserAlertStyle
}

func (s *PAApplicationUnauthorizedAccessConfigBlockContentBlockTextZh) GetBrowserAlertTitle() *string {
	return s.BrowserAlertTitle
}

func (s *PAApplicationUnauthorizedAccessConfigBlockContentBlockTextZh) GetReportButtonText() *string {
	return s.ReportButtonText
}

func (s *PAApplicationUnauthorizedAccessConfigBlockContentBlockTextZh) SetBrowserAlertContent(v string) *PAApplicationUnauthorizedAccessConfigBlockContentBlockTextZh {
	s.BrowserAlertContent = &v
	return s
}

func (s *PAApplicationUnauthorizedAccessConfigBlockContentBlockTextZh) SetBrowserAlertStyle(v string) *PAApplicationUnauthorizedAccessConfigBlockContentBlockTextZh {
	s.BrowserAlertStyle = &v
	return s
}

func (s *PAApplicationUnauthorizedAccessConfigBlockContentBlockTextZh) SetBrowserAlertTitle(v string) *PAApplicationUnauthorizedAccessConfigBlockContentBlockTextZh {
	s.BrowserAlertTitle = &v
	return s
}

func (s *PAApplicationUnauthorizedAccessConfigBlockContentBlockTextZh) SetReportButtonText(v string) *PAApplicationUnauthorizedAccessConfigBlockContentBlockTextZh {
	s.ReportButtonText = &v
	return s
}

func (s *PAApplicationUnauthorizedAccessConfigBlockContentBlockTextZh) Validate() error {
	return dara.Validate(s)
}
