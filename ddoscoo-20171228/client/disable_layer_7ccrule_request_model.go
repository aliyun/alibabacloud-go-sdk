// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableLayer7CCRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *DisableLayer7CCRuleRequest
	GetDomain() *string
	SetResourceGroupId(v string) *DisableLayer7CCRuleRequest
	GetResourceGroupId() *string
	SetSourceIp(v string) *DisableLayer7CCRuleRequest
	GetSourceIp() *string
}

type DisableLayer7CCRuleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// www.aliyun.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// test
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 1.1.1.1
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DisableLayer7CCRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableLayer7CCRuleRequest) GoString() string {
	return s.String()
}

func (s *DisableLayer7CCRuleRequest) GetDomain() *string {
	return s.Domain
}

func (s *DisableLayer7CCRuleRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DisableLayer7CCRuleRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DisableLayer7CCRuleRequest) SetDomain(v string) *DisableLayer7CCRuleRequest {
	s.Domain = &v
	return s
}

func (s *DisableLayer7CCRuleRequest) SetResourceGroupId(v string) *DisableLayer7CCRuleRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DisableLayer7CCRuleRequest) SetSourceIp(v string) *DisableLayer7CCRuleRequest {
	s.SourceIp = &v
	return s
}

func (s *DisableLayer7CCRuleRequest) Validate() error {
	return dara.Validate(s)
}
