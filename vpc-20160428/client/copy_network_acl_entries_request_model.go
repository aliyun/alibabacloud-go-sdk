// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyNetworkAclEntriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CopyNetworkAclEntriesRequest
	GetClientToken() *string
	SetDryRun(v bool) *CopyNetworkAclEntriesRequest
	GetDryRun() *bool
	SetNetworkAclId(v string) *CopyNetworkAclEntriesRequest
	GetNetworkAclId() *string
	SetOwnerAccount(v string) *CopyNetworkAclEntriesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CopyNetworkAclEntriesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CopyNetworkAclEntriesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CopyNetworkAclEntriesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CopyNetworkAclEntriesRequest
	GetResourceOwnerId() *int64
	SetSourceNetworkAclId(v string) *CopyNetworkAclEntriesRequest
	GetSourceNetworkAclId() *string
}

type CopyNetworkAclEntriesRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 123
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the network ACL.
	//
	// This parameter is required.
	//
	// example:
	//
	// nacl-a2do9e413e0spxxxxxxxx
	NetworkAclId *string `json:"NetworkAclId,omitempty" xml:"NetworkAclId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the network ACL. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the network ACL whose rules you want to copy.
	//
	// This parameter is required.
	//
	// example:
	//
	// nacl-ghuo9ehg3e0spxxxxxxxx
	SourceNetworkAclId *string `json:"SourceNetworkAclId,omitempty" xml:"SourceNetworkAclId,omitempty"`
}

func (s CopyNetworkAclEntriesRequest) String() string {
	return dara.Prettify(s)
}

func (s CopyNetworkAclEntriesRequest) GoString() string {
	return s.String()
}

func (s *CopyNetworkAclEntriesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CopyNetworkAclEntriesRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CopyNetworkAclEntriesRequest) GetNetworkAclId() *string {
	return s.NetworkAclId
}

func (s *CopyNetworkAclEntriesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CopyNetworkAclEntriesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CopyNetworkAclEntriesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CopyNetworkAclEntriesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CopyNetworkAclEntriesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CopyNetworkAclEntriesRequest) GetSourceNetworkAclId() *string {
	return s.SourceNetworkAclId
}

func (s *CopyNetworkAclEntriesRequest) SetClientToken(v string) *CopyNetworkAclEntriesRequest {
	s.ClientToken = &v
	return s
}

func (s *CopyNetworkAclEntriesRequest) SetDryRun(v bool) *CopyNetworkAclEntriesRequest {
	s.DryRun = &v
	return s
}

func (s *CopyNetworkAclEntriesRequest) SetNetworkAclId(v string) *CopyNetworkAclEntriesRequest {
	s.NetworkAclId = &v
	return s
}

func (s *CopyNetworkAclEntriesRequest) SetOwnerAccount(v string) *CopyNetworkAclEntriesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CopyNetworkAclEntriesRequest) SetOwnerId(v int64) *CopyNetworkAclEntriesRequest {
	s.OwnerId = &v
	return s
}

func (s *CopyNetworkAclEntriesRequest) SetRegionId(v string) *CopyNetworkAclEntriesRequest {
	s.RegionId = &v
	return s
}

func (s *CopyNetworkAclEntriesRequest) SetResourceOwnerAccount(v string) *CopyNetworkAclEntriesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CopyNetworkAclEntriesRequest) SetResourceOwnerId(v int64) *CopyNetworkAclEntriesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CopyNetworkAclEntriesRequest) SetSourceNetworkAclId(v string) *CopyNetworkAclEntriesRequest {
	s.SourceNetworkAclId = &v
	return s
}

func (s *CopyNetworkAclEntriesRequest) Validate() error {
	return dara.Validate(s)
}
