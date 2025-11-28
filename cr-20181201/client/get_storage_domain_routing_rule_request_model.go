// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStorageDomainRoutingRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetStorageDomainRoutingRuleRequest
	GetInstanceId() *string
	SetRuleId(v string) *GetStorageDomainRoutingRuleRequest
	GetRuleId() *string
}

type GetStorageDomainRoutingRuleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cri-kmsiwlxxdcva****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// crasr-mdbpung4i1rm****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s GetStorageDomainRoutingRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s GetStorageDomainRoutingRuleRequest) GoString() string {
	return s.String()
}

func (s *GetStorageDomainRoutingRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetStorageDomainRoutingRuleRequest) GetRuleId() *string {
	return s.RuleId
}

func (s *GetStorageDomainRoutingRuleRequest) SetInstanceId(v string) *GetStorageDomainRoutingRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *GetStorageDomainRoutingRuleRequest) SetRuleId(v string) *GetStorageDomainRoutingRuleRequest {
	s.RuleId = &v
	return s
}

func (s *GetStorageDomainRoutingRuleRequest) Validate() error {
	return dara.Validate(s)
}
