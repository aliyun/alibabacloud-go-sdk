// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIpamScopeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateIpamScopeRequest
	GetClientToken() *string
	SetDryRun(v bool) *UpdateIpamScopeRequest
	GetDryRun() *bool
	SetIpamScopeDescription(v string) *UpdateIpamScopeRequest
	GetIpamScopeDescription() *string
	SetIpamScopeId(v string) *UpdateIpamScopeRequest
	GetIpamScopeId() *string
	SetIpamScopeName(v string) *UpdateIpamScopeRequest
	GetIpamScopeName() *string
	SetOwnerAccount(v string) *UpdateIpamScopeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UpdateIpamScopeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *UpdateIpamScopeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *UpdateIpamScopeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateIpamScopeRequest
	GetResourceOwnerId() *int64
}

type UpdateIpamScopeRequest struct {
	// The client token that is used to ensure the idempotence of the request. Generate a value for this parameter that is unique among different requests. The token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the RequestId of the request as the ClientToken. The RequestId of each request is unique.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: Performs a dry run. The system checks the request for potential issues, including missing required parameters, incorrect request format, and instance status. If the request fails the dry run, an error is returned. If the request passes the dry run, the DryRunOperation error code is returned.
	//
	// - **false*	- (Default): Sends a normal request. After the request passes the check, the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The description of the IPAM scope.
	//
	// The description must be 1 to 256 characters long. It must start with a letter and cannot start with `http://` or `https://`. If you do not specify this parameter, the value is not modified.
	//
	// example:
	//
	// test description
	IpamScopeDescription *string `json:"IpamScopeDescription,omitempty" xml:"IpamScopeDescription,omitempty"`
	// The ID of the IPAM scope.
	//
	// This parameter is required.
	//
	// example:
	//
	// ipam-scope-glfmcyldpm8lsy****
	IpamScopeId *string `json:"IpamScopeId,omitempty" xml:"IpamScopeId,omitempty"`
	// The name of the IPAM scope.
	//
	// The name must be 1 to 128 characters long and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// test
	IpamScopeName *string `json:"IpamScopeName,omitempty" xml:"IpamScopeName,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the IPAM is hosted. Call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to get the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s UpdateIpamScopeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateIpamScopeRequest) GoString() string {
	return s.String()
}

func (s *UpdateIpamScopeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateIpamScopeRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateIpamScopeRequest) GetIpamScopeDescription() *string {
	return s.IpamScopeDescription
}

func (s *UpdateIpamScopeRequest) GetIpamScopeId() *string {
	return s.IpamScopeId
}

func (s *UpdateIpamScopeRequest) GetIpamScopeName() *string {
	return s.IpamScopeName
}

func (s *UpdateIpamScopeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UpdateIpamScopeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateIpamScopeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateIpamScopeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateIpamScopeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateIpamScopeRequest) SetClientToken(v string) *UpdateIpamScopeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateIpamScopeRequest) SetDryRun(v bool) *UpdateIpamScopeRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateIpamScopeRequest) SetIpamScopeDescription(v string) *UpdateIpamScopeRequest {
	s.IpamScopeDescription = &v
	return s
}

func (s *UpdateIpamScopeRequest) SetIpamScopeId(v string) *UpdateIpamScopeRequest {
	s.IpamScopeId = &v
	return s
}

func (s *UpdateIpamScopeRequest) SetIpamScopeName(v string) *UpdateIpamScopeRequest {
	s.IpamScopeName = &v
	return s
}

func (s *UpdateIpamScopeRequest) SetOwnerAccount(v string) *UpdateIpamScopeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpdateIpamScopeRequest) SetOwnerId(v int64) *UpdateIpamScopeRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateIpamScopeRequest) SetRegionId(v string) *UpdateIpamScopeRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateIpamScopeRequest) SetResourceOwnerAccount(v string) *UpdateIpamScopeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateIpamScopeRequest) SetResourceOwnerId(v int64) *UpdateIpamScopeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateIpamScopeRequest) Validate() error {
	return dara.Validate(s)
}
