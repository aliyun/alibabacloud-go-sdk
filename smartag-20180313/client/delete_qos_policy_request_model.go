// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteQosPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DeleteQosPolicyRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteQosPolicyRequest
	GetOwnerId() *int64
	SetQosId(v string) *DeleteQosPolicyRequest
	GetQosId() *string
	SetQosPolicyId(v string) *DeleteQosPolicyRequest
	GetQosPolicyId() *string
	SetRegionId(v string) *DeleteQosPolicyRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteQosPolicyRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteQosPolicyRequest
	GetResourceOwnerId() *int64
}

type DeleteQosPolicyRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The instance ID of the QoS policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// qos-xitd8690ucu8ro****
	QosId *string `json:"QosId,omitempty" xml:"QosId,omitempty"`
	// The ID of the quintuple rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// qospy-xhwhyuo43l****
	QosPolicyId *string `json:"QosPolicyId,omitempty" xml:"QosPolicyId,omitempty"`
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

func (s DeleteQosPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteQosPolicyRequest) GoString() string {
	return s.String()
}

func (s *DeleteQosPolicyRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteQosPolicyRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteQosPolicyRequest) GetQosId() *string {
	return s.QosId
}

func (s *DeleteQosPolicyRequest) GetQosPolicyId() *string {
	return s.QosPolicyId
}

func (s *DeleteQosPolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteQosPolicyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteQosPolicyRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteQosPolicyRequest) SetOwnerAccount(v string) *DeleteQosPolicyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteQosPolicyRequest) SetOwnerId(v int64) *DeleteQosPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteQosPolicyRequest) SetQosId(v string) *DeleteQosPolicyRequest {
	s.QosId = &v
	return s
}

func (s *DeleteQosPolicyRequest) SetQosPolicyId(v string) *DeleteQosPolicyRequest {
	s.QosPolicyId = &v
	return s
}

func (s *DeleteQosPolicyRequest) SetRegionId(v string) *DeleteQosPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteQosPolicyRequest) SetResourceOwnerAccount(v string) *DeleteQosPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteQosPolicyRequest) SetResourceOwnerId(v int64) *DeleteQosPolicyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteQosPolicyRequest) Validate() error {
	return dara.Validate(s)
}
