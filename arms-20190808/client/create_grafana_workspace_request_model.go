// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGrafanaWorkspaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountNumber(v string) *CreateGrafanaWorkspaceRequest
	GetAccountNumber() *string
	SetAliyunLang(v string) *CreateGrafanaWorkspaceRequest
	GetAliyunLang() *string
	SetAutoRenew(v string) *CreateGrafanaWorkspaceRequest
	GetAutoRenew() *string
	SetCustomAccountNumber(v string) *CreateGrafanaWorkspaceRequest
	GetCustomAccountNumber() *string
	SetDescription(v string) *CreateGrafanaWorkspaceRequest
	GetDescription() *string
	SetDuration(v string) *CreateGrafanaWorkspaceRequest
	GetDuration() *string
	SetGrafanaVersion(v string) *CreateGrafanaWorkspaceRequest
	GetGrafanaVersion() *string
	SetGrafanaWorkspaceEdition(v string) *CreateGrafanaWorkspaceRequest
	GetGrafanaWorkspaceEdition() *string
	SetGrafanaWorkspaceName(v string) *CreateGrafanaWorkspaceRequest
	GetGrafanaWorkspaceName() *string
	SetPassword(v string) *CreateGrafanaWorkspaceRequest
	GetPassword() *string
	SetPricingCycle(v string) *CreateGrafanaWorkspaceRequest
	GetPricingCycle() *string
	SetRegionId(v string) *CreateGrafanaWorkspaceRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateGrafanaWorkspaceRequest
	GetResourceGroupId() *string
	SetTags(v []*CreateGrafanaWorkspaceRequestTags) *CreateGrafanaWorkspaceRequest
	GetTags() []*CreateGrafanaWorkspaceRequestTags
}

type CreateGrafanaWorkspaceRequest struct {
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
	// example:
	//
	// false
	AutoRenew *string `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
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
	// example:
	//
	// 6
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
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
	Tags []*CreateGrafanaWorkspaceRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s CreateGrafanaWorkspaceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateGrafanaWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *CreateGrafanaWorkspaceRequest) GetAccountNumber() *string {
	return s.AccountNumber
}

func (s *CreateGrafanaWorkspaceRequest) GetAliyunLang() *string {
	return s.AliyunLang
}

func (s *CreateGrafanaWorkspaceRequest) GetAutoRenew() *string {
	return s.AutoRenew
}

func (s *CreateGrafanaWorkspaceRequest) GetCustomAccountNumber() *string {
	return s.CustomAccountNumber
}

func (s *CreateGrafanaWorkspaceRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateGrafanaWorkspaceRequest) GetDuration() *string {
	return s.Duration
}

func (s *CreateGrafanaWorkspaceRequest) GetGrafanaVersion() *string {
	return s.GrafanaVersion
}

func (s *CreateGrafanaWorkspaceRequest) GetGrafanaWorkspaceEdition() *string {
	return s.GrafanaWorkspaceEdition
}

func (s *CreateGrafanaWorkspaceRequest) GetGrafanaWorkspaceName() *string {
	return s.GrafanaWorkspaceName
}

func (s *CreateGrafanaWorkspaceRequest) GetPassword() *string {
	return s.Password
}

func (s *CreateGrafanaWorkspaceRequest) GetPricingCycle() *string {
	return s.PricingCycle
}

func (s *CreateGrafanaWorkspaceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateGrafanaWorkspaceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateGrafanaWorkspaceRequest) GetTags() []*CreateGrafanaWorkspaceRequestTags {
	return s.Tags
}

func (s *CreateGrafanaWorkspaceRequest) SetAccountNumber(v string) *CreateGrafanaWorkspaceRequest {
	s.AccountNumber = &v
	return s
}

func (s *CreateGrafanaWorkspaceRequest) SetAliyunLang(v string) *CreateGrafanaWorkspaceRequest {
	s.AliyunLang = &v
	return s
}

func (s *CreateGrafanaWorkspaceRequest) SetAutoRenew(v string) *CreateGrafanaWorkspaceRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateGrafanaWorkspaceRequest) SetCustomAccountNumber(v string) *CreateGrafanaWorkspaceRequest {
	s.CustomAccountNumber = &v
	return s
}

func (s *CreateGrafanaWorkspaceRequest) SetDescription(v string) *CreateGrafanaWorkspaceRequest {
	s.Description = &v
	return s
}

func (s *CreateGrafanaWorkspaceRequest) SetDuration(v string) *CreateGrafanaWorkspaceRequest {
	s.Duration = &v
	return s
}

func (s *CreateGrafanaWorkspaceRequest) SetGrafanaVersion(v string) *CreateGrafanaWorkspaceRequest {
	s.GrafanaVersion = &v
	return s
}

func (s *CreateGrafanaWorkspaceRequest) SetGrafanaWorkspaceEdition(v string) *CreateGrafanaWorkspaceRequest {
	s.GrafanaWorkspaceEdition = &v
	return s
}

func (s *CreateGrafanaWorkspaceRequest) SetGrafanaWorkspaceName(v string) *CreateGrafanaWorkspaceRequest {
	s.GrafanaWorkspaceName = &v
	return s
}

func (s *CreateGrafanaWorkspaceRequest) SetPassword(v string) *CreateGrafanaWorkspaceRequest {
	s.Password = &v
	return s
}

func (s *CreateGrafanaWorkspaceRequest) SetPricingCycle(v string) *CreateGrafanaWorkspaceRequest {
	s.PricingCycle = &v
	return s
}

func (s *CreateGrafanaWorkspaceRequest) SetRegionId(v string) *CreateGrafanaWorkspaceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateGrafanaWorkspaceRequest) SetResourceGroupId(v string) *CreateGrafanaWorkspaceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateGrafanaWorkspaceRequest) SetTags(v []*CreateGrafanaWorkspaceRequestTags) *CreateGrafanaWorkspaceRequest {
	s.Tags = v
	return s
}

func (s *CreateGrafanaWorkspaceRequest) Validate() error {
	return dara.Validate(s)
}

type CreateGrafanaWorkspaceRequestTags struct {
	// The tag key.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateGrafanaWorkspaceRequestTags) String() string {
	return dara.Prettify(s)
}

func (s CreateGrafanaWorkspaceRequestTags) GoString() string {
	return s.String()
}

func (s *CreateGrafanaWorkspaceRequestTags) GetKey() *string {
	return s.Key
}

func (s *CreateGrafanaWorkspaceRequestTags) GetValue() *string {
	return s.Value
}

func (s *CreateGrafanaWorkspaceRequestTags) SetKey(v string) *CreateGrafanaWorkspaceRequestTags {
	s.Key = &v
	return s
}

func (s *CreateGrafanaWorkspaceRequestTags) SetValue(v string) *CreateGrafanaWorkspaceRequestTags {
	s.Value = &v
	return s
}

func (s *CreateGrafanaWorkspaceRequestTags) Validate() error {
	return dara.Validate(s)
}
