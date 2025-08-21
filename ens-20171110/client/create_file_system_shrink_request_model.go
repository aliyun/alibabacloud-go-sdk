// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFileSystemShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderDetailsShrink(v string) *CreateFileSystemShrinkRequest
	GetOrderDetailsShrink() *string
}

type CreateFileSystemShrinkRequest struct {
	// The information about the orders.
	//
	// This parameter is required.
	OrderDetailsShrink *string `json:"OrderDetails,omitempty" xml:"OrderDetails,omitempty"`
}

func (s CreateFileSystemShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFileSystemShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateFileSystemShrinkRequest) GetOrderDetailsShrink() *string {
	return s.OrderDetailsShrink
}

func (s *CreateFileSystemShrinkRequest) SetOrderDetailsShrink(v string) *CreateFileSystemShrinkRequest {
	s.OrderDetailsShrink = &v
	return s
}

func (s *CreateFileSystemShrinkRequest) Validate() error {
	return dara.Validate(s)
}
