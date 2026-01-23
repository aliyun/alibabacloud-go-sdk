// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStorageDomainRoutingRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteStorageDomainRoutingRuleRequest
	GetInstanceId() *string
	SetRuleId(v string) *DeleteStorageDomainRoutingRuleRequest
	GetRuleId() *string
}

type DeleteStorageDomainRoutingRuleRequest struct {
	// The instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-4cdrlqmhn4gm****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The rule ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// crsdr-n6pbhgjxtla***
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DeleteStorageDomainRoutingRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteStorageDomainRoutingRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteStorageDomainRoutingRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteStorageDomainRoutingRuleRequest) GetRuleId() *string {
	return s.RuleId
}

func (s *DeleteStorageDomainRoutingRuleRequest) SetInstanceId(v string) *DeleteStorageDomainRoutingRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteStorageDomainRoutingRuleRequest) SetRuleId(v string) *DeleteStorageDomainRoutingRuleRequest {
	s.RuleId = &v
	return s
}

func (s *DeleteStorageDomainRoutingRuleRequest) Validate() error {
	return dara.Validate(s)
}
