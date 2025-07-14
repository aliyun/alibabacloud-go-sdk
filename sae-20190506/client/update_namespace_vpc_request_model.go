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
	// example:
	//
	// test
	NameSpaceShortId *string `json:"NameSpaceShortId,omitempty" xml:"NameSpaceShortId,omitempty"`
	// vpc-2ze0i263cnn311nvj\\*\\*\\*\\*
	//
	// example:
	//
	// cn-beijing:test
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The ID of the request.
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
