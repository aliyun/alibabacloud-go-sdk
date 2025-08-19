// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVpcBindingsOutput interface {
	dara.Model
	String() string
	GoString() string
	SetVpcIds(v []*string) *ListVpcBindingsOutput
	GetVpcIds() []*string
}

type ListVpcBindingsOutput struct {
	VpcIds []*string `json:"vpcIds" xml:"vpcIds" type:"Repeated"`
}

func (s ListVpcBindingsOutput) String() string {
	return dara.Prettify(s)
}

func (s ListVpcBindingsOutput) GoString() string {
	return s.String()
}

func (s *ListVpcBindingsOutput) GetVpcIds() []*string {
	return s.VpcIds
}

func (s *ListVpcBindingsOutput) SetVpcIds(v []*string) *ListVpcBindingsOutput {
	s.VpcIds = v
	return s
}

func (s *ListVpcBindingsOutput) Validate() error {
	return dara.Validate(s)
}
