// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryInEffectQuthOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizCode(v string) *QueryInEffectQuthOrderRequest
	GetBizCode() *string
	SetChannel(v string) *QueryInEffectQuthOrderRequest
	GetChannel() *string
	SetLanguage(v string) *QueryInEffectQuthOrderRequest
	GetLanguage() *string
	SetOperatorTypeEnum(v string) *QueryInEffectQuthOrderRequest
	GetOperatorTypeEnum() *string
	SetRequestFromApp(v string) *QueryInEffectQuthOrderRequest
	GetRequestFromApp() *string
	SetRequestId(v string) *QueryInEffectQuthOrderRequest
	GetRequestId() *string
	SetRequestWay(v string) *QueryInEffectQuthOrderRequest
	GetRequestWay() *string
	SetUserNo(v string) *QueryInEffectQuthOrderRequest
	GetUserNo() *string
}

type QueryInEffectQuthOrderRequest struct {
	BizCode          *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	Channel          *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	Language         *string `json:"Language,omitempty" xml:"Language,omitempty"`
	OperatorTypeEnum *string `json:"OperatorTypeEnum,omitempty" xml:"OperatorTypeEnum,omitempty"`
	RequestFromApp   *string `json:"RequestFromApp,omitempty" xml:"RequestFromApp,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RequestWay       *string `json:"RequestWay,omitempty" xml:"RequestWay,omitempty"`
	UserNo           *string `json:"UserNo,omitempty" xml:"UserNo,omitempty"`
}

func (s QueryInEffectQuthOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryInEffectQuthOrderRequest) GoString() string {
	return s.String()
}

func (s *QueryInEffectQuthOrderRequest) GetBizCode() *string {
	return s.BizCode
}

func (s *QueryInEffectQuthOrderRequest) GetChannel() *string {
	return s.Channel
}

func (s *QueryInEffectQuthOrderRequest) GetLanguage() *string {
	return s.Language
}

func (s *QueryInEffectQuthOrderRequest) GetOperatorTypeEnum() *string {
	return s.OperatorTypeEnum
}

func (s *QueryInEffectQuthOrderRequest) GetRequestFromApp() *string {
	return s.RequestFromApp
}

func (s *QueryInEffectQuthOrderRequest) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryInEffectQuthOrderRequest) GetRequestWay() *string {
	return s.RequestWay
}

func (s *QueryInEffectQuthOrderRequest) GetUserNo() *string {
	return s.UserNo
}

func (s *QueryInEffectQuthOrderRequest) SetBizCode(v string) *QueryInEffectQuthOrderRequest {
	s.BizCode = &v
	return s
}

func (s *QueryInEffectQuthOrderRequest) SetChannel(v string) *QueryInEffectQuthOrderRequest {
	s.Channel = &v
	return s
}

func (s *QueryInEffectQuthOrderRequest) SetLanguage(v string) *QueryInEffectQuthOrderRequest {
	s.Language = &v
	return s
}

func (s *QueryInEffectQuthOrderRequest) SetOperatorTypeEnum(v string) *QueryInEffectQuthOrderRequest {
	s.OperatorTypeEnum = &v
	return s
}

func (s *QueryInEffectQuthOrderRequest) SetRequestFromApp(v string) *QueryInEffectQuthOrderRequest {
	s.RequestFromApp = &v
	return s
}

func (s *QueryInEffectQuthOrderRequest) SetRequestId(v string) *QueryInEffectQuthOrderRequest {
	s.RequestId = &v
	return s
}

func (s *QueryInEffectQuthOrderRequest) SetRequestWay(v string) *QueryInEffectQuthOrderRequest {
	s.RequestWay = &v
	return s
}

func (s *QueryInEffectQuthOrderRequest) SetUserNo(v string) *QueryInEffectQuthOrderRequest {
	s.UserNo = &v
	return s
}

func (s *QueryInEffectQuthOrderRequest) Validate() error {
	return dara.Validate(s)
}
