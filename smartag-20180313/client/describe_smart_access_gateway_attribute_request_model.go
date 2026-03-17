// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSmartAccessGatewayAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribeSmartAccessGatewayAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeSmartAccessGatewayAttributeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeSmartAccessGatewayAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeSmartAccessGatewayAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeSmartAccessGatewayAttributeRequest
	GetResourceOwnerId() *int64
	SetSmartAGId(v string) *DescribeSmartAccessGatewayAttributeRequest
	GetSmartAGId() *string
}

type DescribeSmartAccessGatewayAttributeRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the SAG instance is deployed.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the SAG instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag-6z21oj0vjjrx****
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
}

func (s DescribeSmartAccessGatewayAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSmartAccessGatewayAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeSmartAccessGatewayAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeSmartAccessGatewayAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeSmartAccessGatewayAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSmartAccessGatewayAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeSmartAccessGatewayAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSmartAccessGatewayAttributeRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *DescribeSmartAccessGatewayAttributeRequest) SetOwnerAccount(v string) *DescribeSmartAccessGatewayAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSmartAccessGatewayAttributeRequest) SetOwnerId(v int64) *DescribeSmartAccessGatewayAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSmartAccessGatewayAttributeRequest) SetRegionId(v string) *DescribeSmartAccessGatewayAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSmartAccessGatewayAttributeRequest) SetResourceOwnerAccount(v string) *DescribeSmartAccessGatewayAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSmartAccessGatewayAttributeRequest) SetResourceOwnerId(v int64) *DescribeSmartAccessGatewayAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSmartAccessGatewayAttributeRequest) SetSmartAGId(v string) *DescribeSmartAccessGatewayAttributeRequest {
	s.SmartAGId = &v
	return s
}

func (s *DescribeSmartAccessGatewayAttributeRequest) Validate() error {
	return dara.Validate(s)
}
