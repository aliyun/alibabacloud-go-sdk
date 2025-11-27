// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBizTraceConfig interface {
	dara.Model
	String() string
	GoString() string
	SetAdvancedConfig(v string) *BizTraceConfig
	GetAdvancedConfig() *string
	SetBizTraceCode(v string) *BizTraceConfig
	GetBizTraceCode() *string
	SetBizTraceId(v string) *BizTraceConfig
	GetBizTraceId() *string
	SetBizTraceName(v string) *BizTraceConfig
	GetBizTraceName() *string
	SetCreateTime(v string) *BizTraceConfig
	GetCreateTime() *string
	SetRegionId(v string) *BizTraceConfig
	GetRegionId() *string
	SetRuleConfig(v string) *BizTraceConfig
	GetRuleConfig() *string
	SetWorkspace(v string) *BizTraceConfig
	GetWorkspace() *string
}

type BizTraceConfig struct {
	// example:
	//
	// {"sample":{"strategy":"BY_APP"}}
	AdvancedConfig *string `json:"advancedConfig,omitempty" xml:"advancedConfig,omitempty"`
	// example:
	//
	// label_env
	BizTraceCode *string `json:"bizTraceCode,omitempty" xml:"bizTraceCode,omitempty"`
	// example:
	//
	// e339260ed64c95d
	BizTraceId *string `json:"bizTraceId,omitempty" xml:"bizTraceId,omitempty"`
	// example:
	//
	// just test
	BizTraceName *string `json:"bizTraceName,omitempty" xml:"bizTraceName,omitempty"`
	// example:
	//
	// 2025-12-12 00:00:00
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// example:
	//
	// [{"entrancePid":"xxxxx@b57c44xx6e86","rpcMatcher":{"matchType":"EQUALS","pattern":"/createApp"},"characteristics":{"operation":"AND","rules":[{"target":"CUSTOM_EXTRACT","matcher":{"matchType":"CONTAINS","pattern":[]}}]}}]
	RuleConfig *string `json:"ruleConfig,omitempty" xml:"ruleConfig,omitempty"`
	// example:
	//
	// default-cms-xxxxxxx-cn-hangzhou
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s BizTraceConfig) String() string {
	return dara.Prettify(s)
}

func (s BizTraceConfig) GoString() string {
	return s.String()
}

func (s *BizTraceConfig) GetAdvancedConfig() *string {
	return s.AdvancedConfig
}

func (s *BizTraceConfig) GetBizTraceCode() *string {
	return s.BizTraceCode
}

func (s *BizTraceConfig) GetBizTraceId() *string {
	return s.BizTraceId
}

func (s *BizTraceConfig) GetBizTraceName() *string {
	return s.BizTraceName
}

func (s *BizTraceConfig) GetCreateTime() *string {
	return s.CreateTime
}

func (s *BizTraceConfig) GetRegionId() *string {
	return s.RegionId
}

func (s *BizTraceConfig) GetRuleConfig() *string {
	return s.RuleConfig
}

func (s *BizTraceConfig) GetWorkspace() *string {
	return s.Workspace
}

func (s *BizTraceConfig) SetAdvancedConfig(v string) *BizTraceConfig {
	s.AdvancedConfig = &v
	return s
}

func (s *BizTraceConfig) SetBizTraceCode(v string) *BizTraceConfig {
	s.BizTraceCode = &v
	return s
}

func (s *BizTraceConfig) SetBizTraceId(v string) *BizTraceConfig {
	s.BizTraceId = &v
	return s
}

func (s *BizTraceConfig) SetBizTraceName(v string) *BizTraceConfig {
	s.BizTraceName = &v
	return s
}

func (s *BizTraceConfig) SetCreateTime(v string) *BizTraceConfig {
	s.CreateTime = &v
	return s
}

func (s *BizTraceConfig) SetRegionId(v string) *BizTraceConfig {
	s.RegionId = &v
	return s
}

func (s *BizTraceConfig) SetRuleConfig(v string) *BizTraceConfig {
	s.RuleConfig = &v
	return s
}

func (s *BizTraceConfig) SetWorkspace(v string) *BizTraceConfig {
	s.Workspace = &v
	return s
}

func (s *BizTraceConfig) Validate() error {
	return dara.Validate(s)
}
