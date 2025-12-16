// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAppGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentVersion(v string) *ModifyAppGroupRequest
	GetCurrentVersion() *string
	SetDescription(v string) *ModifyAppGroupRequest
	GetDescription() *string
	SetDomain(v string) *ModifyAppGroupRequest
	GetDomain() *string
	SetResourceGroupId(v string) *ModifyAppGroupRequest
	GetResourceGroupId() *string
	SetDryRun(v bool) *ModifyAppGroupRequest
	GetDryRun() *bool
}

type ModifyAppGroupRequest struct {
	// The online version of the application.
	//
	// example:
	//
	// 1223232
	CurrentVersion *string `json:"currentVersion,omitempty" xml:"currentVersion,omitempty"`
	// The description of the application.
	//
	// example:
	//
	// "test"
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The type of the industry. Valid values:
	//
	// 	- general: general.
	//
	// 	- ecommerce: e-commerce.
	//
	// 	- education: education.
	//
	// 	- esports: electronic sports.
	//
	// 	- community: content community.
	//
	// example:
	//
	// "ecommerce"
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// The ID of the resource group to which the instance belongs.
	//
	// example:
	//
	// rg-****
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// Specifies whether to verify the application before modification. Valid values: true and false.
	//
	// example:
	//
	// true
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s ModifyAppGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyAppGroupRequest) GetCurrentVersion() *string {
	return s.CurrentVersion
}

func (s *ModifyAppGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyAppGroupRequest) GetDomain() *string {
	return s.Domain
}

func (s *ModifyAppGroupRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyAppGroupRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ModifyAppGroupRequest) SetCurrentVersion(v string) *ModifyAppGroupRequest {
	s.CurrentVersion = &v
	return s
}

func (s *ModifyAppGroupRequest) SetDescription(v string) *ModifyAppGroupRequest {
	s.Description = &v
	return s
}

func (s *ModifyAppGroupRequest) SetDomain(v string) *ModifyAppGroupRequest {
	s.Domain = &v
	return s
}

func (s *ModifyAppGroupRequest) SetResourceGroupId(v string) *ModifyAppGroupRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyAppGroupRequest) SetDryRun(v bool) *ModifyAppGroupRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyAppGroupRequest) Validate() error {
	return dara.Validate(s)
}
