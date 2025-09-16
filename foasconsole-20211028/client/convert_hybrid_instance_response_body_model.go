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
	// example:
	//
	// 000000
	ErrCode   *string                                     `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	OrderInfo *ConvertHybridInstanceResponseBodyOrderInfo `json:"OrderInfo,omitempty" xml:"OrderInfo,omitempty" type:"Struct"`
	// example:
	//
	// 67F33190-946B-1105-B6A1-E2DF0426DD51
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	return dara.Validate(s)
}

type ConvertHybridInstanceResponseBodyOrderInfo struct {
	ElasticInstanceId *string `json:"ElasticInstanceId,omitempty" xml:"ElasticInstanceId,omitempty"`
	// example:
	//
	// f-cn-zvp2q0zik06
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 210406354694567
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
