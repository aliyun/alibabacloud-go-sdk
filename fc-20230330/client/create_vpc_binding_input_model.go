// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVpcBindingInput interface {
	dara.Model
	String() string
	GoString() string
	SetVpcId(v string) *CreateVpcBindingInput
	GetVpcId() *string
}

type CreateVpcBindingInput struct {
	// This parameter is required.
	//
	// example:
	//
	// vpc-8vb8x8dggvr0axxxxxxxx
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s CreateVpcBindingInput) String() string {
	return dara.Prettify(s)
}

func (s CreateVpcBindingInput) GoString() string {
	return s.String()
}

func (s *CreateVpcBindingInput) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateVpcBindingInput) SetVpcId(v string) *CreateVpcBindingInput {
	s.VpcId = &v
	return s
}

func (s *CreateVpcBindingInput) Validate() error {
	return dara.Validate(s)
}
