// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRateLimitPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGwClusterId(v string) *CreateRateLimitPolicyRequest
	GetGwClusterId() *string
	SetRateLimitRpm(v string) *CreateRateLimitPolicyRequest
	GetRateLimitRpm() *string
	SetRateLimitTpm(v string) *CreateRateLimitPolicyRequest
	GetRateLimitTpm() *string
	SetRegionId(v string) *CreateRateLimitPolicyRequest
	GetRegionId() *string
	SetScopeRefId(v string) *CreateRateLimitPolicyRequest
	GetScopeRefId() *string
	SetScopeType(v string) *CreateRateLimitPolicyRequest
	GetScopeType() *string
}

type CreateRateLimitPolicyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pg-xxxxxxx
	GwClusterId *string `json:"GwClusterId,omitempty" xml:"GwClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	RateLimitRpm *string `json:"RateLimitRpm,omitempty" xml:"RateLimitRpm,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	RateLimitTpm *string `json:"RateLimitTpm,omitempty" xml:"RateLimitTpm,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cg-xxxxxxx
	ScopeRefId *string `json:"ScopeRefId,omitempty" xml:"ScopeRefId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ConsumerGroup
	ScopeType *string `json:"ScopeType,omitempty" xml:"ScopeType,omitempty"`
}

func (s CreateRateLimitPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRateLimitPolicyRequest) GoString() string {
	return s.String()
}

func (s *CreateRateLimitPolicyRequest) GetGwClusterId() *string {
	return s.GwClusterId
}

func (s *CreateRateLimitPolicyRequest) GetRateLimitRpm() *string {
	return s.RateLimitRpm
}

func (s *CreateRateLimitPolicyRequest) GetRateLimitTpm() *string {
	return s.RateLimitTpm
}

func (s *CreateRateLimitPolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateRateLimitPolicyRequest) GetScopeRefId() *string {
	return s.ScopeRefId
}

func (s *CreateRateLimitPolicyRequest) GetScopeType() *string {
	return s.ScopeType
}

func (s *CreateRateLimitPolicyRequest) SetGwClusterId(v string) *CreateRateLimitPolicyRequest {
	s.GwClusterId = &v
	return s
}

func (s *CreateRateLimitPolicyRequest) SetRateLimitRpm(v string) *CreateRateLimitPolicyRequest {
	s.RateLimitRpm = &v
	return s
}

func (s *CreateRateLimitPolicyRequest) SetRateLimitTpm(v string) *CreateRateLimitPolicyRequest {
	s.RateLimitTpm = &v
	return s
}

func (s *CreateRateLimitPolicyRequest) SetRegionId(v string) *CreateRateLimitPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *CreateRateLimitPolicyRequest) SetScopeRefId(v string) *CreateRateLimitPolicyRequest {
	s.ScopeRefId = &v
	return s
}

func (s *CreateRateLimitPolicyRequest) SetScopeType(v string) *CreateRateLimitPolicyRequest {
	s.ScopeType = &v
	return s
}

func (s *CreateRateLimitPolicyRequest) Validate() error {
	return dara.Validate(s)
}
