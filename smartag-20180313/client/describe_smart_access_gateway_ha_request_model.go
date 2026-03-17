// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSmartAccessGatewayHaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribeSmartAccessGatewayHaRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeSmartAccessGatewayHaRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeSmartAccessGatewayHaRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeSmartAccessGatewayHaRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeSmartAccessGatewayHaRequest
	GetResourceOwnerId() *int64
	SetSmartAGId(v string) *DescribeSmartAccessGatewayHaRequest
	GetSmartAGId() *string
}

type DescribeSmartAccessGatewayHaRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the SAG instance is deployed.
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
	// sag-pno62188piyc6txxxxx
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
}

func (s DescribeSmartAccessGatewayHaRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSmartAccessGatewayHaRequest) GoString() string {
	return s.String()
}

func (s *DescribeSmartAccessGatewayHaRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeSmartAccessGatewayHaRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeSmartAccessGatewayHaRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSmartAccessGatewayHaRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeSmartAccessGatewayHaRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSmartAccessGatewayHaRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *DescribeSmartAccessGatewayHaRequest) SetOwnerAccount(v string) *DescribeSmartAccessGatewayHaRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSmartAccessGatewayHaRequest) SetOwnerId(v int64) *DescribeSmartAccessGatewayHaRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSmartAccessGatewayHaRequest) SetRegionId(v string) *DescribeSmartAccessGatewayHaRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSmartAccessGatewayHaRequest) SetResourceOwnerAccount(v string) *DescribeSmartAccessGatewayHaRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSmartAccessGatewayHaRequest) SetResourceOwnerId(v int64) *DescribeSmartAccessGatewayHaRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSmartAccessGatewayHaRequest) SetSmartAGId(v string) *DescribeSmartAccessGatewayHaRequest {
	s.SmartAGId = &v
	return s
}

func (s *DescribeSmartAccessGatewayHaRequest) Validate() error {
	return dara.Validate(s)
}
