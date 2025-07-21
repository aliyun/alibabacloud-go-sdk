// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProductRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAfter(v string) *ListProductRequest
	GetAfter() *string
	SetBefore(v string) *ListProductRequest
	GetBefore() *string
	SetCatalogId(v string) *ListProductRequest
	GetCatalogId() *string
	SetCustSpaceId(v string) *ListProductRequest
	GetCustSpaceId() *string
	SetFields(v string) *ListProductRequest
	GetFields() *string
	SetLimit(v int64) *ListProductRequest
	GetLimit() *int64
	SetOwnerId(v int64) *ListProductRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ListProductRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListProductRequest
	GetResourceOwnerId() *int64
	SetWabaId(v string) *ListProductRequest
	GetWabaId() *string
}

type ListProductRequest struct {
	// The cursor that points to the end of the page of the returned data.
	//
	// example:
	//
	// kdkii48jfjjei3
	After *string `json:"After,omitempty" xml:"After,omitempty"`
	// The cursor that points to the beginning of the page of the returned data.
	//
	// example:
	//
	// wiidkd939kek93
	Before *string `json:"Before,omitempty" xml:"Before,omitempty"`
	// The catalog ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 29398389292
	CatalogId *string `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
	// The space ID of the user within the independent software vendor (ISV) account.
	//
	// example:
	//
	// C29398388383
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// The fields. Separate multiple fields with commas (,).
	//
	//  see [product fields](https://help.aliyun.com/document_detail/2579419.html)
	//
	// example:
	//
	// id,name
	Fields *string `json:"Fields,omitempty" xml:"Fields,omitempty"`
	// The number of products to be queried. Valid values: 1 to 1000.
	//
	// example:
	//
	// 73
	Limit                *int64  `json:"Limit,omitempty" xml:"Limit,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the WhatsApp Business account (WABA).
	//
	// This parameter is required.
	//
	// example:
	//
	// 38487474747
	WabaId *string `json:"WabaId,omitempty" xml:"WabaId,omitempty"`
}

func (s ListProductRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProductRequest) GoString() string {
	return s.String()
}

func (s *ListProductRequest) GetAfter() *string {
	return s.After
}

func (s *ListProductRequest) GetBefore() *string {
	return s.Before
}

func (s *ListProductRequest) GetCatalogId() *string {
	return s.CatalogId
}

func (s *ListProductRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *ListProductRequest) GetFields() *string {
	return s.Fields
}

func (s *ListProductRequest) GetLimit() *int64 {
	return s.Limit
}

func (s *ListProductRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListProductRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListProductRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListProductRequest) GetWabaId() *string {
	return s.WabaId
}

func (s *ListProductRequest) SetAfter(v string) *ListProductRequest {
	s.After = &v
	return s
}

func (s *ListProductRequest) SetBefore(v string) *ListProductRequest {
	s.Before = &v
	return s
}

func (s *ListProductRequest) SetCatalogId(v string) *ListProductRequest {
	s.CatalogId = &v
	return s
}

func (s *ListProductRequest) SetCustSpaceId(v string) *ListProductRequest {
	s.CustSpaceId = &v
	return s
}

func (s *ListProductRequest) SetFields(v string) *ListProductRequest {
	s.Fields = &v
	return s
}

func (s *ListProductRequest) SetLimit(v int64) *ListProductRequest {
	s.Limit = &v
	return s
}

func (s *ListProductRequest) SetOwnerId(v int64) *ListProductRequest {
	s.OwnerId = &v
	return s
}

func (s *ListProductRequest) SetResourceOwnerAccount(v string) *ListProductRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListProductRequest) SetResourceOwnerId(v int64) *ListProductRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListProductRequest) SetWabaId(v string) *ListProductRequest {
	s.WabaId = &v
	return s
}

func (s *ListProductRequest) Validate() error {
	return dara.Validate(s)
}
