// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyIpControlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ModifyIpControlRequest
	GetDescription() *string
	SetIpControlId(v string) *ModifyIpControlRequest
	GetIpControlId() *string
	SetIpControlName(v string) *ModifyIpControlRequest
	GetIpControlName() *string
	SetSecurityToken(v string) *ModifyIpControlRequest
	GetSecurityToken() *string
}

type ModifyIpControlRequest struct {
	// The description. The description can be up to 200 characters in length.
	//
	// example:
	//
	// description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the ACL. The ID is unique.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7ea91319a34d48a09b5c9c871d9768b1
	IpControlId *string `json:"IpControlId,omitempty" xml:"IpControlId,omitempty"`
	// The name of the ACL. The name must be 4 to 50 characters in length, and can contain letters, digits, and underscores (_). The name cannot start with an underscore (_).
	//
	// example:
	//
	// testControl11
	IpControlName *string `json:"IpControlName,omitempty" xml:"IpControlName,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ModifyIpControlRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyIpControlRequest) GoString() string {
	return s.String()
}

func (s *ModifyIpControlRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyIpControlRequest) GetIpControlId() *string {
	return s.IpControlId
}

func (s *ModifyIpControlRequest) GetIpControlName() *string {
	return s.IpControlName
}

func (s *ModifyIpControlRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ModifyIpControlRequest) SetDescription(v string) *ModifyIpControlRequest {
	s.Description = &v
	return s
}

func (s *ModifyIpControlRequest) SetIpControlId(v string) *ModifyIpControlRequest {
	s.IpControlId = &v
	return s
}

func (s *ModifyIpControlRequest) SetIpControlName(v string) *ModifyIpControlRequest {
	s.IpControlName = &v
	return s
}

func (s *ModifyIpControlRequest) SetSecurityToken(v string) *ModifyIpControlRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyIpControlRequest) Validate() error {
	return dara.Validate(s)
}
