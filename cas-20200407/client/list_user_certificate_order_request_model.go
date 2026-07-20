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
	// The page number of the current page in a paginated query.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The keyword for fuzzy search. Matches the domain name or the corresponding resource ID.
	//
	// example:
	//
	// cert-instanceId
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The resource type. Default value: **CPACK**. Valid values:
	//
	// - **CPACK**: resource virtual order. Only orders generated from quotas are returned.
	//
	// - **BUY**: purchase order. Only orders generated from purchases are returned. You can ignore this type in most cases.
	//
	// - **UPLOAD**: uploaded certificate. Only uploaded certificates are returned.
	//
	// - **CERT**: certificate. Both issued certificates and uploaded certificates are returned.
	//
	// example:
	//
	// CPACK
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// The resource group ID. You can obtain this ID by calling the [ListResources](https://help.aliyun.com/document_detail/2716559.html) operation.
	//
	// example:
	//
	// rg-ae******4wia
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The number of entries per page in a paginated query. Default value: 50.
	//
	// example:
	//
	// 10
	ShowSize *int64 `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
	// The order status. Valid values:
	//
	// - **PAYED**: Pending application. Valid when OrderType is set to CPACK or BUY.
	//
	// - **CHECKING**: Under review. Valid when OrderType is set to CPACK or BUY.
	//
	// - **CHECKED_FAIL**: Review failed. Valid when OrderType is set to CPACK or BUY.
	//
	// - **ISSUED**: Issued.
	//
	// - **WILLEXPIRED**: About to expire.
	//
	// - **EXPIRED**: Expired.
	//
	// - **NOTACTIVATED**: Not activated. Valid when OrderType is set to CPACK or BUY.
	//
	// - **REVOKED**: Revoked. Valid when OrderType is set to CPACK or BUY.
	//
	// If OrderType is set to CERT or UPLOAD and Status is empty, valid certificates are returned by default, including issued and about-to-expire certificates. If OrderType is set to CPACK or BUY and Status is empty, all orders are returned by default.
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
