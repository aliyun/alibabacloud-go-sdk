// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMaskingRulesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *ModifyMaskingRulesShrinkRequest
	GetDBInstanceName() *string
	SetDBName(v string) *ModifyMaskingRulesShrinkRequest
	GetDBName() *string
	SetDefaultAlgo(v string) *ModifyMaskingRulesShrinkRequest
	GetDefaultAlgo() *string
	SetEnabled(v string) *ModifyMaskingRulesShrinkRequest
	GetEnabled() *string
	SetMaskingAlgo(v string) *ModifyMaskingRulesShrinkRequest
	GetMaskingAlgo() *string
	SetOwnerId(v string) *ModifyMaskingRulesShrinkRequest
	GetOwnerId() *string
	SetRegionId(v string) *ModifyMaskingRulesShrinkRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyMaskingRulesShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyMaskingRulesShrinkRequest
	GetResourceOwnerId() *int64
	SetRuleConfigShrink(v string) *ModifyMaskingRulesShrinkRequest
	GetRuleConfigShrink() *string
	SetRuleName(v string) *ModifyMaskingRulesShrinkRequest
	GetRuleName() *string
}

type ModifyMaskingRulesShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// rm-t4n8t18o******6d5
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// example:
	//
	// myDB
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// example:
	//
	// sm4-128-gcm
	DefaultAlgo *string `json:"DefaultAlgo,omitempty" xml:"DefaultAlgo,omitempty"`
	// example:
	//
	// true
	Enabled *string `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// example:
	//
	// [{"name": "sha256"},
	//
	//         {"name":"sm4-128-gcm"}]
	MaskingAlgo *string `json:"MaskingAlgo,omitempty" xml:"MaskingAlgo,omitempty"`
	OwnerId     *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// ap-southeast-1
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RuleConfigShrink     *string `json:"RuleConfig,omitempty" xml:"RuleConfig,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// rulename1
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s ModifyMaskingRulesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyMaskingRulesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyMaskingRulesShrinkRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *ModifyMaskingRulesShrinkRequest) GetDBName() *string {
	return s.DBName
}

func (s *ModifyMaskingRulesShrinkRequest) GetDefaultAlgo() *string {
	return s.DefaultAlgo
}

func (s *ModifyMaskingRulesShrinkRequest) GetEnabled() *string {
	return s.Enabled
}

func (s *ModifyMaskingRulesShrinkRequest) GetMaskingAlgo() *string {
	return s.MaskingAlgo
}

func (s *ModifyMaskingRulesShrinkRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *ModifyMaskingRulesShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyMaskingRulesShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyMaskingRulesShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyMaskingRulesShrinkRequest) GetRuleConfigShrink() *string {
	return s.RuleConfigShrink
}

func (s *ModifyMaskingRulesShrinkRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *ModifyMaskingRulesShrinkRequest) SetDBInstanceName(v string) *ModifyMaskingRulesShrinkRequest {
	s.DBInstanceName = &v
	return s
}

func (s *ModifyMaskingRulesShrinkRequest) SetDBName(v string) *ModifyMaskingRulesShrinkRequest {
	s.DBName = &v
	return s
}

func (s *ModifyMaskingRulesShrinkRequest) SetDefaultAlgo(v string) *ModifyMaskingRulesShrinkRequest {
	s.DefaultAlgo = &v
	return s
}

func (s *ModifyMaskingRulesShrinkRequest) SetEnabled(v string) *ModifyMaskingRulesShrinkRequest {
	s.Enabled = &v
	return s
}

func (s *ModifyMaskingRulesShrinkRequest) SetMaskingAlgo(v string) *ModifyMaskingRulesShrinkRequest {
	s.MaskingAlgo = &v
	return s
}

func (s *ModifyMaskingRulesShrinkRequest) SetOwnerId(v string) *ModifyMaskingRulesShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyMaskingRulesShrinkRequest) SetRegionId(v string) *ModifyMaskingRulesShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyMaskingRulesShrinkRequest) SetResourceOwnerAccount(v string) *ModifyMaskingRulesShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyMaskingRulesShrinkRequest) SetResourceOwnerId(v int64) *ModifyMaskingRulesShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyMaskingRulesShrinkRequest) SetRuleConfigShrink(v string) *ModifyMaskingRulesShrinkRequest {
	s.RuleConfigShrink = &v
	return s
}

func (s *ModifyMaskingRulesShrinkRequest) SetRuleName(v string) *ModifyMaskingRulesShrinkRequest {
	s.RuleName = &v
	return s
}

func (s *ModifyMaskingRulesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
