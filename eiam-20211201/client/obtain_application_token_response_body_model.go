// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iObtainApplicationTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationToken(v *ObtainApplicationTokenResponseBodyApplicationToken) *ObtainApplicationTokenResponseBody
	GetApplicationToken() *ObtainApplicationTokenResponseBodyApplicationToken
	SetRequestId(v string) *ObtainApplicationTokenResponseBody
	GetRequestId() *string
}

type ObtainApplicationTokenResponseBody struct {
	ApplicationToken *ObtainApplicationTokenResponseBodyApplicationToken `json:"ApplicationToken,omitempty" xml:"ApplicationToken,omitempty" type:"Struct"`
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ObtainApplicationTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ObtainApplicationTokenResponseBody) GoString() string {
	return s.String()
}

func (s *ObtainApplicationTokenResponseBody) GetApplicationToken() *ObtainApplicationTokenResponseBodyApplicationToken {
	return s.ApplicationToken
}

func (s *ObtainApplicationTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ObtainApplicationTokenResponseBody) SetApplicationToken(v *ObtainApplicationTokenResponseBodyApplicationToken) *ObtainApplicationTokenResponseBody {
	s.ApplicationToken = v
	return s
}

func (s *ObtainApplicationTokenResponseBody) SetRequestId(v string) *ObtainApplicationTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *ObtainApplicationTokenResponseBody) Validate() error {
	if s.ApplicationToken != nil {
		if err := s.ApplicationToken.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ObtainApplicationTokenResponseBodyApplicationToken struct {
	// IDaaS EIAM 应用Id
	//
	// example:
	//
	// app_na2r73a65s7o4zbs7nj5gxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// 客户端密钥
	//
	// example:
	//
	// SATFwqX8zxGf83pJcJw78KFGjmrft4erWeZYBGS8oE7NN6qoE217yaJpUdMb1UuuGqhDiF43sCA4CF91CTL5iGntqwyLuaAcS9FJ9HfGadE5a7TjiwVafwrBxxxxx
	ApplicationToken *string `json:"ApplicationToken,omitempty" xml:"ApplicationToken,omitempty"`
	// IDaaS EIAM 客户端ID
	//
	// example:
	//
	// token_m7aso6v4efvu2otfq3jdzxxxx
	ApplicationTokenId *string `json:"ApplicationTokenId,omitempty" xml:"ApplicationTokenId,omitempty"`
	// IDaaS EIAM 客户端密钥Id
	//
	// example:
	//
	// bearer_token
	ApplicationTokenType *string `json:"ApplicationTokenType,omitempty" xml:"ApplicationTokenType,omitempty"`
	// example:
	//
	// 1735610930000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1735610950000
	ExpirationTime *int64 `json:"ExpirationTime,omitempty" xml:"ExpirationTime,omitempty"`
	// IDaaS EIAM 实例Id
	//
	// example:
	//
	// idaas_ki6hd7ihir4ybawogqk6kqxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// IDaaS EIAM 客户端密钥最近使用时间
	//
	// example:
	//
	// 1735610930000
	LastUsedTime *int64 `json:"LastUsedTime,omitempty" xml:"LastUsedTime,omitempty"`
	// IDaaS EIAM 客户端密钥状态
	//
	// example:
	//
	// enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ObtainApplicationTokenResponseBodyApplicationToken) String() string {
	return dara.Prettify(s)
}

func (s ObtainApplicationTokenResponseBodyApplicationToken) GoString() string {
	return s.String()
}

func (s *ObtainApplicationTokenResponseBodyApplicationToken) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ObtainApplicationTokenResponseBodyApplicationToken) GetApplicationToken() *string {
	return s.ApplicationToken
}

func (s *ObtainApplicationTokenResponseBodyApplicationToken) GetApplicationTokenId() *string {
	return s.ApplicationTokenId
}

func (s *ObtainApplicationTokenResponseBodyApplicationToken) GetApplicationTokenType() *string {
	return s.ApplicationTokenType
}

func (s *ObtainApplicationTokenResponseBodyApplicationToken) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ObtainApplicationTokenResponseBodyApplicationToken) GetExpirationTime() *int64 {
	return s.ExpirationTime
}

func (s *ObtainApplicationTokenResponseBodyApplicationToken) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ObtainApplicationTokenResponseBodyApplicationToken) GetLastUsedTime() *int64 {
	return s.LastUsedTime
}

func (s *ObtainApplicationTokenResponseBodyApplicationToken) GetStatus() *string {
	return s.Status
}

func (s *ObtainApplicationTokenResponseBodyApplicationToken) SetApplicationId(v string) *ObtainApplicationTokenResponseBodyApplicationToken {
	s.ApplicationId = &v
	return s
}

func (s *ObtainApplicationTokenResponseBodyApplicationToken) SetApplicationToken(v string) *ObtainApplicationTokenResponseBodyApplicationToken {
	s.ApplicationToken = &v
	return s
}

func (s *ObtainApplicationTokenResponseBodyApplicationToken) SetApplicationTokenId(v string) *ObtainApplicationTokenResponseBodyApplicationToken {
	s.ApplicationTokenId = &v
	return s
}

func (s *ObtainApplicationTokenResponseBodyApplicationToken) SetApplicationTokenType(v string) *ObtainApplicationTokenResponseBodyApplicationToken {
	s.ApplicationTokenType = &v
	return s
}

func (s *ObtainApplicationTokenResponseBodyApplicationToken) SetCreateTime(v int64) *ObtainApplicationTokenResponseBodyApplicationToken {
	s.CreateTime = &v
	return s
}

func (s *ObtainApplicationTokenResponseBodyApplicationToken) SetExpirationTime(v int64) *ObtainApplicationTokenResponseBodyApplicationToken {
	s.ExpirationTime = &v
	return s
}

func (s *ObtainApplicationTokenResponseBodyApplicationToken) SetInstanceId(v string) *ObtainApplicationTokenResponseBodyApplicationToken {
	s.InstanceId = &v
	return s
}

func (s *ObtainApplicationTokenResponseBodyApplicationToken) SetLastUsedTime(v int64) *ObtainApplicationTokenResponseBodyApplicationToken {
	s.LastUsedTime = &v
	return s
}

func (s *ObtainApplicationTokenResponseBodyApplicationToken) SetStatus(v string) *ObtainApplicationTokenResponseBodyApplicationToken {
	s.Status = &v
	return s
}

func (s *ObtainApplicationTokenResponseBodyApplicationToken) Validate() error {
	return dara.Validate(s)
}
