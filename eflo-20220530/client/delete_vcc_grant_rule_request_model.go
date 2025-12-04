// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVccGrantRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetErId(v string) *DeleteVccGrantRuleRequest
	GetErId() *string
	SetGrantRuleId(v string) *DeleteVccGrantRuleRequest
	GetGrantRuleId() *string
	SetInstanceId(v string) *DeleteVccGrantRuleRequest
	GetInstanceId() *string
	SetRegionId(v string) *DeleteVccGrantRuleRequest
	GetRegionId() *string
}

type DeleteVccGrantRuleRequest struct {
	// Lingjun HUB ID
	//
	// example:
	//
	// er-kkopgtne
	ErId *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	// Authorization Entry ID
	//
	// This parameter is required.
	//
	// example:
	//
	// grant-rule-jaj34d75h01
	GrantRuleId *string `json:"GrantRuleId,omitempty" xml:"GrantRuleId,omitempty"`
	// Network Instance ID
	//
	// example:
	//
	// vcc-cn-jaj34d75h01
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteVccGrantRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteVccGrantRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteVccGrantRuleRequest) GetErId() *string {
	return s.ErId
}

func (s *DeleteVccGrantRuleRequest) GetGrantRuleId() *string {
	return s.GrantRuleId
}

func (s *DeleteVccGrantRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteVccGrantRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteVccGrantRuleRequest) SetErId(v string) *DeleteVccGrantRuleRequest {
	s.ErId = &v
	return s
}

func (s *DeleteVccGrantRuleRequest) SetGrantRuleId(v string) *DeleteVccGrantRuleRequest {
	s.GrantRuleId = &v
	return s
}

func (s *DeleteVccGrantRuleRequest) SetInstanceId(v string) *DeleteVccGrantRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteVccGrantRuleRequest) SetRegionId(v string) *DeleteVccGrantRuleRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteVccGrantRuleRequest) Validate() error {
	return dara.Validate(s)
}
