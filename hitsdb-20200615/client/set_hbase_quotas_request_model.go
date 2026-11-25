// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetHBaseQuotasRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDimension(v string) *SetHBaseQuotasRequest
	GetDimension() *string
	SetInstanceId(v string) *SetHBaseQuotasRequest
	GetInstanceId() *string
	SetLimit(v string) *SetHBaseQuotasRequest
	GetLimit() *string
	SetRegionId(v string) *SetHBaseQuotasRequest
	GetRegionId() *string
	SetScope(v string) *SetHBaseQuotasRequest
	GetScope() *string
	SetSecurityToken(v string) *SetHBaseQuotasRequest
	GetSecurityToken() *string
	SetTimeUnit(v string) *SetHBaseQuotasRequest
	GetTimeUnit() *string
	SetType(v string) *SetHBaseQuotasRequest
	GetType() *string
	SetUnit(v string) *SetHBaseQuotasRequest
	GetUnit() *string
}

type SetHBaseQuotasRequest struct {
	Dimension *string `json:"Dimension,omitempty" xml:"Dimension,omitempty"`
	// This parameter is required.
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Limit         *string `json:"Limit,omitempty" xml:"Limit,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Scope         *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	TimeUnit      *string `json:"TimeUnit,omitempty" xml:"TimeUnit,omitempty"`
	Type          *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Unit          *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
}

func (s SetHBaseQuotasRequest) String() string {
	return dara.Prettify(s)
}

func (s SetHBaseQuotasRequest) GoString() string {
	return s.String()
}

func (s *SetHBaseQuotasRequest) GetDimension() *string {
	return s.Dimension
}

func (s *SetHBaseQuotasRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SetHBaseQuotasRequest) GetLimit() *string {
	return s.Limit
}

func (s *SetHBaseQuotasRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SetHBaseQuotasRequest) GetScope() *string {
	return s.Scope
}

func (s *SetHBaseQuotasRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *SetHBaseQuotasRequest) GetTimeUnit() *string {
	return s.TimeUnit
}

func (s *SetHBaseQuotasRequest) GetType() *string {
	return s.Type
}

func (s *SetHBaseQuotasRequest) GetUnit() *string {
	return s.Unit
}

func (s *SetHBaseQuotasRequest) SetDimension(v string) *SetHBaseQuotasRequest {
	s.Dimension = &v
	return s
}

func (s *SetHBaseQuotasRequest) SetInstanceId(v string) *SetHBaseQuotasRequest {
	s.InstanceId = &v
	return s
}

func (s *SetHBaseQuotasRequest) SetLimit(v string) *SetHBaseQuotasRequest {
	s.Limit = &v
	return s
}

func (s *SetHBaseQuotasRequest) SetRegionId(v string) *SetHBaseQuotasRequest {
	s.RegionId = &v
	return s
}

func (s *SetHBaseQuotasRequest) SetScope(v string) *SetHBaseQuotasRequest {
	s.Scope = &v
	return s
}

func (s *SetHBaseQuotasRequest) SetSecurityToken(v string) *SetHBaseQuotasRequest {
	s.SecurityToken = &v
	return s
}

func (s *SetHBaseQuotasRequest) SetTimeUnit(v string) *SetHBaseQuotasRequest {
	s.TimeUnit = &v
	return s
}

func (s *SetHBaseQuotasRequest) SetType(v string) *SetHBaseQuotasRequest {
	s.Type = &v
	return s
}

func (s *SetHBaseQuotasRequest) SetUnit(v string) *SetHBaseQuotasRequest {
	s.Unit = &v
	return s
}

func (s *SetHBaseQuotasRequest) Validate() error {
	return dara.Validate(s)
}
