// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRateLimitPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGwClusterId(v string) *ModifyRateLimitPolicyRequest
	GetGwClusterId() *string
	SetPolicyId(v string) *ModifyRateLimitPolicyRequest
	GetPolicyId() *string
	SetRateLimitRpm(v string) *ModifyRateLimitPolicyRequest
	GetRateLimitRpm() *string
	SetRateLimitTpm(v string) *ModifyRateLimitPolicyRequest
	GetRateLimitTpm() *string
	SetRegionId(v string) *ModifyRateLimitPolicyRequest
	GetRegionId() *string
}

type ModifyRateLimitPolicyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pg-xxxxxxxx
	GwClusterId *string `json:"GwClusterId,omitempty" xml:"GwClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 02eccf7c61cf4d05a543075ee907f3**
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// example:
	//
	// 10
	RateLimitRpm *string `json:"RateLimitRpm,omitempty" xml:"RateLimitRpm,omitempty"`
	// example:
	//
	// 10
	RateLimitTpm *string `json:"RateLimitTpm,omitempty" xml:"RateLimitTpm,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyRateLimitPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyRateLimitPolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifyRateLimitPolicyRequest) GetGwClusterId() *string {
	return s.GwClusterId
}

func (s *ModifyRateLimitPolicyRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *ModifyRateLimitPolicyRequest) GetRateLimitRpm() *string {
	return s.RateLimitRpm
}

func (s *ModifyRateLimitPolicyRequest) GetRateLimitTpm() *string {
	return s.RateLimitTpm
}

func (s *ModifyRateLimitPolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyRateLimitPolicyRequest) SetGwClusterId(v string) *ModifyRateLimitPolicyRequest {
	s.GwClusterId = &v
	return s
}

func (s *ModifyRateLimitPolicyRequest) SetPolicyId(v string) *ModifyRateLimitPolicyRequest {
	s.PolicyId = &v
	return s
}

func (s *ModifyRateLimitPolicyRequest) SetRateLimitRpm(v string) *ModifyRateLimitPolicyRequest {
	s.RateLimitRpm = &v
	return s
}

func (s *ModifyRateLimitPolicyRequest) SetRateLimitTpm(v string) *ModifyRateLimitPolicyRequest {
	s.RateLimitTpm = &v
	return s
}

func (s *ModifyRateLimitPolicyRequest) SetRegionId(v string) *ModifyRateLimitPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyRateLimitPolicyRequest) Validate() error {
	return dara.Validate(s)
}
