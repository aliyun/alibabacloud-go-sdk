// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisassociateQosRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DisassociateQosRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DisassociateQosRequest
	GetOwnerId() *int64
	SetQosId(v string) *DisassociateQosRequest
	GetQosId() *string
	SetRegionId(v string) *DisassociateQosRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DisassociateQosRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DisassociateQosRequest
	GetResourceOwnerId() *int64
	SetSmartAGId(v string) *DisassociateQosRequest
	GetSmartAGId() *string
}

type DisassociateQosRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The instance ID of the QoS policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// qos-1lcl9gv5ew7x****
	QosId *string `json:"QosId,omitempty" xml:"QosId,omitempty"`
	// The ID of the region to which the QoS policy belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the SAG instance to be disassociated.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag-c3m3n1khz58l****
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
}

func (s DisassociateQosRequest) String() string {
	return dara.Prettify(s)
}

func (s DisassociateQosRequest) GoString() string {
	return s.String()
}

func (s *DisassociateQosRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DisassociateQosRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DisassociateQosRequest) GetQosId() *string {
	return s.QosId
}

func (s *DisassociateQosRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DisassociateQosRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DisassociateQosRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DisassociateQosRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *DisassociateQosRequest) SetOwnerAccount(v string) *DisassociateQosRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DisassociateQosRequest) SetOwnerId(v int64) *DisassociateQosRequest {
	s.OwnerId = &v
	return s
}

func (s *DisassociateQosRequest) SetQosId(v string) *DisassociateQosRequest {
	s.QosId = &v
	return s
}

func (s *DisassociateQosRequest) SetRegionId(v string) *DisassociateQosRequest {
	s.RegionId = &v
	return s
}

func (s *DisassociateQosRequest) SetResourceOwnerAccount(v string) *DisassociateQosRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DisassociateQosRequest) SetResourceOwnerId(v int64) *DisassociateQosRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DisassociateQosRequest) SetSmartAGId(v string) *DisassociateQosRequest {
	s.SmartAGId = &v
	return s
}

func (s *DisassociateQosRequest) Validate() error {
	return dara.Validate(s)
}
