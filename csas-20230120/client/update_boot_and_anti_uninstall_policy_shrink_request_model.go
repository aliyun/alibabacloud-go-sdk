// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBootAndAntiUninstallPolicyShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllowReport(v bool) *UpdateBootAndAntiUninstallPolicyShrinkRequest
	GetAllowReport() *bool
	SetBlockContentShrink(v string) *UpdateBootAndAntiUninstallPolicyShrinkRequest
	GetBlockContentShrink() *string
	SetIsAntiUninstall(v bool) *UpdateBootAndAntiUninstallPolicyShrinkRequest
	GetIsAntiUninstall() *bool
	SetIsBoot(v bool) *UpdateBootAndAntiUninstallPolicyShrinkRequest
	GetIsBoot() *bool
	SetUserGroupIds(v []*string) *UpdateBootAndAntiUninstallPolicyShrinkRequest
	GetUserGroupIds() []*string
	SetWhitelistUsers(v []*string) *UpdateBootAndAntiUninstallPolicyShrinkRequest
	GetWhitelistUsers() []*string
}

type UpdateBootAndAntiUninstallPolicyShrinkRequest struct {
	// example:
	//
	// true
	AllowReport        *bool   `json:"AllowReport,omitempty" xml:"AllowReport,omitempty"`
	BlockContentShrink *string `json:"BlockContent,omitempty" xml:"BlockContent,omitempty"`
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

func (s UpdateBootAndAntiUninstallPolicyShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateBootAndAntiUninstallPolicyShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateBootAndAntiUninstallPolicyShrinkRequest) GetAllowReport() *bool {
	return s.AllowReport
}

func (s *UpdateBootAndAntiUninstallPolicyShrinkRequest) GetBlockContentShrink() *string {
	return s.BlockContentShrink
}

func (s *UpdateBootAndAntiUninstallPolicyShrinkRequest) GetIsAntiUninstall() *bool {
	return s.IsAntiUninstall
}

func (s *UpdateBootAndAntiUninstallPolicyShrinkRequest) GetIsBoot() *bool {
	return s.IsBoot
}

func (s *UpdateBootAndAntiUninstallPolicyShrinkRequest) GetUserGroupIds() []*string {
	return s.UserGroupIds
}

func (s *UpdateBootAndAntiUninstallPolicyShrinkRequest) GetWhitelistUsers() []*string {
	return s.WhitelistUsers
}

func (s *UpdateBootAndAntiUninstallPolicyShrinkRequest) SetAllowReport(v bool) *UpdateBootAndAntiUninstallPolicyShrinkRequest {
	s.AllowReport = &v
	return s
}

func (s *UpdateBootAndAntiUninstallPolicyShrinkRequest) SetBlockContentShrink(v string) *UpdateBootAndAntiUninstallPolicyShrinkRequest {
	s.BlockContentShrink = &v
	return s
}

func (s *UpdateBootAndAntiUninstallPolicyShrinkRequest) SetIsAntiUninstall(v bool) *UpdateBootAndAntiUninstallPolicyShrinkRequest {
	s.IsAntiUninstall = &v
	return s
}

func (s *UpdateBootAndAntiUninstallPolicyShrinkRequest) SetIsBoot(v bool) *UpdateBootAndAntiUninstallPolicyShrinkRequest {
	s.IsBoot = &v
	return s
}

func (s *UpdateBootAndAntiUninstallPolicyShrinkRequest) SetUserGroupIds(v []*string) *UpdateBootAndAntiUninstallPolicyShrinkRequest {
	s.UserGroupIds = v
	return s
}

func (s *UpdateBootAndAntiUninstallPolicyShrinkRequest) SetWhitelistUsers(v []*string) *UpdateBootAndAntiUninstallPolicyShrinkRequest {
	s.WhitelistUsers = v
	return s
}

func (s *UpdateBootAndAntiUninstallPolicyShrinkRequest) Validate() error {
	return dara.Validate(s)
}
