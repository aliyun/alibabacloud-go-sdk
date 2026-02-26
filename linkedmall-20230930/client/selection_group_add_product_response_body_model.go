// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSelectionGroupAddProductResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetProductIds(v []*string) *SelectionGroupAddProductResponseBody
	GetProductIds() []*string
}

type SelectionGroupAddProductResponseBody struct {
	ProductIds []*string `json:"productIds,omitempty" xml:"productIds,omitempty" type:"Repeated"`
}

func (s SelectionGroupAddProductResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SelectionGroupAddProductResponseBody) GoString() string {
	return s.String()
}

func (s *SelectionGroupAddProductResponseBody) GetProductIds() []*string {
	return s.ProductIds
}

func (s *SelectionGroupAddProductResponseBody) SetProductIds(v []*string) *SelectionGroupAddProductResponseBody {
	s.ProductIds = v
	return s
}

func (s *SelectionGroupAddProductResponseBody) Validate() error {
	return dara.Validate(s)
}
