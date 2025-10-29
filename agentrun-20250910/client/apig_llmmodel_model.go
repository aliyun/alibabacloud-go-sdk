// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApigLLMModel interface {
	dara.Model
	String() string
	GoString() string
	SetAddress(v string) *ApigLLMModel
	GetAddress() *string
	SetApiKey(v string) *ApigLLMModel
	GetApiKey() *string
	SetCreatedTime(v string) *ApigLLMModel
	GetCreatedTime() *string
	SetDesc(v string) *ApigLLMModel
	GetDesc() *string
	SetGatewayId(v string) *ApigLLMModel
	GetGatewayId() *string
	SetModelId(v string) *ApigLLMModel
	GetModelId() *string
	SetModels(v string) *ApigLLMModel
	GetModels() *string
	SetModelsWeight(v string) *ApigLLMModel
	GetModelsWeight() *string
	SetName(v string) *ApigLLMModel
	GetName() *string
	SetProvider(v string) *ApigLLMModel
	GetProvider() *string
	SetTargetId(v string) *ApigLLMModel
	GetTargetId() *string
	SetTenantId(v string) *ApigLLMModel
	GetTenantId() *string
	SetType(v string) *ApigLLMModel
	GetType() *string
	SetUpdateTime(v string) *ApigLLMModel
	GetUpdateTime() *string
}

type ApigLLMModel struct {
	Address      *string `json:"address,omitempty" xml:"address,omitempty"`
	ApiKey       *string `json:"apiKey,omitempty" xml:"apiKey,omitempty"`
	CreatedTime  *string `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	Desc         *string `json:"desc,omitempty" xml:"desc,omitempty"`
	GatewayId    *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	ModelId      *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	Models       *string `json:"models,omitempty" xml:"models,omitempty"`
	ModelsWeight *string `json:"modelsWeight,omitempty" xml:"modelsWeight,omitempty"`
	Name         *string `json:"name,omitempty" xml:"name,omitempty"`
	Provider     *string `json:"provider,omitempty" xml:"provider,omitempty"`
	TargetId     *string `json:"targetId,omitempty" xml:"targetId,omitempty"`
	TenantId     *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	Type         *string `json:"type,omitempty" xml:"type,omitempty"`
	UpdateTime   *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s ApigLLMModel) String() string {
	return dara.Prettify(s)
}

func (s ApigLLMModel) GoString() string {
	return s.String()
}

func (s *ApigLLMModel) GetAddress() *string {
	return s.Address
}

func (s *ApigLLMModel) GetApiKey() *string {
	return s.ApiKey
}

func (s *ApigLLMModel) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *ApigLLMModel) GetDesc() *string {
	return s.Desc
}

func (s *ApigLLMModel) GetGatewayId() *string {
	return s.GatewayId
}

func (s *ApigLLMModel) GetModelId() *string {
	return s.ModelId
}

func (s *ApigLLMModel) GetModels() *string {
	return s.Models
}

func (s *ApigLLMModel) GetModelsWeight() *string {
	return s.ModelsWeight
}

func (s *ApigLLMModel) GetName() *string {
	return s.Name
}

func (s *ApigLLMModel) GetProvider() *string {
	return s.Provider
}

func (s *ApigLLMModel) GetTargetId() *string {
	return s.TargetId
}

func (s *ApigLLMModel) GetTenantId() *string {
	return s.TenantId
}

func (s *ApigLLMModel) GetType() *string {
	return s.Type
}

func (s *ApigLLMModel) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ApigLLMModel) SetAddress(v string) *ApigLLMModel {
	s.Address = &v
	return s
}

func (s *ApigLLMModel) SetApiKey(v string) *ApigLLMModel {
	s.ApiKey = &v
	return s
}

func (s *ApigLLMModel) SetCreatedTime(v string) *ApigLLMModel {
	s.CreatedTime = &v
	return s
}

func (s *ApigLLMModel) SetDesc(v string) *ApigLLMModel {
	s.Desc = &v
	return s
}

func (s *ApigLLMModel) SetGatewayId(v string) *ApigLLMModel {
	s.GatewayId = &v
	return s
}

func (s *ApigLLMModel) SetModelId(v string) *ApigLLMModel {
	s.ModelId = &v
	return s
}

func (s *ApigLLMModel) SetModels(v string) *ApigLLMModel {
	s.Models = &v
	return s
}

func (s *ApigLLMModel) SetModelsWeight(v string) *ApigLLMModel {
	s.ModelsWeight = &v
	return s
}

func (s *ApigLLMModel) SetName(v string) *ApigLLMModel {
	s.Name = &v
	return s
}

func (s *ApigLLMModel) SetProvider(v string) *ApigLLMModel {
	s.Provider = &v
	return s
}

func (s *ApigLLMModel) SetTargetId(v string) *ApigLLMModel {
	s.TargetId = &v
	return s
}

func (s *ApigLLMModel) SetTenantId(v string) *ApigLLMModel {
	s.TenantId = &v
	return s
}

func (s *ApigLLMModel) SetType(v string) *ApigLLMModel {
	s.Type = &v
	return s
}

func (s *ApigLLMModel) SetUpdateTime(v string) *ApigLLMModel {
	s.UpdateTime = &v
	return s
}

func (s *ApigLLMModel) Validate() error {
	return dara.Validate(s)
}
