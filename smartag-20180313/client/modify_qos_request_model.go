// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyQosRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *ModifyQosRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyQosRequest
	GetOwnerId() *int64
	SetQosDescription(v string) *ModifyQosRequest
	GetQosDescription() *string
	SetQosId(v string) *ModifyQosRequest
	GetQosId() *string
	SetQosName(v string) *ModifyQosRequest
	GetQosName() *string
	SetRegionId(v string) *ModifyQosRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyQosRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyQosRequest
	GetResourceOwnerId() *int64
}

type ModifyQosRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The description of the QoS policy. The description must be 1 to 512 characters in length, and can contain digits, underscores (_), and hyphens (-). It must start with a letter or Chinese character.
	//
	// example:
	//
	// qosdes
	QosDescription *string `json:"QosDescription,omitempty" xml:"QosDescription,omitempty"`
	// The ID of the QoS policy that you want to modify.
	//
	// This parameter is required.
	//
	// example:
	//
	// qos-awfxl1adxeqyk****
	QosId *string `json:"QosId,omitempty" xml:"QosId,omitempty"`
	// The name of the QoS policy.
	//
	// example:
	//
	// doctest
	QosName *string `json:"QosName,omitempty" xml:"QosName,omitempty"`
	// The region where the QoS policy is deployed.
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

func (s ModifyQosRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyQosRequest) GoString() string {
	return s.String()
}

func (s *ModifyQosRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyQosRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyQosRequest) GetQosDescription() *string {
	return s.QosDescription
}

func (s *ModifyQosRequest) GetQosId() *string {
	return s.QosId
}

func (s *ModifyQosRequest) GetQosName() *string {
	return s.QosName
}

func (s *ModifyQosRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyQosRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyQosRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyQosRequest) SetOwnerAccount(v string) *ModifyQosRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyQosRequest) SetOwnerId(v int64) *ModifyQosRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyQosRequest) SetQosDescription(v string) *ModifyQosRequest {
	s.QosDescription = &v
	return s
}

func (s *ModifyQosRequest) SetQosId(v string) *ModifyQosRequest {
	s.QosId = &v
	return s
}

func (s *ModifyQosRequest) SetQosName(v string) *ModifyQosRequest {
	s.QosName = &v
	return s
}

func (s *ModifyQosRequest) SetRegionId(v string) *ModifyQosRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyQosRequest) SetResourceOwnerAccount(v string) *ModifyQosRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyQosRequest) SetResourceOwnerId(v int64) *ModifyQosRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyQosRequest) Validate() error {
	return dara.Validate(s)
}
