// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSmartAccessGatewayVersionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribeSmartAccessGatewayVersionsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeSmartAccessGatewayVersionsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeSmartAccessGatewayVersionsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeSmartAccessGatewayVersionsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeSmartAccessGatewayVersionsRequest
	GetResourceOwnerId() *int64
	SetSmartAGId(v string) *DescribeSmartAccessGatewayVersionsRequest
	GetSmartAGId() *string
	SetSmartAGSn(v string) *DescribeSmartAccessGatewayVersionsRequest
	GetSmartAGSn() *string
	SetVersionType(v string) *DescribeSmartAccessGatewayVersionsRequest
	GetVersionType() *string
}

type DescribeSmartAccessGatewayVersionsRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the SAG instance is deployed.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the SAG instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag-d3m51apgw4po5*****
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
	// The serial number of the SAG device.
	//
	// example:
	//
	// sage62x022502****
	SmartAGSn *string `json:"SmartAGSn,omitempty" xml:"SmartAGSn,omitempty"`
	// The type of software run by the SAG device. Valid values:
	//
	// 	- **Device**: The operating system run by the SAG device. This is the default value.
	//
	// 	- **Dpi**: The signature database used by the SAG device.
	//
	// example:
	//
	// Device
	VersionType *string `json:"VersionType,omitempty" xml:"VersionType,omitempty"`
}

func (s DescribeSmartAccessGatewayVersionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSmartAccessGatewayVersionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSmartAccessGatewayVersionsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeSmartAccessGatewayVersionsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeSmartAccessGatewayVersionsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSmartAccessGatewayVersionsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeSmartAccessGatewayVersionsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSmartAccessGatewayVersionsRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *DescribeSmartAccessGatewayVersionsRequest) GetSmartAGSn() *string {
	return s.SmartAGSn
}

func (s *DescribeSmartAccessGatewayVersionsRequest) GetVersionType() *string {
	return s.VersionType
}

func (s *DescribeSmartAccessGatewayVersionsRequest) SetOwnerAccount(v string) *DescribeSmartAccessGatewayVersionsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSmartAccessGatewayVersionsRequest) SetOwnerId(v int64) *DescribeSmartAccessGatewayVersionsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSmartAccessGatewayVersionsRequest) SetRegionId(v string) *DescribeSmartAccessGatewayVersionsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSmartAccessGatewayVersionsRequest) SetResourceOwnerAccount(v string) *DescribeSmartAccessGatewayVersionsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSmartAccessGatewayVersionsRequest) SetResourceOwnerId(v int64) *DescribeSmartAccessGatewayVersionsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSmartAccessGatewayVersionsRequest) SetSmartAGId(v string) *DescribeSmartAccessGatewayVersionsRequest {
	s.SmartAGId = &v
	return s
}

func (s *DescribeSmartAccessGatewayVersionsRequest) SetSmartAGSn(v string) *DescribeSmartAccessGatewayVersionsRequest {
	s.SmartAGSn = &v
	return s
}

func (s *DescribeSmartAccessGatewayVersionsRequest) SetVersionType(v string) *DescribeSmartAccessGatewayVersionsRequest {
	s.VersionType = &v
	return s
}

func (s *DescribeSmartAccessGatewayVersionsRequest) Validate() error {
	return dara.Validate(s)
}
