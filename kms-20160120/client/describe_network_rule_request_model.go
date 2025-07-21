// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNetworkRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *DescribeNetworkRuleRequest
	GetName() *string
}

type DescribeNetworkRuleRequest struct {
	// The name of the access control rule that you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// networkrule_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeNetworkRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkRuleRequest) GoString() string {
	return s.String()
}

func (s *DescribeNetworkRuleRequest) GetName() *string {
	return s.Name
}

func (s *DescribeNetworkRuleRequest) SetName(v string) *DescribeNetworkRuleRequest {
	s.Name = &v
	return s
}

func (s *DescribeNetworkRuleRequest) Validate() error {
	return dara.Validate(s)
}
