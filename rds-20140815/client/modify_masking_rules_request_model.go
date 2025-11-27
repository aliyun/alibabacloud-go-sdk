// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMaskingRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *ModifyMaskingRulesRequest
	GetDBInstanceName() *string
	SetDBName(v string) *ModifyMaskingRulesRequest
	GetDBName() *string
	SetDefaultAlgo(v string) *ModifyMaskingRulesRequest
	GetDefaultAlgo() *string
	SetEnabled(v string) *ModifyMaskingRulesRequest
	GetEnabled() *string
	SetMaskingAlgo(v string) *ModifyMaskingRulesRequest
	GetMaskingAlgo() *string
	SetOwnerId(v string) *ModifyMaskingRulesRequest
	GetOwnerId() *string
	SetRegionId(v string) *ModifyMaskingRulesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyMaskingRulesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyMaskingRulesRequest
	GetResourceOwnerId() *int64
	SetRuleConfig(v *ModifyMaskingRulesRequestRuleConfig) *ModifyMaskingRulesRequest
	GetRuleConfig() *ModifyMaskingRulesRequestRuleConfig
	SetRuleName(v string) *ModifyMaskingRulesRequest
	GetRuleName() *string
}

type ModifyMaskingRulesRequest struct {
	// This parameter is required.
	DBInstanceName       *string                              `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	DBName               *string                              `json:"DBName,omitempty" xml:"DBName,omitempty"`
	DefaultAlgo          *string                              `json:"DefaultAlgo,omitempty" xml:"DefaultAlgo,omitempty"`
	Enabled              *string                              `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	MaskingAlgo          *string                              `json:"MaskingAlgo,omitempty" xml:"MaskingAlgo,omitempty"`
	OwnerId              *string                              `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string                              `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string                              `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                               `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RuleConfig           *ModifyMaskingRulesRequestRuleConfig `json:"RuleConfig,omitempty" xml:"RuleConfig,omitempty" type:"Struct"`
	// This parameter is required.
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s ModifyMaskingRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyMaskingRulesRequest) GoString() string {
	return s.String()
}

func (s *ModifyMaskingRulesRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *ModifyMaskingRulesRequest) GetDBName() *string {
	return s.DBName
}

func (s *ModifyMaskingRulesRequest) GetDefaultAlgo() *string {
	return s.DefaultAlgo
}

func (s *ModifyMaskingRulesRequest) GetEnabled() *string {
	return s.Enabled
}

func (s *ModifyMaskingRulesRequest) GetMaskingAlgo() *string {
	return s.MaskingAlgo
}

func (s *ModifyMaskingRulesRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *ModifyMaskingRulesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyMaskingRulesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyMaskingRulesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyMaskingRulesRequest) GetRuleConfig() *ModifyMaskingRulesRequestRuleConfig {
	return s.RuleConfig
}

func (s *ModifyMaskingRulesRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *ModifyMaskingRulesRequest) SetDBInstanceName(v string) *ModifyMaskingRulesRequest {
	s.DBInstanceName = &v
	return s
}

func (s *ModifyMaskingRulesRequest) SetDBName(v string) *ModifyMaskingRulesRequest {
	s.DBName = &v
	return s
}

func (s *ModifyMaskingRulesRequest) SetDefaultAlgo(v string) *ModifyMaskingRulesRequest {
	s.DefaultAlgo = &v
	return s
}

func (s *ModifyMaskingRulesRequest) SetEnabled(v string) *ModifyMaskingRulesRequest {
	s.Enabled = &v
	return s
}

func (s *ModifyMaskingRulesRequest) SetMaskingAlgo(v string) *ModifyMaskingRulesRequest {
	s.MaskingAlgo = &v
	return s
}

func (s *ModifyMaskingRulesRequest) SetOwnerId(v string) *ModifyMaskingRulesRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyMaskingRulesRequest) SetRegionId(v string) *ModifyMaskingRulesRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyMaskingRulesRequest) SetResourceOwnerAccount(v string) *ModifyMaskingRulesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyMaskingRulesRequest) SetResourceOwnerId(v int64) *ModifyMaskingRulesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyMaskingRulesRequest) SetRuleConfig(v *ModifyMaskingRulesRequestRuleConfig) *ModifyMaskingRulesRequest {
	s.RuleConfig = v
	return s
}

func (s *ModifyMaskingRulesRequest) SetRuleName(v string) *ModifyMaskingRulesRequest {
	s.RuleName = &v
	return s
}

func (s *ModifyMaskingRulesRequest) Validate() error {
	if s.RuleConfig != nil {
		if err := s.RuleConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyMaskingRulesRequestRuleConfig struct {
	Columns   []*string `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Repeated"`
	Databases []*string `json:"Databases,omitempty" xml:"Databases,omitempty" type:"Repeated"`
	Tables    []*string `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Repeated"`
}

func (s ModifyMaskingRulesRequestRuleConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyMaskingRulesRequestRuleConfig) GoString() string {
	return s.String()
}

func (s *ModifyMaskingRulesRequestRuleConfig) GetColumns() []*string {
	return s.Columns
}

func (s *ModifyMaskingRulesRequestRuleConfig) GetDatabases() []*string {
	return s.Databases
}

func (s *ModifyMaskingRulesRequestRuleConfig) GetTables() []*string {
	return s.Tables
}

func (s *ModifyMaskingRulesRequestRuleConfig) SetColumns(v []*string) *ModifyMaskingRulesRequestRuleConfig {
	s.Columns = v
	return s
}

func (s *ModifyMaskingRulesRequestRuleConfig) SetDatabases(v []*string) *ModifyMaskingRulesRequestRuleConfig {
	s.Databases = v
	return s
}

func (s *ModifyMaskingRulesRequestRuleConfig) SetTables(v []*string) *ModifyMaskingRulesRequestRuleConfig {
	s.Tables = v
	return s
}

func (s *ModifyMaskingRulesRequestRuleConfig) Validate() error {
	return dara.Validate(s)
}
