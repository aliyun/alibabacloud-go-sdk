// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceTypesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceType(v string) *DescribeInstanceTypesRequest
	GetInstanceType() *string
}

type DescribeInstanceTypesRequest struct {
	// The instance type.
	//
	// example:
	//
	// ens.sn1.small
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s DescribeInstanceTypesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceTypesRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTypesRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeInstanceTypesRequest) SetInstanceType(v string) *DescribeInstanceTypesRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeInstanceTypesRequest) Validate() error {
	return dara.Validate(s)
}
