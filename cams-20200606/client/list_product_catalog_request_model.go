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
	// The Business Manager ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 28
	BusinessId *int64 `json:"BusinessId,omitempty" xml:"BusinessId,omitempty"`
	// The space ID of the user within the independent software vendor (ISV) account.
	//
	// example:
	//
	// 293483938849493
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// The fields. Separate multiple fields with commas (,).
	//
	// see  [catalog fields](https://help.aliyun.com/document_detail/2579419.html)
	//
	// example:
	//
	// id,name
	Fields *string `json:"Fields,omitempty" xml:"Fields,omitempty"`
	// The number of catalogs to be queried. Valid values: 1 to 1000.
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
