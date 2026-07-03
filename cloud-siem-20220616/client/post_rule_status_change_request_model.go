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
	// A JSON array of rule IDs.
	//
	// example:
	//
	// [123,345]
	Ids *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	// The status of the rule. Valid values:
	//
	// - true: enabled
	//
	// - false: disabled
	//
	// example:
	//
	// true
	InUse *bool `json:"InUse,omitempty" xml:"InUse,omitempty"`
	// The region where the Data Management center of threat analysis is located. Select a region based on the location of your assets. Valid values:
	//
	// - cn-hangzhou: Your assets are in the Chinese mainland or China (Hong Kong).
	//
	// - ap-southeast-1: Your assets are in a region outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID of a member. An administrator can specify this parameter to switch to the perspective of the member.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view.
	//
	// - 0: the view of the current Alibaba Cloud account.
	//
	// - 1: the view of all accounts that belong to the enterprise.
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// The type of the rule. Valid values:
	//
	// - predefine: predefined rule
	//
	// - customize: custom rule
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
