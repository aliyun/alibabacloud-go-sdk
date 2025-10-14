// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProductResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateProductResponseBody
	GetRequestId() *string
}

type UpdateProductResponseBody struct {
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateProductResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateProductResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateProductResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateProductResponseBody) SetRequestId(v string) *UpdateProductResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateProductResponseBody) Validate() error {
	return dara.Validate(s)
}
