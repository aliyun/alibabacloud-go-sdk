// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProductsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetProducts(v []*ListProductsResponseBodyProducts) *ListProductsResponseBody
	GetProducts() []*ListProductsResponseBodyProducts
	SetRequestId(v string) *ListProductsResponseBody
	GetRequestId() *string
	SetServices(v []*ListProductsResponseBodyServices) *ListProductsResponseBody
	GetServices() []*ListProductsResponseBodyServices
}

type ListProductsResponseBody struct {
	Products []*ListProductsResponseBodyProducts `json:"Products,omitempty" xml:"Products,omitempty" type:"Repeated"`
	// example:
	//
	// 1e195c5116124202371861018d5bde
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Services  []*ListProductsResponseBodyServices `json:"Services,omitempty" xml:"Services,omitempty" type:"Repeated"`
}

func (s ListProductsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListProductsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProductsResponseBody) GetProducts() []*ListProductsResponseBodyProducts {
	return s.Products
}

func (s *ListProductsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListProductsResponseBody) GetServices() []*ListProductsResponseBodyServices {
	return s.Services
}

func (s *ListProductsResponseBody) SetProducts(v []*ListProductsResponseBodyProducts) *ListProductsResponseBody {
	s.Products = v
	return s
}

func (s *ListProductsResponseBody) SetRequestId(v string) *ListProductsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProductsResponseBody) SetServices(v []*ListProductsResponseBodyServices) *ListProductsResponseBody {
	s.Services = v
	return s
}

func (s *ListProductsResponseBody) Validate() error {
	if s.Products != nil {
		for _, item := range s.Products {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Services != nil {
		for _, item := range s.Services {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListProductsResponseBodyProducts struct {
	HasPermissionToPurchase *bool `json:"HasPermissionToPurchase,omitempty" xml:"HasPermissionToPurchase,omitempty"`
	// example:
	//
	// true
	IsPurchased *bool `json:"IsPurchased,omitempty" xml:"IsPurchased,omitempty"`
	// example:
	//
	// DataWorks_isolate
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProductId   *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// example:
	//
	// https://common-buy.aliy
	PurchaseUrl *string `json:"PurchaseUrl,omitempty" xml:"PurchaseUrl,omitempty"`
}

func (s ListProductsResponseBodyProducts) String() string {
	return dara.Prettify(s)
}

func (s ListProductsResponseBodyProducts) GoString() string {
	return s.String()
}

func (s *ListProductsResponseBodyProducts) GetHasPermissionToPurchase() *bool {
	return s.HasPermissionToPurchase
}

func (s *ListProductsResponseBodyProducts) GetIsPurchased() *bool {
	return s.IsPurchased
}

func (s *ListProductsResponseBodyProducts) GetProductCode() *string {
	return s.ProductCode
}

func (s *ListProductsResponseBodyProducts) GetProductId() *string {
	return s.ProductId
}

func (s *ListProductsResponseBodyProducts) GetPurchaseUrl() *string {
	return s.PurchaseUrl
}

func (s *ListProductsResponseBodyProducts) SetHasPermissionToPurchase(v bool) *ListProductsResponseBodyProducts {
	s.HasPermissionToPurchase = &v
	return s
}

func (s *ListProductsResponseBodyProducts) SetIsPurchased(v bool) *ListProductsResponseBodyProducts {
	s.IsPurchased = &v
	return s
}

func (s *ListProductsResponseBodyProducts) SetProductCode(v string) *ListProductsResponseBodyProducts {
	s.ProductCode = &v
	return s
}

func (s *ListProductsResponseBodyProducts) SetProductId(v string) *ListProductsResponseBodyProducts {
	s.ProductId = &v
	return s
}

func (s *ListProductsResponseBodyProducts) SetPurchaseUrl(v string) *ListProductsResponseBodyProducts {
	s.PurchaseUrl = &v
	return s
}

func (s *ListProductsResponseBodyProducts) Validate() error {
	return dara.Validate(s)
}

type ListProductsResponseBodyServices struct {
	// example:
	//
	// true
	IsOpen  *bool   `json:"IsOpen,omitempty" xml:"IsOpen,omitempty"`
	OpenUrl *string `json:"OpenUrl,omitempty" xml:"OpenUrl,omitempty"`
	// example:
	//
	// oss
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
}

func (s ListProductsResponseBodyServices) String() string {
	return dara.Prettify(s)
}

func (s ListProductsResponseBodyServices) GoString() string {
	return s.String()
}

func (s *ListProductsResponseBodyServices) GetIsOpen() *bool {
	return s.IsOpen
}

func (s *ListProductsResponseBodyServices) GetOpenUrl() *string {
	return s.OpenUrl
}

func (s *ListProductsResponseBodyServices) GetServiceCode() *string {
	return s.ServiceCode
}

func (s *ListProductsResponseBodyServices) SetIsOpen(v bool) *ListProductsResponseBodyServices {
	s.IsOpen = &v
	return s
}

func (s *ListProductsResponseBodyServices) SetOpenUrl(v string) *ListProductsResponseBodyServices {
	s.OpenUrl = &v
	return s
}

func (s *ListProductsResponseBodyServices) SetServiceCode(v string) *ListProductsResponseBodyServices {
	s.ServiceCode = &v
	return s
}

func (s *ListProductsResponseBodyServices) Validate() error {
	return dara.Validate(s)
}
