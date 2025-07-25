// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIpamResourceDiscoveryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddOperatingRegion(v []*string) *UpdateIpamResourceDiscoveryRequest
	GetAddOperatingRegion() []*string
	SetClientToken(v string) *UpdateIpamResourceDiscoveryRequest
	GetClientToken() *string
	SetDryRun(v bool) *UpdateIpamResourceDiscoveryRequest
	GetDryRun() *bool
	SetIpamResourceDiscoveryDescription(v string) *UpdateIpamResourceDiscoveryRequest
	GetIpamResourceDiscoveryDescription() *string
	SetIpamResourceDiscoveryId(v string) *UpdateIpamResourceDiscoveryRequest
	GetIpamResourceDiscoveryId() *string
	SetIpamResourceDiscoveryName(v string) *UpdateIpamResourceDiscoveryRequest
	GetIpamResourceDiscoveryName() *string
	SetOwnerAccount(v string) *UpdateIpamResourceDiscoveryRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UpdateIpamResourceDiscoveryRequest
	GetOwnerId() *int64
	SetRegionId(v string) *UpdateIpamResourceDiscoveryRequest
	GetRegionId() *string
	SetRemoveOperatingRegion(v []*string) *UpdateIpamResourceDiscoveryRequest
	GetRemoveOperatingRegion() []*string
	SetResourceOwnerAccount(v string) *UpdateIpamResourceDiscoveryRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateIpamResourceDiscoveryRequest
	GetResourceOwnerId() *int64
}

type UpdateIpamResourceDiscoveryRequest struct {
	// The list of effective regions to add.
	AddOperatingRegion []*string `json:"AddOperatingRegion,omitempty" xml:"AddOperatingRegion,omitempty" type:"Repeated"`
	// The client token used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the request ID as the client token. The request ID may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform the dry run. Valid values:
	//
	// 	- **true**: Performs a dry run without modifying the resource discovery instance. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error code is returned. If the request passes the dry run, the DryRunOperation error code is returned.
	//
	// 	- **false*	- (default): Performs a dry run and the actual request. If the request passes the check, an HTTP 2xx status code is returned and the resource discovery instance is modified.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The description of resource discovery.
	//
	// example:
	//
	// test description
	IpamResourceDiscoveryDescription *string `json:"IpamResourceDiscoveryDescription,omitempty" xml:"IpamResourceDiscoveryDescription,omitempty"`
	// The ID of resource discovery instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ipam-res-disco-jt5f2af2u6nk2z321****
	IpamResourceDiscoveryId *string `json:"IpamResourceDiscoveryId,omitempty" xml:"IpamResourceDiscoveryId,omitempty"`
	// The name of the resource discovery.
	//
	// example:
	//
	// test
	IpamResourceDiscoveryName *string `json:"IpamResourceDiscoveryName,omitempty" xml:"IpamResourceDiscoveryName,omitempty"`
	OwnerAccount              *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The request region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The list of effective regions to remove.
	RemoveOperatingRegion []*string `json:"RemoveOperatingRegion,omitempty" xml:"RemoveOperatingRegion,omitempty" type:"Repeated"`
	ResourceOwnerAccount  *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId       *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s UpdateIpamResourceDiscoveryRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateIpamResourceDiscoveryRequest) GoString() string {
	return s.String()
}

func (s *UpdateIpamResourceDiscoveryRequest) GetAddOperatingRegion() []*string {
	return s.AddOperatingRegion
}

func (s *UpdateIpamResourceDiscoveryRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateIpamResourceDiscoveryRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateIpamResourceDiscoveryRequest) GetIpamResourceDiscoveryDescription() *string {
	return s.IpamResourceDiscoveryDescription
}

func (s *UpdateIpamResourceDiscoveryRequest) GetIpamResourceDiscoveryId() *string {
	return s.IpamResourceDiscoveryId
}

func (s *UpdateIpamResourceDiscoveryRequest) GetIpamResourceDiscoveryName() *string {
	return s.IpamResourceDiscoveryName
}

func (s *UpdateIpamResourceDiscoveryRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UpdateIpamResourceDiscoveryRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateIpamResourceDiscoveryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateIpamResourceDiscoveryRequest) GetRemoveOperatingRegion() []*string {
	return s.RemoveOperatingRegion
}

func (s *UpdateIpamResourceDiscoveryRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateIpamResourceDiscoveryRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateIpamResourceDiscoveryRequest) SetAddOperatingRegion(v []*string) *UpdateIpamResourceDiscoveryRequest {
	s.AddOperatingRegion = v
	return s
}

func (s *UpdateIpamResourceDiscoveryRequest) SetClientToken(v string) *UpdateIpamResourceDiscoveryRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateIpamResourceDiscoveryRequest) SetDryRun(v bool) *UpdateIpamResourceDiscoveryRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateIpamResourceDiscoveryRequest) SetIpamResourceDiscoveryDescription(v string) *UpdateIpamResourceDiscoveryRequest {
	s.IpamResourceDiscoveryDescription = &v
	return s
}

func (s *UpdateIpamResourceDiscoveryRequest) SetIpamResourceDiscoveryId(v string) *UpdateIpamResourceDiscoveryRequest {
	s.IpamResourceDiscoveryId = &v
	return s
}

func (s *UpdateIpamResourceDiscoveryRequest) SetIpamResourceDiscoveryName(v string) *UpdateIpamResourceDiscoveryRequest {
	s.IpamResourceDiscoveryName = &v
	return s
}

func (s *UpdateIpamResourceDiscoveryRequest) SetOwnerAccount(v string) *UpdateIpamResourceDiscoveryRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpdateIpamResourceDiscoveryRequest) SetOwnerId(v int64) *UpdateIpamResourceDiscoveryRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateIpamResourceDiscoveryRequest) SetRegionId(v string) *UpdateIpamResourceDiscoveryRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateIpamResourceDiscoveryRequest) SetRemoveOperatingRegion(v []*string) *UpdateIpamResourceDiscoveryRequest {
	s.RemoveOperatingRegion = v
	return s
}

func (s *UpdateIpamResourceDiscoveryRequest) SetResourceOwnerAccount(v string) *UpdateIpamResourceDiscoveryRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateIpamResourceDiscoveryRequest) SetResourceOwnerId(v int64) *UpdateIpamResourceDiscoveryRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateIpamResourceDiscoveryRequest) Validate() error {
	return dara.Validate(s)
}
