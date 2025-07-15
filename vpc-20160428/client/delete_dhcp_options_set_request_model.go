// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDhcpOptionsSetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteDhcpOptionsSetRequest
	GetClientToken() *string
	SetDhcpOptionsSetId(v string) *DeleteDhcpOptionsSetRequest
	GetDhcpOptionsSetId() *string
	SetDryRun(v bool) *DeleteDhcpOptionsSetRequest
	GetDryRun() *bool
	SetOwnerAccount(v string) *DeleteDhcpOptionsSetRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteDhcpOptionsSetRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteDhcpOptionsSetRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteDhcpOptionsSetRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteDhcpOptionsSetRequest
	GetResourceOwnerId() *int64
}

type DeleteDhcpOptionsSetRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe637760
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the DHCP options set to be deleted.
	//
	// This parameter is required.
	//
	// example:
	//
	// dopt-o6w0df4epg9zo8isy****
	DhcpOptionsSetId *string `json:"DhcpOptionsSetId,omitempty" xml:"DhcpOptionsSetId,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun       *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region of the DHCP options set to be deleted. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
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

func (s DeleteDhcpOptionsSetRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDhcpOptionsSetRequest) GoString() string {
	return s.String()
}

func (s *DeleteDhcpOptionsSetRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteDhcpOptionsSetRequest) GetDhcpOptionsSetId() *string {
	return s.DhcpOptionsSetId
}

func (s *DeleteDhcpOptionsSetRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteDhcpOptionsSetRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteDhcpOptionsSetRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteDhcpOptionsSetRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteDhcpOptionsSetRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteDhcpOptionsSetRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteDhcpOptionsSetRequest) SetClientToken(v string) *DeleteDhcpOptionsSetRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteDhcpOptionsSetRequest) SetDhcpOptionsSetId(v string) *DeleteDhcpOptionsSetRequest {
	s.DhcpOptionsSetId = &v
	return s
}

func (s *DeleteDhcpOptionsSetRequest) SetDryRun(v bool) *DeleteDhcpOptionsSetRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteDhcpOptionsSetRequest) SetOwnerAccount(v string) *DeleteDhcpOptionsSetRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteDhcpOptionsSetRequest) SetOwnerId(v int64) *DeleteDhcpOptionsSetRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteDhcpOptionsSetRequest) SetRegionId(v string) *DeleteDhcpOptionsSetRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteDhcpOptionsSetRequest) SetResourceOwnerAccount(v string) *DeleteDhcpOptionsSetRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteDhcpOptionsSetRequest) SetResourceOwnerId(v int64) *DeleteDhcpOptionsSetRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteDhcpOptionsSetRequest) Validate() error {
	return dara.Validate(s)
}
