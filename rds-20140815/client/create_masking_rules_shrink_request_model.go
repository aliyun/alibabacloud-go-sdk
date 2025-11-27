// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMaskingRulesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *CreateMaskingRulesShrinkRequest
	GetDBInstanceName() *string
	SetDBName(v string) *CreateMaskingRulesShrinkRequest
	GetDBName() *string
	SetDefaultAlgo(v string) *CreateMaskingRulesShrinkRequest
	GetDefaultAlgo() *string
	SetMaskingAlgo(v string) *CreateMaskingRulesShrinkRequest
	GetMaskingAlgo() *string
	SetOwnerId(v string) *CreateMaskingRulesShrinkRequest
	GetOwnerId() *string
	SetRegionId(v string) *CreateMaskingRulesShrinkRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateMaskingRulesShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateMaskingRulesShrinkRequest
	GetResourceOwnerId() *int64
	SetRuleConfigShrink(v string) *CreateMaskingRulesShrinkRequest
	GetRuleConfigShrink() *string
	SetRuleName(v string) *CreateMaskingRulesShrinkRequest
	GetRuleName() *string
}

type CreateMaskingRulesShrinkRequest struct {
	// This parameter is required.
	DBInstanceName       *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	DBName               *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	DefaultAlgo          *string `json:"DefaultAlgo,omitempty" xml:"DefaultAlgo,omitempty"`
	MaskingAlgo          *string `json:"MaskingAlgo,omitempty" xml:"MaskingAlgo,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RuleConfigShrink     *string `json:"RuleConfig,omitempty" xml:"RuleConfig,omitempty"`
	// This parameter is required.
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s CreateMaskingRulesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMaskingRulesShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateMaskingRulesShrinkRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *CreateMaskingRulesShrinkRequest) GetDBName() *string {
	return s.DBName
}

func (s *CreateMaskingRulesShrinkRequest) GetDefaultAlgo() *string {
	return s.DefaultAlgo
}

func (s *CreateMaskingRulesShrinkRequest) GetMaskingAlgo() *string {
	return s.MaskingAlgo
}

func (s *CreateMaskingRulesShrinkRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *CreateMaskingRulesShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateMaskingRulesShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateMaskingRulesShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateMaskingRulesShrinkRequest) GetRuleConfigShrink() *string {
	return s.RuleConfigShrink
}

func (s *CreateMaskingRulesShrinkRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *CreateMaskingRulesShrinkRequest) SetDBInstanceName(v string) *CreateMaskingRulesShrinkRequest {
	s.DBInstanceName = &v
	return s
}

func (s *CreateMaskingRulesShrinkRequest) SetDBName(v string) *CreateMaskingRulesShrinkRequest {
	s.DBName = &v
	return s
}

func (s *CreateMaskingRulesShrinkRequest) SetDefaultAlgo(v string) *CreateMaskingRulesShrinkRequest {
	s.DefaultAlgo = &v
	return s
}

func (s *CreateMaskingRulesShrinkRequest) SetMaskingAlgo(v string) *CreateMaskingRulesShrinkRequest {
	s.MaskingAlgo = &v
	return s
}

func (s *CreateMaskingRulesShrinkRequest) SetOwnerId(v string) *CreateMaskingRulesShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateMaskingRulesShrinkRequest) SetRegionId(v string) *CreateMaskingRulesShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateMaskingRulesShrinkRequest) SetResourceOwnerAccount(v string) *CreateMaskingRulesShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateMaskingRulesShrinkRequest) SetResourceOwnerId(v int64) *CreateMaskingRulesShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateMaskingRulesShrinkRequest) SetRuleConfigShrink(v string) *CreateMaskingRulesShrinkRequest {
	s.RuleConfigShrink = &v
	return s
}

func (s *CreateMaskingRulesShrinkRequest) SetRuleName(v string) *CreateMaskingRulesShrinkRequest {
	s.RuleName = &v
	return s
}

func (s *CreateMaskingRulesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
