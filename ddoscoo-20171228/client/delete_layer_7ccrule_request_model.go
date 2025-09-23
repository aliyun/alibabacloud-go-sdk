// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLayer7CCRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *DeleteLayer7CCRuleRequest
	GetDomain() *string
	SetName(v string) *DeleteLayer7CCRuleRequest
	GetName() *string
	SetResourceGroupId(v string) *DeleteLayer7CCRuleRequest
	GetResourceGroupId() *string
}

type DeleteLayer7CCRuleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// www.aliyun.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testCcRule1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// test
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DeleteLayer7CCRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLayer7CCRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteLayer7CCRuleRequest) GetDomain() *string {
	return s.Domain
}

func (s *DeleteLayer7CCRuleRequest) GetName() *string {
	return s.Name
}

func (s *DeleteLayer7CCRuleRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DeleteLayer7CCRuleRequest) SetDomain(v string) *DeleteLayer7CCRuleRequest {
	s.Domain = &v
	return s
}

func (s *DeleteLayer7CCRuleRequest) SetName(v string) *DeleteLayer7CCRuleRequest {
	s.Name = &v
	return s
}

func (s *DeleteLayer7CCRuleRequest) SetResourceGroupId(v string) *DeleteLayer7CCRuleRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DeleteLayer7CCRuleRequest) Validate() error {
	return dara.Validate(s)
}
