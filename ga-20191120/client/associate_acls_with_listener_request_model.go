// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateAclsWithListenerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclIds(v []*string) *AssociateAclsWithListenerRequest
	GetAclIds() []*string
	SetAclType(v string) *AssociateAclsWithListenerRequest
	GetAclType() *string
	SetClientToken(v string) *AssociateAclsWithListenerRequest
	GetClientToken() *string
	SetDryRun(v bool) *AssociateAclsWithListenerRequest
	GetDryRun() *bool
	SetListenerId(v string) *AssociateAclsWithListenerRequest
	GetListenerId() *string
	SetRegionId(v string) *AssociateAclsWithListenerRequest
	GetRegionId() *string
}

type AssociateAclsWithListenerRequest struct {
	// The ID of the ACL. You can associate up to two ACL IDs.
	//
	// This parameter is required.
	AclIds []*string `json:"AclIds,omitempty" xml:"AclIds,omitempty" type:"Repeated"`
	// The type of the ACL. Valid values:
	//
	// 	- **white**: Only requests from the IP addresses or CIDR blocks in the ACL are forwarded. Whitelists are suitable for scenarios in which you want to allow access from specific IP addresses to an application. If a whitelist is improperly configured, risks may arise. After a whitelist is configured for a listener, only requests from the IP addresses that are added to the whitelist are distributed by the listener. If a whitelist is enabled but no IP address is added to the whitelist, the listener forwards all requests.
	//
	// 	- **black**: All requests from the IP addresses or CIDR blocks in the ACL are rejected. Blacklists are suitable for scenarios in which you want to deny access from specific IP addresses to an application. If the blacklist is enabled but no IP addresses are added to the ACL, the listener forwards all requests.
	//
	// This parameter is required.
	//
	// example:
	//
	// White
	AclType *string `json:"AclType,omitempty" xml:"AclType,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 02fb3da4****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to only precheck the request. Default value: false. Valid values:
	//
	// 	- **true**: prechecks the request without performing the operation. The system checks the required parameters, request syntax, and limits. If the request fails the precheck, an error message is returned. If the request passes the precheck, the `DryRunOperation` error code is returned.
	//
	// 	- **false**: sends the request. If the request passes the precheck, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The listener ID.
	//
	// Only intelligent routing listeners support ACLs.
	//
	// This parameter is required.
	//
	// example:
	//
	// lsr-bp1bpn0kn908w4nbw****
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The region ID of the Global Accelerator (GA) instance. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AssociateAclsWithListenerRequest) String() string {
	return dara.Prettify(s)
}

func (s AssociateAclsWithListenerRequest) GoString() string {
	return s.String()
}

func (s *AssociateAclsWithListenerRequest) GetAclIds() []*string {
	return s.AclIds
}

func (s *AssociateAclsWithListenerRequest) GetAclType() *string {
	return s.AclType
}

func (s *AssociateAclsWithListenerRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AssociateAclsWithListenerRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *AssociateAclsWithListenerRequest) GetListenerId() *string {
	return s.ListenerId
}

func (s *AssociateAclsWithListenerRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AssociateAclsWithListenerRequest) SetAclIds(v []*string) *AssociateAclsWithListenerRequest {
	s.AclIds = v
	return s
}

func (s *AssociateAclsWithListenerRequest) SetAclType(v string) *AssociateAclsWithListenerRequest {
	s.AclType = &v
	return s
}

func (s *AssociateAclsWithListenerRequest) SetClientToken(v string) *AssociateAclsWithListenerRequest {
	s.ClientToken = &v
	return s
}

func (s *AssociateAclsWithListenerRequest) SetDryRun(v bool) *AssociateAclsWithListenerRequest {
	s.DryRun = &v
	return s
}

func (s *AssociateAclsWithListenerRequest) SetListenerId(v string) *AssociateAclsWithListenerRequest {
	s.ListenerId = &v
	return s
}

func (s *AssociateAclsWithListenerRequest) SetRegionId(v string) *AssociateAclsWithListenerRequest {
	s.RegionId = &v
	return s
}

func (s *AssociateAclsWithListenerRequest) Validate() error {
	return dara.Validate(s)
}
