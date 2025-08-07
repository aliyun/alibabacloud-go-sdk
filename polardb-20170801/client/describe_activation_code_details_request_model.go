// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeActivationCodeDetailsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActivationCodeId(v int32) *DescribeActivationCodeDetailsRequest
	GetActivationCodeId() *int32
	SetAliyunOrderId(v string) *DescribeActivationCodeDetailsRequest
	GetAliyunOrderId() *string
	SetOwnerAccount(v string) *DescribeActivationCodeDetailsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeActivationCodeDetailsRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeActivationCodeDetailsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeActivationCodeDetailsRequest
	GetResourceOwnerId() *int64
}

type DescribeActivationCodeDetailsRequest struct {
	// The ID of the activation code.
	//
	// example:
	//
	// 123
	ActivationCodeId *int32 `json:"ActivationCodeId,omitempty" xml:"ActivationCodeId,omitempty"`
	// The Alibaba Cloud order ID (including the virtual order ID).
	//
	// This parameter is required.
	//
	// example:
	//
	// 2233****445566
	AliyunOrderId        *string `json:"AliyunOrderId,omitempty" xml:"AliyunOrderId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeActivationCodeDetailsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeActivationCodeDetailsRequest) GoString() string {
	return s.String()
}

func (s *DescribeActivationCodeDetailsRequest) GetActivationCodeId() *int32 {
	return s.ActivationCodeId
}

func (s *DescribeActivationCodeDetailsRequest) GetAliyunOrderId() *string {
	return s.AliyunOrderId
}

func (s *DescribeActivationCodeDetailsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeActivationCodeDetailsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeActivationCodeDetailsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeActivationCodeDetailsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeActivationCodeDetailsRequest) SetActivationCodeId(v int32) *DescribeActivationCodeDetailsRequest {
	s.ActivationCodeId = &v
	return s
}

func (s *DescribeActivationCodeDetailsRequest) SetAliyunOrderId(v string) *DescribeActivationCodeDetailsRequest {
	s.AliyunOrderId = &v
	return s
}

func (s *DescribeActivationCodeDetailsRequest) SetOwnerAccount(v string) *DescribeActivationCodeDetailsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeActivationCodeDetailsRequest) SetOwnerId(v int64) *DescribeActivationCodeDetailsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeActivationCodeDetailsRequest) SetResourceOwnerAccount(v string) *DescribeActivationCodeDetailsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeActivationCodeDetailsRequest) SetResourceOwnerId(v int64) *DescribeActivationCodeDetailsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeActivationCodeDetailsRequest) Validate() error {
	return dara.Validate(s)
}
