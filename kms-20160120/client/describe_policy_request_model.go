// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *DescribePolicyRequest
	GetName() *string
}

type DescribePolicyRequest struct {
	// The name of the permission policy that you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// policy_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribePolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePolicyRequest) GoString() string {
	return s.String()
}

func (s *DescribePolicyRequest) GetName() *string {
	return s.Name
}

func (s *DescribePolicyRequest) SetName(v string) *DescribePolicyRequest {
	s.Name = &v
	return s
}

func (s *DescribePolicyRequest) Validate() error {
	return dara.Validate(s)
}
