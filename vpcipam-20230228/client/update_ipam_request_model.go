// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIpamRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddOperatingRegion(v []*string) *UpdateIpamRequest
	GetAddOperatingRegion() []*string
	SetClientToken(v string) *UpdateIpamRequest
	GetClientToken() *string
	SetDryRun(v bool) *UpdateIpamRequest
	GetDryRun() *bool
	SetIpamDescription(v string) *UpdateIpamRequest
	GetIpamDescription() *string
	SetIpamId(v string) *UpdateIpamRequest
	GetIpamId() *string
	SetIpamName(v string) *UpdateIpamRequest
	GetIpamName() *string
	SetOwnerAccount(v string) *UpdateIpamRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UpdateIpamRequest
	GetOwnerId() *int64
	SetRegionId(v string) *UpdateIpamRequest
	GetRegionId() *string
	SetRemoveOperatingRegion(v []*string) *UpdateIpamRequest
	GetRemoveOperatingRegion() []*string
	SetResourceOwnerAccount(v string) *UpdateIpamRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateIpamRequest
	GetResourceOwnerId() *int64
}

type UpdateIpamRequest struct {
	// The effective region that you want to add.
	AddOperatingRegion []*string `json:"AddOperatingRegion,omitempty" xml:"AddOperatingRegion,omitempty" type:"Repeated"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the request ID as the client token. The request ID may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the DryRunOperation error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The description of the IPAM.
	//
	// It must be 2 to 256 characters in length and must start with a letter. It cannot start with `http://` or `https://`. The default value is empty.
	//
	// example:
	//
	// test description
	IpamDescription *string `json:"IpamDescription,omitempty" xml:"IpamDescription,omitempty"`
	// The ID of the IPAM.
	//
	// This parameter is required.
	//
	// example:
	//
	// ipam-ccxbnsbhew0d6t****
	IpamId *string `json:"IpamId,omitempty" xml:"IpamId,omitempty"`
	// The name of the IPAM.
	//
	// It must be 1 to 128 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// test
	IpamName     *string `json:"IpamName,omitempty" xml:"IpamName,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the IPAM instance is hosted. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The effective region that you want to remove.
	RemoveOperatingRegion []*string `json:"RemoveOperatingRegion,omitempty" xml:"RemoveOperatingRegion,omitempty" type:"Repeated"`
	ResourceOwnerAccount  *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId       *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s UpdateIpamRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateIpamRequest) GoString() string {
	return s.String()
}

func (s *UpdateIpamRequest) GetAddOperatingRegion() []*string {
	return s.AddOperatingRegion
}

func (s *UpdateIpamRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateIpamRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateIpamRequest) GetIpamDescription() *string {
	return s.IpamDescription
}

func (s *UpdateIpamRequest) GetIpamId() *string {
	return s.IpamId
}

func (s *UpdateIpamRequest) GetIpamName() *string {
	return s.IpamName
}

func (s *UpdateIpamRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UpdateIpamRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateIpamRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateIpamRequest) GetRemoveOperatingRegion() []*string {
	return s.RemoveOperatingRegion
}

func (s *UpdateIpamRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateIpamRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateIpamRequest) SetAddOperatingRegion(v []*string) *UpdateIpamRequest {
	s.AddOperatingRegion = v
	return s
}

func (s *UpdateIpamRequest) SetClientToken(v string) *UpdateIpamRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateIpamRequest) SetDryRun(v bool) *UpdateIpamRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateIpamRequest) SetIpamDescription(v string) *UpdateIpamRequest {
	s.IpamDescription = &v
	return s
}

func (s *UpdateIpamRequest) SetIpamId(v string) *UpdateIpamRequest {
	s.IpamId = &v
	return s
}

func (s *UpdateIpamRequest) SetIpamName(v string) *UpdateIpamRequest {
	s.IpamName = &v
	return s
}

func (s *UpdateIpamRequest) SetOwnerAccount(v string) *UpdateIpamRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpdateIpamRequest) SetOwnerId(v int64) *UpdateIpamRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateIpamRequest) SetRegionId(v string) *UpdateIpamRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateIpamRequest) SetRemoveOperatingRegion(v []*string) *UpdateIpamRequest {
	s.RemoveOperatingRegion = v
	return s
}

func (s *UpdateIpamRequest) SetResourceOwnerAccount(v string) *UpdateIpamRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateIpamRequest) SetResourceOwnerId(v int64) *UpdateIpamRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateIpamRequest) Validate() error {
	return dara.Validate(s)
}
