// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPageListAppInstanceGroupUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppInstanceGroupId(v string) *PageListAppInstanceGroupUserRequest
	GetAppInstanceGroupId() *string
	SetPageNumber(v int32) *PageListAppInstanceGroupUserRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *PageListAppInstanceGroupUserRequest
	GetPageSize() *int32
	SetProductType(v string) *PageListAppInstanceGroupUserRequest
	GetProductType() *string
}

type PageListAppInstanceGroupUserRequest struct {
	// The ID of the delivery group.
	//
	// This parameter is required.
	//
	// example:
	//
	// aig-9ciijz60n4xsv****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The number of the page to return. We recommend that you configure this parameter.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to be return on each page. The value cannot be greater than `100`. We recommend that you configure this parameter.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The product type.
	//
	// Valid value:
	//
	// 	- CloudApp: App Streaming
	//
	// This parameter is required.
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
}

func (s PageListAppInstanceGroupUserRequest) String() string {
	return dara.Prettify(s)
}

func (s PageListAppInstanceGroupUserRequest) GoString() string {
	return s.String()
}

func (s *PageListAppInstanceGroupUserRequest) GetAppInstanceGroupId() *string {
	return s.AppInstanceGroupId
}

func (s *PageListAppInstanceGroupUserRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *PageListAppInstanceGroupUserRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *PageListAppInstanceGroupUserRequest) GetProductType() *string {
	return s.ProductType
}

func (s *PageListAppInstanceGroupUserRequest) SetAppInstanceGroupId(v string) *PageListAppInstanceGroupUserRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *PageListAppInstanceGroupUserRequest) SetPageNumber(v int32) *PageListAppInstanceGroupUserRequest {
	s.PageNumber = &v
	return s
}

func (s *PageListAppInstanceGroupUserRequest) SetPageSize(v int32) *PageListAppInstanceGroupUserRequest {
	s.PageSize = &v
	return s
}

func (s *PageListAppInstanceGroupUserRequest) SetProductType(v string) *PageListAppInstanceGroupUserRequest {
	s.ProductType = &v
	return s
}

func (s *PageListAppInstanceGroupUserRequest) Validate() error {
	return dara.Validate(s)
}
