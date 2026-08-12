// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConvertHybridInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrCode(v string) *ConvertHybridInstanceResponseBody
	GetErrCode() *string
	SetOrderInfo(v *ConvertHybridInstanceResponseBodyOrderInfo) *ConvertHybridInstanceResponseBody
	GetOrderInfo() *ConvertHybridInstanceResponseBodyOrderInfo
	SetRequestId(v string) *ConvertHybridInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ConvertHybridInstanceResponseBody
	GetSuccess() *bool
}

type ConvertHybridInstanceResponseBody struct {
	// The error code.
	//
	// example:
	//
	// 000000
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The order information.
	OrderInfo *ConvertHybridInstanceResponseBodyOrderInfo `json:"OrderInfo,omitempty" xml:"OrderInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 67F33190-946B-1105-B6A1-E2DF042*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ConvertHybridInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConvertHybridInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ConvertHybridInstanceResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ConvertHybridInstanceResponseBody) GetOrderInfo() *ConvertHybridInstanceResponseBodyOrderInfo {
	return s.OrderInfo
}

func (s *ConvertHybridInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConvertHybridInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ConvertHybridInstanceResponseBody) SetErrCode(v string) *ConvertHybridInstanceResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ConvertHybridInstanceResponseBody) SetOrderInfo(v *ConvertHybridInstanceResponseBodyOrderInfo) *ConvertHybridInstanceResponseBody {
	s.OrderInfo = v
	return s
}

func (s *ConvertHybridInstanceResponseBody) SetRequestId(v string) *ConvertHybridInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConvertHybridInstanceResponseBody) SetSuccess(v bool) *ConvertHybridInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *ConvertHybridInstanceResponseBody) Validate() error {
	if s.OrderInfo != nil {
		if err := s.OrderInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ConvertHybridInstanceResponseBodyOrderInfo struct {
	// The instance ID of the pay-as-you-go portion of hybrid billing.
	//
	// example:
	//
	// f-cn-asd***
	ElasticInstanceId *string `json:"ElasticInstanceId,omitempty" xml:"ElasticInstanceId,omitempty"`
	// The instance ID of the subscription portion of hybrid billing.
	//
	// example:
	//
	// f-cn-zvp2q*****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The order ID.
	//
	// example:
	//
	// 2104063546****
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s ConvertHybridInstanceResponseBodyOrderInfo) String() string {
	return dara.Prettify(s)
}

func (s ConvertHybridInstanceResponseBodyOrderInfo) GoString() string {
	return s.String()
}

func (s *ConvertHybridInstanceResponseBodyOrderInfo) GetElasticInstanceId() *string {
	return s.ElasticInstanceId
}

func (s *ConvertHybridInstanceResponseBodyOrderInfo) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ConvertHybridInstanceResponseBodyOrderInfo) GetOrderId() *int64 {
	return s.OrderId
}

func (s *ConvertHybridInstanceResponseBodyOrderInfo) SetElasticInstanceId(v string) *ConvertHybridInstanceResponseBodyOrderInfo {
	s.ElasticInstanceId = &v
	return s
}

func (s *ConvertHybridInstanceResponseBodyOrderInfo) SetInstanceId(v string) *ConvertHybridInstanceResponseBodyOrderInfo {
	s.InstanceId = &v
	return s
}

func (s *ConvertHybridInstanceResponseBodyOrderInfo) SetOrderId(v int64) *ConvertHybridInstanceResponseBodyOrderInfo {
	s.OrderId = &v
	return s
}

func (s *ConvertHybridInstanceResponseBodyOrderInfo) Validate() error {
	return dara.Validate(s)
}
