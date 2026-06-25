// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNamespaceVpcRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNameSpaceShortId(v string) *UpdateNamespaceVpcRequest
	GetNameSpaceShortId() *string
	SetNamespaceId(v string) *UpdateNamespaceVpcRequest
	GetNamespaceId() *string
	SetVpcId(v string) *UpdateNamespaceVpcRequest
	GetVpcId() *string
}

type UpdateNamespaceVpcRequest struct {
	// The short-format namespace ID. You do not need to include the region ID. We recommend that you use this parameter.
	//
	// example:
	//
	// test
	NameSpaceShortId *string `json:"NameSpaceShortId,omitempty" xml:"NameSpaceShortId,omitempty"`
	// The long-format namespace ID. If you set this parameter, it takes precedence over NameSpaceShortId. This parameter is for backward compatibility. Use the short-format namespace ID to simplify the request.
	//
	// example:
	//
	// cn-beijing:test
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The ID of the VPC. This VPC replaces the original VPC.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-2ze0i263cnn311nvj****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s UpdateNamespaceVpcRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateNamespaceVpcRequest) GoString() string {
	return s.String()
}

func (s *UpdateNamespaceVpcRequest) GetNameSpaceShortId() *string {
	return s.NameSpaceShortId
}

func (s *UpdateNamespaceVpcRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *UpdateNamespaceVpcRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *UpdateNamespaceVpcRequest) SetNameSpaceShortId(v string) *UpdateNamespaceVpcRequest {
	s.NameSpaceShortId = &v
	return s
}

func (s *UpdateNamespaceVpcRequest) SetNamespaceId(v string) *UpdateNamespaceVpcRequest {
	s.NamespaceId = &v
	return s
}

func (s *UpdateNamespaceVpcRequest) SetVpcId(v string) *UpdateNamespaceVpcRequest {
	s.VpcId = &v
	return s
}

func (s *UpdateNamespaceVpcRequest) Validate() error {
	return dara.Validate(s)
}
