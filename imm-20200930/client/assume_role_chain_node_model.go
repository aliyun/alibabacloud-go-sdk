// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssumeRoleChainNode interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v string) *AssumeRoleChainNode
	GetOwnerId() *string
	SetRole(v string) *AssumeRoleChainNode
	GetRole() *string
	SetType(v string) *AssumeRoleChainNode
	GetType() *string
}

type AssumeRoleChainNode struct {
	// This parameter is required.
	//
	// example:
	//
	// 1023210024677934
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test-role
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// user
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s AssumeRoleChainNode) String() string {
	return dara.Prettify(s)
}

func (s AssumeRoleChainNode) GoString() string {
	return s.String()
}

func (s *AssumeRoleChainNode) GetOwnerId() *string {
	return s.OwnerId
}

func (s *AssumeRoleChainNode) GetRole() *string {
	return s.Role
}

func (s *AssumeRoleChainNode) GetType() *string {
	return s.Type
}

func (s *AssumeRoleChainNode) SetOwnerId(v string) *AssumeRoleChainNode {
	s.OwnerId = &v
	return s
}

func (s *AssumeRoleChainNode) SetRole(v string) *AssumeRoleChainNode {
	s.Role = &v
	return s
}

func (s *AssumeRoleChainNode) SetType(v string) *AssumeRoleChainNode {
	s.Type = &v
	return s
}

func (s *AssumeRoleChainNode) Validate() error {
	return dara.Validate(s)
}
