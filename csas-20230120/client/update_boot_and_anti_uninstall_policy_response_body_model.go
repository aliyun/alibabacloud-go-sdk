// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBootAndAntiUninstallPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateBootAndAntiUninstallPolicyResponseBody
	GetRequestId() *string
	SetStrategy(v *UpdateBootAndAntiUninstallPolicyResponseBodyStrategy) *UpdateBootAndAntiUninstallPolicyResponseBody
	GetStrategy() *UpdateBootAndAntiUninstallPolicyResponseBodyStrategy
}

type UpdateBootAndAntiUninstallPolicyResponseBody struct {
	// example:
	//
	// CB67D866-1E54-5106-89DF-6D70C73E5989
	RequestId *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Strategy  *UpdateBootAndAntiUninstallPolicyResponseBodyStrategy `json:"Strategy,omitempty" xml:"Strategy,omitempty" type:"Struct"`
}

func (s UpdateBootAndAntiUninstallPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateBootAndAntiUninstallPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateBootAndAntiUninstallPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateBootAndAntiUninstallPolicyResponseBody) GetStrategy() *UpdateBootAndAntiUninstallPolicyResponseBodyStrategy {
	return s.Strategy
}

func (s *UpdateBootAndAntiUninstallPolicyResponseBody) SetRequestId(v string) *UpdateBootAndAntiUninstallPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateBootAndAntiUninstallPolicyResponseBody) SetStrategy(v *UpdateBootAndAntiUninstallPolicyResponseBodyStrategy) *UpdateBootAndAntiUninstallPolicyResponseBody {
	s.Strategy = v
	return s
}

func (s *UpdateBootAndAntiUninstallPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}

type UpdateBootAndAntiUninstallPolicyResponseBodyStrategy struct {
	// example:
	//
	// true
	AllowReport  *bool                                                             `json:"AllowReport,omitempty" xml:"AllowReport,omitempty"`
	BlockContent *UpdateBootAndAntiUninstallPolicyResponseBodyStrategyBlockContent `json:"BlockContent,omitempty" xml:"BlockContent,omitempty" type:"Struct"`
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
	// approval-process-65c255598826****
	ReportProcessId *string `json:"ReportProcessId,omitempty" xml:"ReportProcessId,omitempty"`
	// example:
	//
	// 2024-06-14 10:17:14
	UpdateTime     *string   `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UserGroupIds   []*string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty" type:"Repeated"`
	WhitelistUsers []*string `json:"WhitelistUsers,omitempty" xml:"WhitelistUsers,omitempty" type:"Repeated"`
}

func (s UpdateBootAndAntiUninstallPolicyResponseBodyStrategy) String() string {
	return dara.Prettify(s)
}

func (s UpdateBootAndAntiUninstallPolicyResponseBodyStrategy) GoString() string {
	return s.String()
}

func (s *UpdateBootAndAntiUninstallPolicyResponseBodyStrategy) GetAllowReport() *bool {
	return s.AllowReport
}

func (s *UpdateBootAndAntiUninstallPolicyResponseBodyStrategy) GetBlockContent() *UpdateBootAndAntiUninstallPolicyResponseBodyStrategyBlockContent {
	return s.BlockContent
}

func (s *UpdateBootAndAntiUninstallPolicyResponseBodyStrategy) GetCreateTime() *string {
	return s.CreateTime
}

func (s *UpdateBootAndAntiUninstallPolicyResponseBodyStrategy) GetIsAntiUninstall() *bool {
	return s.IsAntiUninstall
}

func (s *UpdateBootAndAntiUninstallPolicyResponseBodyStrategy) GetIsBoot() *bool {
	return s.IsBoot
}

func (s *UpdateBootAndAntiUninstallPolicyResponseBodyStrategy) GetPolicyId() *string {
	return s.PolicyId
}

func (s *UpdateBootAndAntiUninstallPolicyResponseBodyStrategy) GetReportProcessId() *string {
	return s.ReportProcessId
}

func (s *UpdateBootAndAntiUninstallPolicyResponseBodyStrategy) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *UpdateBootAndAntiUninstallPolicyResponseBodyStrategy) GetUserGroupIds() []*string {
	return s.UserGroupIds
}

func (s *UpdateBootAndAntiUninstallPolicyResponseBodyStrategy) GetWhitelistUsers() []*string {
	return s.WhitelistUsers
}

func (s *UpdateBootAndAntiUninstallPolicyResponseBodyStrategy) SetAllowReport(v bool) *UpdateBootAndAntiUninstallPolicyResponseBodyStrategy {
	s.AllowReport = &v
	return s
}

func (s *UpdateBootAndAntiUninstallPolicyResponseBodyStrategy) SetBlockContent(v *UpdateBootAndAntiUninstallPolicyResponseBodyStrategyBlockContent) *UpdateBootAndAntiUninstallPolicyResponseBodyStrategy {
	s.BlockContent = v
	return s
}

