// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppsByAppInstanceGroupIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppInstanceGroupId(v string) *ListAppsByAppInstanceGroupIdRequest
	GetAppInstanceGroupId() *string
	SetPageNumber(v int32) *ListAppsByAppInstanceGroupIdRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAppsByAppInstanceGroupIdRequest
	GetPageSize() *int32
	SetProductType(v string) *ListAppsByAppInstanceGroupIdRequest
	GetProductType() *string
}

type ListAppsByAppInstanceGroupIdRequest struct {
	// The delivery group ID.
	//
	// - WUYING Cloud Application delivery group: call the [ListAppInstanceGroup](~~ListAppInstanceGroup~~) operation to obtain the ID.
	//
	// - Cloud Browser group: specify the Cloud Browser group ID. Call the [ListBrowserInstanceGroup](~~ListBrowserInstanceGroup~~) operation to obtain the ID.
	//
	// > This parameter is **required**. If it is not specified, the error code `InvalidParameter.AppInstanceGroupId` is returned.
	//
	// example:
	//
	// aig-9ciijz60n4xsv****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The page number. Pages start from page 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of applications to return per page. Valid values: 1 to 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The product type. The value must match the product type of the delivery group specified by AppInstanceGroupId. Otherwise, the error code `InvalidAppInstanceGroup.NotFound` is returned.
	//
	// Valid values:
	//
	// - CloudApp: WUYING Cloud Application.
	//
	// - CloudBrowser: Cloud Browser.
	//
	// This parameter is required.
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
}

func (s ListAppsByAppInstanceGroupIdRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAppsByAppInstanceGroupIdRequest) GoString() string {
	return s.String()
}

func (s *ListAppsByAppInstanceGroupIdRequest) GetAppInstanceGroupId() *string {
	return s.AppInstanceGroupId
}

func (s *ListAppsByAppInstanceGroupIdRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAppsByAppInstanceGroupIdRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAppsByAppInstanceGroupIdRequest) GetProductType() *string {
	return s.ProductType
}

func (s *ListAppsByAppInstanceGroupIdRequest) SetAppInstanceGroupId(v string) *ListAppsByAppInstanceGroupIdRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *ListAppsByAppInstanceGroupIdRequest) SetPageNumber(v int32) *ListAppsByAppInstanceGroupIdRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAppsByAppInstanceGroupIdRequest) SetPageSize(v int32) *ListAppsByAppInstanceGroupIdRequest {
	s.PageSize = &v
	return s
}

func (s *ListAppsByAppInstanceGroupIdRequest) SetProductType(v string) *ListAppsByAppInstanceGroupIdRequest {
	s.ProductType = &v
	return s
}

func (s *ListAppsByAppInstanceGroupIdRequest) Validate() error {
	return dara.Validate(s)
}
