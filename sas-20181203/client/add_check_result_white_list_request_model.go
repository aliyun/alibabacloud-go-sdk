// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCheckResultWhiteListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckIds(v []*int64) *AddCheckResultWhiteListRequest
	GetCheckIds() []*int64
	SetClientToken(v string) *AddCheckResultWhiteListRequest
	GetClientToken() *string
	SetInstanceIds(v []*string) *AddCheckResultWhiteListRequest
	GetInstanceIds() []*string
	SetRemark(v string) *AddCheckResultWhiteListRequest
	GetRemark() *string
	SetRuleType(v string) *AddCheckResultWhiteListRequest
	GetRuleType() *string
}

type AddCheckResultWhiteListRequest struct {
	// The IDs of the check items.
	//
	// > Call the [ListCheckResult](~~ListCheckResult~~) operation to obtain this parameter.
	CheckIds []*int64 `json:"CheckIds,omitempty" xml:"CheckIds,omitempty" type:"Repeated"`
	// The client token that is used to ensure the idempotence of the request. Different requests should use different tokens. The token supports only ASCII characters and cannot exceed 64 characters in length.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instance IDs of the cloud service instances to add to the whitelist. Separate multiple instance IDs with commas (,).
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The remarks. Maximum length: 65,535 bytes.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The rule type. Default value: **WHITE**. Valid values:
	//
	// - **WHITE**: adds to the whitelist.
	//
	// example:
	//
	// WHITE
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
}

func (s AddCheckResultWhiteListRequest) String() string {
	return dara.Prettify(s)
}

func (s AddCheckResultWhiteListRequest) GoString() string {
	return s.String()
}

func (s *AddCheckResultWhiteListRequest) GetCheckIds() []*int64 {
	return s.CheckIds
}

func (s *AddCheckResultWhiteListRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AddCheckResultWhiteListRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *AddCheckResultWhiteListRequest) GetRemark() *string {
	return s.Remark
}

func (s *AddCheckResultWhiteListRequest) GetRuleType() *string {
	return s.RuleType
}

func (s *AddCheckResultWhiteListRequest) SetCheckIds(v []*int64) *AddCheckResultWhiteListRequest {
	s.CheckIds = v
	return s
}

func (s *AddCheckResultWhiteListRequest) SetClientToken(v string) *AddCheckResultWhiteListRequest {
	s.ClientToken = &v
	return s
}

func (s *AddCheckResultWhiteListRequest) SetInstanceIds(v []*string) *AddCheckResultWhiteListRequest {
	s.InstanceIds = v
	return s
}

func (s *AddCheckResultWhiteListRequest) SetRemark(v string) *AddCheckResultWhiteListRequest {
	s.Remark = &v
	return s
}

func (s *AddCheckResultWhiteListRequest) SetRuleType(v string) *AddCheckResultWhiteListRequest {
	s.RuleType = &v
	return s
}

func (s *AddCheckResultWhiteListRequest) Validate() error {
	return dara.Validate(s)
}
