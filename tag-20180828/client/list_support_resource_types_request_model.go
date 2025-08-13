// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSupportResourceTypesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResult(v int32) *ListSupportResourceTypesRequest
	GetMaxResult() *int32
	SetNextToken(v string) *ListSupportResourceTypesRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *ListSupportResourceTypesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListSupportResourceTypesRequest
	GetOwnerId() *int64
	SetProductCode(v string) *ListSupportResourceTypesRequest
	GetProductCode() *string
	SetRegionId(v string) *ListSupportResourceTypesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ListSupportResourceTypesRequest
	GetResourceOwnerAccount() *string
	SetResourceTye(v string) *ListSupportResourceTypesRequest
	GetResourceTye() *string
	SetShowItems(v bool) *ListSupportResourceTypesRequest
	GetShowItems() *bool
	SetSupportCode(v string) *ListSupportResourceTypesRequest
	GetSupportCode() *string
}

type ListSupportResourceTypesRequest struct {
	// The number of entries to return on each page.
	//
	// Maximum value: 1000. Default value: 50.
	//
	// example:
	//
	// 50
	MaxResult *int32 `json:"MaxResult,omitempty" xml:"MaxResult,omitempty"`
	// The token that is used to start the next query.
	//
	// example:
	//
	// AAAAAYws9fJ0Ur4MGm/5OkDoW/Y3wDNwUdssyKODK****
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The service code. This parameter specifies a filter condition for the query.
	//
	// This parameter is obtained from the response.
	//
	// example:
	//
	// ecs
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The region ID.
	//
	// For more information about region IDs, see [Endpoints](https://help.aliyun.com/document_detail/2330902.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The resource type. This parameter specifies a filter condition for the query.
	//
	// This parameter is obtained from the response.
	//
	// example:
	//
	// instance
	ResourceTye *string `json:"ResourceTye,omitempty" xml:"ResourceTye,omitempty"`
	// Specifies whether to return tag-related capability items. Valid values:
	//
	// 	- true: The system returns tag-related capability items.
	//
	// 	- false (default value): The system does not return tag-related capability items.
	//
	// example:
	//
	// false
	ShowItems *bool `json:"ShowItems,omitempty" xml:"ShowItems,omitempty"`
	// The code of the tag-related capability item. This parameter specifies a filter condition for the query.
	//
	// For more information, see **Tag-related capability items**.
	//
	// example:
	//
	// TAG_CONSOLE_SUPPORT
	SupportCode *string `json:"SupportCode,omitempty" xml:"SupportCode,omitempty"`
}

func (s ListSupportResourceTypesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSupportResourceTypesRequest) GoString() string {
	return s.String()
}

func (s *ListSupportResourceTypesRequest) GetMaxResult() *int32 {
	return s.MaxResult
}

func (s *ListSupportResourceTypesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSupportResourceTypesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListSupportResourceTypesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListSupportResourceTypesRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *ListSupportResourceTypesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListSupportResourceTypesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListSupportResourceTypesRequest) GetResourceTye() *string {
	return s.ResourceTye
}

func (s *ListSupportResourceTypesRequest) GetShowItems() *bool {
	return s.ShowItems
}

func (s *ListSupportResourceTypesRequest) GetSupportCode() *string {
	return s.SupportCode
}

func (s *ListSupportResourceTypesRequest) SetMaxResult(v int32) *ListSupportResourceTypesRequest {
	s.MaxResult = &v
	return s
}

func (s *ListSupportResourceTypesRequest) SetNextToken(v string) *ListSupportResourceTypesRequest {
	s.NextToken = &v
	return s
}

func (s *ListSupportResourceTypesRequest) SetOwnerAccount(v string) *ListSupportResourceTypesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListSupportResourceTypesRequest) SetOwnerId(v int64) *ListSupportResourceTypesRequest {
	s.OwnerId = &v
	return s
}

func (s *ListSupportResourceTypesRequest) SetProductCode(v string) *ListSupportResourceTypesRequest {
	s.ProductCode = &v
	return s
}

func (s *ListSupportResourceTypesRequest) SetRegionId(v string) *ListSupportResourceTypesRequest {
	s.RegionId = &v
	return s
}

func (s *ListSupportResourceTypesRequest) SetResourceOwnerAccount(v string) *ListSupportResourceTypesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListSupportResourceTypesRequest) SetResourceTye(v string) *ListSupportResourceTypesRequest {
	s.ResourceTye = &v
	return s
}

func (s *ListSupportResourceTypesRequest) SetShowItems(v bool) *ListSupportResourceTypesRequest {
	s.ShowItems = &v
	return s
}

func (s *ListSupportResourceTypesRequest) SetSupportCode(v string) *ListSupportResourceTypesRequest {
	s.SupportCode = &v
	return s
}

func (s *ListSupportResourceTypesRequest) Validate() error {
	return dara.Validate(s)
}
