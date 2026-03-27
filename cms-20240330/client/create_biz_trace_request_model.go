// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBizTraceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdvancedConfig(v string) *CreateBizTraceRequest
	GetAdvancedConfig() *string
	SetBizTraceCode(v string) *CreateBizTraceRequest
	GetBizTraceCode() *string
	SetBizTraceName(v string) *CreateBizTraceRequest
	GetBizTraceName() *string
	SetRuleConfig(v string) *CreateBizTraceRequest
	GetRuleConfig() *string
	SetWorkspace(v string) *CreateBizTraceRequest
	GetWorkspace() *string
}

type CreateBizTraceRequest struct {
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

func (s CreateBizTraceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateBizTraceRequest) GoString() string {
	return s.String()
}

func (s *CreateBizTraceRequest) GetAdvancedConfig() *string {
	return s.AdvancedConfig
}

func (s *CreateBizTraceRequest) GetBizTraceCode() *string {
	return s.BizTraceCode
}

func (s *CreateBizTraceRequest) GetBizTraceName() *string {
	return s.BizTraceName
}

func (s *CreateBizTraceRequest) GetRuleConfig() *string {
	return s.RuleConfig
}

func (s *CreateBizTraceRequest) GetWorkspace() *string {
	return s.Workspace
}

func (s *CreateBizTraceRequest) SetAdvancedConfig(v string) *CreateBizTraceRequest {
	s.AdvancedConfig = &v
	return s
}

func (s *CreateBizTraceRequest) SetBizTraceCode(v string) *CreateBizTraceRequest {
	s.BizTraceCode = &v
	return s
}

func (s *CreateBizTraceRequest) SetBizTraceName(v string) *CreateBizTraceRequest {
	s.BizTraceName = &v
	return s
}

func (s *CreateBizTraceRequest) SetRuleConfig(v string) *CreateBizTraceRequest {
	s.RuleConfig = &v
	return s
}

func (s *CreateBizTraceRequest) SetWorkspace(v string) *CreateBizTraceRequest {
	s.Workspace = &v
	return s
}

func (s *CreateBizTraceRequest) Validate() error {
	return dara.Validate(s)
}
