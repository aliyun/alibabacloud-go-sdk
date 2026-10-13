// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCategoryProductsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAction(v string) *UpdateCategoryProductsRequest
	GetAction() *string
	SetProductIds(v []*string) *UpdateCategoryProductsRequest
	GetProductIds() []*string
}

type UpdateCategoryProductsRequest struct {
	// The direction of this bulk operation: adds the API products specified in productIds to the category, or removes these API products from the category. The set of valid values is not declared in the API definition.
	//
	// example:
	//
	// action-xxxxxxxx
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// The list of API product IDs to process. Each element is a primary ID of the API product resource.
	ProductIds []*string `json:"productIds,omitempty" xml:"productIds,omitempty" type:"Repeated"`
}

func (s UpdateCategoryProductsRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCategoryProductsRequest) GoString() string {
	return s.String()
}

func (s *UpdateCategoryProductsRequest) GetAction() *string {
	return s.Action
}

func (s *UpdateCategoryProductsRequest) GetProductIds() []*string {
	return s.ProductIds
}

func (s *UpdateCategoryProductsRequest) SetAction(v string) *UpdateCategoryProductsRequest {
	s.Action = &v
	return s
}

func (s *UpdateCategoryProductsRequest) SetProductIds(v []*string) *UpdateCategoryProductsRequest {
	s.ProductIds = v
	return s
}

func (s *UpdateCategoryProductsRequest) Validate() error {
	return dara.Validate(s)
}
