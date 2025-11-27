// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBizTraceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdvancedConfig(v string) *UpdateBizTraceRequest
	GetAdvancedConfig() *string
	SetBizTraceName(v string) *UpdateBizTraceRequest
	GetBizTraceName() *string
	SetRuleConfig(v string) *UpdateBizTraceRequest
	GetRuleConfig() *string
	SetWorkspace(v string) *UpdateBizTraceRequest
	GetWorkspace() *string
}

type UpdateBizTraceRequest struct {
	// example:
	//
	// {"sample":{"strategy":"BY_APP"}}
	AdvancedConfig *string `json:"advancedConfig,omitempty" xml:"advancedConfig,omitempty"`
	// example:
	//
	// just test
	BizTraceName *string `json:"bizTraceName,omitempty" xml:"bizTraceName,omitempty"`
	// example:
	//
	// [{"entrancePid":"xxx@d9w3jd9j3","rpcMatcher":{"matchType":"EQUALS","pattern":"/"},"characteristics":{"operation":"OR","rules":[{"target":"CUSTOM_EXTRACT","id":"oi0b3bb7","key":"biz.test","matcher":{"matchType":"CONTAINS","pattern":["1"]}}]}}]
	RuleConfig *string `json:"ruleConfig,omitempty" xml:"ruleConfig,omitempty"`
	// example:
	//
	// default-cms-xxxxxxx-cn-hangzhou
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s UpdateBizTraceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateBizTraceRequest) GoString() string {
	return s.String()
}

func (s *UpdateBizTraceRequest) GetAdvancedConfig() *string {
	return s.AdvancedConfig
}

func (s *UpdateBizTraceRequest) GetBizTraceName() *string {
	return s.BizTraceName
}

func (s *UpdateBizTraceRequest) GetRuleConfig() *string {
	return s.RuleConfig
}

func (s *UpdateBizTraceRequest) GetWorkspace() *string {
	return s.Workspace
}

func (s *UpdateBizTraceRequest) SetAdvancedConfig(v string) *UpdateBizTraceRequest {
	s.AdvancedConfig = &v
	return s
}

func (s *UpdateBizTraceRequest) SetBizTraceName(v string) *UpdateBizTraceRequest {
	s.BizTraceName = &v
	return s
}

func (s *UpdateBizTraceRequest) SetRuleConfig(v string) *UpdateBizTraceRequest {
	s.RuleConfig = &v
	return s
}

func (s *UpdateBizTraceRequest) SetWorkspace(v string) *UpdateBizTraceRequest {
	s.Workspace = &v
	return s
}

func (s *UpdateBizTraceRequest) Validate() error {
	return dara.Validate(s)
}
