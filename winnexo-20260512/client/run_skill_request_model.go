// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunSkillRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArguments(v map[string]interface{}) *RunSkillRequest
	GetArguments() map[string]interface{}
	SetClientToken(v string) *RunSkillRequest
	GetClientToken() *string
	SetModel(v string) *RunSkillRequest
	GetModel() *string
	SetOperatingObjectName(v string) *RunSkillRequest
	GetOperatingObjectName() *string
	SetSkillCode(v string) *RunSkillRequest
	GetSkillCode() *string
	SetSkillName(v string) *RunSkillRequest
	GetSkillName() *string
	SetTenantId(v string) *RunSkillRequest
	GetTenantId() *string
}

type RunSkillRequest struct {
	Arguments map[string]interface{} `json:"arguments,omitempty" xml:"arguments,omitempty"`
	// 幂等 token，调用方自行生成；当前版本仅记录到 metadata，未做去重
	//
	// example:
	//
	// string_value
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// 抽象模型名（模型档位），不传默认 standard
	//
	// example:
	//
	// quick
	Model *string `json:"model,omitempty" xml:"model,omitempty"`
	// 数字员工名称；用于按绑定关系计算 CodeAgent allowedSkills 白名单
	//
	// example:
	//
	// string_value
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
	// 技能编码（全局唯一），优先级高于 skillName
	//
	// example:
	//
	// string_value
	SkillCode *string `json:"skillCode,omitempty" xml:"skillCode,omitempty"`
	// 技能名称，未传 skillCode 时使用；租户范围内必须唯一
	//
	// example:
	//
	// string_value
	SkillName *string `json:"skillName,omitempty" xml:"skillName,omitempty"`
	// 租户ID，公共参数，缺省时使用调用方默认租户
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s RunSkillRequest) String() string {
	return dara.Prettify(s)
}

func (s RunSkillRequest) GoString() string {
	return s.String()
}

func (s *RunSkillRequest) GetArguments() map[string]interface{} {
	return s.Arguments
}

func (s *RunSkillRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RunSkillRequest) GetModel() *string {
	return s.Model
}

func (s *RunSkillRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *RunSkillRequest) GetSkillCode() *string {
	return s.SkillCode
}

func (s *RunSkillRequest) GetSkillName() *string {
	return s.SkillName
}

func (s *RunSkillRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *RunSkillRequest) SetArguments(v map[string]interface{}) *RunSkillRequest {
	s.Arguments = v
	return s
}

func (s *RunSkillRequest) SetClientToken(v string) *RunSkillRequest {
	s.ClientToken = &v
	return s
}

func (s *RunSkillRequest) SetModel(v string) *RunSkillRequest {
	s.Model = &v
	return s
}

func (s *RunSkillRequest) SetOperatingObjectName(v string) *RunSkillRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *RunSkillRequest) SetSkillCode(v string) *RunSkillRequest {
	s.SkillCode = &v
	return s
}

func (s *RunSkillRequest) SetSkillName(v string) *RunSkillRequest {
	s.SkillName = &v
	return s
}

func (s *RunSkillRequest) SetTenantId(v string) *RunSkillRequest {
	s.TenantId = &v
	return s
}

func (s *RunSkillRequest) Validate() error {
	return dara.Validate(s)
}
