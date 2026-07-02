// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthenticateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *AuthenticateRequest
	GetInstanceId() *string
	SetLanguage(v string) *AuthenticateRequest
	GetLanguage() *string
	SetOperateCode(v string) *AuthenticateRequest
	GetOperateCode() *string
	SetOperatorTypeEnum(v string) *AuthenticateRequest
	GetOperatorTypeEnum() *string
	SetProductCode(v string) *AuthenticateRequest
	GetProductCode() *string
	SetRequestFromApp(v string) *AuthenticateRequest
	GetRequestFromApp() *string
	SetRequestWay(v string) *AuthenticateRequest
	GetRequestWay() *string
	SetUserNo(v string) *AuthenticateRequest
	GetUserNo() *string
}

type AuthenticateRequest struct {
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Language         *string `json:"Language,omitempty" xml:"Language,omitempty"`
	OperateCode      *string `json:"OperateCode,omitempty" xml:"OperateCode,omitempty"`
	OperatorTypeEnum *string `json:"OperatorTypeEnum,omitempty" xml:"OperatorTypeEnum,omitempty"`
	ProductCode      *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	RequestFromApp   *string `json:"RequestFromApp,omitempty" xml:"RequestFromApp,omitempty"`
	RequestWay       *string `json:"RequestWay,omitempty" xml:"RequestWay,omitempty"`
	UserNo           *string `json:"UserNo,omitempty" xml:"UserNo,omitempty"`
}

func (s AuthenticateRequest) String() string {
	return dara.Prettify(s)
}

func (s AuthenticateRequest) GoString() string {
	return s.String()
}

func (s *AuthenticateRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AuthenticateRequest) GetLanguage() *string {
	return s.Language
}

func (s *AuthenticateRequest) GetOperateCode() *string {
	return s.OperateCode
}

func (s *AuthenticateRequest) GetOperatorTypeEnum() *string {
	return s.OperatorTypeEnum
}

func (s *AuthenticateRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *AuthenticateRequest) GetRequestFromApp() *string {
	return s.RequestFromApp
}

func (s *AuthenticateRequest) GetRequestWay() *string {
	return s.RequestWay
}

func (s *AuthenticateRequest) GetUserNo() *string {
	return s.UserNo
}

func (s *AuthenticateRequest) SetInstanceId(v string) *AuthenticateRequest {
	s.InstanceId = &v
	return s
}

func (s *AuthenticateRequest) SetLanguage(v string) *AuthenticateRequest {
	s.Language = &v
	return s
}

func (s *AuthenticateRequest) SetOperateCode(v string) *AuthenticateRequest {
	s.OperateCode = &v
	return s
}

func (s *AuthenticateRequest) SetOperatorTypeEnum(v string) *AuthenticateRequest {
	s.OperatorTypeEnum = &v
	return s
}

func (s *AuthenticateRequest) SetProductCode(v string) *AuthenticateRequest {
	s.ProductCode = &v
	return s
}

func (s *AuthenticateRequest) SetRequestFromApp(v string) *AuthenticateRequest {
	s.RequestFromApp = &v
	return s
}

func (s *AuthenticateRequest) SetRequestWay(v string) *AuthenticateRequest {
	s.RequestWay = &v
	return s
}

func (s *AuthenticateRequest) SetUserNo(v string) *AuthenticateRequest {
	s.UserNo = &v
	return s
}

func (s *AuthenticateRequest) Validate() error {
	return dara.Validate(s)
}
