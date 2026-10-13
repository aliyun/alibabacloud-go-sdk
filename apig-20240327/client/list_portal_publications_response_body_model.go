// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPortalPublicationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListPortalPublicationsResponseBody
	GetCode() *string
	SetData(v *ListPortalPublicationsResponseBodyData) *ListPortalPublicationsResponseBody
	GetData() *ListPortalPublicationsResponseBodyData
	SetMessage(v string) *ListPortalPublicationsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListPortalPublicationsResponseBody
	GetRequestId() *string
}

type ListPortalPublicationsResponseBody struct {
	// The response status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response data.
	Data *ListPortalPublicationsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 393E2630-DBE7-5221-AB35-9E740675491A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListPortalPublicationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPortalPublicationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPortalPublicationsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListPortalPublicationsResponseBody) GetData() *ListPortalPublicationsResponseBodyData {
	return s.Data
}

func (s *ListPortalPublicationsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListPortalPublicationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPortalPublicationsResponseBody) SetCode(v string) *ListPortalPublicationsResponseBody {
	s.Code = &v
	return s
}

func (s *ListPortalPublicationsResponseBody) SetData(v *ListPortalPublicationsResponseBodyData) *ListPortalPublicationsResponseBody {
	s.Data = v
	return s
}

func (s *ListPortalPublicationsResponseBody) SetMessage(v string) *ListPortalPublicationsResponseBody {
	s.Message = &v
	return s
}

func (s *ListPortalPublicationsResponseBody) SetRequestId(v string) *ListPortalPublicationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPortalPublicationsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListPortalPublicationsResponseBodyData struct {
	// The list of published API products on the portal. The list elements are named portalConfig, and each element represents an API product publication record.
	Content []*ListPortalPublicationsResponseBodyDataContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The page size.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The total number of items.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListPortalPublicationsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListPortalPublicationsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListPortalPublicationsResponseBodyData) GetContent() []*ListPortalPublicationsResponseBodyDataContent {
	return s.Content
}

func (s *ListPortalPublicationsResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListPortalPublicationsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPortalPublicationsResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListPortalPublicationsResponseBodyData) SetContent(v []*ListPortalPublicationsResponseBodyDataContent) *ListPortalPublicationsResponseBodyData {
	s.Content = v
	return s
}

func (s *ListPortalPublicationsResponseBodyData) SetPageNumber(v int32) *ListPortalPublicationsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListPortalPublicationsResponseBodyData) SetPageSize(v int32) *ListPortalPublicationsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListPortalPublicationsResponseBodyData) SetTotalCount(v int32) *ListPortalPublicationsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListPortalPublicationsResponseBodyData) Validate() error {
	if s.Content != nil {
		for _, item := range s.Content {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPortalPublicationsResponseBodyDataContent struct {
	// The creation time in ISO 8601 (RFC 3339) format in the UTC time zone, such as 2026-05-15T10:00:00Z.
	//
	// example:
	//
	// 2026-05-15T10:00:00Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The description of the API product.
	//
	// example:
	//
	// API product for model services available to internal developers
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The API product ID. This is the primary ID of the Product resource, which identifies the API product associated with this publication record.
	//
	// example:
	//
	// product-xxxxxxxx
	ProductId *string `json:"productId,omitempty" xml:"productId,omitempty"`
	// The API product name. This is the Name property of the Product resource.
	//
	// example:
	//
	// model-service
	ProductName *string `json:"productName,omitempty" xml:"productName,omitempty"`
	// The API product type. The valid values are not defined in the API specification. Refer to the types displayed in the console.
	//
	// example:
	//
	// HTTP_API
	ProductType *string `json:"productType,omitempty" xml:"productType,omitempty"`
	// The publication record ID of the API product. This ID is returned when the PublishProduct operation publishes the API product to the portal. You can use it in the /himarket/v1/products/{productId}/publications/{publicationId} path to locate the publication record for the UnpublishProduct operation.
	//
	// example:
	//
	// publication-xxxxxxxx
	PublicationId *string `json:"publicationId,omitempty" xml:"publicationId,omitempty"`
}

func (s ListPortalPublicationsResponseBodyDataContent) String() string {
	return dara.Prettify(s)
}

func (s ListPortalPublicationsResponseBodyDataContent) GoString() string {
	return s.String()
}

func (s *ListPortalPublicationsResponseBodyDataContent) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ListPortalPublicationsResponseBodyDataContent) GetDescription() *string {
	return s.Description
}

func (s *ListPortalPublicationsResponseBodyDataContent) GetProductId() *string {
	return s.ProductId
}

func (s *ListPortalPublicationsResponseBodyDataContent) GetProductName() *string {
	return s.ProductName
}

func (s *ListPortalPublicationsResponseBodyDataContent) GetProductType() *string {
	return s.ProductType
}

func (s *ListPortalPublicationsResponseBodyDataContent) GetPublicationId() *string {
	return s.PublicationId
}

func (s *ListPortalPublicationsResponseBodyDataContent) SetCreatedAt(v string) *ListPortalPublicationsResponseBodyDataContent {
	s.CreatedAt = &v
	return s
}

func (s *ListPortalPublicationsResponseBodyDataContent) SetDescription(v string) *ListPortalPublicationsResponseBodyDataContent {
	s.Description = &v
	return s
}

func (s *ListPortalPublicationsResponseBodyDataContent) SetProductId(v string) *ListPortalPublicationsResponseBodyDataContent {
	s.ProductId = &v
	return s
}

func (s *ListPortalPublicationsResponseBodyDataContent) SetProductName(v string) *ListPortalPublicationsResponseBodyDataContent {
	s.ProductName = &v
	return s
}

func (s *ListPortalPublicationsResponseBodyDataContent) SetProductType(v string) *ListPortalPublicationsResponseBodyDataContent {
	s.ProductType = &v
	return s
}

func (s *ListPortalPublicationsResponseBodyDataContent) SetPublicationId(v string) *ListPortalPublicationsResponseBodyDataContent {
	s.PublicationId = &v
	return s
}

func (s *ListPortalPublicationsResponseBodyDataContent) Validate() error {
	return dara.Validate(s)
}
