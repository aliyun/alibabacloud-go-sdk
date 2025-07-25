// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIpamScopeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteIpamScopeRequest
	GetClientToken() *string
	SetDryRun(v bool) *DeleteIpamScopeRequest
	GetDryRun() *bool
	SetIpamScopeId(v string) *DeleteIpamScopeRequest
	GetIpamScopeId() *string
	SetOwnerAccount(v string) *DeleteIpamScopeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteIpamScopeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteIpamScopeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteIpamScopeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteIpamScopeRequest
	GetResourceOwnerId() *int64
}

type DeleteIpamScopeRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the request ID as the client token. The request ID may be different for each request.
	//
	// example:
	//
	// 88144bdb-b190-4813-a3f5-66cc86694162
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the precheck, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the IPAM scope.
	//
	// This parameter is required.
	//
	// example:
	//
	// ipam-scope-glfmcyldpm8lsy****
	IpamScopeId  *string `json:"IpamScopeId,omitempty" xml:"IpamScopeId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the IPAM instance is hosted. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
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

func (s DeleteIpamScopeRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteIpamScopeRequest) GoString() string {
	return s.String()
}

func (s *DeleteIpamScopeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteIpamScopeRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteIpamScopeRequest) GetIpamScopeId() *string {
	return s.IpamScopeId
}

func (s *DeleteIpamScopeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteIpamScopeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteIpamScopeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteIpamScopeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteIpamScopeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteIpamScopeRequest) SetClientToken(v string) *DeleteIpamScopeRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteIpamScopeRequest) SetDryRun(v bool) *DeleteIpamScopeRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteIpamScopeRequest) SetIpamScopeId(v string) *DeleteIpamScopeRequest {
	s.IpamScopeId = &v
	return s
}

func (s *DeleteIpamScopeRequest) SetOwnerAccount(v string) *DeleteIpamScopeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteIpamScopeRequest) SetOwnerId(v int64) *DeleteIpamScopeRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteIpamScopeRequest) SetRegionId(v string) *DeleteIpamScopeRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteIpamScopeRequest) SetResourceOwnerAccount(v string) *DeleteIpamScopeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteIpamScopeRequest) SetResourceOwnerId(v int64) *DeleteIpamScopeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteIpamScopeRequest) Validate() error {
	return dara.Validate(s)
}
