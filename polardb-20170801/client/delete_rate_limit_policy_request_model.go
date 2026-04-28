// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRateLimitPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGwClusterId(v string) *DeleteRateLimitPolicyRequest
	GetGwClusterId() *string
	SetPolicyId(v string) *DeleteRateLimitPolicyRequest
	GetPolicyId() *string
	SetRegionId(v string) *DeleteRateLimitPolicyRequest
	GetRegionId() *string
}

type DeleteRateLimitPolicyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pg-xxxxxxxxxx
	GwClusterId *string `json:"GwClusterId,omitempty" xml:"GwClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 02eccf7c61cf4d05a543075ee907f3**
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteRateLimitPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRateLimitPolicyRequest) GoString() string {
	return s.String()
}

func (s *DeleteRateLimitPolicyRequest) GetGwClusterId() *string {
	return s.GwClusterId
}

func (s *DeleteRateLimitPolicyRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *DeleteRateLimitPolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteRateLimitPolicyRequest) SetGwClusterId(v string) *DeleteRateLimitPolicyRequest {
	s.GwClusterId = &v
	return s
}

func (s *DeleteRateLimitPolicyRequest) SetPolicyId(v string) *DeleteRateLimitPolicyRequest {
	s.PolicyId = &v
	return s
}

func (s *DeleteRateLimitPolicyRequest) SetRegionId(v string) *DeleteRateLimitPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteRateLimitPolicyRequest) Validate() error {
	return dara.Validate(s)
}
