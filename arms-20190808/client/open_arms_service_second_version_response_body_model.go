// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenArmsServiceSecondVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *OpenArmsServiceSecondVersionResponseBody
	GetOrderId() *string
	SetRequestId(v string) *OpenArmsServiceSecondVersionResponseBody
	GetRequestId() *string
}

type OpenArmsServiceSecondVersionResponseBody struct {
	// The service ID returned if the service is activated.
	//
	// example:
	//
	// 20896874992****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9ED50893-F3C4-42DF-ABB2-C200BE******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OpenArmsServiceSecondVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OpenArmsServiceSecondVersionResponseBody) GoString() string {
	return s.String()
}

func (s *OpenArmsServiceSecondVersionResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *OpenArmsServiceSecondVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OpenArmsServiceSecondVersionResponseBody) SetOrderId(v string) *OpenArmsServiceSecondVersionResponseBody {
	s.OrderId = &v
	return s
}

func (s *OpenArmsServiceSecondVersionResponseBody) SetRequestId(v string) *OpenArmsServiceSecondVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *OpenArmsServiceSecondVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
