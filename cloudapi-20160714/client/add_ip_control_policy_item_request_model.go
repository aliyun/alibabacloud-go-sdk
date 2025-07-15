// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddIpControlPolicyItemRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *AddIpControlPolicyItemRequest
	GetAppId() *string
	SetCidrIp(v string) *AddIpControlPolicyItemRequest
	GetCidrIp() *string
	SetIpControlId(v string) *AddIpControlPolicyItemRequest
	GetIpControlId() *string
	SetSecurityToken(v string) *AddIpControlPolicyItemRequest
	GetSecurityToken() *string
}

type AddIpControlPolicyItemRequest struct {
	// The restriction policy on app IDs for a specific policy. You can restrict app IDs only for whitelists. The IpControlType values of whitelists are ALLOW.
	//
	// 	- You can add only one app ID restriction policy at a time.
	//
	// 	- If this parameter is empty, no restriction is imposed on the app IDs.
	//
	// 	- If this parameter is not empty, there is restriction not only on IP addresses, but also on apps.
	//
	// 	- Please note that if this parameter is not empty and the security authentication method of the API is No Authentication, all API calls are restricted.
	//
	// 	- If this parameter is not empty for a blacklist, API Gateway automatically skips this parameter and sets only restriction on IP addresses. The IpControlType value of a blacklist is REFUSE.
	//
	// example:
	//
	// 1111111
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The IP addresses or CIDR blocks involved in the policy. Separate multiple IP addresses or CIDR blocks with semicolons (;). You can specify a maximum of 10 IP addresses or CIDR blocks.
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
	// The security token included in the WebSocket request header. The system uses this token to authenticate the request.
	//
	// example:
	//
	// 4223a10e-eed3-46a6-8b7c-23003f488153
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s AddIpControlPolicyItemRequest) String() string {
	return dara.Prettify(s)
}

func (s AddIpControlPolicyItemRequest) GoString() string {
	return s.String()
}

func (s *AddIpControlPolicyItemRequest) GetAppId() *string {
	return s.AppId
}

func (s *AddIpControlPolicyItemRequest) GetCidrIp() *string {
	return s.CidrIp
}

func (s *AddIpControlPolicyItemRequest) GetIpControlId() *string {
	return s.IpControlId
}

func (s *AddIpControlPolicyItemRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *AddIpControlPolicyItemRequest) SetAppId(v string) *AddIpControlPolicyItemRequest {
	s.AppId = &v
	return s
}

func (s *AddIpControlPolicyItemRequest) SetCidrIp(v string) *AddIpControlPolicyItemRequest {
	s.CidrIp = &v
	return s
}

func (s *AddIpControlPolicyItemRequest) SetIpControlId(v string) *AddIpControlPolicyItemRequest {
	s.IpControlId = &v
	return s
}

func (s *AddIpControlPolicyItemRequest) SetSecurityToken(v string) *AddIpControlPolicyItemRequest {
	s.SecurityToken = &v
	return s
}

func (s *AddIpControlPolicyItemRequest) Validate() error {
	return dara.Validate(s)
}
