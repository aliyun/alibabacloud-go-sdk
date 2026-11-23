// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetError(v *OpenServiceResponseBodyError) *OpenServiceResponseBody
	GetError() *OpenServiceResponseBodyError
	SetOrderId(v string) *OpenServiceResponseBody
	GetOrderId() *string
	SetSpInstanceId(v string) *OpenServiceResponseBody
	GetSpInstanceId() *string
	SetStatus(v string) *OpenServiceResponseBody
	GetStatus() *string
}

type OpenServiceResponseBody struct {
	// 网关拒绝或处理失败时返回。成功响应没有此对象；平台在抵达业务网关前产生的错误遵循 OpenAPI 公共错误结构。
	Error *OpenServiceResponseBodyError `json:"error,omitempty" xml:"error,omitempty" type:"Struct"`
	// OPENING 或 ACTIVE 时返回的订单号；其他状态省略，无默认值。
	//
	// example:
	//
	// order-example-001
	OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
	// ACTIVE 时返回的 SP 实例 ID；其他状态省略，无默认值。
	//
	// example:
	//
	// sp-example-001
	SpInstanceId *string `json:"spInstanceId,omitempty" xml:"spInstanceId,omitempty"`
	// 正常响应必填。TEST_OK：接口验证通过，未发起实际开通；OPENING：已取得订单号，开通仍在处理中；ACTIVE：服务已开通，并返回订单号和实例 ID；ORDER_UNKNOWN：本次开通结果暂不能确认，不表示必定失败；REJECTED：本次因可用额度不足明确拒绝且未返回订单。HTTP 200 不等于已开通。
	//
	// example:
	//
	// TEST_OK
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s OpenServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OpenServiceResponseBody) GoString() string {
	return s.String()
}

func (s *OpenServiceResponseBody) GetError() *OpenServiceResponseBodyError {
	return s.Error
}

func (s *OpenServiceResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *OpenServiceResponseBody) GetSpInstanceId() *string {
	return s.SpInstanceId
}

func (s *OpenServiceResponseBody) GetStatus() *string {
	return s.Status
}

func (s *OpenServiceResponseBody) SetError(v *OpenServiceResponseBodyError) *OpenServiceResponseBody {
	s.Error = v
	return s
}

func (s *OpenServiceResponseBody) SetOrderId(v string) *OpenServiceResponseBody {
	s.OrderId = &v
	return s
}

func (s *OpenServiceResponseBody) SetSpInstanceId(v string) *OpenServiceResponseBody {
	s.SpInstanceId = &v
	return s
}

func (s *OpenServiceResponseBody) SetStatus(v string) *OpenServiceResponseBody {
	s.Status = &v
	return s
}

func (s *OpenServiceResponseBody) Validate() error {
	if s.Error != nil {
		if err := s.Error.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OpenServiceResponseBodyError struct {
	// 错误码列表见接口说明。请求校验、鉴权、授权、配额及开通处理错误都使用此位置。
	//
	// This parameter is required.
	//
	// example:
	//
	// INVALID_REQUEST
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 对应错误的可读说明，不应依赖固定文案判断。开通处理错误可能返回 service activation failed。
	//
	// This parameter is required.
	//
	// example:
	//
	// invalid request
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 用于排查错误；通常对应 X-Request-Id。客户关联号缺失或非法时可返回服务端诊断号，不是订单号。
	//
	// This parameter is required.
	//
	// example:
	//
	// open-service-001
	RequestId *string `json:"request_id,omitempty" xml:"request_id,omitempty"`
}

func (s OpenServiceResponseBodyError) String() string {
	return dara.Prettify(s)
}

func (s OpenServiceResponseBodyError) GoString() string {
	return s.String()
}

func (s *OpenServiceResponseBodyError) GetCode() *string {
	return s.Code
}

func (s *OpenServiceResponseBodyError) GetMessage() *string {
	return s.Message
}

func (s *OpenServiceResponseBodyError) GetRequestId() *string {
	return s.RequestId
}

func (s *OpenServiceResponseBodyError) SetCode(v string) *OpenServiceResponseBodyError {
	s.Code = &v
	return s
}

func (s *OpenServiceResponseBodyError) SetMessage(v string) *OpenServiceResponseBodyError {
	s.Message = &v
	return s
}

func (s *OpenServiceResponseBodyError) SetRequestId(v string) *OpenServiceResponseBodyError {
	s.RequestId = &v
	return s
}

func (s *OpenServiceResponseBodyError) Validate() error {
	return dara.Validate(s)
}
