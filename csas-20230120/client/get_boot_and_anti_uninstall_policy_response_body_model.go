// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBootAndAntiUninstallPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetBootAndAntiUninstallPolicyResponseBody
	GetRequestId() *string
	SetStrategy(v *GetBootAndAntiUninstallPolicyResponseBodyStrategy) *GetBootAndAntiUninstallPolicyResponseBody
	GetStrategy() *GetBootAndAntiUninstallPolicyResponseBodyStrategy
}

type GetBootAndAntiUninstallPolicyResponseBody struct {
	// example:
	//
	// CB67D866-1E54-5106-89DF-6D70C73E5989
	RequestId *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Strategy  *GetBootAndAntiUninstallPolicyResponseBodyStrategy `json:"Strategy,omitempty" xml:"Strategy,omitempty" type:"Struct"`
}

func (s GetBootAndAntiUninstallPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBootAndAntiUninstallPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *GetBootAndAntiUninstallPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetBootAndAntiUninstallPolicyResponseBody) GetStrategy() *GetBootAndAntiUninstallPolicyResponseBodyStrategy {
	return s.Strategy
}

func (s *GetBootAndAntiUninstallPolicyResponseBody) SetRequestId(v string) *GetBootAndAntiUninstallPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBootAndAntiUninstallPolicyResponseBody) SetStrategy(v *GetBootAndAntiUninstallPolicyResponseBodyStrategy) *GetBootAndAntiUninstallPolicyResponseBody {
	s.Strategy = v
	return s
}

func (s *GetBootAndAntiUninstallPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetBootAndAntiUninstallPolicyResponseBodyStrategy struct {
	// example:
	//
	// true
	AllowReport  *bool                                                          `json:"AllowReport,omitempty" xml:"AllowReport,omitempty"`
	BlockContent *GetBootAndAntiUninstallPolicyResponseBodyStrategyBlockContent `json:"BlockContent,omitempty" xml:"BlockContent,omitempty" type:"Struct"`
	// example:
	//
	// 2023-04-16 10:50:05
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// true
	IsAntiUninstall *bool `json:"IsAntiUninstall,omitempty" xml:"IsAntiUninstall,omitempty"`
	// example:
	//
	// true
	IsBoot *bool `json:"IsBoot,omitempty" xml:"IsBoot,omitempty"`
	// example:
	//
	// auto-boot-anti-uninstall-6f6cbf5f6605****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// example:
	//
	// approval-process-300abfb970cc****
	ReportProcessId *string `json:"ReportProcessId,omitempty" xml:"ReportProcessId,omitempty"`
	// example:
	//
	// 2024-06-14 10:17:14
	UpdateTime     *string   `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UserGroupIds   []*string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty" type:"Repeated"`
	WhitelistUsers []*string `json:"WhitelistUsers,omitempty" xml:"WhitelistUsers,omitempty" type:"Repeated"`
}

func (s GetBootAndAntiUninstallPolicyResponseBodyStrategy) String() string {
	return dara.Prettify(s)
}

func (s GetBootAndAntiUninstallPolicyResponseBodyStrategy) GoString() string {
	return s.String()
}

func (s *GetBootAndAntiUninstallPolicyResponseBodyStrategy) GetAllowReport() *bool {
	return s.AllowReport
}

func (s *GetBootAndAntiUninstallPolicyResponseBodyStrategy) GetBlockContent() *GetBootAndAntiUninstallPolicyResponseBodyStrategyBlockContent {
	return s.BlockContent
}

func (s *GetBootAndAntiUninstallPolicyResponseBodyStrategy) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetBootAndAntiUninstallPolicyResponseBodyStrategy) GetIsAntiUninstall() *bool {
	return s.IsAntiUninstall
}

func (s *GetBootAndAntiUninstallPolicyResponseBodyStrategy) GetIsBoot() *bool {
	return s.IsBoot
}

func (s *GetBootAndAntiUninstallPolicyResponseBodyStrategy) GetPolicyId() *string {
	return s.PolicyId
}

func (s *GetBootAndAntiUninstallPolicyResponseBodyStrategy) GetReportProcessId() *string {
	return s.ReportProcessId
}

func (s *GetBootAndAntiUninstallPolicyResponseBodyStrategy) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetBootAndAntiUninstallPolicyResponseBodyStrategy) GetUserGroupIds() []*string {
	return s.UserGroupIds
}

func (s *GetBootAndAntiUninstallPolicyResponseBodyStrategy) GetWhitelistUsers() []*string {
	return s.WhitelistUsers
}

func (s *GetBootAndAntiUninstallPolicyResponseBodyStrategy) SetAllowReport(v bool) *GetBootAndAntiUninstallPolicyResponseBodyStrategy {
	s.AllowReport = &v
	return s
}

