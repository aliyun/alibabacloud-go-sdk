// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBusAppConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRetCode(v int32) *GetBusAppConfigResponseBody
	GetRetCode() *int32
	SetRetMsg(v string) *GetBusAppConfigResponseBody
	GetRetMsg() *string
	SetRetValue(v *GetBusAppConfigResponseBodyRetValue) *GetBusAppConfigResponseBody
	GetRetValue() *GetBusAppConfigResponseBodyRetValue
}

type GetBusAppConfigResponseBody struct {
	// example:
	//
	// 0
	RetCode *int32 `json:"RetCode,omitempty" xml:"RetCode,omitempty"`
	// example:
	//
	// 请求异常
	RetMsg   *string                              `json:"RetMsg,omitempty" xml:"RetMsg,omitempty"`
	RetValue *GetBusAppConfigResponseBodyRetValue `json:"RetValue,omitempty" xml:"RetValue,omitempty" type:"Struct"`
}

func (s GetBusAppConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBusAppConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetBusAppConfigResponseBody) GetRetCode() *int32 {
	return s.RetCode
}

func (s *GetBusAppConfigResponseBody) GetRetMsg() *string {
	return s.RetMsg
}

func (s *GetBusAppConfigResponseBody) GetRetValue() *GetBusAppConfigResponseBodyRetValue {
	return s.RetValue
}

func (s *GetBusAppConfigResponseBody) SetRetCode(v int32) *GetBusAppConfigResponseBody {
	s.RetCode = &v
	return s
}

func (s *GetBusAppConfigResponseBody) SetRetMsg(v string) *GetBusAppConfigResponseBody {
	s.RetMsg = &v
	return s
}

func (s *GetBusAppConfigResponseBody) SetRetValue(v *GetBusAppConfigResponseBodyRetValue) *GetBusAppConfigResponseBody {
	s.RetValue = v
	return s
}

func (s *GetBusAppConfigResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetBusAppConfigResponseBodyRetValue struct {
	Cashier *string `json:"Cashier,omitempty" xml:"Cashier,omitempty"`
	// example:
	//
	// https://******.com/design/******?imageId=xxxxx
	ShoppingBar *string `json:"ShoppingBar,omitempty" xml:"ShoppingBar,omitempty"`
	// example:
	//
	// https://******.com/design/******?imageId=xxxxx
	ShoppingWindow *string `json:"ShoppingWindow,omitempty" xml:"ShoppingWindow,omitempty"`
	// example:
	//
	// https://******.com/design/******?imageId=xxxxx
	VipLabel *string `json:"VipLabel,omitempty" xml:"VipLabel,omitempty"`
}

func (s GetBusAppConfigResponseBodyRetValue) String() string {
	return dara.Prettify(s)
}

func (s GetBusAppConfigResponseBodyRetValue) GoString() string {
	return s.String()
}

func (s *GetBusAppConfigResponseBodyRetValue) GetCashier() *string {
	return s.Cashier
}

func (s *GetBusAppConfigResponseBodyRetValue) GetShoppingBar() *string {
	return s.ShoppingBar
}

func (s *GetBusAppConfigResponseBodyRetValue) GetShoppingWindow() *string {
	return s.ShoppingWindow
}

func (s *GetBusAppConfigResponseBodyRetValue) GetVipLabel() *string {
	return s.VipLabel
}

func (s *GetBusAppConfigResponseBodyRetValue) SetCashier(v string) *GetBusAppConfigResponseBodyRetValue {
	s.Cashier = &v
	return s
}

func (s *GetBusAppConfigResponseBodyRetValue) SetShoppingBar(v string) *GetBusAppConfigResponseBodyRetValue {
	s.ShoppingBar = &v
	return s
}

func (s *GetBusAppConfigResponseBodyRetValue) SetShoppingWindow(v string) *GetBusAppConfigResponseBodyRetValue {
	s.ShoppingWindow = &v
	return s
}

func (s *GetBusAppConfigResponseBodyRetValue) SetVipLabel(v string) *GetBusAppConfigResponseBodyRetValue {
	s.VipLabel = &v
	return s
}

func (s *GetBusAppConfigResponseBodyRetValue) Validate() error {
	return dara.Validate(s)
}
