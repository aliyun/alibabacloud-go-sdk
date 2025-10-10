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
}

type AssociateAclsWithListenerRequest struct {
	// The IDs of the ACLs. You can specify up to three IDs in each call.
	//
	// This parameter is required.
	AclIds []*string `json:"AclIds,omitempty" xml:"AclIds,omitempty" type:"Repeated"`
	// The type of the ACL. Valid values:
	//
	// 	- **White**: a whitelist. Only requests from the IP addresses or CIDR blocks in the ACL are forwarded. The whitelist applies to scenarios in which you want to allow only specific IP addresses to access an application. Your service may be adversely affected if the whitelist is not properly configured. If a whitelist is configured for a listener, only requests from IP addresses that are added to the whitelist are forwarded by the listener. If you enable a whitelist but do not add an IP address to the whitelist, the listener forwards all requests.
	//
	// 	- **Black**: a blacklist. All requests from the IP addresses or CIDR blocks in the ACL are blocked. The blacklist applies to scenarios in which you want to block access from specific IP addresses to an application. If a blacklist is configured for a listener but no IP address is added to the blacklist, the listener forwards all requests.
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
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF3898
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the listener.
	//
	// This parameter is required.
	//
	// example:
	//
	// lsr-bp1bpn0kn908w4nbw****
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
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

func (s *AssociateAclsWithListenerRequest) Validate() error {
	return dara.Validate(s)
}
