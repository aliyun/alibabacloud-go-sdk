// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteQosCarRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DeleteQosCarRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteQosCarRequest
	GetOwnerId() *int64
	SetQosCarId(v string) *DeleteQosCarRequest
	GetQosCarId() *string
	SetQosId(v string) *DeleteQosCarRequest
	GetQosId() *string
	SetRegionId(v string) *DeleteQosCarRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteQosCarRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteQosCarRequest
	GetResourceOwnerId() *int64
}

type DeleteQosCarRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The instance ID of the QoS speed limiting rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// qoscar-n5k8g97lihlph****
	QosCarId *string `json:"QosCarId,omitempty" xml:"QosCarId,omitempty"`
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
}

func (s DeleteQosCarRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteQosCarRequest) GoString() string {
	return s.String()
}

func (s *DeleteQosCarRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteQosCarRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteQosCarRequest) GetQosCarId() *string {
	return s.QosCarId
}

func (s *DeleteQosCarRequest) GetQosId() *string {
	return s.QosId
}

func (s *DeleteQosCarRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteQosCarRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteQosCarRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteQosCarRequest) SetOwnerAccount(v string) *DeleteQosCarRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteQosCarRequest) SetOwnerId(v int64) *DeleteQosCarRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteQosCarRequest) SetQosCarId(v string) *DeleteQosCarRequest {
	s.QosCarId = &v
	return s
}

func (s *DeleteQosCarRequest) SetQosId(v string) *DeleteQosCarRequest {
	s.QosId = &v
	return s
}

func (s *DeleteQosCarRequest) SetRegionId(v string) *DeleteQosCarRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteQosCarRequest) SetResourceOwnerAccount(v string) *DeleteQosCarRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteQosCarRequest) SetResourceOwnerId(v int64) *DeleteQosCarRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteQosCarRequest) Validate() error {
	return dara.Validate(s)
}
