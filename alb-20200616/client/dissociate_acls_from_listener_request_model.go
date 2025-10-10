// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDissociateAclsFromListenerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclIds(v []*string) *DissociateAclsFromListenerRequest
	GetAclIds() []*string
	SetClientToken(v string) *DissociateAclsFromListenerRequest
	GetClientToken() *string
	SetDryRun(v bool) *DissociateAclsFromListenerRequest
	GetDryRun() *bool
	SetListenerId(v string) *DissociateAclsFromListenerRequest
	GetListenerId() *string
}

type DissociateAclsFromListenerRequest struct {
	// The access control list (ACL) IDs. You can disassociate at most three ACLs from a listener in each call.
	//
	// This parameter is required.
	AclIds []*string `json:"AclIds,omitempty" xml:"AclIds,omitempty" type:"Repeated"`
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

func (s DissociateAclsFromListenerRequest) String() string {
	return dara.Prettify(s)
}

func (s DissociateAclsFromListenerRequest) GoString() string {
	return s.String()
}

func (s *DissociateAclsFromListenerRequest) GetAclIds() []*string {
	return s.AclIds
}

func (s *DissociateAclsFromListenerRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DissociateAclsFromListenerRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DissociateAclsFromListenerRequest) GetListenerId() *string {
	return s.ListenerId
}

func (s *DissociateAclsFromListenerRequest) SetAclIds(v []*string) *DissociateAclsFromListenerRequest {
	s.AclIds = v
	return s
}

func (s *DissociateAclsFromListenerRequest) SetClientToken(v string) *DissociateAclsFromListenerRequest {
	s.ClientToken = &v
	return s
}

func (s *DissociateAclsFromListenerRequest) SetDryRun(v bool) *DissociateAclsFromListenerRequest {
	s.DryRun = &v
	return s
}

func (s *DissociateAclsFromListenerRequest) SetListenerId(v string) *DissociateAclsFromListenerRequest {
	s.ListenerId = &v
	return s
}

func (s *DissociateAclsFromListenerRequest) Validate() error {
	return dara.Validate(s)
}