func (s *UpdateBootAndAntiUninstallPolicyResponseBodyStrategy) SetCreateTime(v string) *UpdateBootAndAntiUninstallPolicyResponseBodyStrategy {
	s.CreateTime = &v
	return s
}

func (s *UpdateBootAndAntiUninstallPolicyResponseBodyStrategy) SetIsAntiUninstall(v bool) *UpdateBootAndAntiUninstallPolicyResponseBodyStrategy {
	s.IsAntiUninstall = &v
	return s
}

func (s *UpdateBootAndAntiUninstallPolicyResponseBodyStrategy) SetIsBoot(v bool) *UpdateBootAndAntiUninstallPolicyResponseBodyStrategy {
	s.IsBoot = &v
	return s
}

func (s *UpdateBootAndAntiUninstallPolicyResponseBodyStrategy) SetPolicyId(v string) *UpdateBootAndAntiUninstallPolicyResponseBodyStrategy {
	s.PolicyId = &v
	return s
}

func (s *UpdateBootAndAntiUninstallPolicyResponseBodyStrategy) SetReportProcessId(v string) *UpdateBootAndAntiUninstallPolicyResponseBodyStrategy {
	s.ReportProcessId = &v
	return s
}

func (s *UpdateBootAndAntiUninstallPolicyResponseBodyStrategy) SetUpdateTime(v string) *UpdateBootAndAntiUninstallPolicyResponseBodyStrategy {
	s.UpdateTime = &v
	return s
}

func (s *UpdateBootAndAntiUninstallPolicyResponseBodyStrategy) SetUserGroupIds(v []*string) *UpdateBootAndAntiUninstallPolicyResponseBodyStrategy {
	s.UserGroupIds = v
	return s
}

func (s *UpdateBootAndAntiUninstallPolicyResponseBodyStrategy) SetWhitelistUsers(v []*string) *UpdateBootAndAntiUninstallPolicyResponseBodyStrategy {
	s.WhitelistUsers = v
	return s
}

func (s *UpdateBootAndAntiUninstallPolicyResponseBodyStrategy) Validate() error {
	return dara.Validate(s)
}

type UpdateBootAndAntiUninstallPolicyResponseBodyStrategyBlockContent struct {
	BlockTextEn *UpdateBootAndAntiUninstallPolicyResponseBodyStrategyBlockContentBlockTextEn `json:"BlockTextEn,omitempty" xml:"BlockTextEn,omitempty" type:"Struct"`
	BlockTextZh *UpdateBootAndAntiUninstallPolicyResponseBodyStrategyBlockContentBlockTextZh `json:"BlockTextZh,omitempty" xml:"BlockTextZh,omitempty" type:"Struct"`
}

func (s UpdateBootAndAntiUninstallPolicyResponseBodyStrategyBlockContent) String() string {
	return dara.Prettify(s)
}

func (s UpdateBootAndAntiUninstallPolicyResponseBodyStrategyBlockContent) GoString() string {
	return s.String()
}

func (s *UpdateBootAndAntiUninstallPolicyResponseBodyStrategyBlockContent) GetBlockTextEn() *UpdateBootAndAntiUninstallPolicyResponseBodyStrategyBlockContentBlockTextEn {
	return s.BlockTextEn
}

func (s *UpdateBootAndAntiUninstallPolicyResponseBodyStrategyBlockContent) GetBlockTextZh() *UpdateBootAndAntiUninstallPolicyResponseBodyStrategyBlockContentBlockTextZh {
	return s.BlockTextZh
}

func (s *UpdateBootAndAntiUninstallPolicyResponseBodyStrategyBlockContent) SetBlockTextEn(v *UpdateBootAndAntiUninstallPolicyResponseBodyStrategyBlockContentBlockTextEn) *UpdateBootAndAntiUninstallPolicyResponseBodyStrategyBlockContent {
	s.BlockTextEn = v
	return s
}

func (s *UpdateBootAndAntiUninstallPolicyResponseBodyStrategyBlockContent) SetBlockTextZh(v *UpdateBootAndAntiUninstallPolicyResponseBodyStrategyBlockContentBlockTextZh) *UpdateBootAndAntiUninstallPolicyResponseBodyStrategyBlockContent {
	s.BlockTextZh = v
	return s
}

func (s *UpdateBootAndAntiUninstallPolicyResponseBodyStrategyBlockContent) Validate() error {
	return dara.Validate(s)
}

