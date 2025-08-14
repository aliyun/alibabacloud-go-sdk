// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeChildInstanceRegionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *DescribeChildInstanceRegionsRequest
	GetAcceptLanguage() *string
	SetOwnerAccount(v string) *DescribeChildInstanceRegionsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeChildInstanceRegionsRequest
	GetOwnerId() *int64
	SetProductType(v string) *DescribeChildInstanceRegionsRequest
	GetProductType() *string
	SetResourceOwnerAccount(v string) *DescribeChildInstanceRegionsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeChildInstanceRegionsRequest
	GetResourceOwnerId() *int64
}

type DescribeChildInstanceRegionsRequest struct {
	// The language of the response. Valid values: zh-CN (Chinese, which is the default language), en-US (English), and ja (Japanese).
	//
	// example:
	//
	// zh-CN
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	OwnerAccount   *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The type of the network instance. Valid values:
	//
	// 	- **VPC**: virtual private cloud (VPC)
	//
	// 	- **VBR**: virtual border router (VBR)
	//
	// 	- **CCN**: Cloud Connect Network (CCN) instance
	//
	// example:
	//
	// VPC
	ProductType          *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeChildInstanceRegionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeChildInstanceRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeChildInstanceRegionsRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *DescribeChildInstanceRegionsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeChildInstanceRegionsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeChildInstanceRegionsRequest) GetProductType() *string {
	return s.ProductType
}

func (s *DescribeChildInstanceRegionsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeChildInstanceRegionsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeChildInstanceRegionsRequest) SetAcceptLanguage(v string) *DescribeChildInstanceRegionsRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DescribeChildInstanceRegionsRequest) SetOwnerAccount(v string) *DescribeChildInstanceRegionsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeChildInstanceRegionsRequest) SetOwnerId(v int64) *DescribeChildInstanceRegionsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeChildInstanceRegionsRequest) SetProductType(v string) *DescribeChildInstanceRegionsRequest {
	s.ProductType = &v
	return s
}

func (s *DescribeChildInstanceRegionsRequest) SetResourceOwnerAccount(v string) *DescribeChildInstanceRegionsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeChildInstanceRegionsRequest) SetResourceOwnerId(v int64) *DescribeChildInstanceRegionsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeChildInstanceRegionsRequest) Validate() error {
	return dara.Validate(s)
}
