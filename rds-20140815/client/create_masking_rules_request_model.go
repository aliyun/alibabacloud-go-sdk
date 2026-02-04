// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMaskingRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *CreateMaskingRulesRequest
	GetDBInstanceName() *string
	SetDBName(v string) *CreateMaskingRulesRequest
	GetDBName() *string
	SetDefaultAlgo(v string) *CreateMaskingRulesRequest
	GetDefaultAlgo() *string
	SetMaskingAlgo(v string) *CreateMaskingRulesRequest
	GetMaskingAlgo() *string
	SetOwnerId(v string) *CreateMaskingRulesRequest
	GetOwnerId() *string
	SetRegionId(v string) *CreateMaskingRulesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateMaskingRulesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateMaskingRulesRequest
	GetResourceOwnerId() *int64
	SetRuleConfig(v *CreateMaskingRulesRequestRuleConfig) *CreateMaskingRulesRequest
	GetRuleConfig() *CreateMaskingRulesRequestRuleConfig
	SetRuleName(v string) *CreateMaskingRulesRequest
	GetRuleName() *string
}

type CreateMaskingRulesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// rm-t4n8t18o3*****d5
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// example:
	//
	// testdb
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// example:
	//
	// aes-128-gcm
	DefaultAlgo *string `json:"DefaultAlgo,omitempty" xml:"DefaultAlgo,omitempty"`
	// example:
	//
	// [{"name": "aes-128-gcm"},
	//
	//         {"name":"sm4-128-gcm"}]
	MaskingAlgo *string `json:"MaskingAlgo,omitempty" xml:"MaskingAlgo,omitempty"`
	OwnerId     *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// ap-southeast-1
	RegionId             *string                              `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string                              `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                               `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RuleConfig           *CreateMaskingRulesRequestRuleConfig `json:"RuleConfig,omitempty" xml:"RuleConfig,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// rulename1
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s CreateMaskingRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMaskingRulesRequest) GoString() string {
	return s.String()
}

func (s *CreateMaskingRulesRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *CreateMaskingRulesRequest) GetDBName() *string {
	return s.DBName
}

func (s *CreateMaskingRulesRequest) GetDefaultAlgo() *string {
	return s.DefaultAlgo
}

func (s *CreateMaskingRulesRequest) GetMaskingAlgo() *string {
	return s.MaskingAlgo
}

func (s *CreateMaskingRulesRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *CreateMaskingRulesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateMaskingRulesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateMaskingRulesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateMaskingRulesRequest) GetRuleConfig() *CreateMaskingRulesRequestRuleConfig {
	return s.RuleConfig
}

func (s *CreateMaskingRulesRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *CreateMaskingRulesRequest) SetDBInstanceName(v string) *CreateMaskingRulesRequest {
	s.DBInstanceName = &v
	return s
}

func (s *CreateMaskingRulesRequest) SetDBName(v string) *CreateMaskingRulesRequest {
	s.DBName = &v
	return s
}

func (s *CreateMaskingRulesRequest) SetDefaultAlgo(v string) *CreateMaskingRulesRequest {
	s.DefaultAlgo = &v
	return s
}

func (s *CreateMaskingRulesRequest) SetMaskingAlgo(v string) *CreateMaskingRulesRequest {
	s.MaskingAlgo = &v
	return s
}

func (s *CreateMaskingRulesRequest) SetOwnerId(v string) *CreateMaskingRulesRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateMaskingRulesRequest) SetRegionId(v string) *CreateMaskingRulesRequest {
	s.RegionId = &v
	return s
}

func (s *CreateMaskingRulesRequest) SetResourceOwnerAccount(v string) *CreateMaskingRulesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateMaskingRulesRequest) SetResourceOwnerId(v int64) *CreateMaskingRulesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateMaskingRulesRequest) SetRuleConfig(v *CreateMaskingRulesRequestRuleConfig) *CreateMaskingRulesRequest {
	s.RuleConfig = v
	return s
}

func (s *CreateMaskingRulesRequest) SetRuleName(v string) *CreateMaskingRulesRequest {
	s.RuleName = &v
	return s
}

func (s *CreateMaskingRulesRequest) Validate() error {
	if s.RuleConfig != nil {
		if err := s.RuleConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateMaskingRulesRequestRuleConfig struct {
	Columns   []*string `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Repeated"`
	Databases []*string `json:"Databases,omitempty" xml:"Databases,omitempty" type:"Repeated"`
	Tables    []*string `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Repeated"`
}

func (s CreateMaskingRulesRequestRuleConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateMaskingRulesRequestRuleConfig) GoString() string {
	return s.String()
}

func (s *CreateMaskingRulesRequestRuleConfig) GetColumns() []*string {
	return s.Columns
}

func (s *CreateMaskingRulesRequestRuleConfig) GetDatabases() []*string {
	return s.Databases
}

func (s *CreateMaskingRulesRequestRuleConfig) GetTables() []*string {
	return s.Tables
}

func (s *CreateMaskingRulesRequestRuleConfig) SetColumns(v []*string) *CreateMaskingRulesRequestRuleConfig {
	s.Columns = v
	return s
}

func (s *CreateMaskingRulesRequestRuleConfig) SetDatabases(v []*string) *CreateMaskingRulesRequestRuleConfig {
	s.Databases = v
	return s
}

func (s *CreateMaskingRulesRequestRuleConfig) SetTables(v []*string) *CreateMaskingRulesRequestRuleConfig {
	s.Tables = v
	return s
}

func (s *CreateMaskingRulesRequestRuleConfig) Validate() error {
	return dara.Validate(s)
}
