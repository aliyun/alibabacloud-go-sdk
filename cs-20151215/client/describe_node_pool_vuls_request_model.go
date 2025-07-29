// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNodePoolVulsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNecessity(v string) *DescribeNodePoolVulsRequest
	GetNecessity() *string
}

type DescribeNodePoolVulsRequest struct {
	// The priority to fix the vulnerability. Separate multiple priorities with commas (,). Valid values:
	//
	// 	- `asap`: high
	//
	// 	- `later`: medium
	//
	// 	- `nntf`: low
	//
	// example:
	//
	// asap
	Necessity *string `json:"necessity,omitempty" xml:"necessity,omitempty"`
}

func (s DescribeNodePoolVulsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNodePoolVulsRequest) GoString() string {
	return s.String()
}

func (s *DescribeNodePoolVulsRequest) GetNecessity() *string {
	return s.Necessity
}

func (s *DescribeNodePoolVulsRequest) SetNecessity(v string) *DescribeNodePoolVulsRequest {
	s.Necessity = &v
	return s
}

func (s *DescribeNodePoolVulsRequest) Validate() error {
	return dara.Validate(s)
}
