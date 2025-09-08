// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPostRuleStatusChangeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIds(v string) *PostRuleStatusChangeRequest
	GetIds() *string
	SetInUse(v bool) *PostRuleStatusChangeRequest
	GetInUse() *bool
	SetRegionId(v string) *PostRuleStatusChangeRequest
	GetRegionId() *string
	SetRoleFor(v int64) *PostRuleStatusChangeRequest
	GetRoleFor() *int64
	SetRoleType(v int32) *PostRuleStatusChangeRequest
	GetRoleType() *int32
	SetRuleType(v string) *PostRuleStatusChangeRequest
	GetRuleType() *string
}

type PostRuleStatusChangeRequest struct {
	// The rule IDs. The value is a JSON array.
	//
	// example:
	//
	// [123,345]
	Ids *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	// Specifies whether to enable the rule. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	InUse *bool `json:"InUse,omitempty" xml:"InUse,omitempty"`
	// The region in which the data management center of the threat analysis feature resides. Specify this parameter based on the regions in which your assets reside. Valid values:
	//
	// 	- cn-hangzhou: Your assets reside in regions in China.
	//
	// 	- ap-southeast-1: Your assets reside in regions outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the account that you switch from the management account.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view. Valid values:
	//
	// - 0: the current Alibaba Cloud account
	//
	// - 1: the global account
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// The type of the rule. Valid values:
	//
	// 	- predefine
	//
	// 	- customize
	//
	// example:
	//
	// customize
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
}

func (s PostRuleStatusChangeRequest) String() string {
	return dara.Prettify(s)
}

func (s PostRuleStatusChangeRequest) GoString() string {
	return s.String()
}

func (s *PostRuleStatusChangeRequest) GetIds() *string {
	return s.Ids
}

func (s *PostRuleStatusChangeRequest) GetInUse() *bool {
	return s.InUse
}

func (s *PostRuleStatusChangeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *PostRuleStatusChangeRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *PostRuleStatusChangeRequest) GetRoleType() *int32 {
	return s.RoleType
}

func (s *PostRuleStatusChangeRequest) GetRuleType() *string {
	return s.RuleType
}

func (s *PostRuleStatusChangeRequest) SetIds(v string) *PostRuleStatusChangeRequest {
	s.Ids = &v
	return s
}

func (s *PostRuleStatusChangeRequest) SetInUse(v bool) *PostRuleStatusChangeRequest {
	s.InUse = &v
	return s
}

func (s *PostRuleStatusChangeRequest) SetRegionId(v string) *PostRuleStatusChangeRequest {
	s.RegionId = &v
	return s
}

func (s *PostRuleStatusChangeRequest) SetRoleFor(v int64) *PostRuleStatusChangeRequest {
	s.RoleFor = &v
	return s
}

func (s *PostRuleStatusChangeRequest) SetRoleType(v int32) *PostRuleStatusChangeRequest {
	s.RoleType = &v
	return s
}

func (s *PostRuleStatusChangeRequest) SetRuleType(v string) *PostRuleStatusChangeRequest {
	s.RuleType = &v
	return s
}

func (s *PostRuleStatusChangeRequest) Validate() error {
	return dara.Validate(s)
}
