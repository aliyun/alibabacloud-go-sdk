// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunSkillShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArgumentsShrink(v string) *RunSkillShrinkRequest
	GetArgumentsShrink() *string
	SetClientToken(v string) *RunSkillShrinkRequest
	GetClientToken() *string
	SetModel(v string) *RunSkillShrinkRequest
	GetModel() *string
	SetOperatingObjectName(v string) *RunSkillShrinkRequest
	GetOperatingObjectName() *string
	SetSkillCode(v string) *RunSkillShrinkRequest
	GetSkillCode() *string
	SetSkillName(v string) *RunSkillShrinkRequest
	GetSkillName() *string
	SetTenantId(v string) *RunSkillShrinkRequest
	GetTenantId() *string
}

type RunSkillShrinkRequest struct {
	ArgumentsShrink *string `json:"arguments,omitempty" xml:"arguments,omitempty"`
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

func (s RunSkillShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RunSkillShrinkRequest) GoString() string {
	return s.String()
}

func (s *RunSkillShrinkRequest) GetArgumentsShrink() *string {
	return s.ArgumentsShrink
}

func (s *RunSkillShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RunSkillShrinkRequest) GetModel() *string {
	return s.Model
}

func (s *RunSkillShrinkRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *RunSkillShrinkRequest) GetSkillCode() *string {
	return s.SkillCode
}

func (s *RunSkillShrinkRequest) GetSkillName() *string {
	return s.SkillName
}

func (s *RunSkillShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *RunSkillShrinkRequest) SetArgumentsShrink(v string) *RunSkillShrinkRequest {
	s.ArgumentsShrink = &v
	return s
}

func (s *RunSkillShrinkRequest) SetClientToken(v string) *RunSkillShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *RunSkillShrinkRequest) SetModel(v string) *RunSkillShrinkRequest {
	s.Model = &v
	return s
}

func (s *RunSkillShrinkRequest) SetOperatingObjectName(v string) *RunSkillShrinkRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *RunSkillShrinkRequest) SetSkillCode(v string) *RunSkillShrinkRequest {
	s.SkillCode = &v
	return s
}

func (s *RunSkillShrinkRequest) SetSkillName(v string) *RunSkillShrinkRequest {
	s.SkillName = &v
	return s
}

func (s *RunSkillShrinkRequest) SetTenantId(v string) *RunSkillShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *RunSkillShrinkRequest) Validate() error {
	return dara.Validate(s)
}
