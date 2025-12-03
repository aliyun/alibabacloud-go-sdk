// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveCheckResultWhiteListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckGroupId(v string) *RemoveCheckResultWhiteListRequest
	GetCheckGroupId() *string
	SetCheckIds(v []*int64) *RemoveCheckResultWhiteListRequest
	GetCheckIds() []*int64
	SetInstanceIds(v []*string) *RemoveCheckResultWhiteListRequest
	GetInstanceIds() []*string
	SetRuleId(v int64) *RemoveCheckResultWhiteListRequest
	GetRuleId() *int64
	SetType(v string) *RemoveCheckResultWhiteListRequest
	GetType() *string
}

type RemoveCheckResultWhiteListRequest struct {
	// This parameter is deprecated.
	//
	// example:
	//
	// Deprecated
	CheckGroupId *string `json:"CheckGroupId,omitempty" xml:"CheckGroupId,omitempty"`
	// The IDs of the check items.
	CheckIds []*int64 `json:"CheckIds,omitempty" xml:"CheckIds,omitempty" type:"Repeated"`
	// A set of cloud product instance IDs that require validation.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The ID of the whitelist rule.
	//
	// >  You can call the [ListCheckResult](~~ListCheckResult~~) operation to query the IDs of whitelist rules.
	//
	// example:
	//
	// 22
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// Deprecated
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s RemoveCheckResultWhiteListRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveCheckResultWhiteListRequest) GoString() string {
	return s.String()
}

func (s *RemoveCheckResultWhiteListRequest) GetCheckGroupId() *string {
	return s.CheckGroupId
}

func (s *RemoveCheckResultWhiteListRequest) GetCheckIds() []*int64 {
	return s.CheckIds
}

func (s *RemoveCheckResultWhiteListRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *RemoveCheckResultWhiteListRequest) GetRuleId() *int64 {
	return s.RuleId
}

func (s *RemoveCheckResultWhiteListRequest) GetType() *string {
	return s.Type
}

func (s *RemoveCheckResultWhiteListRequest) SetCheckGroupId(v string) *RemoveCheckResultWhiteListRequest {
	s.CheckGroupId = &v
	return s
}

func (s *RemoveCheckResultWhiteListRequest) SetCheckIds(v []*int64) *RemoveCheckResultWhiteListRequest {
	s.CheckIds = v
	return s
}

func (s *RemoveCheckResultWhiteListRequest) SetInstanceIds(v []*string) *RemoveCheckResultWhiteListRequest {
	s.InstanceIds = v
	return s
}

func (s *RemoveCheckResultWhiteListRequest) SetRuleId(v int64) *RemoveCheckResultWhiteListRequest {
	s.RuleId = &v
	return s
}

func (s *RemoveCheckResultWhiteListRequest) SetType(v string) *RemoveCheckResultWhiteListRequest {
	s.Type = &v
	return s
}

func (s *RemoveCheckResultWhiteListRequest) Validate() error {
	return dara.Validate(s)
}
