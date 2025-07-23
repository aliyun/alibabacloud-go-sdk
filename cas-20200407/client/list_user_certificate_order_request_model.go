// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserCertificateOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int64) *ListUserCertificateOrderRequest
	GetCurrentPage() *int64
	SetKeyword(v string) *ListUserCertificateOrderRequest
	GetKeyword() *string
	SetOrderType(v string) *ListUserCertificateOrderRequest
	GetOrderType() *string
	SetResourceGroupId(v string) *ListUserCertificateOrderRequest
	GetResourceGroupId() *string
	SetShowSize(v int64) *ListUserCertificateOrderRequest
	GetShowSize() *int64
	SetStatus(v string) *ListUserCertificateOrderRequest
	GetStatus() *string
}

type ListUserCertificateOrderRequest struct {
	// The number of the page to return.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The domain name that is bound or the ID of the resource. Fuzzy match is supported.
	//
	// example:
	//
	// cert-instanceId
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The type of the order. Default value: **CPACK**. Valid values:
	//
	// 	- **CPACK**: virtual resource order. If you set OrderType to CPACK, only the information about orders that are generated to consume the certificate quota is returned.
	//
	// 	- **BUY**: purchase order. If you set OrderType to BUY, only the information about purchase orders is returned. In most cases, this type of order can be ignored.
	//
	// 	- **UPLOAD**: uploaded certificate. If you set OrderType to UPLOAD, only uploaded certificates are returned.
	//
	// 	- **CERT**: certificate. If you set OrderType to CERT, both issued certificates and uploaded certificates are returned.
	//
	// example:
	//
	// CPACK
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// The ID of the resource group. You can call the [ListResources](https://help.aliyun.com/document_detail/2716559.html) operation to obtain the ID.
	//
	// example:
	//
	// rg-ae******4wia
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The number of entries to return on each page. Default value: 50.
	//
	// example:
	//
	// 10
	ShowSize *int64 `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
	// The certificate status of the order. Valid values:
	//
	// 	- **PAYED**: pending application. You can set Status to PAYED only if you set OrderType to CPACK or BUY.
	//
	// 	- **CHECKING**: validating. You can set Status to CHECKING only if you set OrderType to CPACK or BUY.
	//
	// 	- **CHECKED_FAIL**: validation failed. You can set Status to CHECKED_FAIL only if you set OrderType to CPACK or BUY.
	//
	// 	- **ISSUED**: issued.
	//
	// 	- **WILLEXPIRED**: about to expire.
	//
	// 	- **EXPIRED**: expired.
	//
	// 	- **NOTACTIVATED**: not activated. You can set Status to NOTACTIVATED only if you set OrderType to CPACK or BUY.
	//
	// 	- **REVOKED**: revoked. You can set Status to REVOKED only if you set OrderType to CPACK or BUY.
	//
	// If you set OrderType to CERT or UPLOAD and Status is left empty, valid certificates are returned by default, including issued certificates and certificates that are about to expire. If you set OrderType to CPACK or BUY and Status is left empty, all orders are returned by default.
	//
	// example:
	//
	// ISSUED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListUserCertificateOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUserCertificateOrderRequest) GoString() string {
	return s.String()
}

func (s *ListUserCertificateOrderRequest) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *ListUserCertificateOrderRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListUserCertificateOrderRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *ListUserCertificateOrderRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListUserCertificateOrderRequest) GetShowSize() *int64 {
	return s.ShowSize
}

func (s *ListUserCertificateOrderRequest) GetStatus() *string {
	return s.Status
}

func (s *ListUserCertificateOrderRequest) SetCurrentPage(v int64) *ListUserCertificateOrderRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListUserCertificateOrderRequest) SetKeyword(v string) *ListUserCertificateOrderRequest {
	s.Keyword = &v
	return s
}

func (s *ListUserCertificateOrderRequest) SetOrderType(v string) *ListUserCertificateOrderRequest {
	s.OrderType = &v
	return s
}

func (s *ListUserCertificateOrderRequest) SetResourceGroupId(v string) *ListUserCertificateOrderRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListUserCertificateOrderRequest) SetShowSize(v int64) *ListUserCertificateOrderRequest {
	s.ShowSize = &v
	return s
}

func (s *ListUserCertificateOrderRequest) SetStatus(v string) *ListUserCertificateOrderRequest {
	s.Status = &v
	return s
}

func (s *ListUserCertificateOrderRequest) Validate() error {
	return dara.Validate(s)
}
