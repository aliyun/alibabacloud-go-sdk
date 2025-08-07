// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLicenseOrderDetailsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunOrderId(v string) *DescribeLicenseOrderDetailsRequest
	GetAliyunOrderId() *string
	SetOwnerAccount(v string) *DescribeLicenseOrderDetailsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeLicenseOrderDetailsRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeLicenseOrderDetailsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeLicenseOrderDetailsRequest
	GetResourceOwnerId() *int64
}

type DescribeLicenseOrderDetailsRequest struct {
	// The Alibaba Cloud order ID (or virtual order ID).
	//
	// This parameter is required.
	//
	// example:
	//
	// 239618016570503
	AliyunOrderId        *string `json:"AliyunOrderId,omitempty" xml:"AliyunOrderId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeLicenseOrderDetailsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLicenseOrderDetailsRequest) GoString() string {
	return s.String()
}

func (s *DescribeLicenseOrderDetailsRequest) GetAliyunOrderId() *string {
	return s.AliyunOrderId
}

func (s *DescribeLicenseOrderDetailsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeLicenseOrderDetailsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLicenseOrderDetailsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeLicenseOrderDetailsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeLicenseOrderDetailsRequest) SetAliyunOrderId(v string) *DescribeLicenseOrderDetailsRequest {
	s.AliyunOrderId = &v
	return s
}

func (s *DescribeLicenseOrderDetailsRequest) SetOwnerAccount(v string) *DescribeLicenseOrderDetailsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeLicenseOrderDetailsRequest) SetOwnerId(v int64) *DescribeLicenseOrderDetailsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLicenseOrderDetailsRequest) SetResourceOwnerAccount(v string) *DescribeLicenseOrderDetailsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeLicenseOrderDetailsRequest) SetResourceOwnerId(v int64) *DescribeLicenseOrderDetailsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeLicenseOrderDetailsRequest) Validate() error {
	return dara.Validate(s)
}
