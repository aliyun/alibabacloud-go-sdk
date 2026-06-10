// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserProfilePathRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeUserProfilePathRulesResponseBody
	GetRequestId() *string
	SetUserProfilePathRule(v *DescribeUserProfilePathRulesResponseBodyUserProfilePathRule) *DescribeUserProfilePathRulesResponseBody
	GetUserProfilePathRule() *DescribeUserProfilePathRulesResponseBodyUserProfilePathRule
}

type DescribeUserProfilePathRulesResponseBody struct {
	RequestId           *string                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UserProfilePathRule *DescribeUserProfilePathRulesResponseBodyUserProfilePathRule `json:"UserProfilePathRule,omitempty" xml:"UserProfilePathRule,omitempty" type:"Struct"`
}

func (s DescribeUserProfilePathRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserProfilePathRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserProfilePathRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUserProfilePathRulesResponseBody) GetUserProfilePathRule() *DescribeUserProfilePathRulesResponseBodyUserProfilePathRule {
	return s.UserProfilePathRule
}

func (s *DescribeUserProfilePathRulesResponseBody) SetRequestId(v string) *DescribeUserProfilePathRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserProfilePathRulesResponseBody) SetUserProfilePathRule(v *DescribeUserProfilePathRulesResponseBodyUserProfilePathRule) *DescribeUserProfilePathRulesResponseBody {
	s.UserProfilePathRule = v
	return s
}

