// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenFlinkAiServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderInfo(v *OpenFlinkAiServiceResponseBodyOrderInfo) *OpenFlinkAiServiceResponseBody
	GetOrderInfo() *OpenFlinkAiServiceResponseBodyOrderInfo
	SetRequestId(v string) *OpenFlinkAiServiceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *OpenFlinkAiServiceResponseBody
	GetSuccess() *bool
}

type OpenFlinkAiServiceResponseBody struct {
	// The order information.
	OrderInfo *OpenFlinkAiServiceResponseBodyOrderInfo `json:"OrderInfo,omitempty" xml:"OrderInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 67F33190-946B-1105-B6A1-E2DF0426DD51
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s OpenFlinkAiServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OpenFlinkAiServiceResponseBody) GoString() string {
	return s.String()
}

func (s *OpenFlinkAiServiceResponseBody) GetOrderInfo() *OpenFlinkAiServiceResponseBodyOrderInfo {
	return s.OrderInfo
}

func (s *OpenFlinkAiServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OpenFlinkAiServiceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *OpenFlinkAiServiceResponseBody) SetOrderInfo(v *OpenFlinkAiServiceResponseBodyOrderInfo) *OpenFlinkAiServiceResponseBody {
	s.OrderInfo = v
	return s
}

func (s *OpenFlinkAiServiceResponseBody) SetRequestId(v string) *OpenFlinkAiServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *OpenFlinkAiServiceResponseBody) SetSuccess(v bool) *OpenFlinkAiServiceResponseBody {
	s.Success = &v
	return s
}

func (s *OpenFlinkAiServiceResponseBody) Validate() error {
	if s.OrderInfo != nil {
		if err := s.OrderInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OpenFlinkAiServiceResponseBodyOrderInfo struct {
	// The instance ID.
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

func (s OpenFlinkAiServiceResponseBodyOrderInfo) String() string {
	return dara.Prettify(s)
}

func (s OpenFlinkAiServiceResponseBodyOrderInfo) GoString() string {
	return s.String()
}

func (s *OpenFlinkAiServiceResponseBodyOrderInfo) GetInstanceId() *string {
	return s.InstanceId
}

func (s *OpenFlinkAiServiceResponseBodyOrderInfo) GetOrderId() *int64 {
	return s.OrderId
}

func (s *OpenFlinkAiServiceResponseBodyOrderInfo) SetInstanceId(v string) *OpenFlinkAiServiceResponseBodyOrderInfo {
	s.InstanceId = &v
	return s
}

func (s *OpenFlinkAiServiceResponseBodyOrderInfo) SetOrderId(v int64) *OpenFlinkAiServiceResponseBodyOrderInfo {
	s.OrderId = &v
	return s
}

func (s *OpenFlinkAiServiceResponseBodyOrderInfo) Validate() error {
	return dara.Validate(s)
}