func (s *GetBootAndAntiUninstallPolicyResponseBodyStrategy) SetBlockContent(v *GetBootAndAntiUninstallPolicyResponseBodyStrategyBlockContent) *GetBootAndAntiUninstallPolicyResponseBodyStrategy {
	s.BlockContent = v
	return s
}

func (s *GetBootAndAntiUninstallPolicyResponseBodyStrategy) SetCreateTime(v string) *GetBootAndAntiUninstallPolicyResponseBodyStrategy {
	s.CreateTime = &v
	return s
}

func (s *GetBootAndAntiUninstallPolicyResponseBodyStrategy) SetIsAntiUninstall(v bool) *GetBootAndAntiUninstallPolicyResponseBodyStrategy {
	s.IsAntiUninstall = &v
	return s
}

func (s *GetBootAndAntiUninstallPolicyResponseBodyStrategy) SetIsBoot(v bool) *GetBootAndAntiUninstallPolicyResponseBodyStrategy {
	s.IsBoot = &v
	return s
}

func (s *GetBootAndAntiUninstallPolicyResponseBodyStrategy) SetPolicyId(v string) *GetBootAndAntiUninstallPolicyResponseBodyStrategy {
	s.PolicyId = &v
	return s
}

func (s *GetBootAndAntiUninstallPolicyResponseBodyStrategy) SetReportProcessId(v string) *GetBootAndAntiUninstallPolicyResponseBodyStrategy {
	s.ReportProcessId = &v
	return s
}

func (s *GetBootAndAntiUninstallPolicyResponseBodyStrategy) SetUpdateTime(v string) *GetBootAndAntiUninstallPolicyResponseBodyStrategy {
	s.UpdateTime = &v
	return s
}

func (s *GetBootAndAntiUninstallPolicyResponseBodyStrategy) SetUserGroupIds(v []*string) *GetBootAndAntiUninstallPolicyResponseBodyStrategy {
	s.UserGroupIds = v
	return s
}

func (s *GetBootAndAntiUninstallPolicyResponseBodyStrategy) SetWhitelistUsers(v []*string) *GetBootAndAntiUninstallPolicyResponseBodyStrategy {
	s.WhitelistUsers = v
	return s
}

func (s *GetBootAndAntiUninstallPolicyResponseBodyStrategy) Validate() error {
	return dara.Validate(s)
}

type GetBootAndAntiUninstallPolicyResponseBodyStrategyBlockContent struct {
	BlockTextEn *GetBootAndAntiUninstallPolicyResponseBodyStrategyBlockContentBlockTextEn `json:"BlockTextEn,omitempty" xml:"BlockTextEn,omitempty" type:"Struct"`
	BlockTextZh *GetBootAndAntiUninstallPolicyResponseBodyStrategyBlockContentBlockTextZh `json:"BlockTextZh,omitempty" xml:"BlockTextZh,omitempty" type:"Struct"`
}

func (s GetBootAndAntiUninstallPolicyResponseBodyStrategyBlockContent) String() string {
	return dara.Prettify(s)
}

func (s GetBootAndAntiUninstallPolicyResponseBodyStrategyBlockContent) GoString() string {
	return s.String()
}

func (s *GetBootAndAntiUninstallPolicyResponseBodyStrategyBlockContent) GetBlockTextEn() *GetBootAndAntiUninstallPolicyResponseBodyStrategyBlockContentBlockTextEn {
	return s.BlockTextEn
}

func (s *GetBootAndAntiUninstallPolicyResponseBodyStrategyBlockContent) GetBlockTextZh() *GetBootAndAntiUninstallPolicyResponseBodyStrategyBlockContentBlockTextZh {
	return s.BlockTextZh
}

func (s *GetBootAndAntiUninstallPolicyResponseBodyStrategyBlockContent) SetBlockTextEn(v *GetBootAndAntiUninstallPolicyResponseBodyStrategyBlockContentBlockTextEn) *GetBootAndAntiUninstallPolicyResponseBodyStrategyBlockContent {
	s.BlockTextEn = v
	return s
}

func (s *GetBootAndAntiUninstallPolicyResponseBodyStrategyBlockContent) SetBlockTextZh(v *GetBootAndAntiUninstallPolicyResponseBodyStrategyBlockContentBlockTextZh) *GetBootAndAntiUninstallPolicyResponseBodyStrategyBlockContent {
	s.BlockTextZh = v
	return s
}

func (s *GetBootAndAntiUninstallPolicyResponseBodyStrategyBlockContent) Validate() error {
	return dara.Validate(s)
}

