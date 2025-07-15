// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApiGroupNetworkPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *ModifyApiGroupNetworkPolicyRequest
	GetGroupId() *string
	SetHttpsPolicy(v string) *ModifyApiGroupNetworkPolicyRequest
	GetHttpsPolicy() *string
	SetInnerDomainEnable(v bool) *ModifyApiGroupNetworkPolicyRequest
	GetInnerDomainEnable() *bool
	SetInternetEnable(v bool) *ModifyApiGroupNetworkPolicyRequest
	GetInternetEnable() *bool
	SetInternetIPV6Enable(v bool) *ModifyApiGroupNetworkPolicyRequest
	GetInternetIPV6Enable() *bool
	SetSecurityToken(v string) *ModifyApiGroupNetworkPolicyRequest
	GetSecurityToken() *string
	SetVpcIntranetEnable(v bool) *ModifyApiGroupNetworkPolicyRequest
	GetVpcIntranetEnable() *bool
	SetVpcSlbIntranetEnable(v bool) *ModifyApiGroupNetworkPolicyRequest
	GetVpcSlbIntranetEnable() *bool
}

type ModifyApiGroupNetworkPolicyRequest struct {
	// The ID of the API group.
	//
	// This parameter is required.
	//
	// example:
	//
	// b0162c75d7d34ff48506f1aff878b05e
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The HTTPS security policy.
	//
	// example:
	//
	// HTTPS1_1_TLS1_0
	HttpsPolicy *string `json:"HttpsPolicy,omitempty" xml:"HttpsPolicy,omitempty"`
	// Specifies whether to disable the public second-level domain name.
	//
	// example:
	//
	// true
	InnerDomainEnable *bool `json:"InnerDomainEnable,omitempty" xml:"InnerDomainEnable,omitempty"`
	// Specifies whether to enable the virtual private cloud (VPC) second-level domain name.
	//
	// example:
	//
	// true
	InternetEnable *bool `json:"InternetEnable,omitempty" xml:"InternetEnable,omitempty"`
	// Specifies whether to enable IPv6. Valid values: **true*	- and **false**.
	//
	// example:
	//
	// true
	InternetIPV6Enable *bool   `json:"InternetIPV6Enable,omitempty" xml:"InternetIPV6Enable,omitempty"`
	SecurityToken      *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// Specifies whether to enable the VPC domain name. Valid values:
	//
	// 	- TRUE
	//
	// 	- FALSE
	//
	// example:
	//
	// false
	VpcIntranetEnable *bool `json:"VpcIntranetEnable,omitempty" xml:"VpcIntranetEnable,omitempty"`
	// Specifies whether to enable the self-calling domain name.
	//
	// example:
	//
	// false
	VpcSlbIntranetEnable *bool `json:"VpcSlbIntranetEnable,omitempty" xml:"VpcSlbIntranetEnable,omitempty"`
}

func (s ModifyApiGroupNetworkPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyApiGroupNetworkPolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifyApiGroupNetworkPolicyRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *ModifyApiGroupNetworkPolicyRequest) GetHttpsPolicy() *string {
	return s.HttpsPolicy
}

func (s *ModifyApiGroupNetworkPolicyRequest) GetInnerDomainEnable() *bool {
	return s.InnerDomainEnable
}

func (s *ModifyApiGroupNetworkPolicyRequest) GetInternetEnable() *bool {
	return s.InternetEnable
}

func (s *ModifyApiGroupNetworkPolicyRequest) GetInternetIPV6Enable() *bool {
	return s.InternetIPV6Enable
}

func (s *ModifyApiGroupNetworkPolicyRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ModifyApiGroupNetworkPolicyRequest) GetVpcIntranetEnable() *bool {
	return s.VpcIntranetEnable
}

func (s *ModifyApiGroupNetworkPolicyRequest) GetVpcSlbIntranetEnable() *bool {
	return s.VpcSlbIntranetEnable
}

func (s *ModifyApiGroupNetworkPolicyRequest) SetGroupId(v string) *ModifyApiGroupNetworkPolicyRequest {
	s.GroupId = &v
	return s
}

func (s *ModifyApiGroupNetworkPolicyRequest) SetHttpsPolicy(v string) *ModifyApiGroupNetworkPolicyRequest {
	s.HttpsPolicy = &v
	return s
}

func (s *ModifyApiGroupNetworkPolicyRequest) SetInnerDomainEnable(v bool) *ModifyApiGroupNetworkPolicyRequest {
	s.InnerDomainEnable = &v
	return s
}

func (s *ModifyApiGroupNetworkPolicyRequest) SetInternetEnable(v bool) *ModifyApiGroupNetworkPolicyRequest {
	s.InternetEnable = &v
	return s
}

func (s *ModifyApiGroupNetworkPolicyRequest) SetInternetIPV6Enable(v bool) *ModifyApiGroupNetworkPolicyRequest {
	s.InternetIPV6Enable = &v
	return s
}

func (s *ModifyApiGroupNetworkPolicyRequest) SetSecurityToken(v string) *ModifyApiGroupNetworkPolicyRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyApiGroupNetworkPolicyRequest) SetVpcIntranetEnable(v bool) *ModifyApiGroupNetworkPolicyRequest {
	s.VpcIntranetEnable = &v
	return s
}

func (s *ModifyApiGroupNetworkPolicyRequest) SetVpcSlbIntranetEnable(v bool) *ModifyApiGroupNetworkPolicyRequest {
	s.VpcSlbIntranetEnable = &v
	return s
}

func (s *ModifyApiGroupNetworkPolicyRequest) Validate() error {
	return dara.Validate(s)
}
