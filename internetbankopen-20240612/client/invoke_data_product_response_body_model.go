// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvokeDataProductResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *InvokeDataProductResponseBody
	GetData() *string
	SetRequestId(v string) *InvokeDataProductResponseBody
	GetRequestId() *string
}

type InvokeDataProductResponseBody struct {
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s InvokeDataProductResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InvokeDataProductResponseBody) GoString() string {
	return s.String()
}

func (s *InvokeDataProductResponseBody) GetData() *string {
	return s.Data
}

func (s *InvokeDataProductResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InvokeDataProductResponseBody) SetData(v string) *InvokeDataProductResponseBody {
	s.Data = &v
	return s
}

func (s *InvokeDataProductResponseBody) SetRequestId(v string) *InvokeDataProductResponseBody {
	s.RequestId = &v
	return s
}

func (s *InvokeDataProductResponseBody) Validate() error {
	return dara.Validate(s)
}
