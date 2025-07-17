// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewaysRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGatewayId(v string) *ListGatewaysRequest
	GetGatewayId() *string
	SetGatewayType(v string) *ListGatewaysRequest
	GetGatewayType() *string
	SetKeyword(v string) *ListGatewaysRequest
	GetKeyword() *string
	SetName(v string) *ListGatewaysRequest
	GetName() *string
	SetPageNumber(v int32) *ListGatewaysRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListGatewaysRequest
	GetPageSize() *int32
	SetResourceGroupId(v string) *ListGatewaysRequest
	GetResourceGroupId() *string
	SetTag(v []*ListGatewaysRequestTag) *ListGatewaysRequest
	GetTag() []*ListGatewaysRequestTag
}

type ListGatewaysRequest struct {
	// The instance ID. If you specify an ID, an exact search is performed.
	//
	// example:
	//
	// gw-cpv4sqdl****
	GatewayId   *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	GatewayType *string `json:"gatewayType,omitempty" xml:"gatewayType,omitempty"`
	// The search keyword. A full match is performed. The search is case-insensitive.
	//
	// example:
	//
	// dev
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// The instance name. If you specify a name, an exact search is performed.
	//
	// example:
	//
	// itemcenter-gateway
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-aekz3wes3hnre5a
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The tags that you want to use for the search.
	Tag []*ListGatewaysRequestTag `json:"tag,omitempty" xml:"tag,omitempty" type:"Repeated"`
}

func (s ListGatewaysRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGatewaysRequest) GoString() string {
	return s.String()
}

func (s *ListGatewaysRequest) GetGatewayId() *string {
	return s.GatewayId
}

func (s *ListGatewaysRequest) GetGatewayType() *string {
	return s.GatewayType
}

func (s *ListGatewaysRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListGatewaysRequest) GetName() *string {
	return s.Name
}

func (s *ListGatewaysRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListGatewaysRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListGatewaysRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListGatewaysRequest) GetTag() []*ListGatewaysRequestTag {
	return s.Tag
}

func (s *ListGatewaysRequest) SetGatewayId(v string) *ListGatewaysRequest {
	s.GatewayId = &v
	return s
}

func (s *ListGatewaysRequest) SetGatewayType(v string) *ListGatewaysRequest {
	s.GatewayType = &v
	return s
}

func (s *ListGatewaysRequest) SetKeyword(v string) *ListGatewaysRequest {
	s.Keyword = &v
	return s
}

func (s *ListGatewaysRequest) SetName(v string) *ListGatewaysRequest {
	s.Name = &v
	return s
}

func (s *ListGatewaysRequest) SetPageNumber(v int32) *ListGatewaysRequest {
	s.PageNumber = &v
	return s
}

func (s *ListGatewaysRequest) SetPageSize(v int32) *ListGatewaysRequest {
	s.PageSize = &v
	return s
}

func (s *ListGatewaysRequest) SetResourceGroupId(v string) *ListGatewaysRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListGatewaysRequest) SetTag(v []*ListGatewaysRequestTag) *ListGatewaysRequest {
	s.Tag = v
	return s
}

func (s *ListGatewaysRequest) Validate() error {
	return dara.Validate(s)
}

type ListGatewaysRequestTag struct {
	// The key of tag N.
	//
	// example:
	//
	// owner
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The value of tag N.
	//
	// example:
	//
	// zhangsan
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListGatewaysRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListGatewaysRequestTag) GoString() string {
	return s.String()
}

func (s *ListGatewaysRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListGatewaysRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListGatewaysRequestTag) SetKey(v string) *ListGatewaysRequestTag {
	s.Key = &v
	return s
}

func (s *ListGatewaysRequestTag) SetValue(v string) *ListGatewaysRequestTag {
	s.Value = &v
	return s
}

func (s *ListGatewaysRequestTag) Validate() error {
	return dara.Validate(s)
}
