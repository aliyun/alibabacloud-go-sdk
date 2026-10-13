// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQuotasRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBillingType(v string) *ListQuotasRequest
	GetBillingType() *string
	SetMarker(v string) *ListQuotasRequest
	GetMarker() *string
	SetMaxItem(v int64) *ListQuotasRequest
	GetMaxItem() *int64
	SetProductId(v string) *ListQuotasRequest
	GetProductId() *string
	SetRegion(v string) *ListQuotasRequest
	GetRegion() *string
	SetSaleTags(v string) *ListQuotasRequest
	GetSaleTags() *string
	SetTenantId(v string) *ListQuotasRequest
	GetTenantId() *string
}

type ListQuotasRequest struct {
	// The billing method.
	//
	// example:
	//
	// subscription
	BillingType *string `json:"billingType,omitempty" xml:"billingType,omitempty"`
	// The marker after which the results are returned in alphabetical order.
	//
	// example:
	//
	// cHlvZHBzX3VkZl8xMDExNV8xNDU3NDI4NDkzKg==
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
	// The maximum number of items to return.
	//
	// example:
	//
	// 100
	MaxItem *int64 `json:"maxItem,omitempty" xml:"maxItem,omitempty"`
	// Deprecated
	ProductId *string `json:"productId,omitempty" xml:"productId,omitempty"`
	// Deprecated
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// The billing tags used to filter quota objects. Billing tags are created by users for cloud products.
	//
	// example:
	//
	// {"tag":"this_is_tag_demo"}
	SaleTags *string `json:"saleTags,omitempty" xml:"saleTags,omitempty"`
	// Deprecated
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListQuotasRequest) String() string {
	return dara.Prettify(s)
}

func (s ListQuotasRequest) GoString() string {
	return s.String()
}

func (s *ListQuotasRequest) GetBillingType() *string {
	return s.BillingType
}

func (s *ListQuotasRequest) GetMarker() *string {
	return s.Marker
}

func (s *ListQuotasRequest) GetMaxItem() *int64 {
	return s.MaxItem
}

func (s *ListQuotasRequest) GetProductId() *string {
	return s.ProductId
}

func (s *ListQuotasRequest) GetRegion() *string {
	return s.Region
}

func (s *ListQuotasRequest) GetSaleTags() *string {
	return s.SaleTags
}

func (s *ListQuotasRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListQuotasRequest) SetBillingType(v string) *ListQuotasRequest {
	s.BillingType = &v
	return s
}

func (s *ListQuotasRequest) SetMarker(v string) *ListQuotasRequest {
	s.Marker = &v
	return s
}

func (s *ListQuotasRequest) SetMaxItem(v int64) *ListQuotasRequest {
	s.MaxItem = &v
	return s
}

func (s *ListQuotasRequest) SetProductId(v string) *ListQuotasRequest {
	s.ProductId = &v
	return s
}

func (s *ListQuotasRequest) SetRegion(v string) *ListQuotasRequest {
	s.Region = &v
	return s
}

func (s *ListQuotasRequest) SetSaleTags(v string) *ListQuotasRequest {
	s.SaleTags = &v
	return s
}

func (s *ListQuotasRequest) SetTenantId(v string) *ListQuotasRequest {
	s.TenantId = &v
	return s
}

func (s *ListQuotasRequest) Validate() error {
	return dara.Validate(s)
}
