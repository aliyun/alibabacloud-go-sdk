// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebootSmartAccessGatewayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *RebootSmartAccessGatewayRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *RebootSmartAccessGatewayRequest
	GetOwnerId() *int64
	SetRegionId(v string) *RebootSmartAccessGatewayRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *RebootSmartAccessGatewayRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RebootSmartAccessGatewayRequest
	GetResourceOwnerId() *int64
	SetSerialNumber(v string) *RebootSmartAccessGatewayRequest
	GetSerialNumber() *string
	SetSmartAGId(v string) *RebootSmartAccessGatewayRequest
	GetSmartAGId() *string
}

type RebootSmartAccessGatewayRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region to which the SAG instance belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The serial number of the SAG device.
	//
	// example:
	//
	// a1b2c3d4e5f6g7h8f9
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	// The ID of the SAG instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag-w9unmktmupcde*****
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
}

func (s RebootSmartAccessGatewayRequest) String() string {
	return dara.Prettify(s)
}

func (s RebootSmartAccessGatewayRequest) GoString() string {
	return s.String()
}

func (s *RebootSmartAccessGatewayRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *RebootSmartAccessGatewayRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RebootSmartAccessGatewayRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RebootSmartAccessGatewayRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RebootSmartAccessGatewayRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RebootSmartAccessGatewayRequest) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *RebootSmartAccessGatewayRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *RebootSmartAccessGatewayRequest) SetOwnerAccount(v string) *RebootSmartAccessGatewayRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RebootSmartAccessGatewayRequest) SetOwnerId(v int64) *RebootSmartAccessGatewayRequest {
	s.OwnerId = &v
	return s
}

func (s *RebootSmartAccessGatewayRequest) SetRegionId(v string) *RebootSmartAccessGatewayRequest {
	s.RegionId = &v
	return s
}

func (s *RebootSmartAccessGatewayRequest) SetResourceOwnerAccount(v string) *RebootSmartAccessGatewayRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RebootSmartAccessGatewayRequest) SetResourceOwnerId(v int64) *RebootSmartAccessGatewayRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RebootSmartAccessGatewayRequest) SetSerialNumber(v string) *RebootSmartAccessGatewayRequest {
	s.SerialNumber = &v
	return s
}

func (s *RebootSmartAccessGatewayRequest) SetSmartAGId(v string) *RebootSmartAccessGatewayRequest {
	s.SmartAGId = &v
	return s
}

func (s *RebootSmartAccessGatewayRequest) Validate() error {
	return dara.Validate(s)
}
