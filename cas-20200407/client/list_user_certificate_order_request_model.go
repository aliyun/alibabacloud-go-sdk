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
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Performs a fuzzy query. The keyword can be a domain name or a resource ID.
	//
	// example:
	//
	// cert-instanceId
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The resource type. Default value: **CPACK**. Valid values:
	//
	// - **CPACK**: An order for a resource plan. Only orders created from a resource plan are returned.
	//
	// - **BUY**: A direct purchase. Only orders created from direct purchases are returned. You can ignore this type in most cases.
	//
	// - **UPLOAD**: An uploaded certificate. Only uploaded certificates are returned.
	//
	// - **CERT**: A certificate. Both issued and uploaded certificates are returned.
	//
	// example:
	//
	// CPACK
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// The ID of the resource group. For more information, see [ListResources](https://help.aliyun.com/document_detail/2716559.html).
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
	// The status of the order. Valid values:
	//
	// - **PAYED**: The certificate is pending application. This value is valid only when OrderType is set to CPACK or BUY.
	//
	// - **CHECKING**: The certificate is under review. This value is valid only when OrderType is set to CPACK or BUY.
	//
	// - **CHECKED_FAIL**: The review failed. This value is valid only when OrderType is set to CPACK or BUY.
	//
	// - **ISSUED**: The certificate is issued.
	//
	// - **WILLEXPIRED**: The certificate is about to expire.
	//
	// - **EXPIRED**: The certificate has expired.
	//
	// - **NOTACTIVATED**: The certificate is not activated. This value is valid only when OrderType is set to CPACK or BUY.
	//
	// - **REVOKED**: The certificate is revoked. This value is valid only when OrderType is set to CPACK or BUY.
	//
	// If OrderType is CERT or UPLOAD and you leave this parameter empty, active certificates are returned by default. Active certificates are those in the ISSUED or WILLEXPIRED state. If OrderType is CPACK or BUY and you leave this parameter empty, all orders are returned by default.
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
