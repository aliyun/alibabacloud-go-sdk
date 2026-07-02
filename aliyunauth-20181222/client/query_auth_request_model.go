// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAuthRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizNo(v string) *QueryAuthRequest
	GetBizNo() *string
	SetChannel(v string) *QueryAuthRequest
	GetChannel() *string
	SetInstanceId(v string) *QueryAuthRequest
	GetInstanceId() *string
	SetLanguage(v string) *QueryAuthRequest
	GetLanguage() *string
	SetOperatorTypeEnum(v string) *QueryAuthRequest
	GetOperatorTypeEnum() *string
	SetOrderVid(v string) *QueryAuthRequest
	GetOrderVid() *string
	SetProductCode(v string) *QueryAuthRequest
	GetProductCode() *string
	SetRequestFromApp(v string) *QueryAuthRequest
	GetRequestFromApp() *string
	SetRequestWay(v string) *QueryAuthRequest
	GetRequestWay() *string
	SetUserNo(v string) *QueryAuthRequest
	GetUserNo() *string
}

type QueryAuthRequest struct {
	BizNo            *string `json:"BizNo,omitempty" xml:"BizNo,omitempty"`
	Channel          *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Language         *string `json:"Language,omitempty" xml:"Language,omitempty"`
	OperatorTypeEnum *string `json:"OperatorTypeEnum,omitempty" xml:"OperatorTypeEnum,omitempty"`
	OrderVid         *string `json:"OrderVid,omitempty" xml:"OrderVid,omitempty"`
	ProductCode      *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	RequestFromApp   *string `json:"RequestFromApp,omitempty" xml:"RequestFromApp,omitempty"`
	RequestWay       *string `json:"RequestWay,omitempty" xml:"RequestWay,omitempty"`
	UserNo           *string `json:"UserNo,omitempty" xml:"UserNo,omitempty"`
}

func (s QueryAuthRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryAuthRequest) GoString() string {
	return s.String()
}

func (s *QueryAuthRequest) GetBizNo() *string {
	return s.BizNo
}

func (s *QueryAuthRequest) GetChannel() *string {
	return s.Channel
}

func (s *QueryAuthRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryAuthRequest) GetLanguage() *string {
	return s.Language
}

func (s *QueryAuthRequest) GetOperatorTypeEnum() *string {
	return s.OperatorTypeEnum
}

func (s *QueryAuthRequest) GetOrderVid() *string {
	return s.OrderVid
}

func (s *QueryAuthRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *QueryAuthRequest) GetRequestFromApp() *string {
	return s.RequestFromApp
}

func (s *QueryAuthRequest) GetRequestWay() *string {
	return s.RequestWay
}

func (s *QueryAuthRequest) GetUserNo() *string {
	return s.UserNo
}

func (s *QueryAuthRequest) SetBizNo(v string) *QueryAuthRequest {
	s.BizNo = &v
	return s
}

func (s *QueryAuthRequest) SetChannel(v string) *QueryAuthRequest {
	s.Channel = &v
	return s
}

func (s *QueryAuthRequest) SetInstanceId(v string) *QueryAuthRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryAuthRequest) SetLanguage(v string) *QueryAuthRequest {
	s.Language = &v
	return s
}

func (s *QueryAuthRequest) SetOperatorTypeEnum(v string) *QueryAuthRequest {
	s.OperatorTypeEnum = &v
	return s
}

func (s *QueryAuthRequest) SetOrderVid(v string) *QueryAuthRequest {
	s.OrderVid = &v
	return s
}

func (s *QueryAuthRequest) SetProductCode(v string) *QueryAuthRequest {
	s.ProductCode = &v
	return s
}

func (s *QueryAuthRequest) SetRequestFromApp(v string) *QueryAuthRequest {
	s.RequestFromApp = &v
	return s
}

func (s *QueryAuthRequest) SetRequestWay(v string) *QueryAuthRequest {
	s.RequestWay = &v
	return s
}

func (s *QueryAuthRequest) SetUserNo(v string) *QueryAuthRequest {
	s.UserNo = &v
	return s
}

func (s *QueryAuthRequest) Validate() error {
	return dara.Validate(s)
}
