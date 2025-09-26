// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModel interface {
	dara.Model
	String() string
	GoString() string
	SetAddress(v string) *Model
	GetAddress() *string
	SetApiKey(v string) *Model
	GetApiKey() *string
	SetCreatedTime(v string) *Model
	GetCreatedTime() *string
	SetDesc(v string) *Model
	GetDesc() *string
	SetGatewayId(v string) *Model
	GetGatewayId() *string
	SetModelId(v string) *Model
	GetModelId() *string
	SetModels(v string) *Model
	GetModels() *string
	SetModelsWeight(v string) *Model
	GetModelsWeight() *string
	SetName(v string) *Model
	GetName() *string
	SetProvider(v string) *Model
	GetProvider() *string
	SetTargetId(v string) *Model
	GetTargetId() *string
	SetTenantId(v string) *Model
	GetTenantId() *string
	SetType(v string) *Model
	GetType() *string
	SetUpdateTime(v string) *Model
	GetUpdateTime() *string
}

type Model struct {
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

func (s Model) String() string {
	return dara.Prettify(s)
}

func (s Model) GoString() string {
	return s.String()
}

func (s *Model) GetAddress() *string {
	return s.Address
}

func (s *Model) GetApiKey() *string {
	return s.ApiKey
}

func (s *Model) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *Model) GetDesc() *string {
	return s.Desc
}

func (s *Model) GetGatewayId() *string {
	return s.GatewayId
}

func (s *Model) GetModelId() *string {
	return s.ModelId
}

func (s *Model) GetModels() *string {
	return s.Models
}

func (s *Model) GetModelsWeight() *string {
	return s.ModelsWeight
}

func (s *Model) GetName() *string {
	return s.Name
}

func (s *Model) GetProvider() *string {
	return s.Provider
}

func (s *Model) GetTargetId() *string {
	return s.TargetId
}

func (s *Model) GetTenantId() *string {
	return s.TenantId
}

func (s *Model) GetType() *string {
	return s.Type
}

func (s *Model) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *Model) SetAddress(v string) *Model {
	s.Address = &v
	return s
}

func (s *Model) SetApiKey(v string) *Model {
	s.ApiKey = &v
	return s
}

func (s *Model) SetCreatedTime(v string) *Model {
	s.CreatedTime = &v
	return s
}

func (s *Model) SetDesc(v string) *Model {
	s.Desc = &v
	return s
}

func (s *Model) SetGatewayId(v string) *Model {
	s.GatewayId = &v
	return s
}

func (s *Model) SetModelId(v string) *Model {
	s.ModelId = &v
	return s
}

func (s *Model) SetModels(v string) *Model {
	s.Models = &v
	return s
}

func (s *Model) SetModelsWeight(v string) *Model {
	s.ModelsWeight = &v
	return s
}

func (s *Model) SetName(v string) *Model {
	s.Name = &v
	return s
}

func (s *Model) SetProvider(v string) *Model {
	s.Provider = &v
	return s
}

func (s *Model) SetTargetId(v string) *Model {
	s.TargetId = &v
	return s
}

func (s *Model) SetTenantId(v string) *Model {
	s.TenantId = &v
	return s
}

func (s *Model) SetType(v string) *Model {
	s.Type = &v
	return s
}

func (s *Model) SetUpdateTime(v string) *Model {
	s.UpdateTime = &v
	return s
}

func (s *Model) Validate() error {
	return dara.Validate(s)
}
