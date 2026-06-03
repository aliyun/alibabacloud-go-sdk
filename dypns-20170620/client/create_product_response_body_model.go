// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProductResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateProductResponseBody
	GetRequestId() *string
	SetCode(v string) *CreateProductResponseBody
	GetCode() *string
	SetData(v bool) *CreateProductResponseBody
	GetData() *bool
}

type CreateProductResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *bool   `json:"data,omitempty" xml:"data,omitempty"`
}

func (s CreateProductResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateProductResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProductResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateProductResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateProductResponseBody) GetData() *bool {
	return s.Data
}

func (s *CreateProductResponseBody) SetRequestId(v string) *CreateProductResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateProductResponseBody) SetCode(v string) *CreateProductResponseBody {
	s.Code = &v
	return s
}

func (s *CreateProductResponseBody) SetData(v bool) *CreateProductResponseBody {
	s.Data = &v
	return s
}

func (s *CreateProductResponseBody) Validate() error {
	return dara.Validate(s)
}
