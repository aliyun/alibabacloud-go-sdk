// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentResp interface {
	dara.Model
	String() string
	GoString() string
	SetActivationKey(v string) *GetAgentResp
	GetActivationKey() *string
	SetAgentEndpoint(v string) *GetAgentResp
	GetAgentEndpoint() *string
	SetCreateTime(v string) *GetAgentResp
	GetCreateTime() *string
	SetDeployMethod(v string) *GetAgentResp
	GetDeployMethod() *string
	SetModifyTime(v string) *GetAgentResp
	GetModifyTime() *string
	SetName(v string) *GetAgentResp
	GetName() *string
	SetOwner(v string) *GetAgentResp
	GetOwner() *string
	SetTags(v string) *GetAgentResp
	GetTags() *string
	SetTunnelId(v string) *GetAgentResp
	GetTunnelId() *string
	SetVersion(v string) *GetAgentResp
	GetVersion() *string
}

type GetAgentResp struct {
	ActivationKey *string `json:"ActivationKey,omitempty" xml:"ActivationKey,omitempty"`
	// example:
	//
	// vpc
	AgentEndpoint *string `json:"AgentEndpoint,omitempty" xml:"AgentEndpoint,omitempty"`
	// example:
	//
	// 2024-05-01 12:00:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// default
	DeployMethod *string `json:"DeployMethod,omitempty" xml:"DeployMethod,omitempty"`
	// example:
	//
	// 2024-05-01 12:00:00
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// <your-agent-name>
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1***90**87***53*
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// K1:V1,K2:V2
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// example:
	//
	// ****13f-****-0567-****-af8d****446665
	TunnelId *string `json:"TunnelId,omitempty" xml:"TunnelId,omitempty"`
	// example:
	//
	// ****33f-****-4740-****-bc4f****332b3
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetAgentResp) String() string {
	return dara.Prettify(s)
}

func (s GetAgentResp) GoString() string {
	return s.String()
}

func (s *GetAgentResp) GetActivationKey() *string {
	return s.ActivationKey
}

func (s *GetAgentResp) GetAgentEndpoint() *string {
	return s.AgentEndpoint
}

func (s *GetAgentResp) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetAgentResp) GetDeployMethod() *string {
	return s.DeployMethod
}

func (s *GetAgentResp) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *GetAgentResp) GetName() *string {
	return s.Name
}

func (s *GetAgentResp) GetOwner() *string {
	return s.Owner
}

func (s *GetAgentResp) GetTags() *string {
	return s.Tags
}

func (s *GetAgentResp) GetTunnelId() *string {
	return s.TunnelId
}

func (s *GetAgentResp) GetVersion() *string {
	return s.Version
}

func (s *GetAgentResp) SetActivationKey(v string) *GetAgentResp {
	s.ActivationKey = &v
	return s
}

func (s *GetAgentResp) SetAgentEndpoint(v string) *GetAgentResp {
	s.AgentEndpoint = &v
	return s
}

func (s *GetAgentResp) SetCreateTime(v string) *GetAgentResp {
	s.CreateTime = &v
	return s
}

func (s *GetAgentResp) SetDeployMethod(v string) *GetAgentResp {
	s.DeployMethod = &v
	return s
}

func (s *GetAgentResp) SetModifyTime(v string) *GetAgentResp {
	s.ModifyTime = &v
	return s
}

func (s *GetAgentResp) SetName(v string) *GetAgentResp {
	s.Name = &v
	return s
}

func (s *GetAgentResp) SetOwner(v string) *GetAgentResp {
	s.Owner = &v
	return s
}

func (s *GetAgentResp) SetTags(v string) *GetAgentResp {
	s.Tags = &v
	return s
}

func (s *GetAgentResp) SetTunnelId(v string) *GetAgentResp {
	s.TunnelId = &v
	return s
}

func (s *GetAgentResp) SetVersion(v string) *GetAgentResp {
	s.Version = &v
	return s
}

func (s *GetAgentResp) Validate() error {
	return dara.Validate(s)
}
