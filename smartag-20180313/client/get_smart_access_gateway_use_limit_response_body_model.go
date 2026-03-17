// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSmartAccessGatewayUseLimitResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetSmartAccessGatewayUseLimitResponseBody
	GetRequestId() *string
	SetTotalAmount(v int32) *GetSmartAccessGatewayUseLimitResponseBody
	GetTotalAmount() *int32
	SetUsedAmount(v int32) *GetSmartAccessGatewayUseLimitResponseBody
	GetUsedAmount() *int32
}

type GetSmartAccessGatewayUseLimitResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 2265DB11-F5CC-496E-ADE7-D043AC37926A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of SAG instances that you can purchase.
	//
	// example:
	//
	// 500
	TotalAmount *int32 `json:"TotalAmount,omitempty" xml:"TotalAmount,omitempty"`
	// The number of SAG instances that you have purchased.
	//
	// example:
	//
	// 47
	UsedAmount *int32 `json:"UsedAmount,omitempty" xml:"UsedAmount,omitempty"`
}

func (s GetSmartAccessGatewayUseLimitResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSmartAccessGatewayUseLimitResponseBody) GoString() string {
	return s.String()
}

func (s *GetSmartAccessGatewayUseLimitResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSmartAccessGatewayUseLimitResponseBody) GetTotalAmount() *int32 {
	return s.TotalAmount
}

func (s *GetSmartAccessGatewayUseLimitResponseBody) GetUsedAmount() *int32 {
	return s.UsedAmount
}

func (s *GetSmartAccessGatewayUseLimitResponseBody) SetRequestId(v string) *GetSmartAccessGatewayUseLimitResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSmartAccessGatewayUseLimitResponseBody) SetTotalAmount(v int32) *GetSmartAccessGatewayUseLimitResponseBody {
	s.TotalAmount = &v
	return s
}

func (s *GetSmartAccessGatewayUseLimitResponseBody) SetUsedAmount(v int32) *GetSmartAccessGatewayUseLimitResponseBody {
	s.UsedAmount = &v
	return s
}

func (s *GetSmartAccessGatewayUseLimitResponseBody) Validate() error {
	return dara.Validate(s)
}
