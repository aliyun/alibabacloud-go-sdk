// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGrafanaWorkspaceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountNumber(v string) *CreateGrafanaWorkspaceShrinkRequest
	GetAccountNumber() *string
	SetAliyunLang(v string) *CreateGrafanaWorkspaceShrinkRequest
	GetAliyunLang() *string
	SetAutoRenew(v string) *CreateGrafanaWorkspaceShrinkRequest
	GetAutoRenew() *string
	SetCustomAccountNumber(v string) *CreateGrafanaWorkspaceShrinkRequest
	GetCustomAccountNumber() *string
	SetDescription(v string) *CreateGrafanaWorkspaceShrinkRequest
	GetDescription() *string
	SetDuration(v string) *CreateGrafanaWorkspaceShrinkRequest
	GetDuration() *string
	SetGrafanaVersion(v string) *CreateGrafanaWorkspaceShrinkRequest
	GetGrafanaVersion() *string
	SetGrafanaWorkspaceEdition(v string) *CreateGrafanaWorkspaceShrinkRequest
	GetGrafanaWorkspaceEdition() *string
	SetGrafanaWorkspaceName(v string) *CreateGrafanaWorkspaceShrinkRequest
	GetGrafanaWorkspaceName() *string
	SetPassword(v string) *CreateGrafanaWorkspaceShrinkRequest
	GetPassword() *string
	SetPricingCycle(v string) *CreateGrafanaWorkspaceShrinkRequest
	GetPricingCycle() *string
	SetRegionId(v string) *CreateGrafanaWorkspaceShrinkRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateGrafanaWorkspaceShrinkRequest
	GetResourceGroupId() *string
	SetTagsShrink(v string) *CreateGrafanaWorkspaceShrinkRequest
	GetTagsShrink() *string
}

