// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindInstance2VpcRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceName(v string) *UnbindInstance2VpcRequest
	GetInstanceName() *string
	SetInstanceVpcName(v string) *UnbindInstance2VpcRequest
	GetInstanceVpcName() *string
}

type UnbindInstance2VpcRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// mkdd-test
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xu6666
	InstanceVpcName *string `json:"InstanceVpcName,omitempty" xml:"InstanceVpcName,omitempty"`
}

func (s UnbindInstance2VpcRequest) String() string {
	return dara.Prettify(s)
}

func (s UnbindInstance2VpcRequest) GoString() string {
	return s.String()
}

func (s *UnbindInstance2VpcRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *UnbindInstance2VpcRequest) GetInstanceVpcName() *string {
	return s.InstanceVpcName
}

func (s *UnbindInstance2VpcRequest) SetInstanceName(v string) *UnbindInstance2VpcRequest {
	s.InstanceName = &v
	return s
}

func (s *UnbindInstance2VpcRequest) SetInstanceVpcName(v string) *UnbindInstance2VpcRequest {
	s.InstanceVpcName = &v
	return s
}

func (s *UnbindInstance2VpcRequest) Validate() error {
	return dara.Validate(s)
}
