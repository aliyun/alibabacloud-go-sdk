// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeSagInstanceFromCcnRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCcnInstanceId(v string) *RevokeSagInstanceFromCcnRequest
	GetCcnInstanceId() *string
	SetOwnerAccount(v string) *RevokeSagInstanceFromCcnRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *RevokeSagInstanceFromCcnRequest
	GetOwnerId() *int64
	SetRegionId(v string) *RevokeSagInstanceFromCcnRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *RevokeSagInstanceFromCcnRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RevokeSagInstanceFromCcnRequest
	GetResourceOwnerId() *int64
	SetSmartAGId(v string) *RevokeSagInstanceFromCcnRequest
	GetSmartAGId() *string
}

type RevokeSagInstanceFromCcnRequest struct {
	// The ID of the CCN instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ccn-jf********
	CcnInstanceId *string `json:"CcnInstanceId,omitempty" xml:"CcnInstanceId,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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
	// sag-hd**************
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
}

func (s RevokeSagInstanceFromCcnRequest) String() string {
	return dara.Prettify(s)
}

func (s RevokeSagInstanceFromCcnRequest) GoString() string {
	return s.String()
}

func (s *RevokeSagInstanceFromCcnRequest) GetCcnInstanceId() *string {
	return s.CcnInstanceId
}

func (s *RevokeSagInstanceFromCcnRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *RevokeSagInstanceFromCcnRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RevokeSagInstanceFromCcnRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RevokeSagInstanceFromCcnRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RevokeSagInstanceFromCcnRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RevokeSagInstanceFromCcnRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *RevokeSagInstanceFromCcnRequest) SetCcnInstanceId(v string) *RevokeSagInstanceFromCcnRequest {
	s.CcnInstanceId = &v
	return s
}

func (s *RevokeSagInstanceFromCcnRequest) SetOwnerAccount(v string) *RevokeSagInstanceFromCcnRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RevokeSagInstanceFromCcnRequest) SetOwnerId(v int64) *RevokeSagInstanceFromCcnRequest {
	s.OwnerId = &v
	return s
}

func (s *RevokeSagInstanceFromCcnRequest) SetRegionId(v string) *RevokeSagInstanceFromCcnRequest {
	s.RegionId = &v
	return s
}

func (s *RevokeSagInstanceFromCcnRequest) SetResourceOwnerAccount(v string) *RevokeSagInstanceFromCcnRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RevokeSagInstanceFromCcnRequest) SetResourceOwnerId(v int64) *RevokeSagInstanceFromCcnRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RevokeSagInstanceFromCcnRequest) SetSmartAGId(v string) *RevokeSagInstanceFromCcnRequest {
	s.SmartAGId = &v
	return s
}

func (s *RevokeSagInstanceFromCcnRequest) Validate() error {
	return dara.Validate(s)
}
