// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceMultiVIPRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddCount(v string) *CreateInstanceMultiVIPRequest
	GetAddCount() *string
	SetInstanceId(v string) *CreateInstanceMultiVIPRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *CreateInstanceMultiVIPRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateInstanceMultiVIPRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CreateInstanceMultiVIPRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateInstanceMultiVIPRequest
	GetResourceOwnerId() *int64
}

type CreateInstanceMultiVIPRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	AddCount *string `json:"AddCount,omitempty" xml:"AddCount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateInstanceMultiVIPRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceMultiVIPRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceMultiVIPRequest) GetAddCount() *string {
	return s.AddCount
}

func (s *CreateInstanceMultiVIPRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateInstanceMultiVIPRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateInstanceMultiVIPRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateInstanceMultiVIPRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateInstanceMultiVIPRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateInstanceMultiVIPRequest) SetAddCount(v string) *CreateInstanceMultiVIPRequest {
	s.AddCount = &v
	return s
}

func (s *CreateInstanceMultiVIPRequest) SetInstanceId(v string) *CreateInstanceMultiVIPRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateInstanceMultiVIPRequest) SetOwnerAccount(v string) *CreateInstanceMultiVIPRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateInstanceMultiVIPRequest) SetOwnerId(v int64) *CreateInstanceMultiVIPRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateInstanceMultiVIPRequest) SetResourceOwnerAccount(v string) *CreateInstanceMultiVIPRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateInstanceMultiVIPRequest) SetResourceOwnerId(v int64) *CreateInstanceMultiVIPRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateInstanceMultiVIPRequest) Validate() error {
	return dara.Validate(s)
}
