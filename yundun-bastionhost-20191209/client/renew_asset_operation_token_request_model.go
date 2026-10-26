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
	SetToken(v string) *RenewAssetOperationTokenRequest
	GetToken() *string
	SetTokenId(v string) *RenewAssetOperationTokenRequest
	GetTokenId() *string
}

type RenewAssetOperationTokenRequest struct {
	// The ID of the Bastionhost instance.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the Bastionhost instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the Bastionhost instance.
	//
	// > For the mappings between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The O&M token to be renewed.
	//
	// example:
	//
	// NmYyMmEzNmMwYzljNGY******
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
	// The ID of the O&M token to be renewed.
	//
	// > You can call the [GenerateAssetOperationToken](https://help.aliyun.com/document_detail/2758861.html) operation to obtain this parameter.
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

func (s *RenewAssetOperationTokenRequest) GetToken() *string {
	return s.Token
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

func (s *RenewAssetOperationTokenRequest) SetToken(v string) *RenewAssetOperationTokenRequest {
	s.Token = &v
	return s
}

func (s *RenewAssetOperationTokenRequest) SetTokenId(v string) *RenewAssetOperationTokenRequest {
	s.TokenId = &v
	return s
}

func (s *RenewAssetOperationTokenRequest) Validate() error {
	return dara.Validate(s)
}
