// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyIpControlPolicyItemRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ModifyIpControlPolicyItemRequest
	GetAppId() *string
	SetCidrIp(v string) *ModifyIpControlPolicyItemRequest
	GetCidrIp() *string
	SetIpControlId(v string) *ModifyIpControlPolicyItemRequest
	GetIpControlId() *string
	SetPolicyItemId(v string) *ModifyIpControlPolicyItemRequest
	GetPolicyItemId() *string
	SetSecurityToken(v string) *ModifyIpControlPolicyItemRequest
	GetSecurityToken() *string
}

type ModifyIpControlPolicyItemRequest struct {
	// The ID of the application that is restricted by the policy. You can configure the AppId parameter only when the value of the IpControlType parameter is ALLOW.
	//
	// 	- You can add only one application ID at a time.
	//
	// 	- If this parameter is empty, no applications are restricted.
	//
	// 	- If this parameter is not empty, not only IP addresses but also applications are restricted.
	//
	// 	- If this parameter is not empty and no security authentication method is specified for the API, all API calls are restricted.
	//
	// 	- If the value of the IpControlType parameter is REFUSE and the AppId parameter is not empty, API Gateway automatically ignores the AppId parameter and restricts only the IP addresses.
	//
	// example:
	//
	// 123
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The IP address or CIDR block that is defined in a policy. Separate multiple IP addresses or CIDR blocks with semicolons (;). You can add a maximum of 10 IP addresses or CIDR blocks.
	//
	// This parameter is required.
	//
	// example:
	//
	// 113.125.1.101;101.11.1.1
	CidrIp *string `json:"CidrIp,omitempty" xml:"CidrIp,omitempty"`
	// The ID of the ACL. The ID is unique.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7ea91319a34d48a09b5c9c871d9768b1
	IpControlId *string `json:"IpControlId,omitempty" xml:"IpControlId,omitempty"`
	// The ID of the policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// P151617000829241
	PolicyItemId  *string `json:"PolicyItemId,omitempty" xml:"PolicyItemId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ModifyIpControlPolicyItemRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyIpControlPolicyItemRequest) GoString() string {
	return s.String()
}

func (s *ModifyIpControlPolicyItemRequest) GetAppId() *string {
	return s.AppId
}

func (s *ModifyIpControlPolicyItemRequest) GetCidrIp() *string {
	return s.CidrIp
}

func (s *ModifyIpControlPolicyItemRequest) GetIpControlId() *string {
	return s.IpControlId
}

func (s *ModifyIpControlPolicyItemRequest) GetPolicyItemId() *string {
	return s.PolicyItemId
}

func (s *ModifyIpControlPolicyItemRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ModifyIpControlPolicyItemRequest) SetAppId(v string) *ModifyIpControlPolicyItemRequest {
	s.AppId = &v
	return s
}

func (s *ModifyIpControlPolicyItemRequest) SetCidrIp(v string) *ModifyIpControlPolicyItemRequest {
	s.CidrIp = &v
	return s
}

func (s *ModifyIpControlPolicyItemRequest) SetIpControlId(v string) *ModifyIpControlPolicyItemRequest {
	s.IpControlId = &v
	return s
}

func (s *ModifyIpControlPolicyItemRequest) SetPolicyItemId(v string) *ModifyIpControlPolicyItemRequest {
	s.PolicyItemId = &v
	return s
}

func (s *ModifyIpControlPolicyItemRequest) SetSecurityToken(v string) *ModifyIpControlPolicyItemRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyIpControlPolicyItemRequest) Validate() error {
	return dara.Validate(s)
}
