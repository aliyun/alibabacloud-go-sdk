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
	// Error code returned. A value of `0` indicates success.
	//
	// example:
	//
	// 0
	RetCode *int32 `json:"RetCode,omitempty" xml:"RetCode,omitempty"`
	// Error message returned. No message is returned upon success.
	//
	// example:
	//
	// 请求异常
	RetMsg *string `json:"RetMsg,omitempty" xml:"RetMsg,omitempty"`
	// List of result objects returned.
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
	if s.RetValue != nil {
		if err := s.RetValue.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAccountForAppResponseBodyRetValue struct {
	// Indicates whether the user is an authorized member.
	//
	// example:
	//
	// true
	IsVip *bool `json:"IsVip,omitempty" xml:"IsVip,omitempty"`
	// Membership expiration time.
	//
	// example:
	//
	// 2022-05-12 15:22:18
	StrVipExpire *string `json:"StrVipExpire,omitempty" xml:"StrVipExpire,omitempty"`
	// Millisecond-level UNIX timestamp.
	//
	// - If the user has activated membership, returns the membership expiration time.
	//
	// - If the user has never activated membership, the value is always `0`.
	//
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