type UpdateBootAndAntiUninstallPolicyResponseBodyStrategyBlockContentBlockTextEn struct {
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

func (s UpdateBootAndAntiUninstallPolicyResponseBodyStrategyBlockContentBlockTextEn) String() string {
	return dara.Prettify(s)
}

func (s UpdateBootAndAntiUninstallPolicyResponseBodyStrategyBlockContentBlockTextEn) GoString() string {
	return s.String()
}

func (s *UpdateBootAndAntiUninstallPolicyResponseBodyStrategyBlockContentBlockTextEn) GetContent() *string {
	return s.Content
}

func (s *UpdateBootAndAntiUninstallPolicyResponseBodyStrategyBlockContentBlockTextEn) GetMainButtonText() *string {
	return s.MainButtonText
}

func (s *UpdateBootAndAntiUninstallPolicyResponseBodyStrategyBlockContentBlockTextEn) GetMinorButtonText() *string {
	return s.MinorButtonText
}

func (s *UpdateBootAndAntiUninstallPolicyResponseBodyStrategyBlockContentBlockTextEn) GetTitle() *string {
	return s.Title
}

func (s *UpdateBootAndAntiUninstallPolicyResponseBodyStrategyBlockContentBlockTextEn) SetContent(v string) *UpdateBootAndAntiUninstallPolicyResponseBodyStrategyBlockContentBlockTextEn {
	s.Content = &v
	return s
}

func (s *UpdateBootAndAntiUninstallPolicyResponseBodyStrategyBlockContentBlockTextEn) SetMainButtonText(v string) *UpdateBootAndAntiUninstallPolicyResponseBodyStrategyBlockContentBlockTextEn {
	s.MainButtonText = &v
	return s
}

func (s *UpdateBootAndAntiUninstallPolicyResponseBodyStrategyBlockContentBlockTextEn) SetMinorButtonText(v string) *UpdateBootAndAntiUninstallPolicyResponseBodyStrategyBlockContentBlockTextEn {
	s.MinorButtonText = &v
	return s
}

func (s *UpdateBootAndAntiUninstallPolicyResponseBodyStrategyBlockContentBlockTextEn) SetTitle(v string) *UpdateBootAndAntiUninstallPolicyResponseBodyStrategyBlockContentBlockTextEn {
	s.Title = &v
	return s
}

func (s *UpdateBootAndAntiUninstallPolicyResponseBodyStrategyBlockContentBlockTextEn) Validate() error {
	return dara.Validate(s)
}

type UpdateBootAndAntiUninstallPolicyResponseBodyStrategyBlockContentBlockTextZh struct {
	Content         *string `json:"Content,omitempty" xml:"Content,omitempty"`
	MainButtonText  *string `json:"MainButtonText,omitempty" xml:"MainButtonText,omitempty"`
	MinorButtonText *string `json:"MinorButtonText,omitempty" xml:"MinorButtonText,omitempty"`
	Title           *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s UpdateBootAndAntiUninstallPolicyResponseBodyStrategyBlockContentBlockTextZh) String() string {
	return dara.Prettify(s)
}

func (s UpdateBootAndAntiUninstallPolicyResponseBodyStrategyBlockContentBlockTextZh) GoString() string {
	return s.String()
}

func (s *UpdateBootAndAntiUninstallPolicyResponseBodyStrategyBlockContentBlockTextZh) GetContent() *string {
	return s.Content
}

func (s *UpdateBootAndAntiUninstallPolicyResponseBodyStrategyBlockContentBlockTextZh) GetMainButtonText() *string {
	return s.MainButtonText
}

func (s *UpdateBootAndAntiUninstallPolicyResponseBodyStrategyBlockContentBlockTextZh) GetMinorButtonText() *string {
	return s.MinorButtonText
}

func (s *UpdateBootAndAntiUninstallPolicyResponseBodyStrategyBlockContentBlockTextZh) GetTitle() *string {
	return s.Title
}

func (s *UpdateBootAndAntiUninstallPolicyResponseBodyStrategyBlockContentBlockTextZh) SetContent(v string) *UpdateBootAndAntiUninstallPolicyResponseBodyStrategyBlockContentBlockTextZh {
	s.Content = &v
	return s
}

func (s *UpdateBootAndAntiUninstallPolicyResponseBodyStrategyBlockContentBlockTextZh) SetMainButtonText(v string) *UpdateBootAndAntiUninstallPolicyResponseBodyStrategyBlockContentBlockTextZh {
	s.MainButtonText = &v
	return s
}

func (s *UpdateBootAndAntiUninstallPolicyResponseBodyStrategyBlockContentBlockTextZh) SetMinorButtonText(v string) *UpdateBootAndAntiUninstallPolicyResponseBodyStrategyBlockContentBlockTextZh {
	s.MinorButtonText = &v
	return s
}

func (s *UpdateBootAndAntiUninstallPolicyResponseBodyStrategyBlockContentBlockTextZh) SetTitle(v string) *UpdateBootAndAntiUninstallPolicyResponseBodyStrategyBlockContentBlockTextZh {
	s.Title = &v
	return s
}

func (s *UpdateBootAndAntiUninstallPolicyResponseBodyStrategyBlockContentBlockTextZh) Validate() error {
	return dara.Validate(s)
}
