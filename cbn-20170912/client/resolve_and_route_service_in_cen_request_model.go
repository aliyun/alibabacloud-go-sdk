// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResolveAndRouteServiceInCenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessRegionIds(v []*string) *ResolveAndRouteServiceInCenRequest
	GetAccessRegionIds() []*string
	SetCenId(v string) *ResolveAndRouteServiceInCenRequest
	GetCenId() *string
	SetClientToken(v string) *ResolveAndRouteServiceInCenRequest
	GetClientToken() *string
	SetDescription(v string) *ResolveAndRouteServiceInCenRequest
	GetDescription() *string
	SetHost(v string) *ResolveAndRouteServiceInCenRequest
	GetHost() *string
	SetHostRegionId(v string) *ResolveAndRouteServiceInCenRequest
	GetHostRegionId() *string
	SetHostVpcId(v string) *ResolveAndRouteServiceInCenRequest
	GetHostVpcId() *string
	SetOwnerAccount(v string) *ResolveAndRouteServiceInCenRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ResolveAndRouteServiceInCenRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ResolveAndRouteServiceInCenRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ResolveAndRouteServiceInCenRequest
	GetResourceOwnerId() *int64
}

type ResolveAndRouteServiceInCenRequest struct {
	// The IDs of the regions where the cloud service is accessed.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	AccessRegionIds []*string `json:"AccessRegionIds,omitempty" xml:"AccessRegionIds,omitempty" type:"Repeated"`
	// The ID of the CEN instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cen-ckwa2hhmuislse****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not set this parameter, **ClientToken*	- is set to the value of **RequestId**. The value of **RequestId*	- for each API request may be different.
	//
	// example:
	//
	// 02fb3da4****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the cloud service.
	//
	// This parameter is optional. If you enter a description, it must be 1 to 256 characters in length and cannot start with http:// or https://.
	//
	// example:
	//
	// descname
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The IP addresses or CIDR blocks of the cloud service.
	//
	// > In most cases, multiple IP addresses or CIDR blocks are assigned to a cloud service. We recommend that you call this operation multiple times to add all IP addresses and CIDR blocks of the cloud service.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100.118.28.0/24
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The ID of the region in which the cloud service is deployed.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	HostRegionId *string `json:"HostRegionId,omitempty" xml:"HostRegionId,omitempty"`
	// The ID of the VPC that is associated with the cloud service.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-o6woh5s494zueq40v****
	HostVpcId            *string `json:"HostVpcId,omitempty" xml:"HostVpcId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ResolveAndRouteServiceInCenRequest) String() string {
	return dara.Prettify(s)
}

func (s ResolveAndRouteServiceInCenRequest) GoString() string {
	return s.String()
}

func (s *ResolveAndRouteServiceInCenRequest) GetAccessRegionIds() []*string {
	return s.AccessRegionIds
}

func (s *ResolveAndRouteServiceInCenRequest) GetCenId() *string {
	return s.CenId
}

func (s *ResolveAndRouteServiceInCenRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ResolveAndRouteServiceInCenRequest) GetDescription() *string {
	return s.Description
}

func (s *ResolveAndRouteServiceInCenRequest) GetHost() *string {
	return s.Host
}

func (s *ResolveAndRouteServiceInCenRequest) GetHostRegionId() *string {
	return s.HostRegionId
}

func (s *ResolveAndRouteServiceInCenRequest) GetHostVpcId() *string {
	return s.HostVpcId
}

func (s *ResolveAndRouteServiceInCenRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ResolveAndRouteServiceInCenRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ResolveAndRouteServiceInCenRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ResolveAndRouteServiceInCenRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ResolveAndRouteServiceInCenRequest) SetAccessRegionIds(v []*string) *ResolveAndRouteServiceInCenRequest {
	s.AccessRegionIds = v
	return s
}

func (s *ResolveAndRouteServiceInCenRequest) SetCenId(v string) *ResolveAndRouteServiceInCenRequest {
	s.CenId = &v
	return s
}

func (s *ResolveAndRouteServiceInCenRequest) SetClientToken(v string) *ResolveAndRouteServiceInCenRequest {
	s.ClientToken = &v
	return s
}

func (s *ResolveAndRouteServiceInCenRequest) SetDescription(v string) *ResolveAndRouteServiceInCenRequest {
	s.Description = &v
	return s
}

func (s *ResolveAndRouteServiceInCenRequest) SetHost(v string) *ResolveAndRouteServiceInCenRequest {
	s.Host = &v
	return s
}

func (s *ResolveAndRouteServiceInCenRequest) SetHostRegionId(v string) *ResolveAndRouteServiceInCenRequest {
	s.HostRegionId = &v
	return s
}

func (s *ResolveAndRouteServiceInCenRequest) SetHostVpcId(v string) *ResolveAndRouteServiceInCenRequest {
	s.HostVpcId = &v
	return s
}

func (s *ResolveAndRouteServiceInCenRequest) SetOwnerAccount(v string) *ResolveAndRouteServiceInCenRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ResolveAndRouteServiceInCenRequest) SetOwnerId(v int64) *ResolveAndRouteServiceInCenRequest {
	s.OwnerId = &v
	return s
}

func (s *ResolveAndRouteServiceInCenRequest) SetResourceOwnerAccount(v string) *ResolveAndRouteServiceInCenRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ResolveAndRouteServiceInCenRequest) SetResourceOwnerId(v int64) *ResolveAndRouteServiceInCenRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ResolveAndRouteServiceInCenRequest) Validate() error {
	return dara.Validate(s)
}