func (s *DescribeUserProfilePathRulesResponseBody) Validate() error {
	if s.UserProfilePathRule != nil {
		if err := s.UserProfilePathRule.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeUserProfilePathRulesResponseBodyUserProfilePathRule struct {
	DesktopGroupId      *string                                                             `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty"`
	Rules               []*DescribeUserProfilePathRulesResponseBodyUserProfilePathRuleRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	UserProfileRuleType *string                                                             `json:"UserProfileRuleType,omitempty" xml:"UserProfileRuleType,omitempty"`
}

func (s DescribeUserProfilePathRulesResponseBodyUserProfilePathRule) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserProfilePathRulesResponseBodyUserProfilePathRule) GoString() string {
	return s.String()
}

func (s *DescribeUserProfilePathRulesResponseBodyUserProfilePathRule) GetDesktopGroupId() *string {
	return s.DesktopGroupId
}

func (s *DescribeUserProfilePathRulesResponseBodyUserProfilePathRule) GetRules() []*DescribeUserProfilePathRulesResponseBodyUserProfilePathRuleRules {
	return s.Rules
}

func (s *DescribeUserProfilePathRulesResponseBodyUserProfilePathRule) GetUserProfileRuleType() *string {
	return s.UserProfileRuleType
}

func (s *DescribeUserProfilePathRulesResponseBodyUserProfilePathRule) SetDesktopGroupId(v string) *DescribeUserProfilePathRulesResponseBodyUserProfilePathRule {
	s.DesktopGroupId = &v
	return s
}

func (s *DescribeUserProfilePathRulesResponseBodyUserProfilePathRule) SetRules(v []*DescribeUserProfilePathRulesResponseBodyUserProfilePathRuleRules) *DescribeUserProfilePathRulesResponseBodyUserProfilePathRule {
	s.Rules = v
	return s
}

func (s *DescribeUserProfilePathRulesResponseBodyUserProfilePathRule) SetUserProfileRuleType(v string) *DescribeUserProfilePathRulesResponseBodyUserProfilePathRule {
	s.UserProfileRuleType = &v
	return s
}

func (s *DescribeUserProfilePathRulesResponseBodyUserProfilePathRule) Validate() error {
	if s.Rules != nil {
		for _, item := range s.Rules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeUserProfilePathRulesResponseBodyUserProfilePathRuleRules struct {
	BlackPath  *DescribeUserProfilePathRulesResponseBodyUserProfilePathRuleRulesBlackPath    `json:"BlackPath,omitempty" xml:"BlackPath,omitempty" type:"Struct"`
	WhitePaths []*DescribeUserProfilePathRulesResponseBodyUserProfilePathRuleRulesWhitePaths `json:"WhitePaths,omitempty" xml:"WhitePaths,omitempty" type:"Repeated"`
}

func (s DescribeUserProfilePathRulesResponseBodyUserProfilePathRuleRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserProfilePathRulesResponseBodyUserProfilePathRuleRules) GoString() string {
	return s.String()
}

func (s *DescribeUserProfilePathRulesResponseBodyUserProfilePathRuleRules) GetBlackPath() *DescribeUserProfilePathRulesResponseBodyUserProfilePathRuleRulesBlackPath {
	return s.BlackPath
}

func (s *DescribeUserProfilePathRulesResponseBodyUserProfilePathRuleRules) GetWhitePaths() []*DescribeUserProfilePathRulesResponseBodyUserProfilePathRuleRulesWhitePaths {
	return s.WhitePaths
}

func (s *DescribeUserProfilePathRulesResponseBodyUserProfilePathRuleRules) SetBlackPath(v *DescribeUserProfilePathRulesResponseBodyUserProfilePathRuleRulesBlackPath) *DescribeUserProfilePathRulesResponseBodyUserProfilePathRuleRules {
	s.BlackPath = v
	return s
}

func (s *DescribeUserProfilePathRulesResponseBodyUserProfilePathRuleRules) SetWhitePaths(v []*DescribeUserProfilePathRulesResponseBodyUserProfilePathRuleRulesWhitePaths) *DescribeUserProfilePathRulesResponseBodyUserProfilePathRuleRules {
	s.WhitePaths = v
	return s
}

func (s *DescribeUserProfilePathRulesResponseBodyUserProfilePathRuleRules) Validate() error {
	if s.BlackPath != nil {
		if err := s.BlackPath.Validate(); err != nil {
			return err
		}
	}
	if s.WhitePaths != nil {
		for _, item := range s.WhitePaths {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeUserProfilePathRulesResponseBodyUserProfilePathRuleRulesBlackPath struct {
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeUserProfilePathRulesResponseBodyUserProfilePathRuleRulesBlackPath) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserProfilePathRulesResponseBodyUserProfilePathRuleRulesBlackPath) GoString() string {
	return s.String()
}

func (s *DescribeUserProfilePathRulesResponseBodyUserProfilePathRuleRulesBlackPath) GetPath() *string {
	return s.Path
}

func (s *DescribeUserProfilePathRulesResponseBodyUserProfilePathRuleRulesBlackPath) GetType() *string {
	return s.Type
}

func (s *DescribeUserProfilePathRulesResponseBodyUserProfilePathRuleRulesBlackPath) SetPath(v string) *DescribeUserProfilePathRulesResponseBodyUserProfilePathRuleRulesBlackPath {
	s.Path = &v
	return s
}

func (s *DescribeUserProfilePathRulesResponseBodyUserProfilePathRuleRulesBlackPath) SetType(v string) *DescribeUserProfilePathRulesResponseBodyUserProfilePathRuleRulesBlackPath {
	s.Type = &v
	return s
}

func (s *DescribeUserProfilePathRulesResponseBodyUserProfilePathRuleRulesBlackPath) Validate() error {
	return dara.Validate(s)
}

type DescribeUserProfilePathRulesResponseBodyUserProfilePathRuleRulesWhitePaths struct {
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeUserProfilePathRulesResponseBodyUserProfilePathRuleRulesWhitePaths) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserProfilePathRulesResponseBodyUserProfilePathRuleRulesWhitePaths) GoString() string {
	return s.String()
}

func (s *DescribeUserProfilePathRulesResponseBodyUserProfilePathRuleRulesWhitePaths) GetPath() *string {
	return s.Path
}

func (s *DescribeUserProfilePathRulesResponseBodyUserProfilePathRuleRulesWhitePaths) GetType() *string {
	return s.Type
}

func (s *DescribeUserProfilePathRulesResponseBodyUserProfilePathRuleRulesWhitePaths) SetPath(v string) *DescribeUserProfilePathRulesResponseBodyUserProfilePathRuleRulesWhitePaths {
	s.Path = &v
	return s
}

func (s *DescribeUserProfilePathRulesResponseBodyUserProfilePathRuleRulesWhitePaths) SetType(v string) *DescribeUserProfilePathRulesResponseBodyUserProfilePathRuleRulesWhitePaths {
	s.Type = &v
	return s
}

func (s *DescribeUserProfilePathRulesResponseBodyUserProfilePathRuleRulesWhitePaths) Validate() error {
	return dara.Validate(s)
}