type CreateGrafanaWorkspaceShrinkRequest struct {
	// Account quantity.
	//
	// **Value description:*	-
	//
	// - If GrafanaWorkspaceEdition is **standard**, this parameter is invalid.
	//
	// - If GrafanaWorkspaceEdition is **personal_edition**, this parameter is invalid. Default Value: 1.
	//
	// - If GrafanaWorkspaceEdition is **experts_edition**, valid values are 10, 30, or 50. Default Value: 10.
	//
	// - If GrafanaWorkspaceEdition is **advanced_edition**, this parameter is invalid. Default Value: 100.
	//
	// example:
	//
	// 10
	AccountNumber *string `json:"AccountNumber,omitempty" xml:"AccountNumber,omitempty"`
	// The language. Default value: zh. Valid values:
	//
	// 	- zh
	//
	// 	- en
	//
	// example:
	//
	// zh
	AliyunLang *string `json:"AliyunLang,omitempty" xml:"AliyunLang,omitempty"`
	// Whether auto-renewal is enabled. Valid values:
	//
	// - true: Auto-renewal is enabled.
	//
	// - false: Auto-renewal is disabled.
	//
	// Default Value: true.
	//
	// example:
	//
	// false
	AutoRenew *string `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// Additional custom account quantity for the User.
	//
	// **Value description:*	-
	//
	// - If GrafanaWorkspaceEdition is **standard**, this parameter is invalid.
	//
	// - If GrafanaWorkspaceEdition is **personal_edition**, this parameter is invalid.
	//
	// - If GrafanaWorkspaceEdition is **experts_edition**, this parameter is invalid.
	//
	// - If GrafanaWorkspaceEdition is **advanced_edition**, the value range is 0 to 2000 and must be a multiple of 10. Default Value: 0.
	//
	// example:
	//
	// 0
	CustomAccountNumber *string `json:"CustomAccountNumber,omitempty" xml:"CustomAccountNumber,omitempty"`
	// The description of the workspace
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Subscription duration of the instance. Valid values:
	//
	// - If PricingCycle is **Month**, indicating monthly billing, the value range is **1*	- to **9**.
	//
	// - If PricingCycle is **Year**, indicating yearly billing, the value range is **1*	- to **3**.
	//
	// Default Value: 1.
	//
	// example:
	//
	// 6
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// Grafana version.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10.0.x
	GrafanaVersion *string `json:"GrafanaVersion,omitempty" xml:"GrafanaVersion,omitempty"`
	// The edition.
	//
	// **Valid values:**
	//
	// 	- standard: `Beta Edition or Standard Edition`
	//
	// 	- personal_edition: Developer Edition
	//
	// 	- experts_edition: Pro Edition
	//
	// 	- advanced_edition: Advanced Edition
	//
	// This parameter is required.
	//
	// example:
	//
	// experts_edition
	GrafanaWorkspaceEdition *string `json:"GrafanaWorkspaceEdition,omitempty" xml:"GrafanaWorkspaceEdition,omitempty"`
	// The name of the Grafana workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// testgrafana
	GrafanaWorkspaceName *string `json:"GrafanaWorkspaceName,omitempty" xml:"GrafanaWorkspaceName,omitempty"`
	// The password of the workspace. The password must be 8 to 30 characters in length. It must include at least three of the following characters types: uppercase letter, lowercase letter, digit, and special character. Special characters include () \\" ~ ! @ # $ % ^ & \\	- - _ + =.
	//
	// example:
	//
	// Test123456!
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// 包年包月的计费周期，取值： Month（默认值）：按月购买。                                 Year：按年购买。
	//
	// example:
	//
	// Month
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmxyexli2****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The list of tags.
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s CreateGrafanaWorkspaceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateGrafanaWorkspaceShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateGrafanaWorkspaceShrinkRequest) GetAccountNumber() *string {
	return s.AccountNumber
}

func (s *CreateGrafanaWorkspaceShrinkRequest) GetAliyunLang() *string {
	return s.AliyunLang
}

func (s *CreateGrafanaWorkspaceShrinkRequest) GetAutoRenew() *string {
	return s.AutoRenew
}

func (s *CreateGrafanaWorkspaceShrinkRequest) GetCustomAccountNumber() *string {
	return s.CustomAccountNumber
}

func (s *CreateGrafanaWorkspaceShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateGrafanaWorkspaceShrinkRequest) GetDuration() *string {
	return s.Duration
}

func (s *CreateGrafanaWorkspaceShrinkRequest) GetGrafanaVersion() *string {
	return s.GrafanaVersion
}

func (s *CreateGrafanaWorkspaceShrinkRequest) GetGrafanaWorkspaceEdition() *string {
	return s.GrafanaWorkspaceEdition
}

func (s *CreateGrafanaWorkspaceShrinkRequest) GetGrafanaWorkspaceName() *string {
	return s.GrafanaWorkspaceName
}

func (s *CreateGrafanaWorkspaceShrinkRequest) GetPassword() *string {
	return s.Password
}

func (s *CreateGrafanaWorkspaceShrinkRequest) GetPricingCycle() *string {
	return s.PricingCycle
}

func (s *CreateGrafanaWorkspaceShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateGrafanaWorkspaceShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateGrafanaWorkspaceShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *CreateGrafanaWorkspaceShrinkRequest) SetAccountNumber(v string) *CreateGrafanaWorkspaceShrinkRequest {
	s.AccountNumber = &v
	return s
}

func (s *CreateGrafanaWorkspaceShrinkRequest) SetAliyunLang(v string) *CreateGrafanaWorkspaceShrinkRequest {
	s.AliyunLang = &v
	return s
}

func (s *CreateGrafanaWorkspaceShrinkRequest) SetAutoRenew(v string) *CreateGrafanaWorkspaceShrinkRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateGrafanaWorkspaceShrinkRequest) SetCustomAccountNumber(v string) *CreateGrafanaWorkspaceShrinkRequest {
	s.CustomAccountNumber = &v
	return s
}

func (s *CreateGrafanaWorkspaceShrinkRequest) SetDescription(v string) *CreateGrafanaWorkspaceShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateGrafanaWorkspaceShrinkRequest) SetDuration(v string) *CreateGrafanaWorkspaceShrinkRequest {
	s.Duration = &v
	return s
}

func (s *CreateGrafanaWorkspaceShrinkRequest) SetGrafanaVersion(v string) *CreateGrafanaWorkspaceShrinkRequest {
	s.GrafanaVersion = &v
	return s
}

func (s *CreateGrafanaWorkspaceShrinkRequest) SetGrafanaWorkspaceEdition(v string) *CreateGrafanaWorkspaceShrinkRequest {
	s.GrafanaWorkspaceEdition = &v
	return s
}

func (s *CreateGrafanaWorkspaceShrinkRequest) SetGrafanaWorkspaceName(v string) *CreateGrafanaWorkspaceShrinkRequest {
	s.GrafanaWorkspaceName = &v
	return s
}

func (s *CreateGrafanaWorkspaceShrinkRequest) SetPassword(v string) *CreateGrafanaWorkspaceShrinkRequest {
	s.Password = &v
	return s
}

func (s *CreateGrafanaWorkspaceShrinkRequest) SetPricingCycle(v string) *CreateGrafanaWorkspaceShrinkRequest {
	s.PricingCycle = &v
	return s
}

func (s *CreateGrafanaWorkspaceShrinkRequest) SetRegionId(v string) *CreateGrafanaWorkspaceShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateGrafanaWorkspaceShrinkRequest) SetResourceGroupId(v string) *CreateGrafanaWorkspaceShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateGrafanaWorkspaceShrinkRequest) SetTagsShrink(v string) *CreateGrafanaWorkspaceShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *CreateGrafanaWorkspaceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
