// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNetworkRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *DeleteNetworkRuleRequest
	GetName() *string
}

type DeleteNetworkRuleRequest struct {
	// The name of the network access rule that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// networkrule_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DeleteNetworkRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteNetworkRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteNetworkRuleRequest) GetName() *string {
	return s.Name
}

func (s *DeleteNetworkRuleRequest) SetName(v string) *DeleteNetworkRuleRequest {
	s.Name = &v
	return s
}

func (s *DeleteNetworkRuleRequest) Validate() error {
	return dara.Validate(s)
}
