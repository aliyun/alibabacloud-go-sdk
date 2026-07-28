// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePublicIpAddressPoolAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdatePublicIpAddressPoolAttributeRequest
	GetClientToken() *string
	SetDescription(v string) *UpdatePublicIpAddressPoolAttributeRequest
	GetDescription() *string
	SetDryRun(v bool) *UpdatePublicIpAddressPoolAttributeRequest
	GetDryRun() *bool
	SetName(v string) *UpdatePublicIpAddressPoolAttributeRequest
	GetName() *string
	SetOwnerAccount(v string) *UpdatePublicIpAddressPoolAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UpdatePublicIpAddressPoolAttributeRequest
	GetOwnerId() *int64
	SetPublicIpAddressPoolId(v string) *UpdatePublicIpAddressPoolAttributeRequest
	GetPublicIpAddressPoolId() *string
	SetRegionId(v string) *UpdatePublicIpAddressPoolAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *UpdatePublicIpAddressPoolAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdatePublicIpAddressPoolAttributeRequest
	GetResourceOwnerId() *int64
}

type UpdatePublicIpAddressPoolAttributeRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The ClientToken value can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- may be different for each API request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the IP address pool.
	//
	// The description must be 0 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// AddressPoolDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: performs a dry run without modifying the IP address pool attributes. The system checks the required parameters, request format, and service limits. If the check fails, the corresponding error is returned. If the check succeeds, the `DryRunOperation` error code is returned.
	//
	// - **false*	- (default): performs a dry run and then sends the request. If the check succeeds, an HTTP 2xx status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The name of the IP address pool.
	//
	// The name must be 0 to 128 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// AddressPoolName
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The instance ID of the IPAM pool.
	//
	// This parameter is required.
	//
	// example:
	//
	// pippool-6wetvn6fumkgycssx****
	PublicIpAddressPoolId *string `json:"PublicIpAddressPoolId,omitempty" xml:"PublicIpAddressPoolId,omitempty"`
	// The region ID of the IP address pool to modify.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-chengdu
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s UpdatePublicIpAddressPoolAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePublicIpAddressPoolAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdatePublicIpAddressPoolAttributeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdatePublicIpAddressPoolAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdatePublicIpAddressPoolAttributeRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdatePublicIpAddressPoolAttributeRequest) GetName() *string {
	return s.Name
}

func (s *UpdatePublicIpAddressPoolAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UpdatePublicIpAddressPoolAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdatePublicIpAddressPoolAttributeRequest) GetPublicIpAddressPoolId() *string {
	return s.PublicIpAddressPoolId
}

func (s *UpdatePublicIpAddressPoolAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdatePublicIpAddressPoolAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdatePublicIpAddressPoolAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdatePublicIpAddressPoolAttributeRequest) SetClientToken(v string) *UpdatePublicIpAddressPoolAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdatePublicIpAddressPoolAttributeRequest) SetDescription(v string) *UpdatePublicIpAddressPoolAttributeRequest {
	s.Description = &v
	return s
}

func (s *UpdatePublicIpAddressPoolAttributeRequest) SetDryRun(v bool) *UpdatePublicIpAddressPoolAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *UpdatePublicIpAddressPoolAttributeRequest) SetName(v string) *UpdatePublicIpAddressPoolAttributeRequest {
	s.Name = &v
	return s
}

func (s *UpdatePublicIpAddressPoolAttributeRequest) SetOwnerAccount(v string) *UpdatePublicIpAddressPoolAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpdatePublicIpAddressPoolAttributeRequest) SetOwnerId(v int64) *UpdatePublicIpAddressPoolAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdatePublicIpAddressPoolAttributeRequest) SetPublicIpAddressPoolId(v string) *UpdatePublicIpAddressPoolAttributeRequest {
	s.PublicIpAddressPoolId = &v
	return s
}

func (s *UpdatePublicIpAddressPoolAttributeRequest) SetRegionId(v string) *UpdatePublicIpAddressPoolAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *UpdatePublicIpAddressPoolAttributeRequest) SetResourceOwnerAccount(v string) *UpdatePublicIpAddressPoolAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdatePublicIpAddressPoolAttributeRequest) SetResourceOwnerId(v int64) *UpdatePublicIpAddressPoolAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdatePublicIpAddressPoolAttributeRequest) Validate() error {
	return dara.Validate(s)
}
