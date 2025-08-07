// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeActivationCodesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunOrderId(v string) *DescribeActivationCodesRequest
	GetAliyunOrderId() *string
	SetMacAddress(v string) *DescribeActivationCodesRequest
	GetMacAddress() *string
	SetOwnerAccount(v string) *DescribeActivationCodesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeActivationCodesRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeActivationCodesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeActivationCodesRequest
	GetPageSize() *int32
	SetResourceOwnerAccount(v string) *DescribeActivationCodesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeActivationCodesRequest
	GetResourceOwnerId() *int64
	SetSystemIdentifier(v string) *DescribeActivationCodesRequest
	GetSystemIdentifier() *string
}

type DescribeActivationCodesRequest struct {
	// The ID of the Alibaba Cloud order. The value can be the ID of a virtual order.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2233****445566
	AliyunOrderId *string `json:"AliyunOrderId,omitempty" xml:"AliyunOrderId,omitempty"`
	// example:
	//
	// aa:bb:cc:dd:ee:ff
	MacAddress   *string `json:"MacAddress,omitempty" xml:"MacAddress,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 30
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// 3312548696141831911
	SystemIdentifier *string `json:"SystemIdentifier,omitempty" xml:"SystemIdentifier,omitempty"`
}

func (s DescribeActivationCodesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeActivationCodesRequest) GoString() string {
	return s.String()
}

func (s *DescribeActivationCodesRequest) GetAliyunOrderId() *string {
	return s.AliyunOrderId
}

func (s *DescribeActivationCodesRequest) GetMacAddress() *string {
	return s.MacAddress
}

func (s *DescribeActivationCodesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeActivationCodesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeActivationCodesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeActivationCodesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeActivationCodesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeActivationCodesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeActivationCodesRequest) GetSystemIdentifier() *string {
	return s.SystemIdentifier
}

func (s *DescribeActivationCodesRequest) SetAliyunOrderId(v string) *DescribeActivationCodesRequest {
	s.AliyunOrderId = &v
	return s
}

func (s *DescribeActivationCodesRequest) SetMacAddress(v string) *DescribeActivationCodesRequest {
	s.MacAddress = &v
	return s
}

func (s *DescribeActivationCodesRequest) SetOwnerAccount(v string) *DescribeActivationCodesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeActivationCodesRequest) SetOwnerId(v int64) *DescribeActivationCodesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeActivationCodesRequest) SetPageNumber(v int32) *DescribeActivationCodesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeActivationCodesRequest) SetPageSize(v int32) *DescribeActivationCodesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeActivationCodesRequest) SetResourceOwnerAccount(v string) *DescribeActivationCodesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeActivationCodesRequest) SetResourceOwnerId(v int64) *DescribeActivationCodesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeActivationCodesRequest) SetSystemIdentifier(v string) *DescribeActivationCodesRequest {
	s.SystemIdentifier = &v
	return s
}

func (s *DescribeActivationCodesRequest) Validate() error {
	return dara.Validate(s)
}
