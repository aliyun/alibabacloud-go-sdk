// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEcologyOpennessAuthenticateResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v int32) *EcologyOpennessAuthenticateResponseBody
  GetCode() *int32 
  SetMessage(v string) *EcologyOpennessAuthenticateResponseBody
  GetMessage() *string 
  SetRequestId(v string) *EcologyOpennessAuthenticateResponseBody
  GetRequestId() *string 
  SetResult(v *EcologyOpennessAuthenticateResponseBodyResult) *EcologyOpennessAuthenticateResponseBody
  GetResult() *EcologyOpennessAuthenticateResponseBodyResult 
  SetSuccess(v bool) *EcologyOpennessAuthenticateResponseBody
  GetSuccess() *bool 
}

type EcologyOpennessAuthenticateResponseBody struct {
  // example:
  // 
  // 200
  Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
  // example:
  // 
  // OK
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // example:
  // 
  // 73****9-175A-1324-8202-9FAAB*****A
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  Result *EcologyOpennessAuthenticateResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
  // example:
  // 
  // true
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EcologyOpennessAuthenticateResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EcologyOpennessAuthenticateResponseBody) GoString() string {
  return s.String()
}

func (s *EcologyOpennessAuthenticateResponseBody) GetCode() *int32  {
  return s.Code
}

func (s *EcologyOpennessAuthenticateResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EcologyOpennessAuthenticateResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EcologyOpennessAuthenticateResponseBody) GetResult() *EcologyOpennessAuthenticateResponseBodyResult  {
  return s.Result
}

func (s *EcologyOpennessAuthenticateResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EcologyOpennessAuthenticateResponseBody) SetCode(v int32) *EcologyOpennessAuthenticateResponseBody {
  s.Code = &v
  return s
}

func (s *EcologyOpennessAuthenticateResponseBody) SetMessage(v string) *EcologyOpennessAuthenticateResponseBody {
  s.Message = &v
  return s
}

func (s *EcologyOpennessAuthenticateResponseBody) SetRequestId(v string) *EcologyOpennessAuthenticateResponseBody {
  s.RequestId = &v
  return s
}

func (s *EcologyOpennessAuthenticateResponseBody) SetResult(v *EcologyOpennessAuthenticateResponseBodyResult) *EcologyOpennessAuthenticateResponseBody {
  s.Result = v
  return s
}

func (s *EcologyOpennessAuthenticateResponseBody) SetSuccess(v bool) *EcologyOpennessAuthenticateResponseBody {
  s.Success = &v
  return s
}

func (s *EcologyOpennessAuthenticateResponseBody) Validate() error {
  return dara.Validate(s)
}

type EcologyOpennessAuthenticateResponseBodyResult struct {
  // example:
  // 
  // 12****7
  EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
  // example:
  // 
  // PROJECT_ID
  EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
  // example:
  // 
  // ******
  SceneCode *string `json:"SceneCode,omitempty" xml:"SceneCode,omitempty"`
  // example:
  // 
  // ******
  ThirdUserIdentifier *string `json:"ThirdUserIdentifier,omitempty" xml:"ThirdUserIdentifier,omitempty"`
  // example:
  // 
  // ******
  ThirdUserType *string `json:"ThirdUserType,omitempty" xml:"ThirdUserType,omitempty"`
  // example:
  // 
  // o****RnNAW/smBNX9By7Zlc3J7iQUXPiUj/6OizU+ifLSzn1vpQL9ZgSp22u7hsxj0UZ2i6urbv9HQ==
  UserOpenId *string `json:"UserOpenId,omitempty" xml:"UserOpenId,omitempty"`
}

func (s EcologyOpennessAuthenticateResponseBodyResult) String() string {
  return dara.Prettify(s)
}

func (s EcologyOpennessAuthenticateResponseBodyResult) GoString() string {
  return s.String()
}

func (s *EcologyOpennessAuthenticateResponseBodyResult) GetEncodeKey() *string  {
  return s.EncodeKey
}

func (s *EcologyOpennessAuthenticateResponseBodyResult) GetEncodeType() *string  {
  return s.EncodeType
}

func (s *EcologyOpennessAuthenticateResponseBodyResult) GetSceneCode() *string  {
  return s.SceneCode
}

func (s *EcologyOpennessAuthenticateResponseBodyResult) GetThirdUserIdentifier() *string  {
  return s.ThirdUserIdentifier
}

func (s *EcologyOpennessAuthenticateResponseBodyResult) GetThirdUserType() *string  {
  return s.ThirdUserType
}

func (s *EcologyOpennessAuthenticateResponseBodyResult) GetUserOpenId() *string  {
  return s.UserOpenId
}

func (s *EcologyOpennessAuthenticateResponseBodyResult) SetEncodeKey(v string) *EcologyOpennessAuthenticateResponseBodyResult {
  s.EncodeKey = &v
  return s
}

func (s *EcologyOpennessAuthenticateResponseBodyResult) SetEncodeType(v string) *EcologyOpennessAuthenticateResponseBodyResult {
  s.EncodeType = &v
  return s
}

func (s *EcologyOpennessAuthenticateResponseBodyResult) SetSceneCode(v string) *EcologyOpennessAuthenticateResponseBodyResult {
  s.SceneCode = &v
  return s
}

func (s *EcologyOpennessAuthenticateResponseBodyResult) SetThirdUserIdentifier(v string) *EcologyOpennessAuthenticateResponseBodyResult {
  s.ThirdUserIdentifier = &v
  return s
}

func (s *EcologyOpennessAuthenticateResponseBodyResult) SetThirdUserType(v string) *EcologyOpennessAuthenticateResponseBodyResult {
  s.ThirdUserType = &v
  return s
}

func (s *EcologyOpennessAuthenticateResponseBodyResult) SetUserOpenId(v string) *EcologyOpennessAuthenticateResponseBodyResult {
  s.UserOpenId = &v
  return s
}

func (s *EcologyOpennessAuthenticateResponseBodyResult) Validate() error {
  return dara.Validate(s)
}

