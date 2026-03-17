// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeInstanceFromCbnRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCcnInstanceId(v string) *RevokeInstanceFromCbnRequest
	GetCcnInstanceId() *string
	SetCenInstanceId(v string) *RevokeInstanceFromCbnRequest
	GetCenInstanceId() *string
	SetOwnerAccount(v string) *RevokeInstanceFromCbnRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *RevokeInstanceFromCbnRequest
	GetOwnerId() *int64
	SetRegionId(v string) *RevokeInstanceFromCbnRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *RevokeInstanceFromCbnRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RevokeInstanceFromCbnRequest
	GetResourceOwnerId() *int64
}

type RevokeInstanceFromCbnRequest struct {
	// The ID of the CCN instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ccn-n2935s1mnwv8i*****
	CcnInstanceId *string `json:"CcnInstanceId,omitempty" xml:"CcnInstanceId,omitempty"`
	// The ID of the CEN instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cen-7qthudw0ll6jm*****
	CenInstanceId *string `json:"CenInstanceId,omitempty" xml:"CenInstanceId,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the CCN instance is deployed.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s RevokeInstanceFromCbnRequest) String() string {
	return dara.Prettify(s)
}

func (s RevokeInstanceFromCbnRequest) GoString() string {
	return s.String()
}

func (s *RevokeInstanceFromCbnRequest) GetCcnInstanceId() *string {
	return s.CcnInstanceId
}

func (s *RevokeInstanceFromCbnRequest) GetCenInstanceId() *string {
	return s.CenInstanceId
}

func (s *RevokeInstanceFromCbnRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *RevokeInstanceFromCbnRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RevokeInstanceFromCbnRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RevokeInstanceFromCbnRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RevokeInstanceFromCbnRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RevokeInstanceFromCbnRequest) SetCcnInstanceId(v string) *RevokeInstanceFromCbnRequest {
	s.CcnInstanceId = &v
	return s
}

func (s *RevokeInstanceFromCbnRequest) SetCenInstanceId(v string) *RevokeInstanceFromCbnRequest {
	s.CenInstanceId = &v
	return s
}

func (s *RevokeInstanceFromCbnRequest) SetOwnerAccount(v string) *RevokeInstanceFromCbnRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RevokeInstanceFromCbnRequest) SetOwnerId(v int64) *RevokeInstanceFromCbnRequest {
	s.OwnerId = &v
	return s
}

func (s *RevokeInstanceFromCbnRequest) SetRegionId(v string) *RevokeInstanceFromCbnRequest {
	s.RegionId = &v
	return s
}

func (s *RevokeInstanceFromCbnRequest) SetResourceOwnerAccount(v string) *RevokeInstanceFromCbnRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RevokeInstanceFromCbnRequest) SetResourceOwnerId(v int64) *RevokeInstanceFromCbnRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RevokeInstanceFromCbnRequest) Validate() error {
	return dara.Validate(s)
}
