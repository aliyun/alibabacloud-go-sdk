// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSelectionGroupRemoveProductResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetProductIds(v []*string) *SelectionGroupRemoveProductResponseBody
	GetProductIds() []*string
}

type SelectionGroupRemoveProductResponseBody struct {
	ProductIds []*string `json:"productIds,omitempty" xml:"productIds,omitempty" type:"Repeated"`
}

func (s SelectionGroupRemoveProductResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SelectionGroupRemoveProductResponseBody) GoString() string {
	return s.String()
}

func (s *SelectionGroupRemoveProductResponseBody) GetProductIds() []*string {
	return s.ProductIds
}

func (s *SelectionGroupRemoveProductResponseBody) SetProductIds(v []*string) *SelectionGroupRemoveProductResponseBody {
	s.ProductIds = v
	return s
}

func (s *SelectionGroupRemoveProductResponseBody) Validate() error {
	return dara.Validate(s)
}
