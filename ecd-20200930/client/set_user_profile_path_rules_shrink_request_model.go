// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetUserProfilePathRulesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDesktopGroupId(v string) *SetUserProfilePathRulesShrinkRequest
	GetDesktopGroupId() *string
	SetRegionId(v string) *SetUserProfilePathRulesShrinkRequest
	GetRegionId() *string
	SetUserProfilePathRuleShrink(v string) *SetUserProfilePathRulesShrinkRequest
	GetUserProfilePathRuleShrink() *string
	SetUserProfileRuleType(v string) *SetUserProfilePathRulesShrinkRequest
	GetUserProfileRuleType() *string
}

type SetUserProfilePathRulesShrinkRequest struct {
	// The desktop group ID.
	//
	// example:
	//
	// dg-2i8qxpv6t1a03****
	DesktopGroupId *string `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The directories that you want to configure in the blacklist and whitelist.
	UserProfilePathRuleShrink *string `json:"UserProfilePathRule,omitempty" xml:"UserProfilePathRule,omitempty"`
	// The directory type that you want to configure.
	//
	// Valid values:
	//
	// 	- Both_Default_DesktopGroup
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- DesktopGroup
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Default
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// DesktopGroup
	UserProfileRuleType *string `json:"UserProfileRuleType,omitempty" xml:"UserProfileRuleType,omitempty"`
}

func (s SetUserProfilePathRulesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SetUserProfilePathRulesShrinkRequest) GoString() string {
	return s.String()
}

func (s *SetUserProfilePathRulesShrinkRequest) GetDesktopGroupId() *string {
	return s.DesktopGroupId
}

func (s *SetUserProfilePathRulesShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SetUserProfilePathRulesShrinkRequest) GetUserProfilePathRuleShrink() *string {
	return s.UserProfilePathRuleShrink
}

func (s *SetUserProfilePathRulesShrinkRequest) GetUserProfileRuleType() *string {
	return s.UserProfileRuleType
}

func (s *SetUserProfilePathRulesShrinkRequest) SetDesktopGroupId(v string) *SetUserProfilePathRulesShrinkRequest {
	s.DesktopGroupId = &v
	return s
}

func (s *SetUserProfilePathRulesShrinkRequest) SetRegionId(v string) *SetUserProfilePathRulesShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *SetUserProfilePathRulesShrinkRequest) SetUserProfilePathRuleShrink(v string) *SetUserProfilePathRulesShrinkRequest {
	s.UserProfilePathRuleShrink = &v
	return s
}

func (s *SetUserProfilePathRulesShrinkRequest) SetUserProfileRuleType(v string) *SetUserProfilePathRulesShrinkRequest {
	s.UserProfileRuleType = &v
	return s
}

func (s *SetUserProfilePathRulesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
