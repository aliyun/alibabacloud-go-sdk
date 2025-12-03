// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIpControlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateIpControlRequest
	GetDescription() *string
	SetIpControlName(v string) *CreateIpControlRequest
	GetIpControlName() *string
	SetIpControlPolicys(v []*CreateIpControlRequestIpControlPolicys) *CreateIpControlRequest
	GetIpControlPolicys() []*CreateIpControlRequestIpControlPolicys
	SetIpControlType(v string) *CreateIpControlRequest
	GetIpControlType() *string
	SetSecurityToken(v string) *CreateIpControlRequest
	GetSecurityToken() *string
}

type CreateIpControlRequest struct {
	// The description. The description can be up to 200 characters in length.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the ACL. The name must be 4 to 50 characters in length, and can contain letters, digits, and underscores (_). The name cannot start with an underscore (_).``
	//
	// This parameter is required.
	//
	// example:
	//
	// controlNameTest
	IpControlName *string `json:"IpControlName,omitempty" xml:"IpControlName,omitempty"`
	// The information about the policies. The information is an array of ipcontrolpolicys data.
	IpControlPolicys []*CreateIpControlRequestIpControlPolicys `json:"IpControlPolicys,omitempty" xml:"IpControlPolicys,omitempty" type:"Repeated"`
	// The type of the ACL. Valid values:
	//
	// 	- **ALLOW**: an IP address whitelist
	//
	// 	- **REFUSE**: an IP address blacklist
	//
	// This parameter is required.
	//
	// example:
	//
	// ALLOW
	IpControlType *string `json:"IpControlType,omitempty" xml:"IpControlType,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s CreateIpControlRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateIpControlRequest) GoString() string {
	return s.String()
}

func (s *CreateIpControlRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateIpControlRequest) GetIpControlName() *string {
	return s.IpControlName
}

func (s *CreateIpControlRequest) GetIpControlPolicys() []*CreateIpControlRequestIpControlPolicys {
	return s.IpControlPolicys
}

func (s *CreateIpControlRequest) GetIpControlType() *string {
	return s.IpControlType
}

func (s *CreateIpControlRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *CreateIpControlRequest) SetDescription(v string) *CreateIpControlRequest {
	s.Description = &v
	return s
}

func (s *CreateIpControlRequest) SetIpControlName(v string) *CreateIpControlRequest {
	s.IpControlName = &v
	return s
}

func (s *CreateIpControlRequest) SetIpControlPolicys(v []*CreateIpControlRequestIpControlPolicys) *CreateIpControlRequest {
	s.IpControlPolicys = v
	return s
}

func (s *CreateIpControlRequest) SetIpControlType(v string) *CreateIpControlRequest {
	s.IpControlType = &v
	return s
}

func (s *CreateIpControlRequest) SetSecurityToken(v string) *CreateIpControlRequest {
	s.SecurityToken = &v
	return s
}

func (s *CreateIpControlRequest) Validate() error {
	if s.IpControlPolicys != nil {
		for _, item := range s.IpControlPolicys {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateIpControlRequestIpControlPolicys struct {
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
	// 	- Valid values of N in IpControlPolicys.N: `[1,100]`.
	//
	// example:
	//
	// 11111
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The IP address or CIDR block involved in a policy.
	//
	// 	- If you want to specify a policy when you create an ACL, this parameter is required.
	//
	// 	- The IP address or CIDR block that is defined in each policy. Separate multiple IP addresses or CIDR blocks with semicolons (;). You can add a maximum of 10 IP addresses or CIDR blocks.
	//
	// 	- Valid values of N in IpControlPolicys.N: `[1,100]`.
	//
	// example:
	//
	// 114.1.1.0/24
	CidrIp *string `json:"CidrIp,omitempty" xml:"CidrIp,omitempty"`
}

func (s CreateIpControlRequestIpControlPolicys) String() string {
	return dara.Prettify(s)
}

func (s CreateIpControlRequestIpControlPolicys) GoString() string {
	return s.String()
}

func (s *CreateIpControlRequestIpControlPolicys) GetAppId() *string {
	return s.AppId
}

func (s *CreateIpControlRequestIpControlPolicys) GetCidrIp() *string {
	return s.CidrIp
}

func (s *CreateIpControlRequestIpControlPolicys) SetAppId(v string) *CreateIpControlRequestIpControlPolicys {
	s.AppId = &v
	return s
}

func (s *CreateIpControlRequestIpControlPolicys) SetCidrIp(v string) *CreateIpControlRequestIpControlPolicys {
	s.CidrIp = &v
	return s
}

func (s *CreateIpControlRequestIpControlPolicys) Validate() error {
	return dara.Validate(s)
}
