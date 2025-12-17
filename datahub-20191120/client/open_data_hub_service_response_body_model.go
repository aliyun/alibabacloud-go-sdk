// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenDataHubServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *OpenDataHubServiceResponseBody
	GetOrderId() *string
	SetRequestId(v string) *OpenDataHubServiceResponseBody
	GetRequestId() *string
}

type OpenDataHubServiceResponseBody struct {
	OrderId   *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OpenDataHubServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OpenDataHubServiceResponseBody) GoString() string {
	return s.String()
}

func (s *OpenDataHubServiceResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *OpenDataHubServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OpenDataHubServiceResponseBody) SetOrderId(v string) *OpenDataHubServiceResponseBody {
	s.OrderId = &v
	return s
}

func (s *OpenDataHubServiceResponseBody) SetRequestId(v string) *OpenDataHubServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *OpenDataHubServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
