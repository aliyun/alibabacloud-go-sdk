// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySagRemoteAccessRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *ModifySagRemoteAccessRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifySagRemoteAccessRequest
	GetOwnerId() *int64
	SetRemoteAccessIp(v string) *ModifySagRemoteAccessRequest
	GetRemoteAccessIp() *string
	SetResourceOwnerAccount(v string) *ModifySagRemoteAccessRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifySagRemoteAccessRequest
	GetResourceOwnerId() *int64
	SetSerialNumber(v string) *ModifySagRemoteAccessRequest
	GetSerialNumber() *string
	SetSmartAGId(v string) *ModifySagRemoteAccessRequest
	GetSmartAGId() *string
}

type ModifySagRemoteAccessRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The remote access IP address of the SAG device.
	//
	// example:
	//
	// 10.10.10.2
	RemoteAccessIp       *string `json:"RemoteAccessIp,omitempty" xml:"RemoteAccessIp,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The serial number (SN) of the SAG device.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag-100-v1p7-9
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	// The ID of the corresponding SAG instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag-r79m060r6oy55******
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
}

func (s ModifySagRemoteAccessRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySagRemoteAccessRequest) GoString() string {
	return s.String()
}

func (s *ModifySagRemoteAccessRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifySagRemoteAccessRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifySagRemoteAccessRequest) GetRemoteAccessIp() *string {
	return s.RemoteAccessIp
}

func (s *ModifySagRemoteAccessRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifySagRemoteAccessRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifySagRemoteAccessRequest) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *ModifySagRemoteAccessRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *ModifySagRemoteAccessRequest) SetOwnerAccount(v string) *ModifySagRemoteAccessRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifySagRemoteAccessRequest) SetOwnerId(v int64) *ModifySagRemoteAccessRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifySagRemoteAccessRequest) SetRemoteAccessIp(v string) *ModifySagRemoteAccessRequest {
	s.RemoteAccessIp = &v
	return s
}

func (s *ModifySagRemoteAccessRequest) SetResourceOwnerAccount(v string) *ModifySagRemoteAccessRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifySagRemoteAccessRequest) SetResourceOwnerId(v int64) *ModifySagRemoteAccessRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifySagRemoteAccessRequest) SetSerialNumber(v string) *ModifySagRemoteAccessRequest {
	s.SerialNumber = &v
	return s
}

func (s *ModifySagRemoteAccessRequest) SetSmartAGId(v string) *ModifySagRemoteAccessRequest {
	s.SmartAGId = &v
	return s
}

func (s *ModifySagRemoteAccessRequest) Validate() error {
	return dara.Validate(s)
}
