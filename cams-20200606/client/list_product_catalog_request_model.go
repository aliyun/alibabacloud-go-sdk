// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProductCatalogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAfter(v string) *ListProductCatalogRequest
	GetAfter() *string
	SetBefore(v string) *ListProductCatalogRequest
	GetBefore() *string
	SetBusinessId(v int64) *ListProductCatalogRequest
	GetBusinessId() *int64
	SetCustSpaceId(v string) *ListProductCatalogRequest
	GetCustSpaceId() *string
	SetFields(v string) *ListProductCatalogRequest
	GetFields() *string
	SetLimit(v int64) *ListProductCatalogRequest
	GetLimit() *int64
	SetOwnerId(v int64) *ListProductCatalogRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ListProductCatalogRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListProductCatalogRequest
	GetResourceOwnerId() *int64
}

type ListProductCatalogRequest struct {
	// The start position of the next page for pagination.
	//
	// example:
	//
	// kdkii48jfjjei3
	After *string `json:"After,omitempty" xml:"After,omitempty"`
	// The end position of the previous page for pagination.
	//
	// example:
	//
	// wiidkd939kek93
	Before *string `json:"Before,omitempty" xml:"Before,omitempty"`
	// The ID of the business platform where the WABA is located. You can view the business platform ID in the Chat App Message Service console after WABA registration is complete.
	//
	// This parameter is required.
	//
	// example:
	//
	// 105048970863****
	BusinessId *int64 `json:"BusinessId,omitempty" xml:"BusinessId,omitempty"`
	// The space ID of the ISV sub-customer.
	//
	// example:
	//
	// C29398882929
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// The list of fields, separated by commas (,).
	//
	// For specific fields, see [Catalog fields](https://help.aliyun.com/document_detail/2579419.html).
	//
	// example:
	//
	// id,name
	Fields *string `json:"Fields,omitempty" xml:"Fields,omitempty"`
	// The number of entries to query. Valid values: 1 to 1000.
	//
	// example:
	//
	// 73
	Limit                *int64  `json:"Limit,omitempty" xml:"Limit,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ListProductCatalogRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProductCatalogRequest) GoString() string {
	return s.String()
}

func (s *ListProductCatalogRequest) GetAfter() *string {
	return s.After
}

func (s *ListProductCatalogRequest) GetBefore() *string {
	return s.Before
}

func (s *ListProductCatalogRequest) GetBusinessId() *int64 {
	return s.BusinessId
}

func (s *ListProductCatalogRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *ListProductCatalogRequest) GetFields() *string {
	return s.Fields
}

func (s *ListProductCatalogRequest) GetLimit() *int64 {
	return s.Limit
}

func (s *ListProductCatalogRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListProductCatalogRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListProductCatalogRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListProductCatalogRequest) SetAfter(v string) *ListProductCatalogRequest {
	s.After = &v
	return s
}

func (s *ListProductCatalogRequest) SetBefore(v string) *ListProductCatalogRequest {
	s.Before = &v
	return s
}

func (s *ListProductCatalogRequest) SetBusinessId(v int64) *ListProductCatalogRequest {
	s.BusinessId = &v
	return s
}

func (s *ListProductCatalogRequest) SetCustSpaceId(v string) *ListProductCatalogRequest {
	s.CustSpaceId = &v
	return s
}

func (s *ListProductCatalogRequest) SetFields(v string) *ListProductCatalogRequest {
	s.Fields = &v
	return s
}

func (s *ListProductCatalogRequest) SetLimit(v int64) *ListProductCatalogRequest {
	s.Limit = &v
	return s
}

func (s *ListProductCatalogRequest) SetOwnerId(v int64) *ListProductCatalogRequest {
	s.OwnerId = &v
	return s
}

func (s *ListProductCatalogRequest) SetResourceOwnerAccount(v string) *ListProductCatalogRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListProductCatalogRequest) SetResourceOwnerId(v int64) *ListProductCatalogRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListProductCatalogRequest) Validate() error {
	return dara.Validate(s)
}
