// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResizeNodeCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *ResizeNodeCountResponseBody
	GetOrderId() *string
	SetRequestId(v string) *ResizeNodeCountResponseBody
	GetRequestId() *string
}

type ResizeNodeCountResponseBody struct {
	// example:
	//
	// 20470860005****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// B288B41F-6681-42A6-8905-47C3C42B19B0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResizeNodeCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResizeNodeCountResponseBody) GoString() string {
	return s.String()
}

func (s *ResizeNodeCountResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *ResizeNodeCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResizeNodeCountResponseBody) SetOrderId(v string) *ResizeNodeCountResponseBody {
	s.OrderId = &v
	return s
}

func (s *ResizeNodeCountResponseBody) SetRequestId(v string) *ResizeNodeCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResizeNodeCountResponseBody) Validate() error {
	return dara.Validate(s)
}
