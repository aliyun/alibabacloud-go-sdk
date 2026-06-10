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
	DesktopGroupId *string `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty"`
	// This parameter is required.
	RegionId                  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UserProfilePathRuleShrink *string `json:"UserProfilePathRule,omitempty" xml:"UserProfilePathRule,omitempty"`
	UserProfileRuleType       *string `json:"UserProfileRuleType,omitempty" xml:"UserProfileRuleType,omitempty"`
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
