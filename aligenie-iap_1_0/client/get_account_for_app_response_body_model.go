// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccountForAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRetCode(v int32) *GetAccountForAppResponseBody
	GetRetCode() *int32
	SetRetMsg(v string) *GetAccountForAppResponseBody
	GetRetMsg() *string
	SetRetValue(v *GetAccountForAppResponseBodyRetValue) *GetAccountForAppResponseBody
	GetRetValue() *GetAccountForAppResponseBodyRetValue
}

type GetAccountForAppResponseBody struct {
	// example:
	//
	// 0
	RetCode  *int32                                `json:"RetCode,omitempty" xml:"RetCode,omitempty"`
	RetMsg   *string                               `json:"RetMsg,omitempty" xml:"RetMsg,omitempty"`
	RetValue *GetAccountForAppResponseBodyRetValue `json:"RetValue,omitempty" xml:"RetValue,omitempty" type:"Struct"`
}

func (s GetAccountForAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAccountForAppResponseBody) GoString() string {
	return s.String()
}

func (s *GetAccountForAppResponseBody) GetRetCode() *int32 {
	return s.RetCode
}

func (s *GetAccountForAppResponseBody) GetRetMsg() *string {
	return s.RetMsg
}

func (s *GetAccountForAppResponseBody) GetRetValue() *GetAccountForAppResponseBodyRetValue {
	return s.RetValue
}

func (s *GetAccountForAppResponseBody) SetRetCode(v int32) *GetAccountForAppResponseBody {
	s.RetCode = &v
	return s
}

func (s *GetAccountForAppResponseBody) SetRetMsg(v string) *GetAccountForAppResponseBody {
	s.RetMsg = &v
	return s
}

func (s *GetAccountForAppResponseBody) SetRetValue(v *GetAccountForAppResponseBodyRetValue) *GetAccountForAppResponseBody {
	s.RetValue = v
	return s
}

func (s *GetAccountForAppResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetAccountForAppResponseBodyRetValue struct {
	// example:
	//
	// true
	IsVip *bool `json:"IsVip,omitempty" xml:"IsVip,omitempty"`
	// example:
	//
	// 2022-05-12 15:22:18
	StrVipExpire *string `json:"StrVipExpire,omitempty" xml:"StrVipExpire,omitempty"`
	// example:
	//
	// 1652340138347
	VipExpireAt *int64 `json:"VipExpireAt,omitempty" xml:"VipExpireAt,omitempty"`
}

func (s GetAccountForAppResponseBodyRetValue) String() string {
	return dara.Prettify(s)
}

func (s GetAccountForAppResponseBodyRetValue) GoString() string {
	return s.String()
}

func (s *GetAccountForAppResponseBodyRetValue) GetIsVip() *bool {
	return s.IsVip
}

func (s *GetAccountForAppResponseBodyRetValue) GetStrVipExpire() *string {
	return s.StrVipExpire
}

func (s *GetAccountForAppResponseBodyRetValue) GetVipExpireAt() *int64 {
	return s.VipExpireAt
}

func (s *GetAccountForAppResponseBodyRetValue) SetIsVip(v bool) *GetAccountForAppResponseBodyRetValue {
	s.IsVip = &v
	return s
}

func (s *GetAccountForAppResponseBodyRetValue) SetStrVipExpire(v string) *GetAccountForAppResponseBodyRetValue {
	s.StrVipExpire = &v
	return s
}

func (s *GetAccountForAppResponseBodyRetValue) SetVipExpireAt(v int64) *GetAccountForAppResponseBodyRetValue {
	s.VipExpireAt = &v
	return s
}

func (s *GetAccountForAppResponseBodyRetValue) Validate() error {
	return dara.Validate(s)
}
