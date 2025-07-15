// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceVpcAttributeForConsoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeleteVpcAccess(v bool) *ModifyInstanceVpcAttributeForConsoleRequest
	GetDeleteVpcAccess() *bool
	SetInstanceId(v string) *ModifyInstanceVpcAttributeForConsoleRequest
	GetInstanceId() *string
	SetToken(v string) *ModifyInstanceVpcAttributeForConsoleRequest
	GetToken() *string
	SetVpcId(v string) *ModifyInstanceVpcAttributeForConsoleRequest
	GetVpcId() *string
	SetVpcOwnerId(v int64) *ModifyInstanceVpcAttributeForConsoleRequest
	GetVpcOwnerId() *int64
	SetVswitchId(v string) *ModifyInstanceVpcAttributeForConsoleRequest
	GetVswitchId() *string
}

type ModifyInstanceVpcAttributeForConsoleRequest struct {
	// Whether delete instance client VPC.
	//
	// - FALSE: set or modify instance client VPC
	//
	// - TRUE: delete instance client VPC
	//
	// example:
	//
	// false
	DeleteVpcAccess *bool `json:"DeleteVpcAccess,omitempty" xml:"DeleteVpcAccess,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// apigateway-bj-f28baxxxxb51
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The token of the request.
	//
	// example:
	//
	// 505959c38776d9324945dbff709582
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
	// The ID of the VPC.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-8vbnnd66xxxx2xb5oig4f
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of the Alibaba Cloud account to which the VPC belongs.
	//
	// example:
	//
	// 1121011712128923
	VpcOwnerId *int64 `json:"VpcOwnerId,omitempty" xml:"VpcOwnerId,omitempty"`
	// The ID of the vSwitch.
	//
	// example:
	//
	// vsw-bj9e2i8w3wz7shkvnuj9a
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
}

func (s ModifyInstanceVpcAttributeForConsoleRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceVpcAttributeForConsoleRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceVpcAttributeForConsoleRequest) GetDeleteVpcAccess() *bool {
	return s.DeleteVpcAccess
}

func (s *ModifyInstanceVpcAttributeForConsoleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyInstanceVpcAttributeForConsoleRequest) GetToken() *string {
	return s.Token
}

func (s *ModifyInstanceVpcAttributeForConsoleRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *ModifyInstanceVpcAttributeForConsoleRequest) GetVpcOwnerId() *int64 {
	return s.VpcOwnerId
}

func (s *ModifyInstanceVpcAttributeForConsoleRequest) GetVswitchId() *string {
	return s.VswitchId
}

func (s *ModifyInstanceVpcAttributeForConsoleRequest) SetDeleteVpcAccess(v bool) *ModifyInstanceVpcAttributeForConsoleRequest {
	s.DeleteVpcAccess = &v
	return s
}

func (s *ModifyInstanceVpcAttributeForConsoleRequest) SetInstanceId(v string) *ModifyInstanceVpcAttributeForConsoleRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceVpcAttributeForConsoleRequest) SetToken(v string) *ModifyInstanceVpcAttributeForConsoleRequest {
	s.Token = &v
	return s
}

func (s *ModifyInstanceVpcAttributeForConsoleRequest) SetVpcId(v string) *ModifyInstanceVpcAttributeForConsoleRequest {
	s.VpcId = &v
	return s
}

func (s *ModifyInstanceVpcAttributeForConsoleRequest) SetVpcOwnerId(v int64) *ModifyInstanceVpcAttributeForConsoleRequest {
	s.VpcOwnerId = &v
	return s
}

func (s *ModifyInstanceVpcAttributeForConsoleRequest) SetVswitchId(v string) *ModifyInstanceVpcAttributeForConsoleRequest {
	s.VswitchId = &v
	return s
}

func (s *ModifyInstanceVpcAttributeForConsoleRequest) Validate() error {
	return dara.Validate(s)
}
