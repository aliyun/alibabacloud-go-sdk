// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewAssetOperationTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *RenewAssetOperationTokenRequest
	GetInstanceId() *string
	SetRegionId(v string) *RenewAssetOperationTokenRequest
	GetRegionId() *string
	SetTokenId(v string) *RenewAssetOperationTokenRequest
	GetTokenId() *string
}

type RenewAssetOperationTokenRequest struct {
	// The bastion host ID.
	//
	// >  You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the bastion host ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the O\\&M token that you want to renew.
	//
	// >  You can call the [GenerateAssetOperationToken](https://help.aliyun.com/document_detail/2758861.html) operation to query the O\\&M token ID.
	//
	// example:
	//
	// NmYyMmEzNmMwYzljNGYxMjh******
	TokenId *string `json:"TokenId,omitempty" xml:"TokenId,omitempty"`
}

func (s RenewAssetOperationTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s RenewAssetOperationTokenRequest) GoString() string {
	return s.String()
}

func (s *RenewAssetOperationTokenRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RenewAssetOperationTokenRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RenewAssetOperationTokenRequest) GetTokenId() *string {
	return s.TokenId
}

func (s *RenewAssetOperationTokenRequest) SetInstanceId(v string) *RenewAssetOperationTokenRequest {
	s.InstanceId = &v
	return s
}

func (s *RenewAssetOperationTokenRequest) SetRegionId(v string) *RenewAssetOperationTokenRequest {
	s.RegionId = &v
	return s
}

func (s *RenewAssetOperationTokenRequest) SetTokenId(v string) *RenewAssetOperationTokenRequest {
	s.TokenId = &v
	return s
}

func (s *RenewAssetOperationTokenRequest) Validate() error {
	return dara.Validate(s)
}
