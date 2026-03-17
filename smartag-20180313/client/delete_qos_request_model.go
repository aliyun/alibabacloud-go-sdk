// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteQosRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DeleteQosRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteQosRequest
	GetOwnerId() *int64
	SetQosId(v string) *DeleteQosRequest
	GetQosId() *string
	SetRegionId(v string) *DeleteQosRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteQosRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteQosRequest
	GetResourceOwnerId() *int64
}

type DeleteQosRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The instance ID of the QoS policy to be deleted.
	//
	// This parameter is required.
	//
	// example:
	//
	// qos-xhwhyuo43l****
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

func (s DeleteQosRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteQosRequest) GoString() string {
	return s.String()
}

func (s *DeleteQosRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteQosRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteQosRequest) GetQosId() *string {
	return s.QosId
}

func (s *DeleteQosRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteQosRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteQosRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteQosRequest) SetOwnerAccount(v string) *DeleteQosRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteQosRequest) SetOwnerId(v int64) *DeleteQosRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteQosRequest) SetQosId(v string) *DeleteQosRequest {
	s.QosId = &v
	return s
}

func (s *DeleteQosRequest) SetRegionId(v string) *DeleteQosRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteQosRequest) SetResourceOwnerAccount(v string) *DeleteQosRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteQosRequest) SetResourceOwnerId(v int64) *DeleteQosRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteQosRequest) Validate() error {
	return dara.Validate(s)
}
