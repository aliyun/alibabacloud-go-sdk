// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetUserProfilePathRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDesktopGroupId(v string) *SetUserProfilePathRulesRequest
	GetDesktopGroupId() *string
	SetRegionId(v string) *SetUserProfilePathRulesRequest
	GetRegionId() *string
	SetUserProfilePathRule(v []*SetUserProfilePathRulesRequestUserProfilePathRule) *SetUserProfilePathRulesRequest
	GetUserProfilePathRule() []*SetUserProfilePathRulesRequestUserProfilePathRule
	SetUserProfileRuleType(v string) *SetUserProfilePathRulesRequest
	GetUserProfileRuleType() *string
}

type SetUserProfilePathRulesRequest struct {
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
	UserProfilePathRule []*SetUserProfilePathRulesRequestUserProfilePathRule `json:"UserProfilePathRule,omitempty" xml:"UserProfilePathRule,omitempty" type:"Repeated"`
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

func (s SetUserProfilePathRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s SetUserProfilePathRulesRequest) GoString() string {
	return s.String()
}

func (s *SetUserProfilePathRulesRequest) GetDesktopGroupId() *string {
	return s.DesktopGroupId
}

func (s *SetUserProfilePathRulesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SetUserProfilePathRulesRequest) GetUserProfilePathRule() []*SetUserProfilePathRulesRequestUserProfilePathRule {
	return s.UserProfilePathRule
}

func (s *SetUserProfilePathRulesRequest) GetUserProfileRuleType() *string {
	return s.UserProfileRuleType
}

func (s *SetUserProfilePathRulesRequest) SetDesktopGroupId(v string) *SetUserProfilePathRulesRequest {
	s.DesktopGroupId = &v
	return s
}

func (s *SetUserProfilePathRulesRequest) SetRegionId(v string) *SetUserProfilePathRulesRequest {
	s.RegionId = &v
	return s
}

func (s *SetUserProfilePathRulesRequest) SetUserProfilePathRule(v []*SetUserProfilePathRulesRequestUserProfilePathRule) *SetUserProfilePathRulesRequest {
	s.UserProfilePathRule = v
	return s
}

func (s *SetUserProfilePathRulesRequest) SetUserProfileRuleType(v string) *SetUserProfilePathRulesRequest {
	s.UserProfileRuleType = &v
	return s
}

func (s *SetUserProfilePathRulesRequest) Validate() error {
	return dara.Validate(s)
}

type SetUserProfilePathRulesRequestUserProfilePathRule struct {
	// The directory in the blacklist.
	BlackPath *SetUserProfilePathRulesRequestUserProfilePathRuleBlackPath `json:"BlackPath,omitempty" xml:"BlackPath,omitempty" type:"Struct"`
	// The directories that you want to configure in the whitelist.
	WhitePaths []*SetUserProfilePathRulesRequestUserProfilePathRuleWhitePaths `json:"WhitePaths,omitempty" xml:"WhitePaths,omitempty" type:"Repeated"`
}

func (s SetUserProfilePathRulesRequestUserProfilePathRule) String() string {
	return dara.Prettify(s)
}

func (s SetUserProfilePathRulesRequestUserProfilePathRule) GoString() string {
	return s.String()
}

func (s *SetUserProfilePathRulesRequestUserProfilePathRule) GetBlackPath() *SetUserProfilePathRulesRequestUserProfilePathRuleBlackPath {
	return s.BlackPath
}

func (s *SetUserProfilePathRulesRequestUserProfilePathRule) GetWhitePaths() []*SetUserProfilePathRulesRequestUserProfilePathRuleWhitePaths {
	return s.WhitePaths
}

func (s *SetUserProfilePathRulesRequestUserProfilePathRule) SetBlackPath(v *SetUserProfilePathRulesRequestUserProfilePathRuleBlackPath) *SetUserProfilePathRulesRequestUserProfilePathRule {
	s.BlackPath = v
	return s
}

func (s *SetUserProfilePathRulesRequestUserProfilePathRule) SetWhitePaths(v []*SetUserProfilePathRulesRequestUserProfilePathRuleWhitePaths) *SetUserProfilePathRulesRequestUserProfilePathRule {
	s.WhitePaths = v
	return s
}

func (s *SetUserProfilePathRulesRequestUserProfilePathRule) Validate() error {
	return dara.Validate(s)
}

type SetUserProfilePathRulesRequestUserProfilePathRuleBlackPath struct {
	// The blacklist path.
	//
	// example:
	//
	// AppLocal/Data
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The path type.
	//
	// Valid values:
	//
	// 	- file
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- folder
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// folder
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SetUserProfilePathRulesRequestUserProfilePathRuleBlackPath) String() string {
	return dara.Prettify(s)
}

func (s SetUserProfilePathRulesRequestUserProfilePathRuleBlackPath) GoString() string {
	return s.String()
}

func (s *SetUserProfilePathRulesRequestUserProfilePathRuleBlackPath) GetPath() *string {
	return s.Path
}

func (s *SetUserProfilePathRulesRequestUserProfilePathRuleBlackPath) GetType() *string {
	return s.Type
}

func (s *SetUserProfilePathRulesRequestUserProfilePathRuleBlackPath) SetPath(v string) *SetUserProfilePathRulesRequestUserProfilePathRuleBlackPath {
	s.Path = &v
	return s
}

func (s *SetUserProfilePathRulesRequestUserProfilePathRuleBlackPath) SetType(v string) *SetUserProfilePathRulesRequestUserProfilePathRuleBlackPath {
	s.Type = &v
	return s
}

func (s *SetUserProfilePathRulesRequestUserProfilePathRuleBlackPath) Validate() error {
	return dara.Validate(s)
}

type SetUserProfilePathRulesRequestUserProfilePathRuleWhitePaths struct {
	// The whitelist path.
	//
	// example:
	//
	// whitePath
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The path type.
	//
	// Valid values:
	//
	// 	- file
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- folder
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// file
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SetUserProfilePathRulesRequestUserProfilePathRuleWhitePaths) String() string {
	return dara.Prettify(s)
}

func (s SetUserProfilePathRulesRequestUserProfilePathRuleWhitePaths) GoString() string {
	return s.String()
}

func (s *SetUserProfilePathRulesRequestUserProfilePathRuleWhitePaths) GetPath() *string {
	return s.Path
}

func (s *SetUserProfilePathRulesRequestUserProfilePathRuleWhitePaths) GetType() *string {
	return s.Type
}

func (s *SetUserProfilePathRulesRequestUserProfilePathRuleWhitePaths) SetPath(v string) *SetUserProfilePathRulesRequestUserProfilePathRuleWhitePaths {
	s.Path = &v
	return s
}

func (s *SetUserProfilePathRulesRequestUserProfilePathRuleWhitePaths) SetType(v string) *SetUserProfilePathRulesRequestUserProfilePathRuleWhitePaths {
	s.Type = &v
	return s
}

func (s *SetUserProfilePathRulesRequestUserProfilePathRuleWhitePaths) Validate() error {
	return dara.Validate(s)
}
