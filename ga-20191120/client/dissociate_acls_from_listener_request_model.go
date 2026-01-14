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
	SetRegionId(v string) *DissociateAclsFromListenerRequest
	GetRegionId() *string
}

type DissociateAclsFromListenerRequest struct {
	// The ID of the ACL. You can disassociate up to two ACLs from a listener.
	//
	// This parameter is required.
	AclIds []*string `json:"AclIds,omitempty" xml:"AclIds,omitempty" type:"Repeated"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF3898
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to only precheck the request. Default value: false. Valid values:
	//
	// 	- **true**: prechecks the request without performing the operation. The system prechecks the required parameters, request syntax, and limits. If the request fails the precheck, an error message is returned. If the request passes the precheck, the `DryRunOperation` error code is returned.
	//
	// 	- **false**: sends the request. If the request passes the precheck, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the listener.
	//
	// This parameter is required.
	//
	// example:
	//
	// lsr-bp1bpn0kn908w4nbw****
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The ID of the region where the GA instance is deployed. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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

func (s *DissociateAclsFromListenerRequest) GetRegionId() *string {
	return s.RegionId
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

func (s *DissociateAclsFromListenerRequest) SetRegionId(v string) *DissociateAclsFromListenerRequest {
	s.RegionId = &v
	return s
}

func (s *DissociateAclsFromListenerRequest) Validate() error {
	return dara.Validate(s)
}
