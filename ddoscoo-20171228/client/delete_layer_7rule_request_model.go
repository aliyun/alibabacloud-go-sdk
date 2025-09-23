// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLayer7RuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *DeleteLayer7RuleRequest
	GetDomain() *string
	SetResourceGroupId(v string) *DeleteLayer7RuleRequest
	GetResourceGroupId() *string
}

type DeleteLayer7RuleRequest struct {
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
}

func (s DeleteLayer7RuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLayer7RuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteLayer7RuleRequest) GetDomain() *string {
	return s.Domain
}

func (s *DeleteLayer7RuleRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DeleteLayer7RuleRequest) SetDomain(v string) *DeleteLayer7RuleRequest {
	s.Domain = &v
	return s
}

func (s *DeleteLayer7RuleRequest) SetResourceGroupId(v string) *DeleteLayer7RuleRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DeleteLayer7RuleRequest) Validate() error {
	return dara.Validate(s)
}
