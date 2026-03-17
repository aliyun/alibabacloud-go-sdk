// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateQosRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *CreateQosRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateQosRequest
	GetOwnerId() *int64
	SetQosDescription(v string) *CreateQosRequest
	GetQosDescription() *string
	SetQosName(v string) *CreateQosRequest
	GetQosName() *string
	SetRegionId(v string) *CreateQosRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateQosRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateQosRequest
	GetResourceOwnerId() *int64
}

type CreateQosRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The description of the QoS policy.
	//
	// The description must be 1 to 512 characters in length and can contain letters, digits, underscores (_), and hyphens (-). It must start with a letter.
	//
	// example:
	//
	// testdesc
	QosDescription *string `json:"QosDescription,omitempty" xml:"QosDescription,omitempty"`
	// The name of the QoS policy.
	//
	// The name must be 2 to 100 characters in length and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It must start with a letter.
	//
	// This parameter is required.
	//
	// example:
	//
	// doctest
	QosName *string `json:"QosName,omitempty" xml:"QosName,omitempty"`
	// The ID of the region where the QoS policy is deployed.
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

func (s CreateQosRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateQosRequest) GoString() string {
	return s.String()
}

func (s *CreateQosRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateQosRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateQosRequest) GetQosDescription() *string {
	return s.QosDescription
}

func (s *CreateQosRequest) GetQosName() *string {
	return s.QosName
}

func (s *CreateQosRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateQosRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateQosRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateQosRequest) SetOwnerAccount(v string) *CreateQosRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateQosRequest) SetOwnerId(v int64) *CreateQosRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateQosRequest) SetQosDescription(v string) *CreateQosRequest {
	s.QosDescription = &v
	return s
}

func (s *CreateQosRequest) SetQosName(v string) *CreateQosRequest {
	s.QosName = &v
	return s
}

func (s *CreateQosRequest) SetRegionId(v string) *CreateQosRequest {
	s.RegionId = &v
	return s
}

func (s *CreateQosRequest) SetResourceOwnerAccount(v string) *CreateQosRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateQosRequest) SetResourceOwnerId(v int64) *CreateQosRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateQosRequest) Validate() error {
	return dara.Validate(s)
}