type GetBootAndAntiUninstallPolicyResponseBodyStrategyBlockContentBlockTextEn struct {
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

func (s GetBootAndAntiUninstallPolicyResponseBodyStrategyBlockContentBlockTextEn) String() string {
	return dara.Prettify(s)
}

func (s GetBootAndAntiUninstallPolicyResponseBodyStrategyBlockContentBlockTextEn) GoString() string {
	return s.String()
}

func (s *GetBootAndAntiUninstallPolicyResponseBodyStrategyBlockContentBlockTextEn) GetContent() *string {
	return s.Content
}

func (s *GetBootAndAntiUninstallPolicyResponseBodyStrategyBlockContentBlockTextEn) GetMainButtonText() *string {
	return s.MainButtonText
}

func (s *GetBootAndAntiUninstallPolicyResponseBodyStrategyBlockContentBlockTextEn) GetMinorButtonText() *string {
	return s.MinorButtonText
}

func (s *GetBootAndAntiUninstallPolicyResponseBodyStrategyBlockContentBlockTextEn) GetTitle() *string {
	return s.Title
}

func (s *GetBootAndAntiUninstallPolicyResponseBodyStrategyBlockContentBlockTextEn) SetContent(v string) *GetBootAndAntiUninstallPolicyResponseBodyStrategyBlockContentBlockTextEn {
	s.Content = &v
	return s
}

func (s *GetBootAndAntiUninstallPolicyResponseBodyStrategyBlockContentBlockTextEn) SetMainButtonText(v string) *GetBootAndAntiUninstallPolicyResponseBodyStrategyBlockContentBlockTextEn {
	s.MainButtonText = &v
	return s
}

func (s *GetBootAndAntiUninstallPolicyResponseBodyStrategyBlockContentBlockTextEn) SetMinorButtonText(v string) *GetBootAndAntiUninstallPolicyResponseBodyStrategyBlockContentBlockTextEn {
	s.MinorButtonText = &v
	return s
}

func (s *GetBootAndAntiUninstallPolicyResponseBodyStrategyBlockContentBlockTextEn) SetTitle(v string) *GetBootAndAntiUninstallPolicyResponseBodyStrategyBlockContentBlockTextEn {
	s.Title = &v
	return s
}

func (s *GetBootAndAntiUninstallPolicyResponseBodyStrategyBlockContentBlockTextEn) Validate() error {
	return dara.Validate(s)
}

type GetBootAndAntiUninstallPolicyResponseBodyStrategyBlockContentBlockTextZh struct {
	Content         *string `json:"Content,omitempty" xml:"Content,omitempty"`
	MainButtonText  *string `json:"MainButtonText,omitempty" xml:"MainButtonText,omitempty"`
	MinorButtonText *string `json:"MinorButtonText,omitempty" xml:"MinorButtonText,omitempty"`
	Title           *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetBootAndAntiUninstallPolicyResponseBodyStrategyBlockContentBlockTextZh) String() string {
	return dara.Prettify(s)
}

func (s GetBootAndAntiUninstallPolicyResponseBodyStrategyBlockContentBlockTextZh) GoString() string {
	return s.String()
}

func (s *GetBootAndAntiUninstallPolicyResponseBodyStrategyBlockContentBlockTextZh) GetContent() *string {
	return s.Content
}

func (s *GetBootAndAntiUninstallPolicyResponseBodyStrategyBlockContentBlockTextZh) GetMainButtonText() *string {
	return s.MainButtonText
}

func (s *GetBootAndAntiUninstallPolicyResponseBodyStrategyBlockContentBlockTextZh) GetMinorButtonText() *string {
	return s.MinorButtonText
}

func (s *GetBootAndAntiUninstallPolicyResponseBodyStrategyBlockContentBlockTextZh) GetTitle() *string {
	return s.Title
}

func (s *GetBootAndAntiUninstallPolicyResponseBodyStrategyBlockContentBlockTextZh) SetContent(v string) *GetBootAndAntiUninstallPolicyResponseBodyStrategyBlockContentBlockTextZh {
	s.Content = &v
	return s
}

func (s *GetBootAndAntiUninstallPolicyResponseBodyStrategyBlockContentBlockTextZh) SetMainButtonText(v string) *GetBootAndAntiUninstallPolicyResponseBodyStrategyBlockContentBlockTextZh {
	s.MainButtonText = &v
	return s
}

func (s *GetBootAndAntiUninstallPolicyResponseBodyStrategyBlockContentBlockTextZh) SetMinorButtonText(v string) *GetBootAndAntiUninstallPolicyResponseBodyStrategyBlockContentBlockTextZh {
	s.MinorButtonText = &v
	return s
}

func (s *GetBootAndAntiUninstallPolicyResponseBodyStrategyBlockContentBlockTextZh) SetTitle(v string) *GetBootAndAntiUninstallPolicyResponseBodyStrategyBlockContentBlockTextZh {
	s.Title = &v
	return s
}

func (s *GetBootAndAntiUninstallPolicyResponseBodyStrategyBlockContentBlockTextZh) Validate() error {
	return dara.Validate(s)
}
