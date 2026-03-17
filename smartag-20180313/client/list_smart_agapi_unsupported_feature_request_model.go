// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSmartAGApiUnsupportedFeatureRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpenApiName(v string) *ListSmartAGApiUnsupportedFeatureRequest
	GetOpenApiName() *string
	SetOwnerAccount(v string) *ListSmartAGApiUnsupportedFeatureRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListSmartAGApiUnsupportedFeatureRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ListSmartAGApiUnsupportedFeatureRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ListSmartAGApiUnsupportedFeatureRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListSmartAGApiUnsupportedFeatureRequest
	GetResourceOwnerId() *int64
	SetSerialNumber(v string) *ListSmartAGApiUnsupportedFeatureRequest
	GetSerialNumber() *string
	SetSmartAGId(v string) *ListSmartAGApiUnsupportedFeatureRequest
	GetSmartAGId() *string
}

type ListSmartAGApiUnsupportedFeatureRequest struct {
	// The API operation for the unsupported feature.
	//
	// This parameter is required.
	//
	// example:
	//
	// ModifySagWan
	OpenApiName  *string `json:"OpenApiName,omitempty" xml:"OpenApiName,omitempty"`
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
	// The serial number of the SAG device.
	//
	// This parameter is required.
	//
	// example:
	//
	// sage62x052614****
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	// The ID of the SAG instance with which the SAG device is associated.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag-4d6i45zess8nj4****
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
}

func (s ListSmartAGApiUnsupportedFeatureRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSmartAGApiUnsupportedFeatureRequest) GoString() string {
	return s.String()
}

func (s *ListSmartAGApiUnsupportedFeatureRequest) GetOpenApiName() *string {
	return s.OpenApiName
}

func (s *ListSmartAGApiUnsupportedFeatureRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListSmartAGApiUnsupportedFeatureRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListSmartAGApiUnsupportedFeatureRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListSmartAGApiUnsupportedFeatureRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListSmartAGApiUnsupportedFeatureRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListSmartAGApiUnsupportedFeatureRequest) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *ListSmartAGApiUnsupportedFeatureRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *ListSmartAGApiUnsupportedFeatureRequest) SetOpenApiName(v string) *ListSmartAGApiUnsupportedFeatureRequest {
	s.OpenApiName = &v
	return s
}

func (s *ListSmartAGApiUnsupportedFeatureRequest) SetOwnerAccount(v string) *ListSmartAGApiUnsupportedFeatureRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListSmartAGApiUnsupportedFeatureRequest) SetOwnerId(v int64) *ListSmartAGApiUnsupportedFeatureRequest {
	s.OwnerId = &v
	return s
}

func (s *ListSmartAGApiUnsupportedFeatureRequest) SetRegionId(v string) *ListSmartAGApiUnsupportedFeatureRequest {
	s.RegionId = &v
	return s
}

func (s *ListSmartAGApiUnsupportedFeatureRequest) SetResourceOwnerAccount(v string) *ListSmartAGApiUnsupportedFeatureRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListSmartAGApiUnsupportedFeatureRequest) SetResourceOwnerId(v int64) *ListSmartAGApiUnsupportedFeatureRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListSmartAGApiUnsupportedFeatureRequest) SetSerialNumber(v string) *ListSmartAGApiUnsupportedFeatureRequest {
	s.SerialNumber = &v
	return s
}

func (s *ListSmartAGApiUnsupportedFeatureRequest) SetSmartAGId(v string) *ListSmartAGApiUnsupportedFeatureRequest {
	s.SmartAGId = &v
	return s
}

func (s *ListSmartAGApiUnsupportedFeatureRequest) Validate() error {
	return dara.Validate(s)
}
