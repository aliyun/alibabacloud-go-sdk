// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateQosRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *AssociateQosRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AssociateQosRequest
	GetOwnerId() *int64
	SetQosId(v string) *AssociateQosRequest
	GetQosId() *string
	SetRegionId(v string) *AssociateQosRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *AssociateQosRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AssociateQosRequest
	GetResourceOwnerId() *int64
	SetSmartAGId(v string) *AssociateQosRequest
	GetSmartAGId() *string
}

type AssociateQosRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The instance ID of the QoS policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// qos-awfxl1adxeqyk****
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
	// The ID of the SAG instance to which the QoS policy is to be applied.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag-c3m3n1khz58l****
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
}

func (s AssociateQosRequest) String() string {
	return dara.Prettify(s)
}

func (s AssociateQosRequest) GoString() string {
	return s.String()
}

func (s *AssociateQosRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AssociateQosRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AssociateQosRequest) GetQosId() *string {
	return s.QosId
}

func (s *AssociateQosRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AssociateQosRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AssociateQosRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AssociateQosRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *AssociateQosRequest) SetOwnerAccount(v string) *AssociateQosRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AssociateQosRequest) SetOwnerId(v int64) *AssociateQosRequest {
	s.OwnerId = &v
	return s
}

func (s *AssociateQosRequest) SetQosId(v string) *AssociateQosRequest {
	s.QosId = &v
	return s
}

func (s *AssociateQosRequest) SetRegionId(v string) *AssociateQosRequest {
	s.RegionId = &v
	return s
}

func (s *AssociateQosRequest) SetResourceOwnerAccount(v string) *AssociateQosRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AssociateQosRequest) SetResourceOwnerId(v int64) *AssociateQosRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AssociateQosRequest) SetSmartAGId(v string) *AssociateQosRequest {
	s.SmartAGId = &v
	return s
}

func (s *AssociateQosRequest) Validate() error {
	return dara.Validate(s)
}
