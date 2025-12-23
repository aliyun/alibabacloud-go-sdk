// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrCode(v int32) *CreateTokenResponseBody
	GetErrCode() *int32
	SetErrMsg(v string) *CreateTokenResponseBody
	GetErrMsg() *string
	SetNlsRequestId(v string) *CreateTokenResponseBody
	GetNlsRequestId() *string
	SetRequestId(v string) *CreateTokenResponseBody
	GetRequestId() *string
	SetToken(v *CreateTokenResponseBodyToken) *CreateTokenResponseBody
	GetToken() *CreateTokenResponseBodyToken
}

type CreateTokenResponseBody struct {
	// example:
	//
	// 50000000
	ErrCode *int32 `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// Specified access key is not found.
	ErrMsg *string `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	// example:
	//
	// dd05a301b40441c99a2671905325****
	NlsRequestId *string `json:"NlsRequestId,omitempty" xml:"NlsRequestId,omitempty"`
	// example:
	//
	// A51587CB-5193-4DB8-9AED-CD4365C2****
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Token     *CreateTokenResponseBodyToken `json:"Token,omitempty" xml:"Token,omitempty" type:"Struct"`
}

func (s CreateTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTokenResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTokenResponseBody) GetErrCode() *int32 {
	return s.ErrCode
}

func (s *CreateTokenResponseBody) GetErrMsg() *string {
	return s.ErrMsg
}

func (s *CreateTokenResponseBody) GetNlsRequestId() *string {
	return s.NlsRequestId
}

func (s *CreateTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTokenResponseBody) GetToken() *CreateTokenResponseBodyToken {
	return s.Token
}

func (s *CreateTokenResponseBody) SetErrCode(v int32) *CreateTokenResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CreateTokenResponseBody) SetErrMsg(v string) *CreateTokenResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *CreateTokenResponseBody) SetNlsRequestId(v string) *CreateTokenResponseBody {
	s.NlsRequestId = &v
	return s
}

func (s *CreateTokenResponseBody) SetRequestId(v string) *CreateTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTokenResponseBody) SetToken(v *CreateTokenResponseBodyToken) *CreateTokenResponseBody {
	s.Token = v
	return s
}

func (s *CreateTokenResponseBody) Validate() error {
	if s.Token != nil {
		if err := s.Token.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateTokenResponseBodyToken struct {
	// example:
	//
	// 1553825814
	ExpireTime *int32 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// fb1c4648a61b426589dab0fe90d1****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 107992689699****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CreateTokenResponseBodyToken) String() string {
	return dara.Prettify(s)
}

func (s CreateTokenResponseBodyToken) GoString() string {
	return s.String()
}

func (s *CreateTokenResponseBodyToken) GetExpireTime() *int32 {
	return s.ExpireTime
}

func (s *CreateTokenResponseBodyToken) GetId() *string {
	return s.Id
}

func (s *CreateTokenResponseBodyToken) GetUserId() *string {
	return s.UserId
}

func (s *CreateTokenResponseBodyToken) SetExpireTime(v int32) *CreateTokenResponseBodyToken {
	s.ExpireTime = &v
	return s
}

func (s *CreateTokenResponseBodyToken) SetId(v string) *CreateTokenResponseBodyToken {
	s.Id = &v
	return s
}

func (s *CreateTokenResponseBodyToken) SetUserId(v string) *CreateTokenResponseBodyToken {
	s.UserId = &v
	return s
}

func (s *CreateTokenResponseBodyToken) Validate() error {
	return dara.Validate(s)
}
