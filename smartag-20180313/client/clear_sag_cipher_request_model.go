// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClearSagCipherRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *ClearSagCipherRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ClearSagCipherRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ClearSagCipherRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ClearSagCipherRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ClearSagCipherRequest
	GetResourceOwnerId() *int64
	SetSagId(v string) *ClearSagCipherRequest
	GetSagId() *string
	SetSnNumber(v string) *ClearSagCipherRequest
	GetSnNumber() *string
}

type ClearSagCipherRequest struct {
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
	// sag-0nnteglltw6z4b****
	SagId *string `json:"SagId,omitempty" xml:"SagId,omitempty"`
	// The serial number of the SAG vCPE device.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag42c3****
	SnNumber *string `json:"SnNumber,omitempty" xml:"SnNumber,omitempty"`
}

func (s ClearSagCipherRequest) String() string {
	return dara.Prettify(s)
}

func (s ClearSagCipherRequest) GoString() string {
	return s.String()
}

func (s *ClearSagCipherRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ClearSagCipherRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ClearSagCipherRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ClearSagCipherRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ClearSagCipherRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ClearSagCipherRequest) GetSagId() *string {
	return s.SagId
}

func (s *ClearSagCipherRequest) GetSnNumber() *string {
	return s.SnNumber
}

func (s *ClearSagCipherRequest) SetOwnerAccount(v string) *ClearSagCipherRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ClearSagCipherRequest) SetOwnerId(v int64) *ClearSagCipherRequest {
	s.OwnerId = &v
	return s
}

func (s *ClearSagCipherRequest) SetRegionId(v string) *ClearSagCipherRequest {
	s.RegionId = &v
	return s
}

func (s *ClearSagCipherRequest) SetResourceOwnerAccount(v string) *ClearSagCipherRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ClearSagCipherRequest) SetResourceOwnerId(v int64) *ClearSagCipherRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ClearSagCipherRequest) SetSagId(v string) *ClearSagCipherRequest {
	s.SagId = &v
	return s
}

func (s *ClearSagCipherRequest) SetSnNumber(v string) *ClearSagCipherRequest {
	s.SnNumber = &v
	return s
}

func (s *ClearSagCipherRequest) Validate() error {
	return dara.Validate(s)
}
