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
	// The operation type for the bulk operation. This parameter specifies whether to add the API products specified by productIds to the category or remove them from the category. The valid values are not declared in the API definition.
	//
	// example:
	//
	// action-xxxxxxxx
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// The list of API product IDs to be processed. Each element is the primary ID of an API product resource.
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
