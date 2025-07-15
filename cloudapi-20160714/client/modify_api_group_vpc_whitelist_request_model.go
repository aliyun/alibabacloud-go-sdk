// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApiGroupVpcWhitelistRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *ModifyApiGroupVpcWhitelistRequest
	GetGroupId() *string
	SetSecurityToken(v string) *ModifyApiGroupVpcWhitelistRequest
	GetSecurityToken() *string
	SetVpcIds(v string) *ModifyApiGroupVpcWhitelistRequest
	GetVpcIds() *string
}

type ModifyApiGroupVpcWhitelistRequest struct {
	// The ID of the API group.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9b80408147724ddab4c4e2703c6ca019
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The ID of the VPC instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp11w979o2s9rcr962w25
	VpcIds *string `json:"VpcIds,omitempty" xml:"VpcIds,omitempty"`
}

func (s ModifyApiGroupVpcWhitelistRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyApiGroupVpcWhitelistRequest) GoString() string {
	return s.String()
}

func (s *ModifyApiGroupVpcWhitelistRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *ModifyApiGroupVpcWhitelistRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ModifyApiGroupVpcWhitelistRequest) GetVpcIds() *string {
	return s.VpcIds
}

func (s *ModifyApiGroupVpcWhitelistRequest) SetGroupId(v string) *ModifyApiGroupVpcWhitelistRequest {
	s.GroupId = &v
	return s
}

func (s *ModifyApiGroupVpcWhitelistRequest) SetSecurityToken(v string) *ModifyApiGroupVpcWhitelistRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyApiGroupVpcWhitelistRequest) SetVpcIds(v string) *ModifyApiGroupVpcWhitelistRequest {
	s.VpcIds = &v
	return s
}

func (s *ModifyApiGroupVpcWhitelistRequest) Validate() error {
	return dara.Validate(s)
}
