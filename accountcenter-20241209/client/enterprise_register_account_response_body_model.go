// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseRegisterAccountResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetAccountInfo(v *EnterpriseRegisterAccountResponseBodyAccountInfo) *EnterpriseRegisterAccountResponseBody
  GetAccountInfo() *EnterpriseRegisterAccountResponseBodyAccountInfo 
  SetCode(v string) *EnterpriseRegisterAccountResponseBody
  GetCode() *string 
  SetMessage(v string) *EnterpriseRegisterAccountResponseBody
  GetMessage() *string 
  SetRequestId(v string) *EnterpriseRegisterAccountResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EnterpriseRegisterAccountResponseBody
  GetSuccess() *bool 
}

type EnterpriseRegisterAccountResponseBody struct {
  AccountInfo *EnterpriseRegisterAccountResponseBodyAccountInfo `json:"AccountInfo,omitempty" xml:"AccountInfo,omitempty" type:"Struct"`
  // example:
  // 
  // 200
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  // example:
  // 
  // The operation is not allowed. Channel state (RELEASED) does not meet expectations (ANSWERED).
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // example:
  // 
  // BDFCF081-7BCD-52D5-9D82-0F58D96EFD92
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // example:
  // 
  // True
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnterpriseRegisterAccountResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseRegisterAccountResponseBody) GoString() string {
  return s.String()
}

func (s *EnterpriseRegisterAccountResponseBody) GetAccountInfo() *EnterpriseRegisterAccountResponseBodyAccountInfo  {
  return s.AccountInfo
}

func (s *EnterpriseRegisterAccountResponseBody) GetCode() *string  {
  return s.Code
}

func (s *EnterpriseRegisterAccountResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EnterpriseRegisterAccountResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnterpriseRegisterAccountResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EnterpriseRegisterAccountResponseBody) SetAccountInfo(v *EnterpriseRegisterAccountResponseBodyAccountInfo) *EnterpriseRegisterAccountResponseBody {
  s.AccountInfo = v
  return s
}

func (s *EnterpriseRegisterAccountResponseBody) SetCode(v string) *EnterpriseRegisterAccountResponseBody {
  s.Code = &v
  return s
}

func (s *EnterpriseRegisterAccountResponseBody) SetMessage(v string) *EnterpriseRegisterAccountResponseBody {
  s.Message = &v
  return s
}

func (s *EnterpriseRegisterAccountResponseBody) SetRequestId(v string) *EnterpriseRegisterAccountResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnterpriseRegisterAccountResponseBody) SetSuccess(v bool) *EnterpriseRegisterAccountResponseBody {
  s.Success = &v
  return s
}

func (s *EnterpriseRegisterAccountResponseBody) Validate() error {
  if s.AccountInfo != nil {
    if err := s.AccountInfo.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type EnterpriseRegisterAccountResponseBodyAccountInfo struct {
  EnterpriseLicenseNo *string `json:"EnterpriseLicenseNo,omitempty" xml:"EnterpriseLicenseNo,omitempty"`
  // example:
  // 
  // 海南屿可网络科技有限公司
  EnterpriseName *string `json:"EnterpriseName,omitempty" xml:"EnterpriseName,omitempty"`
  // example:
  // 
  // 195529
  LoginId *string `json:"LoginId,omitempty" xml:"LoginId,omitempty"`
  // example:
  // 
  // 5190216604405754
  Pk *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
}

func (s EnterpriseRegisterAccountResponseBodyAccountInfo) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseRegisterAccountResponseBodyAccountInfo) GoString() string {
  return s.String()
}

func (s *EnterpriseRegisterAccountResponseBodyAccountInfo) GetEnterpriseLicenseNo() *string  {
  return s.EnterpriseLicenseNo
}

func (s *EnterpriseRegisterAccountResponseBodyAccountInfo) GetEnterpriseName() *string  {
  return s.EnterpriseName
}

func (s *EnterpriseRegisterAccountResponseBodyAccountInfo) GetLoginId() *string  {
  return s.LoginId
}

func (s *EnterpriseRegisterAccountResponseBodyAccountInfo) GetPk() *string  {
  return s.Pk
}

func (s *EnterpriseRegisterAccountResponseBodyAccountInfo) SetEnterpriseLicenseNo(v string) *EnterpriseRegisterAccountResponseBodyAccountInfo {
  s.EnterpriseLicenseNo = &v
  return s
}

func (s *EnterpriseRegisterAccountResponseBodyAccountInfo) SetEnterpriseName(v string) *EnterpriseRegisterAccountResponseBodyAccountInfo {
  s.EnterpriseName = &v
  return s
}

func (s *EnterpriseRegisterAccountResponseBodyAccountInfo) SetLoginId(v string) *EnterpriseRegisterAccountResponseBodyAccountInfo {
  s.LoginId = &v
  return s
}

func (s *EnterpriseRegisterAccountResponseBodyAccountInfo) SetPk(v string) *EnterpriseRegisterAccountResponseBodyAccountInfo {
  s.Pk = &v
  return s
}

func (s *EnterpriseRegisterAccountResponseBodyAccountInfo) Validate() error {
  return dara.Validate(s)
}

